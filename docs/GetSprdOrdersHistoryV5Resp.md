# GetSprdOrdersHistoryV5Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] [default to ""]
**Data** | Pointer to [**[]GetSprdOrdersHistoryArchiveV5RespDataInner**](GetSprdOrdersHistoryArchiveV5RespDataInner.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewGetSprdOrdersHistoryV5Resp

`func NewGetSprdOrdersHistoryV5Resp() *GetSprdOrdersHistoryV5Resp`

NewGetSprdOrdersHistoryV5Resp instantiates a new GetSprdOrdersHistoryV5Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSprdOrdersHistoryV5RespWithDefaults

`func NewGetSprdOrdersHistoryV5RespWithDefaults() *GetSprdOrdersHistoryV5Resp`

NewGetSprdOrdersHistoryV5RespWithDefaults instantiates a new GetSprdOrdersHistoryV5Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetSprdOrdersHistoryV5Resp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetSprdOrdersHistoryV5Resp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetSprdOrdersHistoryV5Resp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetSprdOrdersHistoryV5Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetSprdOrdersHistoryV5Resp) GetData() []GetSprdOrdersHistoryArchiveV5RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSprdOrdersHistoryV5Resp) GetDataOk() (*[]GetSprdOrdersHistoryArchiveV5RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSprdOrdersHistoryV5Resp) SetData(v []GetSprdOrdersHistoryArchiveV5RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetSprdOrdersHistoryV5Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *GetSprdOrdersHistoryV5Resp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetSprdOrdersHistoryV5Resp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetSprdOrdersHistoryV5Resp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetSprdOrdersHistoryV5Resp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


