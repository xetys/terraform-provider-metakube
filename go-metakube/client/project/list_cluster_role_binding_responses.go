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

// ListClusterRoleBindingReader is a Reader for the ListClusterRoleBinding structure.
type ListClusterRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClusterRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClusterRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListClusterRoleBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClusterRoleBindingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListClusterRoleBindingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListClusterRoleBindingOK creates a ListClusterRoleBindingOK with default headers values
func NewListClusterRoleBindingOK() *ListClusterRoleBindingOK {
	return &ListClusterRoleBindingOK{}
}

/*ListClusterRoleBindingOK handles this case with default header values.

ClusterRoleBinding
*/
type ListClusterRoleBindingOK struct {
	Payload []*models.ClusterRoleBinding
}

func (o *ListClusterRoleBindingOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterbindings][%d] listClusterRoleBindingOK  %+v", 200, o.Payload)
}

func (o *ListClusterRoleBindingOK) GetPayload() []*models.ClusterRoleBinding {
	return o.Payload
}

func (o *ListClusterRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClusterRoleBindingUnauthorized creates a ListClusterRoleBindingUnauthorized with default headers values
func NewListClusterRoleBindingUnauthorized() *ListClusterRoleBindingUnauthorized {
	return &ListClusterRoleBindingUnauthorized{}
}

/*ListClusterRoleBindingUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListClusterRoleBindingUnauthorized struct {
}

func (o *ListClusterRoleBindingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterbindings][%d] listClusterRoleBindingUnauthorized ", 401)
}

func (o *ListClusterRoleBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterRoleBindingForbidden creates a ListClusterRoleBindingForbidden with default headers values
func NewListClusterRoleBindingForbidden() *ListClusterRoleBindingForbidden {
	return &ListClusterRoleBindingForbidden{}
}

/*ListClusterRoleBindingForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListClusterRoleBindingForbidden struct {
}

func (o *ListClusterRoleBindingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterbindings][%d] listClusterRoleBindingForbidden ", 403)
}

func (o *ListClusterRoleBindingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClusterRoleBindingDefault creates a ListClusterRoleBindingDefault with default headers values
func NewListClusterRoleBindingDefault(code int) *ListClusterRoleBindingDefault {
	return &ListClusterRoleBindingDefault{
		_statusCode: code,
	}
}

/*ListClusterRoleBindingDefault handles this case with default header values.

errorResponse
*/
type ListClusterRoleBindingDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list cluster role binding default response
func (o *ListClusterRoleBindingDefault) Code() int {
	return o._statusCode
}

func (o *ListClusterRoleBindingDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterbindings][%d] listClusterRoleBinding default  %+v", o._statusCode, o.Payload)
}

func (o *ListClusterRoleBindingDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListClusterRoleBindingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
