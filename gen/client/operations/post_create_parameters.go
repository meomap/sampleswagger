// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/meomap/sampleswagger/gen/models"
)

// NewPostCreateParams creates a new PostCreateParams object
// with the default values initialized.
func NewPostCreateParams() *PostCreateParams {
	var ()
	return &PostCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCreateParamsWithTimeout creates a new PostCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCreateParamsWithTimeout(timeout time.Duration) *PostCreateParams {
	var ()
	return &PostCreateParams{

		timeout: timeout,
	}
}

// NewPostCreateParamsWithContext creates a new PostCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCreateParamsWithContext(ctx context.Context) *PostCreateParams {
	var ()
	return &PostCreateParams{

		Context: ctx,
	}
}

// NewPostCreateParamsWithHTTPClient creates a new PostCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCreateParamsWithHTTPClient(client *http.Client) *PostCreateParams {
	var ()
	return &PostCreateParams{
		HTTPClient: client,
	}
}

/*PostCreateParams contains all the parameters to send to the API endpoint
for the post create operation typically these are written to a http.Request
*/
type PostCreateParams struct {

	/*Body
	  Resource

	*/
	Body *models.Resource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post create params
func (o *PostCreateParams) WithTimeout(timeout time.Duration) *PostCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post create params
func (o *PostCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post create params
func (o *PostCreateParams) WithContext(ctx context.Context) *PostCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post create params
func (o *PostCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post create params
func (o *PostCreateParams) WithHTTPClient(client *http.Client) *PostCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post create params
func (o *PostCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post create params
func (o *PostCreateParams) WithBody(body *models.Resource) *PostCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post create params
func (o *PostCreateParams) SetBody(body *models.Resource) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
