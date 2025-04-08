# \CopyTradingAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCopytradingAlgoOrderV5**](CopyTradingAPI.md#CreateCopytradingAlgoOrderV5) | **Post** /api/v5/copytrading/algo-order | Set TP/SL for the current lead position that are not closed.  
[**CreateCopytradingAmendCopySettingsV5**](CopyTradingAPI.md#CreateCopytradingAmendCopySettingsV5) | **Post** /api/v5/copytrading/amend-copy-settings | You need to use this endpoint to amend copy settings  
[**CreateCopytradingAmendProfitSharingRatioV5**](CopyTradingAPI.md#CreateCopytradingAmendProfitSharingRatioV5) | **Post** /api/v5/copytrading/amend-profit-sharing-ratio | It is used to amend profit sharing ratio.   
[**CreateCopytradingCloseSubpositionV5**](CopyTradingAPI.md#CreateCopytradingCloseSubpositionV5) | **Post** /api/v5/copytrading/close-subposition | You can only close a lead position once a time.    It is required to pass subPosId which can get from .  
[**CreateCopytradingFirstCopySettingsV5**](CopyTradingAPI.md#CreateCopytradingFirstCopySettingsV5) | **Post** /api/v5/copytrading/first-copy-settings | The first copy settings for the certain lead trader. You need to first copy settings after stopping copying.  
[**CreateCopytradingSetInstrumentsV5**](CopyTradingAPI.md#CreateCopytradingSetInstrumentsV5) | **Post** /api/v5/copytrading/set-instruments | The leading trader can amend current leading instruments, need to set initial leading instruments while applying to become a leading trader.   All non-leading instruments can&#39;t have position or pending orders for the current request when setting non-leading instruments as leading instruments.  
[**CreateCopytradingStopCopyTradingV5**](CopyTradingAPI.md#CreateCopytradingStopCopyTradingV5) | **Post** /api/v5/copytrading/stop-copy-trading | You need to use this endpoint to stop copy trading  
[**GetCopytradingConfigV5**](CopyTradingAPI.md#GetCopytradingConfigV5) | **Get** /api/v5/copytrading/config | Retrieve current account configuration related to copy/lead trading.  
[**GetCopytradingCopySettingsV5**](CopyTradingAPI.md#GetCopytradingCopySettingsV5) | **Get** /api/v5/copytrading/copy-settings | Retrieve the copy settings about certain lead trader.  
[**GetCopytradingCurrentLeadTradersV5**](CopyTradingAPI.md#GetCopytradingCurrentLeadTradersV5) | **Get** /api/v5/copytrading/current-lead-traders | Retrieve my lead traders.  
[**GetCopytradingCurrentSubpositionsV5**](CopyTradingAPI.md#GetCopytradingCurrentSubpositionsV5) | **Get** /api/v5/copytrading/current-subpositions | Retrieve lead positions that are not closed.    Returns reverse chronological order with &#x60;openTime&#x60;  
[**GetCopytradingInstrumentsV5**](CopyTradingAPI.md#GetCopytradingInstrumentsV5) | **Get** /api/v5/copytrading/instruments | Retrieve instruments that are supported to lead by the platform.  Retrieve instruments that the lead trader has set.  
[**GetCopytradingProfitSharingDetailsV5**](CopyTradingAPI.md#GetCopytradingProfitSharingDetailsV5) | **Get** /api/v5/copytrading/profit-sharing-details | The leading trader gets profits shared details for the last 3 months.  
[**GetCopytradingPublicConfigV5**](CopyTradingAPI.md#GetCopytradingPublicConfigV5) | **Get** /api/v5/copytrading/public-config | Public endpoint. Retrieve copy trading parameter configuration information of copy settings  
[**GetCopytradingPublicCopyTradersV5**](CopyTradingAPI.md#GetCopytradingPublicCopyTradersV5) | **Get** /api/v5/copytrading/public-copy-traders | Public endpoint. Retrieve copy trader coming from certain lead trader. Return according to &#x60;pnl&#x60; from high to low  
[**GetCopytradingPublicCurrentSubpositionsV5**](CopyTradingAPI.md#GetCopytradingPublicCurrentSubpositionsV5) | **Get** /api/v5/copytrading/public-current-subpositions | Public endpoint. Get current leading positions of lead trader  
[**GetCopytradingPublicLeadTradersV5**](CopyTradingAPI.md#GetCopytradingPublicLeadTradersV5) | **Get** /api/v5/copytrading/public-lead-traders | Public endpoint. Retrieve lead trader ranks.  
[**GetCopytradingPublicPnlV5**](CopyTradingAPI.md#GetCopytradingPublicPnlV5) | **Get** /api/v5/copytrading/public-pnl | Public endpoint. Retrieve lead trader daily pnl. Results are returned in counter chronological order.  
[**GetCopytradingPublicPreferenceCurrencyV5**](CopyTradingAPI.md#GetCopytradingPublicPreferenceCurrencyV5) | **Get** /api/v5/copytrading/public-preference-currency | Public endpoint. The most frequently traded crypto of this lead trader. Results are sorted by ratio from large to small.  
[**GetCopytradingPublicStatsV5**](CopyTradingAPI.md#GetCopytradingPublicStatsV5) | **Get** /api/v5/copytrading/public-stats | Public endpoint. Key data related to lead trader performance.  
[**GetCopytradingPublicSubpositionsHistoryV5**](CopyTradingAPI.md#GetCopytradingPublicSubpositionsHistoryV5) | **Get** /api/v5/copytrading/public-subpositions-history | Public endpoint. Retrieve the lead trader completed leading position of the last 3 months.   Returns reverse chronological order with &#x60;subPosId&#x60;.   
[**GetCopytradingPublicWeeklyPnlV5**](CopyTradingAPI.md#GetCopytradingPublicWeeklyPnlV5) | **Get** /api/v5/copytrading/public-weekly-pnl | Public endpoint. Retrieve lead trader weekly pnl. Results are returned in counter chronological order.  
[**GetCopytradingSubpositionsHistoryV5**](CopyTradingAPI.md#GetCopytradingSubpositionsHistoryV5) | **Get** /api/v5/copytrading/subpositions-history | Retrieve the completed lead position of the last 3 months.   Returns reverse chronological order with &#x60;subPosId&#x60;.   
[**GetCopytradingTotalProfitSharingV5**](CopyTradingAPI.md#GetCopytradingTotalProfitSharingV5) | **Get** /api/v5/copytrading/total-profit-sharing | The leading trader gets the total amount of profit shared since joining the platform.  
[**GetCopytradingTotalUnrealizedProfitSharingV5**](CopyTradingAPI.md#GetCopytradingTotalUnrealizedProfitSharingV5) | **Get** /api/v5/copytrading/total-unrealized-profit-sharing | The leading trader gets the total unrealized amount of profit shared.  
[**GetCopytradingUnrealizedProfitSharingDetailsV5**](CopyTradingAPI.md#GetCopytradingUnrealizedProfitSharingDetailsV5) | **Get** /api/v5/copytrading/unrealized-profit-sharing-details | The leading trader gets the profit sharing details that are expected to be shared in the next settlement cycle.   The unrealized profit sharing details will update once there copy position is closed.  



## CreateCopytradingAlgoOrderV5

> CreateCopytradingAlgoOrderV5Resp CreateCopytradingAlgoOrderV5(ctx).CreateCopytradingAlgoOrderV5Req(createCopytradingAlgoOrderV5Req).Execute()

Set TP/SL for the current lead position that are not closed.  



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
	createCopytradingAlgoOrderV5Req := *openapiclient.NewCreateCopytradingAlgoOrderV5Req("SubPosId_example") // CreateCopytradingAlgoOrderV5Req | The request body for CreateCopytradingAlgoOrderV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.CreateCopytradingAlgoOrderV5(context.Background()).CreateCopytradingAlgoOrderV5Req(createCopytradingAlgoOrderV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.CreateCopytradingAlgoOrderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCopytradingAlgoOrderV5`: CreateCopytradingAlgoOrderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.CreateCopytradingAlgoOrderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCopytradingAlgoOrderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCopytradingAlgoOrderV5Req** | [**CreateCopytradingAlgoOrderV5Req**](CreateCopytradingAlgoOrderV5Req.md) | The request body for CreateCopytradingAlgoOrderV5 | 

### Return type

[**CreateCopytradingAlgoOrderV5Resp**](CreateCopytradingAlgoOrderV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCopytradingAmendCopySettingsV5

> CreateCopytradingAmendCopySettingsV5Resp CreateCopytradingAmendCopySettingsV5(ctx).CreateCopytradingAmendCopySettingsV5Req(createCopytradingAmendCopySettingsV5Req).Execute()

You need to use this endpoint to amend copy settings  



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
	createCopytradingAmendCopySettingsV5Req := *openapiclient.NewCreateCopytradingAmendCopySettingsV5Req("CopyInstIdType_example", "CopyMgnMode_example", "CopyTotalAmt_example", "SubPosCloseType_example", "UniqueCode_example") // CreateCopytradingAmendCopySettingsV5Req | The request body for CreateCopytradingAmendCopySettingsV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.CreateCopytradingAmendCopySettingsV5(context.Background()).CreateCopytradingAmendCopySettingsV5Req(createCopytradingAmendCopySettingsV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.CreateCopytradingAmendCopySettingsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCopytradingAmendCopySettingsV5`: CreateCopytradingAmendCopySettingsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.CreateCopytradingAmendCopySettingsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCopytradingAmendCopySettingsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCopytradingAmendCopySettingsV5Req** | [**CreateCopytradingAmendCopySettingsV5Req**](CreateCopytradingAmendCopySettingsV5Req.md) | The request body for CreateCopytradingAmendCopySettingsV5 | 

### Return type

[**CreateCopytradingAmendCopySettingsV5Resp**](CreateCopytradingAmendCopySettingsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCopytradingAmendProfitSharingRatioV5

> CreateCopytradingAmendProfitSharingRatioV5Resp CreateCopytradingAmendProfitSharingRatioV5(ctx).CreateCopytradingAmendProfitSharingRatioV5Req(createCopytradingAmendProfitSharingRatioV5Req).Execute()

It is used to amend profit sharing ratio.   



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
	createCopytradingAmendProfitSharingRatioV5Req := *openapiclient.NewCreateCopytradingAmendProfitSharingRatioV5Req("ProfitSharingRatio_example") // CreateCopytradingAmendProfitSharingRatioV5Req | The request body for CreateCopytradingAmendProfitSharingRatioV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.CreateCopytradingAmendProfitSharingRatioV5(context.Background()).CreateCopytradingAmendProfitSharingRatioV5Req(createCopytradingAmendProfitSharingRatioV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.CreateCopytradingAmendProfitSharingRatioV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCopytradingAmendProfitSharingRatioV5`: CreateCopytradingAmendProfitSharingRatioV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.CreateCopytradingAmendProfitSharingRatioV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCopytradingAmendProfitSharingRatioV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCopytradingAmendProfitSharingRatioV5Req** | [**CreateCopytradingAmendProfitSharingRatioV5Req**](CreateCopytradingAmendProfitSharingRatioV5Req.md) | The request body for CreateCopytradingAmendProfitSharingRatioV5 | 

### Return type

[**CreateCopytradingAmendProfitSharingRatioV5Resp**](CreateCopytradingAmendProfitSharingRatioV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCopytradingCloseSubpositionV5

> CreateCopytradingCloseSubpositionV5Resp CreateCopytradingCloseSubpositionV5(ctx).CreateCopytradingCloseSubpositionV5Req(createCopytradingCloseSubpositionV5Req).Execute()

You can only close a lead position once a time.    It is required to pass subPosId which can get from .  



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
	createCopytradingCloseSubpositionV5Req := *openapiclient.NewCreateCopytradingCloseSubpositionV5Req("SubPosId_example") // CreateCopytradingCloseSubpositionV5Req | The request body for CreateCopytradingCloseSubpositionV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.CreateCopytradingCloseSubpositionV5(context.Background()).CreateCopytradingCloseSubpositionV5Req(createCopytradingCloseSubpositionV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.CreateCopytradingCloseSubpositionV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCopytradingCloseSubpositionV5`: CreateCopytradingCloseSubpositionV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.CreateCopytradingCloseSubpositionV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCopytradingCloseSubpositionV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCopytradingCloseSubpositionV5Req** | [**CreateCopytradingCloseSubpositionV5Req**](CreateCopytradingCloseSubpositionV5Req.md) | The request body for CreateCopytradingCloseSubpositionV5 | 

### Return type

[**CreateCopytradingCloseSubpositionV5Resp**](CreateCopytradingCloseSubpositionV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCopytradingFirstCopySettingsV5

> CreateCopytradingFirstCopySettingsV5Resp CreateCopytradingFirstCopySettingsV5(ctx).CreateCopytradingFirstCopySettingsV5Req(createCopytradingFirstCopySettingsV5Req).Execute()

The first copy settings for the certain lead trader. You need to first copy settings after stopping copying.  



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
	createCopytradingFirstCopySettingsV5Req := *openapiclient.NewCreateCopytradingFirstCopySettingsV5Req("CopyInstIdType_example", "CopyMgnMode_example", "CopyTotalAmt_example", "SubPosCloseType_example", "UniqueCode_example") // CreateCopytradingFirstCopySettingsV5Req | The request body for CreateCopytradingFirstCopySettingsV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.CreateCopytradingFirstCopySettingsV5(context.Background()).CreateCopytradingFirstCopySettingsV5Req(createCopytradingFirstCopySettingsV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.CreateCopytradingFirstCopySettingsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCopytradingFirstCopySettingsV5`: CreateCopytradingFirstCopySettingsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.CreateCopytradingFirstCopySettingsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCopytradingFirstCopySettingsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCopytradingFirstCopySettingsV5Req** | [**CreateCopytradingFirstCopySettingsV5Req**](CreateCopytradingFirstCopySettingsV5Req.md) | The request body for CreateCopytradingFirstCopySettingsV5 | 

### Return type

[**CreateCopytradingFirstCopySettingsV5Resp**](CreateCopytradingFirstCopySettingsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCopytradingSetInstrumentsV5

> CreateCopytradingSetInstrumentsV5Resp CreateCopytradingSetInstrumentsV5(ctx).CreateCopytradingSetInstrumentsV5Req(createCopytradingSetInstrumentsV5Req).Execute()

The leading trader can amend current leading instruments, need to set initial leading instruments while applying to become a leading trader.   All non-leading instruments can't have position or pending orders for the current request when setting non-leading instruments as leading instruments.  



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
	createCopytradingSetInstrumentsV5Req := *openapiclient.NewCreateCopytradingSetInstrumentsV5Req("InstId_example") // CreateCopytradingSetInstrumentsV5Req | The request body for CreateCopytradingSetInstrumentsV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.CreateCopytradingSetInstrumentsV5(context.Background()).CreateCopytradingSetInstrumentsV5Req(createCopytradingSetInstrumentsV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.CreateCopytradingSetInstrumentsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCopytradingSetInstrumentsV5`: CreateCopytradingSetInstrumentsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.CreateCopytradingSetInstrumentsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCopytradingSetInstrumentsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCopytradingSetInstrumentsV5Req** | [**CreateCopytradingSetInstrumentsV5Req**](CreateCopytradingSetInstrumentsV5Req.md) | The request body for CreateCopytradingSetInstrumentsV5 | 

### Return type

[**CreateCopytradingSetInstrumentsV5Resp**](CreateCopytradingSetInstrumentsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCopytradingStopCopyTradingV5

> CreateCopytradingStopCopyTradingV5Resp CreateCopytradingStopCopyTradingV5(ctx).CreateCopytradingStopCopyTradingV5Req(createCopytradingStopCopyTradingV5Req).Execute()

You need to use this endpoint to stop copy trading  



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
	createCopytradingStopCopyTradingV5Req := *openapiclient.NewCreateCopytradingStopCopyTradingV5Req("SubPosCloseType_example", "UniqueCode_example") // CreateCopytradingStopCopyTradingV5Req | The request body for CreateCopytradingStopCopyTradingV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.CreateCopytradingStopCopyTradingV5(context.Background()).CreateCopytradingStopCopyTradingV5Req(createCopytradingStopCopyTradingV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.CreateCopytradingStopCopyTradingV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCopytradingStopCopyTradingV5`: CreateCopytradingStopCopyTradingV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.CreateCopytradingStopCopyTradingV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCopytradingStopCopyTradingV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCopytradingStopCopyTradingV5Req** | [**CreateCopytradingStopCopyTradingV5Req**](CreateCopytradingStopCopyTradingV5Req.md) | The request body for CreateCopytradingStopCopyTradingV5 | 

### Return type

[**CreateCopytradingStopCopyTradingV5Resp**](CreateCopytradingStopCopyTradingV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingConfigV5

> GetCopytradingConfigV5Resp GetCopytradingConfigV5(ctx).Execute()

Retrieve current account configuration related to copy/lead trading.  



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
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingConfigV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingConfigV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingConfigV5`: GetCopytradingConfigV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingConfigV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingConfigV5Request struct via the builder pattern


### Return type

[**GetCopytradingConfigV5Resp**](GetCopytradingConfigV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingCopySettingsV5

> GetCopytradingCopySettingsV5Resp GetCopytradingCopySettingsV5(ctx).UniqueCode(uniqueCode).InstType(instType).Execute()

Retrieve the copy settings about certain lead trader.  



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
	uniqueCode := "uniqueCode_example" // string | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC (default to "")
	instType := "instType_example" // string | Instrument type  `SWAP` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingCopySettingsV5(context.Background()).UniqueCode(uniqueCode).InstType(instType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingCopySettingsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingCopySettingsV5`: GetCopytradingCopySettingsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingCopySettingsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingCopySettingsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueCode** | **string** | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60; | [default to &quot;&quot;]

### Return type

[**GetCopytradingCopySettingsV5Resp**](GetCopytradingCopySettingsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingCurrentLeadTradersV5

> GetCopytradingCurrentLeadTradersV5Resp GetCopytradingCurrentLeadTradersV5(ctx).InstType(instType).Execute()

Retrieve my lead traders.  



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
	instType := "instType_example" // string | Instrument type  `SWAP`, the default value (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingCurrentLeadTradersV5(context.Background()).InstType(instType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingCurrentLeadTradersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingCurrentLeadTradersV5`: GetCopytradingCurrentLeadTradersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingCurrentLeadTradersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingCurrentLeadTradersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;, the default value | [default to &quot;&quot;]

### Return type

[**GetCopytradingCurrentLeadTradersV5Resp**](GetCopytradingCurrentLeadTradersV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingCurrentSubpositionsV5

> GetCopytradingCurrentSubpositionsV5Resp GetCopytradingCurrentSubpositionsV5(ctx).InstType(instType).InstId(instId).After(after).Before(before).Limit(limit).Execute()

Retrieve lead positions that are not closed.    Returns reverse chronological order with `openTime`  



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
	instType := "instType_example" // string | Instrument type  `SPOT`  `SWAP`  It returns all types by default. (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. BTC-USDT-SWAP (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `subPosId`. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `subPosId`. (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. Maximum is 500. Default is 500. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingCurrentSubpositionsV5(context.Background()).InstType(instType).InstId(instId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingCurrentSubpositionsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingCurrentSubpositionsV5`: GetCopytradingCurrentSubpositionsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingCurrentSubpositionsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingCurrentSubpositionsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;  It returns all types by default. | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. BTC-USDT-SWAP | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;subPosId&#x60;. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;subPosId&#x60;. | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. Maximum is 500. Default is 500. | [default to &quot;&quot;]

### Return type

[**GetCopytradingCurrentSubpositionsV5Resp**](GetCopytradingCurrentSubpositionsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingInstrumentsV5

> GetCopytradingInstrumentsV5Resp GetCopytradingInstrumentsV5(ctx).InstType(instType).Execute()

Retrieve instruments that are supported to lead by the platform.  Retrieve instruments that the lead trader has set.  



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
	instType := "instType_example" // string | Instrument type  `SPOT`  `SWAP`, the default value (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingInstrumentsV5(context.Background()).InstType(instType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingInstrumentsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingInstrumentsV5`: GetCopytradingInstrumentsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingInstrumentsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingInstrumentsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;, the default value | [default to &quot;&quot;]

### Return type

[**GetCopytradingInstrumentsV5Resp**](GetCopytradingInstrumentsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingProfitSharingDetailsV5

> GetCopytradingProfitSharingDetailsV5Resp GetCopytradingProfitSharingDetailsV5(ctx).InstType(instType).After(after).Before(before).Limit(limit).Execute()

The leading trader gets profits shared details for the last 3 months.  



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
	instType := "instType_example" // string | Instrument type  `SPOT`  `SWAP`  It returns all types by default. (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `profitSharingId` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `profitSharingId` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. Maximum is 100. Default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingProfitSharingDetailsV5(context.Background()).InstType(instType).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingProfitSharingDetailsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingProfitSharingDetailsV5`: GetCopytradingProfitSharingDetailsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingProfitSharingDetailsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingProfitSharingDetailsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;  It returns all types by default. | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;profitSharingId&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;profitSharingId&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. Maximum is 100. Default is 100. | [default to &quot;&quot;]

### Return type

[**GetCopytradingProfitSharingDetailsV5Resp**](GetCopytradingProfitSharingDetailsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingPublicConfigV5

> GetCopytradingPublicConfigV5Resp GetCopytradingPublicConfigV5(ctx).InstType(instType).Execute()

Public endpoint. Retrieve copy trading parameter configuration information of copy settings  



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
	instType := "instType_example" // string | Instrument type  `SWAP`, the default value (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingPublicConfigV5(context.Background()).InstType(instType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingPublicConfigV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingPublicConfigV5`: GetCopytradingPublicConfigV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingPublicConfigV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingPublicConfigV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;, the default value | [default to &quot;&quot;]

### Return type

[**GetCopytradingPublicConfigV5Resp**](GetCopytradingPublicConfigV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingPublicCopyTradersV5

> GetCopytradingPublicCopyTradersV5Resp GetCopytradingPublicCopyTradersV5(ctx).UniqueCode(uniqueCode).InstType(instType).Limit(limit).Execute()

Public endpoint. Retrieve copy trader coming from certain lead trader. Return according to `pnl` from high to low  



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
	uniqueCode := "uniqueCode_example" // string | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC (default to "")
	instType := "instType_example" // string | Instrument type  `SWAP`, the default value (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingPublicCopyTradersV5(context.Background()).UniqueCode(uniqueCode).InstType(instType).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingPublicCopyTradersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingPublicCopyTradersV5`: GetCopytradingPublicCopyTradersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingPublicCopyTradersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingPublicCopyTradersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueCode** | **string** | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;, the default value | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetCopytradingPublicCopyTradersV5Resp**](GetCopytradingPublicCopyTradersV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingPublicCurrentSubpositionsV5

> GetCopytradingPublicCurrentSubpositionsV5Resp GetCopytradingPublicCurrentSubpositionsV5(ctx).UniqueCode(uniqueCode).InstType(instType).After(after).Before(before).Limit(limit).Execute()

Public endpoint. Get current leading positions of lead trader  



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
	uniqueCode := "uniqueCode_example" // string | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC (default to "")
	instType := "instType_example" // string | Instrument type  `SWAP`, the default value. (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `subPosId`. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `subPosId`. (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. Maximum is 100. Default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingPublicCurrentSubpositionsV5(context.Background()).UniqueCode(uniqueCode).InstType(instType).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingPublicCurrentSubpositionsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingPublicCurrentSubpositionsV5`: GetCopytradingPublicCurrentSubpositionsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingPublicCurrentSubpositionsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingPublicCurrentSubpositionsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueCode** | **string** | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;, the default value. | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;subPosId&#x60;. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;subPosId&#x60;. | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. Maximum is 100. Default is 100. | [default to &quot;&quot;]

### Return type

[**GetCopytradingPublicCurrentSubpositionsV5Resp**](GetCopytradingPublicCurrentSubpositionsV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingPublicLeadTradersV5

> GetCopytradingPublicLeadTradersV5Resp GetCopytradingPublicLeadTradersV5(ctx).InstType(instType).SortType(sortType).State(state).MinLeadDays(minLeadDays).MinAssets(minAssets).MaxAssets(maxAssets).MinAum(minAum).MaxAum(maxAum).DataVer(dataVer).Page(page).Limit(limit).Execute()

Public endpoint. Retrieve lead trader ranks.  



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
	instType := "instType_example" // string | Instrument type  `SWAP`, the default value (optional) (default to "")
	sortType := "sortType_example" // string | Sort type  `overview`: overview, the default value  `pnl`: profit and loss  `aum`: assets under management  `win_ratio`: win ratio  `pnl_ratio`: pnl ratio  `current_copy_trader_pnl`: current copy trader pnl (optional) (default to "")
	state := "state_example" // string | Lead trader state  `0`: All lead traders, the default, including vacancy and non-vacancy   `1`: lead traders who have vacancy (optional) (default to "")
	minLeadDays := "minLeadDays_example" // string | Minimum lead days  `1`: 7 days  `2`: 30 days  `3`: 90 days  `4`: 180 days (optional) (default to "")
	minAssets := "minAssets_example" // string | Minimum assets in USDT (optional) (default to "")
	maxAssets := "maxAssets_example" // string | Maximum assets in USDT (optional) (default to "")
	minAum := "minAum_example" // string | Minimum assets in USDT under management. (optional) (default to "")
	maxAum := "maxAum_example" // string | Maximum assets in USDT under management. (optional) (default to "")
	dataVer := "dataVer_example" // string | Data version. It is 14 numbers. e.g. 20231010182400. Generally, it is used for pagination   A new version will be generated every 10 minutes. Only last 5 versions are stored  The default is latest version. If it is not exist, error will not be throwed and the latest version will be used. (optional) (default to "")
	page := "page_example" // string | Page for pagination (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 20; the default is 10 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingPublicLeadTradersV5(context.Background()).InstType(instType).SortType(sortType).State(state).MinLeadDays(minLeadDays).MinAssets(minAssets).MaxAssets(maxAssets).MinAum(minAum).MaxAum(maxAum).DataVer(dataVer).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingPublicLeadTradersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingPublicLeadTradersV5`: GetCopytradingPublicLeadTradersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingPublicLeadTradersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingPublicLeadTradersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;, the default value | [default to &quot;&quot;]
 **sortType** | **string** | Sort type  &#x60;overview&#x60;: overview, the default value  &#x60;pnl&#x60;: profit and loss  &#x60;aum&#x60;: assets under management  &#x60;win_ratio&#x60;: win ratio  &#x60;pnl_ratio&#x60;: pnl ratio  &#x60;current_copy_trader_pnl&#x60;: current copy trader pnl | [default to &quot;&quot;]
 **state** | **string** | Lead trader state  &#x60;0&#x60;: All lead traders, the default, including vacancy and non-vacancy   &#x60;1&#x60;: lead traders who have vacancy | [default to &quot;&quot;]
 **minLeadDays** | **string** | Minimum lead days  &#x60;1&#x60;: 7 days  &#x60;2&#x60;: 30 days  &#x60;3&#x60;: 90 days  &#x60;4&#x60;: 180 days | [default to &quot;&quot;]
 **minAssets** | **string** | Minimum assets in USDT | [default to &quot;&quot;]
 **maxAssets** | **string** | Maximum assets in USDT | [default to &quot;&quot;]
 **minAum** | **string** | Minimum assets in USDT under management. | [default to &quot;&quot;]
 **maxAum** | **string** | Maximum assets in USDT under management. | [default to &quot;&quot;]
 **dataVer** | **string** | Data version. It is 14 numbers. e.g. 20231010182400. Generally, it is used for pagination   A new version will be generated every 10 minutes. Only last 5 versions are stored  The default is latest version. If it is not exist, error will not be throwed and the latest version will be used. | [default to &quot;&quot;]
 **page** | **string** | Page for pagination | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 20; the default is 10 | [default to &quot;&quot;]

### Return type

[**GetCopytradingPublicLeadTradersV5Resp**](GetCopytradingPublicLeadTradersV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingPublicPnlV5

> GetCopytradingPublicPnlV5Resp GetCopytradingPublicPnlV5(ctx).UniqueCode(uniqueCode).LastDays(lastDays).InstType(instType).Execute()

Public endpoint. Retrieve lead trader daily pnl. Results are returned in counter chronological order.  



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
	uniqueCode := "uniqueCode_example" // string | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC (default to "")
	lastDays := "lastDays_example" // string | Last days  `1`: last 7 days   `2`: last 30 days  `3`: last 90 days   `4`: last 365 days (default to "")
	instType := "instType_example" // string | Instrument type  `SWAP`, the default value (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingPublicPnlV5(context.Background()).UniqueCode(uniqueCode).LastDays(lastDays).InstType(instType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingPublicPnlV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingPublicPnlV5`: GetCopytradingPublicPnlV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingPublicPnlV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingPublicPnlV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueCode** | **string** | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC | [default to &quot;&quot;]
 **lastDays** | **string** | Last days  &#x60;1&#x60;: last 7 days   &#x60;2&#x60;: last 30 days  &#x60;3&#x60;: last 90 days   &#x60;4&#x60;: last 365 days | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;, the default value | [default to &quot;&quot;]

### Return type

[**GetCopytradingPublicPnlV5Resp**](GetCopytradingPublicPnlV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingPublicPreferenceCurrencyV5

> GetCopytradingPublicPreferenceCurrencyV5Resp GetCopytradingPublicPreferenceCurrencyV5(ctx).UniqueCode(uniqueCode).InstType(instType).Execute()

Public endpoint. The most frequently traded crypto of this lead trader. Results are sorted by ratio from large to small.  



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
	uniqueCode := "uniqueCode_example" // string | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC (default to "")
	instType := "instType_example" // string | Instrument type  `SWAP`, the default value (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingPublicPreferenceCurrencyV5(context.Background()).UniqueCode(uniqueCode).InstType(instType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingPublicPreferenceCurrencyV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingPublicPreferenceCurrencyV5`: GetCopytradingPublicPreferenceCurrencyV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingPublicPreferenceCurrencyV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingPublicPreferenceCurrencyV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueCode** | **string** | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;, the default value | [default to &quot;&quot;]

### Return type

[**GetCopytradingPublicPreferenceCurrencyV5Resp**](GetCopytradingPublicPreferenceCurrencyV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingPublicStatsV5

> GetCopytradingPublicStatsV5Resp GetCopytradingPublicStatsV5(ctx).UniqueCode(uniqueCode).LastDays(lastDays).InstType(instType).Execute()

Public endpoint. Key data related to lead trader performance.  



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
	uniqueCode := "uniqueCode_example" // string | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC (default to "")
	lastDays := "lastDays_example" // string | Last days  `1`: last 7 days   `2`: last 30 days  `3`: last 90 days   `4`: last 365 days (default to "")
	instType := "instType_example" // string | Instrument type  `SWAP`, the default value (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingPublicStatsV5(context.Background()).UniqueCode(uniqueCode).LastDays(lastDays).InstType(instType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingPublicStatsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingPublicStatsV5`: GetCopytradingPublicStatsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingPublicStatsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingPublicStatsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueCode** | **string** | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC | [default to &quot;&quot;]
 **lastDays** | **string** | Last days  &#x60;1&#x60;: last 7 days   &#x60;2&#x60;: last 30 days  &#x60;3&#x60;: last 90 days   &#x60;4&#x60;: last 365 days | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;, the default value | [default to &quot;&quot;]

### Return type

[**GetCopytradingPublicStatsV5Resp**](GetCopytradingPublicStatsV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingPublicSubpositionsHistoryV5

> GetCopytradingPublicSubpositionsHistoryV5Resp GetCopytradingPublicSubpositionsHistoryV5(ctx).UniqueCode(uniqueCode).InstType(instType).After(after).Before(before).Limit(limit).Execute()

Public endpoint. Retrieve the lead trader completed leading position of the last 3 months.   Returns reverse chronological order with `subPosId`.   



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
	uniqueCode := "uniqueCode_example" // string | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC (default to "")
	instType := "instType_example" // string | Instrument type  `SWAP`, the default value. (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `subPosId`. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `subPosId`. (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. Maximum is 100. Default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingPublicSubpositionsHistoryV5(context.Background()).UniqueCode(uniqueCode).InstType(instType).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingPublicSubpositionsHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingPublicSubpositionsHistoryV5`: GetCopytradingPublicSubpositionsHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingPublicSubpositionsHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingPublicSubpositionsHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueCode** | **string** | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;, the default value. | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;subPosId&#x60;. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;subPosId&#x60;. | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. Maximum is 100. Default is 100. | [default to &quot;&quot;]

### Return type

[**GetCopytradingPublicSubpositionsHistoryV5Resp**](GetCopytradingPublicSubpositionsHistoryV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingPublicWeeklyPnlV5

> GetCopytradingPublicWeeklyPnlV5Resp GetCopytradingPublicWeeklyPnlV5(ctx).UniqueCode(uniqueCode).InstType(instType).Execute()

Public endpoint. Retrieve lead trader weekly pnl. Results are returned in counter chronological order.  



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
	uniqueCode := "uniqueCode_example" // string | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC (default to "")
	instType := "instType_example" // string | Instrument type  `SWAP`, the default value (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingPublicWeeklyPnlV5(context.Background()).UniqueCode(uniqueCode).InstType(instType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingPublicWeeklyPnlV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingPublicWeeklyPnlV5`: GetCopytradingPublicWeeklyPnlV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingPublicWeeklyPnlV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingPublicWeeklyPnlV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uniqueCode** | **string** | Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC | [default to &quot;&quot;]
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;, the default value | [default to &quot;&quot;]

### Return type

[**GetCopytradingPublicWeeklyPnlV5Resp**](GetCopytradingPublicWeeklyPnlV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingSubpositionsHistoryV5

> GetCopytradingSubpositionsHistoryV5Resp GetCopytradingSubpositionsHistoryV5(ctx).InstType(instType).InstId(instId).After(after).Before(before).Limit(limit).Execute()

Retrieve the completed lead position of the last 3 months.   Returns reverse chronological order with `subPosId`.   



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
	instType := "instType_example" // string | Instrument type  `SPOT`  `SWAP`  It returns all types by default. (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. BTC-USDT-SWAP (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `subPosId`. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `subPosId`. (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. Maximum is 100. Default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingSubpositionsHistoryV5(context.Background()).InstType(instType).InstId(instId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingSubpositionsHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingSubpositionsHistoryV5`: GetCopytradingSubpositionsHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingSubpositionsHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingSubpositionsHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;  It returns all types by default. | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. BTC-USDT-SWAP | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;subPosId&#x60;. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;subPosId&#x60;. | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. Maximum is 100. Default is 100. | [default to &quot;&quot;]

### Return type

[**GetCopytradingSubpositionsHistoryV5Resp**](GetCopytradingSubpositionsHistoryV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingTotalProfitSharingV5

> GetCopytradingTotalProfitSharingV5Resp GetCopytradingTotalProfitSharingV5(ctx).InstType(instType).Execute()

The leading trader gets the total amount of profit shared since joining the platform.  



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
	instType := "instType_example" // string | Instrument type  `SPOT`  `SWAP`  It returns all types by default. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingTotalProfitSharingV5(context.Background()).InstType(instType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingTotalProfitSharingV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingTotalProfitSharingV5`: GetCopytradingTotalProfitSharingV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingTotalProfitSharingV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingTotalProfitSharingV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;  It returns all types by default. | [default to &quot;&quot;]

### Return type

[**GetCopytradingTotalProfitSharingV5Resp**](GetCopytradingTotalProfitSharingV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingTotalUnrealizedProfitSharingV5

> GetCopytradingTotalUnrealizedProfitSharingV5Resp GetCopytradingTotalUnrealizedProfitSharingV5(ctx).InstType(instType).Execute()

The leading trader gets the total unrealized amount of profit shared.  



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
	instType := "instType_example" // string | Instrument type  `SWAP`, the default value. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingTotalUnrealizedProfitSharingV5(context.Background()).InstType(instType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingTotalUnrealizedProfitSharingV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingTotalUnrealizedProfitSharingV5`: GetCopytradingTotalUnrealizedProfitSharingV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingTotalUnrealizedProfitSharingV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingTotalUnrealizedProfitSharingV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;, the default value. | [default to &quot;&quot;]

### Return type

[**GetCopytradingTotalUnrealizedProfitSharingV5Resp**](GetCopytradingTotalUnrealizedProfitSharingV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCopytradingUnrealizedProfitSharingDetailsV5

> GetCopytradingUnrealizedProfitSharingDetailsV5Resp GetCopytradingUnrealizedProfitSharingDetailsV5(ctx).InstType(instType).Execute()

The leading trader gets the profit sharing details that are expected to be shared in the next settlement cycle.   The unrealized profit sharing details will update once there copy position is closed.  



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
	instType := "instType_example" // string | Instrument type  `SPOT`  `SWAP`  It returns all types by default. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CopyTradingAPI.GetCopytradingUnrealizedProfitSharingDetailsV5(context.Background()).InstType(instType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CopyTradingAPI.GetCopytradingUnrealizedProfitSharingDetailsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCopytradingUnrealizedProfitSharingDetailsV5`: GetCopytradingUnrealizedProfitSharingDetailsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `CopyTradingAPI.GetCopytradingUnrealizedProfitSharingDetailsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCopytradingUnrealizedProfitSharingDetailsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;  It returns all types by default. | [default to &quot;&quot;]

### Return type

[**GetCopytradingUnrealizedProfitSharingDetailsV5Resp**](GetCopytradingUnrealizedProfitSharingDetailsV5Resp.md)

### Authorization

[OK-ACCESS-KEY](../README.md#OK-ACCESS-KEY), [OK-ACCESS-PASSPHRASE](../README.md#OK-ACCESS-PASSPHRASE)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

