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

// checks if the GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner{}

// GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner struct for GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner
type GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner struct {
	// Long/Short open interest ratio
	OiRatio *string `json:"oiRatio,omitempty"`
	// Timestamp of data generation time
	Ts *string `json:"ts,omitempty"`
	// Long/Short trading volume ratio
	VolRatio *string `json:"volRatio,omitempty"`
}

// NewGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner instantiates a new GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner() *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner {
	this := GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner{}
	var oiRatio string = ""
	this.OiRatio = &oiRatio
	var ts string = ""
	this.Ts = &ts
	var volRatio string = ""
	this.VolRatio = &volRatio
	return &this
}

// NewGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInnerWithDefaults instantiates a new GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInnerWithDefaults() *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner {
	this := GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner{}
	var oiRatio string = ""
	this.OiRatio = &oiRatio
	var ts string = ""
	this.Ts = &ts
	var volRatio string = ""
	this.VolRatio = &volRatio
	return &this
}

// GetOiRatio returns the OiRatio field value if set, zero value otherwise.
func (o *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) GetOiRatio() string {
	if o == nil || IsNil(o.OiRatio) {
		var ret string
		return ret
	}
	return *o.OiRatio
}

// GetOiRatioOk returns a tuple with the OiRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) GetOiRatioOk() (*string, bool) {
	if o == nil || IsNil(o.OiRatio) {
		return nil, false
	}
	return o.OiRatio, true
}

// HasOiRatio returns a boolean if a field has been set.
func (o *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) HasOiRatio() bool {
	if o != nil && !IsNil(o.OiRatio) {
		return true
	}

	return false
}

// SetOiRatio gets a reference to the given string and assigns it to the OiRatio field.
func (o *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) SetOiRatio(v string) {
	o.OiRatio = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

// GetVolRatio returns the VolRatio field value if set, zero value otherwise.
func (o *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) GetVolRatio() string {
	if o == nil || IsNil(o.VolRatio) {
		var ret string
		return ret
	}
	return *o.VolRatio
}

// GetVolRatioOk returns a tuple with the VolRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) GetVolRatioOk() (*string, bool) {
	if o == nil || IsNil(o.VolRatio) {
		return nil, false
	}
	return o.VolRatio, true
}

// HasVolRatio returns a boolean if a field has been set.
func (o *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) HasVolRatio() bool {
	if o != nil && !IsNil(o.VolRatio) {
		return true
	}

	return false
}

// SetVolRatio gets a reference to the given string and assigns it to the VolRatio field.
func (o *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) SetVolRatio(v string) {
	o.VolRatio = &v
}

func (o GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OiRatio) {
		toSerialize["oiRatio"] = o.OiRatio
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.VolRatio) {
		toSerialize["volRatio"] = o.VolRatio
	}
	return toSerialize, nil
}

type NullableGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner struct {
	value *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner
	isSet bool
}

func (v NullableGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) Get() *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner {
	return v.value
}

func (v *NullableGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) Set(val *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner(val *GetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) *NullableGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner {
	return &NullableGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRubikStatOptionOpenInterestVolumeRatioV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


