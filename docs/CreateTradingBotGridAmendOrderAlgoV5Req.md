# CreateTradingBotGridAmendOrderAlgoV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | **string** | Algo ID | [default to ""]
**InstId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [default to ""]
**SlRatio** | Pointer to **string** | Stop loss ratio, 0.1 represents 10%, only applicable to contract grid&#x60;  if it is set \&quot;\&quot; means stop-loss ratio is canceled. | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | New stop-loss trigger price  if slTriggerPx is set \&quot;\&quot; means stop-loss trigger price is canceled.  Either &#x60;slTriggerPx&#x60; or &#x60;tpTriggerPx&#x60; is required. | [optional] [default to ""]
**TpRatio** | Pointer to **string** | Take profit ratio, 0.1 represents 10%, only applicable to contract grid  if it is set \&quot;\&quot; means take-profit ratio is canceled. | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | New take-profit trigger price  if tpTriggerPx is set \&quot;\&quot; means take-profit trigger price is canceled. | [optional] [default to ""]
**TriggerParams** | Pointer to [**[]CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner**](CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner.md) | Trigger Parameters | [optional] 

## Methods

### NewCreateTradingBotGridAmendOrderAlgoV5Req

`func NewCreateTradingBotGridAmendOrderAlgoV5Req(algoId string, instId string, ) *CreateTradingBotGridAmendOrderAlgoV5Req`

NewCreateTradingBotGridAmendOrderAlgoV5Req instantiates a new CreateTradingBotGridAmendOrderAlgoV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotGridAmendOrderAlgoV5ReqWithDefaults

`func NewCreateTradingBotGridAmendOrderAlgoV5ReqWithDefaults() *CreateTradingBotGridAmendOrderAlgoV5Req`

NewCreateTradingBotGridAmendOrderAlgoV5ReqWithDefaults instantiates a new CreateTradingBotGridAmendOrderAlgoV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.


### GetInstId

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetSlRatio

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetSlRatio() string`

GetSlRatio returns the SlRatio field if non-nil, zero value otherwise.

### GetSlRatioOk

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetSlRatioOk() (*string, bool)`

GetSlRatioOk returns a tuple with the SlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlRatio

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) SetSlRatio(v string)`

SetSlRatio sets SlRatio field to given value.

### HasSlRatio

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) HasSlRatio() bool`

HasSlRatio returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetTpRatio

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetTpRatio() string`

GetTpRatio returns the TpRatio field if non-nil, zero value otherwise.

### GetTpRatioOk

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetTpRatioOk() (*string, bool)`

GetTpRatioOk returns a tuple with the TpRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpRatio

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) SetTpRatio(v string)`

SetTpRatio sets TpRatio field to given value.

### HasTpRatio

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) HasTpRatio() bool`

HasTpRatio returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTriggerParams

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetTriggerParams() []CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner`

GetTriggerParams returns the TriggerParams field if non-nil, zero value otherwise.

### GetTriggerParamsOk

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) GetTriggerParamsOk() (*[]CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner, bool)`

GetTriggerParamsOk returns a tuple with the TriggerParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerParams

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) SetTriggerParams(v []CreateTradingBotGridAmendOrderAlgoV5ReqTriggerParamsInner)`

SetTriggerParams sets TriggerParams field to given value.

### HasTriggerParams

`func (o *CreateTradingBotGridAmendOrderAlgoV5Req) HasTriggerParams() bool`

HasTriggerParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


