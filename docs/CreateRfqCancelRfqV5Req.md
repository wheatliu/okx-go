# CreateRfqCancelRfqV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClRfqId** | Pointer to **string** | Client-supplied RFQ ID.   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.    Either rfqId or clRfqId is required. If both are passed, rfqId will be used. | [optional] [default to ""]
**RfqId** | Pointer to **string** | RFQ ID created . | [optional] [default to ""]

## Methods

### NewCreateRfqCancelRfqV5Req

`func NewCreateRfqCancelRfqV5Req() *CreateRfqCancelRfqV5Req`

NewCreateRfqCancelRfqV5Req instantiates a new CreateRfqCancelRfqV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCancelRfqV5ReqWithDefaults

`func NewCreateRfqCancelRfqV5ReqWithDefaults() *CreateRfqCancelRfqV5Req`

NewCreateRfqCancelRfqV5ReqWithDefaults instantiates a new CreateRfqCancelRfqV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClRfqId

`func (o *CreateRfqCancelRfqV5Req) GetClRfqId() string`

GetClRfqId returns the ClRfqId field if non-nil, zero value otherwise.

### GetClRfqIdOk

`func (o *CreateRfqCancelRfqV5Req) GetClRfqIdOk() (*string, bool)`

GetClRfqIdOk returns a tuple with the ClRfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClRfqId

`func (o *CreateRfqCancelRfqV5Req) SetClRfqId(v string)`

SetClRfqId sets ClRfqId field to given value.

### HasClRfqId

`func (o *CreateRfqCancelRfqV5Req) HasClRfqId() bool`

HasClRfqId returns a boolean if a field has been set.

### GetRfqId

`func (o *CreateRfqCancelRfqV5Req) GetRfqId() string`

GetRfqId returns the RfqId field if non-nil, zero value otherwise.

### GetRfqIdOk

`func (o *CreateRfqCancelRfqV5Req) GetRfqIdOk() (*string, bool)`

GetRfqIdOk returns a tuple with the RfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqId

`func (o *CreateRfqCancelRfqV5Req) SetRfqId(v string)`

SetRfqId sets RfqId field to given value.

### HasRfqId

`func (o *CreateRfqCancelRfqV5Req) HasRfqId() bool`

HasRfqId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


