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

// NewGetKubeLoginClusterKubeconfigParams creates a new GetKubeLoginClusterKubeconfigParams object
// with the default values initialized.
func NewGetKubeLoginClusterKubeconfigParams() *GetKubeLoginClusterKubeconfigParams {
	var ()
	return &GetKubeLoginClusterKubeconfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubeLoginClusterKubeconfigParamsWithTimeout creates a new GetKubeLoginClusterKubeconfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKubeLoginClusterKubeconfigParamsWithTimeout(timeout time.Duration) *GetKubeLoginClusterKubeconfigParams {
	var ()
	return &GetKubeLoginClusterKubeconfigParams{

		timeout: timeout,
	}
}

// NewGetKubeLoginClusterKubeconfigParamsWithContext creates a new GetKubeLoginClusterKubeconfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKubeLoginClusterKubeconfigParamsWithContext(ctx context.Context) *GetKubeLoginClusterKubeconfigParams {
	var ()
	return &GetKubeLoginClusterKubeconfigParams{

		Context: ctx,
	}
}

// NewGetKubeLoginClusterKubeconfigParamsWithHTTPClient creates a new GetKubeLoginClusterKubeconfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKubeLoginClusterKubeconfigParamsWithHTTPClient(client *http.Client) *GetKubeLoginClusterKubeconfigParams {
	var ()
	return &GetKubeLoginClusterKubeconfigParams{
		HTTPClient: client,
	}
}

/*GetKubeLoginClusterKubeconfigParams contains all the parameters to send to the API endpoint
for the get kube login cluster kubeconfig operation typically these are written to a http.Request
*/
type GetKubeLoginClusterKubeconfigParams struct {

	/*ClusterID*/
	ClusterID string
	/*Dc*/
	DC string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get kube login cluster kubeconfig params
func (o *GetKubeLoginClusterKubeconfigParams) WithTimeout(timeout time.Duration) *GetKubeLoginClusterKubeconfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kube login cluster kubeconfig params
func (o *GetKubeLoginClusterKubeconfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kube login cluster kubeconfig params
func (o *GetKubeLoginClusterKubeconfigParams) WithContext(ctx context.Context) *GetKubeLoginClusterKubeconfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kube login cluster kubeconfig params
func (o *GetKubeLoginClusterKubeconfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kube login cluster kubeconfig params
func (o *GetKubeLoginClusterKubeconfigParams) WithHTTPClient(client *http.Client) *GetKubeLoginClusterKubeconfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kube login cluster kubeconfig params
func (o *GetKubeLoginClusterKubeconfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get kube login cluster kubeconfig params
func (o *GetKubeLoginClusterKubeconfigParams) WithClusterID(clusterID string) *GetKubeLoginClusterKubeconfigParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get kube login cluster kubeconfig params
func (o *GetKubeLoginClusterKubeconfigParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the get kube login cluster kubeconfig params
func (o *GetKubeLoginClusterKubeconfigParams) WithDC(dc string) *GetKubeLoginClusterKubeconfigParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the get kube login cluster kubeconfig params
func (o *GetKubeLoginClusterKubeconfigParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the get kube login cluster kubeconfig params
func (o *GetKubeLoginClusterKubeconfigParams) WithProjectID(projectID string) *GetKubeLoginClusterKubeconfigParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get kube login cluster kubeconfig params
func (o *GetKubeLoginClusterKubeconfigParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubeLoginClusterKubeconfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
