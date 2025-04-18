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

// checks if the CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner{}

// CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner struct for CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner
type CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner struct {
	// The code of the event execution result, `0` means success.
	SCode *string `json:"sCode,omitempty"`
	// Rejection or success message of event execution.
	SMsg *string `json:"sMsg,omitempty"`
	// Order ID
	SignalOrdId *string `json:"signalOrdId,omitempty"`
}

// NewCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner instantiates a new CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner() *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner {
	this := CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner{}
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	var signalOrdId string = ""
	this.SignalOrdId = &signalOrdId
	return &this
}

// NewCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInnerWithDefaults instantiates a new CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInnerWithDefaults() *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner {
	this := CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner{}
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	var signalOrdId string = ""
	this.SignalOrdId = &signalOrdId
	return &this
}

// GetSCode returns the SCode field value if set, zero value otherwise.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) GetSCode() string {
	if o == nil || IsNil(o.SCode) {
		var ret string
		return ret
	}
	return *o.SCode
}

// GetSCodeOk returns a tuple with the SCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) GetSCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SCode) {
		return nil, false
	}
	return o.SCode, true
}

// HasSCode returns a boolean if a field has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) HasSCode() bool {
	if o != nil && !IsNil(o.SCode) {
		return true
	}

	return false
}

// SetSCode gets a reference to the given string and assigns it to the SCode field.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) SetSCode(v string) {
	o.SCode = &v
}

// GetSMsg returns the SMsg field value if set, zero value otherwise.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) GetSMsg() string {
	if o == nil || IsNil(o.SMsg) {
		var ret string
		return ret
	}
	return *o.SMsg
}

// GetSMsgOk returns a tuple with the SMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) GetSMsgOk() (*string, bool) {
	if o == nil || IsNil(o.SMsg) {
		return nil, false
	}
	return o.SMsg, true
}

// HasSMsg returns a boolean if a field has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) HasSMsg() bool {
	if o != nil && !IsNil(o.SMsg) {
		return true
	}

	return false
}

// SetSMsg gets a reference to the given string and assigns it to the SMsg field.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) SetSMsg(v string) {
	o.SMsg = &v
}

// GetSignalOrdId returns the SignalOrdId field value if set, zero value otherwise.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) GetSignalOrdId() string {
	if o == nil || IsNil(o.SignalOrdId) {
		var ret string
		return ret
	}
	return *o.SignalOrdId
}

// GetSignalOrdIdOk returns a tuple with the SignalOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) GetSignalOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.SignalOrdId) {
		return nil, false
	}
	return o.SignalOrdId, true
}

// HasSignalOrdId returns a boolean if a field has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) HasSignalOrdId() bool {
	if o != nil && !IsNil(o.SignalOrdId) {
		return true
	}

	return false
}

// SetSignalOrdId gets a reference to the given string and assigns it to the SignalOrdId field.
func (o *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) SetSignalOrdId(v string) {
	o.SignalOrdId = &v
}

func (o CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) ToMap() (map[string]interface{}, error) {
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

type NullableCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner struct {
	value *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner
	isSet bool
}

func (v NullableCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) Get() *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner {
	return v.value
}

func (v *NullableCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) Set(val *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner(val *CreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) *NullableCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner {
	return &NullableCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner{value: val, isSet: true}
}

func (v NullableCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotSignalCancelSubOrderV5RespDataInnerDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


