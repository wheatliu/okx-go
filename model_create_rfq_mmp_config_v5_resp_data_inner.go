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

// checks if the CreateRfqMmpConfigV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRfqMmpConfigV5RespDataInner{}

// CreateRfqMmpConfigV5RespDataInner struct for CreateRfqMmpConfigV5RespDataInner
type CreateRfqMmpConfigV5RespDataInner struct {
	// Limit in number of execution attempts
	CountLimit *string `json:"countLimit,omitempty"`
	// Frozen period (ms).
	FrozenInterval *string `json:"frozenInterval,omitempty"`
	// Time window (ms). MMP interval where monitoring is done
	TimeInterval *string `json:"timeInterval,omitempty"`
}

// NewCreateRfqMmpConfigV5RespDataInner instantiates a new CreateRfqMmpConfigV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRfqMmpConfigV5RespDataInner() *CreateRfqMmpConfigV5RespDataInner {
	this := CreateRfqMmpConfigV5RespDataInner{}
	var countLimit string = ""
	this.CountLimit = &countLimit
	var frozenInterval string = ""
	this.FrozenInterval = &frozenInterval
	var timeInterval string = ""
	this.TimeInterval = &timeInterval
	return &this
}

// NewCreateRfqMmpConfigV5RespDataInnerWithDefaults instantiates a new CreateRfqMmpConfigV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRfqMmpConfigV5RespDataInnerWithDefaults() *CreateRfqMmpConfigV5RespDataInner {
	this := CreateRfqMmpConfigV5RespDataInner{}
	var countLimit string = ""
	this.CountLimit = &countLimit
	var frozenInterval string = ""
	this.FrozenInterval = &frozenInterval
	var timeInterval string = ""
	this.TimeInterval = &timeInterval
	return &this
}

// GetCountLimit returns the CountLimit field value if set, zero value otherwise.
func (o *CreateRfqMmpConfigV5RespDataInner) GetCountLimit() string {
	if o == nil || IsNil(o.CountLimit) {
		var ret string
		return ret
	}
	return *o.CountLimit
}

// GetCountLimitOk returns a tuple with the CountLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqMmpConfigV5RespDataInner) GetCountLimitOk() (*string, bool) {
	if o == nil || IsNil(o.CountLimit) {
		return nil, false
	}
	return o.CountLimit, true
}

// HasCountLimit returns a boolean if a field has been set.
func (o *CreateRfqMmpConfigV5RespDataInner) HasCountLimit() bool {
	if o != nil && !IsNil(o.CountLimit) {
		return true
	}

	return false
}

// SetCountLimit gets a reference to the given string and assigns it to the CountLimit field.
func (o *CreateRfqMmpConfigV5RespDataInner) SetCountLimit(v string) {
	o.CountLimit = &v
}

// GetFrozenInterval returns the FrozenInterval field value if set, zero value otherwise.
func (o *CreateRfqMmpConfigV5RespDataInner) GetFrozenInterval() string {
	if o == nil || IsNil(o.FrozenInterval) {
		var ret string
		return ret
	}
	return *o.FrozenInterval
}

// GetFrozenIntervalOk returns a tuple with the FrozenInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqMmpConfigV5RespDataInner) GetFrozenIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.FrozenInterval) {
		return nil, false
	}
	return o.FrozenInterval, true
}

// HasFrozenInterval returns a boolean if a field has been set.
func (o *CreateRfqMmpConfigV5RespDataInner) HasFrozenInterval() bool {
	if o != nil && !IsNil(o.FrozenInterval) {
		return true
	}

	return false
}

// SetFrozenInterval gets a reference to the given string and assigns it to the FrozenInterval field.
func (o *CreateRfqMmpConfigV5RespDataInner) SetFrozenInterval(v string) {
	o.FrozenInterval = &v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *CreateRfqMmpConfigV5RespDataInner) GetTimeInterval() string {
	if o == nil || IsNil(o.TimeInterval) {
		var ret string
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqMmpConfigV5RespDataInner) GetTimeIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.TimeInterval) {
		return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *CreateRfqMmpConfigV5RespDataInner) HasTimeInterval() bool {
	if o != nil && !IsNil(o.TimeInterval) {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given string and assigns it to the TimeInterval field.
func (o *CreateRfqMmpConfigV5RespDataInner) SetTimeInterval(v string) {
	o.TimeInterval = &v
}

func (o CreateRfqMmpConfigV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRfqMmpConfigV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountLimit) {
		toSerialize["countLimit"] = o.CountLimit
	}
	if !IsNil(o.FrozenInterval) {
		toSerialize["frozenInterval"] = o.FrozenInterval
	}
	if !IsNil(o.TimeInterval) {
		toSerialize["timeInterval"] = o.TimeInterval
	}
	return toSerialize, nil
}

type NullableCreateRfqMmpConfigV5RespDataInner struct {
	value *CreateRfqMmpConfigV5RespDataInner
	isSet bool
}

func (v NullableCreateRfqMmpConfigV5RespDataInner) Get() *CreateRfqMmpConfigV5RespDataInner {
	return v.value
}

func (v *NullableCreateRfqMmpConfigV5RespDataInner) Set(val *CreateRfqMmpConfigV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRfqMmpConfigV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRfqMmpConfigV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRfqMmpConfigV5RespDataInner(val *CreateRfqMmpConfigV5RespDataInner) *NullableCreateRfqMmpConfigV5RespDataInner {
	return &NullableCreateRfqMmpConfigV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateRfqMmpConfigV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRfqMmpConfigV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


