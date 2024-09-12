// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type PaymentMethodPaypalTag string

const (
	PaymentMethodPaypalTagPaypal PaymentMethodPaypalTag = "paypal"
)

func (e PaymentMethodPaypalTag) ToPointer() *PaymentMethodPaypalTag {
	return &e
}
func (e *PaymentMethodPaypalTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "paypal":
		*e = PaymentMethodPaypalTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodPaypalTag: %v", v)
	}
}

type PaymentMethodPaypalOutput struct {
	DotTag PaymentMethodPaypalTag `json:".tag"`
}

func (o *PaymentMethodPaypalOutput) GetDotTag() PaymentMethodPaypalTag {
	if o == nil {
		return PaymentMethodPaypalTag("")
	}
	return o.DotTag
}

type PaymentMethodPaypal struct {
	DotTag PaymentMethodPaypalTag `json:".tag"`
	// Redirect URL for successful PayPal transaction.
	SuccessURL string `json:"success_url"`
	// Redirect URL for canceled PayPal transaction.
	CancelURL string `json:"cancel_url"`
}

func (o *PaymentMethodPaypal) GetDotTag() PaymentMethodPaypalTag {
	if o == nil {
		return PaymentMethodPaypalTag("")
	}
	return o.DotTag
}

func (o *PaymentMethodPaypal) GetSuccessURL() string {
	if o == nil {
		return ""
	}
	return o.SuccessURL
}

func (o *PaymentMethodPaypal) GetCancelURL() string {
	if o == nil {
		return ""
	}
	return o.CancelURL
}
