# CreateTradeCancelBatchOrdersV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**SCode** | Pointer to **string** | The code of the event execution result, &#x60;0&#x60; means success. | [optional] [default to ""]
**SMsg** | Pointer to **string** | Rejection message if the request is unsuccessful. | [optional] [default to ""]
**Ts** | Pointer to **string** | Timestamp when the order request processing is finished by our system, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewCreateTradeCancelBatchOrdersV5RespDataInner

`func NewCreateTradeCancelBatchOrdersV5RespDataInner() *CreateTradeCancelBatchOrdersV5RespDataInner`

NewCreateTradeCancelBatchOrdersV5RespDataInner instantiates a new CreateTradeCancelBatchOrdersV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeCancelBatchOrdersV5RespDataInnerWithDefaults

`func NewCreateTradeCancelBatchOrdersV5RespDataInnerWithDefaults() *CreateTradeCancelBatchOrdersV5RespDataInner`

NewCreateTradeCancelBatchOrdersV5RespDataInnerWithDefaults instantiates a new CreateTradeCancelBatchOrdersV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClOrdId

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetOrdId

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetSCode

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetSCode() string`

GetSCode returns the SCode field if non-nil, zero value otherwise.

### GetSCodeOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetSCodeOk() (*string, bool)`

GetSCodeOk returns a tuple with the SCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCode

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) SetSCode(v string)`

SetSCode sets SCode field to given value.

### HasSCode

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) HasSCode() bool`

HasSCode returns a boolean if a field has been set.

### GetSMsg

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetSMsg() string`

GetSMsg returns the SMsg field if non-nil, zero value otherwise.

### GetSMsgOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetSMsgOk() (*string, bool)`

GetSMsgOk returns a tuple with the SMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMsg

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) SetSMsg(v string)`

SetSMsg sets SMsg field to given value.

### HasSMsg

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) HasSMsg() bool`

HasSMsg returns a boolean if a field has been set.

### GetTs

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CreateTradeCancelBatchOrdersV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


