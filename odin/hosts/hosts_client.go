// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new hosts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new hosts API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new hosts API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for hosts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1CvesAllIPPage(params *GetV1CvesAllIPPageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1CvesAllIPPageOK, error)

	GetV1HostsCveIP(params *GetV1HostsCveIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1HostsCveIPOK, error)

	GetV1HostsCvesIPCve(params *GetV1HostsCvesIPCveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1HostsCvesIPCveOK, error)

	GetV1HostsExploitsIP(params *GetV1HostsExploitsIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1HostsExploitsIPOK, error)

	GetV1HostsExploitsIPCve(params *GetV1HostsExploitsIPCveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1HostsExploitsIPCveOK, error)

	GetV1HostsIP(params *GetV1HostsIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1HostsIPOK, error)

	PostV1HostsCount(params *PostV1HostsCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1HostsCountOK, error)

	PostV1HostsSearch(params *PostV1HostsSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1HostsSearchOK, error)

	PostV1HostsSummary(params *PostV1HostsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1HostsSummaryOK, error)

	PostV2HostsCount(params *PostV2HostsCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsCountOK, error)

	PostV2HostsIP(params *PostV2HostsIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsIPOK, error)

	PostV2HostsSearch(params *PostV2HostsSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsSearchOK, error)

	PostV2HostsSummary(params *PostV2HostsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsSummaryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetV1CvesAllIPPage gets cve details

Get the detailed cve details for ip based on page
*/
func (a *Client) GetV1CvesAllIPPage(params *GetV1CvesAllIPPageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1CvesAllIPPageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1CvesAllIPPageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1CvesAllIPPage",
		Method:             "GET",
		PathPattern:        "/v1/cves/all/{ip}/{page}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1CvesAllIPPageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1CvesAllIPPageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1CvesAllIPPage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1HostsCveIP gets ip cve details

Get detailed lost of cve ips
*/
func (a *Client) GetV1HostsCveIP(params *GetV1HostsCveIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1HostsCveIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1HostsCveIPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1HostsCveIP",
		Method:             "GET",
		PathPattern:        "/v1/hosts/cve/{ip}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1HostsCveIPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1HostsCveIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1HostsCveIP: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1HostsCvesIPCve gets cve

Get the detailed cve based on ip and cve
*/
func (a *Client) GetV1HostsCvesIPCve(params *GetV1HostsCvesIPCveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1HostsCvesIPCveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1HostsCvesIPCveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1HostsCvesIPCve",
		Method:             "GET",
		PathPattern:        "/v1/hosts/cves/{ip}/{cve}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1HostsCvesIPCveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1HostsCvesIPCveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1HostsCvesIPCve: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1HostsExploitsIP gets exploits for ip

Get the detailed list of exploits for the ip
*/
func (a *Client) GetV1HostsExploitsIP(params *GetV1HostsExploitsIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1HostsExploitsIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1HostsExploitsIPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1HostsExploitsIP",
		Method:             "GET",
		PathPattern:        "/v1/hosts/exploits/{ip}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1HostsExploitsIPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1HostsExploitsIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1HostsExploitsIP: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1HostsExploitsIPCve gets exploits for ip and cve

Get the detailed exploit details for ip and cve
*/
func (a *Client) GetV1HostsExploitsIPCve(params *GetV1HostsExploitsIPCveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1HostsExploitsIPCveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1HostsExploitsIPCveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1HostsExploitsIPCve",
		Method:             "GET",
		PathPattern:        "/v1/hosts/exploits/{ip}/{cve}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1HostsExploitsIPCveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1HostsExploitsIPCveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1HostsExploitsIPCve: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV1HostsIP gets the latest ip details

Get the complete and latest ip details
*/
func (a *Client) GetV1HostsIP(params *GetV1HostsIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetV1HostsIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1HostsIPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1HostsIP",
		Method:             "GET",
		PathPattern:        "/v1/hosts/{ip}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV1HostsIPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1HostsIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1HostsIP: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1HostsCount gets the record count

Returns the total no of records based on query
*/
func (a *Client) PostV1HostsCount(params *PostV1HostsCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1HostsCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1HostsCountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV1HostsCount",
		Method:             "POST",
		PathPattern:        "/v1/hosts/count",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1HostsCountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV1HostsCountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV1HostsCount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1HostsSearch searches hosts

Get the detailed list of hosts based on query
*/
func (a *Client) PostV1HostsSearch(params *PostV1HostsSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1HostsSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1HostsSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV1HostsSearch",
		Method:             "POST",
		PathPattern:        "/v1/hosts/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1HostsSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV1HostsSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV1HostsSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV1HostsSummary gets summary

Get the summary of hosts based on the field
*/
func (a *Client) PostV1HostsSummary(params *PostV1HostsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV1HostsSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1HostsSummaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV1HostsSummary",
		Method:             "POST",
		PathPattern:        "/v1/hosts/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV1HostsSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV1HostsSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV1HostsSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV2HostsCount fetches the record count

Returns the total no of records based on query
*/
func (a *Client) PostV2HostsCount(params *PostV2HostsCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV2HostsCountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV2HostsCount",
		Method:             "POST",
		PathPattern:        "/v2/hosts/count",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV2HostsCountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV2HostsCountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV2HostsCount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV2HostsIP fetches the latest ip details

Returns the complete ip details
*/
func (a *Client) PostV2HostsIP(params *PostV2HostsIPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsIPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV2HostsIPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV2HostsIP",
		Method:             "POST",
		PathPattern:        "/v2/hosts/{ip}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV2HostsIPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV2HostsIPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV2HostsIP: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV2HostsSearch fetches the record based on query

Returns the record based on query and blank query will return all records. It uses es searchafter for the pagination.
*/
func (a *Client) PostV2HostsSearch(params *PostV2HostsSearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV2HostsSearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV2HostsSearch",
		Method:             "POST",
		PathPattern:        "/v2/hosts/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV2HostsSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV2HostsSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV2HostsSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostV2HostsSummary creates the summary of the field based on query

Returns the summary of the field based on query
*/
func (a *Client) PostV2HostsSummary(params *PostV2HostsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostV2HostsSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV2HostsSummaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV2HostsSummary",
		Method:             "POST",
		PathPattern:        "/v2/hosts/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostV2HostsSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV2HostsSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV2HostsSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
