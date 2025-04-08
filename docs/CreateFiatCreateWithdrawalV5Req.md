# CreateFiatCreateWithdrawalV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | **string** | Requested withdrawal amount before fees. Has to be less than or equal to 2 decimal points double | [default to ""]
**Ccy** | **string** | Currency for withdrawal, must match currency allowed for paymentMethod | [default to ""]
**ClientId** | **string** | Client-supplied ID, A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters    e.g. &#x60;194a6975e98246538faeb0fab0d502df&#x60; | [default to ""]
**PaymentAcctId** | **string** | Payment account id to withdraw to, retrieved from get withdrawal payment methods API | [default to ""]
**PaymentMethod** | **string** | Payment method to use for withdrawal  &#x60;TR_BANKS&#x60;  &#x60;PIX&#x60;  &#x60;SEPA&#x60; | [default to ""]

## Methods

### NewCreateFiatCreateWithdrawalV5Req

`func NewCreateFiatCreateWithdrawalV5Req(amt string, ccy string, clientId string, paymentAcctId string, paymentMethod string, ) *CreateFiatCreateWithdrawalV5Req`

NewCreateFiatCreateWithdrawalV5Req instantiates a new CreateFiatCreateWithdrawalV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFiatCreateWithdrawalV5ReqWithDefaults

`func NewCreateFiatCreateWithdrawalV5ReqWithDefaults() *CreateFiatCreateWithdrawalV5Req`

NewCreateFiatCreateWithdrawalV5ReqWithDefaults instantiates a new CreateFiatCreateWithdrawalV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateFiatCreateWithdrawalV5Req) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateFiatCreateWithdrawalV5Req) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateFiatCreateWithdrawalV5Req) SetAmt(v string)`

SetAmt sets Amt field to given value.


### GetCcy

`func (o *CreateFiatCreateWithdrawalV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateFiatCreateWithdrawalV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateFiatCreateWithdrawalV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.


### GetClientId

`func (o *CreateFiatCreateWithdrawalV5Req) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateFiatCreateWithdrawalV5Req) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateFiatCreateWithdrawalV5Req) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetPaymentAcctId

`func (o *CreateFiatCreateWithdrawalV5Req) GetPaymentAcctId() string`

GetPaymentAcctId returns the PaymentAcctId field if non-nil, zero value otherwise.

### GetPaymentAcctIdOk

`func (o *CreateFiatCreateWithdrawalV5Req) GetPaymentAcctIdOk() (*string, bool)`

GetPaymentAcctIdOk returns a tuple with the PaymentAcctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAcctId

`func (o *CreateFiatCreateWithdrawalV5Req) SetPaymentAcctId(v string)`

SetPaymentAcctId sets PaymentAcctId field to given value.


### GetPaymentMethod

`func (o *CreateFiatCreateWithdrawalV5Req) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CreateFiatCreateWithdrawalV5Req) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CreateFiatCreateWithdrawalV5Req) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


