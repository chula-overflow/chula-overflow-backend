package main

import (
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/src/app/middleware"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/src/config"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/src/context"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	*fiber.App
	config *config.Config
}

type Handler = func(*context.Ctx)

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

func NewServer(fiberApp *fiber.App, conf *config.Config) *App {
	return &App{
		fiberApp,
		conf,
	}
}

func (app *App) RegisterHdr(handler Handler, metadata MetaData) {
	if metadata.requiredAuth {
		app.Use(metadata.url, middleware.AuthMiddleWare)
	}

	switch method := metadata.method; method {

	case GET:
		app.Get(metadata.url, func(ctx *fiber.Ctx) error {
			handler(context.NewCtx(ctx))
			return nil
		})

	case POST:
		app.Post(metadata.url, func(ctx *fiber.Ctx) error {
			handler(context.NewCtx(ctx))
			return nil
		})
	}

}
