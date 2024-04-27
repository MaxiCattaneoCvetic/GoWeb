package user

import (
	"context"
	"goWeb/v2/internal/domain"
	"log"
)

type (
	Service interface {
		/*
			Este método toma un contexto (ctx) de tipo context.Context y tres cadenas de texto (firstName, Lastname y email).
			El método está destinado a crear un nuevo usuario con los datos proporcionados y
			devolver un puntero a un objeto User del paquete domain, así como un posible error
			si ocurre algún problema durante la creación del usuario.
		*/

		Create(ctx context.Context, firstName, Lastname, email string) (*domain.User, error) // recibe un contexto y datos de un usuario -> enviamos a doman
		GetAll(ctx context.Context) ([]domain.User, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{ // &service -> me retorna la referencia del puntero de service en memoria
		log:  log,
		repo: repo,
	}
}

func (s service) Create(ctx context.Context, firstName, Lastname, email string) (*domain.User, error) {
	user := &domain.User{
		FirstName: firstName,
		LastName:  Lastname,
		Email:     email,
	}
	err := s.repo.Create(ctx, user)

	if err != nil {
		return nil, err
	}
	s.log.Println("Service log: create user")
	return user, nil

}

func (s service) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}
	s.log.Println("Service log: Get all users")
	return users, nil

}
