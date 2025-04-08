# \MarketDataAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketBooksFullV5**](MarketDataAPI.md#GetMarketBooksFullV5) | **Get** /api/v5/market/books-full | Retrieve order book of the instrument. The data will be updated once a second.  
[**GetMarketBooksV5**](MarketDataAPI.md#GetMarketBooksV5) | **Get** /api/v5/market/books | Retrieve order book of the instrument.  
[**GetMarketCallAuctionDetailsV5**](MarketDataAPI.md#GetMarketCallAuctionDetailsV5) | **Get** /api/v5/market/call-auction-details | Retrieve call auction details.  
[**GetMarketCandlesV5**](MarketDataAPI.md#GetMarketCandlesV5) | **Get** /api/v5/market/candles | Retrieve the candlestick charts. This endpoint can retrieve the latest 1,440 data entries. Charts are returned in groups based on the requested bar.   
[**GetMarketHistoryTradesV5**](MarketDataAPI.md#GetMarketHistoryTradesV5) | **Get** /api/v5/market/history-trades | Retrieve the recent transactions of an instrument from the last 3 months with pagination.  
[**GetMarketOptionInstrumentFamilyTradesV5**](MarketDataAPI.md#GetMarketOptionInstrumentFamilyTradesV5) | **Get** /api/v5/market/option/instrument-family-trades | Retrieve the recent transactions of an instrument under same instFamily. The maximum is 100.  
[**GetMarketPlatform24VolumeV5**](MarketDataAPI.md#GetMarketPlatform24VolumeV5) | **Get** /api/v5/market/platform-24-volume | The 24-hour trading volume is calculated on a rolling basis.  
[**GetMarketTickerV5**](MarketDataAPI.md#GetMarketTickerV5) | **Get** /api/v5/market/ticker | Retrieve the latest price snapshot, best bid/ask price, and trading volume in the last 24 hours.  
[**GetMarketTickersV5**](MarketDataAPI.md#GetMarketTickersV5) | **Get** /api/v5/market/tickers | Retrieve the latest price snapshot, best bid/ask price, and trading volume in the last 24 hours.  
[**GetMarketTradesV5**](MarketDataAPI.md#GetMarketTradesV5) | **Get** /api/v5/market/trades | Retrieve the recent transactions of an instrument.  
[**GetPublicOptionTradesV5**](MarketDataAPI.md#GetPublicOptionTradesV5) | **Get** /api/v5/public/option-trades | The maximum is 100.  



## GetMarketBooksFullV5

> GetMarketBooksFullV5Resp GetMarketBooksFullV5(ctx).InstId(instId).Sz(sz).Execute()

Retrieve order book of the instrument. The data will be updated once a second.  



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
	sz := "sz_example" // string | Order book depth per side. Maximum 5000, e.g. 5000 bids + 5000 asks    Default returns to `1` depth data. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetMarketBooksFullV5(context.Background()).InstId(instId).Sz(sz).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetMarketBooksFullV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketBooksFullV5`: GetMarketBooksFullV5Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetMarketBooksFullV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketBooksFullV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **sz** | **string** | Order book depth per side. Maximum 5000, e.g. 5000 bids + 5000 asks    Default returns to &#x60;1&#x60; depth data. | [default to &quot;&quot;]

### Return type

[**GetMarketBooksFullV5Resp**](GetMarketBooksFullV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketBooksV5

> GetMarketBooksV5Resp GetMarketBooksV5(ctx).InstId(instId).Sz(sz).Execute()

Retrieve order book of the instrument.  



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
	sz := "sz_example" // string | Order book depth per side. Maximum 400, e.g. 400 bids + 400 asks   Default returns to `1` depth data (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetMarketBooksV5(context.Background()).InstId(instId).Sz(sz).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetMarketBooksV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketBooksV5`: GetMarketBooksV5Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetMarketBooksV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketBooksV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **sz** | **string** | Order book depth per side. Maximum 400, e.g. 400 bids + 400 asks   Default returns to &#x60;1&#x60; depth data | [default to &quot;&quot;]

### Return type

[**GetMarketBooksV5Resp**](GetMarketBooksV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketCallAuctionDetailsV5

> GetMarketCallAuctionDetailsV5Resp GetMarketCallAuctionDetailsV5(ctx).InstId(instId).Execute()

Retrieve call auction details.  



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetMarketCallAuctionDetailsV5(context.Background()).InstId(instId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetMarketCallAuctionDetailsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketCallAuctionDetailsV5`: GetMarketCallAuctionDetailsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetMarketCallAuctionDetailsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketCallAuctionDetailsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]

### Return type

[**GetMarketCallAuctionDetailsV5Resp**](GetMarketCallAuctionDetailsV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketCandlesV5

> GetMarketCandlesV5Resp GetMarketCandlesV5(ctx).InstId(instId).Bar(bar).After(after).Before(before).Limit(limit).Execute()

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
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT` (default to "")
	bar := "bar_example" // string | Bar size, the default is `1m`  e.g. [1m/3m/5m/15m/30m/1H/2H/4H]   Hong Kong time opening price k-line: [6H/12H/1D/2D/3D/1W/1M/3M]  UTC time opening price k-line: [/6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/1Wutc/1Mutc/3Mutc] (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ts` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ts`. The latest data will be returned when using `before` individually (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `300`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetMarketCandlesV5(context.Background()).InstId(instId).Bar(bar).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetMarketCandlesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketCandlesV5`: GetMarketCandlesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetMarketCandlesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketCandlesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **bar** | **string** | Bar size, the default is &#x60;1m&#x60;  e.g. [1m/3m/5m/15m/30m/1H/2H/4H]   Hong Kong time opening price k-line: [6H/12H/1D/2D/3D/1W/1M/3M]  UTC time opening price k-line: [/6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/1Wutc/1Mutc/3Mutc] | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ts&#x60;. The latest data will be returned when using &#x60;before&#x60; individually | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;300&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetMarketCandlesV5Resp**](GetMarketCandlesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketHistoryTradesV5

> GetMarketHistoryTradesV5Resp GetMarketHistoryTradesV5(ctx).InstId(instId).Type_(type_).After(after).Before(before).Limit(limit).Execute()

Retrieve the recent transactions of an instrument from the last 3 months with pagination.  



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
	type_ := "type__example" // string | Pagination Type   `1`: tradeId `2`: timestamp  The default is `1` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested tradeId or ts. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested tradeId.   Do not support timestamp for pagination. The latest data will be returned when using `before` individually (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum and default both are `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetMarketHistoryTradesV5(context.Background()).InstId(instId).Type_(type_).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetMarketHistoryTradesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketHistoryTradesV5`: GetMarketHistoryTradesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetMarketHistoryTradesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketHistoryTradesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **type_** | **string** | Pagination Type   &#x60;1&#x60;: tradeId &#x60;2&#x60;: timestamp  The default is &#x60;1&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested tradeId or ts. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested tradeId.   Do not support timestamp for pagination. The latest data will be returned when using &#x60;before&#x60; individually | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum and default both are &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetMarketHistoryTradesV5Resp**](GetMarketHistoryTradesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketOptionInstrumentFamilyTradesV5

> GetMarketOptionInstrumentFamilyTradesV5Resp GetMarketOptionInstrumentFamilyTradesV5(ctx).InstFamily(instFamily).Execute()

Retrieve the recent transactions of an instrument under same instFamily. The maximum is 100.  



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
	instFamily := "instFamily_example" // string | Instrument family, e.g. BTC-USD  Applicable to `OPTION` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetMarketOptionInstrumentFamilyTradesV5(context.Background()).InstFamily(instFamily).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetMarketOptionInstrumentFamilyTradesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketOptionInstrumentFamilyTradesV5`: GetMarketOptionInstrumentFamilyTradesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetMarketOptionInstrumentFamilyTradesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketOptionInstrumentFamilyTradesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instFamily** | **string** | Instrument family, e.g. BTC-USD  Applicable to &#x60;OPTION&#x60; | [default to &quot;&quot;]

### Return type

[**GetMarketOptionInstrumentFamilyTradesV5Resp**](GetMarketOptionInstrumentFamilyTradesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketPlatform24VolumeV5

> GetMarketPlatform24VolumeV5Resp GetMarketPlatform24VolumeV5(ctx).Execute()

The 24-hour trading volume is calculated on a rolling basis.  



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
	resp, r, err := apiClient.MarketDataAPI.GetMarketPlatform24VolumeV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetMarketPlatform24VolumeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketPlatform24VolumeV5`: GetMarketPlatform24VolumeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetMarketPlatform24VolumeV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketPlatform24VolumeV5Request struct via the builder pattern


### Return type

[**GetMarketPlatform24VolumeV5Resp**](GetMarketPlatform24VolumeV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketTickerV5

> GetMarketTickerV5Resp GetMarketTickerV5(ctx).InstId(instId).Execute()

Retrieve the latest price snapshot, best bid/ask price, and trading volume in the last 24 hours.  



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
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USD-SWAP` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetMarketTickerV5(context.Background()).InstId(instId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetMarketTickerV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketTickerV5`: GetMarketTickerV5Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetMarketTickerV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketTickerV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USD-SWAP&#x60; | [default to &quot;&quot;]

### Return type

[**GetMarketTickerV5Resp**](GetMarketTickerV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketTickersV5

> GetMarketTickersV5Resp GetMarketTickersV5(ctx).InstType(instType).Uly(uly).InstFamily(instFamily).Execute()

Retrieve the latest price snapshot, best bid/ask price, and trading volume in the last 24 hours.  



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
	instType := "instType_example" // string | Instrument type  `SPOT`  `SWAP`  `FUTURES`  `OPTION` (default to "")
	uly := "uly_example" // string | Underlying, e.g. `BTC-USD`   Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetMarketTickersV5(context.Background()).InstType(instType).Uly(uly).InstFamily(instFamily).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetMarketTickersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketTickersV5`: GetMarketTickersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetMarketTickersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketTickersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Underlying, e.g. &#x60;BTC-USD&#x60;   Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]

### Return type

[**GetMarketTickersV5Resp**](GetMarketTickersV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketTradesV5

> GetMarketTradesV5Resp GetMarketTradesV5(ctx).InstId(instId).Limit(limit).Execute()

Retrieve the recent transactions of an instrument.  



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
	limit := "limit_example" // string | Number of results per request. The maximum is `500`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetMarketTradesV5(context.Background()).InstId(instId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetMarketTradesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketTradesV5`: GetMarketTradesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetMarketTradesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketTradesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;500&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetMarketTradesV5Resp**](GetMarketTradesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicOptionTradesV5

> GetPublicOptionTradesV5Resp GetPublicOptionTradesV5(ctx).InstId(instId).InstFamily(instFamily).OptType(optType).Execute()

The maximum is 100.  



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
	instId := "instId_example" // string | Instrument ID, e.g. BTC-USD-221230-4000-C, Either `instId` or `instFamily` is required. If both are passed, `instId` will be used. (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family, e.g. BTC-USD (optional) (default to "")
	optType := "optType_example" // string | Option type, `C`: Call  `P`: put (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetPublicOptionTradesV5(context.Background()).InstId(instId).InstFamily(instFamily).OptType(optType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetPublicOptionTradesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicOptionTradesV5`: GetPublicOptionTradesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetPublicOptionTradesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicOptionTradesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. BTC-USD-221230-4000-C, Either &#x60;instId&#x60; or &#x60;instFamily&#x60; is required. If both are passed, &#x60;instId&#x60; will be used. | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family, e.g. BTC-USD | [default to &quot;&quot;]
 **optType** | **string** | Option type, &#x60;C&#x60;: Call  &#x60;P&#x60;: put | [default to &quot;&quot;]

### Return type

[**GetPublicOptionTradesV5Resp**](GetPublicOptionTradesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

