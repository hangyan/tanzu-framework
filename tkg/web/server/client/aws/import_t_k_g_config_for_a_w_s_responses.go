// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// ImportTKGConfigForAWSReader is a Reader for the ImportTKGConfigForAWS structure.
type ImportTKGConfigForAWSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportTKGConfigForAWSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportTKGConfigForAWSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportTKGConfigForAWSBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImportTKGConfigForAWSUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImportTKGConfigForAWSInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewImportTKGConfigForAWSOK creates a ImportTKGConfigForAWSOK with default headers values
func NewImportTKGConfigForAWSOK() *ImportTKGConfigForAWSOK {
	return &ImportTKGConfigForAWSOK{}
}

/*ImportTKGConfigForAWSOK handles this case with default header values.

Generated TKG configuration successfully
*/
type ImportTKGConfigForAWSOK struct {
	Payload *models.AWSRegionalClusterParams
}

func (o *ImportTKGConfigForAWSOK) Error() string {
	return fmt.Sprintf("[POST /api/providers/aws/config/import][%d] importTKGConfigForAWSOK  %+v", 200, o.Payload)
}

func (o *ImportTKGConfigForAWSOK) GetPayload() *models.AWSRegionalClusterParams {
	return o.Payload
}

func (o *ImportTKGConfigForAWSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AWSRegionalClusterParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportTKGConfigForAWSBadRequest creates a ImportTKGConfigForAWSBadRequest with default headers values
func NewImportTKGConfigForAWSBadRequest() *ImportTKGConfigForAWSBadRequest {
	return &ImportTKGConfigForAWSBadRequest{}
}

/*ImportTKGConfigForAWSBadRequest handles this case with default header values.

Bad request
*/
type ImportTKGConfigForAWSBadRequest struct {
	Payload *models.Error
}

func (o *ImportTKGConfigForAWSBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/providers/aws/config/import][%d] importTKGConfigForAWSBadRequest  %+v", 400, o.Payload)
}

func (o *ImportTKGConfigForAWSBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportTKGConfigForAWSBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportTKGConfigForAWSUnauthorized creates a ImportTKGConfigForAWSUnauthorized with default headers values
func NewImportTKGConfigForAWSUnauthorized() *ImportTKGConfigForAWSUnauthorized {
	return &ImportTKGConfigForAWSUnauthorized{}
}

/*ImportTKGConfigForAWSUnauthorized handles this case with default header values.

Incorrect credentials
*/
type ImportTKGConfigForAWSUnauthorized struct {
	Payload *models.Error
}

func (o *ImportTKGConfigForAWSUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/providers/aws/config/import][%d] importTKGConfigForAWSUnauthorized  %+v", 401, o.Payload)
}

func (o *ImportTKGConfigForAWSUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportTKGConfigForAWSUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportTKGConfigForAWSInternalServerError creates a ImportTKGConfigForAWSInternalServerError with default headers values
func NewImportTKGConfigForAWSInternalServerError() *ImportTKGConfigForAWSInternalServerError {
	return &ImportTKGConfigForAWSInternalServerError{}
}

/*ImportTKGConfigForAWSInternalServerError handles this case with default header values.

Internal server error
*/
type ImportTKGConfigForAWSInternalServerError struct {
	Payload *models.Error
}

func (o *ImportTKGConfigForAWSInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/providers/aws/config/import][%d] importTKGConfigForAWSInternalServerError  %+v", 500, o.Payload)
}

func (o *ImportTKGConfigForAWSInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportTKGConfigForAWSInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}