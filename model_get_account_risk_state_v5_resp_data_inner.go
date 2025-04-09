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

// checks if the GetAccountRiskStateV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountRiskStateV5RespDataInner{}

// GetAccountRiskStateV5RespDataInner struct for GetAccountRiskStateV5RespDataInner
type GetAccountRiskStateV5RespDataInner struct {
	// Account risk status in auto-borrow mode   true:  the account is currently in a specific risk state   false:  the account is currently not in a specific risk state
	AtRisk *bool `json:"atRisk,omitempty"`
	// derivatives risk unit list
	AtRiskIdx []string `json:"atRiskIdx,omitempty"`
	// margin risk unit list
	AtRiskMgn []string `json:"atRiskMgn,omitempty"`
	// Unix timestamp format in milliseconds, e.g.`1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewGetAccountRiskStateV5RespDataInner instantiates a new GetAccountRiskStateV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountRiskStateV5RespDataInner() *GetAccountRiskStateV5RespDataInner {
	this := GetAccountRiskStateV5RespDataInner{}
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetAccountRiskStateV5RespDataInnerWithDefaults instantiates a new GetAccountRiskStateV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountRiskStateV5RespDataInnerWithDefaults() *GetAccountRiskStateV5RespDataInner {
	this := GetAccountRiskStateV5RespDataInner{}
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetAtRisk returns the AtRisk field value if set, zero value otherwise.
func (o *GetAccountRiskStateV5RespDataInner) GetAtRisk() bool {
	if o == nil || IsNil(o.AtRisk) {
		var ret bool
		return ret
	}
	return *o.AtRisk
}

// GetAtRiskOk returns a tuple with the AtRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountRiskStateV5RespDataInner) GetAtRiskOk() (*bool, bool) {
	if o == nil || IsNil(o.AtRisk) {
		return nil, false
	}
	return o.AtRisk, true
}

// HasAtRisk returns a boolean if a field has been set.
func (o *GetAccountRiskStateV5RespDataInner) HasAtRisk() bool {
	if o != nil && !IsNil(o.AtRisk) {
		return true
	}

	return false
}

// SetAtRisk gets a reference to the given bool and assigns it to the AtRisk field.
func (o *GetAccountRiskStateV5RespDataInner) SetAtRisk(v bool) {
	o.AtRisk = &v
}

// GetAtRiskIdx returns the AtRiskIdx field value if set, zero value otherwise.
func (o *GetAccountRiskStateV5RespDataInner) GetAtRiskIdx() []string {
	if o == nil || IsNil(o.AtRiskIdx) {
		var ret []string
		return ret
	}
	return o.AtRiskIdx
}

// GetAtRiskIdxOk returns a tuple with the AtRiskIdx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountRiskStateV5RespDataInner) GetAtRiskIdxOk() ([]string, bool) {
	if o == nil || IsNil(o.AtRiskIdx) {
		return nil, false
	}
	return o.AtRiskIdx, true
}

// HasAtRiskIdx returns a boolean if a field has been set.
func (o *GetAccountRiskStateV5RespDataInner) HasAtRiskIdx() bool {
	if o != nil && !IsNil(o.AtRiskIdx) {
		return true
	}

	return false
}

// SetAtRiskIdx gets a reference to the given []string and assigns it to the AtRiskIdx field.
func (o *GetAccountRiskStateV5RespDataInner) SetAtRiskIdx(v []string) {
	o.AtRiskIdx = v
}

// GetAtRiskMgn returns the AtRiskMgn field value if set, zero value otherwise.
func (o *GetAccountRiskStateV5RespDataInner) GetAtRiskMgn() []string {
	if o == nil || IsNil(o.AtRiskMgn) {
		var ret []string
		return ret
	}
	return o.AtRiskMgn
}

// GetAtRiskMgnOk returns a tuple with the AtRiskMgn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountRiskStateV5RespDataInner) GetAtRiskMgnOk() ([]string, bool) {
	if o == nil || IsNil(o.AtRiskMgn) {
		return nil, false
	}
	return o.AtRiskMgn, true
}

// HasAtRiskMgn returns a boolean if a field has been set.
func (o *GetAccountRiskStateV5RespDataInner) HasAtRiskMgn() bool {
	if o != nil && !IsNil(o.AtRiskMgn) {
		return true
	}

	return false
}

// SetAtRiskMgn gets a reference to the given []string and assigns it to the AtRiskMgn field.
func (o *GetAccountRiskStateV5RespDataInner) SetAtRiskMgn(v []string) {
	o.AtRiskMgn = v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetAccountRiskStateV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountRiskStateV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetAccountRiskStateV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetAccountRiskStateV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetAccountRiskStateV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountRiskStateV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AtRisk) {
		toSerialize["atRisk"] = o.AtRisk
	}
	if !IsNil(o.AtRiskIdx) {
		toSerialize["atRiskIdx"] = o.AtRiskIdx
	}
	if !IsNil(o.AtRiskMgn) {
		toSerialize["atRiskMgn"] = o.AtRiskMgn
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetAccountRiskStateV5RespDataInner struct {
	value *GetAccountRiskStateV5RespDataInner
	isSet bool
}

func (v NullableGetAccountRiskStateV5RespDataInner) Get() *GetAccountRiskStateV5RespDataInner {
	return v.value
}

func (v *NullableGetAccountRiskStateV5RespDataInner) Set(val *GetAccountRiskStateV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountRiskStateV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountRiskStateV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountRiskStateV5RespDataInner(val *GetAccountRiskStateV5RespDataInner) *NullableGetAccountRiskStateV5RespDataInner {
	return &NullableGetAccountRiskStateV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetAccountRiskStateV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountRiskStateV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


