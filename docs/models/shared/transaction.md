# Transaction


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Amount`                                           | *float32*                                          | :heavy_check_mark:                                 | N/A                                                | 120.98                                             |
| `Beneficiary`                                      | [*Beneficiary](../../models/shared/beneficiary.md) | :heavy_minus_sign:                                 | N/A                                                |                                                    |
| `Currency`                                         | *string*                                           | :heavy_check_mark:                                 | N/A                                                | EUR                                                |
| `EndToEndID`                                       | **string*                                          | :heavy_minus_sign:                                 | N/A                                                | E2E_ID-1234                                        |
| `ExecutionDate`                                    | [*time.Time](https://pkg.go.dev/time#Time)         | :heavy_minus_sign:                                 | N/A                                                | 2021-07-24T22:00:00.000Z                           |
| `Label`                                            | *string*                                           | :heavy_check_mark:                                 | N/A                                                | Refund 123456                                      |