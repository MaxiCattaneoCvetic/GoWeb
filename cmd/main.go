package main

import (
	"context"
	"fmt"
	"goWeb/v2/internal/domain"
	"goWeb/v2/internal/user"
	"log"
	"net/http"
	"os"
)

func main() {

	db := user.DB{
		Users: []domain.User{
			{
				ID:        1,
				FirstName: "Cristian",
				LastName:  "Vazquez",
				Email:     "XVJfX@example.com",
			},
			{
				ID:        2,
				FirstName: "juan",
				LastName:  "Perezx",
				Email:     "asd@example.com",
			},
			{
				ID:        3,
				FirstName: "Pedro",
				LastName:  "pedrini",
				Email:     "peddre@example.com",
			},
		},
		MaxUserID: 3, // porque el maximo de usuario que tengo es 3, cuando se cree uno se incrementara +1
	}

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)
	/*
		Generamos un contexto para manejar la informacion y enviamos infor a las demas capas
	*/
	ctx := context.Background()

	//generamos un objeto server}
	server := http.NewServeMux()
	server.HandleFunc("/users", user.MakeEndpoints(ctx, service)) // generamos el endpoint
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))

}
