# GetFiatDepositPaymentMethodsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]GetFiatDepositPaymentMethodsV5RespDataAccountsInner**](GetFiatDepositPaymentMethodsV5RespDataAccountsInner.md) | An array containing information about payment accounts associated with the currency and method. | [optional] 
**Ccy** | Pointer to **string** | Fiat currency | [optional] [default to ""]
**FeeRate** | Pointer to **string** | The fee rate for each deposit, expressed as a percentage  e.g. &#x60;0.02&#x60; represents 2 percent fee for each transaction. | [optional] [default to ""]
**Limits** | Pointer to [**GetFiatDepositPaymentMethodsV5RespDataLimits**](GetFiatDepositPaymentMethodsV5RespDataLimits.md) |  | [optional] 
**MinFee** | Pointer to **string** | The minimum fee for each deposit | [optional] [default to ""]
**PaymentMethod** | Pointer to **string** | The payment method associated with the currency  &#x60;TR_BANKS&#x60;  &#x60;PIX&#x60;  &#x60;SEPA&#x60; | [optional] [default to ""]

## Methods

### NewGetFiatDepositPaymentMethodsV5RespData

`func NewGetFiatDepositPaymentMethodsV5RespData() *GetFiatDepositPaymentMethodsV5RespData`

NewGetFiatDepositPaymentMethodsV5RespData instantiates a new GetFiatDepositPaymentMethodsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatDepositPaymentMethodsV5RespDataWithDefaults

`func NewGetFiatDepositPaymentMethodsV5RespDataWithDefaults() *GetFiatDepositPaymentMethodsV5RespData`

NewGetFiatDepositPaymentMethodsV5RespDataWithDefaults instantiates a new GetFiatDepositPaymentMethodsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *GetFiatDepositPaymentMethodsV5RespData) GetAccounts() []GetFiatDepositPaymentMethodsV5RespDataAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetFiatDepositPaymentMethodsV5RespData) GetAccountsOk() (*[]GetFiatDepositPaymentMethodsV5RespDataAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetFiatDepositPaymentMethodsV5RespData) SetAccounts(v []GetFiatDepositPaymentMethodsV5RespDataAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *GetFiatDepositPaymentMethodsV5RespData) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetCcy

`func (o *GetFiatDepositPaymentMethodsV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFiatDepositPaymentMethodsV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFiatDepositPaymentMethodsV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFiatDepositPaymentMethodsV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetFeeRate

`func (o *GetFiatDepositPaymentMethodsV5RespData) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *GetFiatDepositPaymentMethodsV5RespData) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *GetFiatDepositPaymentMethodsV5RespData) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *GetFiatDepositPaymentMethodsV5RespData) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetLimits

`func (o *GetFiatDepositPaymentMethodsV5RespData) GetLimits() GetFiatDepositPaymentMethodsV5RespDataLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *GetFiatDepositPaymentMethodsV5RespData) GetLimitsOk() (*GetFiatDepositPaymentMethodsV5RespDataLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *GetFiatDepositPaymentMethodsV5RespData) SetLimits(v GetFiatDepositPaymentMethodsV5RespDataLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *GetFiatDepositPaymentMethodsV5RespData) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetMinFee

`func (o *GetFiatDepositPaymentMethodsV5RespData) GetMinFee() string`

GetMinFee returns the MinFee field if non-nil, zero value otherwise.

### GetMinFeeOk

`func (o *GetFiatDepositPaymentMethodsV5RespData) GetMinFeeOk() (*string, bool)`

GetMinFeeOk returns a tuple with the MinFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFee

`func (o *GetFiatDepositPaymentMethodsV5RespData) SetMinFee(v string)`

SetMinFee sets MinFee field to given value.

### HasMinFee

`func (o *GetFiatDepositPaymentMethodsV5RespData) HasMinFee() bool`

HasMinFee returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *GetFiatDepositPaymentMethodsV5RespData) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GetFiatDepositPaymentMethodsV5RespData) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GetFiatDepositPaymentMethodsV5RespData) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *GetFiatDepositPaymentMethodsV5RespData) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


