package main

import (
	"library/datastore"
	"library/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	s := datastore.New()
	h := handler.New(s)

	app.GET("/library/{id}", h.GetByID)
	app.POST("/library", h.Create)
	app.PUT("/library/{id}", h.Update)
	app.DELETE("/library/{id}", h.Delete)

	// starting the server on a custom port
	app.Server.HTTP.Port = 9000
	app.Start()
}
