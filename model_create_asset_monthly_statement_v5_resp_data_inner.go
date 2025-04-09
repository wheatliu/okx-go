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

// checks if the CreateAssetMonthlyStatementV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssetMonthlyStatementV5RespDataInner{}

// CreateAssetMonthlyStatementV5RespDataInner struct for CreateAssetMonthlyStatementV5RespDataInner
type CreateAssetMonthlyStatementV5RespDataInner struct {
	// Download link generation time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewCreateAssetMonthlyStatementV5RespDataInner instantiates a new CreateAssetMonthlyStatementV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssetMonthlyStatementV5RespDataInner() *CreateAssetMonthlyStatementV5RespDataInner {
	this := CreateAssetMonthlyStatementV5RespDataInner{}
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewCreateAssetMonthlyStatementV5RespDataInnerWithDefaults instantiates a new CreateAssetMonthlyStatementV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssetMonthlyStatementV5RespDataInnerWithDefaults() *CreateAssetMonthlyStatementV5RespDataInner {
	this := CreateAssetMonthlyStatementV5RespDataInner{}
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *CreateAssetMonthlyStatementV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetMonthlyStatementV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *CreateAssetMonthlyStatementV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *CreateAssetMonthlyStatementV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o CreateAssetMonthlyStatementV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAssetMonthlyStatementV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableCreateAssetMonthlyStatementV5RespDataInner struct {
	value *CreateAssetMonthlyStatementV5RespDataInner
	isSet bool
}

func (v NullableCreateAssetMonthlyStatementV5RespDataInner) Get() *CreateAssetMonthlyStatementV5RespDataInner {
	return v.value
}

func (v *NullableCreateAssetMonthlyStatementV5RespDataInner) Set(val *CreateAssetMonthlyStatementV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssetMonthlyStatementV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssetMonthlyStatementV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssetMonthlyStatementV5RespDataInner(val *CreateAssetMonthlyStatementV5RespDataInner) *NullableCreateAssetMonthlyStatementV5RespDataInner {
	return &NullableCreateAssetMonthlyStatementV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateAssetMonthlyStatementV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAssetMonthlyStatementV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


