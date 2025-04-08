# CreateTradeOrderV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachAlgoOrds** | Pointer to [**[]CreateTradeOrderV5ReqAttachAlgoOrdsInner**](CreateTradeOrderV5ReqAttachAlgoOrdsInner.md) | TP/SL information attached when placing order | [optional] 
**BanAmend** | Pointer to **bool** | Whether to disallow the system from amending the size of the SPOT Market Order.  Valid options: &#x60;true&#x60; or &#x60;false&#x60;. The default value is &#x60;false&#x60;.  If &#x60;true&#x60;, system will not amend and reject the market order if user does not have sufficient funds.   Only applicable to SPOT Market Orders | [optional] 
**Ccy** | Pointer to **string** | Margin currency   Applicable to all &#x60;isolated&#x60; &#x60;MARGIN&#x60; orders and &#x60;cross&#x60; &#x60;MARGIN&#x60; orders in &#x60;Spot and futures mode&#x60;. | [optional] [default to ""]
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client    A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.  Only applicable to general order. It will not be posted to algoId when placing TP/SL order after the general order is filled completely. | [optional] [default to ""]
**InstId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to ""]
**OrdType** | **string** | Order type   &#x60;market&#x60;: Market order   &#x60;limit&#x60;: Limit order   &#x60;post_only&#x60;: Post-only order   &#x60;fok&#x60;: Fill-or-kill order   &#x60;ioc&#x60;: Immediate-or-cancel order    &#x60;optimal_limit_ioc&#x60;: Market order with immediate-or-cancel order (applicable only to Expiry Futures and Perpetual Futures).  &#x60;mmp&#x60;: Market Maker Protection (only applicable to Option in Portfolio Margin mode)   &#x60;mmp_and_post_only&#x60;: Market Maker Protection and Post-only order(only applicable to Option in Portfolio Margin mode) | [default to ""]
**PosSide** | Pointer to **string** | Position side    The default is &#x60;net&#x60; in the &#x60;net&#x60; mode   It is required in the &#x60;long/short&#x60; mode, and can only be &#x60;long&#x60; or &#x60;short&#x60;.   Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;. | [optional] [default to ""]
**Px** | Pointer to **string** | Order price. Only applicable to &#x60;limit&#x60;,&#x60;post_only&#x60;,&#x60;fok&#x60;,&#x60;ioc&#x60;,&#x60;mmp&#x60;,&#x60;mmp_and_post_only&#x60; order.  When placing an option order, one of px/pxUsd/pxVol must be filled in, and only one can be filled in | [optional] [default to ""]
**PxUsd** | Pointer to **string** | Place options orders in &#x60;USD&#x60;   Only applicable to options   When placing an option order, one of px/pxUsd/pxVol must be filled in, and only one can be filled in | [optional] [default to ""]
**PxVol** | Pointer to **string** | Place options orders based on implied volatility, where 1 represents 100%   Only applicable to options   When placing an option order, one of px/pxUsd/pxVol must be filled in, and only one can be filled in | [optional] [default to ""]
**QuickMgnType** | Pointer to **string** | Quick Margin type. Only applicable to Quick Margin Mode of isolated margin   &#x60;manual&#x60;, &#x60;auto_borrow&#x60;, &#x60;auto_repay&#x60;  The default value is &#x60;manual&#x60;(Deprecated) | [optional] [default to ""]
**ReduceOnly** | Pointer to **bool** | Whether orders can only reduce in position size.    Valid options: &#x60;true&#x60; or &#x60;false&#x60;. The default value is &#x60;false&#x60;.  Only applicable to &#x60;MARGIN&#x60; orders, and &#x60;FUTURES&#x60;/&#x60;SWAP&#x60; orders in &#x60;net&#x60; mode   Only applicable to &#x60;Spot and futures mode&#x60; and &#x60;Multi-currency margin&#x60; | [optional] 
**Side** | **string** | Order side, &#x60;buy&#x60; &#x60;sell&#x60; | [default to ""]
**StpId** | Pointer to **string** | Self trade prevention ID. Orders from the same master account with the same ID will be prevented from self trade.  Numerical integers defined by user in the range of 1&lt;&#x3D; x&lt;&#x3D; 999999999 (deprecated) | [optional] [default to ""]
**StpMode** | Pointer to **string** | Self trade prevention mode.   Default to cancel maker   &#x60;cancel_maker&#x60;,&#x60;cancel_taker&#x60;, &#x60;cancel_both&#x60;  Cancel both does not support FOK. | [optional] [default to ""]
**Sz** | **string** | Quantity to buy or sell | [default to ""]
**Tag** | Pointer to **string** | Order tag    A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters. | [optional] [default to ""]
**TdMode** | **string** | Trade mode  Margin mode &#x60;cross&#x60; &#x60;isolated&#x60;  Non-Margin mode &#x60;cash&#x60;  &#x60;spot_isolated&#x60; (only applicable to SPOT lead trading, &#x60;tdMode&#x60; should be &#x60;spot_isolated&#x60; for &#x60;SPOT&#x60; lead trading.) | [default to ""]
**TgtCcy** | Pointer to **string** | Whether the target currency uses the quote or base currency.  &#x60;base_ccy&#x60;: Base currency ,&#x60;quote_ccy&#x60;: Quote currency    Only applicable to &#x60;SPOT&#x60; Market Orders  Default is &#x60;quote_ccy&#x60; for buy, &#x60;base_ccy&#x60; for sell | [optional] [default to ""]

## Methods

### NewCreateTradeOrderV5Req

`func NewCreateTradeOrderV5Req(instId string, ordType string, side string, sz string, tdMode string, ) *CreateTradeOrderV5Req`

NewCreateTradeOrderV5Req instantiates a new CreateTradeOrderV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeOrderV5ReqWithDefaults

`func NewCreateTradeOrderV5ReqWithDefaults() *CreateTradeOrderV5Req`

NewCreateTradeOrderV5ReqWithDefaults instantiates a new CreateTradeOrderV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachAlgoOrds

`func (o *CreateTradeOrderV5Req) GetAttachAlgoOrds() []CreateTradeOrderV5ReqAttachAlgoOrdsInner`

GetAttachAlgoOrds returns the AttachAlgoOrds field if non-nil, zero value otherwise.

### GetAttachAlgoOrdsOk

`func (o *CreateTradeOrderV5Req) GetAttachAlgoOrdsOk() (*[]CreateTradeOrderV5ReqAttachAlgoOrdsInner, bool)`

GetAttachAlgoOrdsOk returns a tuple with the AttachAlgoOrds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoOrds

`func (o *CreateTradeOrderV5Req) SetAttachAlgoOrds(v []CreateTradeOrderV5ReqAttachAlgoOrdsInner)`

SetAttachAlgoOrds sets AttachAlgoOrds field to given value.

### HasAttachAlgoOrds

`func (o *CreateTradeOrderV5Req) HasAttachAlgoOrds() bool`

HasAttachAlgoOrds returns a boolean if a field has been set.

### GetBanAmend

`func (o *CreateTradeOrderV5Req) GetBanAmend() bool`

GetBanAmend returns the BanAmend field if non-nil, zero value otherwise.

### GetBanAmendOk

`func (o *CreateTradeOrderV5Req) GetBanAmendOk() (*bool, bool)`

GetBanAmendOk returns a tuple with the BanAmend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanAmend

`func (o *CreateTradeOrderV5Req) SetBanAmend(v bool)`

SetBanAmend sets BanAmend field to given value.

### HasBanAmend

`func (o *CreateTradeOrderV5Req) HasBanAmend() bool`

HasBanAmend returns a boolean if a field has been set.

### GetCcy

`func (o *CreateTradeOrderV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateTradeOrderV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateTradeOrderV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateTradeOrderV5Req) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClOrdId

`func (o *CreateTradeOrderV5Req) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateTradeOrderV5Req) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateTradeOrderV5Req) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateTradeOrderV5Req) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetInstId

`func (o *CreateTradeOrderV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradeOrderV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradeOrderV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetOrdType

`func (o *CreateTradeOrderV5Req) GetOrdType() string`

GetOrdType returns the OrdType field if non-nil, zero value otherwise.

### GetOrdTypeOk

`func (o *CreateTradeOrderV5Req) GetOrdTypeOk() (*string, bool)`

GetOrdTypeOk returns a tuple with the OrdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdType

`func (o *CreateTradeOrderV5Req) SetOrdType(v string)`

SetOrdType sets OrdType field to given value.


### GetPosSide

`func (o *CreateTradeOrderV5Req) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateTradeOrderV5Req) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateTradeOrderV5Req) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateTradeOrderV5Req) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetPx

`func (o *CreateTradeOrderV5Req) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *CreateTradeOrderV5Req) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *CreateTradeOrderV5Req) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *CreateTradeOrderV5Req) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetPxUsd

`func (o *CreateTradeOrderV5Req) GetPxUsd() string`

GetPxUsd returns the PxUsd field if non-nil, zero value otherwise.

### GetPxUsdOk

`func (o *CreateTradeOrderV5Req) GetPxUsdOk() (*string, bool)`

GetPxUsdOk returns a tuple with the PxUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxUsd

`func (o *CreateTradeOrderV5Req) SetPxUsd(v string)`

SetPxUsd sets PxUsd field to given value.

### HasPxUsd

`func (o *CreateTradeOrderV5Req) HasPxUsd() bool`

HasPxUsd returns a boolean if a field has been set.

### GetPxVol

`func (o *CreateTradeOrderV5Req) GetPxVol() string`

GetPxVol returns the PxVol field if non-nil, zero value otherwise.

### GetPxVolOk

`func (o *CreateTradeOrderV5Req) GetPxVolOk() (*string, bool)`

GetPxVolOk returns a tuple with the PxVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxVol

`func (o *CreateTradeOrderV5Req) SetPxVol(v string)`

SetPxVol sets PxVol field to given value.

### HasPxVol

`func (o *CreateTradeOrderV5Req) HasPxVol() bool`

HasPxVol returns a boolean if a field has been set.

### GetQuickMgnType

`func (o *CreateTradeOrderV5Req) GetQuickMgnType() string`

GetQuickMgnType returns the QuickMgnType field if non-nil, zero value otherwise.

### GetQuickMgnTypeOk

`func (o *CreateTradeOrderV5Req) GetQuickMgnTypeOk() (*string, bool)`

GetQuickMgnTypeOk returns a tuple with the QuickMgnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickMgnType

`func (o *CreateTradeOrderV5Req) SetQuickMgnType(v string)`

SetQuickMgnType sets QuickMgnType field to given value.

### HasQuickMgnType

`func (o *CreateTradeOrderV5Req) HasQuickMgnType() bool`

HasQuickMgnType returns a boolean if a field has been set.

### GetReduceOnly

`func (o *CreateTradeOrderV5Req) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *CreateTradeOrderV5Req) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *CreateTradeOrderV5Req) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *CreateTradeOrderV5Req) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetSide

`func (o *CreateTradeOrderV5Req) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateTradeOrderV5Req) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateTradeOrderV5Req) SetSide(v string)`

SetSide sets Side field to given value.


### GetStpId

`func (o *CreateTradeOrderV5Req) GetStpId() string`

GetStpId returns the StpId field if non-nil, zero value otherwise.

### GetStpIdOk

`func (o *CreateTradeOrderV5Req) GetStpIdOk() (*string, bool)`

GetStpIdOk returns a tuple with the StpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpId

`func (o *CreateTradeOrderV5Req) SetStpId(v string)`

SetStpId sets StpId field to given value.

### HasStpId

`func (o *CreateTradeOrderV5Req) HasStpId() bool`

HasStpId returns a boolean if a field has been set.

### GetStpMode

`func (o *CreateTradeOrderV5Req) GetStpMode() string`

GetStpMode returns the StpMode field if non-nil, zero value otherwise.

### GetStpModeOk

`func (o *CreateTradeOrderV5Req) GetStpModeOk() (*string, bool)`

GetStpModeOk returns a tuple with the StpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpMode

`func (o *CreateTradeOrderV5Req) SetStpMode(v string)`

SetStpMode sets StpMode field to given value.

### HasStpMode

`func (o *CreateTradeOrderV5Req) HasStpMode() bool`

HasStpMode returns a boolean if a field has been set.

### GetSz

`func (o *CreateTradeOrderV5Req) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateTradeOrderV5Req) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateTradeOrderV5Req) SetSz(v string)`

SetSz sets Sz field to given value.


### GetTag

`func (o *CreateTradeOrderV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateTradeOrderV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateTradeOrderV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateTradeOrderV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTdMode

`func (o *CreateTradeOrderV5Req) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *CreateTradeOrderV5Req) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *CreateTradeOrderV5Req) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.


### GetTgtCcy

`func (o *CreateTradeOrderV5Req) GetTgtCcy() string`

GetTgtCcy returns the TgtCcy field if non-nil, zero value otherwise.

### GetTgtCcyOk

`func (o *CreateTradeOrderV5Req) GetTgtCcyOk() (*string, bool)`

GetTgtCcyOk returns a tuple with the TgtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCcy

`func (o *CreateTradeOrderV5Req) SetTgtCcy(v string)`

SetTgtCcy sets TgtCcy field to given value.

### HasTgtCcy

`func (o *CreateTradeOrderV5Req) HasTgtCcy() bool`

HasTgtCcy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


