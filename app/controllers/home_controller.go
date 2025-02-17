package controllers

import (
	"net/http"

	"github.com/unrolled/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
	})

	render.HTML(w, http.StatusOK, "home", map[string]interface{}{
		"title": "Home",
		"message": "Welcome to the Home Page!",
	})
}