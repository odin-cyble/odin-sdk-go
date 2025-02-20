// Code generated by go-swagger; DO NOT EDIT.

package fields

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cybledev/odin-sdk-go/models"
)

// GetV1FieldsHostsCategoryReader is a Reader for the GetV1FieldsHostsCategory structure.
type GetV1FieldsHostsCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1FieldsHostsCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1FieldsHostsCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1FieldsHostsCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewGetV1FieldsHostsCategoryPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetV1FieldsHostsCategoryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1FieldsHostsCategoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/fields/hosts/{category}/] GetV1FieldsHostsCategory", response, response.Code())
	}
}

// NewGetV1FieldsHostsCategoryOK creates a GetV1FieldsHostsCategoryOK with default headers values
func NewGetV1FieldsHostsCategoryOK() *GetV1FieldsHostsCategoryOK {
	return &GetV1FieldsHostsCategoryOK{}
}

/*
GetV1FieldsHostsCategoryOK describes a response with status code 200, with default header values.

OK
*/
type GetV1FieldsHostsCategoryOK struct {
	Payload *GetV1FieldsHostsCategoryOKBody
}

// IsSuccess returns true when this get v1 fields hosts category o k response has a 2xx status code
func (o *GetV1FieldsHostsCategoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 fields hosts category o k response has a 3xx status code
func (o *GetV1FieldsHostsCategoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 fields hosts category o k response has a 4xx status code
func (o *GetV1FieldsHostsCategoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 fields hosts category o k response has a 5xx status code
func (o *GetV1FieldsHostsCategoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 fields hosts category o k response a status code equal to that given
func (o *GetV1FieldsHostsCategoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 fields hosts category o k response
func (o *GetV1FieldsHostsCategoryOK) Code() int {
	return 200
}

func (o *GetV1FieldsHostsCategoryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/hosts/{category}/][%d] getV1FieldsHostsCategoryOK %s", 200, payload)
}

func (o *GetV1FieldsHostsCategoryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/hosts/{category}/][%d] getV1FieldsHostsCategoryOK %s", 200, payload)
}

func (o *GetV1FieldsHostsCategoryOK) GetPayload() *GetV1FieldsHostsCategoryOKBody {
	return o.Payload
}

func (o *GetV1FieldsHostsCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV1FieldsHostsCategoryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1FieldsHostsCategoryBadRequest creates a GetV1FieldsHostsCategoryBadRequest with default headers values
func NewGetV1FieldsHostsCategoryBadRequest() *GetV1FieldsHostsCategoryBadRequest {
	return &GetV1FieldsHostsCategoryBadRequest{}
}

/*
GetV1FieldsHostsCategoryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1FieldsHostsCategoryBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 fields hosts category bad request response has a 2xx status code
func (o *GetV1FieldsHostsCategoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 fields hosts category bad request response has a 3xx status code
func (o *GetV1FieldsHostsCategoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 fields hosts category bad request response has a 4xx status code
func (o *GetV1FieldsHostsCategoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 fields hosts category bad request response has a 5xx status code
func (o *GetV1FieldsHostsCategoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 fields hosts category bad request response a status code equal to that given
func (o *GetV1FieldsHostsCategoryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 fields hosts category bad request response
func (o *GetV1FieldsHostsCategoryBadRequest) Code() int {
	return 400
}

func (o *GetV1FieldsHostsCategoryBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/hosts/{category}/][%d] getV1FieldsHostsCategoryBadRequest %s", 400, payload)
}

func (o *GetV1FieldsHostsCategoryBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/hosts/{category}/][%d] getV1FieldsHostsCategoryBadRequest %s", 400, payload)
}

func (o *GetV1FieldsHostsCategoryBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1FieldsHostsCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1FieldsHostsCategoryPaymentRequired creates a GetV1FieldsHostsCategoryPaymentRequired with default headers values
func NewGetV1FieldsHostsCategoryPaymentRequired() *GetV1FieldsHostsCategoryPaymentRequired {
	return &GetV1FieldsHostsCategoryPaymentRequired{}
}

/*
GetV1FieldsHostsCategoryPaymentRequired describes a response with status code 402, with default header values.

Payment Required
*/
type GetV1FieldsHostsCategoryPaymentRequired struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 fields hosts category payment required response has a 2xx status code
func (o *GetV1FieldsHostsCategoryPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 fields hosts category payment required response has a 3xx status code
func (o *GetV1FieldsHostsCategoryPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 fields hosts category payment required response has a 4xx status code
func (o *GetV1FieldsHostsCategoryPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 fields hosts category payment required response has a 5xx status code
func (o *GetV1FieldsHostsCategoryPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 fields hosts category payment required response a status code equal to that given
func (o *GetV1FieldsHostsCategoryPaymentRequired) IsCode(code int) bool {
	return code == 402
}

// Code gets the status code for the get v1 fields hosts category payment required response
func (o *GetV1FieldsHostsCategoryPaymentRequired) Code() int {
	return 402
}

func (o *GetV1FieldsHostsCategoryPaymentRequired) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/hosts/{category}/][%d] getV1FieldsHostsCategoryPaymentRequired %s", 402, payload)
}

func (o *GetV1FieldsHostsCategoryPaymentRequired) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/hosts/{category}/][%d] getV1FieldsHostsCategoryPaymentRequired %s", 402, payload)
}

func (o *GetV1FieldsHostsCategoryPaymentRequired) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1FieldsHostsCategoryPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1FieldsHostsCategoryRequestTimeout creates a GetV1FieldsHostsCategoryRequestTimeout with default headers values
func NewGetV1FieldsHostsCategoryRequestTimeout() *GetV1FieldsHostsCategoryRequestTimeout {
	return &GetV1FieldsHostsCategoryRequestTimeout{}
}

/*
GetV1FieldsHostsCategoryRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetV1FieldsHostsCategoryRequestTimeout struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 fields hosts category request timeout response has a 2xx status code
func (o *GetV1FieldsHostsCategoryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 fields hosts category request timeout response has a 3xx status code
func (o *GetV1FieldsHostsCategoryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 fields hosts category request timeout response has a 4xx status code
func (o *GetV1FieldsHostsCategoryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 fields hosts category request timeout response has a 5xx status code
func (o *GetV1FieldsHostsCategoryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 fields hosts category request timeout response a status code equal to that given
func (o *GetV1FieldsHostsCategoryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get v1 fields hosts category request timeout response
func (o *GetV1FieldsHostsCategoryRequestTimeout) Code() int {
	return 408
}

func (o *GetV1FieldsHostsCategoryRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/hosts/{category}/][%d] getV1FieldsHostsCategoryRequestTimeout %s", 408, payload)
}

func (o *GetV1FieldsHostsCategoryRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/hosts/{category}/][%d] getV1FieldsHostsCategoryRequestTimeout %s", 408, payload)
}

func (o *GetV1FieldsHostsCategoryRequestTimeout) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1FieldsHostsCategoryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1FieldsHostsCategoryInternalServerError creates a GetV1FieldsHostsCategoryInternalServerError with default headers values
func NewGetV1FieldsHostsCategoryInternalServerError() *GetV1FieldsHostsCategoryInternalServerError {
	return &GetV1FieldsHostsCategoryInternalServerError{}
}

/*
GetV1FieldsHostsCategoryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1FieldsHostsCategoryInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 fields hosts category internal server error response has a 2xx status code
func (o *GetV1FieldsHostsCategoryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 fields hosts category internal server error response has a 3xx status code
func (o *GetV1FieldsHostsCategoryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 fields hosts category internal server error response has a 4xx status code
func (o *GetV1FieldsHostsCategoryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 fields hosts category internal server error response has a 5xx status code
func (o *GetV1FieldsHostsCategoryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 fields hosts category internal server error response a status code equal to that given
func (o *GetV1FieldsHostsCategoryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 fields hosts category internal server error response
func (o *GetV1FieldsHostsCategoryInternalServerError) Code() int {
	return 500
}

func (o *GetV1FieldsHostsCategoryInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/hosts/{category}/][%d] getV1FieldsHostsCategoryInternalServerError %s", 500, payload)
}

func (o *GetV1FieldsHostsCategoryInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/hosts/{category}/][%d] getV1FieldsHostsCategoryInternalServerError %s", 500, payload)
}

func (o *GetV1FieldsHostsCategoryInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1FieldsHostsCategoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetV1FieldsHostsCategoryOKBody get v1 fields hosts category o k body
swagger:model GetV1FieldsHostsCategoryOKBody
*/
type GetV1FieldsHostsCategoryOKBody struct {
	models.APIResponse

	// data
	Data []*models.Field `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetV1FieldsHostsCategoryOKBody) UnmarshalJSON(raw []byte) error {
	// GetV1FieldsHostsCategoryOKBodyAO0
	var getV1FieldsHostsCategoryOKBodyAO0 models.APIResponse
	if err := swag.ReadJSON(raw, &getV1FieldsHostsCategoryOKBodyAO0); err != nil {
		return err
	}
	o.APIResponse = getV1FieldsHostsCategoryOKBodyAO0

	// GetV1FieldsHostsCategoryOKBodyAO1
	var dataGetV1FieldsHostsCategoryOKBodyAO1 struct {
		Data []*models.Field `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataGetV1FieldsHostsCategoryOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetV1FieldsHostsCategoryOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetV1FieldsHostsCategoryOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getV1FieldsHostsCategoryOKBodyAO0, err := swag.WriteJSON(o.APIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getV1FieldsHostsCategoryOKBodyAO0)
	var dataGetV1FieldsHostsCategoryOKBodyAO1 struct {
		Data []*models.Field `json:"data"`
	}

	dataGetV1FieldsHostsCategoryOKBodyAO1.Data = o.Data

	jsonDataGetV1FieldsHostsCategoryOKBodyAO1, errGetV1FieldsHostsCategoryOKBodyAO1 := swag.WriteJSON(dataGetV1FieldsHostsCategoryOKBodyAO1)
	if errGetV1FieldsHostsCategoryOKBodyAO1 != nil {
		return nil, errGetV1FieldsHostsCategoryOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetV1FieldsHostsCategoryOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get v1 fields hosts category o k body
func (o *GetV1FieldsHostsCategoryOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetV1FieldsHostsCategoryOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1FieldsHostsCategoryOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1FieldsHostsCategoryOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get v1 fields hosts category o k body based on the context it is used
func (o *GetV1FieldsHostsCategoryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetV1FieldsHostsCategoryOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1FieldsHostsCategoryOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1FieldsHostsCategoryOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV1FieldsHostsCategoryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV1FieldsHostsCategoryOKBody) UnmarshalBinary(b []byte) error {
	var res GetV1FieldsHostsCategoryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
