// Code generated by go-swagger; DO NOT EDIT.

package domain

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

// PostV1DomainCountReader is a Reader for the PostV1DomainCount structure.
type PostV1DomainCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1DomainCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1DomainCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1DomainCountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostV1DomainCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/domain/count] PostV1DomainCount", response, response.Code())
	}
}

// NewPostV1DomainCountOK creates a PostV1DomainCountOK with default headers values
func NewPostV1DomainCountOK() *PostV1DomainCountOK {
	return &PostV1DomainCountOK{}
}

/*
PostV1DomainCountOK describes a response with status code 200, with default header values.

OK
*/
type PostV1DomainCountOK struct {
	Payload *PostV1DomainCountOKBody
}

// IsSuccess returns true when this post v1 domain count o k response has a 2xx status code
func (o *PostV1DomainCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 domain count o k response has a 3xx status code
func (o *PostV1DomainCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 domain count o k response has a 4xx status code
func (o *PostV1DomainCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 domain count o k response has a 5xx status code
func (o *PostV1DomainCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 domain count o k response a status code equal to that given
func (o *PostV1DomainCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 domain count o k response
func (o *PostV1DomainCountOK) Code() int {
	return 200
}

func (o *PostV1DomainCountOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/domain/count][%d] postV1DomainCountOK %s", 200, payload)
}

func (o *PostV1DomainCountOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/domain/count][%d] postV1DomainCountOK %s", 200, payload)
}

func (o *PostV1DomainCountOK) GetPayload() *PostV1DomainCountOKBody {
	return o.Payload
}

func (o *PostV1DomainCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostV1DomainCountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1DomainCountBadRequest creates a PostV1DomainCountBadRequest with default headers values
func NewPostV1DomainCountBadRequest() *PostV1DomainCountBadRequest {
	return &PostV1DomainCountBadRequest{}
}

/*
PostV1DomainCountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1DomainCountBadRequest struct {
	Payload *models.DNSErrorResponse
}

// IsSuccess returns true when this post v1 domain count bad request response has a 2xx status code
func (o *PostV1DomainCountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 domain count bad request response has a 3xx status code
func (o *PostV1DomainCountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 domain count bad request response has a 4xx status code
func (o *PostV1DomainCountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 domain count bad request response has a 5xx status code
func (o *PostV1DomainCountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 domain count bad request response a status code equal to that given
func (o *PostV1DomainCountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post v1 domain count bad request response
func (o *PostV1DomainCountBadRequest) Code() int {
	return 400
}

func (o *PostV1DomainCountBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/domain/count][%d] postV1DomainCountBadRequest %s", 400, payload)
}

func (o *PostV1DomainCountBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/domain/count][%d] postV1DomainCountBadRequest %s", 400, payload)
}

func (o *PostV1DomainCountBadRequest) GetPayload() *models.DNSErrorResponse {
	return o.Payload
}

func (o *PostV1DomainCountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DNSErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1DomainCountInternalServerError creates a PostV1DomainCountInternalServerError with default headers values
func NewPostV1DomainCountInternalServerError() *PostV1DomainCountInternalServerError {
	return &PostV1DomainCountInternalServerError{}
}

/*
PostV1DomainCountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostV1DomainCountInternalServerError struct {
	Payload *models.DNSErrorResponse
}

// IsSuccess returns true when this post v1 domain count internal server error response has a 2xx status code
func (o *PostV1DomainCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 domain count internal server error response has a 3xx status code
func (o *PostV1DomainCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 domain count internal server error response has a 4xx status code
func (o *PostV1DomainCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 domain count internal server error response has a 5xx status code
func (o *PostV1DomainCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post v1 domain count internal server error response a status code equal to that given
func (o *PostV1DomainCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post v1 domain count internal server error response
func (o *PostV1DomainCountInternalServerError) Code() int {
	return 500
}

func (o *PostV1DomainCountInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/domain/count][%d] postV1DomainCountInternalServerError %s", 500, payload)
}

func (o *PostV1DomainCountInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/domain/count][%d] postV1DomainCountInternalServerError %s", 500, payload)
}

func (o *PostV1DomainCountInternalServerError) GetPayload() *models.DNSErrorResponse {
	return o.Payload
}

func (o *PostV1DomainCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DNSErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostV1DomainCountOKBody post v1 domain count o k body
swagger:model PostV1DomainCountOKBody
*/
type PostV1DomainCountOKBody struct {
	models.DNSAPIResponse

	// data
	Data *models.DNSData `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostV1DomainCountOKBody) UnmarshalJSON(raw []byte) error {
	// PostV1DomainCountOKBodyAO0
	var postV1DomainCountOKBodyAO0 models.DNSAPIResponse
	if err := swag.ReadJSON(raw, &postV1DomainCountOKBodyAO0); err != nil {
		return err
	}
	o.DNSAPIResponse = postV1DomainCountOKBodyAO0

	// PostV1DomainCountOKBodyAO1
	var dataPostV1DomainCountOKBodyAO1 struct {
		Data *models.DNSData `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostV1DomainCountOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostV1DomainCountOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostV1DomainCountOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postV1DomainCountOKBodyAO0, err := swag.WriteJSON(o.DNSAPIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postV1DomainCountOKBodyAO0)
	var dataPostV1DomainCountOKBodyAO1 struct {
		Data *models.DNSData `json:"data,omitempty"`
	}

	dataPostV1DomainCountOKBodyAO1.Data = o.Data

	jsonDataPostV1DomainCountOKBodyAO1, errPostV1DomainCountOKBodyAO1 := swag.WriteJSON(dataPostV1DomainCountOKBodyAO1)
	if errPostV1DomainCountOKBodyAO1 != nil {
		return nil, errPostV1DomainCountOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostV1DomainCountOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post v1 domain count o k body
func (o *PostV1DomainCountOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.DNSAPIResponse
	if err := o.DNSAPIResponse.Validate(formats); err != nil {
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

func (o *PostV1DomainCountOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postV1DomainCountOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postV1DomainCountOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post v1 domain count o k body based on the context it is used
func (o *PostV1DomainCountOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.DNSAPIResponse
	if err := o.DNSAPIResponse.ContextValidate(ctx, formats); err != nil {
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

func (o *PostV1DomainCountOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postV1DomainCountOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postV1DomainCountOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostV1DomainCountOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostV1DomainCountOKBody) UnmarshalBinary(b []byte) error {
	var res PostV1DomainCountOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
