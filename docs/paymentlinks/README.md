# PaymentLinks

## Overview

Payment Links API

### Available Operations

* [Create](#create) - Create a payment link
* [Get](#get) - Retrieve payment link information by ID
* [List](#list) - Retrieve a list of payment links
* [Revoke](#revoke) - Revoke a payment link

## Create

Create a payment link

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bridge-go"
	"github.com/speakeasy-sdks/bridge-go/pkg/models/shared"
	"github.com/speakeasy-sdks/bridge-go/pkg/types"
)

func main() {
    s := bride.New(
        bride.WithSecurity(shared.Security{
            ClientID: bride.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.PaymentLinks.Create(ctx, shared.PaymentLinkRequest{
        BankID: bride.Int64(602763),
        CallbackURL: bride.String("nulla"),
        ClientReference: bride.String("ABCDE_FG-HI_12345"),
        Country: shared.PaymentLinkRequestCountryDe.ToPointer(),
        ExpiredDate: types.MustTimeFromString("2021-07-24T22:00:00.000Z"),
        Transactions: []shared.Transaction{
            shared.Transaction{
                Amount: 120.98,
                Beneficiary: &shared.Beneficiary{
                    CompanyName: bride.String("Acme Inc."),
                    FirstName: bride.String("John"),
                    Iban: "GB29 NWBK 6016 1331 9268 19",
                    LastName: bride.String("Doe"),
                },
                Currency: "EUR",
                EndToEndID: bride.String("E2E_ID-1234"),
                ExecutionDate: types.MustTimeFromString("2021-07-24T22:00:00.000Z"),
                Label: "Refund 123456",
            },
            shared.Transaction{
                Amount: 120.98,
                Beneficiary: &shared.Beneficiary{
                    CompanyName: bride.String("Acme Inc."),
                    FirstName: bride.String("John"),
                    Iban: "GB29 NWBK 6016 1331 9268 19",
                    LastName: bride.String("Doe"),
                },
                Currency: "EUR",
                EndToEndID: bride.String("E2E_ID-1234"),
                ExecutionDate: types.MustTimeFromString("2021-07-24T22:00:00.000Z"),
                Label: "Refund 123456",
            },
            shared.Transaction{
                Amount: 120.98,
                Beneficiary: &shared.Beneficiary{
                    CompanyName: bride.String("Acme Inc."),
                    FirstName: bride.String("John"),
                    Iban: "GB29 NWBK 6016 1331 9268 19",
                    LastName: bride.String("Doe"),
                },
                Currency: "EUR",
                EndToEndID: bride.String("E2E_ID-1234"),
                ExecutionDate: types.MustTimeFromString("2021-07-24T22:00:00.000Z"),
                Label: "Refund 123456",
            },
            shared.Transaction{
                Amount: 120.98,
                Beneficiary: &shared.Beneficiary{
                    CompanyName: bride.String("Acme Inc."),
                    FirstName: bride.String("John"),
                    Iban: "GB29 NWBK 6016 1331 9268 19",
                    LastName: bride.String("Doe"),
                },
                Currency: "EUR",
                EndToEndID: bride.String("E2E_ID-1234"),
                ExecutionDate: types.MustTimeFromString("2021-07-24T22:00:00.000Z"),
                Label: "Refund 123456",
            },
        },
        User: shared.User{
            ExternalReference: bride.String("REF-USER-1234_AZ"),
            FirstName: "Thomas",
            LastName: "Pichet",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePaymentLink200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## Get

Returns the details of a payment link specified by its ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bridge-go"
	"github.com/speakeasy-sdks/bridge-go/pkg/models/operations"
)

func main() {
    s := bride.New(
        bride.WithSecurity(shared.Security{
            ClientID: bride.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.PaymentLinks.Get(ctx, operations.GetPaymentLinkRequest{
        PaymentLinkID: "vel",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentLinkResponse != nil {
        // handle response
    }
}
```

## List

Retrieve a list of payment links

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bridge-go"
	"github.com/speakeasy-sdks/bridge-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bridge-go/pkg/types"
)

func main() {
    s := bride.New(
        bride.WithSecurity(shared.Security{
            ClientID: bride.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.PaymentLinks.List(ctx, operations.ListPaymentLinksRequest{
        After: bride.String("error"),
        Limit: bride.Int64(645894),
        Since: types.MustTimeFromString("2022-07-25T06:44:09.184Z"),
        Status: operations.ListPaymentLinksStatusCompleted.ToPointer(),
        Until: types.MustTimeFromString("2022-10-30T21:34:57.850Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentLinksResponse != nil {
        // handle response
    }
}
```

## Revoke

Revokes the payment link specified by its ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bridge-go"
	"github.com/speakeasy-sdks/bridge-go/pkg/models/operations"
)

func main() {
    s := bride.New(
        bride.WithSecurity(shared.Security{
            ClientID: bride.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.PaymentLinks.Revoke(ctx, operations.RevokePaymentLinkRequest{
        PaymentLinkID: "delectus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
