# GetCopytradingSubpositionsHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Currency | [optional] [default to ""]
**CloseAvgPx** | Pointer to **string** | Average price of closing position | [optional] [default to ""]
**CloseSubPos** | Pointer to **string** | Quantity of positions that is already closed | [optional] [default to ""]
**CloseTime** | Pointer to **string** | Time of closing position | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage | [optional] [default to ""]
**Margin** | Pointer to **string** | Margin | [optional] [default to ""]
**MarkPx** | Pointer to **string** | Latest mark price, only applicable to contract | [optional] [default to ""]
**MgnMode** | Pointer to **string** | Margin mode. &#x60;cross&#x60; &#x60;isolated&#x60; | [optional] [default to ""]
**OpenAvgPx** | Pointer to **string** | Average open price | [optional] [default to ""]
**OpenOrdId** | Pointer to **string** | Order ID for opening position, only applicable to lead position | [optional] [default to ""]
**OpenTime** | Pointer to **string** | Time of opening | [optional] [default to ""]
**Pnl** | Pointer to **string** | Profit and loss | [optional] [default to ""]
**PnlRatio** | Pointer to **string** | P&amp;L ratio | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side  &#x60;long&#x60;   &#x60;short&#x60;   &#x60;net&#x60;  (long position has positive subPos; short position has negative subPos) | [optional] [default to ""]
**ProfitSharingAmt** | Pointer to **string** | Profit sharing amount, only applicable to copy trading. Note: this parameter is already deprecated. | [optional] [default to ""]
**SubPos** | Pointer to **string** | Quantity of positions | [optional] [default to ""]
**SubPosId** | Pointer to **string** | Lead position ID | [optional] [default to ""]
**Type** | Pointer to **string** | The type of closing position  &#x60;1&#x60;：Close position partially;  &#x60;2&#x60;：Close all | [optional] [default to ""]
**UniqueCode** | Pointer to **string** | Lead trader unique code | [optional] [default to ""]

## Methods

### NewGetCopytradingSubpositionsHistoryV5RespDataInner

`func NewGetCopytradingSubpositionsHistoryV5RespDataInner() *GetCopytradingSubpositionsHistoryV5RespDataInner`

NewGetCopytradingSubpositionsHistoryV5RespDataInner instantiates a new GetCopytradingSubpositionsHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingSubpositionsHistoryV5RespDataInnerWithDefaults

`func NewGetCopytradingSubpositionsHistoryV5RespDataInnerWithDefaults() *GetCopytradingSubpositionsHistoryV5RespDataInner`

NewGetCopytradingSubpositionsHistoryV5RespDataInnerWithDefaults instantiates a new GetCopytradingSubpositionsHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetCloseAvgPx

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCloseAvgPx() string`

GetCloseAvgPx returns the CloseAvgPx field if non-nil, zero value otherwise.

### GetCloseAvgPxOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCloseAvgPxOk() (*string, bool)`

GetCloseAvgPxOk returns a tuple with the CloseAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseAvgPx

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetCloseAvgPx(v string)`

SetCloseAvgPx sets CloseAvgPx field to given value.

### HasCloseAvgPx

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasCloseAvgPx() bool`

HasCloseAvgPx returns a boolean if a field has been set.

### GetCloseSubPos

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCloseSubPos() string`

GetCloseSubPos returns the CloseSubPos field if non-nil, zero value otherwise.

### GetCloseSubPosOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCloseSubPosOk() (*string, bool)`

GetCloseSubPosOk returns a tuple with the CloseSubPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseSubPos

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetCloseSubPos(v string)`

SetCloseSubPos sets CloseSubPos field to given value.

### HasCloseSubPos

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasCloseSubPos() bool`

HasCloseSubPos returns a boolean if a field has been set.

### GetCloseTime

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCloseTime() string`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetCloseTimeOk() (*string, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetCloseTime(v string)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetInstId

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLever

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetMargin

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetMargin() string`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetMarginOk() (*string, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetMargin(v string)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetMarkPx

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetMarkPx() string`

GetMarkPx returns the MarkPx field if non-nil, zero value otherwise.

### GetMarkPxOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetMarkPxOk() (*string, bool)`

GetMarkPxOk returns a tuple with the MarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPx

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetMarkPx(v string)`

SetMarkPx sets MarkPx field to given value.

### HasMarkPx

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasMarkPx() bool`

HasMarkPx returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetOpenAvgPx

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetOpenAvgPx() string`

GetOpenAvgPx returns the OpenAvgPx field if non-nil, zero value otherwise.

### GetOpenAvgPxOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetOpenAvgPxOk() (*string, bool)`

GetOpenAvgPxOk returns a tuple with the OpenAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAvgPx

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetOpenAvgPx(v string)`

SetOpenAvgPx sets OpenAvgPx field to given value.

### HasOpenAvgPx

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasOpenAvgPx() bool`

HasOpenAvgPx returns a boolean if a field has been set.

### GetOpenOrdId

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetOpenOrdId() string`

GetOpenOrdId returns the OpenOrdId field if non-nil, zero value otherwise.

### GetOpenOrdIdOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetOpenOrdIdOk() (*string, bool)`

GetOpenOrdIdOk returns a tuple with the OpenOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrdId

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetOpenOrdId(v string)`

SetOpenOrdId sets OpenOrdId field to given value.

### HasOpenOrdId

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasOpenOrdId() bool`

HasOpenOrdId returns a boolean if a field has been set.

### GetOpenTime

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetOpenTime() string`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetOpenTimeOk() (*string, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetOpenTime(v string)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetPnl

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetPnlRatio

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetPnlRatio() string`

GetPnlRatio returns the PnlRatio field if non-nil, zero value otherwise.

### GetPnlRatioOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetPnlRatioOk() (*string, bool)`

GetPnlRatioOk returns a tuple with the PnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlRatio

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetPnlRatio(v string)`

SetPnlRatio sets PnlRatio field to given value.

### HasPnlRatio

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasPnlRatio() bool`

HasPnlRatio returns a boolean if a field has been set.

### GetPosSide

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetProfitSharingAmt

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetProfitSharingAmt() string`

GetProfitSharingAmt returns the ProfitSharingAmt field if non-nil, zero value otherwise.

### GetProfitSharingAmtOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetProfitSharingAmtOk() (*string, bool)`

GetProfitSharingAmtOk returns a tuple with the ProfitSharingAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitSharingAmt

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetProfitSharingAmt(v string)`

SetProfitSharingAmt sets ProfitSharingAmt field to given value.

### HasProfitSharingAmt

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasProfitSharingAmt() bool`

HasProfitSharingAmt returns a boolean if a field has been set.

### GetSubPos

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetSubPos() string`

GetSubPos returns the SubPos field if non-nil, zero value otherwise.

### GetSubPosOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetSubPosOk() (*string, bool)`

GetSubPosOk returns a tuple with the SubPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPos

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetSubPos(v string)`

SetSubPos sets SubPos field to given value.

### HasSubPos

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasSubPos() bool`

HasSubPos returns a boolean if a field has been set.

### GetSubPosId

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetSubPosId() string`

GetSubPosId returns the SubPosId field if non-nil, zero value otherwise.

### GetSubPosIdOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetSubPosIdOk() (*string, bool)`

GetSubPosIdOk returns a tuple with the SubPosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPosId

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetSubPosId(v string)`

SetSubPosId sets SubPosId field to given value.

### HasSubPosId

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasSubPosId() bool`

HasSubPosId returns a boolean if a field has been set.

### GetType

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUniqueCode

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetUniqueCode() string`

GetUniqueCode returns the UniqueCode field if non-nil, zero value otherwise.

### GetUniqueCodeOk

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) GetUniqueCodeOk() (*string, bool)`

GetUniqueCodeOk returns a tuple with the UniqueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCode

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) SetUniqueCode(v string)`

SetUniqueCode sets UniqueCode field to given value.

### HasUniqueCode

`func (o *GetCopytradingSubpositionsHistoryV5RespDataInner) HasUniqueCode() bool`

HasUniqueCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


