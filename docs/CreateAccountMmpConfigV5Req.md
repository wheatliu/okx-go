# CreateAccountMmpConfigV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrozenInterval** | **string** | Frozen period (ms).   \&quot;0\&quot; means the trade will remain frozen until you request \&quot;Reset MMP Status\&quot; to unfrozen | [default to ""]
**InstFamily** | **string** | Instrument family | [default to ""]
**QtyLimit** | **string** | Trade qty limit in number of contracts  Must be &gt; 0 | [default to ""]
**TimeInterval** | **string** | Time window (ms). MMP interval where monitoring is done  \&quot;0\&quot; means disable MMP | [default to ""]

## Methods

### NewCreateAccountMmpConfigV5Req

`func NewCreateAccountMmpConfigV5Req(frozenInterval string, instFamily string, qtyLimit string, timeInterval string, ) *CreateAccountMmpConfigV5Req`

NewCreateAccountMmpConfigV5Req instantiates a new CreateAccountMmpConfigV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountMmpConfigV5ReqWithDefaults

`func NewCreateAccountMmpConfigV5ReqWithDefaults() *CreateAccountMmpConfigV5Req`

NewCreateAccountMmpConfigV5ReqWithDefaults instantiates a new CreateAccountMmpConfigV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrozenInterval

`func (o *CreateAccountMmpConfigV5Req) GetFrozenInterval() string`

GetFrozenInterval returns the FrozenInterval field if non-nil, zero value otherwise.

### GetFrozenIntervalOk

`func (o *CreateAccountMmpConfigV5Req) GetFrozenIntervalOk() (*string, bool)`

GetFrozenIntervalOk returns a tuple with the FrozenInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenInterval

`func (o *CreateAccountMmpConfigV5Req) SetFrozenInterval(v string)`

SetFrozenInterval sets FrozenInterval field to given value.


### GetInstFamily

`func (o *CreateAccountMmpConfigV5Req) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *CreateAccountMmpConfigV5Req) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *CreateAccountMmpConfigV5Req) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.


### GetQtyLimit

`func (o *CreateAccountMmpConfigV5Req) GetQtyLimit() string`

GetQtyLimit returns the QtyLimit field if non-nil, zero value otherwise.

### GetQtyLimitOk

`func (o *CreateAccountMmpConfigV5Req) GetQtyLimitOk() (*string, bool)`

GetQtyLimitOk returns a tuple with the QtyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyLimit

`func (o *CreateAccountMmpConfigV5Req) SetQtyLimit(v string)`

SetQtyLimit sets QtyLimit field to given value.


### GetTimeInterval

`func (o *CreateAccountMmpConfigV5Req) GetTimeInterval() string`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *CreateAccountMmpConfigV5Req) GetTimeIntervalOk() (*string, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *CreateAccountMmpConfigV5Req) SetTimeInterval(v string)`

SetTimeInterval sets TimeInterval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


