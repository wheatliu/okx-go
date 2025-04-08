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

// checks if the CreateRfqCancelBatchQuotesV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRfqCancelBatchQuotesV5RespData{}

// CreateRfqCancelBatchQuotesV5RespData struct for CreateRfqCancelBatchQuotesV5RespData
type CreateRfqCancelBatchQuotesV5RespData struct {
	// The result code, `0` means success.
	Code *string `json:"code,omitempty"`
	// Array of objects containing the results
	Data []CreateRfqCancelBatchQuotesV5RespDataDataInner `json:"data,omitempty"`
	// The error message, not empty if the code is not 0.
	Msg *string `json:"msg,omitempty"`
}

// NewCreateRfqCancelBatchQuotesV5RespData instantiates a new CreateRfqCancelBatchQuotesV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRfqCancelBatchQuotesV5RespData() *CreateRfqCancelBatchQuotesV5RespData {
	this := CreateRfqCancelBatchQuotesV5RespData{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewCreateRfqCancelBatchQuotesV5RespDataWithDefaults instantiates a new CreateRfqCancelBatchQuotesV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRfqCancelBatchQuotesV5RespDataWithDefaults() *CreateRfqCancelBatchQuotesV5RespData {
	this := CreateRfqCancelBatchQuotesV5RespData{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateRfqCancelBatchQuotesV5RespData) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCancelBatchQuotesV5RespData) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateRfqCancelBatchQuotesV5RespData) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreateRfqCancelBatchQuotesV5RespData) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateRfqCancelBatchQuotesV5RespData) GetData() []CreateRfqCancelBatchQuotesV5RespDataDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []CreateRfqCancelBatchQuotesV5RespDataDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCancelBatchQuotesV5RespData) GetDataOk() ([]CreateRfqCancelBatchQuotesV5RespDataDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateRfqCancelBatchQuotesV5RespData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CreateRfqCancelBatchQuotesV5RespDataDataInner and assigns it to the Data field.
func (o *CreateRfqCancelBatchQuotesV5RespData) SetData(v []CreateRfqCancelBatchQuotesV5RespDataDataInner) {
	o.Data = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CreateRfqCancelBatchQuotesV5RespData) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRfqCancelBatchQuotesV5RespData) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CreateRfqCancelBatchQuotesV5RespData) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CreateRfqCancelBatchQuotesV5RespData) SetMsg(v string) {
	o.Msg = &v
}

func (o CreateRfqCancelBatchQuotesV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRfqCancelBatchQuotesV5RespData) ToMap() (map[string]interface{}, error) {
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

type NullableCreateRfqCancelBatchQuotesV5RespData struct {
	value *CreateRfqCancelBatchQuotesV5RespData
	isSet bool
}

func (v NullableCreateRfqCancelBatchQuotesV5RespData) Get() *CreateRfqCancelBatchQuotesV5RespData {
	return v.value
}

func (v *NullableCreateRfqCancelBatchQuotesV5RespData) Set(val *CreateRfqCancelBatchQuotesV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRfqCancelBatchQuotesV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRfqCancelBatchQuotesV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRfqCancelBatchQuotesV5RespData(val *CreateRfqCancelBatchQuotesV5RespData) *NullableCreateRfqCancelBatchQuotesV5RespData {
	return &NullableCreateRfqCancelBatchQuotesV5RespData{value: val, isSet: true}
}

func (v NullableCreateRfqCancelBatchQuotesV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRfqCancelBatchQuotesV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


