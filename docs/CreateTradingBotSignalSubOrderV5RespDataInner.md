# CreateTradingBotSignalSubOrderV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The result code, &#x60;0&#x60; means success | [optional] [default to ""]
**Data** | Pointer to **[]map[string]interface{}** | Array of objects contains the response results | [optional] 
**Msg** | Pointer to **string** | The error message, empty if the code is 0 | [optional] [default to ""]

## Methods

### NewCreateTradingBotSignalSubOrderV5RespDataInner

`func NewCreateTradingBotSignalSubOrderV5RespDataInner() *CreateTradingBotSignalSubOrderV5RespDataInner`

NewCreateTradingBotSignalSubOrderV5RespDataInner instantiates a new CreateTradingBotSignalSubOrderV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotSignalSubOrderV5RespDataInnerWithDefaults

`func NewCreateTradingBotSignalSubOrderV5RespDataInnerWithDefaults() *CreateTradingBotSignalSubOrderV5RespDataInner`

NewCreateTradingBotSignalSubOrderV5RespDataInnerWithDefaults instantiates a new CreateTradingBotSignalSubOrderV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateTradingBotSignalSubOrderV5RespDataInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateTradingBotSignalSubOrderV5RespDataInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateTradingBotSignalSubOrderV5RespDataInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateTradingBotSignalSubOrderV5RespDataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *CreateTradingBotSignalSubOrderV5RespDataInner) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTradingBotSignalSubOrderV5RespDataInner) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTradingBotSignalSubOrderV5RespDataInner) SetData(v []map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CreateTradingBotSignalSubOrderV5RespDataInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CreateTradingBotSignalSubOrderV5RespDataInner) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateTradingBotSignalSubOrderV5RespDataInner) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateTradingBotSignalSubOrderV5RespDataInner) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateTradingBotSignalSubOrderV5RespDataInner) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


