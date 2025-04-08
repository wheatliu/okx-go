# CreateRfqCancelBatchQuotesV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClQuoteIds** | Pointer to **[]string** | Client-supplied Quote IDs. Either &#x60;quoteIds&#x60; or &#x60;clQuoteIds&#x60; is required.If both attributes are sent, &#x60;quoteIds&#x60; will be used as primary identifier. | [optional] 
**QuoteIds** | Pointer to **[]string** | Quote IDs . | [optional] 

## Methods

### NewCreateRfqCancelBatchQuotesV5Req

`func NewCreateRfqCancelBatchQuotesV5Req() *CreateRfqCancelBatchQuotesV5Req`

NewCreateRfqCancelBatchQuotesV5Req instantiates a new CreateRfqCancelBatchQuotesV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCancelBatchQuotesV5ReqWithDefaults

`func NewCreateRfqCancelBatchQuotesV5ReqWithDefaults() *CreateRfqCancelBatchQuotesV5Req`

NewCreateRfqCancelBatchQuotesV5ReqWithDefaults instantiates a new CreateRfqCancelBatchQuotesV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClQuoteIds

`func (o *CreateRfqCancelBatchQuotesV5Req) GetClQuoteIds() []string`

GetClQuoteIds returns the ClQuoteIds field if non-nil, zero value otherwise.

### GetClQuoteIdsOk

`func (o *CreateRfqCancelBatchQuotesV5Req) GetClQuoteIdsOk() (*[]string, bool)`

GetClQuoteIdsOk returns a tuple with the ClQuoteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClQuoteIds

`func (o *CreateRfqCancelBatchQuotesV5Req) SetClQuoteIds(v []string)`

SetClQuoteIds sets ClQuoteIds field to given value.

### HasClQuoteIds

`func (o *CreateRfqCancelBatchQuotesV5Req) HasClQuoteIds() bool`

HasClQuoteIds returns a boolean if a field has been set.

### GetQuoteIds

`func (o *CreateRfqCancelBatchQuotesV5Req) GetQuoteIds() []string`

GetQuoteIds returns the QuoteIds field if non-nil, zero value otherwise.

### GetQuoteIdsOk

`func (o *CreateRfqCancelBatchQuotesV5Req) GetQuoteIdsOk() (*[]string, bool)`

GetQuoteIdsOk returns a tuple with the QuoteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteIds

`func (o *CreateRfqCancelBatchQuotesV5Req) SetQuoteIds(v []string)`

SetQuoteIds sets QuoteIds field to given value.

### HasQuoteIds

`func (o *CreateRfqCancelBatchQuotesV5Req) HasQuoteIds() bool`

HasQuoteIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


