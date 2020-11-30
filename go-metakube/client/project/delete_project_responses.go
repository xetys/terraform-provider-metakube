// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/terraform-provider-metakube/go-metakube/models"
)

// DeleteProjectReader is a Reader for the DeleteProject structure.
type DeleteProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteProjectOK creates a DeleteProjectOK with default headers values
func NewDeleteProjectOK() *DeleteProjectOK {
	return &DeleteProjectOK{}
}

/*DeleteProjectOK handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteProjectOK struct {
}

func (o *DeleteProjectOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}][%d] deleteProjectOK ", 200)
}

func (o *DeleteProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectUnauthorized creates a DeleteProjectUnauthorized with default headers values
func NewDeleteProjectUnauthorized() *DeleteProjectUnauthorized {
	return &DeleteProjectUnauthorized{}
}

/*DeleteProjectUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteProjectUnauthorized struct {
}

func (o *DeleteProjectUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}][%d] deleteProjectUnauthorized ", 401)
}

func (o *DeleteProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectForbidden creates a DeleteProjectForbidden with default headers values
func NewDeleteProjectForbidden() *DeleteProjectForbidden {
	return &DeleteProjectForbidden{}
}

/*DeleteProjectForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteProjectForbidden struct {
}

func (o *DeleteProjectForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}][%d] deleteProjectForbidden ", 403)
}

func (o *DeleteProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectDefault creates a DeleteProjectDefault with default headers values
func NewDeleteProjectDefault(code int) *DeleteProjectDefault {
	return &DeleteProjectDefault{
		_statusCode: code,
	}
}

/*DeleteProjectDefault handles this case with default header values.

errorResponse
*/
type DeleteProjectDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete project default response
func (o *DeleteProjectDefault) Code() int {
	return o._statusCode
}

func (o *DeleteProjectDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}][%d] deleteProject default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteProjectDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
