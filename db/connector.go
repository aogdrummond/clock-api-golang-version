package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var db *sql.DB

func initDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)
	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
}

func Persist(results map[string]interface{}, requestAddress string) error {
	initDB()
	defer db.Close()

	cursor := db
	dB := &dbCursor{cursor: cursor}

	err := dB.persistResult(results, requestAddress)
	return err
}

type dbCursor struct {
	cursor *sql.DB
}

func (dbCursor *dbCursor) persistResult(results map[string]interface{}, requestAddress string) error {

	logger := log.New(os.Stdout, "appLogger ", log.LstdFlags)

	angle := results["angle"].(int)
	hours := results["hours"].(int)
	minutes := results["minutes"].(int)
	query := dbCursor.createQueryToInsert(angle, hours, minutes, requestAddress)

	tx, err := dbCursor.cursor.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return err
	}

	_, err = tx.Exec(dbCursor.queryToCreateTable("requests_archive"))
	if err != nil {
		log.Printf("Error creating table: %v", err)
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec(query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		return err
	}

	logger.Println("Operation persisted successfully.")
	return nil
}

func (dbCursor *dbCursor) createQueryToInsert(angle, hours, minutes int, requestAddress string) string {
	currentDatetime := time.Now().Format("2006-01-02 15:04:05")
	query := fmt.Sprintf("INSERT INTO requests_archive (angle, clock, request_address, requested_at) "+
		"VALUES (%d, ARRAY[%d,%d],'%s','%s');",
		angle, hours, minutes, requestAddress, currentDatetime)

	return query
}

func (dbCursor *dbCursor) queryToCreateTable(tableName string) string {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id SERIAL PRIMARY KEY, angle INTEGER, clock INTEGER[], requested_at TIMESTAMP, request_address VARCHAR(155));", tableName)

	return query
}
