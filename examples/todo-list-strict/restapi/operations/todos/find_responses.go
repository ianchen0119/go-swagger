// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ianchen0119/go-swagger/examples/todo-list-strict/models"
)

// FindOKCode is the HTTP code returned for type FindOK
const FindOKCode int = 200

/*
FindOK OK

swagger:response findOK
*/
type FindOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Item `json:"body,omitempty"`
}

// NewFindOK creates FindOK with default headers values
func NewFindOK() *FindOK {

	return &FindOK{}
}

// WithPayload adds the payload to the find o k response
func (o *FindOK) WithPayload(payload []*models.Item) *FindOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find o k response
func (o *FindOK) SetPayload(payload []*models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Item, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (o *FindOK) FindResponder() {}

/*
FindDefault error

swagger:response findDefault
*/
type FindDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindDefault creates FindDefault with default headers values
func NewFindDefault(code int) *FindDefault {
	if code <= 0 {
		code = 500
	}

	return &FindDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find default response
func (o *FindDefault) WithStatusCode(code int) *FindDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find default response
func (o *FindDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find default response
func (o *FindDefault) WithPayload(payload *models.Error) *FindDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find default response
func (o *FindDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *FindDefault) FindResponder() {}

type FindNotImplementedResponder struct {
	middleware.Responder
}

func (*FindNotImplementedResponder) FindResponder() {}

func FindNotImplemented() FindResponder {
	return &FindNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.Find has not yet been implemented",
		),
	}
}

type FindResponder interface {
	middleware.Responder
	FindResponder()
}
