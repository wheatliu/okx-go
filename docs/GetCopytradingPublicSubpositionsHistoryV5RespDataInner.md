# GetCopytradingPublicSubpositionsHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Currency | [optional] [default to ""]
**CloseAvgPx** | Pointer to **string** | Average price of closing position | [optional] [default to ""]
**CloseTime** | Pointer to **string** | Time of closing position | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage | [optional] [default to ""]
**Margin** | Pointer to **string** | Margin | [optional] [default to ""]
**MgnMode** | Pointer to **string** | Margin mode. &#x60;cross&#x60; &#x60;isolated&#x60; | [optional] [default to ""]
**OpenAvgPx** | Pointer to **string** | Average open price | [optional] [default to ""]
**OpenTime** | Pointer to **string** | Time of opening | [optional] [default to ""]
**Pnl** | Pointer to **string** | Profit and loss | [optional] [default to ""]
**PnlRatio** | Pointer to **string** | P&amp;L ratio | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side  &#x60;long&#x60;   &#x60;short&#x60;   &#x60;net&#x60;  (long position has positive subPos; short position has negative subPos) | [optional] [default to ""]
**SubPos** | Pointer to **string** | Quantity of positions | [optional] [default to ""]
**SubPosId** | Pointer to **string** | Lead position ID | [optional] [default to ""]
**UniqueCode** | Pointer to **string** | Lead trader unique code | [optional] [default to ""]

## Methods

### NewGetCopytradingPublicSubpositionsHistoryV5RespDataInner

`func NewGetCopytradingPublicSubpositionsHistoryV5RespDataInner() *GetCopytradingPublicSubpositionsHistoryV5RespDataInner`

NewGetCopytradingPublicSubpositionsHistoryV5RespDataInner instantiates a new GetCopytradingPublicSubpositionsHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingPublicSubpositionsHistoryV5RespDataInnerWithDefaults

`func NewGetCopytradingPublicSubpositionsHistoryV5RespDataInnerWithDefaults() *GetCopytradingPublicSubpositionsHistoryV5RespDataInner`

NewGetCopytradingPublicSubpositionsHistoryV5RespDataInnerWithDefaults instantiates a new GetCopytradingPublicSubpositionsHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetCloseAvgPx

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetCloseAvgPx() string`

GetCloseAvgPx returns the CloseAvgPx field if non-nil, zero value otherwise.

### GetCloseAvgPxOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetCloseAvgPxOk() (*string, bool)`

GetCloseAvgPxOk returns a tuple with the CloseAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseAvgPx

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetCloseAvgPx(v string)`

SetCloseAvgPx sets CloseAvgPx field to given value.

### HasCloseAvgPx

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasCloseAvgPx() bool`

HasCloseAvgPx returns a boolean if a field has been set.

### GetCloseTime

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetCloseTime() string`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetCloseTimeOk() (*string, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetCloseTime(v string)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetInstId

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLever

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetMargin

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetMargin() string`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetMarginOk() (*string, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetMargin(v string)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetOpenAvgPx

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetOpenAvgPx() string`

GetOpenAvgPx returns the OpenAvgPx field if non-nil, zero value otherwise.

### GetOpenAvgPxOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetOpenAvgPxOk() (*string, bool)`

GetOpenAvgPxOk returns a tuple with the OpenAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAvgPx

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetOpenAvgPx(v string)`

SetOpenAvgPx sets OpenAvgPx field to given value.

### HasOpenAvgPx

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasOpenAvgPx() bool`

HasOpenAvgPx returns a boolean if a field has been set.

### GetOpenTime

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetOpenTime() string`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetOpenTimeOk() (*string, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetOpenTime(v string)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetPnl

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetPnlRatio

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetPnlRatio() string`

GetPnlRatio returns the PnlRatio field if non-nil, zero value otherwise.

### GetPnlRatioOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetPnlRatioOk() (*string, bool)`

GetPnlRatioOk returns a tuple with the PnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlRatio

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetPnlRatio(v string)`

SetPnlRatio sets PnlRatio field to given value.

### HasPnlRatio

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasPnlRatio() bool`

HasPnlRatio returns a boolean if a field has been set.

### GetPosSide

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetSubPos

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetSubPos() string`

GetSubPos returns the SubPos field if non-nil, zero value otherwise.

### GetSubPosOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetSubPosOk() (*string, bool)`

GetSubPosOk returns a tuple with the SubPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPos

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetSubPos(v string)`

SetSubPos sets SubPos field to given value.

### HasSubPos

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasSubPos() bool`

HasSubPos returns a boolean if a field has been set.

### GetSubPosId

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetSubPosId() string`

GetSubPosId returns the SubPosId field if non-nil, zero value otherwise.

### GetSubPosIdOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetSubPosIdOk() (*string, bool)`

GetSubPosIdOk returns a tuple with the SubPosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPosId

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetSubPosId(v string)`

SetSubPosId sets SubPosId field to given value.

### HasSubPosId

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasSubPosId() bool`

HasSubPosId returns a boolean if a field has been set.

### GetUniqueCode

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetUniqueCode() string`

GetUniqueCode returns the UniqueCode field if non-nil, zero value otherwise.

### GetUniqueCodeOk

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) GetUniqueCodeOk() (*string, bool)`

GetUniqueCodeOk returns a tuple with the UniqueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCode

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) SetUniqueCode(v string)`

SetUniqueCode sets UniqueCode field to given value.

### HasUniqueCode

`func (o *GetCopytradingPublicSubpositionsHistoryV5RespDataInner) HasUniqueCode() bool`

HasUniqueCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


