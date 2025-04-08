# GetPublicDeliveryExerciseHistoryV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner**](GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner.md) | Delivery/exercise details | [optional] 
**Ts** | Pointer to **string** | Delivery/exercise time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicDeliveryExerciseHistoryV5RespData

`func NewGetPublicDeliveryExerciseHistoryV5RespData() *GetPublicDeliveryExerciseHistoryV5RespData`

NewGetPublicDeliveryExerciseHistoryV5RespData instantiates a new GetPublicDeliveryExerciseHistoryV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicDeliveryExerciseHistoryV5RespDataWithDefaults

`func NewGetPublicDeliveryExerciseHistoryV5RespDataWithDefaults() *GetPublicDeliveryExerciseHistoryV5RespData`

NewGetPublicDeliveryExerciseHistoryV5RespDataWithDefaults instantiates a new GetPublicDeliveryExerciseHistoryV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *GetPublicDeliveryExerciseHistoryV5RespData) GetDetails() []GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetPublicDeliveryExerciseHistoryV5RespData) GetDetailsOk() (*[]GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetPublicDeliveryExerciseHistoryV5RespData) SetDetails(v []GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetPublicDeliveryExerciseHistoryV5RespData) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicDeliveryExerciseHistoryV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicDeliveryExerciseHistoryV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicDeliveryExerciseHistoryV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicDeliveryExerciseHistoryV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


