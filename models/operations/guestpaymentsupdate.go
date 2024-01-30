// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/BoltApp/bolt-go/models/components"
	"net/http"
)

type GuestPaymentsUpdateSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *GuestPaymentsUpdateSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

type GuestPaymentsUpdateRequest struct {
	// The ID of the guest payment to update
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The publicly viewable identifier used to identify a merchant division.
	XPublishableKey      string                          `header:"style=simple,explode=false,name=X-Publishable-Key"`
	PaymentUpdateRequest components.PaymentUpdateRequest `request:"mediaType=application/json"`
}

func (o *GuestPaymentsUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GuestPaymentsUpdateRequest) GetXPublishableKey() string {
	if o == nil {
		return ""
	}
	return o.XPublishableKey
}

func (o *GuestPaymentsUpdateRequest) GetPaymentUpdateRequest() components.PaymentUpdateRequest {
	if o == nil {
		return components.PaymentUpdateRequest{}
	}
	return o.PaymentUpdateRequest
}

type GuestPaymentsUpdateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The pending payment was successfully updated
	PaymentResponse *components.PaymentResponse
}

func (o *GuestPaymentsUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GuestPaymentsUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GuestPaymentsUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GuestPaymentsUpdateResponse) GetPaymentResponse() *components.PaymentResponse {
	if o == nil {
		return nil
	}
	return o.PaymentResponse
}

func (o *GuestPaymentsUpdateResponse) GetPaymentResponseFinalized() *components.PaymentResponseFinalized {
	if v := o.GetPaymentResponse(); v != nil {
		return v.PaymentResponseFinalized
	}
	return nil
}

func (o *GuestPaymentsUpdateResponse) GetPaymentResponsePending() *components.PaymentResponsePending {
	if v := o.GetPaymentResponse(); v != nil {
		return v.PaymentResponsePending
	}
	return nil
}
