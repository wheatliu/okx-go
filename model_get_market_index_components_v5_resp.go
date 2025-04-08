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

// checks if the GetMarketIndexComponentsV5Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarketIndexComponentsV5Resp{}

// GetMarketIndexComponentsV5Resp struct for GetMarketIndexComponentsV5Resp
type GetMarketIndexComponentsV5Resp struct {
	Code *string `json:"code,omitempty"`
	Data []GetMarketIndexComponentsV5RespDataInner `json:"data,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewGetMarketIndexComponentsV5Resp instantiates a new GetMarketIndexComponentsV5Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketIndexComponentsV5Resp() *GetMarketIndexComponentsV5Resp {
	this := GetMarketIndexComponentsV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewGetMarketIndexComponentsV5RespWithDefaults instantiates a new GetMarketIndexComponentsV5Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketIndexComponentsV5RespWithDefaults() *GetMarketIndexComponentsV5Resp {
	this := GetMarketIndexComponentsV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetMarketIndexComponentsV5Resp) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketIndexComponentsV5Resp) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetMarketIndexComponentsV5Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetMarketIndexComponentsV5Resp) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetMarketIndexComponentsV5Resp) GetData() []GetMarketIndexComponentsV5RespDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetMarketIndexComponentsV5RespDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketIndexComponentsV5Resp) GetDataOk() ([]GetMarketIndexComponentsV5RespDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetMarketIndexComponentsV5Resp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetMarketIndexComponentsV5RespDataInner and assigns it to the Data field.
func (o *GetMarketIndexComponentsV5Resp) SetData(v []GetMarketIndexComponentsV5RespDataInner) {
	o.Data = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *GetMarketIndexComponentsV5Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketIndexComponentsV5Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *GetMarketIndexComponentsV5Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *GetMarketIndexComponentsV5Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o GetMarketIndexComponentsV5Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketIndexComponentsV5Resp) ToMap() (map[string]interface{}, error) {
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

type NullableGetMarketIndexComponentsV5Resp struct {
	value *GetMarketIndexComponentsV5Resp
	isSet bool
}

func (v NullableGetMarketIndexComponentsV5Resp) Get() *GetMarketIndexComponentsV5Resp {
	return v.value
}

func (v *NullableGetMarketIndexComponentsV5Resp) Set(val *GetMarketIndexComponentsV5Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketIndexComponentsV5Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketIndexComponentsV5Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketIndexComponentsV5Resp(val *GetMarketIndexComponentsV5Resp) *NullableGetMarketIndexComponentsV5Resp {
	return &NullableGetMarketIndexComponentsV5Resp{value: val, isSet: true}
}

func (v NullableGetMarketIndexComponentsV5Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketIndexComponentsV5Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


