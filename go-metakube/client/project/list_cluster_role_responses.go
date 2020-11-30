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

// ListClusterRoleReader is a Reader for the ListClusterRole structure.
type ListClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClusterRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListClusterRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListClusterRoleOK creates a ListClusterRoleOK with default headers values
func NewListClusterRoleOK() *ListClusterRoleOK {
	return &ListClusterRoleOK{}
}

/*ListClusterRoleOK handles this case with default header values.

ClusterRole
*/
type ListClusterRoleOK struct {
	Payload []*models.ClusterRole
}

func (o *ListClusterRoleOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleOK  %+v", 200, o.Payload)
}

func (o *ListClusterRoleOK) GetPayload() []*models.ClusterRole {
	return o.Payload
}

func (o *ListClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterRoleUnauthorized creates a ListClusterRoleUnauthorized with default headers values
func NewListClusterRoleUnauthorized() *ListClusterRoleUnauthorized {
	return &ListClusterRoleUnauthorized{}
}

/*ListClusterRoleUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListClusterRoleUnauthorized struct {
}

func (o *ListClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleUnauthorized ", 401)
}

func (o *ListClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterRoleForbidden creates a ListClusterRoleForbidden with default headers values
func NewListClusterRoleForbidden() *ListClusterRoleForbidden {
	return &ListClusterRoleForbidden{}
}

/*ListClusterRoleForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListClusterRoleForbidden struct {
}

func (o *ListClusterRoleForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles][%d] listClusterRoleForbidden ", 403)
}

func (o *ListClusterRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterRoleDefault creates a ListClusterRoleDefault with default headers values
func NewListClusterRoleDefault(code int) *ListClusterRoleDefault {
	return &ListClusterRoleDefault{
		_statusCode: code,
	}
}

/*ListClusterRoleDefault handles this case with default header values.

errorResponse
*/
type ListClusterRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list cluster role default response
func (o *ListClusterRoleDefault) Code() int {
	return o._statusCode
}

func (o *ListClusterRoleDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles][%d] listClusterRole default  %+v", o._statusCode, o.Payload)
}

func (o *ListClusterRoleDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListClusterRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
