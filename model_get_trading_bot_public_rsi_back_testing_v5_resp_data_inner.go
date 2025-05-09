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

// checks if the GetTradingBotPublicRsiBackTestingV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTradingBotPublicRsiBackTestingV5RespDataInner{}

// GetTradingBotPublicRsiBackTestingV5RespDataInner struct for GetTradingBotPublicRsiBackTestingV5RespDataInner
type GetTradingBotPublicRsiBackTestingV5RespDataInner struct {
	// Trigger number
	TriggerNum *string `json:"triggerNum,omitempty"`
}

// NewGetTradingBotPublicRsiBackTestingV5RespDataInner instantiates a new GetTradingBotPublicRsiBackTestingV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTradingBotPublicRsiBackTestingV5RespDataInner() *GetTradingBotPublicRsiBackTestingV5RespDataInner {
	this := GetTradingBotPublicRsiBackTestingV5RespDataInner{}
	var triggerNum string = ""
	this.TriggerNum = &triggerNum
	return &this
}

// NewGetTradingBotPublicRsiBackTestingV5RespDataInnerWithDefaults instantiates a new GetTradingBotPublicRsiBackTestingV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTradingBotPublicRsiBackTestingV5RespDataInnerWithDefaults() *GetTradingBotPublicRsiBackTestingV5RespDataInner {
	this := GetTradingBotPublicRsiBackTestingV5RespDataInner{}
	var triggerNum string = ""
	this.TriggerNum = &triggerNum
	return &this
}

// GetTriggerNum returns the TriggerNum field value if set, zero value otherwise.
func (o *GetTradingBotPublicRsiBackTestingV5RespDataInner) GetTriggerNum() string {
	if o == nil || IsNil(o.TriggerNum) {
		var ret string
		return ret
	}
	return *o.TriggerNum
}

// GetTriggerNumOk returns a tuple with the TriggerNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotPublicRsiBackTestingV5RespDataInner) GetTriggerNumOk() (*string, bool) {
	if o == nil || IsNil(o.TriggerNum) {
		return nil, false
	}
	return o.TriggerNum, true
}

// HasTriggerNum returns a boolean if a field has been set.
func (o *GetTradingBotPublicRsiBackTestingV5RespDataInner) HasTriggerNum() bool {
	if o != nil && !IsNil(o.TriggerNum) {
		return true
	}

	return false
}

// SetTriggerNum gets a reference to the given string and assigns it to the TriggerNum field.
func (o *GetTradingBotPublicRsiBackTestingV5RespDataInner) SetTriggerNum(v string) {
	o.TriggerNum = &v
}

func (o GetTradingBotPublicRsiBackTestingV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTradingBotPublicRsiBackTestingV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TriggerNum) {
		toSerialize["triggerNum"] = o.TriggerNum
	}
	return toSerialize, nil
}

type NullableGetTradingBotPublicRsiBackTestingV5RespDataInner struct {
	value *GetTradingBotPublicRsiBackTestingV5RespDataInner
	isSet bool
}

func (v NullableGetTradingBotPublicRsiBackTestingV5RespDataInner) Get() *GetTradingBotPublicRsiBackTestingV5RespDataInner {
	return v.value
}

func (v *NullableGetTradingBotPublicRsiBackTestingV5RespDataInner) Set(val *GetTradingBotPublicRsiBackTestingV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTradingBotPublicRsiBackTestingV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTradingBotPublicRsiBackTestingV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTradingBotPublicRsiBackTestingV5RespDataInner(val *GetTradingBotPublicRsiBackTestingV5RespDataInner) *NullableGetTradingBotPublicRsiBackTestingV5RespDataInner {
	return &NullableGetTradingBotPublicRsiBackTestingV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetTradingBotPublicRsiBackTestingV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTradingBotPublicRsiBackTestingV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


