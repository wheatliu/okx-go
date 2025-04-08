# GetAssetDepositAddressV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | Pointer to **string** | Deposit address | [optional] [default to ""]
**AddrEx** | Pointer to **map[string]interface{}** | Deposit address attachment (This will not be returned if the currency does not require this)  e.g. &#x60;TONCOIN&#x60; attached tag name is &#x60;comment&#x60;, the return will be &#x60;{&#39;comment&#39;:&#39;123456&#39;}&#x60; | [optional] 
**Ccy** | Pointer to **string** | Currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**Chain** | Pointer to **string** | Chain name, e.g. &#x60;USDT-ERC20&#x60;, &#x60;USDT-TRC20&#x60; | [optional] [default to ""]
**CtAddr** | Pointer to **string** | Last 6 digits of contract address | [optional] [default to ""]
**Memo** | Pointer to **string** | Deposit memo (This will not be returned if the currency does not require a memo for deposit) | [optional] [default to ""]
**PmtId** | Pointer to **string** | Deposit payment ID (This will not be returned if the currency does not require a payment_id for deposit) | [optional] [default to ""]
**Selected** | Pointer to **bool** | Return &#x60;true&#x60; if the current deposit address is selected by the website page | [optional] 
**Tag** | Pointer to **string** | Deposit tag (This will not be returned if the currency does not require a tag for deposit) | [optional] [default to ""]
**To** | Pointer to **string** | The beneficiary account  &#x60;6&#x60;: Funding account &#x60;18&#x60;: Trading account  The users under some entity (e.g. Brazil) only support deposit to trading account. | [optional] [default to ""]
**VerifiedName** | Pointer to **string** | Verified name (for recipient) | [optional] [default to ""]

## Methods

### NewGetAssetDepositAddressV5RespData

`func NewGetAssetDepositAddressV5RespData() *GetAssetDepositAddressV5RespData`

NewGetAssetDepositAddressV5RespData instantiates a new GetAssetDepositAddressV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetDepositAddressV5RespDataWithDefaults

`func NewGetAssetDepositAddressV5RespDataWithDefaults() *GetAssetDepositAddressV5RespData`

NewGetAssetDepositAddressV5RespDataWithDefaults instantiates a new GetAssetDepositAddressV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *GetAssetDepositAddressV5RespData) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *GetAssetDepositAddressV5RespData) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *GetAssetDepositAddressV5RespData) SetAddr(v string)`

SetAddr sets Addr field to given value.

### HasAddr

`func (o *GetAssetDepositAddressV5RespData) HasAddr() bool`

HasAddr returns a boolean if a field has been set.

### GetAddrEx

`func (o *GetAssetDepositAddressV5RespData) GetAddrEx() map[string]interface{}`

GetAddrEx returns the AddrEx field if non-nil, zero value otherwise.

### GetAddrExOk

`func (o *GetAssetDepositAddressV5RespData) GetAddrExOk() (*map[string]interface{}, bool)`

GetAddrExOk returns a tuple with the AddrEx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrEx

`func (o *GetAssetDepositAddressV5RespData) SetAddrEx(v map[string]interface{})`

SetAddrEx sets AddrEx field to given value.

### HasAddrEx

`func (o *GetAssetDepositAddressV5RespData) HasAddrEx() bool`

HasAddrEx returns a boolean if a field has been set.

### GetCcy

`func (o *GetAssetDepositAddressV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAssetDepositAddressV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAssetDepositAddressV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAssetDepositAddressV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetChain

`func (o *GetAssetDepositAddressV5RespData) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GetAssetDepositAddressV5RespData) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GetAssetDepositAddressV5RespData) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *GetAssetDepositAddressV5RespData) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetCtAddr

`func (o *GetAssetDepositAddressV5RespData) GetCtAddr() string`

GetCtAddr returns the CtAddr field if non-nil, zero value otherwise.

### GetCtAddrOk

`func (o *GetAssetDepositAddressV5RespData) GetCtAddrOk() (*string, bool)`

GetCtAddrOk returns a tuple with the CtAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtAddr

`func (o *GetAssetDepositAddressV5RespData) SetCtAddr(v string)`

SetCtAddr sets CtAddr field to given value.

### HasCtAddr

`func (o *GetAssetDepositAddressV5RespData) HasCtAddr() bool`

HasCtAddr returns a boolean if a field has been set.

### GetMemo

`func (o *GetAssetDepositAddressV5RespData) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GetAssetDepositAddressV5RespData) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GetAssetDepositAddressV5RespData) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GetAssetDepositAddressV5RespData) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPmtId

`func (o *GetAssetDepositAddressV5RespData) GetPmtId() string`

GetPmtId returns the PmtId field if non-nil, zero value otherwise.

### GetPmtIdOk

`func (o *GetAssetDepositAddressV5RespData) GetPmtIdOk() (*string, bool)`

GetPmtIdOk returns a tuple with the PmtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmtId

`func (o *GetAssetDepositAddressV5RespData) SetPmtId(v string)`

SetPmtId sets PmtId field to given value.

### HasPmtId

`func (o *GetAssetDepositAddressV5RespData) HasPmtId() bool`

HasPmtId returns a boolean if a field has been set.

### GetSelected

`func (o *GetAssetDepositAddressV5RespData) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *GetAssetDepositAddressV5RespData) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *GetAssetDepositAddressV5RespData) SetSelected(v bool)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *GetAssetDepositAddressV5RespData) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### GetTag

`func (o *GetAssetDepositAddressV5RespData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetAssetDepositAddressV5RespData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetAssetDepositAddressV5RespData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetAssetDepositAddressV5RespData) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTo

`func (o *GetAssetDepositAddressV5RespData) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetAssetDepositAddressV5RespData) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetAssetDepositAddressV5RespData) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GetAssetDepositAddressV5RespData) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetVerifiedName

`func (o *GetAssetDepositAddressV5RespData) GetVerifiedName() string`

GetVerifiedName returns the VerifiedName field if non-nil, zero value otherwise.

### GetVerifiedNameOk

`func (o *GetAssetDepositAddressV5RespData) GetVerifiedNameOk() (*string, bool)`

GetVerifiedNameOk returns a tuple with the VerifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedName

`func (o *GetAssetDepositAddressV5RespData) SetVerifiedName(v string)`

SetVerifiedName sets VerifiedName field to given value.

### HasVerifiedName

`func (o *GetAssetDepositAddressV5RespData) HasVerifiedName() bool`

HasVerifiedName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


