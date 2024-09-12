// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BoltApp/bolt-go/internal/utils"
	"github.com/BoltApp/bolt-go/models/components"
	"net/http"
)

type GuestPaymentsActionResponseBodyType string

const (
	GuestPaymentsActionResponseBodyTypeGenericError GuestPaymentsActionResponseBodyType = "generic-error"
	GuestPaymentsActionResponseBodyTypeFieldError   GuestPaymentsActionResponseBodyType = "field-error"
)

// GuestPaymentsActionResponseBody - An error has occurred, and further details are contained in the response
type GuestPaymentsActionResponseBody struct {
	GenericError *components.GenericError
	FieldError   *components.FieldError

	Type GuestPaymentsActionResponseBodyType

	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
}

var _ error = &GuestPaymentsActionResponseBody{}

func CreateGuestPaymentsActionResponseBodyGenericError(genericError components.GenericError) GuestPaymentsActionResponseBody {
	typ := GuestPaymentsActionResponseBodyTypeGenericError

	return GuestPaymentsActionResponseBody{
		GenericError: &genericError,
		Type:         typ,
	}
}

func CreateGuestPaymentsActionResponseBodyFieldError(fieldError components.FieldError) GuestPaymentsActionResponseBody {
	typ := GuestPaymentsActionResponseBodyTypeFieldError

	return GuestPaymentsActionResponseBody{
		FieldError: &fieldError,
		Type:       typ,
	}
}

func (u *GuestPaymentsActionResponseBody) UnmarshalJSON(data []byte) error {

	var genericError components.GenericError = components.GenericError{}
	if err := utils.UnmarshalJSON(data, &genericError, "", true, true); err == nil {
		u.GenericError = &genericError
		u.Type = GuestPaymentsActionResponseBodyTypeGenericError
		return nil
	}

	var fieldError components.FieldError = components.FieldError{}
	if err := utils.UnmarshalJSON(data, &fieldError, "", true, true); err == nil {
		u.FieldError = &fieldError
		u.Type = GuestPaymentsActionResponseBodyTypeFieldError
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GuestPaymentsActionResponseBody", string(data))
}

func (u GuestPaymentsActionResponseBody) MarshalJSON() ([]byte, error) {
	if u.GenericError != nil {
		return utils.MarshalJSON(u.GenericError, "", true)
	}

	if u.FieldError != nil {
		return utils.MarshalJSON(u.FieldError, "", true)
	}

	return nil, errors.New("could not marshal union type GuestPaymentsActionResponseBody: all fields are null")
}

func (u GuestPaymentsActionResponseBody) Error() string {
	switch u.Type {
	case GuestPaymentsActionResponseBodyTypeGenericError:
		data, _ := json.Marshal(u.GenericError)
		return string(data)
	case GuestPaymentsActionResponseBodyTypeFieldError:
		data, _ := json.Marshal(u.FieldError)
		return string(data)
	default:
		return "unknown error"
	}
}
