// Code generated by go-swagger; DO NOT EDIT.

package domain

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

	"github.com/cybledev/odin-sdk-go/models"
)

// NewPostV1DomainSubdomainCountParams creates a new PostV1DomainSubdomainCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1DomainSubdomainCountParams() *PostV1DomainSubdomainCountParams {
	return &PostV1DomainSubdomainCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1DomainSubdomainCountParamsWithTimeout creates a new PostV1DomainSubdomainCountParams object
// with the ability to set a timeout on a request.
func NewPostV1DomainSubdomainCountParamsWithTimeout(timeout time.Duration) *PostV1DomainSubdomainCountParams {
	return &PostV1DomainSubdomainCountParams{
		timeout: timeout,
	}
}

// NewPostV1DomainSubdomainCountParamsWithContext creates a new PostV1DomainSubdomainCountParams object
// with the ability to set a context for a request.
func NewPostV1DomainSubdomainCountParamsWithContext(ctx context.Context) *PostV1DomainSubdomainCountParams {
	return &PostV1DomainSubdomainCountParams{
		Context: ctx,
	}
}

// NewPostV1DomainSubdomainCountParamsWithHTTPClient creates a new PostV1DomainSubdomainCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1DomainSubdomainCountParamsWithHTTPClient(client *http.Client) *PostV1DomainSubdomainCountParams {
	return &PostV1DomainSubdomainCountParams{
		HTTPClient: client,
	}
}

/*
PostV1DomainSubdomainCountParams contains all the parameters to send to the API endpoint

	for the post v1 domain subdomain count operation.

	Typically these are written to a http.Request.
*/
type PostV1DomainSubdomainCountParams struct {

	/* Query.

	   Query
	*/
	Query *models.DNSDNSCountRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 domain subdomain count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1DomainSubdomainCountParams) WithDefaults() *PostV1DomainSubdomainCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 domain subdomain count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1DomainSubdomainCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 domain subdomain count params
func (o *PostV1DomainSubdomainCountParams) WithTimeout(timeout time.Duration) *PostV1DomainSubdomainCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 domain subdomain count params
func (o *PostV1DomainSubdomainCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 domain subdomain count params
func (o *PostV1DomainSubdomainCountParams) WithContext(ctx context.Context) *PostV1DomainSubdomainCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 domain subdomain count params
func (o *PostV1DomainSubdomainCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 domain subdomain count params
func (o *PostV1DomainSubdomainCountParams) WithHTTPClient(client *http.Client) *PostV1DomainSubdomainCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 domain subdomain count params
func (o *PostV1DomainSubdomainCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post v1 domain subdomain count params
func (o *PostV1DomainSubdomainCountParams) WithQuery(query *models.DNSDNSCountRequest) *PostV1DomainSubdomainCountParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post v1 domain subdomain count params
func (o *PostV1DomainSubdomainCountParams) SetQuery(query *models.DNSDNSCountRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1DomainSubdomainCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Query != nil {
		if err := r.SetBodyParam(o.Query); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
