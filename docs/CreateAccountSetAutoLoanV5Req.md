# CreateAccountSetAutoLoanV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoLoan** | Pointer to **bool** | Whether to automatically make loans   Valid values are &#x60;true&#x60;, &#x60;false&#x60;   The default is &#x60;true&#x60; | [optional] 

## Methods

### NewCreateAccountSetAutoLoanV5Req

`func NewCreateAccountSetAutoLoanV5Req() *CreateAccountSetAutoLoanV5Req`

NewCreateAccountSetAutoLoanV5Req instantiates a new CreateAccountSetAutoLoanV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountSetAutoLoanV5ReqWithDefaults

`func NewCreateAccountSetAutoLoanV5ReqWithDefaults() *CreateAccountSetAutoLoanV5Req`

NewCreateAccountSetAutoLoanV5ReqWithDefaults instantiates a new CreateAccountSetAutoLoanV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoLoan

`func (o *CreateAccountSetAutoLoanV5Req) GetAutoLoan() bool`

GetAutoLoan returns the AutoLoan field if non-nil, zero value otherwise.

### GetAutoLoanOk

`func (o *CreateAccountSetAutoLoanV5Req) GetAutoLoanOk() (*bool, bool)`

GetAutoLoanOk returns a tuple with the AutoLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLoan

`func (o *CreateAccountSetAutoLoanV5Req) SetAutoLoan(v bool)`

SetAutoLoan sets AutoLoan field to given value.

### HasAutoLoan

`func (o *CreateAccountSetAutoLoanV5Req) HasAutoLoan() bool`

HasAutoLoan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


