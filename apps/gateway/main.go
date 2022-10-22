package main

import (
	"log"
	"os"
	"strings"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/config"
)

// @title Chula Overflow Backend Doc
// @version 1.0
// @description Not over engineering at all
// @termsOfService http://swagger.io/terms/
// @host localhost:3000
// @BasePath /
func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	port := strings.Split(conf.GatewayURL, ":")[1]

	app := NewServer(conf)

	if app.config.Deployment == "development" {
		app.RegisterDoc()
	}

	app.Listen(":" + port)
}
