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

// checks if the GetPublicInterestRateLoanQuotaV5RespDataVipInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPublicInterestRateLoanQuotaV5RespDataVipInner{}

// GetPublicInterestRateLoanQuotaV5RespDataVipInner struct for GetPublicInterestRateLoanQuotaV5RespDataVipInner
type GetPublicInterestRateLoanQuotaV5RespDataVipInner struct {
	// Interest rate discount(Deprecated)
	// Deprecated
	IrDiscount *string `json:"irDiscount,omitempty"`
	// VIP Level, e.g. `VIP1`
	Level *string `json:"level,omitempty"`
	// Loan quota coefficient. Loan quota = `quota` * `level`
	LoanQuotaCoef *string `json:"loanQuotaCoef,omitempty"`
}

// NewGetPublicInterestRateLoanQuotaV5RespDataVipInner instantiates a new GetPublicInterestRateLoanQuotaV5RespDataVipInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPublicInterestRateLoanQuotaV5RespDataVipInner() *GetPublicInterestRateLoanQuotaV5RespDataVipInner {
	this := GetPublicInterestRateLoanQuotaV5RespDataVipInner{}
	var irDiscount string = ""
	this.IrDiscount = &irDiscount
	var level string = ""
	this.Level = &level
	var loanQuotaCoef string = ""
	this.LoanQuotaCoef = &loanQuotaCoef
	return &this
}

// NewGetPublicInterestRateLoanQuotaV5RespDataVipInnerWithDefaults instantiates a new GetPublicInterestRateLoanQuotaV5RespDataVipInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPublicInterestRateLoanQuotaV5RespDataVipInnerWithDefaults() *GetPublicInterestRateLoanQuotaV5RespDataVipInner {
	this := GetPublicInterestRateLoanQuotaV5RespDataVipInner{}
	var irDiscount string = ""
	this.IrDiscount = &irDiscount
	var level string = ""
	this.Level = &level
	var loanQuotaCoef string = ""
	this.LoanQuotaCoef = &loanQuotaCoef
	return &this
}

// GetIrDiscount returns the IrDiscount field value if set, zero value otherwise.
// Deprecated
func (o *GetPublicInterestRateLoanQuotaV5RespDataVipInner) GetIrDiscount() string {
	if o == nil || IsNil(o.IrDiscount) {
		var ret string
		return ret
	}
	return *o.IrDiscount
}

// GetIrDiscountOk returns a tuple with the IrDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GetPublicInterestRateLoanQuotaV5RespDataVipInner) GetIrDiscountOk() (*string, bool) {
	if o == nil || IsNil(o.IrDiscount) {
		return nil, false
	}
	return o.IrDiscount, true
}

// HasIrDiscount returns a boolean if a field has been set.
func (o *GetPublicInterestRateLoanQuotaV5RespDataVipInner) HasIrDiscount() bool {
	if o != nil && !IsNil(o.IrDiscount) {
		return true
	}

	return false
}

// SetIrDiscount gets a reference to the given string and assigns it to the IrDiscount field.
// Deprecated
func (o *GetPublicInterestRateLoanQuotaV5RespDataVipInner) SetIrDiscount(v string) {
	o.IrDiscount = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *GetPublicInterestRateLoanQuotaV5RespDataVipInner) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicInterestRateLoanQuotaV5RespDataVipInner) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *GetPublicInterestRateLoanQuotaV5RespDataVipInner) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *GetPublicInterestRateLoanQuotaV5RespDataVipInner) SetLevel(v string) {
	o.Level = &v
}

// GetLoanQuotaCoef returns the LoanQuotaCoef field value if set, zero value otherwise.
func (o *GetPublicInterestRateLoanQuotaV5RespDataVipInner) GetLoanQuotaCoef() string {
	if o == nil || IsNil(o.LoanQuotaCoef) {
		var ret string
		return ret
	}
	return *o.LoanQuotaCoef
}

// GetLoanQuotaCoefOk returns a tuple with the LoanQuotaCoef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicInterestRateLoanQuotaV5RespDataVipInner) GetLoanQuotaCoefOk() (*string, bool) {
	if o == nil || IsNil(o.LoanQuotaCoef) {
		return nil, false
	}
	return o.LoanQuotaCoef, true
}

// HasLoanQuotaCoef returns a boolean if a field has been set.
func (o *GetPublicInterestRateLoanQuotaV5RespDataVipInner) HasLoanQuotaCoef() bool {
	if o != nil && !IsNil(o.LoanQuotaCoef) {
		return true
	}

	return false
}

// SetLoanQuotaCoef gets a reference to the given string and assigns it to the LoanQuotaCoef field.
func (o *GetPublicInterestRateLoanQuotaV5RespDataVipInner) SetLoanQuotaCoef(v string) {
	o.LoanQuotaCoef = &v
}

func (o GetPublicInterestRateLoanQuotaV5RespDataVipInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPublicInterestRateLoanQuotaV5RespDataVipInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IrDiscount) {
		toSerialize["irDiscount"] = o.IrDiscount
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.LoanQuotaCoef) {
		toSerialize["loanQuotaCoef"] = o.LoanQuotaCoef
	}
	return toSerialize, nil
}

type NullableGetPublicInterestRateLoanQuotaV5RespDataVipInner struct {
	value *GetPublicInterestRateLoanQuotaV5RespDataVipInner
	isSet bool
}

func (v NullableGetPublicInterestRateLoanQuotaV5RespDataVipInner) Get() *GetPublicInterestRateLoanQuotaV5RespDataVipInner {
	return v.value
}

func (v *NullableGetPublicInterestRateLoanQuotaV5RespDataVipInner) Set(val *GetPublicInterestRateLoanQuotaV5RespDataVipInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPublicInterestRateLoanQuotaV5RespDataVipInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPublicInterestRateLoanQuotaV5RespDataVipInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPublicInterestRateLoanQuotaV5RespDataVipInner(val *GetPublicInterestRateLoanQuotaV5RespDataVipInner) *NullableGetPublicInterestRateLoanQuotaV5RespDataVipInner {
	return &NullableGetPublicInterestRateLoanQuotaV5RespDataVipInner{value: val, isSet: true}
}

func (v NullableGetPublicInterestRateLoanQuotaV5RespDataVipInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPublicInterestRateLoanQuotaV5RespDataVipInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


