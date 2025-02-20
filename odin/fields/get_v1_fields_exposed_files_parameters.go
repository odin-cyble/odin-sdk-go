// Code generated by go-swagger; DO NOT EDIT.

package fields

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

// NewGetV1FieldsExposedFilesParams creates a new GetV1FieldsExposedFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1FieldsExposedFilesParams() *GetV1FieldsExposedFilesParams {
	return &GetV1FieldsExposedFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1FieldsExposedFilesParamsWithTimeout creates a new GetV1FieldsExposedFilesParams object
// with the ability to set a timeout on a request.
func NewGetV1FieldsExposedFilesParamsWithTimeout(timeout time.Duration) *GetV1FieldsExposedFilesParams {
	return &GetV1FieldsExposedFilesParams{
		timeout: timeout,
	}
}

// NewGetV1FieldsExposedFilesParamsWithContext creates a new GetV1FieldsExposedFilesParams object
// with the ability to set a context for a request.
func NewGetV1FieldsExposedFilesParamsWithContext(ctx context.Context) *GetV1FieldsExposedFilesParams {
	return &GetV1FieldsExposedFilesParams{
		Context: ctx,
	}
}

// NewGetV1FieldsExposedFilesParamsWithHTTPClient creates a new GetV1FieldsExposedFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1FieldsExposedFilesParamsWithHTTPClient(client *http.Client) *GetV1FieldsExposedFilesParams {
	return &GetV1FieldsExposedFilesParams{
		HTTPClient: client,
	}
}

/*
GetV1FieldsExposedFilesParams contains all the parameters to send to the API endpoint

	for the get v1 fields exposed files operation.

	Typically these are written to a http.Request.
*/
type GetV1FieldsExposedFilesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 fields exposed files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1FieldsExposedFilesParams) WithDefaults() *GetV1FieldsExposedFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 fields exposed files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1FieldsExposedFilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 fields exposed files params
func (o *GetV1FieldsExposedFilesParams) WithTimeout(timeout time.Duration) *GetV1FieldsExposedFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 fields exposed files params
func (o *GetV1FieldsExposedFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 fields exposed files params
func (o *GetV1FieldsExposedFilesParams) WithContext(ctx context.Context) *GetV1FieldsExposedFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 fields exposed files params
func (o *GetV1FieldsExposedFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 fields exposed files params
func (o *GetV1FieldsExposedFilesParams) WithHTTPClient(client *http.Client) *GetV1FieldsExposedFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 fields exposed files params
func (o *GetV1FieldsExposedFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1FieldsExposedFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
