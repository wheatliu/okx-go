# CreateRfqCancelAllAfterV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerTime** | Pointer to **string** | The time the cancellation is triggered.  triggerTime&#x3D;0 means Cancel All After is disabled. | [optional] [default to ""]
**Ts** | Pointer to **string** | The time the request is received. | [optional] [default to ""]

## Methods

### NewCreateRfqCancelAllAfterV5RespData

`func NewCreateRfqCancelAllAfterV5RespData() *CreateRfqCancelAllAfterV5RespData`

NewCreateRfqCancelAllAfterV5RespData instantiates a new CreateRfqCancelAllAfterV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCancelAllAfterV5RespDataWithDefaults

`func NewCreateRfqCancelAllAfterV5RespDataWithDefaults() *CreateRfqCancelAllAfterV5RespData`

NewCreateRfqCancelAllAfterV5RespDataWithDefaults instantiates a new CreateRfqCancelAllAfterV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerTime

`func (o *CreateRfqCancelAllAfterV5RespData) GetTriggerTime() string`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *CreateRfqCancelAllAfterV5RespData) GetTriggerTimeOk() (*string, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *CreateRfqCancelAllAfterV5RespData) SetTriggerTime(v string)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *CreateRfqCancelAllAfterV5RespData) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetTs

`func (o *CreateRfqCancelAllAfterV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CreateRfqCancelAllAfterV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CreateRfqCancelAllAfterV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CreateRfqCancelAllAfterV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


