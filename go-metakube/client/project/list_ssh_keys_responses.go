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

// ListSSHKeysReader is a Reader for the ListSSHKeys structure.
type ListSSHKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSSHKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSSHKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListSSHKeysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListSSHKeysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListSSHKeysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSSHKeysOK creates a ListSSHKeysOK with default headers values
func NewListSSHKeysOK() *ListSSHKeysOK {
	return &ListSSHKeysOK{}
}

/*ListSSHKeysOK handles this case with default header values.

SSHKey
*/
type ListSSHKeysOK struct {
	Payload []*models.SSHKey
}

func (o *ListSSHKeysOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/sshkeys][%d] listSshKeysOK  %+v", 200, o.Payload)
}

func (o *ListSSHKeysOK) GetPayload() []*models.SSHKey {
	return o.Payload
}

func (o *ListSSHKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSSHKeysUnauthorized creates a ListSSHKeysUnauthorized with default headers values
func NewListSSHKeysUnauthorized() *ListSSHKeysUnauthorized {
	return &ListSSHKeysUnauthorized{}
}

/*ListSSHKeysUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListSSHKeysUnauthorized struct {
}

func (o *ListSSHKeysUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/sshkeys][%d] listSshKeysUnauthorized ", 401)
}

func (o *ListSSHKeysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSSHKeysForbidden creates a ListSSHKeysForbidden with default headers values
func NewListSSHKeysForbidden() *ListSSHKeysForbidden {
	return &ListSSHKeysForbidden{}
}

/*ListSSHKeysForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListSSHKeysForbidden struct {
}

func (o *ListSSHKeysForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/sshkeys][%d] listSshKeysForbidden ", 403)
}

func (o *ListSSHKeysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSSHKeysDefault creates a ListSSHKeysDefault with default headers values
func NewListSSHKeysDefault(code int) *ListSSHKeysDefault {
	return &ListSSHKeysDefault{
		_statusCode: code,
	}
}

/*ListSSHKeysDefault handles this case with default header values.

errorResponse
*/
type ListSSHKeysDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list SSH keys default response
func (o *ListSSHKeysDefault) Code() int {
	return o._statusCode
}

func (o *ListSSHKeysDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/sshkeys][%d] listSSHKeys default  %+v", o._statusCode, o.Payload)
}

func (o *ListSSHKeysDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListSSHKeysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
