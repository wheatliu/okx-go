# CreateRfqMmpConfigV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountLimit** | Pointer to **string** | Limit in number of execution attempts | [optional] [default to ""]
**FrozenInterval** | Pointer to **string** | Frozen period (ms). | [optional] [default to ""]
**TimeInterval** | Pointer to **string** | Time window (ms). MMP interval where monitoring is done | [optional] [default to ""]

## Methods

### NewCreateRfqMmpConfigV5RespData

`func NewCreateRfqMmpConfigV5RespData() *CreateRfqMmpConfigV5RespData`

NewCreateRfqMmpConfigV5RespData instantiates a new CreateRfqMmpConfigV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqMmpConfigV5RespDataWithDefaults

`func NewCreateRfqMmpConfigV5RespDataWithDefaults() *CreateRfqMmpConfigV5RespData`

NewCreateRfqMmpConfigV5RespDataWithDefaults instantiates a new CreateRfqMmpConfigV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountLimit

`func (o *CreateRfqMmpConfigV5RespData) GetCountLimit() string`

GetCountLimit returns the CountLimit field if non-nil, zero value otherwise.

### GetCountLimitOk

`func (o *CreateRfqMmpConfigV5RespData) GetCountLimitOk() (*string, bool)`

GetCountLimitOk returns a tuple with the CountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountLimit

`func (o *CreateRfqMmpConfigV5RespData) SetCountLimit(v string)`

SetCountLimit sets CountLimit field to given value.

### HasCountLimit

`func (o *CreateRfqMmpConfigV5RespData) HasCountLimit() bool`

HasCountLimit returns a boolean if a field has been set.

### GetFrozenInterval

`func (o *CreateRfqMmpConfigV5RespData) GetFrozenInterval() string`

GetFrozenInterval returns the FrozenInterval field if non-nil, zero value otherwise.

### GetFrozenIntervalOk

`func (o *CreateRfqMmpConfigV5RespData) GetFrozenIntervalOk() (*string, bool)`

GetFrozenIntervalOk returns a tuple with the FrozenInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenInterval

`func (o *CreateRfqMmpConfigV5RespData) SetFrozenInterval(v string)`

SetFrozenInterval sets FrozenInterval field to given value.

### HasFrozenInterval

`func (o *CreateRfqMmpConfigV5RespData) HasFrozenInterval() bool`

HasFrozenInterval returns a boolean if a field has been set.

### GetTimeInterval

`func (o *CreateRfqMmpConfigV5RespData) GetTimeInterval() string`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *CreateRfqMmpConfigV5RespData) GetTimeIntervalOk() (*string, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *CreateRfqMmpConfigV5RespData) SetTimeInterval(v string)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *CreateRfqMmpConfigV5RespData) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


