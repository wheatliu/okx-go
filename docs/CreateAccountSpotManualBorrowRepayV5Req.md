# CreateAccountSpotManualBorrowRepayV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | **string** | Amount | [default to ""]
**Ccy** | **string** | Currency, e.g. &#x60;BTC&#x60; | [default to ""]
**Side** | **string** | Side  &#x60;borrow&#x60;  &#x60;repay&#x60; | [default to ""]

## Methods

### NewCreateAccountSpotManualBorrowRepayV5Req

`func NewCreateAccountSpotManualBorrowRepayV5Req(amt string, ccy string, side string, ) *CreateAccountSpotManualBorrowRepayV5Req`

NewCreateAccountSpotManualBorrowRepayV5Req instantiates a new CreateAccountSpotManualBorrowRepayV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountSpotManualBorrowRepayV5ReqWithDefaults

`func NewCreateAccountSpotManualBorrowRepayV5ReqWithDefaults() *CreateAccountSpotManualBorrowRepayV5Req`

NewCreateAccountSpotManualBorrowRepayV5ReqWithDefaults instantiates a new CreateAccountSpotManualBorrowRepayV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAccountSpotManualBorrowRepayV5Req) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAccountSpotManualBorrowRepayV5Req) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAccountSpotManualBorrowRepayV5Req) SetAmt(v string)`

SetAmt sets Amt field to given value.


### GetCcy

`func (o *CreateAccountSpotManualBorrowRepayV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateAccountSpotManualBorrowRepayV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateAccountSpotManualBorrowRepayV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.


### GetSide

`func (o *CreateAccountSpotManualBorrowRepayV5Req) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateAccountSpotManualBorrowRepayV5Req) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateAccountSpotManualBorrowRepayV5Req) SetSide(v string)`

SetSide sets Side field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


