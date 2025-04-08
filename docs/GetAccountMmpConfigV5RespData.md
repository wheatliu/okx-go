# GetAccountMmpConfigV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrozenInterval** | Pointer to **string** | Frozen period (ms). If it is \&quot;0\&quot;, the trade will remain frozen until manually reset and &#x60;mmpFrozenUntil&#x60; will be \&quot;\&quot;. | [optional] [default to ""]
**InstFamily** | Pointer to **string** | Instrument Family | [optional] [default to ""]
**MmpFrozen** | Pointer to **bool** | Whether MMP is currently triggered. &#x60;true&#x60; or &#x60;false&#x60; | [optional] 
**MmpFrozenUntil** | Pointer to **string** | If frozenInterval is configured and mmpFrozen &#x3D; True, it is the time interval (in ms) when MMP is no longer triggered, otherwise \&quot;\&quot;. | [optional] [default to ""]
**QtyLimit** | Pointer to **string** | Trade qty limit in number of contracts | [optional] [default to ""]
**TimeInterval** | Pointer to **string** | Time window (ms). MMP interval where monitoring is done | [optional] [default to ""]

## Methods

### NewGetAccountMmpConfigV5RespData

`func NewGetAccountMmpConfigV5RespData() *GetAccountMmpConfigV5RespData`

NewGetAccountMmpConfigV5RespData instantiates a new GetAccountMmpConfigV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountMmpConfigV5RespDataWithDefaults

`func NewGetAccountMmpConfigV5RespDataWithDefaults() *GetAccountMmpConfigV5RespData`

NewGetAccountMmpConfigV5RespDataWithDefaults instantiates a new GetAccountMmpConfigV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrozenInterval

`func (o *GetAccountMmpConfigV5RespData) GetFrozenInterval() string`

GetFrozenInterval returns the FrozenInterval field if non-nil, zero value otherwise.

### GetFrozenIntervalOk

`func (o *GetAccountMmpConfigV5RespData) GetFrozenIntervalOk() (*string, bool)`

GetFrozenIntervalOk returns a tuple with the FrozenInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenInterval

`func (o *GetAccountMmpConfigV5RespData) SetFrozenInterval(v string)`

SetFrozenInterval sets FrozenInterval field to given value.

### HasFrozenInterval

`func (o *GetAccountMmpConfigV5RespData) HasFrozenInterval() bool`

HasFrozenInterval returns a boolean if a field has been set.

### GetInstFamily

`func (o *GetAccountMmpConfigV5RespData) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetAccountMmpConfigV5RespData) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetAccountMmpConfigV5RespData) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetAccountMmpConfigV5RespData) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetMmpFrozen

`func (o *GetAccountMmpConfigV5RespData) GetMmpFrozen() bool`

GetMmpFrozen returns the MmpFrozen field if non-nil, zero value otherwise.

### GetMmpFrozenOk

`func (o *GetAccountMmpConfigV5RespData) GetMmpFrozenOk() (*bool, bool)`

GetMmpFrozenOk returns a tuple with the MmpFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmpFrozen

`func (o *GetAccountMmpConfigV5RespData) SetMmpFrozen(v bool)`

SetMmpFrozen sets MmpFrozen field to given value.

### HasMmpFrozen

`func (o *GetAccountMmpConfigV5RespData) HasMmpFrozen() bool`

HasMmpFrozen returns a boolean if a field has been set.

### GetMmpFrozenUntil

`func (o *GetAccountMmpConfigV5RespData) GetMmpFrozenUntil() string`

GetMmpFrozenUntil returns the MmpFrozenUntil field if non-nil, zero value otherwise.

### GetMmpFrozenUntilOk

`func (o *GetAccountMmpConfigV5RespData) GetMmpFrozenUntilOk() (*string, bool)`

GetMmpFrozenUntilOk returns a tuple with the MmpFrozenUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmpFrozenUntil

`func (o *GetAccountMmpConfigV5RespData) SetMmpFrozenUntil(v string)`

SetMmpFrozenUntil sets MmpFrozenUntil field to given value.

### HasMmpFrozenUntil

`func (o *GetAccountMmpConfigV5RespData) HasMmpFrozenUntil() bool`

HasMmpFrozenUntil returns a boolean if a field has been set.

### GetQtyLimit

`func (o *GetAccountMmpConfigV5RespData) GetQtyLimit() string`

GetQtyLimit returns the QtyLimit field if non-nil, zero value otherwise.

### GetQtyLimitOk

`func (o *GetAccountMmpConfigV5RespData) GetQtyLimitOk() (*string, bool)`

GetQtyLimitOk returns a tuple with the QtyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyLimit

`func (o *GetAccountMmpConfigV5RespData) SetQtyLimit(v string)`

SetQtyLimit sets QtyLimit field to given value.

### HasQtyLimit

`func (o *GetAccountMmpConfigV5RespData) HasQtyLimit() bool`

HasQtyLimit returns a boolean if a field has been set.

### GetTimeInterval

`func (o *GetAccountMmpConfigV5RespData) GetTimeInterval() string`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *GetAccountMmpConfigV5RespData) GetTimeIntervalOk() (*string, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *GetAccountMmpConfigV5RespData) SetTimeInterval(v string)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *GetAccountMmpConfigV5RespData) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


