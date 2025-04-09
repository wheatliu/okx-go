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

// checks if the GetPublicOpenInterestV5Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPublicOpenInterestV5Resp{}

// GetPublicOpenInterestV5Resp struct for GetPublicOpenInterestV5Resp
type GetPublicOpenInterestV5Resp struct {
	Code *string `json:"code,omitempty"`
	Data []GetPublicOpenInterestV5RespDataInner `json:"data,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewGetPublicOpenInterestV5Resp instantiates a new GetPublicOpenInterestV5Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPublicOpenInterestV5Resp() *GetPublicOpenInterestV5Resp {
	this := GetPublicOpenInterestV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewGetPublicOpenInterestV5RespWithDefaults instantiates a new GetPublicOpenInterestV5Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPublicOpenInterestV5RespWithDefaults() *GetPublicOpenInterestV5Resp {
	this := GetPublicOpenInterestV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetPublicOpenInterestV5Resp) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicOpenInterestV5Resp) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetPublicOpenInterestV5Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetPublicOpenInterestV5Resp) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetPublicOpenInterestV5Resp) GetData() []GetPublicOpenInterestV5RespDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetPublicOpenInterestV5RespDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicOpenInterestV5Resp) GetDataOk() ([]GetPublicOpenInterestV5RespDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetPublicOpenInterestV5Resp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetPublicOpenInterestV5RespDataInner and assigns it to the Data field.
func (o *GetPublicOpenInterestV5Resp) SetData(v []GetPublicOpenInterestV5RespDataInner) {
	o.Data = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *GetPublicOpenInterestV5Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicOpenInterestV5Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *GetPublicOpenInterestV5Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *GetPublicOpenInterestV5Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o GetPublicOpenInterestV5Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPublicOpenInterestV5Resp) ToMap() (map[string]interface{}, error) {
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

type NullableGetPublicOpenInterestV5Resp struct {
	value *GetPublicOpenInterestV5Resp
	isSet bool
}

func (v NullableGetPublicOpenInterestV5Resp) Get() *GetPublicOpenInterestV5Resp {
	return v.value
}

func (v *NullableGetPublicOpenInterestV5Resp) Set(val *GetPublicOpenInterestV5Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPublicOpenInterestV5Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPublicOpenInterestV5Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPublicOpenInterestV5Resp(val *GetPublicOpenInterestV5Resp) *NullableGetPublicOpenInterestV5Resp {
	return &NullableGetPublicOpenInterestV5Resp{value: val, isSet: true}
}

func (v NullableGetPublicOpenInterestV5Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPublicOpenInterestV5Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


