# CreateRfqCancelBatchRfqsV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClRfqIds** | Pointer to **[]string** | Client-supplied RFQ IDs.   Either &#x60;rfqIds&#x60; or &#x60;clRfqIds&#x60; is required.   If both attributes are sent, &#x60;rfqIds&#x60; will be used as primary identifier. | [optional] 
**RfqIds** | Pointer to **[]string** | RFQ IDs . | [optional] 

## Methods

### NewCreateRfqCancelBatchRfqsV5Req

`func NewCreateRfqCancelBatchRfqsV5Req() *CreateRfqCancelBatchRfqsV5Req`

NewCreateRfqCancelBatchRfqsV5Req instantiates a new CreateRfqCancelBatchRfqsV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCancelBatchRfqsV5ReqWithDefaults

`func NewCreateRfqCancelBatchRfqsV5ReqWithDefaults() *CreateRfqCancelBatchRfqsV5Req`

NewCreateRfqCancelBatchRfqsV5ReqWithDefaults instantiates a new CreateRfqCancelBatchRfqsV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClRfqIds

`func (o *CreateRfqCancelBatchRfqsV5Req) GetClRfqIds() []string`

GetClRfqIds returns the ClRfqIds field if non-nil, zero value otherwise.

### GetClRfqIdsOk

`func (o *CreateRfqCancelBatchRfqsV5Req) GetClRfqIdsOk() (*[]string, bool)`

GetClRfqIdsOk returns a tuple with the ClRfqIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClRfqIds

`func (o *CreateRfqCancelBatchRfqsV5Req) SetClRfqIds(v []string)`

SetClRfqIds sets ClRfqIds field to given value.

### HasClRfqIds

`func (o *CreateRfqCancelBatchRfqsV5Req) HasClRfqIds() bool`

HasClRfqIds returns a boolean if a field has been set.

### GetRfqIds

`func (o *CreateRfqCancelBatchRfqsV5Req) GetRfqIds() []string`

GetRfqIds returns the RfqIds field if non-nil, zero value otherwise.

### GetRfqIdsOk

`func (o *CreateRfqCancelBatchRfqsV5Req) GetRfqIdsOk() (*[]string, bool)`

GetRfqIdsOk returns a tuple with the RfqIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqIds

`func (o *CreateRfqCancelBatchRfqsV5Req) SetRfqIds(v []string)`

SetRfqIds sets RfqIds field to given value.

### HasRfqIds

`func (o *CreateRfqCancelBatchRfqsV5Req) HasRfqIds() bool`

HasRfqIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


