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

// checks if the GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner{}

// GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner struct for GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner
type GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner struct {
	// Currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Redeeming amount
	RedeemingAmt *string `json:"redeemingAmt,omitempty"`
}

// NewGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner instantiates a new GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner() *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner {
	this := GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var redeemingAmt string = ""
	this.RedeemingAmt = &redeemingAmt
	return &this
}

// NewGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInnerWithDefaults instantiates a new GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInnerWithDefaults() *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner {
	this := GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var redeemingAmt string = ""
	this.RedeemingAmt = &redeemingAmt
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetRedeemingAmt returns the RedeemingAmt field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) GetRedeemingAmt() string {
	if o == nil || IsNil(o.RedeemingAmt) {
		var ret string
		return ret
	}
	return *o.RedeemingAmt
}

// GetRedeemingAmtOk returns a tuple with the RedeemingAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) GetRedeemingAmtOk() (*string, bool) {
	if o == nil || IsNil(o.RedeemingAmt) {
		return nil, false
	}
	return o.RedeemingAmt, true
}

// HasRedeemingAmt returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) HasRedeemingAmt() bool {
	if o != nil && !IsNil(o.RedeemingAmt) {
		return true
	}

	return false
}

// SetRedeemingAmt gets a reference to the given string and assigns it to the RedeemingAmt field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) SetRedeemingAmt(v string) {
	o.RedeemingAmt = &v
}

func (o GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.RedeemingAmt) {
		toSerialize["redeemingAmt"] = o.RedeemingAmt
	}
	return toSerialize, nil
}

type NullableGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner struct {
	value *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner
	isSet bool
}

func (v NullableGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) Get() *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner {
	return v.value
}

func (v *NullableGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) Set(val *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner(val *GetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) *NullableGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner {
	return &NullableGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner{value: val, isSet: true}
}

func (v NullableGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceStakingDefiOrdersActiveV5RespDataInnerFastRedemptionDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


