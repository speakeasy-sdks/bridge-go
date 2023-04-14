<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/bridge-go"
    "github.com/speakeasy-sdks/bridge-go/pkg/models/shared"
    "github.com/speakeasy-sdks/bridge-go/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            ClientID: sdk.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()    
    req := shared.PaymentLinkRequest{
        BankID: 548814,
        CallbackURL: "provident",
        ClientReference: "ABCDE_FG-HI_12345",
        Country: "de",
        ExpiredDate: "2021-07-24T22:00:00.000Z",
        Transactions: []shared.Transaction{
            shared.Transaction{
                Amount: 120.98,
                Beneficiary: &shared.Beneficiary{
                    CompanyName: "Acme Inc.",
                    FirstName: "John",
                    Iban: "GB29 NWBK 6016 1331 9268 19",
                    LastName: "Doe",
                },
                Currency: "EUR",
                EndToEndID: "E2E_ID-1234",
                ExecutionDate: "2021-07-24T22:00:00.000Z",
                Label: "Refund 123456",
            },
            shared.Transaction{
                Amount: 120.98,
                Beneficiary: &shared.Beneficiary{
                    CompanyName: "Acme Inc.",
                    FirstName: "John",
                    Iban: "GB29 NWBK 6016 1331 9268 19",
                    LastName: "Doe",
                },
                Currency: "EUR",
                EndToEndID: "E2E_ID-1234",
                ExecutionDate: "2021-07-24T22:00:00.000Z",
                Label: "Refund 123456",
            },
            shared.Transaction{
                Amount: 120.98,
                Beneficiary: &shared.Beneficiary{
                    CompanyName: "Acme Inc.",
                    FirstName: "John",
                    Iban: "GB29 NWBK 6016 1331 9268 19",
                    LastName: "Doe",
                },
                Currency: "EUR",
                EndToEndID: "E2E_ID-1234",
                ExecutionDate: "2021-07-24T22:00:00.000Z",
                Label: "Refund 123456",
            },
            shared.Transaction{
                Amount: 120.98,
                Beneficiary: &shared.Beneficiary{
                    CompanyName: "Acme Inc.",
                    FirstName: "John",
                    Iban: "GB29 NWBK 6016 1331 9268 19",
                    LastName: "Doe",
                },
                Currency: "EUR",
                EndToEndID: "E2E_ID-1234",
                ExecutionDate: "2021-07-24T22:00:00.000Z",
                Label: "Refund 123456",
            },
        },
        User: shared.User{
            ExternalReference: "REF-USER-1234_AZ",
            FirstName: "Thomas",
            LastName: "Pichet",
        },
    }

    res, err := s.PaymentLinks.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePaymentLink200ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->