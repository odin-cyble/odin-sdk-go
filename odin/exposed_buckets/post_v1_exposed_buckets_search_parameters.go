// Code generated by go-swagger; DO NOT EDIT.

package exposed_buckets

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

// NewPostV1ExposedBucketsSearchParams creates a new PostV1ExposedBucketsSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1ExposedBucketsSearchParams() *PostV1ExposedBucketsSearchParams {
	return &PostV1ExposedBucketsSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1ExposedBucketsSearchParamsWithTimeout creates a new PostV1ExposedBucketsSearchParams object
// with the ability to set a timeout on a request.
func NewPostV1ExposedBucketsSearchParamsWithTimeout(timeout time.Duration) *PostV1ExposedBucketsSearchParams {
	return &PostV1ExposedBucketsSearchParams{
		timeout: timeout,
	}
}

// NewPostV1ExposedBucketsSearchParamsWithContext creates a new PostV1ExposedBucketsSearchParams object
// with the ability to set a context for a request.
func NewPostV1ExposedBucketsSearchParamsWithContext(ctx context.Context) *PostV1ExposedBucketsSearchParams {
	return &PostV1ExposedBucketsSearchParams{
		Context: ctx,
	}
}

// NewPostV1ExposedBucketsSearchParamsWithHTTPClient creates a new PostV1ExposedBucketsSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1ExposedBucketsSearchParamsWithHTTPClient(client *http.Client) *PostV1ExposedBucketsSearchParams {
	return &PostV1ExposedBucketsSearchParams{
		HTTPClient: client,
	}
}

/*
PostV1ExposedBucketsSearchParams contains all the parameters to send to the API endpoint

	for the post v1 exposed buckets search operation.

	Typically these are written to a http.Request.
*/
type PostV1ExposedBucketsSearchParams struct {

	/* Query.

	   Search Query
	*/
	Query *models.ExposedSearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 exposed buckets search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ExposedBucketsSearchParams) WithDefaults() *PostV1ExposedBucketsSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 exposed buckets search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ExposedBucketsSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 exposed buckets search params
func (o *PostV1ExposedBucketsSearchParams) WithTimeout(timeout time.Duration) *PostV1ExposedBucketsSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 exposed buckets search params
func (o *PostV1ExposedBucketsSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 exposed buckets search params
func (o *PostV1ExposedBucketsSearchParams) WithContext(ctx context.Context) *PostV1ExposedBucketsSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 exposed buckets search params
func (o *PostV1ExposedBucketsSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 exposed buckets search params
func (o *PostV1ExposedBucketsSearchParams) WithHTTPClient(client *http.Client) *PostV1ExposedBucketsSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 exposed buckets search params
func (o *PostV1ExposedBucketsSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post v1 exposed buckets search params
func (o *PostV1ExposedBucketsSearchParams) WithQuery(query *models.ExposedSearchRequest) *PostV1ExposedBucketsSearchParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post v1 exposed buckets search params
func (o *PostV1ExposedBucketsSearchParams) SetQuery(query *models.ExposedSearchRequest) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1ExposedBucketsSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
