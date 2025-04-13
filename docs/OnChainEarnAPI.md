# \OnChainEarnAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFinanceStakingDefiCancelV5**](OnChainEarnAPI.md#CreateFinanceStakingDefiCancelV5) | **Post** /api/v5/finance/staking-defi/cancel | POST / Cancel purchases/redemptions
[**CreateFinanceStakingDefiPurchaseV5**](OnChainEarnAPI.md#CreateFinanceStakingDefiPurchaseV5) | **Post** /api/v5/finance/staking-defi/purchase | POST / Purchase
[**CreateFinanceStakingDefiRedeemV5**](OnChainEarnAPI.md#CreateFinanceStakingDefiRedeemV5) | **Post** /api/v5/finance/staking-defi/redeem | POST / Redeem
[**GetFinanceStakingDefiOffersV5**](OnChainEarnAPI.md#GetFinanceStakingDefiOffersV5) | **Get** /api/v5/finance/staking-defi/offers | GET / Offers
[**GetFinanceStakingDefiOrdersActiveV5**](OnChainEarnAPI.md#GetFinanceStakingDefiOrdersActiveV5) | **Get** /api/v5/finance/staking-defi/orders-active | GET / Active orders
[**GetFinanceStakingDefiOrdersHistoryV5**](OnChainEarnAPI.md#GetFinanceStakingDefiOrdersHistoryV5) | **Get** /api/v5/finance/staking-defi/orders-history | GET / Order history



## CreateFinanceStakingDefiCancelV5

> CreateFinanceStakingDefiCancelV5Resp CreateFinanceStakingDefiCancelV5(ctx).CreateFinanceStakingDefiCancelV5Req(createFinanceStakingDefiCancelV5Req).Execute()

POST / Cancel purchases/redemptions



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
	createFinanceStakingDefiCancelV5Req := *openapiclient.NewCreateFinanceStakingDefiCancelV5Req("OrdId_example", "ProtocolType_example") // CreateFinanceStakingDefiCancelV5Req | The request body for CreateFinanceStakingDefiCancelV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnChainEarnAPI.CreateFinanceStakingDefiCancelV5(context.Background()).CreateFinanceStakingDefiCancelV5Req(createFinanceStakingDefiCancelV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnChainEarnAPI.CreateFinanceStakingDefiCancelV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFinanceStakingDefiCancelV5`: CreateFinanceStakingDefiCancelV5Resp
	fmt.Fprintf(os.Stdout, "Response from `OnChainEarnAPI.CreateFinanceStakingDefiCancelV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFinanceStakingDefiCancelV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFinanceStakingDefiCancelV5Req** | [**CreateFinanceStakingDefiCancelV5Req**](CreateFinanceStakingDefiCancelV5Req.md) | The request body for CreateFinanceStakingDefiCancelV5 | 

### Return type

[**CreateFinanceStakingDefiCancelV5Resp**](CreateFinanceStakingDefiCancelV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFinanceStakingDefiPurchaseV5

> CreateFinanceStakingDefiPurchaseV5Resp CreateFinanceStakingDefiPurchaseV5(ctx).CreateFinanceStakingDefiPurchaseV5Req(createFinanceStakingDefiPurchaseV5Req).Execute()

POST / Purchase



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
	createFinanceStakingDefiPurchaseV5Req := *openapiclient.NewCreateFinanceStakingDefiPurchaseV5Req([]openapiclient.CreateFinanceStakingDefiPurchaseV5ReqInvestDataInner{*openapiclient.NewCreateFinanceStakingDefiPurchaseV5ReqInvestDataInner()}, "ProductId_example") // CreateFinanceStakingDefiPurchaseV5Req | The request body for CreateFinanceStakingDefiPurchaseV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnChainEarnAPI.CreateFinanceStakingDefiPurchaseV5(context.Background()).CreateFinanceStakingDefiPurchaseV5Req(createFinanceStakingDefiPurchaseV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnChainEarnAPI.CreateFinanceStakingDefiPurchaseV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFinanceStakingDefiPurchaseV5`: CreateFinanceStakingDefiPurchaseV5Resp
	fmt.Fprintf(os.Stdout, "Response from `OnChainEarnAPI.CreateFinanceStakingDefiPurchaseV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFinanceStakingDefiPurchaseV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFinanceStakingDefiPurchaseV5Req** | [**CreateFinanceStakingDefiPurchaseV5Req**](CreateFinanceStakingDefiPurchaseV5Req.md) | The request body for CreateFinanceStakingDefiPurchaseV5 | 

### Return type

[**CreateFinanceStakingDefiPurchaseV5Resp**](CreateFinanceStakingDefiPurchaseV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFinanceStakingDefiRedeemV5

> CreateFinanceStakingDefiRedeemV5Resp CreateFinanceStakingDefiRedeemV5(ctx).CreateFinanceStakingDefiRedeemV5Req(createFinanceStakingDefiRedeemV5Req).Execute()

POST / Redeem



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
	createFinanceStakingDefiRedeemV5Req := *openapiclient.NewCreateFinanceStakingDefiRedeemV5Req("OrdId_example", "ProtocolType_example") // CreateFinanceStakingDefiRedeemV5Req | The request body for CreateFinanceStakingDefiRedeemV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnChainEarnAPI.CreateFinanceStakingDefiRedeemV5(context.Background()).CreateFinanceStakingDefiRedeemV5Req(createFinanceStakingDefiRedeemV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnChainEarnAPI.CreateFinanceStakingDefiRedeemV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFinanceStakingDefiRedeemV5`: CreateFinanceStakingDefiRedeemV5Resp
	fmt.Fprintf(os.Stdout, "Response from `OnChainEarnAPI.CreateFinanceStakingDefiRedeemV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFinanceStakingDefiRedeemV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFinanceStakingDefiRedeemV5Req** | [**CreateFinanceStakingDefiRedeemV5Req**](CreateFinanceStakingDefiRedeemV5Req.md) | The request body for CreateFinanceStakingDefiRedeemV5 | 

### Return type

[**CreateFinanceStakingDefiRedeemV5Resp**](CreateFinanceStakingDefiRedeemV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceStakingDefiOffersV5

> GetFinanceStakingDefiOffersV5Resp GetFinanceStakingDefiOffersV5(ctx).ProductId(productId).ProtocolType(protocolType).Ccy(ccy).Execute()

GET / Offers



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
	productId := "productId_example" // string | Product ID (optional) (default to "")
	protocolType := "protocolType_example" // string | Protocol type  `defi`: on-chain earn (optional) (default to "")
	ccy := "ccy_example" // string | Investment currency, e.g. `BTC` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnChainEarnAPI.GetFinanceStakingDefiOffersV5(context.Background()).ProductId(productId).ProtocolType(protocolType).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnChainEarnAPI.GetFinanceStakingDefiOffersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceStakingDefiOffersV5`: GetFinanceStakingDefiOffersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `OnChainEarnAPI.GetFinanceStakingDefiOffersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceStakingDefiOffersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Product ID | [default to &quot;&quot;]
 **protocolType** | **string** | Protocol type  &#x60;defi&#x60;: on-chain earn | [default to &quot;&quot;]
 **ccy** | **string** | Investment currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]

### Return type

[**GetFinanceStakingDefiOffersV5Resp**](GetFinanceStakingDefiOffersV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceStakingDefiOrdersActiveV5

> GetFinanceStakingDefiOrdersActiveV5Resp GetFinanceStakingDefiOrdersActiveV5(ctx).ProductId(productId).ProtocolType(protocolType).Ccy(ccy).State(state).Execute()

GET / Active orders



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
	productId := "productId_example" // string | Product ID (optional) (default to "")
	protocolType := "protocolType_example" // string | Protocol type  `defi`: on-chain earn (optional) (default to "")
	ccy := "ccy_example" // string | Investment currency, e.g. `BTC` (optional) (default to "")
	state := "state_example" // string | Order state  `8`: Pending   `13`: Cancelling   `9`: Onchain   `1`: Earning   `2`: Redeeming (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnChainEarnAPI.GetFinanceStakingDefiOrdersActiveV5(context.Background()).ProductId(productId).ProtocolType(protocolType).Ccy(ccy).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnChainEarnAPI.GetFinanceStakingDefiOrdersActiveV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceStakingDefiOrdersActiveV5`: GetFinanceStakingDefiOrdersActiveV5Resp
	fmt.Fprintf(os.Stdout, "Response from `OnChainEarnAPI.GetFinanceStakingDefiOrdersActiveV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceStakingDefiOrdersActiveV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Product ID | [default to &quot;&quot;]
 **protocolType** | **string** | Protocol type  &#x60;defi&#x60;: on-chain earn | [default to &quot;&quot;]
 **ccy** | **string** | Investment currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]
 **state** | **string** | Order state  &#x60;8&#x60;: Pending   &#x60;13&#x60;: Cancelling   &#x60;9&#x60;: Onchain   &#x60;1&#x60;: Earning   &#x60;2&#x60;: Redeeming | [default to &quot;&quot;]

### Return type

[**GetFinanceStakingDefiOrdersActiveV5Resp**](GetFinanceStakingDefiOrdersActiveV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinanceStakingDefiOrdersHistoryV5

> GetFinanceStakingDefiOrdersHistoryV5Resp GetFinanceStakingDefiOrdersHistoryV5(ctx).ProductId(productId).ProtocolType(protocolType).Ccy(ccy).After(after).Before(before).Limit(limit).Execute()

GET / Order history



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
	productId := "productId_example" // string | Product ID (optional) (default to "")
	protocolType := "protocolType_example" // string | Protocol type  `defi`: on-chain earn (optional) (default to "")
	ccy := "ccy_example" // string | Investment currency, e.g. `BTC` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested ID. The value passed is the corresponding `ordId` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested ID. The value passed is the corresponding `ordId` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The default is `100`. The maximum is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnChainEarnAPI.GetFinanceStakingDefiOrdersHistoryV5(context.Background()).ProductId(productId).ProtocolType(protocolType).Ccy(ccy).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnChainEarnAPI.GetFinanceStakingDefiOrdersHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinanceStakingDefiOrdersHistoryV5`: GetFinanceStakingDefiOrdersHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `OnChainEarnAPI.GetFinanceStakingDefiOrdersHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinanceStakingDefiOrdersHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Product ID | [default to &quot;&quot;]
 **protocolType** | **string** | Protocol type  &#x60;defi&#x60;: on-chain earn | [default to &quot;&quot;]
 **ccy** | **string** | Investment currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested ID. The value passed is the corresponding &#x60;ordId&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested ID. The value passed is the corresponding &#x60;ordId&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The default is &#x60;100&#x60;. The maximum is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetFinanceStakingDefiOrdersHistoryV5Resp**](GetFinanceStakingDefiOrdersHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

