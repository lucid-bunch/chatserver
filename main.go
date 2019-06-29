package main

import (
	"chatserver/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := app.NewApp()
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, os.Interrupt)
	go func() {
		sig := <-gracefulStop
		app.Close(sig)
	}()
	app.Listen(3000)
}
