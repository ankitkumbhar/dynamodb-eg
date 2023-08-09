package main

import (
	"dynamodb-eg/internal/database"
	"dynamodb-eg/internal/handler"
	"dynamodb-eg/internal/routes"
	"net/http"
)

var app *handler.Application

// init
func init() {
	// initialize database
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	app = handler.New(db)
}

func main() {
	r := routes.InitRouter(app)

	http.ListenAndServe(":8080", r)
}
