# GetRfqQuotesV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The result code, &#x60;0&#x60; means success. | [optional] [default to ""]
**Data** | Pointer to [**[]GetRfqQuotesV5RespDataInnerDataInner**](GetRfqQuotesV5RespDataInnerDataInner.md) | Array of objects containing the results of the Quote creation. | [optional] 
**Msg** | Pointer to **string** | The error message, not empty if the code is not 0. | [optional] [default to ""]

## Methods

### NewGetRfqQuotesV5RespDataInner

`func NewGetRfqQuotesV5RespDataInner() *GetRfqQuotesV5RespDataInner`

NewGetRfqQuotesV5RespDataInner instantiates a new GetRfqQuotesV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRfqQuotesV5RespDataInnerWithDefaults

`func NewGetRfqQuotesV5RespDataInnerWithDefaults() *GetRfqQuotesV5RespDataInner`

NewGetRfqQuotesV5RespDataInnerWithDefaults instantiates a new GetRfqQuotesV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetRfqQuotesV5RespDataInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetRfqQuotesV5RespDataInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetRfqQuotesV5RespDataInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetRfqQuotesV5RespDataInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *GetRfqQuotesV5RespDataInner) GetData() []GetRfqQuotesV5RespDataInnerDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRfqQuotesV5RespDataInner) GetDataOk() (*[]GetRfqQuotesV5RespDataInnerDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRfqQuotesV5RespDataInner) SetData(v []GetRfqQuotesV5RespDataInnerDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetRfqQuotesV5RespDataInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *GetRfqQuotesV5RespDataInner) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetRfqQuotesV5RespDataInner) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetRfqQuotesV5RespDataInner) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetRfqQuotesV5RespDataInner) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


