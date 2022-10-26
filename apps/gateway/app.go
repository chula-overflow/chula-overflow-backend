package main

import (
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	*fiber.App
	config *config.Config
}

type MetaData struct {
	url          string
	requiredAuth bool
	method       Method
}

func NewServer(conf *config.Config) *App {
	fiberApp := fiber.New()

	app := &App{
		fiberApp,
		conf,
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))

	app.Use(recover.New())

	app.RegisterRoute()

	return app
}
