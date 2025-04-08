# CreateTradeAmendOrderV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachAlgoOrds** | Pointer to [**[]CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner**](CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner.md) | TP/SL information attached when placing order | [optional] 
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**CxlOnFail** | Pointer to **bool** | Whether the order needs to be automatically canceled when the order amendment fails    Valid options: &#x60;false&#x60; or &#x60;true&#x60;, the default is &#x60;false&#x60;. | [optional] 
**InstId** | **string** | Instrument ID | [default to ""]
**NewPx** | Pointer to **string** | New price after amendment.   When modifying options orders, users can only fill in one of the following: newPx, newPxUsd, or newPxVol. It must be consistent with parameters when placing orders. For example, if users placed the order using px, they should use newPx when modifying the order. | [optional] [default to ""]
**NewPxUsd** | Pointer to **string** | Modify options orders using USD prices   Only applicable to options.   When modifying options orders, users can only fill in one of the following: newPx, newPxUsd, or newPxVol. | [optional] [default to ""]
**NewPxVol** | Pointer to **string** | Modify options orders based on implied volatility, where 1 represents 100%   Only applicable to options.   When modifying options orders, users can only fill in one of the following: newPx, newPxUsd, or newPxVol. | [optional] [default to ""]
**NewSz** | Pointer to **string** | New quantity after amendment and it has to be larger than 0. When amending a partially-filled order, the &#x60;newSz&#x60; should include the amount that has been filled. | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID    Either &#x60;ordId&#x60; or &#x60;clOrdId&#x60; is required. If both are passed, &#x60;ordId&#x60; will be used. | [optional] [default to ""]
**ReqId** | Pointer to **string** | Client Request ID as assigned by the client for order amendment   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.   The response will include the corresponding &#x60;reqId&#x60; to help you identify the request if you provide it in the request. | [optional] [default to ""]

## Methods

### NewCreateTradeAmendOrderV5Req

`func NewCreateTradeAmendOrderV5Req(instId string, ) *CreateTradeAmendOrderV5Req`

NewCreateTradeAmendOrderV5Req instantiates a new CreateTradeAmendOrderV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeAmendOrderV5ReqWithDefaults

`func NewCreateTradeAmendOrderV5ReqWithDefaults() *CreateTradeAmendOrderV5Req`

NewCreateTradeAmendOrderV5ReqWithDefaults instantiates a new CreateTradeAmendOrderV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachAlgoOrds

`func (o *CreateTradeAmendOrderV5Req) GetAttachAlgoOrds() []CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner`

GetAttachAlgoOrds returns the AttachAlgoOrds field if non-nil, zero value otherwise.

### GetAttachAlgoOrdsOk

`func (o *CreateTradeAmendOrderV5Req) GetAttachAlgoOrdsOk() (*[]CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner, bool)`

GetAttachAlgoOrdsOk returns a tuple with the AttachAlgoOrds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlgoOrds

`func (o *CreateTradeAmendOrderV5Req) SetAttachAlgoOrds(v []CreateTradeAmendBatchOrdersV5ReqAttachAlgoOrdsInner)`

SetAttachAlgoOrds sets AttachAlgoOrds field to given value.

### HasAttachAlgoOrds

`func (o *CreateTradeAmendOrderV5Req) HasAttachAlgoOrds() bool`

HasAttachAlgoOrds returns a boolean if a field has been set.

### GetClOrdId

`func (o *CreateTradeAmendOrderV5Req) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateTradeAmendOrderV5Req) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateTradeAmendOrderV5Req) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateTradeAmendOrderV5Req) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetCxlOnFail

`func (o *CreateTradeAmendOrderV5Req) GetCxlOnFail() bool`

GetCxlOnFail returns the CxlOnFail field if non-nil, zero value otherwise.

### GetCxlOnFailOk

`func (o *CreateTradeAmendOrderV5Req) GetCxlOnFailOk() (*bool, bool)`

GetCxlOnFailOk returns a tuple with the CxlOnFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCxlOnFail

`func (o *CreateTradeAmendOrderV5Req) SetCxlOnFail(v bool)`

SetCxlOnFail sets CxlOnFail field to given value.

### HasCxlOnFail

`func (o *CreateTradeAmendOrderV5Req) HasCxlOnFail() bool`

HasCxlOnFail returns a boolean if a field has been set.

### GetInstId

`func (o *CreateTradeAmendOrderV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradeAmendOrderV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradeAmendOrderV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetNewPx

`func (o *CreateTradeAmendOrderV5Req) GetNewPx() string`

GetNewPx returns the NewPx field if non-nil, zero value otherwise.

### GetNewPxOk

`func (o *CreateTradeAmendOrderV5Req) GetNewPxOk() (*string, bool)`

GetNewPxOk returns a tuple with the NewPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPx

`func (o *CreateTradeAmendOrderV5Req) SetNewPx(v string)`

SetNewPx sets NewPx field to given value.

### HasNewPx

`func (o *CreateTradeAmendOrderV5Req) HasNewPx() bool`

HasNewPx returns a boolean if a field has been set.

### GetNewPxUsd

`func (o *CreateTradeAmendOrderV5Req) GetNewPxUsd() string`

GetNewPxUsd returns the NewPxUsd field if non-nil, zero value otherwise.

### GetNewPxUsdOk

`func (o *CreateTradeAmendOrderV5Req) GetNewPxUsdOk() (*string, bool)`

GetNewPxUsdOk returns a tuple with the NewPxUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPxUsd

`func (o *CreateTradeAmendOrderV5Req) SetNewPxUsd(v string)`

SetNewPxUsd sets NewPxUsd field to given value.

### HasNewPxUsd

`func (o *CreateTradeAmendOrderV5Req) HasNewPxUsd() bool`

HasNewPxUsd returns a boolean if a field has been set.

### GetNewPxVol

`func (o *CreateTradeAmendOrderV5Req) GetNewPxVol() string`

GetNewPxVol returns the NewPxVol field if non-nil, zero value otherwise.

### GetNewPxVolOk

`func (o *CreateTradeAmendOrderV5Req) GetNewPxVolOk() (*string, bool)`

GetNewPxVolOk returns a tuple with the NewPxVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPxVol

`func (o *CreateTradeAmendOrderV5Req) SetNewPxVol(v string)`

SetNewPxVol sets NewPxVol field to given value.

### HasNewPxVol

`func (o *CreateTradeAmendOrderV5Req) HasNewPxVol() bool`

HasNewPxVol returns a boolean if a field has been set.

### GetNewSz

`func (o *CreateTradeAmendOrderV5Req) GetNewSz() string`

GetNewSz returns the NewSz field if non-nil, zero value otherwise.

### GetNewSzOk

`func (o *CreateTradeAmendOrderV5Req) GetNewSzOk() (*string, bool)`

GetNewSzOk returns a tuple with the NewSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSz

`func (o *CreateTradeAmendOrderV5Req) SetNewSz(v string)`

SetNewSz sets NewSz field to given value.

### HasNewSz

`func (o *CreateTradeAmendOrderV5Req) HasNewSz() bool`

HasNewSz returns a boolean if a field has been set.

### GetOrdId

`func (o *CreateTradeAmendOrderV5Req) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateTradeAmendOrderV5Req) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateTradeAmendOrderV5Req) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *CreateTradeAmendOrderV5Req) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetReqId

`func (o *CreateTradeAmendOrderV5Req) GetReqId() string`

GetReqId returns the ReqId field if non-nil, zero value otherwise.

### GetReqIdOk

`func (o *CreateTradeAmendOrderV5Req) GetReqIdOk() (*string, bool)`

GetReqIdOk returns a tuple with the ReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqId

`func (o *CreateTradeAmendOrderV5Req) SetReqId(v string)`

SetReqId sets ReqId field to given value.

### HasReqId

`func (o *CreateTradeAmendOrderV5Req) HasReqId() bool`

HasReqId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


