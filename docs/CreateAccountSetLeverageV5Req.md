# CreateAccountSetLeverageV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Currency used for margin, used for the leverage setting for the currency in auto borrow.  Only applicable to &#x60;cross&#x60; &#x60;MARGIN&#x60; of &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60;  Required when setting the leverage for automatically borrowing coin. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID  Under cross mode, either &#x60;instId&#x60; or &#x60;ccy&#x60; is required; if both are passed, &#x60;instId&#x60; will be used by default. | [optional] [default to ""]
**Lever** | **string** | Leverage | [default to ""]
**MgnMode** | **string** | Margin mode  &#x60;isolated&#x60; &#x60;cross&#x60;   Can only be &#x60;cross&#x60; if &#x60;ccy&#x60; is passed. | [default to ""]
**PosSide** | Pointer to **string** | Position side  &#x60;long&#x60; &#x60;short&#x60;  Only required when margin mode is &#x60;isolated&#x60; in &#x60;long/short&#x60; mode for &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;. | [optional] [default to ""]

## Methods

### NewCreateAccountSetLeverageV5Req

`func NewCreateAccountSetLeverageV5Req(lever string, mgnMode string, ) *CreateAccountSetLeverageV5Req`

NewCreateAccountSetLeverageV5Req instantiates a new CreateAccountSetLeverageV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountSetLeverageV5ReqWithDefaults

`func NewCreateAccountSetLeverageV5ReqWithDefaults() *CreateAccountSetLeverageV5Req`

NewCreateAccountSetLeverageV5ReqWithDefaults instantiates a new CreateAccountSetLeverageV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *CreateAccountSetLeverageV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateAccountSetLeverageV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateAccountSetLeverageV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateAccountSetLeverageV5Req) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *CreateAccountSetLeverageV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateAccountSetLeverageV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateAccountSetLeverageV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateAccountSetLeverageV5Req) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLever

`func (o *CreateAccountSetLeverageV5Req) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *CreateAccountSetLeverageV5Req) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *CreateAccountSetLeverageV5Req) SetLever(v string)`

SetLever sets Lever field to given value.


### GetMgnMode

`func (o *CreateAccountSetLeverageV5Req) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *CreateAccountSetLeverageV5Req) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *CreateAccountSetLeverageV5Req) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.


### GetPosSide

`func (o *CreateAccountSetLeverageV5Req) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateAccountSetLeverageV5Req) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateAccountSetLeverageV5Req) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateAccountSetLeverageV5Req) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


