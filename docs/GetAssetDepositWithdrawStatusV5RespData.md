# GetAssetDepositWithdrawStatusV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstCompleteTime** | Pointer to **string** | Estimated complete time  The timezone is &#x60;UTC+8&#x60;. The format is MM/dd/yyyy, h:mm:ss AM/PM   estCompleteTime is only an approximate estimated time, for reference only. | [optional] [default to ""]
**State** | Pointer to **string** | The detailed stage and status of the deposit/withdrawal   The message in front of the colon is the stage; the message after the colon is the ongoing status. | [optional] [default to ""]
**TxId** | Pointer to **string** | Hash record on-chain  For withdrawal, if the &#x60;txId&#x60; has already been generated, it will return the value, otherwise, it will return \&quot;\&quot;. | [optional] [default to ""]
**WdId** | Pointer to **string** | Withdrawal ID  When retrieving deposit status, wdId returns blank \&quot;\&quot;. | [optional] [default to ""]

## Methods

### NewGetAssetDepositWithdrawStatusV5RespData

`func NewGetAssetDepositWithdrawStatusV5RespData() *GetAssetDepositWithdrawStatusV5RespData`

NewGetAssetDepositWithdrawStatusV5RespData instantiates a new GetAssetDepositWithdrawStatusV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetDepositWithdrawStatusV5RespDataWithDefaults

`func NewGetAssetDepositWithdrawStatusV5RespDataWithDefaults() *GetAssetDepositWithdrawStatusV5RespData`

NewGetAssetDepositWithdrawStatusV5RespDataWithDefaults instantiates a new GetAssetDepositWithdrawStatusV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstCompleteTime

`func (o *GetAssetDepositWithdrawStatusV5RespData) GetEstCompleteTime() string`

GetEstCompleteTime returns the EstCompleteTime field if non-nil, zero value otherwise.

### GetEstCompleteTimeOk

`func (o *GetAssetDepositWithdrawStatusV5RespData) GetEstCompleteTimeOk() (*string, bool)`

GetEstCompleteTimeOk returns a tuple with the EstCompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstCompleteTime

`func (o *GetAssetDepositWithdrawStatusV5RespData) SetEstCompleteTime(v string)`

SetEstCompleteTime sets EstCompleteTime field to given value.

### HasEstCompleteTime

`func (o *GetAssetDepositWithdrawStatusV5RespData) HasEstCompleteTime() bool`

HasEstCompleteTime returns a boolean if a field has been set.

### GetState

`func (o *GetAssetDepositWithdrawStatusV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetAssetDepositWithdrawStatusV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetAssetDepositWithdrawStatusV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetAssetDepositWithdrawStatusV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTxId

`func (o *GetAssetDepositWithdrawStatusV5RespData) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *GetAssetDepositWithdrawStatusV5RespData) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *GetAssetDepositWithdrawStatusV5RespData) SetTxId(v string)`

SetTxId sets TxId field to given value.

### HasTxId

`func (o *GetAssetDepositWithdrawStatusV5RespData) HasTxId() bool`

HasTxId returns a boolean if a field has been set.

### GetWdId

`func (o *GetAssetDepositWithdrawStatusV5RespData) GetWdId() string`

GetWdId returns the WdId field if non-nil, zero value otherwise.

### GetWdIdOk

`func (o *GetAssetDepositWithdrawStatusV5RespData) GetWdIdOk() (*string, bool)`

GetWdIdOk returns a tuple with the WdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdId

`func (o *GetAssetDepositWithdrawStatusV5RespData) SetWdId(v string)`

SetWdId sets WdId field to given value.

### HasWdId

`func (o *GetAssetDepositWithdrawStatusV5RespData) HasWdId() bool`

HasWdId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


