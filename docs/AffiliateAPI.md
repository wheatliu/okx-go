# \AffiliateAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAffiliateInviteeDetailV5**](AffiliateAPI.md#GetAffiliateInviteeDetailV5) | **Get** /api/v5/affiliate/invitee/detail | Get the invitee&#39;s detail
[**GetSupportAnnouncementTypesV5**](AffiliateAPI.md#GetSupportAnnouncementTypesV5) | **Get** /api/v5/support/announcement-types | Get the user&#39;s affiliate rebate information
[**GetUsersPartnerIfRebateV5**](AffiliateAPI.md#GetUsersPartnerIfRebateV5) | **Get** /api/v5/users/partner/if-rebate | Get the user&#39;s affiliate rebate information



## GetAffiliateInviteeDetailV5

> GetAffiliateInviteeDetailV5Resp GetAffiliateInviteeDetailV5(ctx).Uid(uid).Execute()

Get the invitee's detail



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
	uid := "uid_example" // string | UID of the invitee. Only applicable to the UID of invitee master account.    The data returned covers invitee master account and invitee sub-accounts. (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliateAPI.GetAffiliateInviteeDetailV5(context.Background()).Uid(uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.GetAffiliateInviteeDetailV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAffiliateInviteeDetailV5`: GetAffiliateInviteeDetailV5Resp
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.GetAffiliateInviteeDetailV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAffiliateInviteeDetailV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **string** | UID of the invitee. Only applicable to the UID of invitee master account.    The data returned covers invitee master account and invitee sub-accounts. | [default to &quot;&quot;]

### Return type

[**GetAffiliateInviteeDetailV5Resp**](GetAffiliateInviteeDetailV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportAnnouncementTypesV5

> GetSupportAnnouncementTypesV5Resp GetSupportAnnouncementTypesV5(ctx).ApiKey(apiKey).Execute()

Get the user's affiliate rebate information



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
	apiKey := "apiKey_example" // string | The user's API key. Only applicable to the API key of invitee master account (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliateAPI.GetSupportAnnouncementTypesV5(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.GetSupportAnnouncementTypesV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportAnnouncementTypesV5`: GetSupportAnnouncementTypesV5Resp
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.GetSupportAnnouncementTypesV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportAnnouncementTypesV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | The user&#39;s API key. Only applicable to the API key of invitee master account | [default to &quot;&quot;]

### Return type

[**GetSupportAnnouncementTypesV5Resp**](GetSupportAnnouncementTypesV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersPartnerIfRebateV5

> GetUsersPartnerIfRebateV5Resp GetUsersPartnerIfRebateV5(ctx).ApiKey(apiKey).Execute()

Get the user's affiliate rebate information



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
	apiKey := "apiKey_example" // string | The user's API key. Only applicable to the API key of invitee master account (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliateAPI.GetUsersPartnerIfRebateV5(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.GetUsersPartnerIfRebateV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersPartnerIfRebateV5`: GetUsersPartnerIfRebateV5Resp
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.GetUsersPartnerIfRebateV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersPartnerIfRebateV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **string** | The user&#39;s API key. Only applicable to the API key of invitee master account | [default to &quot;&quot;]

### Return type

[**GetUsersPartnerIfRebateV5Resp**](GetUsersPartnerIfRebateV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

