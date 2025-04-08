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

// checks if the GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails{}

// GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails The details of available amount across all sub-accounts  The value of `surplusLmt` is the minimum value within this array. It can help you judge the reason that `surplusLmt` is not enough.  Only applicable to `VIP loans`
type GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails struct {
	// Total remaining quota for master account and sub-accounts
	AllAcctRemainingQuota *string `json:"allAcctRemainingQuota,omitempty"`
	// The remaining quota for the current account.  Only applicable to the case in which the sub-account is assigned the loan allocation
	CurAcctRemainingQuota *string `json:"curAcctRemainingQuota,omitempty"`
	// Remaining quota for the platform.  The format like  \"600\" will be returned when it is more than `curAcctRemainingQuota` or `allAcctRemainingQuota`
	PlatRemainingQuota *string `json:"platRemainingQuota,omitempty"`
}

// NewGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails instantiates a new GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails() *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails {
	this := GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails{}
	var allAcctRemainingQuota string = ""
	this.AllAcctRemainingQuota = &allAcctRemainingQuota
	var curAcctRemainingQuota string = ""
	this.CurAcctRemainingQuota = &curAcctRemainingQuota
	var platRemainingQuota string = ""
	this.PlatRemainingQuota = &platRemainingQuota
	return &this
}

// NewGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetailsWithDefaults instantiates a new GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetailsWithDefaults() *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails {
	this := GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails{}
	var allAcctRemainingQuota string = ""
	this.AllAcctRemainingQuota = &allAcctRemainingQuota
	var curAcctRemainingQuota string = ""
	this.CurAcctRemainingQuota = &curAcctRemainingQuota
	var platRemainingQuota string = ""
	this.PlatRemainingQuota = &platRemainingQuota
	return &this
}

// GetAllAcctRemainingQuota returns the AllAcctRemainingQuota field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) GetAllAcctRemainingQuota() string {
	if o == nil || IsNil(o.AllAcctRemainingQuota) {
		var ret string
		return ret
	}
	return *o.AllAcctRemainingQuota
}

// GetAllAcctRemainingQuotaOk returns a tuple with the AllAcctRemainingQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) GetAllAcctRemainingQuotaOk() (*string, bool) {
	if o == nil || IsNil(o.AllAcctRemainingQuota) {
		return nil, false
	}
	return o.AllAcctRemainingQuota, true
}

// HasAllAcctRemainingQuota returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) HasAllAcctRemainingQuota() bool {
	if o != nil && !IsNil(o.AllAcctRemainingQuota) {
		return true
	}

	return false
}

// SetAllAcctRemainingQuota gets a reference to the given string and assigns it to the AllAcctRemainingQuota field.
func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) SetAllAcctRemainingQuota(v string) {
	o.AllAcctRemainingQuota = &v
}

// GetCurAcctRemainingQuota returns the CurAcctRemainingQuota field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) GetCurAcctRemainingQuota() string {
	if o == nil || IsNil(o.CurAcctRemainingQuota) {
		var ret string
		return ret
	}
	return *o.CurAcctRemainingQuota
}

// GetCurAcctRemainingQuotaOk returns a tuple with the CurAcctRemainingQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) GetCurAcctRemainingQuotaOk() (*string, bool) {
	if o == nil || IsNil(o.CurAcctRemainingQuota) {
		return nil, false
	}
	return o.CurAcctRemainingQuota, true
}

// HasCurAcctRemainingQuota returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) HasCurAcctRemainingQuota() bool {
	if o != nil && !IsNil(o.CurAcctRemainingQuota) {
		return true
	}

	return false
}

// SetCurAcctRemainingQuota gets a reference to the given string and assigns it to the CurAcctRemainingQuota field.
func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) SetCurAcctRemainingQuota(v string) {
	o.CurAcctRemainingQuota = &v
}

// GetPlatRemainingQuota returns the PlatRemainingQuota field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) GetPlatRemainingQuota() string {
	if o == nil || IsNil(o.PlatRemainingQuota) {
		var ret string
		return ret
	}
	return *o.PlatRemainingQuota
}

// GetPlatRemainingQuotaOk returns a tuple with the PlatRemainingQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) GetPlatRemainingQuotaOk() (*string, bool) {
	if o == nil || IsNil(o.PlatRemainingQuota) {
		return nil, false
	}
	return o.PlatRemainingQuota, true
}

// HasPlatRemainingQuota returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) HasPlatRemainingQuota() bool {
	if o != nil && !IsNil(o.PlatRemainingQuota) {
		return true
	}

	return false
}

// SetPlatRemainingQuota gets a reference to the given string and assigns it to the PlatRemainingQuota field.
func (o *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) SetPlatRemainingQuota(v string) {
	o.PlatRemainingQuota = &v
}

func (o GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllAcctRemainingQuota) {
		toSerialize["allAcctRemainingQuota"] = o.AllAcctRemainingQuota
	}
	if !IsNil(o.CurAcctRemainingQuota) {
		toSerialize["curAcctRemainingQuota"] = o.CurAcctRemainingQuota
	}
	if !IsNil(o.PlatRemainingQuota) {
		toSerialize["platRemainingQuota"] = o.PlatRemainingQuota
	}
	return toSerialize, nil
}

type NullableGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails struct {
	value *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails
	isSet bool
}

func (v NullableGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) Get() *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails {
	return v.value
}

func (v *NullableGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) Set(val *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails(val *GetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) *NullableGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails {
	return &NullableGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails{value: val, isSet: true}
}

func (v NullableGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountInterestLimitsV5RespDataRecordsInnerSurplusLmtDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


