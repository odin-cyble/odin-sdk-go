// Code generated by go-swagger; DO NOT EDIT.

package certificate

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

// NewPostV1CertificatesCountParams creates a new PostV1CertificatesCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1CertificatesCountParams() *PostV1CertificatesCountParams {
	return &PostV1CertificatesCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1CertificatesCountParamsWithTimeout creates a new PostV1CertificatesCountParams object
// with the ability to set a timeout on a request.
func NewPostV1CertificatesCountParamsWithTimeout(timeout time.Duration) *PostV1CertificatesCountParams {
	return &PostV1CertificatesCountParams{
		timeout: timeout,
	}
}

// NewPostV1CertificatesCountParamsWithContext creates a new PostV1CertificatesCountParams object
// with the ability to set a context for a request.
func NewPostV1CertificatesCountParamsWithContext(ctx context.Context) *PostV1CertificatesCountParams {
	return &PostV1CertificatesCountParams{
		Context: ctx,
	}
}

// NewPostV1CertificatesCountParamsWithHTTPClient creates a new PostV1CertificatesCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1CertificatesCountParamsWithHTTPClient(client *http.Client) *PostV1CertificatesCountParams {
	return &PostV1CertificatesCountParams{
		HTTPClient: client,
	}
}

/*
PostV1CertificatesCountParams contains all the parameters to send to the API endpoint

	for the post v1 certificates count operation.

	Typically these are written to a http.Request.
*/
type PostV1CertificatesCountParams struct {

	/* Query.

	   Count Query
	*/
	Query *models.CertificateCertCountRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 certificates count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1CertificatesCountParams) WithDefaults() *PostV1CertificatesCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 certificates count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1CertificatesCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 certificates count params
func (o *PostV1CertificatesCountParams) WithTimeout(timeout time.Duration) *PostV1CertificatesCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 certificates count params
func (o *PostV1CertificatesCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 certificates count params
func (o *PostV1CertificatesCountParams) WithContext(ctx context.Context) *PostV1CertificatesCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 certificates count params
func (o *PostV1CertificatesCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 certificates count params
func (o *PostV1CertificatesCountParams) WithHTTPClient(client *http.Client) *PostV1CertificatesCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 certificates count params
func (o *PostV1CertificatesCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post v1 certificates count params
func (o *PostV1CertificatesCountParams) WithQuery(query *models.CertificateCertCountRequest) *PostV1CertificatesCountParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post v1 certificates count params
func (o *PostV1CertificatesCountParams) SetQuery(query *models.CertificateCertCountRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1CertificatesCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
