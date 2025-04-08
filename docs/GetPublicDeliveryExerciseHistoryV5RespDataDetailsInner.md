# GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InsId** | Pointer to **string** | Delivery/exercise contract ID | [optional] [default to ""]
**Px** | Pointer to **string** | Delivery/exercise price | [optional] [default to ""]
**Type** | Pointer to **string** | Type    &#x60;delivery&#x60;   &#x60;exercised&#x60;   &#x60;expired_otm&#x60;:Out of the money | [optional] [default to ""]

## Methods

### NewGetPublicDeliveryExerciseHistoryV5RespDataDetailsInner

`func NewGetPublicDeliveryExerciseHistoryV5RespDataDetailsInner() *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner`

NewGetPublicDeliveryExerciseHistoryV5RespDataDetailsInner instantiates a new GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicDeliveryExerciseHistoryV5RespDataDetailsInnerWithDefaults

`func NewGetPublicDeliveryExerciseHistoryV5RespDataDetailsInnerWithDefaults() *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner`

NewGetPublicDeliveryExerciseHistoryV5RespDataDetailsInnerWithDefaults instantiates a new GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsId

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner) GetInsId() string`

GetInsId returns the InsId field if non-nil, zero value otherwise.

### GetInsIdOk

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner) GetInsIdOk() (*string, bool)`

GetInsIdOk returns a tuple with the InsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsId

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner) SetInsId(v string)`

SetInsId sets InsId field to given value.

### HasInsId

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner) HasInsId() bool`

HasInsId returns a boolean if a field has been set.

### GetPx

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetType

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


