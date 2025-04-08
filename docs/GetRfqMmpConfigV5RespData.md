# GetRfqMmpConfigV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountLimit** | Pointer to **string** | Limit in number of execution attempts | [optional] [default to ""]
**FrozenInterval** | Pointer to **string** | Frozen period (ms). If it is \&quot;0\&quot;, the trade will remain frozen until manually reset and &#x60;mmpFrozenUntil&#x60; will be \&quot;\&quot;. | [optional] [default to ""]
**MmpFrozen** | Pointer to **bool** | Whether MMP is currently triggered. &#x60;true&#x60; or &#x60;false&#x60; | [optional] 
**MmpFrozenUntil** | Pointer to **string** | If frozenInterval is not \&quot;0\&quot; and mmpFrozen &#x3D; True, it is the time interval (in ms) when MMP is no longer triggered, otherwise \&quot;\&quot; | [optional] [default to ""]
**TimeInterval** | Pointer to **string** | Time window (ms). MMP interval where monitoring is done  \&quot;0\&quot; means MMP is diabled | [optional] [default to ""]

## Methods

### NewGetRfqMmpConfigV5RespData

`func NewGetRfqMmpConfigV5RespData() *GetRfqMmpConfigV5RespData`

NewGetRfqMmpConfigV5RespData instantiates a new GetRfqMmpConfigV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRfqMmpConfigV5RespDataWithDefaults

`func NewGetRfqMmpConfigV5RespDataWithDefaults() *GetRfqMmpConfigV5RespData`

NewGetRfqMmpConfigV5RespDataWithDefaults instantiates a new GetRfqMmpConfigV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountLimit

`func (o *GetRfqMmpConfigV5RespData) GetCountLimit() string`

GetCountLimit returns the CountLimit field if non-nil, zero value otherwise.

### GetCountLimitOk

`func (o *GetRfqMmpConfigV5RespData) GetCountLimitOk() (*string, bool)`

GetCountLimitOk returns a tuple with the CountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountLimit

`func (o *GetRfqMmpConfigV5RespData) SetCountLimit(v string)`

SetCountLimit sets CountLimit field to given value.

### HasCountLimit

`func (o *GetRfqMmpConfigV5RespData) HasCountLimit() bool`

HasCountLimit returns a boolean if a field has been set.

### GetFrozenInterval

`func (o *GetRfqMmpConfigV5RespData) GetFrozenInterval() string`

GetFrozenInterval returns the FrozenInterval field if non-nil, zero value otherwise.

### GetFrozenIntervalOk

`func (o *GetRfqMmpConfigV5RespData) GetFrozenIntervalOk() (*string, bool)`

GetFrozenIntervalOk returns a tuple with the FrozenInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenInterval

`func (o *GetRfqMmpConfigV5RespData) SetFrozenInterval(v string)`

SetFrozenInterval sets FrozenInterval field to given value.

### HasFrozenInterval

`func (o *GetRfqMmpConfigV5RespData) HasFrozenInterval() bool`

HasFrozenInterval returns a boolean if a field has been set.

### GetMmpFrozen

`func (o *GetRfqMmpConfigV5RespData) GetMmpFrozen() bool`

GetMmpFrozen returns the MmpFrozen field if non-nil, zero value otherwise.

### GetMmpFrozenOk

`func (o *GetRfqMmpConfigV5RespData) GetMmpFrozenOk() (*bool, bool)`

GetMmpFrozenOk returns a tuple with the MmpFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmpFrozen

`func (o *GetRfqMmpConfigV5RespData) SetMmpFrozen(v bool)`

SetMmpFrozen sets MmpFrozen field to given value.

### HasMmpFrozen

`func (o *GetRfqMmpConfigV5RespData) HasMmpFrozen() bool`

HasMmpFrozen returns a boolean if a field has been set.

### GetMmpFrozenUntil

`func (o *GetRfqMmpConfigV5RespData) GetMmpFrozenUntil() string`

GetMmpFrozenUntil returns the MmpFrozenUntil field if non-nil, zero value otherwise.

### GetMmpFrozenUntilOk

`func (o *GetRfqMmpConfigV5RespData) GetMmpFrozenUntilOk() (*string, bool)`

GetMmpFrozenUntilOk returns a tuple with the MmpFrozenUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmpFrozenUntil

`func (o *GetRfqMmpConfigV5RespData) SetMmpFrozenUntil(v string)`

SetMmpFrozenUntil sets MmpFrozenUntil field to given value.

### HasMmpFrozenUntil

`func (o *GetRfqMmpConfigV5RespData) HasMmpFrozenUntil() bool`

HasMmpFrozenUntil returns a boolean if a field has been set.

### GetTimeInterval

`func (o *GetRfqMmpConfigV5RespData) GetTimeInterval() string`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *GetRfqMmpConfigV5RespData) GetTimeIntervalOk() (*string, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *GetRfqMmpConfigV5RespData) SetTimeInterval(v string)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *GetRfqMmpConfigV5RespData) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


