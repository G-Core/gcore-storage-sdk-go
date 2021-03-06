// Code generated by go-swagger; DO NOT EDIT.

package storage

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
	"github.com/go-openapi/swag"
)

// NewKeyLinkHTTPParams creates a new KeyLinkHTTPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKeyLinkHTTPParams() *KeyLinkHTTPParams {
	return &KeyLinkHTTPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKeyLinkHTTPParamsWithTimeout creates a new KeyLinkHTTPParams object
// with the ability to set a timeout on a request.
func NewKeyLinkHTTPParamsWithTimeout(timeout time.Duration) *KeyLinkHTTPParams {
	return &KeyLinkHTTPParams{
		timeout: timeout,
	}
}

// NewKeyLinkHTTPParamsWithContext creates a new KeyLinkHTTPParams object
// with the ability to set a context for a request.
func NewKeyLinkHTTPParamsWithContext(ctx context.Context) *KeyLinkHTTPParams {
	return &KeyLinkHTTPParams{
		Context: ctx,
	}
}

// NewKeyLinkHTTPParamsWithHTTPClient creates a new KeyLinkHTTPParams object
// with the ability to set a custom HTTPClient for a request.
func NewKeyLinkHTTPParamsWithHTTPClient(client *http.Client) *KeyLinkHTTPParams {
	return &KeyLinkHTTPParams{
		HTTPClient: client,
	}
}

/* KeyLinkHTTPParams contains all the parameters to send to the API endpoint
   for the key link Http operation.

   Typically these are written to a http.Request.
*/
type KeyLinkHTTPParams struct {

	// ID.
	//
	// Format: int64
	ID int64

	// KeyID.
	//
	// Format: int64
	KeyID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the key link Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeyLinkHTTPParams) WithDefaults() *KeyLinkHTTPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the key link Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeyLinkHTTPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the key link Http params
func (o *KeyLinkHTTPParams) WithTimeout(timeout time.Duration) *KeyLinkHTTPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the key link Http params
func (o *KeyLinkHTTPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the key link Http params
func (o *KeyLinkHTTPParams) WithContext(ctx context.Context) *KeyLinkHTTPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the key link Http params
func (o *KeyLinkHTTPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the key link Http params
func (o *KeyLinkHTTPParams) WithHTTPClient(client *http.Client) *KeyLinkHTTPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the key link Http params
func (o *KeyLinkHTTPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the key link Http params
func (o *KeyLinkHTTPParams) WithID(id int64) *KeyLinkHTTPParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the key link Http params
func (o *KeyLinkHTTPParams) SetID(id int64) {
	o.ID = id
}

// WithKeyID adds the keyID to the key link Http params
func (o *KeyLinkHTTPParams) WithKeyID(keyID int64) *KeyLinkHTTPParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the key link Http params
func (o *KeyLinkHTTPParams) SetKeyID(keyID int64) {
	o.KeyID = keyID
}

// WriteToRequest writes these params to a swagger request
func (o *KeyLinkHTTPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param key_id
	if err := r.SetPathParam("key_id", swag.FormatInt64(o.KeyID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
