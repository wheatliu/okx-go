# CreateTradingBotSignalCancelSubOrderV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SCode** | Pointer to **string** | The code of the event execution result, &#x60;0&#x60; means success. | [optional] [default to ""]
**SMsg** | Pointer to **string** | Rejection or success message of event execution. | [optional] [default to ""]
**SignalOrdId** | Pointer to **string** | Order ID | [optional] [default to ""]

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

### GetSCode

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) GetSCode() string`

GetSCode returns the SCode field if non-nil, zero value otherwise.

### GetSCodeOk

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) GetSCodeOk() (*string, bool)`

GetSCodeOk returns a tuple with the SCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCode

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) SetSCode(v string)`

SetSCode sets SCode field to given value.

### HasSCode

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) HasSCode() bool`

HasSCode returns a boolean if a field has been set.

### GetSMsg

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) GetSMsg() string`

GetSMsg returns the SMsg field if non-nil, zero value otherwise.

### GetSMsgOk

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) GetSMsgOk() (*string, bool)`

GetSMsgOk returns a tuple with the SMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMsg

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) SetSMsg(v string)`

SetSMsg sets SMsg field to given value.

### HasSMsg

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) HasSMsg() bool`

HasSMsg returns a boolean if a field has been set.

### GetSignalOrdId

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) GetSignalOrdId() string`

GetSignalOrdId returns the SignalOrdId field if non-nil, zero value otherwise.

### GetSignalOrdIdOk

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) GetSignalOrdIdOk() (*string, bool)`

GetSignalOrdIdOk returns a tuple with the SignalOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalOrdId

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) SetSignalOrdId(v string)`

SetSignalOrdId sets SignalOrdId field to given value.

### HasSignalOrdId

`func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInner) HasSignalOrdId() bool`

HasSignalOrdId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


