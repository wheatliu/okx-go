# GetAssetDepositWithdrawStatusV5Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] [default to ""]
**Data** | Pointer to [**[]GetAssetDepositWithdrawStatusV5RespDataInner**](GetAssetDepositWithdrawStatusV5RespDataInner.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewGetAssetDepositWithdrawStatusV5Resp

`func NewGetAssetDepositWithdrawStatusV5Resp() *GetAssetDepositWithdrawStatusV5Resp`

NewGetAssetDepositWithdrawStatusV5Resp instantiates a new GetAssetDepositWithdrawStatusV5Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetDepositWithdrawStatusV5RespWithDefaults

`func NewGetAssetDepositWithdrawStatusV5RespWithDefaults() *GetAssetDepositWithdrawStatusV5Resp`

NewGetAssetDepositWithdrawStatusV5RespWithDefaults instantiates a new GetAssetDepositWithdrawStatusV5Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetAssetDepositWithdrawStatusV5Resp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetAssetDepositWithdrawStatusV5Resp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetAssetDepositWithdrawStatusV5Resp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetAssetDepositWithdrawStatusV5Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetAssetDepositWithdrawStatusV5Resp) GetData() []GetAssetDepositWithdrawStatusV5RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAssetDepositWithdrawStatusV5Resp) GetDataOk() (*[]GetAssetDepositWithdrawStatusV5RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAssetDepositWithdrawStatusV5Resp) SetData(v []GetAssetDepositWithdrawStatusV5RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetAssetDepositWithdrawStatusV5Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *GetAssetDepositWithdrawStatusV5Resp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetAssetDepositWithdrawStatusV5Resp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetAssetDepositWithdrawStatusV5Resp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetAssetDepositWithdrawStatusV5Resp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


