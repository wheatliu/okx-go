# CreateTradeCancelAdvanceAlgosV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | Pointer to **string** | Order ID | [optional] [default to ""]
**SCode** | Pointer to **string** | The code of the event execution result, &#x60;0&#x60; means success. | [optional] [default to ""]
**SMsg** | Pointer to **string** | Rejection message if the request is unsuccessful. | [optional] [default to ""]

## Methods

### NewCreateTradeCancelAdvanceAlgosV5RespData

`func NewCreateTradeCancelAdvanceAlgosV5RespData() *CreateTradeCancelAdvanceAlgosV5RespData`

NewCreateTradeCancelAdvanceAlgosV5RespData instantiates a new CreateTradeCancelAdvanceAlgosV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeCancelAdvanceAlgosV5RespDataWithDefaults

`func NewCreateTradeCancelAdvanceAlgosV5RespDataWithDefaults() *CreateTradeCancelAdvanceAlgosV5RespData`

NewCreateTradeCancelAdvanceAlgosV5RespDataWithDefaults instantiates a new CreateTradeCancelAdvanceAlgosV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *CreateTradeCancelAdvanceAlgosV5RespData) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradeCancelAdvanceAlgosV5RespData) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradeCancelAdvanceAlgosV5RespData) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *CreateTradeCancelAdvanceAlgosV5RespData) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetSCode

`func (o *CreateTradeCancelAdvanceAlgosV5RespData) GetSCode() string`

GetSCode returns the SCode field if non-nil, zero value otherwise.

### GetSCodeOk

`func (o *CreateTradeCancelAdvanceAlgosV5RespData) GetSCodeOk() (*string, bool)`

GetSCodeOk returns a tuple with the SCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCode

`func (o *CreateTradeCancelAdvanceAlgosV5RespData) SetSCode(v string)`

SetSCode sets SCode field to given value.

### HasSCode

`func (o *CreateTradeCancelAdvanceAlgosV5RespData) HasSCode() bool`

HasSCode returns a boolean if a field has been set.

### GetSMsg

`func (o *CreateTradeCancelAdvanceAlgosV5RespData) GetSMsg() string`

GetSMsg returns the SMsg field if non-nil, zero value otherwise.

### GetSMsgOk

`func (o *CreateTradeCancelAdvanceAlgosV5RespData) GetSMsgOk() (*string, bool)`

GetSMsgOk returns a tuple with the SMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMsg

`func (o *CreateTradeCancelAdvanceAlgosV5RespData) SetSMsg(v string)`

SetSMsg sets SMsg field to given value.

### HasSMsg

`func (o *CreateTradeCancelAdvanceAlgosV5RespData) HasSMsg() bool`

HasSMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


