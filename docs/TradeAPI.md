# \TradeAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTradeAmendBatchOrdersV5**](TradeAPI.md#CreateTradeAmendBatchOrdersV5) | **Post** /api/v5/trade/amend-batch-orders | POST / Amend multiple orders
[**CreateTradeAmendOrderV5**](TradeAPI.md#CreateTradeAmendOrderV5) | **Post** /api/v5/trade/amend-order | POST / Amend order
[**CreateTradeBatchOrdersV5**](TradeAPI.md#CreateTradeBatchOrdersV5) | **Post** /api/v5/trade/batch-orders | POST / Place multiple orders
[**CreateTradeCancelAllAfterV5**](TradeAPI.md#CreateTradeCancelAllAfterV5) | **Post** /api/v5/trade/cancel-all-after | POST / Cancel All After
[**CreateTradeCancelBatchOrdersV5**](TradeAPI.md#CreateTradeCancelBatchOrdersV5) | **Post** /api/v5/trade/cancel-batch-orders | POST / Cancel multiple orders
[**CreateTradeCancelOrderV5**](TradeAPI.md#CreateTradeCancelOrderV5) | **Post** /api/v5/trade/cancel-order | POST / Cancel order
[**CreateTradeClosePositionV5**](TradeAPI.md#CreateTradeClosePositionV5) | **Post** /api/v5/trade/close-position | POST / Close positions
[**CreateTradeEasyConvertV5**](TradeAPI.md#CreateTradeEasyConvertV5) | **Post** /api/v5/trade/easy-convert | POST / Place easy convert
[**CreateTradeMassCancelV5**](TradeAPI.md#CreateTradeMassCancelV5) | **Post** /api/v5/trade/mass-cancel | POST / Mass cancel order
[**CreateTradeOneClickRepayV2V5**](TradeAPI.md#CreateTradeOneClickRepayV2V5) | **Post** /api/v5/trade/one-click-repay-v2 | POST / Trade one-click repay (New)
[**CreateTradeOneClickRepayV5**](TradeAPI.md#CreateTradeOneClickRepayV5) | **Post** /api/v5/trade/one-click-repay | POST / Trade one-click repay
[**CreateTradeOrderPrecheckV5**](TradeAPI.md#CreateTradeOrderPrecheckV5) | **Post** /api/v5/trade/order-precheck | POST / Order precheck
[**CreateTradeOrderV5**](TradeAPI.md#CreateTradeOrderV5) | **Post** /api/v5/trade/order | POST / Place order
[**GetTradeAccountRateLimitV5**](TradeAPI.md#GetTradeAccountRateLimitV5) | **Get** /api/v5/trade/account-rate-limit | GET / Account rate limit
[**GetTradeEasyConvertCurrencyListV5**](TradeAPI.md#GetTradeEasyConvertCurrencyListV5) | **Get** /api/v5/trade/easy-convert-currency-list | GET / Easy convert currency list
[**GetTradeEasyConvertHistoryV5**](TradeAPI.md#GetTradeEasyConvertHistoryV5) | **Get** /api/v5/trade/easy-convert-history | GET / Easy convert history
[**GetTradeFillsHistoryV5**](TradeAPI.md#GetTradeFillsHistoryV5) | **Get** /api/v5/trade/fills-history | GET / Transaction details (last 3 months)
[**GetTradeFillsV5**](TradeAPI.md#GetTradeFillsV5) | **Get** /api/v5/trade/fills | GET / Transaction details (last 3 days)
[**GetTradeOneClickRepayCurrencyListV2V5**](TradeAPI.md#GetTradeOneClickRepayCurrencyListV2V5) | **Get** /api/v5/trade/one-click-repay-currency-list-v2 | GET / One-click repay currency list (New)
[**GetTradeOneClickRepayCurrencyListV5**](TradeAPI.md#GetTradeOneClickRepayCurrencyListV5) | **Get** /api/v5/trade/one-click-repay-currency-list | GET / One-click repay currency list
[**GetTradeOneClickRepayHistoryV2V5**](TradeAPI.md#GetTradeOneClickRepayHistoryV2V5) | **Get** /api/v5/trade/one-click-repay-history-v2 | GET / One-click repay history (New)
[**GetTradeOneClickRepayHistoryV5**](TradeAPI.md#GetTradeOneClickRepayHistoryV5) | **Get** /api/v5/trade/one-click-repay-history | GET / One-click repay history
[**GetTradeOrderV5**](TradeAPI.md#GetTradeOrderV5) | **Get** /api/v5/trade/order | GET / Order details
[**GetTradeOrdersHistoryArchiveV5**](TradeAPI.md#GetTradeOrdersHistoryArchiveV5) | **Get** /api/v5/trade/orders-history-archive | GET / Order history (last 3 months)
[**GetTradeOrdersHistoryV5**](TradeAPI.md#GetTradeOrdersHistoryV5) | **Get** /api/v5/trade/orders-history | GET / Order history (last 7 days)
[**GetTradeOrdersPendingV5**](TradeAPI.md#GetTradeOrdersPendingV5) | **Get** /api/v5/trade/orders-pending | GET / Order List



## CreateTradeAmendBatchOrdersV5

> CreateTradeAmendBatchOrdersV5Resp CreateTradeAmendBatchOrdersV5(ctx).CreateTradeAmendBatchOrdersV5Req(createTradeAmendBatchOrdersV5Req).Execute()

POST / Amend multiple orders



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
	createTradeAmendBatchOrdersV5Req := *openapiclient.NewCreateTradeAmendBatchOrdersV5Req("InstId_example") // CreateTradeAmendBatchOrdersV5Req | The request body for CreateTradeAmendBatchOrdersV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.CreateTradeAmendBatchOrdersV5(context.Background()).CreateTradeAmendBatchOrdersV5Req(createTradeAmendBatchOrdersV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.CreateTradeAmendBatchOrdersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeAmendBatchOrdersV5`: CreateTradeAmendBatchOrdersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.CreateTradeAmendBatchOrdersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeAmendBatchOrdersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeAmendBatchOrdersV5Req** | [**CreateTradeAmendBatchOrdersV5Req**](CreateTradeAmendBatchOrdersV5Req.md) | The request body for CreateTradeAmendBatchOrdersV5 | 

### Return type

[**CreateTradeAmendBatchOrdersV5Resp**](CreateTradeAmendBatchOrdersV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeAmendOrderV5

> CreateTradeAmendOrderV5Resp CreateTradeAmendOrderV5(ctx).CreateTradeAmendOrderV5Req(createTradeAmendOrderV5Req).Execute()

POST / Amend order



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
	createTradeAmendOrderV5Req := *openapiclient.NewCreateTradeAmendOrderV5Req("InstId_example") // CreateTradeAmendOrderV5Req | The request body for CreateTradeAmendOrderV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.CreateTradeAmendOrderV5(context.Background()).CreateTradeAmendOrderV5Req(createTradeAmendOrderV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.CreateTradeAmendOrderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeAmendOrderV5`: CreateTradeAmendOrderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.CreateTradeAmendOrderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeAmendOrderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeAmendOrderV5Req** | [**CreateTradeAmendOrderV5Req**](CreateTradeAmendOrderV5Req.md) | The request body for CreateTradeAmendOrderV5 | 

### Return type

[**CreateTradeAmendOrderV5Resp**](CreateTradeAmendOrderV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeBatchOrdersV5

> CreateTradeBatchOrdersV5Resp CreateTradeBatchOrdersV5(ctx).CreateTradeBatchOrdersV5Req(createTradeBatchOrdersV5Req).Execute()

POST / Place multiple orders



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
	createTradeBatchOrdersV5Req := *openapiclient.NewCreateTradeBatchOrdersV5Req("InstId_example", "OrdType_example", "Side_example", "Sz_example", "TdMode_example") // CreateTradeBatchOrdersV5Req | The request body for CreateTradeBatchOrdersV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.CreateTradeBatchOrdersV5(context.Background()).CreateTradeBatchOrdersV5Req(createTradeBatchOrdersV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.CreateTradeBatchOrdersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeBatchOrdersV5`: CreateTradeBatchOrdersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.CreateTradeBatchOrdersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeBatchOrdersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeBatchOrdersV5Req** | [**CreateTradeBatchOrdersV5Req**](CreateTradeBatchOrdersV5Req.md) | The request body for CreateTradeBatchOrdersV5 | 

### Return type

[**CreateTradeBatchOrdersV5Resp**](CreateTradeBatchOrdersV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeCancelAllAfterV5

> CreateTradeCancelAllAfterV5Resp CreateTradeCancelAllAfterV5(ctx).CreateTradeCancelAllAfterV5Req(createTradeCancelAllAfterV5Req).Execute()

POST / Cancel All After



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
	createTradeCancelAllAfterV5Req := *openapiclient.NewCreateTradeCancelAllAfterV5Req("TimeOut_example") // CreateTradeCancelAllAfterV5Req | The request body for CreateTradeCancelAllAfterV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.CreateTradeCancelAllAfterV5(context.Background()).CreateTradeCancelAllAfterV5Req(createTradeCancelAllAfterV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.CreateTradeCancelAllAfterV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeCancelAllAfterV5`: CreateTradeCancelAllAfterV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.CreateTradeCancelAllAfterV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeCancelAllAfterV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeCancelAllAfterV5Req** | [**CreateTradeCancelAllAfterV5Req**](CreateTradeCancelAllAfterV5Req.md) | The request body for CreateTradeCancelAllAfterV5 | 

### Return type

[**CreateTradeCancelAllAfterV5Resp**](CreateTradeCancelAllAfterV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeCancelBatchOrdersV5

> CreateTradeCancelBatchOrdersV5Resp CreateTradeCancelBatchOrdersV5(ctx).CreateTradeCancelBatchOrdersV5Req(createTradeCancelBatchOrdersV5Req).Execute()

POST / Cancel multiple orders



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
	createTradeCancelBatchOrdersV5Req := *openapiclient.NewCreateTradeCancelBatchOrdersV5Req("InstId_example") // CreateTradeCancelBatchOrdersV5Req | The request body for CreateTradeCancelBatchOrdersV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.CreateTradeCancelBatchOrdersV5(context.Background()).CreateTradeCancelBatchOrdersV5Req(createTradeCancelBatchOrdersV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.CreateTradeCancelBatchOrdersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeCancelBatchOrdersV5`: CreateTradeCancelBatchOrdersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.CreateTradeCancelBatchOrdersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeCancelBatchOrdersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeCancelBatchOrdersV5Req** | [**CreateTradeCancelBatchOrdersV5Req**](CreateTradeCancelBatchOrdersV5Req.md) | The request body for CreateTradeCancelBatchOrdersV5 | 

### Return type

[**CreateTradeCancelBatchOrdersV5Resp**](CreateTradeCancelBatchOrdersV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeCancelOrderV5

> CreateTradeCancelOrderV5Resp CreateTradeCancelOrderV5(ctx).CreateTradeCancelOrderV5Req(createTradeCancelOrderV5Req).Execute()

POST / Cancel order



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
	createTradeCancelOrderV5Req := *openapiclient.NewCreateTradeCancelOrderV5Req("InstId_example") // CreateTradeCancelOrderV5Req | The request body for CreateTradeCancelOrderV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.CreateTradeCancelOrderV5(context.Background()).CreateTradeCancelOrderV5Req(createTradeCancelOrderV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.CreateTradeCancelOrderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeCancelOrderV5`: CreateTradeCancelOrderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.CreateTradeCancelOrderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeCancelOrderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeCancelOrderV5Req** | [**CreateTradeCancelOrderV5Req**](CreateTradeCancelOrderV5Req.md) | The request body for CreateTradeCancelOrderV5 | 

### Return type

[**CreateTradeCancelOrderV5Resp**](CreateTradeCancelOrderV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeClosePositionV5

> CreateTradeClosePositionV5Resp CreateTradeClosePositionV5(ctx).CreateTradeClosePositionV5Req(createTradeClosePositionV5Req).Execute()

POST / Close positions



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
	createTradeClosePositionV5Req := *openapiclient.NewCreateTradeClosePositionV5Req("InstId_example", "MgnMode_example") // CreateTradeClosePositionV5Req | The request body for CreateTradeClosePositionV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.CreateTradeClosePositionV5(context.Background()).CreateTradeClosePositionV5Req(createTradeClosePositionV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.CreateTradeClosePositionV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeClosePositionV5`: CreateTradeClosePositionV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.CreateTradeClosePositionV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeClosePositionV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeClosePositionV5Req** | [**CreateTradeClosePositionV5Req**](CreateTradeClosePositionV5Req.md) | The request body for CreateTradeClosePositionV5 | 

### Return type

[**CreateTradeClosePositionV5Resp**](CreateTradeClosePositionV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeEasyConvertV5

> CreateTradeEasyConvertV5Resp CreateTradeEasyConvertV5(ctx).CreateTradeEasyConvertV5Req(createTradeEasyConvertV5Req).Execute()

POST / Place easy convert



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
	createTradeEasyConvertV5Req := *openapiclient.NewCreateTradeEasyConvertV5Req([]string{"FromCcy_example"}, "ToCcy_example") // CreateTradeEasyConvertV5Req | The request body for CreateTradeEasyConvertV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.CreateTradeEasyConvertV5(context.Background()).CreateTradeEasyConvertV5Req(createTradeEasyConvertV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.CreateTradeEasyConvertV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeEasyConvertV5`: CreateTradeEasyConvertV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.CreateTradeEasyConvertV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeEasyConvertV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeEasyConvertV5Req** | [**CreateTradeEasyConvertV5Req**](CreateTradeEasyConvertV5Req.md) | The request body for CreateTradeEasyConvertV5 | 

### Return type

[**CreateTradeEasyConvertV5Resp**](CreateTradeEasyConvertV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeMassCancelV5

> CreateTradeMassCancelV5Resp CreateTradeMassCancelV5(ctx).CreateTradeMassCancelV5Req(createTradeMassCancelV5Req).Execute()

POST / Mass cancel order



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
	createTradeMassCancelV5Req := *openapiclient.NewCreateTradeMassCancelV5Req("InstFamily_example", "InstType_example") // CreateTradeMassCancelV5Req | The request body for CreateTradeMassCancelV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.CreateTradeMassCancelV5(context.Background()).CreateTradeMassCancelV5Req(createTradeMassCancelV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.CreateTradeMassCancelV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeMassCancelV5`: CreateTradeMassCancelV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.CreateTradeMassCancelV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeMassCancelV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeMassCancelV5Req** | [**CreateTradeMassCancelV5Req**](CreateTradeMassCancelV5Req.md) | The request body for CreateTradeMassCancelV5 | 

### Return type

[**CreateTradeMassCancelV5Resp**](CreateTradeMassCancelV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeOneClickRepayV2V5

> CreateTradeOneClickRepayV2V5Resp CreateTradeOneClickRepayV2V5(ctx).CreateTradeOneClickRepayV2V5Req(createTradeOneClickRepayV2V5Req).Execute()

POST / Trade one-click repay (New)



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
	createTradeOneClickRepayV2V5Req := *openapiclient.NewCreateTradeOneClickRepayV2V5Req("DebtCcy_example", []string{"RepayCcyList_example"}) // CreateTradeOneClickRepayV2V5Req | The request body for CreateTradeOneClickRepayV2V5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.CreateTradeOneClickRepayV2V5(context.Background()).CreateTradeOneClickRepayV2V5Req(createTradeOneClickRepayV2V5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.CreateTradeOneClickRepayV2V5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeOneClickRepayV2V5`: CreateTradeOneClickRepayV2V5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.CreateTradeOneClickRepayV2V5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeOneClickRepayV2V5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeOneClickRepayV2V5Req** | [**CreateTradeOneClickRepayV2V5Req**](CreateTradeOneClickRepayV2V5Req.md) | The request body for CreateTradeOneClickRepayV2V5 | 

### Return type

[**CreateTradeOneClickRepayV2V5Resp**](CreateTradeOneClickRepayV2V5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeOneClickRepayV5

> CreateTradeOneClickRepayV5Resp CreateTradeOneClickRepayV5(ctx).CreateTradeOneClickRepayV5Req(createTradeOneClickRepayV5Req).Execute()

POST / Trade one-click repay



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
	createTradeOneClickRepayV5Req := *openapiclient.NewCreateTradeOneClickRepayV5Req([]string{"DebtCcy_example"}, "RepayCcy_example") // CreateTradeOneClickRepayV5Req | The request body for CreateTradeOneClickRepayV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.CreateTradeOneClickRepayV5(context.Background()).CreateTradeOneClickRepayV5Req(createTradeOneClickRepayV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.CreateTradeOneClickRepayV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeOneClickRepayV5`: CreateTradeOneClickRepayV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.CreateTradeOneClickRepayV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeOneClickRepayV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeOneClickRepayV5Req** | [**CreateTradeOneClickRepayV5Req**](CreateTradeOneClickRepayV5Req.md) | The request body for CreateTradeOneClickRepayV5 | 

### Return type

[**CreateTradeOneClickRepayV5Resp**](CreateTradeOneClickRepayV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeOrderPrecheckV5

> CreateTradeOrderPrecheckV5Resp CreateTradeOrderPrecheckV5(ctx).CreateTradeOrderPrecheckV5Req(createTradeOrderPrecheckV5Req).Execute()

POST / Order precheck



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
	createTradeOrderPrecheckV5Req := *openapiclient.NewCreateTradeOrderPrecheckV5Req("InstId_example", "OrdType_example", "Side_example", "Sz_example", "TdMode_example") // CreateTradeOrderPrecheckV5Req | The request body for CreateTradeOrderPrecheckV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.CreateTradeOrderPrecheckV5(context.Background()).CreateTradeOrderPrecheckV5Req(createTradeOrderPrecheckV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.CreateTradeOrderPrecheckV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeOrderPrecheckV5`: CreateTradeOrderPrecheckV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.CreateTradeOrderPrecheckV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeOrderPrecheckV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeOrderPrecheckV5Req** | [**CreateTradeOrderPrecheckV5Req**](CreateTradeOrderPrecheckV5Req.md) | The request body for CreateTradeOrderPrecheckV5 | 

### Return type

[**CreateTradeOrderPrecheckV5Resp**](CreateTradeOrderPrecheckV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTradeOrderV5

> CreateTradeOrderV5Resp CreateTradeOrderV5(ctx).CreateTradeOrderV5Req(createTradeOrderV5Req).Execute()

POST / Place order



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
	createTradeOrderV5Req := *openapiclient.NewCreateTradeOrderV5Req("InstId_example", "OrdType_example", "Side_example", "Sz_example", "TdMode_example") // CreateTradeOrderV5Req | The request body for CreateTradeOrderV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.CreateTradeOrderV5(context.Background()).CreateTradeOrderV5Req(createTradeOrderV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.CreateTradeOrderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTradeOrderV5`: CreateTradeOrderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.CreateTradeOrderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeOrderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeOrderV5Req** | [**CreateTradeOrderV5Req**](CreateTradeOrderV5Req.md) | The request body for CreateTradeOrderV5 | 

### Return type

[**CreateTradeOrderV5Resp**](CreateTradeOrderV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeAccountRateLimitV5

> GetTradeAccountRateLimitV5Resp GetTradeAccountRateLimitV5(ctx).Execute()

GET / Account rate limit



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
	resp, r, err := apiClient.TradeAPI.GetTradeAccountRateLimitV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTradeAccountRateLimitV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeAccountRateLimitV5`: GetTradeAccountRateLimitV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTradeAccountRateLimitV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeAccountRateLimitV5Request struct via the builder pattern


### Return type

[**GetTradeAccountRateLimitV5Resp**](GetTradeAccountRateLimitV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeEasyConvertCurrencyListV5

> GetTradeEasyConvertCurrencyListV5Resp GetTradeEasyConvertCurrencyListV5(ctx).Source(source).Execute()

GET / Easy convert currency list



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
	source := "source_example" // string | Funding source  `1`: Trading account  `2`: Funding account  The default is `1`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTradeEasyConvertCurrencyListV5(context.Background()).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTradeEasyConvertCurrencyListV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeEasyConvertCurrencyListV5`: GetTradeEasyConvertCurrencyListV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTradeEasyConvertCurrencyListV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeEasyConvertCurrencyListV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** | Funding source  &#x60;1&#x60;: Trading account  &#x60;2&#x60;: Funding account  The default is &#x60;1&#x60;. | [default to &quot;&quot;]

### Return type

[**GetTradeEasyConvertCurrencyListV5Resp**](GetTradeEasyConvertCurrencyListV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeEasyConvertHistoryV5

> GetTradeEasyConvertHistoryV5Resp GetTradeEasyConvertHistoryV5(ctx).After(after).Before(before).Limit(limit).Execute()

GET / Easy convert history



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
	after := "after_example" // string | Pagination of data to return records earlier than the requested time (exclude), Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested time (exclude), Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTradeEasyConvertHistoryV5(context.Background()).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTradeEasyConvertHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeEasyConvertHistoryV5`: GetTradeEasyConvertHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTradeEasyConvertHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeEasyConvertHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | Pagination of data to return records earlier than the requested time (exclude), Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested time (exclude), Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetTradeEasyConvertHistoryV5Resp**](GetTradeEasyConvertHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeFillsHistoryV5

> GetTradeFillsHistoryV5Resp GetTradeFillsHistoryV5(ctx).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).OrdId(ordId).SubType(subType).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()

GET / Transaction details (last 3 months)



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
	instType := "instType_example" // string | Instrument type  `SPOT`  `MARGIN`  `SWAP`  `FUTURES`  `OPTION` (default to "")
	uly := "uly_example" // string | Underlying  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT` (optional) (default to "")
	ordId := "ordId_example" // string | Order ID (optional) (default to "")
	subType := "subType_example" // string | Transaction type   `1`: Buy  `2`: Sell  `3`: Open long  `4`: Open short  `5`: Close long  `6`: Close short   `100`: Partial liquidation close long  `101`: Partial liquidation close short  `102`: Partial liquidation buy  `103`: Partial liquidation sell  `104`: Liquidation long  `105`: Liquidation short  `106`: Liquidation buy   `107`: Liquidation sell   `110`: Liquidation transfer in  `111`: Liquidation transfer out   `118`: System token conversion transfer in  `119`: System token conversion transfer out   `125`: ADL close long  `126`: ADL close short  `127`: ADL buy  `128`: ADL sell   `212`: Auto borrow of quick margin  `213`: Auto repay of quick margin   `204`: block trade buy  `205`: block trade sell  `206`: block trade open long  `207`: block trade open short  `208`: block trade close long  `209`: block trade close short  `236`: Easy convert in  `237`: Easy convert out  `270`: Spread trading buy  `271`: Spread trading sell  `272`: Spread trading open long  `273`: Spread trading open short  `274`: Spread trading close long  `275`: Spread trading close short (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `billId` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `billId` (optional) (default to "")
	begin := "begin_example" // string | Filter with a begin timestamp `ts`. Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	end := "end_example" // string | Filter with an end timestamp `ts`. Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTradeFillsHistoryV5(context.Background()).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).OrdId(ordId).SubType(subType).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTradeFillsHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeFillsHistoryV5`: GetTradeFillsHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTradeFillsHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeFillsHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Underlying  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **ordId** | **string** | Order ID | [default to &quot;&quot;]
 **subType** | **string** | Transaction type   &#x60;1&#x60;: Buy  &#x60;2&#x60;: Sell  &#x60;3&#x60;: Open long  &#x60;4&#x60;: Open short  &#x60;5&#x60;: Close long  &#x60;6&#x60;: Close short   &#x60;100&#x60;: Partial liquidation close long  &#x60;101&#x60;: Partial liquidation close short  &#x60;102&#x60;: Partial liquidation buy  &#x60;103&#x60;: Partial liquidation sell  &#x60;104&#x60;: Liquidation long  &#x60;105&#x60;: Liquidation short  &#x60;106&#x60;: Liquidation buy   &#x60;107&#x60;: Liquidation sell   &#x60;110&#x60;: Liquidation transfer in  &#x60;111&#x60;: Liquidation transfer out   &#x60;118&#x60;: System token conversion transfer in  &#x60;119&#x60;: System token conversion transfer out   &#x60;125&#x60;: ADL close long  &#x60;126&#x60;: ADL close short  &#x60;127&#x60;: ADL buy  &#x60;128&#x60;: ADL sell   &#x60;212&#x60;: Auto borrow of quick margin  &#x60;213&#x60;: Auto repay of quick margin   &#x60;204&#x60;: block trade buy  &#x60;205&#x60;: block trade sell  &#x60;206&#x60;: block trade open long  &#x60;207&#x60;: block trade open short  &#x60;208&#x60;: block trade close long  &#x60;209&#x60;: block trade close short  &#x60;236&#x60;: Easy convert in  &#x60;237&#x60;: Easy convert out  &#x60;270&#x60;: Spread trading buy  &#x60;271&#x60;: Spread trading sell  &#x60;272&#x60;: Spread trading open long  &#x60;273&#x60;: Spread trading open short  &#x60;274&#x60;: Spread trading close long  &#x60;275&#x60;: Spread trading close short | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;billId&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;billId&#x60; | [default to &quot;&quot;]
 **begin** | **string** | Filter with a begin timestamp &#x60;ts&#x60;. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **end** | **string** | Filter with an end timestamp &#x60;ts&#x60;. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetTradeFillsHistoryV5Resp**](GetTradeFillsHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeFillsV5

> GetTradeFillsV5Resp GetTradeFillsV5(ctx).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).OrdId(ordId).SubType(subType).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()

GET / Transaction details (last 3 days)



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
	instType := "instType_example" // string | Instrument type  `SPOT`  `MARGIN`  `SWAP`  `FUTURES`  `OPTION` (optional) (default to "")
	uly := "uly_example" // string | Underlying  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT` (optional) (default to "")
	ordId := "ordId_example" // string | Order ID (optional) (default to "")
	subType := "subType_example" // string | Transaction type   `1`: Buy  `2`: Sell  `3`: Open long  `4`: Open short  `5`: Close long  `6`: Close short   `100`: Partial liquidation close long  `101`: Partial liquidation close short  `102`: Partial liquidation buy  `103`: Partial liquidation sell  `104`: Liquidation long  `105`: Liquidation short  `106`: Liquidation buy   `107`: Liquidation sell   `110`: Liquidation transfer in  `111`: Liquidation transfer out   `118`: System token conversion transfer in  `119`: System token conversion transfer out   `125`: ADL close long  `126`: ADL close short  `127`: ADL buy  `128`: ADL sell   `212`: Auto borrow of quick margin  `213`: Auto repay of quick margin   `204`: block trade buy  `205`: block trade sell  `206`: block trade open long  `207`: block trade open short  `208`: block trade close long  `209`: block trade close short  `236`: Easy convert in  `237`: Easy convert out  `270`: Spread trading buy  `271`: Spread trading sell  `272`: Spread trading open long  `273`: Spread trading open short  `274`: Spread trading close long  `275`: Spread trading close short (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `billId` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `billId` (optional) (default to "")
	begin := "begin_example" // string | Filter with a begin timestamp  `ts`. Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	end := "end_example" // string | Filter with an end timestamp  `ts`. Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTradeFillsV5(context.Background()).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).OrdId(ordId).SubType(subType).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTradeFillsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeFillsV5`: GetTradeFillsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTradeFillsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeFillsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Underlying  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **ordId** | **string** | Order ID | [default to &quot;&quot;]
 **subType** | **string** | Transaction type   &#x60;1&#x60;: Buy  &#x60;2&#x60;: Sell  &#x60;3&#x60;: Open long  &#x60;4&#x60;: Open short  &#x60;5&#x60;: Close long  &#x60;6&#x60;: Close short   &#x60;100&#x60;: Partial liquidation close long  &#x60;101&#x60;: Partial liquidation close short  &#x60;102&#x60;: Partial liquidation buy  &#x60;103&#x60;: Partial liquidation sell  &#x60;104&#x60;: Liquidation long  &#x60;105&#x60;: Liquidation short  &#x60;106&#x60;: Liquidation buy   &#x60;107&#x60;: Liquidation sell   &#x60;110&#x60;: Liquidation transfer in  &#x60;111&#x60;: Liquidation transfer out   &#x60;118&#x60;: System token conversion transfer in  &#x60;119&#x60;: System token conversion transfer out   &#x60;125&#x60;: ADL close long  &#x60;126&#x60;: ADL close short  &#x60;127&#x60;: ADL buy  &#x60;128&#x60;: ADL sell   &#x60;212&#x60;: Auto borrow of quick margin  &#x60;213&#x60;: Auto repay of quick margin   &#x60;204&#x60;: block trade buy  &#x60;205&#x60;: block trade sell  &#x60;206&#x60;: block trade open long  &#x60;207&#x60;: block trade open short  &#x60;208&#x60;: block trade close long  &#x60;209&#x60;: block trade close short  &#x60;236&#x60;: Easy convert in  &#x60;237&#x60;: Easy convert out  &#x60;270&#x60;: Spread trading buy  &#x60;271&#x60;: Spread trading sell  &#x60;272&#x60;: Spread trading open long  &#x60;273&#x60;: Spread trading open short  &#x60;274&#x60;: Spread trading close long  &#x60;275&#x60;: Spread trading close short | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;billId&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;billId&#x60; | [default to &quot;&quot;]
 **begin** | **string** | Filter with a begin timestamp  &#x60;ts&#x60;. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **end** | **string** | Filter with an end timestamp  &#x60;ts&#x60;. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetTradeFillsV5Resp**](GetTradeFillsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeOneClickRepayCurrencyListV2V5

> GetTradeOneClickRepayCurrencyListV2V5Resp GetTradeOneClickRepayCurrencyListV2V5(ctx).Execute()

GET / One-click repay currency list (New)



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
	resp, r, err := apiClient.TradeAPI.GetTradeOneClickRepayCurrencyListV2V5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTradeOneClickRepayCurrencyListV2V5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeOneClickRepayCurrencyListV2V5`: GetTradeOneClickRepayCurrencyListV2V5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTradeOneClickRepayCurrencyListV2V5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeOneClickRepayCurrencyListV2V5Request struct via the builder pattern


### Return type

[**GetTradeOneClickRepayCurrencyListV2V5Resp**](GetTradeOneClickRepayCurrencyListV2V5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeOneClickRepayCurrencyListV5

> GetTradeOneClickRepayCurrencyListV5Resp GetTradeOneClickRepayCurrencyListV5(ctx).DebtType(debtType).Execute()

GET / One-click repay currency list



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
	debtType := "debtType_example" // string | Debt type   `cross`: cross   `isolated`: isolated (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTradeOneClickRepayCurrencyListV5(context.Background()).DebtType(debtType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTradeOneClickRepayCurrencyListV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeOneClickRepayCurrencyListV5`: GetTradeOneClickRepayCurrencyListV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTradeOneClickRepayCurrencyListV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeOneClickRepayCurrencyListV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **debtType** | **string** | Debt type   &#x60;cross&#x60;: cross   &#x60;isolated&#x60;: isolated | [default to &quot;&quot;]

### Return type

[**GetTradeOneClickRepayCurrencyListV5Resp**](GetTradeOneClickRepayCurrencyListV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeOneClickRepayHistoryV2V5

> GetTradeOneClickRepayHistoryV2V5Resp GetTradeOneClickRepayHistoryV2V5(ctx).After(after).Before(before).Limit(limit).Execute()

GET / One-click repay history (New)



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
	after := "after_example" // string | Pagination of data to return records earlier than (included) the requested time `ts` , Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than (included) the requested time `ts`, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTradeOneClickRepayHistoryV2V5(context.Background()).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTradeOneClickRepayHistoryV2V5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeOneClickRepayHistoryV2V5`: GetTradeOneClickRepayHistoryV2V5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTradeOneClickRepayHistoryV2V5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeOneClickRepayHistoryV2V5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | Pagination of data to return records earlier than (included) the requested time &#x60;ts&#x60; , Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than (included) the requested time &#x60;ts&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetTradeOneClickRepayHistoryV2V5Resp**](GetTradeOneClickRepayHistoryV2V5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeOneClickRepayHistoryV5

> GetTradeOneClickRepayHistoryV5Resp GetTradeOneClickRepayHistoryV5(ctx).After(after).Before(before).Limit(limit).Execute()

GET / One-click repay history



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
	after := "after_example" // string | Pagination of data to return records earlier than the requested time, Unix timestamp format in milliseconds, e.g. 1597026383085 (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested time, Unix timestamp format in milliseconds, e.g. 1597026383085 (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTradeOneClickRepayHistoryV5(context.Background()).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTradeOneClickRepayHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeOneClickRepayHistoryV5`: GetTradeOneClickRepayHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTradeOneClickRepayHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeOneClickRepayHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | Pagination of data to return records earlier than the requested time, Unix timestamp format in milliseconds, e.g. 1597026383085 | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested time, Unix timestamp format in milliseconds, e.g. 1597026383085 | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetTradeOneClickRepayHistoryV5Resp**](GetTradeOneClickRepayHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeOrderV5

> GetTradeOrderV5Resp GetTradeOrderV5(ctx).InstId(instId).OrdId(ordId).ClOrdId(clOrdId).Execute()

GET / Order details



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
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT`  Only applicable to live instruments (default to "")
	ordId := "ordId_example" // string | Order ID   Either `ordId` or `clOrdId` is required, if both are passed, `ordId` will be used (optional) (default to "")
	clOrdId := "clOrdId_example" // string | Client Order ID as assigned by the client  If the `clOrdId` is associated with multiple orders, only the latest one will be returned. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTradeOrderV5(context.Background()).InstId(instId).OrdId(ordId).ClOrdId(clOrdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTradeOrderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeOrderV5`: GetTradeOrderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTradeOrderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeOrderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60;  Only applicable to live instruments | [default to &quot;&quot;]
 **ordId** | **string** | Order ID   Either &#x60;ordId&#x60; or &#x60;clOrdId&#x60; is required, if both are passed, &#x60;ordId&#x60; will be used | [default to &quot;&quot;]
 **clOrdId** | **string** | Client Order ID as assigned by the client  If the &#x60;clOrdId&#x60; is associated with multiple orders, only the latest one will be returned. | [default to &quot;&quot;]

### Return type

[**GetTradeOrderV5Resp**](GetTradeOrderV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeOrdersHistoryArchiveV5

> GetTradeOrdersHistoryArchiveV5Resp GetTradeOrdersHistoryArchiveV5(ctx).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).OrdType(ordType).State(state).Category(category).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()

GET / Order history (last 3 months)



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
	instType := "instType_example" // string | Instrument type  `SPOT`  `MARGIN`  `SWAP`  `FUTURES`  `OPTION` (default to "")
	uly := "uly_example" // string | Underlying  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USD-200927` (optional) (default to "")
	ordType := "ordType_example" // string | Order type   `market`: Market order   `limit`: Limit order   `post_only`: Post-only order   `fok`: Fill-or-kill order   `ioc`: Immediate-or-cancel order    `optimal_limit_ioc`: Market order with immediate-or-cancel order  `mmp`: Market Maker Protection (only applicable to Option in Portfolio Margin mode)  `mmp_and_post_only`: Market Maker Protection and Post-only order(only applicable to Option in Portfolio Margin mode)   `op_fok`: Simple options (fok) (optional) (default to "")
	state := "state_example" // string | State  `canceled`   `filled`  `mmp_canceled`: Order canceled automatically due to Market Maker Protection (optional) (default to "")
	category := "category_example" // string | Category   `twap`   `adl`  `full_liquidation`  `partial_liquidation`  `delivery`    `ddh`: Delta dynamic hedge (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ordId` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ordId` (optional) (default to "")
	begin := "begin_example" // string | Filter with a begin timestamp `cTime`. Unix timestamp format in milliseconds, e.g. 1597026383085 (optional) (default to "")
	end := "end_example" // string | Filter with an end timestamp `cTime`. Unix timestamp format in milliseconds, e.g. 1597026383085 (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTradeOrdersHistoryArchiveV5(context.Background()).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).OrdType(ordType).State(state).Category(category).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTradeOrdersHistoryArchiveV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeOrdersHistoryArchiveV5`: GetTradeOrdersHistoryArchiveV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTradeOrdersHistoryArchiveV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeOrdersHistoryArchiveV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Underlying  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USD-200927&#x60; | [default to &quot;&quot;]
 **ordType** | **string** | Order type   &#x60;market&#x60;: Market order   &#x60;limit&#x60;: Limit order   &#x60;post_only&#x60;: Post-only order   &#x60;fok&#x60;: Fill-or-kill order   &#x60;ioc&#x60;: Immediate-or-cancel order    &#x60;optimal_limit_ioc&#x60;: Market order with immediate-or-cancel order  &#x60;mmp&#x60;: Market Maker Protection (only applicable to Option in Portfolio Margin mode)  &#x60;mmp_and_post_only&#x60;: Market Maker Protection and Post-only order(only applicable to Option in Portfolio Margin mode)   &#x60;op_fok&#x60;: Simple options (fok) | [default to &quot;&quot;]
 **state** | **string** | State  &#x60;canceled&#x60;   &#x60;filled&#x60;  &#x60;mmp_canceled&#x60;: Order canceled automatically due to Market Maker Protection | [default to &quot;&quot;]
 **category** | **string** | Category   &#x60;twap&#x60;   &#x60;adl&#x60;  &#x60;full_liquidation&#x60;  &#x60;partial_liquidation&#x60;  &#x60;delivery&#x60;    &#x60;ddh&#x60;: Delta dynamic hedge | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ordId&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ordId&#x60; | [default to &quot;&quot;]
 **begin** | **string** | Filter with a begin timestamp &#x60;cTime&#x60;. Unix timestamp format in milliseconds, e.g. 1597026383085 | [default to &quot;&quot;]
 **end** | **string** | Filter with an end timestamp &#x60;cTime&#x60;. Unix timestamp format in milliseconds, e.g. 1597026383085 | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetTradeOrdersHistoryArchiveV5Resp**](GetTradeOrdersHistoryArchiveV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeOrdersHistoryV5

> GetTradeOrdersHistoryV5Resp GetTradeOrdersHistoryV5(ctx).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).OrdType(ordType).State(state).Category(category).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()

GET / Order history (last 7 days)



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
	instType := "instType_example" // string | Instrument type  `SPOT`  `MARGIN`  `SWAP`  `FUTURES`  `OPTION` (default to "")
	uly := "uly_example" // string | Underlying  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT` (optional) (default to "")
	ordType := "ordType_example" // string | Order type  `market`: market order   `limit`: limit order   `post_only`: Post-only order   `fok`: Fill-or-kill order   `ioc`: Immediate-or-cancel order    `optimal_limit_ioc`: Market order with immediate-or-cancel order  `mmp`: Market Maker Protection (only applicable to Option in Portfolio Margin mode)  `mmp_and_post_only`: Market Maker Protection and Post-only order(only applicable to Option in Portfolio Margin mode)   `op_fok`: Simple options (fok) (optional) (default to "")
	state := "state_example" // string | State  `canceled`  `filled`  `mmp_canceled`: Order canceled automatically due to Market Maker Protection (optional) (default to "")
	category := "category_example" // string | Category   `twap`   `adl`  `full_liquidation`  `partial_liquidation`   `delivery`    `ddh`: Delta dynamic hedge (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ordId` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ordId` (optional) (default to "")
	begin := "begin_example" // string | Filter with a begin timestamp `cTime`. Unix timestamp format in milliseconds, e.g. 1597026383085 (optional) (default to "")
	end := "end_example" // string | Filter with an end timestamp `cTime`. Unix timestamp format in milliseconds, e.g. 1597026383085 (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTradeOrdersHistoryV5(context.Background()).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).OrdType(ordType).State(state).Category(category).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTradeOrdersHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeOrdersHistoryV5`: GetTradeOrdersHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTradeOrdersHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeOrdersHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Underlying  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **ordType** | **string** | Order type  &#x60;market&#x60;: market order   &#x60;limit&#x60;: limit order   &#x60;post_only&#x60;: Post-only order   &#x60;fok&#x60;: Fill-or-kill order   &#x60;ioc&#x60;: Immediate-or-cancel order    &#x60;optimal_limit_ioc&#x60;: Market order with immediate-or-cancel order  &#x60;mmp&#x60;: Market Maker Protection (only applicable to Option in Portfolio Margin mode)  &#x60;mmp_and_post_only&#x60;: Market Maker Protection and Post-only order(only applicable to Option in Portfolio Margin mode)   &#x60;op_fok&#x60;: Simple options (fok) | [default to &quot;&quot;]
 **state** | **string** | State  &#x60;canceled&#x60;  &#x60;filled&#x60;  &#x60;mmp_canceled&#x60;: Order canceled automatically due to Market Maker Protection | [default to &quot;&quot;]
 **category** | **string** | Category   &#x60;twap&#x60;   &#x60;adl&#x60;  &#x60;full_liquidation&#x60;  &#x60;partial_liquidation&#x60;   &#x60;delivery&#x60;    &#x60;ddh&#x60;: Delta dynamic hedge | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ordId&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ordId&#x60; | [default to &quot;&quot;]
 **begin** | **string** | Filter with a begin timestamp &#x60;cTime&#x60;. Unix timestamp format in milliseconds, e.g. 1597026383085 | [default to &quot;&quot;]
 **end** | **string** | Filter with an end timestamp &#x60;cTime&#x60;. Unix timestamp format in milliseconds, e.g. 1597026383085 | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetTradeOrdersHistoryV5Resp**](GetTradeOrdersHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTradeOrdersPendingV5

> GetTradeOrdersPendingV5Resp GetTradeOrdersPendingV5(ctx).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).OrdType(ordType).State(state).After(after).Before(before).Limit(limit).Execute()

GET / Order List



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
	instType := "instType_example" // string | Instrument type  `SPOT`  `MARGIN`  `SWAP`  `FUTURES`  `OPTION` (optional) (default to "")
	uly := "uly_example" // string | Underlying  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USD-200927` (optional) (default to "")
	ordType := "ordType_example" // string | Order type   `market`: Market order   `limit`: Limit order   `post_only`: Post-only order   `fok`: Fill-or-kill order   `ioc`: Immediate-or-cancel order    `optimal_limit_ioc`: Market order with immediate-or-cancel order   `mmp`: Market Maker Protection (only applicable to Option in Portfolio Margin mode)  `mmp_and_post_only`: Market Maker Protection and Post-only order(only applicable to Option in Portfolio Margin mode)   `op_fok`: Simple options (fok) (optional) (default to "")
	state := "state_example" // string | State  `live`   `partially_filled` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ordId` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ordId` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradeAPI.GetTradeOrdersPendingV5(context.Background()).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).OrdType(ordType).State(state).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradeAPI.GetTradeOrdersPendingV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTradeOrdersPendingV5`: GetTradeOrdersPendingV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradeAPI.GetTradeOrdersPendingV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeOrdersPendingV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Underlying  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USD-200927&#x60; | [default to &quot;&quot;]
 **ordType** | **string** | Order type   &#x60;market&#x60;: Market order   &#x60;limit&#x60;: Limit order   &#x60;post_only&#x60;: Post-only order   &#x60;fok&#x60;: Fill-or-kill order   &#x60;ioc&#x60;: Immediate-or-cancel order    &#x60;optimal_limit_ioc&#x60;: Market order with immediate-or-cancel order   &#x60;mmp&#x60;: Market Maker Protection (only applicable to Option in Portfolio Margin mode)  &#x60;mmp_and_post_only&#x60;: Market Maker Protection and Post-only order(only applicable to Option in Portfolio Margin mode)   &#x60;op_fok&#x60;: Simple options (fok) | [default to &quot;&quot;]
 **state** | **string** | State  &#x60;live&#x60;   &#x60;partially_filled&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ordId&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ordId&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetTradeOrdersPendingV5Resp**](GetTradeOrdersPendingV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

