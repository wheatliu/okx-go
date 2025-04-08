# GetUsersPartnerIfRebateV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **bool** | Whether the user is invited by the current affiliate. &#x60;true&#x60;, &#x60;false&#x60; | [optional] 
**Type** | Pointer to **string** | Whether there is affiliate rebate.  &#x60;0&#x60; There is affiliate rebate  &#x60;1&#x60; There is no affiliate rebate. Because the account which is requesting this endpoint is not affiliate   &#x60;2&#x60; There is no affiliate rebate. Because there is no relationship of invitation or recall, e.g. api key does not exist   &#x60;4&#x60; There is no affiliate rebate. Because the user level is equal to or more than VIP6 | [optional] [default to ""]

## Methods

### NewGetUsersPartnerIfRebateV5RespDataInner

`func NewGetUsersPartnerIfRebateV5RespDataInner() *GetUsersPartnerIfRebateV5RespDataInner`

NewGetUsersPartnerIfRebateV5RespDataInner instantiates a new GetUsersPartnerIfRebateV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsersPartnerIfRebateV5RespDataInnerWithDefaults

`func NewGetUsersPartnerIfRebateV5RespDataInnerWithDefaults() *GetUsersPartnerIfRebateV5RespDataInner`

NewGetUsersPartnerIfRebateV5RespDataInnerWithDefaults instantiates a new GetUsersPartnerIfRebateV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetUsersPartnerIfRebateV5RespDataInner) GetResult() bool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetUsersPartnerIfRebateV5RespDataInner) GetResultOk() (*bool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetUsersPartnerIfRebateV5RespDataInner) SetResult(v bool)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetUsersPartnerIfRebateV5RespDataInner) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetType

`func (o *GetUsersPartnerIfRebateV5RespDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetUsersPartnerIfRebateV5RespDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetUsersPartnerIfRebateV5RespDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetUsersPartnerIfRebateV5RespDataInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


