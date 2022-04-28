package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hypebeast/go-osc/osc"
)

var num int

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) Transmitter(address string, value string, valuetype string) string {
	client := osc.NewClient("localhost", 9000)
	msg := osc.NewMessage(address)
	if valuetype == "bool" {
		boolvalue, _ := strconv.ParseBool(value)
		msg.Append(boolvalue)
		client.Send(msg)
	} else if valuetype == "int" {
		intvalue, _ := strconv.Atoi(value)
		msg.Append(int32(intvalue))
		client.Send(msg)
	} else if valuetype == "float" {
		floatvalue, _ := strconv.ParseFloat(value, 32)
		msg.Append(float32(floatvalue))
		client.Send(msg)
	}
	return fmt.Sprintf("> Transmission to \"%s\" complete.<br>> Data transmitted:\"%s\"", address, value)
}
