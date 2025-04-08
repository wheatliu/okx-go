# \SimpleEarnFlexibleAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFinanceSavingsPurchaseRedemptV5**](SimpleEarnFlexibleAPI.md#CreateFinanceSavingsPurchaseRedemptV5) | **Post** /api/v5/finance/savings/purchase-redempt | Only the assets in the funding account can be used for saving.  
[**CreateFinanceSavingsSetLendingRateV5**](SimpleEarnFlexibleAPI.md#CreateFinanceSavingsSetLendingRateV5) | **Post** /api/v5/finance/savings/set-lending-rate | 
[**GetFinanceSavingsBalanceV5**](SimpleEarnFlexibleAPI.md#GetFinanceSavingsBalanceV5) | **Get** /api/v5/finance/savings/balance | 
[**GetFinanceSavingsLendingHistoryV5**](SimpleEarnFlexibleAPI.md#GetFinanceSavingsLendingHistoryV5) | **Get** /api/v5/finance/savings/lending-history | Return data in the past month.  
[**GetFinanceSavingsLendingRateHistoryV5**](SimpleEarnFlexibleAPI.md#GetFinanceSavingsLendingRateHistoryV5) | **Get** /api/v5/finance/savings/lending-rate-history | Authentication is not required for this public endpoint.   Only returned records after December 14, 2021.  
[**GetFinanceSavingsLendingRateSummaryV5**](SimpleEarnFlexibleAPI.md#GetFinanceSavingsLendingRateSummaryV5) | **Get** /api/v5/finance/savings/lending-rate-summary | Authentication is not required for this public endpoint.  



## CreateFinanceSavingsPurchaseRedemptV5

> CreateFinanceSavingsPurchaseRedemptV5Resp CreateFinanceSavingsPurchaseRedemptV5(ctx).CreateFinanceSavingsPurchaseRedemptV5Req(createFinanceSavingsPurchaseRedemptV5Req).Execute()

Only the assets in the funding account can be used for saving.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wheatliu/okx-go"
)

func main() {
	createFinanceSavingsPurchaseRedemptV5Req := *openapiclient.NewCreateFinanceSavingsPurchaseRedemptV5Req("Amt_example", "Ccy_example", "Rate_example", "Side_example") // CreateFinanceSavingsPurchaseRedemptV5Req | The request body for CreateFinanceSavingsPurchaseRedemptV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnFlexibleAPI.CreateFinanceSavingsPurchaseRedemptV5(context.Background()).CreateFinanceSavingsPurchaseRedemptV5Req(createFinanceSavingsPurchaseRedemptV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnFlexibleAPI.CreateFinanceSavingsPurchaseRedemptV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFinanceSavingsPurchaseRedemptV5`: CreateFinanceSavingsPurchaseRedemptV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnFlexibleAPI.CreateFinanceSavingsPurchaseRedemptV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFinanceSavingsPurchaseRedemptV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFinanceSavingsPurchaseRedemptV5Req** | [**CreateFinanceSavingsPurchaseRedemptV5Req**](CreateFinanceSavingsPurchaseRedemptV5Req.md) | The request body for CreateFinanceSavingsPurchaseRedemptV5 | 

### Return type

[**CreateFinanceSavingsPurchaseRedemptV5Resp**](CreateFinanceSavingsPurchaseRedemptV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFinanceSavingsSetLendingRateV5

> CreateFinanceSavingsSetLendingRateV5Resp CreateFinanceSavingsSetLendingRateV5(ctx).CreateFinanceSavingsSetLendingRateV5Req(createFinanceSavingsSetLendingRateV5Req).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wheatliu/okx-go"
)

func main() {
	createFinanceSavingsSetLendingRateV5Req := *openapiclient.NewCreateFinanceSavingsSetLendingRateV5Req("Ccy_example", "Rate_example") // CreateFinanceSavingsSetLendingRateV5Req | The request body for CreateFinanceSavingsSetLendingRateV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnFlexibleAPI.CreateFinanceSavingsSetLendingRateV5(context.Background()).CreateFinanceSavingsSetLendingRateV5Req(createFinanceSavingsSetLendingRateV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnFlexibleAPI.CreateFinanceSavingsSetLendingRateV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFinanceSavingsSetLendingRateV5`: CreateFinanceSavingsSetLendingRateV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnFlexibleAPI.CreateFinanceSavingsSetLendingRateV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFinanceSavingsSetLendingRateV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFinanceSavingsSetLendingRateV5Req** | [**CreateFinanceSavingsSetLendingRateV5Req**](CreateFinanceSavingsSetLendingRateV5Req.md) | The request body for CreateFinanceSavingsSetLendingRateV5 | 

### Return type

[**CreateFinanceSavingsSetLendingRateV5Resp**](CreateFinanceSavingsSetLendingRateV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceSavingsBalanceV5

> GetFinanceSavingsBalanceV5Resp GetFinanceSavingsBalanceV5(ctx).Ccy(ccy).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wheatliu/okx-go"
)

func main() {
	ccy := "ccy_example" // string | Currency, e.g. `BTC` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnFlexibleAPI.GetFinanceSavingsBalanceV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnFlexibleAPI.GetFinanceSavingsBalanceV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceSavingsBalanceV5`: GetFinanceSavingsBalanceV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnFlexibleAPI.GetFinanceSavingsBalanceV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceSavingsBalanceV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]

### Return type

[**GetFinanceSavingsBalanceV5Resp**](GetFinanceSavingsBalanceV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceSavingsLendingHistoryV5

> GetFinanceSavingsLendingHistoryV5Resp GetFinanceSavingsLendingHistoryV5(ctx).Ccy(ccy).After(after).Before(before).Limit(limit).Execute()

Return data in the past month.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wheatliu/okx-go"
)

func main() {
	ccy := "ccy_example" // string | Currency, e.g. `BTC` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ts`, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ts`, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnFlexibleAPI.GetFinanceSavingsLendingHistoryV5(context.Background()).Ccy(ccy).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnFlexibleAPI.GetFinanceSavingsLendingHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceSavingsLendingHistoryV5`: GetFinanceSavingsLendingHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnFlexibleAPI.GetFinanceSavingsLendingHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceSavingsLendingHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ts&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ts&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetFinanceSavingsLendingHistoryV5Resp**](GetFinanceSavingsLendingHistoryV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceSavingsLendingRateHistoryV5

> GetFinanceSavingsLendingRateHistoryV5Resp GetFinanceSavingsLendingRateHistoryV5(ctx).Ccy(ccy).After(after).Before(before).Limit(limit).Execute()

Authentication is not required for this public endpoint.   Only returned records after December 14, 2021.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wheatliu/okx-go"
)

func main() {
	ccy := "ccy_example" // string | Currency, e.g. `BTC` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ts`, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ts`, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`.  If `ccy` is not specified, all data under the same `ts` will be returned, not limited by `limit` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnFlexibleAPI.GetFinanceSavingsLendingRateHistoryV5(context.Background()).Ccy(ccy).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnFlexibleAPI.GetFinanceSavingsLendingRateHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceSavingsLendingRateHistoryV5`: GetFinanceSavingsLendingRateHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnFlexibleAPI.GetFinanceSavingsLendingRateHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceSavingsLendingRateHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ts&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ts&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;.  If &#x60;ccy&#x60; is not specified, all data under the same &#x60;ts&#x60; will be returned, not limited by &#x60;limit&#x60; | [default to &quot;&quot;]

### Return type

[**GetFinanceSavingsLendingRateHistoryV5Resp**](GetFinanceSavingsLendingRateHistoryV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceSavingsLendingRateSummaryV5

> GetFinanceSavingsLendingRateSummaryV5Resp GetFinanceSavingsLendingRateSummaryV5(ctx).Ccy(ccy).Execute()

Authentication is not required for this public endpoint.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wheatliu/okx-go"
)

func main() {
	ccy := "ccy_example" // string | Currency, e.g. `BTC` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimpleEarnFlexibleAPI.GetFinanceSavingsLendingRateSummaryV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimpleEarnFlexibleAPI.GetFinanceSavingsLendingRateSummaryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceSavingsLendingRateSummaryV5`: GetFinanceSavingsLendingRateSummaryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SimpleEarnFlexibleAPI.GetFinanceSavingsLendingRateSummaryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceSavingsLendingRateSummaryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]

### Return type

[**GetFinanceSavingsLendingRateSummaryV5Resp**](GetFinanceSavingsLendingRateSummaryV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

