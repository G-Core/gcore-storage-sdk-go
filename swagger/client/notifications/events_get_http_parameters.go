// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

// NewEventsGetHTTPParams creates a new EventsGetHTTPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEventsGetHTTPParams() *EventsGetHTTPParams {
	return &EventsGetHTTPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEventsGetHTTPParamsWithTimeout creates a new EventsGetHTTPParams object
// with the ability to set a timeout on a request.
func NewEventsGetHTTPParamsWithTimeout(timeout time.Duration) *EventsGetHTTPParams {
	return &EventsGetHTTPParams{
		timeout: timeout,
	}
}

// NewEventsGetHTTPParamsWithContext creates a new EventsGetHTTPParams object
// with the ability to set a context for a request.
func NewEventsGetHTTPParamsWithContext(ctx context.Context) *EventsGetHTTPParams {
	return &EventsGetHTTPParams{
		Context: ctx,
	}
}

// NewEventsGetHTTPParamsWithHTTPClient creates a new EventsGetHTTPParams object
// with the ability to set a custom HTTPClient for a request.
func NewEventsGetHTTPParamsWithHTTPClient(client *http.Client) *EventsGetHTTPParams {
	return &EventsGetHTTPParams{
		HTTPClient: client,
	}
}

/* EventsGetHTTPParams contains all the parameters to send to the API endpoint
   for the events get Http operation.

   Typically these are written to a http.Request.
*/
type EventsGetHTTPParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the events get Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EventsGetHTTPParams) WithDefaults() *EventsGetHTTPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the events get Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EventsGetHTTPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the events get Http params
func (o *EventsGetHTTPParams) WithTimeout(timeout time.Duration) *EventsGetHTTPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the events get Http params
func (o *EventsGetHTTPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the events get Http params
func (o *EventsGetHTTPParams) WithContext(ctx context.Context) *EventsGetHTTPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the events get Http params
func (o *EventsGetHTTPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the events get Http params
func (o *EventsGetHTTPParams) WithHTTPClient(client *http.Client) *EventsGetHTTPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the events get Http params
func (o *EventsGetHTTPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *EventsGetHTTPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
