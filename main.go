package main

import (
	"log"
	"net/http"

	"github.com/bockbone/taskmanager/common"
	"github.com/bockbone/taskmanager/routers"
	"github.com/codegangsta/negroni"
)

//Entry of program
func main() {

	//Call startup logic
	common.StartUp()

	//Get the mux router object
	router := routers.InitRoutes()

	//Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr: common.AppConfig.Server,
		Handler: n,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}