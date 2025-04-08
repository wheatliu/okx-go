# CreateTradingBotGridOrderAlgoV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [optional] [default to ""]
**AlgoOrdType** | **string** | Algo order type  &#x60;grid&#x60;: Spot grid  &#x60;contract_grid&#x60;: Contract grid | [default to ""]
**GridNum** | **string** | Grid quantity | [default to ""]
**InstId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [default to ""]
**MaxPx** | **string** | Upper price of price range | [default to ""]
**MinPx** | **string** | Lower price of price range | [default to ""]
**ProfitSharingRatio** | Pointer to **string** | Profit sharing ratio, it only supports these values  &#x60;0&#x60;,&#x60;0.1&#x60;,&#x60;0.2&#x60;,&#x60;0.3&#x60;   0.1 represents 10% | [optional] [default to ""]
**RunType** | Pointer to **string** | Grid type  &#x60;1&#x60;: Arithmetic, &#x60;2&#x60;: Geometric  Default is Arithmetic | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | SL tigger price  Applicable to &#x60;Spot grid&#x60;/&#x60;Contract grid&#x60; | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | TP tigger price  Applicable to &#x60;Spot grid&#x60;/&#x60;Contract grid&#x60; | [optional] [default to ""]
**TriggerParams** | Pointer to [**[]CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner**](CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner.md) | Trigger Parameters  Applicable to &#x60;Spot grid&#x60;/&#x60;Contract grid&#x60; | [optional] 

## Methods

### NewCreateTradingBotGridOrderAlgoV5Req

`func NewCreateTradingBotGridOrderAlgoV5Req(algoOrdType string, gridNum string, instId string, maxPx string, minPx string, ) *CreateTradingBotGridOrderAlgoV5Req`

NewCreateTradingBotGridOrderAlgoV5Req instantiates a new CreateTradingBotGridOrderAlgoV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotGridOrderAlgoV5ReqWithDefaults

`func NewCreateTradingBotGridOrderAlgoV5ReqWithDefaults() *CreateTradingBotGridOrderAlgoV5Req`

NewCreateTradingBotGridOrderAlgoV5ReqWithDefaults instantiates a new CreateTradingBotGridOrderAlgoV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoClOrdId

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *CreateTradingBotGridOrderAlgoV5Req) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *CreateTradingBotGridOrderAlgoV5Req) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoOrdType

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *CreateTradingBotGridOrderAlgoV5Req) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.


### GetGridNum

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetGridNum() string`

GetGridNum returns the GridNum field if non-nil, zero value otherwise.

### GetGridNumOk

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetGridNumOk() (*string, bool)`

GetGridNumOk returns a tuple with the GridNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridNum

`func (o *CreateTradingBotGridOrderAlgoV5Req) SetGridNum(v string)`

SetGridNum sets GridNum field to given value.


### GetInstId

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradingBotGridOrderAlgoV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetMaxPx

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetMaxPx() string`

GetMaxPx returns the MaxPx field if non-nil, zero value otherwise.

### GetMaxPxOk

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetMaxPxOk() (*string, bool)`

GetMaxPxOk returns a tuple with the MaxPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPx

`func (o *CreateTradingBotGridOrderAlgoV5Req) SetMaxPx(v string)`

SetMaxPx sets MaxPx field to given value.


### GetMinPx

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetMinPx() string`

GetMinPx returns the MinPx field if non-nil, zero value otherwise.

### GetMinPxOk

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetMinPxOk() (*string, bool)`

GetMinPxOk returns a tuple with the MinPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPx

`func (o *CreateTradingBotGridOrderAlgoV5Req) SetMinPx(v string)`

SetMinPx sets MinPx field to given value.


### GetProfitSharingRatio

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetProfitSharingRatio() string`

GetProfitSharingRatio returns the ProfitSharingRatio field if non-nil, zero value otherwise.

### GetProfitSharingRatioOk

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetProfitSharingRatioOk() (*string, bool)`

GetProfitSharingRatioOk returns a tuple with the ProfitSharingRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitSharingRatio

`func (o *CreateTradingBotGridOrderAlgoV5Req) SetProfitSharingRatio(v string)`

SetProfitSharingRatio sets ProfitSharingRatio field to given value.

### HasProfitSharingRatio

`func (o *CreateTradingBotGridOrderAlgoV5Req) HasProfitSharingRatio() bool`

HasProfitSharingRatio returns a boolean if a field has been set.

### GetRunType

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *CreateTradingBotGridOrderAlgoV5Req) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *CreateTradingBotGridOrderAlgoV5Req) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *CreateTradingBotGridOrderAlgoV5Req) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *CreateTradingBotGridOrderAlgoV5Req) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetTag

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateTradingBotGridOrderAlgoV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateTradingBotGridOrderAlgoV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *CreateTradingBotGridOrderAlgoV5Req) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *CreateTradingBotGridOrderAlgoV5Req) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetTriggerParams

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetTriggerParams() []CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner`

GetTriggerParams returns the TriggerParams field if non-nil, zero value otherwise.

### GetTriggerParamsOk

`func (o *CreateTradingBotGridOrderAlgoV5Req) GetTriggerParamsOk() (*[]CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner, bool)`

GetTriggerParamsOk returns a tuple with the TriggerParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerParams

`func (o *CreateTradingBotGridOrderAlgoV5Req) SetTriggerParams(v []CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner)`

SetTriggerParams sets TriggerParams field to given value.

### HasTriggerParams

`func (o *CreateTradingBotGridOrderAlgoV5Req) HasTriggerParams() bool`

HasTriggerParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


