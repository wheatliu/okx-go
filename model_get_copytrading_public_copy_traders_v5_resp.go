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

// checks if the GetCopytradingPublicCopyTradersV5Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCopytradingPublicCopyTradersV5Resp{}

// GetCopytradingPublicCopyTradersV5Resp struct for GetCopytradingPublicCopyTradersV5Resp
type GetCopytradingPublicCopyTradersV5Resp struct {
	Code *string `json:"code,omitempty"`
	Data *GetCopytradingPublicCopyTradersV5RespData `json:"data,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewGetCopytradingPublicCopyTradersV5Resp instantiates a new GetCopytradingPublicCopyTradersV5Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCopytradingPublicCopyTradersV5Resp() *GetCopytradingPublicCopyTradersV5Resp {
	this := GetCopytradingPublicCopyTradersV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewGetCopytradingPublicCopyTradersV5RespWithDefaults instantiates a new GetCopytradingPublicCopyTradersV5Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCopytradingPublicCopyTradersV5RespWithDefaults() *GetCopytradingPublicCopyTradersV5Resp {
	this := GetCopytradingPublicCopyTradersV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetCopytradingPublicCopyTradersV5Resp) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicCopyTradersV5Resp) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetCopytradingPublicCopyTradersV5Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetCopytradingPublicCopyTradersV5Resp) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetCopytradingPublicCopyTradersV5Resp) GetData() GetCopytradingPublicCopyTradersV5RespData {
	if o == nil || IsNil(o.Data) {
		var ret GetCopytradingPublicCopyTradersV5RespData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicCopyTradersV5Resp) GetDataOk() (*GetCopytradingPublicCopyTradersV5RespData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetCopytradingPublicCopyTradersV5Resp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetCopytradingPublicCopyTradersV5RespData and assigns it to the Data field.
func (o *GetCopytradingPublicCopyTradersV5Resp) SetData(v GetCopytradingPublicCopyTradersV5RespData) {
	o.Data = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *GetCopytradingPublicCopyTradersV5Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicCopyTradersV5Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *GetCopytradingPublicCopyTradersV5Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *GetCopytradingPublicCopyTradersV5Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o GetCopytradingPublicCopyTradersV5Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCopytradingPublicCopyTradersV5Resp) ToMap() (map[string]interface{}, error) {
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

type NullableGetCopytradingPublicCopyTradersV5Resp struct {
	value *GetCopytradingPublicCopyTradersV5Resp
	isSet bool
}

func (v NullableGetCopytradingPublicCopyTradersV5Resp) Get() *GetCopytradingPublicCopyTradersV5Resp {
	return v.value
}

func (v *NullableGetCopytradingPublicCopyTradersV5Resp) Set(val *GetCopytradingPublicCopyTradersV5Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCopytradingPublicCopyTradersV5Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCopytradingPublicCopyTradersV5Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCopytradingPublicCopyTradersV5Resp(val *GetCopytradingPublicCopyTradersV5Resp) *NullableGetCopytradingPublicCopyTradersV5Resp {
	return &NullableGetCopytradingPublicCopyTradersV5Resp{value: val, isSet: true}
}

func (v NullableGetCopytradingPublicCopyTradersV5Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCopytradingPublicCopyTradersV5Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


