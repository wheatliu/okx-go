# GetFiatDepositPaymentMethodsV5RespDataLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyLimit** | Pointer to **string** | The daily transaction limit | [optional] [default to ""]
**DailyLimitRemaining** | Pointer to **string** | The remaining daily transaction limit | [optional] [default to ""]
**LifetimeLimit** | Pointer to **string** | The lifetime transaction limit. Return the configured value, \&quot;\&quot; if not configured | [optional] [default to ""]
**MaxAmt** | Pointer to **string** | The maximum amount allowed per transaction | [optional] [default to ""]
**MinAmt** | Pointer to **string** | The minimum amount allowed per transaction | [optional] [default to ""]
**MonthlyLimit** | Pointer to **string** | The monthly transaction limit | [optional] [default to ""]
**MonthlyLimitRemaining** | Pointer to **string** | The remaining monthly transaction limit | [optional] [default to ""]
**WeeklyLimit** | Pointer to **string** | The weekly transaction limit | [optional] [default to ""]
**WeeklyLimitRemaining** | Pointer to **string** | The remaining weekly transaction limit | [optional] [default to ""]

## Methods

### NewGetFiatDepositPaymentMethodsV5RespDataLimits

`func NewGetFiatDepositPaymentMethodsV5RespDataLimits() *GetFiatDepositPaymentMethodsV5RespDataLimits`

NewGetFiatDepositPaymentMethodsV5RespDataLimits instantiates a new GetFiatDepositPaymentMethodsV5RespDataLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatDepositPaymentMethodsV5RespDataLimitsWithDefaults

`func NewGetFiatDepositPaymentMethodsV5RespDataLimitsWithDefaults() *GetFiatDepositPaymentMethodsV5RespDataLimits`

NewGetFiatDepositPaymentMethodsV5RespDataLimitsWithDefaults instantiates a new GetFiatDepositPaymentMethodsV5RespDataLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetDailyLimit() string`

GetDailyLimit returns the DailyLimit field if non-nil, zero value otherwise.

### GetDailyLimitOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetDailyLimitOk() (*string, bool)`

GetDailyLimitOk returns a tuple with the DailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) SetDailyLimit(v string)`

SetDailyLimit sets DailyLimit field to given value.

### HasDailyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) HasDailyLimit() bool`

HasDailyLimit returns a boolean if a field has been set.

### GetDailyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetDailyLimitRemaining() string`

GetDailyLimitRemaining returns the DailyLimitRemaining field if non-nil, zero value otherwise.

### GetDailyLimitRemainingOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetDailyLimitRemainingOk() (*string, bool)`

GetDailyLimitRemainingOk returns a tuple with the DailyLimitRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) SetDailyLimitRemaining(v string)`

SetDailyLimitRemaining sets DailyLimitRemaining field to given value.

### HasDailyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) HasDailyLimitRemaining() bool`

HasDailyLimitRemaining returns a boolean if a field has been set.

### GetLifetimeLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetLifetimeLimit() string`

GetLifetimeLimit returns the LifetimeLimit field if non-nil, zero value otherwise.

### GetLifetimeLimitOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetLifetimeLimitOk() (*string, bool)`

GetLifetimeLimitOk returns a tuple with the LifetimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) SetLifetimeLimit(v string)`

SetLifetimeLimit sets LifetimeLimit field to given value.

### HasLifetimeLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) HasLifetimeLimit() bool`

HasLifetimeLimit returns a boolean if a field has been set.

### GetMaxAmt

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetMaxAmt() string`

GetMaxAmt returns the MaxAmt field if non-nil, zero value otherwise.

### GetMaxAmtOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetMaxAmtOk() (*string, bool)`

GetMaxAmtOk returns a tuple with the MaxAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmt

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) SetMaxAmt(v string)`

SetMaxAmt sets MaxAmt field to given value.

### HasMaxAmt

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) HasMaxAmt() bool`

HasMaxAmt returns a boolean if a field has been set.

### GetMinAmt

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetMinAmt() string`

GetMinAmt returns the MinAmt field if non-nil, zero value otherwise.

### GetMinAmtOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetMinAmtOk() (*string, bool)`

GetMinAmtOk returns a tuple with the MinAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmt

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) SetMinAmt(v string)`

SetMinAmt sets MinAmt field to given value.

### HasMinAmt

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) HasMinAmt() bool`

HasMinAmt returns a boolean if a field has been set.

### GetMonthlyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetMonthlyLimit() string`

GetMonthlyLimit returns the MonthlyLimit field if non-nil, zero value otherwise.

### GetMonthlyLimitOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetMonthlyLimitOk() (*string, bool)`

GetMonthlyLimitOk returns a tuple with the MonthlyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) SetMonthlyLimit(v string)`

SetMonthlyLimit sets MonthlyLimit field to given value.

### HasMonthlyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) HasMonthlyLimit() bool`

HasMonthlyLimit returns a boolean if a field has been set.

### GetMonthlyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetMonthlyLimitRemaining() string`

GetMonthlyLimitRemaining returns the MonthlyLimitRemaining field if non-nil, zero value otherwise.

### GetMonthlyLimitRemainingOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetMonthlyLimitRemainingOk() (*string, bool)`

GetMonthlyLimitRemainingOk returns a tuple with the MonthlyLimitRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) SetMonthlyLimitRemaining(v string)`

SetMonthlyLimitRemaining sets MonthlyLimitRemaining field to given value.

### HasMonthlyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) HasMonthlyLimitRemaining() bool`

HasMonthlyLimitRemaining returns a boolean if a field has been set.

### GetWeeklyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetWeeklyLimit() string`

GetWeeklyLimit returns the WeeklyLimit field if non-nil, zero value otherwise.

### GetWeeklyLimitOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetWeeklyLimitOk() (*string, bool)`

GetWeeklyLimitOk returns a tuple with the WeeklyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) SetWeeklyLimit(v string)`

SetWeeklyLimit sets WeeklyLimit field to given value.

### HasWeeklyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) HasWeeklyLimit() bool`

HasWeeklyLimit returns a boolean if a field has been set.

### GetWeeklyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetWeeklyLimitRemaining() string`

GetWeeklyLimitRemaining returns the WeeklyLimitRemaining field if non-nil, zero value otherwise.

### GetWeeklyLimitRemainingOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) GetWeeklyLimitRemainingOk() (*string, bool)`

GetWeeklyLimitRemainingOk returns a tuple with the WeeklyLimitRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) SetWeeklyLimitRemaining(v string)`

SetWeeklyLimitRemaining sets WeeklyLimitRemaining field to given value.

### HasWeeklyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataLimits) HasWeeklyLimitRemaining() bool`

HasWeeklyLimitRemaining returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


