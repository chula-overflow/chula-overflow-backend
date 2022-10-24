package main

import (
	docs "github.com/chula-overflow/chula-overflow-backend/apps/gateway/docs"
	"github.com/gofiber/swagger"
)

func (app *App) RegisterDoc() {
	docs.SwaggerInfo.Host = app.config.GatewayURL

	app.Get("/swagger/*", swagger.HandlerDefault) // default
}
