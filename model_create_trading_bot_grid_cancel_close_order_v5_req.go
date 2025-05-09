/*
OKX v5 API

OpenAPI specification for Okx exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateTradingBotGridCancelCloseOrderV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotGridCancelCloseOrderV5Req{}

// CreateTradingBotGridCancelCloseOrderV5Req struct for CreateTradingBotGridCancelCloseOrderV5Req
type CreateTradingBotGridCancelCloseOrderV5Req struct {
	// Algo ID
	AlgoId string `json:"algoId"`
	// Close position order ID
	OrdId string `json:"ordId"`
}

type _CreateTradingBotGridCancelCloseOrderV5Req CreateTradingBotGridCancelCloseOrderV5Req

// NewCreateTradingBotGridCancelCloseOrderV5Req instantiates a new CreateTradingBotGridCancelCloseOrderV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotGridCancelCloseOrderV5Req(algoId string, ordId string) *CreateTradingBotGridCancelCloseOrderV5Req {
	this := CreateTradingBotGridCancelCloseOrderV5Req{}
	this.AlgoId = algoId
	this.OrdId = ordId
	return &this
}

// NewCreateTradingBotGridCancelCloseOrderV5ReqWithDefaults instantiates a new CreateTradingBotGridCancelCloseOrderV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotGridCancelCloseOrderV5ReqWithDefaults() *CreateTradingBotGridCancelCloseOrderV5Req {
	this := CreateTradingBotGridCancelCloseOrderV5Req{}
	var algoId string = ""
	this.AlgoId = algoId
	var ordId string = ""
	this.OrdId = ordId
	return &this
}

// GetAlgoId returns the AlgoId field value
func (o *CreateTradingBotGridCancelCloseOrderV5Req) GetAlgoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridCancelCloseOrderV5Req) GetAlgoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlgoId, true
}

// SetAlgoId sets field value
func (o *CreateTradingBotGridCancelCloseOrderV5Req) SetAlgoId(v string) {
	o.AlgoId = v
}

// GetOrdId returns the OrdId field value
func (o *CreateTradingBotGridCancelCloseOrderV5Req) GetOrdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrdId
}

// GetOrdIdOk returns a tuple with the OrdId field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridCancelCloseOrderV5Req) GetOrdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrdId, true
}

// SetOrdId sets field value
func (o *CreateTradingBotGridCancelCloseOrderV5Req) SetOrdId(v string) {
	o.OrdId = v
}

func (o CreateTradingBotGridCancelCloseOrderV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotGridCancelCloseOrderV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["algoId"] = o.AlgoId
	toSerialize["ordId"] = o.OrdId
	return toSerialize, nil
}

func (o *CreateTradingBotGridCancelCloseOrderV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"algoId",
		"ordId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateTradingBotGridCancelCloseOrderV5Req := _CreateTradingBotGridCancelCloseOrderV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTradingBotGridCancelCloseOrderV5Req)

	if err != nil {
		return err
	}

	*o = CreateTradingBotGridCancelCloseOrderV5Req(varCreateTradingBotGridCancelCloseOrderV5Req)

	return err
}

type NullableCreateTradingBotGridCancelCloseOrderV5Req struct {
	value *CreateTradingBotGridCancelCloseOrderV5Req
	isSet bool
}

func (v NullableCreateTradingBotGridCancelCloseOrderV5Req) Get() *CreateTradingBotGridCancelCloseOrderV5Req {
	return v.value
}

func (v *NullableCreateTradingBotGridCancelCloseOrderV5Req) Set(val *CreateTradingBotGridCancelCloseOrderV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotGridCancelCloseOrderV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotGridCancelCloseOrderV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotGridCancelCloseOrderV5Req(val *CreateTradingBotGridCancelCloseOrderV5Req) *NullableCreateTradingBotGridCancelCloseOrderV5Req {
	return &NullableCreateTradingBotGridCancelCloseOrderV5Req{value: val, isSet: true}
}

func (v NullableCreateTradingBotGridCancelCloseOrderV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotGridCancelCloseOrderV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


