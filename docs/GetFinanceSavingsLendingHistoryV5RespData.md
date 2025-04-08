# GetFinanceSavingsLendingHistoryV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | Lending amount | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**Earnings** | Pointer to **string** | Currency earnings | [optional] [default to ""]
**Rate** | Pointer to **string** | Lending annual interest rate | [optional] [default to ""]
**Ts** | Pointer to **string** | Lending time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetFinanceSavingsLendingHistoryV5RespData

`func NewGetFinanceSavingsLendingHistoryV5RespData() *GetFinanceSavingsLendingHistoryV5RespData`

NewGetFinanceSavingsLendingHistoryV5RespData instantiates a new GetFinanceSavingsLendingHistoryV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFinanceSavingsLendingHistoryV5RespDataWithDefaults

`func NewGetFinanceSavingsLendingHistoryV5RespDataWithDefaults() *GetFinanceSavingsLendingHistoryV5RespData`

NewGetFinanceSavingsLendingHistoryV5RespDataWithDefaults instantiates a new GetFinanceSavingsLendingHistoryV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *GetFinanceSavingsLendingHistoryV5RespData) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetFinanceSavingsLendingHistoryV5RespData) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetFinanceSavingsLendingHistoryV5RespData) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetFinanceSavingsLendingHistoryV5RespData) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCcy

`func (o *GetFinanceSavingsLendingHistoryV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFinanceSavingsLendingHistoryV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFinanceSavingsLendingHistoryV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFinanceSavingsLendingHistoryV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetEarnings

`func (o *GetFinanceSavingsLendingHistoryV5RespData) GetEarnings() string`

GetEarnings returns the Earnings field if non-nil, zero value otherwise.

### GetEarningsOk

`func (o *GetFinanceSavingsLendingHistoryV5RespData) GetEarningsOk() (*string, bool)`

GetEarningsOk returns a tuple with the Earnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnings

`func (o *GetFinanceSavingsLendingHistoryV5RespData) SetEarnings(v string)`

SetEarnings sets Earnings field to given value.

### HasEarnings

`func (o *GetFinanceSavingsLendingHistoryV5RespData) HasEarnings() bool`

HasEarnings returns a boolean if a field has been set.

### GetRate

`func (o *GetFinanceSavingsLendingHistoryV5RespData) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetFinanceSavingsLendingHistoryV5RespData) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetFinanceSavingsLendingHistoryV5RespData) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *GetFinanceSavingsLendingHistoryV5RespData) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetTs

`func (o *GetFinanceSavingsLendingHistoryV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetFinanceSavingsLendingHistoryV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetFinanceSavingsLendingHistoryV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetFinanceSavingsLendingHistoryV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


