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

// checks if the CreateRfqCancelAllRfqsV5RespDataInnerDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRfqCancelAllRfqsV5RespDataInnerDataInner{}

// CreateRfqCancelAllRfqsV5RespDataInnerDataInner struct for CreateRfqCancelAllRfqsV5RespDataInnerDataInner
type CreateRfqCancelAllRfqsV5RespDataInnerDataInner struct {
	// The timestamp of successful cancellation. Unix timestamp format in milliseconds, e.g. 1597026383085.
	Ts *string `json:"ts,omitempty"`
}

// NewCreateRfqCancelAllRfqsV5RespDataInnerDataInner instantiates a new CreateRfqCancelAllRfqsV5RespDataInnerDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRfqCancelAllRfqsV5RespDataInnerDataInner() *CreateRfqCancelAllRfqsV5RespDataInnerDataInner {
	this := CreateRfqCancelAllRfqsV5RespDataInnerDataInner{}
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewCreateRfqCancelAllRfqsV5RespDataInnerDataInnerWithDefaults instantiates a new CreateRfqCancelAllRfqsV5RespDataInnerDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRfqCancelAllRfqsV5RespDataInnerDataInnerWithDefaults() *CreateRfqCancelAllRfqsV5RespDataInnerDataInner {
	this := CreateRfqCancelAllRfqsV5RespDataInnerDataInner{}
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *CreateRfqCancelAllRfqsV5RespDataInnerDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCancelAllRfqsV5RespDataInnerDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *CreateRfqCancelAllRfqsV5RespDataInnerDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *CreateRfqCancelAllRfqsV5RespDataInnerDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o CreateRfqCancelAllRfqsV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRfqCancelAllRfqsV5RespDataInnerDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableCreateRfqCancelAllRfqsV5RespDataInnerDataInner struct {
	value *CreateRfqCancelAllRfqsV5RespDataInnerDataInner
	isSet bool
}

func (v NullableCreateRfqCancelAllRfqsV5RespDataInnerDataInner) Get() *CreateRfqCancelAllRfqsV5RespDataInnerDataInner {
	return v.value
}

func (v *NullableCreateRfqCancelAllRfqsV5RespDataInnerDataInner) Set(val *CreateRfqCancelAllRfqsV5RespDataInnerDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRfqCancelAllRfqsV5RespDataInnerDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRfqCancelAllRfqsV5RespDataInnerDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRfqCancelAllRfqsV5RespDataInnerDataInner(val *CreateRfqCancelAllRfqsV5RespDataInnerDataInner) *NullableCreateRfqCancelAllRfqsV5RespDataInnerDataInner {
	return &NullableCreateRfqCancelAllRfqsV5RespDataInnerDataInner{value: val, isSet: true}
}

func (v NullableCreateRfqCancelAllRfqsV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRfqCancelAllRfqsV5RespDataInnerDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


