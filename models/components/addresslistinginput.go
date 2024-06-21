// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// AddressListingInput - An address saved on an account, i.e. a physical address plus any additional account-specific metadata.
type AddressListingInput struct {
	// The company associated with this address.
	Company *string `json:"company,omitempty"`
	// The country (in its ISO 3166 alpha-2 format) associated with this address.
	CountryCode CountryCode `json:"country_code"`
	// The email address associated with this address.
	Email *string `json:"email,omitempty"`
	// The first name of the person associated with this address.
	FirstName string `json:"first_name"`
	// The last name of the person associated with this address.
	LastName string `json:"last_name"`
	// The locality (e.g. city, town, etc...) associated with this address.
	Locality string `json:"locality"`
	// The phone number associated with this address.
	Phone *string `json:"phone,omitempty"`
	// The postal code associated with this address.
	PostalCode string `json:"postal_code"`
	// The region or administrative area (e.g. state, province, county, etc...) associated with this address.
	Region *string `json:"region,omitempty"`
	// The street address associated with this address.
	StreetAddress1 string `json:"street_address1"`
	// Any additional, optional, street address information associated with this address.
	StreetAddress2 *string `json:"street_address2,omitempty"`
}

func (o *AddressListingInput) GetCompany() *string {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *AddressListingInput) GetCountryCode() CountryCode {
	if o == nil {
		return CountryCode("")
	}
	return o.CountryCode
}

func (o *AddressListingInput) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AddressListingInput) GetFirstName() string {
	if o == nil {
		return ""
	}
	return o.FirstName
}

func (o *AddressListingInput) GetLastName() string {
	if o == nil {
		return ""
	}
	return o.LastName
}

func (o *AddressListingInput) GetLocality() string {
	if o == nil {
		return ""
	}
	return o.Locality
}

func (o *AddressListingInput) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *AddressListingInput) GetPostalCode() string {
	if o == nil {
		return ""
	}
	return o.PostalCode
}

func (o *AddressListingInput) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *AddressListingInput) GetStreetAddress1() string {
	if o == nil {
		return ""
	}
	return o.StreetAddress1
}

func (o *AddressListingInput) GetStreetAddress2() *string {
	if o == nil {
		return nil
	}
	return o.StreetAddress2
}
