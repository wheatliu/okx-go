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

// checks if the GetFinanceSavingsLendingHistoryV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceSavingsLendingHistoryV5RespData{}

// GetFinanceSavingsLendingHistoryV5RespData struct for GetFinanceSavingsLendingHistoryV5RespData
type GetFinanceSavingsLendingHistoryV5RespData struct {
	// Lending amount
	Amt *string `json:"amt,omitempty"`
	// Currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Currency earnings
	Earnings *string `json:"earnings,omitempty"`
	// Lending annual interest rate
	Rate *string `json:"rate,omitempty"`
	// Lending time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewGetFinanceSavingsLendingHistoryV5RespData instantiates a new GetFinanceSavingsLendingHistoryV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceSavingsLendingHistoryV5RespData() *GetFinanceSavingsLendingHistoryV5RespData {
	this := GetFinanceSavingsLendingHistoryV5RespData{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var earnings string = ""
	this.Earnings = &earnings
	var rate string = ""
	this.Rate = &rate
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetFinanceSavingsLendingHistoryV5RespDataWithDefaults instantiates a new GetFinanceSavingsLendingHistoryV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceSavingsLendingHistoryV5RespDataWithDefaults() *GetFinanceSavingsLendingHistoryV5RespData {
	this := GetFinanceSavingsLendingHistoryV5RespData{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var earnings string = ""
	this.Earnings = &earnings
	var rate string = ""
	this.Rate = &rate
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *GetFinanceSavingsLendingHistoryV5RespData) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceSavingsLendingHistoryV5RespData) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *GetFinanceSavingsLendingHistoryV5RespData) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *GetFinanceSavingsLendingHistoryV5RespData) SetAmt(v string) {
	o.Amt = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFinanceSavingsLendingHistoryV5RespData) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceSavingsLendingHistoryV5RespData) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFinanceSavingsLendingHistoryV5RespData) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFinanceSavingsLendingHistoryV5RespData) SetCcy(v string) {
	o.Ccy = &v
}

// GetEarnings returns the Earnings field value if set, zero value otherwise.
func (o *GetFinanceSavingsLendingHistoryV5RespData) GetEarnings() string {
	if o == nil || IsNil(o.Earnings) {
		var ret string
		return ret
	}
	return *o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceSavingsLendingHistoryV5RespData) GetEarningsOk() (*string, bool) {
	if o == nil || IsNil(o.Earnings) {
		return nil, false
	}
	return o.Earnings, true
}

// HasEarnings returns a boolean if a field has been set.
func (o *GetFinanceSavingsLendingHistoryV5RespData) HasEarnings() bool {
	if o != nil && !IsNil(o.Earnings) {
		return true
	}

	return false
}

// SetEarnings gets a reference to the given string and assigns it to the Earnings field.
func (o *GetFinanceSavingsLendingHistoryV5RespData) SetEarnings(v string) {
	o.Earnings = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *GetFinanceSavingsLendingHistoryV5RespData) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceSavingsLendingHistoryV5RespData) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *GetFinanceSavingsLendingHistoryV5RespData) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *GetFinanceSavingsLendingHistoryV5RespData) SetRate(v string) {
	o.Rate = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetFinanceSavingsLendingHistoryV5RespData) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceSavingsLendingHistoryV5RespData) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetFinanceSavingsLendingHistoryV5RespData) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetFinanceSavingsLendingHistoryV5RespData) SetTs(v string) {
	o.Ts = &v
}

func (o GetFinanceSavingsLendingHistoryV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceSavingsLendingHistoryV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.Earnings) {
		toSerialize["earnings"] = o.Earnings
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetFinanceSavingsLendingHistoryV5RespData struct {
	value *GetFinanceSavingsLendingHistoryV5RespData
	isSet bool
}

func (v NullableGetFinanceSavingsLendingHistoryV5RespData) Get() *GetFinanceSavingsLendingHistoryV5RespData {
	return v.value
}

func (v *NullableGetFinanceSavingsLendingHistoryV5RespData) Set(val *GetFinanceSavingsLendingHistoryV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceSavingsLendingHistoryV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceSavingsLendingHistoryV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceSavingsLendingHistoryV5RespData(val *GetFinanceSavingsLendingHistoryV5RespData) *NullableGetFinanceSavingsLendingHistoryV5RespData {
	return &NullableGetFinanceSavingsLendingHistoryV5RespData{value: val, isSet: true}
}

func (v NullableGetFinanceSavingsLendingHistoryV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceSavingsLendingHistoryV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


