# GetAccountPositionsV5Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] [default to ""]
**Data** | Pointer to [**[]GetAccountPositionsV5RespDataInner**](GetAccountPositionsV5RespDataInner.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewGetAccountPositionsV5Resp

`func NewGetAccountPositionsV5Resp() *GetAccountPositionsV5Resp`

NewGetAccountPositionsV5Resp instantiates a new GetAccountPositionsV5Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountPositionsV5RespWithDefaults

`func NewGetAccountPositionsV5RespWithDefaults() *GetAccountPositionsV5Resp`

NewGetAccountPositionsV5RespWithDefaults instantiates a new GetAccountPositionsV5Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetAccountPositionsV5Resp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetAccountPositionsV5Resp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetAccountPositionsV5Resp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetAccountPositionsV5Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetAccountPositionsV5Resp) GetData() []GetAccountPositionsV5RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAccountPositionsV5Resp) GetDataOk() (*[]GetAccountPositionsV5RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAccountPositionsV5Resp) SetData(v []GetAccountPositionsV5RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetAccountPositionsV5Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *GetAccountPositionsV5Resp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetAccountPositionsV5Resp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetAccountPositionsV5Resp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetAccountPositionsV5Resp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


