# \GridTradingAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTradingBotGridAdjustInvestmentV5**](GridTradingAPI.md#CreateTradingBotGridAdjustInvestmentV5) | **Post** /api/v5/tradingBot/grid/adjust-investment | It is used to add investment and only applicable to contract gird.  
[**CreateTradingBotGridAmendOrderAlgoV5**](GridTradingAPI.md#CreateTradingBotGridAmendOrderAlgoV5) | **Post** /api/v5/tradingBot/grid/amend-order-algo | 
[**CreateTradingBotGridCancelCloseOrderV5**](GridTradingAPI.md#CreateTradingBotGridCancelCloseOrderV5) | **Post** /api/v5/tradingBot/grid/cancel-close-order | 
[**CreateTradingBotGridClosePositionV5**](GridTradingAPI.md#CreateTradingBotGridClosePositionV5) | **Post** /api/v5/tradingBot/grid/close-position | Close position when the contract grid stop type is &#39;keep position&#39;.  
[**CreateTradingBotGridComputeMarginBalanceV5**](GridTradingAPI.md#CreateTradingBotGridComputeMarginBalanceV5) | **Post** /api/v5/tradingBot/grid/compute-margin-balance | 
[**CreateTradingBotGridMarginBalanceV5**](GridTradingAPI.md#CreateTradingBotGridMarginBalanceV5) | **Post** /api/v5/tradingBot/grid/margin-balance | 
[**CreateTradingBotGridMinInvestmentV5**](GridTradingAPI.md#CreateTradingBotGridMinInvestmentV5) | **Post** /api/v5/tradingBot/grid/min-investment | Authentication is not required for this public endpoint.  
[**CreateTradingBotGridOrderAlgoV5**](GridTradingAPI.md#CreateTradingBotGridOrderAlgoV5) | **Post** /api/v5/tradingBot/grid/order-algo | 
[**CreateTradingBotGridOrderInstantTriggerV5**](GridTradingAPI.md#CreateTradingBotGridOrderInstantTriggerV5) | **Post** /api/v5/tradingBot/grid/order-instant-trigger | 
[**CreateTradingBotGridStopOrderAlgoV5**](GridTradingAPI.md#CreateTradingBotGridStopOrderAlgoV5) | **Post** /api/v5/tradingBot/grid/stop-order-algo | A maximum of 10 orders can be stopped per request.  
[**CreateTradingBotGridWithdrawIncomeV5**](GridTradingAPI.md#CreateTradingBotGridWithdrawIncomeV5) | **Post** /api/v5/tradingBot/grid/withdraw-income | 
[**GetTradingBotGridAiParamV5**](GridTradingAPI.md#GetTradingBotGridAiParamV5) | **Get** /api/v5/tradingBot/grid/ai-param | Authentication is not required for this public endpoint.  
[**GetTradingBotGridGridQuantityV5**](GridTradingAPI.md#GetTradingBotGridGridQuantityV5) | **Get** /api/v5/tradingBot/grid/grid-quantity | Authentication is not required for this public endpoint.    Maximum grid quantity can be retrieved from this endpoint. Minimum grid quantity always is 2.  
[**GetTradingBotGridOrdersAlgoDetailsV5**](GridTradingAPI.md#GetTradingBotGridOrdersAlgoDetailsV5) | **Get** /api/v5/tradingBot/grid/orders-algo-details | 
[**GetTradingBotGridOrdersAlgoHistoryV5**](GridTradingAPI.md#GetTradingBotGridOrdersAlgoHistoryV5) | **Get** /api/v5/tradingBot/grid/orders-algo-history | 
[**GetTradingBotGridOrdersAlgoPendingV5**](GridTradingAPI.md#GetTradingBotGridOrdersAlgoPendingV5) | **Get** /api/v5/tradingBot/grid/orders-algo-pending | 
[**GetTradingBotGridPositionsV5**](GridTradingAPI.md#GetTradingBotGridPositionsV5) | **Get** /api/v5/tradingBot/grid/positions | 
[**GetTradingBotGridSubOrdersV5**](GridTradingAPI.md#GetTradingBotGridSubOrdersV5) | **Get** /api/v5/tradingBot/grid/sub-orders | 
[**GetTradingBotPublicRsiBackTestingV5**](GridTradingAPI.md#GetTradingBotPublicRsiBackTestingV5) | **Get** /api/v5/tradingBot/public/rsi-back-testing | Authentication is not required for this public endpoint.  



## CreateTradingBotGridAdjustInvestmentV5

> CreateTradingBotGridAdjustInvestmentV5Resp CreateTradingBotGridAdjustInvestmentV5(ctx).CreateTradingBotGridAdjustInvestmentV5Req(createTradingBotGridAdjustInvestmentV5Req).Execute()

It is used to add investment and only applicable to contract gird.  



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
	createTradingBotGridAdjustInvestmentV5Req := *openapiclient.NewCreateTradingBotGridAdjustInvestmentV5Req("AlgoId_example", "Amt_example") // CreateTradingBotGridAdjustInvestmentV5Req | The request body for CreateTradingBotGridAdjustInvestmentV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.CreateTradingBotGridAdjustInvestmentV5(context.Background()).CreateTradingBotGridAdjustInvestmentV5Req(createTradingBotGridAdjustInvestmentV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.CreateTradingBotGridAdjustInvestmentV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotGridAdjustInvestmentV5`: CreateTradingBotGridAdjustInvestmentV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.CreateTradingBotGridAdjustInvestmentV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotGridAdjustInvestmentV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotGridAdjustInvestmentV5Req** | [**CreateTradingBotGridAdjustInvestmentV5Req**](CreateTradingBotGridAdjustInvestmentV5Req.md) | The request body for CreateTradingBotGridAdjustInvestmentV5 | 

### Return type

[**CreateTradingBotGridAdjustInvestmentV5Resp**](CreateTradingBotGridAdjustInvestmentV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotGridAmendOrderAlgoV5

> CreateTradingBotGridAmendOrderAlgoV5Resp CreateTradingBotGridAmendOrderAlgoV5(ctx).CreateTradingBotGridAmendOrderAlgoV5Req(createTradingBotGridAmendOrderAlgoV5Req).Execute()





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
	createTradingBotGridAmendOrderAlgoV5Req := *openapiclient.NewCreateTradingBotGridAmendOrderAlgoV5Req("AlgoId_example", "InstId_example") // CreateTradingBotGridAmendOrderAlgoV5Req | The request body for CreateTradingBotGridAmendOrderAlgoV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.CreateTradingBotGridAmendOrderAlgoV5(context.Background()).CreateTradingBotGridAmendOrderAlgoV5Req(createTradingBotGridAmendOrderAlgoV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.CreateTradingBotGridAmendOrderAlgoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotGridAmendOrderAlgoV5`: CreateTradingBotGridAmendOrderAlgoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.CreateTradingBotGridAmendOrderAlgoV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotGridAmendOrderAlgoV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotGridAmendOrderAlgoV5Req** | [**CreateTradingBotGridAmendOrderAlgoV5Req**](CreateTradingBotGridAmendOrderAlgoV5Req.md) | The request body for CreateTradingBotGridAmendOrderAlgoV5 | 

### Return type

[**CreateTradingBotGridAmendOrderAlgoV5Resp**](CreateTradingBotGridAmendOrderAlgoV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotGridCancelCloseOrderV5

> CreateTradingBotGridCancelCloseOrderV5Resp CreateTradingBotGridCancelCloseOrderV5(ctx).CreateTradingBotGridCancelCloseOrderV5Req(createTradingBotGridCancelCloseOrderV5Req).Execute()





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
	createTradingBotGridCancelCloseOrderV5Req := *openapiclient.NewCreateTradingBotGridCancelCloseOrderV5Req("AlgoId_example", "OrdId_example") // CreateTradingBotGridCancelCloseOrderV5Req | The request body for CreateTradingBotGridCancelCloseOrderV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.CreateTradingBotGridCancelCloseOrderV5(context.Background()).CreateTradingBotGridCancelCloseOrderV5Req(createTradingBotGridCancelCloseOrderV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.CreateTradingBotGridCancelCloseOrderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotGridCancelCloseOrderV5`: CreateTradingBotGridCancelCloseOrderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.CreateTradingBotGridCancelCloseOrderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotGridCancelCloseOrderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotGridCancelCloseOrderV5Req** | [**CreateTradingBotGridCancelCloseOrderV5Req**](CreateTradingBotGridCancelCloseOrderV5Req.md) | The request body for CreateTradingBotGridCancelCloseOrderV5 | 

### Return type

[**CreateTradingBotGridCancelCloseOrderV5Resp**](CreateTradingBotGridCancelCloseOrderV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotGridClosePositionV5

> CreateTradingBotGridClosePositionV5Resp CreateTradingBotGridClosePositionV5(ctx).CreateTradingBotGridClosePositionV5Req(createTradingBotGridClosePositionV5Req).Execute()

Close position when the contract grid stop type is 'keep position'.  



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
	createTradingBotGridClosePositionV5Req := *openapiclient.NewCreateTradingBotGridClosePositionV5Req("AlgoId_example", false) // CreateTradingBotGridClosePositionV5Req | The request body for CreateTradingBotGridClosePositionV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.CreateTradingBotGridClosePositionV5(context.Background()).CreateTradingBotGridClosePositionV5Req(createTradingBotGridClosePositionV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.CreateTradingBotGridClosePositionV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotGridClosePositionV5`: CreateTradingBotGridClosePositionV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.CreateTradingBotGridClosePositionV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotGridClosePositionV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotGridClosePositionV5Req** | [**CreateTradingBotGridClosePositionV5Req**](CreateTradingBotGridClosePositionV5Req.md) | The request body for CreateTradingBotGridClosePositionV5 | 

### Return type

[**CreateTradingBotGridClosePositionV5Resp**](CreateTradingBotGridClosePositionV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotGridComputeMarginBalanceV5

> CreateTradingBotGridComputeMarginBalanceV5Resp CreateTradingBotGridComputeMarginBalanceV5(ctx).CreateTradingBotGridComputeMarginBalanceV5Req(createTradingBotGridComputeMarginBalanceV5Req).Execute()





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
	createTradingBotGridComputeMarginBalanceV5Req := *openapiclient.NewCreateTradingBotGridComputeMarginBalanceV5Req("AlgoId_example", "Type_example") // CreateTradingBotGridComputeMarginBalanceV5Req | The request body for CreateTradingBotGridComputeMarginBalanceV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.CreateTradingBotGridComputeMarginBalanceV5(context.Background()).CreateTradingBotGridComputeMarginBalanceV5Req(createTradingBotGridComputeMarginBalanceV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.CreateTradingBotGridComputeMarginBalanceV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotGridComputeMarginBalanceV5`: CreateTradingBotGridComputeMarginBalanceV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.CreateTradingBotGridComputeMarginBalanceV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotGridComputeMarginBalanceV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotGridComputeMarginBalanceV5Req** | [**CreateTradingBotGridComputeMarginBalanceV5Req**](CreateTradingBotGridComputeMarginBalanceV5Req.md) | The request body for CreateTradingBotGridComputeMarginBalanceV5 | 

### Return type

[**CreateTradingBotGridComputeMarginBalanceV5Resp**](CreateTradingBotGridComputeMarginBalanceV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotGridMarginBalanceV5

> CreateTradingBotGridMarginBalanceV5Resp CreateTradingBotGridMarginBalanceV5(ctx).CreateTradingBotGridMarginBalanceV5Req(createTradingBotGridMarginBalanceV5Req).Execute()





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
	createTradingBotGridMarginBalanceV5Req := *openapiclient.NewCreateTradingBotGridMarginBalanceV5Req("AlgoId_example", "Type_example") // CreateTradingBotGridMarginBalanceV5Req | The request body for CreateTradingBotGridMarginBalanceV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.CreateTradingBotGridMarginBalanceV5(context.Background()).CreateTradingBotGridMarginBalanceV5Req(createTradingBotGridMarginBalanceV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.CreateTradingBotGridMarginBalanceV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotGridMarginBalanceV5`: CreateTradingBotGridMarginBalanceV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.CreateTradingBotGridMarginBalanceV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotGridMarginBalanceV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotGridMarginBalanceV5Req** | [**CreateTradingBotGridMarginBalanceV5Req**](CreateTradingBotGridMarginBalanceV5Req.md) | The request body for CreateTradingBotGridMarginBalanceV5 | 

### Return type

[**CreateTradingBotGridMarginBalanceV5Resp**](CreateTradingBotGridMarginBalanceV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotGridMinInvestmentV5

> CreateTradingBotGridMinInvestmentV5Resp CreateTradingBotGridMinInvestmentV5(ctx).CreateTradingBotGridMinInvestmentV5Req(createTradingBotGridMinInvestmentV5Req).Execute()

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
	createTradingBotGridMinInvestmentV5Req := *openapiclient.NewCreateTradingBotGridMinInvestmentV5Req("AlgoOrdType_example", "GridNum_example", "InstId_example", "MaxPx_example", "MinPx_example", "RunType_example") // CreateTradingBotGridMinInvestmentV5Req | The request body for CreateTradingBotGridMinInvestmentV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.CreateTradingBotGridMinInvestmentV5(context.Background()).CreateTradingBotGridMinInvestmentV5Req(createTradingBotGridMinInvestmentV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.CreateTradingBotGridMinInvestmentV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotGridMinInvestmentV5`: CreateTradingBotGridMinInvestmentV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.CreateTradingBotGridMinInvestmentV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotGridMinInvestmentV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotGridMinInvestmentV5Req** | [**CreateTradingBotGridMinInvestmentV5Req**](CreateTradingBotGridMinInvestmentV5Req.md) | The request body for CreateTradingBotGridMinInvestmentV5 | 

### Return type

[**CreateTradingBotGridMinInvestmentV5Resp**](CreateTradingBotGridMinInvestmentV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotGridOrderAlgoV5

> CreateTradingBotGridOrderAlgoV5Resp CreateTradingBotGridOrderAlgoV5(ctx).CreateTradingBotGridOrderAlgoV5Req(createTradingBotGridOrderAlgoV5Req).Execute()





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
	createTradingBotGridOrderAlgoV5Req := *openapiclient.NewCreateTradingBotGridOrderAlgoV5Req("AlgoOrdType_example", "GridNum_example", "InstId_example", "MaxPx_example", "MinPx_example") // CreateTradingBotGridOrderAlgoV5Req | The request body for CreateTradingBotGridOrderAlgoV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.CreateTradingBotGridOrderAlgoV5(context.Background()).CreateTradingBotGridOrderAlgoV5Req(createTradingBotGridOrderAlgoV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.CreateTradingBotGridOrderAlgoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotGridOrderAlgoV5`: CreateTradingBotGridOrderAlgoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.CreateTradingBotGridOrderAlgoV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotGridOrderAlgoV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotGridOrderAlgoV5Req** | [**CreateTradingBotGridOrderAlgoV5Req**](CreateTradingBotGridOrderAlgoV5Req.md) | The request body for CreateTradingBotGridOrderAlgoV5 | 

### Return type

[**CreateTradingBotGridOrderAlgoV5Resp**](CreateTradingBotGridOrderAlgoV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotGridOrderInstantTriggerV5

> CreateTradingBotGridOrderInstantTriggerV5Resp CreateTradingBotGridOrderInstantTriggerV5(ctx).CreateTradingBotGridOrderInstantTriggerV5Req(createTradingBotGridOrderInstantTriggerV5Req).Execute()





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
	createTradingBotGridOrderInstantTriggerV5Req := *openapiclient.NewCreateTradingBotGridOrderInstantTriggerV5Req("AlgoId_example") // CreateTradingBotGridOrderInstantTriggerV5Req | The request body for CreateTradingBotGridOrderInstantTriggerV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.CreateTradingBotGridOrderInstantTriggerV5(context.Background()).CreateTradingBotGridOrderInstantTriggerV5Req(createTradingBotGridOrderInstantTriggerV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.CreateTradingBotGridOrderInstantTriggerV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotGridOrderInstantTriggerV5`: CreateTradingBotGridOrderInstantTriggerV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.CreateTradingBotGridOrderInstantTriggerV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotGridOrderInstantTriggerV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotGridOrderInstantTriggerV5Req** | [**CreateTradingBotGridOrderInstantTriggerV5Req**](CreateTradingBotGridOrderInstantTriggerV5Req.md) | The request body for CreateTradingBotGridOrderInstantTriggerV5 | 

### Return type

[**CreateTradingBotGridOrderInstantTriggerV5Resp**](CreateTradingBotGridOrderInstantTriggerV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotGridStopOrderAlgoV5

> CreateTradingBotGridStopOrderAlgoV5Resp CreateTradingBotGridStopOrderAlgoV5(ctx).CreateTradingBotGridStopOrderAlgoV5Req(createTradingBotGridStopOrderAlgoV5Req).Execute()

A maximum of 10 orders can be stopped per request.  



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
	createTradingBotGridStopOrderAlgoV5Req := *openapiclient.NewCreateTradingBotGridStopOrderAlgoV5Req("AlgoId_example", "AlgoOrdType_example", "InstId_example", "StopType_example") // CreateTradingBotGridStopOrderAlgoV5Req | The request body for CreateTradingBotGridStopOrderAlgoV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.CreateTradingBotGridStopOrderAlgoV5(context.Background()).CreateTradingBotGridStopOrderAlgoV5Req(createTradingBotGridStopOrderAlgoV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.CreateTradingBotGridStopOrderAlgoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotGridStopOrderAlgoV5`: CreateTradingBotGridStopOrderAlgoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.CreateTradingBotGridStopOrderAlgoV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotGridStopOrderAlgoV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotGridStopOrderAlgoV5Req** | [**CreateTradingBotGridStopOrderAlgoV5Req**](CreateTradingBotGridStopOrderAlgoV5Req.md) | The request body for CreateTradingBotGridStopOrderAlgoV5 | 

### Return type

[**CreateTradingBotGridStopOrderAlgoV5Resp**](CreateTradingBotGridStopOrderAlgoV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotGridWithdrawIncomeV5

> CreateTradingBotGridWithdrawIncomeV5Resp CreateTradingBotGridWithdrawIncomeV5(ctx).CreateTradingBotGridWithdrawIncomeV5Req(createTradingBotGridWithdrawIncomeV5Req).Execute()





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
	createTradingBotGridWithdrawIncomeV5Req := *openapiclient.NewCreateTradingBotGridWithdrawIncomeV5Req("AlgoId_example") // CreateTradingBotGridWithdrawIncomeV5Req | The request body for CreateTradingBotGridWithdrawIncomeV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.CreateTradingBotGridWithdrawIncomeV5(context.Background()).CreateTradingBotGridWithdrawIncomeV5Req(createTradingBotGridWithdrawIncomeV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.CreateTradingBotGridWithdrawIncomeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotGridWithdrawIncomeV5`: CreateTradingBotGridWithdrawIncomeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.CreateTradingBotGridWithdrawIncomeV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotGridWithdrawIncomeV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotGridWithdrawIncomeV5Req** | [**CreateTradingBotGridWithdrawIncomeV5Req**](CreateTradingBotGridWithdrawIncomeV5Req.md) | The request body for CreateTradingBotGridWithdrawIncomeV5 | 

### Return type

[**CreateTradingBotGridWithdrawIncomeV5Resp**](CreateTradingBotGridWithdrawIncomeV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotGridAiParamV5

> GetTradingBotGridAiParamV5Resp GetTradingBotGridAiParamV5(ctx).AlgoOrdType(algoOrdType).InstId(instId).Direction(direction).Duration(duration).Execute()

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
	algoOrdType := "algoOrdType_example" // string | Algo order type  `grid`: Spot grid  `contract_grid`: Contract grid (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT` (default to "")
	direction := "direction_example" // string | Contract grid type  `long`,`short`,`neutral`  Required in the case of `contract_grid` (optional) (default to "")
	duration := "duration_example" // string | Back testing duration  `7D`: 7 Days, `30D`: 30 Days, `180D`: 180 Days  The default is `7D` for `Spot grid`  Only `7D` is available for `Contract grid` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.GetTradingBotGridAiParamV5(context.Background()).AlgoOrdType(algoOrdType).InstId(instId).Direction(direction).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.GetTradingBotGridAiParamV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotGridAiParamV5`: GetTradingBotGridAiParamV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.GetTradingBotGridAiParamV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotGridAiParamV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoOrdType** | **string** | Algo order type  &#x60;grid&#x60;: Spot grid  &#x60;contract_grid&#x60;: Contract grid | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **direction** | **string** | Contract grid type  &#x60;long&#x60;,&#x60;short&#x60;,&#x60;neutral&#x60;  Required in the case of &#x60;contract_grid&#x60; | [default to &quot;&quot;]
 **duration** | **string** | Back testing duration  &#x60;7D&#x60;: 7 Days, &#x60;30D&#x60;: 30 Days, &#x60;180D&#x60;: 180 Days  The default is &#x60;7D&#x60; for &#x60;Spot grid&#x60;  Only &#x60;7D&#x60; is available for &#x60;Contract grid&#x60; | [default to &quot;&quot;]

### Return type

[**GetTradingBotGridAiParamV5Resp**](GetTradingBotGridAiParamV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotGridGridQuantityV5

> GetTradingBotGridGridQuantityV5Resp GetTradingBotGridGridQuantityV5(ctx).InstId(instId).RunType(runType).AlgoOrdType(algoOrdType).MaxPx(maxPx).MinPx(minPx).Lever(lever).Execute()

Authentication is not required for this public endpoint.    Maximum grid quantity can be retrieved from this endpoint. Minimum grid quantity always is 2.  



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
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT` (default to "")
	runType := "runType_example" // string | Grid type  `1`: Arithmetic   `2`: Geometric (default to "")
	algoOrdType := "algoOrdType_example" // string | Algo order type  `grid`: Spot grid  `contract_grid`: Contract grid (default to "")
	maxPx := "maxPx_example" // string | Upper price of price range (default to "")
	minPx := "minPx_example" // string | Lower price of price range (default to "")
	lever := "lever_example" // string | Leverage, it is required for contract grid (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.GetTradingBotGridGridQuantityV5(context.Background()).InstId(instId).RunType(runType).AlgoOrdType(algoOrdType).MaxPx(maxPx).MinPx(minPx).Lever(lever).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.GetTradingBotGridGridQuantityV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotGridGridQuantityV5`: GetTradingBotGridGridQuantityV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.GetTradingBotGridGridQuantityV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotGridGridQuantityV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **runType** | **string** | Grid type  &#x60;1&#x60;: Arithmetic   &#x60;2&#x60;: Geometric | [default to &quot;&quot;]
 **algoOrdType** | **string** | Algo order type  &#x60;grid&#x60;: Spot grid  &#x60;contract_grid&#x60;: Contract grid | [default to &quot;&quot;]
 **maxPx** | **string** | Upper price of price range | [default to &quot;&quot;]
 **minPx** | **string** | Lower price of price range | [default to &quot;&quot;]
 **lever** | **string** | Leverage, it is required for contract grid | [default to &quot;&quot;]

### Return type

[**GetTradingBotGridGridQuantityV5Resp**](GetTradingBotGridGridQuantityV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotGridOrdersAlgoDetailsV5

> GetTradingBotGridOrdersAlgoDetailsV5Resp GetTradingBotGridOrdersAlgoDetailsV5(ctx).AlgoOrdType(algoOrdType).AlgoId(algoId).Execute()





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
	algoOrdType := "algoOrdType_example" // string | Algo order type  `grid`: Spot grid  `contract_grid`: Contract grid (default to "")
	algoId := "algoId_example" // string | Algo ID (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.GetTradingBotGridOrdersAlgoDetailsV5(context.Background()).AlgoOrdType(algoOrdType).AlgoId(algoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.GetTradingBotGridOrdersAlgoDetailsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotGridOrdersAlgoDetailsV5`: GetTradingBotGridOrdersAlgoDetailsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.GetTradingBotGridOrdersAlgoDetailsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotGridOrdersAlgoDetailsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoOrdType** | **string** | Algo order type  &#x60;grid&#x60;: Spot grid  &#x60;contract_grid&#x60;: Contract grid | [default to &quot;&quot;]
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]

### Return type

[**GetTradingBotGridOrdersAlgoDetailsV5Resp**](GetTradingBotGridOrdersAlgoDetailsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotGridOrdersAlgoHistoryV5

> GetTradingBotGridOrdersAlgoHistoryV5Resp GetTradingBotGridOrdersAlgoHistoryV5(ctx).AlgoOrdType(algoOrdType).AlgoId(algoId).InstId(instId).InstType(instType).After(after).Before(before).Limit(limit).Execute()





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
	algoOrdType := "algoOrdType_example" // string | Algo order type  `grid`: Spot grid  `contract_grid`: Contract grid (default to "")
	algoId := "algoId_example" // string | Algo ID (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT` (optional) (default to "")
	instType := "instType_example" // string | Instrument type  `SPOT`  `MARGIN`  `FUTURES`  `SWAP` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `algoId`. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `algoId`. (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.GetTradingBotGridOrdersAlgoHistoryV5(context.Background()).AlgoOrdType(algoOrdType).AlgoId(algoId).InstId(instId).InstType(instType).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.GetTradingBotGridOrdersAlgoHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotGridOrdersAlgoHistoryV5`: GetTradingBotGridOrdersAlgoHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.GetTradingBotGridOrdersAlgoHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotGridOrdersAlgoHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoOrdType** | **string** | Algo order type  &#x60;grid&#x60;: Spot grid  &#x60;contract_grid&#x60;: Contract grid | [default to &quot;&quot;]
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;MARGIN&#x60;  &#x60;FUTURES&#x60;  &#x60;SWAP&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;algoId&#x60;. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;algoId&#x60;. | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetTradingBotGridOrdersAlgoHistoryV5Resp**](GetTradingBotGridOrdersAlgoHistoryV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotGridOrdersAlgoPendingV5

> GetTradingBotGridOrdersAlgoPendingV5Resp GetTradingBotGridOrdersAlgoPendingV5(ctx).AlgoOrdType(algoOrdType).AlgoId(algoId).InstId(instId).InstType(instType).After(after).Before(before).Limit(limit).Execute()





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
	algoOrdType := "algoOrdType_example" // string | Algo order type  `grid`: Spot grid  `contract_grid`: Contract grid (default to "")
	algoId := "algoId_example" // string | Algo ID (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT` (optional) (default to "")
	instType := "instType_example" // string | Instrument type  `SPOT`  `MARGIN`  `FUTURES`  `SWAP` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `algoId`. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `algoId`. (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.GetTradingBotGridOrdersAlgoPendingV5(context.Background()).AlgoOrdType(algoOrdType).AlgoId(algoId).InstId(instId).InstType(instType).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.GetTradingBotGridOrdersAlgoPendingV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotGridOrdersAlgoPendingV5`: GetTradingBotGridOrdersAlgoPendingV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.GetTradingBotGridOrdersAlgoPendingV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotGridOrdersAlgoPendingV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoOrdType** | **string** | Algo order type  &#x60;grid&#x60;: Spot grid  &#x60;contract_grid&#x60;: Contract grid | [default to &quot;&quot;]
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;MARGIN&#x60;  &#x60;FUTURES&#x60;  &#x60;SWAP&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;algoId&#x60;. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;algoId&#x60;. | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100 | [default to &quot;&quot;]

### Return type

[**GetTradingBotGridOrdersAlgoPendingV5Resp**](GetTradingBotGridOrdersAlgoPendingV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotGridPositionsV5

> GetTradingBotGridPositionsV5Resp GetTradingBotGridPositionsV5(ctx).AlgoOrdType(algoOrdType).AlgoId(algoId).Execute()





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
	algoOrdType := "algoOrdType_example" // string | Algo order type  `contract_grid`: Contract grid (default to "")
	algoId := "algoId_example" // string | Algo ID (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.GetTradingBotGridPositionsV5(context.Background()).AlgoOrdType(algoOrdType).AlgoId(algoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.GetTradingBotGridPositionsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotGridPositionsV5`: GetTradingBotGridPositionsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.GetTradingBotGridPositionsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotGridPositionsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoOrdType** | **string** | Algo order type  &#x60;contract_grid&#x60;: Contract grid | [default to &quot;&quot;]
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]

### Return type

[**GetTradingBotGridPositionsV5Resp**](GetTradingBotGridPositionsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotGridSubOrdersV5

> GetTradingBotGridSubOrdersV5Resp GetTradingBotGridSubOrdersV5(ctx).AlgoOrdType(algoOrdType).AlgoId(algoId).Type_(type_).GroupId(groupId).After(after).Before(before).Limit(limit).Execute()





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
	algoOrdType := "algoOrdType_example" // string | Algo order type  `grid`: Spot grid  `contract_grid`: Contract grid (default to "")
	algoId := "algoId_example" // string | Algo ID (default to "")
	type_ := "type__example" // string | Sub order state  `live`  `filled` (default to "")
	groupId := "groupId_example" // string | Group ID (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ordId`. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ordId`. (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.GetTradingBotGridSubOrdersV5(context.Background()).AlgoOrdType(algoOrdType).AlgoId(algoId).Type_(type_).GroupId(groupId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.GetTradingBotGridSubOrdersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotGridSubOrdersV5`: GetTradingBotGridSubOrdersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.GetTradingBotGridSubOrdersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotGridSubOrdersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoOrdType** | **string** | Algo order type  &#x60;grid&#x60;: Spot grid  &#x60;contract_grid&#x60;: Contract grid | [default to &quot;&quot;]
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]
 **type_** | **string** | Sub order state  &#x60;live&#x60;  &#x60;filled&#x60; | [default to &quot;&quot;]
 **groupId** | **string** | Group ID | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ordId&#x60;. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ordId&#x60;. | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100 | [default to &quot;&quot;]

### Return type

[**GetTradingBotGridSubOrdersV5Resp**](GetTradingBotGridSubOrdersV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotPublicRsiBackTestingV5

> GetTradingBotPublicRsiBackTestingV5Resp GetTradingBotPublicRsiBackTestingV5(ctx).InstId(instId).Timeframe(timeframe).Thold(thold).TimePeriod(timePeriod).TriggerCond(triggerCond).Duration(duration).Execute()

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
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT`  Only applicable to `SPOT` (default to "")
	timeframe := "timeframe_example" // string | K-line type  `3m`, `5m`, `15m`, `30m` (`m`: minute)  `1H`, `4H` (`H`: hour)  `1D` (`D`: day) (default to "")
	thold := "thold_example" // string | Threshold  The value should be an integer between 1 to 100 (default to "")
	timePeriod := "timePeriod_example" // string | Time Period  `14` (default to "")
	triggerCond := "triggerCond_example" // string | Trigger condition  `cross_up`  `cross_down`  `above`  `below`  `cross`  Default is `cross_down` (optional) (default to "")
	duration := "duration_example" // string | Back testing duration  `1M` (`M`: month)  Default is `1M` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GridTradingAPI.GetTradingBotPublicRsiBackTestingV5(context.Background()).InstId(instId).Timeframe(timeframe).Thold(thold).TimePeriod(timePeriod).TriggerCond(triggerCond).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridTradingAPI.GetTradingBotPublicRsiBackTestingV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotPublicRsiBackTestingV5`: GetTradingBotPublicRsiBackTestingV5Resp
	fmt.Fprintf(os.Stdout, "Response from `GridTradingAPI.GetTradingBotPublicRsiBackTestingV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotPublicRsiBackTestingV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60;  Only applicable to &#x60;SPOT&#x60; | [default to &quot;&quot;]
 **timeframe** | **string** | K-line type  &#x60;3m&#x60;, &#x60;5m&#x60;, &#x60;15m&#x60;, &#x60;30m&#x60; (&#x60;m&#x60;: minute)  &#x60;1H&#x60;, &#x60;4H&#x60; (&#x60;H&#x60;: hour)  &#x60;1D&#x60; (&#x60;D&#x60;: day) | [default to &quot;&quot;]
 **thold** | **string** | Threshold  The value should be an integer between 1 to 100 | [default to &quot;&quot;]
 **timePeriod** | **string** | Time Period  &#x60;14&#x60; | [default to &quot;&quot;]
 **triggerCond** | **string** | Trigger condition  &#x60;cross_up&#x60;  &#x60;cross_down&#x60;  &#x60;above&#x60;  &#x60;below&#x60;  &#x60;cross&#x60;  Default is &#x60;cross_down&#x60; | [default to &quot;&quot;]
 **duration** | **string** | Back testing duration  &#x60;1M&#x60; (&#x60;M&#x60;: month)  Default is &#x60;1M&#x60; | [default to &quot;&quot;]

### Return type

[**GetTradingBotPublicRsiBackTestingV5Resp**](GetTradingBotPublicRsiBackTestingV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

