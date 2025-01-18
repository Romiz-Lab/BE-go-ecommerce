package app

import (
	"net/http"

	"github.com/Romiz-Lab/BE-go-ecommerce/app/controllers"
	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/api", controllers.Home).Methods("GET")

	staticDir := http.Dir("./assets/")
	staticDirHandler := http.StripPrefix("/public/", http.FileServer(staticDir))
	server.Router.PathPrefix("/public/").Handler(staticDirHandler).Methods("GET")
}