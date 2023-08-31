package main

import (
	"fmt"
	"log"
	"myproject/db"
	"myproject/src"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func calculateClockAngle(w http.ResponseWriter, r *http.Request) {

	logger := src.PrepareLogger()

	normalized, err := src.ValidateParameters(r)

	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		logger.Fatal("Error: ", err)
	} else {
		logger.Println("Parameters correctly validated.")
	}

	result := src.CalculateAngleBetweenArrows(normalized)
	results := src.ParseResults(normalized, result)
	requestAddress := r.Host
	errPersistance := db.Persist(results, requestAddress)
	if errPersistance != nil {
		logger.Println("There was found an error during data persistance: ", err)
	}
	fmt.Fprintf(w, `{"angle": %d}`, result)
	logger.Println("Angle between arrows: ", result)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	r := mux.NewRouter()

	r.HandleFunc("/calculate-clock/{hours:[0-9]+}/{minutes:[0-9]+}", calculateClockAngle).Methods("GET")
	r.HandleFunc("/calculate-clock/{hours:[0-9]+}", calculateClockAngle).Methods("GET")

	apiPort := os.Getenv("API_PORT")
	apiHost := os.Getenv("API_HOST")

	log.Println("Server is running on", apiHost+":"+apiPort)
	log.Fatal(http.ListenAndServe(apiHost+":"+apiPort, r))
}
