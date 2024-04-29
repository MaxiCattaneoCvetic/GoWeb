package user

/*
RESPONSE 1

*/import (
	"context"
)

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	//Esto recibimos desde el decode de la request
	GetReq struct {
		ID uint64
	}

	Endpoints struct {
		Create Controller
		GetAll Controller
		Get    Controller
		Update Controller
	}

	CreateRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}

	/*Generamos una estructura de la request del cliente
	Lo generamos como un puntero ya que si el cliente no envia nada devolvera nil,
	pero si fuese un string devolvera el string vacio

	*/
	UpdateReq struct {
		ID        uint64  `json:"id"`
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Email     *string `json:"email"`
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
		Get:    makeGetEndopoint(s),
		Update: makeUpdateEndpoint(s),
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

		if req.FirstName == "" {
			return nil, ErrFirstNameRequired
		}
		if req.LastName == "" {
			return nil, ErrLastNameRequired
		}
		if req.Email == "" {
			return nil, ErrEmailRequired
		}
		user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func makeGetEndopoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetReq)
		user, err := s.Get(ctx, req.ID)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func makeUpdateEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateReq)

		if req.FirstName != nil && *req.FirstName == "" {
			return nil, ErrFirstNameRequired
		}
		if req.LastName != nil && *req.LastName == "" {
			return nil, ErrLastNameRequired
		}
		if req.Email != nil && *req.Email == "" {
			return nil, ErrEmailRequired
		}

		if err := s.Update(ctx, req.ID, req.FirstName, req.LastName, req.Email); err != nil {
			return nil, err
		}

		return nil, nil
	}
}
