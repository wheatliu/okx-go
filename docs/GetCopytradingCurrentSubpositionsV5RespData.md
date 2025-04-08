# GetCopytradingCurrentSubpositionsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | Pointer to **string** | Stop order ID | [optional] [default to ""]
**AvailSubPos** | Pointer to **string** | Quantity of positions that can be closed | [optional] [default to ""]
**Ccy** | Pointer to **string** | Margin currency | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage | [optional] [default to ""]
**Margin** | Pointer to **string** | Margin | [optional] [default to ""]
**MarkPx** | Pointer to **string** | Latest mark price, only applicable to contract | [optional] [default to ""]
**MgnMode** | Pointer to **string** | Margin mode. &#x60;cross&#x60; &#x60;isolated&#x60; | [optional] [default to ""]
**OpenAvgPx** | Pointer to **string** | Average open price | [optional] [default to ""]
**OpenOrdId** | Pointer to **string** | Order ID for opening position, only applicable to lead position | [optional] [default to ""]
**OpenTime** | Pointer to **string** | Open time | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side  &#x60;long&#x60;   &#x60;short&#x60;   &#x60;net&#x60;  (Long positions have positive subPos; short positions have negative subPos) | [optional] [default to ""]
**SlOrdPx** | Pointer to **string** | Stop-loss order price, it is -1 for market price | [optional] [default to ""]
**SlTriggerPx** | Pointer to **string** | Stop-loss trigger price. | [optional] [default to ""]
**SubPos** | Pointer to **string** | Quantity of positions | [optional] [default to ""]
**SubPosId** | Pointer to **string** | Lead position ID | [optional] [default to ""]
**TpOrdPx** | Pointer to **string** | Take-profit order price, it is -1 for market price | [optional] [default to ""]
**TpTriggerPx** | Pointer to **string** | Take-profit trigger price. | [optional] [default to ""]
**UniqueCode** | Pointer to **string** | Lead trader unique code | [optional] [default to ""]
**Upl** | Pointer to **string** | Unrealized profit and loss | [optional] [default to ""]
**UplRatio** | Pointer to **string** | Unrealized profit and loss ratio | [optional] [default to ""]

## Methods

### NewGetCopytradingCurrentSubpositionsV5RespData

`func NewGetCopytradingCurrentSubpositionsV5RespData() *GetCopytradingCurrentSubpositionsV5RespData`

NewGetCopytradingCurrentSubpositionsV5RespData instantiates a new GetCopytradingCurrentSubpositionsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingCurrentSubpositionsV5RespDataWithDefaults

`func NewGetCopytradingCurrentSubpositionsV5RespDataWithDefaults() *GetCopytradingCurrentSubpositionsV5RespData`

NewGetCopytradingCurrentSubpositionsV5RespDataWithDefaults instantiates a new GetCopytradingCurrentSubpositionsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetAvailSubPos

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetAvailSubPos() string`

GetAvailSubPos returns the AvailSubPos field if non-nil, zero value otherwise.

### GetAvailSubPosOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetAvailSubPosOk() (*string, bool)`

GetAvailSubPosOk returns a tuple with the AvailSubPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailSubPos

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetAvailSubPos(v string)`

SetAvailSubPos sets AvailSubPos field to given value.

### HasAvailSubPos

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasAvailSubPos() bool`

HasAvailSubPos returns a boolean if a field has been set.

### GetCcy

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLever

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetMargin

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetMargin() string`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetMarginOk() (*string, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetMargin(v string)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetMarkPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetMarkPx() string`

GetMarkPx returns the MarkPx field if non-nil, zero value otherwise.

### GetMarkPxOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetMarkPxOk() (*string, bool)`

GetMarkPxOk returns a tuple with the MarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetMarkPx(v string)`

SetMarkPx sets MarkPx field to given value.

### HasMarkPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasMarkPx() bool`

HasMarkPx returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetOpenAvgPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetOpenAvgPx() string`

GetOpenAvgPx returns the OpenAvgPx field if non-nil, zero value otherwise.

### GetOpenAvgPxOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetOpenAvgPxOk() (*string, bool)`

GetOpenAvgPxOk returns a tuple with the OpenAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAvgPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetOpenAvgPx(v string)`

SetOpenAvgPx sets OpenAvgPx field to given value.

### HasOpenAvgPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasOpenAvgPx() bool`

HasOpenAvgPx returns a boolean if a field has been set.

### GetOpenOrdId

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetOpenOrdId() string`

GetOpenOrdId returns the OpenOrdId field if non-nil, zero value otherwise.

### GetOpenOrdIdOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetOpenOrdIdOk() (*string, bool)`

GetOpenOrdIdOk returns a tuple with the OpenOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenOrdId

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetOpenOrdId(v string)`

SetOpenOrdId sets OpenOrdId field to given value.

### HasOpenOrdId

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasOpenOrdId() bool`

HasOpenOrdId returns a boolean if a field has been set.

### GetOpenTime

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetOpenTime() string`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetOpenTimeOk() (*string, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetOpenTime(v string)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetPosSide

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetSlOrdPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetSlOrdPx() string`

GetSlOrdPx returns the SlOrdPx field if non-nil, zero value otherwise.

### GetSlOrdPxOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetSlOrdPxOk() (*string, bool)`

GetSlOrdPxOk returns a tuple with the SlOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlOrdPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetSlOrdPx(v string)`

SetSlOrdPx sets SlOrdPx field to given value.

### HasSlOrdPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasSlOrdPx() bool`

HasSlOrdPx returns a boolean if a field has been set.

### GetSlTriggerPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetSlTriggerPx() string`

GetSlTriggerPx returns the SlTriggerPx field if non-nil, zero value otherwise.

### GetSlTriggerPxOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetSlTriggerPxOk() (*string, bool)`

GetSlTriggerPxOk returns a tuple with the SlTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetSlTriggerPx(v string)`

SetSlTriggerPx sets SlTriggerPx field to given value.

### HasSlTriggerPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasSlTriggerPx() bool`

HasSlTriggerPx returns a boolean if a field has been set.

### GetSubPos

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetSubPos() string`

GetSubPos returns the SubPos field if non-nil, zero value otherwise.

### GetSubPosOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetSubPosOk() (*string, bool)`

GetSubPosOk returns a tuple with the SubPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPos

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetSubPos(v string)`

SetSubPos sets SubPos field to given value.

### HasSubPos

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasSubPos() bool`

HasSubPos returns a boolean if a field has been set.

### GetSubPosId

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetSubPosId() string`

GetSubPosId returns the SubPosId field if non-nil, zero value otherwise.

### GetSubPosIdOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetSubPosIdOk() (*string, bool)`

GetSubPosIdOk returns a tuple with the SubPosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPosId

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetSubPosId(v string)`

SetSubPosId sets SubPosId field to given value.

### HasSubPosId

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasSubPosId() bool`

HasSubPosId returns a boolean if a field has been set.

### GetTpOrdPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetTpOrdPx() string`

GetTpOrdPx returns the TpOrdPx field if non-nil, zero value otherwise.

### GetTpOrdPxOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetTpOrdPxOk() (*string, bool)`

GetTpOrdPxOk returns a tuple with the TpOrdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrdPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetTpOrdPx(v string)`

SetTpOrdPx sets TpOrdPx field to given value.

### HasTpOrdPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasTpOrdPx() bool`

HasTpOrdPx returns a boolean if a field has been set.

### GetTpTriggerPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetTpTriggerPx() string`

GetTpTriggerPx returns the TpTriggerPx field if non-nil, zero value otherwise.

### GetTpTriggerPxOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetTpTriggerPxOk() (*string, bool)`

GetTpTriggerPxOk returns a tuple with the TpTriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetTpTriggerPx(v string)`

SetTpTriggerPx sets TpTriggerPx field to given value.

### HasTpTriggerPx

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasTpTriggerPx() bool`

HasTpTriggerPx returns a boolean if a field has been set.

### GetUniqueCode

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetUniqueCode() string`

GetUniqueCode returns the UniqueCode field if non-nil, zero value otherwise.

### GetUniqueCodeOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetUniqueCodeOk() (*string, bool)`

GetUniqueCodeOk returns a tuple with the UniqueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCode

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetUniqueCode(v string)`

SetUniqueCode sets UniqueCode field to given value.

### HasUniqueCode

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasUniqueCode() bool`

HasUniqueCode returns a boolean if a field has been set.

### GetUpl

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasUpl() bool`

HasUpl returns a boolean if a field has been set.

### GetUplRatio

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetUplRatio() string`

GetUplRatio returns the UplRatio field if non-nil, zero value otherwise.

### GetUplRatioOk

`func (o *GetCopytradingCurrentSubpositionsV5RespData) GetUplRatioOk() (*string, bool)`

GetUplRatioOk returns a tuple with the UplRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplRatio

`func (o *GetCopytradingCurrentSubpositionsV5RespData) SetUplRatio(v string)`

SetUplRatio sets UplRatio field to given value.

### HasUplRatio

`func (o *GetCopytradingCurrentSubpositionsV5RespData) HasUplRatio() bool`

HasUplRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


