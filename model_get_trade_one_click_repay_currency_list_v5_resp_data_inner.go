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

// checks if the GetTradeOneClickRepayCurrencyListV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTradeOneClickRepayCurrencyListV5RespDataInner{}

// GetTradeOneClickRepayCurrencyListV5RespDataInner struct for GetTradeOneClickRepayCurrencyListV5RespDataInner
type GetTradeOneClickRepayCurrencyListV5RespDataInner struct {
	// Debt currency data list
	DebtData []GetTradeOneClickRepayCurrencyListV5RespDataInnerDebtDataInner `json:"debtData,omitempty"`
	// Debt type   `cross`: cross   `isolated`: isolated
	DebtType *string `json:"debtType,omitempty"`
	// Repay currency data list
	RepayData []GetTradeOneClickRepayCurrencyListV2V5RespDataInnerRepayDataInner `json:"repayData,omitempty"`
}

// NewGetTradeOneClickRepayCurrencyListV5RespDataInner instantiates a new GetTradeOneClickRepayCurrencyListV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTradeOneClickRepayCurrencyListV5RespDataInner() *GetTradeOneClickRepayCurrencyListV5RespDataInner {
	this := GetTradeOneClickRepayCurrencyListV5RespDataInner{}
	var debtType string = ""
	this.DebtType = &debtType
	return &this
}

// NewGetTradeOneClickRepayCurrencyListV5RespDataInnerWithDefaults instantiates a new GetTradeOneClickRepayCurrencyListV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTradeOneClickRepayCurrencyListV5RespDataInnerWithDefaults() *GetTradeOneClickRepayCurrencyListV5RespDataInner {
	this := GetTradeOneClickRepayCurrencyListV5RespDataInner{}
	var debtType string = ""
	this.DebtType = &debtType
	return &this
}

// GetDebtData returns the DebtData field value if set, zero value otherwise.
func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) GetDebtData() []GetTradeOneClickRepayCurrencyListV5RespDataInnerDebtDataInner {
	if o == nil || IsNil(o.DebtData) {
		var ret []GetTradeOneClickRepayCurrencyListV5RespDataInnerDebtDataInner
		return ret
	}
	return o.DebtData
}

// GetDebtDataOk returns a tuple with the DebtData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) GetDebtDataOk() ([]GetTradeOneClickRepayCurrencyListV5RespDataInnerDebtDataInner, bool) {
	if o == nil || IsNil(o.DebtData) {
		return nil, false
	}
	return o.DebtData, true
}

// HasDebtData returns a boolean if a field has been set.
func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) HasDebtData() bool {
	if o != nil && !IsNil(o.DebtData) {
		return true
	}

	return false
}

// SetDebtData gets a reference to the given []GetTradeOneClickRepayCurrencyListV5RespDataInnerDebtDataInner and assigns it to the DebtData field.
func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) SetDebtData(v []GetTradeOneClickRepayCurrencyListV5RespDataInnerDebtDataInner) {
	o.DebtData = v
}

// GetDebtType returns the DebtType field value if set, zero value otherwise.
func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) GetDebtType() string {
	if o == nil || IsNil(o.DebtType) {
		var ret string
		return ret
	}
	return *o.DebtType
}

// GetDebtTypeOk returns a tuple with the DebtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) GetDebtTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DebtType) {
		return nil, false
	}
	return o.DebtType, true
}

// HasDebtType returns a boolean if a field has been set.
func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) HasDebtType() bool {
	if o != nil && !IsNil(o.DebtType) {
		return true
	}

	return false
}

// SetDebtType gets a reference to the given string and assigns it to the DebtType field.
func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) SetDebtType(v string) {
	o.DebtType = &v
}

// GetRepayData returns the RepayData field value if set, zero value otherwise.
func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) GetRepayData() []GetTradeOneClickRepayCurrencyListV2V5RespDataInnerRepayDataInner {
	if o == nil || IsNil(o.RepayData) {
		var ret []GetTradeOneClickRepayCurrencyListV2V5RespDataInnerRepayDataInner
		return ret
	}
	return o.RepayData
}

// GetRepayDataOk returns a tuple with the RepayData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) GetRepayDataOk() ([]GetTradeOneClickRepayCurrencyListV2V5RespDataInnerRepayDataInner, bool) {
	if o == nil || IsNil(o.RepayData) {
		return nil, false
	}
	return o.RepayData, true
}

// HasRepayData returns a boolean if a field has been set.
func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) HasRepayData() bool {
	if o != nil && !IsNil(o.RepayData) {
		return true
	}

	return false
}

// SetRepayData gets a reference to the given []GetTradeOneClickRepayCurrencyListV2V5RespDataInnerRepayDataInner and assigns it to the RepayData field.
func (o *GetTradeOneClickRepayCurrencyListV5RespDataInner) SetRepayData(v []GetTradeOneClickRepayCurrencyListV2V5RespDataInnerRepayDataInner) {
	o.RepayData = v
}

func (o GetTradeOneClickRepayCurrencyListV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTradeOneClickRepayCurrencyListV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DebtData) {
		toSerialize["debtData"] = o.DebtData
	}
	if !IsNil(o.DebtType) {
		toSerialize["debtType"] = o.DebtType
	}
	if !IsNil(o.RepayData) {
		toSerialize["repayData"] = o.RepayData
	}
	return toSerialize, nil
}

type NullableGetTradeOneClickRepayCurrencyListV5RespDataInner struct {
	value *GetTradeOneClickRepayCurrencyListV5RespDataInner
	isSet bool
}

func (v NullableGetTradeOneClickRepayCurrencyListV5RespDataInner) Get() *GetTradeOneClickRepayCurrencyListV5RespDataInner {
	return v.value
}

func (v *NullableGetTradeOneClickRepayCurrencyListV5RespDataInner) Set(val *GetTradeOneClickRepayCurrencyListV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTradeOneClickRepayCurrencyListV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTradeOneClickRepayCurrencyListV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTradeOneClickRepayCurrencyListV5RespDataInner(val *GetTradeOneClickRepayCurrencyListV5RespDataInner) *NullableGetTradeOneClickRepayCurrencyListV5RespDataInner {
	return &NullableGetTradeOneClickRepayCurrencyListV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetTradeOneClickRepayCurrencyListV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTradeOneClickRepayCurrencyListV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


