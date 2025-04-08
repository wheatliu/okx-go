# CreateTradeAmendBatchOrdersV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The result code, &#x60;0&#x60; means success | [optional] [default to ""]
**Data** | Pointer to [**[]CreateTradeAmendBatchOrdersV5RespDataDataInner**](CreateTradeAmendBatchOrdersV5RespDataDataInner.md) | Array of objects contains the response results | [optional] 
**InTime** | Pointer to **string** | Timestamp at REST gateway when the request is received, Unix timestamp format in microseconds, e.g. &#x60;1597026383085123&#x60;   The time is recorded after authentication. | [optional] [default to ""]
**Msg** | Pointer to **string** | The error message, empty if the code is 0 | [optional] [default to ""]
**OutTime** | Pointer to **string** | Timestamp at REST gateway when the response is sent, Unix timestamp format in microseconds, e.g. &#x60;1597026383085123&#x60; | [optional] [default to ""]

## Methods

### NewCreateTradeAmendBatchOrdersV5RespData

`func NewCreateTradeAmendBatchOrdersV5RespData() *CreateTradeAmendBatchOrdersV5RespData`

NewCreateTradeAmendBatchOrdersV5RespData instantiates a new CreateTradeAmendBatchOrdersV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeAmendBatchOrdersV5RespDataWithDefaults

`func NewCreateTradeAmendBatchOrdersV5RespDataWithDefaults() *CreateTradeAmendBatchOrdersV5RespData`

NewCreateTradeAmendBatchOrdersV5RespDataWithDefaults instantiates a new CreateTradeAmendBatchOrdersV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateTradeAmendBatchOrdersV5RespData) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateTradeAmendBatchOrdersV5RespData) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateTradeAmendBatchOrdersV5RespData) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateTradeAmendBatchOrdersV5RespData) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *CreateTradeAmendBatchOrdersV5RespData) GetData() []CreateTradeAmendBatchOrdersV5RespDataDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTradeAmendBatchOrdersV5RespData) GetDataOk() (*[]CreateTradeAmendBatchOrdersV5RespDataDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTradeAmendBatchOrdersV5RespData) SetData(v []CreateTradeAmendBatchOrdersV5RespDataDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *CreateTradeAmendBatchOrdersV5RespData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetInTime

`func (o *CreateTradeAmendBatchOrdersV5RespData) GetInTime() string`

GetInTime returns the InTime field if non-nil, zero value otherwise.

### GetInTimeOk

`func (o *CreateTradeAmendBatchOrdersV5RespData) GetInTimeOk() (*string, bool)`

GetInTimeOk returns a tuple with the InTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInTime

`func (o *CreateTradeAmendBatchOrdersV5RespData) SetInTime(v string)`

SetInTime sets InTime field to given value.

### HasInTime

`func (o *CreateTradeAmendBatchOrdersV5RespData) HasInTime() bool`

HasInTime returns a boolean if a field has been set.

### GetMsg

`func (o *CreateTradeAmendBatchOrdersV5RespData) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateTradeAmendBatchOrdersV5RespData) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateTradeAmendBatchOrdersV5RespData) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateTradeAmendBatchOrdersV5RespData) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetOutTime

`func (o *CreateTradeAmendBatchOrdersV5RespData) GetOutTime() string`

GetOutTime returns the OutTime field if non-nil, zero value otherwise.

### GetOutTimeOk

`func (o *CreateTradeAmendBatchOrdersV5RespData) GetOutTimeOk() (*string, bool)`

GetOutTimeOk returns a tuple with the OutTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutTime

`func (o *CreateTradeAmendBatchOrdersV5RespData) SetOutTime(v string)`

SetOutTime sets OutTime field to given value.

### HasOutTime

`func (o *CreateTradeAmendBatchOrdersV5RespData) HasOutTime() bool`

HasOutTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


