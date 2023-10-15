package main

import (
	"golang-grpc/app"
	"golang-grpc/config"
	"log"
	"os"
)

func main() {

	var err error
	// initialize config (database, cache, log, etc)
	if err = config.Application.InitConfig(); err != nil {
		os.Exit(1)
	}
	app.Server.New()
	app.Server.Register()
	if err = app.Server.Run(); err != nil {
		log.Fatal(err)
	}
}
