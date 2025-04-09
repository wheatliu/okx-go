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

// checks if the CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner{}

// CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner struct for CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner
type CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner struct {
	// Amount
	Amt *string `json:"amt,omitempty"`
	// Currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
}

// NewCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner instantiates a new CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner() *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner {
	this := CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	return &this
}

// NewCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInnerWithDefaults instantiates a new CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInnerWithDefaults() *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner {
	this := CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	return &this
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) SetAmt(v string) {
	o.Amt = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) SetCcy(v string) {
	o.Ccy = &v
}

func (o CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	return toSerialize, nil
}

type NullableCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner struct {
	value *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner
	isSet bool
}

func (v NullableCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) Get() *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner {
	return v.value
}

func (v *NullableCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) Set(val *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner(val *CreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) *NullableCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner {
	return &NullableCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner{value: val, isSet: true}
}

func (v NullableCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFinanceFlexibleLoanMaxLoanV5ReqSupCollateralInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


