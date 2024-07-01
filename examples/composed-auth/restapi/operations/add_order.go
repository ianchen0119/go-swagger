// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/ianchen0119/go-swagger/examples/composed-auth/models"
)

// AddOrderHandlerFunc turns a function with the right signature into a add order handler
type AddOrderHandlerFunc func(AddOrderParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn AddOrderHandlerFunc) Handle(params AddOrderParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// AddOrderHandler interface for that can handle valid add order params
type AddOrderHandler interface {
	Handle(AddOrderParams, *models.Principal) middleware.Responder
}

// NewAddOrder creates a new http.Handler for the add order operation
func NewAddOrder(ctx *middleware.Context, handler AddOrderHandler) *AddOrder {
	return &AddOrder{Context: ctx, Handler: handler}
}

/*
	AddOrder swagger:route POST /order/add addOrder

post a new order

Registered customers should be able to add purchase orders.
Registered inventory managers should be able to add replenishment orders.
*/
type AddOrder struct {
	Context *middleware.Context
	Handler AddOrderHandler
}

func (o *AddOrder) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddOrderParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
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
