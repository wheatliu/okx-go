# CreateUsersSubaccountModifyApikeyV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | Sub-account APIKey | [default to ""]
**Ip** | Pointer to **string** | Sub-account API Key linked IP addresses, separate with commas if more than one. Support up to 20 IP addresses.  The IP will be reset if this is passed through.  If &#x60;ip&#x60; is set to \&quot;\&quot;, then no IP addresses is linked to the APIKey. | [optional] [default to ""]
**Label** | Pointer to **string** | Sub-account API Key label. The label will be reset if this is passed through. | [optional] [default to ""]
**Perm** | Pointer to **string** | Sub-account API Key permissions  &#x60;read_only&#x60;: Read  &#x60;trade&#x60;: Trade  Separate with commas if more than one.   The permission will be reset if this is passed through. | [optional] [default to ""]
**SubAcct** | **string** | Sub-account name | [default to ""]

## Methods

### NewCreateUsersSubaccountModifyApikeyV5Req

`func NewCreateUsersSubaccountModifyApikeyV5Req(apiKey string, subAcct string, ) *CreateUsersSubaccountModifyApikeyV5Req`

NewCreateUsersSubaccountModifyApikeyV5Req instantiates a new CreateUsersSubaccountModifyApikeyV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUsersSubaccountModifyApikeyV5ReqWithDefaults

`func NewCreateUsersSubaccountModifyApikeyV5ReqWithDefaults() *CreateUsersSubaccountModifyApikeyV5Req`

NewCreateUsersSubaccountModifyApikeyV5ReqWithDefaults instantiates a new CreateUsersSubaccountModifyApikeyV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *CreateUsersSubaccountModifyApikeyV5Req) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateUsersSubaccountModifyApikeyV5Req) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateUsersSubaccountModifyApikeyV5Req) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetIp

`func (o *CreateUsersSubaccountModifyApikeyV5Req) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CreateUsersSubaccountModifyApikeyV5Req) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CreateUsersSubaccountModifyApikeyV5Req) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CreateUsersSubaccountModifyApikeyV5Req) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLabel

`func (o *CreateUsersSubaccountModifyApikeyV5Req) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateUsersSubaccountModifyApikeyV5Req) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateUsersSubaccountModifyApikeyV5Req) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateUsersSubaccountModifyApikeyV5Req) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPerm

`func (o *CreateUsersSubaccountModifyApikeyV5Req) GetPerm() string`

GetPerm returns the Perm field if non-nil, zero value otherwise.

### GetPermOk

`func (o *CreateUsersSubaccountModifyApikeyV5Req) GetPermOk() (*string, bool)`

GetPermOk returns a tuple with the Perm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerm

`func (o *CreateUsersSubaccountModifyApikeyV5Req) SetPerm(v string)`

SetPerm sets Perm field to given value.

### HasPerm

`func (o *CreateUsersSubaccountModifyApikeyV5Req) HasPerm() bool`

HasPerm returns a boolean if a field has been set.

### GetSubAcct

`func (o *CreateUsersSubaccountModifyApikeyV5Req) GetSubAcct() string`

GetSubAcct returns the SubAcct field if non-nil, zero value otherwise.

### GetSubAcctOk

`func (o *CreateUsersSubaccountModifyApikeyV5Req) GetSubAcctOk() (*string, bool)`

GetSubAcctOk returns a tuple with the SubAcct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAcct

`func (o *CreateUsersSubaccountModifyApikeyV5Req) SetSubAcct(v string)`

SetSubAcct sets SubAcct field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


