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

// NewKeyUnlinkHTTPParams creates a new KeyUnlinkHTTPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKeyUnlinkHTTPParams() *KeyUnlinkHTTPParams {
	return &KeyUnlinkHTTPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKeyUnlinkHTTPParamsWithTimeout creates a new KeyUnlinkHTTPParams object
// with the ability to set a timeout on a request.
func NewKeyUnlinkHTTPParamsWithTimeout(timeout time.Duration) *KeyUnlinkHTTPParams {
	return &KeyUnlinkHTTPParams{
		timeout: timeout,
	}
}

// NewKeyUnlinkHTTPParamsWithContext creates a new KeyUnlinkHTTPParams object
// with the ability to set a context for a request.
func NewKeyUnlinkHTTPParamsWithContext(ctx context.Context) *KeyUnlinkHTTPParams {
	return &KeyUnlinkHTTPParams{
		Context: ctx,
	}
}

// NewKeyUnlinkHTTPParamsWithHTTPClient creates a new KeyUnlinkHTTPParams object
// with the ability to set a custom HTTPClient for a request.
func NewKeyUnlinkHTTPParamsWithHTTPClient(client *http.Client) *KeyUnlinkHTTPParams {
	return &KeyUnlinkHTTPParams{
		HTTPClient: client,
	}
}

/* KeyUnlinkHTTPParams contains all the parameters to send to the API endpoint
   for the key unlink Http operation.

   Typically these are written to a http.Request.
*/
type KeyUnlinkHTTPParams struct {

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

// WithDefaults hydrates default values in the key unlink Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeyUnlinkHTTPParams) WithDefaults() *KeyUnlinkHTTPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the key unlink Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeyUnlinkHTTPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the key unlink Http params
func (o *KeyUnlinkHTTPParams) WithTimeout(timeout time.Duration) *KeyUnlinkHTTPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the key unlink Http params
func (o *KeyUnlinkHTTPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the key unlink Http params
func (o *KeyUnlinkHTTPParams) WithContext(ctx context.Context) *KeyUnlinkHTTPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the key unlink Http params
func (o *KeyUnlinkHTTPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the key unlink Http params
func (o *KeyUnlinkHTTPParams) WithHTTPClient(client *http.Client) *KeyUnlinkHTTPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the key unlink Http params
func (o *KeyUnlinkHTTPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the key unlink Http params
func (o *KeyUnlinkHTTPParams) WithID(id int64) *KeyUnlinkHTTPParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the key unlink Http params
func (o *KeyUnlinkHTTPParams) SetID(id int64) {
	o.ID = id
}

// WithKeyID adds the keyID to the key unlink Http params
func (o *KeyUnlinkHTTPParams) WithKeyID(keyID int64) *KeyUnlinkHTTPParams {
	o.SetKeyID(keyID)
	return o
}

// SetKeyID adds the keyId to the key unlink Http params
func (o *KeyUnlinkHTTPParams) SetKeyID(keyID int64) {
	o.KeyID = keyID
}

// WriteToRequest writes these params to a swagger request
func (o *KeyUnlinkHTTPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
