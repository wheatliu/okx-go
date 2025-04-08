# GetPublicDeliveryExerciseHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner**](GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner.md) | Delivery/exercise details | [optional] 
**Ts** | Pointer to **string** | Delivery/exercise time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicDeliveryExerciseHistoryV5RespDataInner

`func NewGetPublicDeliveryExerciseHistoryV5RespDataInner() *GetPublicDeliveryExerciseHistoryV5RespDataInner`

NewGetPublicDeliveryExerciseHistoryV5RespDataInner instantiates a new GetPublicDeliveryExerciseHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicDeliveryExerciseHistoryV5RespDataInnerWithDefaults

`func NewGetPublicDeliveryExerciseHistoryV5RespDataInnerWithDefaults() *GetPublicDeliveryExerciseHistoryV5RespDataInner`

NewGetPublicDeliveryExerciseHistoryV5RespDataInnerWithDefaults instantiates a new GetPublicDeliveryExerciseHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataInner) GetDetails() []GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataInner) GetDetailsOk() (*[]GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataInner) SetDetails(v []GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicDeliveryExerciseHistoryV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


