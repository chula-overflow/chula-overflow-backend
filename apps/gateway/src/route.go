package main

import (
	"log"
	"os"

	authHdr "github.com/chula-overflow/chula-overflow-backend/apps/gateway/src/app/handler/auth"
	authSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/src/app/service/auth"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/src/config"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/src/proto"
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
	app.RegisterHdr(auth.Login, metadata)

	metadata = MetaData{
		url:          "/auth/revoke",
		requiredAuth: true,
		method:       GET,
	}
	app.RegisterHdr(auth.Revoke, metadata)
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
