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

// checks if the GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner{}

// GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner struct for GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner
type GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner struct {
	// Available amount
	Amt *string `json:"amt,omitempty"`
	// Currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Notional value in `USD`
	NotionalUsd *string `json:"notionalUsd,omitempty"`
}

// NewGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner instantiates a new GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner() *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner {
	this := GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var notionalUsd string = ""
	this.NotionalUsd = &notionalUsd
	return &this
}

// NewGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInnerWithDefaults instantiates a new GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInnerWithDefaults() *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner {
	this := GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var notionalUsd string = ""
	this.NotionalUsd = &notionalUsd
	return &this
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) SetAmt(v string) {
	o.Amt = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetNotionalUsd returns the NotionalUsd field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) GetNotionalUsd() string {
	if o == nil || IsNil(o.NotionalUsd) {
		var ret string
		return ret
	}
	return *o.NotionalUsd
}

// GetNotionalUsdOk returns a tuple with the NotionalUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) GetNotionalUsdOk() (*string, bool) {
	if o == nil || IsNil(o.NotionalUsd) {
		return nil, false
	}
	return o.NotionalUsd, true
}

// HasNotionalUsd returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) HasNotionalUsd() bool {
	if o != nil && !IsNil(o.NotionalUsd) {
		return true
	}

	return false
}

// SetNotionalUsd gets a reference to the given string and assigns it to the NotionalUsd field.
func (o *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) SetNotionalUsd(v string) {
	o.NotionalUsd = &v
}

func (o GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.NotionalUsd) {
		toSerialize["notionalUsd"] = o.NotionalUsd
	}
	return toSerialize, nil
}

type NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner struct {
	value *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner
	isSet bool
}

func (v NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) Get() *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner {
	return v.value
}

func (v *NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) Set(val *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner(val *GetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) *NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner {
	return &NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner{value: val, isSet: true}
}

func (v NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceFlexibleLoanCollateralAssetsV5RespDataInnerAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


