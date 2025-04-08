# CreateRfqCreateRfqV5RespDataInnerLegsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Margin currency.   Only applicable to &#x60;cross&#x60; &#x60;MARGIN&#x60; orders in &#x60;Spot and futures mode&#x60;. The parameter will be ignored in other scenarios. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side.   The default is &#x60;net&#x60; in the net mode. If not specified, return \&quot;\&quot;, which is equivalent to net.   It can only be &#x60;long&#x60; or &#x60;short&#x60; in the long/short mode. If not specified, return \&quot;\&quot;, which corresponds to the direction that opens new positions for the trade (buy &#x3D;&gt; long, sell &#x3D;&gt; short).   Only applicable to FUTURES/SWAP. | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of the leg. Valid values can be buy or sell. | [optional] [default to ""]
**Sz** | Pointer to **string** | Size of the leg in contracts or spot. | [optional] [default to ""]
**TdMode** | Pointer to **string** | Trade mode   Margin mode: &#x60;cross&#x60; &#x60;isolated&#x60;   Non-Margin mode: &#x60;cash&#x60;.   If not provided, tdMode will inherit default values set by the system shown below:   Spot and futures mode &amp; SPOT: &#x60;cash&#x60;   Buy options in Spot and futures mode and Multi-currency Margin: &#x60;isolated&#x60;   Other cases: &#x60;cross&#x60; | [optional] [default to ""]
**TgtCcy** | Pointer to **string** | Defines the unit of the “sz” attribute.   Only applicable to instType &#x3D; SPOT.   The valid enumerations are &#x60;base_ccy&#x60; and &#x60;quote_ccy&#x60;. When not specified this is equal to &#x60;base_ccy&#x60; by default. | [optional] [default to ""]

## Methods

### NewCreateRfqCreateRfqV5RespDataInnerLegsInner

`func NewCreateRfqCreateRfqV5RespDataInnerLegsInner() *CreateRfqCreateRfqV5RespDataInnerLegsInner`

NewCreateRfqCreateRfqV5RespDataInnerLegsInner instantiates a new CreateRfqCreateRfqV5RespDataInnerLegsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCreateRfqV5RespDataInnerLegsInnerWithDefaults

`func NewCreateRfqCreateRfqV5RespDataInnerLegsInnerWithDefaults() *CreateRfqCreateRfqV5RespDataInnerLegsInner`

NewCreateRfqCreateRfqV5RespDataInnerLegsInnerWithDefaults instantiates a new CreateRfqCreateRfqV5RespDataInnerLegsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetPosSide

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetSide

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetSz

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTdMode

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.

### HasTdMode

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) HasTdMode() bool`

HasTdMode returns a boolean if a field has been set.

### GetTgtCcy

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetTgtCcy() string`

GetTgtCcy returns the TgtCcy field if non-nil, zero value otherwise.

### GetTgtCcyOk

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) GetTgtCcyOk() (*string, bool)`

GetTgtCcyOk returns a tuple with the TgtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCcy

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) SetTgtCcy(v string)`

SetTgtCcy sets TgtCcy field to given value.

### HasTgtCcy

`func (o *CreateRfqCreateRfqV5RespDataInnerLegsInner) HasTgtCcy() bool`

HasTgtCcy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


