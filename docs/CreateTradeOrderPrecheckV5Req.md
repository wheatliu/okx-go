# CreateTradeOrderPrecheckV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachAlgoOrds** | Pointer to [**[]CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner**](CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner.md) | TP/SL information attached when placing order | [optional] 
**InstId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to ""]
**OrdType** | **string** | Order type   &#x60;market&#x60;: Market order   &#x60;limit&#x60;: Limit order   &#x60;post_only&#x60;: Post-only order   &#x60;fok&#x60;: Fill-or-kill order   &#x60;ioc&#x60;: Immediate-or-cancel order    &#x60;optimal_limit_ioc&#x60;: Market order with immediate-or-cancel order (applicable only to Expiry Futures and Perpetual Futures). | [default to ""]
**PosSide** | Pointer to **string** | Position side    The default is &#x60;net&#x60; in the &#x60;net&#x60; mode   It is required in the &#x60;long/short&#x60; mode, and can only be &#x60;long&#x60; or &#x60;short&#x60;.   Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;. | [optional] [default to ""]
**Px** | Pointer to **string** | Order price. Only applicable to &#x60;limit&#x60;,&#x60;post_only&#x60;,&#x60;fok&#x60;,&#x60;ioc&#x60;,&#x60;mmp&#x60;,&#x60;mmp_and_post_only&#x60; order. | [optional] [default to ""]
**ReduceOnly** | Pointer to **bool** | Whether orders can only reduce in position size.    Valid options: &#x60;true&#x60; or &#x60;false&#x60;. The default value is &#x60;false&#x60;.  Only applicable to &#x60;MARGIN&#x60; orders, and &#x60;FUTURES&#x60;/&#x60;SWAP&#x60; orders in &#x60;net&#x60; mode   Only applicable to &#x60;Spot and futures mode&#x60; and &#x60;Multi-currency margin&#x60; | [optional] 
**Side** | **string** | Order side, &#x60;buy&#x60; &#x60;sell&#x60; | [default to ""]
**Sz** | **string** | Quantity to buy or sell | [default to ""]
**TdMode** | **string** | Trade mode  Margin mode &#x60;cross&#x60; &#x60;isolated&#x60;  Non-Margin mode &#x60;cash&#x60;  &#x60;spot_isolated&#x60; (only applicable to SPOT lead trading, &#x60;tdMode&#x60; should be &#x60;spot_isolated&#x60; for &#x60;SPOT&#x60; lead trading.) | [default to ""]
**TgtCcy** | Pointer to **string** | Whether the target currency uses the quote or base currency.  &#x60;base_ccy&#x60;: Base currency ,&#x60;quote_ccy&#x60;: Quote currency    Only applicable to &#x60;SPOT&#x60; Market Orders  Default is &#x60;quote_ccy&#x60; for buy, &#x60;base_ccy&#x60; for sell | [optional] [default to ""]

## Methods

### NewCreateTradeOrderPrecheckV5Req

`func NewCreateTradeOrderPrecheckV5Req(instId string, ordType string, side string, sz string, tdMode string, ) *CreateTradeOrderPrecheckV5Req`

NewCreateTradeOrderPrecheckV5Req instantiates a new CreateTradeOrderPrecheckV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeOrderPrecheckV5ReqWithDefaults

`func NewCreateTradeOrderPrecheckV5ReqWithDefaults() *CreateTradeOrderPrecheckV5Req`

NewCreateTradeOrderPrecheckV5ReqWithDefaults instantiates a new CreateTradeOrderPrecheckV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachAlgoOrds

`func (o *CreateTradeOrderPrecheckV5Req) GetAttachAlgoOrds() []CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner`

GetAttachAlgoOrds returns the AttachAlgoOrds field if non-nil, zero value otherwise.

### GetAttachAlgoOrdsOk

`func (o *CreateTradeOrderPrecheckV5Req) GetAttachAlgoOrdsOk() (*[]CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner, bool)`

GetAttachAlgoOrdsOk returns a tuple with the AttachAlgoOrds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoOrds

`func (o *CreateTradeOrderPrecheckV5Req) SetAttachAlgoOrds(v []CreateTradeOrderPrecheckV5ReqAttachAlgoOrdsInner)`

SetAttachAlgoOrds sets AttachAlgoOrds field to given value.

### HasAttachAlgoOrds

`func (o *CreateTradeOrderPrecheckV5Req) HasAttachAlgoOrds() bool`

HasAttachAlgoOrds returns a boolean if a field has been set.

### GetInstId

`func (o *CreateTradeOrderPrecheckV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradeOrderPrecheckV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradeOrderPrecheckV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetOrdType

`func (o *CreateTradeOrderPrecheckV5Req) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *CreateTradeOrderPrecheckV5Req) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *CreateTradeOrderPrecheckV5Req) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.


### GetPosSide

`func (o *CreateTradeOrderPrecheckV5Req) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateTradeOrderPrecheckV5Req) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateTradeOrderPrecheckV5Req) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateTradeOrderPrecheckV5Req) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetPx

`func (o *CreateTradeOrderPrecheckV5Req) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *CreateTradeOrderPrecheckV5Req) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *CreateTradeOrderPrecheckV5Req) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *CreateTradeOrderPrecheckV5Req) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetReduceOnly

`func (o *CreateTradeOrderPrecheckV5Req) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *CreateTradeOrderPrecheckV5Req) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *CreateTradeOrderPrecheckV5Req) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *CreateTradeOrderPrecheckV5Req) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *CreateTradeOrderPrecheckV5Req) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateTradeOrderPrecheckV5Req) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateTradeOrderPrecheckV5Req) SetSide(v string)`

SetSide sets Side field to given value.


### GetSz

`func (o *CreateTradeOrderPrecheckV5Req) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateTradeOrderPrecheckV5Req) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateTradeOrderPrecheckV5Req) SetSz(v string)`

SetSz sets Sz field to given value.


### GetTdMode

`func (o *CreateTradeOrderPrecheckV5Req) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *CreateTradeOrderPrecheckV5Req) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *CreateTradeOrderPrecheckV5Req) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.


### GetTgtCcy

`func (o *CreateTradeOrderPrecheckV5Req) GetTgtCcy() string`

GetTgtCcy returns the TgtCcy field if non-nil, zero value otherwise.

### GetTgtCcyOk

`func (o *CreateTradeOrderPrecheckV5Req) GetTgtCcyOk() (*string, bool)`

GetTgtCcyOk returns a tuple with the TgtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCcy

`func (o *CreateTradeOrderPrecheckV5Req) SetTgtCcy(v string)`

SetTgtCcy sets TgtCcy field to given value.

### HasTgtCcy

`func (o *CreateTradeOrderPrecheckV5Req) HasTgtCcy() bool`

HasTgtCcy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


