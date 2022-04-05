package main

import (
	"api/cmd"
	"api/pkg/logging"
)

func main() {
	logging.Init()
	var app cmd.App

	logger := logging.GetLogger()
	if err := app.Run(logger); err != nil {
		logger.Error(err)
	}
}
