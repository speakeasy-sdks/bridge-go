# PaymentLinkResponse

OK


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `ClientReference`                                   | **string*                                           | :heavy_minus_sign:                                  | The client reference of the payment link.           |
| `CreatedAt`                                         | [*time.Time](https://pkg.go.dev/time#Time)          | :heavy_minus_sign:                                  | The creation date and time of the payment link.     |
| `ExpiredAt`                                         | [*time.Time](https://pkg.go.dev/time#Time)          | :heavy_minus_sign:                                  | The expiration date and time of the payment link.   |
| `ID`                                                | **string*                                           | :heavy_minus_sign:                                  | The unique identifier of the payment link.          |
| `Link`                                              | **string*                                           | :heavy_minus_sign:                                  | The URL of the payment link.                        |
| `Status`                                            | **string*                                           | :heavy_minus_sign:                                  | The status of the payment link.                     |
| `Transactions`                                      | [][Transaction](../../models/shared/transaction.md) | :heavy_minus_sign:                                  | The transactions associated with the payment link.  |
| `UpdatedAt`                                         | [*time.Time](https://pkg.go.dev/time#Time)          | :heavy_minus_sign:                                  | The last update date and time of the payment link.  |
| `User`                                              | [*User](../../models/shared/user.md)                | :heavy_minus_sign:                                  | N/A                                                 |