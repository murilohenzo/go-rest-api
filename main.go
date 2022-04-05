package main

import (
	Comment "api/internal/comment/services"
	"api/internal/database"
	transportHTTP "api/internal/transport/http"
	"api/pkg/logging"
	"fmt"
	"net/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	logger := logging.GetLogger()
	logger.Info("Settings Up Our App")

	fmt.Println("Listen: http://localhost:8080")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := Comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService, logger)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	logging.Init()
	fmt.Println("GO REST API")
	var app App
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
