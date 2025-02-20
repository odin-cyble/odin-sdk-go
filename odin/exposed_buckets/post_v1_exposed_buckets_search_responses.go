// Code generated by go-swagger; DO NOT EDIT.

package exposed_buckets

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

// PostV1ExposedBucketsSearchReader is a Reader for the PostV1ExposedBucketsSearch structure.
type PostV1ExposedBucketsSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ExposedBucketsSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1ExposedBucketsSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1ExposedBucketsSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostV1ExposedBucketsSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/exposed/buckets/search] PostV1ExposedBucketsSearch", response, response.Code())
	}
}

// NewPostV1ExposedBucketsSearchOK creates a PostV1ExposedBucketsSearchOK with default headers values
func NewPostV1ExposedBucketsSearchOK() *PostV1ExposedBucketsSearchOK {
	return &PostV1ExposedBucketsSearchOK{}
}

/*
PostV1ExposedBucketsSearchOK describes a response with status code 200, with default header values.

OK
*/
type PostV1ExposedBucketsSearchOK struct {
	Payload *models.ExposedBucketAPIResponse
}

// IsSuccess returns true when this post v1 exposed buckets search o k response has a 2xx status code
func (o *PostV1ExposedBucketsSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 exposed buckets search o k response has a 3xx status code
func (o *PostV1ExposedBucketsSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 exposed buckets search o k response has a 4xx status code
func (o *PostV1ExposedBucketsSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 exposed buckets search o k response has a 5xx status code
func (o *PostV1ExposedBucketsSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 exposed buckets search o k response a status code equal to that given
func (o *PostV1ExposedBucketsSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 exposed buckets search o k response
func (o *PostV1ExposedBucketsSearchOK) Code() int {
	return 200
}

func (o *PostV1ExposedBucketsSearchOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/buckets/search][%d] postV1ExposedBucketsSearchOK %s", 200, payload)
}

func (o *PostV1ExposedBucketsSearchOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/buckets/search][%d] postV1ExposedBucketsSearchOK %s", 200, payload)
}

func (o *PostV1ExposedBucketsSearchOK) GetPayload() *models.ExposedBucketAPIResponse {
	return o.Payload
}

func (o *PostV1ExposedBucketsSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExposedBucketAPIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ExposedBucketsSearchBadRequest creates a PostV1ExposedBucketsSearchBadRequest with default headers values
func NewPostV1ExposedBucketsSearchBadRequest() *PostV1ExposedBucketsSearchBadRequest {
	return &PostV1ExposedBucketsSearchBadRequest{}
}

/*
PostV1ExposedBucketsSearchBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1ExposedBucketsSearchBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this post v1 exposed buckets search bad request response has a 2xx status code
func (o *PostV1ExposedBucketsSearchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 exposed buckets search bad request response has a 3xx status code
func (o *PostV1ExposedBucketsSearchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 exposed buckets search bad request response has a 4xx status code
func (o *PostV1ExposedBucketsSearchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 exposed buckets search bad request response has a 5xx status code
func (o *PostV1ExposedBucketsSearchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 exposed buckets search bad request response a status code equal to that given
func (o *PostV1ExposedBucketsSearchBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post v1 exposed buckets search bad request response
func (o *PostV1ExposedBucketsSearchBadRequest) Code() int {
	return 400
}

func (o *PostV1ExposedBucketsSearchBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/buckets/search][%d] postV1ExposedBucketsSearchBadRequest %s", 400, payload)
}

func (o *PostV1ExposedBucketsSearchBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/buckets/search][%d] postV1ExposedBucketsSearchBadRequest %s", 400, payload)
}

func (o *PostV1ExposedBucketsSearchBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PostV1ExposedBucketsSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ExposedBucketsSearchInternalServerError creates a PostV1ExposedBucketsSearchInternalServerError with default headers values
func NewPostV1ExposedBucketsSearchInternalServerError() *PostV1ExposedBucketsSearchInternalServerError {
	return &PostV1ExposedBucketsSearchInternalServerError{}
}

/*
PostV1ExposedBucketsSearchInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostV1ExposedBucketsSearchInternalServerError struct {
	Payload interface{}
}

// IsSuccess returns true when this post v1 exposed buckets search internal server error response has a 2xx status code
func (o *PostV1ExposedBucketsSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 exposed buckets search internal server error response has a 3xx status code
func (o *PostV1ExposedBucketsSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 exposed buckets search internal server error response has a 4xx status code
func (o *PostV1ExposedBucketsSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 exposed buckets search internal server error response has a 5xx status code
func (o *PostV1ExposedBucketsSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post v1 exposed buckets search internal server error response a status code equal to that given
func (o *PostV1ExposedBucketsSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post v1 exposed buckets search internal server error response
func (o *PostV1ExposedBucketsSearchInternalServerError) Code() int {
	return 500
}

func (o *PostV1ExposedBucketsSearchInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/buckets/search][%d] postV1ExposedBucketsSearchInternalServerError %s", 500, payload)
}

func (o *PostV1ExposedBucketsSearchInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/exposed/buckets/search][%d] postV1ExposedBucketsSearchInternalServerError %s", 500, payload)
}

func (o *PostV1ExposedBucketsSearchInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *PostV1ExposedBucketsSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
