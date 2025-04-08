# CreateTradeOrderAlgoV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [optional] [default to ""]
**Ccy** | Pointer to **string** | Margin currency   Applicable to all &#x60;isolated&#x60; &#x60;MARGIN&#x60; orders and &#x60;cross&#x60; &#x60;MARGIN&#x60; orders in &#x60;Spot and futures mode&#x60;. | [optional] [default to ""]
**CloseFraction** | Pointer to **string** | Fraction of position to be closed when the algo order is triggered.   Currently the system supports fully closing the position only so the only accepted value is &#x60;1&#x60;. For the same position, only one TPSL pending order for fully closing the position is supported.   This is only applicable to &#x60;FUTURES&#x60; or &#x60;SWAP&#x60; instruments.  If &#x60;posSide&#x60; is &#x60;net&#x60;, &#x60;reduceOnly&#x60; must be &#x60;true&#x60;.  This is only applicable if &#x60;ordType&#x60; is &#x60;conditional&#x60; or &#x60;oco&#x60;.  This is only applicable if the stop loss and take profit order is executed as market order.  This is not supported in Portfolio Margin mode.  Either &#x60;sz&#x60; or &#x60;closeFraction&#x60; is required. | [optional] [default to ""]
**InstId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to ""]
**OrdType** | **string** | Order type    &#x60;conditional&#x60;: One-way stop order  &#x60;oco&#x60;: One-cancels-the-other order  &#x60;chase&#x60;: chase order, only applicable to FUTURES and SWAP  &#x60;trigger&#x60;: Trigger order  &#x60;move_order_stop&#x60;: Trailing order  &#x60;twap&#x60;: TWAP order | [default to ""]
**PosSide** | Pointer to **string** | Position side   Required in &#x60;long/short&#x60; mode and only be &#x60;long&#x60; or &#x60;short&#x60; | [optional] [default to ""]
**Side** | **string** | Order side, &#x60;buy&#x60; &#x60;sell&#x60; | [default to ""]
**Sz** | Pointer to **string** | Quantity to buy or sell  Either &#x60;sz&#x60; or &#x60;closeFraction&#x60; is required. | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag    A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters. | [optional] [default to ""]
**TdMode** | **string** | Trade mode  Margin mode &#x60;cross&#x60; &#x60;isolated&#x60;  Non-Margin mode &#x60;cash&#x60;  &#x60;spot_isolated&#x60; (only applicable to SPOT lead trading) | [default to ""]
**TgtCcy** | Pointer to **string** | Order quantity unit setting for &#x60;sz&#x60;  &#x60;base_ccy&#x60;: Base currency ,&#x60;quote_ccy&#x60;: Quote currency    Only applicable to &#x60;SPOT&#x60; traded with Market buy &#x60;conditional&#x60; order  Default is &#x60;quote_ccy&#x60; for buy, &#x60;base_ccy&#x60; for sell | [optional] [default to ""]

## Methods

### NewCreateTradeOrderAlgoV5Req

`func NewCreateTradeOrderAlgoV5Req(instId string, ordType string, side string, tdMode string, ) *CreateTradeOrderAlgoV5Req`

NewCreateTradeOrderAlgoV5Req instantiates a new CreateTradeOrderAlgoV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeOrderAlgoV5ReqWithDefaults

`func NewCreateTradeOrderAlgoV5ReqWithDefaults() *CreateTradeOrderAlgoV5Req`

NewCreateTradeOrderAlgoV5ReqWithDefaults instantiates a new CreateTradeOrderAlgoV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoClOrdId

`func (o *CreateTradeOrderAlgoV5Req) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *CreateTradeOrderAlgoV5Req) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *CreateTradeOrderAlgoV5Req) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *CreateTradeOrderAlgoV5Req) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetCcy

`func (o *CreateTradeOrderAlgoV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateTradeOrderAlgoV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateTradeOrderAlgoV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateTradeOrderAlgoV5Req) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetCloseFraction

`func (o *CreateTradeOrderAlgoV5Req) GetCloseFraction() string`

GetCloseFraction returns the CloseFraction field if non-nil, zero value otherwise.

### GetCloseFractionOk

`func (o *CreateTradeOrderAlgoV5Req) GetCloseFractionOk() (*string, bool)`

GetCloseFractionOk returns a tuple with the CloseFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseFraction

`func (o *CreateTradeOrderAlgoV5Req) SetCloseFraction(v string)`

SetCloseFraction sets CloseFraction field to given value.

### HasCloseFraction

`func (o *CreateTradeOrderAlgoV5Req) HasCloseFraction() bool`

HasCloseFraction returns a boolean if a field has been set.

### GetInstId

`func (o *CreateTradeOrderAlgoV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradeOrderAlgoV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradeOrderAlgoV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetOrdType

`func (o *CreateTradeOrderAlgoV5Req) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *CreateTradeOrderAlgoV5Req) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *CreateTradeOrderAlgoV5Req) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.


### GetPosSide

`func (o *CreateTradeOrderAlgoV5Req) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateTradeOrderAlgoV5Req) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateTradeOrderAlgoV5Req) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateTradeOrderAlgoV5Req) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetSide

`func (o *CreateTradeOrderAlgoV5Req) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateTradeOrderAlgoV5Req) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateTradeOrderAlgoV5Req) SetSide(v string)`

SetSide sets Side field to given value.


### GetSz

`func (o *CreateTradeOrderAlgoV5Req) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateTradeOrderAlgoV5Req) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateTradeOrderAlgoV5Req) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateTradeOrderAlgoV5Req) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *CreateTradeOrderAlgoV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateTradeOrderAlgoV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateTradeOrderAlgoV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateTradeOrderAlgoV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTdMode

`func (o *CreateTradeOrderAlgoV5Req) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *CreateTradeOrderAlgoV5Req) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *CreateTradeOrderAlgoV5Req) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.


### GetTgtCcy

`func (o *CreateTradeOrderAlgoV5Req) GetTgtCcy() string`

GetTgtCcy returns the TgtCcy field if non-nil, zero value otherwise.

### GetTgtCcyOk

`func (o *CreateTradeOrderAlgoV5Req) GetTgtCcyOk() (*string, bool)`

GetTgtCcyOk returns a tuple with the TgtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCcy

`func (o *CreateTradeOrderAlgoV5Req) SetTgtCcy(v string)`

SetTgtCcy sets TgtCcy field to given value.

### HasTgtCcy

`func (o *CreateTradeOrderAlgoV5Req) HasTgtCcy() bool`

HasTgtCcy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


