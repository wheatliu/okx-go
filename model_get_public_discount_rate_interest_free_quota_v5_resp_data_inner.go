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

// checks if the GetPublicDiscountRateInterestFreeQuotaV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPublicDiscountRateInterestFreeQuotaV5RespDataInner{}

// GetPublicDiscountRateInterestFreeQuotaV5RespDataInner struct for GetPublicDiscountRateInterestFreeQuotaV5RespDataInner
type GetPublicDiscountRateInterestFreeQuotaV5RespDataInner struct {
	// Interest-free quota
	Amt *string `json:"amt,omitempty"`
	// Currency
	Ccy *string `json:"ccy,omitempty"`
	// New discount details.
	Details []GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner `json:"details,omitempty"`
	// Discount rate level.(Deprecated)
	// Deprecated
	DiscountLv *string `json:"discountLv,omitempty"`
	// Minimum discount rate when it exceeds the maximum amount of the last tier.
	MinDiscountRate *string `json:"minDiscountRate,omitempty"`
}

// NewGetPublicDiscountRateInterestFreeQuotaV5RespDataInner instantiates a new GetPublicDiscountRateInterestFreeQuotaV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPublicDiscountRateInterestFreeQuotaV5RespDataInner() *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner {
	this := GetPublicDiscountRateInterestFreeQuotaV5RespDataInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var discountLv string = ""
	this.DiscountLv = &discountLv
	var minDiscountRate string = ""
	this.MinDiscountRate = &minDiscountRate
	return &this
}

// NewGetPublicDiscountRateInterestFreeQuotaV5RespDataInnerWithDefaults instantiates a new GetPublicDiscountRateInterestFreeQuotaV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPublicDiscountRateInterestFreeQuotaV5RespDataInnerWithDefaults() *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner {
	this := GetPublicDiscountRateInterestFreeQuotaV5RespDataInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var discountLv string = ""
	this.DiscountLv = &discountLv
	var minDiscountRate string = ""
	this.MinDiscountRate = &minDiscountRate
	return &this
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) SetAmt(v string) {
	o.Amt = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) GetDetails() []GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner {
	if o == nil || IsNil(o.Details) {
		var ret []GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) GetDetailsOk() ([]GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner and assigns it to the Details field.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) SetDetails(v []GetPublicDiscountRateInterestFreeQuotaV5RespDataInnerDetailsInner) {
	o.Details = v
}

// GetDiscountLv returns the DiscountLv field value if set, zero value otherwise.
// Deprecated
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) GetDiscountLv() string {
	if o == nil || IsNil(o.DiscountLv) {
		var ret string
		return ret
	}
	return *o.DiscountLv
}

// GetDiscountLvOk returns a tuple with the DiscountLv field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) GetDiscountLvOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountLv) {
		return nil, false
	}
	return o.DiscountLv, true
}

// HasDiscountLv returns a boolean if a field has been set.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) HasDiscountLv() bool {
	if o != nil && !IsNil(o.DiscountLv) {
		return true
	}

	return false
}

// SetDiscountLv gets a reference to the given string and assigns it to the DiscountLv field.
// Deprecated
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) SetDiscountLv(v string) {
	o.DiscountLv = &v
}

// GetMinDiscountRate returns the MinDiscountRate field value if set, zero value otherwise.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) GetMinDiscountRate() string {
	if o == nil || IsNil(o.MinDiscountRate) {
		var ret string
		return ret
	}
	return *o.MinDiscountRate
}

// GetMinDiscountRateOk returns a tuple with the MinDiscountRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) GetMinDiscountRateOk() (*string, bool) {
	if o == nil || IsNil(o.MinDiscountRate) {
		return nil, false
	}
	return o.MinDiscountRate, true
}

// HasMinDiscountRate returns a boolean if a field has been set.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) HasMinDiscountRate() bool {
	if o != nil && !IsNil(o.MinDiscountRate) {
		return true
	}

	return false
}

// SetMinDiscountRate gets a reference to the given string and assigns it to the MinDiscountRate field.
func (o *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) SetMinDiscountRate(v string) {
	o.MinDiscountRate = &v
}

func (o GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.DiscountLv) {
		toSerialize["discountLv"] = o.DiscountLv
	}
	if !IsNil(o.MinDiscountRate) {
		toSerialize["minDiscountRate"] = o.MinDiscountRate
	}
	return toSerialize, nil
}

type NullableGetPublicDiscountRateInterestFreeQuotaV5RespDataInner struct {
	value *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner
	isSet bool
}

func (v NullableGetPublicDiscountRateInterestFreeQuotaV5RespDataInner) Get() *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner {
	return v.value
}

func (v *NullableGetPublicDiscountRateInterestFreeQuotaV5RespDataInner) Set(val *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPublicDiscountRateInterestFreeQuotaV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPublicDiscountRateInterestFreeQuotaV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPublicDiscountRateInterestFreeQuotaV5RespDataInner(val *GetPublicDiscountRateInterestFreeQuotaV5RespDataInner) *NullableGetPublicDiscountRateInterestFreeQuotaV5RespDataInner {
	return &NullableGetPublicDiscountRateInterestFreeQuotaV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetPublicDiscountRateInterestFreeQuotaV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPublicDiscountRateInterestFreeQuotaV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


