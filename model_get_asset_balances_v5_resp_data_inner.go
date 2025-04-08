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

// checks if the GetAssetBalancesV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAssetBalancesV5RespDataInner{}

// GetAssetBalancesV5RespDataInner struct for GetAssetBalancesV5RespDataInner
type GetAssetBalancesV5RespDataInner struct {
	// Available balance
	AvailBal *string `json:"availBal,omitempty"`
	// Balance
	Bal *string `json:"bal,omitempty"`
	// Currency
	Ccy *string `json:"ccy,omitempty"`
	// Frozen balance
	FrozenBal *string `json:"frozenBal,omitempty"`
}

// NewGetAssetBalancesV5RespDataInner instantiates a new GetAssetBalancesV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssetBalancesV5RespDataInner() *GetAssetBalancesV5RespDataInner {
	this := GetAssetBalancesV5RespDataInner{}
	var availBal string = ""
	this.AvailBal = &availBal
	var bal string = ""
	this.Bal = &bal
	var ccy string = ""
	this.Ccy = &ccy
	var frozenBal string = ""
	this.FrozenBal = &frozenBal
	return &this
}

// NewGetAssetBalancesV5RespDataInnerWithDefaults instantiates a new GetAssetBalancesV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssetBalancesV5RespDataInnerWithDefaults() *GetAssetBalancesV5RespDataInner {
	this := GetAssetBalancesV5RespDataInner{}
	var availBal string = ""
	this.AvailBal = &availBal
	var bal string = ""
	this.Bal = &bal
	var ccy string = ""
	this.Ccy = &ccy
	var frozenBal string = ""
	this.FrozenBal = &frozenBal
	return &this
}

// GetAvailBal returns the AvailBal field value if set, zero value otherwise.
func (o *GetAssetBalancesV5RespDataInner) GetAvailBal() string {
	if o == nil || IsNil(o.AvailBal) {
		var ret string
		return ret
	}
	return *o.AvailBal
}

// GetAvailBalOk returns a tuple with the AvailBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetBalancesV5RespDataInner) GetAvailBalOk() (*string, bool) {
	if o == nil || IsNil(o.AvailBal) {
		return nil, false
	}
	return o.AvailBal, true
}

// HasAvailBal returns a boolean if a field has been set.
func (o *GetAssetBalancesV5RespDataInner) HasAvailBal() bool {
	if o != nil && !IsNil(o.AvailBal) {
		return true
	}

	return false
}

// SetAvailBal gets a reference to the given string and assigns it to the AvailBal field.
func (o *GetAssetBalancesV5RespDataInner) SetAvailBal(v string) {
	o.AvailBal = &v
}

// GetBal returns the Bal field value if set, zero value otherwise.
func (o *GetAssetBalancesV5RespDataInner) GetBal() string {
	if o == nil || IsNil(o.Bal) {
		var ret string
		return ret
	}
	return *o.Bal
}

// GetBalOk returns a tuple with the Bal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetBalancesV5RespDataInner) GetBalOk() (*string, bool) {
	if o == nil || IsNil(o.Bal) {
		return nil, false
	}
	return o.Bal, true
}

// HasBal returns a boolean if a field has been set.
func (o *GetAssetBalancesV5RespDataInner) HasBal() bool {
	if o != nil && !IsNil(o.Bal) {
		return true
	}

	return false
}

// SetBal gets a reference to the given string and assigns it to the Bal field.
func (o *GetAssetBalancesV5RespDataInner) SetBal(v string) {
	o.Bal = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetAssetBalancesV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetBalancesV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetAssetBalancesV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetAssetBalancesV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetFrozenBal returns the FrozenBal field value if set, zero value otherwise.
func (o *GetAssetBalancesV5RespDataInner) GetFrozenBal() string {
	if o == nil || IsNil(o.FrozenBal) {
		var ret string
		return ret
	}
	return *o.FrozenBal
}

// GetFrozenBalOk returns a tuple with the FrozenBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetBalancesV5RespDataInner) GetFrozenBalOk() (*string, bool) {
	if o == nil || IsNil(o.FrozenBal) {
		return nil, false
	}
	return o.FrozenBal, true
}

// HasFrozenBal returns a boolean if a field has been set.
func (o *GetAssetBalancesV5RespDataInner) HasFrozenBal() bool {
	if o != nil && !IsNil(o.FrozenBal) {
		return true
	}

	return false
}

// SetFrozenBal gets a reference to the given string and assigns it to the FrozenBal field.
func (o *GetAssetBalancesV5RespDataInner) SetFrozenBal(v string) {
	o.FrozenBal = &v
}

func (o GetAssetBalancesV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssetBalancesV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailBal) {
		toSerialize["availBal"] = o.AvailBal
	}
	if !IsNil(o.Bal) {
		toSerialize["bal"] = o.Bal
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.FrozenBal) {
		toSerialize["frozenBal"] = o.FrozenBal
	}
	return toSerialize, nil
}

type NullableGetAssetBalancesV5RespDataInner struct {
	value *GetAssetBalancesV5RespDataInner
	isSet bool
}

func (v NullableGetAssetBalancesV5RespDataInner) Get() *GetAssetBalancesV5RespDataInner {
	return v.value
}

func (v *NullableGetAssetBalancesV5RespDataInner) Set(val *GetAssetBalancesV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetBalancesV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetBalancesV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetBalancesV5RespDataInner(val *GetAssetBalancesV5RespDataInner) *NullableGetAssetBalancesV5RespDataInner {
	return &NullableGetAssetBalancesV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetAssetBalancesV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetBalancesV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


