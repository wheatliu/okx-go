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

// checks if the CreateTradingBotSignalCancelSubOrderV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotSignalCancelSubOrderV5RespData{}

// CreateTradingBotSignalCancelSubOrderV5RespData struct for CreateTradingBotSignalCancelSubOrderV5RespData
type CreateTradingBotSignalCancelSubOrderV5RespData struct {
	// The result code, `0` means success
	Code *string `json:"code,omitempty"`
	// Array of objects contains the response results
	Data []CreateTradingBotSignalCancelSubOrderV5RespDataDataInner `json:"data,omitempty"`
	// The error message, empty if the code is 0
	Msg *string `json:"msg,omitempty"`
}

// NewCreateTradingBotSignalCancelSubOrderV5RespData instantiates a new CreateTradingBotSignalCancelSubOrderV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotSignalCancelSubOrderV5RespData() *CreateTradingBotSignalCancelSubOrderV5RespData {
	this := CreateTradingBotSignalCancelSubOrderV5RespData{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewCreateTradingBotSignalCancelSubOrderV5RespDataWithDefaults instantiates a new CreateTradingBotSignalCancelSubOrderV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotSignalCancelSubOrderV5RespDataWithDefaults() *CreateTradingBotSignalCancelSubOrderV5RespData {
	this := CreateTradingBotSignalCancelSubOrderV5RespData{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateTradingBotSignalCancelSubOrderV5RespData) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespData) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespData) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreateTradingBotSignalCancelSubOrderV5RespData) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateTradingBotSignalCancelSubOrderV5RespData) GetData() []CreateTradingBotSignalCancelSubOrderV5RespDataDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []CreateTradingBotSignalCancelSubOrderV5RespDataDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespData) GetDataOk() ([]CreateTradingBotSignalCancelSubOrderV5RespDataDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CreateTradingBotSignalCancelSubOrderV5RespDataDataInner and assigns it to the Data field.
func (o *CreateTradingBotSignalCancelSubOrderV5RespData) SetData(v []CreateTradingBotSignalCancelSubOrderV5RespDataDataInner) {
	o.Data = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CreateTradingBotSignalCancelSubOrderV5RespData) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespData) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CreateTradingBotSignalCancelSubOrderV5RespData) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CreateTradingBotSignalCancelSubOrderV5RespData) SetMsg(v string) {
	o.Msg = &v
}

func (o CreateTradingBotSignalCancelSubOrderV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotSignalCancelSubOrderV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return toSerialize, nil
}

type NullableCreateTradingBotSignalCancelSubOrderV5RespData struct {
	value *CreateTradingBotSignalCancelSubOrderV5RespData
	isSet bool
}

func (v NullableCreateTradingBotSignalCancelSubOrderV5RespData) Get() *CreateTradingBotSignalCancelSubOrderV5RespData {
	return v.value
}

func (v *NullableCreateTradingBotSignalCancelSubOrderV5RespData) Set(val *CreateTradingBotSignalCancelSubOrderV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotSignalCancelSubOrderV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotSignalCancelSubOrderV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotSignalCancelSubOrderV5RespData(val *CreateTradingBotSignalCancelSubOrderV5RespData) *NullableCreateTradingBotSignalCancelSubOrderV5RespData {
	return &NullableCreateTradingBotSignalCancelSubOrderV5RespData{value: val, isSet: true}
}

func (v NullableCreateTradingBotSignalCancelSubOrderV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotSignalCancelSubOrderV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


