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

// checks if the GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner{}

// GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner struct for GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner
type GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner struct {
	// Delivery/exercise contract ID
	InsId *string `json:"insId,omitempty"`
	// Delivery/exercise price
	Px *string `json:"px,omitempty"`
	// Type    `delivery`   `exercised`   `expired_otm`:Out of the money
	Type *string `json:"type,omitempty"`
}

// NewGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner instantiates a new GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner() *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner {
	this := GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner{}
	var insId string = ""
	this.InsId = &insId
	var px string = ""
	this.Px = &px
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInnerWithDefaults instantiates a new GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInnerWithDefaults() *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner {
	this := GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner{}
	var insId string = ""
	this.InsId = &insId
	var px string = ""
	this.Px = &px
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetInsId returns the InsId field value if set, zero value otherwise.
func (o *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) GetInsId() string {
	if o == nil || IsNil(o.InsId) {
		var ret string
		return ret
	}
	return *o.InsId
}

// GetInsIdOk returns a tuple with the InsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) GetInsIdOk() (*string, bool) {
	if o == nil || IsNil(o.InsId) {
		return nil, false
	}
	return o.InsId, true
}

// HasInsId returns a boolean if a field has been set.
func (o *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) HasInsId() bool {
	if o != nil && !IsNil(o.InsId) {
		return true
	}

	return false
}

// SetInsId gets a reference to the given string and assigns it to the InsId field.
func (o *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) SetInsId(v string) {
	o.InsId = &v
}

// GetPx returns the Px field value if set, zero value otherwise.
func (o *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) GetPx() string {
	if o == nil || IsNil(o.Px) {
		var ret string
		return ret
	}
	return *o.Px
}

// GetPxOk returns a tuple with the Px field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) GetPxOk() (*string, bool) {
	if o == nil || IsNil(o.Px) {
		return nil, false
	}
	return o.Px, true
}

// HasPx returns a boolean if a field has been set.
func (o *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) HasPx() bool {
	if o != nil && !IsNil(o.Px) {
		return true
	}

	return false
}

// SetPx gets a reference to the given string and assigns it to the Px field.
func (o *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) SetPx(v string) {
	o.Px = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) SetType(v string) {
	o.Type = &v
}

func (o GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InsId) {
		toSerialize["insId"] = o.InsId
	}
	if !IsNil(o.Px) {
		toSerialize["px"] = o.Px
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner struct {
	value *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner
	isSet bool
}

func (v NullableGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) Get() *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner {
	return v.value
}

func (v *NullableGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) Set(val *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner(val *GetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) *NullableGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner {
	return &NullableGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner{value: val, isSet: true}
}

func (v NullableGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPublicDeliveryExerciseHistoryV5RespDataInnerDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


