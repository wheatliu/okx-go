# CreateSprdAmendOrderV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client | [optional] [default to ""]
**NewPx** | Pointer to **string** | New price after amendment   Either &#x60;newSz&#x60; or &#x60;newPx&#x60; is required. | [optional] [default to ""]
**NewSz** | Pointer to **string** | New quantity after amendment   Either &#x60;newSz&#x60; or &#x60;newPx&#x60; is required.   When amending a partially-filled order, the newSz should include the amount that has been filled. | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID   Either &#x60;ordId&#x60; or &#x60;clOrdId&#x60; is required. If both are passed, ordId will be used. | [optional] [default to ""]
**ReqId** | Pointer to **string** | Client Request ID as assigned by the client for order amendment   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.   The response will include the corresponding reqId to help you identify the request if you provide it in the request. | [optional] [default to ""]

## Methods

### NewCreateSprdAmendOrderV5RespData

`func NewCreateSprdAmendOrderV5RespData() *CreateSprdAmendOrderV5RespData`

NewCreateSprdAmendOrderV5RespData instantiates a new CreateSprdAmendOrderV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSprdAmendOrderV5RespDataWithDefaults

`func NewCreateSprdAmendOrderV5RespDataWithDefaults() *CreateSprdAmendOrderV5RespData`

NewCreateSprdAmendOrderV5RespDataWithDefaults instantiates a new CreateSprdAmendOrderV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClOrdId

`func (o *CreateSprdAmendOrderV5RespData) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *CreateSprdAmendOrderV5RespData) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *CreateSprdAmendOrderV5RespData) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *CreateSprdAmendOrderV5RespData) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetNewPx

`func (o *CreateSprdAmendOrderV5RespData) GetNewPx() string`

GetNewPx returns the NewPx field if non-nil, zero value otherwise.

### GetNewPxOk

`func (o *CreateSprdAmendOrderV5RespData) GetNewPxOk() (*string, bool)`

GetNewPxOk returns a tuple with the NewPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPx

`func (o *CreateSprdAmendOrderV5RespData) SetNewPx(v string)`

SetNewPx sets NewPx field to given value.

### HasNewPx

`func (o *CreateSprdAmendOrderV5RespData) HasNewPx() bool`

HasNewPx returns a boolean if a field has been set.

### GetNewSz

`func (o *CreateSprdAmendOrderV5RespData) GetNewSz() string`

GetNewSz returns the NewSz field if non-nil, zero value otherwise.

### GetNewSzOk

`func (o *CreateSprdAmendOrderV5RespData) GetNewSzOk() (*string, bool)`

GetNewSzOk returns a tuple with the NewSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSz

`func (o *CreateSprdAmendOrderV5RespData) SetNewSz(v string)`

SetNewSz sets NewSz field to given value.

### HasNewSz

`func (o *CreateSprdAmendOrderV5RespData) HasNewSz() bool`

HasNewSz returns a boolean if a field has been set.

### GetOrdId

`func (o *CreateSprdAmendOrderV5RespData) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *CreateSprdAmendOrderV5RespData) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *CreateSprdAmendOrderV5RespData) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *CreateSprdAmendOrderV5RespData) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetReqId

`func (o *CreateSprdAmendOrderV5RespData) GetReqId() string`

GetReqId returns the ReqId field if non-nil, zero value otherwise.

### GetReqIdOk

`func (o *CreateSprdAmendOrderV5RespData) GetReqIdOk() (*string, bool)`

GetReqIdOk returns a tuple with the ReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqId

`func (o *CreateSprdAmendOrderV5RespData) SetReqId(v string)`

SetReqId sets ReqId field to given value.

### HasReqId

`func (o *CreateSprdAmendOrderV5RespData) HasReqId() bool`

HasReqId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


