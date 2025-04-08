# CreateAssetWithdrawalV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amt** | **string** | Withdrawal amount  Withdrawal fee is not included in withdrawal amount. Please reserve sufficient transaction fees when withdrawing.  You can get fee amount by .  For &#x60;internal transfer&#x60;, transaction fee is always &#x60;0&#x60;. | [default to ""]
**AreaCode** | Pointer to **string** | Area code for the phone number, e.g. &#x60;86&#x60;  If &#x60;toAddr&#x60; is a phone number, this parameter is required.  Apply to &#x60;internal transfer&#x60; | [optional] [default to ""]
**Ccy** | **string** | Currency, e.g. &#x60;USDT&#x60; | [default to ""]
**Chain** | Pointer to **string** | Chain name  There are multiple chains under some currencies, such as &#x60;USDT&#x60; has &#x60;USDT-ERC20&#x60;, &#x60;USDT-TRC20&#x60;  If the parameter is not filled in, the default will be the main chain.  When you withdrawal the non-tradable asset, if the parameter is not filled in, the default will be the unique withdrawal chain.  Apply to &#x60;on-chain withdrawal&#x60;.  You can get supported chain name by the endpoint of . | [optional] [default to ""]
**ClientId** | Pointer to **string** | Client-supplied ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [optional] [default to ""]
**Dest** | **string** | Withdrawal method  &#x60;3&#x60;: internal transfer  &#x60;4&#x60;: on-chain withdrawal | [default to ""]
**RcvrInfo** | Pointer to [**CreateAssetWithdrawalV5ReqRcvrInfo**](CreateAssetWithdrawalV5ReqRcvrInfo.md) |  | [optional] 
**ToAddr** | **string** | &#x60;toAddr&#x60; should be a trusted address/account.   If your &#x60;dest&#x60; is &#x60;4&#x60;, some crypto currency addresses are formatted as &#x60;&#39;address:tag&#39;&#x60;, e.g. &#x60;&#39;ARDOR-7JF3-8F2E-QUWZ-CAN7F:123456&#39;&#x60;  If your &#x60;dest&#x60; is &#x60;3&#x60;,&#x60;toAddr&#x60; should be a recipient address which can be email, phone or login account name (account name is only for sub-account). | [default to ""]

## Methods

### NewCreateAssetWithdrawalV5Req

`func NewCreateAssetWithdrawalV5Req(amt string, ccy string, dest string, toAddr string, ) *CreateAssetWithdrawalV5Req`

NewCreateAssetWithdrawalV5Req instantiates a new CreateAssetWithdrawalV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssetWithdrawalV5ReqWithDefaults

`func NewCreateAssetWithdrawalV5ReqWithDefaults() *CreateAssetWithdrawalV5Req`

NewCreateAssetWithdrawalV5ReqWithDefaults instantiates a new CreateAssetWithdrawalV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmt

`func (o *CreateAssetWithdrawalV5Req) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateAssetWithdrawalV5Req) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateAssetWithdrawalV5Req) SetAmt(v string)`

SetAmt sets Amt field to given value.


### GetAreaCode

`func (o *CreateAssetWithdrawalV5Req) GetAreaCode() string`

GetAreaCode returns the AreaCode field if non-nil, zero value otherwise.

### GetAreaCodeOk

`func (o *CreateAssetWithdrawalV5Req) GetAreaCodeOk() (*string, bool)`

GetAreaCodeOk returns a tuple with the AreaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaCode

`func (o *CreateAssetWithdrawalV5Req) SetAreaCode(v string)`

SetAreaCode sets AreaCode field to given value.

### HasAreaCode

`func (o *CreateAssetWithdrawalV5Req) HasAreaCode() bool`

HasAreaCode returns a boolean if a field has been set.

### GetCcy

`func (o *CreateAssetWithdrawalV5Req) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateAssetWithdrawalV5Req) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateAssetWithdrawalV5Req) SetCcy(v string)`

SetCcy sets Ccy field to given value.


### GetChain

`func (o *CreateAssetWithdrawalV5Req) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *CreateAssetWithdrawalV5Req) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *CreateAssetWithdrawalV5Req) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *CreateAssetWithdrawalV5Req) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetClientId

`func (o *CreateAssetWithdrawalV5Req) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateAssetWithdrawalV5Req) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateAssetWithdrawalV5Req) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CreateAssetWithdrawalV5Req) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetDest

`func (o *CreateAssetWithdrawalV5Req) GetDest() string`

GetDest returns the Dest field if non-nil, zero value otherwise.

### GetDestOk

`func (o *CreateAssetWithdrawalV5Req) GetDestOk() (*string, bool)`

GetDestOk returns a tuple with the Dest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDest

`func (o *CreateAssetWithdrawalV5Req) SetDest(v string)`

SetDest sets Dest field to given value.


### GetRcvrInfo

`func (o *CreateAssetWithdrawalV5Req) GetRcvrInfo() CreateAssetWithdrawalV5ReqRcvrInfo`

GetRcvrInfo returns the RcvrInfo field if non-nil, zero value otherwise.

### GetRcvrInfoOk

`func (o *CreateAssetWithdrawalV5Req) GetRcvrInfoOk() (*CreateAssetWithdrawalV5ReqRcvrInfo, bool)`

GetRcvrInfoOk returns a tuple with the RcvrInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcvrInfo

`func (o *CreateAssetWithdrawalV5Req) SetRcvrInfo(v CreateAssetWithdrawalV5ReqRcvrInfo)`

SetRcvrInfo sets RcvrInfo field to given value.

### HasRcvrInfo

`func (o *CreateAssetWithdrawalV5Req) HasRcvrInfo() bool`

HasRcvrInfo returns a boolean if a field has been set.

### GetToAddr

`func (o *CreateAssetWithdrawalV5Req) GetToAddr() string`

GetToAddr returns the ToAddr field if non-nil, zero value otherwise.

### GetToAddrOk

`func (o *CreateAssetWithdrawalV5Req) GetToAddrOk() (*string, bool)`

GetToAddrOk returns a tuple with the ToAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddr

`func (o *CreateAssetWithdrawalV5Req) SetToAddr(v string)`

SetToAddr sets ToAddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


