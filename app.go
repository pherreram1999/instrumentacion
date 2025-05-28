package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"go.bug.st/serial"
	"log"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetSerialPorts() ([]string, error) {
	return serial.GetPortsList()
}

func (a *App) OpenPort(portName string) error {

	mode := serial.Mode{
		BaudRate: 9600,
	}

	port, err := serial.Open(portName, &mode)
	reader := bufio.NewReader(port)
	if err != nil {
		return err
	}

	go func() {
		for {
			lineBytes, err := reader.ReadBytes('\n')
			if err != nil {
				log.Fatal(err)
			}

			runtime.EventsEmit(a.ctx, "voltage", string(lineBytes))
		}
	}()

	return nil
}
