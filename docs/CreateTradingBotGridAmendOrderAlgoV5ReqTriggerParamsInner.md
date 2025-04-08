# CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StopType** | Pointer to **string** | Stop type  Spot grid &#x60;1&#x60;: Sell base currency &#x60;2&#x60;: Keep base currency  Contract grid &#x60;1&#x60;: Market Close All positions &#x60;2&#x60;: Keep positions  This field is only valid when &#x60;triggerAction&#x60; is &#x60;stop&#x60; | [optional] [default to ""]
**TriggerAction** | Pointer to **string** | Trigger action  &#x60;start&#x60;  &#x60;stop&#x60; | [optional] [default to ""]
**TriggerPx** | Pointer to **string** | Trigger Price  This field is only valid when &#x60;triggerStrategy&#x60; is &#x60;price&#x60; | [optional] [default to ""]
**TriggerStrategy** | Pointer to **string** | Trigger strategy  &#x60;instant&#x60;  &#x60;price&#x60;  &#x60;rsi&#x60; | [optional] [default to ""]

## Methods

### NewCreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner

`func NewCreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner() *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner`

NewCreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner instantiates a new CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInnerWithDefaults

`func NewCreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInnerWithDefaults() *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner`

NewCreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInnerWithDefaults instantiates a new CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStopType

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) GetStopType() string`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) GetStopTypeOk() (*string, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) SetStopType(v string)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### GetTriggerAction

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) GetTriggerAction() string`

GetTriggerAction returns the TriggerAction field if non-nil, zero value otherwise.

### GetTriggerActionOk

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) GetTriggerActionOk() (*string, bool)`

GetTriggerActionOk returns a tuple with the TriggerAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAction

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) SetTriggerAction(v string)`

SetTriggerAction sets TriggerAction field to given value.

### HasTriggerAction

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) HasTriggerAction() bool`

HasTriggerAction returns a boolean if a field has been set.

### GetTriggerPx

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) GetTriggerPx() string`

GetTriggerPx returns the TriggerPx field if non-nil, zero value otherwise.

### GetTriggerPxOk

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) GetTriggerPxOk() (*string, bool)`

GetTriggerPxOk returns a tuple with the TriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPx

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) SetTriggerPx(v string)`

SetTriggerPx sets TriggerPx field to given value.

### HasTriggerPx

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) HasTriggerPx() bool`

HasTriggerPx returns a boolean if a field has been set.

### GetTriggerStrategy

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) GetTriggerStrategy() string`

GetTriggerStrategy returns the TriggerStrategy field if non-nil, zero value otherwise.

### GetTriggerStrategyOk

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) GetTriggerStrategyOk() (*string, bool)`

GetTriggerStrategyOk returns a tuple with the TriggerStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerStrategy

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) SetTriggerStrategy(v string)`

SetTriggerStrategy sets TriggerStrategy field to given value.

### HasTriggerStrategy

`func (o *CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner) HasTriggerStrategy() bool`

HasTriggerStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


