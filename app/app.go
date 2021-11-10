package app

import (
	"errors"
	"fmt"

	"github.com/arshabbir/gmux/src/controller"
)

type app struct {
	ctrl controller.Controller
}

type App interface {
	StartApp() error
}

func NewApp(ctrl controller.Controller) App {
	return &app{ctrl: ctrl}
}

func (a *app) StartApp() error {

	appStatus := make(chan int)
	fmt.Println("Starting the app....")
	go a.ctrl.Start(appStatus)
	for {
		select {
		case v, ok := <-appStatus:
			if !ok {
				return errors.New("Channel Closed")
			}
			if v == 1 {
				return nil
			}
			if v == 2 {
				return errors.New("Error Starting the application")
			}
		default:

		}

	}

}
