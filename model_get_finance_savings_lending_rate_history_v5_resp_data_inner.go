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

// checks if the GetFinanceSavingsLendingRateHistoryV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceSavingsLendingRateHistoryV5RespDataInner{}

// GetFinanceSavingsLendingRateHistoryV5RespDataInner struct for GetFinanceSavingsLendingRateHistoryV5RespDataInner
type GetFinanceSavingsLendingRateHistoryV5RespDataInner struct {
	// Lending amount
	Amt *string `json:"amt,omitempty"`
	// Currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Lending annual interest rate
	Rate *string `json:"rate,omitempty"`
	// Time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewGetFinanceSavingsLendingRateHistoryV5RespDataInner instantiates a new GetFinanceSavingsLendingRateHistoryV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceSavingsLendingRateHistoryV5RespDataInner() *GetFinanceSavingsLendingRateHistoryV5RespDataInner {
	this := GetFinanceSavingsLendingRateHistoryV5RespDataInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var rate string = ""
	this.Rate = &rate
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetFinanceSavingsLendingRateHistoryV5RespDataInnerWithDefaults instantiates a new GetFinanceSavingsLendingRateHistoryV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceSavingsLendingRateHistoryV5RespDataInnerWithDefaults() *GetFinanceSavingsLendingRateHistoryV5RespDataInner {
	this := GetFinanceSavingsLendingRateHistoryV5RespDataInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var rate string = ""
	this.Rate = &rate
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) SetAmt(v string) {
	o.Amt = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) SetRate(v string) {
	o.Rate = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetFinanceSavingsLendingRateHistoryV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetFinanceSavingsLendingRateHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceSavingsLendingRateHistoryV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetFinanceSavingsLendingRateHistoryV5RespDataInner struct {
	value *GetFinanceSavingsLendingRateHistoryV5RespDataInner
	isSet bool
}

func (v NullableGetFinanceSavingsLendingRateHistoryV5RespDataInner) Get() *GetFinanceSavingsLendingRateHistoryV5RespDataInner {
	return v.value
}

func (v *NullableGetFinanceSavingsLendingRateHistoryV5RespDataInner) Set(val *GetFinanceSavingsLendingRateHistoryV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceSavingsLendingRateHistoryV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceSavingsLendingRateHistoryV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceSavingsLendingRateHistoryV5RespDataInner(val *GetFinanceSavingsLendingRateHistoryV5RespDataInner) *NullableGetFinanceSavingsLendingRateHistoryV5RespDataInner {
	return &NullableGetFinanceSavingsLendingRateHistoryV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetFinanceSavingsLendingRateHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceSavingsLendingRateHistoryV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


