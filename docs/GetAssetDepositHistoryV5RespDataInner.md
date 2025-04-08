# GetAssetDepositHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualDepBlkConfirm** | Pointer to **string** | The actual amount of blockchain confirmed in a single deposit | [optional] [default to ""]
**Amt** | Pointer to **string** | Deposit amount | [optional] [default to ""]
**AreaCodeFrom** | Pointer to **string** | If &#x60;from&#x60; is a phone number, this parameter return area code of the phone number | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency | [optional] [default to ""]
**Chain** | Pointer to **string** | Chain name | [optional] [default to ""]
**DepId** | Pointer to **string** | Deposit ID | [optional] [default to ""]
**From** | Pointer to **string** | Deposit account  If the deposit comes from an internal transfer, this field displays the account information of the internal transfer initiator, which can be a mobile phone number, email address, account name, and will return \&quot;\&quot; in other cases | [optional] [default to ""]
**FromWdId** | Pointer to **string** | Internal transfer initiator&#39;s withdrawal ID  If the deposit comes from internal transfer, this field displays the withdrawal ID of the internal transfer initiator, and will return \&quot;\&quot; in other cases | [optional] [default to ""]
**State** | Pointer to **string** | Status of deposit  &#x60;0&#x60;: Waiting for confirmation  &#x60;1&#x60;: Deposit credited    &#x60;2&#x60;: Deposit successful   &#x60;8&#x60;: Pending due to temporary deposit suspension on this crypto currency  &#x60;11&#x60;: Match the address blacklist  &#x60;12&#x60;: Account or deposit is frozen  &#x60;13&#x60;: Sub-account deposit interception  &#x60;14&#x60;: KYC limit | [optional] [default to ""]
**To** | Pointer to **string** | Deposit address  If the deposit comes from the on-chain, this field displays the on-chain address, and will return \&quot;\&quot; in other cases | [optional] [default to ""]
**Ts** | Pointer to **string** | The timestamp that the deposit record is created, Unix timestamp format in milliseconds, e.g. &#x60;1655251200000&#x60; | [optional] [default to ""]
**TxId** | Pointer to **string** | Hash record of the deposit | [optional] [default to ""]

## Methods

### NewGetAssetDepositHistoryV5RespDataInner

`func NewGetAssetDepositHistoryV5RespDataInner() *GetAssetDepositHistoryV5RespDataInner`

NewGetAssetDepositHistoryV5RespDataInner instantiates a new GetAssetDepositHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetDepositHistoryV5RespDataInnerWithDefaults

`func NewGetAssetDepositHistoryV5RespDataInnerWithDefaults() *GetAssetDepositHistoryV5RespDataInner`

NewGetAssetDepositHistoryV5RespDataInnerWithDefaults instantiates a new GetAssetDepositHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualDepBlkConfirm

`func (o *GetAssetDepositHistoryV5RespDataInner) GetActualDepBlkConfirm() string`

GetActualDepBlkConfirm returns the ActualDepBlkConfirm field if non-nil, zero value otherwise.

### GetActualDepBlkConfirmOk

`func (o *GetAssetDepositHistoryV5RespDataInner) GetActualDepBlkConfirmOk() (*string, bool)`

GetActualDepBlkConfirmOk returns a tuple with the ActualDepBlkConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDepBlkConfirm

`func (o *GetAssetDepositHistoryV5RespDataInner) SetActualDepBlkConfirm(v string)`

SetActualDepBlkConfirm sets ActualDepBlkConfirm field to given value.

### HasActualDepBlkConfirm

`func (o *GetAssetDepositHistoryV5RespDataInner) HasActualDepBlkConfirm() bool`

HasActualDepBlkConfirm returns a boolean if a field has been set.

### GetAmt

`func (o *GetAssetDepositHistoryV5RespDataInner) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *GetAssetDepositHistoryV5RespDataInner) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *GetAssetDepositHistoryV5RespDataInner) SetAmt(v string)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *GetAssetDepositHistoryV5RespDataInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.

### GetAreaCodeFrom

`func (o *GetAssetDepositHistoryV5RespDataInner) GetAreaCodeFrom() string`

GetAreaCodeFrom returns the AreaCodeFrom field if non-nil, zero value otherwise.

### GetAreaCodeFromOk

`func (o *GetAssetDepositHistoryV5RespDataInner) GetAreaCodeFromOk() (*string, bool)`

GetAreaCodeFromOk returns a tuple with the AreaCodeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaCodeFrom

`func (o *GetAssetDepositHistoryV5RespDataInner) SetAreaCodeFrom(v string)`

SetAreaCodeFrom sets AreaCodeFrom field to given value.

### HasAreaCodeFrom

`func (o *GetAssetDepositHistoryV5RespDataInner) HasAreaCodeFrom() bool`

HasAreaCodeFrom returns a boolean if a field has been set.

### GetCcy

`func (o *GetAssetDepositHistoryV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAssetDepositHistoryV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAssetDepositHistoryV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAssetDepositHistoryV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetChain

`func (o *GetAssetDepositHistoryV5RespDataInner) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GetAssetDepositHistoryV5RespDataInner) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GetAssetDepositHistoryV5RespDataInner) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *GetAssetDepositHistoryV5RespDataInner) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetDepId

`func (o *GetAssetDepositHistoryV5RespDataInner) GetDepId() string`

GetDepId returns the DepId field if non-nil, zero value otherwise.

### GetDepIdOk

`func (o *GetAssetDepositHistoryV5RespDataInner) GetDepIdOk() (*string, bool)`

GetDepIdOk returns a tuple with the DepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepId

`func (o *GetAssetDepositHistoryV5RespDataInner) SetDepId(v string)`

SetDepId sets DepId field to given value.

### HasDepId

`func (o *GetAssetDepositHistoryV5RespDataInner) HasDepId() bool`

HasDepId returns a boolean if a field has been set.

### GetFrom

`func (o *GetAssetDepositHistoryV5RespDataInner) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetAssetDepositHistoryV5RespDataInner) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetAssetDepositHistoryV5RespDataInner) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetAssetDepositHistoryV5RespDataInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetFromWdId

`func (o *GetAssetDepositHistoryV5RespDataInner) GetFromWdId() string`

GetFromWdId returns the FromWdId field if non-nil, zero value otherwise.

### GetFromWdIdOk

`func (o *GetAssetDepositHistoryV5RespDataInner) GetFromWdIdOk() (*string, bool)`

GetFromWdIdOk returns a tuple with the FromWdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWdId

`func (o *GetAssetDepositHistoryV5RespDataInner) SetFromWdId(v string)`

SetFromWdId sets FromWdId field to given value.

### HasFromWdId

`func (o *GetAssetDepositHistoryV5RespDataInner) HasFromWdId() bool`

HasFromWdId returns a boolean if a field has been set.

### GetState

`func (o *GetAssetDepositHistoryV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetAssetDepositHistoryV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetAssetDepositHistoryV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetAssetDepositHistoryV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTo

`func (o *GetAssetDepositHistoryV5RespDataInner) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetAssetDepositHistoryV5RespDataInner) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetAssetDepositHistoryV5RespDataInner) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GetAssetDepositHistoryV5RespDataInner) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTs

`func (o *GetAssetDepositHistoryV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAssetDepositHistoryV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAssetDepositHistoryV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAssetDepositHistoryV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetTxId

`func (o *GetAssetDepositHistoryV5RespDataInner) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *GetAssetDepositHistoryV5RespDataInner) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *GetAssetDepositHistoryV5RespDataInner) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *GetAssetDepositHistoryV5RespDataInner) HasTxId() bool`

HasTxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


