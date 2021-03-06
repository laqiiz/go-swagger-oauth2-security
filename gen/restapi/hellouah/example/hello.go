// Code generated by go-swagger; DO NOT EDIT.

package example

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/laqiiz/go-swagger-oauth2-security/gen/models"
)

// HelloHandlerFunc turns a function with the right signature into a hello handler
type HelloHandlerFunc func(HelloParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn HelloHandlerFunc) Handle(params HelloParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// HelloHandler interface for that can handle valid hello params
type HelloHandler interface {
	Handle(HelloParams, *models.Principal) middleware.Responder
}

// NewHello creates a new http.Handler for the hello operation
func NewHello(ctx *middleware.Context, handler HelloHandler) *Hello {
	return &Hello{Context: ctx, Handler: handler}
}

/*Hello swagger:route GET /hello Example hello

hello api

*/
type Hello struct {
	Context *middleware.Context
	Handler HelloHandler
}

func (o *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewHelloParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
