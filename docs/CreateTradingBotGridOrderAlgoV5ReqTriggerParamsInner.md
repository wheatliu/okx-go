# CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner

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
**TriggerStrategy** | Pointer to **string** | Trigger strategy  &#x60;instant&#x60;  &#x60;price&#x60;  &#x60;rsi&#x60;  Default is &#x60;instant&#x60; | [optional] [default to ""]

## Methods

### NewCreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner

`func NewCreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner() *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner`

NewCreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner instantiates a new CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotGridOrderAlgoV5ReqTriggerParamsInnerWithDefaults

`func NewCreateTradingBotGridOrderAlgoV5ReqTriggerParamsInnerWithDefaults() *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner`

NewCreateTradingBotGridOrderAlgoV5ReqTriggerParamsInnerWithDefaults instantiates a new CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelaySeconds

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetDelaySeconds() string`

GetDelaySeconds returns the DelaySeconds field if non-nil, zero value otherwise.

### GetDelaySecondsOk

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetDelaySecondsOk() (*string, bool)`

GetDelaySecondsOk returns a tuple with the DelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelaySeconds

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) SetDelaySeconds(v string)`

SetDelaySeconds sets DelaySeconds field to given value.

### HasDelaySeconds

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) HasDelaySeconds() bool`

HasDelaySeconds returns a boolean if a field has been set.

### GetStopType

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetStopType() string`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetStopTypeOk() (*string, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) SetStopType(v string)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### GetThold

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetThold() string`

GetThold returns the Thold field if non-nil, zero value otherwise.

### GetTholdOk

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetTholdOk() (*string, bool)`

GetTholdOk returns a tuple with the Thold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThold

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) SetThold(v string)`

SetThold sets Thold field to given value.

### HasThold

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) HasThold() bool`

HasThold returns a boolean if a field has been set.

### GetTimePeriod

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetTimePeriod() string`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetTimePeriodOk() (*string, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) SetTimePeriod(v string)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetTimeframe

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetTimeframe() string`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetTimeframeOk() (*string, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) SetTimeframe(v string)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetTriggerAction

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetTriggerAction() string`

GetTriggerAction returns the TriggerAction field if non-nil, zero value otherwise.

### GetTriggerActionOk

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetTriggerActionOk() (*string, bool)`

GetTriggerActionOk returns a tuple with the TriggerAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAction

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) SetTriggerAction(v string)`

SetTriggerAction sets TriggerAction field to given value.

### HasTriggerAction

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) HasTriggerAction() bool`

HasTriggerAction returns a boolean if a field has been set.

### GetTriggerCond

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetTriggerCond() string`

GetTriggerCond returns the TriggerCond field if non-nil, zero value otherwise.

### GetTriggerCondOk

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetTriggerCondOk() (*string, bool)`

GetTriggerCondOk returns a tuple with the TriggerCond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerCond

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) SetTriggerCond(v string)`

SetTriggerCond sets TriggerCond field to given value.

### HasTriggerCond

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) HasTriggerCond() bool`

HasTriggerCond returns a boolean if a field has been set.

### GetTriggerPx

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetTriggerPx() string`

GetTriggerPx returns the TriggerPx field if non-nil, zero value otherwise.

### GetTriggerPxOk

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetTriggerPxOk() (*string, bool)`

GetTriggerPxOk returns a tuple with the TriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPx

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) SetTriggerPx(v string)`

SetTriggerPx sets TriggerPx field to given value.

### HasTriggerPx

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) HasTriggerPx() bool`

HasTriggerPx returns a boolean if a field has been set.

### GetTriggerStrategy

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetTriggerStrategy() string`

GetTriggerStrategy returns the TriggerStrategy field if non-nil, zero value otherwise.

### GetTriggerStrategyOk

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) GetTriggerStrategyOk() (*string, bool)`

GetTriggerStrategyOk returns a tuple with the TriggerStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerStrategy

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) SetTriggerStrategy(v string)`

SetTriggerStrategy sets TriggerStrategy field to given value.

### HasTriggerStrategy

`func (o *CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) HasTriggerStrategy() bool`

HasTriggerStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


