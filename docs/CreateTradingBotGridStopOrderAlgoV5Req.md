# CreateTradingBotGridStopOrderAlgoV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | **string** | Algo ID | [default to ""]
**AlgoOrdType** | **string** | Algo order type  &#x60;grid&#x60;: Spot grid  &#x60;contract_grid&#x60;: Contract grid | [default to ""]
**InstId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to ""]
**StopType** | **string** | Stop type  Spot grid &#x60;1&#x60;: Sell base currency &#x60;2&#x60;: Keep base currency  Contract grid &#x60;1&#x60;: Market Close All positions &#x60;2&#x60;: Keep positions | [default to ""]

## Methods

### NewCreateTradingBotGridStopOrderAlgoV5Req

`func NewCreateTradingBotGridStopOrderAlgoV5Req(algoId string, algoOrdType string, instId string, stopType string, ) *CreateTradingBotGridStopOrderAlgoV5Req`

NewCreateTradingBotGridStopOrderAlgoV5Req instantiates a new CreateTradingBotGridStopOrderAlgoV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotGridStopOrderAlgoV5ReqWithDefaults

`func NewCreateTradingBotGridStopOrderAlgoV5ReqWithDefaults() *CreateTradingBotGridStopOrderAlgoV5Req`

NewCreateTradingBotGridStopOrderAlgoV5ReqWithDefaults instantiates a new CreateTradingBotGridStopOrderAlgoV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *CreateTradingBotGridStopOrderAlgoV5Req) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradingBotGridStopOrderAlgoV5Req) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradingBotGridStopOrderAlgoV5Req) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.


### GetAlgoOrdType

`func (o *CreateTradingBotGridStopOrderAlgoV5Req) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *CreateTradingBotGridStopOrderAlgoV5Req) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *CreateTradingBotGridStopOrderAlgoV5Req) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.


### GetInstId

`func (o *CreateTradingBotGridStopOrderAlgoV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradingBotGridStopOrderAlgoV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradingBotGridStopOrderAlgoV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetStopType

`func (o *CreateTradingBotGridStopOrderAlgoV5Req) GetStopType() string`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *CreateTradingBotGridStopOrderAlgoV5Req) GetStopTypeOk() (*string, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *CreateTradingBotGridStopOrderAlgoV5Req) SetStopType(v string)`

SetStopType sets StopType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


