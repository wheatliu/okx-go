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

// checks if the GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner{}

// GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner struct for GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner
type GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner struct {
	// Earning currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Earning type  `0`: Estimated earning  `1`: Cumulative earning
	EarningType *string `json:"earningType,omitempty"`
}

// NewGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner instantiates a new GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner() *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner {
	this := GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var earningType string = ""
	this.EarningType = &earningType
	return &this
}

// NewGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInnerWithDefaults instantiates a new GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInnerWithDefaults() *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner {
	this := GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var earningType string = ""
	this.EarningType = &earningType
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetEarningType returns the EarningType field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) GetEarningType() string {
	if o == nil || IsNil(o.EarningType) {
		var ret string
		return ret
	}
	return *o.EarningType
}

// GetEarningTypeOk returns a tuple with the EarningType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) GetEarningTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EarningType) {
		return nil, false
	}
	return o.EarningType, true
}

// HasEarningType returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) HasEarningType() bool {
	if o != nil && !IsNil(o.EarningType) {
		return true
	}

	return false
}

// SetEarningType gets a reference to the given string and assigns it to the EarningType field.
func (o *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) SetEarningType(v string) {
	o.EarningType = &v
}

func (o GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.EarningType) {
		toSerialize["earningType"] = o.EarningType
	}
	return toSerialize, nil
}

type NullableGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner struct {
	value *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner
	isSet bool
}

func (v NullableGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) Get() *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner {
	return v.value
}

func (v *NullableGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) Set(val *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner(val *GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) *NullableGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner {
	return &NullableGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner{value: val, isSet: true}
}

func (v NullableGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


