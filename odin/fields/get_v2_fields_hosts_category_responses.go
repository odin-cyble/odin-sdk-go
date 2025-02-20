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

// GetV2FieldsHostsCategoryReader is a Reader for the GetV2FieldsHostsCategory structure.
type GetV2FieldsHostsCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV2FieldsHostsCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV2FieldsHostsCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV2FieldsHostsCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewGetV2FieldsHostsCategoryPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetV2FieldsHostsCategoryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV2FieldsHostsCategoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v2/fields/hosts/{category}/] GetV2FieldsHostsCategory", response, response.Code())
	}
}

// NewGetV2FieldsHostsCategoryOK creates a GetV2FieldsHostsCategoryOK with default headers values
func NewGetV2FieldsHostsCategoryOK() *GetV2FieldsHostsCategoryOK {
	return &GetV2FieldsHostsCategoryOK{}
}

/*
GetV2FieldsHostsCategoryOK describes a response with status code 200, with default header values.

OK
*/
type GetV2FieldsHostsCategoryOK struct {
	Payload *GetV2FieldsHostsCategoryOKBody
}

// IsSuccess returns true when this get v2 fields hosts category o k response has a 2xx status code
func (o *GetV2FieldsHostsCategoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v2 fields hosts category o k response has a 3xx status code
func (o *GetV2FieldsHostsCategoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v2 fields hosts category o k response has a 4xx status code
func (o *GetV2FieldsHostsCategoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v2 fields hosts category o k response has a 5xx status code
func (o *GetV2FieldsHostsCategoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v2 fields hosts category o k response a status code equal to that given
func (o *GetV2FieldsHostsCategoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v2 fields hosts category o k response
func (o *GetV2FieldsHostsCategoryOK) Code() int {
	return 200
}

func (o *GetV2FieldsHostsCategoryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/fields/hosts/{category}/][%d] getV2FieldsHostsCategoryOK %s", 200, payload)
}

func (o *GetV2FieldsHostsCategoryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/fields/hosts/{category}/][%d] getV2FieldsHostsCategoryOK %s", 200, payload)
}

func (o *GetV2FieldsHostsCategoryOK) GetPayload() *GetV2FieldsHostsCategoryOKBody {
	return o.Payload
}

func (o *GetV2FieldsHostsCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV2FieldsHostsCategoryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2FieldsHostsCategoryBadRequest creates a GetV2FieldsHostsCategoryBadRequest with default headers values
func NewGetV2FieldsHostsCategoryBadRequest() *GetV2FieldsHostsCategoryBadRequest {
	return &GetV2FieldsHostsCategoryBadRequest{}
}

/*
GetV2FieldsHostsCategoryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV2FieldsHostsCategoryBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v2 fields hosts category bad request response has a 2xx status code
func (o *GetV2FieldsHostsCategoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v2 fields hosts category bad request response has a 3xx status code
func (o *GetV2FieldsHostsCategoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v2 fields hosts category bad request response has a 4xx status code
func (o *GetV2FieldsHostsCategoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v2 fields hosts category bad request response has a 5xx status code
func (o *GetV2FieldsHostsCategoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v2 fields hosts category bad request response a status code equal to that given
func (o *GetV2FieldsHostsCategoryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v2 fields hosts category bad request response
func (o *GetV2FieldsHostsCategoryBadRequest) Code() int {
	return 400
}

func (o *GetV2FieldsHostsCategoryBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/fields/hosts/{category}/][%d] getV2FieldsHostsCategoryBadRequest %s", 400, payload)
}

func (o *GetV2FieldsHostsCategoryBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/fields/hosts/{category}/][%d] getV2FieldsHostsCategoryBadRequest %s", 400, payload)
}

func (o *GetV2FieldsHostsCategoryBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV2FieldsHostsCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2FieldsHostsCategoryPaymentRequired creates a GetV2FieldsHostsCategoryPaymentRequired with default headers values
func NewGetV2FieldsHostsCategoryPaymentRequired() *GetV2FieldsHostsCategoryPaymentRequired {
	return &GetV2FieldsHostsCategoryPaymentRequired{}
}

/*
GetV2FieldsHostsCategoryPaymentRequired describes a response with status code 402, with default header values.

Payment Required
*/
type GetV2FieldsHostsCategoryPaymentRequired struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v2 fields hosts category payment required response has a 2xx status code
func (o *GetV2FieldsHostsCategoryPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v2 fields hosts category payment required response has a 3xx status code
func (o *GetV2FieldsHostsCategoryPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v2 fields hosts category payment required response has a 4xx status code
func (o *GetV2FieldsHostsCategoryPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v2 fields hosts category payment required response has a 5xx status code
func (o *GetV2FieldsHostsCategoryPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this get v2 fields hosts category payment required response a status code equal to that given
func (o *GetV2FieldsHostsCategoryPaymentRequired) IsCode(code int) bool {
	return code == 402
}

// Code gets the status code for the get v2 fields hosts category payment required response
func (o *GetV2FieldsHostsCategoryPaymentRequired) Code() int {
	return 402
}

func (o *GetV2FieldsHostsCategoryPaymentRequired) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/fields/hosts/{category}/][%d] getV2FieldsHostsCategoryPaymentRequired %s", 402, payload)
}

func (o *GetV2FieldsHostsCategoryPaymentRequired) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/fields/hosts/{category}/][%d] getV2FieldsHostsCategoryPaymentRequired %s", 402, payload)
}

func (o *GetV2FieldsHostsCategoryPaymentRequired) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV2FieldsHostsCategoryPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2FieldsHostsCategoryRequestTimeout creates a GetV2FieldsHostsCategoryRequestTimeout with default headers values
func NewGetV2FieldsHostsCategoryRequestTimeout() *GetV2FieldsHostsCategoryRequestTimeout {
	return &GetV2FieldsHostsCategoryRequestTimeout{}
}

/*
GetV2FieldsHostsCategoryRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetV2FieldsHostsCategoryRequestTimeout struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v2 fields hosts category request timeout response has a 2xx status code
func (o *GetV2FieldsHostsCategoryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v2 fields hosts category request timeout response has a 3xx status code
func (o *GetV2FieldsHostsCategoryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v2 fields hosts category request timeout response has a 4xx status code
func (o *GetV2FieldsHostsCategoryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v2 fields hosts category request timeout response has a 5xx status code
func (o *GetV2FieldsHostsCategoryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get v2 fields hosts category request timeout response a status code equal to that given
func (o *GetV2FieldsHostsCategoryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get v2 fields hosts category request timeout response
func (o *GetV2FieldsHostsCategoryRequestTimeout) Code() int {
	return 408
}

func (o *GetV2FieldsHostsCategoryRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/fields/hosts/{category}/][%d] getV2FieldsHostsCategoryRequestTimeout %s", 408, payload)
}

func (o *GetV2FieldsHostsCategoryRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/fields/hosts/{category}/][%d] getV2FieldsHostsCategoryRequestTimeout %s", 408, payload)
}

func (o *GetV2FieldsHostsCategoryRequestTimeout) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV2FieldsHostsCategoryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2FieldsHostsCategoryInternalServerError creates a GetV2FieldsHostsCategoryInternalServerError with default headers values
func NewGetV2FieldsHostsCategoryInternalServerError() *GetV2FieldsHostsCategoryInternalServerError {
	return &GetV2FieldsHostsCategoryInternalServerError{}
}

/*
GetV2FieldsHostsCategoryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV2FieldsHostsCategoryInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v2 fields hosts category internal server error response has a 2xx status code
func (o *GetV2FieldsHostsCategoryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v2 fields hosts category internal server error response has a 3xx status code
func (o *GetV2FieldsHostsCategoryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v2 fields hosts category internal server error response has a 4xx status code
func (o *GetV2FieldsHostsCategoryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v2 fields hosts category internal server error response has a 5xx status code
func (o *GetV2FieldsHostsCategoryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v2 fields hosts category internal server error response a status code equal to that given
func (o *GetV2FieldsHostsCategoryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v2 fields hosts category internal server error response
func (o *GetV2FieldsHostsCategoryInternalServerError) Code() int {
	return 500
}

func (o *GetV2FieldsHostsCategoryInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/fields/hosts/{category}/][%d] getV2FieldsHostsCategoryInternalServerError %s", 500, payload)
}

func (o *GetV2FieldsHostsCategoryInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v2/fields/hosts/{category}/][%d] getV2FieldsHostsCategoryInternalServerError %s", 500, payload)
}

func (o *GetV2FieldsHostsCategoryInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV2FieldsHostsCategoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetV2FieldsHostsCategoryOKBody get v2 fields hosts category o k body
swagger:model GetV2FieldsHostsCategoryOKBody
*/
type GetV2FieldsHostsCategoryOKBody struct {
	models.APIResponse

	// data
	Data []*models.Field `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetV2FieldsHostsCategoryOKBody) UnmarshalJSON(raw []byte) error {
	// GetV2FieldsHostsCategoryOKBodyAO0
	var getV2FieldsHostsCategoryOKBodyAO0 models.APIResponse
	if err := swag.ReadJSON(raw, &getV2FieldsHostsCategoryOKBodyAO0); err != nil {
		return err
	}
	o.APIResponse = getV2FieldsHostsCategoryOKBodyAO0

	// GetV2FieldsHostsCategoryOKBodyAO1
	var dataGetV2FieldsHostsCategoryOKBodyAO1 struct {
		Data []*models.Field `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataGetV2FieldsHostsCategoryOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetV2FieldsHostsCategoryOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetV2FieldsHostsCategoryOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getV2FieldsHostsCategoryOKBodyAO0, err := swag.WriteJSON(o.APIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getV2FieldsHostsCategoryOKBodyAO0)
	var dataGetV2FieldsHostsCategoryOKBodyAO1 struct {
		Data []*models.Field `json:"data"`
	}

	dataGetV2FieldsHostsCategoryOKBodyAO1.Data = o.Data

	jsonDataGetV2FieldsHostsCategoryOKBodyAO1, errGetV2FieldsHostsCategoryOKBodyAO1 := swag.WriteJSON(dataGetV2FieldsHostsCategoryOKBodyAO1)
	if errGetV2FieldsHostsCategoryOKBodyAO1 != nil {
		return nil, errGetV2FieldsHostsCategoryOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetV2FieldsHostsCategoryOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get v2 fields hosts category o k body
func (o *GetV2FieldsHostsCategoryOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetV2FieldsHostsCategoryOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("getV2FieldsHostsCategoryOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV2FieldsHostsCategoryOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get v2 fields hosts category o k body based on the context it is used
func (o *GetV2FieldsHostsCategoryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetV2FieldsHostsCategoryOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV2FieldsHostsCategoryOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV2FieldsHostsCategoryOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV2FieldsHostsCategoryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV2FieldsHostsCategoryOKBody) UnmarshalBinary(b []byte) error {
	var res GetV2FieldsHostsCategoryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
