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

// PostV1HostsCountReader is a Reader for the PostV1HostsCount structure.
type PostV1HostsCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1HostsCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1HostsCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1HostsCountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewPostV1HostsCountPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostV1HostsCountRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostV1HostsCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/hosts/count] PostV1HostsCount", response, response.Code())
	}
}

// NewPostV1HostsCountOK creates a PostV1HostsCountOK with default headers values
func NewPostV1HostsCountOK() *PostV1HostsCountOK {
	return &PostV1HostsCountOK{}
}

/*
PostV1HostsCountOK describes a response with status code 200, with default header values.

OK
*/
type PostV1HostsCountOK struct {
	Payload *PostV1HostsCountOKBody
}

// IsSuccess returns true when this post v1 hosts count o k response has a 2xx status code
func (o *PostV1HostsCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 hosts count o k response has a 3xx status code
func (o *PostV1HostsCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 hosts count o k response has a 4xx status code
func (o *PostV1HostsCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 hosts count o k response has a 5xx status code
func (o *PostV1HostsCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 hosts count o k response a status code equal to that given
func (o *PostV1HostsCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 hosts count o k response
func (o *PostV1HostsCountOK) Code() int {
	return 200
}

func (o *PostV1HostsCountOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/hosts/count][%d] postV1HostsCountOK %s", 200, payload)
}

func (o *PostV1HostsCountOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/hosts/count][%d] postV1HostsCountOK %s", 200, payload)
}

func (o *PostV1HostsCountOK) GetPayload() *PostV1HostsCountOKBody {
	return o.Payload
}

func (o *PostV1HostsCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostV1HostsCountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1HostsCountBadRequest creates a PostV1HostsCountBadRequest with default headers values
func NewPostV1HostsCountBadRequest() *PostV1HostsCountBadRequest {
	return &PostV1HostsCountBadRequest{}
}

/*
PostV1HostsCountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1HostsCountBadRequest struct {
	Payload *PostV1HostsCountBadRequestBody
}

// IsSuccess returns true when this post v1 hosts count bad request response has a 2xx status code
func (o *PostV1HostsCountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 hosts count bad request response has a 3xx status code
func (o *PostV1HostsCountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 hosts count bad request response has a 4xx status code
func (o *PostV1HostsCountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 hosts count bad request response has a 5xx status code
func (o *PostV1HostsCountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 hosts count bad request response a status code equal to that given
func (o *PostV1HostsCountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post v1 hosts count bad request response
func (o *PostV1HostsCountBadRequest) Code() int {
	return 400
}

func (o *PostV1HostsCountBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/hosts/count][%d] postV1HostsCountBadRequest %s", 400, payload)
}

func (o *PostV1HostsCountBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/hosts/count][%d] postV1HostsCountBadRequest %s", 400, payload)
}

func (o *PostV1HostsCountBadRequest) GetPayload() *PostV1HostsCountBadRequestBody {
	return o.Payload
}

func (o *PostV1HostsCountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostV1HostsCountBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1HostsCountPaymentRequired creates a PostV1HostsCountPaymentRequired with default headers values
func NewPostV1HostsCountPaymentRequired() *PostV1HostsCountPaymentRequired {
	return &PostV1HostsCountPaymentRequired{}
}

/*
PostV1HostsCountPaymentRequired describes a response with status code 402, with default header values.

Payment Required
*/
type PostV1HostsCountPaymentRequired struct {
	Payload *PostV1HostsCountPaymentRequiredBody
}

// IsSuccess returns true when this post v1 hosts count payment required response has a 2xx status code
func (o *PostV1HostsCountPaymentRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 hosts count payment required response has a 3xx status code
func (o *PostV1HostsCountPaymentRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 hosts count payment required response has a 4xx status code
func (o *PostV1HostsCountPaymentRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 hosts count payment required response has a 5xx status code
func (o *PostV1HostsCountPaymentRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 hosts count payment required response a status code equal to that given
func (o *PostV1HostsCountPaymentRequired) IsCode(code int) bool {
	return code == 402
}

// Code gets the status code for the post v1 hosts count payment required response
func (o *PostV1HostsCountPaymentRequired) Code() int {
	return 402
}

func (o *PostV1HostsCountPaymentRequired) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/hosts/count][%d] postV1HostsCountPaymentRequired %s", 402, payload)
}

func (o *PostV1HostsCountPaymentRequired) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/hosts/count][%d] postV1HostsCountPaymentRequired %s", 402, payload)
}

func (o *PostV1HostsCountPaymentRequired) GetPayload() *PostV1HostsCountPaymentRequiredBody {
	return o.Payload
}

func (o *PostV1HostsCountPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostV1HostsCountPaymentRequiredBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1HostsCountRequestTimeout creates a PostV1HostsCountRequestTimeout with default headers values
func NewPostV1HostsCountRequestTimeout() *PostV1HostsCountRequestTimeout {
	return &PostV1HostsCountRequestTimeout{}
}

/*
PostV1HostsCountRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PostV1HostsCountRequestTimeout struct {
	Payload *PostV1HostsCountRequestTimeoutBody
}

// IsSuccess returns true when this post v1 hosts count request timeout response has a 2xx status code
func (o *PostV1HostsCountRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 hosts count request timeout response has a 3xx status code
func (o *PostV1HostsCountRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 hosts count request timeout response has a 4xx status code
func (o *PostV1HostsCountRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 hosts count request timeout response has a 5xx status code
func (o *PostV1HostsCountRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 hosts count request timeout response a status code equal to that given
func (o *PostV1HostsCountRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the post v1 hosts count request timeout response
func (o *PostV1HostsCountRequestTimeout) Code() int {
	return 408
}

func (o *PostV1HostsCountRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/hosts/count][%d] postV1HostsCountRequestTimeout %s", 408, payload)
}

func (o *PostV1HostsCountRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/hosts/count][%d] postV1HostsCountRequestTimeout %s", 408, payload)
}

func (o *PostV1HostsCountRequestTimeout) GetPayload() *PostV1HostsCountRequestTimeoutBody {
	return o.Payload
}

func (o *PostV1HostsCountRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostV1HostsCountRequestTimeoutBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1HostsCountInternalServerError creates a PostV1HostsCountInternalServerError with default headers values
func NewPostV1HostsCountInternalServerError() *PostV1HostsCountInternalServerError {
	return &PostV1HostsCountInternalServerError{}
}

/*
PostV1HostsCountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostV1HostsCountInternalServerError struct {
	Payload *PostV1HostsCountInternalServerErrorBody
}

// IsSuccess returns true when this post v1 hosts count internal server error response has a 2xx status code
func (o *PostV1HostsCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 hosts count internal server error response has a 3xx status code
func (o *PostV1HostsCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 hosts count internal server error response has a 4xx status code
func (o *PostV1HostsCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 hosts count internal server error response has a 5xx status code
func (o *PostV1HostsCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post v1 hosts count internal server error response a status code equal to that given
func (o *PostV1HostsCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post v1 hosts count internal server error response
func (o *PostV1HostsCountInternalServerError) Code() int {
	return 500
}

func (o *PostV1HostsCountInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/hosts/count][%d] postV1HostsCountInternalServerError %s", 500, payload)
}

func (o *PostV1HostsCountInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/hosts/count][%d] postV1HostsCountInternalServerError %s", 500, payload)
}

func (o *PostV1HostsCountInternalServerError) GetPayload() *PostV1HostsCountInternalServerErrorBody {
	return o.Payload
}

func (o *PostV1HostsCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostV1HostsCountInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostV1HostsCountBadRequestBody post v1 hosts count bad request body
swagger:model PostV1HostsCountBadRequestBody
*/
type PostV1HostsCountBadRequestBody struct {
	models.ErrorResponse

	// error
	Error string `json:" error,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostV1HostsCountBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PostV1HostsCountBadRequestBodyAO0
	var postV1HostsCountBadRequestBodyAO0 models.ErrorResponse
	if err := swag.ReadJSON(raw, &postV1HostsCountBadRequestBodyAO0); err != nil {
		return err
	}
	o.ErrorResponse = postV1HostsCountBadRequestBodyAO0

	// PostV1HostsCountBadRequestBodyAO1
	var dataPostV1HostsCountBadRequestBodyAO1 struct {
		Error string `json:" error,omitempty"`

		Success bool `json:"success,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostV1HostsCountBadRequestBodyAO1); err != nil {
		return err
	}

	o.Error = dataPostV1HostsCountBadRequestBodyAO1.Error

	o.Success = dataPostV1HostsCountBadRequestBodyAO1.Success

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostV1HostsCountBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postV1HostsCountBadRequestBodyAO0, err := swag.WriteJSON(o.ErrorResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postV1HostsCountBadRequestBodyAO0)
	var dataPostV1HostsCountBadRequestBodyAO1 struct {
		Error string `json:" error,omitempty"`

		Success bool `json:"success,omitempty"`
	}

	dataPostV1HostsCountBadRequestBodyAO1.Error = o.Error

	dataPostV1HostsCountBadRequestBodyAO1.Success = o.Success

	jsonDataPostV1HostsCountBadRequestBodyAO1, errPostV1HostsCountBadRequestBodyAO1 := swag.WriteJSON(dataPostV1HostsCountBadRequestBodyAO1)
	if errPostV1HostsCountBadRequestBodyAO1 != nil {
		return nil, errPostV1HostsCountBadRequestBodyAO1
	}
	_parts = append(_parts, jsonDataPostV1HostsCountBadRequestBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post v1 hosts count bad request body
func (o *PostV1HostsCountBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ErrorResponse
	if err := o.ErrorResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this post v1 hosts count bad request body based on the context it is used
func (o *PostV1HostsCountBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ErrorResponse
	if err := o.ErrorResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostV1HostsCountBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostV1HostsCountBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostV1HostsCountBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostV1HostsCountInternalServerErrorBody post v1 hosts count internal server error body
swagger:model PostV1HostsCountInternalServerErrorBody
*/
type PostV1HostsCountInternalServerErrorBody struct {
	models.ErrorResponse

	// error
	Error string `json:" error,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostV1HostsCountInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// PostV1HostsCountInternalServerErrorBodyAO0
	var postV1HostsCountInternalServerErrorBodyAO0 models.ErrorResponse
	if err := swag.ReadJSON(raw, &postV1HostsCountInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.ErrorResponse = postV1HostsCountInternalServerErrorBodyAO0

	// PostV1HostsCountInternalServerErrorBodyAO1
	var dataPostV1HostsCountInternalServerErrorBodyAO1 struct {
		Error string `json:" error,omitempty"`

		Success bool `json:"success,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostV1HostsCountInternalServerErrorBodyAO1); err != nil {
		return err
	}

	o.Error = dataPostV1HostsCountInternalServerErrorBodyAO1.Error

	o.Success = dataPostV1HostsCountInternalServerErrorBodyAO1.Success

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostV1HostsCountInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postV1HostsCountInternalServerErrorBodyAO0, err := swag.WriteJSON(o.ErrorResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postV1HostsCountInternalServerErrorBodyAO0)
	var dataPostV1HostsCountInternalServerErrorBodyAO1 struct {
		Error string `json:" error,omitempty"`

		Success bool `json:"success,omitempty"`
	}

	dataPostV1HostsCountInternalServerErrorBodyAO1.Error = o.Error

	dataPostV1HostsCountInternalServerErrorBodyAO1.Success = o.Success

	jsonDataPostV1HostsCountInternalServerErrorBodyAO1, errPostV1HostsCountInternalServerErrorBodyAO1 := swag.WriteJSON(dataPostV1HostsCountInternalServerErrorBodyAO1)
	if errPostV1HostsCountInternalServerErrorBodyAO1 != nil {
		return nil, errPostV1HostsCountInternalServerErrorBodyAO1
	}
	_parts = append(_parts, jsonDataPostV1HostsCountInternalServerErrorBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post v1 hosts count internal server error body
func (o *PostV1HostsCountInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ErrorResponse
	if err := o.ErrorResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this post v1 hosts count internal server error body based on the context it is used
func (o *PostV1HostsCountInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ErrorResponse
	if err := o.ErrorResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostV1HostsCountInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostV1HostsCountInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostV1HostsCountInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostV1HostsCountOKBody post v1 hosts count o k body
swagger:model PostV1HostsCountOKBody
*/
type PostV1HostsCountOKBody struct {
	models.APIResponse

	// meta
	Meta *models.PaginationStruct `json:" meta,omitempty"`

	// data
	Data *models.CertCount `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostV1HostsCountOKBody) UnmarshalJSON(raw []byte) error {
	// PostV1HostsCountOKBodyAO0
	var postV1HostsCountOKBodyAO0 models.APIResponse
	if err := swag.ReadJSON(raw, &postV1HostsCountOKBodyAO0); err != nil {
		return err
	}
	o.APIResponse = postV1HostsCountOKBodyAO0

	// PostV1HostsCountOKBodyAO1
	var dataPostV1HostsCountOKBodyAO1 struct {
		Meta *models.PaginationStruct `json:" meta,omitempty"`

		Data *models.CertCount `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostV1HostsCountOKBodyAO1); err != nil {
		return err
	}

	o.Meta = dataPostV1HostsCountOKBodyAO1.Meta

	o.Data = dataPostV1HostsCountOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostV1HostsCountOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postV1HostsCountOKBodyAO0, err := swag.WriteJSON(o.APIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postV1HostsCountOKBodyAO0)
	var dataPostV1HostsCountOKBodyAO1 struct {
		Meta *models.PaginationStruct `json:" meta,omitempty"`

		Data *models.CertCount `json:"data,omitempty"`
	}

	dataPostV1HostsCountOKBodyAO1.Meta = o.Meta

	dataPostV1HostsCountOKBodyAO1.Data = o.Data

	jsonDataPostV1HostsCountOKBodyAO1, errPostV1HostsCountOKBodyAO1 := swag.WriteJSON(dataPostV1HostsCountOKBodyAO1)
	if errPostV1HostsCountOKBodyAO1 != nil {
		return nil, errPostV1HostsCountOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostV1HostsCountOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post v1 hosts count o k body
func (o *PostV1HostsCountOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.APIResponse
	if err := o.APIResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMeta(formats); err != nil {
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

func (o *PostV1HostsCountOKBody) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(o.Meta) { // not required
		return nil
	}

	if o.Meta != nil {
		if err := o.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postV1HostsCountOK" + "." + " meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postV1HostsCountOK" + "." + " meta")
			}
			return err
		}
	}

	return nil
}

func (o *PostV1HostsCountOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postV1HostsCountOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postV1HostsCountOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post v1 hosts count o k body based on the context it is used
func (o *PostV1HostsCountOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.APIResponse
	if err := o.APIResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMeta(ctx, formats); err != nil {
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

func (o *PostV1HostsCountOKBody) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if o.Meta != nil {

		if swag.IsZero(o.Meta) { // not required
			return nil
		}

		if err := o.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postV1HostsCountOK" + "." + " meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postV1HostsCountOK" + "." + " meta")
			}
			return err
		}
	}

	return nil
}

func (o *PostV1HostsCountOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postV1HostsCountOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postV1HostsCountOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostV1HostsCountOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostV1HostsCountOKBody) UnmarshalBinary(b []byte) error {
	var res PostV1HostsCountOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostV1HostsCountPaymentRequiredBody post v1 hosts count payment required body
swagger:model PostV1HostsCountPaymentRequiredBody
*/
type PostV1HostsCountPaymentRequiredBody struct {
	models.ErrorResponse

	// error
	Error string `json:" error,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostV1HostsCountPaymentRequiredBody) UnmarshalJSON(raw []byte) error {
	// PostV1HostsCountPaymentRequiredBodyAO0
	var postV1HostsCountPaymentRequiredBodyAO0 models.ErrorResponse
	if err := swag.ReadJSON(raw, &postV1HostsCountPaymentRequiredBodyAO0); err != nil {
		return err
	}
	o.ErrorResponse = postV1HostsCountPaymentRequiredBodyAO0

	// PostV1HostsCountPaymentRequiredBodyAO1
	var dataPostV1HostsCountPaymentRequiredBodyAO1 struct {
		Error string `json:" error,omitempty"`

		Success bool `json:"success,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostV1HostsCountPaymentRequiredBodyAO1); err != nil {
		return err
	}

	o.Error = dataPostV1HostsCountPaymentRequiredBodyAO1.Error

	o.Success = dataPostV1HostsCountPaymentRequiredBodyAO1.Success

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostV1HostsCountPaymentRequiredBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postV1HostsCountPaymentRequiredBodyAO0, err := swag.WriteJSON(o.ErrorResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postV1HostsCountPaymentRequiredBodyAO0)
	var dataPostV1HostsCountPaymentRequiredBodyAO1 struct {
		Error string `json:" error,omitempty"`

		Success bool `json:"success,omitempty"`
	}

	dataPostV1HostsCountPaymentRequiredBodyAO1.Error = o.Error

	dataPostV1HostsCountPaymentRequiredBodyAO1.Success = o.Success

	jsonDataPostV1HostsCountPaymentRequiredBodyAO1, errPostV1HostsCountPaymentRequiredBodyAO1 := swag.WriteJSON(dataPostV1HostsCountPaymentRequiredBodyAO1)
	if errPostV1HostsCountPaymentRequiredBodyAO1 != nil {
		return nil, errPostV1HostsCountPaymentRequiredBodyAO1
	}
	_parts = append(_parts, jsonDataPostV1HostsCountPaymentRequiredBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post v1 hosts count payment required body
func (o *PostV1HostsCountPaymentRequiredBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ErrorResponse
	if err := o.ErrorResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this post v1 hosts count payment required body based on the context it is used
func (o *PostV1HostsCountPaymentRequiredBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ErrorResponse
	if err := o.ErrorResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostV1HostsCountPaymentRequiredBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostV1HostsCountPaymentRequiredBody) UnmarshalBinary(b []byte) error {
	var res PostV1HostsCountPaymentRequiredBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostV1HostsCountRequestTimeoutBody post v1 hosts count request timeout body
swagger:model PostV1HostsCountRequestTimeoutBody
*/
type PostV1HostsCountRequestTimeoutBody struct {
	models.ErrorResponse

	// error
	Error string `json:" error,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostV1HostsCountRequestTimeoutBody) UnmarshalJSON(raw []byte) error {
	// PostV1HostsCountRequestTimeoutBodyAO0
	var postV1HostsCountRequestTimeoutBodyAO0 models.ErrorResponse
	if err := swag.ReadJSON(raw, &postV1HostsCountRequestTimeoutBodyAO0); err != nil {
		return err
	}
	o.ErrorResponse = postV1HostsCountRequestTimeoutBodyAO0

	// PostV1HostsCountRequestTimeoutBodyAO1
	var dataPostV1HostsCountRequestTimeoutBodyAO1 struct {
		Error string `json:" error,omitempty"`

		Success bool `json:"success,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostV1HostsCountRequestTimeoutBodyAO1); err != nil {
		return err
	}

	o.Error = dataPostV1HostsCountRequestTimeoutBodyAO1.Error

	o.Success = dataPostV1HostsCountRequestTimeoutBodyAO1.Success

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostV1HostsCountRequestTimeoutBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postV1HostsCountRequestTimeoutBodyAO0, err := swag.WriteJSON(o.ErrorResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postV1HostsCountRequestTimeoutBodyAO0)
	var dataPostV1HostsCountRequestTimeoutBodyAO1 struct {
		Error string `json:" error,omitempty"`

		Success bool `json:"success,omitempty"`
	}

	dataPostV1HostsCountRequestTimeoutBodyAO1.Error = o.Error

	dataPostV1HostsCountRequestTimeoutBodyAO1.Success = o.Success

	jsonDataPostV1HostsCountRequestTimeoutBodyAO1, errPostV1HostsCountRequestTimeoutBodyAO1 := swag.WriteJSON(dataPostV1HostsCountRequestTimeoutBodyAO1)
	if errPostV1HostsCountRequestTimeoutBodyAO1 != nil {
		return nil, errPostV1HostsCountRequestTimeoutBodyAO1
	}
	_parts = append(_parts, jsonDataPostV1HostsCountRequestTimeoutBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post v1 hosts count request timeout body
func (o *PostV1HostsCountRequestTimeoutBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ErrorResponse
	if err := o.ErrorResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this post v1 hosts count request timeout body based on the context it is used
func (o *PostV1HostsCountRequestTimeoutBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ErrorResponse
	if err := o.ErrorResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostV1HostsCountRequestTimeoutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostV1HostsCountRequestTimeoutBody) UnmarshalBinary(b []byte) error {
	var res PostV1HostsCountRequestTimeoutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
