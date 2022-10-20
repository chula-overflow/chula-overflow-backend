package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/src/config"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/src/proto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func getServer(conf *config.Config) *fiber.App {
	app := fiber.New()

	flag.Parse()
	conn, err := grpc.Dial(conf.Auth.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(2)
	}

	authClient := proto.NewAuthClient(conn)

	app.Get("/", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		result, err := authClient.Login(ctx, &proto.AuthRequest{
			Name: "Hello",
		})

		if err != nil {
			log.Fatal(err)
			return err
		}

		return c.SendString(result.Token)
	})

	return app
}
