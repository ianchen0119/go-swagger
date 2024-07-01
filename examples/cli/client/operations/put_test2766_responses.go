// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ianchen0119/go-swagger/examples/cli/models"
)

// PutTest2766Reader is a Reader for the PutTest2766 structure.
type PutTest2766Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTest2766Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutTest2766OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /test2766] PutTest2766", response, response.Code())
	}
}

// NewPutTest2766OK creates a PutTest2766OK with default headers values
func NewPutTest2766OK() *PutTest2766OK {
	return &PutTest2766OK{}
}

/*
PutTest2766OK describes a response with status code 200, with default header values.

test issue #2766
*/
type PutTest2766OK struct {
	Payload *models.GithubReactions
}

// IsSuccess returns true when this put test2766 o k response has a 2xx status code
func (o *PutTest2766OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put test2766 o k response has a 3xx status code
func (o *PutTest2766OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put test2766 o k response has a 4xx status code
func (o *PutTest2766OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put test2766 o k response has a 5xx status code
func (o *PutTest2766OK) IsServerError() bool {
	return false
}

// IsCode returns true when this put test2766 o k response a status code equal to that given
func (o *PutTest2766OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put test2766 o k response
func (o *PutTest2766OK) Code() int {
	return 200
}

func (o *PutTest2766OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /test2766][%d] putTest2766OK %s", 200, payload)
}

func (o *PutTest2766OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /test2766][%d] putTest2766OK %s", 200, payload)
}

func (o *PutTest2766OK) GetPayload() *models.GithubReactions {
	return o.Payload
}

func (o *PutTest2766OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubReactions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
