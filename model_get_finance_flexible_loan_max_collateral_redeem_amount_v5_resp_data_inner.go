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

// checks if the GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner{}

// GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner struct for GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner
type GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner struct {
	// Collateral currency, e.g. `USDT`
	Ccy *string `json:"ccy,omitempty"`
	// Maximum collateral redeem amount
	MaxRedeemAmt *string `json:"maxRedeemAmt,omitempty"`
}

// NewGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner instantiates a new GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner() *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner {
	this := GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var maxRedeemAmt string = ""
	this.MaxRedeemAmt = &maxRedeemAmt
	return &this
}

// NewGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInnerWithDefaults instantiates a new GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInnerWithDefaults() *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner {
	this := GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var maxRedeemAmt string = ""
	this.MaxRedeemAmt = &maxRedeemAmt
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetMaxRedeemAmt returns the MaxRedeemAmt field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) GetMaxRedeemAmt() string {
	if o == nil || IsNil(o.MaxRedeemAmt) {
		var ret string
		return ret
	}
	return *o.MaxRedeemAmt
}

// GetMaxRedeemAmtOk returns a tuple with the MaxRedeemAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) GetMaxRedeemAmtOk() (*string, bool) {
	if o == nil || IsNil(o.MaxRedeemAmt) {
		return nil, false
	}
	return o.MaxRedeemAmt, true
}

// HasMaxRedeemAmt returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) HasMaxRedeemAmt() bool {
	if o != nil && !IsNil(o.MaxRedeemAmt) {
		return true
	}

	return false
}

// SetMaxRedeemAmt gets a reference to the given string and assigns it to the MaxRedeemAmt field.
func (o *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) SetMaxRedeemAmt(v string) {
	o.MaxRedeemAmt = &v
}

func (o GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.MaxRedeemAmt) {
		toSerialize["maxRedeemAmt"] = o.MaxRedeemAmt
	}
	return toSerialize, nil
}

type NullableGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner struct {
	value *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner
	isSet bool
}

func (v NullableGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) Get() *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner {
	return v.value
}

func (v *NullableGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) Set(val *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner(val *GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) *NullableGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner {
	return &NullableGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceFlexibleLoanMaxCollateralRedeemAmountV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


