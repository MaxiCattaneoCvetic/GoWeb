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
	"log"
	"net/http"
	"strconv"
)

func NewUserHttpServer(ctx context.Context, router *http.ServeMux, endpoints user.Endpoints) {
	router.HandleFunc("/users/", UserServer(ctx, endpoints)) // generamos el endpoint

}

func UserServer(ctx context.Context, endpoints user.Endpoints) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		//PathVariable
		url := r.URL.Path
		log.Println(r.Method, ": ", url)
		path, pathSise := transport.Clean(url)

		params := make(map[string]string)

		if pathSise == 4 && path[2] != "" {
			params["userID"] = path[2]
		}

		//generamos la capa transporte

		tran := transport.NewTransport(w, r, context.WithValue(ctx, "params", params))

		var end user.Controller
		var deco func(ctx context.Context, r *http.Request) (interface{}, error)

		switch r.Method {

		case http.MethodGet:
			switch pathSise {
			case 3:
				end = endpoints.GetAll
				deco = decodeGetAllUsers
				// encodeResponse,
				// encodeError,
			case 4:
				end = endpoints.Get
				deco = decodeGetUser
			}
		case http.MethodPost:
			switch pathSise {
			case 3:
				end = endpoints.Create
				deco = decodeCreateUser
			}
		case http.MethodPatch:
			switch pathSise {
			case 4:
				end = endpoints.Update
				deco = decoUpdateUser
			}
		}
		if end != nil && deco != nil {
			tran.Server(
				transport.Endpoint(end),
				deco,
				encodeResponse,
				encodeError,
			)
		} else {
			InvalidMethod(w)
		}
	}
}

func decodeGetUser(ctx context.Context, r *http.Request) (interface{}, error) {
	params := ctx.Value("params").(map[string]string)
	//Convertimos el string a int  -> 10 = base 10 y 64 = 64 bits
	id, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		return nil, err
	}

	return user.GetReq{
		ID: id,
	}, nil
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

//UPDATE USER

func decoUpdateUser(ctx context.Context, request *http.Request) (interface{}, error) {
	var req user.UpdateReq
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("invalid request format: %v", err.Error())
	}

	params := ctx.Value("params").(map[string]string)

	//Convertimos el string a int  -> 10 = base 10 y 64 = 64 bits
	id, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		return nil, err
	}
	req.ID = id
	return req, nil

}
