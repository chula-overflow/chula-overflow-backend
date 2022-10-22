package main

import (
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/middleware"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/config"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	*fiber.App
	config *config.Config
}

type Handler = func(*context.Ctx) error

type Method int

const (
	GET Method = iota
	POST
)

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
		AllowOrigins: "*",
	}))

	app.Use(recover.New())

	app.RegisterRoute()

	return app
}

func (app *App) AddHdr(handler Handler, metadata MetaData) {
	if metadata.requiredAuth {
		app.Use(metadata.url, middleware.AuthMiddleWare)
	}

	switch method := metadata.method; method {

	case GET:
		app.Get(metadata.url, func(ctx *fiber.Ctx) error {
			return handler(context.NewCtx(ctx))
		})

	case POST:
		app.Post(metadata.url, func(ctx *fiber.Ctx) error {
			return handler(context.NewCtx(ctx))
		})
	}

}
