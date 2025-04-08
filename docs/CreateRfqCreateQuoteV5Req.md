# CreateRfqCreateQuoteV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anonymous** | Pointer to **bool** | Submit Quote on a disclosed or anonymous basis.   Valid value is &#x60;true&#x60; or &#x60;false&#x60;. &#x60;false&#x60; by default. | [optional] 
**ClQuoteId** | Pointer to **string** | Client-supplied Quote ID.   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [optional] [default to ""]
**ExpiresIn** | Pointer to **string** | Seconds that a quote expires in.   Must be an integer between 10-120. Default is 60. | [optional] [default to ""]
**Legs** | [**[]CreateRfqCreateQuoteV5ReqLegsInner**](CreateRfqCreateQuoteV5ReqLegsInner.md) | The legs of the Quote. | 
**QuoteSide** | **string** | The trading direction of the Quote. Its value can be &#x60;buy&#x60; or &#x60;sell&#x60;.   For example, if quoteSide is &#x60;buy&#x60;, all the legs are executed in their leg sides; otherwise, all the legs are executed in the opposite of their leg sides. | [default to ""]
**RfqId** | **string** | RFQ ID . | [default to ""]
**Tag** | Pointer to **string** | Quote tag.   The block trade associated with the Quote will have the same tag.   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters. | [optional] [default to ""]

## Methods

### NewCreateRfqCreateQuoteV5Req

`func NewCreateRfqCreateQuoteV5Req(legs []CreateRfqCreateQuoteV5ReqLegsInner, quoteSide string, rfqId string, ) *CreateRfqCreateQuoteV5Req`

NewCreateRfqCreateQuoteV5Req instantiates a new CreateRfqCreateQuoteV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCreateQuoteV5ReqWithDefaults

`func NewCreateRfqCreateQuoteV5ReqWithDefaults() *CreateRfqCreateQuoteV5Req`

NewCreateRfqCreateQuoteV5ReqWithDefaults instantiates a new CreateRfqCreateQuoteV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymous

`func (o *CreateRfqCreateQuoteV5Req) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *CreateRfqCreateQuoteV5Req) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *CreateRfqCreateQuoteV5Req) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *CreateRfqCreateQuoteV5Req) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetClQuoteId

`func (o *CreateRfqCreateQuoteV5Req) GetClQuoteId() string`

GetClQuoteId returns the ClQuoteId field if non-nil, zero value otherwise.

### GetClQuoteIdOk

`func (o *CreateRfqCreateQuoteV5Req) GetClQuoteIdOk() (*string, bool)`

GetClQuoteIdOk returns a tuple with the ClQuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClQuoteId

`func (o *CreateRfqCreateQuoteV5Req) SetClQuoteId(v string)`

SetClQuoteId sets ClQuoteId field to given value.

### HasClQuoteId

`func (o *CreateRfqCreateQuoteV5Req) HasClQuoteId() bool`

HasClQuoteId returns a boolean if a field has been set.

### GetExpiresIn

`func (o *CreateRfqCreateQuoteV5Req) GetExpiresIn() string`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CreateRfqCreateQuoteV5Req) GetExpiresInOk() (*string, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CreateRfqCreateQuoteV5Req) SetExpiresIn(v string)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CreateRfqCreateQuoteV5Req) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetLegs

`func (o *CreateRfqCreateQuoteV5Req) GetLegs() []CreateRfqCreateQuoteV5ReqLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *CreateRfqCreateQuoteV5Req) GetLegsOk() (*[]CreateRfqCreateQuoteV5ReqLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *CreateRfqCreateQuoteV5Req) SetLegs(v []CreateRfqCreateQuoteV5ReqLegsInner)`

SetLegs sets Legs field to given value.


### GetQuoteSide

`func (o *CreateRfqCreateQuoteV5Req) GetQuoteSide() string`

GetQuoteSide returns the QuoteSide field if non-nil, zero value otherwise.

### GetQuoteSideOk

`func (o *CreateRfqCreateQuoteV5Req) GetQuoteSideOk() (*string, bool)`

GetQuoteSideOk returns a tuple with the QuoteSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteSide

`func (o *CreateRfqCreateQuoteV5Req) SetQuoteSide(v string)`

SetQuoteSide sets QuoteSide field to given value.


### GetRfqId

`func (o *CreateRfqCreateQuoteV5Req) GetRfqId() string`

GetRfqId returns the RfqId field if non-nil, zero value otherwise.

### GetRfqIdOk

`func (o *CreateRfqCreateQuoteV5Req) GetRfqIdOk() (*string, bool)`

GetRfqIdOk returns a tuple with the RfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqId

`func (o *CreateRfqCreateQuoteV5Req) SetRfqId(v string)`

SetRfqId sets RfqId field to given value.


### GetTag

`func (o *CreateRfqCreateQuoteV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateRfqCreateQuoteV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateRfqCreateQuoteV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateRfqCreateQuoteV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


