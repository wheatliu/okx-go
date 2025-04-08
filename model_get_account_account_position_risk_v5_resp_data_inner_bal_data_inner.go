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

// checks if the GetAccountAccountPositionRiskV5RespDataInnerBalDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountAccountPositionRiskV5RespDataInnerBalDataInner{}

// GetAccountAccountPositionRiskV5RespDataInnerBalDataInner struct for GetAccountAccountPositionRiskV5RespDataInnerBalDataInner
type GetAccountAccountPositionRiskV5RespDataInnerBalDataInner struct {
	// Currency
	Ccy *string `json:"ccy,omitempty"`
	// Discount equity of currency in `USD`.
	DisEq *string `json:"disEq,omitempty"`
	// Equity of currency
	Eq *string `json:"eq,omitempty"`
}

// NewGetAccountAccountPositionRiskV5RespDataInnerBalDataInner instantiates a new GetAccountAccountPositionRiskV5RespDataInnerBalDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountAccountPositionRiskV5RespDataInnerBalDataInner() *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner {
	this := GetAccountAccountPositionRiskV5RespDataInnerBalDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var disEq string = ""
	this.DisEq = &disEq
	var eq string = ""
	this.Eq = &eq
	return &this
}

// NewGetAccountAccountPositionRiskV5RespDataInnerBalDataInnerWithDefaults instantiates a new GetAccountAccountPositionRiskV5RespDataInnerBalDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountAccountPositionRiskV5RespDataInnerBalDataInnerWithDefaults() *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner {
	this := GetAccountAccountPositionRiskV5RespDataInnerBalDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var disEq string = ""
	this.DisEq = &disEq
	var eq string = ""
	this.Eq = &eq
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetDisEq returns the DisEq field value if set, zero value otherwise.
func (o *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) GetDisEq() string {
	if o == nil || IsNil(o.DisEq) {
		var ret string
		return ret
	}
	return *o.DisEq
}

// GetDisEqOk returns a tuple with the DisEq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) GetDisEqOk() (*string, bool) {
	if o == nil || IsNil(o.DisEq) {
		return nil, false
	}
	return o.DisEq, true
}

// HasDisEq returns a boolean if a field has been set.
func (o *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) HasDisEq() bool {
	if o != nil && !IsNil(o.DisEq) {
		return true
	}

	return false
}

// SetDisEq gets a reference to the given string and assigns it to the DisEq field.
func (o *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) SetDisEq(v string) {
	o.DisEq = &v
}

// GetEq returns the Eq field value if set, zero value otherwise.
func (o *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) GetEq() string {
	if o == nil || IsNil(o.Eq) {
		var ret string
		return ret
	}
	return *o.Eq
}

// GetEqOk returns a tuple with the Eq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) GetEqOk() (*string, bool) {
	if o == nil || IsNil(o.Eq) {
		return nil, false
	}
	return o.Eq, true
}

// HasEq returns a boolean if a field has been set.
func (o *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) HasEq() bool {
	if o != nil && !IsNil(o.Eq) {
		return true
	}

	return false
}

// SetEq gets a reference to the given string and assigns it to the Eq field.
func (o *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) SetEq(v string) {
	o.Eq = &v
}

func (o GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.DisEq) {
		toSerialize["disEq"] = o.DisEq
	}
	if !IsNil(o.Eq) {
		toSerialize["eq"] = o.Eq
	}
	return toSerialize, nil
}

type NullableGetAccountAccountPositionRiskV5RespDataInnerBalDataInner struct {
	value *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner
	isSet bool
}

func (v NullableGetAccountAccountPositionRiskV5RespDataInnerBalDataInner) Get() *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner {
	return v.value
}

func (v *NullableGetAccountAccountPositionRiskV5RespDataInnerBalDataInner) Set(val *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountAccountPositionRiskV5RespDataInnerBalDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountAccountPositionRiskV5RespDataInnerBalDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountAccountPositionRiskV5RespDataInnerBalDataInner(val *GetAccountAccountPositionRiskV5RespDataInnerBalDataInner) *NullableGetAccountAccountPositionRiskV5RespDataInnerBalDataInner {
	return &NullableGetAccountAccountPositionRiskV5RespDataInnerBalDataInner{value: val, isSet: true}
}

func (v NullableGetAccountAccountPositionRiskV5RespDataInnerBalDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountAccountPositionRiskV5RespDataInnerBalDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


