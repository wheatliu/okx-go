# GetAccountAccountPositionRiskV5RespDataPosDataInner

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

### NewGetAccountAccountPositionRiskV5RespDataPosDataInner

`func NewGetAccountAccountPositionRiskV5RespDataPosDataInner() *GetAccountAccountPositionRiskV5RespDataPosDataInner`

NewGetAccountAccountPositionRiskV5RespDataPosDataInner instantiates a new GetAccountAccountPositionRiskV5RespDataPosDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountAccountPositionRiskV5RespDataPosDataInnerWithDefaults

`func NewGetAccountAccountPositionRiskV5RespDataPosDataInnerWithDefaults() *GetAccountAccountPositionRiskV5RespDataPosDataInner`

NewGetAccountAccountPositionRiskV5RespDataPosDataInnerWithDefaults instantiates a new GetAccountAccountPositionRiskV5RespDataPosDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseBal

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetBaseBal() string`

GetBaseBal returns the BaseBal field if non-nil, zero value otherwise.

### GetBaseBalOk

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetBaseBalOk() (*string, bool)`

GetBaseBalOk returns a tuple with the BaseBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBal

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) SetBaseBal(v string)`

SetBaseBal sets BaseBal field to given value.

### HasBaseBal

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) HasBaseBal() bool`

HasBaseBal returns a boolean if a field has been set.

### GetCcy

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetNotionalCcy

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetNotionalCcy() string`

GetNotionalCcy returns the NotionalCcy field if non-nil, zero value otherwise.

### GetNotionalCcyOk

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetNotionalCcyOk() (*string, bool)`

GetNotionalCcyOk returns a tuple with the NotionalCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalCcy

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) SetNotionalCcy(v string)`

SetNotionalCcy sets NotionalCcy field to given value.

### HasNotionalCcy

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) HasNotionalCcy() bool`

HasNotionalCcy returns a boolean if a field has been set.

### GetNotionalUsd

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetNotionalUsd() string`

GetNotionalUsd returns the NotionalUsd field if non-nil, zero value otherwise.

### GetNotionalUsdOk

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetNotionalUsdOk() (*string, bool)`

GetNotionalUsdOk returns a tuple with the NotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsd

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) SetNotionalUsd(v string)`

SetNotionalUsd sets NotionalUsd field to given value.

### HasNotionalUsd

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) HasNotionalUsd() bool`

HasNotionalUsd returns a boolean if a field has been set.

### GetPos

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetPos() string`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetPosOk() (*string, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) SetPos(v string)`

SetPos sets Pos field to given value.

### HasPos

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) HasPos() bool`

HasPos returns a boolean if a field has been set.

### GetPosCcy

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetPosCcy() string`

GetPosCcy returns the PosCcy field if non-nil, zero value otherwise.

### GetPosCcyOk

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetPosCcyOk() (*string, bool)`

GetPosCcyOk returns a tuple with the PosCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosCcy

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) SetPosCcy(v string)`

SetPosCcy sets PosCcy field to given value.

### HasPosCcy

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) HasPosCcy() bool`

HasPosCcy returns a boolean if a field has been set.

### GetPosId

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetPosId() string`

GetPosId returns the PosId field if non-nil, zero value otherwise.

### GetPosIdOk

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetPosIdOk() (*string, bool)`

GetPosIdOk returns a tuple with the PosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosId

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) SetPosId(v string)`

SetPosId sets PosId field to given value.

### HasPosId

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) HasPosId() bool`

HasPosId returns a boolean if a field has been set.

### GetPosSide

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetQuoteBal

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetQuoteBal() string`

GetQuoteBal returns the QuoteBal field if non-nil, zero value otherwise.

### GetQuoteBalOk

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) GetQuoteBalOk() (*string, bool)`

GetQuoteBalOk returns a tuple with the QuoteBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteBal

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) SetQuoteBal(v string)`

SetQuoteBal sets QuoteBal field to given value.

### HasQuoteBal

`func (o *GetAccountAccountPositionRiskV5RespDataPosDataInner) HasQuoteBal() bool`

HasQuoteBal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


