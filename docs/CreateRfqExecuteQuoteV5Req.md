# CreateRfqExecuteQuoteV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Legs** | Pointer to [**[]CreateRfqExecuteQuoteV5ReqLegsInner**](CreateRfqExecuteQuoteV5ReqLegsInner.md) | An Array of objects containing the execution size of each leg of the RFQ.   The ratio of the leg sizes needs to be the same as the RFQ.   *Note: &#x60;tgtCcy&#x60; and &#x60;side&#x60; of each leg will be same as ones in the RFQ. &#x60;px&#x60; will be the same as the ones in the Quote. | [optional] 
**QuoteId** | **string** | Quote ID. | [default to ""]
**RfqId** | **string** | RFQ ID . | [default to ""]

## Methods

### NewCreateRfqExecuteQuoteV5Req

`func NewCreateRfqExecuteQuoteV5Req(quoteId string, rfqId string, ) *CreateRfqExecuteQuoteV5Req`

NewCreateRfqExecuteQuoteV5Req instantiates a new CreateRfqExecuteQuoteV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqExecuteQuoteV5ReqWithDefaults

`func NewCreateRfqExecuteQuoteV5ReqWithDefaults() *CreateRfqExecuteQuoteV5Req`

NewCreateRfqExecuteQuoteV5ReqWithDefaults instantiates a new CreateRfqExecuteQuoteV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegs

`func (o *CreateRfqExecuteQuoteV5Req) GetLegs() []CreateRfqExecuteQuoteV5ReqLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *CreateRfqExecuteQuoteV5Req) GetLegsOk() (*[]CreateRfqExecuteQuoteV5ReqLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *CreateRfqExecuteQuoteV5Req) SetLegs(v []CreateRfqExecuteQuoteV5ReqLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *CreateRfqExecuteQuoteV5Req) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetQuoteId

`func (o *CreateRfqExecuteQuoteV5Req) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *CreateRfqExecuteQuoteV5Req) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *CreateRfqExecuteQuoteV5Req) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetRfqId

`func (o *CreateRfqExecuteQuoteV5Req) GetRfqId() string`

GetRfqId returns the RfqId field if non-nil, zero value otherwise.

### GetRfqIdOk

`func (o *CreateRfqExecuteQuoteV5Req) GetRfqIdOk() (*string, bool)`

GetRfqIdOk returns a tuple with the RfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqId

`func (o *CreateRfqExecuteQuoteV5Req) SetRfqId(v string)`

SetRfqId sets RfqId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


