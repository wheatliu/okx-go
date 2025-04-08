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

// checks if the GetMarketPlatform24VolumeV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarketPlatform24VolumeV5RespDataInner{}

// GetMarketPlatform24VolumeV5RespDataInner struct for GetMarketPlatform24VolumeV5RespDataInner
type GetMarketPlatform24VolumeV5RespDataInner struct {
	// Data return time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
	// 24-hour total trading volume from the order book trading in \"CNY\"
	VolCny *string `json:"volCny,omitempty"`
	// 24-hour total trading volume from the order book trading in \"USD\"
	VolUsd *string `json:"volUsd,omitempty"`
}

// NewGetMarketPlatform24VolumeV5RespDataInner instantiates a new GetMarketPlatform24VolumeV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketPlatform24VolumeV5RespDataInner() *GetMarketPlatform24VolumeV5RespDataInner {
	this := GetMarketPlatform24VolumeV5RespDataInner{}
	var ts string = ""
	this.Ts = &ts
	var volCny string = ""
	this.VolCny = &volCny
	var volUsd string = ""
	this.VolUsd = &volUsd
	return &this
}

// NewGetMarketPlatform24VolumeV5RespDataInnerWithDefaults instantiates a new GetMarketPlatform24VolumeV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketPlatform24VolumeV5RespDataInnerWithDefaults() *GetMarketPlatform24VolumeV5RespDataInner {
	this := GetMarketPlatform24VolumeV5RespDataInner{}
	var ts string = ""
	this.Ts = &ts
	var volCny string = ""
	this.VolCny = &volCny
	var volUsd string = ""
	this.VolUsd = &volUsd
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetMarketPlatform24VolumeV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketPlatform24VolumeV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetMarketPlatform24VolumeV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetMarketPlatform24VolumeV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

// GetVolCny returns the VolCny field value if set, zero value otherwise.
func (o *GetMarketPlatform24VolumeV5RespDataInner) GetVolCny() string {
	if o == nil || IsNil(o.VolCny) {
		var ret string
		return ret
	}
	return *o.VolCny
}

// GetVolCnyOk returns a tuple with the VolCny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketPlatform24VolumeV5RespDataInner) GetVolCnyOk() (*string, bool) {
	if o == nil || IsNil(o.VolCny) {
		return nil, false
	}
	return o.VolCny, true
}

// HasVolCny returns a boolean if a field has been set.
func (o *GetMarketPlatform24VolumeV5RespDataInner) HasVolCny() bool {
	if o != nil && !IsNil(o.VolCny) {
		return true
	}

	return false
}

// SetVolCny gets a reference to the given string and assigns it to the VolCny field.
func (o *GetMarketPlatform24VolumeV5RespDataInner) SetVolCny(v string) {
	o.VolCny = &v
}

// GetVolUsd returns the VolUsd field value if set, zero value otherwise.
func (o *GetMarketPlatform24VolumeV5RespDataInner) GetVolUsd() string {
	if o == nil || IsNil(o.VolUsd) {
		var ret string
		return ret
	}
	return *o.VolUsd
}

// GetVolUsdOk returns a tuple with the VolUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketPlatform24VolumeV5RespDataInner) GetVolUsdOk() (*string, bool) {
	if o == nil || IsNil(o.VolUsd) {
		return nil, false
	}
	return o.VolUsd, true
}

// HasVolUsd returns a boolean if a field has been set.
func (o *GetMarketPlatform24VolumeV5RespDataInner) HasVolUsd() bool {
	if o != nil && !IsNil(o.VolUsd) {
		return true
	}

	return false
}

// SetVolUsd gets a reference to the given string and assigns it to the VolUsd field.
func (o *GetMarketPlatform24VolumeV5RespDataInner) SetVolUsd(v string) {
	o.VolUsd = &v
}

func (o GetMarketPlatform24VolumeV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketPlatform24VolumeV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.VolCny) {
		toSerialize["volCny"] = o.VolCny
	}
	if !IsNil(o.VolUsd) {
		toSerialize["volUsd"] = o.VolUsd
	}
	return toSerialize, nil
}

type NullableGetMarketPlatform24VolumeV5RespDataInner struct {
	value *GetMarketPlatform24VolumeV5RespDataInner
	isSet bool
}

func (v NullableGetMarketPlatform24VolumeV5RespDataInner) Get() *GetMarketPlatform24VolumeV5RespDataInner {
	return v.value
}

func (v *NullableGetMarketPlatform24VolumeV5RespDataInner) Set(val *GetMarketPlatform24VolumeV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketPlatform24VolumeV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketPlatform24VolumeV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketPlatform24VolumeV5RespDataInner(val *GetMarketPlatform24VolumeV5RespDataInner) *NullableGetMarketPlatform24VolumeV5RespDataInner {
	return &NullableGetMarketPlatform24VolumeV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetMarketPlatform24VolumeV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketPlatform24VolumeV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


