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

// checks if the GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner{}

// GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner struct for GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner
type GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner struct {
	// Available equity of currency
	AvailEq *string `json:"availEq,omitempty"`
	// Currency
	Ccy *string `json:"ccy,omitempty"`
	// Margin ratio of currency
	MgnRatio *string `json:"mgnRatio,omitempty"`
}

// NewGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner instantiates a new GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner() *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner {
	this := GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner{}
	var availEq string = ""
	this.AvailEq = &availEq
	var ccy string = ""
	this.Ccy = &ccy
	var mgnRatio string = ""
	this.MgnRatio = &mgnRatio
	return &this
}

// NewGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInnerWithDefaults instantiates a new GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInnerWithDefaults() *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner {
	this := GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner{}
	var availEq string = ""
	this.AvailEq = &availEq
	var ccy string = ""
	this.Ccy = &ccy
	var mgnRatio string = ""
	this.MgnRatio = &mgnRatio
	return &this
}

// GetAvailEq returns the AvailEq field value if set, zero value otherwise.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) GetAvailEq() string {
	if o == nil || IsNil(o.AvailEq) {
		var ret string
		return ret
	}
	return *o.AvailEq
}

// GetAvailEqOk returns a tuple with the AvailEq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) GetAvailEqOk() (*string, bool) {
	if o == nil || IsNil(o.AvailEq) {
		return nil, false
	}
	return o.AvailEq, true
}

// HasAvailEq returns a boolean if a field has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) HasAvailEq() bool {
	if o != nil && !IsNil(o.AvailEq) {
		return true
	}

	return false
}

// SetAvailEq gets a reference to the given string and assigns it to the AvailEq field.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) SetAvailEq(v string) {
	o.AvailEq = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetMgnRatio returns the MgnRatio field value if set, zero value otherwise.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) GetMgnRatio() string {
	if o == nil || IsNil(o.MgnRatio) {
		var ret string
		return ret
	}
	return *o.MgnRatio
}

// GetMgnRatioOk returns a tuple with the MgnRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) GetMgnRatioOk() (*string, bool) {
	if o == nil || IsNil(o.MgnRatio) {
		return nil, false
	}
	return o.MgnRatio, true
}

// HasMgnRatio returns a boolean if a field has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) HasMgnRatio() bool {
	if o != nil && !IsNil(o.MgnRatio) {
		return true
	}

	return false
}

// SetMgnRatio gets a reference to the given string and assigns it to the MgnRatio field.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) SetMgnRatio(v string) {
	o.MgnRatio = &v
}

func (o GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailEq) {
		toSerialize["availEq"] = o.AvailEq
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.MgnRatio) {
		toSerialize["mgnRatio"] = o.MgnRatio
	}
	return toSerialize, nil
}

type NullableGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner struct {
	value *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner
	isSet bool
}

func (v NullableGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) Get() *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner {
	return v.value
}

func (v *NullableGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) Set(val *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner(val *GetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) *NullableGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner {
	return &NullableGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner{value: val, isSet: true}
}

func (v NullableGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountSetAccountSwitchPrecheckV5RespDataMgnAftDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


