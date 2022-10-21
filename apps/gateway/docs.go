package main

import (
	_ "github.com/chula-overflow/chula-overflow-backend/apps/gateway/docs"
	"github.com/gofiber/swagger"
)

func (app *App) RegisterDoc() {
	app.Get("/swagger/*", swagger.HandlerDefault) // default
}
