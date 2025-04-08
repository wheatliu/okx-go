# CreateTradeCancelAllAfterV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **string** | CAA order tag   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters. | [optional] [default to ""]
**TimeOut** | **string** | The countdown for order cancellation, with second as the unit.  Range of value can be 0, [10, 120].   Setting timeOut to 0 disables Cancel All After. | [default to ""]

## Methods

### NewCreateTradeCancelAllAfterV5Req

`func NewCreateTradeCancelAllAfterV5Req(timeOut string, ) *CreateTradeCancelAllAfterV5Req`

NewCreateTradeCancelAllAfterV5Req instantiates a new CreateTradeCancelAllAfterV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeCancelAllAfterV5ReqWithDefaults

`func NewCreateTradeCancelAllAfterV5ReqWithDefaults() *CreateTradeCancelAllAfterV5Req`

NewCreateTradeCancelAllAfterV5ReqWithDefaults instantiates a new CreateTradeCancelAllAfterV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *CreateTradeCancelAllAfterV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateTradeCancelAllAfterV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateTradeCancelAllAfterV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateTradeCancelAllAfterV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTimeOut

`func (o *CreateTradeCancelAllAfterV5Req) GetTimeOut() string`

GetTimeOut returns the TimeOut field if non-nil, zero value otherwise.

### GetTimeOutOk

`func (o *CreateTradeCancelAllAfterV5Req) GetTimeOutOk() (*string, bool)`

GetTimeOutOk returns a tuple with the TimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOut

`func (o *CreateTradeCancelAllAfterV5Req) SetTimeOut(v string)`

SetTimeOut sets TimeOut field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


