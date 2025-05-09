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

// checks if the GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner{}

// GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner struct for GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner
type GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner struct {
	// Collateral assets data
	Assets []GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner `json:"assets,omitempty"`
}

// NewGetFinanceFlexibleLoanCollateralAssetsV5RespDataInner instantiates a new GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceFlexibleLoanCollateralAssetsV5RespDataInner() *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner {
	this := GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner{}
	return &this
}

// NewGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerWithDefaults instantiates a new GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerWithDefaults() *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner {
	this := GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) GetAssets() []GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner {
	if o == nil || IsNil(o.Assets) {
		var ret []GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) GetAssetsOk() ([]GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner and assigns it to the Assets field.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) SetAssets(v []GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) {
	o.Assets = v
}

func (o GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	return toSerialize, nil
}

type NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInner struct {
	value *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner
	isSet bool
}

func (v NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) Get() *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner {
	return v.value
}

func (v *NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) Set(val *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInner(val *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) *NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInner {
	return &NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


