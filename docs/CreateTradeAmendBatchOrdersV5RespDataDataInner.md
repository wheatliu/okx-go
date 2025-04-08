# CreateTradeAmendBatchOrdersV5RespDataDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**ReqId** | Pointer to **string** | Client Request ID as assigned by the client for order amendment. | [optional] [default to ""]
**SCode** | Pointer to **string** | The code of the event execution result, &#x60;0&#x60; means success. | [optional] [default to ""]
**SMsg** | Pointer to **string** | Rejection message if the request is unsuccessful. | [optional] [default to ""]
**Ts** | Pointer to **string** | Timestamp when the order request processing is finished by our system, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewCreateTradeAmendBatchOrdersV5RespDataDataInner

`func NewCreateTradeAmendBatchOrdersV5RespDataDataInner() *CreateTradeAmendBatchOrdersV5RespDataDataInner`

NewCreateTradeAmendBatchOrdersV5RespDataDataInner instantiates a new CreateTradeAmendBatchOrdersV5RespDataDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeAmendBatchOrdersV5RespDataDataInnerWithDefaults

`func NewCreateTradeAmendBatchOrdersV5RespDataDataInnerWithDefaults() *CreateTradeAmendBatchOrdersV5RespDataDataInner`

NewCreateTradeAmendBatchOrdersV5RespDataDataInnerWithDefaults instantiates a new CreateTradeAmendBatchOrdersV5RespDataDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClOrdId

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetOrdId

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetReqId

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) GetReqId() string`

GetReqId returns the ReqId field if non-nil, zero value otherwise.

### GetReqIdOk

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) GetReqIdOk() (*string, bool)`

GetReqIdOk returns a tuple with the ReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqId

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) SetReqId(v string)`

SetReqId sets ReqId field to given value.

### HasReqId

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) HasReqId() bool`

HasReqId returns a boolean if a field has been set.

### GetSCode

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) GetSCode() string`

GetSCode returns the SCode field if non-nil, zero value otherwise.

### GetSCodeOk

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) GetSCodeOk() (*string, bool)`

GetSCodeOk returns a tuple with the SCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCode

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) SetSCode(v string)`

SetSCode sets SCode field to given value.

### HasSCode

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) HasSCode() bool`

HasSCode returns a boolean if a field has been set.

### GetSMsg

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) GetSMsg() string`

GetSMsg returns the SMsg field if non-nil, zero value otherwise.

### GetSMsgOk

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) GetSMsgOk() (*string, bool)`

GetSMsgOk returns a tuple with the SMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMsg

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) SetSMsg(v string)`

SetSMsg sets SMsg field to given value.

### HasSMsg

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) HasSMsg() bool`

HasSMsg returns a boolean if a field has been set.

### GetTs

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CreateTradeAmendBatchOrdersV5RespDataDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


