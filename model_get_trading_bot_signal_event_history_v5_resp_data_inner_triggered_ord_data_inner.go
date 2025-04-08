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

// checks if the GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner{}

// GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner struct for GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner
type GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner struct {
	// Sub order client-supplied id
	ClOrdId *string `json:"clOrdId,omitempty"`
}

// NewGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner instantiates a new GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner() *GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner {
	this := GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner{}
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	return &this
}

// NewGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInnerWithDefaults instantiates a new GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInnerWithDefaults() *GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner {
	this := GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner{}
	var clOrdId string = ""
	this.ClOrdId = &clOrdId
	return &this
}

// GetClOrdId returns the ClOrdId field value if set, zero value otherwise.
func (o *GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) GetClOrdId() string {
	if o == nil || IsNil(o.ClOrdId) {
		var ret string
		return ret
	}
	return *o.ClOrdId
}

// GetClOrdIdOk returns a tuple with the ClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) GetClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClOrdId) {
		return nil, false
	}
	return o.ClOrdId, true
}

// HasClOrdId returns a boolean if a field has been set.
func (o *GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) HasClOrdId() bool {
	if o != nil && !IsNil(o.ClOrdId) {
		return true
	}

	return false
}

// SetClOrdId gets a reference to the given string and assigns it to the ClOrdId field.
func (o *GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) SetClOrdId(v string) {
	o.ClOrdId = &v
}

func (o GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClOrdId) {
		toSerialize["clOrdId"] = o.ClOrdId
	}
	return toSerialize, nil
}

type NullableGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner struct {
	value *GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner
	isSet bool
}

func (v NullableGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) Get() *GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner {
	return v.value
}

func (v *NullableGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) Set(val *GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner(val *GetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) *NullableGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner {
	return &NullableGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner{value: val, isSet: true}
}

func (v NullableGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTradingBotSignalEventHistoryV5RespDataInnerTriggeredOrdDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


