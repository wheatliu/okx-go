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

// checks if the CreateTradeOneClickRepayV2V5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeOneClickRepayV2V5RespDataInner{}

// CreateTradeOneClickRepayV2V5RespDataInner struct for CreateTradeOneClickRepayV2V5RespDataInner
type CreateTradeOneClickRepayV2V5RespDataInner struct {
	// Debt currency
	DebtCcy *string `json:"debtCcy,omitempty"`
	// Repay currency list, e.g. [\"USDC\",\"BTC\"]
	RepayCcyList []string `json:"repayCcyList,omitempty"`
	// Request time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewCreateTradeOneClickRepayV2V5RespDataInner instantiates a new CreateTradeOneClickRepayV2V5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeOneClickRepayV2V5RespDataInner() *CreateTradeOneClickRepayV2V5RespDataInner {
	this := CreateTradeOneClickRepayV2V5RespDataInner{}
	var debtCcy string = ""
	this.DebtCcy = &debtCcy
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewCreateTradeOneClickRepayV2V5RespDataInnerWithDefaults instantiates a new CreateTradeOneClickRepayV2V5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeOneClickRepayV2V5RespDataInnerWithDefaults() *CreateTradeOneClickRepayV2V5RespDataInner {
	this := CreateTradeOneClickRepayV2V5RespDataInner{}
	var debtCcy string = ""
	this.DebtCcy = &debtCcy
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetDebtCcy returns the DebtCcy field value if set, zero value otherwise.
func (o *CreateTradeOneClickRepayV2V5RespDataInner) GetDebtCcy() string {
	if o == nil || IsNil(o.DebtCcy) {
		var ret string
		return ret
	}
	return *o.DebtCcy
}

// GetDebtCcyOk returns a tuple with the DebtCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeOneClickRepayV2V5RespDataInner) GetDebtCcyOk() (*string, bool) {
	if o == nil || IsNil(o.DebtCcy) {
		return nil, false
	}
	return o.DebtCcy, true
}

// HasDebtCcy returns a boolean if a field has been set.
func (o *CreateTradeOneClickRepayV2V5RespDataInner) HasDebtCcy() bool {
	if o != nil && !IsNil(o.DebtCcy) {
		return true
	}

	return false
}

// SetDebtCcy gets a reference to the given string and assigns it to the DebtCcy field.
func (o *CreateTradeOneClickRepayV2V5RespDataInner) SetDebtCcy(v string) {
	o.DebtCcy = &v
}

// GetRepayCcyList returns the RepayCcyList field value if set, zero value otherwise.
func (o *CreateTradeOneClickRepayV2V5RespDataInner) GetRepayCcyList() []string {
	if o == nil || IsNil(o.RepayCcyList) {
		var ret []string
		return ret
	}
	return o.RepayCcyList
}

// GetRepayCcyListOk returns a tuple with the RepayCcyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeOneClickRepayV2V5RespDataInner) GetRepayCcyListOk() ([]string, bool) {
	if o == nil || IsNil(o.RepayCcyList) {
		return nil, false
	}
	return o.RepayCcyList, true
}

// HasRepayCcyList returns a boolean if a field has been set.
func (o *CreateTradeOneClickRepayV2V5RespDataInner) HasRepayCcyList() bool {
	if o != nil && !IsNil(o.RepayCcyList) {
		return true
	}

	return false
}

// SetRepayCcyList gets a reference to the given []string and assigns it to the RepayCcyList field.
func (o *CreateTradeOneClickRepayV2V5RespDataInner) SetRepayCcyList(v []string) {
	o.RepayCcyList = v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *CreateTradeOneClickRepayV2V5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeOneClickRepayV2V5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *CreateTradeOneClickRepayV2V5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *CreateTradeOneClickRepayV2V5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o CreateTradeOneClickRepayV2V5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeOneClickRepayV2V5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DebtCcy) {
		toSerialize["debtCcy"] = o.DebtCcy
	}
	if !IsNil(o.RepayCcyList) {
		toSerialize["repayCcyList"] = o.RepayCcyList
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableCreateTradeOneClickRepayV2V5RespDataInner struct {
	value *CreateTradeOneClickRepayV2V5RespDataInner
	isSet bool
}

func (v NullableCreateTradeOneClickRepayV2V5RespDataInner) Get() *CreateTradeOneClickRepayV2V5RespDataInner {
	return v.value
}

func (v *NullableCreateTradeOneClickRepayV2V5RespDataInner) Set(val *CreateTradeOneClickRepayV2V5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeOneClickRepayV2V5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeOneClickRepayV2V5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeOneClickRepayV2V5RespDataInner(val *CreateTradeOneClickRepayV2V5RespDataInner) *NullableCreateTradeOneClickRepayV2V5RespDataInner {
	return &NullableCreateTradeOneClickRepayV2V5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateTradeOneClickRepayV2V5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeOneClickRepayV2V5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


