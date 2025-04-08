# CreateRfqCreateQuoteV5RespDataDataInnerLegsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Margin currency.   Only applicable to &#x60;cross&#x60; &#x60;MARGIN&#x60; orders in &#x60;Spot and futures mode&#x60;. The parameter will be ignored in other scenarios. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side.   The default is &#x60;net&#x60; in the net mode. If not specified, return \&quot;\&quot;, which is equivalent to net.   It can only be &#x60;long&#x60; or &#x60;short&#x60; in the long/short mode. If not specified, return \&quot;\&quot;, which corresponds to the direction that opens new positions for the trade (buy &#x3D;&gt; long, sell &#x3D;&gt; short).   Only applicable to FUTURES/SWAP. | [optional] [default to ""]
**Px** | Pointer to **string** | The price of the leg. | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of the leg. Valid values can be buy or sell. | [optional] [default to ""]
**Sz** | Pointer to **string** | Size of the leg in contracts or spot. | [optional] [default to ""]
**TdMode** | Pointer to **string** | Trade mode   Margin mode: &#x60;cross&#x60; &#x60;isolated&#x60;   Non-Margin mode: &#x60;cash&#x60;.   If not provided, tdMode will inherit default values set by the system shown below:   Spot and futures mode &amp; SPOT: &#x60;cash&#x60;   Buy options in Spot and futures mode and Multi-currency Margin: &#x60;isolated&#x60;   Other cases: &#x60;cross&#x60; | [optional] [default to ""]
**TgtCcy** | Pointer to **string** | Defines the unit of the “sz” attribute.   Only applicable to instType &#x3D; SPOT.   The valid enumerations are &#x60;base_ccy&#x60; and &#x60;quote_ccy&#x60;. When not specified this is equal to &#x60;base_ccy&#x60; by default. | [optional] [default to ""]

## Methods

### NewCreateRfqCreateQuoteV5RespDataDataInnerLegsInner

`func NewCreateRfqCreateQuoteV5RespDataDataInnerLegsInner() *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner`

NewCreateRfqCreateQuoteV5RespDataDataInnerLegsInner instantiates a new CreateRfqCreateQuoteV5RespDataDataInnerLegsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCreateQuoteV5RespDataDataInnerLegsInnerWithDefaults

`func NewCreateRfqCreateQuoteV5RespDataDataInnerLegsInnerWithDefaults() *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner`

NewCreateRfqCreateQuoteV5RespDataDataInnerLegsInnerWithDefaults instantiates a new CreateRfqCreateQuoteV5RespDataDataInnerLegsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetPosSide

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetPx

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSide

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTdMode

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.

### HasTdMode

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) HasTdMode() bool`

HasTdMode returns a boolean if a field has been set.

### GetTgtCcy

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetTgtCcy() string`

GetTgtCcy returns the TgtCcy field if non-nil, zero value otherwise.

### GetTgtCcyOk

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) GetTgtCcyOk() (*string, bool)`

GetTgtCcyOk returns a tuple with the TgtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCcy

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) SetTgtCcy(v string)`

SetTgtCcy sets TgtCcy field to given value.

### HasTgtCcy

`func (o *CreateRfqCreateQuoteV5RespDataDataInnerLegsInner) HasTgtCcy() bool`

HasTgtCcy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


