# GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelaySeconds** | Pointer to **string** | Delay seconds after action triggered | [optional] [default to ""]
**StopType** | Pointer to **string** | Stop type  Spot grid &#x60;1&#x60;: Sell base currency &#x60;2&#x60;: Keep base currency  Contract grid &#x60;1&#x60;: Market Close All positions &#x60;2&#x60;: Keep positions  This field is only valid when &#x60;triggerAction&#x60; is &#x60;stop&#x60; | [optional] [default to ""]
**Thold** | Pointer to **string** | Threshold  The value should be an integer between 1 to 100  This field is only valid when &#x60;triggerStrategy&#x60; is &#x60;rsi&#x60; | [optional] [default to ""]
**TimePeriod** | Pointer to **string** | Time Period  &#x60;14&#x60;  This field is only valid when &#x60;triggerStrategy&#x60; is &#x60;rsi&#x60; | [optional] [default to ""]
**Timeframe** | Pointer to **string** | K-line type  &#x60;3m&#x60;, &#x60;5m&#x60;, &#x60;15m&#x60;, &#x60;30m&#x60; (&#x60;m&#x60;: minute)  &#x60;1H&#x60;, &#x60;4H&#x60; (&#x60;H&#x60;: hour)  &#x60;1D&#x60; (&#x60;D&#x60;: day)  This field is only valid when &#x60;triggerStrategy&#x60; is &#x60;rsi&#x60; | [optional] [default to ""]
**TriggerAction** | Pointer to **string** | Trigger action  &#x60;start&#x60;  &#x60;stop&#x60; | [optional] [default to ""]
**TriggerCond** | Pointer to **string** | Trigger condition  &#x60;cross_up&#x60;  &#x60;cross_down&#x60;  &#x60;above&#x60;  &#x60;below&#x60;  &#x60;cross&#x60;  This field is only valid when &#x60;triggerStrategy&#x60; is &#x60;rsi&#x60; | [optional] [default to ""]
**TriggerPx** | Pointer to **string** | Trigger Price  This field is only valid when &#x60;triggerStrategy&#x60; is &#x60;price&#x60; | [optional] [default to ""]
**TriggerStrategy** | Pointer to **string** | Trigger strategy  &#x60;instant&#x60;  &#x60;price&#x60;  &#x60;rsi&#x60; | [optional] [default to ""]
**TriggerTime** | Pointer to **string** | Actual action triggered time, unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**TriggerType** | Pointer to **string** | Actual action triggered type  &#x60;manual&#x60;  &#x60;auto&#x60; | [optional] [default to ""]

## Methods

### NewGetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner

`func NewGetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner() *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner`

NewGetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner instantiates a new GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInnerWithDefaults

`func NewGetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInnerWithDefaults() *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner`

NewGetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInnerWithDefaults instantiates a new GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelaySeconds

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetDelaySeconds() string`

GetDelaySeconds returns the DelaySeconds field if non-nil, zero value otherwise.

### GetDelaySecondsOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetDelaySecondsOk() (*string, bool)`

GetDelaySecondsOk returns a tuple with the DelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelaySeconds

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) SetDelaySeconds(v string)`

SetDelaySeconds sets DelaySeconds field to given value.

### HasDelaySeconds

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) HasDelaySeconds() bool`

HasDelaySeconds returns a boolean if a field has been set.

### GetStopType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetStopType() string`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetStopTypeOk() (*string, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) SetStopType(v string)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### GetThold

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetThold() string`

GetThold returns the Thold field if non-nil, zero value otherwise.

### GetTholdOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTholdOk() (*string, bool)`

GetTholdOk returns a tuple with the Thold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThold

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) SetThold(v string)`

SetThold sets Thold field to given value.

### HasThold

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) HasThold() bool`

HasThold returns a boolean if a field has been set.

### GetTimePeriod

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTimePeriod() string`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTimePeriodOk() (*string, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) SetTimePeriod(v string)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetTimeframe

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTimeframe() string`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTimeframeOk() (*string, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) SetTimeframe(v string)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetTriggerAction

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTriggerAction() string`

GetTriggerAction returns the TriggerAction field if non-nil, zero value otherwise.

### GetTriggerActionOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTriggerActionOk() (*string, bool)`

GetTriggerActionOk returns a tuple with the TriggerAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAction

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) SetTriggerAction(v string)`

SetTriggerAction sets TriggerAction field to given value.

### HasTriggerAction

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) HasTriggerAction() bool`

HasTriggerAction returns a boolean if a field has been set.

### GetTriggerCond

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTriggerCond() string`

GetTriggerCond returns the TriggerCond field if non-nil, zero value otherwise.

### GetTriggerCondOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTriggerCondOk() (*string, bool)`

GetTriggerCondOk returns a tuple with the TriggerCond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCond

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) SetTriggerCond(v string)`

SetTriggerCond sets TriggerCond field to given value.

### HasTriggerCond

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) HasTriggerCond() bool`

HasTriggerCond returns a boolean if a field has been set.

### GetTriggerPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTriggerPx() string`

GetTriggerPx returns the TriggerPx field if non-nil, zero value otherwise.

### GetTriggerPxOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTriggerPxOk() (*string, bool)`

GetTriggerPxOk returns a tuple with the TriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) SetTriggerPx(v string)`

SetTriggerPx sets TriggerPx field to given value.

### HasTriggerPx

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) HasTriggerPx() bool`

HasTriggerPx returns a boolean if a field has been set.

### GetTriggerStrategy

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTriggerStrategy() string`

GetTriggerStrategy returns the TriggerStrategy field if non-nil, zero value otherwise.

### GetTriggerStrategyOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTriggerStrategyOk() (*string, bool)`

GetTriggerStrategyOk returns a tuple with the TriggerStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerStrategy

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) SetTriggerStrategy(v string)`

SetTriggerStrategy sets TriggerStrategy field to given value.

### HasTriggerStrategy

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) HasTriggerStrategy() bool`

HasTriggerStrategy returns a boolean if a field has been set.

### GetTriggerTime

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTriggerTime() string`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTriggerTimeOk() (*string, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) SetTriggerTime(v string)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetTriggerType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *GetTradingBotGridOrdersAlgoDetailsV5RespDataInnerTriggerParamsInner) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


