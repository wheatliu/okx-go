# CreateAssetWithdrawalV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | Pointer to **string** | Withdrawal amount | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency | [optional] [default to ""]
**Chain** | Pointer to **string** | Chain name, e.g. &#x60;USDT-ERC20&#x60;, &#x60;USDT-TRC20&#x60; | [optional] [default to ""]
**ClientId** | Pointer to **string** | Client-supplied ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [optional] [default to ""]
**WdId** | Pointer to **string** | Withdrawal ID | [optional] [default to ""]

## Methods

### NewCreateAssetWithdrawalV5RespData

`func NewCreateAssetWithdrawalV5RespData() *CreateAssetWithdrawalV5RespData`

NewCreateAssetWithdrawalV5RespData instantiates a new CreateAssetWithdrawalV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssetWithdrawalV5RespDataWithDefaults

`func NewCreateAssetWithdrawalV5RespDataWithDefaults() *CreateAssetWithdrawalV5RespData`

NewCreateAssetWithdrawalV5RespDataWithDefaults instantiates a new CreateAssetWithdrawalV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAssetWithdrawalV5RespData) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAssetWithdrawalV5RespData) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAssetWithdrawalV5RespData) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *CreateAssetWithdrawalV5RespData) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetCcy

`func (o *CreateAssetWithdrawalV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateAssetWithdrawalV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateAssetWithdrawalV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateAssetWithdrawalV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetChain

`func (o *CreateAssetWithdrawalV5RespData) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *CreateAssetWithdrawalV5RespData) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *CreateAssetWithdrawalV5RespData) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *CreateAssetWithdrawalV5RespData) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetClientId

`func (o *CreateAssetWithdrawalV5RespData) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateAssetWithdrawalV5RespData) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateAssetWithdrawalV5RespData) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CreateAssetWithdrawalV5RespData) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetWdId

`func (o *CreateAssetWithdrawalV5RespData) GetWdId() string`

GetWdId returns the WdId field if non-nil, zero value otherwise.

### GetWdIdOk

`func (o *CreateAssetWithdrawalV5RespData) GetWdIdOk() (*string, bool)`

GetWdIdOk returns a tuple with the WdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdId

`func (o *CreateAssetWithdrawalV5RespData) SetWdId(v string)`

SetWdId sets WdId field to given value.

### HasWdId

`func (o *CreateAssetWithdrawalV5RespData) HasWdId() bool`

HasWdId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


