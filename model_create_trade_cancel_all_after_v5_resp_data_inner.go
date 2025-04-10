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

// checks if the CreateTradeCancelAllAfterV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeCancelAllAfterV5RespDataInner{}

// CreateTradeCancelAllAfterV5RespDataInner struct for CreateTradeCancelAllAfterV5RespDataInner
type CreateTradeCancelAllAfterV5RespDataInner struct {
	// CAA order tag
	Tag *string `json:"tag,omitempty"`
	// The time the cancellation is triggered.  triggerTime=0 means Cancel All After is disabled.
	TriggerTime *string `json:"triggerTime,omitempty"`
	// The time the request is received.
	Ts *string `json:"ts,omitempty"`
}

// NewCreateTradeCancelAllAfterV5RespDataInner instantiates a new CreateTradeCancelAllAfterV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeCancelAllAfterV5RespDataInner() *CreateTradeCancelAllAfterV5RespDataInner {
	this := CreateTradeCancelAllAfterV5RespDataInner{}
	var tag string = ""
	this.Tag = &tag
	var triggerTime string = ""
	this.TriggerTime = &triggerTime
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewCreateTradeCancelAllAfterV5RespDataInnerWithDefaults instantiates a new CreateTradeCancelAllAfterV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeCancelAllAfterV5RespDataInnerWithDefaults() *CreateTradeCancelAllAfterV5RespDataInner {
	this := CreateTradeCancelAllAfterV5RespDataInner{}
	var tag string = ""
	this.Tag = &tag
	var triggerTime string = ""
	this.TriggerTime = &triggerTime
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateTradeCancelAllAfterV5RespDataInner) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelAllAfterV5RespDataInner) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateTradeCancelAllAfterV5RespDataInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *CreateTradeCancelAllAfterV5RespDataInner) SetTag(v string) {
	o.Tag = &v
}

// GetTriggerTime returns the TriggerTime field value if set, zero value otherwise.
func (o *CreateTradeCancelAllAfterV5RespDataInner) GetTriggerTime() string {
	if o == nil || IsNil(o.TriggerTime) {
		var ret string
		return ret
	}
	return *o.TriggerTime
}

// GetTriggerTimeOk returns a tuple with the TriggerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelAllAfterV5RespDataInner) GetTriggerTimeOk() (*string, bool) {
	if o == nil || IsNil(o.TriggerTime) {
		return nil, false
	}
	return o.TriggerTime, true
}

// HasTriggerTime returns a boolean if a field has been set.
func (o *CreateTradeCancelAllAfterV5RespDataInner) HasTriggerTime() bool {
	if o != nil && !IsNil(o.TriggerTime) {
		return true
	}

	return false
}

// SetTriggerTime gets a reference to the given string and assigns it to the TriggerTime field.
func (o *CreateTradeCancelAllAfterV5RespDataInner) SetTriggerTime(v string) {
	o.TriggerTime = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *CreateTradeCancelAllAfterV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelAllAfterV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *CreateTradeCancelAllAfterV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *CreateTradeCancelAllAfterV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o CreateTradeCancelAllAfterV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeCancelAllAfterV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TriggerTime) {
		toSerialize["triggerTime"] = o.TriggerTime
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableCreateTradeCancelAllAfterV5RespDataInner struct {
	value *CreateTradeCancelAllAfterV5RespDataInner
	isSet bool
}

func (v NullableCreateTradeCancelAllAfterV5RespDataInner) Get() *CreateTradeCancelAllAfterV5RespDataInner {
	return v.value
}

func (v *NullableCreateTradeCancelAllAfterV5RespDataInner) Set(val *CreateTradeCancelAllAfterV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeCancelAllAfterV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeCancelAllAfterV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeCancelAllAfterV5RespDataInner(val *CreateTradeCancelAllAfterV5RespDataInner) *NullableCreateTradeCancelAllAfterV5RespDataInner {
	return &NullableCreateTradeCancelAllAfterV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateTradeCancelAllAfterV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeCancelAllAfterV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


