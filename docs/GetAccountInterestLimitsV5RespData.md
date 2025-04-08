# GetAccountInterestLimitsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Debt** | Pointer to **string** | Current debt in &#x60;USDT&#x60; | [optional] [default to ""]
**Interest** | Pointer to **string** | Current interest in &#x60;USDT&#x60;, the unit is &#x60;USDT&#x60;  Only applicable to &#x60;Market loans&#x60; | [optional] [default to ""]
**LoanAlloc** | Pointer to **string** | VIP Loan allocation for the current trading account  1. The unit is percent(%). Range is [0, 100]. Precision is 0.01%  2. If master account did not assign anything, then \&quot;0\&quot;  3. \&quot;\&quot; if shared between master and sub-account | [optional] [default to ""]
**NextDiscountTime** | Pointer to **string** | Next deduct time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**NextInterestTime** | Pointer to **string** | Next accrual time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Records** | Pointer to [**[]GetAccountInterestLimitsV5RespDataRecordsInner**](GetAccountInterestLimitsV5RespDataRecordsInner.md) | Details for currencies | [optional] 

## Methods

### NewGetAccountInterestLimitsV5RespData

`func NewGetAccountInterestLimitsV5RespData() *GetAccountInterestLimitsV5RespData`

NewGetAccountInterestLimitsV5RespData instantiates a new GetAccountInterestLimitsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountInterestLimitsV5RespDataWithDefaults

`func NewGetAccountInterestLimitsV5RespDataWithDefaults() *GetAccountInterestLimitsV5RespData`

NewGetAccountInterestLimitsV5RespDataWithDefaults instantiates a new GetAccountInterestLimitsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebt

`func (o *GetAccountInterestLimitsV5RespData) GetDebt() string`

GetDebt returns the Debt field if non-nil, zero value otherwise.

### GetDebtOk

`func (o *GetAccountInterestLimitsV5RespData) GetDebtOk() (*string, bool)`

GetDebtOk returns a tuple with the Debt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebt

`func (o *GetAccountInterestLimitsV5RespData) SetDebt(v string)`

SetDebt sets Debt field to given value.

### HasDebt

`func (o *GetAccountInterestLimitsV5RespData) HasDebt() bool`

HasDebt returns a boolean if a field has been set.

### GetInterest

`func (o *GetAccountInterestLimitsV5RespData) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *GetAccountInterestLimitsV5RespData) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *GetAccountInterestLimitsV5RespData) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *GetAccountInterestLimitsV5RespData) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetLoanAlloc

`func (o *GetAccountInterestLimitsV5RespData) GetLoanAlloc() string`

GetLoanAlloc returns the LoanAlloc field if non-nil, zero value otherwise.

### GetLoanAllocOk

`func (o *GetAccountInterestLimitsV5RespData) GetLoanAllocOk() (*string, bool)`

GetLoanAllocOk returns a tuple with the LoanAlloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanAlloc

`func (o *GetAccountInterestLimitsV5RespData) SetLoanAlloc(v string)`

SetLoanAlloc sets LoanAlloc field to given value.

### HasLoanAlloc

`func (o *GetAccountInterestLimitsV5RespData) HasLoanAlloc() bool`

HasLoanAlloc returns a boolean if a field has been set.

### GetNextDiscountTime

`func (o *GetAccountInterestLimitsV5RespData) GetNextDiscountTime() string`

GetNextDiscountTime returns the NextDiscountTime field if non-nil, zero value otherwise.

### GetNextDiscountTimeOk

`func (o *GetAccountInterestLimitsV5RespData) GetNextDiscountTimeOk() (*string, bool)`

GetNextDiscountTimeOk returns a tuple with the NextDiscountTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDiscountTime

`func (o *GetAccountInterestLimitsV5RespData) SetNextDiscountTime(v string)`

SetNextDiscountTime sets NextDiscountTime field to given value.

### HasNextDiscountTime

`func (o *GetAccountInterestLimitsV5RespData) HasNextDiscountTime() bool`

HasNextDiscountTime returns a boolean if a field has been set.

### GetNextInterestTime

`func (o *GetAccountInterestLimitsV5RespData) GetNextInterestTime() string`

GetNextInterestTime returns the NextInterestTime field if non-nil, zero value otherwise.

### GetNextInterestTimeOk

`func (o *GetAccountInterestLimitsV5RespData) GetNextInterestTimeOk() (*string, bool)`

GetNextInterestTimeOk returns a tuple with the NextInterestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInterestTime

`func (o *GetAccountInterestLimitsV5RespData) SetNextInterestTime(v string)`

SetNextInterestTime sets NextInterestTime field to given value.

### HasNextInterestTime

`func (o *GetAccountInterestLimitsV5RespData) HasNextInterestTime() bool`

HasNextInterestTime returns a boolean if a field has been set.

### GetRecords

`func (o *GetAccountInterestLimitsV5RespData) GetRecords() []GetAccountInterestLimitsV5RespDataRecordsInner`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *GetAccountInterestLimitsV5RespData) GetRecordsOk() (*[]GetAccountInterestLimitsV5RespDataRecordsInner, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *GetAccountInterestLimitsV5RespData) SetRecords(v []GetAccountInterestLimitsV5RespDataRecordsInner)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *GetAccountInterestLimitsV5RespData) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


