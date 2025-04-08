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

// checks if the CreateCopytradingSetInstrumentsV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCopytradingSetInstrumentsV5RespData{}

// CreateCopytradingSetInstrumentsV5RespData struct for CreateCopytradingSetInstrumentsV5RespData
type CreateCopytradingSetInstrumentsV5RespData struct {
	// Whether you set it successfully   `true` or `false`
	Enabled *bool `json:"enabled,omitempty"`
	// Instrument ID, e.g. BTC-USDT-SWAP
	InstId *string `json:"instId,omitempty"`
}

// NewCreateCopytradingSetInstrumentsV5RespData instantiates a new CreateCopytradingSetInstrumentsV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCopytradingSetInstrumentsV5RespData() *CreateCopytradingSetInstrumentsV5RespData {
	this := CreateCopytradingSetInstrumentsV5RespData{}
	var instId string = ""
	this.InstId = &instId
	return &this
}

// NewCreateCopytradingSetInstrumentsV5RespDataWithDefaults instantiates a new CreateCopytradingSetInstrumentsV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCopytradingSetInstrumentsV5RespDataWithDefaults() *CreateCopytradingSetInstrumentsV5RespData {
	this := CreateCopytradingSetInstrumentsV5RespData{}
	var instId string = ""
	this.InstId = &instId
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateCopytradingSetInstrumentsV5RespData) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingSetInstrumentsV5RespData) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateCopytradingSetInstrumentsV5RespData) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateCopytradingSetInstrumentsV5RespData) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *CreateCopytradingSetInstrumentsV5RespData) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingSetInstrumentsV5RespData) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *CreateCopytradingSetInstrumentsV5RespData) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *CreateCopytradingSetInstrumentsV5RespData) SetInstId(v string) {
	o.InstId = &v
}

func (o CreateCopytradingSetInstrumentsV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCopytradingSetInstrumentsV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	return toSerialize, nil
}

type NullableCreateCopytradingSetInstrumentsV5RespData struct {
	value *CreateCopytradingSetInstrumentsV5RespData
	isSet bool
}

func (v NullableCreateCopytradingSetInstrumentsV5RespData) Get() *CreateCopytradingSetInstrumentsV5RespData {
	return v.value
}

func (v *NullableCreateCopytradingSetInstrumentsV5RespData) Set(val *CreateCopytradingSetInstrumentsV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCopytradingSetInstrumentsV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCopytradingSetInstrumentsV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCopytradingSetInstrumentsV5RespData(val *CreateCopytradingSetInstrumentsV5RespData) *NullableCreateCopytradingSetInstrumentsV5RespData {
	return &NullableCreateCopytradingSetInstrumentsV5RespData{value: val, isSet: true}
}

func (v NullableCreateCopytradingSetInstrumentsV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCopytradingSetInstrumentsV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


