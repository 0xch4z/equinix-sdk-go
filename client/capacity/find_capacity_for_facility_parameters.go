// Code generated by go-swagger; DO NOT EDIT.

package capacity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewFindCapacityForFacilityParams creates a new FindCapacityForFacilityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindCapacityForFacilityParams() *FindCapacityForFacilityParams {
	return &FindCapacityForFacilityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindCapacityForFacilityParamsWithTimeout creates a new FindCapacityForFacilityParams object
// with the ability to set a timeout on a request.
func NewFindCapacityForFacilityParamsWithTimeout(timeout time.Duration) *FindCapacityForFacilityParams {
	return &FindCapacityForFacilityParams{
		timeout: timeout,
	}
}

// NewFindCapacityForFacilityParamsWithContext creates a new FindCapacityForFacilityParams object
// with the ability to set a context for a request.
func NewFindCapacityForFacilityParamsWithContext(ctx context.Context) *FindCapacityForFacilityParams {
	return &FindCapacityForFacilityParams{
		Context: ctx,
	}
}

// NewFindCapacityForFacilityParamsWithHTTPClient creates a new FindCapacityForFacilityParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindCapacityForFacilityParamsWithHTTPClient(client *http.Client) *FindCapacityForFacilityParams {
	return &FindCapacityForFacilityParams{
		HTTPClient: client,
	}
}

/* FindCapacityForFacilityParams contains all the parameters to send to the API endpoint
   for the find capacity for facility operation.

   Typically these are written to a http.Request.
*/
type FindCapacityForFacilityParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find capacity for facility params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindCapacityForFacilityParams) WithDefaults() *FindCapacityForFacilityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find capacity for facility params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindCapacityForFacilityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find capacity for facility params
func (o *FindCapacityForFacilityParams) WithTimeout(timeout time.Duration) *FindCapacityForFacilityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find capacity for facility params
func (o *FindCapacityForFacilityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find capacity for facility params
func (o *FindCapacityForFacilityParams) WithContext(ctx context.Context) *FindCapacityForFacilityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find capacity for facility params
func (o *FindCapacityForFacilityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find capacity for facility params
func (o *FindCapacityForFacilityParams) WithHTTPClient(client *http.Client) *FindCapacityForFacilityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find capacity for facility params
func (o *FindCapacityForFacilityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindCapacityForFacilityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
