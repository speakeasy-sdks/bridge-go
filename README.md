<div align="center">
    <picture>
        <source srcset="https://user-images.githubusercontent.com/6267663/232034887-48fd151e-8ba5-466f-99bf-8c67aeeec0de.png" media="(prefers-color-scheme: dark)" width="100">
        <img src="https://user-images.githubusercontent.com/6267663/232034887-48fd151e-8ba5-466f-99bf-8c67aeeec0de.png" width="100">
    </picture>
    <h1>Bridge Go SDK</h1>
   <p>Leverage bank transfers and improve payment experience with Payment Links.</p>
   <a href="https://docs.bridgeapi.io/docs"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=5444e4&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/bridge-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/bridge-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/bridge-go/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/bridge-go?sort=semver&style=for-the-badge" /></a>
</div>


<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/bridge-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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
            ClientID: bride.String("YOUR_API_KEY_HERE"),
        }),
    )

    ctx := context.Background()
    res, err := s.PaymentLinks.Create(ctx, shared.PaymentLinkRequest{
        BankID: bride.Int64(548814),
        CallbackURL: bride.String("provident"),
        ClientReference: bride.String("ABCDE_FG-HI_12345"),
        Country: shared.PaymentLinkRequestCountryEnumDe.ToPointer(),
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [PaymentLinks](docs/paymentlinks/README.md)

* [Create](docs/paymentlinks/README.md#create) - Create a payment link
* [Get](docs/paymentlinks/README.md#get) - Retrieve payment link information by ID
* [List](docs/paymentlinks/README.md#list) - Retrieve a list of payment links
* [Revoke](docs/paymentlinks/README.md#revoke) - Revoke a payment link
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
