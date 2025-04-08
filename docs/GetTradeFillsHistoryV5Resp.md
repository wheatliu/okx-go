# GetTradeFillsHistoryV5Resp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] [default to ""]
**Data** | Pointer to [**[]GetTradeFillsHistoryV5RespDataInner**](GetTradeFillsHistoryV5RespDataInner.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewGetTradeFillsHistoryV5Resp

`func NewGetTradeFillsHistoryV5Resp() *GetTradeFillsHistoryV5Resp`

NewGetTradeFillsHistoryV5Resp instantiates a new GetTradeFillsHistoryV5Resp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradeFillsHistoryV5RespWithDefaults

`func NewGetTradeFillsHistoryV5RespWithDefaults() *GetTradeFillsHistoryV5Resp`

NewGetTradeFillsHistoryV5RespWithDefaults instantiates a new GetTradeFillsHistoryV5Resp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetTradeFillsHistoryV5Resp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetTradeFillsHistoryV5Resp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetTradeFillsHistoryV5Resp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetTradeFillsHistoryV5Resp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetTradeFillsHistoryV5Resp) GetData() []GetTradeFillsHistoryV5RespDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTradeFillsHistoryV5Resp) GetDataOk() (*[]GetTradeFillsHistoryV5RespDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTradeFillsHistoryV5Resp) SetData(v []GetTradeFillsHistoryV5RespDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetTradeFillsHistoryV5Resp) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *GetTradeFillsHistoryV5Resp) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetTradeFillsHistoryV5Resp) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetTradeFillsHistoryV5Resp) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetTradeFillsHistoryV5Resp) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


