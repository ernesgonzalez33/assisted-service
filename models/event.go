// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Event event
//
// swagger:model event
type Event struct {

	// Unique identifier of the object this event relates to.
	// Required: true
	// Format: uuid
	EntityID *strfmt.UUID `json:"entity_id" gorm:"index"`

	// event time
	// Required: true
	// Format: date-time
	EventTime *strfmt.DateTime `json:"event_time" gorm:"type:datetime"`

	// message
	// Required: true
	Message *string `json:"message" gorm:"type:varchar(4096)"`

	// Unique identifier for the request that caused this event to occure
	// Format: uuid
	RequestID strfmt.UUID `json:"request_id,omitempty"`

	// severity
	// Required: true
	// Enum: [info warning error critical]
	Severity *string `json:"severity"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateEntityID(formats strfmt.Registry) error {

	if err := validate.Required("entity_id", "body", m.EntityID); err != nil {
		return err
	}

	if err := validate.FormatOf("entity_id", "body", "uuid", m.EntityID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateEventTime(formats strfmt.Registry) error {

	if err := validate.Required("event_time", "body", m.EventTime); err != nil {
		return err
	}

	if err := validate.FormatOf("event_time", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateRequestID(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestID) { // not required
		return nil
	}

	if err := validate.FormatOf("request_id", "body", "uuid", m.RequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

var eventTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["info","warning","error","critical"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventTypeSeverityPropEnum = append(eventTypeSeverityPropEnum, v)
	}
}

const (

	// EventSeverityInfo captures enum value "info"
	EventSeverityInfo string = "info"

	// EventSeverityWarning captures enum value "warning"
	EventSeverityWarning string = "warning"

	// EventSeverityError captures enum value "error"
	EventSeverityError string = "error"

	// EventSeverityCritical captures enum value "critical"
	EventSeverityCritical string = "critical"
)

// prop value enum
func (m *Event) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, eventTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Event) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", *m.Severity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
