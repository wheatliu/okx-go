# GetAssetWithdrawalHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddrEx** | Pointer to **map[string]interface{}** | Withdrawal address attachment (This will not be returned if the currency does not require this) e.g. TONCOIN attached tag name is comment, the return will be {&#39;comment&#39;:&#39;123456&#39;} | [optional] 
**Amt** | Pointer to **string** | Withdrawal amount | [optional] [default to ""]
**AreaCodeFrom** | Pointer to **string** | Area code for the phone number  If &#x60;from&#x60; is a phone number, this parameter returns the area code for the phone number | [optional] [default to ""]
**AreaCodeTo** | Pointer to **string** | Area code for the phone number  If &#x60;to&#x60; is a phone number, this parameter returns the area code for the phone number | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency | [optional] [default to ""]
**Chain** | Pointer to **string** | Chain name, e.g. &#x60;USDT-ERC20&#x60;, &#x60;USDT-TRC20&#x60; | [optional] [default to ""]
**ClientId** | Pointer to **string** | Client-supplied ID | [optional] [default to ""]
**Fee** | Pointer to **string** | Withdrawal fee amount | [optional] [default to ""]
**FeeCcy** | Pointer to **string** | Withdrawal fee currency, e.g. &#x60;USDT&#x60; | [optional] [default to ""]
**From** | Pointer to **string** | Withdrawal account   It can be &#x60;email&#x60;/&#x60;phone&#x60;/&#x60;sub-account name&#x60; | [optional] [default to ""]
**Memo** | Pointer to **string** | Some currencies require this parameter for withdrawals. This is not returned if not required. | [optional] [default to ""]
**NonTradableAsset** | Pointer to **bool** | Whether it is a non-tradable asset or not  &#x60;true&#x60;: non-tradable asset, &#x60;false&#x60;: tradable asset | [optional] 
**Note** | Pointer to **string** | Withdrawal note | [optional] [default to ""]
**PmtId** | Pointer to **string** | Some currencies require a payment ID for withdrawals. This is not returned if not required. | [optional] [default to ""]
**State** | Pointer to **string** | Status of withdrawal | [optional] [default to ""]
**Tag** | Pointer to **string** | Some currencies require a tag for withdrawals. This is not returned if not required. | [optional] [default to ""]
**To** | Pointer to **string** | Receiving address | [optional] [default to ""]
**Ts** | Pointer to **string** | Time the withdrawal request was submitted, Unix timestamp format in milliseconds, e.g. &#x60;1655251200000&#x60;. | [optional] [default to ""]
**TxId** | Pointer to **string** | Hash record of the withdrawal  This parameter will return \&quot;\&quot; for internal transfers. | [optional] [default to ""]
**WdId** | Pointer to **string** | Withdrawal ID | [optional] [default to ""]

## Methods

### NewGetAssetWithdrawalHistoryV5RespDataInner

`func NewGetAssetWithdrawalHistoryV5RespDataInner() *GetAssetWithdrawalHistoryV5RespDataInner`

NewGetAssetWithdrawalHistoryV5RespDataInner instantiates a new GetAssetWithdrawalHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetWithdrawalHistoryV5RespDataInnerWithDefaults

`func NewGetAssetWithdrawalHistoryV5RespDataInnerWithDefaults() *GetAssetWithdrawalHistoryV5RespDataInner`

NewGetAssetWithdrawalHistoryV5RespDataInnerWithDefaults instantiates a new GetAssetWithdrawalHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddrEx

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetAddrEx() map[string]interface{}`

GetAddrEx returns the AddrEx field if non-nil, zero value otherwise.

### GetAddrExOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetAddrExOk() (*map[string]interface{}, bool)`

GetAddrExOk returns a tuple with the AddrEx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrEx

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetAddrEx(v map[string]interface{})`

SetAddrEx sets AddrEx field to given value.

### HasAddrEx

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasAddrEx() bool`

HasAddrEx returns a boolean if a field has been set.

### GetAmt

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetAreaCodeFrom

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetAreaCodeFrom() string`

GetAreaCodeFrom returns the AreaCodeFrom field if non-nil, zero value otherwise.

### GetAreaCodeFromOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetAreaCodeFromOk() (*string, bool)`

GetAreaCodeFromOk returns a tuple with the AreaCodeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaCodeFrom

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetAreaCodeFrom(v string)`

SetAreaCodeFrom sets AreaCodeFrom field to given value.

### HasAreaCodeFrom

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasAreaCodeFrom() bool`

HasAreaCodeFrom returns a boolean if a field has been set.

### GetAreaCodeTo

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetAreaCodeTo() string`

GetAreaCodeTo returns the AreaCodeTo field if non-nil, zero value otherwise.

### GetAreaCodeToOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetAreaCodeToOk() (*string, bool)`

GetAreaCodeToOk returns a tuple with the AreaCodeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaCodeTo

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetAreaCodeTo(v string)`

SetAreaCodeTo sets AreaCodeTo field to given value.

### HasAreaCodeTo

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasAreaCodeTo() bool`

HasAreaCodeTo returns a boolean if a field has been set.

### GetCcy

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetChain

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetClientId

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetFee

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeCcy

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetFeeCcy() string`

GetFeeCcy returns the FeeCcy field if non-nil, zero value otherwise.

### GetFeeCcyOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetFeeCcyOk() (*string, bool)`

GetFeeCcyOk returns a tuple with the FeeCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCcy

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetFeeCcy(v string)`

SetFeeCcy sets FeeCcy field to given value.

### HasFeeCcy

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasFeeCcy() bool`

HasFeeCcy returns a boolean if a field has been set.

### GetFrom

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetMemo

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetNonTradableAsset

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetNonTradableAsset() bool`

GetNonTradableAsset returns the NonTradableAsset field if non-nil, zero value otherwise.

### GetNonTradableAssetOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetNonTradableAssetOk() (*bool, bool)`

GetNonTradableAssetOk returns a tuple with the NonTradableAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonTradableAsset

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetNonTradableAsset(v bool)`

SetNonTradableAsset sets NonTradableAsset field to given value.

### HasNonTradableAsset

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasNonTradableAsset() bool`

HasNonTradableAsset returns a boolean if a field has been set.

### GetNote

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPmtId

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetPmtId() string`

GetPmtId returns the PmtId field if non-nil, zero value otherwise.

### GetPmtIdOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetPmtIdOk() (*string, bool)`

GetPmtIdOk returns a tuple with the PmtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmtId

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetPmtId(v string)`

SetPmtId sets PmtId field to given value.

### HasPmtId

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasPmtId() bool`

HasPmtId returns a boolean if a field has been set.

### GetState

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTag

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTo

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTs

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetTxId

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetWdId

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetWdId() string`

GetWdId returns the WdId field if non-nil, zero value otherwise.

### GetWdIdOk

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) GetWdIdOk() (*string, bool)`

GetWdIdOk returns a tuple with the WdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdId

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) SetWdId(v string)`

SetWdId sets WdId field to given value.

### HasWdId

`func (o *GetAssetWithdrawalHistoryV5RespDataInner) HasWdId() bool`

HasWdId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


