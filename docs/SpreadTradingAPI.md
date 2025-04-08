# \SpreadTradingAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSprdAmendOrderV5**](SpreadTradingAPI.md#CreateSprdAmendOrderV5) | **Post** /api/v5/sprd/amend-order | Amend an incomplete order.  
[**CreateSprdCancelAllAfterV5**](SpreadTradingAPI.md#CreateSprdCancelAllAfterV5) | **Post** /api/v5/sprd/cancel-all-after | Cancel all pending orders after the countdown timeout. Only applicable to spread trading.  
[**CreateSprdCancelOrderV5**](SpreadTradingAPI.md#CreateSprdCancelOrderV5) | **Post** /api/v5/sprd/cancel-order | Cancel an incomplete order.  
[**CreateSprdMassCancelV5**](SpreadTradingAPI.md#CreateSprdMassCancelV5) | **Post** /api/v5/sprd/mass-cancel | Cancel all pending orders.  
[**CreateSprdOrderV5**](SpreadTradingAPI.md#CreateSprdOrderV5) | **Post** /api/v5/sprd/order | Place a new order  
[**GetMarketSprdCandlesV5**](SpreadTradingAPI.md#GetMarketSprdCandlesV5) | **Get** /api/v5/market/sprd-candles | Retrieve the candlestick charts. This endpoint can retrieve the latest 1,440 data entries. Charts are returned in groups based on the requested bar.  
[**GetMarketSprdHistoryCandlesV5**](SpreadTradingAPI.md#GetMarketSprdHistoryCandlesV5) | **Get** /api/v5/market/sprd-history-candles | Retrieve history candlestick charts from recent years.  
[**GetMarketSprdTickerV5**](SpreadTradingAPI.md#GetMarketSprdTickerV5) | **Get** /api/v5/market/sprd-ticker | Retrieve the latest price snapshot, best bid/ask price and quantity.  
[**GetSprdBooksV5**](SpreadTradingAPI.md#GetSprdBooksV5) | **Get** /api/v5/sprd/books | Retrieve the order book of the spread.  
[**GetSprdOrderV5**](SpreadTradingAPI.md#GetSprdOrderV5) | **Get** /api/v5/sprd/order | Retrieve order details.  
[**GetSprdOrdersHistoryArchiveV5**](SpreadTradingAPI.md#GetSprdOrdersHistoryArchiveV5) | **Get** /api/v5/sprd/orders-history-archive | Retrieve the completed order data for the last 3 months, including those placed 3 months ago but completed in the last 3 months. Results are returned in counter chronological order.  
[**GetSprdOrdersHistoryV5**](SpreadTradingAPI.md#GetSprdOrdersHistoryV5) | **Get** /api/v5/sprd/orders-history | Retrieve the completed order data for the last 21 days, and the incomplete orders (filledSz &#x3D;0 &amp; state &#x3D; canceled) that have been canceled are only reserved for 2 hours. Results are returned in counter chronological order of orders creation.  
[**GetSprdOrdersPendingV5**](SpreadTradingAPI.md#GetSprdOrdersPendingV5) | **Get** /api/v5/sprd/orders-pending | Retrieve all incomplete orders under the current account.  
[**GetSprdPublicTradesV5**](SpreadTradingAPI.md#GetSprdPublicTradesV5) | **Get** /api/v5/sprd/public-trades | Retrieve the recent transactions of an instrument (at most 500 records per request). Results are returned in counter chronological order.   
[**GetSprdSpreadsV5**](SpreadTradingAPI.md#GetSprdSpreadsV5) | **Get** /api/v5/sprd/spreads | Retrieve all available spreads based on the request parameters.  
[**GetSprdTradesV5**](SpreadTradingAPI.md#GetSprdTradesV5) | **Get** /api/v5/sprd/trades | Retrieve historical transaction details . Results are returned in counter chronological order.  



## CreateSprdAmendOrderV5

> CreateSprdAmendOrderV5Resp CreateSprdAmendOrderV5(ctx).Execute()

Amend an incomplete order.  



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
	resp, r, err := apiClient.SpreadTradingAPI.CreateSprdAmendOrderV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.CreateSprdAmendOrderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSprdAmendOrderV5`: CreateSprdAmendOrderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.CreateSprdAmendOrderV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSprdAmendOrderV5Request struct via the builder pattern


### Return type

[**CreateSprdAmendOrderV5Resp**](CreateSprdAmendOrderV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSprdCancelAllAfterV5

> CreateSprdCancelAllAfterV5Resp CreateSprdCancelAllAfterV5(ctx).CreateSprdCancelAllAfterV5Req(createSprdCancelAllAfterV5Req).Execute()

Cancel all pending orders after the countdown timeout. Only applicable to spread trading.  



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
	createSprdCancelAllAfterV5Req := *openapiclient.NewCreateSprdCancelAllAfterV5Req("TimeOut_example") // CreateSprdCancelAllAfterV5Req | The request body for CreateSprdCancelAllAfterV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.CreateSprdCancelAllAfterV5(context.Background()).CreateSprdCancelAllAfterV5Req(createSprdCancelAllAfterV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.CreateSprdCancelAllAfterV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSprdCancelAllAfterV5`: CreateSprdCancelAllAfterV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.CreateSprdCancelAllAfterV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSprdCancelAllAfterV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSprdCancelAllAfterV5Req** | [**CreateSprdCancelAllAfterV5Req**](CreateSprdCancelAllAfterV5Req.md) | The request body for CreateSprdCancelAllAfterV5 | 

### Return type

[**CreateSprdCancelAllAfterV5Resp**](CreateSprdCancelAllAfterV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSprdCancelOrderV5

> CreateSprdCancelOrderV5Resp CreateSprdCancelOrderV5(ctx).CreateSprdCancelOrderV5Req(createSprdCancelOrderV5Req).Execute()

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
	createSprdCancelOrderV5Req := *openapiclient.NewCreateSprdCancelOrderV5Req() // CreateSprdCancelOrderV5Req | The request body for CreateSprdCancelOrderV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.CreateSprdCancelOrderV5(context.Background()).CreateSprdCancelOrderV5Req(createSprdCancelOrderV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.CreateSprdCancelOrderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSprdCancelOrderV5`: CreateSprdCancelOrderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.CreateSprdCancelOrderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSprdCancelOrderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSprdCancelOrderV5Req** | [**CreateSprdCancelOrderV5Req**](CreateSprdCancelOrderV5Req.md) | The request body for CreateSprdCancelOrderV5 | 

### Return type

[**CreateSprdCancelOrderV5Resp**](CreateSprdCancelOrderV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSprdMassCancelV5

> CreateSprdMassCancelV5Resp CreateSprdMassCancelV5(ctx).CreateSprdMassCancelV5Req(createSprdMassCancelV5Req).Execute()

Cancel all pending orders.  



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
	createSprdMassCancelV5Req := *openapiclient.NewCreateSprdMassCancelV5Req() // CreateSprdMassCancelV5Req | The request body for CreateSprdMassCancelV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.CreateSprdMassCancelV5(context.Background()).CreateSprdMassCancelV5Req(createSprdMassCancelV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.CreateSprdMassCancelV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSprdMassCancelV5`: CreateSprdMassCancelV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.CreateSprdMassCancelV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSprdMassCancelV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSprdMassCancelV5Req** | [**CreateSprdMassCancelV5Req**](CreateSprdMassCancelV5Req.md) | The request body for CreateSprdMassCancelV5 | 

### Return type

[**CreateSprdMassCancelV5Resp**](CreateSprdMassCancelV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSprdOrderV5

> CreateSprdOrderV5Resp CreateSprdOrderV5(ctx).CreateSprdOrderV5Req(createSprdOrderV5Req).Execute()

Place a new order  



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
	createSprdOrderV5Req := *openapiclient.NewCreateSprdOrderV5Req("OrdType_example", "Px_example", "Side_example", "SprdId_example", "Sz_example") // CreateSprdOrderV5Req | The request body for CreateSprdOrderV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.CreateSprdOrderV5(context.Background()).CreateSprdOrderV5Req(createSprdOrderV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.CreateSprdOrderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSprdOrderV5`: CreateSprdOrderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.CreateSprdOrderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSprdOrderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSprdOrderV5Req** | [**CreateSprdOrderV5Req**](CreateSprdOrderV5Req.md) | The request body for CreateSprdOrderV5 | 

### Return type

[**CreateSprdOrderV5Resp**](CreateSprdOrderV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketSprdCandlesV5

> GetMarketSprdCandlesV5Resp GetMarketSprdCandlesV5(ctx).SprdId(sprdId).Bar(bar).After(after).Before(before).Limit(limit).Execute()

Retrieve the candlestick charts. This endpoint can retrieve the latest 1,440 data entries. Charts are returned in groups based on the requested bar.  



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
	sprdId := "sprdId_example" // string | Spread ID (default to "")
	bar := "bar_example" // string | Bar size, the default is 1m, e.g. [1m/3m/5m/15m/30m/1H/2H/4H]   Hong Kong time opening price k-line:[6H/12H/1D/2D/3D/1W/1M/3M]   UTC time opening price k-line:[/6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/1Wutc/1Mutc/3Mutc] (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested ts (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested ts. The latest data will be returned when using before individually (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 300. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.GetMarketSprdCandlesV5(context.Background()).SprdId(sprdId).Bar(bar).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.GetMarketSprdCandlesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketSprdCandlesV5`: GetMarketSprdCandlesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.GetMarketSprdCandlesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketSprdCandlesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sprdId** | **string** | Spread ID | [default to &quot;&quot;]
 **bar** | **string** | Bar size, the default is 1m, e.g. [1m/3m/5m/15m/30m/1H/2H/4H]   Hong Kong time opening price k-line:[6H/12H/1D/2D/3D/1W/1M/3M]   UTC time opening price k-line:[/6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/1Wutc/1Mutc/3Mutc] | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested ts | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested ts. The latest data will be returned when using before individually | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 300. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetMarketSprdCandlesV5Resp**](GetMarketSprdCandlesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketSprdHistoryCandlesV5

> GetMarketSprdHistoryCandlesV5Resp GetMarketSprdHistoryCandlesV5(ctx).SprdId(sprdId).After(after).Before(before).Bar(bar).Limit(limit).Execute()

Retrieve history candlestick charts from recent years.  



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
	sprdId := "sprdId_example" // string | Spread ID (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested ts (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested ts. The latest data will be returned when using before individually (optional) (default to "")
	bar := "bar_example" // string | Bar size, the default is 1m, e.g. [1m/3m/5m/15m/30m/1H/2H/4H]   Hong Kong time opening price k-line:[6H/12H/1D/2D/3D/1W/1M/3M]   UTC time opening price k-line:[6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/1Wutc/1Mutc/3Mutc] (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.GetMarketSprdHistoryCandlesV5(context.Background()).SprdId(sprdId).After(after).Before(before).Bar(bar).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.GetMarketSprdHistoryCandlesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketSprdHistoryCandlesV5`: GetMarketSprdHistoryCandlesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.GetMarketSprdHistoryCandlesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketSprdHistoryCandlesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sprdId** | **string** | Spread ID | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested ts | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested ts. The latest data will be returned when using before individually | [default to &quot;&quot;]
 **bar** | **string** | Bar size, the default is 1m, e.g. [1m/3m/5m/15m/30m/1H/2H/4H]   Hong Kong time opening price k-line:[6H/12H/1D/2D/3D/1W/1M/3M]   UTC time opening price k-line:[6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/1Wutc/1Mutc/3Mutc] | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetMarketSprdHistoryCandlesV5Resp**](GetMarketSprdHistoryCandlesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketSprdTickerV5

> GetMarketSprdTickerV5Resp GetMarketSprdTickerV5(ctx).SprdId(sprdId).Execute()

Retrieve the latest price snapshot, best bid/ask price and quantity.  



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
	sprdId := "sprdId_example" // string | spread ID, e.g. BTC-USDT_BTC-USDT-SWAP (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.GetMarketSprdTickerV5(context.Background()).SprdId(sprdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.GetMarketSprdTickerV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketSprdTickerV5`: GetMarketSprdTickerV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.GetMarketSprdTickerV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketSprdTickerV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sprdId** | **string** | spread ID, e.g. BTC-USDT_BTC-USDT-SWAP | [default to &quot;&quot;]

### Return type

[**GetMarketSprdTickerV5Resp**](GetMarketSprdTickerV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSprdBooksV5

> GetSprdBooksV5Resp GetSprdBooksV5(ctx).SprdId(sprdId).Sz(sz).Execute()

Retrieve the order book of the spread.  



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
	sprdId := "sprdId_example" // string | spread ID, e.g. BTC-USDT_BTC-USDT-SWAP (default to "")
	sz := "sz_example" // string | Order book depth per side. Maximum value is 400. Default value is 5. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.GetSprdBooksV5(context.Background()).SprdId(sprdId).Sz(sz).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.GetSprdBooksV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSprdBooksV5`: GetSprdBooksV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.GetSprdBooksV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSprdBooksV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sprdId** | **string** | spread ID, e.g. BTC-USDT_BTC-USDT-SWAP | [default to &quot;&quot;]
 **sz** | **string** | Order book depth per side. Maximum value is 400. Default value is 5. | [default to &quot;&quot;]

### Return type

[**GetSprdBooksV5Resp**](GetSprdBooksV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSprdOrderV5

> GetSprdOrderV5Resp GetSprdOrderV5(ctx).OrdId(ordId).ClOrdId(clOrdId).Execute()

Retrieve order details.  



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
	ordId := "ordId_example" // string | Order ID   Either `ordId` or `clOrdId` is required, if both are passed, `ordId` will be used (optional) (default to "")
	clOrdId := "clOrdId_example" // string | Client Order ID as assigned by the client. The latest order will be returned. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.GetSprdOrderV5(context.Background()).OrdId(ordId).ClOrdId(clOrdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.GetSprdOrderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSprdOrderV5`: GetSprdOrderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.GetSprdOrderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSprdOrderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordId** | **string** | Order ID   Either &#x60;ordId&#x60; or &#x60;clOrdId&#x60; is required, if both are passed, &#x60;ordId&#x60; will be used | [default to &quot;&quot;]
 **clOrdId** | **string** | Client Order ID as assigned by the client. The latest order will be returned. | [default to &quot;&quot;]

### Return type

[**GetSprdOrderV5Resp**](GetSprdOrderV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSprdOrdersHistoryArchiveV5

> GetSprdOrdersHistoryArchiveV5Resp GetSprdOrdersHistoryArchiveV5(ctx).SprdId(sprdId).OrdType(ordType).State(state).InstType(instType).InstFamily(instFamily).BeginId(beginId).EndId(endId).Begin(begin).End(end).Limit(limit).Execute()

Retrieve the completed order data for the last 3 months, including those placed 3 months ago but completed in the last 3 months. Results are returned in counter chronological order.  



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
	sprdId := "sprdId_example" // string | spread ID, e.g. (optional) (default to "")
	ordType := "ordType_example" // string | Order type  `market`: Market order   `limit`: limit order   `post_only`: Post-only order   `ioc`: Immediate-or-cancel order (optional) (default to "")
	state := "state_example" // string | State   `canceled`   `filled` (optional) (default to "")
	instType := "instType_example" // string | Instrument type  `SPOT`  `FUTURES`  `SWAP`   Any orders with spreads containing the specified instrument type in any legs will be returned (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family, e.g. BTC-USDT. Any orders with spreads containing the specified instrument family in any legs will be returned (optional) (default to "")
	beginId := "beginId_example" // string | Start order ID the request to begin with. Pagination of data to return records newer than the requested order Id, not including beginId (optional) (default to "")
	endId := "endId_example" // string | End order ID the request to end with. Pagination of data to return records earlier than the requested order Id, not including endId (optional) (default to "")
	begin := "begin_example" // string | Filter with a begin timestamp. Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	end := "end_example" // string | Filter with an end timestamp. Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.GetSprdOrdersHistoryArchiveV5(context.Background()).SprdId(sprdId).OrdType(ordType).State(state).InstType(instType).InstFamily(instFamily).BeginId(beginId).EndId(endId).Begin(begin).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.GetSprdOrdersHistoryArchiveV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSprdOrdersHistoryArchiveV5`: GetSprdOrdersHistoryArchiveV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.GetSprdOrdersHistoryArchiveV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSprdOrdersHistoryArchiveV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sprdId** | **string** | spread ID, e.g. | [default to &quot;&quot;]
 **ordType** | **string** | Order type  &#x60;market&#x60;: Market order   &#x60;limit&#x60;: limit order   &#x60;post_only&#x60;: Post-only order   &#x60;ioc&#x60;: Immediate-or-cancel order | [default to &quot;&quot;]
 **state** | **string** | State   &#x60;canceled&#x60;   &#x60;filled&#x60; | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;FUTURES&#x60;  &#x60;SWAP&#x60;   Any orders with spreads containing the specified instrument type in any legs will be returned | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family, e.g. BTC-USDT. Any orders with spreads containing the specified instrument family in any legs will be returned | [default to &quot;&quot;]
 **beginId** | **string** | Start order ID the request to begin with. Pagination of data to return records newer than the requested order Id, not including beginId | [default to &quot;&quot;]
 **endId** | **string** | End order ID the request to end with. Pagination of data to return records earlier than the requested order Id, not including endId | [default to &quot;&quot;]
 **begin** | **string** | Filter with a begin timestamp. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **end** | **string** | Filter with an end timestamp. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100 | [default to &quot;&quot;]

### Return type

[**GetSprdOrdersHistoryArchiveV5Resp**](GetSprdOrdersHistoryArchiveV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSprdOrdersHistoryV5

> GetSprdOrdersHistoryV5Resp GetSprdOrdersHistoryV5(ctx).SprdId(sprdId).OrdType(ordType).State(state).BeginId(beginId).EndId(endId).Begin(begin).End(end).Limit(limit).Execute()

Retrieve the completed order data for the last 21 days, and the incomplete orders (filledSz =0 & state = canceled) that have been canceled are only reserved for 2 hours. Results are returned in counter chronological order of orders creation.  



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
	sprdId := "sprdId_example" // string | spread ID, e.g. (optional) (default to "")
	ordType := "ordType_example" // string | Order type  `market`: Market order   `limit`: limit order   `post_only`: Post-only order   `ioc`: Immediate-or-cancel order (optional) (default to "")
	state := "state_example" // string | State   `canceled`   `filled` (optional) (default to "")
	beginId := "beginId_example" // string | Start order ID the request to begin with. Pagination of data to return records newer than the requested order Id, not including beginId (optional) (default to "")
	endId := "endId_example" // string | End order ID the request to end with. Pagination of data to return records earlier than the requested order Id, not including endId (optional) (default to "")
	begin := "begin_example" // string | Filter with a begin timestamp. Unix timestamp format in milliseconds, e.g. `1597026383085`. Date older than 7 days will be truncated. (optional) (default to "")
	end := "end_example" // string | Filter with an end timestamp. Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.GetSprdOrdersHistoryV5(context.Background()).SprdId(sprdId).OrdType(ordType).State(state).BeginId(beginId).EndId(endId).Begin(begin).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.GetSprdOrdersHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSprdOrdersHistoryV5`: GetSprdOrdersHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.GetSprdOrdersHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSprdOrdersHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sprdId** | **string** | spread ID, e.g. | [default to &quot;&quot;]
 **ordType** | **string** | Order type  &#x60;market&#x60;: Market order   &#x60;limit&#x60;: limit order   &#x60;post_only&#x60;: Post-only order   &#x60;ioc&#x60;: Immediate-or-cancel order | [default to &quot;&quot;]
 **state** | **string** | State   &#x60;canceled&#x60;   &#x60;filled&#x60; | [default to &quot;&quot;]
 **beginId** | **string** | Start order ID the request to begin with. Pagination of data to return records newer than the requested order Id, not including beginId | [default to &quot;&quot;]
 **endId** | **string** | End order ID the request to end with. Pagination of data to return records earlier than the requested order Id, not including endId | [default to &quot;&quot;]
 **begin** | **string** | Filter with a begin timestamp. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;. Date older than 7 days will be truncated. | [default to &quot;&quot;]
 **end** | **string** | Filter with an end timestamp. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100 | [default to &quot;&quot;]

### Return type

[**GetSprdOrdersHistoryV5Resp**](GetSprdOrdersHistoryV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSprdOrdersPendingV5

> GetSprdOrdersPendingV5Resp GetSprdOrdersPendingV5(ctx).SprdId(sprdId).OrdType(ordType).State(state).BeginId(beginId).EndId(endId).Limit(limit).Execute()

Retrieve all incomplete orders under the current account.  



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
	sprdId := "sprdId_example" // string | spread ID, e.g. (optional) (default to "")
	ordType := "ordType_example" // string | Order type  `market`: Market order   `limit`: Limit order   `post_only`: Post-only order   `ioc`: Immediate-or-cancel order (optional) (default to "")
	state := "state_example" // string | State   `live`   `partially_filled` (optional) (default to "")
	beginId := "beginId_example" // string | Start order ID the request to begin with. Pagination of data to return records newer than the requested order Id, not including beginId (optional) (default to "")
	endId := "endId_example" // string | End order ID the request to end with. Pagination of data to return records earlier than the requested order Id, not including endId (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.GetSprdOrdersPendingV5(context.Background()).SprdId(sprdId).OrdType(ordType).State(state).BeginId(beginId).EndId(endId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.GetSprdOrdersPendingV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSprdOrdersPendingV5`: GetSprdOrdersPendingV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.GetSprdOrdersPendingV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSprdOrdersPendingV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sprdId** | **string** | spread ID, e.g. | [default to &quot;&quot;]
 **ordType** | **string** | Order type  &#x60;market&#x60;: Market order   &#x60;limit&#x60;: Limit order   &#x60;post_only&#x60;: Post-only order   &#x60;ioc&#x60;: Immediate-or-cancel order | [default to &quot;&quot;]
 **state** | **string** | State   &#x60;live&#x60;   &#x60;partially_filled&#x60; | [default to &quot;&quot;]
 **beginId** | **string** | Start order ID the request to begin with. Pagination of data to return records newer than the requested order Id, not including beginId | [default to &quot;&quot;]
 **endId** | **string** | End order ID the request to end with. Pagination of data to return records earlier than the requested order Id, not including endId | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100 | [default to &quot;&quot;]

### Return type

[**GetSprdOrdersPendingV5Resp**](GetSprdOrdersPendingV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSprdPublicTradesV5

> GetSprdPublicTradesV5Resp GetSprdPublicTradesV5(ctx).SprdId(sprdId).Execute()

Retrieve the recent transactions of an instrument (at most 500 records per request). Results are returned in counter chronological order.   



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
	sprdId := "sprdId_example" // string | Spread ID, e.g. BTC-USDT_BTC-USDT-SWAP (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.GetSprdPublicTradesV5(context.Background()).SprdId(sprdId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.GetSprdPublicTradesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSprdPublicTradesV5`: GetSprdPublicTradesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.GetSprdPublicTradesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSprdPublicTradesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sprdId** | **string** | Spread ID, e.g. BTC-USDT_BTC-USDT-SWAP | [default to &quot;&quot;]

### Return type

[**GetSprdPublicTradesV5Resp**](GetSprdPublicTradesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSprdSpreadsV5

> GetSprdSpreadsV5Resp GetSprdSpreadsV5(ctx).BaseCcy(baseCcy).InstId(instId).SprdId(sprdId).State(state).Execute()

Retrieve all available spreads based on the request parameters.  



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
	baseCcy := "baseCcy_example" // string | Currency instrument is based in, e.g. BTC, ETH (optional) (default to "")
	instId := "instId_example" // string | The instrument ID to be included in the spread. (optional) (default to "")
	sprdId := "sprdId_example" // string | The spread ID (optional) (default to "")
	state := "state_example" // string | Spreads which are available to trade, suspened or expired. Valid values include `live`, `suspend` and `expired`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.GetSprdSpreadsV5(context.Background()).BaseCcy(baseCcy).InstId(instId).SprdId(sprdId).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.GetSprdSpreadsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSprdSpreadsV5`: GetSprdSpreadsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.GetSprdSpreadsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSprdSpreadsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseCcy** | **string** | Currency instrument is based in, e.g. BTC, ETH | [default to &quot;&quot;]
 **instId** | **string** | The instrument ID to be included in the spread. | [default to &quot;&quot;]
 **sprdId** | **string** | The spread ID | [default to &quot;&quot;]
 **state** | **string** | Spreads which are available to trade, suspened or expired. Valid values include &#x60;live&#x60;, &#x60;suspend&#x60; and &#x60;expired&#x60;. | [default to &quot;&quot;]

### Return type

[**GetSprdSpreadsV5Resp**](GetSprdSpreadsV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSprdTradesV5

> GetSprdTradesV5Resp GetSprdTradesV5(ctx).SprdId(sprdId).TradeId(tradeId).OrdId(ordId).BeginId(beginId).EndId(endId).Begin(begin).End(end).Limit(limit).Execute()

Retrieve historical transaction details . Results are returned in counter chronological order.  



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
	sprdId := "sprdId_example" // string | spread ID, e.g. (optional) (default to "")
	tradeId := "tradeId_example" // string | Trade ID (optional) (default to "")
	ordId := "ordId_example" // string | Order ID (optional) (default to "")
	beginId := "beginId_example" // string | Start trade ID the request to begin with. Pagination of data to return records newer than the requested tradeId, not including beginId (optional) (default to "")
	endId := "endId_example" // string | End trade ID the request to end with. Pagination of data to return records earlier than the requested tradeId, not including endId (optional) (default to "")
	begin := "begin_example" // string | Filter with a begin timestamp. Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	end := "end_example" // string | Filter with an end timestamp. Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpreadTradingAPI.GetSprdTradesV5(context.Background()).SprdId(sprdId).TradeId(tradeId).OrdId(ordId).BeginId(beginId).EndId(endId).Begin(begin).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpreadTradingAPI.GetSprdTradesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSprdTradesV5`: GetSprdTradesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SpreadTradingAPI.GetSprdTradesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSprdTradesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sprdId** | **string** | spread ID, e.g. | [default to &quot;&quot;]
 **tradeId** | **string** | Trade ID | [default to &quot;&quot;]
 **ordId** | **string** | Order ID | [default to &quot;&quot;]
 **beginId** | **string** | Start trade ID the request to begin with. Pagination of data to return records newer than the requested tradeId, not including beginId | [default to &quot;&quot;]
 **endId** | **string** | End trade ID the request to end with. Pagination of data to return records earlier than the requested tradeId, not including endId | [default to &quot;&quot;]
 **begin** | **string** | Filter with a begin timestamp. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **end** | **string** | Filter with an end timestamp. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100 | [default to &quot;&quot;]

### Return type

[**GetSprdTradesV5Resp**](GetSprdTradesV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

