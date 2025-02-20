// Code generated by go-swagger; DO NOT EDIT.

package exposed_buckets

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

// PostV1ExposedBucketsSummaryReader is a Reader for the PostV1ExposedBucketsSummary structure.
type PostV1ExposedBucketsSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ExposedBucketsSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1ExposedBucketsSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1ExposedBucketsSummaryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostV1ExposedBucketsSummaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/exposed/buckets/summary] PostV1ExposedBucketsSummary", response, response.Code())
	}
}

// NewPostV1ExposedBucketsSummaryOK creates a PostV1ExposedBucketsSummaryOK with default headers values
func NewPostV1ExposedBucketsSummaryOK() *PostV1ExposedBucketsSummaryOK {
	return &PostV1ExposedBucketsSummaryOK{}
}

/*
PostV1ExposedBucketsSummaryOK describes a response with status code 200, with default header values.

OK
*/
type PostV1ExposedBucketsSummaryOK struct {
	Payload *PostV1ExposedBucketsSummaryOKBody
}

// IsSuccess returns true when this post v1 exposed buckets summary o k response has a 2xx status code
func (o *PostV1ExposedBucketsSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 exposed buckets summary o k response has a 3xx status code
func (o *PostV1ExposedBucketsSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 exposed buckets summary o k response has a 4xx status code
func (o *PostV1ExposedBucketsSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 exposed buckets summary o k response has a 5xx status code
func (o *PostV1ExposedBucketsSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 exposed buckets summary o k response a status code equal to that given
func (o *PostV1ExposedBucketsSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 exposed buckets summary o k response
func (o *PostV1ExposedBucketsSummaryOK) Code() int {
	return 200
}

func (o *PostV1ExposedBucketsSummaryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/buckets/summary][%d] postV1ExposedBucketsSummaryOK %s", 200, payload)
}

func (o *PostV1ExposedBucketsSummaryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/buckets/summary][%d] postV1ExposedBucketsSummaryOK %s", 200, payload)
}

func (o *PostV1ExposedBucketsSummaryOK) GetPayload() *PostV1ExposedBucketsSummaryOKBody {
	return o.Payload
}

func (o *PostV1ExposedBucketsSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostV1ExposedBucketsSummaryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ExposedBucketsSummaryBadRequest creates a PostV1ExposedBucketsSummaryBadRequest with default headers values
func NewPostV1ExposedBucketsSummaryBadRequest() *PostV1ExposedBucketsSummaryBadRequest {
	return &PostV1ExposedBucketsSummaryBadRequest{}
}

/*
PostV1ExposedBucketsSummaryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1ExposedBucketsSummaryBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this post v1 exposed buckets summary bad request response has a 2xx status code
func (o *PostV1ExposedBucketsSummaryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 exposed buckets summary bad request response has a 3xx status code
func (o *PostV1ExposedBucketsSummaryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 exposed buckets summary bad request response has a 4xx status code
func (o *PostV1ExposedBucketsSummaryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 exposed buckets summary bad request response has a 5xx status code
func (o *PostV1ExposedBucketsSummaryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 exposed buckets summary bad request response a status code equal to that given
func (o *PostV1ExposedBucketsSummaryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post v1 exposed buckets summary bad request response
func (o *PostV1ExposedBucketsSummaryBadRequest) Code() int {
	return 400
}

func (o *PostV1ExposedBucketsSummaryBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/buckets/summary][%d] postV1ExposedBucketsSummaryBadRequest %s", 400, payload)
}

func (o *PostV1ExposedBucketsSummaryBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/buckets/summary][%d] postV1ExposedBucketsSummaryBadRequest %s", 400, payload)
}

func (o *PostV1ExposedBucketsSummaryBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PostV1ExposedBucketsSummaryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ExposedBucketsSummaryInternalServerError creates a PostV1ExposedBucketsSummaryInternalServerError with default headers values
func NewPostV1ExposedBucketsSummaryInternalServerError() *PostV1ExposedBucketsSummaryInternalServerError {
	return &PostV1ExposedBucketsSummaryInternalServerError{}
}

/*
PostV1ExposedBucketsSummaryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostV1ExposedBucketsSummaryInternalServerError struct {
	Payload interface{}
}

// IsSuccess returns true when this post v1 exposed buckets summary internal server error response has a 2xx status code
func (o *PostV1ExposedBucketsSummaryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 exposed buckets summary internal server error response has a 3xx status code
func (o *PostV1ExposedBucketsSummaryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 exposed buckets summary internal server error response has a 4xx status code
func (o *PostV1ExposedBucketsSummaryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 exposed buckets summary internal server error response has a 5xx status code
func (o *PostV1ExposedBucketsSummaryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post v1 exposed buckets summary internal server error response a status code equal to that given
func (o *PostV1ExposedBucketsSummaryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post v1 exposed buckets summary internal server error response
func (o *PostV1ExposedBucketsSummaryInternalServerError) Code() int {
	return 500
}

func (o *PostV1ExposedBucketsSummaryInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/buckets/summary][%d] postV1ExposedBucketsSummaryInternalServerError %s", 500, payload)
}

func (o *PostV1ExposedBucketsSummaryInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/buckets/summary][%d] postV1ExposedBucketsSummaryInternalServerError %s", 500, payload)
}

func (o *PostV1ExposedBucketsSummaryInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *PostV1ExposedBucketsSummaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostV1ExposedBucketsSummaryOKBody post v1 exposed buckets summary o k body
swagger:model PostV1ExposedBucketsSummaryOKBody
*/
type PostV1ExposedBucketsSummaryOKBody struct {
	models.ExposedAPIResponse

	// data
	Data *models.ExposedAggregate `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostV1ExposedBucketsSummaryOKBody) UnmarshalJSON(raw []byte) error {
	// PostV1ExposedBucketsSummaryOKBodyAO0
	var postV1ExposedBucketsSummaryOKBodyAO0 models.ExposedAPIResponse
	if err := swag.ReadJSON(raw, &postV1ExposedBucketsSummaryOKBodyAO0); err != nil {
		return err
	}
	o.ExposedAPIResponse = postV1ExposedBucketsSummaryOKBodyAO0

	// PostV1ExposedBucketsSummaryOKBodyAO1
	var dataPostV1ExposedBucketsSummaryOKBodyAO1 struct {
		Data *models.ExposedAggregate `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostV1ExposedBucketsSummaryOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostV1ExposedBucketsSummaryOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostV1ExposedBucketsSummaryOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postV1ExposedBucketsSummaryOKBodyAO0, err := swag.WriteJSON(o.ExposedAPIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postV1ExposedBucketsSummaryOKBodyAO0)
	var dataPostV1ExposedBucketsSummaryOKBodyAO1 struct {
		Data *models.ExposedAggregate `json:"data,omitempty"`
	}

	dataPostV1ExposedBucketsSummaryOKBodyAO1.Data = o.Data

	jsonDataPostV1ExposedBucketsSummaryOKBodyAO1, errPostV1ExposedBucketsSummaryOKBodyAO1 := swag.WriteJSON(dataPostV1ExposedBucketsSummaryOKBodyAO1)
	if errPostV1ExposedBucketsSummaryOKBodyAO1 != nil {
		return nil, errPostV1ExposedBucketsSummaryOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostV1ExposedBucketsSummaryOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post v1 exposed buckets summary o k body
func (o *PostV1ExposedBucketsSummaryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ExposedAPIResponse
	if err := o.ExposedAPIResponse.Validate(formats); err != nil {
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

func (o *PostV1ExposedBucketsSummaryOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postV1ExposedBucketsSummaryOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postV1ExposedBucketsSummaryOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post v1 exposed buckets summary o k body based on the context it is used
func (o *PostV1ExposedBucketsSummaryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ExposedAPIResponse
	if err := o.ExposedAPIResponse.ContextValidate(ctx, formats); err != nil {
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

func (o *PostV1ExposedBucketsSummaryOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postV1ExposedBucketsSummaryOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postV1ExposedBucketsSummaryOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostV1ExposedBucketsSummaryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostV1ExposedBucketsSummaryOKBody) UnmarshalBinary(b []byte) error {
	var res PostV1ExposedBucketsSummaryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
