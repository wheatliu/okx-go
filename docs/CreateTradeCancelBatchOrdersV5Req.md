# CreateTradeCancelBatchOrdersV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**InstId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to ""]
**OrdId** | Pointer to **string** | Order ID  Either &#x60;ordId&#x60; or &#x60;clOrdId&#x60; is required. If both are passed, &#x60;ordId&#x60; will be used. | [optional] [default to ""]

## Methods

### NewCreateTradeCancelBatchOrdersV5Req

`func NewCreateTradeCancelBatchOrdersV5Req(instId string, ) *CreateTradeCancelBatchOrdersV5Req`

NewCreateTradeCancelBatchOrdersV5Req instantiates a new CreateTradeCancelBatchOrdersV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeCancelBatchOrdersV5ReqWithDefaults

`func NewCreateTradeCancelBatchOrdersV5ReqWithDefaults() *CreateTradeCancelBatchOrdersV5Req`

NewCreateTradeCancelBatchOrdersV5ReqWithDefaults instantiates a new CreateTradeCancelBatchOrdersV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClOrdId

`func (o *CreateTradeCancelBatchOrdersV5Req) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateTradeCancelBatchOrdersV5Req) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateTradeCancelBatchOrdersV5Req) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateTradeCancelBatchOrdersV5Req) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetInstId

`func (o *CreateTradeCancelBatchOrdersV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradeCancelBatchOrdersV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradeCancelBatchOrdersV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetOrdId

`func (o *CreateTradeCancelBatchOrdersV5Req) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateTradeCancelBatchOrdersV5Req) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateTradeCancelBatchOrdersV5Req) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *CreateTradeCancelBatchOrdersV5Req) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


