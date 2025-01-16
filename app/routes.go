package app

import "github.com/Romiz-Lab/BE-go-ecommerce/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/api", controllers.Home).Methods("GET")
}