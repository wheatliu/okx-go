# \AlgoTradingAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTradeAmendAlgosV5**](AlgoTradingAPI.md#CreateTradeAmendAlgosV5) | **Post** /api/v5/trade/amend-algos | POST / Amend algo order
[**CreateTradeCancelAdvanceAlgosV5**](AlgoTradingAPI.md#CreateTradeCancelAdvanceAlgosV5) | **Post** /api/v5/trade/cancel-advance-algos | POST / Cancel advance algo order
[**CreateTradeCancelAlgosV5**](AlgoTradingAPI.md#CreateTradeCancelAlgosV5) | **Post** /api/v5/trade/cancel-algos | POST / Cancel algo order
[**CreateTradeOrderAlgoV5**](AlgoTradingAPI.md#CreateTradeOrderAlgoV5) | **Post** /api/v5/trade/order-algo | POST / Place algo order
[**GetTradeOrderAlgoV5**](AlgoTradingAPI.md#GetTradeOrderAlgoV5) | **Get** /api/v5/trade/order-algo | GET / Algo order details
[**GetTradeOrdersAlgoHistoryV5**](AlgoTradingAPI.md#GetTradeOrdersAlgoHistoryV5) | **Get** /api/v5/trade/orders-algo-history | GET / Algo order history
[**GetTradeOrdersAlgoPendingV5**](AlgoTradingAPI.md#GetTradeOrdersAlgoPendingV5) | **Get** /api/v5/trade/orders-algo-pending | GET / Algo order list



## CreateTradeAmendAlgosV5

> CreateTradeAmendAlgosV5Resp CreateTradeAmendAlgosV5(ctx).CreateTradeAmendAlgosV5Req(createTradeAmendAlgosV5Req).Execute()

POST / Amend algo order



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
	createTradeAmendAlgosV5Req := *openapiclient.NewCreateTradeAmendAlgosV5Req("InstId_example") // CreateTradeAmendAlgosV5Req | The request body for CreateTradeAmendAlgosV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.CreateTradeAmendAlgosV5(context.Background()).CreateTradeAmendAlgosV5Req(createTradeAmendAlgosV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.CreateTradeAmendAlgosV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeAmendAlgosV5`: CreateTradeAmendAlgosV5Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.CreateTradeAmendAlgosV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeAmendAlgosV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeAmendAlgosV5Req** | [**CreateTradeAmendAlgosV5Req**](CreateTradeAmendAlgosV5Req.md) | The request body for CreateTradeAmendAlgosV5 | 

### Return type

[**CreateTradeAmendAlgosV5Resp**](CreateTradeAmendAlgosV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeCancelAdvanceAlgosV5

> CreateTradeCancelAdvanceAlgosV5Resp CreateTradeCancelAdvanceAlgosV5(ctx).CreateTradeCancelAdvanceAlgosV5Req(createTradeCancelAdvanceAlgosV5Req).Execute()

POST / Cancel advance algo order



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
	createTradeCancelAdvanceAlgosV5Req := *openapiclient.NewCreateTradeCancelAdvanceAlgosV5Req("AlgoId_example", "InstId_example") // CreateTradeCancelAdvanceAlgosV5Req | The request body for CreateTradeCancelAdvanceAlgosV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.CreateTradeCancelAdvanceAlgosV5(context.Background()).CreateTradeCancelAdvanceAlgosV5Req(createTradeCancelAdvanceAlgosV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.CreateTradeCancelAdvanceAlgosV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeCancelAdvanceAlgosV5`: CreateTradeCancelAdvanceAlgosV5Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.CreateTradeCancelAdvanceAlgosV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeCancelAdvanceAlgosV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeCancelAdvanceAlgosV5Req** | [**CreateTradeCancelAdvanceAlgosV5Req**](CreateTradeCancelAdvanceAlgosV5Req.md) | The request body for CreateTradeCancelAdvanceAlgosV5 | 

### Return type

[**CreateTradeCancelAdvanceAlgosV5Resp**](CreateTradeCancelAdvanceAlgosV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeCancelAlgosV5

> CreateTradeCancelAlgosV5Resp CreateTradeCancelAlgosV5(ctx).CreateTradeCancelAlgosV5Req(createTradeCancelAlgosV5Req).Execute()

POST / Cancel algo order



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
	createTradeCancelAlgosV5Req := *openapiclient.NewCreateTradeCancelAlgosV5Req("AlgoId_example", "InstId_example") // CreateTradeCancelAlgosV5Req | The request body for CreateTradeCancelAlgosV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.CreateTradeCancelAlgosV5(context.Background()).CreateTradeCancelAlgosV5Req(createTradeCancelAlgosV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.CreateTradeCancelAlgosV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeCancelAlgosV5`: CreateTradeCancelAlgosV5Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.CreateTradeCancelAlgosV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeCancelAlgosV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeCancelAlgosV5Req** | [**CreateTradeCancelAlgosV5Req**](CreateTradeCancelAlgosV5Req.md) | The request body for CreateTradeCancelAlgosV5 | 

### Return type

[**CreateTradeCancelAlgosV5Resp**](CreateTradeCancelAlgosV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeOrderAlgoV5

> CreateTradeOrderAlgoV5Resp CreateTradeOrderAlgoV5(ctx).CreateTradeOrderAlgoV5Req(createTradeOrderAlgoV5Req).Execute()

POST / Place algo order



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
	createTradeOrderAlgoV5Req := *openapiclient.NewCreateTradeOrderAlgoV5Req("InstId_example", "OrdType_example", "Side_example", "TdMode_example") // CreateTradeOrderAlgoV5Req | The request body for CreateTradeOrderAlgoV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.CreateTradeOrderAlgoV5(context.Background()).CreateTradeOrderAlgoV5Req(createTradeOrderAlgoV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.CreateTradeOrderAlgoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeOrderAlgoV5`: CreateTradeOrderAlgoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.CreateTradeOrderAlgoV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeOrderAlgoV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeOrderAlgoV5Req** | [**CreateTradeOrderAlgoV5Req**](CreateTradeOrderAlgoV5Req.md) | The request body for CreateTradeOrderAlgoV5 | 

### Return type

[**CreateTradeOrderAlgoV5Resp**](CreateTradeOrderAlgoV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeOrderAlgoV5

> GetTradeOrderAlgoV5Resp GetTradeOrderAlgoV5(ctx).AlgoId(algoId).AlgoClOrdId(algoClOrdId).Execute()

GET / Algo order details



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
	algoId := "algoId_example" // string | Algo ID  Either `algoId` or `algoClOrdId` is required.If both are passed, `algoId` will be used. (optional) (default to "")
	algoClOrdId := "algoClOrdId_example" // string | Client-supplied Algo ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.GetTradeOrderAlgoV5(context.Background()).AlgoId(algoId).AlgoClOrdId(algoClOrdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.GetTradeOrderAlgoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeOrderAlgoV5`: GetTradeOrderAlgoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.GetTradeOrderAlgoV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeOrderAlgoV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **algoId** | **string** | Algo ID  Either &#x60;algoId&#x60; or &#x60;algoClOrdId&#x60; is required.If both are passed, &#x60;algoId&#x60; will be used. | [default to &quot;&quot;]
 **algoClOrdId** | **string** | Client-supplied Algo ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [default to &quot;&quot;]

### Return type

[**GetTradeOrderAlgoV5Resp**](GetTradeOrderAlgoV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeOrdersAlgoHistoryV5

> GetTradeOrdersAlgoHistoryV5Resp GetTradeOrdersAlgoHistoryV5(ctx).OrdType(ordType).State(state).AlgoId(algoId).InstType(instType).InstId(instId).After(after).Before(before).Limit(limit).Execute()

GET / Algo order history



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
	ordType := "ordType_example" // string | Order type    `conditional`: One-way stop order    `oco`: One-cancels-the-other order   `chase`: chase order, only applicable to FUTURES and SWAP  `trigger`: Trigger order   `move_order_stop`: Trailing order   `iceberg`: Iceberg order   `twap`: TWAP order   For every request, unlike other ordType which only can use one type, `conditional` and `oco` both can be used and separated with comma. (default to "")
	state := "state_example" // string | State  `effective`  `canceled`  `order_failed`  Either `state` or `algoId` is required (optional) (default to "")
	algoId := "algoId_example" // string | Algo ID   Either `state` or `algoId` is required. (optional) (default to "")
	instType := "instType_example" // string | Instrument type  `SPOT`  `SWAP`  `FUTURES`  `MARGIN` (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `algoId` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records new than the requested `algoId` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.GetTradeOrdersAlgoHistoryV5(context.Background()).OrdType(ordType).State(state).AlgoId(algoId).InstType(instType).InstId(instId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.GetTradeOrdersAlgoHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeOrdersAlgoHistoryV5`: GetTradeOrdersAlgoHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.GetTradeOrdersAlgoHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeOrdersAlgoHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordType** | **string** | Order type    &#x60;conditional&#x60;: One-way stop order    &#x60;oco&#x60;: One-cancels-the-other order   &#x60;chase&#x60;: chase order, only applicable to FUTURES and SWAP  &#x60;trigger&#x60;: Trigger order   &#x60;move_order_stop&#x60;: Trailing order   &#x60;iceberg&#x60;: Iceberg order   &#x60;twap&#x60;: TWAP order   For every request, unlike other ordType which only can use one type, &#x60;conditional&#x60; and &#x60;oco&#x60; both can be used and separated with comma. | [default to &quot;&quot;]
 **state** | **string** | State  &#x60;effective&#x60;  &#x60;canceled&#x60;  &#x60;order_failed&#x60;  Either &#x60;state&#x60; or &#x60;algoId&#x60; is required | [default to &quot;&quot;]
 **algoId** | **string** | Algo ID   Either &#x60;state&#x60; or &#x60;algoId&#x60; is required. | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;MARGIN&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;algoId&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records new than the requested &#x60;algoId&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetTradeOrdersAlgoHistoryV5Resp**](GetTradeOrdersAlgoHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeOrdersAlgoPendingV5

> GetTradeOrdersAlgoPendingV5Resp GetTradeOrdersAlgoPendingV5(ctx).OrdType(ordType).AlgoId(algoId).AlgoClOrdId(algoClOrdId).InstType(instType).InstId(instId).After(after).Before(before).Limit(limit).Execute()

GET / Algo order list



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
	ordType := "ordType_example" // string | Order type  `conditional`: One-way stop order    `oco`: One-cancels-the-other order   `chase`: chase order, only applicable to FUTURES and SWAP  `trigger`: Trigger order   `move_order_stop`: Trailing order   `iceberg`: Iceberg order   `twap`: TWAP order  For every request, unlike other ordType which only can use one type, `conditional` and `oco` both can be used and separated with comma. (default to "")
	algoId := "algoId_example" // string | Algo ID (optional) (default to "")
	algoClOrdId := "algoClOrdId_example" // string | Client-supplied Algo ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. (optional) (default to "")
	instType := "instType_example" // string | Instrument type  `SPOT`  `SWAP`  `FUTURES`  `MARGIN` (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `algoId`. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `algoId`. (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlgoTradingAPI.GetTradeOrdersAlgoPendingV5(context.Background()).OrdType(ordType).AlgoId(algoId).AlgoClOrdId(algoClOrdId).InstType(instType).InstId(instId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlgoTradingAPI.GetTradeOrdersAlgoPendingV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeOrdersAlgoPendingV5`: GetTradeOrdersAlgoPendingV5Resp
	fmt.Fprintf(os.Stdout, "Response from `AlgoTradingAPI.GetTradeOrdersAlgoPendingV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeOrdersAlgoPendingV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordType** | **string** | Order type  &#x60;conditional&#x60;: One-way stop order    &#x60;oco&#x60;: One-cancels-the-other order   &#x60;chase&#x60;: chase order, only applicable to FUTURES and SWAP  &#x60;trigger&#x60;: Trigger order   &#x60;move_order_stop&#x60;: Trailing order   &#x60;iceberg&#x60;: Iceberg order   &#x60;twap&#x60;: TWAP order  For every request, unlike other ordType which only can use one type, &#x60;conditional&#x60; and &#x60;oco&#x60; both can be used and separated with comma. | [default to &quot;&quot;]
 **algoId** | **string** | Algo ID | [default to &quot;&quot;]
 **algoClOrdId** | **string** | Client-supplied Algo ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;MARGIN&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;algoId&#x60;. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;algoId&#x60;. | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetTradeOrdersAlgoPendingV5Resp**](GetTradeOrdersAlgoPendingV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

