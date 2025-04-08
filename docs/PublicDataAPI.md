# \PublicDataAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketExchangeRateV5**](PublicDataAPI.md#GetMarketExchangeRateV5) | **Get** /api/v5/market/exchange-rate | This interface provides the average exchange rate data for 2 weeks  
[**GetMarketHistoryIndexCandlesV5**](PublicDataAPI.md#GetMarketHistoryIndexCandlesV5) | **Get** /api/v5/market/history-index-candles | Retrieve the candlestick charts of the index from recent years.  
[**GetMarketHistoryMarkPriceCandlesV5**](PublicDataAPI.md#GetMarketHistoryMarkPriceCandlesV5) | **Get** /api/v5/market/history-mark-price-candles | Retrieve the candlestick charts of mark price from recent years.  
[**GetMarketIndexCandlesV5**](PublicDataAPI.md#GetMarketIndexCandlesV5) | **Get** /api/v5/market/index-candles | Retrieve the candlestick charts of the index. This endpoint can retrieve the latest 1,440 data entries. Charts are returned in groups based on the requested bar.   
[**GetMarketIndexComponentsV5**](PublicDataAPI.md#GetMarketIndexComponentsV5) | **Get** /api/v5/market/index-components | Get the index component information data on the market  
[**GetMarketIndexTickersV5**](PublicDataAPI.md#GetMarketIndexTickersV5) | **Get** /api/v5/market/index-tickers | Retrieve index tickers.  
[**GetMarketMarkPriceCandlesV5**](PublicDataAPI.md#GetMarketMarkPriceCandlesV5) | **Get** /api/v5/market/mark-price-candles | Retrieve the candlestick charts of mark price. This endpoint can retrieve the latest 1,440 data entries. Charts are returned in groups based on the requested bar.  
[**GetPublicConvertContractCoinV5**](PublicDataAPI.md#GetPublicConvertContractCoinV5) | **Get** /api/v5/public/convert-contract-coin | Convert the crypto value to the number of contracts, or vice versa  
[**GetPublicDeliveryExerciseHistoryV5**](PublicDataAPI.md#GetPublicDeliveryExerciseHistoryV5) | **Get** /api/v5/public/delivery-exercise-history | Retrieve delivery records of Futures and exercise records of Options in the last 3 months.  
[**GetPublicDiscountRateInterestFreeQuotaV5**](PublicDataAPI.md#GetPublicDiscountRateInterestFreeQuotaV5) | **Get** /api/v5/public/discount-rate-interest-free-quota | Retrieve discount rate level and interest-free quota.  
[**GetPublicEconomicCalendarV5**](PublicDataAPI.md#GetPublicEconomicCalendarV5) | **Get** /api/v5/public/economic-calendar | Get the macro-economic calendar data within 3 months. Historical data from 3 months ago is only available to users with trading fee tier VIP1 and above.  
[**GetPublicEstimatedPriceV5**](PublicDataAPI.md#GetPublicEstimatedPriceV5) | **Get** /api/v5/public/estimated-price | Retrieve the estimated delivery price which will only have a return value one hour before the delivery/exercise.  
[**GetPublicEstimatedSettlementInfoV5**](PublicDataAPI.md#GetPublicEstimatedSettlementInfoV5) | **Get** /api/v5/public/estimated-settlement-info | Retrieve the estimated settlement price which will only have a return value one hour before the settlement.  
[**GetPublicFundingRateHistoryV5**](PublicDataAPI.md#GetPublicFundingRateHistoryV5) | **Get** /api/v5/public/funding-rate-history | Retrieve funding rate history. This endpoint can retrieve data from the last 3 months.  
[**GetPublicFundingRateV5**](PublicDataAPI.md#GetPublicFundingRateV5) | **Get** /api/v5/public/funding-rate | Retrieve funding rate.  
[**GetPublicInstrumentTickBandsV5**](PublicDataAPI.md#GetPublicInstrumentTickBandsV5) | **Get** /api/v5/public/instrument-tick-bands | Get option tick bands information  
[**GetPublicInstrumentsV5**](PublicDataAPI.md#GetPublicInstrumentsV5) | **Get** /api/v5/public/instruments | Retrieve a list of instruments with open contracts for OKX. Retrieve available instruments info of current account, please refer to .  
[**GetPublicInsuranceFundV5**](PublicDataAPI.md#GetPublicInsuranceFundV5) | **Get** /api/v5/public/insurance-fund | Get insurance fund balance information  
[**GetPublicInterestRateLoanQuotaV5**](PublicDataAPI.md#GetPublicInterestRateLoanQuotaV5) | **Get** /api/v5/public/interest-rate-loan-quota | Retrieve interest rate  
[**GetPublicMarkPriceV5**](PublicDataAPI.md#GetPublicMarkPriceV5) | **Get** /api/v5/public/mark-price | Retrieve mark price.  We set the mark price based on the SPOT index and at a reasonable basis to prevent individual users from manipulating the market and causing the contract price to fluctuate.  
[**GetPublicOpenInterestV5**](PublicDataAPI.md#GetPublicOpenInterestV5) | **Get** /api/v5/public/open-interest | Retrieve the total open interest for contracts on OKX.  
[**GetPublicOptSummaryV5**](PublicDataAPI.md#GetPublicOptSummaryV5) | **Get** /api/v5/public/opt-summary | Retrieve option market data.  
[**GetPublicPositionTiersV5**](PublicDataAPI.md#GetPublicPositionTiersV5) | **Get** /api/v5/public/position-tiers | Retrieve position tiers information, maximum leverage depends on your borrowings and margin ratio.  
[**GetPublicPremiumHistoryV5**](PublicDataAPI.md#GetPublicPremiumHistoryV5) | **Get** /api/v5/public/premium-history | It will return premium data in the past 6 months.  
[**GetPublicPriceLimitV5**](PublicDataAPI.md#GetPublicPriceLimitV5) | **Get** /api/v5/public/price-limit | Retrieve the highest buy limit and lowest sell limit of the instrument.  
[**GetPublicSettlementHistoryV5**](PublicDataAPI.md#GetPublicSettlementHistoryV5) | **Get** /api/v5/public/settlement-history | Retrieve settlement records of futures in the last 3 months.  
[**GetPublicTimeV5**](PublicDataAPI.md#GetPublicTimeV5) | **Get** /api/v5/public/time | Retrieve API server time.  
[**GetPublicUnderlyingV5**](PublicDataAPI.md#GetPublicUnderlyingV5) | **Get** /api/v5/public/underlying | 



## GetMarketExchangeRateV5

> GetMarketExchangeRateV5Resp GetMarketExchangeRateV5(ctx).Execute()

This interface provides the average exchange rate data for 2 weeks  



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
	resp, r, err := apiClient.PublicDataAPI.GetMarketExchangeRateV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetMarketExchangeRateV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketExchangeRateV5`: GetMarketExchangeRateV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetMarketExchangeRateV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketExchangeRateV5Request struct via the builder pattern


### Return type

[**GetMarketExchangeRateV5Resp**](GetMarketExchangeRateV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketHistoryIndexCandlesV5

> GetMarketHistoryIndexCandlesV5Resp GetMarketHistoryIndexCandlesV5(ctx).InstId(instId).After(after).Before(before).Bar(bar).Limit(limit).Execute()

Retrieve the candlestick charts of the index from recent years.  



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
	instId := "instId_example" // string | Index, e.g. `BTC-USD` (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ts` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ts`. The latest data will be returned when using `before` individually (optional) (default to "")
	bar := "bar_example" // string | Bar size, the default is `1m`  e.g. [1m/3m/5m/15m/30m/1H/2H/4H]   Hong Kong time opening price k-line: [6H/12H/1D/1W/1M]  UTC time opening price k-line: [/6Hutc/12Hutc/1Dutc/1Wutc/1Mutc] (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetMarketHistoryIndexCandlesV5(context.Background()).InstId(instId).After(after).Before(before).Bar(bar).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetMarketHistoryIndexCandlesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketHistoryIndexCandlesV5`: GetMarketHistoryIndexCandlesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetMarketHistoryIndexCandlesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketHistoryIndexCandlesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Index, e.g. &#x60;BTC-USD&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ts&#x60;. The latest data will be returned when using &#x60;before&#x60; individually | [default to &quot;&quot;]
 **bar** | **string** | Bar size, the default is &#x60;1m&#x60;  e.g. [1m/3m/5m/15m/30m/1H/2H/4H]   Hong Kong time opening price k-line: [6H/12H/1D/1W/1M]  UTC time opening price k-line: [/6Hutc/12Hutc/1Dutc/1Wutc/1Mutc] | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetMarketHistoryIndexCandlesV5Resp**](GetMarketHistoryIndexCandlesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketHistoryMarkPriceCandlesV5

> GetMarketHistoryMarkPriceCandlesV5Resp GetMarketHistoryMarkPriceCandlesV5(ctx).InstId(instId).After(after).Before(before).Bar(bar).Limit(limit).Execute()

Retrieve the candlestick charts of mark price from recent years.  



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
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ts` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ts`. The latest data will be returned when using `before` individually (optional) (default to "")
	bar := "bar_example" // string | Bar size, the default is `1m`  e.g. [1m/3m/5m/15m/30m/1H/2H/4H]   Hong Kong time opening price k-line: [6H/12H/1D/1W/1M]  UTC time opening price k-line: [6Hutc/12Hutc/1Dutc/1Wutc/1Mutc] (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetMarketHistoryMarkPriceCandlesV5(context.Background()).InstId(instId).After(after).Before(before).Bar(bar).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetMarketHistoryMarkPriceCandlesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketHistoryMarkPriceCandlesV5`: GetMarketHistoryMarkPriceCandlesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetMarketHistoryMarkPriceCandlesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketHistoryMarkPriceCandlesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USD-SWAP&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ts&#x60;. The latest data will be returned when using &#x60;before&#x60; individually | [default to &quot;&quot;]
 **bar** | **string** | Bar size, the default is &#x60;1m&#x60;  e.g. [1m/3m/5m/15m/30m/1H/2H/4H]   Hong Kong time opening price k-line: [6H/12H/1D/1W/1M]  UTC time opening price k-line: [6Hutc/12Hutc/1Dutc/1Wutc/1Mutc] | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetMarketHistoryMarkPriceCandlesV5Resp**](GetMarketHistoryMarkPriceCandlesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketIndexCandlesV5

> GetMarketIndexCandlesV5Resp GetMarketIndexCandlesV5(ctx).InstId(instId).After(after).Before(before).Bar(bar).Limit(limit).Execute()

Retrieve the candlestick charts of the index. This endpoint can retrieve the latest 1,440 data entries. Charts are returned in groups based on the requested bar.   



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
	instId := "instId_example" // string | Index, e.g. `BTC-USD` (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ts` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ts`. The latest data will be returned when using `before` individually (optional) (default to "")
	bar := "bar_example" // string | Bar size, the default is `1m`  e.g. [`1m`/`3m`/`5m`/`15m`/`30m`/`1H`/`2H`/`4H`]   Hong Kong time opening price k-line: [`6H`/`12H`/`1D`/`1W`/`1M`/`3M`]  UTC time opening price k-line: [`6Hutc`/`12Hutc`/`1Dutc`/`1Wutc`/`1Mutc`/`3Mutc`] (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetMarketIndexCandlesV5(context.Background()).InstId(instId).After(after).Before(before).Bar(bar).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetMarketIndexCandlesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketIndexCandlesV5`: GetMarketIndexCandlesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetMarketIndexCandlesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketIndexCandlesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Index, e.g. &#x60;BTC-USD&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ts&#x60;. The latest data will be returned when using &#x60;before&#x60; individually | [default to &quot;&quot;]
 **bar** | **string** | Bar size, the default is &#x60;1m&#x60;  e.g. [&#x60;1m&#x60;/&#x60;3m&#x60;/&#x60;5m&#x60;/&#x60;15m&#x60;/&#x60;30m&#x60;/&#x60;1H&#x60;/&#x60;2H&#x60;/&#x60;4H&#x60;]   Hong Kong time opening price k-line: [&#x60;6H&#x60;/&#x60;12H&#x60;/&#x60;1D&#x60;/&#x60;1W&#x60;/&#x60;1M&#x60;/&#x60;3M&#x60;]  UTC time opening price k-line: [&#x60;6Hutc&#x60;/&#x60;12Hutc&#x60;/&#x60;1Dutc&#x60;/&#x60;1Wutc&#x60;/&#x60;1Mutc&#x60;/&#x60;3Mutc&#x60;] | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetMarketIndexCandlesV5Resp**](GetMarketIndexCandlesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketIndexComponentsV5

> GetMarketIndexComponentsV5Resp GetMarketIndexComponentsV5(ctx).Index(index).Execute()

Get the index component information data on the market  



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
	index := "index_example" // string | index, e.g `BTC-USDT` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetMarketIndexComponentsV5(context.Background()).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetMarketIndexComponentsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketIndexComponentsV5`: GetMarketIndexComponentsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetMarketIndexComponentsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketIndexComponentsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **string** | index, e.g &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]

### Return type

[**GetMarketIndexComponentsV5Resp**](GetMarketIndexComponentsV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketIndexTickersV5

> GetMarketIndexTickersV5Resp GetMarketIndexTickersV5(ctx).QuoteCcy(quoteCcy).InstId(instId).Execute()

Retrieve index tickers.  



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
	quoteCcy := "quoteCcy_example" // string | Quote currency   Currently there is only an index with `USD/USDT/BTC/USDC` as the quote currency. (optional) (default to "")
	instId := "instId_example" // string | Index, e.g. `BTC-USD`  Either `quoteCcy` or `instId` is required. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetMarketIndexTickersV5(context.Background()).QuoteCcy(quoteCcy).InstId(instId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetMarketIndexTickersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketIndexTickersV5`: GetMarketIndexTickersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetMarketIndexTickersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketIndexTickersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quoteCcy** | **string** | Quote currency   Currently there is only an index with &#x60;USD/USDT/BTC/USDC&#x60; as the quote currency. | [default to &quot;&quot;]
 **instId** | **string** | Index, e.g. &#x60;BTC-USD&#x60;  Either &#x60;quoteCcy&#x60; or &#x60;instId&#x60; is required. | [default to &quot;&quot;]

### Return type

[**GetMarketIndexTickersV5Resp**](GetMarketIndexTickersV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketMarkPriceCandlesV5

> GetMarketMarkPriceCandlesV5Resp GetMarketMarkPriceCandlesV5(ctx).InstId(instId).After(after).Before(before).Bar(bar).Limit(limit).Execute()

Retrieve the candlestick charts of mark price. This endpoint can retrieve the latest 1,440 data entries. Charts are returned in groups based on the requested bar.  



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
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ts` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ts`. The latest data will be returned when using `before` individually (optional) (default to "")
	bar := "bar_example" // string | Bar size, the default is `1m`  e.g. [1m/3m/5m/15m/30m/1H/2H/4H]   Hong Kong time opening price k-line: [6H/12H/1D/1W/1M/3M]  UTC time opening price k-line: [6Hutc/12Hutc/1Dutc/1Wutc/1Mutc/3Mutc] (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetMarketMarkPriceCandlesV5(context.Background()).InstId(instId).After(after).Before(before).Bar(bar).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetMarketMarkPriceCandlesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketMarkPriceCandlesV5`: GetMarketMarkPriceCandlesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetMarketMarkPriceCandlesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketMarkPriceCandlesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USD-SWAP&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ts&#x60;. The latest data will be returned when using &#x60;before&#x60; individually | [default to &quot;&quot;]
 **bar** | **string** | Bar size, the default is &#x60;1m&#x60;  e.g. [1m/3m/5m/15m/30m/1H/2H/4H]   Hong Kong time opening price k-line: [6H/12H/1D/1W/1M/3M]  UTC time opening price k-line: [6Hutc/12Hutc/1Dutc/1Wutc/1Mutc/3Mutc] | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetMarketMarkPriceCandlesV5Resp**](GetMarketMarkPriceCandlesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicConvertContractCoinV5

> GetPublicConvertContractCoinV5Resp GetPublicConvertContractCoinV5(ctx).InstId(instId).Sz(sz).Type_(type_).Px(px).Unit(unit).OpType(opType).Execute()

Convert the crypto value to the number of contracts, or vice versa  



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
	instId := "instId_example" // string | Instrument ID  only applicable to `FUTURES`/`SWAP`/`OPTION` (default to "")
	sz := "sz_example" // string | Quantity to buy or sell  It is quantity of currency while converting currency to contract;   It is quantity of contract while converting contract to currency. (default to "")
	type_ := "type__example" // string | Convert type  `1`: Convert currency to contract  `2`: Convert contract to currency  The default is `1` (optional) (default to "")
	px := "px_example" // string | Order price  For crypto-margined contracts, it is necessary while converting.  For USDT-margined contracts, it is necessary while converting between usdt and contract.  It is optional while converting between coin and contract.   For OPTION, it is optional. (optional) (default to "")
	unit := "unit_example" // string | The unit of currency  `coin`  `usds`: USDT/USDC  only applicable to USDⓈ-margined contracts from `FUTURES`/`SWAP` (optional) (default to "")
	opType := "opType_example" // string | Order type  `open`: round down sz when opening positions   `close`: round sz to the nearest when closing positions   The default is `close`   Applicable to `FUTURES` `SWAP` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicConvertContractCoinV5(context.Background()).InstId(instId).Sz(sz).Type_(type_).Px(px).Unit(unit).OpType(opType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicConvertContractCoinV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicConvertContractCoinV5`: GetPublicConvertContractCoinV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicConvertContractCoinV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicConvertContractCoinV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID  only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **sz** | **string** | Quantity to buy or sell  It is quantity of currency while converting currency to contract;   It is quantity of contract while converting contract to currency. | [default to &quot;&quot;]
 **type_** | **string** | Convert type  &#x60;1&#x60;: Convert currency to contract  &#x60;2&#x60;: Convert contract to currency  The default is &#x60;1&#x60; | [default to &quot;&quot;]
 **px** | **string** | Order price  For crypto-margined contracts, it is necessary while converting.  For USDT-margined contracts, it is necessary while converting between usdt and contract.  It is optional while converting between coin and contract.   For OPTION, it is optional. | [default to &quot;&quot;]
 **unit** | **string** | The unit of currency  &#x60;coin&#x60;  &#x60;usds&#x60;: USDT/USDC  only applicable to USDⓈ-margined contracts from &#x60;FUTURES&#x60;/&#x60;SWAP&#x60; | [default to &quot;&quot;]
 **opType** | **string** | Order type  &#x60;open&#x60;: round down sz when opening positions   &#x60;close&#x60;: round sz to the nearest when closing positions   The default is &#x60;close&#x60;   Applicable to &#x60;FUTURES&#x60; &#x60;SWAP&#x60; | [default to &quot;&quot;]

### Return type

[**GetPublicConvertContractCoinV5Resp**](GetPublicConvertContractCoinV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicDeliveryExerciseHistoryV5

> GetPublicDeliveryExerciseHistoryV5Resp GetPublicDeliveryExerciseHistoryV5(ctx).InstType(instType).Uly(uly).InstFamily(instFamily).After(after).Before(before).Limit(limit).Execute()

Retrieve delivery records of Futures and exercise records of Options in the last 3 months.  



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
	instType := "instType_example" // string | Instrument type  `FUTURES`  `OPTION` (default to "")
	uly := "uly_example" // string | Underlying, only applicable to `FUTURES`/`OPTION`  Either `uly` or `instFamily` is required. If both are passed, `instFamily` will be used. (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family, only applicable to `FUTURES`/`OPTION`  Either `uly` or `instFamily` is required. If both are passed, `instFamily` will be used. (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ts` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ts` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicDeliveryExerciseHistoryV5(context.Background()).InstType(instType).Uly(uly).InstFamily(instFamily).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicDeliveryExerciseHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicDeliveryExerciseHistoryV5`: GetPublicDeliveryExerciseHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicDeliveryExerciseHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicDeliveryExerciseHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Underlying, only applicable to &#x60;FUTURES&#x60;/&#x60;OPTION&#x60;  Either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. If both are passed, &#x60;instFamily&#x60; will be used. | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family, only applicable to &#x60;FUTURES&#x60;/&#x60;OPTION&#x60;  Either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. If both are passed, &#x60;instFamily&#x60; will be used. | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetPublicDeliveryExerciseHistoryV5Resp**](GetPublicDeliveryExerciseHistoryV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicDiscountRateInterestFreeQuotaV5

> GetPublicDiscountRateInterestFreeQuotaV5Resp GetPublicDiscountRateInterestFreeQuotaV5(ctx).Ccy(ccy).DiscountLv(discountLv).Execute()

Retrieve discount rate level and interest-free quota.  



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
	ccy := "ccy_example" // string | Currency (optional) (default to "")
	discountLv := "discountLv_example" // string | Discount level (Deprecated)  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicDiscountRateInterestFreeQuotaV5(context.Background()).Ccy(ccy).DiscountLv(discountLv).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicDiscountRateInterestFreeQuotaV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicDiscountRateInterestFreeQuotaV5`: GetPublicDiscountRateInterestFreeQuotaV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicDiscountRateInterestFreeQuotaV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicDiscountRateInterestFreeQuotaV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency | [default to &quot;&quot;]
 **discountLv** | **string** | Discount level (Deprecated)  | [default to &quot;&quot;]

### Return type

[**GetPublicDiscountRateInterestFreeQuotaV5Resp**](GetPublicDiscountRateInterestFreeQuotaV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicEconomicCalendarV5

> GetPublicEconomicCalendarV5Resp GetPublicEconomicCalendarV5(ctx).Region(region).Importance(importance).Before(before).After(after).Limit(limit).Execute()

Get the macro-economic calendar data within 3 months. Historical data from 3 months ago is only available to users with trading fee tier VIP1 and above.  



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
	region := "region_example" // string | Country, region or entity   `afghanistan`, `albania`, `algeria`, `andorra`, `angola`, `antigua_and_barbuda`, `argentina`, `armenia`, `aruba`, `australia`, `austria`, `azerbaijan`, `bahamas`, `bahrain`, `bangladesh`, `barbados`, `belarus`, `belgium`, `belize`, `benin`, `bermuda`, `bhutan`, `bolivia`, `bosnia_and_herzegovina`, `botswana`, `brazil`, `brunei`, `bulgaria`, `burkina_faso`, `burundi`, `cambodia`, `cameroon`, `canada`, `cape_verde`, `cayman_islands`, `central_african_republic`, `chad`, `chile`, `china`, `colombia`, `comoros`, `congo`, `costa_rica`, `croatia`, `cuba`, `cyprus`, `czech_republic`, `denmark`, `djibouti`, `dominica`, `dominican_republic`, `east_timor`, `ecuador`, `egypt`, `el_salvador`, `equatorial_guinea`, `eritrea`, `estonia`, `ethiopia`, `euro_area`, `european_union`, `faroe_islands`, `fiji`, `finland`, `france`, `g20`, `g7`, `gabon`, `gambia`, `georgia`, `germany`, `ghana`, `greece`, `greenland`, `grenada`, `guatemala`, `guinea`, `guinea_bissau`, `guyana`, `hungary`, `haiti`, `honduras`, `hong_kong`, `hungary`, `imf`, `indonesia`, `iceland`, `india`, `indonesia`, `iran`, `iraq`, `ireland`, `isle_of_man`, `israel`, `italy`, `ivory_coast`, `jamaica`, `japan`, `jordan`, `kazakhstan`, `kenya`, `kiribati`, `kosovo`, `kuwait`, `kyrgyzstan`, `laos`, `latvia`, `lebanon`, `lesotho`, `liberia`, `libya`, `liechtenstein`, `lithuania`, `luxembourg`, `macau`, `macedonia`, `madagascar`, `malawi`, `malaysia`, `maldives`, `mali`, `malta`, `mauritania`, `mauritius`, `mexico`, `micronesia`, `moldova`, `monaco`, `mongolia`, `montenegro`, `morocco`, `mozambique`, `myanmar`, `namibia`, `nepal`, `netherlands`, `new_caledonia`, `new_zealand`, `nicaragua`, `niger`, `nigeria`, `north_korea`, `northern_mariana_islands`, `norway`, `opec`, `oman`, `pakistan`, `palau`, `palestine`, `panama`, `papua_new_guinea`, `paraguay`, `peru`, `philippines`, `poland`, `portugal`, `puerto_rico`, `qatar`, `russia`, `republic_of_the_congo`, `romania`, `russia`, `rwanda`, `slovakia`, `samoa`, `san_marino`, `sao_tome_and_principe`, `saudi_arabia`, `senegal`, `serbia`, `seychelles`, `sierra_leone`, `singapore`, `slovakia`, `slovenia`, `solomon_islands`, `somalia`, `south_africa`, `south_korea`, `south_sudan`, `spain`, `sri_lanka`, `st_kitts_and_nevis`, `st_lucia`, `sudan`, `suriname`, `swaziland`, `sweden`, `switzerland`, `syria`, `taiwan`, `tajikistan`, `tanzania`, `thailand`, `togo`, `tonga`, `trinidad_and_tobago`, `tunisia`, `turkey`, `turkmenistan`, `uganda`, `ukraine`, `united_arab_emirates`, `united_kingdom`, `united_states`, `uruguay`, `uzbekistan`, `vanuatu`, `venezuela`, `vietnam`, `world`, `yemen`, `zambia`, `zimbabwe` (optional) (default to "")
	importance := "importance_example" // string | Level of importance   `1`: low   `2`: medium   `3`: high (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested ts based on the date parameter. Unix timestamp format in milliseconds. (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested ts based on the date parameter. Unix timestamp format in milliseconds. The default is the timestamp of the request moment. (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicEconomicCalendarV5(context.Background()).Region(region).Importance(importance).Before(before).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicEconomicCalendarV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicEconomicCalendarV5`: GetPublicEconomicCalendarV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicEconomicCalendarV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicEconomicCalendarV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | Country, region or entity   &#x60;afghanistan&#x60;, &#x60;albania&#x60;, &#x60;algeria&#x60;, &#x60;andorra&#x60;, &#x60;angola&#x60;, &#x60;antigua_and_barbuda&#x60;, &#x60;argentina&#x60;, &#x60;armenia&#x60;, &#x60;aruba&#x60;, &#x60;australia&#x60;, &#x60;austria&#x60;, &#x60;azerbaijan&#x60;, &#x60;bahamas&#x60;, &#x60;bahrain&#x60;, &#x60;bangladesh&#x60;, &#x60;barbados&#x60;, &#x60;belarus&#x60;, &#x60;belgium&#x60;, &#x60;belize&#x60;, &#x60;benin&#x60;, &#x60;bermuda&#x60;, &#x60;bhutan&#x60;, &#x60;bolivia&#x60;, &#x60;bosnia_and_herzegovina&#x60;, &#x60;botswana&#x60;, &#x60;brazil&#x60;, &#x60;brunei&#x60;, &#x60;bulgaria&#x60;, &#x60;burkina_faso&#x60;, &#x60;burundi&#x60;, &#x60;cambodia&#x60;, &#x60;cameroon&#x60;, &#x60;canada&#x60;, &#x60;cape_verde&#x60;, &#x60;cayman_islands&#x60;, &#x60;central_african_republic&#x60;, &#x60;chad&#x60;, &#x60;chile&#x60;, &#x60;china&#x60;, &#x60;colombia&#x60;, &#x60;comoros&#x60;, &#x60;congo&#x60;, &#x60;costa_rica&#x60;, &#x60;croatia&#x60;, &#x60;cuba&#x60;, &#x60;cyprus&#x60;, &#x60;czech_republic&#x60;, &#x60;denmark&#x60;, &#x60;djibouti&#x60;, &#x60;dominica&#x60;, &#x60;dominican_republic&#x60;, &#x60;east_timor&#x60;, &#x60;ecuador&#x60;, &#x60;egypt&#x60;, &#x60;el_salvador&#x60;, &#x60;equatorial_guinea&#x60;, &#x60;eritrea&#x60;, &#x60;estonia&#x60;, &#x60;ethiopia&#x60;, &#x60;euro_area&#x60;, &#x60;european_union&#x60;, &#x60;faroe_islands&#x60;, &#x60;fiji&#x60;, &#x60;finland&#x60;, &#x60;france&#x60;, &#x60;g20&#x60;, &#x60;g7&#x60;, &#x60;gabon&#x60;, &#x60;gambia&#x60;, &#x60;georgia&#x60;, &#x60;germany&#x60;, &#x60;ghana&#x60;, &#x60;greece&#x60;, &#x60;greenland&#x60;, &#x60;grenada&#x60;, &#x60;guatemala&#x60;, &#x60;guinea&#x60;, &#x60;guinea_bissau&#x60;, &#x60;guyana&#x60;, &#x60;hungary&#x60;, &#x60;haiti&#x60;, &#x60;honduras&#x60;, &#x60;hong_kong&#x60;, &#x60;hungary&#x60;, &#x60;imf&#x60;, &#x60;indonesia&#x60;, &#x60;iceland&#x60;, &#x60;india&#x60;, &#x60;indonesia&#x60;, &#x60;iran&#x60;, &#x60;iraq&#x60;, &#x60;ireland&#x60;, &#x60;isle_of_man&#x60;, &#x60;israel&#x60;, &#x60;italy&#x60;, &#x60;ivory_coast&#x60;, &#x60;jamaica&#x60;, &#x60;japan&#x60;, &#x60;jordan&#x60;, &#x60;kazakhstan&#x60;, &#x60;kenya&#x60;, &#x60;kiribati&#x60;, &#x60;kosovo&#x60;, &#x60;kuwait&#x60;, &#x60;kyrgyzstan&#x60;, &#x60;laos&#x60;, &#x60;latvia&#x60;, &#x60;lebanon&#x60;, &#x60;lesotho&#x60;, &#x60;liberia&#x60;, &#x60;libya&#x60;, &#x60;liechtenstein&#x60;, &#x60;lithuania&#x60;, &#x60;luxembourg&#x60;, &#x60;macau&#x60;, &#x60;macedonia&#x60;, &#x60;madagascar&#x60;, &#x60;malawi&#x60;, &#x60;malaysia&#x60;, &#x60;maldives&#x60;, &#x60;mali&#x60;, &#x60;malta&#x60;, &#x60;mauritania&#x60;, &#x60;mauritius&#x60;, &#x60;mexico&#x60;, &#x60;micronesia&#x60;, &#x60;moldova&#x60;, &#x60;monaco&#x60;, &#x60;mongolia&#x60;, &#x60;montenegro&#x60;, &#x60;morocco&#x60;, &#x60;mozambique&#x60;, &#x60;myanmar&#x60;, &#x60;namibia&#x60;, &#x60;nepal&#x60;, &#x60;netherlands&#x60;, &#x60;new_caledonia&#x60;, &#x60;new_zealand&#x60;, &#x60;nicaragua&#x60;, &#x60;niger&#x60;, &#x60;nigeria&#x60;, &#x60;north_korea&#x60;, &#x60;northern_mariana_islands&#x60;, &#x60;norway&#x60;, &#x60;opec&#x60;, &#x60;oman&#x60;, &#x60;pakistan&#x60;, &#x60;palau&#x60;, &#x60;palestine&#x60;, &#x60;panama&#x60;, &#x60;papua_new_guinea&#x60;, &#x60;paraguay&#x60;, &#x60;peru&#x60;, &#x60;philippines&#x60;, &#x60;poland&#x60;, &#x60;portugal&#x60;, &#x60;puerto_rico&#x60;, &#x60;qatar&#x60;, &#x60;russia&#x60;, &#x60;republic_of_the_congo&#x60;, &#x60;romania&#x60;, &#x60;russia&#x60;, &#x60;rwanda&#x60;, &#x60;slovakia&#x60;, &#x60;samoa&#x60;, &#x60;san_marino&#x60;, &#x60;sao_tome_and_principe&#x60;, &#x60;saudi_arabia&#x60;, &#x60;senegal&#x60;, &#x60;serbia&#x60;, &#x60;seychelles&#x60;, &#x60;sierra_leone&#x60;, &#x60;singapore&#x60;, &#x60;slovakia&#x60;, &#x60;slovenia&#x60;, &#x60;solomon_islands&#x60;, &#x60;somalia&#x60;, &#x60;south_africa&#x60;, &#x60;south_korea&#x60;, &#x60;south_sudan&#x60;, &#x60;spain&#x60;, &#x60;sri_lanka&#x60;, &#x60;st_kitts_and_nevis&#x60;, &#x60;st_lucia&#x60;, &#x60;sudan&#x60;, &#x60;suriname&#x60;, &#x60;swaziland&#x60;, &#x60;sweden&#x60;, &#x60;switzerland&#x60;, &#x60;syria&#x60;, &#x60;taiwan&#x60;, &#x60;tajikistan&#x60;, &#x60;tanzania&#x60;, &#x60;thailand&#x60;, &#x60;togo&#x60;, &#x60;tonga&#x60;, &#x60;trinidad_and_tobago&#x60;, &#x60;tunisia&#x60;, &#x60;turkey&#x60;, &#x60;turkmenistan&#x60;, &#x60;uganda&#x60;, &#x60;ukraine&#x60;, &#x60;united_arab_emirates&#x60;, &#x60;united_kingdom&#x60;, &#x60;united_states&#x60;, &#x60;uruguay&#x60;, &#x60;uzbekistan&#x60;, &#x60;vanuatu&#x60;, &#x60;venezuela&#x60;, &#x60;vietnam&#x60;, &#x60;world&#x60;, &#x60;yemen&#x60;, &#x60;zambia&#x60;, &#x60;zimbabwe&#x60; | [default to &quot;&quot;]
 **importance** | **string** | Level of importance   &#x60;1&#x60;: low   &#x60;2&#x60;: medium   &#x60;3&#x60;: high | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested ts based on the date parameter. Unix timestamp format in milliseconds. | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested ts based on the date parameter. Unix timestamp format in milliseconds. The default is the timestamp of the request moment. | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetPublicEconomicCalendarV5Resp**](GetPublicEconomicCalendarV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicEstimatedPriceV5

> GetPublicEstimatedPriceV5Resp GetPublicEstimatedPriceV5(ctx).InstId(instId).Execute()

Retrieve the estimated delivery price which will only have a return value one hour before the delivery/exercise.  



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
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USD-200214`   only applicable to `FUTURES`/`OPTION` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicEstimatedPriceV5(context.Background()).InstId(instId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicEstimatedPriceV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicEstimatedPriceV5`: GetPublicEstimatedPriceV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicEstimatedPriceV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicEstimatedPriceV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USD-200214&#x60;   only applicable to &#x60;FUTURES&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]

### Return type

[**GetPublicEstimatedPriceV5Resp**](GetPublicEstimatedPriceV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicEstimatedSettlementInfoV5

> GetPublicEstimatedSettlementInfoV5Resp GetPublicEstimatedSettlementInfoV5(ctx).InstId(instId).Execute()

Retrieve the estimated settlement price which will only have a return value one hour before the settlement.  



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
	instId := "instId_example" // string | Instrument ID, e.g. `XRP-USDT-250307`   only applicable to `FUTURES` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicEstimatedSettlementInfoV5(context.Background()).InstId(instId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicEstimatedSettlementInfoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicEstimatedSettlementInfoV5`: GetPublicEstimatedSettlementInfoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicEstimatedSettlementInfoV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicEstimatedSettlementInfoV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;XRP-USDT-250307&#x60;   only applicable to &#x60;FUTURES&#x60; | [default to &quot;&quot;]

### Return type

[**GetPublicEstimatedSettlementInfoV5Resp**](GetPublicEstimatedSettlementInfoV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicFundingRateHistoryV5

> GetPublicFundingRateHistoryV5Resp GetPublicFundingRateHistoryV5(ctx).InstId(instId).Before(before).After(after).Limit(limit).Execute()

Retrieve funding rate history. This endpoint can retrieve data from the last 3 months.  



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
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USD-SWAP`   only applicable to `SWAP` (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `fundingTime` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `fundingTime` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicFundingRateHistoryV5(context.Background()).InstId(instId).Before(before).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicFundingRateHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicFundingRateHistoryV5`: GetPublicFundingRateHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicFundingRateHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicFundingRateHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USD-SWAP&#x60;   only applicable to &#x60;SWAP&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;fundingTime&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;fundingTime&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetPublicFundingRateHistoryV5Resp**](GetPublicFundingRateHistoryV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicFundingRateV5

> GetPublicFundingRateV5Resp GetPublicFundingRateV5(ctx).InstId(instId).Execute()

Retrieve funding rate.  



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
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USD-SWAP`   only applicable to `SWAP` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicFundingRateV5(context.Background()).InstId(instId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicFundingRateV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicFundingRateV5`: GetPublicFundingRateV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicFundingRateV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicFundingRateV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USD-SWAP&#x60;   only applicable to &#x60;SWAP&#x60; | [default to &quot;&quot;]

### Return type

[**GetPublicFundingRateV5Resp**](GetPublicFundingRateV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicInstrumentTickBandsV5

> GetPublicInstrumentTickBandsV5Resp GetPublicInstrumentTickBandsV5(ctx).InstType(instType).InstFamily(instFamily).Execute()

Get option tick bands information  



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
	instType := "instType_example" // string | Instrument type  `OPTION` (default to "")
	instFamily := "instFamily_example" // string | Instrument family  Only applicable to OPTION (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicInstrumentTickBandsV5(context.Background()).InstType(instType).InstFamily(instFamily).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicInstrumentTickBandsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicInstrumentTickBandsV5`: GetPublicInstrumentTickBandsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicInstrumentTickBandsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicInstrumentTickBandsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family  Only applicable to OPTION | [default to &quot;&quot;]

### Return type

[**GetPublicInstrumentTickBandsV5Resp**](GetPublicInstrumentTickBandsV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicInstrumentsV5

> GetPublicInstrumentsV5Resp GetPublicInstrumentsV5(ctx).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).Execute()

Retrieve a list of instruments with open contracts for OKX. Retrieve available instruments info of current account, please refer to .  



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
	instType := "instType_example" // string | Instrument type  `SPOT`: Spot  `MARGIN`: Margin  `SWAP`: Perpetual Futures  `FUTURES`: Expiry Futures  `OPTION`: Option (default to "")
	uly := "uly_example" // string | Underlying   Only applicable to `FUTURES`/`SWAP`/`OPTION`.If instType is `OPTION`, either `uly` or `instFamily` is required. (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family  Only applicable to `FUTURES`/`SWAP`/`OPTION`. If instType is `OPTION`, either `uly` or `instFamily` is required. (optional) (default to "")
	instId := "instId_example" // string | Instrument ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicInstrumentsV5(context.Background()).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicInstrumentsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicInstrumentsV5`: GetPublicInstrumentsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicInstrumentsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicInstrumentsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;: Spot  &#x60;MARGIN&#x60;: Margin  &#x60;SWAP&#x60;: Perpetual Futures  &#x60;FUTURES&#x60;: Expiry Futures  &#x60;OPTION&#x60;: Option | [default to &quot;&quot;]
 **uly** | **string** | Underlying   Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;.If instType is &#x60;OPTION&#x60;, either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family  Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;. If instType is &#x60;OPTION&#x60;, either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID | [default to &quot;&quot;]

### Return type

[**GetPublicInstrumentsV5Resp**](GetPublicInstrumentsV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicInsuranceFundV5

> GetPublicInsuranceFundV5Resp GetPublicInsuranceFundV5(ctx).InstType(instType).Type_(type_).Uly(uly).InstFamily(instFamily).Ccy(ccy).Before(before).After(after).Limit(limit).Execute()

Get insurance fund balance information  



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
	instType := "instType_example" // string | Instrument type  `MARGIN`  `SWAP`  `FUTURES`  `OPTION` (default to "")
	type_ := "type__example" // string | Type  `regular_update`   `liquidation_balance_deposit`  `bankruptcy_loss`  `platform_revenue`   `adl`: ADL historical data   The default is `all type` (optional) (default to "")
	uly := "uly_example" // string | Underlying  Required for `FUTURES`/`SWAP`/`OPTION`  Either `uly` or `instFamily` is required. If both are passed, `instFamily` will be used. (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family  Required for `FUTURES`/`SWAP`/`OPTION`  Either `uly` or `instFamily` is required. If both are passed, `instFamily` will be used. (optional) (default to "")
	ccy := "ccy_example" // string | Currency, only applicable to `MARGIN` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ts` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ts` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicInsuranceFundV5(context.Background()).InstType(instType).Type_(type_).Uly(uly).InstFamily(instFamily).Ccy(ccy).Before(before).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicInsuranceFundV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicInsuranceFundV5`: GetPublicInsuranceFundV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicInsuranceFundV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicInsuranceFundV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **type_** | **string** | Type  &#x60;regular_update&#x60;   &#x60;liquidation_balance_deposit&#x60;  &#x60;bankruptcy_loss&#x60;  &#x60;platform_revenue&#x60;   &#x60;adl&#x60;: ADL historical data   The default is &#x60;all type&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Underlying  Required for &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;  Either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. If both are passed, &#x60;instFamily&#x60; will be used. | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family  Required for &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;  Either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. If both are passed, &#x60;instFamily&#x60; will be used. | [default to &quot;&quot;]
 **ccy** | **string** | Currency, only applicable to &#x60;MARGIN&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetPublicInsuranceFundV5Resp**](GetPublicInsuranceFundV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicInterestRateLoanQuotaV5

> GetPublicInterestRateLoanQuotaV5Resp GetPublicInterestRateLoanQuotaV5(ctx).Execute()

Retrieve interest rate  



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
	resp, r, err := apiClient.PublicDataAPI.GetPublicInterestRateLoanQuotaV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicInterestRateLoanQuotaV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicInterestRateLoanQuotaV5`: GetPublicInterestRateLoanQuotaV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicInterestRateLoanQuotaV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicInterestRateLoanQuotaV5Request struct via the builder pattern


### Return type

[**GetPublicInterestRateLoanQuotaV5Resp**](GetPublicInterestRateLoanQuotaV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicMarkPriceV5

> GetPublicMarkPriceV5Resp GetPublicMarkPriceV5(ctx).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).Execute()

Retrieve mark price.  We set the mark price based on the SPOT index and at a reasonable basis to prevent individual users from manipulating the market and causing the contract price to fluctuate.  



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
	instType := "instType_example" // string | Instrument type  `MARGIN`  `SWAP`  `FUTURES`  `OPTION` (default to "")
	uly := "uly_example" // string | Underlying  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USD-SWAP` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicMarkPriceV5(context.Background()).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicMarkPriceV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicMarkPriceV5`: GetPublicMarkPriceV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicMarkPriceV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicMarkPriceV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Underlying  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USD-SWAP&#x60; | [default to &quot;&quot;]

### Return type

[**GetPublicMarkPriceV5Resp**](GetPublicMarkPriceV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicOpenInterestV5

> GetPublicOpenInterestV5Resp GetPublicOpenInterestV5(ctx).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).Execute()

Retrieve the total open interest for contracts on OKX.  



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
	instType := "instType_example" // string | Instrument type  `SWAP`  `FUTURES`  `OPTION` (default to "")
	uly := "uly_example" // string | Underlying  Applicable to `FUTURES`/`SWAP`/`OPTION`.   If instType is `OPTION`, either `uly` or `instFamily` is required. (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family  Applicable to `FUTURES`/`SWAP`/`OPTION`  If instType is `OPTION`, either `uly` or `instFamily` is required. (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT-SWAP`  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicOpenInterestV5(context.Background()).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicOpenInterestV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicOpenInterestV5`: GetPublicOpenInterestV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicOpenInterestV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicOpenInterestV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Underlying  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;.   If instType is &#x60;OPTION&#x60;, either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;  If instType is &#x60;OPTION&#x60;, either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60;  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]

### Return type

[**GetPublicOpenInterestV5Resp**](GetPublicOpenInterestV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicOptSummaryV5

> GetPublicOptSummaryV5Resp GetPublicOptSummaryV5(ctx).Uly(uly).InstFamily(instFamily).ExpTime(expTime).Execute()

Retrieve option market data.  



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
	uly := "uly_example" // string | Underlying, only applicable to `OPTION`  Either `uly` or `instFamily` is required. If both are passed, `instFamily` will be used. (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family, only applicable to `OPTION`  Either `uly` or `instFamily` is required. If both are passed, `instFamily` will be used. (optional) (default to "")
	expTime := "expTime_example" // string | Contract expiry date, the format is \"YYMMDD\", e.g. \"200527\" (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicOptSummaryV5(context.Background()).Uly(uly).InstFamily(instFamily).ExpTime(expTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicOptSummaryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicOptSummaryV5`: GetPublicOptSummaryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicOptSummaryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicOptSummaryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uly** | **string** | Underlying, only applicable to &#x60;OPTION&#x60;  Either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. If both are passed, &#x60;instFamily&#x60; will be used. | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family, only applicable to &#x60;OPTION&#x60;  Either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. If both are passed, &#x60;instFamily&#x60; will be used. | [default to &quot;&quot;]
 **expTime** | **string** | Contract expiry date, the format is \&quot;YYMMDD\&quot;, e.g. \&quot;200527\&quot; | [default to &quot;&quot;]

### Return type

[**GetPublicOptSummaryV5Resp**](GetPublicOptSummaryV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicPositionTiersV5

> GetPublicPositionTiersV5Resp GetPublicPositionTiersV5(ctx).InstType(instType).TdMode(tdMode).Uly(uly).InstFamily(instFamily).InstId(instId).Ccy(ccy).Tier(tier).Execute()

Retrieve position tiers information, maximum leverage depends on your borrowings and margin ratio.  



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
	instType := "instType_example" // string | Instrument type  `MARGIN`  `SWAP`  `FUTURES`  `OPTION` (default to "")
	tdMode := "tdMode_example" // string | Trade mode  Margin mode `cross` `isolated` (default to "")
	uly := "uly_example" // string | Single underlying or multiple underlyings (no more than 3) separated with comma.  If instType is `SWAP/FUTURES/OPTION`, either `uly` or `instFamily` is required.  If both are passed, `instFamily` will be used. (optional) (default to "")
	instFamily := "instFamily_example" // string | Single instrument familiy or multiple instrument families (no more than 5) separated with comma.  If instType is `SWAP/FUTURES/OPTION`, either `uly` or `instFamily` is required.  If both are passed, `instFamily` will be used. (optional) (default to "")
	instId := "instId_example" // string | Single instrument or multiple instruments (no more than 5) separated with comma.  Either instId or ccy is required, if both are passed, instId will be used, ignore when instType is one of `SWAP`,`FUTURES`,`OPTION` (optional) (default to "")
	ccy := "ccy_example" // string | Margin currency  Only applicable to cross MARGIN. It will return borrowing amount for `Multi-currency margin` and `Portfolio margin` when `ccy` takes effect. (optional) (default to "")
	tier := "tier_example" // string | Tiers (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicPositionTiersV5(context.Background()).InstType(instType).TdMode(tdMode).Uly(uly).InstFamily(instFamily).InstId(instId).Ccy(ccy).Tier(tier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicPositionTiersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicPositionTiersV5`: GetPublicPositionTiersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicPositionTiersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicPositionTiersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **tdMode** | **string** | Trade mode  Margin mode &#x60;cross&#x60; &#x60;isolated&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Single underlying or multiple underlyings (no more than 3) separated with comma.  If instType is &#x60;SWAP/FUTURES/OPTION&#x60;, either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required.  If both are passed, &#x60;instFamily&#x60; will be used. | [default to &quot;&quot;]
 **instFamily** | **string** | Single instrument familiy or multiple instrument families (no more than 5) separated with comma.  If instType is &#x60;SWAP/FUTURES/OPTION&#x60;, either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required.  If both are passed, &#x60;instFamily&#x60; will be used. | [default to &quot;&quot;]
 **instId** | **string** | Single instrument or multiple instruments (no more than 5) separated with comma.  Either instId or ccy is required, if both are passed, instId will be used, ignore when instType is one of &#x60;SWAP&#x60;,&#x60;FUTURES&#x60;,&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **ccy** | **string** | Margin currency  Only applicable to cross MARGIN. It will return borrowing amount for &#x60;Multi-currency margin&#x60; and &#x60;Portfolio margin&#x60; when &#x60;ccy&#x60; takes effect. | [default to &quot;&quot;]
 **tier** | **string** | Tiers | [default to &quot;&quot;]

### Return type

[**GetPublicPositionTiersV5Resp**](GetPublicPositionTiersV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicPremiumHistoryV5

> GetPublicPremiumHistoryV5Resp GetPublicPremiumHistoryV5(ctx).InstId(instId).After(after).Before(before).Limit(limit).Execute()

It will return premium data in the past 6 months.  



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
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT-SWAP`  Applicable to `SWAP` (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested ts(not included) (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested ts(not included) (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicPremiumHistoryV5(context.Background()).InstId(instId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicPremiumHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicPremiumHistoryV5`: GetPublicPremiumHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicPremiumHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicPremiumHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60;  Applicable to &#x60;SWAP&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested ts(not included) | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested ts(not included) | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetPublicPremiumHistoryV5Resp**](GetPublicPremiumHistoryV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicPriceLimitV5

> GetPublicPriceLimitV5Resp GetPublicPriceLimitV5(ctx).InstId(instId).Execute()

Retrieve the highest buy limit and lowest sell limit of the instrument.  



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
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT-SWAP` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicPriceLimitV5(context.Background()).InstId(instId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicPriceLimitV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicPriceLimitV5`: GetPublicPriceLimitV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicPriceLimitV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicPriceLimitV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [default to &quot;&quot;]

### Return type

[**GetPublicPriceLimitV5Resp**](GetPublicPriceLimitV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicSettlementHistoryV5

> GetPublicSettlementHistoryV5Resp GetPublicSettlementHistoryV5(ctx).InstFamily(instFamily).After(after).Before(before).Limit(limit).Execute()

Retrieve settlement records of futures in the last 3 months.  



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
	instFamily := "instFamily_example" // string | Instrument family (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than (not include) the requested `ts` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than (not include) the requested `ts` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicSettlementHistoryV5(context.Background()).InstFamily(instFamily).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicSettlementHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicSettlementHistoryV5`: GetPublicSettlementHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicSettlementHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicSettlementHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instFamily** | **string** | Instrument family | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than (not include) the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than (not include) the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetPublicSettlementHistoryV5Resp**](GetPublicSettlementHistoryV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicTimeV5

> GetPublicTimeV5Resp GetPublicTimeV5(ctx).Execute()

Retrieve API server time.  



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
	resp, r, err := apiClient.PublicDataAPI.GetPublicTimeV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicTimeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicTimeV5`: GetPublicTimeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicTimeV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicTimeV5Request struct via the builder pattern


### Return type

[**GetPublicTimeV5Resp**](GetPublicTimeV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicUnderlyingV5

> GetPublicUnderlyingV5Resp GetPublicUnderlyingV5(ctx).InstType(instType).Execute()





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
	instType := "instType_example" // string | Instrument type  `SWAP`  `FUTURES`  `OPTION` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicDataAPI.GetPublicUnderlyingV5(context.Background()).InstType(instType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicDataAPI.GetPublicUnderlyingV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicUnderlyingV5`: GetPublicUnderlyingV5Resp
	fmt.Fprintf(os.Stdout, "Response from `PublicDataAPI.GetPublicUnderlyingV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicUnderlyingV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]

### Return type

[**GetPublicUnderlyingV5Resp**](GetPublicUnderlyingV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

