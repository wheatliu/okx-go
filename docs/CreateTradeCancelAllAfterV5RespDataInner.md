# CreateTradeCancelAllAfterV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **string** | CAA order tag | [optional] [default to ""]
**TriggerTime** | Pointer to **string** | The time the cancellation is triggered.  triggerTime&#x3D;0 means Cancel All After is disabled. | [optional] [default to ""]
**Ts** | Pointer to **string** | The time the request is received. | [optional] [default to ""]

## Methods

### NewCreateTradeCancelAllAfterV5RespDataInner

`func NewCreateTradeCancelAllAfterV5RespDataInner() *CreateTradeCancelAllAfterV5RespDataInner`

NewCreateTradeCancelAllAfterV5RespDataInner instantiates a new CreateTradeCancelAllAfterV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeCancelAllAfterV5RespDataInnerWithDefaults

`func NewCreateTradeCancelAllAfterV5RespDataInnerWithDefaults() *CreateTradeCancelAllAfterV5RespDataInner`

NewCreateTradeCancelAllAfterV5RespDataInnerWithDefaults instantiates a new CreateTradeCancelAllAfterV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *CreateTradeCancelAllAfterV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateTradeCancelAllAfterV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateTradeCancelAllAfterV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateTradeCancelAllAfterV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTriggerTime

`func (o *CreateTradeCancelAllAfterV5RespDataInner) GetTriggerTime() string`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *CreateTradeCancelAllAfterV5RespDataInner) GetTriggerTimeOk() (*string, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *CreateTradeCancelAllAfterV5RespDataInner) SetTriggerTime(v string)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *CreateTradeCancelAllAfterV5RespDataInner) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetTs

`func (o *CreateTradeCancelAllAfterV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CreateTradeCancelAllAfterV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CreateTradeCancelAllAfterV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CreateTradeCancelAllAfterV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


