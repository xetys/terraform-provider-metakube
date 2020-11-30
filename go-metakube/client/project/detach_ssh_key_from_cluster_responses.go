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

// DetachSSHKeyFromClusterReader is a Reader for the DetachSSHKeyFromCluster structure.
type DetachSSHKeyFromClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetachSSHKeyFromClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetachSSHKeyFromClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetachSSHKeyFromClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDetachSSHKeyFromClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDetachSSHKeyFromClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDetachSSHKeyFromClusterOK creates a DetachSSHKeyFromClusterOK with default headers values
func NewDetachSSHKeyFromClusterOK() *DetachSSHKeyFromClusterOK {
	return &DetachSSHKeyFromClusterOK{}
}

/*DetachSSHKeyFromClusterOK handles this case with default header values.

EmptyResponse is a empty response
*/
type DetachSSHKeyFromClusterOK struct {
}

func (o *DetachSSHKeyFromClusterOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] detachSshKeyFromClusterOK ", 200)
}

func (o *DetachSSHKeyFromClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDetachSSHKeyFromClusterUnauthorized creates a DetachSSHKeyFromClusterUnauthorized with default headers values
func NewDetachSSHKeyFromClusterUnauthorized() *DetachSSHKeyFromClusterUnauthorized {
	return &DetachSSHKeyFromClusterUnauthorized{}
}

/*DetachSSHKeyFromClusterUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type DetachSSHKeyFromClusterUnauthorized struct {
}

func (o *DetachSSHKeyFromClusterUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] detachSshKeyFromClusterUnauthorized ", 401)
}

func (o *DetachSSHKeyFromClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDetachSSHKeyFromClusterForbidden creates a DetachSSHKeyFromClusterForbidden with default headers values
func NewDetachSSHKeyFromClusterForbidden() *DetachSSHKeyFromClusterForbidden {
	return &DetachSSHKeyFromClusterForbidden{}
}

/*DetachSSHKeyFromClusterForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type DetachSSHKeyFromClusterForbidden struct {
}

func (o *DetachSSHKeyFromClusterForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] detachSshKeyFromClusterForbidden ", 403)
}

func (o *DetachSSHKeyFromClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDetachSSHKeyFromClusterDefault creates a DetachSSHKeyFromClusterDefault with default headers values
func NewDetachSSHKeyFromClusterDefault(code int) *DetachSSHKeyFromClusterDefault {
	return &DetachSSHKeyFromClusterDefault{
		_statusCode: code,
	}
}

/*DetachSSHKeyFromClusterDefault handles this case with default header values.

errorResponse
*/
type DetachSSHKeyFromClusterDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the detach SSH key from cluster default response
func (o *DetachSSHKeyFromClusterDefault) Code() int {
	return o._statusCode
}

func (o *DetachSSHKeyFromClusterDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] detachSSHKeyFromCluster default  %+v", o._statusCode, o.Payload)
}

func (o *DetachSSHKeyFromClusterDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DetachSSHKeyFromClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
