// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListOpenstackImagesParams creates a new ListOpenstackImagesParams object
// with the default values initialized.
func NewListOpenstackImagesParams() *ListOpenstackImagesParams {

	return &ListOpenstackImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOpenstackImagesParamsWithTimeout creates a new ListOpenstackImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOpenstackImagesParamsWithTimeout(timeout time.Duration) *ListOpenstackImagesParams {

	return &ListOpenstackImagesParams{

		timeout: timeout,
	}
}

// NewListOpenstackImagesParamsWithContext creates a new ListOpenstackImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOpenstackImagesParamsWithContext(ctx context.Context) *ListOpenstackImagesParams {

	return &ListOpenstackImagesParams{

		Context: ctx,
	}
}

// NewListOpenstackImagesParamsWithHTTPClient creates a new ListOpenstackImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOpenstackImagesParamsWithHTTPClient(client *http.Client) *ListOpenstackImagesParams {

	return &ListOpenstackImagesParams{
		HTTPClient: client,
	}
}

/*ListOpenstackImagesParams contains all the parameters to send to the API endpoint
for the list openstack images operation typically these are written to a http.Request
*/
type ListOpenstackImagesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list openstack images params
func (o *ListOpenstackImagesParams) WithTimeout(timeout time.Duration) *ListOpenstackImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list openstack images params
func (o *ListOpenstackImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list openstack images params
func (o *ListOpenstackImagesParams) WithContext(ctx context.Context) *ListOpenstackImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list openstack images params
func (o *ListOpenstackImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list openstack images params
func (o *ListOpenstackImagesParams) WithHTTPClient(client *http.Client) *ListOpenstackImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list openstack images params
func (o *ListOpenstackImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListOpenstackImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
