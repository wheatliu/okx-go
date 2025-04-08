# GetFiatDepositV5RespData

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

### NewGetFiatDepositV5RespData

`func NewGetFiatDepositV5RespData() *GetFiatDepositV5RespData`

NewGetFiatDepositV5RespData instantiates a new GetFiatDepositV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatDepositV5RespDataWithDefaults

`func NewGetFiatDepositV5RespDataWithDefaults() *GetFiatDepositV5RespData`

NewGetFiatDepositV5RespDataWithDefaults instantiates a new GetFiatDepositV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *GetFiatDepositV5RespData) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetFiatDepositV5RespData) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetFiatDepositV5RespData) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetFiatDepositV5RespData) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCTime

`func (o *GetFiatDepositV5RespData) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetFiatDepositV5RespData) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetFiatDepositV5RespData) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetFiatDepositV5RespData) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCcy

`func (o *GetFiatDepositV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFiatDepositV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFiatDepositV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFiatDepositV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClientId

`func (o *GetFiatDepositV5RespData) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetFiatDepositV5RespData) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetFiatDepositV5RespData) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetFiatDepositV5RespData) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetFee

`func (o *GetFiatDepositV5RespData) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetFiatDepositV5RespData) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetFiatDepositV5RespData) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetFiatDepositV5RespData) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetOrdId

`func (o *GetFiatDepositV5RespData) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetFiatDepositV5RespData) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetFiatDepositV5RespData) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetFiatDepositV5RespData) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetPaymentAcctId

`func (o *GetFiatDepositV5RespData) GetPaymentAcctId() string`

GetPaymentAcctId returns the PaymentAcctId field if non-nil, zero value otherwise.

### GetPaymentAcctIdOk

`func (o *GetFiatDepositV5RespData) GetPaymentAcctIdOk() (*string, bool)`

GetPaymentAcctIdOk returns a tuple with the PaymentAcctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAcctId

`func (o *GetFiatDepositV5RespData) SetPaymentAcctId(v string)`

SetPaymentAcctId sets PaymentAcctId field to given value.

### HasPaymentAcctId

`func (o *GetFiatDepositV5RespData) HasPaymentAcctId() bool`

HasPaymentAcctId returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *GetFiatDepositV5RespData) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GetFiatDepositV5RespData) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GetFiatDepositV5RespData) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *GetFiatDepositV5RespData) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetState

`func (o *GetFiatDepositV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetFiatDepositV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetFiatDepositV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetFiatDepositV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUTime

`func (o *GetFiatDepositV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetFiatDepositV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetFiatDepositV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetFiatDepositV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


