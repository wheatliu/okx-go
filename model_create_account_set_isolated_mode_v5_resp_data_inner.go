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

// checks if the CreateAccountSetIsolatedModeV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountSetIsolatedModeV5RespDataInner{}

// CreateAccountSetIsolatedModeV5RespDataInner struct for CreateAccountSetIsolatedModeV5RespDataInner
type CreateAccountSetIsolatedModeV5RespDataInner struct {
	// Isolated margin trading settings  `automatic`: Auto transfers
	IsoMode *string `json:"isoMode,omitempty"`
}

// NewCreateAccountSetIsolatedModeV5RespDataInner instantiates a new CreateAccountSetIsolatedModeV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountSetIsolatedModeV5RespDataInner() *CreateAccountSetIsolatedModeV5RespDataInner {
	this := CreateAccountSetIsolatedModeV5RespDataInner{}
	var isoMode string = ""
	this.IsoMode = &isoMode
	return &this
}

// NewCreateAccountSetIsolatedModeV5RespDataInnerWithDefaults instantiates a new CreateAccountSetIsolatedModeV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountSetIsolatedModeV5RespDataInnerWithDefaults() *CreateAccountSetIsolatedModeV5RespDataInner {
	this := CreateAccountSetIsolatedModeV5RespDataInner{}
	var isoMode string = ""
	this.IsoMode = &isoMode
	return &this
}

// GetIsoMode returns the IsoMode field value if set, zero value otherwise.
func (o *CreateAccountSetIsolatedModeV5RespDataInner) GetIsoMode() string {
	if o == nil || IsNil(o.IsoMode) {
		var ret string
		return ret
	}
	return *o.IsoMode
}

// GetIsoModeOk returns a tuple with the IsoMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountSetIsolatedModeV5RespDataInner) GetIsoModeOk() (*string, bool) {
	if o == nil || IsNil(o.IsoMode) {
		return nil, false
	}
	return o.IsoMode, true
}

// HasIsoMode returns a boolean if a field has been set.
func (o *CreateAccountSetIsolatedModeV5RespDataInner) HasIsoMode() bool {
	if o != nil && !IsNil(o.IsoMode) {
		return true
	}

	return false
}

// SetIsoMode gets a reference to the given string and assigns it to the IsoMode field.
func (o *CreateAccountSetIsolatedModeV5RespDataInner) SetIsoMode(v string) {
	o.IsoMode = &v
}

func (o CreateAccountSetIsolatedModeV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountSetIsolatedModeV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsoMode) {
		toSerialize["isoMode"] = o.IsoMode
	}
	return toSerialize, nil
}

type NullableCreateAccountSetIsolatedModeV5RespDataInner struct {
	value *CreateAccountSetIsolatedModeV5RespDataInner
	isSet bool
}

func (v NullableCreateAccountSetIsolatedModeV5RespDataInner) Get() *CreateAccountSetIsolatedModeV5RespDataInner {
	return v.value
}

func (v *NullableCreateAccountSetIsolatedModeV5RespDataInner) Set(val *CreateAccountSetIsolatedModeV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountSetIsolatedModeV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountSetIsolatedModeV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountSetIsolatedModeV5RespDataInner(val *CreateAccountSetIsolatedModeV5RespDataInner) *NullableCreateAccountSetIsolatedModeV5RespDataInner {
	return &NullableCreateAccountSetIsolatedModeV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateAccountSetIsolatedModeV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountSetIsolatedModeV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


