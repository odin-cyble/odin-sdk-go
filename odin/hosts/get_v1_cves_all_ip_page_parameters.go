// Code generated by go-swagger; DO NOT EDIT.

package hosts

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

// NewGetV1CvesAllIPPageParams creates a new GetV1CvesAllIPPageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1CvesAllIPPageParams() *GetV1CvesAllIPPageParams {
	return &GetV1CvesAllIPPageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1CvesAllIPPageParamsWithTimeout creates a new GetV1CvesAllIPPageParams object
// with the ability to set a timeout on a request.
func NewGetV1CvesAllIPPageParamsWithTimeout(timeout time.Duration) *GetV1CvesAllIPPageParams {
	return &GetV1CvesAllIPPageParams{
		timeout: timeout,
	}
}

// NewGetV1CvesAllIPPageParamsWithContext creates a new GetV1CvesAllIPPageParams object
// with the ability to set a context for a request.
func NewGetV1CvesAllIPPageParamsWithContext(ctx context.Context) *GetV1CvesAllIPPageParams {
	return &GetV1CvesAllIPPageParams{
		Context: ctx,
	}
}

// NewGetV1CvesAllIPPageParamsWithHTTPClient creates a new GetV1CvesAllIPPageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1CvesAllIPPageParamsWithHTTPClient(client *http.Client) *GetV1CvesAllIPPageParams {
	return &GetV1CvesAllIPPageParams{
		HTTPClient: client,
	}
}

/*
GetV1CvesAllIPPageParams contains all the parameters to send to the API endpoint

	for the get v1 cves all IP page operation.

	Typically these are written to a http.Request.
*/
type GetV1CvesAllIPPageParams struct {

	/* IP.

	   get the ip
	*/
	IP string

	/* Page.

	   get the page
	*/
	Page string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 cves all IP page params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1CvesAllIPPageParams) WithDefaults() *GetV1CvesAllIPPageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 cves all IP page params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1CvesAllIPPageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 cves all IP page params
func (o *GetV1CvesAllIPPageParams) WithTimeout(timeout time.Duration) *GetV1CvesAllIPPageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 cves all IP page params
func (o *GetV1CvesAllIPPageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 cves all IP page params
func (o *GetV1CvesAllIPPageParams) WithContext(ctx context.Context) *GetV1CvesAllIPPageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 cves all IP page params
func (o *GetV1CvesAllIPPageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 cves all IP page params
func (o *GetV1CvesAllIPPageParams) WithHTTPClient(client *http.Client) *GetV1CvesAllIPPageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 cves all IP page params
func (o *GetV1CvesAllIPPageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the get v1 cves all IP page params
func (o *GetV1CvesAllIPPageParams) WithIP(ip string) *GetV1CvesAllIPPageParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the get v1 cves all IP page params
func (o *GetV1CvesAllIPPageParams) SetIP(ip string) {
	o.IP = ip
}

// WithPage adds the page to the get v1 cves all IP page params
func (o *GetV1CvesAllIPPageParams) WithPage(page string) *GetV1CvesAllIPPageParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 cves all IP page params
func (o *GetV1CvesAllIPPageParams) SetPage(page string) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1CvesAllIPPageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ip
	if err := r.SetPathParam("ip", o.IP); err != nil {
		return err
	}

	// path param page
	if err := r.SetPathParam("page", o.Page); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
