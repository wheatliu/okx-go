# CreateFinanceFlexibleLoanMaxLoanV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BorrowCcy** | Pointer to **string** | Currency to borrow, e.g. &#x60;USDT&#x60; | [optional] [default to ""]
**MaxLoan** | Pointer to **string** | Maximum available loan | [optional] [default to ""]
**NotionalUsd** | Pointer to **string** | Maximum available loan notional value, unit in &#x60;USD&#x60; | [optional] [default to ""]
**RemainingQuota** | Pointer to **string** | Remaining quota, unit in &#x60;borrowCcy&#x60; | [optional] [default to ""]

## Methods

### NewCreateFinanceFlexibleLoanMaxLoanV5RespData

`func NewCreateFinanceFlexibleLoanMaxLoanV5RespData() *CreateFinanceFlexibleLoanMaxLoanV5RespData`

NewCreateFinanceFlexibleLoanMaxLoanV5RespData instantiates a new CreateFinanceFlexibleLoanMaxLoanV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFinanceFlexibleLoanMaxLoanV5RespDataWithDefaults

`func NewCreateFinanceFlexibleLoanMaxLoanV5RespDataWithDefaults() *CreateFinanceFlexibleLoanMaxLoanV5RespData`

NewCreateFinanceFlexibleLoanMaxLoanV5RespDataWithDefaults instantiates a new CreateFinanceFlexibleLoanMaxLoanV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBorrowCcy

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) GetBorrowCcy() string`

GetBorrowCcy returns the BorrowCcy field if non-nil, zero value otherwise.

### GetBorrowCcyOk

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) GetBorrowCcyOk() (*string, bool)`

GetBorrowCcyOk returns a tuple with the BorrowCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowCcy

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) SetBorrowCcy(v string)`

SetBorrowCcy sets BorrowCcy field to given value.

### HasBorrowCcy

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) HasBorrowCcy() bool`

HasBorrowCcy returns a boolean if a field has been set.

### GetMaxLoan

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) GetMaxLoan() string`

GetMaxLoan returns the MaxLoan field if non-nil, zero value otherwise.

### GetMaxLoanOk

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) GetMaxLoanOk() (*string, bool)`

GetMaxLoanOk returns a tuple with the MaxLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoan

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) SetMaxLoan(v string)`

SetMaxLoan sets MaxLoan field to given value.

### HasMaxLoan

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) HasMaxLoan() bool`

HasMaxLoan returns a boolean if a field has been set.

### GetNotionalUsd

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) GetNotionalUsd() string`

GetNotionalUsd returns the NotionalUsd field if non-nil, zero value otherwise.

### GetNotionalUsdOk

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) GetNotionalUsdOk() (*string, bool)`

GetNotionalUsdOk returns a tuple with the NotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsd

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) SetNotionalUsd(v string)`

SetNotionalUsd sets NotionalUsd field to given value.

### HasNotionalUsd

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) HasNotionalUsd() bool`

HasNotionalUsd returns a boolean if a field has been set.

### GetRemainingQuota

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) GetRemainingQuota() string`

GetRemainingQuota returns the RemainingQuota field if non-nil, zero value otherwise.

### GetRemainingQuotaOk

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) GetRemainingQuotaOk() (*string, bool)`

GetRemainingQuotaOk returns a tuple with the RemainingQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuota

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) SetRemainingQuota(v string)`

SetRemainingQuota sets RemainingQuota field to given value.

### HasRemainingQuota

`func (o *CreateFinanceFlexibleLoanMaxLoanV5RespData) HasRemainingQuota() bool`

HasRemainingQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


