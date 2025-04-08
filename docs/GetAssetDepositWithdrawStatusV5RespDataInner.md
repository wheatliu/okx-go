# GetAssetDepositWithdrawStatusV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstCompleteTime** | Pointer to **string** | Estimated complete time  The timezone is &#x60;UTC+8&#x60;. The format is MM/dd/yyyy, h:mm:ss AM/PM   estCompleteTime is only an approximate estimated time, for reference only. | [optional] [default to ""]
**State** | Pointer to **string** | The detailed stage and status of the deposit/withdrawal   The message in front of the colon is the stage; the message after the colon is the ongoing status. | [optional] [default to ""]
**TxId** | Pointer to **string** | Hash record on-chain  For withdrawal, if the &#x60;txId&#x60; has already been generated, it will return the value, otherwise, it will return \&quot;\&quot;. | [optional] [default to ""]
**WdId** | Pointer to **string** | Withdrawal ID  When retrieving deposit status, wdId returns blank \&quot;\&quot;. | [optional] [default to ""]

## Methods

### NewGetAssetDepositWithdrawStatusV5RespDataInner

`func NewGetAssetDepositWithdrawStatusV5RespDataInner() *GetAssetDepositWithdrawStatusV5RespDataInner`

NewGetAssetDepositWithdrawStatusV5RespDataInner instantiates a new GetAssetDepositWithdrawStatusV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetDepositWithdrawStatusV5RespDataInnerWithDefaults

`func NewGetAssetDepositWithdrawStatusV5RespDataInnerWithDefaults() *GetAssetDepositWithdrawStatusV5RespDataInner`

NewGetAssetDepositWithdrawStatusV5RespDataInnerWithDefaults instantiates a new GetAssetDepositWithdrawStatusV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstCompleteTime

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) GetEstCompleteTime() string`

GetEstCompleteTime returns the EstCompleteTime field if non-nil, zero value otherwise.

### GetEstCompleteTimeOk

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) GetEstCompleteTimeOk() (*string, bool)`

GetEstCompleteTimeOk returns a tuple with the EstCompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstCompleteTime

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) SetEstCompleteTime(v string)`

SetEstCompleteTime sets EstCompleteTime field to given value.

### HasEstCompleteTime

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) HasEstCompleteTime() bool`

HasEstCompleteTime returns a boolean if a field has been set.

### GetState

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTxId

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetWdId

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) GetWdId() string`

GetWdId returns the WdId field if non-nil, zero value otherwise.

### GetWdIdOk

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) GetWdIdOk() (*string, bool)`

GetWdIdOk returns a tuple with the WdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdId

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) SetWdId(v string)`

SetWdId sets WdId field to given value.

### HasWdId

`func (o *GetAssetDepositWithdrawStatusV5RespDataInner) HasWdId() bool`

HasWdId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


