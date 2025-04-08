# GetFiatDepositPaymentMethodsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner**](GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner.md) | An array containing information about payment accounts associated with the currency and method. | [optional] 
**Ccy** | Pointer to **string** | Fiat currency | [optional] [default to ""]
**FeeRate** | Pointer to **string** | The fee rate for each deposit, expressed as a percentage  e.g. &#x60;0.02&#x60; represents 2 percent fee for each transaction. | [optional] [default to ""]
**Limits** | Pointer to [**GetFiatDepositPaymentMethodsV5RespDataInnerLimits**](GetFiatDepositPaymentMethodsV5RespDataInnerLimits.md) |  | [optional] 
**MinFee** | Pointer to **string** | The minimum fee for each deposit | [optional] [default to ""]
**PaymentMethod** | Pointer to **string** | The payment method associated with the currency  &#x60;TR_BANKS&#x60;  &#x60;PIX&#x60;  &#x60;SEPA&#x60; | [optional] [default to ""]

## Methods

### NewGetFiatDepositPaymentMethodsV5RespDataInner

`func NewGetFiatDepositPaymentMethodsV5RespDataInner() *GetFiatDepositPaymentMethodsV5RespDataInner`

NewGetFiatDepositPaymentMethodsV5RespDataInner instantiates a new GetFiatDepositPaymentMethodsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatDepositPaymentMethodsV5RespDataInnerWithDefaults

`func NewGetFiatDepositPaymentMethodsV5RespDataInnerWithDefaults() *GetFiatDepositPaymentMethodsV5RespDataInner`

NewGetFiatDepositPaymentMethodsV5RespDataInnerWithDefaults instantiates a new GetFiatDepositPaymentMethodsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) GetAccounts() []GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) GetAccountsOk() (*[]GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) SetAccounts(v []GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetCcy

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetFeeRate

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) GetFeeRate() string`

GetFeeRate returns the FeeRate field if non-nil, zero value otherwise.

### GetFeeRateOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) GetFeeRateOk() (*string, bool)`

GetFeeRateOk returns a tuple with the FeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRate

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) SetFeeRate(v string)`

SetFeeRate sets FeeRate field to given value.

### HasFeeRate

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) HasFeeRate() bool`

HasFeeRate returns a boolean if a field has been set.

### GetLimits

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) GetLimits() GetFiatDepositPaymentMethodsV5RespDataInnerLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) GetLimitsOk() (*GetFiatDepositPaymentMethodsV5RespDataInnerLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) SetLimits(v GetFiatDepositPaymentMethodsV5RespDataInnerLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetMinFee

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) GetMinFee() string`

GetMinFee returns the MinFee field if non-nil, zero value otherwise.

### GetMinFeeOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) GetMinFeeOk() (*string, bool)`

GetMinFeeOk returns a tuple with the MinFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFee

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) SetMinFee(v string)`

SetMinFee sets MinFee field to given value.

### HasMinFee

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) HasMinFee() bool`

HasMinFee returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *GetFiatDepositPaymentMethodsV5RespDataInner) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


