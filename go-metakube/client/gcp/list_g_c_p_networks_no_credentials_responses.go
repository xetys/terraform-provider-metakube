// Code generated by go-swagger; DO NOT EDIT.

package gcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/terraform-provider-metakube/go-metakube/models"
)

// ListGCPNetworksNoCredentialsReader is a Reader for the ListGCPNetworksNoCredentials structure.
type ListGCPNetworksNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPNetworksNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGCPNetworksNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGCPNetworksNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPNetworksNoCredentialsOK creates a ListGCPNetworksNoCredentialsOK with default headers values
func NewListGCPNetworksNoCredentialsOK() *ListGCPNetworksNoCredentialsOK {
	return &ListGCPNetworksNoCredentialsOK{}
}

/*ListGCPNetworksNoCredentialsOK handles this case with default header values.

GCPNetworkList
*/
type ListGCPNetworksNoCredentialsOK struct {
	Payload models.GCPNetworkList
}

func (o *ListGCPNetworksNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/networks][%d] listGCPNetworksNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListGCPNetworksNoCredentialsOK) GetPayload() models.GCPNetworkList {
	return o.Payload
}

func (o *ListGCPNetworksNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPNetworksNoCredentialsDefault creates a ListGCPNetworksNoCredentialsDefault with default headers values
func NewListGCPNetworksNoCredentialsDefault(code int) *ListGCPNetworksNoCredentialsDefault {
	return &ListGCPNetworksNoCredentialsDefault{
		_statusCode: code,
	}
}

/*ListGCPNetworksNoCredentialsDefault handles this case with default header values.

errorResponse
*/
type ListGCPNetworksNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list g c p networks no credentials default response
func (o *ListGCPNetworksNoCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListGCPNetworksNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/networks][%d] listGCPNetworksNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPNetworksNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListGCPNetworksNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
