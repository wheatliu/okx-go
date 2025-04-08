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

// checks if the CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner{}

// CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner struct for CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner
type CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner struct {
	// Invest amount
	Amt *string `json:"amt,omitempty"`
	// Invest currency
	Ccy *string `json:"ccy,omitempty"`
}

// NewCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner instantiates a new CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner() *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner {
	this := CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	return &this
}

// NewCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInnerWithDefaults instantiates a new CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInnerWithDefaults() *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner {
	this := CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	return &this
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) SetAmt(v string) {
	o.Amt = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) SetCcy(v string) {
	o.Ccy = &v
}

func (o CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	return toSerialize, nil
}

type NullableCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner struct {
	value *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner
	isSet bool
}

func (v NullableCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) Get() *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner {
	return v.value
}

func (v *NullableCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) Set(val *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner(val *CreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) *NullableCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner {
	return &NullableCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner{value: val, isSet: true}
}

func (v NullableCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotGridMinInvestmentV5ReqInvestmentDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


