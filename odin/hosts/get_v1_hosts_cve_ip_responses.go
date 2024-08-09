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

// GetV1HostsCveIPReader is a Reader for the GetV1HostsCveIP structure.
type GetV1HostsCveIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1HostsCveIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1HostsCveIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1HostsCveIPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewGetV1HostsCveIPPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetV1HostsCveIPRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1HostsCveIPInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/hosts/cve/{ip}/] GetV1HostsCveIP", response, response.Code())
	}
}

// NewGetV1HostsCveIPOK creates a GetV1HostsCveIPOK with default headers values
func NewGetV1HostsCveIPOK() *GetV1HostsCveIPOK {
	return &GetV1HostsCveIPOK{}
}

/*
GetV1HostsCveIPOK describes a response with status code 200, with default header values.

OK
*/
type GetV1HostsCveIPOK struct {
	Payload *models.IpservicesIPCveResponse
}

// IsSuccess returns true when this get v1 hosts cve Ip o k response has a 2xx status code
func (o *GetV1HostsCveIPOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 hosts cve Ip o k response has a 3xx status code
func (o *GetV1HostsCveIPOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts cve Ip o k response has a 4xx status code
func (o *GetV1HostsCveIPOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 hosts cve Ip o k response has a 5xx status code
func (o *GetV1HostsCveIPOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 hosts cve Ip o k response a status code equal to that given
func (o *GetV1HostsCveIPOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 hosts cve Ip o k response
func (o *GetV1HostsCveIPOK) Code() int {
	return 200
}

func (o *GetV1HostsCveIPOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cve/{ip}/][%d] getV1HostsCveIpOK %s", 200, payload)
}

func (o *GetV1HostsCveIPOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cve/{ip}/][%d] getV1HostsCveIpOK %s", 200, payload)
}

func (o *GetV1HostsCveIPOK) GetPayload() *models.IpservicesIPCveResponse {
	return o.Payload
}

func (o *GetV1HostsCveIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpservicesIPCveResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1HostsCveIPBadRequest creates a GetV1HostsCveIPBadRequest with default headers values
func NewGetV1HostsCveIPBadRequest() *GetV1HostsCveIPBadRequest {
	return &GetV1HostsCveIPBadRequest{}
}

/*
GetV1HostsCveIPBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1HostsCveIPBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 hosts cve Ip bad request response has a 2xx status code
func (o *GetV1HostsCveIPBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 hosts cve Ip bad request response has a 3xx status code
func (o *GetV1HostsCveIPBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts cve Ip bad request response has a 4xx status code
func (o *GetV1HostsCveIPBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 hosts cve Ip bad request response has a 5xx status code
func (o *GetV1HostsCveIPBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 hosts cve Ip bad request response a status code equal to that given
func (o *GetV1HostsCveIPBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 hosts cve Ip bad request response
func (o *GetV1HostsCveIPBadRequest) Code() int {
	return 400
}

func (o *GetV1HostsCveIPBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cve/{ip}/][%d] getV1HostsCveIpBadRequest %s", 400, payload)
}

func (o *GetV1HostsCveIPBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cve/{ip}/][%d] getV1HostsCveIpBadRequest %s", 400, payload)
}

func (o *GetV1HostsCveIPBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1HostsCveIPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1HostsCveIPPaymentRequired creates a GetV1HostsCveIPPaymentRequired with default headers values
func NewGetV1HostsCveIPPaymentRequired() *GetV1HostsCveIPPaymentRequired {
	return &GetV1HostsCveIPPaymentRequired{}
}

/*
GetV1HostsCveIPPaymentRequired describes a response with status code 402, with default header values.

Payment Required
*/
type GetV1HostsCveIPPaymentRequired struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 hosts cve Ip payment required response has a 2xx status code
func (o *GetV1HostsCveIPPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 hosts cve Ip payment required response has a 3xx status code
func (o *GetV1HostsCveIPPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts cve Ip payment required response has a 4xx status code
func (o *GetV1HostsCveIPPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 hosts cve Ip payment required response has a 5xx status code
func (o *GetV1HostsCveIPPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 hosts cve Ip payment required response a status code equal to that given
func (o *GetV1HostsCveIPPaymentRequired) IsCode(code int) bool {
	return code == 402
}

// Code gets the status code for the get v1 hosts cve Ip payment required response
func (o *GetV1HostsCveIPPaymentRequired) Code() int {
	return 402
}

func (o *GetV1HostsCveIPPaymentRequired) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cve/{ip}/][%d] getV1HostsCveIpPaymentRequired %s", 402, payload)
}

func (o *GetV1HostsCveIPPaymentRequired) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cve/{ip}/][%d] getV1HostsCveIpPaymentRequired %s", 402, payload)
}

func (o *GetV1HostsCveIPPaymentRequired) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1HostsCveIPPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1HostsCveIPRequestTimeout creates a GetV1HostsCveIPRequestTimeout with default headers values
func NewGetV1HostsCveIPRequestTimeout() *GetV1HostsCveIPRequestTimeout {
	return &GetV1HostsCveIPRequestTimeout{}
}

/*
GetV1HostsCveIPRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetV1HostsCveIPRequestTimeout struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 hosts cve Ip request timeout response has a 2xx status code
func (o *GetV1HostsCveIPRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 hosts cve Ip request timeout response has a 3xx status code
func (o *GetV1HostsCveIPRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts cve Ip request timeout response has a 4xx status code
func (o *GetV1HostsCveIPRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 hosts cve Ip request timeout response has a 5xx status code
func (o *GetV1HostsCveIPRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 hosts cve Ip request timeout response a status code equal to that given
func (o *GetV1HostsCveIPRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get v1 hosts cve Ip request timeout response
func (o *GetV1HostsCveIPRequestTimeout) Code() int {
	return 408
}

func (o *GetV1HostsCveIPRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cve/{ip}/][%d] getV1HostsCveIpRequestTimeout %s", 408, payload)
}

func (o *GetV1HostsCveIPRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cve/{ip}/][%d] getV1HostsCveIpRequestTimeout %s", 408, payload)
}

func (o *GetV1HostsCveIPRequestTimeout) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1HostsCveIPRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1HostsCveIPInternalServerError creates a GetV1HostsCveIPInternalServerError with default headers values
func NewGetV1HostsCveIPInternalServerError() *GetV1HostsCveIPInternalServerError {
	return &GetV1HostsCveIPInternalServerError{}
}

/*
GetV1HostsCveIPInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1HostsCveIPInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 hosts cve Ip internal server error response has a 2xx status code
func (o *GetV1HostsCveIPInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 hosts cve Ip internal server error response has a 3xx status code
func (o *GetV1HostsCveIPInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts cve Ip internal server error response has a 4xx status code
func (o *GetV1HostsCveIPInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 hosts cve Ip internal server error response has a 5xx status code
func (o *GetV1HostsCveIPInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 hosts cve Ip internal server error response a status code equal to that given
func (o *GetV1HostsCveIPInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 hosts cve Ip internal server error response
func (o *GetV1HostsCveIPInternalServerError) Code() int {
	return 500
}

func (o *GetV1HostsCveIPInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cve/{ip}/][%d] getV1HostsCveIpInternalServerError %s", 500, payload)
}

func (o *GetV1HostsCveIPInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/cve/{ip}/][%d] getV1HostsCveIpInternalServerError %s", 500, payload)
}

func (o *GetV1HostsCveIPInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1HostsCveIPInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
