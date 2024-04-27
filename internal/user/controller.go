package user

import (
	"context"
	"encoding/json" // con eso transformamos una estructura a json
	"fmt"
	"goWeb/v2/internal/domain"
	"net/http"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

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

func MakeEndpoints(ctx context.Context, s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAllUsers(ctx, s, w)
		case http.MethodPost:
			decode := json.NewDecoder(r.Body) // decodificamos el body
			var user domain.User
			if err := decode.Decode(&user); err != nil {
				MsgResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			PostUser(ctx, s, w, user)
		default:
			InvalidMethod(w)
		}
	}

}

/* METODO GET*/
func GetAllUsers(ctx context.Context, s Service, w http.ResponseWriter) {
	users, err := s.GetAll(ctx)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	DataResponse(w, http.StatusOK, users)
}

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var users []User
var MaxID uint64

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, err := json.Marshal(users) // me transforma 1 entidad a un json
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, value)
}

/* METODO POST*/

func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, message)

}

func PostUser(ctx context.Context, s Service, w http.ResponseWriter, data interface{}) {
	req := data.(CreateRequest) // casteo el valor de data a un User
	if req.FirstName == "" || req.LastName == "" || req.Email == "" {
		MsgResponse(w, http.StatusBadRequest, "Faltan datos")
		return
	}
	user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	DataResponse(w, http.StatusCreated, user)

}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "Method doesnt exist"}`, status)
}
