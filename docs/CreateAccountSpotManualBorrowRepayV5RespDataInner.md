# CreateAccountSpotManualBorrowRepayV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | Actual amount | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**Side** | Pointer to **string** | Side  &#x60;borrow&#x60;  &#x60;repay&#x60; | [optional] [default to ""]

## Methods

### NewCreateAccountSpotManualBorrowRepayV5RespDataInner

`func NewCreateAccountSpotManualBorrowRepayV5RespDataInner() *CreateAccountSpotManualBorrowRepayV5RespDataInner`

NewCreateAccountSpotManualBorrowRepayV5RespDataInner instantiates a new CreateAccountSpotManualBorrowRepayV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountSpotManualBorrowRepayV5RespDataInnerWithDefaults

`func NewCreateAccountSpotManualBorrowRepayV5RespDataInnerWithDefaults() *CreateAccountSpotManualBorrowRepayV5RespDataInner`

NewCreateAccountSpotManualBorrowRepayV5RespDataInnerWithDefaults instantiates a new CreateAccountSpotManualBorrowRepayV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAccountSpotManualBorrowRepayV5RespDataInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAccountSpotManualBorrowRepayV5RespDataInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAccountSpotManualBorrowRepayV5RespDataInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *CreateAccountSpotManualBorrowRepayV5RespDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCcy

`func (o *CreateAccountSpotManualBorrowRepayV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateAccountSpotManualBorrowRepayV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateAccountSpotManualBorrowRepayV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateAccountSpotManualBorrowRepayV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetSide

`func (o *CreateAccountSpotManualBorrowRepayV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateAccountSpotManualBorrowRepayV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateAccountSpotManualBorrowRepayV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CreateAccountSpotManualBorrowRepayV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


