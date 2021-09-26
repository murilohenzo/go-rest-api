package main

import (
	Comment "api/internal/comment/services"
	"api/internal/database"
	transportHTTP "api/internal/transport/http"
	"fmt"
	"net/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct {}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Settings Up Our App")

	var err error

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	commentService := Comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("GO REST API")
	var app App
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
