# \SignalBotTradingAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTradingBotSignalAmendTPSLV5**](SignalBotTradingAPI.md#CreateTradingBotSignalAmendTPSLV5) | **Post** /api/v5/tradingBot/signal/amendTPSL | 
[**CreateTradingBotSignalCancelSubOrderV5**](SignalBotTradingAPI.md#CreateTradingBotSignalCancelSubOrderV5) | **Post** /api/v5/tradingBot/signal/cancel-sub-order | Cancel an incomplete order.  
[**CreateTradingBotSignalClosePositionV5**](SignalBotTradingAPI.md#CreateTradingBotSignalClosePositionV5) | **Post** /api/v5/tradingBot/signal/close-position | Close the position of an instrument via a market order.  
[**CreateTradingBotSignalCreateSignalV5**](SignalBotTradingAPI.md#CreateTradingBotSignalCreateSignalV5) | **Post** /api/v5/tradingBot/signal/create-signal | 
[**CreateTradingBotSignalMarginBalanceV5**](SignalBotTradingAPI.md#CreateTradingBotSignalMarginBalanceV5) | **Post** /api/v5/tradingBot/signal/margin-balance | 
[**CreateTradingBotSignalOrderAlgoV5**](SignalBotTradingAPI.md#CreateTradingBotSignalOrderAlgoV5) | **Post** /api/v5/tradingBot/signal/order-algo | 
[**CreateTradingBotSignalSetInstrumentsV5**](SignalBotTradingAPI.md#CreateTradingBotSignalSetInstrumentsV5) | **Post** /api/v5/tradingBot/signal/set-instruments | 
[**CreateTradingBotSignalStopOrderAlgoV5**](SignalBotTradingAPI.md#CreateTradingBotSignalStopOrderAlgoV5) | **Post** /api/v5/tradingBot/signal/stop-order-algo | A maximum of 10 orders can be stopped per request.  
[**CreateTradingBotSignalSubOrderV5**](SignalBotTradingAPI.md#CreateTradingBotSignalSubOrderV5) | **Post** /api/v5/tradingBot/signal/sub-order | You can place an order only if you have sufficient funds.      
[**GetTradingBotSignalEventHistoryV5**](SignalBotTradingAPI.md#GetTradingBotSignalEventHistoryV5) | **Get** /api/v5/tradingBot/signal/event-history | 
[**GetTradingBotSignalOrdersAlgoDetailsV5**](SignalBotTradingAPI.md#GetTradingBotSignalOrdersAlgoDetailsV5) | **Get** /api/v5/tradingBot/signal/orders-algo-details | 
[**GetTradingBotSignalOrdersAlgoHistoryV5**](SignalBotTradingAPI.md#GetTradingBotSignalOrdersAlgoHistoryV5) | **Get** /api/v5/tradingBot/signal/orders-algo-history | 
[**GetTradingBotSignalOrdersAlgoPendingV5**](SignalBotTradingAPI.md#GetTradingBotSignalOrdersAlgoPendingV5) | **Get** /api/v5/tradingBot/signal/orders-algo-pending | 
[**GetTradingBotSignalPositionsHistoryV5**](SignalBotTradingAPI.md#GetTradingBotSignalPositionsHistoryV5) | **Get** /api/v5/tradingBot/signal/positions-history | Retrieve the updated position data for the last 3 months. Return in reverse chronological order using utime.  
[**GetTradingBotSignalPositionsV5**](SignalBotTradingAPI.md#GetTradingBotSignalPositionsV5) | **Get** /api/v5/tradingBot/signal/positions | 
[**GetTradingBotSignalSignalsV5**](SignalBotTradingAPI.md#GetTradingBotSignalSignalsV5) | **Get** /api/v5/tradingBot/signal/signals | 
[**GetTradingBotSignalSubOrdersV5**](SignalBotTradingAPI.md#GetTradingBotSignalSubOrdersV5) | **Get** /api/v5/tradingBot/signal/sub-orders | 



## CreateTradingBotSignalAmendTPSLV5

> CreateTradingBotSignalAmendTPSLV5Resp CreateTradingBotSignalAmendTPSLV5(ctx).CreateTradingBotSignalAmendTPSLV5Req(createTradingBotSignalAmendTPSLV5Req).Execute()





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
	createTradingBotSignalAmendTPSLV5Req := *openapiclient.NewCreateTradingBotSignalAmendTPSLV5Req("AlgoId_example", "ExitSettingParam_example") // CreateTradingBotSignalAmendTPSLV5Req | The request body for CreateTradingBotSignalAmendTPSLV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.CreateTradingBotSignalAmendTPSLV5(context.Background()).CreateTradingBotSignalAmendTPSLV5Req(createTradingBotSignalAmendTPSLV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.CreateTradingBotSignalAmendTPSLV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotSignalAmendTPSLV5`: CreateTradingBotSignalAmendTPSLV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.CreateTradingBotSignalAmendTPSLV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotSignalAmendTPSLV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotSignalAmendTPSLV5Req** | [**CreateTradingBotSignalAmendTPSLV5Req**](CreateTradingBotSignalAmendTPSLV5Req.md) | The request body for CreateTradingBotSignalAmendTPSLV5 | 

### Return type

[**CreateTradingBotSignalAmendTPSLV5Resp**](CreateTradingBotSignalAmendTPSLV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotSignalCancelSubOrderV5

> CreateTradingBotSignalCancelSubOrderV5Resp CreateTradingBotSignalCancelSubOrderV5(ctx).CreateTradingBotSignalCancelSubOrderV5Req(createTradingBotSignalCancelSubOrderV5Req).Execute()

Cancel an incomplete order.  



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
	createTradingBotSignalCancelSubOrderV5Req := *openapiclient.NewCreateTradingBotSignalCancelSubOrderV5Req("AlgoId_example", "InstId_example", "SignalOrdId_example") // CreateTradingBotSignalCancelSubOrderV5Req | The request body for CreateTradingBotSignalCancelSubOrderV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.CreateTradingBotSignalCancelSubOrderV5(context.Background()).CreateTradingBotSignalCancelSubOrderV5Req(createTradingBotSignalCancelSubOrderV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.CreateTradingBotSignalCancelSubOrderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotSignalCancelSubOrderV5`: CreateTradingBotSignalCancelSubOrderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.CreateTradingBotSignalCancelSubOrderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotSignalCancelSubOrderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotSignalCancelSubOrderV5Req** | [**CreateTradingBotSignalCancelSubOrderV5Req**](CreateTradingBotSignalCancelSubOrderV5Req.md) | The request body for CreateTradingBotSignalCancelSubOrderV5 | 

### Return type

[**CreateTradingBotSignalCancelSubOrderV5Resp**](CreateTradingBotSignalCancelSubOrderV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotSignalClosePositionV5

> CreateTradingBotSignalClosePositionV5Resp CreateTradingBotSignalClosePositionV5(ctx).CreateTradingBotSignalClosePositionV5Req(createTradingBotSignalClosePositionV5Req).Execute()

Close the position of an instrument via a market order.  



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
	createTradingBotSignalClosePositionV5Req := *openapiclient.NewCreateTradingBotSignalClosePositionV5Req("AlgoId_example", "InstId_example") // CreateTradingBotSignalClosePositionV5Req | The request body for CreateTradingBotSignalClosePositionV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.CreateTradingBotSignalClosePositionV5(context.Background()).CreateTradingBotSignalClosePositionV5Req(createTradingBotSignalClosePositionV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.CreateTradingBotSignalClosePositionV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotSignalClosePositionV5`: CreateTradingBotSignalClosePositionV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.CreateTradingBotSignalClosePositionV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotSignalClosePositionV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotSignalClosePositionV5Req** | [**CreateTradingBotSignalClosePositionV5Req**](CreateTradingBotSignalClosePositionV5Req.md) | The request body for CreateTradingBotSignalClosePositionV5 | 

### Return type

[**CreateTradingBotSignalClosePositionV5Resp**](CreateTradingBotSignalClosePositionV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotSignalCreateSignalV5

> CreateTradingBotSignalCreateSignalV5Resp CreateTradingBotSignalCreateSignalV5(ctx).CreateTradingBotSignalCreateSignalV5Req(createTradingBotSignalCreateSignalV5Req).Execute()





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
	createTradingBotSignalCreateSignalV5Req := *openapiclient.NewCreateTradingBotSignalCreateSignalV5Req("SignalChanName_example") // CreateTradingBotSignalCreateSignalV5Req | The request body for CreateTradingBotSignalCreateSignalV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.CreateTradingBotSignalCreateSignalV5(context.Background()).CreateTradingBotSignalCreateSignalV5Req(createTradingBotSignalCreateSignalV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.CreateTradingBotSignalCreateSignalV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotSignalCreateSignalV5`: CreateTradingBotSignalCreateSignalV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.CreateTradingBotSignalCreateSignalV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotSignalCreateSignalV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotSignalCreateSignalV5Req** | [**CreateTradingBotSignalCreateSignalV5Req**](CreateTradingBotSignalCreateSignalV5Req.md) | The request body for CreateTradingBotSignalCreateSignalV5 | 

### Return type

[**CreateTradingBotSignalCreateSignalV5Resp**](CreateTradingBotSignalCreateSignalV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotSignalMarginBalanceV5

> CreateTradingBotSignalMarginBalanceV5Resp CreateTradingBotSignalMarginBalanceV5(ctx).CreateTradingBotSignalMarginBalanceV5Req(createTradingBotSignalMarginBalanceV5Req).Execute()





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
	createTradingBotSignalMarginBalanceV5Req := *openapiclient.NewCreateTradingBotSignalMarginBalanceV5Req("AlgoId_example", "Amt_example", "Type_example") // CreateTradingBotSignalMarginBalanceV5Req | The request body for CreateTradingBotSignalMarginBalanceV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.CreateTradingBotSignalMarginBalanceV5(context.Background()).CreateTradingBotSignalMarginBalanceV5Req(createTradingBotSignalMarginBalanceV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.CreateTradingBotSignalMarginBalanceV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotSignalMarginBalanceV5`: CreateTradingBotSignalMarginBalanceV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.CreateTradingBotSignalMarginBalanceV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotSignalMarginBalanceV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotSignalMarginBalanceV5Req** | [**CreateTradingBotSignalMarginBalanceV5Req**](CreateTradingBotSignalMarginBalanceV5Req.md) | The request body for CreateTradingBotSignalMarginBalanceV5 | 

### Return type

[**CreateTradingBotSignalMarginBalanceV5Resp**](CreateTradingBotSignalMarginBalanceV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotSignalOrderAlgoV5

> CreateTradingBotSignalOrderAlgoV5Resp CreateTradingBotSignalOrderAlgoV5(ctx).CreateTradingBotSignalOrderAlgoV5Req(createTradingBotSignalOrderAlgoV5Req).Execute()





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
	createTradingBotSignalOrderAlgoV5Req := *openapiclient.NewCreateTradingBotSignalOrderAlgoV5Req("InvestAmt_example", "Lever_example", "SignalChanId_example", "SubOrdType_example") // CreateTradingBotSignalOrderAlgoV5Req | The request body for CreateTradingBotSignalOrderAlgoV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.CreateTradingBotSignalOrderAlgoV5(context.Background()).CreateTradingBotSignalOrderAlgoV5Req(createTradingBotSignalOrderAlgoV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.CreateTradingBotSignalOrderAlgoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotSignalOrderAlgoV5`: CreateTradingBotSignalOrderAlgoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.CreateTradingBotSignalOrderAlgoV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotSignalOrderAlgoV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotSignalOrderAlgoV5Req** | [**CreateTradingBotSignalOrderAlgoV5Req**](CreateTradingBotSignalOrderAlgoV5Req.md) | The request body for CreateTradingBotSignalOrderAlgoV5 | 

### Return type

[**CreateTradingBotSignalOrderAlgoV5Resp**](CreateTradingBotSignalOrderAlgoV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotSignalSetInstrumentsV5

> CreateTradingBotSignalSetInstrumentsV5Resp CreateTradingBotSignalSetInstrumentsV5(ctx).CreateTradingBotSignalSetInstrumentsV5Req(createTradingBotSignalSetInstrumentsV5Req).Execute()





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
	createTradingBotSignalSetInstrumentsV5Req := *openapiclient.NewCreateTradingBotSignalSetInstrumentsV5Req("AlgoId_example", false, []string{"InstIds_example"}) // CreateTradingBotSignalSetInstrumentsV5Req | The request body for CreateTradingBotSignalSetInstrumentsV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.CreateTradingBotSignalSetInstrumentsV5(context.Background()).CreateTradingBotSignalSetInstrumentsV5Req(createTradingBotSignalSetInstrumentsV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.CreateTradingBotSignalSetInstrumentsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotSignalSetInstrumentsV5`: CreateTradingBotSignalSetInstrumentsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.CreateTradingBotSignalSetInstrumentsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotSignalSetInstrumentsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotSignalSetInstrumentsV5Req** | [**CreateTradingBotSignalSetInstrumentsV5Req**](CreateTradingBotSignalSetInstrumentsV5Req.md) | The request body for CreateTradingBotSignalSetInstrumentsV5 | 

### Return type

[**CreateTradingBotSignalSetInstrumentsV5Resp**](CreateTradingBotSignalSetInstrumentsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotSignalStopOrderAlgoV5

> CreateTradingBotSignalStopOrderAlgoV5Resp CreateTradingBotSignalStopOrderAlgoV5(ctx).CreateTradingBotSignalStopOrderAlgoV5Req(createTradingBotSignalStopOrderAlgoV5Req).Execute()

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
	createTradingBotSignalStopOrderAlgoV5Req := *openapiclient.NewCreateTradingBotSignalStopOrderAlgoV5Req("AlgoId_example") // CreateTradingBotSignalStopOrderAlgoV5Req | The request body for CreateTradingBotSignalStopOrderAlgoV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.CreateTradingBotSignalStopOrderAlgoV5(context.Background()).CreateTradingBotSignalStopOrderAlgoV5Req(createTradingBotSignalStopOrderAlgoV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.CreateTradingBotSignalStopOrderAlgoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotSignalStopOrderAlgoV5`: CreateTradingBotSignalStopOrderAlgoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.CreateTradingBotSignalStopOrderAlgoV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotSignalStopOrderAlgoV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotSignalStopOrderAlgoV5Req** | [**CreateTradingBotSignalStopOrderAlgoV5Req**](CreateTradingBotSignalStopOrderAlgoV5Req.md) | The request body for CreateTradingBotSignalStopOrderAlgoV5 | 

### Return type

[**CreateTradingBotSignalStopOrderAlgoV5Resp**](CreateTradingBotSignalStopOrderAlgoV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotSignalSubOrderV5

> CreateTradingBotSignalSubOrderV5Resp CreateTradingBotSignalSubOrderV5(ctx).CreateTradingBotSignalSubOrderV5Req(createTradingBotSignalSubOrderV5Req).Execute()

You can place an order only if you have sufficient funds.      



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
	createTradingBotSignalSubOrderV5Req := *openapiclient.NewCreateTradingBotSignalSubOrderV5Req("AlgoId_example", "InstId_example", "OrdType_example", "Side_example", "Sz_example") // CreateTradingBotSignalSubOrderV5Req | The request body for CreateTradingBotSignalSubOrderV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.CreateTradingBotSignalSubOrderV5(context.Background()).CreateTradingBotSignalSubOrderV5Req(createTradingBotSignalSubOrderV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.CreateTradingBotSignalSubOrderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotSignalSubOrderV5`: CreateTradingBotSignalSubOrderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.CreateTradingBotSignalSubOrderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotSignalSubOrderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotSignalSubOrderV5Req** | [**CreateTradingBotSignalSubOrderV5Req**](CreateTradingBotSignalSubOrderV5Req.md) | The request body for CreateTradingBotSignalSubOrderV5 | 

### Return type

[**CreateTradingBotSignalSubOrderV5Resp**](CreateTradingBotSignalSubOrderV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotSignalEventHistoryV5

> GetTradingBotSignalEventHistoryV5Resp GetTradingBotSignalEventHistoryV5(ctx).AlgoId(algoId).After(after).Before(before).Limit(limit).Execute()





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
	algoId := "algoId_example" // string | Algo ID (default to "")
	after := "after_example" // string | Pagination of data to return records `eventCtime` earlier than the requested timestamp, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records `eventCtime` newer than the requested timestamp, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.GetTradingBotSignalEventHistoryV5(context.Background()).AlgoId(algoId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.GetTradingBotSignalEventHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotSignalEventHistoryV5`: GetTradingBotSignalEventHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.GetTradingBotSignalEventHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotSignalEventHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records &#x60;eventCtime&#x60; earlier than the requested timestamp, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records &#x60;eventCtime&#x60; newer than the requested timestamp, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetTradingBotSignalEventHistoryV5Resp**](GetTradingBotSignalEventHistoryV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotSignalOrdersAlgoDetailsV5

> GetTradingBotSignalOrdersAlgoDetailsV5Resp GetTradingBotSignalOrdersAlgoDetailsV5(ctx).AlgoOrdType(algoOrdType).AlgoId(algoId).Execute()





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
	algoOrdType := "algoOrdType_example" // string | Algo order type  `contract`: Contract signal (default to "")
	algoId := "algoId_example" // string | Algo ID (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.GetTradingBotSignalOrdersAlgoDetailsV5(context.Background()).AlgoOrdType(algoOrdType).AlgoId(algoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.GetTradingBotSignalOrdersAlgoDetailsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotSignalOrdersAlgoDetailsV5`: GetTradingBotSignalOrdersAlgoDetailsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.GetTradingBotSignalOrdersAlgoDetailsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotSignalOrdersAlgoDetailsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoOrdType** | **string** | Algo order type  &#x60;contract&#x60;: Contract signal | [default to &quot;&quot;]
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]

### Return type

[**GetTradingBotSignalOrdersAlgoDetailsV5Resp**](GetTradingBotSignalOrdersAlgoDetailsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotSignalOrdersAlgoHistoryV5

> GetTradingBotSignalOrdersAlgoHistoryV5Resp GetTradingBotSignalOrdersAlgoHistoryV5(ctx).AlgoOrdType(algoOrdType).AlgoId(algoId).After(after).Before(before).Limit(limit).Execute()





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
	algoOrdType := "algoOrdType_example" // string | Algo order type  `contract`: Contract signal (default to "")
	algoId := "algoId_example" // string | Algo ID (default to "")
	after := "after_example" // string | Pagination of data to return records `algoId` earlier than the requested timestamp, Unix timestamp format in milliseconds, e.g. `1597026383085` (default to "")
	before := "before_example" // string | Pagination of data to return records `algoId` newer than the requested timestamp, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.GetTradingBotSignalOrdersAlgoHistoryV5(context.Background()).AlgoOrdType(algoOrdType).AlgoId(algoId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.GetTradingBotSignalOrdersAlgoHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotSignalOrdersAlgoHistoryV5`: GetTradingBotSignalOrdersAlgoHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.GetTradingBotSignalOrdersAlgoHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotSignalOrdersAlgoHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoOrdType** | **string** | Algo order type  &#x60;contract&#x60;: Contract signal | [default to &quot;&quot;]
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records &#x60;algoId&#x60; earlier than the requested timestamp, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records &#x60;algoId&#x60; newer than the requested timestamp, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetTradingBotSignalOrdersAlgoHistoryV5Resp**](GetTradingBotSignalOrdersAlgoHistoryV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotSignalOrdersAlgoPendingV5

> GetTradingBotSignalOrdersAlgoPendingV5Resp GetTradingBotSignalOrdersAlgoPendingV5(ctx).AlgoOrdType(algoOrdType).After(after).AlgoId(algoId).Before(before).Limit(limit).Execute()





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
	algoOrdType := "algoOrdType_example" // string | Algo order type  `contract`: Contract signal (default to "")
	after := "after_example" // string | Pagination of data to return records `algoId` earlier than the requested timestamp, Unix timestamp format in milliseconds, e.g. `1597026383085` (default to "")
	algoId := "algoId_example" // string | Algo ID (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records `algoId` newer than the requested timestamp, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.GetTradingBotSignalOrdersAlgoPendingV5(context.Background()).AlgoOrdType(algoOrdType).After(after).AlgoId(algoId).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.GetTradingBotSignalOrdersAlgoPendingV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotSignalOrdersAlgoPendingV5`: GetTradingBotSignalOrdersAlgoPendingV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.GetTradingBotSignalOrdersAlgoPendingV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotSignalOrdersAlgoPendingV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoOrdType** | **string** | Algo order type  &#x60;contract&#x60;: Contract signal | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records &#x60;algoId&#x60; earlier than the requested timestamp, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records &#x60;algoId&#x60; newer than the requested timestamp, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetTradingBotSignalOrdersAlgoPendingV5Resp**](GetTradingBotSignalOrdersAlgoPendingV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotSignalPositionsHistoryV5

> GetTradingBotSignalPositionsHistoryV5Resp GetTradingBotSignalPositionsHistoryV5(ctx).AlgoId(algoId).InstId(instId).After(after).Before(before).Limit(limit).Execute()

Retrieve the updated position data for the last 3 months. Return in reverse chronological order using utime.  



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
	algoId := "algoId_example" // string | Algo ID (default to "")
	instId := "instId_example" // string | Instrument ID, e.g.：`BTC-USD-SWAP` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `uTime`, Unix timestamp format in milliseconds, e.g.`1597026383085` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `uTime`, Unix timestamp format in milliseconds, e.g `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.GetTradingBotSignalPositionsHistoryV5(context.Background()).AlgoId(algoId).InstId(instId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.GetTradingBotSignalPositionsHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotSignalPositionsHistoryV5`: GetTradingBotSignalPositionsHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.GetTradingBotSignalPositionsHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotSignalPositionsHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g.：&#x60;BTC-USD-SWAP&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;uTime&#x60;, Unix timestamp format in milliseconds, e.g.&#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;uTime&#x60;, Unix timestamp format in milliseconds, e.g &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetTradingBotSignalPositionsHistoryV5Resp**](GetTradingBotSignalPositionsHistoryV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotSignalPositionsV5

> GetTradingBotSignalPositionsV5Resp GetTradingBotSignalPositionsV5(ctx).AlgoOrdType(algoOrdType).AlgoId(algoId).Execute()





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
	algoOrdType := "algoOrdType_example" // string | Algo order type  `contract`: Contract signal (default to "")
	algoId := "algoId_example" // string | Algo ID (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.GetTradingBotSignalPositionsV5(context.Background()).AlgoOrdType(algoOrdType).AlgoId(algoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.GetTradingBotSignalPositionsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotSignalPositionsV5`: GetTradingBotSignalPositionsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.GetTradingBotSignalPositionsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotSignalPositionsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoOrdType** | **string** | Algo order type  &#x60;contract&#x60;: Contract signal | [default to &quot;&quot;]
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]

### Return type

[**GetTradingBotSignalPositionsV5Resp**](GetTradingBotSignalPositionsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotSignalSignalsV5

> GetTradingBotSignalSignalsV5Resp GetTradingBotSignalSignalsV5(ctx).SignalSourceType(signalSourceType).SignalChanId(signalChanId).After(after).Before(before).Limit(limit).Execute()





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
	signalSourceType := "signalSourceType_example" // string | Signal source type  `1`: Created by yourself  `2`: Subscribe  `3`: Free signal (default to "")
	signalChanId := "signalChanId_example" // string | Signal channel id (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records `signalChanId` earlier than the requested timestamp, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records `signalChanId` newer than the requested timestamp, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.GetTradingBotSignalSignalsV5(context.Background()).SignalSourceType(signalSourceType).SignalChanId(signalChanId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.GetTradingBotSignalSignalsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotSignalSignalsV5`: GetTradingBotSignalSignalsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.GetTradingBotSignalSignalsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotSignalSignalsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signalSourceType** | **string** | Signal source type  &#x60;1&#x60;: Created by yourself  &#x60;2&#x60;: Subscribe  &#x60;3&#x60;: Free signal | [default to &quot;&quot;]
 **signalChanId** | **string** | Signal channel id | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records &#x60;signalChanId&#x60; earlier than the requested timestamp, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records &#x60;signalChanId&#x60; newer than the requested timestamp, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetTradingBotSignalSignalsV5Resp**](GetTradingBotSignalSignalsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotSignalSubOrdersV5

> GetTradingBotSignalSubOrdersV5Resp GetTradingBotSignalSubOrdersV5(ctx).AlgoId(algoId).AlgoOrdType(algoOrdType).State(state).SignalOrdId(signalOrdId).After(after).Before(before).Begin(begin).End(end).Limit(limit).Type_(type_).ClOrdId(clOrdId).Execute()





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
	algoId := "algoId_example" // string | Algo ID (default to "")
	algoOrdType := "algoOrdType_example" // string | Algo order type  `contract`: Contract signal (default to "")
	state := "state_example" // string | Sub order state  `live`  `partially_filled`  `filled`  `cancelled`  Either `state` or `signalOrdId` is required, if both are passed in, only `state` is valid. (optional) (default to "")
	signalOrdId := "signalOrdId_example" // string | Sub order ID (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ordId` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ordId`. (optional) (default to "")
	begin := "begin_example" // string | Return records of `ctime` after than the requested timestamp (include), Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	end := "end_example" // string | Return records of `ctime` before than the requested timestamp (include), Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")
	type_ := "type__example" // string | Sub order type   `live`  `filled`  Either `type` or `clOrdId` is required, if both are passed in, only `clOrdId` is valid. (optional) (default to "")
	clOrdId := "clOrdId_example" // string | Sub order client-supplied ID.   `It will be deprecated soon` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignalBotTradingAPI.GetTradingBotSignalSubOrdersV5(context.Background()).AlgoId(algoId).AlgoOrdType(algoOrdType).State(state).SignalOrdId(signalOrdId).After(after).Before(before).Begin(begin).End(end).Limit(limit).Type_(type_).ClOrdId(clOrdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignalBotTradingAPI.GetTradingBotSignalSubOrdersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotSignalSubOrdersV5`: GetTradingBotSignalSubOrdersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SignalBotTradingAPI.GetTradingBotSignalSubOrdersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotSignalSubOrdersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]
 **algoOrdType** | **string** | Algo order type  &#x60;contract&#x60;: Contract signal | [default to &quot;&quot;]
 **state** | **string** | Sub order state  &#x60;live&#x60;  &#x60;partially_filled&#x60;  &#x60;filled&#x60;  &#x60;cancelled&#x60;  Either &#x60;state&#x60; or &#x60;signalOrdId&#x60; is required, if both are passed in, only &#x60;state&#x60; is valid. | [default to &quot;&quot;]
 **signalOrdId** | **string** | Sub order ID | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ordId&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ordId&#x60;. | [default to &quot;&quot;]
 **begin** | **string** | Return records of &#x60;ctime&#x60; after than the requested timestamp (include), Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **end** | **string** | Return records of &#x60;ctime&#x60; before than the requested timestamp (include), Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]
 **type_** | **string** | Sub order type   &#x60;live&#x60;  &#x60;filled&#x60;  Either &#x60;type&#x60; or &#x60;clOrdId&#x60; is required, if both are passed in, only &#x60;clOrdId&#x60; is valid. | [default to &quot;&quot;]
 **clOrdId** | **string** | Sub order client-supplied ID.   &#x60;It will be deprecated soon&#x60; | [default to &quot;&quot;]

### Return type

[**GetTradingBotSignalSubOrdersV5Resp**](GetTradingBotSignalSubOrdersV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

