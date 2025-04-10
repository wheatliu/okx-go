/*
OKX v5 API

OpenAPI specification for Okx exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
)

// checks if the CreateAccountMmpResetV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountMmpResetV5RespDataInner{}

// CreateAccountMmpResetV5RespDataInner struct for CreateAccountMmpResetV5RespDataInner
type CreateAccountMmpResetV5RespDataInner struct {
	// Result of the request `true`, `false`
	Result *bool `json:"result,omitempty"`
}

// NewCreateAccountMmpResetV5RespDataInner instantiates a new CreateAccountMmpResetV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountMmpResetV5RespDataInner() *CreateAccountMmpResetV5RespDataInner {
	this := CreateAccountMmpResetV5RespDataInner{}
	return &this
}

// NewCreateAccountMmpResetV5RespDataInnerWithDefaults instantiates a new CreateAccountMmpResetV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountMmpResetV5RespDataInnerWithDefaults() *CreateAccountMmpResetV5RespDataInner {
	this := CreateAccountMmpResetV5RespDataInner{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateAccountMmpResetV5RespDataInner) GetResult() bool {
	if o == nil || IsNil(o.Result) {
		var ret bool
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountMmpResetV5RespDataInner) GetResultOk() (*bool, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateAccountMmpResetV5RespDataInner) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given bool and assigns it to the Result field.
func (o *CreateAccountMmpResetV5RespDataInner) SetResult(v bool) {
	o.Result = &v
}

func (o CreateAccountMmpResetV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountMmpResetV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableCreateAccountMmpResetV5RespDataInner struct {
	value *CreateAccountMmpResetV5RespDataInner
	isSet bool
}

func (v NullableCreateAccountMmpResetV5RespDataInner) Get() *CreateAccountMmpResetV5RespDataInner {
	return v.value
}

func (v *NullableCreateAccountMmpResetV5RespDataInner) Set(val *CreateAccountMmpResetV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountMmpResetV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountMmpResetV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountMmpResetV5RespDataInner(val *CreateAccountMmpResetV5RespDataInner) *NullableCreateAccountMmpResetV5RespDataInner {
	return &NullableCreateAccountMmpResetV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateAccountMmpResetV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountMmpResetV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


