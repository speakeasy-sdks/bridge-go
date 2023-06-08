<!-- Start SDK Example Usage -->
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
        BankID: bride.Int64(548814),
        CallbackURL: bride.String("provident"),
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
<!-- End SDK Example Usage -->