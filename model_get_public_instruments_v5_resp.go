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

// checks if the GetPublicInstrumentsV5Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPublicInstrumentsV5Resp{}

// GetPublicInstrumentsV5Resp struct for GetPublicInstrumentsV5Resp
type GetPublicInstrumentsV5Resp struct {
	Code *string `json:"code,omitempty"`
	Data []GetPublicInstrumentsV5RespDataInner `json:"data,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewGetPublicInstrumentsV5Resp instantiates a new GetPublicInstrumentsV5Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPublicInstrumentsV5Resp() *GetPublicInstrumentsV5Resp {
	this := GetPublicInstrumentsV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewGetPublicInstrumentsV5RespWithDefaults instantiates a new GetPublicInstrumentsV5Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPublicInstrumentsV5RespWithDefaults() *GetPublicInstrumentsV5Resp {
	this := GetPublicInstrumentsV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetPublicInstrumentsV5Resp) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicInstrumentsV5Resp) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetPublicInstrumentsV5Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetPublicInstrumentsV5Resp) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetPublicInstrumentsV5Resp) GetData() []GetPublicInstrumentsV5RespDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetPublicInstrumentsV5RespDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicInstrumentsV5Resp) GetDataOk() ([]GetPublicInstrumentsV5RespDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetPublicInstrumentsV5Resp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetPublicInstrumentsV5RespDataInner and assigns it to the Data field.
func (o *GetPublicInstrumentsV5Resp) SetData(v []GetPublicInstrumentsV5RespDataInner) {
	o.Data = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *GetPublicInstrumentsV5Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicInstrumentsV5Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *GetPublicInstrumentsV5Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *GetPublicInstrumentsV5Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o GetPublicInstrumentsV5Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPublicInstrumentsV5Resp) ToMap() (map[string]interface{}, error) {
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

type NullableGetPublicInstrumentsV5Resp struct {
	value *GetPublicInstrumentsV5Resp
	isSet bool
}

func (v NullableGetPublicInstrumentsV5Resp) Get() *GetPublicInstrumentsV5Resp {
	return v.value
}

func (v *NullableGetPublicInstrumentsV5Resp) Set(val *GetPublicInstrumentsV5Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPublicInstrumentsV5Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPublicInstrumentsV5Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPublicInstrumentsV5Resp(val *GetPublicInstrumentsV5Resp) *NullableGetPublicInstrumentsV5Resp {
	return &NullableGetPublicInstrumentsV5Resp{value: val, isSet: true}
}

func (v NullableGetPublicInstrumentsV5Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPublicInstrumentsV5Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


