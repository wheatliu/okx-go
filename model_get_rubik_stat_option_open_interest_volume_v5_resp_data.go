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

// checks if the GetRubikStatOptionOpenInterestVolumeV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRubikStatOptionOpenInterestVolumeV5RespData{}

// GetRubikStatOptionOpenInterestVolumeV5RespData struct for GetRubikStatOptionOpenInterestVolumeV5RespData
type GetRubikStatOptionOpenInterestVolumeV5RespData struct {
	// Total open interest , unit in `ccy` (in request parameter)
	Oi *string `json:"oi,omitempty"`
	// Timestamp
	Ts *string `json:"ts,omitempty"`
	// Total trading volume , unit in `ccy` (in request parameter)
	Vol *string `json:"vol,omitempty"`
}

// NewGetRubikStatOptionOpenInterestVolumeV5RespData instantiates a new GetRubikStatOptionOpenInterestVolumeV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRubikStatOptionOpenInterestVolumeV5RespData() *GetRubikStatOptionOpenInterestVolumeV5RespData {
	this := GetRubikStatOptionOpenInterestVolumeV5RespData{}
	var oi string = ""
	this.Oi = &oi
	var ts string = ""
	this.Ts = &ts
	var vol string = ""
	this.Vol = &vol
	return &this
}

// NewGetRubikStatOptionOpenInterestVolumeV5RespDataWithDefaults instantiates a new GetRubikStatOptionOpenInterestVolumeV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRubikStatOptionOpenInterestVolumeV5RespDataWithDefaults() *GetRubikStatOptionOpenInterestVolumeV5RespData {
	this := GetRubikStatOptionOpenInterestVolumeV5RespData{}
	var oi string = ""
	this.Oi = &oi
	var ts string = ""
	this.Ts = &ts
	var vol string = ""
	this.Vol = &vol
	return &this
}

// GetOi returns the Oi field value if set, zero value otherwise.
func (o *GetRubikStatOptionOpenInterestVolumeV5RespData) GetOi() string {
	if o == nil || IsNil(o.Oi) {
		var ret string
		return ret
	}
	return *o.Oi
}

// GetOiOk returns a tuple with the Oi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatOptionOpenInterestVolumeV5RespData) GetOiOk() (*string, bool) {
	if o == nil || IsNil(o.Oi) {
		return nil, false
	}
	return o.Oi, true
}

// HasOi returns a boolean if a field has been set.
func (o *GetRubikStatOptionOpenInterestVolumeV5RespData) HasOi() bool {
	if o != nil && !IsNil(o.Oi) {
		return true
	}

	return false
}

// SetOi gets a reference to the given string and assigns it to the Oi field.
func (o *GetRubikStatOptionOpenInterestVolumeV5RespData) SetOi(v string) {
	o.Oi = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetRubikStatOptionOpenInterestVolumeV5RespData) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatOptionOpenInterestVolumeV5RespData) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetRubikStatOptionOpenInterestVolumeV5RespData) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetRubikStatOptionOpenInterestVolumeV5RespData) SetTs(v string) {
	o.Ts = &v
}

// GetVol returns the Vol field value if set, zero value otherwise.
func (o *GetRubikStatOptionOpenInterestVolumeV5RespData) GetVol() string {
	if o == nil || IsNil(o.Vol) {
		var ret string
		return ret
	}
	return *o.Vol
}

// GetVolOk returns a tuple with the Vol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatOptionOpenInterestVolumeV5RespData) GetVolOk() (*string, bool) {
	if o == nil || IsNil(o.Vol) {
		return nil, false
	}
	return o.Vol, true
}

// HasVol returns a boolean if a field has been set.
func (o *GetRubikStatOptionOpenInterestVolumeV5RespData) HasVol() bool {
	if o != nil && !IsNil(o.Vol) {
		return true
	}

	return false
}

// SetVol gets a reference to the given string and assigns it to the Vol field.
func (o *GetRubikStatOptionOpenInterestVolumeV5RespData) SetVol(v string) {
	o.Vol = &v
}

func (o GetRubikStatOptionOpenInterestVolumeV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRubikStatOptionOpenInterestVolumeV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Oi) {
		toSerialize["oi"] = o.Oi
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Vol) {
		toSerialize["vol"] = o.Vol
	}
	return toSerialize, nil
}

type NullableGetRubikStatOptionOpenInterestVolumeV5RespData struct {
	value *GetRubikStatOptionOpenInterestVolumeV5RespData
	isSet bool
}

func (v NullableGetRubikStatOptionOpenInterestVolumeV5RespData) Get() *GetRubikStatOptionOpenInterestVolumeV5RespData {
	return v.value
}

func (v *NullableGetRubikStatOptionOpenInterestVolumeV5RespData) Set(val *GetRubikStatOptionOpenInterestVolumeV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRubikStatOptionOpenInterestVolumeV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRubikStatOptionOpenInterestVolumeV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRubikStatOptionOpenInterestVolumeV5RespData(val *GetRubikStatOptionOpenInterestVolumeV5RespData) *NullableGetRubikStatOptionOpenInterestVolumeV5RespData {
	return &NullableGetRubikStatOptionOpenInterestVolumeV5RespData{value: val, isSet: true}
}

func (v NullableGetRubikStatOptionOpenInterestVolumeV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRubikStatOptionOpenInterestVolumeV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


