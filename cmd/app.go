package cmd

import (
	Comment "api/internal/comment/services"
	"api/internal/database"
	transportHTTP "api/internal/transport/http"
	"api/pkg/logging"
	"net/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run(logger logging.Logger) error {
	logger.Info("http://localhost:8080")

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
		logger.Error("failed to set up server")
		return err
	}
	return nil
}
