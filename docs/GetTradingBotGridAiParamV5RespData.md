# GetTradingBotGridAiParamV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoOrdType** | Pointer to **string** | Algo order type  &#x60;grid&#x60;: Spot grid  &#x60;contract_grid&#x60;: Contract grid | [optional] [default to ""]
**AnnualizedRate** | Pointer to **string** | Grid annualized rate | [optional] [default to ""]
**Ccy** | Pointer to **string** | The invest currency | [optional] [default to ""]
**Direction** | Pointer to **string** | Contract grid type  &#x60;long&#x60;,&#x60;short&#x60;,&#x60;neutral&#x60;  Only applicable to contract grid | [optional] [default to ""]
**Duration** | Pointer to **string** | Back testing duration  &#x60;7D&#x60;: 7 Days, &#x60;30D&#x60;: 30 Days, &#x60;180D&#x60;: 180 Days | [optional] [default to ""]
**GridNum** | Pointer to **string** | Grid quantity | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage  Only applicable to contract grid | [optional] [default to ""]
**MaxPx** | Pointer to **string** | Upper price of price range | [optional] [default to ""]
**MinInvestment** | Pointer to **string** | The minimum invest amount | [optional] [default to ""]
**MinPx** | Pointer to **string** | Lower price of price range | [optional] [default to ""]
**PerGridProfitRatio** | Pointer to **string** | Per grid profit ratio | [optional] [default to ""]
**PerMaxProfitRate** | Pointer to **string** | Estimated maximum Profit margin per grid | [optional] [default to ""]
**PerMinProfitRate** | Pointer to **string** | Estimated minimum Profit margin per grid | [optional] [default to ""]
**RunType** | Pointer to **string** | Grid type  &#x60;1&#x60;: Arithmetic, &#x60;2&#x60;: Geometric | [optional] [default to ""]
**SourceCcy** | Pointer to **string** | Source currency | [optional] [default to ""]

## Methods

### NewGetTradingBotGridAiParamV5RespData

`func NewGetTradingBotGridAiParamV5RespData() *GetTradingBotGridAiParamV5RespData`

NewGetTradingBotGridAiParamV5RespData instantiates a new GetTradingBotGridAiParamV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotGridAiParamV5RespDataWithDefaults

`func NewGetTradingBotGridAiParamV5RespDataWithDefaults() *GetTradingBotGridAiParamV5RespData`

NewGetTradingBotGridAiParamV5RespDataWithDefaults instantiates a new GetTradingBotGridAiParamV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoOrdType

`func (o *GetTradingBotGridAiParamV5RespData) GetAlgoOrdType() string`

GetAlgoOrdType returns the AlgoOrdType field if non-nil, zero value otherwise.

### GetAlgoOrdTypeOk

`func (o *GetTradingBotGridAiParamV5RespData) GetAlgoOrdTypeOk() (*string, bool)`

GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoOrdType

`func (o *GetTradingBotGridAiParamV5RespData) SetAlgoOrdType(v string)`

SetAlgoOrdType sets AlgoOrdType field to given value.

### HasAlgoOrdType

`func (o *GetTradingBotGridAiParamV5RespData) HasAlgoOrdType() bool`

HasAlgoOrdType returns a boolean if a field has been set.

### GetAnnualizedRate

`func (o *GetTradingBotGridAiParamV5RespData) GetAnnualizedRate() string`

GetAnnualizedRate returns the AnnualizedRate field if non-nil, zero value otherwise.

### GetAnnualizedRateOk

`func (o *GetTradingBotGridAiParamV5RespData) GetAnnualizedRateOk() (*string, bool)`

GetAnnualizedRateOk returns a tuple with the AnnualizedRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualizedRate

`func (o *GetTradingBotGridAiParamV5RespData) SetAnnualizedRate(v string)`

SetAnnualizedRate sets AnnualizedRate field to given value.

### HasAnnualizedRate

`func (o *GetTradingBotGridAiParamV5RespData) HasAnnualizedRate() bool`

HasAnnualizedRate returns a boolean if a field has been set.

### GetCcy

`func (o *GetTradingBotGridAiParamV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetTradingBotGridAiParamV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetTradingBotGridAiParamV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetTradingBotGridAiParamV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetDirection

`func (o *GetTradingBotGridAiParamV5RespData) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetTradingBotGridAiParamV5RespData) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetTradingBotGridAiParamV5RespData) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetTradingBotGridAiParamV5RespData) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDuration

`func (o *GetTradingBotGridAiParamV5RespData) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetTradingBotGridAiParamV5RespData) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetTradingBotGridAiParamV5RespData) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetTradingBotGridAiParamV5RespData) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetGridNum

`func (o *GetTradingBotGridAiParamV5RespData) GetGridNum() string`

GetGridNum returns the GridNum field if non-nil, zero value otherwise.

### GetGridNumOk

`func (o *GetTradingBotGridAiParamV5RespData) GetGridNumOk() (*string, bool)`

GetGridNumOk returns a tuple with the GridNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridNum

`func (o *GetTradingBotGridAiParamV5RespData) SetGridNum(v string)`

SetGridNum sets GridNum field to given value.

### HasGridNum

`func (o *GetTradingBotGridAiParamV5RespData) HasGridNum() bool`

HasGridNum returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradingBotGridAiParamV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradingBotGridAiParamV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradingBotGridAiParamV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradingBotGridAiParamV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLever

`func (o *GetTradingBotGridAiParamV5RespData) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradingBotGridAiParamV5RespData) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradingBotGridAiParamV5RespData) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradingBotGridAiParamV5RespData) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetMaxPx

`func (o *GetTradingBotGridAiParamV5RespData) GetMaxPx() string`

GetMaxPx returns the MaxPx field if non-nil, zero value otherwise.

### GetMaxPxOk

`func (o *GetTradingBotGridAiParamV5RespData) GetMaxPxOk() (*string, bool)`

GetMaxPxOk returns a tuple with the MaxPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPx

`func (o *GetTradingBotGridAiParamV5RespData) SetMaxPx(v string)`

SetMaxPx sets MaxPx field to given value.

### HasMaxPx

`func (o *GetTradingBotGridAiParamV5RespData) HasMaxPx() bool`

HasMaxPx returns a boolean if a field has been set.

### GetMinInvestment

`func (o *GetTradingBotGridAiParamV5RespData) GetMinInvestment() string`

GetMinInvestment returns the MinInvestment field if non-nil, zero value otherwise.

### GetMinInvestmentOk

`func (o *GetTradingBotGridAiParamV5RespData) GetMinInvestmentOk() (*string, bool)`

GetMinInvestmentOk returns a tuple with the MinInvestment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInvestment

`func (o *GetTradingBotGridAiParamV5RespData) SetMinInvestment(v string)`

SetMinInvestment sets MinInvestment field to given value.

### HasMinInvestment

`func (o *GetTradingBotGridAiParamV5RespData) HasMinInvestment() bool`

HasMinInvestment returns a boolean if a field has been set.

### GetMinPx

`func (o *GetTradingBotGridAiParamV5RespData) GetMinPx() string`

GetMinPx returns the MinPx field if non-nil, zero value otherwise.

### GetMinPxOk

`func (o *GetTradingBotGridAiParamV5RespData) GetMinPxOk() (*string, bool)`

GetMinPxOk returns a tuple with the MinPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPx

`func (o *GetTradingBotGridAiParamV5RespData) SetMinPx(v string)`

SetMinPx sets MinPx field to given value.

### HasMinPx

`func (o *GetTradingBotGridAiParamV5RespData) HasMinPx() bool`

HasMinPx returns a boolean if a field has been set.

### GetPerGridProfitRatio

`func (o *GetTradingBotGridAiParamV5RespData) GetPerGridProfitRatio() string`

GetPerGridProfitRatio returns the PerGridProfitRatio field if non-nil, zero value otherwise.

### GetPerGridProfitRatioOk

`func (o *GetTradingBotGridAiParamV5RespData) GetPerGridProfitRatioOk() (*string, bool)`

GetPerGridProfitRatioOk returns a tuple with the PerGridProfitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerGridProfitRatio

`func (o *GetTradingBotGridAiParamV5RespData) SetPerGridProfitRatio(v string)`

SetPerGridProfitRatio sets PerGridProfitRatio field to given value.

### HasPerGridProfitRatio

`func (o *GetTradingBotGridAiParamV5RespData) HasPerGridProfitRatio() bool`

HasPerGridProfitRatio returns a boolean if a field has been set.

### GetPerMaxProfitRate

`func (o *GetTradingBotGridAiParamV5RespData) GetPerMaxProfitRate() string`

GetPerMaxProfitRate returns the PerMaxProfitRate field if non-nil, zero value otherwise.

### GetPerMaxProfitRateOk

`func (o *GetTradingBotGridAiParamV5RespData) GetPerMaxProfitRateOk() (*string, bool)`

GetPerMaxProfitRateOk returns a tuple with the PerMaxProfitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerMaxProfitRate

`func (o *GetTradingBotGridAiParamV5RespData) SetPerMaxProfitRate(v string)`

SetPerMaxProfitRate sets PerMaxProfitRate field to given value.

### HasPerMaxProfitRate

`func (o *GetTradingBotGridAiParamV5RespData) HasPerMaxProfitRate() bool`

HasPerMaxProfitRate returns a boolean if a field has been set.

### GetPerMinProfitRate

`func (o *GetTradingBotGridAiParamV5RespData) GetPerMinProfitRate() string`

GetPerMinProfitRate returns the PerMinProfitRate field if non-nil, zero value otherwise.

### GetPerMinProfitRateOk

`func (o *GetTradingBotGridAiParamV5RespData) GetPerMinProfitRateOk() (*string, bool)`

GetPerMinProfitRateOk returns a tuple with the PerMinProfitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerMinProfitRate

`func (o *GetTradingBotGridAiParamV5RespData) SetPerMinProfitRate(v string)`

SetPerMinProfitRate sets PerMinProfitRate field to given value.

### HasPerMinProfitRate

`func (o *GetTradingBotGridAiParamV5RespData) HasPerMinProfitRate() bool`

HasPerMinProfitRate returns a boolean if a field has been set.

### GetRunType

`func (o *GetTradingBotGridAiParamV5RespData) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *GetTradingBotGridAiParamV5RespData) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *GetTradingBotGridAiParamV5RespData) SetRunType(v string)`

SetRunType sets RunType field to given value.

### HasRunType

`func (o *GetTradingBotGridAiParamV5RespData) HasRunType() bool`

HasRunType returns a boolean if a field has been set.

### GetSourceCcy

`func (o *GetTradingBotGridAiParamV5RespData) GetSourceCcy() string`

GetSourceCcy returns the SourceCcy field if non-nil, zero value otherwise.

### GetSourceCcyOk

`func (o *GetTradingBotGridAiParamV5RespData) GetSourceCcyOk() (*string, bool)`

GetSourceCcyOk returns a tuple with the SourceCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCcy

`func (o *GetTradingBotGridAiParamV5RespData) SetSourceCcy(v string)`

SetSourceCcy sets SourceCcy field to given value.

### HasSourceCcy

`func (o *GetTradingBotGridAiParamV5RespData) HasSourceCcy() bool`

HasSourceCcy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


