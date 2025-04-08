# \EthStakingAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFinanceStakingDefiEthPurchaseV5**](EthStakingAPI.md#CreateFinanceStakingDefiEthPurchaseV5) | **Post** /api/v5/finance/staking-defi/eth/purchase | Staking ETH for BETH   Only the assets in the funding account can be used.    
[**CreateFinanceStakingDefiEthRedeemV5**](EthStakingAPI.md#CreateFinanceStakingDefiEthRedeemV5) | **Post** /api/v5/finance/staking-defi/eth/redeem | Only the assets in the funding account can be used. If your BETH is in your trading account, you can make funding transfer first.    
[**GetFinanceStakingDefiEthApyHistoryV5**](EthStakingAPI.md#GetFinanceStakingDefiEthApyHistoryV5) | **Get** /api/v5/finance/staking-defi/eth/apy-history | Public endpoints don&#39;t need authorization.  
[**GetFinanceStakingDefiEthBalanceV5**](EthStakingAPI.md#GetFinanceStakingDefiEthBalanceV5) | **Get** /api/v5/finance/staking-defi/eth/balance | The balance is a snapshot summarized all BETH assets (including assets in redeeming) in account.  
[**GetFinanceStakingDefiEthProductInfoV5**](EthStakingAPI.md#GetFinanceStakingDefiEthProductInfoV5) | **Get** /api/v5/finance/staking-defi/eth/product-info | 
[**GetFinanceStakingDefiEthPurchaseRedeemHistoryV5**](EthStakingAPI.md#GetFinanceStakingDefiEthPurchaseRedeemHistoryV5) | **Get** /api/v5/finance/staking-defi/eth/purchase-redeem-history | 



## CreateFinanceStakingDefiEthPurchaseV5

> CreateFinanceStakingDefiEthPurchaseV5Resp CreateFinanceStakingDefiEthPurchaseV5(ctx).CreateFinanceStakingDefiEthPurchaseV5Req(createFinanceStakingDefiEthPurchaseV5Req).Execute()

Staking ETH for BETH   Only the assets in the funding account can be used.    



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
	createFinanceStakingDefiEthPurchaseV5Req := *openapiclient.NewCreateFinanceStakingDefiEthPurchaseV5Req("Amt_example") // CreateFinanceStakingDefiEthPurchaseV5Req | The request body for CreateFinanceStaking-DefiEthPurchaseV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EthStakingAPI.CreateFinanceStakingDefiEthPurchaseV5(context.Background()).CreateFinanceStakingDefiEthPurchaseV5Req(createFinanceStakingDefiEthPurchaseV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EthStakingAPI.CreateFinanceStakingDefiEthPurchaseV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFinanceStakingDefiEthPurchaseV5`: CreateFinanceStakingDefiEthPurchaseV5Resp
	fmt.Fprintf(os.Stdout, "Response from `EthStakingAPI.CreateFinanceStakingDefiEthPurchaseV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFinanceStakingDefiEthPurchaseV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFinanceStakingDefiEthPurchaseV5Req** | [**CreateFinanceStakingDefiEthPurchaseV5Req**](CreateFinanceStakingDefiEthPurchaseV5Req.md) | The request body for CreateFinanceStaking-DefiEthPurchaseV5 | 

### Return type

[**CreateFinanceStakingDefiEthPurchaseV5Resp**](CreateFinanceStakingDefiEthPurchaseV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFinanceStakingDefiEthRedeemV5

> CreateFinanceStakingDefiEthRedeemV5Resp CreateFinanceStakingDefiEthRedeemV5(ctx).CreateFinanceStakingDefiEthRedeemV5Req(createFinanceStakingDefiEthRedeemV5Req).Execute()

Only the assets in the funding account can be used. If your BETH is in your trading account, you can make funding transfer first.    



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
	createFinanceStakingDefiEthRedeemV5Req := *openapiclient.NewCreateFinanceStakingDefiEthRedeemV5Req("Amt_example") // CreateFinanceStakingDefiEthRedeemV5Req | The request body for CreateFinanceStaking-DefiEthRedeemV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EthStakingAPI.CreateFinanceStakingDefiEthRedeemV5(context.Background()).CreateFinanceStakingDefiEthRedeemV5Req(createFinanceStakingDefiEthRedeemV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EthStakingAPI.CreateFinanceStakingDefiEthRedeemV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFinanceStakingDefiEthRedeemV5`: CreateFinanceStakingDefiEthRedeemV5Resp
	fmt.Fprintf(os.Stdout, "Response from `EthStakingAPI.CreateFinanceStakingDefiEthRedeemV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFinanceStakingDefiEthRedeemV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFinanceStakingDefiEthRedeemV5Req** | [**CreateFinanceStakingDefiEthRedeemV5Req**](CreateFinanceStakingDefiEthRedeemV5Req.md) | The request body for CreateFinanceStaking-DefiEthRedeemV5 | 

### Return type

[**CreateFinanceStakingDefiEthRedeemV5Resp**](CreateFinanceStakingDefiEthRedeemV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceStakingDefiEthApyHistoryV5

> GetFinanceStakingDefiEthApyHistoryV5Resp GetFinanceStakingDefiEthApyHistoryV5(ctx).Days(days).Execute()

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
	resp, r, err := apiClient.EthStakingAPI.GetFinanceStakingDefiEthApyHistoryV5(context.Background()).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EthStakingAPI.GetFinanceStakingDefiEthApyHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceStakingDefiEthApyHistoryV5`: GetFinanceStakingDefiEthApyHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `EthStakingAPI.GetFinanceStakingDefiEthApyHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceStakingDefiEthApyHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **string** | Get the days of APY(Annual percentage yield) history record in the past  No more than 365 days | [default to &quot;&quot;]

### Return type

[**GetFinanceStakingDefiEthApyHistoryV5Resp**](GetFinanceStakingDefiEthApyHistoryV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceStakingDefiEthBalanceV5

> GetFinanceStakingDefiEthBalanceV5Resp GetFinanceStakingDefiEthBalanceV5(ctx).Execute()

The balance is a snapshot summarized all BETH assets (including assets in redeeming) in account.  



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
	resp, r, err := apiClient.EthStakingAPI.GetFinanceStakingDefiEthBalanceV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EthStakingAPI.GetFinanceStakingDefiEthBalanceV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceStakingDefiEthBalanceV5`: GetFinanceStakingDefiEthBalanceV5Resp
	fmt.Fprintf(os.Stdout, "Response from `EthStakingAPI.GetFinanceStakingDefiEthBalanceV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceStakingDefiEthBalanceV5Request struct via the builder pattern


### Return type

[**GetFinanceStakingDefiEthBalanceV5Resp**](GetFinanceStakingDefiEthBalanceV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceStakingDefiEthProductInfoV5

> GetFinanceStakingDefiEthProductInfoV5Resp GetFinanceStakingDefiEthProductInfoV5(ctx).Execute()





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
	resp, r, err := apiClient.EthStakingAPI.GetFinanceStakingDefiEthProductInfoV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EthStakingAPI.GetFinanceStakingDefiEthProductInfoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceStakingDefiEthProductInfoV5`: GetFinanceStakingDefiEthProductInfoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `EthStakingAPI.GetFinanceStakingDefiEthProductInfoV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceStakingDefiEthProductInfoV5Request struct via the builder pattern


### Return type

[**GetFinanceStakingDefiEthProductInfoV5Resp**](GetFinanceStakingDefiEthProductInfoV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceStakingDefiEthPurchaseRedeemHistoryV5

> GetFinanceStakingDefiEthPurchaseRedeemHistoryV5Resp GetFinanceStakingDefiEthPurchaseRedeemHistoryV5(ctx).Type_(type_).Status(status).After(after).Before(before).Limit(limit).Execute()





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
	resp, r, err := apiClient.EthStakingAPI.GetFinanceStakingDefiEthPurchaseRedeemHistoryV5(context.Background()).Type_(type_).Status(status).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EthStakingAPI.GetFinanceStakingDefiEthPurchaseRedeemHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceStakingDefiEthPurchaseRedeemHistoryV5`: GetFinanceStakingDefiEthPurchaseRedeemHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `EthStakingAPI.GetFinanceStakingDefiEthPurchaseRedeemHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceStakingDefiEthPurchaseRedeemHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Type  &#x60;purchase&#x60;  &#x60;redeem&#x60; | [default to &quot;&quot;]
 **status** | **string** | Status  &#x60;pending&#x60;  &#x60;success&#x60;  &#x60;failed&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the &#x60;requestTime&#x60;. The value passed is the corresponding &#x60;timestamp&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the &#x60;requestTime&#x60;. The value passed is the corresponding &#x60;timestamp&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The default is &#x60;100&#x60;. The maximum is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetFinanceStakingDefiEthPurchaseRedeemHistoryV5Resp**](GetFinanceStakingDefiEthPurchaseRedeemHistoryV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

