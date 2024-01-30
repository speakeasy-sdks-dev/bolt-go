// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

// TestingCreditCardGetResponseBody - An error has occurred, and further details are contained in the response
type TestingCreditCardGetResponseBody struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &TestingCreditCardGetResponseBody{}

func (e *TestingCreditCardGetResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
