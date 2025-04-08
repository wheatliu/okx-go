# CreateRfqCreateRfqV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPartialExecution** | Pointer to **bool** | Whether the RFQ can be partially filled provided that the shape of legs stays the same. Valid values are &#x60;true&#x60; or &#x60;false&#x60;.   &#x60;false&#x60; by default. | [optional] 
**Anonymous** | Pointer to **bool** | Submit RFQ on a disclosed or anonymous basis. Valid values are &#x60;true&#x60; or &#x60;false&#x60;.   If not specified, the default value is &#x60;false&#x60;.   When anonymous &#x3D; true, the taker’s identify is not disclosed to maker even after trade execution. | [optional] 
**ClRfqId** | Pointer to **string** | Client-supplied RFQ ID.   A combination of case-sensitive alpha-numeric, all numbers, or all letters of up to 32 characters. | [optional] [default to ""]
**Counterparties** | **[]string** | The trader code(s) of the counterparties who receive the RFQ. Can be found via /api/v5/rfq/counterparties/ | 
**Legs** | [**[]CreateRfqCreateRfqV5ReqLegsInner**](CreateRfqCreateRfqV5ReqLegsInner.md) | An Array of objects containing each leg of the RFQ. Maximum 15 legs can be placed per request | 
**Tag** | Pointer to **string** | RFQ tag.   The block trade associated with the RFQ will have the same tag.   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters. | [optional] [default to ""]

## Methods

### NewCreateRfqCreateRfqV5Req

`func NewCreateRfqCreateRfqV5Req(counterparties []string, legs []CreateRfqCreateRfqV5ReqLegsInner, ) *CreateRfqCreateRfqV5Req`

NewCreateRfqCreateRfqV5Req instantiates a new CreateRfqCreateRfqV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCreateRfqV5ReqWithDefaults

`func NewCreateRfqCreateRfqV5ReqWithDefaults() *CreateRfqCreateRfqV5Req`

NewCreateRfqCreateRfqV5ReqWithDefaults instantiates a new CreateRfqCreateRfqV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPartialExecution

`func (o *CreateRfqCreateRfqV5Req) GetAllowPartialExecution() bool`

GetAllowPartialExecution returns the AllowPartialExecution field if non-nil, zero value otherwise.

### GetAllowPartialExecutionOk

`func (o *CreateRfqCreateRfqV5Req) GetAllowPartialExecutionOk() (*bool, bool)`

GetAllowPartialExecutionOk returns a tuple with the AllowPartialExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPartialExecution

`func (o *CreateRfqCreateRfqV5Req) SetAllowPartialExecution(v bool)`

SetAllowPartialExecution sets AllowPartialExecution field to given value.

### HasAllowPartialExecution

`func (o *CreateRfqCreateRfqV5Req) HasAllowPartialExecution() bool`

HasAllowPartialExecution returns a boolean if a field has been set.

### GetAnonymous

`func (o *CreateRfqCreateRfqV5Req) GetAnonymous() bool`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *CreateRfqCreateRfqV5Req) GetAnonymousOk() (*bool, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *CreateRfqCreateRfqV5Req) SetAnonymous(v bool)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *CreateRfqCreateRfqV5Req) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetClRfqId

`func (o *CreateRfqCreateRfqV5Req) GetClRfqId() string`

GetClRfqId returns the ClRfqId field if non-nil, zero value otherwise.

### GetClRfqIdOk

`func (o *CreateRfqCreateRfqV5Req) GetClRfqIdOk() (*string, bool)`

GetClRfqIdOk returns a tuple with the ClRfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClRfqId

`func (o *CreateRfqCreateRfqV5Req) SetClRfqId(v string)`

SetClRfqId sets ClRfqId field to given value.

### HasClRfqId

`func (o *CreateRfqCreateRfqV5Req) HasClRfqId() bool`

HasClRfqId returns a boolean if a field has been set.

### GetCounterparties

`func (o *CreateRfqCreateRfqV5Req) GetCounterparties() []string`

GetCounterparties returns the Counterparties field if non-nil, zero value otherwise.

### GetCounterpartiesOk

`func (o *CreateRfqCreateRfqV5Req) GetCounterpartiesOk() (*[]string, bool)`

GetCounterpartiesOk returns a tuple with the Counterparties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterparties

`func (o *CreateRfqCreateRfqV5Req) SetCounterparties(v []string)`

SetCounterparties sets Counterparties field to given value.


### GetLegs

`func (o *CreateRfqCreateRfqV5Req) GetLegs() []CreateRfqCreateRfqV5ReqLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *CreateRfqCreateRfqV5Req) GetLegsOk() (*[]CreateRfqCreateRfqV5ReqLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *CreateRfqCreateRfqV5Req) SetLegs(v []CreateRfqCreateRfqV5ReqLegsInner)`

SetLegs sets Legs field to given value.


### GetTag

`func (o *CreateRfqCreateRfqV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateRfqCreateRfqV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateRfqCreateRfqV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateRfqCreateRfqV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


