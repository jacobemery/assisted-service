// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// CompleteInstallationReader is a Reader for the CompleteInstallation structure.
type CompleteInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCompleteInstallationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCompleteInstallationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCompleteInstallationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCompleteInstallationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCompleteInstallationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCompleteInstallationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCompleteInstallationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCompleteInstallationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCompleteInstallationAccepted creates a CompleteInstallationAccepted with default headers values
func NewCompleteInstallationAccepted() *CompleteInstallationAccepted {
	return &CompleteInstallationAccepted{}
}

/*CompleteInstallationAccepted handles this case with default header values.

Success.
*/
type CompleteInstallationAccepted struct {
	Payload *models.Cluster
}

func (o *CompleteInstallationAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/actions/complete_installation][%d] completeInstallationAccepted  %+v", 202, o.Payload)
}

func (o *CompleteInstallationAccepted) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *CompleteInstallationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteInstallationUnauthorized creates a CompleteInstallationUnauthorized with default headers values
func NewCompleteInstallationUnauthorized() *CompleteInstallationUnauthorized {
	return &CompleteInstallationUnauthorized{}
}

/*CompleteInstallationUnauthorized handles this case with default header values.

Unauthorized.
*/
type CompleteInstallationUnauthorized struct {
	Payload *models.InfraError
}

func (o *CompleteInstallationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/actions/complete_installation][%d] completeInstallationUnauthorized  %+v", 401, o.Payload)
}

func (o *CompleteInstallationUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *CompleteInstallationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteInstallationForbidden creates a CompleteInstallationForbidden with default headers values
func NewCompleteInstallationForbidden() *CompleteInstallationForbidden {
	return &CompleteInstallationForbidden{}
}

/*CompleteInstallationForbidden handles this case with default header values.

Forbidden.
*/
type CompleteInstallationForbidden struct {
	Payload *models.InfraError
}

func (o *CompleteInstallationForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/actions/complete_installation][%d] completeInstallationForbidden  %+v", 403, o.Payload)
}

func (o *CompleteInstallationForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *CompleteInstallationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteInstallationNotFound creates a CompleteInstallationNotFound with default headers values
func NewCompleteInstallationNotFound() *CompleteInstallationNotFound {
	return &CompleteInstallationNotFound{}
}

/*CompleteInstallationNotFound handles this case with default header values.

Error.
*/
type CompleteInstallationNotFound struct {
	Payload *models.Error
}

func (o *CompleteInstallationNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/actions/complete_installation][%d] completeInstallationNotFound  %+v", 404, o.Payload)
}

func (o *CompleteInstallationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CompleteInstallationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteInstallationMethodNotAllowed creates a CompleteInstallationMethodNotAllowed with default headers values
func NewCompleteInstallationMethodNotAllowed() *CompleteInstallationMethodNotAllowed {
	return &CompleteInstallationMethodNotAllowed{}
}

/*CompleteInstallationMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type CompleteInstallationMethodNotAllowed struct {
	Payload *models.Error
}

func (o *CompleteInstallationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/actions/complete_installation][%d] completeInstallationMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CompleteInstallationMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *CompleteInstallationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteInstallationConflict creates a CompleteInstallationConflict with default headers values
func NewCompleteInstallationConflict() *CompleteInstallationConflict {
	return &CompleteInstallationConflict{}
}

/*CompleteInstallationConflict handles this case with default header values.

Error.
*/
type CompleteInstallationConflict struct {
	Payload *models.Error
}

func (o *CompleteInstallationConflict) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/actions/complete_installation][%d] completeInstallationConflict  %+v", 409, o.Payload)
}

func (o *CompleteInstallationConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CompleteInstallationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteInstallationInternalServerError creates a CompleteInstallationInternalServerError with default headers values
func NewCompleteInstallationInternalServerError() *CompleteInstallationInternalServerError {
	return &CompleteInstallationInternalServerError{}
}

/*CompleteInstallationInternalServerError handles this case with default header values.

Error.
*/
type CompleteInstallationInternalServerError struct {
	Payload *models.Error
}

func (o *CompleteInstallationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/actions/complete_installation][%d] completeInstallationInternalServerError  %+v", 500, o.Payload)
}

func (o *CompleteInstallationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CompleteInstallationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteInstallationServiceUnavailable creates a CompleteInstallationServiceUnavailable with default headers values
func NewCompleteInstallationServiceUnavailable() *CompleteInstallationServiceUnavailable {
	return &CompleteInstallationServiceUnavailable{}
}

/*CompleteInstallationServiceUnavailable handles this case with default header values.

Unavailable.
*/
type CompleteInstallationServiceUnavailable struct {
	Payload *models.Error
}

func (o *CompleteInstallationServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/actions/complete_installation][%d] completeInstallationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CompleteInstallationServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *CompleteInstallationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}