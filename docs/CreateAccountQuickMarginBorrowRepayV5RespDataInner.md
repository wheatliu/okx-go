# CreateAccountQuickMarginBorrowRepayV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | borrow/repay amount | [optional] [default to ""]
**Ccy** | Pointer to **string** | Loan currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT | [optional] [default to ""]
**Side** | Pointer to **string** | &#x60;borrow&#x60;  &#x60;repay&#x60; | [optional] [default to ""]

## Methods

### NewCreateAccountQuickMarginBorrowRepayV5RespDataInner

`func NewCreateAccountQuickMarginBorrowRepayV5RespDataInner() *CreateAccountQuickMarginBorrowRepayV5RespDataInner`

NewCreateAccountQuickMarginBorrowRepayV5RespDataInner instantiates a new CreateAccountQuickMarginBorrowRepayV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountQuickMarginBorrowRepayV5RespDataInnerWithDefaults

`func NewCreateAccountQuickMarginBorrowRepayV5RespDataInnerWithDefaults() *CreateAccountQuickMarginBorrowRepayV5RespDataInner`

NewCreateAccountQuickMarginBorrowRepayV5RespDataInnerWithDefaults instantiates a new CreateAccountQuickMarginBorrowRepayV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCcy

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInstId

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetSide

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CreateAccountQuickMarginBorrowRepayV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


