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

// checks if the GetAccountAccountPositionRiskV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountAccountPositionRiskV5RespData{}

// GetAccountAccountPositionRiskV5RespData struct for GetAccountAccountPositionRiskV5RespData
type GetAccountAccountPositionRiskV5RespData struct {
	// Adjusted / Effective equity in `USD`  Applicable to `Multi-currency margin` and `Portfolio margin`
	AdjEq *string `json:"adjEq,omitempty"`
	// Detailed asset information in all currencies
	BalData []GetAccountAccountPositionRiskV5RespDataBalDataInner `json:"balData,omitempty"`
	// Detailed position information in all currencies
	PosData []GetAccountAccountPositionRiskV5RespDataPosDataInner `json:"posData,omitempty"`
	// Update time of account information, millisecond format of Unix timestamp, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewGetAccountAccountPositionRiskV5RespData instantiates a new GetAccountAccountPositionRiskV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountAccountPositionRiskV5RespData() *GetAccountAccountPositionRiskV5RespData {
	this := GetAccountAccountPositionRiskV5RespData{}
	var adjEq string = ""
	this.AdjEq = &adjEq
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetAccountAccountPositionRiskV5RespDataWithDefaults instantiates a new GetAccountAccountPositionRiskV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountAccountPositionRiskV5RespDataWithDefaults() *GetAccountAccountPositionRiskV5RespData {
	this := GetAccountAccountPositionRiskV5RespData{}
	var adjEq string = ""
	this.AdjEq = &adjEq
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetAdjEq returns the AdjEq field value if set, zero value otherwise.
func (o *GetAccountAccountPositionRiskV5RespData) GetAdjEq() string {
	if o == nil || IsNil(o.AdjEq) {
		var ret string
		return ret
	}
	return *o.AdjEq
}

// GetAdjEqOk returns a tuple with the AdjEq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountAccountPositionRiskV5RespData) GetAdjEqOk() (*string, bool) {
	if o == nil || IsNil(o.AdjEq) {
		return nil, false
	}
	return o.AdjEq, true
}

// HasAdjEq returns a boolean if a field has been set.
func (o *GetAccountAccountPositionRiskV5RespData) HasAdjEq() bool {
	if o != nil && !IsNil(o.AdjEq) {
		return true
	}

	return false
}

// SetAdjEq gets a reference to the given string and assigns it to the AdjEq field.
func (o *GetAccountAccountPositionRiskV5RespData) SetAdjEq(v string) {
	o.AdjEq = &v
}

// GetBalData returns the BalData field value if set, zero value otherwise.
func (o *GetAccountAccountPositionRiskV5RespData) GetBalData() []GetAccountAccountPositionRiskV5RespDataBalDataInner {
	if o == nil || IsNil(o.BalData) {
		var ret []GetAccountAccountPositionRiskV5RespDataBalDataInner
		return ret
	}
	return o.BalData
}

// GetBalDataOk returns a tuple with the BalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountAccountPositionRiskV5RespData) GetBalDataOk() ([]GetAccountAccountPositionRiskV5RespDataBalDataInner, bool) {
	if o == nil || IsNil(o.BalData) {
		return nil, false
	}
	return o.BalData, true
}

// HasBalData returns a boolean if a field has been set.
func (o *GetAccountAccountPositionRiskV5RespData) HasBalData() bool {
	if o != nil && !IsNil(o.BalData) {
		return true
	}

	return false
}

// SetBalData gets a reference to the given []GetAccountAccountPositionRiskV5RespDataBalDataInner and assigns it to the BalData field.
func (o *GetAccountAccountPositionRiskV5RespData) SetBalData(v []GetAccountAccountPositionRiskV5RespDataBalDataInner) {
	o.BalData = v
}

// GetPosData returns the PosData field value if set, zero value otherwise.
func (o *GetAccountAccountPositionRiskV5RespData) GetPosData() []GetAccountAccountPositionRiskV5RespDataPosDataInner {
	if o == nil || IsNil(o.PosData) {
		var ret []GetAccountAccountPositionRiskV5RespDataPosDataInner
		return ret
	}
	return o.PosData
}

// GetPosDataOk returns a tuple with the PosData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountAccountPositionRiskV5RespData) GetPosDataOk() ([]GetAccountAccountPositionRiskV5RespDataPosDataInner, bool) {
	if o == nil || IsNil(o.PosData) {
		return nil, false
	}
	return o.PosData, true
}

// HasPosData returns a boolean if a field has been set.
func (o *GetAccountAccountPositionRiskV5RespData) HasPosData() bool {
	if o != nil && !IsNil(o.PosData) {
		return true
	}

	return false
}

// SetPosData gets a reference to the given []GetAccountAccountPositionRiskV5RespDataPosDataInner and assigns it to the PosData field.
func (o *GetAccountAccountPositionRiskV5RespData) SetPosData(v []GetAccountAccountPositionRiskV5RespDataPosDataInner) {
	o.PosData = v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetAccountAccountPositionRiskV5RespData) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountAccountPositionRiskV5RespData) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetAccountAccountPositionRiskV5RespData) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetAccountAccountPositionRiskV5RespData) SetTs(v string) {
	o.Ts = &v
}

func (o GetAccountAccountPositionRiskV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountAccountPositionRiskV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdjEq) {
		toSerialize["adjEq"] = o.AdjEq
	}
	if !IsNil(o.BalData) {
		toSerialize["balData"] = o.BalData
	}
	if !IsNil(o.PosData) {
		toSerialize["posData"] = o.PosData
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetAccountAccountPositionRiskV5RespData struct {
	value *GetAccountAccountPositionRiskV5RespData
	isSet bool
}

func (v NullableGetAccountAccountPositionRiskV5RespData) Get() *GetAccountAccountPositionRiskV5RespData {
	return v.value
}

func (v *NullableGetAccountAccountPositionRiskV5RespData) Set(val *GetAccountAccountPositionRiskV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountAccountPositionRiskV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountAccountPositionRiskV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountAccountPositionRiskV5RespData(val *GetAccountAccountPositionRiskV5RespData) *NullableGetAccountAccountPositionRiskV5RespData {
	return &NullableGetAccountAccountPositionRiskV5RespData{value: val, isSet: true}
}

func (v NullableGetAccountAccountPositionRiskV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountAccountPositionRiskV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


