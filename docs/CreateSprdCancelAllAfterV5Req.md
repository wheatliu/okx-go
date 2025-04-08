# CreateSprdCancelAllAfterV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeOut** | **string** | The countdown for order cancellation, with second as the unit.  Range of value can be 0, [10, 120].   Setting timeOut to 0 disables Cancel All After. | [default to ""]

## Methods

### NewCreateSprdCancelAllAfterV5Req

`func NewCreateSprdCancelAllAfterV5Req(timeOut string, ) *CreateSprdCancelAllAfterV5Req`

NewCreateSprdCancelAllAfterV5Req instantiates a new CreateSprdCancelAllAfterV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSprdCancelAllAfterV5ReqWithDefaults

`func NewCreateSprdCancelAllAfterV5ReqWithDefaults() *CreateSprdCancelAllAfterV5Req`

NewCreateSprdCancelAllAfterV5ReqWithDefaults instantiates a new CreateSprdCancelAllAfterV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeOut

`func (o *CreateSprdCancelAllAfterV5Req) GetTimeOut() string`

GetTimeOut returns the TimeOut field if non-nil, zero value otherwise.

### GetTimeOutOk

`func (o *CreateSprdCancelAllAfterV5Req) GetTimeOutOk() (*string, bool)`

GetTimeOutOk returns a tuple with the TimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOut

`func (o *CreateSprdCancelAllAfterV5Req) SetTimeOut(v string)`

SetTimeOut sets TimeOut field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


