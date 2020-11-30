// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewListClustersForProjectParams creates a new ListClustersForProjectParams object
// with the default values initialized.
func NewListClustersForProjectParams() *ListClustersForProjectParams {
	var ()
	return &ListClustersForProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListClustersForProjectParamsWithTimeout creates a new ListClustersForProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListClustersForProjectParamsWithTimeout(timeout time.Duration) *ListClustersForProjectParams {
	var ()
	return &ListClustersForProjectParams{

		timeout: timeout,
	}
}

// NewListClustersForProjectParamsWithContext creates a new ListClustersForProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewListClustersForProjectParamsWithContext(ctx context.Context) *ListClustersForProjectParams {
	var ()
	return &ListClustersForProjectParams{

		Context: ctx,
	}
}

// NewListClustersForProjectParamsWithHTTPClient creates a new ListClustersForProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListClustersForProjectParamsWithHTTPClient(client *http.Client) *ListClustersForProjectParams {
	var ()
	return &ListClustersForProjectParams{
		HTTPClient: client,
	}
}

/*ListClustersForProjectParams contains all the parameters to send to the API endpoint
for the list clusters for project operation typically these are written to a http.Request
*/
type ListClustersForProjectParams struct {

	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list clusters for project params
func (o *ListClustersForProjectParams) WithTimeout(timeout time.Duration) *ListClustersForProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list clusters for project params
func (o *ListClustersForProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list clusters for project params
func (o *ListClustersForProjectParams) WithContext(ctx context.Context) *ListClustersForProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list clusters for project params
func (o *ListClustersForProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list clusters for project params
func (o *ListClustersForProjectParams) WithHTTPClient(client *http.Client) *ListClustersForProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list clusters for project params
func (o *ListClustersForProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the list clusters for project params
func (o *ListClustersForProjectParams) WithProjectID(projectID string) *ListClustersForProjectParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list clusters for project params
func (o *ListClustersForProjectParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListClustersForProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
