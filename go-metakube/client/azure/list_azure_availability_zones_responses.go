// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/syseleven/terraform-provider-metakube/go-metakube/models"
)

// ListAzureAvailabilityZonesReader is a Reader for the ListAzureAvailabilityZones structure.
type ListAzureAvailabilityZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAzureAvailabilityZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAzureAvailabilityZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAzureAvailabilityZonesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAzureAvailabilityZonesOK creates a ListAzureAvailabilityZonesOK with default headers values
func NewListAzureAvailabilityZonesOK() *ListAzureAvailabilityZonesOK {
	return &ListAzureAvailabilityZonesOK{}
}

/*ListAzureAvailabilityZonesOK handles this case with default header values.

AzureAvailabilityZonesList
*/
type ListAzureAvailabilityZonesOK struct {
	Payload *models.AzureAvailabilityZonesList
}

func (o *ListAzureAvailabilityZonesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/azure/availabilityzones][%d] listAzureAvailabilityZonesOK  %+v", 200, o.Payload)
}

func (o *ListAzureAvailabilityZonesOK) GetPayload() *models.AzureAvailabilityZonesList {
	return o.Payload
}

func (o *ListAzureAvailabilityZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureAvailabilityZonesList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAzureAvailabilityZonesDefault creates a ListAzureAvailabilityZonesDefault with default headers values
func NewListAzureAvailabilityZonesDefault(code int) *ListAzureAvailabilityZonesDefault {
	return &ListAzureAvailabilityZonesDefault{
		_statusCode: code,
	}
}

/*ListAzureAvailabilityZonesDefault handles this case with default header values.

errorResponse
*/
type ListAzureAvailabilityZonesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list azure availability zones default response
func (o *ListAzureAvailabilityZonesDefault) Code() int {
	return o._statusCode
}

func (o *ListAzureAvailabilityZonesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/azure/availabilityzones][%d] listAzureAvailabilityZones default  %+v", o._statusCode, o.Payload)
}

func (o *ListAzureAvailabilityZonesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAzureAvailabilityZonesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
