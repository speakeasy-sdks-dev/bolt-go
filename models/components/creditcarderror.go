// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// CreditCardErrorTag - The type of error returned
type CreditCardErrorTag string

const (
	CreditCardErrorTagDeclined                CreditCardErrorTag = "declined"
	CreditCardErrorTagDeclinedInvalidAmount   CreditCardErrorTag = "declined_invalid_amount"
	CreditCardErrorTagDeclinedInvalidCvv      CreditCardErrorTag = "declined_invalid_cvv"
	CreditCardErrorTagDeclinedInvalidMerchant CreditCardErrorTag = "declined_invalid_merchant"
	CreditCardErrorTagDeclinedInvalidNumber   CreditCardErrorTag = "declined_invalid_number"
	CreditCardErrorTagDeclinedExpired         CreditCardErrorTag = "declined_expired"
	CreditCardErrorTagDeclinedCallIssuer      CreditCardErrorTag = "declined_call_issuer"
	CreditCardErrorTagDeclinedUnsupported     CreditCardErrorTag = "declined_unsupported"
)

func (e CreditCardErrorTag) ToPointer() *CreditCardErrorTag {
	return &e
}
func (e *CreditCardErrorTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "declined":
		fallthrough
	case "declined_invalid_amount":
		fallthrough
	case "declined_invalid_cvv":
		fallthrough
	case "declined_invalid_merchant":
		fallthrough
	case "declined_invalid_number":
		fallthrough
	case "declined_expired":
		fallthrough
	case "declined_call_issuer":
		fallthrough
	case "declined_unsupported":
		*e = CreditCardErrorTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreditCardErrorTag: %v", v)
	}
}

type CreditCardError struct {
	// The type of error returned
	DotTag CreditCardErrorTag `json:".tag"`
	// A human-readable error message, which might include information specific to the request that was made.
	Message string `json:"message"`
}

func (o *CreditCardError) GetDotTag() CreditCardErrorTag {
	if o == nil {
		return CreditCardErrorTag("")
	}
	return o.DotTag
}

func (o *CreditCardError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}
