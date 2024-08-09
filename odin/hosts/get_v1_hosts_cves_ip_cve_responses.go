// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cybledev/odin-sdk-go/models"
)

// GetV1HostsCvesIPCveReader is a Reader for the GetV1HostsCvesIPCve structure.
type GetV1HostsCvesIPCveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1HostsCvesIPCveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1HostsCvesIPCveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1HostsCvesIPCveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetV1HostsCvesIPCveRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1HostsCvesIPCveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/hosts/cves/{ip}/{cve}] GetV1HostsCvesIPCve", response, response.Code())
	}
}

// NewGetV1HostsCvesIPCveOK creates a GetV1HostsCvesIPCveOK with default headers values
func NewGetV1HostsCvesIPCveOK() *GetV1HostsCvesIPCveOK {
	return &GetV1HostsCvesIPCveOK{}
}

/*
GetV1HostsCvesIPCveOK describes a response with status code 200, with default header values.

OK
*/
type GetV1HostsCvesIPCveOK struct {
	Payload *models.IpservicesIPCveResponse
}

// IsSuccess returns true when this get v1 hosts cves Ip cve o k response has a 2xx status code
func (o *GetV1HostsCvesIPCveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 hosts cves Ip cve o k response has a 3xx status code
func (o *GetV1HostsCvesIPCveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts cves Ip cve o k response has a 4xx status code
func (o *GetV1HostsCvesIPCveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 hosts cves Ip cve o k response has a 5xx status code
func (o *GetV1HostsCvesIPCveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 hosts cves Ip cve o k response a status code equal to that given
func (o *GetV1HostsCvesIPCveOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 hosts cves Ip cve o k response
func (o *GetV1HostsCvesIPCveOK) Code() int {
	return 200
}

func (o *GetV1HostsCvesIPCveOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cves/{ip}/{cve}][%d] getV1HostsCvesIpCveOK %s", 200, payload)
}

func (o *GetV1HostsCvesIPCveOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cves/{ip}/{cve}][%d] getV1HostsCvesIpCveOK %s", 200, payload)
}

func (o *GetV1HostsCvesIPCveOK) GetPayload() *models.IpservicesIPCveResponse {
	return o.Payload
}

func (o *GetV1HostsCvesIPCveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpservicesIPCveResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1HostsCvesIPCveBadRequest creates a GetV1HostsCvesIPCveBadRequest with default headers values
func NewGetV1HostsCvesIPCveBadRequest() *GetV1HostsCvesIPCveBadRequest {
	return &GetV1HostsCvesIPCveBadRequest{}
}

/*
GetV1HostsCvesIPCveBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1HostsCvesIPCveBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 hosts cves Ip cve bad request response has a 2xx status code
func (o *GetV1HostsCvesIPCveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 hosts cves Ip cve bad request response has a 3xx status code
func (o *GetV1HostsCvesIPCveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts cves Ip cve bad request response has a 4xx status code
func (o *GetV1HostsCvesIPCveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 hosts cves Ip cve bad request response has a 5xx status code
func (o *GetV1HostsCvesIPCveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 hosts cves Ip cve bad request response a status code equal to that given
func (o *GetV1HostsCvesIPCveBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 hosts cves Ip cve bad request response
func (o *GetV1HostsCvesIPCveBadRequest) Code() int {
	return 400
}

func (o *GetV1HostsCvesIPCveBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cves/{ip}/{cve}][%d] getV1HostsCvesIpCveBadRequest %s", 400, payload)
}

func (o *GetV1HostsCvesIPCveBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cves/{ip}/{cve}][%d] getV1HostsCvesIpCveBadRequest %s", 400, payload)
}

func (o *GetV1HostsCvesIPCveBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1HostsCvesIPCveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1HostsCvesIPCveRequestTimeout creates a GetV1HostsCvesIPCveRequestTimeout with default headers values
func NewGetV1HostsCvesIPCveRequestTimeout() *GetV1HostsCvesIPCveRequestTimeout {
	return &GetV1HostsCvesIPCveRequestTimeout{}
}

/*
GetV1HostsCvesIPCveRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetV1HostsCvesIPCveRequestTimeout struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 hosts cves Ip cve request timeout response has a 2xx status code
func (o *GetV1HostsCvesIPCveRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 hosts cves Ip cve request timeout response has a 3xx status code
func (o *GetV1HostsCvesIPCveRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts cves Ip cve request timeout response has a 4xx status code
func (o *GetV1HostsCvesIPCveRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 hosts cves Ip cve request timeout response has a 5xx status code
func (o *GetV1HostsCvesIPCveRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 hosts cves Ip cve request timeout response a status code equal to that given
func (o *GetV1HostsCvesIPCveRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get v1 hosts cves Ip cve request timeout response
func (o *GetV1HostsCvesIPCveRequestTimeout) Code() int {
	return 408
}

func (o *GetV1HostsCvesIPCveRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cves/{ip}/{cve}][%d] getV1HostsCvesIpCveRequestTimeout %s", 408, payload)
}

func (o *GetV1HostsCvesIPCveRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cves/{ip}/{cve}][%d] getV1HostsCvesIpCveRequestTimeout %s", 408, payload)
}

func (o *GetV1HostsCvesIPCveRequestTimeout) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1HostsCvesIPCveRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1HostsCvesIPCveInternalServerError creates a GetV1HostsCvesIPCveInternalServerError with default headers values
func NewGetV1HostsCvesIPCveInternalServerError() *GetV1HostsCvesIPCveInternalServerError {
	return &GetV1HostsCvesIPCveInternalServerError{}
}

/*
GetV1HostsCvesIPCveInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1HostsCvesIPCveInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 hosts cves Ip cve internal server error response has a 2xx status code
func (o *GetV1HostsCvesIPCveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 hosts cves Ip cve internal server error response has a 3xx status code
func (o *GetV1HostsCvesIPCveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts cves Ip cve internal server error response has a 4xx status code
func (o *GetV1HostsCvesIPCveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 hosts cves Ip cve internal server error response has a 5xx status code
func (o *GetV1HostsCvesIPCveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 hosts cves Ip cve internal server error response a status code equal to that given
func (o *GetV1HostsCvesIPCveInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 hosts cves Ip cve internal server error response
func (o *GetV1HostsCvesIPCveInternalServerError) Code() int {
	return 500
}

func (o *GetV1HostsCvesIPCveInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cves/{ip}/{cve}][%d] getV1HostsCvesIpCveInternalServerError %s", 500, payload)
}

func (o *GetV1HostsCvesIPCveInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cves/{ip}/{cve}][%d] getV1HostsCvesIpCveInternalServerError %s", 500, payload)
}

func (o *GetV1HostsCvesIPCveInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1HostsCvesIPCveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
