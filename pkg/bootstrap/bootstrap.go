package boostrap

/*
Este package se encarga de crear la base de datos y el logger


*/

import (
	"goWeb/v2/internal/domain"
	"goWeb/v2/internal/user"
	"log"
	"os"
)

// esta funcion se encarga de retornar el logger
func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func NewBD() user.DB {
	return user.DB{
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
}
