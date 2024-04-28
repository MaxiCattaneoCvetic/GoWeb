/*
En este caso tenemos un solo dominio, pero en el caso de tener mas dominios necesitariamos un handler para cada uno de ellos
*/

package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"goWeb/v2/internal/user"
	"goWeb/v2/pkg/transport"
	"net/http"
)

func NewUserHttpServer(ctx context.Context, router *http.ServeMux, endpoints user.Endpoints) {
	router.HandleFunc("/users", UserServer(ctx, endpoints)) // generamos el endpoint

}

func UserServer(ctx context.Context, endpoints user.Endpoints) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//generamos la capa transporte
		tran := transport.NewTransport(w, r, ctx)
		switch r.Method {
		case http.MethodGet:
			tran.Server(
				transport.Endpoint(endpoints.GetAll),
				decodeGetAllUsers,
				encodeResponse,
				encodeError)
			return
		case http.MethodPost:
			tran.Server(
				transport.Endpoint(endpoints.Create),
				decodeCreateUser,
				encodeResponse,
				encodeError)
			return
		}
		InvalidMethod(w)
	}
}

func decodeGetAllUsers(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	data, err := json.Marshal(response) // me transforma 1 entidad a un json
	if err != nil {
		return err
	}
	status := http.StatusOK
	w.WriteHeader(status)
	w.Header().Set("content-type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, data)
	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	status := http.StatusInternalServerError
	w.WriteHeader(status)
	w.Header().Set("content-type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{"status": %d, "message": %s}`, status, err.Error())
}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "Method doesnt exist"}`, status)
}

func decodeCreateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req user.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("invalid request format: %v", err.Error())
	}
	return req, nil

}
