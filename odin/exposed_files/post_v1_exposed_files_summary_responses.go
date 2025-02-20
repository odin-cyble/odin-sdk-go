// Code generated by go-swagger; DO NOT EDIT.

package exposed_files

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

// PostV1ExposedFilesSummaryReader is a Reader for the PostV1ExposedFilesSummary structure.
type PostV1ExposedFilesSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ExposedFilesSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1ExposedFilesSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1ExposedFilesSummaryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostV1ExposedFilesSummaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/exposed/files/summary] PostV1ExposedFilesSummary", response, response.Code())
	}
}

// NewPostV1ExposedFilesSummaryOK creates a PostV1ExposedFilesSummaryOK with default headers values
func NewPostV1ExposedFilesSummaryOK() *PostV1ExposedFilesSummaryOK {
	return &PostV1ExposedFilesSummaryOK{}
}

/*
PostV1ExposedFilesSummaryOK describes a response with status code 200, with default header values.

OK
*/
type PostV1ExposedFilesSummaryOK struct {
	Payload *PostV1ExposedFilesSummaryOKBody
}

// IsSuccess returns true when this post v1 exposed files summary o k response has a 2xx status code
func (o *PostV1ExposedFilesSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 exposed files summary o k response has a 3xx status code
func (o *PostV1ExposedFilesSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 exposed files summary o k response has a 4xx status code
func (o *PostV1ExposedFilesSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 exposed files summary o k response has a 5xx status code
func (o *PostV1ExposedFilesSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 exposed files summary o k response a status code equal to that given
func (o *PostV1ExposedFilesSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 exposed files summary o k response
func (o *PostV1ExposedFilesSummaryOK) Code() int {
	return 200
}

func (o *PostV1ExposedFilesSummaryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/files/summary][%d] postV1ExposedFilesSummaryOK %s", 200, payload)
}

func (o *PostV1ExposedFilesSummaryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/files/summary][%d] postV1ExposedFilesSummaryOK %s", 200, payload)
}

func (o *PostV1ExposedFilesSummaryOK) GetPayload() *PostV1ExposedFilesSummaryOKBody {
	return o.Payload
}

func (o *PostV1ExposedFilesSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostV1ExposedFilesSummaryOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ExposedFilesSummaryBadRequest creates a PostV1ExposedFilesSummaryBadRequest with default headers values
func NewPostV1ExposedFilesSummaryBadRequest() *PostV1ExposedFilesSummaryBadRequest {
	return &PostV1ExposedFilesSummaryBadRequest{}
}

/*
PostV1ExposedFilesSummaryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1ExposedFilesSummaryBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this post v1 exposed files summary bad request response has a 2xx status code
func (o *PostV1ExposedFilesSummaryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 exposed files summary bad request response has a 3xx status code
func (o *PostV1ExposedFilesSummaryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 exposed files summary bad request response has a 4xx status code
func (o *PostV1ExposedFilesSummaryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 exposed files summary bad request response has a 5xx status code
func (o *PostV1ExposedFilesSummaryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 exposed files summary bad request response a status code equal to that given
func (o *PostV1ExposedFilesSummaryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post v1 exposed files summary bad request response
func (o *PostV1ExposedFilesSummaryBadRequest) Code() int {
	return 400
}

func (o *PostV1ExposedFilesSummaryBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/files/summary][%d] postV1ExposedFilesSummaryBadRequest %s", 400, payload)
}

func (o *PostV1ExposedFilesSummaryBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/files/summary][%d] postV1ExposedFilesSummaryBadRequest %s", 400, payload)
}

func (o *PostV1ExposedFilesSummaryBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PostV1ExposedFilesSummaryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ExposedFilesSummaryInternalServerError creates a PostV1ExposedFilesSummaryInternalServerError with default headers values
func NewPostV1ExposedFilesSummaryInternalServerError() *PostV1ExposedFilesSummaryInternalServerError {
	return &PostV1ExposedFilesSummaryInternalServerError{}
}

/*
PostV1ExposedFilesSummaryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostV1ExposedFilesSummaryInternalServerError struct {
	Payload interface{}
}

// IsSuccess returns true when this post v1 exposed files summary internal server error response has a 2xx status code
func (o *PostV1ExposedFilesSummaryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 exposed files summary internal server error response has a 3xx status code
func (o *PostV1ExposedFilesSummaryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 exposed files summary internal server error response has a 4xx status code
func (o *PostV1ExposedFilesSummaryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 exposed files summary internal server error response has a 5xx status code
func (o *PostV1ExposedFilesSummaryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post v1 exposed files summary internal server error response a status code equal to that given
func (o *PostV1ExposedFilesSummaryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post v1 exposed files summary internal server error response
func (o *PostV1ExposedFilesSummaryInternalServerError) Code() int {
	return 500
}

func (o *PostV1ExposedFilesSummaryInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/files/summary][%d] postV1ExposedFilesSummaryInternalServerError %s", 500, payload)
}

func (o *PostV1ExposedFilesSummaryInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/files/summary][%d] postV1ExposedFilesSummaryInternalServerError %s", 500, payload)
}

func (o *PostV1ExposedFilesSummaryInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *PostV1ExposedFilesSummaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostV1ExposedFilesSummaryOKBody post v1 exposed files summary o k body
swagger:model PostV1ExposedFilesSummaryOKBody
*/
type PostV1ExposedFilesSummaryOKBody struct {
	models.ExposedAPIResponse

	// data
	Data *models.ExposedAggregate `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostV1ExposedFilesSummaryOKBody) UnmarshalJSON(raw []byte) error {
	// PostV1ExposedFilesSummaryOKBodyAO0
	var postV1ExposedFilesSummaryOKBodyAO0 models.ExposedAPIResponse
	if err := swag.ReadJSON(raw, &postV1ExposedFilesSummaryOKBodyAO0); err != nil {
		return err
	}
	o.ExposedAPIResponse = postV1ExposedFilesSummaryOKBodyAO0

	// PostV1ExposedFilesSummaryOKBodyAO1
	var dataPostV1ExposedFilesSummaryOKBodyAO1 struct {
		Data *models.ExposedAggregate `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostV1ExposedFilesSummaryOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostV1ExposedFilesSummaryOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostV1ExposedFilesSummaryOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postV1ExposedFilesSummaryOKBodyAO0, err := swag.WriteJSON(o.ExposedAPIResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postV1ExposedFilesSummaryOKBodyAO0)
	var dataPostV1ExposedFilesSummaryOKBodyAO1 struct {
		Data *models.ExposedAggregate `json:"data,omitempty"`
	}

	dataPostV1ExposedFilesSummaryOKBodyAO1.Data = o.Data

	jsonDataPostV1ExposedFilesSummaryOKBodyAO1, errPostV1ExposedFilesSummaryOKBodyAO1 := swag.WriteJSON(dataPostV1ExposedFilesSummaryOKBodyAO1)
	if errPostV1ExposedFilesSummaryOKBodyAO1 != nil {
		return nil, errPostV1ExposedFilesSummaryOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostV1ExposedFilesSummaryOKBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post v1 exposed files summary o k body
func (o *PostV1ExposedFilesSummaryOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostV1ExposedFilesSummaryOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postV1ExposedFilesSummaryOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postV1ExposedFilesSummaryOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post v1 exposed files summary o k body based on the context it is used
func (o *PostV1ExposedFilesSummaryOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostV1ExposedFilesSummaryOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {

		if swag.IsZero(o.Data) { // not required
			return nil
		}

		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postV1ExposedFilesSummaryOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postV1ExposedFilesSummaryOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostV1ExposedFilesSummaryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostV1ExposedFilesSummaryOKBody) UnmarshalBinary(b []byte) error {
	var res PostV1ExposedFilesSummaryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
