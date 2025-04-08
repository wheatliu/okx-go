# CreateTradeAmendAlgosV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID  Either &#x60;algoId&#x60; or &#x60;algoClOrdId&#x60; is required. If both are passed, &#x60;algoId&#x60; will be used. | [optional] [default to ""]
**AlgoId** | Pointer to **string** | Algo ID  Either &#x60;algoId&#x60; or &#x60;algoClOrdId&#x60; is required. If both are passed, &#x60;algoId&#x60; will be used. | [optional] [default to ""]
**CxlOnFail** | Pointer to **bool** | Whether the order needs to be automatically canceled when the order amendment fails    Valid options: &#x60;false&#x60; or &#x60;true&#x60;, the default is &#x60;false&#x60;. | [optional] 
**InstId** | **string** | Instrument ID | [default to ""]
**NewSz** | Pointer to **string** | New quantity after amendment and it has to be larger than 0. | [optional] [default to ""]
**ReqId** | Pointer to **string** | Client Request ID as assigned by the client for order amendment   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.   The response will include the corresponding &#x60;reqId&#x60; to help you identify the request if you provide it in the request. | [optional] [default to ""]

## Methods

### NewCreateTradeAmendAlgosV5Req

`func NewCreateTradeAmendAlgosV5Req(instId string, ) *CreateTradeAmendAlgosV5Req`

NewCreateTradeAmendAlgosV5Req instantiates a new CreateTradeAmendAlgosV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeAmendAlgosV5ReqWithDefaults

`func NewCreateTradeAmendAlgosV5ReqWithDefaults() *CreateTradeAmendAlgosV5Req`

NewCreateTradeAmendAlgosV5ReqWithDefaults instantiates a new CreateTradeAmendAlgosV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoClOrdId

`func (o *CreateTradeAmendAlgosV5Req) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *CreateTradeAmendAlgosV5Req) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *CreateTradeAmendAlgosV5Req) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *CreateTradeAmendAlgosV5Req) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAlgoId

`func (o *CreateTradeAmendAlgosV5Req) GetAlgoId() string`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *CreateTradeAmendAlgosV5Req) GetAlgoIdOk() (*string, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *CreateTradeAmendAlgosV5Req) SetAlgoId(v string)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *CreateTradeAmendAlgosV5Req) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetCxlOnFail

`func (o *CreateTradeAmendAlgosV5Req) GetCxlOnFail() bool`

GetCxlOnFail returns the CxlOnFail field if non-nil, zero value otherwise.

### GetCxlOnFailOk

`func (o *CreateTradeAmendAlgosV5Req) GetCxlOnFailOk() (*bool, bool)`

GetCxlOnFailOk returns a tuple with the CxlOnFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCxlOnFail

`func (o *CreateTradeAmendAlgosV5Req) SetCxlOnFail(v bool)`

SetCxlOnFail sets CxlOnFail field to given value.

### HasCxlOnFail

`func (o *CreateTradeAmendAlgosV5Req) HasCxlOnFail() bool`

HasCxlOnFail returns a boolean if a field has been set.

### GetInstId

`func (o *CreateTradeAmendAlgosV5Req) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateTradeAmendAlgosV5Req) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateTradeAmendAlgosV5Req) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetNewSz

`func (o *CreateTradeAmendAlgosV5Req) GetNewSz() string`

GetNewSz returns the NewSz field if non-nil, zero value otherwise.

### GetNewSzOk

`func (o *CreateTradeAmendAlgosV5Req) GetNewSzOk() (*string, bool)`

GetNewSzOk returns a tuple with the NewSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSz

`func (o *CreateTradeAmendAlgosV5Req) SetNewSz(v string)`

SetNewSz sets NewSz field to given value.

### HasNewSz

`func (o *CreateTradeAmendAlgosV5Req) HasNewSz() bool`

HasNewSz returns a boolean if a field has been set.

### GetReqId

`func (o *CreateTradeAmendAlgosV5Req) GetReqId() string`

GetReqId returns the ReqId field if non-nil, zero value otherwise.

### GetReqIdOk

`func (o *CreateTradeAmendAlgosV5Req) GetReqIdOk() (*string, bool)`

GetReqIdOk returns a tuple with the ReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqId

`func (o *CreateTradeAmendAlgosV5Req) SetReqId(v string)`

SetReqId sets ReqId field to given value.

### HasReqId

`func (o *CreateTradeAmendAlgosV5Req) HasReqId() bool`

HasReqId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


