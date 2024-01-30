// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/BoltApp/bolt-go/models/components"
	"net/http"
)

type TestingCreditCardGetSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *TestingCreditCardGetSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

// Type - The expected authorization result when using the generated token for a payment.
type Type string

const (
	TypeApprove Type = "approve"
	TypeDecline Type = "decline"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "approve":
		fallthrough
	case "decline":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type TestingCreditCardGetRequestBody struct {
	// The expected authorization result when using the generated token for a payment.
	Type Type `json:"type"`
}

func (o *TestingCreditCardGetRequestBody) GetType() Type {
	if o == nil {
		return Type("")
	}
	return o.Type
}

type TestingCreditCardGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully generated test credit card details
	TestCreditCard *components.TestCreditCard
}

func (o *TestingCreditCardGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TestingCreditCardGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TestingCreditCardGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TestingCreditCardGetResponse) GetTestCreditCard() *components.TestCreditCard {
	if o == nil {
		return nil
	}
	return o.TestCreditCard
}
