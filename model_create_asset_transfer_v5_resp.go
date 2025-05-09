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

// checks if the CreateAssetTransferV5Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssetTransferV5Resp{}

// CreateAssetTransferV5Resp struct for CreateAssetTransferV5Resp
type CreateAssetTransferV5Resp struct {
	Code *string `json:"code,omitempty"`
	Data []CreateAssetTransferV5RespDataInner `json:"data,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewCreateAssetTransferV5Resp instantiates a new CreateAssetTransferV5Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssetTransferV5Resp() *CreateAssetTransferV5Resp {
	this := CreateAssetTransferV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewCreateAssetTransferV5RespWithDefaults instantiates a new CreateAssetTransferV5Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssetTransferV5RespWithDefaults() *CreateAssetTransferV5Resp {
	this := CreateAssetTransferV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CreateAssetTransferV5Resp) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetTransferV5Resp) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CreateAssetTransferV5Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CreateAssetTransferV5Resp) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateAssetTransferV5Resp) GetData() []CreateAssetTransferV5RespDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []CreateAssetTransferV5RespDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetTransferV5Resp) GetDataOk() ([]CreateAssetTransferV5RespDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateAssetTransferV5Resp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CreateAssetTransferV5RespDataInner and assigns it to the Data field.
func (o *CreateAssetTransferV5Resp) SetData(v []CreateAssetTransferV5RespDataInner) {
	o.Data = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CreateAssetTransferV5Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetTransferV5Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CreateAssetTransferV5Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CreateAssetTransferV5Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o CreateAssetTransferV5Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAssetTransferV5Resp) ToMap() (map[string]interface{}, error) {
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

type NullableCreateAssetTransferV5Resp struct {
	value *CreateAssetTransferV5Resp
	isSet bool
}

func (v NullableCreateAssetTransferV5Resp) Get() *CreateAssetTransferV5Resp {
	return v.value
}

func (v *NullableCreateAssetTransferV5Resp) Set(val *CreateAssetTransferV5Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssetTransferV5Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssetTransferV5Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssetTransferV5Resp(val *CreateAssetTransferV5Resp) *NullableCreateAssetTransferV5Resp {
	return &NullableCreateAssetTransferV5Resp{value: val, isSet: true}
}

func (v NullableCreateAssetTransferV5Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAssetTransferV5Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


