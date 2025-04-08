# GetTradingBotGridPositionsV5RespDataInner

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

### NewGetTradingBotGridPositionsV5RespDataInner

`func NewGetTradingBotGridPositionsV5RespDataInner() *GetTradingBotGridPositionsV5RespDataInner`

NewGetTradingBotGridPositionsV5RespDataInner instantiates a new GetTradingBotGridPositionsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotGridPositionsV5RespDataInnerWithDefaults

`func NewGetTradingBotGridPositionsV5RespDataInnerWithDefaults() *GetTradingBotGridPositionsV5RespDataInner`

NewGetTradingBotGridPositionsV5RespDataInnerWithDefaults instantiates a new GetTradingBotGridPositionsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdl

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetAdl() string`

GetAdl returns the Adl field if non-nil, zero value otherwise.

### GetAdlOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetAdlOk() (*string, bool)`

GetAdlOk returns a tuple with the Adl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdl

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetAdl(v string)`

SetAdl sets Adl field to given value.

### HasAdl

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasAdl() bool`

HasAdl returns a boolean if a field has been set.

### GetAlgoClOrdId

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAvgPx

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetCTime

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCcy

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetImr

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetImr() string`

GetImr returns the Imr field if non-nil, zero value otherwise.

### GetImrOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetImrOk() (*string, bool)`

GetImrOk returns a tuple with the Imr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImr

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetImr(v string)`

SetImr sets Imr field to given value.

### HasImr

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasImr() bool`

HasImr returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLast

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLever

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetLiqPx

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetLiqPx() string`

GetLiqPx returns the LiqPx field if non-nil, zero value otherwise.

### GetLiqPxOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetLiqPxOk() (*string, bool)`

GetLiqPxOk returns a tuple with the LiqPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqPx

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetLiqPx(v string)`

SetLiqPx sets LiqPx field to given value.

### HasLiqPx

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasLiqPx() bool`

HasLiqPx returns a boolean if a field has been set.

### GetMarkPx

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetMarkPx() string`

GetMarkPx returns the MarkPx field if non-nil, zero value otherwise.

### GetMarkPxOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetMarkPxOk() (*string, bool)`

GetMarkPxOk returns a tuple with the MarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPx

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetMarkPx(v string)`

SetMarkPx sets MarkPx field to given value.

### HasMarkPx

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasMarkPx() bool`

HasMarkPx returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetMgnRatio

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetMgnRatio() string`

GetMgnRatio returns the MgnRatio field if non-nil, zero value otherwise.

### GetMgnRatioOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetMgnRatioOk() (*string, bool)`

GetMgnRatioOk returns a tuple with the MgnRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnRatio

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetMgnRatio(v string)`

SetMgnRatio sets MgnRatio field to given value.

### HasMgnRatio

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasMgnRatio() bool`

HasMgnRatio returns a boolean if a field has been set.

### GetMmr

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetMmr() string`

GetMmr returns the Mmr field if non-nil, zero value otherwise.

### GetMmrOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetMmrOk() (*string, bool)`

GetMmrOk returns a tuple with the Mmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmr

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetMmr(v string)`

SetMmr sets Mmr field to given value.

### HasMmr

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasMmr() bool`

HasMmr returns a boolean if a field has been set.

### GetNotionalUsd

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetNotionalUsd() string`

GetNotionalUsd returns the NotionalUsd field if non-nil, zero value otherwise.

### GetNotionalUsdOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetNotionalUsdOk() (*string, bool)`

GetNotionalUsdOk returns a tuple with the NotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsd

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetNotionalUsd(v string)`

SetNotionalUsd sets NotionalUsd field to given value.

### HasNotionalUsd

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasNotionalUsd() bool`

HasNotionalUsd returns a boolean if a field has been set.

### GetPos

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetPos() string`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetPosOk() (*string, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetPos(v string)`

SetPos sets Pos field to given value.

### HasPos

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasPos() bool`

HasPos returns a boolean if a field has been set.

### GetPosSide

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUpl

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasUpl() bool`

HasUpl returns a boolean if a field has been set.

### GetUplRatio

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetUplRatio() string`

GetUplRatio returns the UplRatio field if non-nil, zero value otherwise.

### GetUplRatioOk

`func (o *GetTradingBotGridPositionsV5RespDataInner) GetUplRatioOk() (*string, bool)`

GetUplRatioOk returns a tuple with the UplRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplRatio

`func (o *GetTradingBotGridPositionsV5RespDataInner) SetUplRatio(v string)`

SetUplRatio sets UplRatio field to given value.

### HasUplRatio

`func (o *GetTradingBotGridPositionsV5RespDataInner) HasUplRatio() bool`

HasUplRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


