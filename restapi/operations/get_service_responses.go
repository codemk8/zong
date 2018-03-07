// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*GetServiceDefault successful operation

swagger:response getServiceDefault
*/
type GetServiceDefault struct {
	_statusCode int
}

// NewGetServiceDefault creates GetServiceDefault with default headers values
func NewGetServiceDefault(code int) *GetServiceDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServiceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get service default response
func (o *GetServiceDefault) WithStatusCode(code int) *GetServiceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get service default response
func (o *GetServiceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *GetServiceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
