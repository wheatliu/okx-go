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

// checks if the GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner{}

// GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner struct for GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner
type GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner struct {
	// Leverage of cross margin contract positions after switch
	Lever *string `json:"lever,omitempty"`
	// Position ID
	PosId *string `json:"posId,omitempty"`
}

// NewGetAccountSetAccountSwitchPrecheckV5RespDataPosListInner instantiates a new GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountSetAccountSwitchPrecheckV5RespDataPosListInner() *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner {
	this := GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner{}
	var lever string = ""
	this.Lever = &lever
	var posId string = ""
	this.PosId = &posId
	return &this
}

// NewGetAccountSetAccountSwitchPrecheckV5RespDataPosListInnerWithDefaults instantiates a new GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountSetAccountSwitchPrecheckV5RespDataPosListInnerWithDefaults() *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner {
	this := GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner{}
	var lever string = ""
	this.Lever = &lever
	var posId string = ""
	this.PosId = &posId
	return &this
}

// GetLever returns the Lever field value if set, zero value otherwise.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) GetLever() string {
	if o == nil || IsNil(o.Lever) {
		var ret string
		return ret
	}
	return *o.Lever
}

// GetLeverOk returns a tuple with the Lever field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) GetLeverOk() (*string, bool) {
	if o == nil || IsNil(o.Lever) {
		return nil, false
	}
	return o.Lever, true
}

// HasLever returns a boolean if a field has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) HasLever() bool {
	if o != nil && !IsNil(o.Lever) {
		return true
	}

	return false
}

// SetLever gets a reference to the given string and assigns it to the Lever field.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) SetLever(v string) {
	o.Lever = &v
}

// GetPosId returns the PosId field value if set, zero value otherwise.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) GetPosId() string {
	if o == nil || IsNil(o.PosId) {
		var ret string
		return ret
	}
	return *o.PosId
}

// GetPosIdOk returns a tuple with the PosId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) GetPosIdOk() (*string, bool) {
	if o == nil || IsNil(o.PosId) {
		return nil, false
	}
	return o.PosId, true
}

// HasPosId returns a boolean if a field has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) HasPosId() bool {
	if o != nil && !IsNil(o.PosId) {
		return true
	}

	return false
}

// SetPosId gets a reference to the given string and assigns it to the PosId field.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) SetPosId(v string) {
	o.PosId = &v
}

func (o GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lever) {
		toSerialize["lever"] = o.Lever
	}
	if !IsNil(o.PosId) {
		toSerialize["posId"] = o.PosId
	}
	return toSerialize, nil
}

type NullableGetAccountSetAccountSwitchPrecheckV5RespDataPosListInner struct {
	value *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner
	isSet bool
}

func (v NullableGetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) Get() *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner {
	return v.value
}

func (v *NullableGetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) Set(val *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountSetAccountSwitchPrecheckV5RespDataPosListInner(val *GetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) *NullableGetAccountSetAccountSwitchPrecheckV5RespDataPosListInner {
	return &NullableGetAccountSetAccountSwitchPrecheckV5RespDataPosListInner{value: val, isSet: true}
}

func (v NullableGetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountSetAccountSwitchPrecheckV5RespDataPosListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


