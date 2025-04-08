# CreateFinanceFlexibleLoanMaxLoanV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BorrowCcy** | **string** | Currency to borrow, e.g. &#x60;USDT&#x60; | [default to ""]
**SupCollateral** | Pointer to [**[]CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner**](CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner.md) | Supplementary collateral assets | [optional] 

## Methods

### NewCreateFinanceFlexibleLoanMaxLoanV5Req

`func NewCreateFinanceFlexibleLoanMaxLoanV5Req(borrowCcy string, ) *CreateFinanceFlexibleLoanMaxLoanV5Req`

NewCreateFinanceFlexibleLoanMaxLoanV5Req instantiates a new CreateFinanceFlexibleLoanMaxLoanV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFinanceFlexibleLoanMaxLoanV5ReqWithDefaults

`func NewCreateFinanceFlexibleLoanMaxLoanV5ReqWithDefaults() *CreateFinanceFlexibleLoanMaxLoanV5Req`

NewCreateFinanceFlexibleLoanMaxLoanV5ReqWithDefaults instantiates a new CreateFinanceFlexibleLoanMaxLoanV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBorrowCcy

`func (o *CreateFinanceFlexibleLoanMaxLoanV5Req) GetBorrowCcy() string`

GetBorrowCcy returns the BorrowCcy field if non-nil, zero value otherwise.

### GetBorrowCcyOk

`func (o *CreateFinanceFlexibleLoanMaxLoanV5Req) GetBorrowCcyOk() (*string, bool)`

GetBorrowCcyOk returns a tuple with the BorrowCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowCcy

`func (o *CreateFinanceFlexibleLoanMaxLoanV5Req) SetBorrowCcy(v string)`

SetBorrowCcy sets BorrowCcy field to given value.


### GetSupCollateral

`func (o *CreateFinanceFlexibleLoanMaxLoanV5Req) GetSupCollateral() []CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner`

GetSupCollateral returns the SupCollateral field if non-nil, zero value otherwise.

### GetSupCollateralOk

`func (o *CreateFinanceFlexibleLoanMaxLoanV5Req) GetSupCollateralOk() (*[]CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner, bool)`

GetSupCollateralOk returns a tuple with the SupCollateral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupCollateral

`func (o *CreateFinanceFlexibleLoanMaxLoanV5Req) SetSupCollateral(v []CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner)`

SetSupCollateral sets SupCollateral field to given value.

### HasSupCollateral

`func (o *CreateFinanceFlexibleLoanMaxLoanV5Req) HasSupCollateral() bool`

HasSupCollateral returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


