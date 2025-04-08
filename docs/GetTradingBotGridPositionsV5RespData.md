# GetTradingBotGridPositionsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adl** | Pointer to **string** | Auto decrease line, signal area  Divided into 5 levels, from 1 to 5, the smaller the number, the weaker the adl intensity. | [optional] [default to ""]
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID | [optional] [default to ""]
**AvgPx** | Pointer to **string** | Average open price | [optional] [default to ""]
**CTime** | Pointer to **string** | Algo order created time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Ccy** | Pointer to **string** | Margin currency | [optional] [default to ""]
**Imr** | Pointer to **string** | Initial margin requirement | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Last** | Pointer to **string** | Latest traded price | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage | [optional] [default to ""]
**LiqPx** | Pointer to **string** | Estimated liquidation price | [optional] [default to ""]
**MarkPx** | Pointer to **string** | Mark price | [optional] [default to ""]
**MgnMode** | Pointer to **string** | Margin mode  &#x60;cross&#x60;  &#x60;isolated&#x60; | [optional] [default to ""]
**MgnRatio** | Pointer to **string** | Margin ratio | [optional] [default to ""]
**Mmr** | Pointer to **string** | Maintenance margin requirement | [optional] [default to ""]
**NotionalUsd** | Pointer to **string** | Notional value of positions in &#x60;USD&#x60; | [optional] [default to ""]
**Pos** | Pointer to **string** | Quantity of positions | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side  &#x60;net&#x60; | [optional] [default to ""]
**UTime** | Pointer to **string** | Algo order updated time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Upl** | Pointer to **string** | Unrealized profit and loss | [optional] [default to ""]
**UplRatio** | Pointer to **string** | Unrealized profit and loss ratio | [optional] [default to ""]

## Methods

### NewGetTradingBotGridPositionsV5RespData

`func NewGetTradingBotGridPositionsV5RespData() *GetTradingBotGridPositionsV5RespData`

NewGetTradingBotGridPositionsV5RespData instantiates a new GetTradingBotGridPositionsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotGridPositionsV5RespDataWithDefaults

`func NewGetTradingBotGridPositionsV5RespDataWithDefaults() *GetTradingBotGridPositionsV5RespData`

NewGetTradingBotGridPositionsV5RespDataWithDefaults instantiates a new GetTradingBotGridPositionsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdl

`func (o *GetTradingBotGridPositionsV5RespData) GetAdl() string`

GetAdl returns the Adl field if non-nil, zero value otherwise.

### GetAdlOk

`func (o *GetTradingBotGridPositionsV5RespData) GetAdlOk() (*string, bool)`

GetAdlOk returns a tuple with the Adl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdl

`func (o *GetTradingBotGridPositionsV5RespData) SetAdl(v string)`

SetAdl sets Adl field to given value.

### HasAdl

`func (o *GetTradingBotGridPositionsV5RespData) HasAdl() bool`

HasAdl returns a boolean if a field has been set.

### GetAlgoClOrdId

`func (o *GetTradingBotGridPositionsV5RespData) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotGridPositionsV5RespData) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotGridPositionsV5RespData) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotGridPositionsV5RespData) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotGridPositionsV5RespData) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotGridPositionsV5RespData) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotGridPositionsV5RespData) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotGridPositionsV5RespData) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAvgPx

`func (o *GetTradingBotGridPositionsV5RespData) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *GetTradingBotGridPositionsV5RespData) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *GetTradingBotGridPositionsV5RespData) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *GetTradingBotGridPositionsV5RespData) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotGridPositionsV5RespData) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotGridPositionsV5RespData) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotGridPositionsV5RespData) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotGridPositionsV5RespData) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCcy

`func (o *GetTradingBotGridPositionsV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetTradingBotGridPositionsV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetTradingBotGridPositionsV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetTradingBotGridPositionsV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetImr

`func (o *GetTradingBotGridPositionsV5RespData) GetImr() string`

GetImr returns the Imr field if non-nil, zero value otherwise.

### GetImrOk

`func (o *GetTradingBotGridPositionsV5RespData) GetImrOk() (*string, bool)`

GetImrOk returns a tuple with the Imr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImr

`func (o *GetTradingBotGridPositionsV5RespData) SetImr(v string)`

SetImr sets Imr field to given value.

### HasImr

`func (o *GetTradingBotGridPositionsV5RespData) HasImr() bool`

HasImr returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradingBotGridPositionsV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradingBotGridPositionsV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradingBotGridPositionsV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradingBotGridPositionsV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotGridPositionsV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotGridPositionsV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotGridPositionsV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotGridPositionsV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLast

`func (o *GetTradingBotGridPositionsV5RespData) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetTradingBotGridPositionsV5RespData) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetTradingBotGridPositionsV5RespData) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GetTradingBotGridPositionsV5RespData) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLever

`func (o *GetTradingBotGridPositionsV5RespData) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradingBotGridPositionsV5RespData) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradingBotGridPositionsV5RespData) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradingBotGridPositionsV5RespData) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetLiqPx

`func (o *GetTradingBotGridPositionsV5RespData) GetLiqPx() string`

GetLiqPx returns the LiqPx field if non-nil, zero value otherwise.

### GetLiqPxOk

`func (o *GetTradingBotGridPositionsV5RespData) GetLiqPxOk() (*string, bool)`

GetLiqPxOk returns a tuple with the LiqPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqPx

`func (o *GetTradingBotGridPositionsV5RespData) SetLiqPx(v string)`

SetLiqPx sets LiqPx field to given value.

### HasLiqPx

`func (o *GetTradingBotGridPositionsV5RespData) HasLiqPx() bool`

HasLiqPx returns a boolean if a field has been set.

### GetMarkPx

`func (o *GetTradingBotGridPositionsV5RespData) GetMarkPx() string`

GetMarkPx returns the MarkPx field if non-nil, zero value otherwise.

### GetMarkPxOk

`func (o *GetTradingBotGridPositionsV5RespData) GetMarkPxOk() (*string, bool)`

GetMarkPxOk returns a tuple with the MarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPx

`func (o *GetTradingBotGridPositionsV5RespData) SetMarkPx(v string)`

SetMarkPx sets MarkPx field to given value.

### HasMarkPx

`func (o *GetTradingBotGridPositionsV5RespData) HasMarkPx() bool`

HasMarkPx returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetTradingBotGridPositionsV5RespData) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetTradingBotGridPositionsV5RespData) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetTradingBotGridPositionsV5RespData) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetTradingBotGridPositionsV5RespData) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetMgnRatio

`func (o *GetTradingBotGridPositionsV5RespData) GetMgnRatio() string`

GetMgnRatio returns the MgnRatio field if non-nil, zero value otherwise.

### GetMgnRatioOk

`func (o *GetTradingBotGridPositionsV5RespData) GetMgnRatioOk() (*string, bool)`

GetMgnRatioOk returns a tuple with the MgnRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnRatio

`func (o *GetTradingBotGridPositionsV5RespData) SetMgnRatio(v string)`

SetMgnRatio sets MgnRatio field to given value.

### HasMgnRatio

`func (o *GetTradingBotGridPositionsV5RespData) HasMgnRatio() bool`

HasMgnRatio returns a boolean if a field has been set.

### GetMmr

`func (o *GetTradingBotGridPositionsV5RespData) GetMmr() string`

GetMmr returns the Mmr field if non-nil, zero value otherwise.

### GetMmrOk

`func (o *GetTradingBotGridPositionsV5RespData) GetMmrOk() (*string, bool)`

GetMmrOk returns a tuple with the Mmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmr

`func (o *GetTradingBotGridPositionsV5RespData) SetMmr(v string)`

SetMmr sets Mmr field to given value.

### HasMmr

`func (o *GetTradingBotGridPositionsV5RespData) HasMmr() bool`

HasMmr returns a boolean if a field has been set.

### GetNotionalUsd

`func (o *GetTradingBotGridPositionsV5RespData) GetNotionalUsd() string`

GetNotionalUsd returns the NotionalUsd field if non-nil, zero value otherwise.

### GetNotionalUsdOk

`func (o *GetTradingBotGridPositionsV5RespData) GetNotionalUsdOk() (*string, bool)`

GetNotionalUsdOk returns a tuple with the NotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsd

`func (o *GetTradingBotGridPositionsV5RespData) SetNotionalUsd(v string)`

SetNotionalUsd sets NotionalUsd field to given value.

### HasNotionalUsd

`func (o *GetTradingBotGridPositionsV5RespData) HasNotionalUsd() bool`

HasNotionalUsd returns a boolean if a field has been set.

### GetPos

`func (o *GetTradingBotGridPositionsV5RespData) GetPos() string`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *GetTradingBotGridPositionsV5RespData) GetPosOk() (*string, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *GetTradingBotGridPositionsV5RespData) SetPos(v string)`

SetPos sets Pos field to given value.

### HasPos

`func (o *GetTradingBotGridPositionsV5RespData) HasPos() bool`

HasPos returns a boolean if a field has been set.

### GetPosSide

`func (o *GetTradingBotGridPositionsV5RespData) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetTradingBotGridPositionsV5RespData) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetTradingBotGridPositionsV5RespData) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetTradingBotGridPositionsV5RespData) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotGridPositionsV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotGridPositionsV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotGridPositionsV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotGridPositionsV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUpl

`func (o *GetTradingBotGridPositionsV5RespData) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *GetTradingBotGridPositionsV5RespData) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *GetTradingBotGridPositionsV5RespData) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *GetTradingBotGridPositionsV5RespData) HasUpl() bool`

HasUpl returns a boolean if a field has been set.

### GetUplRatio

`func (o *GetTradingBotGridPositionsV5RespData) GetUplRatio() string`

GetUplRatio returns the UplRatio field if non-nil, zero value otherwise.

### GetUplRatioOk

`func (o *GetTradingBotGridPositionsV5RespData) GetUplRatioOk() (*string, bool)`

GetUplRatioOk returns a tuple with the UplRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplRatio

`func (o *GetTradingBotGridPositionsV5RespData) SetUplRatio(v string)`

SetUplRatio sets UplRatio field to given value.

### HasUplRatio

`func (o *GetTradingBotGridPositionsV5RespData) HasUplRatio() bool`

HasUplRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


