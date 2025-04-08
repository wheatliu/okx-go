# CreateAccountQuickMarginBorrowRepayV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | borrow/repay amount | [optional] [default to ""]
**Ccy** | Pointer to **string** | Loan currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT | [optional] [default to ""]
**Side** | Pointer to **string** | &#x60;borrow&#x60;  &#x60;repay&#x60; | [optional] [default to ""]

## Methods

### NewCreateAccountQuickMarginBorrowRepayV5RespData

`func NewCreateAccountQuickMarginBorrowRepayV5RespData() *CreateAccountQuickMarginBorrowRepayV5RespData`

NewCreateAccountQuickMarginBorrowRepayV5RespData instantiates a new CreateAccountQuickMarginBorrowRepayV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountQuickMarginBorrowRepayV5RespDataWithDefaults

`func NewCreateAccountQuickMarginBorrowRepayV5RespDataWithDefaults() *CreateAccountQuickMarginBorrowRepayV5RespData`

NewCreateAccountQuickMarginBorrowRepayV5RespDataWithDefaults instantiates a new CreateAccountQuickMarginBorrowRepayV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCcy

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetSide

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CreateAccountQuickMarginBorrowRepayV5RespData) HasSide() bool`

HasSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


