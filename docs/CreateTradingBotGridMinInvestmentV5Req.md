# CreateTradingBotGridMinInvestmentV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoOrdType** | **string** | Algo order type  &#x60;grid&#x60;: Spot grid  &#x60;contract_grid&#x60;: Contract grid | [default to ""]
**BasePos** | Pointer to **bool** | Whether or not open a position when the strategy activates  Default is &#x60;false&#x60;  Neutral contract grid should omit the parameter  Only applicable to &#x60;contract grid&#x60; | [optional] 
**Direction** | Pointer to **string** | Contract grid type  &#x60;long&#x60;,&#x60;short&#x60;,&#x60;neutral&#x60;  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**GridNum** | **string** | Grid quantity | [default to ""]
**InstId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [default to ""]
**InvestmentData** | Pointer to [**[]CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner**](CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner.md) | Invest Data | [optional] 
**InvestmentType** | Pointer to **string** | Investment type, only applicable to &#x60;grid&#x60;   &#x60;quote&#x60;  &#x60;base&#x60;  &#x60;dual&#x60; | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage  Only applicable to &#x60;contract grid&#x60; | [optional] [default to ""]
**MaxPx** | **string** | Upper price of price range | [default to ""]
**MinPx** | **string** | Lower price of price range | [default to ""]
**RunType** | **string** | Grid type  &#x60;1&#x60;: Arithmetic, &#x60;2&#x60;: Geometric | [default to ""]
**TriggerStrategy** | Pointer to **string** | Trigger stragety,    &#x60;instant&#x60;  &#x60;price&#x60;  &#x60;rsi&#x60; | [optional] [default to ""]

## Methods

### NewCreateTradingBotGridMinInvestmentV5Req

`func NewCreateTradingBotGridMinInvestmentV5Req(algoOrdType string, gridNum string, instId string, maxPx string, minPx string, runType string, ) *CreateTradingBotGridMinInvestmentV5Req`

NewCreateTradingBotGridMinInvestmentV5Req instantiates a new CreateTradingBotGridMinInvestmentV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotGridMinInvestmentV5ReqWithDefaults

`func NewCreateTradingBotGridMinInvestmentV5ReqWithDefaults() *CreateTradingBotGridMinInvestmentV5Req`

NewCreateTradingBotGridMinInvestmentV5ReqWithDefaults instantiates a new CreateTradingBotGridMinInvestmentV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoOrdType

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *CreateTradingBotGridMinInvestmentV5Req) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.


### GetBasePos

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetBasePos() bool`

GetBasePos returns the BasePos field if non-nil, zero value otherwise.

### GetBasePosOk

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetBasePosOk() (*bool, bool)`

GetBasePosOk returns a tuple with the BasePos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePos

`func (o *CreateTradingBotGridMinInvestmentV5Req) SetBasePos(v bool)`

SetBasePos sets BasePos field to given value.

### HasBasePos

`func (o *CreateTradingBotGridMinInvestmentV5Req) HasBasePos() bool`

HasBasePos returns a boolean if a field has been set.

### GetDirection

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CreateTradingBotGridMinInvestmentV5Req) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CreateTradingBotGridMinInvestmentV5Req) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetGridNum

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetGridNum() string`

GetGridNum returns the GridNum field if non-nil, zero value otherwise.

### GetGridNumOk

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetGridNumOk() (*string, bool)`

GetGridNumOk returns a tuple with the GridNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridNum

`func (o *CreateTradingBotGridMinInvestmentV5Req) SetGridNum(v string)`

SetGridNum sets GridNum field to given value.


### GetInstId

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradingBotGridMinInvestmentV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetInvestmentData

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetInvestmentData() []CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner`

GetInvestmentData returns the InvestmentData field if non-nil, zero value otherwise.

### GetInvestmentDataOk

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetInvestmentDataOk() (*[]CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner, bool)`

GetInvestmentDataOk returns a tuple with the InvestmentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentData

`func (o *CreateTradingBotGridMinInvestmentV5Req) SetInvestmentData(v []CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner)`

SetInvestmentData sets InvestmentData field to given value.

### HasInvestmentData

`func (o *CreateTradingBotGridMinInvestmentV5Req) HasInvestmentData() bool`

HasInvestmentData returns a boolean if a field has been set.

### GetInvestmentType

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetInvestmentType() string`

GetInvestmentType returns the InvestmentType field if non-nil, zero value otherwise.

### GetInvestmentTypeOk

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetInvestmentTypeOk() (*string, bool)`

GetInvestmentTypeOk returns a tuple with the InvestmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentType

`func (o *CreateTradingBotGridMinInvestmentV5Req) SetInvestmentType(v string)`

SetInvestmentType sets InvestmentType field to given value.

### HasInvestmentType

`func (o *CreateTradingBotGridMinInvestmentV5Req) HasInvestmentType() bool`

HasInvestmentType returns a boolean if a field has been set.

### GetLever

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *CreateTradingBotGridMinInvestmentV5Req) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *CreateTradingBotGridMinInvestmentV5Req) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetMaxPx

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetMaxPx() string`

GetMaxPx returns the MaxPx field if non-nil, zero value otherwise.

### GetMaxPxOk

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetMaxPxOk() (*string, bool)`

GetMaxPxOk returns a tuple with the MaxPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPx

`func (o *CreateTradingBotGridMinInvestmentV5Req) SetMaxPx(v string)`

SetMaxPx sets MaxPx field to given value.


### GetMinPx

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetMinPx() string`

GetMinPx returns the MinPx field if non-nil, zero value otherwise.

### GetMinPxOk

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetMinPxOk() (*string, bool)`

GetMinPxOk returns a tuple with the MinPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPx

`func (o *CreateTradingBotGridMinInvestmentV5Req) SetMinPx(v string)`

SetMinPx sets MinPx field to given value.


### GetRunType

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *CreateTradingBotGridMinInvestmentV5Req) SetRunType(v string)`

SetRunType sets RunType field to given value.


### GetTriggerStrategy

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetTriggerStrategy() string`

GetTriggerStrategy returns the TriggerStrategy field if non-nil, zero value otherwise.

### GetTriggerStrategyOk

`func (o *CreateTradingBotGridMinInvestmentV5Req) GetTriggerStrategyOk() (*string, bool)`

GetTriggerStrategyOk returns a tuple with the TriggerStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerStrategy

`func (o *CreateTradingBotGridMinInvestmentV5Req) SetTriggerStrategy(v string)`

SetTriggerStrategy sets TriggerStrategy field to given value.

### HasTriggerStrategy

`func (o *CreateTradingBotGridMinInvestmentV5Req) HasTriggerStrategy() bool`

HasTriggerStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


