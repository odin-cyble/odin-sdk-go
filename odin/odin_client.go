// Code generated by go-swagger; DO NOT EDIT.

package odin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cybledev/odin-sdk-go/odin/certificate"
	"github.com/cybledev/odin-sdk-go/odin/domain"
	"github.com/cybledev/odin-sdk-go/odin/exposed_buckets"
	"github.com/cybledev/odin-sdk-go/odin/exposed_files"
	"github.com/cybledev/odin-sdk-go/odin/fields"
	"github.com/cybledev/odin-sdk-go/odin/health"
	"github.com/cybledev/odin-sdk-go/odin/hosts"
)

// Default odin HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.odin.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new odin HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Odin {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new odin HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Odin {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new odin client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Odin {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Odin)
	cli.Transport = transport
	cli.Certificate = certificate.New(transport, formats)
	cli.Domain = domain.New(transport, formats)
	cli.ExposedBuckets = exposed_buckets.New(transport, formats)
	cli.ExposedFiles = exposed_files.New(transport, formats)
	cli.Fields = fields.New(transport, formats)
	cli.Health = health.New(transport, formats)
	cli.Hosts = hosts.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Odin is a client for odin
type Odin struct {
	Certificate certificate.ClientService

	Domain domain.ClientService

	ExposedBuckets exposed_buckets.ClientService

	ExposedFiles exposed_files.ClientService

	Fields fields.ClientService

	Health health.ClientService

	Hosts hosts.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Odin) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Certificate.SetTransport(transport)
	c.Domain.SetTransport(transport)
	c.ExposedBuckets.SetTransport(transport)
	c.ExposedFiles.SetTransport(transport)
	c.Fields.SetTransport(transport)
	c.Health.SetTransport(transport)
	c.Hosts.SetTransport(transport)
}
