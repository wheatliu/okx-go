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

// checks if the CreateTradeCancelAdvanceAlgosV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradeCancelAdvanceAlgosV5RespDataInner{}

// CreateTradeCancelAdvanceAlgosV5RespDataInner struct for CreateTradeCancelAdvanceAlgosV5RespDataInner
type CreateTradeCancelAdvanceAlgosV5RespDataInner struct {
	// Order ID
	AlgoId *string `json:"algoId,omitempty"`
	// The code of the event execution result, `0` means success.
	SCode *string `json:"sCode,omitempty"`
	// Rejection message if the request is unsuccessful.
	SMsg *string `json:"sMsg,omitempty"`
}

// NewCreateTradeCancelAdvanceAlgosV5RespDataInner instantiates a new CreateTradeCancelAdvanceAlgosV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeCancelAdvanceAlgosV5RespDataInner() *CreateTradeCancelAdvanceAlgosV5RespDataInner {
	this := CreateTradeCancelAdvanceAlgosV5RespDataInner{}
	var algoId string = ""
	this.AlgoId = &algoId
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	return &this
}

// NewCreateTradeCancelAdvanceAlgosV5RespDataInnerWithDefaults instantiates a new CreateTradeCancelAdvanceAlgosV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeCancelAdvanceAlgosV5RespDataInnerWithDefaults() *CreateTradeCancelAdvanceAlgosV5RespDataInner {
	this := CreateTradeCancelAdvanceAlgosV5RespDataInner{}
	var algoId string = ""
	this.AlgoId = &algoId
	var sCode string = ""
	this.SCode = &sCode
	var sMsg string = ""
	this.SMsg = &sMsg
	return &this
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *CreateTradeCancelAdvanceAlgosV5RespDataInner) GetAlgoId() string {
	if o == nil || IsNil(o.AlgoId) {
		var ret string
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelAdvanceAlgosV5RespDataInner) GetAlgoIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *CreateTradeCancelAdvanceAlgosV5RespDataInner) HasAlgoId() bool {
	if o != nil && !IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given string and assigns it to the AlgoId field.
func (o *CreateTradeCancelAdvanceAlgosV5RespDataInner) SetAlgoId(v string) {
	o.AlgoId = &v
}

// GetSCode returns the SCode field value if set, zero value otherwise.
func (o *CreateTradeCancelAdvanceAlgosV5RespDataInner) GetSCode() string {
	if o == nil || IsNil(o.SCode) {
		var ret string
		return ret
	}
	return *o.SCode
}

// GetSCodeOk returns a tuple with the SCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelAdvanceAlgosV5RespDataInner) GetSCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SCode) {
		return nil, false
	}
	return o.SCode, true
}

// HasSCode returns a boolean if a field has been set.
func (o *CreateTradeCancelAdvanceAlgosV5RespDataInner) HasSCode() bool {
	if o != nil && !IsNil(o.SCode) {
		return true
	}

	return false
}

// SetSCode gets a reference to the given string and assigns it to the SCode field.
func (o *CreateTradeCancelAdvanceAlgosV5RespDataInner) SetSCode(v string) {
	o.SCode = &v
}

// GetSMsg returns the SMsg field value if set, zero value otherwise.
func (o *CreateTradeCancelAdvanceAlgosV5RespDataInner) GetSMsg() string {
	if o == nil || IsNil(o.SMsg) {
		var ret string
		return ret
	}
	return *o.SMsg
}

// GetSMsgOk returns a tuple with the SMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeCancelAdvanceAlgosV5RespDataInner) GetSMsgOk() (*string, bool) {
	if o == nil || IsNil(o.SMsg) {
		return nil, false
	}
	return o.SMsg, true
}

// HasSMsg returns a boolean if a field has been set.
func (o *CreateTradeCancelAdvanceAlgosV5RespDataInner) HasSMsg() bool {
	if o != nil && !IsNil(o.SMsg) {
		return true
	}

	return false
}

// SetSMsg gets a reference to the given string and assigns it to the SMsg field.
func (o *CreateTradeCancelAdvanceAlgosV5RespDataInner) SetSMsg(v string) {
	o.SMsg = &v
}

func (o CreateTradeCancelAdvanceAlgosV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradeCancelAdvanceAlgosV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableCreateTradeCancelAdvanceAlgosV5RespDataInner struct {
	value *CreateTradeCancelAdvanceAlgosV5RespDataInner
	isSet bool
}

func (v NullableCreateTradeCancelAdvanceAlgosV5RespDataInner) Get() *CreateTradeCancelAdvanceAlgosV5RespDataInner {
	return v.value
}

func (v *NullableCreateTradeCancelAdvanceAlgosV5RespDataInner) Set(val *CreateTradeCancelAdvanceAlgosV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeCancelAdvanceAlgosV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeCancelAdvanceAlgosV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeCancelAdvanceAlgosV5RespDataInner(val *CreateTradeCancelAdvanceAlgosV5RespDataInner) *NullableCreateTradeCancelAdvanceAlgosV5RespDataInner {
	return &NullableCreateTradeCancelAdvanceAlgosV5RespDataInner{value: val, isSet: true}
}

func (v NullableCreateTradeCancelAdvanceAlgosV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeCancelAdvanceAlgosV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


