# CreateAccountMmpConfigV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrozenInterval** | Pointer to **string** | Frozen period (ms). | [optional] [default to ""]
**InstFamily** | Pointer to **string** | Instrument family | [optional] [default to ""]
**QtyLimit** | Pointer to **string** | Trade qty limit in number of contracts | [optional] [default to ""]
**TimeInterval** | Pointer to **string** | Time window (ms). MMP interval where monitoring is done | [optional] [default to ""]

## Methods

### NewCreateAccountMmpConfigV5RespDataInner

`func NewCreateAccountMmpConfigV5RespDataInner() *CreateAccountMmpConfigV5RespDataInner`

NewCreateAccountMmpConfigV5RespDataInner instantiates a new CreateAccountMmpConfigV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountMmpConfigV5RespDataInnerWithDefaults

`func NewCreateAccountMmpConfigV5RespDataInnerWithDefaults() *CreateAccountMmpConfigV5RespDataInner`

NewCreateAccountMmpConfigV5RespDataInnerWithDefaults instantiates a new CreateAccountMmpConfigV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrozenInterval

`func (o *CreateAccountMmpConfigV5RespDataInner) GetFrozenInterval() string`

GetFrozenInterval returns the FrozenInterval field if non-nil, zero value otherwise.

### GetFrozenIntervalOk

`func (o *CreateAccountMmpConfigV5RespDataInner) GetFrozenIntervalOk() (*string, bool)`

GetFrozenIntervalOk returns a tuple with the FrozenInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenInterval

`func (o *CreateAccountMmpConfigV5RespDataInner) SetFrozenInterval(v string)`

SetFrozenInterval sets FrozenInterval field to given value.

### HasFrozenInterval

`func (o *CreateAccountMmpConfigV5RespDataInner) HasFrozenInterval() bool`

HasFrozenInterval returns a boolean if a field has been set.

### GetInstFamily

`func (o *CreateAccountMmpConfigV5RespDataInner) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *CreateAccountMmpConfigV5RespDataInner) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *CreateAccountMmpConfigV5RespDataInner) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *CreateAccountMmpConfigV5RespDataInner) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetQtyLimit

`func (o *CreateAccountMmpConfigV5RespDataInner) GetQtyLimit() string`

GetQtyLimit returns the QtyLimit field if non-nil, zero value otherwise.

### GetQtyLimitOk

`func (o *CreateAccountMmpConfigV5RespDataInner) GetQtyLimitOk() (*string, bool)`

GetQtyLimitOk returns a tuple with the QtyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyLimit

`func (o *CreateAccountMmpConfigV5RespDataInner) SetQtyLimit(v string)`

SetQtyLimit sets QtyLimit field to given value.

### HasQtyLimit

`func (o *CreateAccountMmpConfigV5RespDataInner) HasQtyLimit() bool`

HasQtyLimit returns a boolean if a field has been set.

### GetTimeInterval

`func (o *CreateAccountMmpConfigV5RespDataInner) GetTimeInterval() string`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *CreateAccountMmpConfigV5RespDataInner) GetTimeIntervalOk() (*string, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *CreateAccountMmpConfigV5RespDataInner) SetTimeInterval(v string)`

SetTimeInterval sets TimeInterval field to given value.

### HasTimeInterval

`func (o *CreateAccountMmpConfigV5RespDataInner) HasTimeInterval() bool`

HasTimeInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


