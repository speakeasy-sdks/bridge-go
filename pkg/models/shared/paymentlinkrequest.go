// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type PaymentLinkRequestCountryEnum string

const (
	PaymentLinkRequestCountryEnumFr PaymentLinkRequestCountryEnum = "fr"
	PaymentLinkRequestCountryEnumEs PaymentLinkRequestCountryEnum = "es"
	PaymentLinkRequestCountryEnumDe PaymentLinkRequestCountryEnum = "de"
	PaymentLinkRequestCountryEnumGb PaymentLinkRequestCountryEnum = "gb"
)

func (e PaymentLinkRequestCountryEnum) ToPointer() *PaymentLinkRequestCountryEnum {
	return &e
}

func (e *PaymentLinkRequestCountryEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "fr":
		fallthrough
	case "es":
		fallthrough
	case "de":
		fallthrough
	case "gb":
		*e = PaymentLinkRequestCountryEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentLinkRequestCountryEnum: %s", s)
	}
}

type PaymentLinkRequest struct {
	BankID          *int64                         `json:"bank_id,omitempty"`
	CallbackURL     *string                        `json:"callback_url,omitempty"`
	ClientReference *string                        `json:"client_reference,omitempty"`
	Country         *PaymentLinkRequestCountryEnum `json:"country,omitempty"`
	ExpiredDate     *time.Time                     `json:"expired_date,omitempty"`
	Transactions    []Transaction                  `json:"transactions"`
	User            User                           `json:"user"`
}
