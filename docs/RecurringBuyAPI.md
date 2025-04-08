# \RecurringBuyAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTradingBotRecurringAmendOrderAlgoV5**](RecurringBuyAPI.md#CreateTradingBotRecurringAmendOrderAlgoV5) | **Post** /api/v5/tradingBot/recurring/amend-order-algo | 
[**CreateTradingBotRecurringOrderAlgoV5**](RecurringBuyAPI.md#CreateTradingBotRecurringOrderAlgoV5) | **Post** /api/v5/tradingBot/recurring/order-algo | 
[**CreateTradingBotRecurringStopOrderAlgoV5**](RecurringBuyAPI.md#CreateTradingBotRecurringStopOrderAlgoV5) | **Post** /api/v5/tradingBot/recurring/stop-order-algo | A maximum of 10 orders can be stopped per request.  
[**GetTradingBotRecurringOrdersAlgoDetailsV5**](RecurringBuyAPI.md#GetTradingBotRecurringOrdersAlgoDetailsV5) | **Get** /api/v5/tradingBot/recurring/orders-algo-details | 
[**GetTradingBotRecurringOrdersAlgoHistoryV5**](RecurringBuyAPI.md#GetTradingBotRecurringOrdersAlgoHistoryV5) | **Get** /api/v5/tradingBot/recurring/orders-algo-history | 
[**GetTradingBotRecurringOrdersAlgoPendingV5**](RecurringBuyAPI.md#GetTradingBotRecurringOrdersAlgoPendingV5) | **Get** /api/v5/tradingBot/recurring/orders-algo-pending | 
[**GetTradingBotRecurringSubOrdersV5**](RecurringBuyAPI.md#GetTradingBotRecurringSubOrdersV5) | **Get** /api/v5/tradingBot/recurring/sub-orders | 



## CreateTradingBotRecurringAmendOrderAlgoV5

> CreateTradingBotRecurringAmendOrderAlgoV5Resp CreateTradingBotRecurringAmendOrderAlgoV5(ctx).CreateTradingBotRecurringAmendOrderAlgoV5Req(createTradingBotRecurringAmendOrderAlgoV5Req).Execute()





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
	createTradingBotRecurringAmendOrderAlgoV5Req := *openapiclient.NewCreateTradingBotRecurringAmendOrderAlgoV5Req("AlgoId_example", "StgyName_example") // CreateTradingBotRecurringAmendOrderAlgoV5Req | The request body for CreateTradingBotRecurringAmendOrderAlgoV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringBuyAPI.CreateTradingBotRecurringAmendOrderAlgoV5(context.Background()).CreateTradingBotRecurringAmendOrderAlgoV5Req(createTradingBotRecurringAmendOrderAlgoV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringBuyAPI.CreateTradingBotRecurringAmendOrderAlgoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotRecurringAmendOrderAlgoV5`: CreateTradingBotRecurringAmendOrderAlgoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `RecurringBuyAPI.CreateTradingBotRecurringAmendOrderAlgoV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotRecurringAmendOrderAlgoV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotRecurringAmendOrderAlgoV5Req** | [**CreateTradingBotRecurringAmendOrderAlgoV5Req**](CreateTradingBotRecurringAmendOrderAlgoV5Req.md) | The request body for CreateTradingBotRecurringAmendOrderAlgoV5 | 

### Return type

[**CreateTradingBotRecurringAmendOrderAlgoV5Resp**](CreateTradingBotRecurringAmendOrderAlgoV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotRecurringOrderAlgoV5

> CreateTradingBotRecurringOrderAlgoV5Resp CreateTradingBotRecurringOrderAlgoV5(ctx).CreateTradingBotRecurringOrderAlgoV5Req(createTradingBotRecurringOrderAlgoV5Req).Execute()





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
	createTradingBotRecurringOrderAlgoV5Req := *openapiclient.NewCreateTradingBotRecurringOrderAlgoV5Req("Amt_example", "InvestmentCcy_example", "Period_example", []openapiclient.CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner{*openapiclient.NewCreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner()}, "RecurringTime_example", "StgyName_example", "TdMode_example", "TimeZone_example") // CreateTradingBotRecurringOrderAlgoV5Req | The request body for CreateTradingBotRecurringOrderAlgoV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringBuyAPI.CreateTradingBotRecurringOrderAlgoV5(context.Background()).CreateTradingBotRecurringOrderAlgoV5Req(createTradingBotRecurringOrderAlgoV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringBuyAPI.CreateTradingBotRecurringOrderAlgoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotRecurringOrderAlgoV5`: CreateTradingBotRecurringOrderAlgoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `RecurringBuyAPI.CreateTradingBotRecurringOrderAlgoV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotRecurringOrderAlgoV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotRecurringOrderAlgoV5Req** | [**CreateTradingBotRecurringOrderAlgoV5Req**](CreateTradingBotRecurringOrderAlgoV5Req.md) | The request body for CreateTradingBotRecurringOrderAlgoV5 | 

### Return type

[**CreateTradingBotRecurringOrderAlgoV5Resp**](CreateTradingBotRecurringOrderAlgoV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradingBotRecurringStopOrderAlgoV5

> CreateTradingBotRecurringStopOrderAlgoV5Resp CreateTradingBotRecurringStopOrderAlgoV5(ctx).CreateTradingBotRecurringStopOrderAlgoV5Req(createTradingBotRecurringStopOrderAlgoV5Req).Execute()

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
	createTradingBotRecurringStopOrderAlgoV5Req := *openapiclient.NewCreateTradingBotRecurringStopOrderAlgoV5Req("AlgoId_example") // CreateTradingBotRecurringStopOrderAlgoV5Req | The request body for CreateTradingBotRecurringStopOrderAlgoV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringBuyAPI.CreateTradingBotRecurringStopOrderAlgoV5(context.Background()).CreateTradingBotRecurringStopOrderAlgoV5Req(createTradingBotRecurringStopOrderAlgoV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringBuyAPI.CreateTradingBotRecurringStopOrderAlgoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradingBotRecurringStopOrderAlgoV5`: CreateTradingBotRecurringStopOrderAlgoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `RecurringBuyAPI.CreateTradingBotRecurringStopOrderAlgoV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradingBotRecurringStopOrderAlgoV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradingBotRecurringStopOrderAlgoV5Req** | [**CreateTradingBotRecurringStopOrderAlgoV5Req**](CreateTradingBotRecurringStopOrderAlgoV5Req.md) | The request body for CreateTradingBotRecurringStopOrderAlgoV5 | 

### Return type

[**CreateTradingBotRecurringStopOrderAlgoV5Resp**](CreateTradingBotRecurringStopOrderAlgoV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotRecurringOrdersAlgoDetailsV5

> GetTradingBotRecurringOrdersAlgoDetailsV5Resp GetTradingBotRecurringOrdersAlgoDetailsV5(ctx).AlgoId(algoId).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringBuyAPI.GetTradingBotRecurringOrdersAlgoDetailsV5(context.Background()).AlgoId(algoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringBuyAPI.GetTradingBotRecurringOrdersAlgoDetailsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotRecurringOrdersAlgoDetailsV5`: GetTradingBotRecurringOrdersAlgoDetailsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `RecurringBuyAPI.GetTradingBotRecurringOrdersAlgoDetailsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotRecurringOrdersAlgoDetailsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]

### Return type

[**GetTradingBotRecurringOrdersAlgoDetailsV5Resp**](GetTradingBotRecurringOrdersAlgoDetailsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotRecurringOrdersAlgoHistoryV5

> GetTradingBotRecurringOrdersAlgoHistoryV5Resp GetTradingBotRecurringOrdersAlgoHistoryV5(ctx).AlgoId(algoId).After(after).Before(before).Limit(limit).Execute()





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
	algoId := "algoId_example" // string | Algo ID (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `algoId`. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `algoId`. (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringBuyAPI.GetTradingBotRecurringOrdersAlgoHistoryV5(context.Background()).AlgoId(algoId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringBuyAPI.GetTradingBotRecurringOrdersAlgoHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotRecurringOrdersAlgoHistoryV5`: GetTradingBotRecurringOrdersAlgoHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `RecurringBuyAPI.GetTradingBotRecurringOrdersAlgoHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotRecurringOrdersAlgoHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;algoId&#x60;. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;algoId&#x60;. | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100 | [default to &quot;&quot;]

### Return type

[**GetTradingBotRecurringOrdersAlgoHistoryV5Resp**](GetTradingBotRecurringOrdersAlgoHistoryV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotRecurringOrdersAlgoPendingV5

> GetTradingBotRecurringOrdersAlgoPendingV5Resp GetTradingBotRecurringOrdersAlgoPendingV5(ctx).AlgoId(algoId).After(after).Before(before).Limit(limit).Execute()





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
	algoId := "algoId_example" // string | Algo ID (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `algoId`. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `algoId`. (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringBuyAPI.GetTradingBotRecurringOrdersAlgoPendingV5(context.Background()).AlgoId(algoId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringBuyAPI.GetTradingBotRecurringOrdersAlgoPendingV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotRecurringOrdersAlgoPendingV5`: GetTradingBotRecurringOrdersAlgoPendingV5Resp
	fmt.Fprintf(os.Stdout, "Response from `RecurringBuyAPI.GetTradingBotRecurringOrdersAlgoPendingV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotRecurringOrdersAlgoPendingV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;algoId&#x60;. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;algoId&#x60;. | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100 | [default to &quot;&quot;]

### Return type

[**GetTradingBotRecurringOrdersAlgoPendingV5Resp**](GetTradingBotRecurringOrdersAlgoPendingV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradingBotRecurringSubOrdersV5

> GetTradingBotRecurringSubOrdersV5Resp GetTradingBotRecurringSubOrdersV5(ctx).AlgoId(algoId).OrdId(ordId).After(after).Before(before).Limit(limit).Execute()





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
	ordId := "ordId_example" // string | Sub order ID (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `algoId`. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `algoId`. (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringBuyAPI.GetTradingBotRecurringSubOrdersV5(context.Background()).AlgoId(algoId).OrdId(ordId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringBuyAPI.GetTradingBotRecurringSubOrdersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradingBotRecurringSubOrdersV5`: GetTradingBotRecurringSubOrdersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `RecurringBuyAPI.GetTradingBotRecurringSubOrdersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradingBotRecurringSubOrdersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]
 **ordId** | **string** | Sub order ID | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;algoId&#x60;. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;algoId&#x60;. | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100 | [default to &quot;&quot;]

### Return type

[**GetTradingBotRecurringSubOrdersV5Resp**](GetTradingBotRecurringSubOrdersV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

