# CreateRfqCancelQuoteV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClQuoteId** | Pointer to **string** | Client-supplied Quote ID. Either &#x60;quoteId&#x60; or &#x60;clQuoteId&#x60; is required. If both &#x60;clQuoteId&#x60; and &#x60;quoteId&#x60; are passed, &#x60;quoteId&#x60; will be treated as primary identifier. | [optional] [default to ""]
**QuoteId** | Pointer to **string** | Quote ID. | [optional] [default to ""]
**RfqId** | Pointer to **string** | RFQ ID. | [optional] [default to ""]

## Methods

### NewCreateRfqCancelQuoteV5Req

`func NewCreateRfqCancelQuoteV5Req() *CreateRfqCancelQuoteV5Req`

NewCreateRfqCancelQuoteV5Req instantiates a new CreateRfqCancelQuoteV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCancelQuoteV5ReqWithDefaults

`func NewCreateRfqCancelQuoteV5ReqWithDefaults() *CreateRfqCancelQuoteV5Req`

NewCreateRfqCancelQuoteV5ReqWithDefaults instantiates a new CreateRfqCancelQuoteV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClQuoteId

`func (o *CreateRfqCancelQuoteV5Req) GetClQuoteId() string`

GetClQuoteId returns the ClQuoteId field if non-nil, zero value otherwise.

### GetClQuoteIdOk

`func (o *CreateRfqCancelQuoteV5Req) GetClQuoteIdOk() (*string, bool)`

GetClQuoteIdOk returns a tuple with the ClQuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClQuoteId

`func (o *CreateRfqCancelQuoteV5Req) SetClQuoteId(v string)`

SetClQuoteId sets ClQuoteId field to given value.

### HasClQuoteId

`func (o *CreateRfqCancelQuoteV5Req) HasClQuoteId() bool`

HasClQuoteId returns a boolean if a field has been set.

### GetQuoteId

`func (o *CreateRfqCancelQuoteV5Req) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *CreateRfqCancelQuoteV5Req) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *CreateRfqCancelQuoteV5Req) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *CreateRfqCancelQuoteV5Req) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetRfqId

`func (o *CreateRfqCancelQuoteV5Req) GetRfqId() string`

GetRfqId returns the RfqId field if non-nil, zero value otherwise.

### GetRfqIdOk

`func (o *CreateRfqCancelQuoteV5Req) GetRfqIdOk() (*string, bool)`

GetRfqIdOk returns a tuple with the RfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqId

`func (o *CreateRfqCancelQuoteV5Req) SetRfqId(v string)`

SetRfqId sets RfqId field to given value.

### HasRfqId

`func (o *CreateRfqCancelQuoteV5Req) HasRfqId() bool`

HasRfqId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


