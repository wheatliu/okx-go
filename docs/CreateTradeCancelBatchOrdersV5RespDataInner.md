# CreateTradeCancelBatchOrdersV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The result code, &#x60;0&#x60; means success | [optional] [default to ""]
**Data** | Pointer to [**[]CreateTradeCancelBatchOrdersV5RespDataInnerDataInner**](CreateTradeCancelBatchOrdersV5RespDataInnerDataInner.md) | Array of objects contains the response results | [optional] 
**InTime** | Pointer to **string** | Timestamp at REST gateway when the request is received, Unix timestamp format in microseconds, e.g. &#x60;1597026383085123&#x60;   The time is recorded after authentication. | [optional] [default to ""]
**Msg** | Pointer to **string** | The error message, empty if the code is 0 | [optional] [default to ""]
**OutTime** | Pointer to **string** | Timestamp at REST gateway when the response is sent, Unix timestamp format in microseconds, e.g. &#x60;1597026383085123&#x60; | [optional] [default to ""]

## Methods

### NewCreateTradeCancelBatchOrdersV5RespDataInner

`func NewCreateTradeCancelBatchOrdersV5RespDataInner() *CreateTradeCancelBatchOrdersV5RespDataInner`

NewCreateTradeCancelBatchOrdersV5RespDataInner instantiates a new CreateTradeCancelBatchOrdersV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeCancelBatchOrdersV5RespDataInnerWithDefaults

`func NewCreateTradeCancelBatchOrdersV5RespDataInnerWithDefaults() *CreateTradeCancelBatchOrdersV5RespDataInner`

NewCreateTradeCancelBatchOrdersV5RespDataInnerWithDefaults instantiates a new CreateTradeCancelBatchOrdersV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetData() []CreateTradeCancelBatchOrdersV5RespDataInnerDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetDataOk() (*[]CreateTradeCancelBatchOrdersV5RespDataInnerDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) SetData(v []CreateTradeCancelBatchOrdersV5RespDataInnerDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetInTime

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetInTime() string`

GetInTime returns the InTime field if non-nil, zero value otherwise.

### GetInTimeOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetInTimeOk() (*string, bool)`

GetInTimeOk returns a tuple with the InTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTime

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) SetInTime(v string)`

SetInTime sets InTime field to given value.

### HasInTime

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) HasInTime() bool`

HasInTime returns a boolean if a field has been set.

### GetMsg

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetOutTime

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetOutTime() string`

GetOutTime returns the OutTime field if non-nil, zero value otherwise.

### GetOutTimeOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetOutTimeOk() (*string, bool)`

GetOutTimeOk returns a tuple with the OutTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutTime

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) SetOutTime(v string)`

SetOutTime sets OutTime field to given value.

### HasOutTime

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) HasOutTime() bool`

HasOutTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


