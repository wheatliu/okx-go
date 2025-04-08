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

// checks if the GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner{}

// GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner struct for GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner
type GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner struct {
	// Loan amount
	Amt *string `json:"amt,omitempty"`
	// Loan currency, e.g. `USDT`
	Ccy *string `json:"ccy,omitempty"`
}

// NewGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner instantiates a new GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner() *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner {
	this := GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	return &this
}

// NewGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInnerWithDefaults instantiates a new GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInnerWithDefaults() *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner {
	this := GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	return &this
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) SetAmt(v string) {
	o.Amt = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) SetCcy(v string) {
	o.Ccy = &v
}

func (o GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	return toSerialize, nil
}

type NullableGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner struct {
	value *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner
	isSet bool
}

func (v NullableGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) Get() *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner {
	return v.value
}

func (v *NullableGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) Set(val *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner(val *GetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) *NullableGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner {
	return &NullableGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner{value: val, isSet: true}
}

func (v NullableGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceFlexibleLoanLoanInfoV5RespDataInnerLoanDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


