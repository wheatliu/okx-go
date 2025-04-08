# CreateTradeBatchOrdersV5RespDataInnerDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID | [optional] [default to ""]
**SCode** | Pointer to **string** | The code of the event execution result, &#x60;0&#x60; means success. | [optional] [default to ""]
**SMsg** | Pointer to **string** | Rejection or success message of event execution. | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**Ts** | Pointer to **string** | Timestamp when the order request processing is finished by our system, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewCreateTradeBatchOrdersV5RespDataInnerDataInner

`func NewCreateTradeBatchOrdersV5RespDataInnerDataInner() *CreateTradeBatchOrdersV5RespDataInnerDataInner`

NewCreateTradeBatchOrdersV5RespDataInnerDataInner instantiates a new CreateTradeBatchOrdersV5RespDataInnerDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeBatchOrdersV5RespDataInnerDataInnerWithDefaults

`func NewCreateTradeBatchOrdersV5RespDataInnerDataInnerWithDefaults() *CreateTradeBatchOrdersV5RespDataInnerDataInner`

NewCreateTradeBatchOrdersV5RespDataInnerDataInnerWithDefaults instantiates a new CreateTradeBatchOrdersV5RespDataInnerDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClOrdId

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetOrdId

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetSCode

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) GetSCode() string`

GetSCode returns the SCode field if non-nil, zero value otherwise.

### GetSCodeOk

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) GetSCodeOk() (*string, bool)`

GetSCodeOk returns a tuple with the SCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCode

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) SetSCode(v string)`

SetSCode sets SCode field to given value.

### HasSCode

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) HasSCode() bool`

HasSCode returns a boolean if a field has been set.

### GetSMsg

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) GetSMsg() string`

GetSMsg returns the SMsg field if non-nil, zero value otherwise.

### GetSMsgOk

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) GetSMsgOk() (*string, bool)`

GetSMsgOk returns a tuple with the SMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMsg

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) SetSMsg(v string)`

SetSMsg sets SMsg field to given value.

### HasSMsg

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) HasSMsg() bool`

HasSMsg returns a boolean if a field has been set.

### GetTag

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTs

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *CreateTradeBatchOrdersV5RespDataInnerDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


