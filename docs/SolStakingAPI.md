# \SolStakingAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFinanceStakingDefiSolPurchaseV5**](SolStakingAPI.md#CreateFinanceStakingDefiSolPurchaseV5) | **Post** /api/v5/finance/staking-defi/sol/purchase | Staking SOL for OKSOL   Only the assets in the funding account can be used.    
[**CreateFinanceStakingDefiSolRedeemV5**](SolStakingAPI.md#CreateFinanceStakingDefiSolRedeemV5) | **Post** /api/v5/finance/staking-defi/sol/redeem | Only the assets in the funding account can be used. If your OKSOL is in your trading account, you can make funding transfer first.    
[**GetFinanceStakingDefiSolApyHistoryV5**](SolStakingAPI.md#GetFinanceStakingDefiSolApyHistoryV5) | **Get** /api/v5/finance/staking-defi/sol/apy-history | Public endpoints don&#39;t need authorization.  
[**GetFinanceStakingDefiSolBalanceV5**](SolStakingAPI.md#GetFinanceStakingDefiSolBalanceV5) | **Get** /api/v5/finance/staking-defi/sol/balance | The balance is summarized all OKSOL assets (including assets in redeeming) in account.  
[**GetFinanceStakingDefiSolPurchaseRedeemHistoryV5**](SolStakingAPI.md#GetFinanceStakingDefiSolPurchaseRedeemHistoryV5) | **Get** /api/v5/finance/staking-defi/sol/purchase-redeem-history | 



## CreateFinanceStakingDefiSolPurchaseV5

> CreateFinanceStakingDefiSolPurchaseV5Resp CreateFinanceStakingDefiSolPurchaseV5(ctx).CreateFinanceStakingDefiSolPurchaseV5Req(createFinanceStakingDefiSolPurchaseV5Req).Execute()

Staking SOL for OKSOL   Only the assets in the funding account can be used.    



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
	createFinanceStakingDefiSolPurchaseV5Req := *openapiclient.NewCreateFinanceStakingDefiSolPurchaseV5Req("Amt_example") // CreateFinanceStakingDefiSolPurchaseV5Req | The request body for CreateFinanceStaking-DefiSolPurchaseV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SolStakingAPI.CreateFinanceStakingDefiSolPurchaseV5(context.Background()).CreateFinanceStakingDefiSolPurchaseV5Req(createFinanceStakingDefiSolPurchaseV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SolStakingAPI.CreateFinanceStakingDefiSolPurchaseV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFinanceStakingDefiSolPurchaseV5`: CreateFinanceStakingDefiSolPurchaseV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SolStakingAPI.CreateFinanceStakingDefiSolPurchaseV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFinanceStakingDefiSolPurchaseV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFinanceStakingDefiSolPurchaseV5Req** | [**CreateFinanceStakingDefiSolPurchaseV5Req**](CreateFinanceStakingDefiSolPurchaseV5Req.md) | The request body for CreateFinanceStaking-DefiSolPurchaseV5 | 

### Return type

[**CreateFinanceStakingDefiSolPurchaseV5Resp**](CreateFinanceStakingDefiSolPurchaseV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFinanceStakingDefiSolRedeemV5

> CreateFinanceStakingDefiSolRedeemV5Resp CreateFinanceStakingDefiSolRedeemV5(ctx).CreateFinanceStakingDefiSolRedeemV5Req(createFinanceStakingDefiSolRedeemV5Req).Execute()

Only the assets in the funding account can be used. If your OKSOL is in your trading account, you can make funding transfer first.    



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
	createFinanceStakingDefiSolRedeemV5Req := *openapiclient.NewCreateFinanceStakingDefiSolRedeemV5Req("Amt_example") // CreateFinanceStakingDefiSolRedeemV5Req | The request body for CreateFinanceStaking-DefiSolRedeemV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SolStakingAPI.CreateFinanceStakingDefiSolRedeemV5(context.Background()).CreateFinanceStakingDefiSolRedeemV5Req(createFinanceStakingDefiSolRedeemV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SolStakingAPI.CreateFinanceStakingDefiSolRedeemV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFinanceStakingDefiSolRedeemV5`: CreateFinanceStakingDefiSolRedeemV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SolStakingAPI.CreateFinanceStakingDefiSolRedeemV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFinanceStakingDefiSolRedeemV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFinanceStakingDefiSolRedeemV5Req** | [**CreateFinanceStakingDefiSolRedeemV5Req**](CreateFinanceStakingDefiSolRedeemV5Req.md) | The request body for CreateFinanceStaking-DefiSolRedeemV5 | 

### Return type

[**CreateFinanceStakingDefiSolRedeemV5Resp**](CreateFinanceStakingDefiSolRedeemV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceStakingDefiSolApyHistoryV5

> GetFinanceStakingDefiSolApyHistoryV5Resp GetFinanceStakingDefiSolApyHistoryV5(ctx).Days(days).Execute()

Public endpoints don't need authorization.  



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
	days := "days_example" // string | Get the days of APY(Annual percentage yield) history record in the past  No more than 365 days (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SolStakingAPI.GetFinanceStakingDefiSolApyHistoryV5(context.Background()).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SolStakingAPI.GetFinanceStakingDefiSolApyHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceStakingDefiSolApyHistoryV5`: GetFinanceStakingDefiSolApyHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SolStakingAPI.GetFinanceStakingDefiSolApyHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceStakingDefiSolApyHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **string** | Get the days of APY(Annual percentage yield) history record in the past  No more than 365 days | [default to &quot;&quot;]

### Return type

[**GetFinanceStakingDefiSolApyHistoryV5Resp**](GetFinanceStakingDefiSolApyHistoryV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceStakingDefiSolBalanceV5

> GetFinanceStakingDefiSolBalanceV5Resp GetFinanceStakingDefiSolBalanceV5(ctx).Execute()

The balance is summarized all OKSOL assets (including assets in redeeming) in account.  



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SolStakingAPI.GetFinanceStakingDefiSolBalanceV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SolStakingAPI.GetFinanceStakingDefiSolBalanceV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceStakingDefiSolBalanceV5`: GetFinanceStakingDefiSolBalanceV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SolStakingAPI.GetFinanceStakingDefiSolBalanceV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceStakingDefiSolBalanceV5Request struct via the builder pattern


### Return type

[**GetFinanceStakingDefiSolBalanceV5Resp**](GetFinanceStakingDefiSolBalanceV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceStakingDefiSolPurchaseRedeemHistoryV5

> GetFinanceStakingDefiSolPurchaseRedeemHistoryV5Resp GetFinanceStakingDefiSolPurchaseRedeemHistoryV5(ctx).Type_(type_).Status(status).After(after).Before(before).Limit(limit).Execute()





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
	type_ := "type__example" // string | Type  `purchase`  `redeem` (optional) (default to "")
	status := "status_example" // string | Status  `pending`  `success`  `failed` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the `requestTime`. The value passed is the corresponding `timestamp` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the `requestTime`. The value passed is the corresponding `timestamp` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The default is `100`. The maximum is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SolStakingAPI.GetFinanceStakingDefiSolPurchaseRedeemHistoryV5(context.Background()).Type_(type_).Status(status).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SolStakingAPI.GetFinanceStakingDefiSolPurchaseRedeemHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceStakingDefiSolPurchaseRedeemHistoryV5`: GetFinanceStakingDefiSolPurchaseRedeemHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SolStakingAPI.GetFinanceStakingDefiSolPurchaseRedeemHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceStakingDefiSolPurchaseRedeemHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Type  &#x60;purchase&#x60;  &#x60;redeem&#x60; | [default to &quot;&quot;]
 **status** | **string** | Status  &#x60;pending&#x60;  &#x60;success&#x60;  &#x60;failed&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the &#x60;requestTime&#x60;. The value passed is the corresponding &#x60;timestamp&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the &#x60;requestTime&#x60;. The value passed is the corresponding &#x60;timestamp&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The default is &#x60;100&#x60;. The maximum is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetFinanceStakingDefiSolPurchaseRedeemHistoryV5Resp**](GetFinanceStakingDefiSolPurchaseRedeemHistoryV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

