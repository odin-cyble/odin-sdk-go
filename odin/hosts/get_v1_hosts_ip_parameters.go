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

// NewGetV1HostsIPParams creates a new GetV1HostsIPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1HostsIPParams() *GetV1HostsIPParams {
	return &GetV1HostsIPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1HostsIPParamsWithTimeout creates a new GetV1HostsIPParams object
// with the ability to set a timeout on a request.
func NewGetV1HostsIPParamsWithTimeout(timeout time.Duration) *GetV1HostsIPParams {
	return &GetV1HostsIPParams{
		timeout: timeout,
	}
}

// NewGetV1HostsIPParamsWithContext creates a new GetV1HostsIPParams object
// with the ability to set a context for a request.
func NewGetV1HostsIPParamsWithContext(ctx context.Context) *GetV1HostsIPParams {
	return &GetV1HostsIPParams{
		Context: ctx,
	}
}

// NewGetV1HostsIPParamsWithHTTPClient creates a new GetV1HostsIPParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1HostsIPParamsWithHTTPClient(client *http.Client) *GetV1HostsIPParams {
	return &GetV1HostsIPParams{
		HTTPClient: client,
	}
}

/*
GetV1HostsIPParams contains all the parameters to send to the API endpoint

	for the get v1 hosts IP operation.

	Typically these are written to a http.Request.
*/
type GetV1HostsIPParams struct {

	/* IP.

	   get the ip
	*/
	IP string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 hosts IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1HostsIPParams) WithDefaults() *GetV1HostsIPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 hosts IP params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1HostsIPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 hosts IP params
func (o *GetV1HostsIPParams) WithTimeout(timeout time.Duration) *GetV1HostsIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 hosts IP params
func (o *GetV1HostsIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 hosts IP params
func (o *GetV1HostsIPParams) WithContext(ctx context.Context) *GetV1HostsIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 hosts IP params
func (o *GetV1HostsIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 hosts IP params
func (o *GetV1HostsIPParams) WithHTTPClient(client *http.Client) *GetV1HostsIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 hosts IP params
func (o *GetV1HostsIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIP adds the ip to the get v1 hosts IP params
func (o *GetV1HostsIPParams) WithIP(ip string) *GetV1HostsIPParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the get v1 hosts IP params
func (o *GetV1HostsIPParams) SetIP(ip string) {
	o.IP = ip
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1HostsIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ip
	if err := r.SetPathParam("ip", o.IP); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
