package main

import (
	"log"
	"os"

	authHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/handler/auth"
	authSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/auth"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/config"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (app *App) RegisterRoute() {
	auth := GetHandler(app.config)

	metadata := MetaData{
		url:          "/auth/login",
		requiredAuth: false,
		method:       POST,
	}
	app.AddHdr(auth.Login, metadata)

	metadata = MetaData{
		url:          "/auth/revoke",
		requiredAuth: true,
		method:       GET,
	}
	app.AddHdr(auth.Revoke, metadata)

	metadata = MetaData{
		url:          "/auth/me",
		requiredAuth: true,
		method:       GET,
	}
	app.AddHdr(auth.Me, metadata)
}

func GetHandler(conf *config.Config) authHdr.Handler {
	conn, err := grpc.Dial(conf.Auth.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(2)
	}
	authClient := proto.NewAuthClient(conn)
	authService := authSrv.NewService(authClient)
	auth := authHdr.NewHandler(&authService)

	return auth
}
