# CreateRfqMmpConfigV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountLimit** | **string** | Limit in number of execution attempts. | [default to ""]
**FrozenInterval** | **string** | Frozen period (ms).   \&quot;0\&quot; means the trade will remain frozen until you request \&quot;Reset MMP Status\&quot; to unfrozen. | [default to ""]
**TimeInterval** | **string** | Time window (ms). MMP interval where monitoring is done.  \&quot;0\&quot; means disable MMP. Maximum time interval is 600,000. | [default to ""]

## Methods

### NewCreateRfqMmpConfigV5Req

`func NewCreateRfqMmpConfigV5Req(countLimit string, frozenInterval string, timeInterval string, ) *CreateRfqMmpConfigV5Req`

NewCreateRfqMmpConfigV5Req instantiates a new CreateRfqMmpConfigV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqMmpConfigV5ReqWithDefaults

`func NewCreateRfqMmpConfigV5ReqWithDefaults() *CreateRfqMmpConfigV5Req`

NewCreateRfqMmpConfigV5ReqWithDefaults instantiates a new CreateRfqMmpConfigV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountLimit

`func (o *CreateRfqMmpConfigV5Req) GetCountLimit() string`

GetCountLimit returns the CountLimit field if non-nil, zero value otherwise.

### GetCountLimitOk

`func (o *CreateRfqMmpConfigV5Req) GetCountLimitOk() (*string, bool)`

GetCountLimitOk returns a tuple with the CountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountLimit

`func (o *CreateRfqMmpConfigV5Req) SetCountLimit(v string)`

SetCountLimit sets CountLimit field to given value.


### GetFrozenInterval

`func (o *CreateRfqMmpConfigV5Req) GetFrozenInterval() string`

GetFrozenInterval returns the FrozenInterval field if non-nil, zero value otherwise.

### GetFrozenIntervalOk

`func (o *CreateRfqMmpConfigV5Req) GetFrozenIntervalOk() (*string, bool)`

GetFrozenIntervalOk returns a tuple with the FrozenInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenInterval

`func (o *CreateRfqMmpConfigV5Req) SetFrozenInterval(v string)`

SetFrozenInterval sets FrozenInterval field to given value.


### GetTimeInterval

`func (o *CreateRfqMmpConfigV5Req) GetTimeInterval() string`

GetTimeInterval returns the TimeInterval field if non-nil, zero value otherwise.

### GetTimeIntervalOk

`func (o *CreateRfqMmpConfigV5Req) GetTimeIntervalOk() (*string, bool)`

GetTimeIntervalOk returns a tuple with the TimeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterval

`func (o *CreateRfqMmpConfigV5Req) SetTimeInterval(v string)`

SetTimeInterval sets TimeInterval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


