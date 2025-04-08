# GetTradingBotSignalEventHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertMsg** | Pointer to **string** | Alert message | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID | [optional] [default to ""]
**EventCtime** | Pointer to **string** | Event timestamp of creation. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**EventProcessMsg** | Pointer to **string** | Event process message | [optional] [default to ""]
**EventStatus** | Pointer to **string** | Event status  &#x60;success&#x60;  &#x60;failure&#x60; | [optional] [default to ""]
**EventType** | Pointer to **string** | Event type  &#x60;system_action&#x60;  &#x60;user_action&#x60;  &#x60;signal_processing&#x60; | [optional] [default to ""]
**EventUtime** | Pointer to **string** | Event timestamp of update. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**TriggeredOrdData** | Pointer to [**[]GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner**](GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner.md) | Triggered sub order data | [optional] 

## Methods

### NewGetTradingBotSignalEventHistoryV5RespDataInner

`func NewGetTradingBotSignalEventHistoryV5RespDataInner() *GetTradingBotSignalEventHistoryV5RespDataInner`

NewGetTradingBotSignalEventHistoryV5RespDataInner instantiates a new GetTradingBotSignalEventHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotSignalEventHistoryV5RespDataInnerWithDefaults

`func NewGetTradingBotSignalEventHistoryV5RespDataInnerWithDefaults() *GetTradingBotSignalEventHistoryV5RespDataInner`

NewGetTradingBotSignalEventHistoryV5RespDataInnerWithDefaults instantiates a new GetTradingBotSignalEventHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertMsg

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetAlertMsg() string`

GetAlertMsg returns the AlertMsg field if non-nil, zero value otherwise.

### GetAlertMsgOk

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetAlertMsgOk() (*string, bool)`

GetAlertMsgOk returns a tuple with the AlertMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertMsg

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetAlertMsg(v string)`

SetAlertMsg sets AlertMsg field to given value.

### HasAlertMsg

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasAlertMsg() bool`

HasAlertMsg returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetEventCtime

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventCtime() string`

GetEventCtime returns the EventCtime field if non-nil, zero value otherwise.

### GetEventCtimeOk

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventCtimeOk() (*string, bool)`

GetEventCtimeOk returns a tuple with the EventCtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCtime

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetEventCtime(v string)`

SetEventCtime sets EventCtime field to given value.

### HasEventCtime

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasEventCtime() bool`

HasEventCtime returns a boolean if a field has been set.

### GetEventProcessMsg

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventProcessMsg() string`

GetEventProcessMsg returns the EventProcessMsg field if non-nil, zero value otherwise.

### GetEventProcessMsgOk

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventProcessMsgOk() (*string, bool)`

GetEventProcessMsgOk returns a tuple with the EventProcessMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventProcessMsg

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetEventProcessMsg(v string)`

SetEventProcessMsg sets EventProcessMsg field to given value.

### HasEventProcessMsg

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasEventProcessMsg() bool`

HasEventProcessMsg returns a boolean if a field has been set.

### GetEventStatus

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventStatus() string`

GetEventStatus returns the EventStatus field if non-nil, zero value otherwise.

### GetEventStatusOk

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventStatusOk() (*string, bool)`

GetEventStatusOk returns a tuple with the EventStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStatus

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetEventStatus(v string)`

SetEventStatus sets EventStatus field to given value.

### HasEventStatus

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasEventStatus() bool`

HasEventStatus returns a boolean if a field has been set.

### GetEventType

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetEventUtime

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventUtime() string`

GetEventUtime returns the EventUtime field if non-nil, zero value otherwise.

### GetEventUtimeOk

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetEventUtimeOk() (*string, bool)`

GetEventUtimeOk returns a tuple with the EventUtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUtime

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetEventUtime(v string)`

SetEventUtime sets EventUtime field to given value.

### HasEventUtime

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasEventUtime() bool`

HasEventUtime returns a boolean if a field has been set.

### GetTriggeredOrdData

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetTriggeredOrdData() []GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner`

GetTriggeredOrdData returns the TriggeredOrdData field if non-nil, zero value otherwise.

### GetTriggeredOrdDataOk

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) GetTriggeredOrdDataOk() (*[]GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner, bool)`

GetTriggeredOrdDataOk returns a tuple with the TriggeredOrdData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredOrdData

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) SetTriggeredOrdData(v []GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner)`

SetTriggeredOrdData sets TriggeredOrdData field to given value.

### HasTriggeredOrdData

`func (o *GetTradingBotSignalEventHistoryV5RespDataInner) HasTriggeredOrdData() bool`

HasTriggeredOrdData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


