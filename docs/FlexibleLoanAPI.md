# \FlexibleLoanAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFinanceFlexibleLoanAdjustCollateralV5**](FlexibleLoanAPI.md#CreateFinanceFlexibleLoanAdjustCollateralV5) | **Post** /api/v5/finance/flexible-loan/adjust-collateral | POST / Adjust collateral
[**CreateFinanceFlexibleLoanMaxLoanV5**](FlexibleLoanAPI.md#CreateFinanceFlexibleLoanMaxLoanV5) | **Post** /api/v5/finance/flexible-loan/max-loan | POST / Maximum loan amount
[**GetFinanceFlexibleLoanBorrowCurrenciesV5**](FlexibleLoanAPI.md#GetFinanceFlexibleLoanBorrowCurrenciesV5) | **Get** /api/v5/finance/flexible-loan/borrow-currencies | GET / Borrowable currencies
[**GetFinanceFlexibleLoanCollateralAssetsV5**](FlexibleLoanAPI.md#GetFinanceFlexibleLoanCollateralAssetsV5) | **Get** /api/v5/finance/flexible-loan/collateral-assets | GET / Collateral assets
[**GetFinanceFlexibleLoanInterestAccruedV5**](FlexibleLoanAPI.md#GetFinanceFlexibleLoanInterestAccruedV5) | **Get** /api/v5/finance/flexible-loan/interest-accrued | GET / Accrued interest
[**GetFinanceFlexibleLoanLoanHistoryV5**](FlexibleLoanAPI.md#GetFinanceFlexibleLoanLoanHistoryV5) | **Get** /api/v5/finance/flexible-loan/loan-history | GET / Loan history
[**GetFinanceFlexibleLoanLoanInfoV5**](FlexibleLoanAPI.md#GetFinanceFlexibleLoanLoanInfoV5) | **Get** /api/v5/finance/flexible-loan/loan-info | GET / Loan info
[**GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5**](FlexibleLoanAPI.md#GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5) | **Get** /api/v5/finance/flexible-loan/max-collateral-redeem-amount | GET / Maximum collateral redeem amount



## CreateFinanceFlexibleLoanAdjustCollateralV5

> CreateFinanceFlexibleLoanAdjustCollateralV5Resp CreateFinanceFlexibleLoanAdjustCollateralV5(ctx).CreateFinanceFlexibleLoanAdjustCollateralV5Req(createFinanceFlexibleLoanAdjustCollateralV5Req).Execute()

POST / Adjust collateral



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createFinanceFlexibleLoanAdjustCollateralV5Req := *openapiclient.NewCreateFinanceFlexibleLoanAdjustCollateralV5Req("CollateralAmt_example", "CollateralCcy_example", "Type_example") // CreateFinanceFlexibleLoanAdjustCollateralV5Req | The request body for CreateFinanceFlexibleLoanAdjustCollateralV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlexibleLoanAPI.CreateFinanceFlexibleLoanAdjustCollateralV5(context.Background()).CreateFinanceFlexibleLoanAdjustCollateralV5Req(createFinanceFlexibleLoanAdjustCollateralV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlexibleLoanAPI.CreateFinanceFlexibleLoanAdjustCollateralV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFinanceFlexibleLoanAdjustCollateralV5`: CreateFinanceFlexibleLoanAdjustCollateralV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FlexibleLoanAPI.CreateFinanceFlexibleLoanAdjustCollateralV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFinanceFlexibleLoanAdjustCollateralV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFinanceFlexibleLoanAdjustCollateralV5Req** | [**CreateFinanceFlexibleLoanAdjustCollateralV5Req**](CreateFinanceFlexibleLoanAdjustCollateralV5Req.md) | The request body for CreateFinanceFlexibleLoanAdjustCollateralV5 | 

### Return type

[**CreateFinanceFlexibleLoanAdjustCollateralV5Resp**](CreateFinanceFlexibleLoanAdjustCollateralV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFinanceFlexibleLoanMaxLoanV5

> CreateFinanceFlexibleLoanMaxLoanV5Resp CreateFinanceFlexibleLoanMaxLoanV5(ctx).CreateFinanceFlexibleLoanMaxLoanV5Req(createFinanceFlexibleLoanMaxLoanV5Req).Execute()

POST / Maximum loan amount



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createFinanceFlexibleLoanMaxLoanV5Req := *openapiclient.NewCreateFinanceFlexibleLoanMaxLoanV5Req("BorrowCcy_example") // CreateFinanceFlexibleLoanMaxLoanV5Req | The request body for CreateFinanceFlexibleLoanMaxLoanV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlexibleLoanAPI.CreateFinanceFlexibleLoanMaxLoanV5(context.Background()).CreateFinanceFlexibleLoanMaxLoanV5Req(createFinanceFlexibleLoanMaxLoanV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlexibleLoanAPI.CreateFinanceFlexibleLoanMaxLoanV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFinanceFlexibleLoanMaxLoanV5`: CreateFinanceFlexibleLoanMaxLoanV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FlexibleLoanAPI.CreateFinanceFlexibleLoanMaxLoanV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFinanceFlexibleLoanMaxLoanV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFinanceFlexibleLoanMaxLoanV5Req** | [**CreateFinanceFlexibleLoanMaxLoanV5Req**](CreateFinanceFlexibleLoanMaxLoanV5Req.md) | The request body for CreateFinanceFlexibleLoanMaxLoanV5 | 

### Return type

[**CreateFinanceFlexibleLoanMaxLoanV5Resp**](CreateFinanceFlexibleLoanMaxLoanV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceFlexibleLoanBorrowCurrenciesV5

> GetFinanceFlexibleLoanBorrowCurrenciesV5Resp GetFinanceFlexibleLoanBorrowCurrenciesV5(ctx).Execute()

GET / Borrowable currencies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlexibleLoanAPI.GetFinanceFlexibleLoanBorrowCurrenciesV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlexibleLoanAPI.GetFinanceFlexibleLoanBorrowCurrenciesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceFlexibleLoanBorrowCurrenciesV5`: GetFinanceFlexibleLoanBorrowCurrenciesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FlexibleLoanAPI.GetFinanceFlexibleLoanBorrowCurrenciesV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceFlexibleLoanBorrowCurrenciesV5Request struct via the builder pattern


### Return type

[**GetFinanceFlexibleLoanBorrowCurrenciesV5Resp**](GetFinanceFlexibleLoanBorrowCurrenciesV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceFlexibleLoanCollateralAssetsV5

> GetFinanceFlexibleLoanCollateralAssetsV5Resp GetFinanceFlexibleLoanCollateralAssetsV5(ctx).Ccy(ccy).Execute()

GET / Collateral assets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	ccy := "ccy_example" // string | Collateral currency, e.g. `BTC` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlexibleLoanAPI.GetFinanceFlexibleLoanCollateralAssetsV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlexibleLoanAPI.GetFinanceFlexibleLoanCollateralAssetsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceFlexibleLoanCollateralAssetsV5`: GetFinanceFlexibleLoanCollateralAssetsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FlexibleLoanAPI.GetFinanceFlexibleLoanCollateralAssetsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceFlexibleLoanCollateralAssetsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Collateral currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]

### Return type

[**GetFinanceFlexibleLoanCollateralAssetsV5Resp**](GetFinanceFlexibleLoanCollateralAssetsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceFlexibleLoanInterestAccruedV5

> GetFinanceFlexibleLoanInterestAccruedV5Resp GetFinanceFlexibleLoanInterestAccruedV5(ctx).Ccy(ccy).After(after).Before(before).Limit(limit).Execute()

GET / Accrued interest



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	ccy := "ccy_example" // string | Loan currency, e.g. `BTC` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `refId`(not include) (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `refId`(not include) (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlexibleLoanAPI.GetFinanceFlexibleLoanInterestAccruedV5(context.Background()).Ccy(ccy).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlexibleLoanAPI.GetFinanceFlexibleLoanInterestAccruedV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceFlexibleLoanInterestAccruedV5`: GetFinanceFlexibleLoanInterestAccruedV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FlexibleLoanAPI.GetFinanceFlexibleLoanInterestAccruedV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceFlexibleLoanInterestAccruedV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Loan currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;refId&#x60;(not include) | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;refId&#x60;(not include) | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetFinanceFlexibleLoanInterestAccruedV5Resp**](GetFinanceFlexibleLoanInterestAccruedV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceFlexibleLoanLoanHistoryV5

> GetFinanceFlexibleLoanLoanHistoryV5Resp GetFinanceFlexibleLoanLoanHistoryV5(ctx).Type_(type_).After(after).Before(before).Limit(limit).Execute()

GET / Loan history



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	type_ := "type__example" // string | Action type  `borrowed`  `repaid`  `collateral_locked`  `collateral_released`  `forced_repayment_buy`  `forced_repayment_sell`  `forced_liquidation`  `partial_liquidation` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `refId`(not include) (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `refId`(not include) (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlexibleLoanAPI.GetFinanceFlexibleLoanLoanHistoryV5(context.Background()).Type_(type_).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlexibleLoanAPI.GetFinanceFlexibleLoanLoanHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceFlexibleLoanLoanHistoryV5`: GetFinanceFlexibleLoanLoanHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FlexibleLoanAPI.GetFinanceFlexibleLoanLoanHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceFlexibleLoanLoanHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Action type  &#x60;borrowed&#x60;  &#x60;repaid&#x60;  &#x60;collateral_locked&#x60;  &#x60;collateral_released&#x60;  &#x60;forced_repayment_buy&#x60;  &#x60;forced_repayment_sell&#x60;  &#x60;forced_liquidation&#x60;  &#x60;partial_liquidation&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;refId&#x60;(not include) | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;refId&#x60;(not include) | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetFinanceFlexibleLoanLoanHistoryV5Resp**](GetFinanceFlexibleLoanLoanHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceFlexibleLoanLoanInfoV5

> GetFinanceFlexibleLoanLoanInfoV5Resp GetFinanceFlexibleLoanLoanInfoV5(ctx).Execute()

GET / Loan info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlexibleLoanAPI.GetFinanceFlexibleLoanLoanInfoV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlexibleLoanAPI.GetFinanceFlexibleLoanLoanInfoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceFlexibleLoanLoanInfoV5`: GetFinanceFlexibleLoanLoanInfoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FlexibleLoanAPI.GetFinanceFlexibleLoanLoanInfoV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceFlexibleLoanLoanInfoV5Request struct via the builder pattern


### Return type

[**GetFinanceFlexibleLoanLoanInfoV5Resp**](GetFinanceFlexibleLoanLoanInfoV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5

> GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5Resp GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5(ctx).Ccy(ccy).Execute()

GET / Maximum collateral redeem amount



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	ccy := "ccy_example" // string | Collateral currency, e.g. `USDT` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlexibleLoanAPI.GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlexibleLoanAPI.GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5`: GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FlexibleLoanAPI.GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Collateral currency, e.g. &#x60;USDT&#x60; | [default to &quot;&quot;]

### Return type

[**GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5Resp**](GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

