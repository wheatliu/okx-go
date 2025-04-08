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

// checks if the CreateRfqMmpResetV5RespDataInnerDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRfqMmpResetV5RespDataInnerDataInner{}

// CreateRfqMmpResetV5RespDataInnerDataInner struct for CreateRfqMmpResetV5RespDataInnerDataInner
type CreateRfqMmpResetV5RespDataInnerDataInner struct {
	// The timestamp of re-setting successfully. Unix timestamp format in milliseconds, e.g. `1597026383085`.
	Ts *string `json:"ts,omitempty"`
}

// NewCreateRfqMmpResetV5RespDataInnerDataInner instantiates a new CreateRfqMmpResetV5RespDataInnerDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRfqMmpResetV5RespDataInnerDataInner() *CreateRfqMmpResetV5RespDataInnerDataInner {
	this := CreateRfqMmpResetV5RespDataInnerDataInner{}
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewCreateRfqMmpResetV5RespDataInnerDataInnerWithDefaults instantiates a new CreateRfqMmpResetV5RespDataInnerDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRfqMmpResetV5RespDataInnerDataInnerWithDefaults() *CreateRfqMmpResetV5RespDataInnerDataInner {
	this := CreateRfqMmpResetV5RespDataInnerDataInner{}
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *CreateRfqMmpResetV5RespDataInnerDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqMmpResetV5RespDataInnerDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *CreateRfqMmpResetV5RespDataInnerDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *CreateRfqMmpResetV5RespDataInnerDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o CreateRfqMmpResetV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRfqMmpResetV5RespDataInnerDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableCreateRfqMmpResetV5RespDataInnerDataInner struct {
	value *CreateRfqMmpResetV5RespDataInnerDataInner
	isSet bool
}

func (v NullableCreateRfqMmpResetV5RespDataInnerDataInner) Get() *CreateRfqMmpResetV5RespDataInnerDataInner {
	return v.value
}

func (v *NullableCreateRfqMmpResetV5RespDataInnerDataInner) Set(val *CreateRfqMmpResetV5RespDataInnerDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRfqMmpResetV5RespDataInnerDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRfqMmpResetV5RespDataInnerDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRfqMmpResetV5RespDataInnerDataInner(val *CreateRfqMmpResetV5RespDataInnerDataInner) *NullableCreateRfqMmpResetV5RespDataInnerDataInner {
	return &NullableCreateRfqMmpResetV5RespDataInnerDataInner{value: val, isSet: true}
}

func (v NullableCreateRfqMmpResetV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRfqMmpResetV5RespDataInnerDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


