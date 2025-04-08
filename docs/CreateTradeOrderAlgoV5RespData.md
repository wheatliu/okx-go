# CreateTradeOrderAlgoV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID | [optional] [default to ""]
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client(Deprecated) | [optional] [default to ""]
**SCode** | Pointer to **string** | The code of the event execution result, &#x60;0&#x60; means success. | [optional] [default to ""]
**SMsg** | Pointer to **string** | Rejection message if the request is unsuccessful. | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]

## Methods

### NewCreateTradeOrderAlgoV5RespData

`func NewCreateTradeOrderAlgoV5RespData() *CreateTradeOrderAlgoV5RespData`

NewCreateTradeOrderAlgoV5RespData instantiates a new CreateTradeOrderAlgoV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeOrderAlgoV5RespDataWithDefaults

`func NewCreateTradeOrderAlgoV5RespDataWithDefaults() *CreateTradeOrderAlgoV5RespData`

NewCreateTradeOrderAlgoV5RespDataWithDefaults instantiates a new CreateTradeOrderAlgoV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoClOrdId

`func (o *CreateTradeOrderAlgoV5RespData) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *CreateTradeOrderAlgoV5RespData) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *CreateTradeOrderAlgoV5RespData) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *CreateTradeOrderAlgoV5RespData) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *CreateTradeOrderAlgoV5RespData) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradeOrderAlgoV5RespData) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradeOrderAlgoV5RespData) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *CreateTradeOrderAlgoV5RespData) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetClOrdId

`func (o *CreateTradeOrderAlgoV5RespData) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateTradeOrderAlgoV5RespData) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateTradeOrderAlgoV5RespData) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateTradeOrderAlgoV5RespData) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetSCode

`func (o *CreateTradeOrderAlgoV5RespData) GetSCode() string`

GetSCode returns the SCode field if non-nil, zero value otherwise.

### GetSCodeOk

`func (o *CreateTradeOrderAlgoV5RespData) GetSCodeOk() (*string, bool)`

GetSCodeOk returns a tuple with the SCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCode

`func (o *CreateTradeOrderAlgoV5RespData) SetSCode(v string)`

SetSCode sets SCode field to given value.

### HasSCode

`func (o *CreateTradeOrderAlgoV5RespData) HasSCode() bool`

HasSCode returns a boolean if a field has been set.

### GetSMsg

`func (o *CreateTradeOrderAlgoV5RespData) GetSMsg() string`

GetSMsg returns the SMsg field if non-nil, zero value otherwise.

### GetSMsgOk

`func (o *CreateTradeOrderAlgoV5RespData) GetSMsgOk() (*string, bool)`

GetSMsgOk returns a tuple with the SMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMsg

`func (o *CreateTradeOrderAlgoV5RespData) SetSMsg(v string)`

SetSMsg sets SMsg field to given value.

### HasSMsg

`func (o *CreateTradeOrderAlgoV5RespData) HasSMsg() bool`

HasSMsg returns a boolean if a field has been set.

### GetTag

`func (o *CreateTradeOrderAlgoV5RespData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateTradeOrderAlgoV5RespData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateTradeOrderAlgoV5RespData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateTradeOrderAlgoV5RespData) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


