// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/meomap/sampleswagger/gen/models"
)

// PostCreateReader is a Reader for the PostCreate structure.
type PostCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCreateCreated creates a PostCreateCreated with default headers values
func NewPostCreateCreated() *PostCreateCreated {
	return &PostCreateCreated{}
}

/*PostCreateCreated handles this case with default header values.

Resource created
*/
type PostCreateCreated struct {
	Payload *models.Resource
}

func (o *PostCreateCreated) Error() string {
	return fmt.Sprintf("[POST /create][%d] postCreateCreated  %+v", 201, o.Payload)
}

func (o *PostCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Resource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
