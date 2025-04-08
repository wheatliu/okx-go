# CreateAccountQuickMarginBorrowRepayV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | **string** | borrow/repay amount | [default to ""]
**Ccy** | **string** | Loan currency, e.g. &#x60;BTC&#x60; | [default to ""]
**InstId** | **string** | Instrument ID, e.g. BTC-USDT | [default to ""]
**Side** | **string** | &#x60;borrow&#x60;  &#x60;repay&#x60; | [default to ""]

## Methods

### NewCreateAccountQuickMarginBorrowRepayV5Req

`func NewCreateAccountQuickMarginBorrowRepayV5Req(amt string, ccy string, instId string, side string, ) *CreateAccountQuickMarginBorrowRepayV5Req`

NewCreateAccountQuickMarginBorrowRepayV5Req instantiates a new CreateAccountQuickMarginBorrowRepayV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountQuickMarginBorrowRepayV5ReqWithDefaults

`func NewCreateAccountQuickMarginBorrowRepayV5ReqWithDefaults() *CreateAccountQuickMarginBorrowRepayV5Req`

NewCreateAccountQuickMarginBorrowRepayV5ReqWithDefaults instantiates a new CreateAccountQuickMarginBorrowRepayV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAccountQuickMarginBorrowRepayV5Req) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAccountQuickMarginBorrowRepayV5Req) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAccountQuickMarginBorrowRepayV5Req) SetAmt(v string)`

SetAmt sets Amt field to given value.


### GetCcy

`func (o *CreateAccountQuickMarginBorrowRepayV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateAccountQuickMarginBorrowRepayV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateAccountQuickMarginBorrowRepayV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.


### GetInstId

`func (o *CreateAccountQuickMarginBorrowRepayV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateAccountQuickMarginBorrowRepayV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateAccountQuickMarginBorrowRepayV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetSide

`func (o *CreateAccountQuickMarginBorrowRepayV5Req) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateAccountQuickMarginBorrowRepayV5Req) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateAccountQuickMarginBorrowRepayV5Req) SetSide(v string)`

SetSide sets Side field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


