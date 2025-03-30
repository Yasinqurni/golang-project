package main

import (
	"fmt"
	di "golang-project/internal"
	"golang-project/pkg/config"
	"golang-project/pkg/db"
	"golang-project/pkg/httpserver"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//setup config
	config, err := config.NewConfig("./")
	if err != nil {
		log.Fatal("config error ", err)
	}

	//setup db
	db.NewGorm(config)

	//dependency injection
	mux := di.NewInternal(config)

	//start server
	httpServer := httpserver.New(mux, httpserver.Port(config.App.Port))
	log.Println("server is running on port:", config.App.Port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: ", s.String())
	case err = <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err).Error())
	}
	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}
}
