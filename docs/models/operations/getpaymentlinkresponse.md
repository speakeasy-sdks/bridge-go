# GetPaymentLinkResponse


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ContentType`                                                             | *string*                                                                  | :heavy_check_mark:                                                        | N/A                                                                       |
| `InvalidBodyContent`                                                      | [*shared.InvalidBodyContent](../../models/shared/invalidbodycontent.md)   | :heavy_minus_sign:                                                        | Invalid body content                                                      |
| `PaymentLinkResponse`                                                     | [*shared.PaymentLinkResponse](../../models/shared/paymentlinkresponse.md) | :heavy_minus_sign:                                                        | OK                                                                        |
| `StatusCode`                                                              | *int*                                                                     | :heavy_check_mark:                                                        | N/A                                                                       |
| `RawResponse`                                                             | [*http.Response](https://pkg.go.dev/net/http#Response)                    | :heavy_minus_sign:                                                        | N/A                                                                       |