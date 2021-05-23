// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FioPerfCheckRequest fio perf check request
//
// swagger:model fio_perf_check_request
type FioPerfCheckRequest struct {

	// The maximal fdatasync duration in ms that is considered acceptable.
	// Required: true
	DurationThresholdMs *int64 `json:"duration_threshold_ms"`

	// Exit code to return in case of an error.
	// Required: true
	ExitCode *int64 `json:"exit_code"`

	// --filename argument for fio (expects a file or a block device path).
	// Required: true
	Path *string `json:"path"`
}

// Validate validates this fio perf check request
func (m *FioPerfCheckRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDurationThresholdMs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExitCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FioPerfCheckRequest) validateDurationThresholdMs(formats strfmt.Registry) error {

	if err := validate.Required("duration_threshold_ms", "body", m.DurationThresholdMs); err != nil {
		return err
	}

	return nil
}

func (m *FioPerfCheckRequest) validateExitCode(formats strfmt.Registry) error {

	if err := validate.Required("exit_code", "body", m.ExitCode); err != nil {
		return err
	}

	return nil
}

func (m *FioPerfCheckRequest) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FioPerfCheckRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FioPerfCheckRequest) UnmarshalBinary(b []byte) error {
	var res FioPerfCheckRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}