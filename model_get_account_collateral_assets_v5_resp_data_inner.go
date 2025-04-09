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

// checks if the GetAccountCollateralAssetsV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountCollateralAssetsV5RespDataInner{}

// GetAccountCollateralAssetsV5RespDataInner struct for GetAccountCollateralAssetsV5RespDataInner
type GetAccountCollateralAssetsV5RespDataInner struct {
	// Currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Whether or not to be a collateral asset
	CollateralEnabled *bool `json:"collateralEnabled,omitempty"`
}

// NewGetAccountCollateralAssetsV5RespDataInner instantiates a new GetAccountCollateralAssetsV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountCollateralAssetsV5RespDataInner() *GetAccountCollateralAssetsV5RespDataInner {
	this := GetAccountCollateralAssetsV5RespDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	return &this
}

// NewGetAccountCollateralAssetsV5RespDataInnerWithDefaults instantiates a new GetAccountCollateralAssetsV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountCollateralAssetsV5RespDataInnerWithDefaults() *GetAccountCollateralAssetsV5RespDataInner {
	this := GetAccountCollateralAssetsV5RespDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetAccountCollateralAssetsV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountCollateralAssetsV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetAccountCollateralAssetsV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetAccountCollateralAssetsV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetCollateralEnabled returns the CollateralEnabled field value if set, zero value otherwise.
func (o *GetAccountCollateralAssetsV5RespDataInner) GetCollateralEnabled() bool {
	if o == nil || IsNil(o.CollateralEnabled) {
		var ret bool
		return ret
	}
	return *o.CollateralEnabled
}

// GetCollateralEnabledOk returns a tuple with the CollateralEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountCollateralAssetsV5RespDataInner) GetCollateralEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CollateralEnabled) {
		return nil, false
	}
	return o.CollateralEnabled, true
}

// HasCollateralEnabled returns a boolean if a field has been set.
func (o *GetAccountCollateralAssetsV5RespDataInner) HasCollateralEnabled() bool {
	if o != nil && !IsNil(o.CollateralEnabled) {
		return true
	}

	return false
}

// SetCollateralEnabled gets a reference to the given bool and assigns it to the CollateralEnabled field.
func (o *GetAccountCollateralAssetsV5RespDataInner) SetCollateralEnabled(v bool) {
	o.CollateralEnabled = &v
}

func (o GetAccountCollateralAssetsV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountCollateralAssetsV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.CollateralEnabled) {
		toSerialize["collateralEnabled"] = o.CollateralEnabled
	}
	return toSerialize, nil
}

type NullableGetAccountCollateralAssetsV5RespDataInner struct {
	value *GetAccountCollateralAssetsV5RespDataInner
	isSet bool
}

func (v NullableGetAccountCollateralAssetsV5RespDataInner) Get() *GetAccountCollateralAssetsV5RespDataInner {
	return v.value
}

func (v *NullableGetAccountCollateralAssetsV5RespDataInner) Set(val *GetAccountCollateralAssetsV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountCollateralAssetsV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountCollateralAssetsV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountCollateralAssetsV5RespDataInner(val *GetAccountCollateralAssetsV5RespDataInner) *NullableGetAccountCollateralAssetsV5RespDataInner {
	return &NullableGetAccountCollateralAssetsV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetAccountCollateralAssetsV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountCollateralAssetsV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


