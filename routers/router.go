package routers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	//Routes for User Entity
	router = SetUserRoutes(router)
	//Routes for Task Entity
	router = SetTasksRoutes(router)
	//Routes for Note Entity
	router = SetNotesRoutes(router)
	return router
}