// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// FieldErrorTag - The type of error returned
type FieldErrorTag string

const (
	FieldErrorTagInvalidInputParameter FieldErrorTag = "invalid_input_parameter"
)

func (e FieldErrorTag) ToPointer() *FieldErrorTag {
	return &e
}
func (e *FieldErrorTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "invalid_input_parameter":
		*e = FieldErrorTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FieldErrorTag: %v", v)
	}
}

// FieldError - An error that pertains to validation of a specific field in the request.
type FieldError struct {
	// The type of error returned
	DotTag FieldErrorTag `json:".tag"`
	// A human-readable error message, which might include information specific to the request that was made.
	Message string `json:"message"`
	// The field (in its hierarchical form) that is failing validation.
	Field string `json:"field"`
}

func (o *FieldError) GetDotTag() FieldErrorTag {
	if o == nil {
		return FieldErrorTag("")
	}
	return o.DotTag
}

func (o *FieldError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *FieldError) GetField() string {
	if o == nil {
		return ""
	}
	return o.Field
}
