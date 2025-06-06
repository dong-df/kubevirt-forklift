// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// LocalCifsGroupDeleteReader is a Reader for the LocalCifsGroupDelete structure.
type LocalCifsGroupDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsGroupDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalCifsGroupDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsGroupDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsGroupDeleteOK creates a LocalCifsGroupDeleteOK with default headers values
func NewLocalCifsGroupDeleteOK() *LocalCifsGroupDeleteOK {
	return &LocalCifsGroupDeleteOK{}
}

/*
LocalCifsGroupDeleteOK describes a response with status code 200, with default header values.

OK
*/
type LocalCifsGroupDeleteOK struct {
}

// IsSuccess returns true when this local cifs group delete o k response has a 2xx status code
func (o *LocalCifsGroupDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local cifs group delete o k response has a 3xx status code
func (o *LocalCifsGroupDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local cifs group delete o k response has a 4xx status code
func (o *LocalCifsGroupDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this local cifs group delete o k response has a 5xx status code
func (o *LocalCifsGroupDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this local cifs group delete o k response a status code equal to that given
func (o *LocalCifsGroupDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *LocalCifsGroupDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/local-groups/{svm.uuid}/{sid}][%d] localCifsGroupDeleteOK ", 200)
}

func (o *LocalCifsGroupDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/local-groups/{svm.uuid}/{sid}][%d] localCifsGroupDeleteOK ", 200)
}

func (o *LocalCifsGroupDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLocalCifsGroupDeleteDefault creates a LocalCifsGroupDeleteDefault with default headers values
func NewLocalCifsGroupDeleteDefault(code int) *LocalCifsGroupDeleteDefault {
	return &LocalCifsGroupDeleteDefault{
		_statusCode: code,
	}
}

/*
LocalCifsGroupDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type LocalCifsGroupDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the local cifs group delete default response
func (o *LocalCifsGroupDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this local cifs group delete default response has a 2xx status code
func (o *LocalCifsGroupDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local cifs group delete default response has a 3xx status code
func (o *LocalCifsGroupDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local cifs group delete default response has a 4xx status code
func (o *LocalCifsGroupDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local cifs group delete default response has a 5xx status code
func (o *LocalCifsGroupDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local cifs group delete default response a status code equal to that given
func (o *LocalCifsGroupDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *LocalCifsGroupDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/local-groups/{svm.uuid}/{sid}][%d] local_cifs_group_delete default  %+v", o._statusCode, o.Payload)
}

func (o *LocalCifsGroupDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/local-groups/{svm.uuid}/{sid}][%d] local_cifs_group_delete default  %+v", o._statusCode, o.Payload)
}

func (o *LocalCifsGroupDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsGroupDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
