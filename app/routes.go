package app

import (
	"github.com/Romiz-Lab/BE-go-ecommerce/app/controllers"
	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/api", controllers.Home).Methods("GET")
}