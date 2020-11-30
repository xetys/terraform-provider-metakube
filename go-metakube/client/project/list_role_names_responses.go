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

// ListRoleNamesReader is a Reader for the ListRoleNames structure.
type ListRoleNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRoleNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRoleNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRoleNamesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRoleNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListRoleNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRoleNamesOK creates a ListRoleNamesOK with default headers values
func NewListRoleNamesOK() *ListRoleNamesOK {
	return &ListRoleNamesOK{}
}

/*ListRoleNamesOK handles this case with default header values.

RoleName
*/
type ListRoleNamesOK struct {
	Payload []*models.RoleName
}

func (o *ListRoleNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames][%d] listRoleNamesOK  %+v", 200, o.Payload)
}

func (o *ListRoleNamesOK) GetPayload() []*models.RoleName {
	return o.Payload
}

func (o *ListRoleNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRoleNamesUnauthorized creates a ListRoleNamesUnauthorized with default headers values
func NewListRoleNamesUnauthorized() *ListRoleNamesUnauthorized {
	return &ListRoleNamesUnauthorized{}
}

/*ListRoleNamesUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListRoleNamesUnauthorized struct {
}

func (o *ListRoleNamesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames][%d] listRoleNamesUnauthorized ", 401)
}

func (o *ListRoleNamesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRoleNamesForbidden creates a ListRoleNamesForbidden with default headers values
func NewListRoleNamesForbidden() *ListRoleNamesForbidden {
	return &ListRoleNamesForbidden{}
}

/*ListRoleNamesForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListRoleNamesForbidden struct {
}

func (o *ListRoleNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames][%d] listRoleNamesForbidden ", 403)
}

func (o *ListRoleNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRoleNamesDefault creates a ListRoleNamesDefault with default headers values
func NewListRoleNamesDefault(code int) *ListRoleNamesDefault {
	return &ListRoleNamesDefault{
		_statusCode: code,
	}
}

/*ListRoleNamesDefault handles this case with default header values.

errorResponse
*/
type ListRoleNamesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list role names default response
func (o *ListRoleNamesDefault) Code() int {
	return o._statusCode
}

func (o *ListRoleNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/rolenames][%d] listRoleNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListRoleNamesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListRoleNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
