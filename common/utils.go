package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)


type configuration struct {
	Server, MongoDBHost, DBUser, DBPwd, Database string
}

type (
	appError struct {
		Error	string		`json:"error"`
		Message	string		`json:"message"`
		HttpStatus	int		`json:"status"`
	}
	errorResource struct {
		Data appError	`json:"data"`
	}
)

//AppConfig holds the configuration value from config.json file
var AppConfig configuration

//Initialize AppConfig
func AppConfig() {
	loadAppConfig()
}

//Reads config.json and decode into AppConfig
func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
}

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:handlerError.Error(),
		Message: message,
		HttpStatus: code,
	}

	log.Printf("[AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}