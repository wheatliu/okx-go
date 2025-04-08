# GetAccountAccountPositionRiskV5RespDataInnerPosDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseBal** | Pointer to **string** | Base currency balance, only applicable to &#x60;MARGIN&#x60;（Quick Margin Mode）(Deprecated) | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency used for margin | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**MgnMode** | Pointer to **string** | Margin mode  &#x60;cross&#x60;    &#x60;isolated&#x60; | [optional] [default to ""]
**NotionalCcy** | Pointer to **string** | Notional value of positions in &#x60;coin&#x60; | [optional] [default to ""]
**NotionalUsd** | Pointer to **string** | Notional value of positions in &#x60;USD&#x60; | [optional] [default to ""]
**Pos** | Pointer to **string** | Quantity of positions &#x60;contract&#x60;. In the isolated margin mode, when doing manual transfers, a position with pos of &#x60;0&#x60; will be generated after the deposit is transferred | [optional] [default to ""]
**PosCcy** | Pointer to **string** | Position currency, only applicable to &#x60;MARGIN&#x60; positions. | [optional] [default to ""]
**PosId** | Pointer to **string** | Position ID | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side  &#x60;long&#x60;   &#x60;short&#x60;   &#x60;net&#x60; (&#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;: positive &#x60;pos&#x60; means long position and negative &#x60;pos&#x60; means short position. &#x60;MARGIN&#x60;: &#x60;posCcy&#x60; being base currency means long position, &#x60;posCcy&#x60; being quote currency means short position.) | [optional] [default to ""]
**QuoteBal** | Pointer to **string** | Quote currency balance, only applicable to &#x60;MARGIN&#x60;（Quick Margin Mode）(Deprecated) | [optional] [default to ""]

## Methods

### NewGetAccountAccountPositionRiskV5RespDataInnerPosDataInner

`func NewGetAccountAccountPositionRiskV5RespDataInnerPosDataInner() *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner`

NewGetAccountAccountPositionRiskV5RespDataInnerPosDataInner instantiates a new GetAccountAccountPositionRiskV5RespDataInnerPosDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountAccountPositionRiskV5RespDataInnerPosDataInnerWithDefaults

`func NewGetAccountAccountPositionRiskV5RespDataInnerPosDataInnerWithDefaults() *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner`

NewGetAccountAccountPositionRiskV5RespDataInnerPosDataInnerWithDefaults instantiates a new GetAccountAccountPositionRiskV5RespDataInnerPosDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseBal

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetBaseBal() string`

GetBaseBal returns the BaseBal field if non-nil, zero value otherwise.

### GetBaseBalOk

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetBaseBalOk() (*string, bool)`

GetBaseBalOk returns a tuple with the BaseBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBal

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) SetBaseBal(v string)`

SetBaseBal sets BaseBal field to given value.

### HasBaseBal

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) HasBaseBal() bool`

HasBaseBal returns a boolean if a field has been set.

### GetCcy

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetNotionalCcy

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetNotionalCcy() string`

GetNotionalCcy returns the NotionalCcy field if non-nil, zero value otherwise.

### GetNotionalCcyOk

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetNotionalCcyOk() (*string, bool)`

GetNotionalCcyOk returns a tuple with the NotionalCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalCcy

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) SetNotionalCcy(v string)`

SetNotionalCcy sets NotionalCcy field to given value.

### HasNotionalCcy

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) HasNotionalCcy() bool`

HasNotionalCcy returns a boolean if a field has been set.

### GetNotionalUsd

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetNotionalUsd() string`

GetNotionalUsd returns the NotionalUsd field if non-nil, zero value otherwise.

### GetNotionalUsdOk

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetNotionalUsdOk() (*string, bool)`

GetNotionalUsdOk returns a tuple with the NotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsd

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) SetNotionalUsd(v string)`

SetNotionalUsd sets NotionalUsd field to given value.

### HasNotionalUsd

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) HasNotionalUsd() bool`

HasNotionalUsd returns a boolean if a field has been set.

### GetPos

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetPos() string`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetPosOk() (*string, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) SetPos(v string)`

SetPos sets Pos field to given value.

### HasPos

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) HasPos() bool`

HasPos returns a boolean if a field has been set.

### GetPosCcy

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetPosCcy() string`

GetPosCcy returns the PosCcy field if non-nil, zero value otherwise.

### GetPosCcyOk

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetPosCcyOk() (*string, bool)`

GetPosCcyOk returns a tuple with the PosCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosCcy

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) SetPosCcy(v string)`

SetPosCcy sets PosCcy field to given value.

### HasPosCcy

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) HasPosCcy() bool`

HasPosCcy returns a boolean if a field has been set.

### GetPosId

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetPosId() string`

GetPosId returns the PosId field if non-nil, zero value otherwise.

### GetPosIdOk

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetPosIdOk() (*string, bool)`

GetPosIdOk returns a tuple with the PosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosId

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) SetPosId(v string)`

SetPosId sets PosId field to given value.

### HasPosId

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) HasPosId() bool`

HasPosId returns a boolean if a field has been set.

### GetPosSide

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetQuoteBal

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetQuoteBal() string`

GetQuoteBal returns the QuoteBal field if non-nil, zero value otherwise.

### GetQuoteBalOk

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) GetQuoteBalOk() (*string, bool)`

GetQuoteBalOk returns a tuple with the QuoteBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteBal

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) SetQuoteBal(v string)`

SetQuoteBal sets QuoteBal field to given value.

### HasQuoteBal

`func (o *GetAccountAccountPositionRiskV5RespDataInnerPosDataInner) HasQuoteBal() bool`

HasQuoteBal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


