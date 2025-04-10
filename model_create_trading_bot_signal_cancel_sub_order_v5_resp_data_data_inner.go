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

// checks if the CreateTradingBotSignalCancelSubOrderV5RespDataDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotSignalCancelSubOrderV5RespDataDataInner{}

// CreateTradingBotSignalCancelSubOrderV5RespDataDataInner struct for CreateTradingBotSignalCancelSubOrderV5RespDataDataInner
type CreateTradingBotSignalCancelSubOrderV5RespDataDataInner struct {
	// The code of the event execution result, `0` means success.
	SCode *string `json:"sCode,omitempty"`
	// Rejection or success message of event execution.
	SMsg *string `json:"sMsg,omitempty"`
	// Order ID
	SignalOrdId *string `json:"signalOrdId,omitempty"`
}

// NewCreateTradingBotSignalCancelSubOrderV5RespDataDataInner instantiates a new CreateTradingBotSignalCancelSubOrderV5RespDataDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotSignalCancelSubOrderV5RespDataDataInner() *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner {
	this := CreateTradingBotSignalCancelSubOrderV5RespDataDataInner{}
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	var signalOrdId string = ""
	this.SignalOrdId = &signalOrdId
	return &this
}

// NewCreateTradingBotSignalCancelSubOrderV5RespDataDataInnerWithDefaults instantiates a new CreateTradingBotSignalCancelSubOrderV5RespDataDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotSignalCancelSubOrderV5RespDataDataInnerWithDefaults() *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner {
	this := CreateTradingBotSignalCancelSubOrderV5RespDataDataInner{}
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	var signalOrdId string = ""
	this.SignalOrdId = &signalOrdId
	return &this
}

// GetSCode returns the SCode field value if set, zero value otherwise.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) GetSCode() string {
	if o == nil || IsNil(o.SCode) {
		var ret string
		return ret
	}
	return *o.SCode
}

// GetSCodeOk returns a tuple with the SCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) GetSCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SCode) {
		return nil, false
	}
	return o.SCode, true
}

// HasSCode returns a boolean if a field has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) HasSCode() bool {
	if o != nil && !IsNil(o.SCode) {
		return true
	}

	return false
}

// SetSCode gets a reference to the given string and assigns it to the SCode field.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) SetSCode(v string) {
	o.SCode = &v
}

// GetSMsg returns the SMsg field value if set, zero value otherwise.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) GetSMsg() string {
	if o == nil || IsNil(o.SMsg) {
		var ret string
		return ret
	}
	return *o.SMsg
}

// GetSMsgOk returns a tuple with the SMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) GetSMsgOk() (*string, bool) {
	if o == nil || IsNil(o.SMsg) {
		return nil, false
	}
	return o.SMsg, true
}

// HasSMsg returns a boolean if a field has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) HasSMsg() bool {
	if o != nil && !IsNil(o.SMsg) {
		return true
	}

	return false
}

// SetSMsg gets a reference to the given string and assigns it to the SMsg field.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) SetSMsg(v string) {
	o.SMsg = &v
}

// GetSignalOrdId returns the SignalOrdId field value if set, zero value otherwise.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) GetSignalOrdId() string {
	if o == nil || IsNil(o.SignalOrdId) {
		var ret string
		return ret
	}
	return *o.SignalOrdId
}

// GetSignalOrdIdOk returns a tuple with the SignalOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) GetSignalOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.SignalOrdId) {
		return nil, false
	}
	return o.SignalOrdId, true
}

// HasSignalOrdId returns a boolean if a field has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) HasSignalOrdId() bool {
	if o != nil && !IsNil(o.SignalOrdId) {
		return true
	}

	return false
}

// SetSignalOrdId gets a reference to the given string and assigns it to the SignalOrdId field.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) SetSignalOrdId(v string) {
	o.SignalOrdId = &v
}

func (o CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SCode) {
		toSerialize["sCode"] = o.SCode
	}
	if !IsNil(o.SMsg) {
		toSerialize["sMsg"] = o.SMsg
	}
	if !IsNil(o.SignalOrdId) {
		toSerialize["signalOrdId"] = o.SignalOrdId
	}
	return toSerialize, nil
}

type NullableCreateTradingBotSignalCancelSubOrderV5RespDataDataInner struct {
	value *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner
	isSet bool
}

func (v NullableCreateTradingBotSignalCancelSubOrderV5RespDataDataInner) Get() *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner {
	return v.value
}

func (v *NullableCreateTradingBotSignalCancelSubOrderV5RespDataDataInner) Set(val *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotSignalCancelSubOrderV5RespDataDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotSignalCancelSubOrderV5RespDataDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotSignalCancelSubOrderV5RespDataDataInner(val *CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) *NullableCreateTradingBotSignalCancelSubOrderV5RespDataDataInner {
	return &NullableCreateTradingBotSignalCancelSubOrderV5RespDataDataInner{value: val, isSet: true}
}

func (v NullableCreateTradingBotSignalCancelSubOrderV5RespDataDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotSignalCancelSubOrderV5RespDataDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


