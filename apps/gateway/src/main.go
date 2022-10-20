package main

import (
	"log"
	"os"
	"strings"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/src/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	port := strings.Split(conf.Gateway.Addr, ":")[1]

	app := getServer(conf)

	app.Listen(":" + port)
}

func getServer(conf *config.Config) *App {
	fiberApp := fiber.New()

	app := NewServer(fiberApp, conf)
	app.RegisterRoute()

	return app
}
