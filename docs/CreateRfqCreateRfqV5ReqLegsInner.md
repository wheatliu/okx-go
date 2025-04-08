# CreateRfqCreateRfqV5ReqLegsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Margin currency.   Only applicable to &#x60;cross&#x60; &#x60;MARGIN&#x60; orders in &#x60;Spot and futures mode&#x60;. The parameter will be ignored in other scenarios. | [optional] [default to ""]
**InstId** | Pointer to **string** | The Instrument ID of each leg. Example : \&quot;BTC-USDT-SWAP\&quot; | [optional] [default to ""]
**LmtPx** | Pointer to **string** | Taker expected price for the RFQ     If provided, RFQ trade will be automatically executed if the price from the quote is better than or equal to the price specified until the RFQ is canceled or expired.   This field has to be provided for all legs to have the RFQ automatically executed, or leave empty for all legs, otherwise request will be rejected.   The auto execution side depends on the leg side of the RFQ.   For &#x60;SPOT/MARGIN/FUTURES/SWAP&#x60;, lmtPx will be in unit of the quote ccy.   For &#x60;OPTION&#x60;, lmtPx will be in unit of settle ccy.   The field will not be disclosed to counterparties. | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side.   The default is &#x60;net&#x60; in the net mode. It can only be &#x60;long&#x60; or &#x60;short&#x60; in the long/short mode.   If not specified, users in long/short mode always open new positions.   Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;. | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of each leg. Valid values can be &#x60;buy&#x60; or &#x60;sell&#x60;. | [optional] [default to ""]
**Sz** | Pointer to **string** | The size of each leg | [optional] [default to ""]
**TdMode** | Pointer to **string** | Trade mode   Margin mode: &#x60;cross&#x60; &#x60;isolated&#x60;   Non-Margin mode: &#x60;cash&#x60;.   If not provided, tdMode will inherit default values set by the system shown below:   Spot and futures mode &amp; SPOT: &#x60;cash&#x60;   Buy options in Spot and futures mode and Multi-currency Margin: &#x60;isolated&#x60;   Other cases: &#x60;cross&#x60; | [optional] [default to ""]
**TgtCcy** | Pointer to **string** | Defines the unit of the “sz” attribute.   Only applicable to instType &#x3D; SPOT.   The valid enumerations are &#x60;base_ccy&#x60; and &#x60;quote_ccy&#x60;. When not specified, this is equal to &#x60;base_ccy&#x60; by default. | [optional] [default to ""]

## Methods

### NewCreateRfqCreateRfqV5ReqLegsInner

`func NewCreateRfqCreateRfqV5ReqLegsInner() *CreateRfqCreateRfqV5ReqLegsInner`

NewCreateRfqCreateRfqV5ReqLegsInner instantiates a new CreateRfqCreateRfqV5ReqLegsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCreateRfqV5ReqLegsInnerWithDefaults

`func NewCreateRfqCreateRfqV5ReqLegsInnerWithDefaults() *CreateRfqCreateRfqV5ReqLegsInner`

NewCreateRfqCreateRfqV5ReqLegsInnerWithDefaults instantiates a new CreateRfqCreateRfqV5ReqLegsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateRfqCreateRfqV5ReqLegsInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateRfqCreateRfqV5ReqLegsInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateRfqCreateRfqV5ReqLegsInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateRfqCreateRfqV5ReqLegsInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLmtPx

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetLmtPx() string`

GetLmtPx returns the LmtPx field if non-nil, zero value otherwise.

### GetLmtPxOk

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetLmtPxOk() (*string, bool)`

GetLmtPxOk returns a tuple with the LmtPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmtPx

`func (o *CreateRfqCreateRfqV5ReqLegsInner) SetLmtPx(v string)`

SetLmtPx sets LmtPx field to given value.

### HasLmtPx

`func (o *CreateRfqCreateRfqV5ReqLegsInner) HasLmtPx() bool`

HasLmtPx returns a boolean if a field has been set.

### GetPosSide

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateRfqCreateRfqV5ReqLegsInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateRfqCreateRfqV5ReqLegsInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetSide

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateRfqCreateRfqV5ReqLegsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CreateRfqCreateRfqV5ReqLegsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateRfqCreateRfqV5ReqLegsInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateRfqCreateRfqV5ReqLegsInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTdMode

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *CreateRfqCreateRfqV5ReqLegsInner) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.

### HasTdMode

`func (o *CreateRfqCreateRfqV5ReqLegsInner) HasTdMode() bool`

HasTdMode returns a boolean if a field has been set.

### GetTgtCcy

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetTgtCcy() string`

GetTgtCcy returns the TgtCcy field if non-nil, zero value otherwise.

### GetTgtCcyOk

`func (o *CreateRfqCreateRfqV5ReqLegsInner) GetTgtCcyOk() (*string, bool)`

GetTgtCcyOk returns a tuple with the TgtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCcy

`func (o *CreateRfqCreateRfqV5ReqLegsInner) SetTgtCcy(v string)`

SetTgtCcy sets TgtCcy field to given value.

### HasTgtCcy

`func (o *CreateRfqCreateRfqV5ReqLegsInner) HasTgtCcy() bool`

HasTgtCcy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


