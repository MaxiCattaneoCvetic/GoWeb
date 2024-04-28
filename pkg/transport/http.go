/*
Generamos un orden de ejecucion
*/

package transport

import (
	"context"
	"net/http"
)

type Transport interface {
	Server(
		endopoind Endpoint,
		decode func(ctx context.Context, r *http.Request) (interface{}, error), // esto se encarga de tomar la request del cliente y transformarlo para pasarselo al controllador
		encode func(ctx context.Context, w http.ResponseWriter, response interface{}) error, //Este en encargado de responder, cuando el controller devuelve la respuesta el encode le devuelve la response al cliente
		encodeError func(ctx context.Context, err error, w http.ResponseWriter),
	)
}

type Endpoint func(ctx context.Context, request interface{}) (interface{}, error)

type transport struct {
	w   http.ResponseWriter
	r   *http.Request
	ctx context.Context
}

// generamos la funcion que se encarga de generar la estructura transport

func NewTransport(w http.ResponseWriter, r *http.Request, ctx context.Context) Transport {
	return &transport{w: w, r: r, ctx: ctx}
}

//definimos metodos de la estructura

func (t *transport) Server(
	endopoind Endpoint,
	decode func(ctx context.Context, r *http.Request) (interface{}, error), // esto se encarga de tomar la request del cliente y transformarlo para pasarselo al controllador
	encode func(ctx context.Context, w http.ResponseWriter, response interface{}) error, //Este en encargado de responder, cuando el controller devuelve la respuesta el encode le devuelve la response al cliente
	encodeError func(ctx context.Context, err error, w http.ResponseWriter),
) {
	//generamos el midleware
	data, err := decode(t.ctx, t.r)

	if err != nil {
		encodeError(t.ctx, err, t.w)
		return
	}

	res, err := endopoind(t.ctx, data)
	if err != nil {
		encodeError(t.ctx, err, t.w)
		return
	}

	// recibimos la info del controller y lka enviamos al cliente

	if err := encode(t.ctx, t.w, res); err != nil {
		encodeError(t.ctx, err, t.w)
		return
	}

}
