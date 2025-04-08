# CreateRfqExecuteQuoteV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The result code, &#x60;0&#x60; means success. | [optional] [default to ""]
**Data** | Pointer to [**[]CreateRfqExecuteQuoteV5RespDataDataInner**](CreateRfqExecuteQuoteV5RespDataDataInner.md) | Array of objects containing the results | [optional] 
**Msg** | Pointer to **string** | The error message, not empty if the code is not 0. | [optional] [default to ""]

## Methods

### NewCreateRfqExecuteQuoteV5RespData

`func NewCreateRfqExecuteQuoteV5RespData() *CreateRfqExecuteQuoteV5RespData`

NewCreateRfqExecuteQuoteV5RespData instantiates a new CreateRfqExecuteQuoteV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqExecuteQuoteV5RespDataWithDefaults

`func NewCreateRfqExecuteQuoteV5RespDataWithDefaults() *CreateRfqExecuteQuoteV5RespData`

NewCreateRfqExecuteQuoteV5RespDataWithDefaults instantiates a new CreateRfqExecuteQuoteV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateRfqExecuteQuoteV5RespData) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateRfqExecuteQuoteV5RespData) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateRfqExecuteQuoteV5RespData) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateRfqExecuteQuoteV5RespData) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *CreateRfqExecuteQuoteV5RespData) GetData() []CreateRfqExecuteQuoteV5RespDataDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateRfqExecuteQuoteV5RespData) GetDataOk() (*[]CreateRfqExecuteQuoteV5RespDataDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateRfqExecuteQuoteV5RespData) SetData(v []CreateRfqExecuteQuoteV5RespDataDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *CreateRfqExecuteQuoteV5RespData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CreateRfqExecuteQuoteV5RespData) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateRfqExecuteQuoteV5RespData) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateRfqExecuteQuoteV5RespData) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateRfqExecuteQuoteV5RespData) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


