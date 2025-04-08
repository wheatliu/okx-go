# GetAccountInstrumentsV5Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] [default to ""]
**Data** | Pointer to [**[]GetAccountInstrumentsV5RespDataInner**](GetAccountInstrumentsV5RespDataInner.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewGetAccountInstrumentsV5Resp

`func NewGetAccountInstrumentsV5Resp() *GetAccountInstrumentsV5Resp`

NewGetAccountInstrumentsV5Resp instantiates a new GetAccountInstrumentsV5Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountInstrumentsV5RespWithDefaults

`func NewGetAccountInstrumentsV5RespWithDefaults() *GetAccountInstrumentsV5Resp`

NewGetAccountInstrumentsV5RespWithDefaults instantiates a new GetAccountInstrumentsV5Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetAccountInstrumentsV5Resp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetAccountInstrumentsV5Resp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetAccountInstrumentsV5Resp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetAccountInstrumentsV5Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetAccountInstrumentsV5Resp) GetData() []GetAccountInstrumentsV5RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAccountInstrumentsV5Resp) GetDataOk() (*[]GetAccountInstrumentsV5RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAccountInstrumentsV5Resp) SetData(v []GetAccountInstrumentsV5RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetAccountInstrumentsV5Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *GetAccountInstrumentsV5Resp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetAccountInstrumentsV5Resp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetAccountInstrumentsV5Resp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetAccountInstrumentsV5Resp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


