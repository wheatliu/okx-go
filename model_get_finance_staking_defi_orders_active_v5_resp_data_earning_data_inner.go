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

// checks if the GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner{}

// GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner struct for GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner
type GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner struct {
	// Earning currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Earning type  `0`: Estimated earning  `1`: Cumulative earning
	EarningType *string `json:"earningType,omitempty"`
	// Earning amount
	Earnings *string `json:"earnings,omitempty"`
}

// NewGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner instantiates a new GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner() *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner {
	this := GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var earningType string = ""
	this.EarningType = &earningType
	var earnings string = ""
	this.Earnings = &earnings
	return &this
}

// NewGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInnerWithDefaults instantiates a new GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInnerWithDefaults() *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner {
	this := GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var earningType string = ""
	this.EarningType = &earningType
	var earnings string = ""
	this.Earnings = &earnings
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetEarningType returns the EarningType field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) GetEarningType() string {
	if o == nil || IsNil(o.EarningType) {
		var ret string
		return ret
	}
	return *o.EarningType
}

// GetEarningTypeOk returns a tuple with the EarningType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) GetEarningTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EarningType) {
		return nil, false
	}
	return o.EarningType, true
}

// HasEarningType returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) HasEarningType() bool {
	if o != nil && !IsNil(o.EarningType) {
		return true
	}

	return false
}

// SetEarningType gets a reference to the given string and assigns it to the EarningType field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) SetEarningType(v string) {
	o.EarningType = &v
}

// GetEarnings returns the Earnings field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) GetEarnings() string {
	if o == nil || IsNil(o.Earnings) {
		var ret string
		return ret
	}
	return *o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) GetEarningsOk() (*string, bool) {
	if o == nil || IsNil(o.Earnings) {
		return nil, false
	}
	return o.Earnings, true
}

// HasEarnings returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) HasEarnings() bool {
	if o != nil && !IsNil(o.Earnings) {
		return true
	}

	return false
}

// SetEarnings gets a reference to the given string and assigns it to the Earnings field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) SetEarnings(v string) {
	o.Earnings = &v
}

func (o GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.EarningType) {
		toSerialize["earningType"] = o.EarningType
	}
	if !IsNil(o.Earnings) {
		toSerialize["earnings"] = o.Earnings
	}
	return toSerialize, nil
}

type NullableGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner struct {
	value *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner
	isSet bool
}

func (v NullableGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) Get() *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner {
	return v.value
}

func (v *NullableGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) Set(val *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner(val *GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) *NullableGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner {
	return &NullableGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner{value: val, isSet: true}
}

func (v NullableGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


