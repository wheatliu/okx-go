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

// checks if the CreateTradingBotGridWithdrawIncomeV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotGridWithdrawIncomeV5RespDataInner{}

// CreateTradingBotGridWithdrawIncomeV5RespDataInner struct for CreateTradingBotGridWithdrawIncomeV5RespDataInner
type CreateTradingBotGridWithdrawIncomeV5RespDataInner struct {
	// Client-supplied Algo ID
	AlgoClOrdId *string `json:"algoClOrdId,omitempty"`
	// Algo ID
	AlgoId *string `json:"algoId,omitempty"`
	// Withdraw profit
	Profit *string `json:"profit,omitempty"`
}

// NewCreateTradingBotGridWithdrawIncomeV5RespDataInner instantiates a new CreateTradingBotGridWithdrawIncomeV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotGridWithdrawIncomeV5RespDataInner() *CreateTradingBotGridWithdrawIncomeV5RespDataInner {
	this := CreateTradingBotGridWithdrawIncomeV5RespDataInner{}
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	var profit string = ""
	this.Profit = &profit
	return &this
}

// NewCreateTradingBotGridWithdrawIncomeV5RespDataInnerWithDefaults instantiates a new CreateTradingBotGridWithdrawIncomeV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotGridWithdrawIncomeV5RespDataInnerWithDefaults() *CreateTradingBotGridWithdrawIncomeV5RespDataInner {
	this := CreateTradingBotGridWithdrawIncomeV5RespDataInner{}
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	var profit string = ""
	this.Profit = &profit
	return &this
}

// GetAlgoClOrdId returns the AlgoClOrdId field value if set, zero value otherwise.
func (o *CreateTradingBotGridWithdrawIncomeV5RespDataInner) GetAlgoClOrdId() string {
	if o == nil || IsNil(o.AlgoClOrdId) {
		var ret string
		return ret
	}
	return *o.AlgoClOrdId
}

// GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridWithdrawIncomeV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoClOrdId) {
		return nil, false
	}
	return o.AlgoClOrdId, true
}

// HasAlgoClOrdId returns a boolean if a field has been set.
func (o *CreateTradingBotGridWithdrawIncomeV5RespDataInner) HasAlgoClOrdId() bool {
	if o != nil && !IsNil(o.AlgoClOrdId) {
		return true
	}

	return false
}

// SetAlgoClOrdId gets a reference to the given string and assigns it to the AlgoClOrdId field.
func (o *CreateTradingBotGridWithdrawIncomeV5RespDataInner) SetAlgoClOrdId(v string) {
	o.AlgoClOrdId = &v
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *CreateTradingBotGridWithdrawIncomeV5RespDataInner) GetAlgoId() string {
	if o == nil || IsNil(o.AlgoId) {
		var ret string
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridWithdrawIncomeV5RespDataInner) GetAlgoIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *CreateTradingBotGridWithdrawIncomeV5RespDataInner) HasAlgoId() bool {
	if o != nil && !IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given string and assigns it to the AlgoId field.
func (o *CreateTradingBotGridWithdrawIncomeV5RespDataInner) SetAlgoId(v string) {
	o.AlgoId = &v
}

// GetProfit returns the Profit field value if set, zero value otherwise.
func (o *CreateTradingBotGridWithdrawIncomeV5RespDataInner) GetProfit() string {
	if o == nil || IsNil(o.Profit) {
		var ret string
		return ret
	}
	return *o.Profit
}

// GetProfitOk returns a tuple with the Profit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridWithdrawIncomeV5RespDataInner) GetProfitOk() (*string, bool) {
	if o == nil || IsNil(o.Profit) {
		return nil, false
	}
	return o.Profit, true
}

// HasProfit returns a boolean if a field has been set.
func (o *CreateTradingBotGridWithdrawIncomeV5RespDataInner) HasProfit() bool {
	if o != nil && !IsNil(o.Profit) {
		return true
	}

	return false
}

// SetProfit gets a reference to the given string and assigns it to the Profit field.
func (o *CreateTradingBotGridWithdrawIncomeV5RespDataInner) SetProfit(v string) {
	o.Profit = &v
}

func (o CreateTradingBotGridWithdrawIncomeV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotGridWithdrawIncomeV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlgoClOrdId) {
		toSerialize["algoClOrdId"] = o.AlgoClOrdId
	}
	if !IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !IsNil(o.Profit) {
		toSerialize["profit"] = o.Profit
	}
	return toSerialize, nil
}

type NullableCreateTradingBotGridWithdrawIncomeV5RespDataInner struct {
	value *CreateTradingBotGridWithdrawIncomeV5RespDataInner
	isSet bool
}

func (v NullableCreateTradingBotGridWithdrawIncomeV5RespDataInner) Get() *CreateTradingBotGridWithdrawIncomeV5RespDataInner {
	return v.value
}

func (v *NullableCreateTradingBotGridWithdrawIncomeV5RespDataInner) Set(val *CreateTradingBotGridWithdrawIncomeV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotGridWithdrawIncomeV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotGridWithdrawIncomeV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotGridWithdrawIncomeV5RespDataInner(val *CreateTradingBotGridWithdrawIncomeV5RespDataInner) *NullableCreateTradingBotGridWithdrawIncomeV5RespDataInner {
	return &NullableCreateTradingBotGridWithdrawIncomeV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateTradingBotGridWithdrawIncomeV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotGridWithdrawIncomeV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


