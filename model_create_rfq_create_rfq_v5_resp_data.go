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

// checks if the CreateRfqCreateRfqV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRfqCreateRfqV5RespData{}

// CreateRfqCreateRfqV5RespData struct for CreateRfqCreateRfqV5RespData
type CreateRfqCreateRfqV5RespData struct {
	// The result code, `0` means success.
	Code *string `json:"code,omitempty"`
	// Array of objects containing the results of the RFQ creation.
	Data []CreateRfqCreateRfqV5RespDataDataInner `json:"data,omitempty"`
	// The error message, not empty if the code is not 0.
	Msg *string `json:"msg,omitempty"`
}

// NewCreateRfqCreateRfqV5RespData instantiates a new CreateRfqCreateRfqV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRfqCreateRfqV5RespData() *CreateRfqCreateRfqV5RespData {
	this := CreateRfqCreateRfqV5RespData{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewCreateRfqCreateRfqV5RespDataWithDefaults instantiates a new CreateRfqCreateRfqV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRfqCreateRfqV5RespDataWithDefaults() *CreateRfqCreateRfqV5RespData {
	this := CreateRfqCreateRfqV5RespData{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespData) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespData) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespData) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreateRfqCreateRfqV5RespData) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespData) GetData() []CreateRfqCreateRfqV5RespDataDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []CreateRfqCreateRfqV5RespDataDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespData) GetDataOk() ([]CreateRfqCreateRfqV5RespDataDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CreateRfqCreateRfqV5RespDataDataInner and assigns it to the Data field.
func (o *CreateRfqCreateRfqV5RespData) SetData(v []CreateRfqCreateRfqV5RespDataDataInner) {
	o.Data = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CreateRfqCreateRfqV5RespData) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCreateRfqV5RespData) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CreateRfqCreateRfqV5RespData) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CreateRfqCreateRfqV5RespData) SetMsg(v string) {
	o.Msg = &v
}

func (o CreateRfqCreateRfqV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRfqCreateRfqV5RespData) ToMap() (map[string]interface{}, error) {
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

type NullableCreateRfqCreateRfqV5RespData struct {
	value *CreateRfqCreateRfqV5RespData
	isSet bool
}

func (v NullableCreateRfqCreateRfqV5RespData) Get() *CreateRfqCreateRfqV5RespData {
	return v.value
}

func (v *NullableCreateRfqCreateRfqV5RespData) Set(val *CreateRfqCreateRfqV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRfqCreateRfqV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRfqCreateRfqV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRfqCreateRfqV5RespData(val *CreateRfqCreateRfqV5RespData) *NullableCreateRfqCreateRfqV5RespData {
	return &NullableCreateRfqCreateRfqV5RespData{value: val, isSet: true}
}

func (v NullableCreateRfqCreateRfqV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRfqCreateRfqV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


