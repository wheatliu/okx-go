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

// checks if the GetAccountInterestRateV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountInterestRateV5RespData{}

// GetAccountInterestRateV5RespData struct for GetAccountInterestRateV5RespData
type GetAccountInterestRateV5RespData struct {
	// currency
	Ccy *string `json:"ccy,omitempty"`
	// interest rate(the current hour)
	InterestRate *string `json:"interestRate,omitempty"`
}

// NewGetAccountInterestRateV5RespData instantiates a new GetAccountInterestRateV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountInterestRateV5RespData() *GetAccountInterestRateV5RespData {
	this := GetAccountInterestRateV5RespData{}
	var ccy string = ""
	this.Ccy = &ccy
	var interestRate string = ""
	this.InterestRate = &interestRate
	return &this
}

// NewGetAccountInterestRateV5RespDataWithDefaults instantiates a new GetAccountInterestRateV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountInterestRateV5RespDataWithDefaults() *GetAccountInterestRateV5RespData {
	this := GetAccountInterestRateV5RespData{}
	var ccy string = ""
	this.Ccy = &ccy
	var interestRate string = ""
	this.InterestRate = &interestRate
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetAccountInterestRateV5RespData) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestRateV5RespData) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetAccountInterestRateV5RespData) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetAccountInterestRateV5RespData) SetCcy(v string) {
	o.Ccy = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *GetAccountInterestRateV5RespData) GetInterestRate() string {
	if o == nil || IsNil(o.InterestRate) {
		var ret string
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestRateV5RespData) GetInterestRateOk() (*string, bool) {
	if o == nil || IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *GetAccountInterestRateV5RespData) HasInterestRate() bool {
	if o != nil && !IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given string and assigns it to the InterestRate field.
func (o *GetAccountInterestRateV5RespData) SetInterestRate(v string) {
	o.InterestRate = &v
}

func (o GetAccountInterestRateV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountInterestRateV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.InterestRate) {
		toSerialize["interestRate"] = o.InterestRate
	}
	return toSerialize, nil
}

type NullableGetAccountInterestRateV5RespData struct {
	value *GetAccountInterestRateV5RespData
	isSet bool
}

func (v NullableGetAccountInterestRateV5RespData) Get() *GetAccountInterestRateV5RespData {
	return v.value
}

func (v *NullableGetAccountInterestRateV5RespData) Set(val *GetAccountInterestRateV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountInterestRateV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountInterestRateV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountInterestRateV5RespData(val *GetAccountInterestRateV5RespData) *NullableGetAccountInterestRateV5RespData {
	return &NullableGetAccountInterestRateV5RespData{value: val, isSet: true}
}

func (v NullableGetAccountInterestRateV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountInterestRateV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


