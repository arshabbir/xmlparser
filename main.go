package main

import (
	"fmt"
	"os"

	"github.com/arshabbir/gmux/src/app"
	"github.com/arshabbir/gmux/src/controller"
	"github.com/arshabbir/gmux/src/services"
)

func main() {

	if err := app.NewApp(controller.NewController(services.NewService())).StartApp(); err != nil {
		fmt.Println("Error starting the application")
		os.Exit(1)
	}
}
