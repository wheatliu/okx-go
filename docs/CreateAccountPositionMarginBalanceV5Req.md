# CreateAccountPositionMarginBalanceV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | **string** | Amount to be increased or decreased. | [default to ""]
**Ccy** | Pointer to **string** | Currency   Applicable to &#x60;isolated&#x60; &#x60;MARGIN&#x60; orders | [optional] [default to ""]
**InstId** | **string** | Instrument ID | [default to ""]
**PosSide** | **string** | Position side, the default is &#x60;net&#x60;  &#x60;long&#x60;   &#x60;short&#x60;   &#x60;net&#x60; | [default to ""]
**Type** | **string** | &#x60;add&#x60;: add margin   &#x60;reduce&#x60;: reduce margin | [default to ""]

## Methods

### NewCreateAccountPositionMarginBalanceV5Req

`func NewCreateAccountPositionMarginBalanceV5Req(amt string, instId string, posSide string, type_ string, ) *CreateAccountPositionMarginBalanceV5Req`

NewCreateAccountPositionMarginBalanceV5Req instantiates a new CreateAccountPositionMarginBalanceV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountPositionMarginBalanceV5ReqWithDefaults

`func NewCreateAccountPositionMarginBalanceV5ReqWithDefaults() *CreateAccountPositionMarginBalanceV5Req`

NewCreateAccountPositionMarginBalanceV5ReqWithDefaults instantiates a new CreateAccountPositionMarginBalanceV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAccountPositionMarginBalanceV5Req) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAccountPositionMarginBalanceV5Req) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAccountPositionMarginBalanceV5Req) SetAmt(v string)`

SetAmt sets Amt field to given value.


### GetCcy

`func (o *CreateAccountPositionMarginBalanceV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateAccountPositionMarginBalanceV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateAccountPositionMarginBalanceV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateAccountPositionMarginBalanceV5Req) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *CreateAccountPositionMarginBalanceV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateAccountPositionMarginBalanceV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateAccountPositionMarginBalanceV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetPosSide

`func (o *CreateAccountPositionMarginBalanceV5Req) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateAccountPositionMarginBalanceV5Req) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateAccountPositionMarginBalanceV5Req) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.


### GetType

`func (o *CreateAccountPositionMarginBalanceV5Req) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAccountPositionMarginBalanceV5Req) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAccountPositionMarginBalanceV5Req) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


