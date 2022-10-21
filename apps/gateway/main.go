package main

import (
	"log"
	"os"
	"strings"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/config"
	"github.com/gofiber/fiber/v2"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	port := strings.Split(conf.Gateway.Addr, ":")[1]

	app := getServer(conf)

	if app.config.Mode == "development" {
		app.RegisterDoc()
	}

	app.Listen(":" + port)
}

func getServer(conf *config.Config) *App {
	fiberApp := fiber.New()

	app := NewServer(fiberApp, conf)
	app.RegisterRoute()

	return app
}
