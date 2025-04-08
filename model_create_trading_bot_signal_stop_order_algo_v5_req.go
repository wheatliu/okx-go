/*
Okx Rest API

OpenAPI specification for Okx cryptocurrency exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateTradingBotSignalStopOrderAlgoV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotSignalStopOrderAlgoV5Req{}

// CreateTradingBotSignalStopOrderAlgoV5Req struct for CreateTradingBotSignalStopOrderAlgoV5Req
type CreateTradingBotSignalStopOrderAlgoV5Req struct {
	// Algo ID
	AlgoId string `json:"algoId"`
}

type _CreateTradingBotSignalStopOrderAlgoV5Req CreateTradingBotSignalStopOrderAlgoV5Req

// NewCreateTradingBotSignalStopOrderAlgoV5Req instantiates a new CreateTradingBotSignalStopOrderAlgoV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotSignalStopOrderAlgoV5Req(algoId string) *CreateTradingBotSignalStopOrderAlgoV5Req {
	this := CreateTradingBotSignalStopOrderAlgoV5Req{}
	this.AlgoId = algoId
	return &this
}

// NewCreateTradingBotSignalStopOrderAlgoV5ReqWithDefaults instantiates a new CreateTradingBotSignalStopOrderAlgoV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotSignalStopOrderAlgoV5ReqWithDefaults() *CreateTradingBotSignalStopOrderAlgoV5Req {
	this := CreateTradingBotSignalStopOrderAlgoV5Req{}
	var algoId string = ""
	this.AlgoId = algoId
	return &this
}

// GetAlgoId returns the AlgoId field value
func (o *CreateTradingBotSignalStopOrderAlgoV5Req) GetAlgoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalStopOrderAlgoV5Req) GetAlgoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlgoId, true
}

// SetAlgoId sets field value
func (o *CreateTradingBotSignalStopOrderAlgoV5Req) SetAlgoId(v string) {
	o.AlgoId = v
}

func (o CreateTradingBotSignalStopOrderAlgoV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotSignalStopOrderAlgoV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["algoId"] = o.AlgoId
	return toSerialize, nil
}

func (o *CreateTradingBotSignalStopOrderAlgoV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"algoId",
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

	varCreateTradingBotSignalStopOrderAlgoV5Req := _CreateTradingBotSignalStopOrderAlgoV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTradingBotSignalStopOrderAlgoV5Req)

	if err != nil {
		return err
	}

	*o = CreateTradingBotSignalStopOrderAlgoV5Req(varCreateTradingBotSignalStopOrderAlgoV5Req)

	return err
}

type NullableCreateTradingBotSignalStopOrderAlgoV5Req struct {
	value *CreateTradingBotSignalStopOrderAlgoV5Req
	isSet bool
}

func (v NullableCreateTradingBotSignalStopOrderAlgoV5Req) Get() *CreateTradingBotSignalStopOrderAlgoV5Req {
	return v.value
}

func (v *NullableCreateTradingBotSignalStopOrderAlgoV5Req) Set(val *CreateTradingBotSignalStopOrderAlgoV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotSignalStopOrderAlgoV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotSignalStopOrderAlgoV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotSignalStopOrderAlgoV5Req(val *CreateTradingBotSignalStopOrderAlgoV5Req) *NullableCreateTradingBotSignalStopOrderAlgoV5Req {
	return &NullableCreateTradingBotSignalStopOrderAlgoV5Req{value: val, isSet: true}
}

func (v NullableCreateTradingBotSignalStopOrderAlgoV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotSignalStopOrderAlgoV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


