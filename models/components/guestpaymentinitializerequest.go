// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type GuestPaymentInitializeRequest struct {
	Cart          Cart                `json:"cart"`
	PaymentMethod PaymentMethodInput  `json:"payment_method"`
	Profile       ProfileCreationData `json:"profile"`
}

func (o *GuestPaymentInitializeRequest) GetCart() Cart {
	if o == nil {
		return Cart{}
	}
	return o.Cart
}

func (o *GuestPaymentInitializeRequest) GetPaymentMethod() PaymentMethodInput {
	if o == nil {
		return PaymentMethodInput{}
	}
	return o.PaymentMethod
}

func (o *GuestPaymentInitializeRequest) GetPaymentMethodAffirm() *PaymentMethodAffirm {
	return o.GetPaymentMethod().PaymentMethodAffirm
}

func (o *GuestPaymentInitializeRequest) GetPaymentMethodAfterpay() *PaymentMethodAfterpay {
	return o.GetPaymentMethod().PaymentMethodAfterpay
}

func (o *GuestPaymentInitializeRequest) GetPaymentMethodCreditCard() *PaymentMethodCreditCardInput {
	return o.GetPaymentMethod().PaymentMethodCreditCardInput
}

func (o *GuestPaymentInitializeRequest) GetPaymentMethodKlarna() *PaymentMethodKlarna {
	return o.GetPaymentMethod().PaymentMethodKlarna
}

func (o *GuestPaymentInitializeRequest) GetPaymentMethodKlarnaAccount() *PaymentMethodKlarnaAccount {
	return o.GetPaymentMethod().PaymentMethodKlarnaAccount
}

func (o *GuestPaymentInitializeRequest) GetPaymentMethodKlarnaPaynow() *PaymentMethodKlarnaPaynow {
	return o.GetPaymentMethod().PaymentMethodKlarnaPaynow
}

func (o *GuestPaymentInitializeRequest) GetPaymentMethodPaypal() *PaymentMethodPaypal {
	return o.GetPaymentMethod().PaymentMethodPaypal
}

func (o *GuestPaymentInitializeRequest) GetProfile() ProfileCreationData {
	if o == nil {
		return ProfileCreationData{}
	}
	return o.Profile
}
