// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteIcecreamParams creates a new DeleteIcecreamParams object
// with the default values initialized.
func NewDeleteIcecreamParams() *DeleteIcecreamParams {
	var ()
	return &DeleteIcecreamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIcecreamParamsWithTimeout creates a new DeleteIcecreamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteIcecreamParamsWithTimeout(timeout time.Duration) *DeleteIcecreamParams {
	var ()
	return &DeleteIcecreamParams{

		timeout: timeout,
	}
}

// NewDeleteIcecreamParamsWithContext creates a new DeleteIcecreamParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteIcecreamParamsWithContext(ctx context.Context) *DeleteIcecreamParams {
	var ()
	return &DeleteIcecreamParams{

		Context: ctx,
	}
}

// NewDeleteIcecreamParamsWithHTTPClient creates a new DeleteIcecreamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteIcecreamParamsWithHTTPClient(client *http.Client) *DeleteIcecreamParams {
	var ()
	return &DeleteIcecreamParams{
		HTTPClient: client,
	}
}

/*DeleteIcecreamParams contains all the parameters to send to the API endpoint
for the delete icecream operation typically these are written to a http.Request
*/
type DeleteIcecreamParams struct {

	/*ProductID
	  Product id

	*/
	ProductID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete icecream params
func (o *DeleteIcecreamParams) WithTimeout(timeout time.Duration) *DeleteIcecreamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete icecream params
func (o *DeleteIcecreamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete icecream params
func (o *DeleteIcecreamParams) WithContext(ctx context.Context) *DeleteIcecreamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete icecream params
func (o *DeleteIcecreamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete icecream params
func (o *DeleteIcecreamParams) WithHTTPClient(client *http.Client) *DeleteIcecreamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete icecream params
func (o *DeleteIcecreamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProductID adds the productID to the delete icecream params
func (o *DeleteIcecreamParams) WithProductID(productID string) *DeleteIcecreamParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the delete icecream params
func (o *DeleteIcecreamParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIcecreamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param product_id
	qrProductID := o.ProductID
	qProductID := qrProductID
	if qProductID != "" {
		if err := r.SetQueryParam("product_id", qProductID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}