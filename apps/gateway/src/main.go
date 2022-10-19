package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/src/proto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:3001", "the address to connect to")
	name = flag.String("name", "world", "Name to greet")
)

func main() {
	app := fiber.New()

	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	authClient := proto.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.Get("/", func(c *fiber.Ctx) error {
		result, err := authClient.Login(ctx, &proto.AuthRequest{
			Name: "Hello",
		})
		if err != nil {
			log.Fatal(err)
			return err
		}

		return c.SendString(result.Token)
	})

	app.Listen(":3000")
}
