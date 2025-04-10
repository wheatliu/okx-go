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

// checks if the GetFinanceStakingDefiEthApyHistoryV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceStakingDefiEthApyHistoryV5RespDataInner{}

// GetFinanceStakingDefiEthApyHistoryV5RespDataInner struct for GetFinanceStakingDefiEthApyHistoryV5RespDataInner
type GetFinanceStakingDefiEthApyHistoryV5RespDataInner struct {
	// APY(Annual percentage yield), e.g. `0.01` represents `1%`
	Rate *string `json:"rate,omitempty"`
	// Data time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewGetFinanceStakingDefiEthApyHistoryV5RespDataInner instantiates a new GetFinanceStakingDefiEthApyHistoryV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceStakingDefiEthApyHistoryV5RespDataInner() *GetFinanceStakingDefiEthApyHistoryV5RespDataInner {
	this := GetFinanceStakingDefiEthApyHistoryV5RespDataInner{}
	var rate string = ""
	this.Rate = &rate
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetFinanceStakingDefiEthApyHistoryV5RespDataInnerWithDefaults instantiates a new GetFinanceStakingDefiEthApyHistoryV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceStakingDefiEthApyHistoryV5RespDataInnerWithDefaults() *GetFinanceStakingDefiEthApyHistoryV5RespDataInner {
	this := GetFinanceStakingDefiEthApyHistoryV5RespDataInner{}
	var rate string = ""
	this.Rate = &rate
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiEthApyHistoryV5RespDataInner) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiEthApyHistoryV5RespDataInner) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiEthApyHistoryV5RespDataInner) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *GetFinanceStakingDefiEthApyHistoryV5RespDataInner) SetRate(v string) {
	o.Rate = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiEthApyHistoryV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiEthApyHistoryV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiEthApyHistoryV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetFinanceStakingDefiEthApyHistoryV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetFinanceStakingDefiEthApyHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceStakingDefiEthApyHistoryV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetFinanceStakingDefiEthApyHistoryV5RespDataInner struct {
	value *GetFinanceStakingDefiEthApyHistoryV5RespDataInner
	isSet bool
}

func (v NullableGetFinanceStakingDefiEthApyHistoryV5RespDataInner) Get() *GetFinanceStakingDefiEthApyHistoryV5RespDataInner {
	return v.value
}

func (v *NullableGetFinanceStakingDefiEthApyHistoryV5RespDataInner) Set(val *GetFinanceStakingDefiEthApyHistoryV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceStakingDefiEthApyHistoryV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceStakingDefiEthApyHistoryV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceStakingDefiEthApyHistoryV5RespDataInner(val *GetFinanceStakingDefiEthApyHistoryV5RespDataInner) *NullableGetFinanceStakingDefiEthApyHistoryV5RespDataInner {
	return &NullableGetFinanceStakingDefiEthApyHistoryV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetFinanceStakingDefiEthApyHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceStakingDefiEthApyHistoryV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


