// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/bridge-go/pkg/models/shared"
	"net/http"
	"time"
)

// ListPaymentLinksStatusEnum - Filter payment links by status
type ListPaymentLinksStatusEnum string

const (
	ListPaymentLinksStatusEnumValid     ListPaymentLinksStatusEnum = "VALID"
	ListPaymentLinksStatusEnumCompleted ListPaymentLinksStatusEnum = "COMPLETED"
	ListPaymentLinksStatusEnumExpired   ListPaymentLinksStatusEnum = "EXPIRED"
	ListPaymentLinksStatusEnumRevoked   ListPaymentLinksStatusEnum = "REVOKED"
)

func (e ListPaymentLinksStatusEnum) ToPointer() *ListPaymentLinksStatusEnum {
	return &e
}

func (e *ListPaymentLinksStatusEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "VALID":
		fallthrough
	case "COMPLETED":
		fallthrough
	case "EXPIRED":
		fallthrough
	case "REVOKED":
		*e = ListPaymentLinksStatusEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPaymentLinksStatusEnum: %v", v)
	}
}

type ListPaymentLinksRequest struct {
	// Cursor pointing to the start of the desired set
	After *string `queryParam:"style=form,explode=true,name=after"`
	// Number of records to return
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Limit to payment links created after the specified date
	Since *time.Time `queryParam:"style=form,explode=true,name=since"`
	// Filter payment links by status
	Status *ListPaymentLinksStatusEnum `queryParam:"style=form,explode=true,name=status"`
	// Limit to transactions created before the specified date
	Until *time.Time `queryParam:"style=form,explode=true,name=until"`
}

type ListPaymentLinksResponse struct {
	ContentType string
	// Invalid body content
	InvalidBodyContent *shared.InvalidBodyContent
	// OK
	PaymentLinksResponse *shared.PaymentLinksResponse
	StatusCode           int
	RawResponse          *http.Response
}
