# GetCopytradingPublicCurrentSubpositionsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Currency | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage | [optional] [default to ""]
**Margin** | Pointer to **string** | Margin | [optional] [default to ""]
**MarkPx** | Pointer to **string** | Latest mark price, only applicable to contract | [optional] [default to ""]
**MgnMode** | Pointer to **string** | Margin mode. &#x60;cross&#x60; &#x60;isolated&#x60; | [optional] [default to ""]
**OpenAvgPx** | Pointer to **string** | Average open price | [optional] [default to ""]
**OpenTime** | Pointer to **string** | Open time | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side  &#x60;long&#x60;   &#x60;short&#x60;   &#x60;net&#x60;  (Long positions have positive subPos; short positions have negative subPos) | [optional] [default to ""]
**SubPos** | Pointer to **string** | Quantity of positions | [optional] [default to ""]
**SubPosId** | Pointer to **string** | Lead position ID | [optional] [default to ""]
**UniqueCode** | Pointer to **string** | Lead trader unique code | [optional] [default to ""]
**Upl** | Pointer to **string** | Unrealized profit and loss | [optional] [default to ""]
**UplRatio** | Pointer to **string** | Unrealized profit and loss ratio | [optional] [default to ""]

## Methods

### NewGetCopytradingPublicCurrentSubpositionsV5RespDataInner

`func NewGetCopytradingPublicCurrentSubpositionsV5RespDataInner() *GetCopytradingPublicCurrentSubpositionsV5RespDataInner`

NewGetCopytradingPublicCurrentSubpositionsV5RespDataInner instantiates a new GetCopytradingPublicCurrentSubpositionsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingPublicCurrentSubpositionsV5RespDataInnerWithDefaults

`func NewGetCopytradingPublicCurrentSubpositionsV5RespDataInnerWithDefaults() *GetCopytradingPublicCurrentSubpositionsV5RespDataInner`

NewGetCopytradingPublicCurrentSubpositionsV5RespDataInnerWithDefaults instantiates a new GetCopytradingPublicCurrentSubpositionsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLever

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetMargin

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetMargin() string`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetMarginOk() (*string, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetMargin(v string)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetMarkPx

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetMarkPx() string`

GetMarkPx returns the MarkPx field if non-nil, zero value otherwise.

### GetMarkPxOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetMarkPxOk() (*string, bool)`

GetMarkPxOk returns a tuple with the MarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPx

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetMarkPx(v string)`

SetMarkPx sets MarkPx field to given value.

### HasMarkPx

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasMarkPx() bool`

HasMarkPx returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetOpenAvgPx

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetOpenAvgPx() string`

GetOpenAvgPx returns the OpenAvgPx field if non-nil, zero value otherwise.

### GetOpenAvgPxOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetOpenAvgPxOk() (*string, bool)`

GetOpenAvgPxOk returns a tuple with the OpenAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAvgPx

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetOpenAvgPx(v string)`

SetOpenAvgPx sets OpenAvgPx field to given value.

### HasOpenAvgPx

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasOpenAvgPx() bool`

HasOpenAvgPx returns a boolean if a field has been set.

### GetOpenTime

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetOpenTime() string`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetOpenTimeOk() (*string, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetOpenTime(v string)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetPosSide

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetSubPos

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetSubPos() string`

GetSubPos returns the SubPos field if non-nil, zero value otherwise.

### GetSubPosOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetSubPosOk() (*string, bool)`

GetSubPosOk returns a tuple with the SubPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPos

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetSubPos(v string)`

SetSubPos sets SubPos field to given value.

### HasSubPos

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasSubPos() bool`

HasSubPos returns a boolean if a field has been set.

### GetSubPosId

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetSubPosId() string`

GetSubPosId returns the SubPosId field if non-nil, zero value otherwise.

### GetSubPosIdOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetSubPosIdOk() (*string, bool)`

GetSubPosIdOk returns a tuple with the SubPosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPosId

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetSubPosId(v string)`

SetSubPosId sets SubPosId field to given value.

### HasSubPosId

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasSubPosId() bool`

HasSubPosId returns a boolean if a field has been set.

### GetUniqueCode

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetUniqueCode() string`

GetUniqueCode returns the UniqueCode field if non-nil, zero value otherwise.

### GetUniqueCodeOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetUniqueCodeOk() (*string, bool)`

GetUniqueCodeOk returns a tuple with the UniqueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCode

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetUniqueCode(v string)`

SetUniqueCode sets UniqueCode field to given value.

### HasUniqueCode

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasUniqueCode() bool`

HasUniqueCode returns a boolean if a field has been set.

### GetUpl

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasUpl() bool`

HasUpl returns a boolean if a field has been set.

### GetUplRatio

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetUplRatio() string`

GetUplRatio returns the UplRatio field if non-nil, zero value otherwise.

### GetUplRatioOk

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) GetUplRatioOk() (*string, bool)`

GetUplRatioOk returns a tuple with the UplRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplRatio

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) SetUplRatio(v string)`

SetUplRatio sets UplRatio field to given value.

### HasUplRatio

`func (o *GetCopytradingPublicCurrentSubpositionsV5RespDataInner) HasUplRatio() bool`

HasUplRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


