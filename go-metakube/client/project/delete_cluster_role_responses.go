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

// DeleteClusterRoleReader is a Reader for the DeleteClusterRole structure.
type DeleteClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteClusterRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteClusterRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteClusterRoleOK creates a DeleteClusterRoleOK with default headers values
func NewDeleteClusterRoleOK() *DeleteClusterRoleOK {
	return &DeleteClusterRoleOK{}
}

/*DeleteClusterRoleOK handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteClusterRoleOK struct {
}

func (o *DeleteClusterRoleOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] deleteClusterRoleOK ", 200)
}

func (o *DeleteClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterRoleUnauthorized creates a DeleteClusterRoleUnauthorized with default headers values
func NewDeleteClusterRoleUnauthorized() *DeleteClusterRoleUnauthorized {
	return &DeleteClusterRoleUnauthorized{}
}

/*DeleteClusterRoleUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteClusterRoleUnauthorized struct {
}

func (o *DeleteClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] deleteClusterRoleUnauthorized ", 401)
}

func (o *DeleteClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterRoleForbidden creates a DeleteClusterRoleForbidden with default headers values
func NewDeleteClusterRoleForbidden() *DeleteClusterRoleForbidden {
	return &DeleteClusterRoleForbidden{}
}

/*DeleteClusterRoleForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteClusterRoleForbidden struct {
}

func (o *DeleteClusterRoleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] deleteClusterRoleForbidden ", 403)
}

func (o *DeleteClusterRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterRoleDefault creates a DeleteClusterRoleDefault with default headers values
func NewDeleteClusterRoleDefault(code int) *DeleteClusterRoleDefault {
	return &DeleteClusterRoleDefault{
		_statusCode: code,
	}
}

/*DeleteClusterRoleDefault handles this case with default header values.

errorResponse
*/
type DeleteClusterRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete cluster role default response
func (o *DeleteClusterRoleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteClusterRoleDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] deleteClusterRole default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteClusterRoleDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteClusterRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
