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

// checks if the CreateTradingBotGridMarginBalanceV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotGridMarginBalanceV5RespDataInner{}

// CreateTradingBotGridMarginBalanceV5RespDataInner struct for CreateTradingBotGridMarginBalanceV5RespDataInner
type CreateTradingBotGridMarginBalanceV5RespDataInner struct {
	// Client-supplied Algo ID
	AlgoClOrdId *string `json:"algoClOrdId,omitempty"`
	// Algo ID
	AlgoId *string `json:"algoId,omitempty"`
}

// NewCreateTradingBotGridMarginBalanceV5RespDataInner instantiates a new CreateTradingBotGridMarginBalanceV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotGridMarginBalanceV5RespDataInner() *CreateTradingBotGridMarginBalanceV5RespDataInner {
	this := CreateTradingBotGridMarginBalanceV5RespDataInner{}
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	return &this
}

// NewCreateTradingBotGridMarginBalanceV5RespDataInnerWithDefaults instantiates a new CreateTradingBotGridMarginBalanceV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotGridMarginBalanceV5RespDataInnerWithDefaults() *CreateTradingBotGridMarginBalanceV5RespDataInner {
	this := CreateTradingBotGridMarginBalanceV5RespDataInner{}
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	return &this
}

// GetAlgoClOrdId returns the AlgoClOrdId field value if set, zero value otherwise.
func (o *CreateTradingBotGridMarginBalanceV5RespDataInner) GetAlgoClOrdId() string {
	if o == nil || IsNil(o.AlgoClOrdId) {
		var ret string
		return ret
	}
	return *o.AlgoClOrdId
}

// GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridMarginBalanceV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoClOrdId) {
		return nil, false
	}
	return o.AlgoClOrdId, true
}

// HasAlgoClOrdId returns a boolean if a field has been set.
func (o *CreateTradingBotGridMarginBalanceV5RespDataInner) HasAlgoClOrdId() bool {
	if o != nil && !IsNil(o.AlgoClOrdId) {
		return true
	}

	return false
}

// SetAlgoClOrdId gets a reference to the given string and assigns it to the AlgoClOrdId field.
func (o *CreateTradingBotGridMarginBalanceV5RespDataInner) SetAlgoClOrdId(v string) {
	o.AlgoClOrdId = &v
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *CreateTradingBotGridMarginBalanceV5RespDataInner) GetAlgoId() string {
	if o == nil || IsNil(o.AlgoId) {
		var ret string
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridMarginBalanceV5RespDataInner) GetAlgoIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *CreateTradingBotGridMarginBalanceV5RespDataInner) HasAlgoId() bool {
	if o != nil && !IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given string and assigns it to the AlgoId field.
func (o *CreateTradingBotGridMarginBalanceV5RespDataInner) SetAlgoId(v string) {
	o.AlgoId = &v
}

func (o CreateTradingBotGridMarginBalanceV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotGridMarginBalanceV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlgoClOrdId) {
		toSerialize["algoClOrdId"] = o.AlgoClOrdId
	}
	if !IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	return toSerialize, nil
}

type NullableCreateTradingBotGridMarginBalanceV5RespDataInner struct {
	value *CreateTradingBotGridMarginBalanceV5RespDataInner
	isSet bool
}

func (v NullableCreateTradingBotGridMarginBalanceV5RespDataInner) Get() *CreateTradingBotGridMarginBalanceV5RespDataInner {
	return v.value
}

func (v *NullableCreateTradingBotGridMarginBalanceV5RespDataInner) Set(val *CreateTradingBotGridMarginBalanceV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotGridMarginBalanceV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotGridMarginBalanceV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotGridMarginBalanceV5RespDataInner(val *CreateTradingBotGridMarginBalanceV5RespDataInner) *NullableCreateTradingBotGridMarginBalanceV5RespDataInner {
	return &NullableCreateTradingBotGridMarginBalanceV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateTradingBotGridMarginBalanceV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotGridMarginBalanceV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


