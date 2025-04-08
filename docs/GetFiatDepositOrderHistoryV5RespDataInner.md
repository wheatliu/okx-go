# GetFiatDepositOrderHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | Final amount of the transaction | [optional] [default to ""]
**CTime** | Pointer to **string** | Creation time of the transaction, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency of the transaction | [optional] [default to ""]
**ClientId** | Pointer to **string** | Client Id of the transaction | [optional] [default to ""]
**Fee** | Pointer to **string** | Transaction fee | [optional] [default to ""]
**OrdId** | Pointer to **string** | Unique Order ID | [optional] [default to ""]
**PaymentAcctId** | Pointer to **string** | ID of the payment account used | [optional] [default to ""]
**PaymentMethod** | Pointer to **string** | Payment Method, e.g. &#x60;TR_BANKS&#x60; | [optional] [default to ""]
**State** | Pointer to **string** | State of the transaction  &#x60;completed&#x60;  &#x60;failed&#x60;  &#x60;pending&#x60;  &#x60;canceled&#x60;  &#x60;inqueue&#x60; | [optional] [default to ""]
**UTime** | Pointer to **string** | Update time of the transaction, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetFiatDepositOrderHistoryV5RespDataInner

`func NewGetFiatDepositOrderHistoryV5RespDataInner() *GetFiatDepositOrderHistoryV5RespDataInner`

NewGetFiatDepositOrderHistoryV5RespDataInner instantiates a new GetFiatDepositOrderHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatDepositOrderHistoryV5RespDataInnerWithDefaults

`func NewGetFiatDepositOrderHistoryV5RespDataInnerWithDefaults() *GetFiatDepositOrderHistoryV5RespDataInner`

NewGetFiatDepositOrderHistoryV5RespDataInnerWithDefaults instantiates a new GetFiatDepositOrderHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCTime

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCcy

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClientId

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetFee

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetOrdId

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetPaymentAcctId

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetPaymentAcctId() string`

GetPaymentAcctId returns the PaymentAcctId field if non-nil, zero value otherwise.

### GetPaymentAcctIdOk

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetPaymentAcctIdOk() (*string, bool)`

GetPaymentAcctIdOk returns a tuple with the PaymentAcctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAcctId

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) SetPaymentAcctId(v string)`

SetPaymentAcctId sets PaymentAcctId field to given value.

### HasPaymentAcctId

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) HasPaymentAcctId() bool`

HasPaymentAcctId returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetState

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUTime

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetFiatDepositOrderHistoryV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


