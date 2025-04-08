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

// checks if the CreateRfqMmpResetV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRfqMmpResetV5RespDataInner{}

// CreateRfqMmpResetV5RespDataInner struct for CreateRfqMmpResetV5RespDataInner
type CreateRfqMmpResetV5RespDataInner struct {
	// The timestamp of re-setting successfully. Unix timestamp format in milliseconds, e.g. `1597026383085`.
	Ts *string `json:"ts,omitempty"`
}

// NewCreateRfqMmpResetV5RespDataInner instantiates a new CreateRfqMmpResetV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRfqMmpResetV5RespDataInner() *CreateRfqMmpResetV5RespDataInner {
	this := CreateRfqMmpResetV5RespDataInner{}
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewCreateRfqMmpResetV5RespDataInnerWithDefaults instantiates a new CreateRfqMmpResetV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRfqMmpResetV5RespDataInnerWithDefaults() *CreateRfqMmpResetV5RespDataInner {
	this := CreateRfqMmpResetV5RespDataInner{}
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *CreateRfqMmpResetV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqMmpResetV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *CreateRfqMmpResetV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *CreateRfqMmpResetV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o CreateRfqMmpResetV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRfqMmpResetV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableCreateRfqMmpResetV5RespDataInner struct {
	value *CreateRfqMmpResetV5RespDataInner
	isSet bool
}

func (v NullableCreateRfqMmpResetV5RespDataInner) Get() *CreateRfqMmpResetV5RespDataInner {
	return v.value
}

func (v *NullableCreateRfqMmpResetV5RespDataInner) Set(val *CreateRfqMmpResetV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRfqMmpResetV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRfqMmpResetV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRfqMmpResetV5RespDataInner(val *CreateRfqMmpResetV5RespDataInner) *NullableCreateRfqMmpResetV5RespDataInner {
	return &NullableCreateRfqMmpResetV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateRfqMmpResetV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRfqMmpResetV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


