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

// checks if the CreateTradingBotSignalCreateSignalV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotSignalCreateSignalV5RespDataInner{}

// CreateTradingBotSignalCreateSignalV5RespDataInner struct for CreateTradingBotSignalCreateSignalV5RespDataInner
type CreateTradingBotSignalCreateSignalV5RespDataInner struct {
	// Signal channel Id
	SignalChanId *string `json:"signalChanId,omitempty"`
	// User identify when placing orders via signal
	SignalChanToken *string `json:"signalChanToken,omitempty"`
}

// NewCreateTradingBotSignalCreateSignalV5RespDataInner instantiates a new CreateTradingBotSignalCreateSignalV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotSignalCreateSignalV5RespDataInner() *CreateTradingBotSignalCreateSignalV5RespDataInner {
	this := CreateTradingBotSignalCreateSignalV5RespDataInner{}
	var signalChanId string = ""
	this.SignalChanId = &signalChanId
	var signalChanToken string = ""
	this.SignalChanToken = &signalChanToken
	return &this
}

// NewCreateTradingBotSignalCreateSignalV5RespDataInnerWithDefaults instantiates a new CreateTradingBotSignalCreateSignalV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotSignalCreateSignalV5RespDataInnerWithDefaults() *CreateTradingBotSignalCreateSignalV5RespDataInner {
	this := CreateTradingBotSignalCreateSignalV5RespDataInner{}
	var signalChanId string = ""
	this.SignalChanId = &signalChanId
	var signalChanToken string = ""
	this.SignalChanToken = &signalChanToken
	return &this
}

// GetSignalChanId returns the SignalChanId field value if set, zero value otherwise.
func (o *CreateTradingBotSignalCreateSignalV5RespDataInner) GetSignalChanId() string {
	if o == nil || IsNil(o.SignalChanId) {
		var ret string
		return ret
	}
	return *o.SignalChanId
}

// GetSignalChanIdOk returns a tuple with the SignalChanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCreateSignalV5RespDataInner) GetSignalChanIdOk() (*string, bool) {
	if o == nil || IsNil(o.SignalChanId) {
		return nil, false
	}
	return o.SignalChanId, true
}

// HasSignalChanId returns a boolean if a field has been set.
func (o *CreateTradingBotSignalCreateSignalV5RespDataInner) HasSignalChanId() bool {
	if o != nil && !IsNil(o.SignalChanId) {
		return true
	}

	return false
}

// SetSignalChanId gets a reference to the given string and assigns it to the SignalChanId field.
func (o *CreateTradingBotSignalCreateSignalV5RespDataInner) SetSignalChanId(v string) {
	o.SignalChanId = &v
}

// GetSignalChanToken returns the SignalChanToken field value if set, zero value otherwise.
func (o *CreateTradingBotSignalCreateSignalV5RespDataInner) GetSignalChanToken() string {
	if o == nil || IsNil(o.SignalChanToken) {
		var ret string
		return ret
	}
	return *o.SignalChanToken
}

// GetSignalChanTokenOk returns a tuple with the SignalChanToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCreateSignalV5RespDataInner) GetSignalChanTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SignalChanToken) {
		return nil, false
	}
	return o.SignalChanToken, true
}

// HasSignalChanToken returns a boolean if a field has been set.
func (o *CreateTradingBotSignalCreateSignalV5RespDataInner) HasSignalChanToken() bool {
	if o != nil && !IsNil(o.SignalChanToken) {
		return true
	}

	return false
}

// SetSignalChanToken gets a reference to the given string and assigns it to the SignalChanToken field.
func (o *CreateTradingBotSignalCreateSignalV5RespDataInner) SetSignalChanToken(v string) {
	o.SignalChanToken = &v
}

func (o CreateTradingBotSignalCreateSignalV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotSignalCreateSignalV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SignalChanId) {
		toSerialize["signalChanId"] = o.SignalChanId
	}
	if !IsNil(o.SignalChanToken) {
		toSerialize["signalChanToken"] = o.SignalChanToken
	}
	return toSerialize, nil
}

type NullableCreateTradingBotSignalCreateSignalV5RespDataInner struct {
	value *CreateTradingBotSignalCreateSignalV5RespDataInner
	isSet bool
}

func (v NullableCreateTradingBotSignalCreateSignalV5RespDataInner) Get() *CreateTradingBotSignalCreateSignalV5RespDataInner {
	return v.value
}

func (v *NullableCreateTradingBotSignalCreateSignalV5RespDataInner) Set(val *CreateTradingBotSignalCreateSignalV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotSignalCreateSignalV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotSignalCreateSignalV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotSignalCreateSignalV5RespDataInner(val *CreateTradingBotSignalCreateSignalV5RespDataInner) *NullableCreateTradingBotSignalCreateSignalV5RespDataInner {
	return &NullableCreateTradingBotSignalCreateSignalV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateTradingBotSignalCreateSignalV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotSignalCreateSignalV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


