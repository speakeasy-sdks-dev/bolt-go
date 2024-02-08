// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type PaymentMethodAffirmTag string

const (
	PaymentMethodAffirmTagAffirm PaymentMethodAffirmTag = "affirm"
)

func (e PaymentMethodAffirmTag) ToPointer() *PaymentMethodAffirmTag {
	return &e
}

func (e *PaymentMethodAffirmTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "affirm":
		*e = PaymentMethodAffirmTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodAffirmTag: %v", v)
	}
}

type PaymentMethodAffirmOutput struct {
	DotTag PaymentMethodAffirmTag `json:".tag"`
}

func (o *PaymentMethodAffirmOutput) GetDotTag() PaymentMethodAffirmTag {
	if o == nil {
		return PaymentMethodAffirmTag("")
	}
	return o.DotTag
}

type PaymentMethodAffirm struct {
	DotTag PaymentMethodAffirmTag `json:".tag"`
	// Return URL to return to after payment completion in Affirm.
	ReturnURL string `json:"return_url"`
}

func (o *PaymentMethodAffirm) GetDotTag() PaymentMethodAffirmTag {
	if o == nil {
		return PaymentMethodAffirmTag("")
	}
	return o.DotTag
}

func (o *PaymentMethodAffirm) GetReturnURL() string {
	if o == nil {
		return ""
	}
	return o.ReturnURL
}