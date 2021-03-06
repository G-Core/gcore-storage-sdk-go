// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Events Events for storage notification
//
// swagger:model Events
type Events struct {

	// reaching files quantity
	ReachingFilesQuantity []*LimitByCount `json:"reaching_files_quantity"`

	// reaching used requests
	ReachingUsedRequests []*LimitByCount `json:"reaching_used_requests"`

	// reaching used space
	ReachingUsedSpace []*LimitByBytes `json:"reaching_used_space"`

	// reaching used traffic
	ReachingUsedTraffic []*LimitByBytes `json:"reaching_used_traffic"`
}

// Validate validates this events
func (m *Events) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReachingFilesQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReachingUsedRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReachingUsedSpace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReachingUsedTraffic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Events) validateReachingFilesQuantity(formats strfmt.Registry) error {
	if swag.IsZero(m.ReachingFilesQuantity) { // not required
		return nil
	}

	for i := 0; i < len(m.ReachingFilesQuantity); i++ {
		if swag.IsZero(m.ReachingFilesQuantity[i]) { // not required
			continue
		}

		if m.ReachingFilesQuantity[i] != nil {
			if err := m.ReachingFilesQuantity[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reaching_files_quantity" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reaching_files_quantity" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Events) validateReachingUsedRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.ReachingUsedRequests) { // not required
		return nil
	}

	for i := 0; i < len(m.ReachingUsedRequests); i++ {
		if swag.IsZero(m.ReachingUsedRequests[i]) { // not required
			continue
		}

		if m.ReachingUsedRequests[i] != nil {
			if err := m.ReachingUsedRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reaching_used_requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reaching_used_requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Events) validateReachingUsedSpace(formats strfmt.Registry) error {
	if swag.IsZero(m.ReachingUsedSpace) { // not required
		return nil
	}

	for i := 0; i < len(m.ReachingUsedSpace); i++ {
		if swag.IsZero(m.ReachingUsedSpace[i]) { // not required
			continue
		}

		if m.ReachingUsedSpace[i] != nil {
			if err := m.ReachingUsedSpace[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reaching_used_space" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reaching_used_space" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Events) validateReachingUsedTraffic(formats strfmt.Registry) error {
	if swag.IsZero(m.ReachingUsedTraffic) { // not required
		return nil
	}

	for i := 0; i < len(m.ReachingUsedTraffic); i++ {
		if swag.IsZero(m.ReachingUsedTraffic[i]) { // not required
			continue
		}

		if m.ReachingUsedTraffic[i] != nil {
			if err := m.ReachingUsedTraffic[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reaching_used_traffic" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reaching_used_traffic" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this events based on the context it is used
func (m *Events) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReachingFilesQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReachingUsedRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReachingUsedSpace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReachingUsedTraffic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Events) contextValidateReachingFilesQuantity(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReachingFilesQuantity); i++ {

		if m.ReachingFilesQuantity[i] != nil {
			if err := m.ReachingFilesQuantity[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reaching_files_quantity" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reaching_files_quantity" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Events) contextValidateReachingUsedRequests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReachingUsedRequests); i++ {

		if m.ReachingUsedRequests[i] != nil {
			if err := m.ReachingUsedRequests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reaching_used_requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reaching_used_requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Events) contextValidateReachingUsedSpace(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReachingUsedSpace); i++ {

		if m.ReachingUsedSpace[i] != nil {
			if err := m.ReachingUsedSpace[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reaching_used_space" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reaching_used_space" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Events) contextValidateReachingUsedTraffic(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReachingUsedTraffic); i++ {

		if m.ReachingUsedTraffic[i] != nil {
			if err := m.ReachingUsedTraffic[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reaching_used_traffic" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reaching_used_traffic" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Events) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Events) UnmarshalBinary(b []byte) error {
	var res Events
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
