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

// checks if the CreateTradingBotSignalSetInstrumentsV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotSignalSetInstrumentsV5Req{}

// CreateTradingBotSignalSetInstrumentsV5Req struct for CreateTradingBotSignalSetInstrumentsV5Req
type CreateTradingBotSignalSetInstrumentsV5Req struct {
	// Algo ID
	AlgoId string `json:"algoId"`
	// Whether to include all USDT-margined contract.The default value is `false`. `true`: include `false` : exclude
	IncludeAll bool `json:"includeAll"`
	// Instrument IDs. When `includeAll` is `true`, it is ignored
	InstIds []string `json:"instIds"`
}

type _CreateTradingBotSignalSetInstrumentsV5Req CreateTradingBotSignalSetInstrumentsV5Req

// NewCreateTradingBotSignalSetInstrumentsV5Req instantiates a new CreateTradingBotSignalSetInstrumentsV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotSignalSetInstrumentsV5Req(algoId string, includeAll bool, instIds []string) *CreateTradingBotSignalSetInstrumentsV5Req {
	this := CreateTradingBotSignalSetInstrumentsV5Req{}
	this.AlgoId = algoId
	this.IncludeAll = includeAll
	this.InstIds = instIds
	return &this
}

// NewCreateTradingBotSignalSetInstrumentsV5ReqWithDefaults instantiates a new CreateTradingBotSignalSetInstrumentsV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotSignalSetInstrumentsV5ReqWithDefaults() *CreateTradingBotSignalSetInstrumentsV5Req {
	this := CreateTradingBotSignalSetInstrumentsV5Req{}
	var algoId string = ""
	this.AlgoId = algoId
	return &this
}

// GetAlgoId returns the AlgoId field value
func (o *CreateTradingBotSignalSetInstrumentsV5Req) GetAlgoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalSetInstrumentsV5Req) GetAlgoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlgoId, true
}

// SetAlgoId sets field value
func (o *CreateTradingBotSignalSetInstrumentsV5Req) SetAlgoId(v string) {
	o.AlgoId = v
}

// GetIncludeAll returns the IncludeAll field value
func (o *CreateTradingBotSignalSetInstrumentsV5Req) GetIncludeAll() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeAll
}

// GetIncludeAllOk returns a tuple with the IncludeAll field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalSetInstrumentsV5Req) GetIncludeAllOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeAll, true
}

// SetIncludeAll sets field value
func (o *CreateTradingBotSignalSetInstrumentsV5Req) SetIncludeAll(v bool) {
	o.IncludeAll = v
}

// GetInstIds returns the InstIds field value
func (o *CreateTradingBotSignalSetInstrumentsV5Req) GetInstIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.InstIds
}

// GetInstIdsOk returns a tuple with the InstIds field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalSetInstrumentsV5Req) GetInstIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstIds, true
}

// SetInstIds sets field value
func (o *CreateTradingBotSignalSetInstrumentsV5Req) SetInstIds(v []string) {
	o.InstIds = v
}

func (o CreateTradingBotSignalSetInstrumentsV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotSignalSetInstrumentsV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["algoId"] = o.AlgoId
	toSerialize["includeAll"] = o.IncludeAll
	toSerialize["instIds"] = o.InstIds
	return toSerialize, nil
}

func (o *CreateTradingBotSignalSetInstrumentsV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"algoId",
		"includeAll",
		"instIds",
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

	varCreateTradingBotSignalSetInstrumentsV5Req := _CreateTradingBotSignalSetInstrumentsV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTradingBotSignalSetInstrumentsV5Req)

	if err != nil {
		return err
	}

	*o = CreateTradingBotSignalSetInstrumentsV5Req(varCreateTradingBotSignalSetInstrumentsV5Req)

	return err
}

type NullableCreateTradingBotSignalSetInstrumentsV5Req struct {
	value *CreateTradingBotSignalSetInstrumentsV5Req
	isSet bool
}

func (v NullableCreateTradingBotSignalSetInstrumentsV5Req) Get() *CreateTradingBotSignalSetInstrumentsV5Req {
	return v.value
}

func (v *NullableCreateTradingBotSignalSetInstrumentsV5Req) Set(val *CreateTradingBotSignalSetInstrumentsV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotSignalSetInstrumentsV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotSignalSetInstrumentsV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotSignalSetInstrumentsV5Req(val *CreateTradingBotSignalSetInstrumentsV5Req) *NullableCreateTradingBotSignalSetInstrumentsV5Req {
	return &NullableCreateTradingBotSignalSetInstrumentsV5Req{value: val, isSet: true}
}

func (v NullableCreateTradingBotSignalSetInstrumentsV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotSignalSetInstrumentsV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


