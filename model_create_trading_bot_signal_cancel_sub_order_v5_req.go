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

// checks if the CreateTradingBotSignalCancelSubOrderV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotSignalCancelSubOrderV5Req{}

// CreateTradingBotSignalCancelSubOrderV5Req struct for CreateTradingBotSignalCancelSubOrderV5Req
type CreateTradingBotSignalCancelSubOrderV5Req struct {
	// Algo ID
	AlgoId string `json:"algoId"`
	// Instrument ID, e.g. BTC-USDT-SWAP
	InstId string `json:"instId"`
	// Order ID
	SignalOrdId string `json:"signalOrdId"`
}

type _CreateTradingBotSignalCancelSubOrderV5Req CreateTradingBotSignalCancelSubOrderV5Req

// NewCreateTradingBotSignalCancelSubOrderV5Req instantiates a new CreateTradingBotSignalCancelSubOrderV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotSignalCancelSubOrderV5Req(algoId string, instId string, signalOrdId string) *CreateTradingBotSignalCancelSubOrderV5Req {
	this := CreateTradingBotSignalCancelSubOrderV5Req{}
	this.AlgoId = algoId
	this.InstId = instId
	this.SignalOrdId = signalOrdId
	return &this
}

// NewCreateTradingBotSignalCancelSubOrderV5ReqWithDefaults instantiates a new CreateTradingBotSignalCancelSubOrderV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotSignalCancelSubOrderV5ReqWithDefaults() *CreateTradingBotSignalCancelSubOrderV5Req {
	this := CreateTradingBotSignalCancelSubOrderV5Req{}
	var algoId string = ""
	this.AlgoId = algoId
	var instId string = ""
	this.InstId = instId
	var signalOrdId string = ""
	this.SignalOrdId = signalOrdId
	return &this
}

// GetAlgoId returns the AlgoId field value
func (o *CreateTradingBotSignalCancelSubOrderV5Req) GetAlgoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5Req) GetAlgoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlgoId, true
}

// SetAlgoId sets field value
func (o *CreateTradingBotSignalCancelSubOrderV5Req) SetAlgoId(v string) {
	o.AlgoId = v
}

// GetInstId returns the InstId field value
func (o *CreateTradingBotSignalCancelSubOrderV5Req) GetInstId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5Req) GetInstIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstId, true
}

// SetInstId sets field value
func (o *CreateTradingBotSignalCancelSubOrderV5Req) SetInstId(v string) {
	o.InstId = v
}

// GetSignalOrdId returns the SignalOrdId field value
func (o *CreateTradingBotSignalCancelSubOrderV5Req) GetSignalOrdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignalOrdId
}

// GetSignalOrdIdOk returns a tuple with the SignalOrdId field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5Req) GetSignalOrdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalOrdId, true
}

// SetSignalOrdId sets field value
func (o *CreateTradingBotSignalCancelSubOrderV5Req) SetSignalOrdId(v string) {
	o.SignalOrdId = v
}

func (o CreateTradingBotSignalCancelSubOrderV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotSignalCancelSubOrderV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["algoId"] = o.AlgoId
	toSerialize["instId"] = o.InstId
	toSerialize["signalOrdId"] = o.SignalOrdId
	return toSerialize, nil
}

func (o *CreateTradingBotSignalCancelSubOrderV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"algoId",
		"instId",
		"signalOrdId",
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

	varCreateTradingBotSignalCancelSubOrderV5Req := _CreateTradingBotSignalCancelSubOrderV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTradingBotSignalCancelSubOrderV5Req)

	if err != nil {
		return err
	}

	*o = CreateTradingBotSignalCancelSubOrderV5Req(varCreateTradingBotSignalCancelSubOrderV5Req)

	return err
}

type NullableCreateTradingBotSignalCancelSubOrderV5Req struct {
	value *CreateTradingBotSignalCancelSubOrderV5Req
	isSet bool
}

func (v NullableCreateTradingBotSignalCancelSubOrderV5Req) Get() *CreateTradingBotSignalCancelSubOrderV5Req {
	return v.value
}

func (v *NullableCreateTradingBotSignalCancelSubOrderV5Req) Set(val *CreateTradingBotSignalCancelSubOrderV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotSignalCancelSubOrderV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotSignalCancelSubOrderV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotSignalCancelSubOrderV5Req(val *CreateTradingBotSignalCancelSubOrderV5Req) *NullableCreateTradingBotSignalCancelSubOrderV5Req {
	return &NullableCreateTradingBotSignalCancelSubOrderV5Req{value: val, isSet: true}
}

func (v NullableCreateTradingBotSignalCancelSubOrderV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotSignalCancelSubOrderV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


