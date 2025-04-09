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

// checks if the GetCopytradingInstrumentsV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCopytradingInstrumentsV5RespDataInner{}

// GetCopytradingInstrumentsV5RespDataInner struct for GetCopytradingInstrumentsV5RespDataInner
type GetCopytradingInstrumentsV5RespDataInner struct {
	// Whether instrument is a lead instrument. `true` or `false`
	Enabled *bool `json:"enabled,omitempty"`
	// Instrument ID, e.g. BTC-USDT-SWAP
	InstId *string `json:"instId,omitempty"`
}

// NewGetCopytradingInstrumentsV5RespDataInner instantiates a new GetCopytradingInstrumentsV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCopytradingInstrumentsV5RespDataInner() *GetCopytradingInstrumentsV5RespDataInner {
	this := GetCopytradingInstrumentsV5RespDataInner{}
	var instId string = ""
	this.InstId = &instId
	return &this
}

// NewGetCopytradingInstrumentsV5RespDataInnerWithDefaults instantiates a new GetCopytradingInstrumentsV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCopytradingInstrumentsV5RespDataInnerWithDefaults() *GetCopytradingInstrumentsV5RespDataInner {
	this := GetCopytradingInstrumentsV5RespDataInner{}
	var instId string = ""
	this.InstId = &instId
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetCopytradingInstrumentsV5RespDataInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingInstrumentsV5RespDataInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetCopytradingInstrumentsV5RespDataInner) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetCopytradingInstrumentsV5RespDataInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetCopytradingInstrumentsV5RespDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingInstrumentsV5RespDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetCopytradingInstrumentsV5RespDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetCopytradingInstrumentsV5RespDataInner) SetInstId(v string) {
	o.InstId = &v
}

func (o GetCopytradingInstrumentsV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCopytradingInstrumentsV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	return toSerialize, nil
}

type NullableGetCopytradingInstrumentsV5RespDataInner struct {
	value *GetCopytradingInstrumentsV5RespDataInner
	isSet bool
}

func (v NullableGetCopytradingInstrumentsV5RespDataInner) Get() *GetCopytradingInstrumentsV5RespDataInner {
	return v.value
}

func (v *NullableGetCopytradingInstrumentsV5RespDataInner) Set(val *GetCopytradingInstrumentsV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCopytradingInstrumentsV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCopytradingInstrumentsV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCopytradingInstrumentsV5RespDataInner(val *GetCopytradingInstrumentsV5RespDataInner) *NullableGetCopytradingInstrumentsV5RespDataInner {
	return &NullableGetCopytradingInstrumentsV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetCopytradingInstrumentsV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCopytradingInstrumentsV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


