# \SubAccountAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssetSubaccountTransferV5**](SubAccountAPI.md#CreateAssetSubaccountTransferV5) | **Post** /api/v5/asset/subaccount/transfer | Applies to master accounts only.   Only API keys with &#x60;Trade&#x60; privilege can call this endpoint.  
[**CreateUsersSubaccountModifyApikeyV5**](SubAccountAPI.md#CreateUsersSubaccountModifyApikeyV5) | **Post** /api/v5/users/subaccount/modify-apikey | Applies to master accounts only and master accounts API Key must be linked to IP addresses. Only API keys with &#x60;Trade&#x60; privilege can call this endpoint.  
[**CreateUsersSubaccountSetTransferOutV5**](SubAccountAPI.md#CreateUsersSubaccountSetTransferOutV5) | **Post** /api/v5/users/subaccount/set-transfer-out | Set permission of transfer out for sub-account (only applicable to master account API key). Sub-account can transfer out to master account by default.  
[**GetAccountSubaccountBalancesV5**](SubAccountAPI.md#GetAccountSubaccountBalancesV5) | **Get** /api/v5/account/subaccount/balances | Query detailed balance info of Trading Account of a sub-account via the master account (applies to master accounts only)  
[**GetAccountSubaccountMaxWithdrawalV5**](SubAccountAPI.md#GetAccountSubaccountMaxWithdrawalV5) | **Get** /api/v5/account/subaccount/max-withdrawal | Retrieve the maximum withdrawal information of a sub-account via the master account (applies to master accounts only). If no currency is specified, the transferable amount of all owned currencies will be returned.  
[**GetAssetSubaccountBalancesV5**](SubAccountAPI.md#GetAssetSubaccountBalancesV5) | **Get** /api/v5/asset/subaccount/balances | Query detailed balance info of Funding Account of a sub-account via the master account (applies to master accounts only)  
[**GetAssetSubaccountBillsV5**](SubAccountAPI.md#GetAssetSubaccountBillsV5) | **Get** /api/v5/asset/subaccount/bills | Applies to master accounts only.  
[**GetAssetSubaccountManagedSubaccountBillsV5**](SubAccountAPI.md#GetAssetSubaccountManagedSubaccountBillsV5) | **Get** /api/v5/asset/subaccount/managed-subaccount-bills | Only applicable to the trading team&#39;s master account to getting transfer records of managed sub accounts entrusted to oneself.  
[**GetUsersEntrustSubaccountListV5**](SubAccountAPI.md#GetUsersEntrustSubaccountListV5) | **Get** /api/v5/users/entrust-subaccount-list | The trading team uses this interface to view the list of sub-accounts currently under escrow  
[**GetUsersSubaccountListV5**](SubAccountAPI.md#GetUsersSubaccountListV5) | **Get** /api/v5/users/subaccount/list | Applies to master accounts only  



## CreateAssetSubaccountTransferV5

> CreateAssetSubaccountTransferV5Resp CreateAssetSubaccountTransferV5(ctx).CreateAssetSubaccountTransferV5Req(createAssetSubaccountTransferV5Req).Execute()

Applies to master accounts only.   Only API keys with `Trade` privilege can call this endpoint.  



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
	createAssetSubaccountTransferV5Req := *openapiclient.NewCreateAssetSubaccountTransferV5Req("Amt_example", "Ccy_example", "From_example", "FromSubAccount_example", "To_example", "ToSubAccount_example") // CreateAssetSubaccountTransferV5Req | The request body for CreateAssetSubaccountTransferV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateAssetSubaccountTransferV5(context.Background()).CreateAssetSubaccountTransferV5Req(createAssetSubaccountTransferV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateAssetSubaccountTransferV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetSubaccountTransferV5`: CreateAssetSubaccountTransferV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateAssetSubaccountTransferV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetSubaccountTransferV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAssetSubaccountTransferV5Req** | [**CreateAssetSubaccountTransferV5Req**](CreateAssetSubaccountTransferV5Req.md) | The request body for CreateAssetSubaccountTransferV5 | 

### Return type

[**CreateAssetSubaccountTransferV5Resp**](CreateAssetSubaccountTransferV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUsersSubaccountModifyApikeyV5

> CreateUsersSubaccountModifyApikeyV5Resp CreateUsersSubaccountModifyApikeyV5(ctx).CreateUsersSubaccountModifyApikeyV5Req(createUsersSubaccountModifyApikeyV5Req).Execute()

Applies to master accounts only and master accounts API Key must be linked to IP addresses. Only API keys with `Trade` privilege can call this endpoint.  



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
	createUsersSubaccountModifyApikeyV5Req := *openapiclient.NewCreateUsersSubaccountModifyApikeyV5Req("ApiKey_example", "SubAcct_example") // CreateUsersSubaccountModifyApikeyV5Req | The request body for CreateUsersSubaccountModifyApikeyV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateUsersSubaccountModifyApikeyV5(context.Background()).CreateUsersSubaccountModifyApikeyV5Req(createUsersSubaccountModifyApikeyV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateUsersSubaccountModifyApikeyV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUsersSubaccountModifyApikeyV5`: CreateUsersSubaccountModifyApikeyV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateUsersSubaccountModifyApikeyV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsersSubaccountModifyApikeyV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUsersSubaccountModifyApikeyV5Req** | [**CreateUsersSubaccountModifyApikeyV5Req**](CreateUsersSubaccountModifyApikeyV5Req.md) | The request body for CreateUsersSubaccountModifyApikeyV5 | 

### Return type

[**CreateUsersSubaccountModifyApikeyV5Resp**](CreateUsersSubaccountModifyApikeyV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUsersSubaccountSetTransferOutV5

> CreateUsersSubaccountSetTransferOutV5Resp CreateUsersSubaccountSetTransferOutV5(ctx).CreateUsersSubaccountSetTransferOutV5Req(createUsersSubaccountSetTransferOutV5Req).Execute()

Set permission of transfer out for sub-account (only applicable to master account API key). Sub-account can transfer out to master account by default.  



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
	createUsersSubaccountSetTransferOutV5Req := *openapiclient.NewCreateUsersSubaccountSetTransferOutV5Req("SubAcct_example") // CreateUsersSubaccountSetTransferOutV5Req | The request body for CreateUsersSubaccountSetTransferOutV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.CreateUsersSubaccountSetTransferOutV5(context.Background()).CreateUsersSubaccountSetTransferOutV5Req(createUsersSubaccountSetTransferOutV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.CreateUsersSubaccountSetTransferOutV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUsersSubaccountSetTransferOutV5`: CreateUsersSubaccountSetTransferOutV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.CreateUsersSubaccountSetTransferOutV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsersSubaccountSetTransferOutV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUsersSubaccountSetTransferOutV5Req** | [**CreateUsersSubaccountSetTransferOutV5Req**](CreateUsersSubaccountSetTransferOutV5Req.md) | The request body for CreateUsersSubaccountSetTransferOutV5 | 

### Return type

[**CreateUsersSubaccountSetTransferOutV5Resp**](CreateUsersSubaccountSetTransferOutV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountSubaccountBalancesV5

> GetAccountSubaccountBalancesV5Resp GetAccountSubaccountBalancesV5(ctx).SubAcct(subAcct).Execute()

Query detailed balance info of Trading Account of a sub-account via the master account (applies to master accounts only)  



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
	subAcct := "subAcct_example" // string | Sub-account name (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetAccountSubaccountBalancesV5(context.Background()).SubAcct(subAcct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetAccountSubaccountBalancesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountSubaccountBalancesV5`: GetAccountSubaccountBalancesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetAccountSubaccountBalancesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSubaccountBalancesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subAcct** | **string** | Sub-account name | [default to &quot;&quot;]

### Return type

[**GetAccountSubaccountBalancesV5Resp**](GetAccountSubaccountBalancesV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountSubaccountMaxWithdrawalV5

> GetAccountSubaccountMaxWithdrawalV5Resp GetAccountSubaccountMaxWithdrawalV5(ctx).SubAcct(subAcct).Ccy(ccy).Execute()

Retrieve the maximum withdrawal information of a sub-account via the master account (applies to master accounts only). If no currency is specified, the transferable amount of all owned currencies will be returned.  



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
	subAcct := "subAcct_example" // string | Sub-account name (default to "")
	ccy := "ccy_example" // string | Single currency or multiple currencies (no more than 20) separated with comma, e.g. `BTC` or `BTC,ETH`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetAccountSubaccountMaxWithdrawalV5(context.Background()).SubAcct(subAcct).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetAccountSubaccountMaxWithdrawalV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountSubaccountMaxWithdrawalV5`: GetAccountSubaccountMaxWithdrawalV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetAccountSubaccountMaxWithdrawalV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSubaccountMaxWithdrawalV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subAcct** | **string** | Sub-account name | [default to &quot;&quot;]
 **ccy** | **string** | Single currency or multiple currencies (no more than 20) separated with comma, e.g. &#x60;BTC&#x60; or &#x60;BTC,ETH&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAccountSubaccountMaxWithdrawalV5Resp**](GetAccountSubaccountMaxWithdrawalV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetSubaccountBalancesV5

> GetAssetSubaccountBalancesV5Resp GetAssetSubaccountBalancesV5(ctx).SubAcct(subAcct).Ccy(ccy).Execute()

Query detailed balance info of Funding Account of a sub-account via the master account (applies to master accounts only)  



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
	subAcct := "subAcct_example" // string | Sub-account name (default to "")
	ccy := "ccy_example" // string | Single currency or multiple currencies (no more than 20) separated with comma, e.g. `BTC` or `BTC,ETH`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetAssetSubaccountBalancesV5(context.Background()).SubAcct(subAcct).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetAssetSubaccountBalancesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetSubaccountBalancesV5`: GetAssetSubaccountBalancesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetAssetSubaccountBalancesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetSubaccountBalancesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subAcct** | **string** | Sub-account name | [default to &quot;&quot;]
 **ccy** | **string** | Single currency or multiple currencies (no more than 20) separated with comma, e.g. &#x60;BTC&#x60; or &#x60;BTC,ETH&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAssetSubaccountBalancesV5Resp**](GetAssetSubaccountBalancesV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetSubaccountBillsV5

> GetAssetSubaccountBillsV5Resp GetAssetSubaccountBillsV5(ctx).Ccy(ccy).Type_(type_).SubAcct(subAcct).After(after).Before(before).Limit(limit).Execute()

Applies to master accounts only.  



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
	ccy := "ccy_example" // string | Currency, such as BTC (optional) (default to "")
	type_ := "type__example" // string | Transfer type  `0`: Transfers from master account to sub-account  `1` : Transfers from sub-account to master account. (optional) (default to "")
	subAcct := "subAcct_example" // string | Sub-account name (optional) (default to "")
	after := "after_example" // string | Query the data prior to the requested bill ID creation time (exclude), the value should be a Unix timestamp in millisecond format. e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Query the data after the requested bill ID creation time (exclude), the value should be a Unix timestamp in millisecond format. e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetAssetSubaccountBillsV5(context.Background()).Ccy(ccy).Type_(type_).SubAcct(subAcct).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetAssetSubaccountBillsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetSubaccountBillsV5`: GetAssetSubaccountBillsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetAssetSubaccountBillsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetSubaccountBillsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency, such as BTC | [default to &quot;&quot;]
 **type_** | **string** | Transfer type  &#x60;0&#x60;: Transfers from master account to sub-account  &#x60;1&#x60; : Transfers from sub-account to master account. | [default to &quot;&quot;]
 **subAcct** | **string** | Sub-account name | [default to &quot;&quot;]
 **after** | **string** | Query the data prior to the requested bill ID creation time (exclude), the value should be a Unix timestamp in millisecond format. e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Query the data after the requested bill ID creation time (exclude), the value should be a Unix timestamp in millisecond format. e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetAssetSubaccountBillsV5Resp**](GetAssetSubaccountBillsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetSubaccountManagedSubaccountBillsV5

> GetAssetSubaccountManagedSubaccountBillsV5Resp GetAssetSubaccountManagedSubaccountBillsV5(ctx).Ccy(ccy).Type_(type_).SubAcct(subAcct).SubUid(subUid).After(after).Before(before).Limit(limit).Execute()

Only applicable to the trading team's master account to getting transfer records of managed sub accounts entrusted to oneself.  



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
	ccy := "ccy_example" // string | Currency, e.g `BTC` (optional) (default to "")
	type_ := "type__example" // string | Transfer type  `0`: Transfers from master account to sub-account  `1`: Transfers from sub-account to master account (optional) (default to "")
	subAcct := "subAcct_example" // string | Sub-account name (optional) (default to "")
	subUid := "subUid_example" // string | Sub-account UID (optional) (default to "")
	after := "after_example" // string | Query the data prior to the requested bill ID creation time (exclude), Unix timestamp in millisecond format, e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Query the data after the requested bill ID creation time (exclude), Unix timestamp in millisecond format, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetAssetSubaccountManagedSubaccountBillsV5(context.Background()).Ccy(ccy).Type_(type_).SubAcct(subAcct).SubUid(subUid).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetAssetSubaccountManagedSubaccountBillsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetSubaccountManagedSubaccountBillsV5`: GetAssetSubaccountManagedSubaccountBillsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetAssetSubaccountManagedSubaccountBillsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetSubaccountManagedSubaccountBillsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency, e.g &#x60;BTC&#x60; | [default to &quot;&quot;]
 **type_** | **string** | Transfer type  &#x60;0&#x60;: Transfers from master account to sub-account  &#x60;1&#x60;: Transfers from sub-account to master account | [default to &quot;&quot;]
 **subAcct** | **string** | Sub-account name | [default to &quot;&quot;]
 **subUid** | **string** | Sub-account UID | [default to &quot;&quot;]
 **after** | **string** | Query the data prior to the requested bill ID creation time (exclude), Unix timestamp in millisecond format, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Query the data after the requested bill ID creation time (exclude), Unix timestamp in millisecond format, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetAssetSubaccountManagedSubaccountBillsV5Resp**](GetAssetSubaccountManagedSubaccountBillsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersEntrustSubaccountListV5

> GetUsersEntrustSubaccountListV5Resp GetUsersEntrustSubaccountListV5(ctx).SubAcct(subAcct).Execute()

The trading team uses this interface to view the list of sub-accounts currently under escrow  



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
	subAcct := "subAcct_example" // string | Sub-account name (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetUsersEntrustSubaccountListV5(context.Background()).SubAcct(subAcct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetUsersEntrustSubaccountListV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersEntrustSubaccountListV5`: GetUsersEntrustSubaccountListV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetUsersEntrustSubaccountListV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersEntrustSubaccountListV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subAcct** | **string** | Sub-account name | [default to &quot;&quot;]

### Return type

[**GetUsersEntrustSubaccountListV5Resp**](GetUsersEntrustSubaccountListV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersSubaccountListV5

> GetUsersSubaccountListV5Resp GetUsersSubaccountListV5(ctx).Enable(enable).SubAcct(subAcct).After(after).Before(before).Limit(limit).Execute()

Applies to master accounts only  



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
	enable := "enable_example" // string | Sub-account status   `true`: Normal  `false`: Frozen (optional) (default to "")
	subAcct := "subAcct_example" // string | Sub-account name (optional) (default to "")
	after := "after_example" // string | Query the data earlier than the requested subaccount creation timestamp, the value should be a Unix timestamp in millisecond format. e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Query the data newer than the requested subaccount creation timestamp, the value should be a Unix timestamp in millisecond format. e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubAccountAPI.GetUsersSubaccountListV5(context.Background()).Enable(enable).SubAcct(subAcct).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAccountAPI.GetUsersSubaccountListV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersSubaccountListV5`: GetUsersSubaccountListV5Resp
	fmt.Fprintf(os.Stdout, "Response from `SubAccountAPI.GetUsersSubaccountListV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersSubaccountListV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enable** | **string** | Sub-account status   &#x60;true&#x60;: Normal  &#x60;false&#x60;: Frozen | [default to &quot;&quot;]
 **subAcct** | **string** | Sub-account name | [default to &quot;&quot;]
 **after** | **string** | Query the data earlier than the requested subaccount creation timestamp, the value should be a Unix timestamp in millisecond format. e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Query the data newer than the requested subaccount creation timestamp, the value should be a Unix timestamp in millisecond format. e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetUsersSubaccountListV5Resp**](GetUsersSubaccountListV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

