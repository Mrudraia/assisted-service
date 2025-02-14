// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2ListEventsOKCode is the HTTP code returned for type V2ListEventsOK
const V2ListEventsOKCode int = 200

/*V2ListEventsOK Success.

swagger:response v2ListEventsOK
*/
type V2ListEventsOK struct {

	/*
	  In: Body
	*/
	Payload models.EventList `json:"body,omitempty"`
}

// NewV2ListEventsOK creates V2ListEventsOK with default headers values
func NewV2ListEventsOK() *V2ListEventsOK {

	return &V2ListEventsOK{}
}

// WithPayload adds the payload to the v2 list events o k response
func (o *V2ListEventsOK) WithPayload(payload models.EventList) *V2ListEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list events o k response
func (o *V2ListEventsOK) SetPayload(payload models.EventList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.EventList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// V2ListEventsUnauthorizedCode is the HTTP code returned for type V2ListEventsUnauthorized
const V2ListEventsUnauthorizedCode int = 401

/*V2ListEventsUnauthorized Unauthorized.

swagger:response v2ListEventsUnauthorized
*/
type V2ListEventsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2ListEventsUnauthorized creates V2ListEventsUnauthorized with default headers values
func NewV2ListEventsUnauthorized() *V2ListEventsUnauthorized {

	return &V2ListEventsUnauthorized{}
}

// WithPayload adds the payload to the v2 list events unauthorized response
func (o *V2ListEventsUnauthorized) WithPayload(payload *models.InfraError) *V2ListEventsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list events unauthorized response
func (o *V2ListEventsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListEventsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ListEventsForbiddenCode is the HTTP code returned for type V2ListEventsForbidden
const V2ListEventsForbiddenCode int = 403

/*V2ListEventsForbidden Forbidden.

swagger:response v2ListEventsForbidden
*/
type V2ListEventsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2ListEventsForbidden creates V2ListEventsForbidden with default headers values
func NewV2ListEventsForbidden() *V2ListEventsForbidden {

	return &V2ListEventsForbidden{}
}

// WithPayload adds the payload to the v2 list events forbidden response
func (o *V2ListEventsForbidden) WithPayload(payload *models.InfraError) *V2ListEventsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list events forbidden response
func (o *V2ListEventsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListEventsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ListEventsNotFoundCode is the HTTP code returned for type V2ListEventsNotFound
const V2ListEventsNotFoundCode int = 404

/*V2ListEventsNotFound Error.

swagger:response v2ListEventsNotFound
*/
type V2ListEventsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ListEventsNotFound creates V2ListEventsNotFound with default headers values
func NewV2ListEventsNotFound() *V2ListEventsNotFound {

	return &V2ListEventsNotFound{}
}

// WithPayload adds the payload to the v2 list events not found response
func (o *V2ListEventsNotFound) WithPayload(payload *models.Error) *V2ListEventsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list events not found response
func (o *V2ListEventsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListEventsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ListEventsMethodNotAllowedCode is the HTTP code returned for type V2ListEventsMethodNotAllowed
const V2ListEventsMethodNotAllowedCode int = 405

/*V2ListEventsMethodNotAllowed Method Not Allowed.

swagger:response v2ListEventsMethodNotAllowed
*/
type V2ListEventsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ListEventsMethodNotAllowed creates V2ListEventsMethodNotAllowed with default headers values
func NewV2ListEventsMethodNotAllowed() *V2ListEventsMethodNotAllowed {

	return &V2ListEventsMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 list events method not allowed response
func (o *V2ListEventsMethodNotAllowed) WithPayload(payload *models.Error) *V2ListEventsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list events method not allowed response
func (o *V2ListEventsMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListEventsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ListEventsInternalServerErrorCode is the HTTP code returned for type V2ListEventsInternalServerError
const V2ListEventsInternalServerErrorCode int = 500

/*V2ListEventsInternalServerError Error.

swagger:response v2ListEventsInternalServerError
*/
type V2ListEventsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ListEventsInternalServerError creates V2ListEventsInternalServerError with default headers values
func NewV2ListEventsInternalServerError() *V2ListEventsInternalServerError {

	return &V2ListEventsInternalServerError{}
}

// WithPayload adds the payload to the v2 list events internal server error response
func (o *V2ListEventsInternalServerError) WithPayload(payload *models.Error) *V2ListEventsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list events internal server error response
func (o *V2ListEventsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListEventsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
