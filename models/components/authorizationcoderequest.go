// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// GrantType - The type of OAuth 2.0 grant being utilized.
type GrantType string

const (
	GrantTypeAuthorizationCode GrantType = "authorization_code"
)

func (e GrantType) ToPointer() *GrantType {
	return &e
}
func (e *GrantType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "authorization_code":
		*e = GrantType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GrantType: %v", v)
	}
}

type Scope string

const (
	ScopeBoltAccountManage Scope = "bolt.account.manage"
	ScopeBoltAccountView   Scope = "bolt.account.view"
	ScopeOpenid            Scope = "openid"
)

func (e Scope) ToPointer() *Scope {
	return &e
}
func (e *Scope) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bolt.account.manage":
		fallthrough
	case "bolt.account.view":
		fallthrough
	case "openid":
		*e = Scope(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Scope: %v", v)
	}
}

type AuthorizationCodeRequest struct {
	// Fetched value using OTP value from the Authorization Modal.
	Code string `form:"name=code"`
	// The type of OAuth 2.0 grant being utilized.
	GrantType GrantType `form:"name=grant_type"`
	// The OAuth client ID, which corresponds to the merchant publishable key, which can be retrieved
	// in the Merchant Dashboard.
	//
	ClientID string `form:"name=client_id"`
	// The OAuth client secret, which corresponds the merchant API key, which can be retrieved in the
	// Merchant Dashboard.
	//
	ClientSecret string `form:"name=client_secret"`
	// The requested scopes. If the request is successful, the OAuth client will be able to perform operations requiring these scopes.
	//
	Scope []Scope `form:"name=scope"`
	// A randomly generated string sent along with an authorization code. This must be included, if provided,
	// in order to prevent CSRF attacks. used to prevent CSRF attacks.
	//
	State *string `form:"name=state"`
}

func (o *AuthorizationCodeRequest) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *AuthorizationCodeRequest) GetGrantType() GrantType {
	if o == nil {
		return GrantType("")
	}
	return o.GrantType
}

func (o *AuthorizationCodeRequest) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *AuthorizationCodeRequest) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *AuthorizationCodeRequest) GetScope() []Scope {
	if o == nil {
		return []Scope{}
	}
	return o.Scope
}

func (o *AuthorizationCodeRequest) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}
