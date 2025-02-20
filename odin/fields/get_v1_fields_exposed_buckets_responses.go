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

// GetV1FieldsExposedBucketsReader is a Reader for the GetV1FieldsExposedBuckets structure.
type GetV1FieldsExposedBucketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1FieldsExposedBucketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1FieldsExposedBucketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1FieldsExposedBucketsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewGetV1FieldsExposedBucketsPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetV1FieldsExposedBucketsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1FieldsExposedBucketsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/fields/exposed/buckets/] GetV1FieldsExposedBuckets", response, response.Code())
	}
}

// NewGetV1FieldsExposedBucketsOK creates a GetV1FieldsExposedBucketsOK with default headers values
func NewGetV1FieldsExposedBucketsOK() *GetV1FieldsExposedBucketsOK {
	return &GetV1FieldsExposedBucketsOK{}
}

/*
GetV1FieldsExposedBucketsOK describes a response with status code 200, with default header values.

OK
*/
type GetV1FieldsExposedBucketsOK struct {
	Payload *GetV1FieldsExposedBucketsOKBody
}

// IsSuccess returns true when this get v1 fields exposed buckets o k response has a 2xx status code
func (o *GetV1FieldsExposedBucketsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 fields exposed buckets o k response has a 3xx status code
func (o *GetV1FieldsExposedBucketsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 fields exposed buckets o k response has a 4xx status code
func (o *GetV1FieldsExposedBucketsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 fields exposed buckets o k response has a 5xx status code
func (o *GetV1FieldsExposedBucketsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 fields exposed buckets o k response a status code equal to that given
func (o *GetV1FieldsExposedBucketsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get v1 fields exposed buckets o k response
func (o *GetV1FieldsExposedBucketsOK) Code() int {
	return 200
}

func (o *GetV1FieldsExposedBucketsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/exposed/buckets/][%d] getV1FieldsExposedBucketsOK %s", 200, payload)
}

func (o *GetV1FieldsExposedBucketsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/exposed/buckets/][%d] getV1FieldsExposedBucketsOK %s", 200, payload)
}

func (o *GetV1FieldsExposedBucketsOK) GetPayload() *GetV1FieldsExposedBucketsOKBody {
	return o.Payload
}

func (o *GetV1FieldsExposedBucketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetV1FieldsExposedBucketsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1FieldsExposedBucketsBadRequest creates a GetV1FieldsExposedBucketsBadRequest with default headers values
func NewGetV1FieldsExposedBucketsBadRequest() *GetV1FieldsExposedBucketsBadRequest {
	return &GetV1FieldsExposedBucketsBadRequest{}
}

/*
GetV1FieldsExposedBucketsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1FieldsExposedBucketsBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 fields exposed buckets bad request response has a 2xx status code
func (o *GetV1FieldsExposedBucketsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 fields exposed buckets bad request response has a 3xx status code
func (o *GetV1FieldsExposedBucketsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 fields exposed buckets bad request response has a 4xx status code
func (o *GetV1FieldsExposedBucketsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 fields exposed buckets bad request response has a 5xx status code
func (o *GetV1FieldsExposedBucketsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 fields exposed buckets bad request response a status code equal to that given
func (o *GetV1FieldsExposedBucketsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 fields exposed buckets bad request response
func (o *GetV1FieldsExposedBucketsBadRequest) Code() int {
	return 400
}

func (o *GetV1FieldsExposedBucketsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/exposed/buckets/][%d] getV1FieldsExposedBucketsBadRequest %s", 400, payload)
}

func (o *GetV1FieldsExposedBucketsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/exposed/buckets/][%d] getV1FieldsExposedBucketsBadRequest %s", 400, payload)
}

func (o *GetV1FieldsExposedBucketsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1FieldsExposedBucketsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1FieldsExposedBucketsPaymentRequired creates a GetV1FieldsExposedBucketsPaymentRequired with default headers values
func NewGetV1FieldsExposedBucketsPaymentRequired() *GetV1FieldsExposedBucketsPaymentRequired {
	return &GetV1FieldsExposedBucketsPaymentRequired{}
}

/*
GetV1FieldsExposedBucketsPaymentRequired describes a response with status code 402, with default header values.

Payment Required
*/
type GetV1FieldsExposedBucketsPaymentRequired struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 fields exposed buckets payment required response has a 2xx status code
func (o *GetV1FieldsExposedBucketsPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 fields exposed buckets payment required response has a 3xx status code
func (o *GetV1FieldsExposedBucketsPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 fields exposed buckets payment required response has a 4xx status code
func (o *GetV1FieldsExposedBucketsPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 fields exposed buckets payment required response has a 5xx status code
func (o *GetV1FieldsExposedBucketsPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 fields exposed buckets payment required response a status code equal to that given
func (o *GetV1FieldsExposedBucketsPaymentRequired) IsCode(code int) bool {
	return code == 402
}

// Code gets the status code for the get v1 fields exposed buckets payment required response
func (o *GetV1FieldsExposedBucketsPaymentRequired) Code() int {
	return 402
}

func (o *GetV1FieldsExposedBucketsPaymentRequired) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/exposed/buckets/][%d] getV1FieldsExposedBucketsPaymentRequired %s", 402, payload)
}

func (o *GetV1FieldsExposedBucketsPaymentRequired) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/exposed/buckets/][%d] getV1FieldsExposedBucketsPaymentRequired %s", 402, payload)
}

func (o *GetV1FieldsExposedBucketsPaymentRequired) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1FieldsExposedBucketsPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1FieldsExposedBucketsRequestTimeout creates a GetV1FieldsExposedBucketsRequestTimeout with default headers values
func NewGetV1FieldsExposedBucketsRequestTimeout() *GetV1FieldsExposedBucketsRequestTimeout {
	return &GetV1FieldsExposedBucketsRequestTimeout{}
}

/*
GetV1FieldsExposedBucketsRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type GetV1FieldsExposedBucketsRequestTimeout struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 fields exposed buckets request timeout response has a 2xx status code
func (o *GetV1FieldsExposedBucketsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 fields exposed buckets request timeout response has a 3xx status code
func (o *GetV1FieldsExposedBucketsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 fields exposed buckets request timeout response has a 4xx status code
func (o *GetV1FieldsExposedBucketsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 fields exposed buckets request timeout response has a 5xx status code
func (o *GetV1FieldsExposedBucketsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 fields exposed buckets request timeout response a status code equal to that given
func (o *GetV1FieldsExposedBucketsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the get v1 fields exposed buckets request timeout response
func (o *GetV1FieldsExposedBucketsRequestTimeout) Code() int {
	return 408
}

func (o *GetV1FieldsExposedBucketsRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/exposed/buckets/][%d] getV1FieldsExposedBucketsRequestTimeout %s", 408, payload)
}

func (o *GetV1FieldsExposedBucketsRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/exposed/buckets/][%d] getV1FieldsExposedBucketsRequestTimeout %s", 408, payload)
}

func (o *GetV1FieldsExposedBucketsRequestTimeout) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1FieldsExposedBucketsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1FieldsExposedBucketsInternalServerError creates a GetV1FieldsExposedBucketsInternalServerError with default headers values
func NewGetV1FieldsExposedBucketsInternalServerError() *GetV1FieldsExposedBucketsInternalServerError {
	return &GetV1FieldsExposedBucketsInternalServerError{}
}

/*
GetV1FieldsExposedBucketsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1FieldsExposedBucketsInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get v1 fields exposed buckets internal server error response has a 2xx status code
func (o *GetV1FieldsExposedBucketsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 fields exposed buckets internal server error response has a 3xx status code
func (o *GetV1FieldsExposedBucketsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 fields exposed buckets internal server error response has a 4xx status code
func (o *GetV1FieldsExposedBucketsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 fields exposed buckets internal server error response has a 5xx status code
func (o *GetV1FieldsExposedBucketsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 fields exposed buckets internal server error response a status code equal to that given
func (o *GetV1FieldsExposedBucketsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 fields exposed buckets internal server error response
func (o *GetV1FieldsExposedBucketsInternalServerError) Code() int {
	return 500
}

func (o *GetV1FieldsExposedBucketsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/exposed/buckets/][%d] getV1FieldsExposedBucketsInternalServerError %s", 500, payload)
}

func (o *GetV1FieldsExposedBucketsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/fields/exposed/buckets/][%d] getV1FieldsExposedBucketsInternalServerError %s", 500, payload)
}

func (o *GetV1FieldsExposedBucketsInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetV1FieldsExposedBucketsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetV1FieldsExposedBucketsOKBody get v1 fields exposed buckets o k body
swagger:model GetV1FieldsExposedBucketsOKBody
*/
type GetV1FieldsExposedBucketsOKBody struct {
	models.APIResponse

	// data
	Data []*models.Field `json:"Data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetV1FieldsExposedBucketsOKBody) UnmarshalJSON(raw []byte) error {
	// GetV1FieldsExposedBucketsOKBodyAO0
	var getV1FieldsExposedBucketsOKBodyAO0 models.APIResponse
	if err := swag.ReadJSON(raw, &getV1FieldsExposedBucketsOKBodyAO0); err != nil {
		return err
	}
	o.APIResponse = getV1FieldsExposedBucketsOKBodyAO0

	// GetV1FieldsExposedBucketsOKBodyAO1
	var dataGetV1FieldsExposedBucketsOKBodyAO1 struct {
		Data []*models.Field `json:"Data"`
	}
	if err := swag.ReadJSON(raw, &dataGetV1FieldsExposedBucketsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetV1FieldsExposedBucketsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetV1FieldsExposedBucketsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getV1FieldsExposedBucketsOKBodyAO0, err := swag.WriteJSON(o.APIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getV1FieldsExposedBucketsOKBodyAO0)
	var dataGetV1FieldsExposedBucketsOKBodyAO1 struct {
		Data []*models.Field `json:"Data"`
	}

	dataGetV1FieldsExposedBucketsOKBodyAO1.Data = o.Data

	jsonDataGetV1FieldsExposedBucketsOKBodyAO1, errGetV1FieldsExposedBucketsOKBodyAO1 := swag.WriteJSON(dataGetV1FieldsExposedBucketsOKBodyAO1)
	if errGetV1FieldsExposedBucketsOKBodyAO1 != nil {
		return nil, errGetV1FieldsExposedBucketsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetV1FieldsExposedBucketsOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get v1 fields exposed buckets o k body
func (o *GetV1FieldsExposedBucketsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetV1FieldsExposedBucketsOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("getV1FieldsExposedBucketsOK" + "." + "Data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1FieldsExposedBucketsOK" + "." + "Data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get v1 fields exposed buckets o k body based on the context it is used
func (o *GetV1FieldsExposedBucketsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *GetV1FieldsExposedBucketsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getV1FieldsExposedBucketsOK" + "." + "Data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getV1FieldsExposedBucketsOK" + "." + "Data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetV1FieldsExposedBucketsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetV1FieldsExposedBucketsOKBody) UnmarshalBinary(b []byte) error {
	var res GetV1FieldsExposedBucketsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
