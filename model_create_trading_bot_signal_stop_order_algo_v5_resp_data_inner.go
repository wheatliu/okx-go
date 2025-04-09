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

// checks if the CreateTradingBotSignalStopOrderAlgoV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotSignalStopOrderAlgoV5RespDataInner{}

// CreateTradingBotSignalStopOrderAlgoV5RespDataInner struct for CreateTradingBotSignalStopOrderAlgoV5RespDataInner
type CreateTradingBotSignalStopOrderAlgoV5RespDataInner struct {
	// Client-supplied Algo ID
	AlgoClOrdId *string `json:"algoClOrdId,omitempty"`
	// Algo ID
	AlgoId *string `json:"algoId,omitempty"`
	// The code of the event execution result, `0` means success.
	SCode *string `json:"sCode,omitempty"`
	// Rejection or success message of event execution.
	SMsg *string `json:"sMsg,omitempty"`
}

// NewCreateTradingBotSignalStopOrderAlgoV5RespDataInner instantiates a new CreateTradingBotSignalStopOrderAlgoV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotSignalStopOrderAlgoV5RespDataInner() *CreateTradingBotSignalStopOrderAlgoV5RespDataInner {
	this := CreateTradingBotSignalStopOrderAlgoV5RespDataInner{}
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	return &this
}

// NewCreateTradingBotSignalStopOrderAlgoV5RespDataInnerWithDefaults instantiates a new CreateTradingBotSignalStopOrderAlgoV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotSignalStopOrderAlgoV5RespDataInnerWithDefaults() *CreateTradingBotSignalStopOrderAlgoV5RespDataInner {
	this := CreateTradingBotSignalStopOrderAlgoV5RespDataInner{}
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoId string = ""
	this.AlgoId = &algoId
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	return &this
}

// GetAlgoClOrdId returns the AlgoClOrdId field value if set, zero value otherwise.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) GetAlgoClOrdId() string {
	if o == nil || IsNil(o.AlgoClOrdId) {
		var ret string
		return ret
	}
	return *o.AlgoClOrdId
}

// GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) GetAlgoClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoClOrdId) {
		return nil, false
	}
	return o.AlgoClOrdId, true
}

// HasAlgoClOrdId returns a boolean if a field has been set.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) HasAlgoClOrdId() bool {
	if o != nil && !IsNil(o.AlgoClOrdId) {
		return true
	}

	return false
}

// SetAlgoClOrdId gets a reference to the given string and assigns it to the AlgoClOrdId field.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) SetAlgoClOrdId(v string) {
	o.AlgoClOrdId = &v
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) GetAlgoId() string {
	if o == nil || IsNil(o.AlgoId) {
		var ret string
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) GetAlgoIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) HasAlgoId() bool {
	if o != nil && !IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given string and assigns it to the AlgoId field.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) SetAlgoId(v string) {
	o.AlgoId = &v
}

// GetSCode returns the SCode field value if set, zero value otherwise.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) GetSCode() string {
	if o == nil || IsNil(o.SCode) {
		var ret string
		return ret
	}
	return *o.SCode
}

// GetSCodeOk returns a tuple with the SCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) GetSCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SCode) {
		return nil, false
	}
	return o.SCode, true
}

// HasSCode returns a boolean if a field has been set.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) HasSCode() bool {
	if o != nil && !IsNil(o.SCode) {
		return true
	}

	return false
}

// SetSCode gets a reference to the given string and assigns it to the SCode field.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) SetSCode(v string) {
	o.SCode = &v
}

// GetSMsg returns the SMsg field value if set, zero value otherwise.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) GetSMsg() string {
	if o == nil || IsNil(o.SMsg) {
		var ret string
		return ret
	}
	return *o.SMsg
}

// GetSMsgOk returns a tuple with the SMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) GetSMsgOk() (*string, bool) {
	if o == nil || IsNil(o.SMsg) {
		return nil, false
	}
	return o.SMsg, true
}

// HasSMsg returns a boolean if a field has been set.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) HasSMsg() bool {
	if o != nil && !IsNil(o.SMsg) {
		return true
	}

	return false
}

// SetSMsg gets a reference to the given string and assigns it to the SMsg field.
func (o *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) SetSMsg(v string) {
	o.SMsg = &v
}

func (o CreateTradingBotSignalStopOrderAlgoV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotSignalStopOrderAlgoV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlgoClOrdId) {
		toSerialize["algoClOrdId"] = o.AlgoClOrdId
	}
	if !IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !IsNil(o.SCode) {
		toSerialize["sCode"] = o.SCode
	}
	if !IsNil(o.SMsg) {
		toSerialize["sMsg"] = o.SMsg
	}
	return toSerialize, nil
}

type NullableCreateTradingBotSignalStopOrderAlgoV5RespDataInner struct {
	value *CreateTradingBotSignalStopOrderAlgoV5RespDataInner
	isSet bool
}

func (v NullableCreateTradingBotSignalStopOrderAlgoV5RespDataInner) Get() *CreateTradingBotSignalStopOrderAlgoV5RespDataInner {
	return v.value
}

func (v *NullableCreateTradingBotSignalStopOrderAlgoV5RespDataInner) Set(val *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotSignalStopOrderAlgoV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotSignalStopOrderAlgoV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotSignalStopOrderAlgoV5RespDataInner(val *CreateTradingBotSignalStopOrderAlgoV5RespDataInner) *NullableCreateTradingBotSignalStopOrderAlgoV5RespDataInner {
	return &NullableCreateTradingBotSignalStopOrderAlgoV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateTradingBotSignalStopOrderAlgoV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotSignalStopOrderAlgoV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


