# \FundingAccountAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssetCancelWithdrawalV5**](FundingAccountAPI.md#CreateAssetCancelWithdrawalV5) | **Post** /api/v5/asset/cancel-withdrawal | You can cancel normal withdrawal requests, but you cannot cancel withdrawal requests on Lightning.  
[**CreateAssetConvertEstimateQuoteV5**](FundingAccountAPI.md#CreateAssetConvertEstimateQuoteV5) | **Post** /api/v5/asset/convert/estimate-quote | 
[**CreateAssetConvertTradeV5**](FundingAccountAPI.md#CreateAssetConvertTradeV5) | **Post** /api/v5/asset/convert/trade | You should make  before convert trade.   For the same side (buy/sell), there&#39;s a trading limit of 1 request per 5 seconds.  
[**CreateAssetMonthlyStatementV5**](FundingAccountAPI.md#CreateAssetMonthlyStatementV5) | **Post** /api/v5/asset/monthly-statement | Apply for monthly statement in the past year.  
[**CreateAssetTransferV5**](FundingAccountAPI.md#CreateAssetTransferV5) | **Post** /api/v5/asset/transfer | Only API keys with &#x60;Trade&#x60; privilege can call this endpoint.  This endpoint supports the transfer of funds between your funding account and trading account, and from the master account to sub-accounts.  Sub-account can transfer out to master account by default. Need to call  to grant privilege first if you want sub-account transferring to another sub-account (sub-accounts need to belong to same master account.)  
[**CreateAssetWithdrawalV5**](FundingAccountAPI.md#CreateAssetWithdrawalV5) | **Post** /api/v5/asset/withdrawal | Only supported withdrawal of assets from funding account. Common sub-account does not support withdrawal.   
[**CreateFiatCancelWithdrawalV5**](FundingAccountAPI.md#CreateFiatCancelWithdrawalV5) | **Post** /api/v5/fiat/cancel-withdrawal | Cancel a pending fiat withdrawal order, currently only applicable to TRY  
[**CreateFiatCreateWithdrawalV5**](FundingAccountAPI.md#CreateFiatCreateWithdrawalV5) | **Post** /api/v5/fiat/create-withdrawal | Initiate a fiat withdrawal request (Authenticated endpoint, Only for API keys with \&quot;Withdrawal\&quot; access)   Only supported withdrawal of assets from funding account.  
[**GetAssetAssetValuationV5**](FundingAccountAPI.md#GetAssetAssetValuationV5) | **Get** /api/v5/asset/asset-valuation | View account asset valuation  
[**GetAssetBalancesV5**](FundingAccountAPI.md#GetAssetBalancesV5) | **Get** /api/v5/asset/balances | Retrieve the funding account balances of all the assets and the amount that is available or on hold.  
[**GetAssetBillsV5**](FundingAccountAPI.md#GetAssetBillsV5) | **Get** /api/v5/asset/bills | Query the billing record in the past month.  
[**GetAssetConvertCurrenciesV5**](FundingAccountAPI.md#GetAssetConvertCurrenciesV5) | **Get** /api/v5/asset/convert/currencies | 
[**GetAssetConvertCurrencyPairV5**](FundingAccountAPI.md#GetAssetConvertCurrencyPairV5) | **Get** /api/v5/asset/convert/currency-pair | 
[**GetAssetConvertHistoryV5**](FundingAccountAPI.md#GetAssetConvertHistoryV5) | **Get** /api/v5/asset/convert/history | 
[**GetAssetCurrenciesV5**](FundingAccountAPI.md#GetAssetCurrenciesV5) | **Get** /api/v5/asset/currencies | Retrieve a list of all currencies available which are related to the current account&#39;s KYC entity.  
[**GetAssetDepositAddressV5**](FundingAccountAPI.md#GetAssetDepositAddressV5) | **Get** /api/v5/asset/deposit-address | Retrieve the deposit addresses of currencies, including previously-used addresses.  
[**GetAssetDepositHistoryV5**](FundingAccountAPI.md#GetAssetDepositHistoryV5) | **Get** /api/v5/asset/deposit-history | Retrieve the deposit records according to the currency, deposit status, and time range in reverse chronological order. The 100 most recent records are returned by default.   Websocket API is also available, refer to .  
[**GetAssetDepositWithdrawStatusV5**](FundingAccountAPI.md#GetAssetDepositWithdrawStatusV5) | **Get** /api/v5/asset/deposit-withdraw-status | Retrieve deposit&#39;s and withdrawal&#39;s detailed status and estimated complete time.  
[**GetAssetExchangeListV5**](FundingAccountAPI.md#GetAssetExchangeListV5) | **Get** /api/v5/asset/exchange-list | Authentication is not required for this public endpoint.  
[**GetAssetMonthlyStatementV5**](FundingAccountAPI.md#GetAssetMonthlyStatementV5) | **Get** /api/v5/asset/monthly-statement | Retrieve monthly statement in the past year.  
[**GetAssetNonTradableAssetsV5**](FundingAccountAPI.md#GetAssetNonTradableAssetsV5) | **Get** /api/v5/asset/non-tradable-assets | 
[**GetAssetTransferStateV5**](FundingAccountAPI.md#GetAssetTransferStateV5) | **Get** /api/v5/asset/transfer-state | Retrieve the transfer state data of the last 2 weeks.  
[**GetAssetWithdrawalHistoryV5**](FundingAccountAPI.md#GetAssetWithdrawalHistoryV5) | **Get** /api/v5/asset/withdrawal-history | Retrieve the withdrawal records according to the currency, withdrawal status, and time range in reverse chronological order. The 100 most recent records are returned by default.   Websocket API is also available, refer to .  
[**GetFiatDepositOrderHistoryV5**](FundingAccountAPI.md#GetFiatDepositOrderHistoryV5) | **Get** /api/v5/fiat/deposit-order-history | Get fiat deposit order history  
[**GetFiatDepositPaymentMethodsV5**](FundingAccountAPI.md#GetFiatDepositPaymentMethodsV5) | **Get** /api/v5/fiat/deposit-payment-methods | To display all the available fiat deposit payment methods  
[**GetFiatDepositV5**](FundingAccountAPI.md#GetFiatDepositV5) | **Get** /api/v5/fiat/deposit | Get fiat deposit order detail  
[**GetFiatWithdrawalOrderHistoryV5**](FundingAccountAPI.md#GetFiatWithdrawalOrderHistoryV5) | **Get** /api/v5/fiat/withdrawal-order-history | Get fiat withdrawal order history  
[**GetFiatWithdrawalPaymentMethodsV5**](FundingAccountAPI.md#GetFiatWithdrawalPaymentMethodsV5) | **Get** /api/v5/fiat/withdrawal-payment-methods | To display all the available fiat withdrawal payment methods  
[**GetFiatWithdrawalV5**](FundingAccountAPI.md#GetFiatWithdrawalV5) | **Get** /api/v5/fiat/withdrawal | Get fiat withdraw order detail  



## CreateAssetCancelWithdrawalV5

> CreateAssetCancelWithdrawalV5Resp CreateAssetCancelWithdrawalV5(ctx).CreateAssetCancelWithdrawalV5Req(createAssetCancelWithdrawalV5Req).Execute()

You can cancel normal withdrawal requests, but you cannot cancel withdrawal requests on Lightning.  



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
	createAssetCancelWithdrawalV5Req := *openapiclient.NewCreateAssetCancelWithdrawalV5Req("WdId_example") // CreateAssetCancelWithdrawalV5Req | The request body for CreateAssetCancelWithdrawalV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.CreateAssetCancelWithdrawalV5(context.Background()).CreateAssetCancelWithdrawalV5Req(createAssetCancelWithdrawalV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.CreateAssetCancelWithdrawalV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetCancelWithdrawalV5`: CreateAssetCancelWithdrawalV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.CreateAssetCancelWithdrawalV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetCancelWithdrawalV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAssetCancelWithdrawalV5Req** | [**CreateAssetCancelWithdrawalV5Req**](CreateAssetCancelWithdrawalV5Req.md) | The request body for CreateAssetCancelWithdrawalV5 | 

### Return type

[**CreateAssetCancelWithdrawalV5Resp**](CreateAssetCancelWithdrawalV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetConvertEstimateQuoteV5

> CreateAssetConvertEstimateQuoteV5Resp CreateAssetConvertEstimateQuoteV5(ctx).CreateAssetConvertEstimateQuoteV5Req(createAssetConvertEstimateQuoteV5Req).Execute()





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
	createAssetConvertEstimateQuoteV5Req := *openapiclient.NewCreateAssetConvertEstimateQuoteV5Req("BaseCcy_example", "QuoteCcy_example", "RfqSz_example", "RfqSzCcy_example", "Side_example") // CreateAssetConvertEstimateQuoteV5Req | The request body for CreateAssetConvertEstimateQuoteV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.CreateAssetConvertEstimateQuoteV5(context.Background()).CreateAssetConvertEstimateQuoteV5Req(createAssetConvertEstimateQuoteV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.CreateAssetConvertEstimateQuoteV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetConvertEstimateQuoteV5`: CreateAssetConvertEstimateQuoteV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.CreateAssetConvertEstimateQuoteV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetConvertEstimateQuoteV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAssetConvertEstimateQuoteV5Req** | [**CreateAssetConvertEstimateQuoteV5Req**](CreateAssetConvertEstimateQuoteV5Req.md) | The request body for CreateAssetConvertEstimateQuoteV5 | 

### Return type

[**CreateAssetConvertEstimateQuoteV5Resp**](CreateAssetConvertEstimateQuoteV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetConvertTradeV5

> CreateAssetConvertTradeV5Resp CreateAssetConvertTradeV5(ctx).CreateAssetConvertTradeV5Req(createAssetConvertTradeV5Req).Execute()

You should make  before convert trade.   For the same side (buy/sell), there's a trading limit of 1 request per 5 seconds.  



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
	createAssetConvertTradeV5Req := *openapiclient.NewCreateAssetConvertTradeV5Req("BaseCcy_example", "QuoteCcy_example", "QuoteId_example", "Side_example", "Sz_example", "SzCcy_example") // CreateAssetConvertTradeV5Req | The request body for CreateAssetConvertTradeV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.CreateAssetConvertTradeV5(context.Background()).CreateAssetConvertTradeV5Req(createAssetConvertTradeV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.CreateAssetConvertTradeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetConvertTradeV5`: CreateAssetConvertTradeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.CreateAssetConvertTradeV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetConvertTradeV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAssetConvertTradeV5Req** | [**CreateAssetConvertTradeV5Req**](CreateAssetConvertTradeV5Req.md) | The request body for CreateAssetConvertTradeV5 | 

### Return type

[**CreateAssetConvertTradeV5Resp**](CreateAssetConvertTradeV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetMonthlyStatementV5

> CreateAssetMonthlyStatementV5Resp CreateAssetMonthlyStatementV5(ctx).CreateAssetMonthlyStatementV5Req(createAssetMonthlyStatementV5Req).Execute()

Apply for monthly statement in the past year.  



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
	createAssetMonthlyStatementV5Req := *openapiclient.NewCreateAssetMonthlyStatementV5Req() // CreateAssetMonthlyStatementV5Req | The request body for CreateAssetMonthlyStatementV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.CreateAssetMonthlyStatementV5(context.Background()).CreateAssetMonthlyStatementV5Req(createAssetMonthlyStatementV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.CreateAssetMonthlyStatementV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetMonthlyStatementV5`: CreateAssetMonthlyStatementV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.CreateAssetMonthlyStatementV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetMonthlyStatementV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAssetMonthlyStatementV5Req** | [**CreateAssetMonthlyStatementV5Req**](CreateAssetMonthlyStatementV5Req.md) | The request body for CreateAssetMonthlyStatementV5 | 

### Return type

[**CreateAssetMonthlyStatementV5Resp**](CreateAssetMonthlyStatementV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetTransferV5

> CreateAssetTransferV5Resp CreateAssetTransferV5(ctx).CreateAssetTransferV5Req(createAssetTransferV5Req).Execute()

Only API keys with `Trade` privilege can call this endpoint.  This endpoint supports the transfer of funds between your funding account and trading account, and from the master account to sub-accounts.  Sub-account can transfer out to master account by default. Need to call  to grant privilege first if you want sub-account transferring to another sub-account (sub-accounts need to belong to same master account.)  



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
	createAssetTransferV5Req := *openapiclient.NewCreateAssetTransferV5Req("Amt_example", "Ccy_example", "From_example", "To_example") // CreateAssetTransferV5Req | The request body for CreateAssetTransferV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.CreateAssetTransferV5(context.Background()).CreateAssetTransferV5Req(createAssetTransferV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.CreateAssetTransferV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetTransferV5`: CreateAssetTransferV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.CreateAssetTransferV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetTransferV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAssetTransferV5Req** | [**CreateAssetTransferV5Req**](CreateAssetTransferV5Req.md) | The request body for CreateAssetTransferV5 | 

### Return type

[**CreateAssetTransferV5Resp**](CreateAssetTransferV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssetWithdrawalV5

> CreateAssetWithdrawalV5Resp CreateAssetWithdrawalV5(ctx).CreateAssetWithdrawalV5Req(createAssetWithdrawalV5Req).Execute()

Only supported withdrawal of assets from funding account. Common sub-account does not support withdrawal.   



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
	createAssetWithdrawalV5Req := *openapiclient.NewCreateAssetWithdrawalV5Req("Amt_example", "Ccy_example", "Dest_example", "ToAddr_example") // CreateAssetWithdrawalV5Req | The request body for CreateAssetWithdrawalV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.CreateAssetWithdrawalV5(context.Background()).CreateAssetWithdrawalV5Req(createAssetWithdrawalV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.CreateAssetWithdrawalV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetWithdrawalV5`: CreateAssetWithdrawalV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.CreateAssetWithdrawalV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetWithdrawalV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAssetWithdrawalV5Req** | [**CreateAssetWithdrawalV5Req**](CreateAssetWithdrawalV5Req.md) | The request body for CreateAssetWithdrawalV5 | 

### Return type

[**CreateAssetWithdrawalV5Resp**](CreateAssetWithdrawalV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFiatCancelWithdrawalV5

> CreateFiatCancelWithdrawalV5Resp CreateFiatCancelWithdrawalV5(ctx).CreateFiatCancelWithdrawalV5Req(createFiatCancelWithdrawalV5Req).Execute()

Cancel a pending fiat withdrawal order, currently only applicable to TRY  



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
	createFiatCancelWithdrawalV5Req := *openapiclient.NewCreateFiatCancelWithdrawalV5Req("OrdId_example") // CreateFiatCancelWithdrawalV5Req | The request body for CreateFiatCancelWithdrawalV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.CreateFiatCancelWithdrawalV5(context.Background()).CreateFiatCancelWithdrawalV5Req(createFiatCancelWithdrawalV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.CreateFiatCancelWithdrawalV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFiatCancelWithdrawalV5`: CreateFiatCancelWithdrawalV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.CreateFiatCancelWithdrawalV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFiatCancelWithdrawalV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFiatCancelWithdrawalV5Req** | [**CreateFiatCancelWithdrawalV5Req**](CreateFiatCancelWithdrawalV5Req.md) | The request body for CreateFiatCancelWithdrawalV5 | 

### Return type

[**CreateFiatCancelWithdrawalV5Resp**](CreateFiatCancelWithdrawalV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFiatCreateWithdrawalV5

> CreateFiatCreateWithdrawalV5Resp CreateFiatCreateWithdrawalV5(ctx).CreateFiatCreateWithdrawalV5Req(createFiatCreateWithdrawalV5Req).Execute()

Initiate a fiat withdrawal request (Authenticated endpoint, Only for API keys with \"Withdrawal\" access)   Only supported withdrawal of assets from funding account.  



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
	createFiatCreateWithdrawalV5Req := *openapiclient.NewCreateFiatCreateWithdrawalV5Req("Amt_example", "Ccy_example", "ClientId_example", "PaymentAcctId_example", "PaymentMethod_example") // CreateFiatCreateWithdrawalV5Req | The request body for CreateFiatCreateWithdrawalV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.CreateFiatCreateWithdrawalV5(context.Background()).CreateFiatCreateWithdrawalV5Req(createFiatCreateWithdrawalV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.CreateFiatCreateWithdrawalV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFiatCreateWithdrawalV5`: CreateFiatCreateWithdrawalV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.CreateFiatCreateWithdrawalV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFiatCreateWithdrawalV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFiatCreateWithdrawalV5Req** | [**CreateFiatCreateWithdrawalV5Req**](CreateFiatCreateWithdrawalV5Req.md) | The request body for CreateFiatCreateWithdrawalV5 | 

### Return type

[**CreateFiatCreateWithdrawalV5Resp**](CreateFiatCreateWithdrawalV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetAssetValuationV5

> GetAssetAssetValuationV5Resp GetAssetAssetValuationV5(ctx).Ccy(ccy).Execute()

View account asset valuation  



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
	ccy := "ccy_example" // string | Asset valuation calculation unit   BTC, USDT  USD, CNY, JP, KRW, RUB, EUR  VND, IDR, INR, PHP, THB, TRY   AUD, SGD, ARS, SAR, AED, IQD   The default is the valuation in BTC. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetAssetAssetValuationV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetAssetValuationV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetAssetValuationV5`: GetAssetAssetValuationV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetAssetValuationV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetAssetValuationV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Asset valuation calculation unit   BTC, USDT  USD, CNY, JP, KRW, RUB, EUR  VND, IDR, INR, PHP, THB, TRY   AUD, SGD, ARS, SAR, AED, IQD   The default is the valuation in BTC. | [default to &quot;&quot;]

### Return type

[**GetAssetAssetValuationV5Resp**](GetAssetAssetValuationV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetBalancesV5

> GetAssetBalancesV5Resp GetAssetBalancesV5(ctx).Ccy(ccy).Execute()

Retrieve the funding account balances of all the assets and the amount that is available or on hold.  



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
	ccy := "ccy_example" // string | Single currency or multiple currencies (no more than 20) separated with comma, e.g. `BTC` or `BTC,ETH`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetAssetBalancesV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetBalancesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetBalancesV5`: GetAssetBalancesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetBalancesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetBalancesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Single currency or multiple currencies (no more than 20) separated with comma, e.g. &#x60;BTC&#x60; or &#x60;BTC,ETH&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAssetBalancesV5Resp**](GetAssetBalancesV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetBillsV5

> GetAssetBillsV5Resp GetAssetBillsV5(ctx).Ccy(ccy).Type_(type_).ClientId(clientId).After(after).Before(before).Limit(limit).Execute()

Query the billing record in the past month.  



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
	ccy := "ccy_example" // string | Currency (optional) (default to "")
	type_ := "type__example" // string | Bill type  `1`: Deposit  `2`: Withdrawal  `13`: Canceled withdrawal  `20`: Transfer to sub account (for master account)  `21`: Transfer from sub account (for master account)  `22`: Transfer out from sub to master account (for sub-account)  `23`: Transfer in from master to sub account (for sub-account)  `28`: Manually claimed Airdrop  `47`: System reversal  `48`: Event Reward  `49`: Event Giveaway  `68`: Fee rebate (by rebate card)  `72`: Token received  `73`: Token given away  `74`: Token refunded  `75`: [Simple earn flexible] Subscription  `76`: [Simple earn flexible] Redemption  `77`: Jumpstart distribute  `78`: Jumpstart lock up  `80`: DEFI/Staking subscription  `82`: DEFI/Staking redemption  `83`: Staking yield  `84`: Violation fee  `89`: Deposit yield  `116`: [Fiat] Place an order  `117`: [Fiat] Fulfill an order  `118`: [Fiat] Cancel an order  `124`: Jumpstart unlocking  `130`: Transferred from Trading account  `131`: Transferred to Trading account  `132`: [P2P] Frozen by customer service  `133`: [P2P] Unfrozen by customer service  `134`: [P2P] Transferred by customer service  `135`: Cross chain exchange  `137`: [ETH Staking] Subscription  `138`: [ETH Staking] Swapping  `139`: [ETH Staking] Earnings  `146`: Customer feedback  `150`: Affiliate commission  `151`: Referral reward  `152`: Broker reward  `160`: Dual Investment subscribe  `161`: Dual Investment collection  `162`: Dual Investment profit  `163`: Dual Investment refund  `172`: [Affiliate] Sub-affiliate commission  `173`: [Affiliate] Fee rebate (by trading fee)  `174`: Jumpstart Pay  `175`: Locked collateral  `176`: Loan  `177`: Added collateral  `178`: Returned collateral  `179`: Repayment  `180`: Unlocked collateral  `181`: Airdrop payment  `185`: [Broker] Convert reward  `187`: [Broker] Convert transfer  `189`: Mystery box bonus  `195`: Untradable asset withdrawal  `196`: Untradable asset withdrawal revoked  `197`: Untradable asset deposit  `198`: Untradable asset collection reduce  `199`: Untradable asset collection increase  `200`: Buy  `202`: Price Lock Subscribe  `203`: Price Lock Collection  `204`: Price Lock Profit  `205`: Price Lock Refund  `207`: Dual Investment Lite Subscribe  `208`: Dual Investment Lite Collection  `209`: Dual Investment Lite Profit  `210`: Dual Investment Lite Refund  `212`: [Flexible loan] Multi-collateral loan collateral locked  `215`: [Flexible loan] Multi-collateral loan collateral released  `217`: [Flexible loan] Multi-collateral loan borrowed  `218`: [Flexible loan] Multi-collateral loan repaid  `232`: [Flexible loan] Subsidized interest received  `220`: Delisted crypto  `221`: Blockchain's withdrawal fee  `222`: Withdrawal fee refund  `223`: SWAP lead trading profit share  `225`: Shark Fin subscribe  `226`: Shark Fin collection  `227`: Shark Fin profit  `228`: Shark Fin refund  `229`: Airdrop  `232`: Subsidized interest received  `233`: Broker rebate compensation  `240`: Snowball subscribe  `241`: Snowball refund  `242`: Snowball profit  `243`: Snowball trading failed  `249`: Seagull subscribe  `250`: Seagull collection  `251`: Seagull profit  `252`: Seagull refund  `263`: Strategy bots profit share  `265`: Signal revenue  `266`: SPOT lead trading profit share  `270`: DCD broker transfer  `271`: DCD broker rebate  `272`: [Convert] Buy Crypto/Fiat  `273`: [Convert] Sell Crypto/Fiat  `284`: [Custody] Transfer out trading sub-account  `285`: [Custody] Transfer in trading sub-account  `286`: [Custody] Transfer out custody funding account  `287`: [Custody] Transfer in custody funding account  `288`: [Custody] Fund delegation   `289`: [Custody] Fund undelegation  `299`: Affiliate recommendation commission  `300`: Fee discount rebate  `303`: Snowball market maker transfer  `304`: [Simple Earn Fixed] Order submission  `305`: [Simple Earn Fixed] Order redemption  `306`: [Simple Earn Fixed] Principal distribution  `307`: [Simple Earn Fixed] Interest distribution (early termination compensation)  `308`: [Simple Earn Fixed] Interest distribution  `309`: [Simple Earn Fixed] Interest distribution (extension compensation)   `311`: Crypto dust auto-transfer in  `313`: Sent by gift  `314`: Received from gift  `315`: Refunded from gift  `328`: [SOL staking] Send Liquidity Staking Token reward  `329`: [SOL staking] Subscribe Liquidity Staking Token staking  `330`: [SOL staking] Mint Liquidity Staking Token  `331`: [SOL staking] Redeem Liquidity Staking Token order  `332`: [SOL staking] Settle Liquidity Staking Token order  `333`: Trial fund reward  `336`: [Credit line] Loan Forced Repayment  `338`: [Credit line] Forced Repayment Refund  `354`: Copy and bot rewards  `361`: Deposit from closed sub-account (optional) (default to "")
	clientId := "clientId_example" // string | Client-supplied ID for transfer or withdrawal  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ts`, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ts`, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetAssetBillsV5(context.Background()).Ccy(ccy).Type_(type_).ClientId(clientId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetBillsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetBillsV5`: GetAssetBillsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetBillsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetBillsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency | [default to &quot;&quot;]
 **type_** | **string** | Bill type  &#x60;1&#x60;: Deposit  &#x60;2&#x60;: Withdrawal  &#x60;13&#x60;: Canceled withdrawal  &#x60;20&#x60;: Transfer to sub account (for master account)  &#x60;21&#x60;: Transfer from sub account (for master account)  &#x60;22&#x60;: Transfer out from sub to master account (for sub-account)  &#x60;23&#x60;: Transfer in from master to sub account (for sub-account)  &#x60;28&#x60;: Manually claimed Airdrop  &#x60;47&#x60;: System reversal  &#x60;48&#x60;: Event Reward  &#x60;49&#x60;: Event Giveaway  &#x60;68&#x60;: Fee rebate (by rebate card)  &#x60;72&#x60;: Token received  &#x60;73&#x60;: Token given away  &#x60;74&#x60;: Token refunded  &#x60;75&#x60;: [Simple earn flexible] Subscription  &#x60;76&#x60;: [Simple earn flexible] Redemption  &#x60;77&#x60;: Jumpstart distribute  &#x60;78&#x60;: Jumpstart lock up  &#x60;80&#x60;: DEFI/Staking subscription  &#x60;82&#x60;: DEFI/Staking redemption  &#x60;83&#x60;: Staking yield  &#x60;84&#x60;: Violation fee  &#x60;89&#x60;: Deposit yield  &#x60;116&#x60;: [Fiat] Place an order  &#x60;117&#x60;: [Fiat] Fulfill an order  &#x60;118&#x60;: [Fiat] Cancel an order  &#x60;124&#x60;: Jumpstart unlocking  &#x60;130&#x60;: Transferred from Trading account  &#x60;131&#x60;: Transferred to Trading account  &#x60;132&#x60;: [P2P] Frozen by customer service  &#x60;133&#x60;: [P2P] Unfrozen by customer service  &#x60;134&#x60;: [P2P] Transferred by customer service  &#x60;135&#x60;: Cross chain exchange  &#x60;137&#x60;: [ETH Staking] Subscription  &#x60;138&#x60;: [ETH Staking] Swapping  &#x60;139&#x60;: [ETH Staking] Earnings  &#x60;146&#x60;: Customer feedback  &#x60;150&#x60;: Affiliate commission  &#x60;151&#x60;: Referral reward  &#x60;152&#x60;: Broker reward  &#x60;160&#x60;: Dual Investment subscribe  &#x60;161&#x60;: Dual Investment collection  &#x60;162&#x60;: Dual Investment profit  &#x60;163&#x60;: Dual Investment refund  &#x60;172&#x60;: [Affiliate] Sub-affiliate commission  &#x60;173&#x60;: [Affiliate] Fee rebate (by trading fee)  &#x60;174&#x60;: Jumpstart Pay  &#x60;175&#x60;: Locked collateral  &#x60;176&#x60;: Loan  &#x60;177&#x60;: Added collateral  &#x60;178&#x60;: Returned collateral  &#x60;179&#x60;: Repayment  &#x60;180&#x60;: Unlocked collateral  &#x60;181&#x60;: Airdrop payment  &#x60;185&#x60;: [Broker] Convert reward  &#x60;187&#x60;: [Broker] Convert transfer  &#x60;189&#x60;: Mystery box bonus  &#x60;195&#x60;: Untradable asset withdrawal  &#x60;196&#x60;: Untradable asset withdrawal revoked  &#x60;197&#x60;: Untradable asset deposit  &#x60;198&#x60;: Untradable asset collection reduce  &#x60;199&#x60;: Untradable asset collection increase  &#x60;200&#x60;: Buy  &#x60;202&#x60;: Price Lock Subscribe  &#x60;203&#x60;: Price Lock Collection  &#x60;204&#x60;: Price Lock Profit  &#x60;205&#x60;: Price Lock Refund  &#x60;207&#x60;: Dual Investment Lite Subscribe  &#x60;208&#x60;: Dual Investment Lite Collection  &#x60;209&#x60;: Dual Investment Lite Profit  &#x60;210&#x60;: Dual Investment Lite Refund  &#x60;212&#x60;: [Flexible loan] Multi-collateral loan collateral locked  &#x60;215&#x60;: [Flexible loan] Multi-collateral loan collateral released  &#x60;217&#x60;: [Flexible loan] Multi-collateral loan borrowed  &#x60;218&#x60;: [Flexible loan] Multi-collateral loan repaid  &#x60;232&#x60;: [Flexible loan] Subsidized interest received  &#x60;220&#x60;: Delisted crypto  &#x60;221&#x60;: Blockchain&#39;s withdrawal fee  &#x60;222&#x60;: Withdrawal fee refund  &#x60;223&#x60;: SWAP lead trading profit share  &#x60;225&#x60;: Shark Fin subscribe  &#x60;226&#x60;: Shark Fin collection  &#x60;227&#x60;: Shark Fin profit  &#x60;228&#x60;: Shark Fin refund  &#x60;229&#x60;: Airdrop  &#x60;232&#x60;: Subsidized interest received  &#x60;233&#x60;: Broker rebate compensation  &#x60;240&#x60;: Snowball subscribe  &#x60;241&#x60;: Snowball refund  &#x60;242&#x60;: Snowball profit  &#x60;243&#x60;: Snowball trading failed  &#x60;249&#x60;: Seagull subscribe  &#x60;250&#x60;: Seagull collection  &#x60;251&#x60;: Seagull profit  &#x60;252&#x60;: Seagull refund  &#x60;263&#x60;: Strategy bots profit share  &#x60;265&#x60;: Signal revenue  &#x60;266&#x60;: SPOT lead trading profit share  &#x60;270&#x60;: DCD broker transfer  &#x60;271&#x60;: DCD broker rebate  &#x60;272&#x60;: [Convert] Buy Crypto/Fiat  &#x60;273&#x60;: [Convert] Sell Crypto/Fiat  &#x60;284&#x60;: [Custody] Transfer out trading sub-account  &#x60;285&#x60;: [Custody] Transfer in trading sub-account  &#x60;286&#x60;: [Custody] Transfer out custody funding account  &#x60;287&#x60;: [Custody] Transfer in custody funding account  &#x60;288&#x60;: [Custody] Fund delegation   &#x60;289&#x60;: [Custody] Fund undelegation  &#x60;299&#x60;: Affiliate recommendation commission  &#x60;300&#x60;: Fee discount rebate  &#x60;303&#x60;: Snowball market maker transfer  &#x60;304&#x60;: [Simple Earn Fixed] Order submission  &#x60;305&#x60;: [Simple Earn Fixed] Order redemption  &#x60;306&#x60;: [Simple Earn Fixed] Principal distribution  &#x60;307&#x60;: [Simple Earn Fixed] Interest distribution (early termination compensation)  &#x60;308&#x60;: [Simple Earn Fixed] Interest distribution  &#x60;309&#x60;: [Simple Earn Fixed] Interest distribution (extension compensation)   &#x60;311&#x60;: Crypto dust auto-transfer in  &#x60;313&#x60;: Sent by gift  &#x60;314&#x60;: Received from gift  &#x60;315&#x60;: Refunded from gift  &#x60;328&#x60;: [SOL staking] Send Liquidity Staking Token reward  &#x60;329&#x60;: [SOL staking] Subscribe Liquidity Staking Token staking  &#x60;330&#x60;: [SOL staking] Mint Liquidity Staking Token  &#x60;331&#x60;: [SOL staking] Redeem Liquidity Staking Token order  &#x60;332&#x60;: [SOL staking] Settle Liquidity Staking Token order  &#x60;333&#x60;: Trial fund reward  &#x60;336&#x60;: [Credit line] Loan Forced Repayment  &#x60;338&#x60;: [Credit line] Forced Repayment Refund  &#x60;354&#x60;: Copy and bot rewards  &#x60;361&#x60;: Deposit from closed sub-account | [default to &quot;&quot;]
 **clientId** | **string** | Client-supplied ID for transfer or withdrawal  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ts&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ts&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAssetBillsV5Resp**](GetAssetBillsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetConvertCurrenciesV5

> GetAssetConvertCurrenciesV5Resp GetAssetConvertCurrenciesV5(ctx).Execute()





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
	resp, r, err := apiClient.FundingAccountAPI.GetAssetConvertCurrenciesV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetConvertCurrenciesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetConvertCurrenciesV5`: GetAssetConvertCurrenciesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetConvertCurrenciesV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetConvertCurrenciesV5Request struct via the builder pattern


### Return type

[**GetAssetConvertCurrenciesV5Resp**](GetAssetConvertCurrenciesV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetConvertCurrencyPairV5

> GetAssetConvertCurrencyPairV5Resp GetAssetConvertCurrencyPairV5(ctx).Execute()





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
	resp, r, err := apiClient.FundingAccountAPI.GetAssetConvertCurrencyPairV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetConvertCurrencyPairV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetConvertCurrencyPairV5`: GetAssetConvertCurrencyPairV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetConvertCurrencyPairV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetConvertCurrencyPairV5Request struct via the builder pattern


### Return type

[**GetAssetConvertCurrencyPairV5Resp**](GetAssetConvertCurrencyPairV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetConvertHistoryV5

> GetAssetConvertHistoryV5Resp GetAssetConvertHistoryV5(ctx).ClTReqId(clTReqId).After(after).Before(before).Limit(limit).Tag(tag).Execute()





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
	clTReqId := "clTReqId_example" // string | Client Order ID as assigned by the client  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ts`, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ts`, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")
	tag := "tag_example" // string | Order tag  Applicable to broker user  If the convert trading used `tag`, this parameter is also required. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetAssetConvertHistoryV5(context.Background()).ClTReqId(clTReqId).After(after).Before(before).Limit(limit).Tag(tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetConvertHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetConvertHistoryV5`: GetAssetConvertHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetConvertHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetConvertHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clTReqId** | **string** | Client Order ID as assigned by the client  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ts&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ts&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]
 **tag** | **string** | Order tag  Applicable to broker user  If the convert trading used &#x60;tag&#x60;, this parameter is also required. | [default to &quot;&quot;]

### Return type

[**GetAssetConvertHistoryV5Resp**](GetAssetConvertHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetCurrenciesV5

> GetAssetCurrenciesV5Resp GetAssetCurrenciesV5(ctx).Ccy(ccy).Execute()

Retrieve a list of all currencies available which are related to the current account's KYC entity.  



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
	ccy := "ccy_example" // string | Single currency or multiple currencies separated with comma, e.g. `BTC` or `BTC,ETH`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetAssetCurrenciesV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetCurrenciesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetCurrenciesV5`: GetAssetCurrenciesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetCurrenciesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetCurrenciesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Single currency or multiple currencies separated with comma, e.g. &#x60;BTC&#x60; or &#x60;BTC,ETH&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAssetCurrenciesV5Resp**](GetAssetCurrenciesV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDepositAddressV5

> GetAssetDepositAddressV5Resp GetAssetDepositAddressV5(ctx).Ccy(ccy).Execute()

Retrieve the deposit addresses of currencies, including previously-used addresses.  



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
	ccy := "ccy_example" // string | Currency, e.g. `BTC` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetAssetDepositAddressV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetDepositAddressV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetDepositAddressV5`: GetAssetDepositAddressV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetDepositAddressV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDepositAddressV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]

### Return type

[**GetAssetDepositAddressV5Resp**](GetAssetDepositAddressV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDepositHistoryV5

> GetAssetDepositHistoryV5Resp GetAssetDepositHistoryV5(ctx).Ccy(ccy).DepId(depId).FromWdId(fromWdId).TxId(txId).Type_(type_).State(state).After(after).Before(before).Limit(limit).Execute()

Retrieve the deposit records according to the currency, deposit status, and time range in reverse chronological order. The 100 most recent records are returned by default.   Websocket API is also available, refer to .  



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
	ccy := "ccy_example" // string | Currency, e.g. `BTC` (optional) (default to "")
	depId := "depId_example" // string | Deposit ID (optional) (default to "")
	fromWdId := "fromWdId_example" // string | Internal transfer initiator's withdrawal ID  If the deposit comes from internal transfer, this field displays the withdrawal ID of the internal transfer initiator (optional) (default to "")
	txId := "txId_example" // string | Hash record of the deposit (optional) (default to "")
	type_ := "type__example" // string | Deposit Type  `3`: internal transfer  `4`: deposit from chain (optional) (default to "")
	state := "state_example" // string | Status of deposit    `0`: waiting for confirmation  `1`: deposit credited    `2`: deposit successful   `8`: pending due to temporary deposit suspension on this crypto currency  `11`: match the address blacklist  `12`: account or deposit is frozen  `13`: sub-account deposit interception  `14`: KYC limit (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested ts, Unix timestamp format in milliseconds, e.g. `1654041600000` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested ts, Unix timestamp format in milliseconds, e.g. `1656633600000` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetAssetDepositHistoryV5(context.Background()).Ccy(ccy).DepId(depId).FromWdId(fromWdId).TxId(txId).Type_(type_).State(state).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetDepositHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetDepositHistoryV5`: GetAssetDepositHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetDepositHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDepositHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]
 **depId** | **string** | Deposit ID | [default to &quot;&quot;]
 **fromWdId** | **string** | Internal transfer initiator&#39;s withdrawal ID  If the deposit comes from internal transfer, this field displays the withdrawal ID of the internal transfer initiator | [default to &quot;&quot;]
 **txId** | **string** | Hash record of the deposit | [default to &quot;&quot;]
 **type_** | **string** | Deposit Type  &#x60;3&#x60;: internal transfer  &#x60;4&#x60;: deposit from chain | [default to &quot;&quot;]
 **state** | **string** | Status of deposit    &#x60;0&#x60;: waiting for confirmation  &#x60;1&#x60;: deposit credited    &#x60;2&#x60;: deposit successful   &#x60;8&#x60;: pending due to temporary deposit suspension on this crypto currency  &#x60;11&#x60;: match the address blacklist  &#x60;12&#x60;: account or deposit is frozen  &#x60;13&#x60;: sub-account deposit interception  &#x60;14&#x60;: KYC limit | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested ts, Unix timestamp format in milliseconds, e.g. &#x60;1654041600000&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested ts, Unix timestamp format in milliseconds, e.g. &#x60;1656633600000&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetAssetDepositHistoryV5Resp**](GetAssetDepositHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetDepositWithdrawStatusV5

> GetAssetDepositWithdrawStatusV5Resp GetAssetDepositWithdrawStatusV5(ctx).WdId(wdId).TxId(txId).Ccy(ccy).To(to).Chain(chain).Execute()

Retrieve deposit's and withdrawal's detailed status and estimated complete time.  



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
	wdId := "wdId_example" // string | Withdrawal ID, use to retrieve withdrawal status   Required to input one and only one of `wdId` and `txId` (optional) (default to "")
	txId := "txId_example" // string | Hash record of the deposit, use to retrieve deposit status   Required to input one and only one of `wdId` and `txId` (optional) (default to "")
	ccy := "ccy_example" // string | Currency type, e.g. `USDT`   Required when retrieving deposit status with `txId` (optional) (default to "")
	to := "to_example" // string | To address, the destination address in deposit   Required when retrieving deposit status with `txId` (optional) (default to "")
	chain := "chain_example" // string | Currency chain information, e.g. USDT-ERC20   Required when retrieving deposit status with `txId` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetAssetDepositWithdrawStatusV5(context.Background()).WdId(wdId).TxId(txId).Ccy(ccy).To(to).Chain(chain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetDepositWithdrawStatusV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetDepositWithdrawStatusV5`: GetAssetDepositWithdrawStatusV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetDepositWithdrawStatusV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetDepositWithdrawStatusV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wdId** | **string** | Withdrawal ID, use to retrieve withdrawal status   Required to input one and only one of &#x60;wdId&#x60; and &#x60;txId&#x60; | [default to &quot;&quot;]
 **txId** | **string** | Hash record of the deposit, use to retrieve deposit status   Required to input one and only one of &#x60;wdId&#x60; and &#x60;txId&#x60; | [default to &quot;&quot;]
 **ccy** | **string** | Currency type, e.g. &#x60;USDT&#x60;   Required when retrieving deposit status with &#x60;txId&#x60; | [default to &quot;&quot;]
 **to** | **string** | To address, the destination address in deposit   Required when retrieving deposit status with &#x60;txId&#x60; | [default to &quot;&quot;]
 **chain** | **string** | Currency chain information, e.g. USDT-ERC20   Required when retrieving deposit status with &#x60;txId&#x60; | [default to &quot;&quot;]

### Return type

[**GetAssetDepositWithdrawStatusV5Resp**](GetAssetDepositWithdrawStatusV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetExchangeListV5

> GetAssetExchangeListV5Resp GetAssetExchangeListV5(ctx).Execute()

Authentication is not required for this public endpoint.  



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
	resp, r, err := apiClient.FundingAccountAPI.GetAssetExchangeListV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetExchangeListV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetExchangeListV5`: GetAssetExchangeListV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetExchangeListV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetExchangeListV5Request struct via the builder pattern


### Return type

[**GetAssetExchangeListV5Resp**](GetAssetExchangeListV5Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetMonthlyStatementV5

> GetAssetMonthlyStatementV5Resp GetAssetMonthlyStatementV5(ctx).Month(month).Execute()

Retrieve monthly statement in the past year.  



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
	month := "month_example" // string | Month, valid value is `Jan`, `Feb`, `Mar`, `Apr`,`May`, `Jun`, `Jul`,`Aug`, `Sep`,`Oct`,`Nov`,`Dec` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetAssetMonthlyStatementV5(context.Background()).Month(month).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetMonthlyStatementV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetMonthlyStatementV5`: GetAssetMonthlyStatementV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetMonthlyStatementV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetMonthlyStatementV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **month** | **string** | Month, valid value is &#x60;Jan&#x60;, &#x60;Feb&#x60;, &#x60;Mar&#x60;, &#x60;Apr&#x60;,&#x60;May&#x60;, &#x60;Jun&#x60;, &#x60;Jul&#x60;,&#x60;Aug&#x60;, &#x60;Sep&#x60;,&#x60;Oct&#x60;,&#x60;Nov&#x60;,&#x60;Dec&#x60; | [default to &quot;&quot;]

### Return type

[**GetAssetMonthlyStatementV5Resp**](GetAssetMonthlyStatementV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetNonTradableAssetsV5

> GetAssetNonTradableAssetsV5Resp GetAssetNonTradableAssetsV5(ctx).Ccy(ccy).Execute()





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
	ccy := "ccy_example" // string | Single currency or multiple currencies (no more than 20) separated with comma, e.g. `BTC` or `BTC,ETH`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetAssetNonTradableAssetsV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetNonTradableAssetsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetNonTradableAssetsV5`: GetAssetNonTradableAssetsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetNonTradableAssetsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetNonTradableAssetsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Single currency or multiple currencies (no more than 20) separated with comma, e.g. &#x60;BTC&#x60; or &#x60;BTC,ETH&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAssetNonTradableAssetsV5Resp**](GetAssetNonTradableAssetsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetTransferStateV5

> GetAssetTransferStateV5Resp GetAssetTransferStateV5(ctx).TransId(transId).ClientId(clientId).Type_(type_).Execute()

Retrieve the transfer state data of the last 2 weeks.  



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
	transId := "transId_example" // string | Transfer ID  Either transId or clientId is required. If both are passed, transId will be used. (optional) (default to "")
	clientId := "clientId_example" // string | Client-supplied ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. (optional) (default to "")
	type_ := "type__example" // string | Transfer type  `0`: transfer within account   `1`: master account to sub-account (Only applicable to API Key from master account)   `2`: sub-account to master account (Only applicable to API Key from master account)  `3`: sub-account to master account (Only applicable to APIKey from sub-account)  `4`: sub-account to sub-account (Only applicable to APIKey from sub-account, and target account needs to be another sub-account which belongs to same master account)  The default is `0`.  For Custody accounts, can choose not to pass this parameter or pass `0`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetAssetTransferStateV5(context.Background()).TransId(transId).ClientId(clientId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetTransferStateV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetTransferStateV5`: GetAssetTransferStateV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetTransferStateV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetTransferStateV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transId** | **string** | Transfer ID  Either transId or clientId is required. If both are passed, transId will be used. | [default to &quot;&quot;]
 **clientId** | **string** | Client-supplied ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [default to &quot;&quot;]
 **type_** | **string** | Transfer type  &#x60;0&#x60;: transfer within account   &#x60;1&#x60;: master account to sub-account (Only applicable to API Key from master account)   &#x60;2&#x60;: sub-account to master account (Only applicable to API Key from master account)  &#x60;3&#x60;: sub-account to master account (Only applicable to APIKey from sub-account)  &#x60;4&#x60;: sub-account to sub-account (Only applicable to APIKey from sub-account, and target account needs to be another sub-account which belongs to same master account)  The default is &#x60;0&#x60;.  For Custody accounts, can choose not to pass this parameter or pass &#x60;0&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAssetTransferStateV5Resp**](GetAssetTransferStateV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetWithdrawalHistoryV5

> GetAssetWithdrawalHistoryV5Resp GetAssetWithdrawalHistoryV5(ctx).Ccy(ccy).WdId(wdId).ClientId(clientId).TxId(txId).Type_(type_).State(state).After(after).Before(before).Limit(limit).Execute()

Retrieve the withdrawal records according to the currency, withdrawal status, and time range in reverse chronological order. The 100 most recent records are returned by default.   Websocket API is also available, refer to .  



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
	ccy := "ccy_example" // string | Currency, e.g. `BTC` (optional) (default to "")
	wdId := "wdId_example" // string | Withdrawal ID (optional) (default to "")
	clientId := "clientId_example" // string | Client-supplied ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. (optional) (default to "")
	txId := "txId_example" // string | Hash record of the deposit (optional) (default to "")
	type_ := "type__example" // string | Withdrawal type  `3`: Internal transfer  `4`: On-chain withdrawal (optional) (default to "")
	state := "state_example" // string | Status of withdrawal    `17`: Pending response from Travel Rule vendor  `10`: Waiting transfer  `0`: Waiting withdrawal  `4`/`5`/`6`/`8`/`9`/`12`: Waiting manual review  `7`: Approved    `1`: Broadcasting your transaction to chain  `15`: Pending transaction validation  `16`: Due to local laws and regulations, your withdrawal may take up to 24 hours to arrive  `-3`: Canceling     `-2`: Canceled   `-1`: Failed  `2`: Success (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested ts, Unix timestamp format in milliseconds, e.g. `1654041600000` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested ts, Unix timestamp format in milliseconds, e.g. `1656633600000` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetAssetWithdrawalHistoryV5(context.Background()).Ccy(ccy).WdId(wdId).ClientId(clientId).TxId(txId).Type_(type_).State(state).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetAssetWithdrawalHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetWithdrawalHistoryV5`: GetAssetWithdrawalHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetAssetWithdrawalHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetWithdrawalHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]
 **wdId** | **string** | Withdrawal ID | [default to &quot;&quot;]
 **clientId** | **string** | Client-supplied ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [default to &quot;&quot;]
 **txId** | **string** | Hash record of the deposit | [default to &quot;&quot;]
 **type_** | **string** | Withdrawal type  &#x60;3&#x60;: Internal transfer  &#x60;4&#x60;: On-chain withdrawal | [default to &quot;&quot;]
 **state** | **string** | Status of withdrawal    &#x60;17&#x60;: Pending response from Travel Rule vendor  &#x60;10&#x60;: Waiting transfer  &#x60;0&#x60;: Waiting withdrawal  &#x60;4&#x60;/&#x60;5&#x60;/&#x60;6&#x60;/&#x60;8&#x60;/&#x60;9&#x60;/&#x60;12&#x60;: Waiting manual review  &#x60;7&#x60;: Approved    &#x60;1&#x60;: Broadcasting your transaction to chain  &#x60;15&#x60;: Pending transaction validation  &#x60;16&#x60;: Due to local laws and regulations, your withdrawal may take up to 24 hours to arrive  &#x60;-3&#x60;: Canceling     &#x60;-2&#x60;: Canceled   &#x60;-1&#x60;: Failed  &#x60;2&#x60;: Success | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested ts, Unix timestamp format in milliseconds, e.g. &#x60;1654041600000&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested ts, Unix timestamp format in milliseconds, e.g. &#x60;1656633600000&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetAssetWithdrawalHistoryV5Resp**](GetAssetWithdrawalHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiatDepositOrderHistoryV5

> GetFiatDepositOrderHistoryV5Resp GetFiatDepositOrderHistoryV5(ctx).Ccy(ccy).PaymentMethod(paymentMethod).State(state).After(after).Before(before).Limit(limit).Execute()

Get fiat deposit order history  



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
	ccy := "ccy_example" // string | ISO-4217 3 digit currency code (optional) (default to "")
	paymentMethod := "paymentMethod_example" // string | Payment Method  `TR_BANKS`  `PIX`  `SEPA` (optional) (default to "")
	state := "state_example" // string | State of the order  `completed`  `failed`  `pending`  `canceled`  `inqueue`  `processing` (optional) (default to "")
	after := "after_example" // string | Filter with a begin timestamp. Unix timestamp format in milliseconds (inclusive), e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Filter with an end timestamp. Unix timestamp format in milliseconds (inclusive), e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. Maximum and default is 100 (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetFiatDepositOrderHistoryV5(context.Background()).Ccy(ccy).PaymentMethod(paymentMethod).State(state).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetFiatDepositOrderHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiatDepositOrderHistoryV5`: GetFiatDepositOrderHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetFiatDepositOrderHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFiatDepositOrderHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | ISO-4217 3 digit currency code | [default to &quot;&quot;]
 **paymentMethod** | **string** | Payment Method  &#x60;TR_BANKS&#x60;  &#x60;PIX&#x60;  &#x60;SEPA&#x60; | [default to &quot;&quot;]
 **state** | **string** | State of the order  &#x60;completed&#x60;  &#x60;failed&#x60;  &#x60;pending&#x60;  &#x60;canceled&#x60;  &#x60;inqueue&#x60;  &#x60;processing&#x60; | [default to &quot;&quot;]
 **after** | **string** | Filter with a begin timestamp. Unix timestamp format in milliseconds (inclusive), e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Filter with an end timestamp. Unix timestamp format in milliseconds (inclusive), e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. Maximum and default is 100 | [default to &quot;&quot;]

### Return type

[**GetFiatDepositOrderHistoryV5Resp**](GetFiatDepositOrderHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiatDepositPaymentMethodsV5

> GetFiatDepositPaymentMethodsV5Resp GetFiatDepositPaymentMethodsV5(ctx).Ccy(ccy).Execute()

To display all the available fiat deposit payment methods  



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
	ccy := "ccy_example" // string | Fiat currency, ISO-4217 3 digit currency code, e.g. `TRY` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetFiatDepositPaymentMethodsV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetFiatDepositPaymentMethodsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiatDepositPaymentMethodsV5`: GetFiatDepositPaymentMethodsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetFiatDepositPaymentMethodsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFiatDepositPaymentMethodsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Fiat currency, ISO-4217 3 digit currency code, e.g. &#x60;TRY&#x60; | [default to &quot;&quot;]

### Return type

[**GetFiatDepositPaymentMethodsV5Resp**](GetFiatDepositPaymentMethodsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiatDepositV5

> GetFiatDepositV5Resp GetFiatDepositV5(ctx).OrdId(ordId).Execute()

Get fiat deposit order detail  



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
	ordId := "ordId_example" // string | Order ID (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetFiatDepositV5(context.Background()).OrdId(ordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetFiatDepositV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiatDepositV5`: GetFiatDepositV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetFiatDepositV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFiatDepositV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordId** | **string** | Order ID | [default to &quot;&quot;]

### Return type

[**GetFiatDepositV5Resp**](GetFiatDepositV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiatWithdrawalOrderHistoryV5

> GetFiatWithdrawalOrderHistoryV5Resp GetFiatWithdrawalOrderHistoryV5(ctx).Ccy(ccy).PaymentMethod(paymentMethod).State(state).After(after).Before(before).Limit(limit).Execute()

Get fiat withdrawal order history  



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
	ccy := "ccy_example" // string | Fiat currency, ISO-4217 3 digit currency code, e.g. `TRY` (optional) (default to "")
	paymentMethod := "paymentMethod_example" // string | Payment Method  `TR_BANKS`  `PIX`  `SEPA` (optional) (default to "")
	state := "state_example" // string | State of the order  `completed`  `failed`  `pending`  `canceled`  `inqueue`  `processing` (optional) (default to "")
	after := "after_example" // string | Filter with a begin timestamp. Unix timestamp format in milliseconds (inclusive), e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Filter with an end timestamp. Unix timestamp format in milliseconds (inclusive), e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. Maximum and default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetFiatWithdrawalOrderHistoryV5(context.Background()).Ccy(ccy).PaymentMethod(paymentMethod).State(state).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetFiatWithdrawalOrderHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiatWithdrawalOrderHistoryV5`: GetFiatWithdrawalOrderHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetFiatWithdrawalOrderHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFiatWithdrawalOrderHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Fiat currency, ISO-4217 3 digit currency code, e.g. &#x60;TRY&#x60; | [default to &quot;&quot;]
 **paymentMethod** | **string** | Payment Method  &#x60;TR_BANKS&#x60;  &#x60;PIX&#x60;  &#x60;SEPA&#x60; | [default to &quot;&quot;]
 **state** | **string** | State of the order  &#x60;completed&#x60;  &#x60;failed&#x60;  &#x60;pending&#x60;  &#x60;canceled&#x60;  &#x60;inqueue&#x60;  &#x60;processing&#x60; | [default to &quot;&quot;]
 **after** | **string** | Filter with a begin timestamp. Unix timestamp format in milliseconds (inclusive), e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Filter with an end timestamp. Unix timestamp format in milliseconds (inclusive), e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. Maximum and default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetFiatWithdrawalOrderHistoryV5Resp**](GetFiatWithdrawalOrderHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiatWithdrawalPaymentMethodsV5

> GetFiatWithdrawalPaymentMethodsV5Resp GetFiatWithdrawalPaymentMethodsV5(ctx).Ccy(ccy).Execute()

To display all the available fiat withdrawal payment methods  



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
	ccy := "ccy_example" // string | Fiat currency, ISO-4217 3 digit currency code. e.g. `TRY` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetFiatWithdrawalPaymentMethodsV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetFiatWithdrawalPaymentMethodsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiatWithdrawalPaymentMethodsV5`: GetFiatWithdrawalPaymentMethodsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetFiatWithdrawalPaymentMethodsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFiatWithdrawalPaymentMethodsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Fiat currency, ISO-4217 3 digit currency code. e.g. &#x60;TRY&#x60; | [default to &quot;&quot;]

### Return type

[**GetFiatWithdrawalPaymentMethodsV5Resp**](GetFiatWithdrawalPaymentMethodsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiatWithdrawalV5

> GetFiatWithdrawalV5Resp GetFiatWithdrawalV5(ctx).OrdId(ordId).Execute()

Get fiat withdraw order detail  



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
	ordId := "ordId_example" // string | Order ID (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FundingAccountAPI.GetFiatWithdrawalV5(context.Background()).OrdId(ordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FundingAccountAPI.GetFiatWithdrawalV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiatWithdrawalV5`: GetFiatWithdrawalV5Resp
	fmt.Fprintf(os.Stdout, "Response from `FundingAccountAPI.GetFiatWithdrawalV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFiatWithdrawalV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordId** | **string** | Order ID | [default to &quot;&quot;]

### Return type

[**GetFiatWithdrawalV5Resp**](GetFiatWithdrawalV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

