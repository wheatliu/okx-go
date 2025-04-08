# GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctNum** | Pointer to **string** | The account number, which can be an IBAN or other bank account number. | [optional] [default to ""]
**BankCode** | Pointer to **string** | The SWIFT code / BIC / bank code associated with the account | [optional] [default to ""]
**BankName** | Pointer to **string** | The name of the bank associated with the account | [optional] [default to ""]
**PaymentAcctId** | Pointer to **string** | The account ID for withdrawal | [optional] [default to ""]
**RecipientName** | Pointer to **string** | The name of the recipient | [optional] [default to ""]
**State** | Pointer to **string** | The state of the account  &#x60;active&#x60; | [optional] [default to ""]

## Methods

### NewGetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner

`func NewGetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner() *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner`

NewGetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner instantiates a new GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatDepositPaymentMethodsV5RespDataInnerAccountsInnerWithDefaults

`func NewGetFiatDepositPaymentMethodsV5RespDataInnerAccountsInnerWithDefaults() *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner`

NewGetFiatDepositPaymentMethodsV5RespDataInnerAccountsInnerWithDefaults instantiates a new GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctNum

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) GetAcctNum() string`

GetAcctNum returns the AcctNum field if non-nil, zero value otherwise.

### GetAcctNumOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) GetAcctNumOk() (*string, bool)`

GetAcctNumOk returns a tuple with the AcctNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctNum

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) SetAcctNum(v string)`

SetAcctNum sets AcctNum field to given value.

### HasAcctNum

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) HasAcctNum() bool`

HasAcctNum returns a boolean if a field has been set.

### GetBankCode

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.

### HasBankCode

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) HasBankCode() bool`

HasBankCode returns a boolean if a field has been set.

### GetBankName

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetPaymentAcctId

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) GetPaymentAcctId() string`

GetPaymentAcctId returns the PaymentAcctId field if non-nil, zero value otherwise.

### GetPaymentAcctIdOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) GetPaymentAcctIdOk() (*string, bool)`

GetPaymentAcctIdOk returns a tuple with the PaymentAcctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAcctId

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) SetPaymentAcctId(v string)`

SetPaymentAcctId sets PaymentAcctId field to given value.

### HasPaymentAcctId

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) HasPaymentAcctId() bool`

HasPaymentAcctId returns a boolean if a field has been set.

### GetRecipientName

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### GetState

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


