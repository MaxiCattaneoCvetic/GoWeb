package main

import (
	"context"
	"fmt"
	"goWeb/v2/internal/user"
	boostrap "goWeb/v2/pkg/bootstrap"
	"goWeb/v2/pkg/handler"
	"log"
	"net/http"
)

func main() {

	//generamos un objeto server
	server := http.NewServeMux()
	db := boostrap.NewBD()
	logger := boostrap.NewLogger()

	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)
	/*
		Generamos un contexto para manejar la informacion y enviamos infor a las demas capas
	*/
	ctx := context.Background()
	handler.NewUserHttpServer(ctx, server, user.MakeEndpoints(ctx, service))

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))

}
