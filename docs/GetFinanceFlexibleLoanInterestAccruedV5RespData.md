# GetFinanceFlexibleLoanInterestAccruedV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ccy** | Pointer to **string** | Loan currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**Interest** | Pointer to **string** | Interest | [optional] [default to ""]
**InterestRate** | Pointer to **string** | Hourly APY, e.g. &#x60;0.01&#x60; represents &#x60;1%&#x60; | [optional] [default to ""]
**Loan** | Pointer to **string** | Loan when calculated interest | [optional] [default to ""]
**RefId** | Pointer to **string** | Reference ID | [optional] [default to ""]
**Ts** | Pointer to **string** | Timestamp to calculated interest, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetFinanceFlexibleLoanInterestAccruedV5RespData

`func NewGetFinanceFlexibleLoanInterestAccruedV5RespData() *GetFinanceFlexibleLoanInterestAccruedV5RespData`

NewGetFinanceFlexibleLoanInterestAccruedV5RespData instantiates a new GetFinanceFlexibleLoanInterestAccruedV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceFlexibleLoanInterestAccruedV5RespDataWithDefaults

`func NewGetFinanceFlexibleLoanInterestAccruedV5RespDataWithDefaults() *GetFinanceFlexibleLoanInterestAccruedV5RespData`

NewGetFinanceFlexibleLoanInterestAccruedV5RespDataWithDefaults instantiates a new GetFinanceFlexibleLoanInterestAccruedV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcy

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetInterest

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetInterestRate

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) GetInterestRate() string`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) GetInterestRateOk() (*string, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) SetInterestRate(v string)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetLoan

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) GetLoan() string`

GetLoan returns the Loan field if non-nil, zero value otherwise.

### GetLoanOk

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) GetLoanOk() (*string, bool)`

GetLoanOk returns a tuple with the Loan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoan

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) SetLoan(v string)`

SetLoan sets Loan field to given value.

### HasLoan

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) HasLoan() bool`

HasLoan returns a boolean if a field has been set.

### GetRefId

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetTs

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetFinanceFlexibleLoanInterestAccruedV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


