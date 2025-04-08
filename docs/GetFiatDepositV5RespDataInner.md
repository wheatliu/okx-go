# GetFiatDepositV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | Amount of the transaction | [optional] [default to ""]
**CTime** | Pointer to **string** | The creation time of the transaction, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Ccy** | Pointer to **string** | The currency of the transaction | [optional] [default to ""]
**ClientId** | Pointer to **string** | The original request ID associated with the transaction. If it&#39;s a deposit, it&#39;s most likely an empty string (\&quot;\&quot;). | [optional] [default to ""]
**Fee** | Pointer to **string** | The transaction fee | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**PaymentAcctId** | Pointer to **string** | The ID of the payment account used | [optional] [default to ""]
**PaymentMethod** | Pointer to **string** | Payment method, e.g.&#x60;TR_BANKS&#x60; | [optional] [default to ""]
**State** | Pointer to **string** | The state of the transaction  &#x60;completed&#x60;  &#x60;failed&#x60;  &#x60;pending&#x60;  &#x60;canceled&#x60;  &#x60;inqueue&#x60;  &#x60;processing&#x60; | [optional] [default to ""]
**UTime** | Pointer to **string** | The update time of the transaction, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetFiatDepositV5RespDataInner

`func NewGetFiatDepositV5RespDataInner() *GetFiatDepositV5RespDataInner`

NewGetFiatDepositV5RespDataInner instantiates a new GetFiatDepositV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatDepositV5RespDataInnerWithDefaults

`func NewGetFiatDepositV5RespDataInnerWithDefaults() *GetFiatDepositV5RespDataInner`

NewGetFiatDepositV5RespDataInnerWithDefaults instantiates a new GetFiatDepositV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *GetFiatDepositV5RespDataInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetFiatDepositV5RespDataInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetFiatDepositV5RespDataInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetFiatDepositV5RespDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCTime

`func (o *GetFiatDepositV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetFiatDepositV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetFiatDepositV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetFiatDepositV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCcy

`func (o *GetFiatDepositV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFiatDepositV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFiatDepositV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFiatDepositV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClientId

`func (o *GetFiatDepositV5RespDataInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetFiatDepositV5RespDataInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetFiatDepositV5RespDataInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetFiatDepositV5RespDataInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetFee

`func (o *GetFiatDepositV5RespDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetFiatDepositV5RespDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetFiatDepositV5RespDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetFiatDepositV5RespDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetOrdId

`func (o *GetFiatDepositV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetFiatDepositV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetFiatDepositV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetFiatDepositV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetPaymentAcctId

`func (o *GetFiatDepositV5RespDataInner) GetPaymentAcctId() string`

GetPaymentAcctId returns the PaymentAcctId field if non-nil, zero value otherwise.

### GetPaymentAcctIdOk

`func (o *GetFiatDepositV5RespDataInner) GetPaymentAcctIdOk() (*string, bool)`

GetPaymentAcctIdOk returns a tuple with the PaymentAcctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAcctId

`func (o *GetFiatDepositV5RespDataInner) SetPaymentAcctId(v string)`

SetPaymentAcctId sets PaymentAcctId field to given value.

### HasPaymentAcctId

`func (o *GetFiatDepositV5RespDataInner) HasPaymentAcctId() bool`

HasPaymentAcctId returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *GetFiatDepositV5RespDataInner) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GetFiatDepositV5RespDataInner) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GetFiatDepositV5RespDataInner) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *GetFiatDepositV5RespDataInner) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetState

`func (o *GetFiatDepositV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetFiatDepositV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetFiatDepositV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetFiatDepositV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUTime

`func (o *GetFiatDepositV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetFiatDepositV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetFiatDepositV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetFiatDepositV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


