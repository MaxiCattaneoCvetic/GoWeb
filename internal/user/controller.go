package user

import (
	"context"
	"errors"
)

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		Create Controller
		GetAll Controller
	}

	CreateRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var users []User
var MaxID uint64

// MakeEndpotins -> Me devuelve una estructura de endopoints
func MakeEndpoints(ctx context.Context, s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndopoins(s),
		GetAll: makeGetAllEndopoins(s),
	}
}

func makeGetAllEndopoins(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
}

func makeCreateEndopoins(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest) // casteo el valor de data a un User

		if req.FirstName == "" || req.LastName == "" || req.Email == "" {
			return nil, errors.New("Faltan datos")
		}
		user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

}
