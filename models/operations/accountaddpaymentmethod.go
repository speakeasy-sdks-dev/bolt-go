// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/BoltApp/bolt-go/models/components"
	"net/http"
)

type AccountAddPaymentMethodRequest struct {
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey string                        `header:"style=simple,explode=false,name=X-Publishable-Key"`
	PaymentMethod   components.PaymentMethodInput `request:"mediaType=application/json"`
}

func (o *AccountAddPaymentMethodRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

func (o *AccountAddPaymentMethodRequest) GetPaymentMethod() components.PaymentMethodInput {
	if o == nil {
		return components.PaymentMethodInput{}
	}
	return o.PaymentMethod
}

func (o *AccountAddPaymentMethodRequest) GetPaymentMethodCreditCard() *components.PaymentMethodCreditCardInput {
	return o.GetPaymentMethod().PaymentMethodCreditCardInput
}

func (o *AccountAddPaymentMethodRequest) GetPaymentMethodPaypal() *components.PaymentMethodPaypal {
	return o.GetPaymentMethod().PaymentMethodPaypal
}

func (o *AccountAddPaymentMethodRequest) GetPaymentMethodAffirm() *components.PaymentMethodAffirm {
	return o.GetPaymentMethod().PaymentMethodAffirm
}

func (o *AccountAddPaymentMethodRequest) GetPaymentMethodAfterpay() *components.PaymentMethodAfterpay {
	return o.GetPaymentMethod().PaymentMethodAfterpay
}

func (o *AccountAddPaymentMethodRequest) GetPaymentMethodKlarna() *components.PaymentMethodKlarna {
	return o.GetPaymentMethod().PaymentMethodKlarna
}

func (o *AccountAddPaymentMethodRequest) GetPaymentMethodKlarnaAccount() *components.PaymentMethodKlarnaAccount {
	return o.GetPaymentMethod().PaymentMethodKlarnaAccount
}

func (o *AccountAddPaymentMethodRequest) GetPaymentMethodKlarnaPaynow() *components.PaymentMethodKlarnaPaynow {
	return o.GetPaymentMethod().PaymentMethodKlarnaPaynow
}

type AccountAddPaymentMethodResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The payment method was successfully added
	PaymentMethod *components.PaymentMethod
}

func (o *AccountAddPaymentMethodResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AccountAddPaymentMethodResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AccountAddPaymentMethodResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AccountAddPaymentMethodResponse) GetPaymentMethod() *components.PaymentMethod {
	if o == nil {
		return nil
	}
	return o.PaymentMethod
}

func (o *AccountAddPaymentMethodResponse) GetPaymentMethodCreditCard() *components.PaymentMethodCreditCard {
	if v := o.GetPaymentMethod(); v != nil {
		return v.PaymentMethodCreditCard
	}
	return nil
}

func (o *AccountAddPaymentMethodResponse) GetPaymentMethodPaypal() *components.PaymentMethodPaypalOutput {
	if v := o.GetPaymentMethod(); v != nil {
		return v.PaymentMethodPaypalOutput
	}
	return nil
}

func (o *AccountAddPaymentMethodResponse) GetPaymentMethodAffirm() *components.PaymentMethodAffirmOutput {
	if v := o.GetPaymentMethod(); v != nil {
		return v.PaymentMethodAffirmOutput
	}
	return nil
}

func (o *AccountAddPaymentMethodResponse) GetPaymentMethodAfterpay() *components.PaymentMethodAfterpayOutput {
	if v := o.GetPaymentMethod(); v != nil {
		return v.PaymentMethodAfterpayOutput
	}
	return nil
}

func (o *AccountAddPaymentMethodResponse) GetPaymentMethodKlarna() *components.PaymentMethodKlarnaOutput {
	if v := o.GetPaymentMethod(); v != nil {
		return v.PaymentMethodKlarnaOutput
	}
	return nil
}

func (o *AccountAddPaymentMethodResponse) GetPaymentMethodKlarnaAccount() *components.PaymentMethodKlarnaAccountOutput {
	if v := o.GetPaymentMethod(); v != nil {
		return v.PaymentMethodKlarnaAccountOutput
	}
	return nil
}

func (o *AccountAddPaymentMethodResponse) GetPaymentMethodKlarnaPaynow() *components.PaymentMethodKlarnaPaynowOutput {
	if v := o.GetPaymentMethod(); v != nil {
		return v.PaymentMethodKlarnaPaynowOutput
	}
	return nil
}
