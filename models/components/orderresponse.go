// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type OrderResponse struct {
	ID          string      `json:"id"`
	Transaction Transaction `json:"transaction"`
}

func (o *OrderResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *OrderResponse) GetTransaction() Transaction {
	if o == nil {
		return Transaction{}
	}
	return o.Transaction
}
