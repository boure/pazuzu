package features

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"swaggen/models"
)

// GetAPIFeaturesNameReader is a Reader for the GetAPIFeaturesName structure.
type GetAPIFeaturesNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIFeaturesNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIFeaturesNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAPIFeaturesNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAPIFeaturesNameOK creates a GetAPIFeaturesNameOK with default headers values
func NewGetAPIFeaturesNameOK() *GetAPIFeaturesNameOK {
	return &GetAPIFeaturesNameOK{}
}

/*GetAPIFeaturesNameOK handles this case with default header values.

a singel feature
*/
type GetAPIFeaturesNameOK struct {
	Payload *models.Feature
}

func (o *GetAPIFeaturesNameOK) Error() string {
	return fmt.Sprintf("[GET /api/features/{name}][%d] getApiFeaturesNameOK  %+v", 200, o.Payload)
}

func (o *GetAPIFeaturesNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Feature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIFeaturesNameDefault creates a GetAPIFeaturesNameDefault with default headers values
func NewGetAPIFeaturesNameDefault(code int) *GetAPIFeaturesNameDefault {
	return &GetAPIFeaturesNameDefault{
		_statusCode: code,
	}
}

/*GetAPIFeaturesNameDefault handles this case with default header values.

Unexpected error
*/
type GetAPIFeaturesNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get API features name default response
func (o *GetAPIFeaturesNameDefault) Code() int {
	return o._statusCode
}

func (o *GetAPIFeaturesNameDefault) Error() string {
	return fmt.Sprintf("[GET /api/features/{name}][%d] GetAPIFeaturesName default  %+v", o._statusCode, o.Payload)
}

func (o *GetAPIFeaturesNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
