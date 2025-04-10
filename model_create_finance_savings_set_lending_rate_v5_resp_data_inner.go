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

// checks if the CreateFinanceSavingsSetLendingRateV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFinanceSavingsSetLendingRateV5RespDataInner{}

// CreateFinanceSavingsSetLendingRateV5RespDataInner struct for CreateFinanceSavingsSetLendingRateV5RespDataInner
type CreateFinanceSavingsSetLendingRateV5RespDataInner struct {
	// Currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Annual lending rate
	Rate *string `json:"rate,omitempty"`
}

// NewCreateFinanceSavingsSetLendingRateV5RespDataInner instantiates a new CreateFinanceSavingsSetLendingRateV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFinanceSavingsSetLendingRateV5RespDataInner() *CreateFinanceSavingsSetLendingRateV5RespDataInner {
	this := CreateFinanceSavingsSetLendingRateV5RespDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var rate string = ""
	this.Rate = &rate
	return &this
}

// NewCreateFinanceSavingsSetLendingRateV5RespDataInnerWithDefaults instantiates a new CreateFinanceSavingsSetLendingRateV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFinanceSavingsSetLendingRateV5RespDataInnerWithDefaults() *CreateFinanceSavingsSetLendingRateV5RespDataInner {
	this := CreateFinanceSavingsSetLendingRateV5RespDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var rate string = ""
	this.Rate = &rate
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *CreateFinanceSavingsSetLendingRateV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFinanceSavingsSetLendingRateV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *CreateFinanceSavingsSetLendingRateV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *CreateFinanceSavingsSetLendingRateV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *CreateFinanceSavingsSetLendingRateV5RespDataInner) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFinanceSavingsSetLendingRateV5RespDataInner) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *CreateFinanceSavingsSetLendingRateV5RespDataInner) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *CreateFinanceSavingsSetLendingRateV5RespDataInner) SetRate(v string) {
	o.Rate = &v
}

func (o CreateFinanceSavingsSetLendingRateV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFinanceSavingsSetLendingRateV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	return toSerialize, nil
}

type NullableCreateFinanceSavingsSetLendingRateV5RespDataInner struct {
	value *CreateFinanceSavingsSetLendingRateV5RespDataInner
	isSet bool
}

func (v NullableCreateFinanceSavingsSetLendingRateV5RespDataInner) Get() *CreateFinanceSavingsSetLendingRateV5RespDataInner {
	return v.value
}

func (v *NullableCreateFinanceSavingsSetLendingRateV5RespDataInner) Set(val *CreateFinanceSavingsSetLendingRateV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFinanceSavingsSetLendingRateV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFinanceSavingsSetLendingRateV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFinanceSavingsSetLendingRateV5RespDataInner(val *CreateFinanceSavingsSetLendingRateV5RespDataInner) *NullableCreateFinanceSavingsSetLendingRateV5RespDataInner {
	return &NullableCreateFinanceSavingsSetLendingRateV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateFinanceSavingsSetLendingRateV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFinanceSavingsSetLendingRateV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


