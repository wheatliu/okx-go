# \TradingStatisticsAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRubikStatContractsLongShortAccountRatioContractTopTraderV5**](TradingStatisticsAPI.md#GetRubikStatContractsLongShortAccountRatioContractTopTraderV5) | **Get** /api/v5/rubik/stat/contracts/long-short-account-ratio-contract-top-trader | Get top traders contract long/short ratio
[**GetRubikStatContractsLongShortAccountRatioContractV5**](TradingStatisticsAPI.md#GetRubikStatContractsLongShortAccountRatioContractV5) | **Get** /api/v5/rubik/stat/contracts/long-short-account-ratio-contract | Get contract long/short ratio
[**GetRubikStatContractsLongShortAccountRatioV5**](TradingStatisticsAPI.md#GetRubikStatContractsLongShortAccountRatioV5) | **Get** /api/v5/rubik/stat/contracts/long-short-account-ratio | Get long/short ratio
[**GetRubikStatContractsLongShortPositionRatioContractTopTraderV5**](TradingStatisticsAPI.md#GetRubikStatContractsLongShortPositionRatioContractTopTraderV5) | **Get** /api/v5/rubik/stat/contracts/long-short-position-ratio-contract-top-trader | Get top traders contract long/short ratio (by position)
[**GetRubikStatContractsOpenInterestHistoryV5**](TradingStatisticsAPI.md#GetRubikStatContractsOpenInterestHistoryV5) | **Get** /api/v5/rubik/stat/contracts/open-interest-history | Get contract open interest history
[**GetRubikStatContractsOpenInterestVolumeV5**](TradingStatisticsAPI.md#GetRubikStatContractsOpenInterestVolumeV5) | **Get** /api/v5/rubik/stat/contracts/open-interest-volume | Get contracts open interest and volume
[**GetRubikStatMarginLoanRatioV5**](TradingStatisticsAPI.md#GetRubikStatMarginLoanRatioV5) | **Get** /api/v5/rubik/stat/margin/loan-ratio | Get margin long/short ratio
[**GetRubikStatOptionOpenInterestVolumeExpiryV5**](TradingStatisticsAPI.md#GetRubikStatOptionOpenInterestVolumeExpiryV5) | **Get** /api/v5/rubik/stat/option/open-interest-volume-expiry | Get open interest and volume (expiry)
[**GetRubikStatOptionOpenInterestVolumeRatioV5**](TradingStatisticsAPI.md#GetRubikStatOptionOpenInterestVolumeRatioV5) | **Get** /api/v5/rubik/stat/option/open-interest-volume-ratio | Get put/call ratio
[**GetRubikStatOptionOpenInterestVolumeStrikeV5**](TradingStatisticsAPI.md#GetRubikStatOptionOpenInterestVolumeStrikeV5) | **Get** /api/v5/rubik/stat/option/open-interest-volume-strike | Get open interest and volume (strike)
[**GetRubikStatOptionOpenInterestVolumeV5**](TradingStatisticsAPI.md#GetRubikStatOptionOpenInterestVolumeV5) | **Get** /api/v5/rubik/stat/option/open-interest-volume | Get options open interest and volume
[**GetRubikStatOptionTakerBlockVolumeV5**](TradingStatisticsAPI.md#GetRubikStatOptionTakerBlockVolumeV5) | **Get** /api/v5/rubik/stat/option/taker-block-volume | Get taker flow
[**GetRubikStatTakerVolumeContractV5**](TradingStatisticsAPI.md#GetRubikStatTakerVolumeContractV5) | **Get** /api/v5/rubik/stat/taker-volume-contract | Get contract taker volume
[**GetRubikStatTakerVolumeV5**](TradingStatisticsAPI.md#GetRubikStatTakerVolumeV5) | **Get** /api/v5/rubik/stat/taker-volume | Get taker volume
[**GetRubikStatTradingDataSupportCoinV5**](TradingStatisticsAPI.md#GetRubikStatTradingDataSupportCoinV5) | **Get** /api/v5/rubik/stat/trading-data/support-coin | Get support coin



## GetRubikStatContractsLongShortAccountRatioContractTopTraderV5

> GetRubikStatContractsLongShortAccountRatioContractTopTraderV5Resp GetRubikStatContractsLongShortAccountRatioContractTopTraderV5(ctx).InstId(instId).Period(period).End(end).Begin(begin).Limit(limit).Execute()

Get top traders contract long/short ratio



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
	instId := "instId_example" // string | Instrument ID, eg: BTC-USDT-SWAP   Only applicable to `FUTURES`, `SWAP` (default to "")
	period := "period_example" // string | Bar size, the default is `5m`, e.g. [`5m/15m/30m/1H/2H/4H`]   Hong Kong time opening price k-line: [`6H/12H/1D/2D/3D/5D/1W/1M/3M`]   UTC time opening price k-line: [`6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/5Dutc/1Wutc/1Mutc/3Mutc`] (optional) (default to "")
	end := "end_example" // string | return records earlier than the requested `ts` (optional) (default to "")
	begin := "begin_example" // string | return records newer than the requested `ts` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatContractsLongShortAccountRatioContractTopTraderV5(context.Background()).InstId(instId).Period(period).End(end).Begin(begin).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatContractsLongShortAccountRatioContractTopTraderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatContractsLongShortAccountRatioContractTopTraderV5`: GetRubikStatContractsLongShortAccountRatioContractTopTraderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatContractsLongShortAccountRatioContractTopTraderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatContractsLongShortAccountRatioContractTopTraderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, eg: BTC-USDT-SWAP   Only applicable to &#x60;FUTURES&#x60;, &#x60;SWAP&#x60; | [default to &quot;&quot;]
 **period** | **string** | Bar size, the default is &#x60;5m&#x60;, e.g. [&#x60;5m/15m/30m/1H/2H/4H&#x60;]   Hong Kong time opening price k-line: [&#x60;6H/12H/1D/2D/3D/5D/1W/1M/3M&#x60;]   UTC time opening price k-line: [&#x60;6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/5Dutc/1Wutc/1Mutc/3Mutc&#x60;] | [default to &quot;&quot;]
 **end** | **string** | return records earlier than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **begin** | **string** | return records newer than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetRubikStatContractsLongShortAccountRatioContractTopTraderV5Resp**](GetRubikStatContractsLongShortAccountRatioContractTopTraderV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatContractsLongShortAccountRatioContractV5

> GetRubikStatContractsLongShortAccountRatioContractV5Resp GetRubikStatContractsLongShortAccountRatioContractV5(ctx).InstId(instId).Period(period).End(end).Begin(begin).Limit(limit).Execute()

Get contract long/short ratio



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
	instId := "instId_example" // string | Instrument ID, eg: BTC-USDT-SWAP   Only applicable to `FUTURES`, `SWAP` (default to "")
	period := "period_example" // string | Bar size, the default is `5m`, e.g. [`5m/15m/30m/1H/2H/4H`]   Hong Kong time opening price k-line:[`6H/12H/1D/2D/3D/5D/1W/1M/3M`]   UTC time opening price k-line: [`6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/5Dutc/1Wutc/1Mutc/3Mutc`] (optional) (default to "")
	end := "end_example" // string | return records earlier than the requested `ts` (optional) (default to "")
	begin := "begin_example" // string | return records newer than the requested `ts` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatContractsLongShortAccountRatioContractV5(context.Background()).InstId(instId).Period(period).End(end).Begin(begin).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatContractsLongShortAccountRatioContractV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatContractsLongShortAccountRatioContractV5`: GetRubikStatContractsLongShortAccountRatioContractV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatContractsLongShortAccountRatioContractV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatContractsLongShortAccountRatioContractV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, eg: BTC-USDT-SWAP   Only applicable to &#x60;FUTURES&#x60;, &#x60;SWAP&#x60; | [default to &quot;&quot;]
 **period** | **string** | Bar size, the default is &#x60;5m&#x60;, e.g. [&#x60;5m/15m/30m/1H/2H/4H&#x60;]   Hong Kong time opening price k-line:[&#x60;6H/12H/1D/2D/3D/5D/1W/1M/3M&#x60;]   UTC time opening price k-line: [&#x60;6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/5Dutc/1Wutc/1Mutc/3Mutc&#x60;] | [default to &quot;&quot;]
 **end** | **string** | return records earlier than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **begin** | **string** | return records newer than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetRubikStatContractsLongShortAccountRatioContractV5Resp**](GetRubikStatContractsLongShortAccountRatioContractV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatContractsLongShortAccountRatioV5

> GetRubikStatContractsLongShortAccountRatioV5Resp GetRubikStatContractsLongShortAccountRatioV5(ctx).Ccy(ccy).Begin(begin).End(end).Period(period).Execute()

Get long/short ratio



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
	ccy := "ccy_example" // string | Currency (default to "")
	begin := "begin_example" // string | Begin time, e.g. `1597026383085` (optional) (default to "")
	end := "end_example" // string | End time, e.g. `1597026383011` (optional) (default to "")
	period := "period_example" // string | Period, the default is `5m`, e.g. [`5m/1H/1D`]    `5m` granularity can only query data within two days at most  `1H` granularity can only query data within 30 days at most   `1D` granularity can only query data within 180 days at most (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatContractsLongShortAccountRatioV5(context.Background()).Ccy(ccy).Begin(begin).End(end).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatContractsLongShortAccountRatioV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatContractsLongShortAccountRatioV5`: GetRubikStatContractsLongShortAccountRatioV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatContractsLongShortAccountRatioV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatContractsLongShortAccountRatioV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency | [default to &quot;&quot;]
 **begin** | **string** | Begin time, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **end** | **string** | End time, e.g. &#x60;1597026383011&#x60; | [default to &quot;&quot;]
 **period** | **string** | Period, the default is &#x60;5m&#x60;, e.g. [&#x60;5m/1H/1D&#x60;]    &#x60;5m&#x60; granularity can only query data within two days at most  &#x60;1H&#x60; granularity can only query data within 30 days at most   &#x60;1D&#x60; granularity can only query data within 180 days at most | [default to &quot;&quot;]

### Return type

[**GetRubikStatContractsLongShortAccountRatioV5Resp**](GetRubikStatContractsLongShortAccountRatioV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatContractsLongShortPositionRatioContractTopTraderV5

> GetRubikStatContractsLongShortPositionRatioContractTopTraderV5Resp GetRubikStatContractsLongShortPositionRatioContractTopTraderV5(ctx).InstId(instId).Period(period).End(end).Begin(begin).Limit(limit).Execute()

Get top traders contract long/short ratio (by position)



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
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT-SWAP`   Only applicable to `FUTURES`/`SWAP` (default to "")
	period := "period_example" // string | Bar size, the default is `5m`, e.g. [`5m/15m/30m/1H/2H/4H`]   Hong Kong time opening price k-line: [`6H/12H/1D/2D/3D/5D/1W/1M/3M`]   UTC time opening price k-line: [`6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/5Dutc/1Wutc/1Mutc/3Mutc`] (optional) (default to "")
	end := "end_example" // string | return records earlier than the requested `ts` (optional) (default to "")
	begin := "begin_example" // string | return records newer than the requested `ts` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatContractsLongShortPositionRatioContractTopTraderV5(context.Background()).InstId(instId).Period(period).End(end).Begin(begin).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatContractsLongShortPositionRatioContractTopTraderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatContractsLongShortPositionRatioContractTopTraderV5`: GetRubikStatContractsLongShortPositionRatioContractTopTraderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatContractsLongShortPositionRatioContractTopTraderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatContractsLongShortPositionRatioContractTopTraderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60;   Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60; | [default to &quot;&quot;]
 **period** | **string** | Bar size, the default is &#x60;5m&#x60;, e.g. [&#x60;5m/15m/30m/1H/2H/4H&#x60;]   Hong Kong time opening price k-line: [&#x60;6H/12H/1D/2D/3D/5D/1W/1M/3M&#x60;]   UTC time opening price k-line: [&#x60;6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/5Dutc/1Wutc/1Mutc/3Mutc&#x60;] | [default to &quot;&quot;]
 **end** | **string** | return records earlier than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **begin** | **string** | return records newer than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetRubikStatContractsLongShortPositionRatioContractTopTraderV5Resp**](GetRubikStatContractsLongShortPositionRatioContractTopTraderV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatContractsOpenInterestHistoryV5

> GetRubikStatContractsOpenInterestHistoryV5Resp GetRubikStatContractsOpenInterestHistoryV5(ctx).InstId(instId).Period(period).End(end).Begin(begin).Limit(limit).Execute()

Get contract open interest history



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
	instId := "instId_example" // string | Instrument ID, eg: BTC-USDT-SWAP   Only applicable to `FUTURES`, `SWAP` (default to "")
	period := "period_example" // string | Bar size, the default is `5m`, e.g. [`5m/15m/30m/1H/2H/4H`]   Hong Kong time opening price k-line: [`6H/12H/1D/2D/3D/5D/1W/1M/3M`]   UTC time opening price k-line: [`6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/5Dutc/1Wutc/1Mutc/3Mutc`] (optional) (default to "")
	end := "end_example" // string | Pagination of data to return records earlier than the requested `ts` (optional) (default to "")
	begin := "begin_example" // string | return records newer than the requested `ts` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatContractsOpenInterestHistoryV5(context.Background()).InstId(instId).Period(period).End(end).Begin(begin).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatContractsOpenInterestHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatContractsOpenInterestHistoryV5`: GetRubikStatContractsOpenInterestHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatContractsOpenInterestHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatContractsOpenInterestHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, eg: BTC-USDT-SWAP   Only applicable to &#x60;FUTURES&#x60;, &#x60;SWAP&#x60; | [default to &quot;&quot;]
 **period** | **string** | Bar size, the default is &#x60;5m&#x60;, e.g. [&#x60;5m/15m/30m/1H/2H/4H&#x60;]   Hong Kong time opening price k-line: [&#x60;6H/12H/1D/2D/3D/5D/1W/1M/3M&#x60;]   UTC time opening price k-line: [&#x60;6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/5Dutc/1Wutc/1Mutc/3Mutc&#x60;] | [default to &quot;&quot;]
 **end** | **string** | Pagination of data to return records earlier than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **begin** | **string** | return records newer than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetRubikStatContractsOpenInterestHistoryV5Resp**](GetRubikStatContractsOpenInterestHistoryV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatContractsOpenInterestVolumeV5

> GetRubikStatContractsOpenInterestVolumeV5Resp GetRubikStatContractsOpenInterestVolumeV5(ctx).Ccy(ccy).Begin(begin).End(end).Period(period).Execute()

Get contracts open interest and volume



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
	ccy := "ccy_example" // string | Currency (default to "")
	begin := "begin_example" // string | Begin time, e.g. `1597026383085` (optional) (default to "")
	end := "end_example" // string | End time, e.g. `1597026383011` (optional) (default to "")
	period := "period_example" // string | Period, the default is `5m`, e.g. [`5m/1H/1D`]    `5m` granularity can only query data within two days at most  `1H` granularity can only query data within 30 days at most   `1D` granularity can only query data within 180 days at most (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatContractsOpenInterestVolumeV5(context.Background()).Ccy(ccy).Begin(begin).End(end).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatContractsOpenInterestVolumeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatContractsOpenInterestVolumeV5`: GetRubikStatContractsOpenInterestVolumeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatContractsOpenInterestVolumeV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatContractsOpenInterestVolumeV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency | [default to &quot;&quot;]
 **begin** | **string** | Begin time, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **end** | **string** | End time, e.g. &#x60;1597026383011&#x60; | [default to &quot;&quot;]
 **period** | **string** | Period, the default is &#x60;5m&#x60;, e.g. [&#x60;5m/1H/1D&#x60;]    &#x60;5m&#x60; granularity can only query data within two days at most  &#x60;1H&#x60; granularity can only query data within 30 days at most   &#x60;1D&#x60; granularity can only query data within 180 days at most | [default to &quot;&quot;]

### Return type

[**GetRubikStatContractsOpenInterestVolumeV5Resp**](GetRubikStatContractsOpenInterestVolumeV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatMarginLoanRatioV5

> GetRubikStatMarginLoanRatioV5Resp GetRubikStatMarginLoanRatioV5(ctx).Ccy(ccy).Begin(begin).End(end).Period(period).Execute()

Get margin long/short ratio



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
	ccy := "ccy_example" // string | Currency (default to "")
	begin := "begin_example" // string | Begin time, e.g. `1597026383085` (optional) (default to "")
	end := "end_example" // string | End time, e.g. `1597026383085` (optional) (default to "")
	period := "period_example" // string | Period  `m`: Minute, `H`: Hour, `D`: Day  the default is `5m`, e.g. [`5m`/`1H`/`1D`]    `5m` granularity can only query data within two days at most  `1H` granularity can only query data within 30 days at most  `1D` granularity can only query data within 180 days at most (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatMarginLoanRatioV5(context.Background()).Ccy(ccy).Begin(begin).End(end).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatMarginLoanRatioV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatMarginLoanRatioV5`: GetRubikStatMarginLoanRatioV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatMarginLoanRatioV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatMarginLoanRatioV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency | [default to &quot;&quot;]
 **begin** | **string** | Begin time, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **end** | **string** | End time, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **period** | **string** | Period  &#x60;m&#x60;: Minute, &#x60;H&#x60;: Hour, &#x60;D&#x60;: Day  the default is &#x60;5m&#x60;, e.g. [&#x60;5m&#x60;/&#x60;1H&#x60;/&#x60;1D&#x60;]    &#x60;5m&#x60; granularity can only query data within two days at most  &#x60;1H&#x60; granularity can only query data within 30 days at most  &#x60;1D&#x60; granularity can only query data within 180 days at most | [default to &quot;&quot;]

### Return type

[**GetRubikStatMarginLoanRatioV5Resp**](GetRubikStatMarginLoanRatioV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatOptionOpenInterestVolumeExpiryV5

> GetRubikStatOptionOpenInterestVolumeExpiryV5Resp GetRubikStatOptionOpenInterestVolumeExpiryV5(ctx).Ccy(ccy).Period(period).Execute()

Get open interest and volume (expiry)



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
	ccy := "ccy_example" // string | Currency (default to "")
	period := "period_example" // string | Period, the default is `8H`. e.g. [`8H/1D`]    Each granularity can provide only one latest piece of data (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatOptionOpenInterestVolumeExpiryV5(context.Background()).Ccy(ccy).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatOptionOpenInterestVolumeExpiryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatOptionOpenInterestVolumeExpiryV5`: GetRubikStatOptionOpenInterestVolumeExpiryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatOptionOpenInterestVolumeExpiryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatOptionOpenInterestVolumeExpiryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency | [default to &quot;&quot;]
 **period** | **string** | Period, the default is &#x60;8H&#x60;. e.g. [&#x60;8H/1D&#x60;]    Each granularity can provide only one latest piece of data | [default to &quot;&quot;]

### Return type

[**GetRubikStatOptionOpenInterestVolumeExpiryV5Resp**](GetRubikStatOptionOpenInterestVolumeExpiryV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatOptionOpenInterestVolumeRatioV5

> GetRubikStatOptionOpenInterestVolumeRatioV5Resp GetRubikStatOptionOpenInterestVolumeRatioV5(ctx).Ccy(ccy).Period(period).Execute()

Get put/call ratio



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
	ccy := "ccy_example" // string | Currency (default to "")
	period := "period_example" // string | Period, the default is `8H`. e.g. [`8H/1D`]    Each granularity can only query 72 pieces of data at the earliest (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatOptionOpenInterestVolumeRatioV5(context.Background()).Ccy(ccy).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatOptionOpenInterestVolumeRatioV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatOptionOpenInterestVolumeRatioV5`: GetRubikStatOptionOpenInterestVolumeRatioV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatOptionOpenInterestVolumeRatioV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatOptionOpenInterestVolumeRatioV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency | [default to &quot;&quot;]
 **period** | **string** | Period, the default is &#x60;8H&#x60;. e.g. [&#x60;8H/1D&#x60;]    Each granularity can only query 72 pieces of data at the earliest | [default to &quot;&quot;]

### Return type

[**GetRubikStatOptionOpenInterestVolumeRatioV5Resp**](GetRubikStatOptionOpenInterestVolumeRatioV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatOptionOpenInterestVolumeStrikeV5

> GetRubikStatOptionOpenInterestVolumeStrikeV5Resp GetRubikStatOptionOpenInterestVolumeStrikeV5(ctx).Ccy(ccy).ExpTime(expTime).Period(period).Execute()

Get open interest and volume (strike)



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
	ccy := "ccy_example" // string | Currency (default to "")
	expTime := "expTime_example" // string | Contract expiry date, the format is `YYYYMMdd`, e.g. `20210623` (default to "")
	period := "period_example" // string | Period, the default is `8H`. e.g. [`8H/1D`]    Each granularity can provide only one latest piece of data (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatOptionOpenInterestVolumeStrikeV5(context.Background()).Ccy(ccy).ExpTime(expTime).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatOptionOpenInterestVolumeStrikeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatOptionOpenInterestVolumeStrikeV5`: GetRubikStatOptionOpenInterestVolumeStrikeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatOptionOpenInterestVolumeStrikeV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatOptionOpenInterestVolumeStrikeV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency | [default to &quot;&quot;]
 **expTime** | **string** | Contract expiry date, the format is &#x60;YYYYMMdd&#x60;, e.g. &#x60;20210623&#x60; | [default to &quot;&quot;]
 **period** | **string** | Period, the default is &#x60;8H&#x60;. e.g. [&#x60;8H/1D&#x60;]    Each granularity can provide only one latest piece of data | [default to &quot;&quot;]

### Return type

[**GetRubikStatOptionOpenInterestVolumeStrikeV5Resp**](GetRubikStatOptionOpenInterestVolumeStrikeV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatOptionOpenInterestVolumeV5

> GetRubikStatOptionOpenInterestVolumeV5Resp GetRubikStatOptionOpenInterestVolumeV5(ctx).Ccy(ccy).Period(period).Execute()

Get options open interest and volume



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
	ccy := "ccy_example" // string | Currency (default to "")
	period := "period_example" // string | Period, the default is `8H`. e.g. [`8H/1D`]    Each granularity can only query 72 pieces of data at the earliest (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatOptionOpenInterestVolumeV5(context.Background()).Ccy(ccy).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatOptionOpenInterestVolumeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatOptionOpenInterestVolumeV5`: GetRubikStatOptionOpenInterestVolumeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatOptionOpenInterestVolumeV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatOptionOpenInterestVolumeV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency | [default to &quot;&quot;]
 **period** | **string** | Period, the default is &#x60;8H&#x60;. e.g. [&#x60;8H/1D&#x60;]    Each granularity can only query 72 pieces of data at the earliest | [default to &quot;&quot;]

### Return type

[**GetRubikStatOptionOpenInterestVolumeV5Resp**](GetRubikStatOptionOpenInterestVolumeV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatOptionTakerBlockVolumeV5

> GetRubikStatOptionTakerBlockVolumeV5Resp GetRubikStatOptionTakerBlockVolumeV5(ctx).Ccy(ccy).Period(period).Execute()

Get taker flow



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
	ccy := "ccy_example" // string | currency (default to "")
	period := "period_example" // string | period, the default is `8H`. e.g. [`8H/1D`]    Each granularity can provide only one latest piece of data (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatOptionTakerBlockVolumeV5(context.Background()).Ccy(ccy).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatOptionTakerBlockVolumeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatOptionTakerBlockVolumeV5`: GetRubikStatOptionTakerBlockVolumeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatOptionTakerBlockVolumeV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatOptionTakerBlockVolumeV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | currency | [default to &quot;&quot;]
 **period** | **string** | period, the default is &#x60;8H&#x60;. e.g. [&#x60;8H/1D&#x60;]    Each granularity can provide only one latest piece of data | [default to &quot;&quot;]

### Return type

[**GetRubikStatOptionTakerBlockVolumeV5Resp**](GetRubikStatOptionTakerBlockVolumeV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatTakerVolumeContractV5

> GetRubikStatTakerVolumeContractV5Resp GetRubikStatTakerVolumeContractV5(ctx).InstId(instId).Period(period).Unit(unit).End(end).Begin(begin).Limit(limit).Execute()

Get contract taker volume



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
	instId := "instId_example" // string | Instrument ID, eg: BTC-USDT-SWAP   Only applicable to `FUTURES`, `SWAP` (default to "")
	period := "period_example" // string | Bar size, the default is `5m`, e.g. [`5m/15m/30m/1H/2H/4H`]   Hong Kong time opening price k-line:[`6H/12H/1D/2D/3D/5D/1W/1M/3M`]   UTC time opening price k-line: [`6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/5Dutc/1Wutc/1Mutc/3Mutc`] (optional) (default to "")
	unit := "unit_example" // string | The unit of buy/sell volume, the default is `1`   `0`: Crypto   `1`: Contracts   `2`: U (optional) (default to "")
	end := "end_example" // string | return records earlier than the requested `ts` (optional) (default to "")
	begin := "begin_example" // string | return records newer than the requested `ts` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatTakerVolumeContractV5(context.Background()).InstId(instId).Period(period).Unit(unit).End(end).Begin(begin).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatTakerVolumeContractV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatTakerVolumeContractV5`: GetRubikStatTakerVolumeContractV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatTakerVolumeContractV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatTakerVolumeContractV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, eg: BTC-USDT-SWAP   Only applicable to &#x60;FUTURES&#x60;, &#x60;SWAP&#x60; | [default to &quot;&quot;]
 **period** | **string** | Bar size, the default is &#x60;5m&#x60;, e.g. [&#x60;5m/15m/30m/1H/2H/4H&#x60;]   Hong Kong time opening price k-line:[&#x60;6H/12H/1D/2D/3D/5D/1W/1M/3M&#x60;]   UTC time opening price k-line: [&#x60;6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/5Dutc/1Wutc/1Mutc/3Mutc&#x60;] | [default to &quot;&quot;]
 **unit** | **string** | The unit of buy/sell volume, the default is &#x60;1&#x60;   &#x60;0&#x60;: Crypto   &#x60;1&#x60;: Contracts   &#x60;2&#x60;: U | [default to &quot;&quot;]
 **end** | **string** | return records earlier than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **begin** | **string** | return records newer than the requested &#x60;ts&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetRubikStatTakerVolumeContractV5Resp**](GetRubikStatTakerVolumeContractV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatTakerVolumeV5

> GetRubikStatTakerVolumeV5Resp GetRubikStatTakerVolumeV5(ctx).Ccy(ccy).InstType(instType).Begin(begin).End(end).Period(period).Execute()

Get taker volume



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
	ccy := "ccy_example" // string | Currency (default to "")
	instType := "instType_example" // string | Instrument type  `SPOT`  `CONTRACTS` (default to "")
	begin := "begin_example" // string | Begin time, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	end := "end_example" // string | End time, Unix timestamp format in milliseconds, e.g. `1597026383011` (optional) (default to "")
	period := "period_example" // string | Period, the default is `5m`, e.g. [`5m`/`1H`/`1D`]    `5m` granularity can only query data within two days at most  `1H` granularity can only query data within 30 days at most   `1D` granularity can only query data within 180 days at most (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatTakerVolumeV5(context.Background()).Ccy(ccy).InstType(instType).Begin(begin).End(end).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatTakerVolumeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatTakerVolumeV5`: GetRubikStatTakerVolumeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatTakerVolumeV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatTakerVolumeV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;CONTRACTS&#x60; | [default to &quot;&quot;]
 **begin** | **string** | Begin time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **end** | **string** | End time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383011&#x60; | [default to &quot;&quot;]
 **period** | **string** | Period, the default is &#x60;5m&#x60;, e.g. [&#x60;5m&#x60;/&#x60;1H&#x60;/&#x60;1D&#x60;]    &#x60;5m&#x60; granularity can only query data within two days at most  &#x60;1H&#x60; granularity can only query data within 30 days at most   &#x60;1D&#x60; granularity can only query data within 180 days at most | [default to &quot;&quot;]

### Return type

[**GetRubikStatTakerVolumeV5Resp**](GetRubikStatTakerVolumeV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRubikStatTradingDataSupportCoinV5

> GetRubikStatTradingDataSupportCoinV5Resp GetRubikStatTradingDataSupportCoinV5(ctx).Execute()

Get support coin



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
	resp, r, err := apiClient.TradingStatisticsAPI.GetRubikStatTradingDataSupportCoinV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingStatisticsAPI.GetRubikStatTradingDataSupportCoinV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRubikStatTradingDataSupportCoinV5`: GetRubikStatTradingDataSupportCoinV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingStatisticsAPI.GetRubikStatTradingDataSupportCoinV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRubikStatTradingDataSupportCoinV5Request struct via the builder pattern


### Return type

[**GetRubikStatTradingDataSupportCoinV5Resp**](GetRubikStatTradingDataSupportCoinV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

