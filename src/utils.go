package src

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func ValidateParameters(r *http.Request) (map[string]int, error) {
	params := mux.Vars(r)
	normalized := map[string]int{
		"hours":   0,
		"minutes": 0,
	}
	hours, err := strconv.Atoi(params["hours"])
	if err != nil {
		log.Println("Error")
		return normalized, err
	}
	if hours <= 0 || hours > 12 {
		log.Println("Error")
		return normalized, errors.New("the hours value must be between 1 and 12")
	}
	_, ok := params["minutes"]
	if ok {
		minutes, err := strconv.Atoi(params["minutes"])
		if err != nil {
			log.Println("Error:", err)
			return normalized, err
		}
		if minutes < 0 || minutes >= 60 {
			log.Println("Error")
			return normalized, errors.New("the minutes value must be between 0 and 60")
		}
		normalized["minutes"] = minutes
	}

	normalized["hours"] = hours

	return normalized, nil

}

func PrepareLogger() *log.Logger {
	filename := "logs/logfile.txt"
	_, err := os.Stat(filename)
	if err != nil {
		_, err := os.Create(filename)
		if err != nil {
			log.Fatal("Error creating log file:", err)
		}
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	defer file.Close()

	return log.New(file, "custom-prefix: ", log.LstdFlags)
}

func ParseResults(parameters map[string]int, result int) map[string]interface{} {

	results := map[string]interface{}{
		"angle":   result,
		"hours":   parameters["hours"],
		"minutes": parameters["minutes"],
	}
	return results
}

/*
Next steps:
Testes
*/
