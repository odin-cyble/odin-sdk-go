// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cybledev/odin-sdk-go/models"
)

// GetV1HostsIPReader is a Reader for the GetV1HostsIP structure.
type GetV1HostsIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1HostsIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1HostsIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1HostsIPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewGetV1HostsIPPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetV1HostsIPRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1HostsIPInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/hosts/{ip}/] GetV1HostsIP", response, response.Code())
	}
}

// NewGetV1HostsIPOK creates a GetV1HostsIPOK with default headers values
func NewGetV1HostsIPOK() *GetV1HostsIPOK {
	return &GetV1HostsIPOK{}
}

/*
GetV1HostsIPOK describes a response with status code 200, with default header values.

OK
*/
type GetV1HostsIPOK struct {
	Payload *GetV1HostsIPOKBody
}

// IsSuccess returns true when this get v1 hosts Ip o k response has a 2xx status code
func (o *GetV1HostsIPOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 hosts Ip o k response has a 3xx status code
func (o *GetV1HostsIPOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts Ip o k response has a 4xx status code
func (o *GetV1HostsIPOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 hosts Ip o k response has a 5xx status code
func (o *GetV1HostsIPOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 hosts Ip o k response a status code equal to that given
func (o *GetV1HostsIPOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 hosts Ip o k response
func (o *GetV1HostsIPOK) Code() int {
	return 200
}

func (o *GetV1HostsIPOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{ip}/][%d] getV1HostsIpOK %s", 200, payload)
}

func (o *GetV1HostsIPOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{ip}/][%d] getV1HostsIpOK %s", 200, payload)
}

func (o *GetV1HostsIPOK) GetPayload() *GetV1HostsIPOKBody {
	return o.Payload
}

func (o *GetV1HostsIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV1HostsIPOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1HostsIPBadRequest creates a GetV1HostsIPBadRequest with default headers values
func NewGetV1HostsIPBadRequest() *GetV1HostsIPBadRequest {
	return &GetV1HostsIPBadRequest{}
}

/*
GetV1HostsIPBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1HostsIPBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 hosts Ip bad request response has a 2xx status code
func (o *GetV1HostsIPBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 hosts Ip bad request response has a 3xx status code
func (o *GetV1HostsIPBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts Ip bad request response has a 4xx status code
func (o *GetV1HostsIPBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 hosts Ip bad request response has a 5xx status code
func (o *GetV1HostsIPBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 hosts Ip bad request response a status code equal to that given
func (o *GetV1HostsIPBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 hosts Ip bad request response
func (o *GetV1HostsIPBadRequest) Code() int {
	return 400
}

func (o *GetV1HostsIPBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{ip}/][%d] getV1HostsIpBadRequest %s", 400, payload)
}

func (o *GetV1HostsIPBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{ip}/][%d] getV1HostsIpBadRequest %s", 400, payload)
}

func (o *GetV1HostsIPBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1HostsIPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1HostsIPPaymentRequired creates a GetV1HostsIPPaymentRequired with default headers values
func NewGetV1HostsIPPaymentRequired() *GetV1HostsIPPaymentRequired {
	return &GetV1HostsIPPaymentRequired{}
}

/*
GetV1HostsIPPaymentRequired describes a response with status code 402, with default header values.

Payment Required
*/
type GetV1HostsIPPaymentRequired struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 hosts Ip payment required response has a 2xx status code
func (o *GetV1HostsIPPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 hosts Ip payment required response has a 3xx status code
func (o *GetV1HostsIPPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts Ip payment required response has a 4xx status code
func (o *GetV1HostsIPPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 hosts Ip payment required response has a 5xx status code
func (o *GetV1HostsIPPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 hosts Ip payment required response a status code equal to that given
func (o *GetV1HostsIPPaymentRequired) IsCode(code int) bool {
	return code == 402
}

// Code gets the status code for the get v1 hosts Ip payment required response
func (o *GetV1HostsIPPaymentRequired) Code() int {
	return 402
}

func (o *GetV1HostsIPPaymentRequired) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{ip}/][%d] getV1HostsIpPaymentRequired %s", 402, payload)
}

func (o *GetV1HostsIPPaymentRequired) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{ip}/][%d] getV1HostsIpPaymentRequired %s", 402, payload)
}

func (o *GetV1HostsIPPaymentRequired) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1HostsIPPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1HostsIPRequestTimeout creates a GetV1HostsIPRequestTimeout with default headers values
func NewGetV1HostsIPRequestTimeout() *GetV1HostsIPRequestTimeout {
	return &GetV1HostsIPRequestTimeout{}
}

/*
GetV1HostsIPRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetV1HostsIPRequestTimeout struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 hosts Ip request timeout response has a 2xx status code
func (o *GetV1HostsIPRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 hosts Ip request timeout response has a 3xx status code
func (o *GetV1HostsIPRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts Ip request timeout response has a 4xx status code
func (o *GetV1HostsIPRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 hosts Ip request timeout response has a 5xx status code
func (o *GetV1HostsIPRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 hosts Ip request timeout response a status code equal to that given
func (o *GetV1HostsIPRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get v1 hosts Ip request timeout response
func (o *GetV1HostsIPRequestTimeout) Code() int {
	return 408
}

func (o *GetV1HostsIPRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{ip}/][%d] getV1HostsIpRequestTimeout %s", 408, payload)
}

func (o *GetV1HostsIPRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{ip}/][%d] getV1HostsIpRequestTimeout %s", 408, payload)
}

func (o *GetV1HostsIPRequestTimeout) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1HostsIPRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1HostsIPInternalServerError creates a GetV1HostsIPInternalServerError with default headers values
func NewGetV1HostsIPInternalServerError() *GetV1HostsIPInternalServerError {
	return &GetV1HostsIPInternalServerError{}
}

/*
GetV1HostsIPInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1HostsIPInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 hosts Ip internal server error response has a 2xx status code
func (o *GetV1HostsIPInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 hosts Ip internal server error response has a 3xx status code
func (o *GetV1HostsIPInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 hosts Ip internal server error response has a 4xx status code
func (o *GetV1HostsIPInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 hosts Ip internal server error response has a 5xx status code
func (o *GetV1HostsIPInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 hosts Ip internal server error response a status code equal to that given
func (o *GetV1HostsIPInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 hosts Ip internal server error response
func (o *GetV1HostsIPInternalServerError) Code() int {
	return 500
}

func (o *GetV1HostsIPInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{ip}/][%d] getV1HostsIpInternalServerError %s", 500, payload)
}

func (o *GetV1HostsIPInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/hosts/{ip}/][%d] getV1HostsIpInternalServerError %s", 500, payload)
}

func (o *GetV1HostsIPInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1HostsIPInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetV1HostsIPOKBody get v1 hosts IP o k body
swagger:model GetV1HostsIPOKBody
*/
type GetV1HostsIPOKBody struct {
	models.APIResponse

	// data
	Data *models.IpservicesIPSummaryData `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetV1HostsIPOKBody) UnmarshalJSON(raw []byte) error {
	// GetV1HostsIPOKBodyAO0
	var getV1HostsIPOKBodyAO0 models.APIResponse
	if err := swag.ReadJSON(raw, &getV1HostsIPOKBodyAO0); err != nil {
		return err
	}
	o.APIResponse = getV1HostsIPOKBodyAO0

	// GetV1HostsIPOKBodyAO1
	var dataGetV1HostsIPOKBodyAO1 struct {
		Data *models.IpservicesIPSummaryData `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetV1HostsIPOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetV1HostsIPOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetV1HostsIPOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getV1HostsIPOKBodyAO0, err := swag.WriteJSON(o.APIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getV1HostsIPOKBodyAO0)
	var dataGetV1HostsIPOKBodyAO1 struct {
		Data *models.IpservicesIPSummaryData `json:"data,omitempty"`
	}

	dataGetV1HostsIPOKBodyAO1.Data = o.Data

	jsonDataGetV1HostsIPOKBodyAO1, errGetV1HostsIPOKBodyAO1 := swag.WriteJSON(dataGetV1HostsIPOKBodyAO1)
	if errGetV1HostsIPOKBodyAO1 != nil {
		return nil, errGetV1HostsIPOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetV1HostsIPOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get v1 hosts IP o k body
func (o *GetV1HostsIPOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.APIResponse
	if err := o.APIResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetV1HostsIPOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getV1HostsIpOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getV1HostsIpOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get v1 hosts IP o k body based on the context it is used
func (o *GetV1HostsIPOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.APIResponse
	if err := o.APIResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetV1HostsIPOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getV1HostsIpOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getV1HostsIpOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV1HostsIPOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV1HostsIPOKBody) UnmarshalBinary(b []byte) error {
	var res GetV1HostsIPOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
