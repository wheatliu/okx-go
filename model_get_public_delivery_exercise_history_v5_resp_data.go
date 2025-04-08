/*
Okx Rest API

OpenAPI specification for Okx cryptocurrency exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
)

// checks if the GetPublicDeliveryExerciseHistoryV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPublicDeliveryExerciseHistoryV5RespData{}

// GetPublicDeliveryExerciseHistoryV5RespData struct for GetPublicDeliveryExerciseHistoryV5RespData
type GetPublicDeliveryExerciseHistoryV5RespData struct {
	// Delivery/exercise details
	Details []GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner `json:"details,omitempty"`
	// Delivery/exercise time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewGetPublicDeliveryExerciseHistoryV5RespData instantiates a new GetPublicDeliveryExerciseHistoryV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPublicDeliveryExerciseHistoryV5RespData() *GetPublicDeliveryExerciseHistoryV5RespData {
	this := GetPublicDeliveryExerciseHistoryV5RespData{}
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetPublicDeliveryExerciseHistoryV5RespDataWithDefaults instantiates a new GetPublicDeliveryExerciseHistoryV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPublicDeliveryExerciseHistoryV5RespDataWithDefaults() *GetPublicDeliveryExerciseHistoryV5RespData {
	this := GetPublicDeliveryExerciseHistoryV5RespData{}
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GetPublicDeliveryExerciseHistoryV5RespData) GetDetails() []GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner {
	if o == nil || IsNil(o.Details) {
		var ret []GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicDeliveryExerciseHistoryV5RespData) GetDetailsOk() ([]GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GetPublicDeliveryExerciseHistoryV5RespData) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner and assigns it to the Details field.
func (o *GetPublicDeliveryExerciseHistoryV5RespData) SetDetails(v []GetPublicDeliveryExerciseHistoryV5RespDataDetailsInner) {
	o.Details = v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetPublicDeliveryExerciseHistoryV5RespData) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicDeliveryExerciseHistoryV5RespData) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetPublicDeliveryExerciseHistoryV5RespData) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetPublicDeliveryExerciseHistoryV5RespData) SetTs(v string) {
	o.Ts = &v
}

func (o GetPublicDeliveryExerciseHistoryV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPublicDeliveryExerciseHistoryV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetPublicDeliveryExerciseHistoryV5RespData struct {
	value *GetPublicDeliveryExerciseHistoryV5RespData
	isSet bool
}

func (v NullableGetPublicDeliveryExerciseHistoryV5RespData) Get() *GetPublicDeliveryExerciseHistoryV5RespData {
	return v.value
}

func (v *NullableGetPublicDeliveryExerciseHistoryV5RespData) Set(val *GetPublicDeliveryExerciseHistoryV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPublicDeliveryExerciseHistoryV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPublicDeliveryExerciseHistoryV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPublicDeliveryExerciseHistoryV5RespData(val *GetPublicDeliveryExerciseHistoryV5RespData) *NullableGetPublicDeliveryExerciseHistoryV5RespData {
	return &NullableGetPublicDeliveryExerciseHistoryV5RespData{value: val, isSet: true}
}

func (v NullableGetPublicDeliveryExerciseHistoryV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPublicDeliveryExerciseHistoryV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


