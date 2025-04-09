# \BlockTradingAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRfqCancelAllAfterV5**](BlockTradingAPI.md#CreateRfqCancelAllAfterV5) | **Post** /api/v5/rfq/cancel-all-after | Cancel all quotes after the countdown timeout.  
[**CreateRfqCancelAllQuotesV5**](BlockTradingAPI.md#CreateRfqCancelAllQuotesV5) | **Post** /api/v5/rfq/cancel-all-quotes | Cancels all active Quotes.  
[**CreateRfqCancelAllRfqsV5**](BlockTradingAPI.md#CreateRfqCancelAllRfqsV5) | **Post** /api/v5/rfq/cancel-all-rfqs | Cancels all active RFQs.  
[**CreateRfqCancelBatchQuotesV5**](BlockTradingAPI.md#CreateRfqCancelBatchQuotesV5) | **Post** /api/v5/rfq/cancel-batch-quotes | Cancel one or multiple active Quote(s) in a single batch. Maximum 100 quote orders can be canceled per request.  
[**CreateRfqCancelBatchRfqsV5**](BlockTradingAPI.md#CreateRfqCancelBatchRfqsV5) | **Post** /api/v5/rfq/cancel-batch-rfqs | Cancel one or multiple active RFQ(s) in a single batch. Maximum 100 RFQ orders can be canceled per request.  
[**CreateRfqCancelQuoteV5**](BlockTradingAPI.md#CreateRfqCancelQuoteV5) | **Post** /api/v5/rfq/cancel-quote | Cancels an existing active Quote you have created in response to an RFQ.  
[**CreateRfqCancelRfqV5**](BlockTradingAPI.md#CreateRfqCancelRfqV5) | **Post** /api/v5/rfq/cancel-rfq | Cancel an existing active RFQ that you have created previously.  
[**CreateRfqCreateQuoteV5**](BlockTradingAPI.md#CreateRfqCreateQuoteV5) | **Post** /api/v5/rfq/create-quote | Allows the user to Quote an RFQ that they are a counterparty to. The user MUST quote the entire RFQ and not part of the legs or part of the quantity. Partial quoting is not allowed.   
[**CreateRfqCreateRfqV5**](BlockTradingAPI.md#CreateRfqCreateRfqV5) | **Post** /api/v5/rfq/create-rfq | Creates a new RFQ  To learn more, please visit   
[**CreateRfqExecuteQuoteV5**](BlockTradingAPI.md#CreateRfqExecuteQuoteV5) | **Post** /api/v5/rfq/execute-quote | Executes a Quote. It is only used by the creator of the RFQ  
[**CreateRfqMakerInstrumentSettingsV5**](BlockTradingAPI.md#CreateRfqMakerInstrumentSettingsV5) | **Post** /api/v5/rfq/maker-instrument-settings | Customize the products which makers want to quote and receive RFQs for, and the corresponding price and size limit.   
[**CreateRfqMmpConfigV5**](BlockTradingAPI.md#CreateRfqMmpConfigV5) | **Post** /api/v5/rfq/mmp-config | This endpoint is used to set MMP configure and only applicable to block trading makers    
[**CreateRfqMmpResetV5**](BlockTradingAPI.md#CreateRfqMmpResetV5) | **Post** /api/v5/rfq/mmp-reset | Reset the MMP status to be inactive.  
[**GetMarketBlockTickerV5**](BlockTradingAPI.md#GetMarketBlockTickerV5) | **Get** /api/v5/market/block-ticker | Retrieve the latest block trading volume in the last 24 hours.  
[**GetMarketBlockTickersV5**](BlockTradingAPI.md#GetMarketBlockTickersV5) | **Get** /api/v5/market/block-tickers | Retrieve the latest block trading volume in the last 24 hours.  
[**GetPublicBlockTradesV5**](BlockTradingAPI.md#GetPublicBlockTradesV5) | **Get** /api/v5/public/block-trades | Retrieve the recent block trading transactions of an instrument. Descending order by tradeId. The data will be updated 15 minutes after the block trade execution.  
[**GetRfqCounterpartiesV5**](BlockTradingAPI.md#GetRfqCounterpartiesV5) | **Get** /api/v5/rfq/counterparties | Retrieves the list of counterparties that the user is permitted to trade with.   
[**GetRfqMakerInstrumentSettingsV5**](BlockTradingAPI.md#GetRfqMakerInstrumentSettingsV5) | **Get** /api/v5/rfq/maker-instrument-settings | Retrieve the products which makers want to quote and receive RFQs for, and the corresponding price and size limit.   
[**GetRfqMmpConfigV5**](BlockTradingAPI.md#GetRfqMmpConfigV5) | **Get** /api/v5/rfq/mmp-config | This endpoint is used to get MMP configure information and only applicable to block trading market makers    
[**GetRfqPublicTradesV5**](BlockTradingAPI.md#GetRfqPublicTradesV5) | **Get** /api/v5/rfq/public-trades | Retrieves the executed block trades. The data will be updated 15 minutes after the block trade execution.  
[**GetRfqQuotesV5**](BlockTradingAPI.md#GetRfqQuotesV5) | **Get** /api/v5/rfq/quotes | Retrieve all Quotes that the user is a counterparty to (either as the creator or the receiver).  
[**GetRfqRfqsV5**](BlockTradingAPI.md#GetRfqRfqsV5) | **Get** /api/v5/rfq/rfqs | Retrieves details of RFQs that the user is a counterparty to (either as the creator or the receiver of the RFQ).   
[**GetRfqTradesV5**](BlockTradingAPI.md#GetRfqTradesV5) | **Get** /api/v5/rfq/trades | Retrieves the executed trades that the user is a counterparty to (either as the creator or the receiver).  



## CreateRfqCancelAllAfterV5

> CreateRfqCancelAllAfterV5Resp CreateRfqCancelAllAfterV5(ctx).CreateRfqCancelAllAfterV5Req(createRfqCancelAllAfterV5Req).Execute()

Cancel all quotes after the countdown timeout.  



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
	createRfqCancelAllAfterV5Req := *openapiclient.NewCreateRfqCancelAllAfterV5Req("TimeOut_example") // CreateRfqCancelAllAfterV5Req | The request body for CreateRfqCancelAllAfterV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.CreateRfqCancelAllAfterV5(context.Background()).CreateRfqCancelAllAfterV5Req(createRfqCancelAllAfterV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.CreateRfqCancelAllAfterV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRfqCancelAllAfterV5`: CreateRfqCancelAllAfterV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.CreateRfqCancelAllAfterV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRfqCancelAllAfterV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRfqCancelAllAfterV5Req** | [**CreateRfqCancelAllAfterV5Req**](CreateRfqCancelAllAfterV5Req.md) | The request body for CreateRfqCancelAllAfterV5 | 

### Return type

[**CreateRfqCancelAllAfterV5Resp**](CreateRfqCancelAllAfterV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRfqCancelAllQuotesV5

> CreateRfqCancelAllQuotesV5Resp CreateRfqCancelAllQuotesV5(ctx).Execute()

Cancels all active Quotes.  



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
	resp, r, err := apiClient.BlockTradingAPI.CreateRfqCancelAllQuotesV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.CreateRfqCancelAllQuotesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRfqCancelAllQuotesV5`: CreateRfqCancelAllQuotesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.CreateRfqCancelAllQuotesV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRfqCancelAllQuotesV5Request struct via the builder pattern


### Return type

[**CreateRfqCancelAllQuotesV5Resp**](CreateRfqCancelAllQuotesV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRfqCancelAllRfqsV5

> CreateRfqCancelAllRfqsV5Resp CreateRfqCancelAllRfqsV5(ctx).Execute()

Cancels all active RFQs.  



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
	resp, r, err := apiClient.BlockTradingAPI.CreateRfqCancelAllRfqsV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.CreateRfqCancelAllRfqsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRfqCancelAllRfqsV5`: CreateRfqCancelAllRfqsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.CreateRfqCancelAllRfqsV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRfqCancelAllRfqsV5Request struct via the builder pattern


### Return type

[**CreateRfqCancelAllRfqsV5Resp**](CreateRfqCancelAllRfqsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRfqCancelBatchQuotesV5

> CreateRfqCancelBatchQuotesV5Resp CreateRfqCancelBatchQuotesV5(ctx).CreateRfqCancelBatchQuotesV5Req(createRfqCancelBatchQuotesV5Req).Execute()

Cancel one or multiple active Quote(s) in a single batch. Maximum 100 quote orders can be canceled per request.  



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
	createRfqCancelBatchQuotesV5Req := *openapiclient.NewCreateRfqCancelBatchQuotesV5Req() // CreateRfqCancelBatchQuotesV5Req | The request body for CreateRfqCancelBatchQuotesV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.CreateRfqCancelBatchQuotesV5(context.Background()).CreateRfqCancelBatchQuotesV5Req(createRfqCancelBatchQuotesV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.CreateRfqCancelBatchQuotesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRfqCancelBatchQuotesV5`: CreateRfqCancelBatchQuotesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.CreateRfqCancelBatchQuotesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRfqCancelBatchQuotesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRfqCancelBatchQuotesV5Req** | [**CreateRfqCancelBatchQuotesV5Req**](CreateRfqCancelBatchQuotesV5Req.md) | The request body for CreateRfqCancelBatchQuotesV5 | 

### Return type

[**CreateRfqCancelBatchQuotesV5Resp**](CreateRfqCancelBatchQuotesV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRfqCancelBatchRfqsV5

> CreateRfqCancelBatchRfqsV5Resp CreateRfqCancelBatchRfqsV5(ctx).CreateRfqCancelBatchRfqsV5Req(createRfqCancelBatchRfqsV5Req).Execute()

Cancel one or multiple active RFQ(s) in a single batch. Maximum 100 RFQ orders can be canceled per request.  



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
	createRfqCancelBatchRfqsV5Req := *openapiclient.NewCreateRfqCancelBatchRfqsV5Req() // CreateRfqCancelBatchRfqsV5Req | The request body for CreateRfqCancelBatchRfqsV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.CreateRfqCancelBatchRfqsV5(context.Background()).CreateRfqCancelBatchRfqsV5Req(createRfqCancelBatchRfqsV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.CreateRfqCancelBatchRfqsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRfqCancelBatchRfqsV5`: CreateRfqCancelBatchRfqsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.CreateRfqCancelBatchRfqsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRfqCancelBatchRfqsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRfqCancelBatchRfqsV5Req** | [**CreateRfqCancelBatchRfqsV5Req**](CreateRfqCancelBatchRfqsV5Req.md) | The request body for CreateRfqCancelBatchRfqsV5 | 

### Return type

[**CreateRfqCancelBatchRfqsV5Resp**](CreateRfqCancelBatchRfqsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRfqCancelQuoteV5

> CreateRfqCancelQuoteV5Resp CreateRfqCancelQuoteV5(ctx).CreateRfqCancelQuoteV5Req(createRfqCancelQuoteV5Req).Execute()

Cancels an existing active Quote you have created in response to an RFQ.  



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
	createRfqCancelQuoteV5Req := *openapiclient.NewCreateRfqCancelQuoteV5Req() // CreateRfqCancelQuoteV5Req | The request body for CreateRfqCancelQuoteV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.CreateRfqCancelQuoteV5(context.Background()).CreateRfqCancelQuoteV5Req(createRfqCancelQuoteV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.CreateRfqCancelQuoteV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRfqCancelQuoteV5`: CreateRfqCancelQuoteV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.CreateRfqCancelQuoteV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRfqCancelQuoteV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRfqCancelQuoteV5Req** | [**CreateRfqCancelQuoteV5Req**](CreateRfqCancelQuoteV5Req.md) | The request body for CreateRfqCancelQuoteV5 | 

### Return type

[**CreateRfqCancelQuoteV5Resp**](CreateRfqCancelQuoteV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRfqCancelRfqV5

> CreateRfqCancelRfqV5Resp CreateRfqCancelRfqV5(ctx).CreateRfqCancelRfqV5Req(createRfqCancelRfqV5Req).Execute()

Cancel an existing active RFQ that you have created previously.  



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
	createRfqCancelRfqV5Req := *openapiclient.NewCreateRfqCancelRfqV5Req() // CreateRfqCancelRfqV5Req | The request body for CreateRfqCancelRfqV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.CreateRfqCancelRfqV5(context.Background()).CreateRfqCancelRfqV5Req(createRfqCancelRfqV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.CreateRfqCancelRfqV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRfqCancelRfqV5`: CreateRfqCancelRfqV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.CreateRfqCancelRfqV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRfqCancelRfqV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRfqCancelRfqV5Req** | [**CreateRfqCancelRfqV5Req**](CreateRfqCancelRfqV5Req.md) | The request body for CreateRfqCancelRfqV5 | 

### Return type

[**CreateRfqCancelRfqV5Resp**](CreateRfqCancelRfqV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRfqCreateQuoteV5

> CreateRfqCreateQuoteV5Resp CreateRfqCreateQuoteV5(ctx).CreateRfqCreateQuoteV5Req(createRfqCreateQuoteV5Req).Execute()

Allows the user to Quote an RFQ that they are a counterparty to. The user MUST quote the entire RFQ and not part of the legs or part of the quantity. Partial quoting is not allowed.   



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
	createRfqCreateQuoteV5Req := *openapiclient.NewCreateRfqCreateQuoteV5Req([]openapiclient.CreateRfqCreateQuoteV5ReqLegsInner{*openapiclient.NewCreateRfqCreateQuoteV5ReqLegsInner()}, "QuoteSide_example", "RfqId_example") // CreateRfqCreateQuoteV5Req | The request body for CreateRfqCreateQuoteV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.CreateRfqCreateQuoteV5(context.Background()).CreateRfqCreateQuoteV5Req(createRfqCreateQuoteV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.CreateRfqCreateQuoteV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRfqCreateQuoteV5`: CreateRfqCreateQuoteV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.CreateRfqCreateQuoteV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRfqCreateQuoteV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRfqCreateQuoteV5Req** | [**CreateRfqCreateQuoteV5Req**](CreateRfqCreateQuoteV5Req.md) | The request body for CreateRfqCreateQuoteV5 | 

### Return type

[**CreateRfqCreateQuoteV5Resp**](CreateRfqCreateQuoteV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRfqCreateRfqV5

> CreateRfqCreateRfqV5Resp CreateRfqCreateRfqV5(ctx).CreateRfqCreateRfqV5Req(createRfqCreateRfqV5Req).Execute()

Creates a new RFQ  To learn more, please visit   



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
	createRfqCreateRfqV5Req := *openapiclient.NewCreateRfqCreateRfqV5Req([]string{"Counterparties_example"}, []openapiclient.CreateRfqCreateRfqV5ReqLegsInner{*openapiclient.NewCreateRfqCreateRfqV5ReqLegsInner()}) // CreateRfqCreateRfqV5Req | The request body for CreateRfqCreateRfqV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.CreateRfqCreateRfqV5(context.Background()).CreateRfqCreateRfqV5Req(createRfqCreateRfqV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.CreateRfqCreateRfqV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRfqCreateRfqV5`: CreateRfqCreateRfqV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.CreateRfqCreateRfqV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRfqCreateRfqV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRfqCreateRfqV5Req** | [**CreateRfqCreateRfqV5Req**](CreateRfqCreateRfqV5Req.md) | The request body for CreateRfqCreateRfqV5 | 

### Return type

[**CreateRfqCreateRfqV5Resp**](CreateRfqCreateRfqV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRfqExecuteQuoteV5

> CreateRfqExecuteQuoteV5Resp CreateRfqExecuteQuoteV5(ctx).CreateRfqExecuteQuoteV5Req(createRfqExecuteQuoteV5Req).Execute()

Executes a Quote. It is only used by the creator of the RFQ  



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
	createRfqExecuteQuoteV5Req := *openapiclient.NewCreateRfqExecuteQuoteV5Req("QuoteId_example", "RfqId_example") // CreateRfqExecuteQuoteV5Req | The request body for CreateRfqExecuteQuoteV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.CreateRfqExecuteQuoteV5(context.Background()).CreateRfqExecuteQuoteV5Req(createRfqExecuteQuoteV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.CreateRfqExecuteQuoteV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRfqExecuteQuoteV5`: CreateRfqExecuteQuoteV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.CreateRfqExecuteQuoteV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRfqExecuteQuoteV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRfqExecuteQuoteV5Req** | [**CreateRfqExecuteQuoteV5Req**](CreateRfqExecuteQuoteV5Req.md) | The request body for CreateRfqExecuteQuoteV5 | 

### Return type

[**CreateRfqExecuteQuoteV5Resp**](CreateRfqExecuteQuoteV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRfqMakerInstrumentSettingsV5

> CreateRfqMakerInstrumentSettingsV5Resp CreateRfqMakerInstrumentSettingsV5(ctx).CreateRfqMakerInstrumentSettingsV5Req(createRfqMakerInstrumentSettingsV5Req).Execute()

Customize the products which makers want to quote and receive RFQs for, and the corresponding price and size limit.   



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
	createRfqMakerInstrumentSettingsV5Req := *openapiclient.NewCreateRfqMakerInstrumentSettingsV5Req([]openapiclient.CreateRfqMakerInstrumentSettingsV5ReqDataInner{*openapiclient.NewCreateRfqMakerInstrumentSettingsV5ReqDataInner()}, "InstType_example") // CreateRfqMakerInstrumentSettingsV5Req | The request body for CreateRfqMakerInstrumentSettingsV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.CreateRfqMakerInstrumentSettingsV5(context.Background()).CreateRfqMakerInstrumentSettingsV5Req(createRfqMakerInstrumentSettingsV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.CreateRfqMakerInstrumentSettingsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRfqMakerInstrumentSettingsV5`: CreateRfqMakerInstrumentSettingsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.CreateRfqMakerInstrumentSettingsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRfqMakerInstrumentSettingsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRfqMakerInstrumentSettingsV5Req** | [**CreateRfqMakerInstrumentSettingsV5Req**](CreateRfqMakerInstrumentSettingsV5Req.md) | The request body for CreateRfqMakerInstrumentSettingsV5 | 

### Return type

[**CreateRfqMakerInstrumentSettingsV5Resp**](CreateRfqMakerInstrumentSettingsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRfqMmpConfigV5

> CreateRfqMmpConfigV5Resp CreateRfqMmpConfigV5(ctx).CreateRfqMmpConfigV5Req(createRfqMmpConfigV5Req).Execute()

This endpoint is used to set MMP configure and only applicable to block trading makers    



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
	createRfqMmpConfigV5Req := *openapiclient.NewCreateRfqMmpConfigV5Req("CountLimit_example", "FrozenInterval_example", "TimeInterval_example") // CreateRfqMmpConfigV5Req | The request body for CreateRfqMmpConfigV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.CreateRfqMmpConfigV5(context.Background()).CreateRfqMmpConfigV5Req(createRfqMmpConfigV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.CreateRfqMmpConfigV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRfqMmpConfigV5`: CreateRfqMmpConfigV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.CreateRfqMmpConfigV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRfqMmpConfigV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRfqMmpConfigV5Req** | [**CreateRfqMmpConfigV5Req**](CreateRfqMmpConfigV5Req.md) | The request body for CreateRfqMmpConfigV5 | 

### Return type

[**CreateRfqMmpConfigV5Resp**](CreateRfqMmpConfigV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRfqMmpResetV5

> CreateRfqMmpResetV5Resp CreateRfqMmpResetV5(ctx).Execute()

Reset the MMP status to be inactive.  



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
	resp, r, err := apiClient.BlockTradingAPI.CreateRfqMmpResetV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.CreateRfqMmpResetV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRfqMmpResetV5`: CreateRfqMmpResetV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.CreateRfqMmpResetV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRfqMmpResetV5Request struct via the builder pattern


### Return type

[**CreateRfqMmpResetV5Resp**](CreateRfqMmpResetV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketBlockTickerV5

> GetMarketBlockTickerV5Resp GetMarketBlockTickerV5(ctx).InstId(instId).Execute()

Retrieve the latest block trading volume in the last 24 hours.  



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
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USD-SWAP` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.GetMarketBlockTickerV5(context.Background()).InstId(instId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.GetMarketBlockTickerV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketBlockTickerV5`: GetMarketBlockTickerV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.GetMarketBlockTickerV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketBlockTickerV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USD-SWAP&#x60; | [default to &quot;&quot;]

### Return type

[**GetMarketBlockTickerV5Resp**](GetMarketBlockTickerV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketBlockTickersV5

> GetMarketBlockTickersV5Resp GetMarketBlockTickersV5(ctx).InstType(instType).Uly(uly).InstFamily(instFamily).Execute()

Retrieve the latest block trading volume in the last 24 hours.  



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
	instType := "instType_example" // string | Instrument type  `SPOT`  `SWAP`  `FUTURES`  `OPTION` (default to "")
	uly := "uly_example" // string | Underlying, e.g. `BTC-USD`   Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family, e.g. `BTC-USD`  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.GetMarketBlockTickersV5(context.Background()).InstType(instType).Uly(uly).InstFamily(instFamily).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.GetMarketBlockTickersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketBlockTickersV5`: GetMarketBlockTickersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.GetMarketBlockTickersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketBlockTickersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Underlying, e.g. &#x60;BTC-USD&#x60;   Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family, e.g. &#x60;BTC-USD&#x60;  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]

### Return type

[**GetMarketBlockTickersV5Resp**](GetMarketBlockTickersV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicBlockTradesV5

> GetPublicBlockTradesV5Resp GetPublicBlockTradesV5(ctx).InstId(instId).Execute()

Retrieve the recent block trading transactions of an instrument. Descending order by tradeId. The data will be updated 15 minutes after the block trade execution.  



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
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.GetPublicBlockTradesV5(context.Background()).InstId(instId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.GetPublicBlockTradesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicBlockTradesV5`: GetPublicBlockTradesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.GetPublicBlockTradesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicBlockTradesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]

### Return type

[**GetPublicBlockTradesV5Resp**](GetPublicBlockTradesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRfqCounterpartiesV5

> GetRfqCounterpartiesV5Resp GetRfqCounterpartiesV5(ctx).Execute()

Retrieves the list of counterparties that the user is permitted to trade with.   



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
	resp, r, err := apiClient.BlockTradingAPI.GetRfqCounterpartiesV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.GetRfqCounterpartiesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRfqCounterpartiesV5`: GetRfqCounterpartiesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.GetRfqCounterpartiesV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRfqCounterpartiesV5Request struct via the builder pattern


### Return type

[**GetRfqCounterpartiesV5Resp**](GetRfqCounterpartiesV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRfqMakerInstrumentSettingsV5

> GetRfqMakerInstrumentSettingsV5Resp GetRfqMakerInstrumentSettingsV5(ctx).Execute()

Retrieve the products which makers want to quote and receive RFQs for, and the corresponding price and size limit.   



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
	resp, r, err := apiClient.BlockTradingAPI.GetRfqMakerInstrumentSettingsV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.GetRfqMakerInstrumentSettingsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRfqMakerInstrumentSettingsV5`: GetRfqMakerInstrumentSettingsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.GetRfqMakerInstrumentSettingsV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRfqMakerInstrumentSettingsV5Request struct via the builder pattern


### Return type

[**GetRfqMakerInstrumentSettingsV5Resp**](GetRfqMakerInstrumentSettingsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRfqMmpConfigV5

> GetRfqMmpConfigV5Resp GetRfqMmpConfigV5(ctx).Execute()

This endpoint is used to get MMP configure information and only applicable to block trading market makers    



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
	resp, r, err := apiClient.BlockTradingAPI.GetRfqMmpConfigV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.GetRfqMmpConfigV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRfqMmpConfigV5`: GetRfqMmpConfigV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.GetRfqMmpConfigV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRfqMmpConfigV5Request struct via the builder pattern


### Return type

[**GetRfqMmpConfigV5Resp**](GetRfqMmpConfigV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRfqPublicTradesV5

> GetRfqPublicTradesV5Resp GetRfqPublicTradesV5(ctx).BeginId(beginId).EndId(endId).Limit(limit).Execute()

Retrieves the executed block trades. The data will be updated 15 minutes after the block trade execution.  



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
	beginId := "beginId_example" // string | The starting blockTdId the request to begin with. Pagination of data to return records newer than the requested `blockTdId`, not including beginId. (optional) (default to "")
	endId := "endId_example" // string | The last blockTdId the request to end with. Pagination of data to return records earlier than the requested `blockTdId`, not including endId. (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100 which is also the default value. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.GetRfqPublicTradesV5(context.Background()).BeginId(beginId).EndId(endId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.GetRfqPublicTradesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRfqPublicTradesV5`: GetRfqPublicTradesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.GetRfqPublicTradesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRfqPublicTradesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **beginId** | **string** | The starting blockTdId the request to begin with. Pagination of data to return records newer than the requested &#x60;blockTdId&#x60;, not including beginId. | [default to &quot;&quot;]
 **endId** | **string** | The last blockTdId the request to end with. Pagination of data to return records earlier than the requested &#x60;blockTdId&#x60;, not including endId. | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100 which is also the default value. | [default to &quot;&quot;]

### Return type

[**GetRfqPublicTradesV5Resp**](GetRfqPublicTradesV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRfqQuotesV5

> GetRfqQuotesV5Resp GetRfqQuotesV5(ctx).RfqId(rfqId).ClRfqId(clRfqId).QuoteId(quoteId).ClQuoteId(clQuoteId).State(state).BeginId(beginId).EndId(endId).Limit(limit).Execute()

Retrieve all Quotes that the user is a counterparty to (either as the creator or the receiver).  



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
	rfqId := "rfqId_example" // string | RFQ ID . (optional) (default to "")
	clRfqId := "clRfqId_example" // string | Client-supplied RFQ ID. If both `clRfqId` and `rfqId` are passed, `rfqId` will be be treated as primary identifier. (optional) (default to "")
	quoteId := "quoteId_example" // string | Quote ID (optional) (default to "")
	clQuoteId := "clQuoteId_example" // string | Client-supplied Quote ID. If both clQuoteId and quoteId are passed, quoteId will be treated as primary identifier (optional) (default to "")
	state := "state_example" // string | The status of the quote. Valid values can be `active` `canceled` `pending_fill` `filled` `expired` or `failed`. (optional) (default to "")
	beginId := "beginId_example" // string | Start quote id the request to begin with. Pagination of data to return records newer than the requested quoteId, not including beginId (optional) (default to "")
	endId := "endId_example" // string | End quote id the request to end with. Pagination of data to return records earlier than the requested quoteId, not including endId (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100 which is also the default value. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.GetRfqQuotesV5(context.Background()).RfqId(rfqId).ClRfqId(clRfqId).QuoteId(quoteId).ClQuoteId(clQuoteId).State(state).BeginId(beginId).EndId(endId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.GetRfqQuotesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRfqQuotesV5`: GetRfqQuotesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.GetRfqQuotesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRfqQuotesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rfqId** | **string** | RFQ ID . | [default to &quot;&quot;]
 **clRfqId** | **string** | Client-supplied RFQ ID. If both &#x60;clRfqId&#x60; and &#x60;rfqId&#x60; are passed, &#x60;rfqId&#x60; will be be treated as primary identifier. | [default to &quot;&quot;]
 **quoteId** | **string** | Quote ID | [default to &quot;&quot;]
 **clQuoteId** | **string** | Client-supplied Quote ID. If both clQuoteId and quoteId are passed, quoteId will be treated as primary identifier | [default to &quot;&quot;]
 **state** | **string** | The status of the quote. Valid values can be &#x60;active&#x60; &#x60;canceled&#x60; &#x60;pending_fill&#x60; &#x60;filled&#x60; &#x60;expired&#x60; or &#x60;failed&#x60;. | [default to &quot;&quot;]
 **beginId** | **string** | Start quote id the request to begin with. Pagination of data to return records newer than the requested quoteId, not including beginId | [default to &quot;&quot;]
 **endId** | **string** | End quote id the request to end with. Pagination of data to return records earlier than the requested quoteId, not including endId | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100 which is also the default value. | [default to &quot;&quot;]

### Return type

[**GetRfqQuotesV5Resp**](GetRfqQuotesV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRfqRfqsV5

> GetRfqRfqsV5Resp GetRfqRfqsV5(ctx).RfqId(rfqId).ClRfqId(clRfqId).State(state).BeginId(beginId).EndId(endId).Limit(limit).Execute()

Retrieves details of RFQs that the user is a counterparty to (either as the creator or the receiver of the RFQ).   



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
	rfqId := "rfqId_example" // string | RFQ ID . (optional) (default to "")
	clRfqId := "clRfqId_example" // string | Client-supplied RFQ ID. If both `clRfqId` and `rfqId` are passed, `rfqId` will be treated as primary identifier (optional) (default to "")
	state := "state_example" // string | The status of the RFQ.   Valid values can be `active` `canceled` `pending_fill` `filled` `expired` `failed` `traded_away`.   `traded_away` only applies to Maker (optional) (default to "")
	beginId := "beginId_example" // string | Start rfq id the request to begin with. Pagination of data to return records newer than the requested rfqId, not including beginId (optional) (default to "")
	endId := "endId_example" // string | End rfq id the request to end with. Pagination of data to return records earlier than the requested rfqId, not including endId (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100 which is also the default value. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.GetRfqRfqsV5(context.Background()).RfqId(rfqId).ClRfqId(clRfqId).State(state).BeginId(beginId).EndId(endId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.GetRfqRfqsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRfqRfqsV5`: GetRfqRfqsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.GetRfqRfqsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRfqRfqsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rfqId** | **string** | RFQ ID . | [default to &quot;&quot;]
 **clRfqId** | **string** | Client-supplied RFQ ID. If both &#x60;clRfqId&#x60; and &#x60;rfqId&#x60; are passed, &#x60;rfqId&#x60; will be treated as primary identifier | [default to &quot;&quot;]
 **state** | **string** | The status of the RFQ.   Valid values can be &#x60;active&#x60; &#x60;canceled&#x60; &#x60;pending_fill&#x60; &#x60;filled&#x60; &#x60;expired&#x60; &#x60;failed&#x60; &#x60;traded_away&#x60;.   &#x60;traded_away&#x60; only applies to Maker | [default to &quot;&quot;]
 **beginId** | **string** | Start rfq id the request to begin with. Pagination of data to return records newer than the requested rfqId, not including beginId | [default to &quot;&quot;]
 **endId** | **string** | End rfq id the request to end with. Pagination of data to return records earlier than the requested rfqId, not including endId | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100 which is also the default value. | [default to &quot;&quot;]

### Return type

[**GetRfqRfqsV5Resp**](GetRfqRfqsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRfqTradesV5

> GetRfqTradesV5Resp GetRfqTradesV5(ctx).RfqId(rfqId).ClRfqId(clRfqId).QuoteId(quoteId).BlockTdId(blockTdId).ClQuoteId(clQuoteId).BeginId(beginId).EndId(endId).BeginTs(beginTs).EndTs(endTs).Limit(limit).IsSuccessful(isSuccessful).Execute()

Retrieves the executed trades that the user is a counterparty to (either as the creator or the receiver).  



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
	rfqId := "rfqId_example" // string | RFQ ID . (optional) (default to "")
	clRfqId := "clRfqId_example" // string | Client-supplied RFQ ID. If both `clRfqId` and `rfqId` are passed, `rfqId` will be treated as primary identifier (optional) (default to "")
	quoteId := "quoteId_example" // string | Quote ID (optional) (default to "")
	blockTdId := "blockTdId_example" // string | Block trade ID (optional) (default to "")
	clQuoteId := "clQuoteId_example" // string | Client-supplied Quote ID. If both `clQuoteId` and `quoteId` are passed, `quoteId` will be treated as primary identifier (optional) (default to "")
	beginId := "beginId_example" // string | The starting rfq id the request to begin with. Pagination of data to return records newer than the requested blockTdId, not including beginId. (optional) (default to "")
	endId := "endId_example" // string | The last rfq id the request to end withPagination of data to return records earlier than the requested blockTdId, not including endId. (optional) (default to "")
	beginTs := "beginTs_example" // string | Filter trade execution time with a begin timestamp (UTC timezone). Unix timestamp format in milliseconds, e.g. 1597026383085 (optional) (default to "")
	endTs := "endTs_example" // string | Filter trade execution time with an end timestamp (UTC timezone). Unix timestamp format in milliseconds, e.g. 1597026383085 (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100 which is also the default value.   If the number of trades in the requested range is bigger than 100, the latest 100 trades in the range will be returned. (optional) (default to "")
	isSuccessful := true // bool | Whether the trade is filled successfully.  `true`: the default value. `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockTradingAPI.GetRfqTradesV5(context.Background()).RfqId(rfqId).ClRfqId(clRfqId).QuoteId(quoteId).BlockTdId(blockTdId).ClQuoteId(clQuoteId).BeginId(beginId).EndId(endId).BeginTs(beginTs).EndTs(endTs).Limit(limit).IsSuccessful(isSuccessful).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockTradingAPI.GetRfqTradesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRfqTradesV5`: GetRfqTradesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `BlockTradingAPI.GetRfqTradesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRfqTradesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rfqId** | **string** | RFQ ID . | [default to &quot;&quot;]
 **clRfqId** | **string** | Client-supplied RFQ ID. If both &#x60;clRfqId&#x60; and &#x60;rfqId&#x60; are passed, &#x60;rfqId&#x60; will be treated as primary identifier | [default to &quot;&quot;]
 **quoteId** | **string** | Quote ID | [default to &quot;&quot;]
 **blockTdId** | **string** | Block trade ID | [default to &quot;&quot;]
 **clQuoteId** | **string** | Client-supplied Quote ID. If both &#x60;clQuoteId&#x60; and &#x60;quoteId&#x60; are passed, &#x60;quoteId&#x60; will be treated as primary identifier | [default to &quot;&quot;]
 **beginId** | **string** | The starting rfq id the request to begin with. Pagination of data to return records newer than the requested blockTdId, not including beginId. | [default to &quot;&quot;]
 **endId** | **string** | The last rfq id the request to end withPagination of data to return records earlier than the requested blockTdId, not including endId. | [default to &quot;&quot;]
 **beginTs** | **string** | Filter trade execution time with a begin timestamp (UTC timezone). Unix timestamp format in milliseconds, e.g. 1597026383085 | [default to &quot;&quot;]
 **endTs** | **string** | Filter trade execution time with an end timestamp (UTC timezone). Unix timestamp format in milliseconds, e.g. 1597026383085 | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100 which is also the default value.   If the number of trades in the requested range is bigger than 100, the latest 100 trades in the range will be returned. | [default to &quot;&quot;]
 **isSuccessful** | **bool** | Whether the trade is filled successfully.  &#x60;true&#x60;: the default value. &#x60;false&#x60;. | 

### Return type

[**GetRfqTradesV5Resp**](GetRfqTradesV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

