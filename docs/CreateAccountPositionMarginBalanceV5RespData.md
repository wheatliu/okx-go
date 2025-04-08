# CreateAccountPositionMarginBalanceV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | Amount to be increase or decrease | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**Leverage** | Pointer to **string** | Real leverage after the margin adjustment | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side, &#x60;long&#x60;  &#x60;short&#x60; | [optional] [default to ""]
**Type** | Pointer to **string** | &#x60;add&#x60;: add margin  &#x60;reduce&#x60;: reduce margin | [optional] [default to ""]

## Methods

### NewCreateAccountPositionMarginBalanceV5RespData

`func NewCreateAccountPositionMarginBalanceV5RespData() *CreateAccountPositionMarginBalanceV5RespData`

NewCreateAccountPositionMarginBalanceV5RespData instantiates a new CreateAccountPositionMarginBalanceV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountPositionMarginBalanceV5RespDataWithDefaults

`func NewCreateAccountPositionMarginBalanceV5RespDataWithDefaults() *CreateAccountPositionMarginBalanceV5RespData`

NewCreateAccountPositionMarginBalanceV5RespDataWithDefaults instantiates a new CreateAccountPositionMarginBalanceV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAccountPositionMarginBalanceV5RespData) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAccountPositionMarginBalanceV5RespData) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAccountPositionMarginBalanceV5RespData) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *CreateAccountPositionMarginBalanceV5RespData) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCcy

`func (o *CreateAccountPositionMarginBalanceV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateAccountPositionMarginBalanceV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateAccountPositionMarginBalanceV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateAccountPositionMarginBalanceV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *CreateAccountPositionMarginBalanceV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateAccountPositionMarginBalanceV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateAccountPositionMarginBalanceV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateAccountPositionMarginBalanceV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLeverage

`func (o *CreateAccountPositionMarginBalanceV5RespData) GetLeverage() string`

GetLeverage returns the Leverage field if non-nil, zero value otherwise.

### GetLeverageOk

`func (o *CreateAccountPositionMarginBalanceV5RespData) GetLeverageOk() (*string, bool)`

GetLeverageOk returns a tuple with the Leverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverage

`func (o *CreateAccountPositionMarginBalanceV5RespData) SetLeverage(v string)`

SetLeverage sets Leverage field to given value.

### HasLeverage

`func (o *CreateAccountPositionMarginBalanceV5RespData) HasLeverage() bool`

HasLeverage returns a boolean if a field has been set.

### GetPosSide

`func (o *CreateAccountPositionMarginBalanceV5RespData) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateAccountPositionMarginBalanceV5RespData) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateAccountPositionMarginBalanceV5RespData) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateAccountPositionMarginBalanceV5RespData) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetType

`func (o *CreateAccountPositionMarginBalanceV5RespData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAccountPositionMarginBalanceV5RespData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAccountPositionMarginBalanceV5RespData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateAccountPositionMarginBalanceV5RespData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


