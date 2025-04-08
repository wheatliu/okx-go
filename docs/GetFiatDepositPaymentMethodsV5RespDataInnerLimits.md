# GetFiatDepositPaymentMethodsV5RespDataInnerLimits

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

### NewGetFiatDepositPaymentMethodsV5RespDataInnerLimits

`func NewGetFiatDepositPaymentMethodsV5RespDataInnerLimits() *GetFiatDepositPaymentMethodsV5RespDataInnerLimits`

NewGetFiatDepositPaymentMethodsV5RespDataInnerLimits instantiates a new GetFiatDepositPaymentMethodsV5RespDataInnerLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatDepositPaymentMethodsV5RespDataInnerLimitsWithDefaults

`func NewGetFiatDepositPaymentMethodsV5RespDataInnerLimitsWithDefaults() *GetFiatDepositPaymentMethodsV5RespDataInnerLimits`

NewGetFiatDepositPaymentMethodsV5RespDataInnerLimitsWithDefaults instantiates a new GetFiatDepositPaymentMethodsV5RespDataInnerLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetDailyLimit() string`

GetDailyLimit returns the DailyLimit field if non-nil, zero value otherwise.

### GetDailyLimitOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetDailyLimitOk() (*string, bool)`

GetDailyLimitOk returns a tuple with the DailyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) SetDailyLimit(v string)`

SetDailyLimit sets DailyLimit field to given value.

### HasDailyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) HasDailyLimit() bool`

HasDailyLimit returns a boolean if a field has been set.

### GetDailyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetDailyLimitRemaining() string`

GetDailyLimitRemaining returns the DailyLimitRemaining field if non-nil, zero value otherwise.

### GetDailyLimitRemainingOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetDailyLimitRemainingOk() (*string, bool)`

GetDailyLimitRemainingOk returns a tuple with the DailyLimitRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) SetDailyLimitRemaining(v string)`

SetDailyLimitRemaining sets DailyLimitRemaining field to given value.

### HasDailyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) HasDailyLimitRemaining() bool`

HasDailyLimitRemaining returns a boolean if a field has been set.

### GetLifetimeLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetLifetimeLimit() string`

GetLifetimeLimit returns the LifetimeLimit field if non-nil, zero value otherwise.

### GetLifetimeLimitOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetLifetimeLimitOk() (*string, bool)`

GetLifetimeLimitOk returns a tuple with the LifetimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) SetLifetimeLimit(v string)`

SetLifetimeLimit sets LifetimeLimit field to given value.

### HasLifetimeLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) HasLifetimeLimit() bool`

HasLifetimeLimit returns a boolean if a field has been set.

### GetMaxAmt

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetMaxAmt() string`

GetMaxAmt returns the MaxAmt field if non-nil, zero value otherwise.

### GetMaxAmtOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetMaxAmtOk() (*string, bool)`

GetMaxAmtOk returns a tuple with the MaxAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmt

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) SetMaxAmt(v string)`

SetMaxAmt sets MaxAmt field to given value.

### HasMaxAmt

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) HasMaxAmt() bool`

HasMaxAmt returns a boolean if a field has been set.

### GetMinAmt

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetMinAmt() string`

GetMinAmt returns the MinAmt field if non-nil, zero value otherwise.

### GetMinAmtOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetMinAmtOk() (*string, bool)`

GetMinAmtOk returns a tuple with the MinAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmt

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) SetMinAmt(v string)`

SetMinAmt sets MinAmt field to given value.

### HasMinAmt

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) HasMinAmt() bool`

HasMinAmt returns a boolean if a field has been set.

### GetMonthlyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetMonthlyLimit() string`

GetMonthlyLimit returns the MonthlyLimit field if non-nil, zero value otherwise.

### GetMonthlyLimitOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetMonthlyLimitOk() (*string, bool)`

GetMonthlyLimitOk returns a tuple with the MonthlyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) SetMonthlyLimit(v string)`

SetMonthlyLimit sets MonthlyLimit field to given value.

### HasMonthlyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) HasMonthlyLimit() bool`

HasMonthlyLimit returns a boolean if a field has been set.

### GetMonthlyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetMonthlyLimitRemaining() string`

GetMonthlyLimitRemaining returns the MonthlyLimitRemaining field if non-nil, zero value otherwise.

### GetMonthlyLimitRemainingOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetMonthlyLimitRemainingOk() (*string, bool)`

GetMonthlyLimitRemainingOk returns a tuple with the MonthlyLimitRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) SetMonthlyLimitRemaining(v string)`

SetMonthlyLimitRemaining sets MonthlyLimitRemaining field to given value.

### HasMonthlyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) HasMonthlyLimitRemaining() bool`

HasMonthlyLimitRemaining returns a boolean if a field has been set.

### GetWeeklyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetWeeklyLimit() string`

GetWeeklyLimit returns the WeeklyLimit field if non-nil, zero value otherwise.

### GetWeeklyLimitOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetWeeklyLimitOk() (*string, bool)`

GetWeeklyLimitOk returns a tuple with the WeeklyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) SetWeeklyLimit(v string)`

SetWeeklyLimit sets WeeklyLimit field to given value.

### HasWeeklyLimit

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) HasWeeklyLimit() bool`

HasWeeklyLimit returns a boolean if a field has been set.

### GetWeeklyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetWeeklyLimitRemaining() string`

GetWeeklyLimitRemaining returns the WeeklyLimitRemaining field if non-nil, zero value otherwise.

### GetWeeklyLimitRemainingOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) GetWeeklyLimitRemainingOk() (*string, bool)`

GetWeeklyLimitRemainingOk returns a tuple with the WeeklyLimitRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) SetWeeklyLimitRemaining(v string)`

SetWeeklyLimitRemaining sets WeeklyLimitRemaining field to given value.

### HasWeeklyLimitRemaining

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerLimits) HasWeeklyLimitRemaining() bool`

HasWeeklyLimitRemaining returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


