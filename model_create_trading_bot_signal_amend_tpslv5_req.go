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

// checks if the CreateTradingBotSignalAmendTPSLV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotSignalAmendTPSLV5Req{}

// CreateTradingBotSignalAmendTPSLV5Req struct for CreateTradingBotSignalAmendTPSLV5Req
type CreateTradingBotSignalAmendTPSLV5Req struct {
	// Algo ID
	AlgoId string `json:"algoId"`
	// Exit setting
	ExitSettingParam string `json:"exitSettingParam"`
}

type _CreateTradingBotSignalAmendTPSLV5Req CreateTradingBotSignalAmendTPSLV5Req

// NewCreateTradingBotSignalAmendTPSLV5Req instantiates a new CreateTradingBotSignalAmendTPSLV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotSignalAmendTPSLV5Req(algoId string, exitSettingParam string) *CreateTradingBotSignalAmendTPSLV5Req {
	this := CreateTradingBotSignalAmendTPSLV5Req{}
	this.AlgoId = algoId
	this.ExitSettingParam = exitSettingParam
	return &this
}

// NewCreateTradingBotSignalAmendTPSLV5ReqWithDefaults instantiates a new CreateTradingBotSignalAmendTPSLV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotSignalAmendTPSLV5ReqWithDefaults() *CreateTradingBotSignalAmendTPSLV5Req {
	this := CreateTradingBotSignalAmendTPSLV5Req{}
	var algoId string = ""
	this.AlgoId = algoId
	var exitSettingParam string = ""
	this.ExitSettingParam = exitSettingParam
	return &this
}

// GetAlgoId returns the AlgoId field value
func (o *CreateTradingBotSignalAmendTPSLV5Req) GetAlgoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalAmendTPSLV5Req) GetAlgoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlgoId, true
}

// SetAlgoId sets field value
func (o *CreateTradingBotSignalAmendTPSLV5Req) SetAlgoId(v string) {
	o.AlgoId = v
}

// GetExitSettingParam returns the ExitSettingParam field value
func (o *CreateTradingBotSignalAmendTPSLV5Req) GetExitSettingParam() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExitSettingParam
}

// GetExitSettingParamOk returns a tuple with the ExitSettingParam field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalAmendTPSLV5Req) GetExitSettingParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExitSettingParam, true
}

// SetExitSettingParam sets field value
func (o *CreateTradingBotSignalAmendTPSLV5Req) SetExitSettingParam(v string) {
	o.ExitSettingParam = v
}

func (o CreateTradingBotSignalAmendTPSLV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotSignalAmendTPSLV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["algoId"] = o.AlgoId
	toSerialize["exitSettingParam"] = o.ExitSettingParam
	return toSerialize, nil
}

func (o *CreateTradingBotSignalAmendTPSLV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"algoId",
		"exitSettingParam",
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

	varCreateTradingBotSignalAmendTPSLV5Req := _CreateTradingBotSignalAmendTPSLV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTradingBotSignalAmendTPSLV5Req)

	if err != nil {
		return err
	}

	*o = CreateTradingBotSignalAmendTPSLV5Req(varCreateTradingBotSignalAmendTPSLV5Req)

	return err
}

type NullableCreateTradingBotSignalAmendTPSLV5Req struct {
	value *CreateTradingBotSignalAmendTPSLV5Req
	isSet bool
}

func (v NullableCreateTradingBotSignalAmendTPSLV5Req) Get() *CreateTradingBotSignalAmendTPSLV5Req {
	return v.value
}

func (v *NullableCreateTradingBotSignalAmendTPSLV5Req) Set(val *CreateTradingBotSignalAmendTPSLV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotSignalAmendTPSLV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotSignalAmendTPSLV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotSignalAmendTPSLV5Req(val *CreateTradingBotSignalAmendTPSLV5Req) *NullableCreateTradingBotSignalAmendTPSLV5Req {
	return &NullableCreateTradingBotSignalAmendTPSLV5Req{value: val, isSet: true}
}

func (v NullableCreateTradingBotSignalAmendTPSLV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotSignalAmendTPSLV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


