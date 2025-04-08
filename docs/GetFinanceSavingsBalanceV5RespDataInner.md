# GetFinanceSavingsBalanceV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | Currency amount | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency | [optional] [default to ""]
**Earnings** | Pointer to **string** | Currency earnings | [optional] [default to ""]
**LoanAmt** | Pointer to **string** | Lending amount | [optional] [default to ""]
**PendingAmt** | Pointer to **string** | Pending amount | [optional] [default to ""]
**Rate** | Pointer to **string** | Lending rate | [optional] [default to ""]
**RedemptAmt** | Pointer to **string** | Redempting amount (Deprecated) | [optional] [default to ""]

## Methods

### NewGetFinanceSavingsBalanceV5RespDataInner

`func NewGetFinanceSavingsBalanceV5RespDataInner() *GetFinanceSavingsBalanceV5RespDataInner`

NewGetFinanceSavingsBalanceV5RespDataInner instantiates a new GetFinanceSavingsBalanceV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceSavingsBalanceV5RespDataInnerWithDefaults

`func NewGetFinanceSavingsBalanceV5RespDataInnerWithDefaults() *GetFinanceSavingsBalanceV5RespDataInner`

NewGetFinanceSavingsBalanceV5RespDataInnerWithDefaults instantiates a new GetFinanceSavingsBalanceV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetFinanceSavingsBalanceV5RespDataInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetFinanceSavingsBalanceV5RespDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCcy

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFinanceSavingsBalanceV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFinanceSavingsBalanceV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetEarnings

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetEarnings() string`

GetEarnings returns the Earnings field if non-nil, zero value otherwise.

### GetEarningsOk

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetEarningsOk() (*string, bool)`

GetEarningsOk returns a tuple with the Earnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnings

`func (o *GetFinanceSavingsBalanceV5RespDataInner) SetEarnings(v string)`

SetEarnings sets Earnings field to given value.

### HasEarnings

`func (o *GetFinanceSavingsBalanceV5RespDataInner) HasEarnings() bool`

HasEarnings returns a boolean if a field has been set.

### GetLoanAmt

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetLoanAmt() string`

GetLoanAmt returns the LoanAmt field if non-nil, zero value otherwise.

### GetLoanAmtOk

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetLoanAmtOk() (*string, bool)`

GetLoanAmtOk returns a tuple with the LoanAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanAmt

`func (o *GetFinanceSavingsBalanceV5RespDataInner) SetLoanAmt(v string)`

SetLoanAmt sets LoanAmt field to given value.

### HasLoanAmt

`func (o *GetFinanceSavingsBalanceV5RespDataInner) HasLoanAmt() bool`

HasLoanAmt returns a boolean if a field has been set.

### GetPendingAmt

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetPendingAmt() string`

GetPendingAmt returns the PendingAmt field if non-nil, zero value otherwise.

### GetPendingAmtOk

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetPendingAmtOk() (*string, bool)`

GetPendingAmtOk returns a tuple with the PendingAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAmt

`func (o *GetFinanceSavingsBalanceV5RespDataInner) SetPendingAmt(v string)`

SetPendingAmt sets PendingAmt field to given value.

### HasPendingAmt

`func (o *GetFinanceSavingsBalanceV5RespDataInner) HasPendingAmt() bool`

HasPendingAmt returns a boolean if a field has been set.

### GetRate

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetFinanceSavingsBalanceV5RespDataInner) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *GetFinanceSavingsBalanceV5RespDataInner) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRedemptAmt

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetRedemptAmt() string`

GetRedemptAmt returns the RedemptAmt field if non-nil, zero value otherwise.

### GetRedemptAmtOk

`func (o *GetFinanceSavingsBalanceV5RespDataInner) GetRedemptAmtOk() (*string, bool)`

GetRedemptAmtOk returns a tuple with the RedemptAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptAmt

`func (o *GetFinanceSavingsBalanceV5RespDataInner) SetRedemptAmt(v string)`

SetRedemptAmt sets RedemptAmt field to given value.

### HasRedemptAmt

`func (o *GetFinanceSavingsBalanceV5RespDataInner) HasRedemptAmt() bool`

HasRedemptAmt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


