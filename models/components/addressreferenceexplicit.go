// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// AddressReferenceExplicitTag - The type of address reference
type AddressReferenceExplicitTag string

const (
	AddressReferenceExplicitTagExplicit AddressReferenceExplicitTag = "explicit"
)

func (e AddressReferenceExplicitTag) ToPointer() *AddressReferenceExplicitTag {
	return &e
}
func (e *AddressReferenceExplicitTag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "explicit":
		*e = AddressReferenceExplicitTag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddressReferenceExplicitTag: %v", v)
	}
}

// AddressReferenceExplicit - A physical address.
type AddressReferenceExplicit struct {
	// The type of address reference
	DotTag AddressReferenceExplicitTag `json:".tag"`
	// The address's unique identifier.
	ID *string `json:"id,omitempty"`
	// The first name of the person associated with this address.
	FirstName string `json:"first_name"`
	// The last name of the person associated with this address.
	LastName string `json:"last_name"`
	// The company associated with this address.
	Company *string `json:"company,omitempty"`
	// The street address associated with this address.
	StreetAddress1 string `json:"street_address1"`
	// Any additional, optional, street address information associated with this address.
	StreetAddress2 *string `json:"street_address2,omitempty"`
	// The locality (e.g. city, town, etc...) associated with this address.
	Locality string `json:"locality"`
	// The postal code associated with this address.
	PostalCode string `json:"postal_code"`
	// The region or administrative area (e.g. state, province, county, etc...) associated with this address.
	Region *string `json:"region,omitempty"`
	// The country (in its ISO 3166 alpha-2 format) associated with this address.
	CountryCode CountryCode `json:"country_code"`
	// The email address associated with this address.
	Email *string `json:"email,omitempty"`
	// The phone number associated with this address.
	Phone *string `json:"phone,omitempty"`
}

func (o *AddressReferenceExplicit) GetDotTag() AddressReferenceExplicitTag {
	if o == nil {
		return AddressReferenceExplicitTag("")
	}
	return o.DotTag
}

func (o *AddressReferenceExplicit) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AddressReferenceExplicit) GetFirstName() string {
	if o == nil {
		return ""
	}
	return o.FirstName
}

func (o *AddressReferenceExplicit) GetLastName() string {
	if o == nil {
		return ""
	}
	return o.LastName
}

func (o *AddressReferenceExplicit) GetCompany() *string {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *AddressReferenceExplicit) GetStreetAddress1() string {
	if o == nil {
		return ""
	}
	return o.StreetAddress1
}

func (o *AddressReferenceExplicit) GetStreetAddress2() *string {
	if o == nil {
		return nil
	}
	return o.StreetAddress2
}

func (o *AddressReferenceExplicit) GetLocality() string {
	if o == nil {
		return ""
	}
	return o.Locality
}

func (o *AddressReferenceExplicit) GetPostalCode() string {
	if o == nil {
		return ""
	}
	return o.PostalCode
}

func (o *AddressReferenceExplicit) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *AddressReferenceExplicit) GetCountryCode() CountryCode {
	if o == nil {
		return CountryCode("")
	}
	return o.CountryCode
}

func (o *AddressReferenceExplicit) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AddressReferenceExplicit) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

// AddressReferenceExplicitInput - A physical address.
type AddressReferenceExplicitInput struct {
	// The type of address reference
	DotTag AddressReferenceExplicitTag `json:".tag"`
	// The first name of the person associated with this address.
	FirstName string `json:"first_name"`
	// The last name of the person associated with this address.
	LastName string `json:"last_name"`
	// The company associated with this address.
	Company *string `json:"company,omitempty"`
	// The street address associated with this address.
	StreetAddress1 string `json:"street_address1"`
	// Any additional, optional, street address information associated with this address.
	StreetAddress2 *string `json:"street_address2,omitempty"`
	// The locality (e.g. city, town, etc...) associated with this address.
	Locality string `json:"locality"`
	// The postal code associated with this address.
	PostalCode string `json:"postal_code"`
	// The region or administrative area (e.g. state, province, county, etc...) associated with this address.
	Region *string `json:"region,omitempty"`
	// The country (in its ISO 3166 alpha-2 format) associated with this address.
	CountryCode CountryCode `json:"country_code"`
	// The email address associated with this address.
	Email *string `json:"email,omitempty"`
	// The phone number associated with this address.
	Phone *string `json:"phone,omitempty"`
}

func (o *AddressReferenceExplicitInput) GetDotTag() AddressReferenceExplicitTag {
	if o == nil {
		return AddressReferenceExplicitTag("")
	}
	return o.DotTag
}

func (o *AddressReferenceExplicitInput) GetFirstName() string {
	if o == nil {
		return ""
	}
	return o.FirstName
}

func (o *AddressReferenceExplicitInput) GetLastName() string {
	if o == nil {
		return ""
	}
	return o.LastName
}

func (o *AddressReferenceExplicitInput) GetCompany() *string {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *AddressReferenceExplicitInput) GetStreetAddress1() string {
	if o == nil {
		return ""
	}
	return o.StreetAddress1
}

func (o *AddressReferenceExplicitInput) GetStreetAddress2() *string {
	if o == nil {
		return nil
	}
	return o.StreetAddress2
}

func (o *AddressReferenceExplicitInput) GetLocality() string {
	if o == nil {
		return ""
	}
	return o.Locality
}

func (o *AddressReferenceExplicitInput) GetPostalCode() string {
	if o == nil {
		return ""
	}
	return o.PostalCode
}

func (o *AddressReferenceExplicitInput) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *AddressReferenceExplicitInput) GetCountryCode() CountryCode {
	if o == nil {
		return CountryCode("")
	}
	return o.CountryCode
}

func (o *AddressReferenceExplicitInput) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AddressReferenceExplicitInput) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}
