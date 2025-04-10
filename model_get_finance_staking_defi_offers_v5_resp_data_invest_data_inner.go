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

// checks if the GetFinanceStakingDefiOffersV5RespDataInvestDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceStakingDefiOffersV5RespDataInvestDataInner{}

// GetFinanceStakingDefiOffersV5RespDataInvestDataInner struct for GetFinanceStakingDefiOffersV5RespDataInvestDataInner
type GetFinanceStakingDefiOffersV5RespDataInvestDataInner struct {
	// Available balance to invest
	Bal *string `json:"bal,omitempty"`
	// Investment currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Maximum available subscription amount
	MaxAmt *string `json:"maxAmt,omitempty"`
	// Minimum subscription amount
	MinAmt *string `json:"minAmt,omitempty"`
}

// NewGetFinanceStakingDefiOffersV5RespDataInvestDataInner instantiates a new GetFinanceStakingDefiOffersV5RespDataInvestDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceStakingDefiOffersV5RespDataInvestDataInner() *GetFinanceStakingDefiOffersV5RespDataInvestDataInner {
	this := GetFinanceStakingDefiOffersV5RespDataInvestDataInner{}
	var bal string = ""
	this.Bal = &bal
	var ccy string = ""
	this.Ccy = &ccy
	var maxAmt string = ""
	this.MaxAmt = &maxAmt
	var minAmt string = ""
	this.MinAmt = &minAmt
	return &this
}

// NewGetFinanceStakingDefiOffersV5RespDataInvestDataInnerWithDefaults instantiates a new GetFinanceStakingDefiOffersV5RespDataInvestDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceStakingDefiOffersV5RespDataInvestDataInnerWithDefaults() *GetFinanceStakingDefiOffersV5RespDataInvestDataInner {
	this := GetFinanceStakingDefiOffersV5RespDataInvestDataInner{}
	var bal string = ""
	this.Bal = &bal
	var ccy string = ""
	this.Ccy = &ccy
	var maxAmt string = ""
	this.MaxAmt = &maxAmt
	var minAmt string = ""
	this.MinAmt = &minAmt
	return &this
}

// GetBal returns the Bal field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) GetBal() string {
	if o == nil || IsNil(o.Bal) {
		var ret string
		return ret
	}
	return *o.Bal
}

// GetBalOk returns a tuple with the Bal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) GetBalOk() (*string, bool) {
	if o == nil || IsNil(o.Bal) {
		return nil, false
	}
	return o.Bal, true
}

// HasBal returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) HasBal() bool {
	if o != nil && !IsNil(o.Bal) {
		return true
	}

	return false
}

// SetBal gets a reference to the given string and assigns it to the Bal field.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) SetBal(v string) {
	o.Bal = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetMaxAmt returns the MaxAmt field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) GetMaxAmt() string {
	if o == nil || IsNil(o.MaxAmt) {
		var ret string
		return ret
	}
	return *o.MaxAmt
}

// GetMaxAmtOk returns a tuple with the MaxAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) GetMaxAmtOk() (*string, bool) {
	if o == nil || IsNil(o.MaxAmt) {
		return nil, false
	}
	return o.MaxAmt, true
}

// HasMaxAmt returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) HasMaxAmt() bool {
	if o != nil && !IsNil(o.MaxAmt) {
		return true
	}

	return false
}

// SetMaxAmt gets a reference to the given string and assigns it to the MaxAmt field.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) SetMaxAmt(v string) {
	o.MaxAmt = &v
}

// GetMinAmt returns the MinAmt field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) GetMinAmt() string {
	if o == nil || IsNil(o.MinAmt) {
		var ret string
		return ret
	}
	return *o.MinAmt
}

// GetMinAmtOk returns a tuple with the MinAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) GetMinAmtOk() (*string, bool) {
	if o == nil || IsNil(o.MinAmt) {
		return nil, false
	}
	return o.MinAmt, true
}

// HasMinAmt returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) HasMinAmt() bool {
	if o != nil && !IsNil(o.MinAmt) {
		return true
	}

	return false
}

// SetMinAmt gets a reference to the given string and assigns it to the MinAmt field.
func (o *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) SetMinAmt(v string) {
	o.MinAmt = &v
}

func (o GetFinanceStakingDefiOffersV5RespDataInvestDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceStakingDefiOffersV5RespDataInvestDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bal) {
		toSerialize["bal"] = o.Bal
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.MaxAmt) {
		toSerialize["maxAmt"] = o.MaxAmt
	}
	if !IsNil(o.MinAmt) {
		toSerialize["minAmt"] = o.MinAmt
	}
	return toSerialize, nil
}

type NullableGetFinanceStakingDefiOffersV5RespDataInvestDataInner struct {
	value *GetFinanceStakingDefiOffersV5RespDataInvestDataInner
	isSet bool
}

func (v NullableGetFinanceStakingDefiOffersV5RespDataInvestDataInner) Get() *GetFinanceStakingDefiOffersV5RespDataInvestDataInner {
	return v.value
}

func (v *NullableGetFinanceStakingDefiOffersV5RespDataInvestDataInner) Set(val *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceStakingDefiOffersV5RespDataInvestDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceStakingDefiOffersV5RespDataInvestDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceStakingDefiOffersV5RespDataInvestDataInner(val *GetFinanceStakingDefiOffersV5RespDataInvestDataInner) *NullableGetFinanceStakingDefiOffersV5RespDataInvestDataInner {
	return &NullableGetFinanceStakingDefiOffersV5RespDataInvestDataInner{value: val, isSet: true}
}

func (v NullableGetFinanceStakingDefiOffersV5RespDataInvestDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceStakingDefiOffersV5RespDataInvestDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


