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

// checks if the GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner{}

// GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner struct for GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner
type GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner struct {
	// Earning currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Earning type  `0`: Estimated earning  `1`: Cumulative earning
	EarningType *string `json:"earningType,omitempty"`
	// Cumulative earning of redeemed orders  This field is just valid when the order is in redemption state
	RealizedEarnings *string `json:"realizedEarnings,omitempty"`
}

// NewGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner instantiates a new GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner() *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner {
	this := GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var earningType string = ""
	this.EarningType = &earningType
	var realizedEarnings string = ""
	this.RealizedEarnings = &realizedEarnings
	return &this
}

// NewGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInnerWithDefaults instantiates a new GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInnerWithDefaults() *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner {
	this := GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var earningType string = ""
	this.EarningType = &earningType
	var realizedEarnings string = ""
	this.RealizedEarnings = &realizedEarnings
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetEarningType returns the EarningType field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) GetEarningType() string {
	if o == nil || IsNil(o.EarningType) {
		var ret string
		return ret
	}
	return *o.EarningType
}

// GetEarningTypeOk returns a tuple with the EarningType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) GetEarningTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EarningType) {
		return nil, false
	}
	return o.EarningType, true
}

// HasEarningType returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) HasEarningType() bool {
	if o != nil && !IsNil(o.EarningType) {
		return true
	}

	return false
}

// SetEarningType gets a reference to the given string and assigns it to the EarningType field.
func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) SetEarningType(v string) {
	o.EarningType = &v
}

// GetRealizedEarnings returns the RealizedEarnings field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) GetRealizedEarnings() string {
	if o == nil || IsNil(o.RealizedEarnings) {
		var ret string
		return ret
	}
	return *o.RealizedEarnings
}

// GetRealizedEarningsOk returns a tuple with the RealizedEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) GetRealizedEarningsOk() (*string, bool) {
	if o == nil || IsNil(o.RealizedEarnings) {
		return nil, false
	}
	return o.RealizedEarnings, true
}

// HasRealizedEarnings returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) HasRealizedEarnings() bool {
	if o != nil && !IsNil(o.RealizedEarnings) {
		return true
	}

	return false
}

// SetRealizedEarnings gets a reference to the given string and assigns it to the RealizedEarnings field.
func (o *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) SetRealizedEarnings(v string) {
	o.RealizedEarnings = &v
}

func (o GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.EarningType) {
		toSerialize["earningType"] = o.EarningType
	}
	if !IsNil(o.RealizedEarnings) {
		toSerialize["realizedEarnings"] = o.RealizedEarnings
	}
	return toSerialize, nil
}

type NullableGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner struct {
	value *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner
	isSet bool
}

func (v NullableGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) Get() *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner {
	return v.value
}

func (v *NullableGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) Set(val *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner(val *GetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) *NullableGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner {
	return &NullableGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner{value: val, isSet: true}
}

func (v NullableGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceStakingDefiOrdersHistoryV5RespDataInnerEarningDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


