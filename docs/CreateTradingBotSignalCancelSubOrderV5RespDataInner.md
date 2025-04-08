# CreateTradingBotSignalCancelSubOrderV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The result code, &#x60;0&#x60; means success | [optional] [default to ""]
**Data** | Pointer to [**[]CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner**](CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner.md) | Array of objects contains the response results | [optional] 
**Msg** | Pointer to **string** | The error message, empty if the code is 0 | [optional] [default to ""]

## Methods

### NewCreateTradingBotSignalCancelSubOrderV5RespDataInner

`func NewCreateTradingBotSignalCancelSubOrderV5RespDataInner() *CreateTradingBotSignalCancelSubOrderV5RespDataInner`

NewCreateTradingBotSignalCancelSubOrderV5RespDataInner instantiates a new CreateTradingBotSignalCancelSubOrderV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotSignalCancelSubOrderV5RespDataInnerWithDefaults

`func NewCreateTradingBotSignalCancelSubOrderV5RespDataInnerWithDefaults() *CreateTradingBotSignalCancelSubOrderV5RespDataInner`

NewCreateTradingBotSignalCancelSubOrderV5RespDataInnerWithDefaults instantiates a new CreateTradingBotSignalCancelSubOrderV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) GetData() []CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) GetDataOk() (*[]CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) SetData(v []CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


