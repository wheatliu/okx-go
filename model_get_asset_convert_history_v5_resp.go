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

// checks if the GetAssetConvertHistoryV5Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAssetConvertHistoryV5Resp{}

// GetAssetConvertHistoryV5Resp struct for GetAssetConvertHistoryV5Resp
type GetAssetConvertHistoryV5Resp struct {
	Code *string `json:"code,omitempty"`
	Data *GetAssetConvertHistoryV5RespData `json:"data,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewGetAssetConvertHistoryV5Resp instantiates a new GetAssetConvertHistoryV5Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssetConvertHistoryV5Resp() *GetAssetConvertHistoryV5Resp {
	this := GetAssetConvertHistoryV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewGetAssetConvertHistoryV5RespWithDefaults instantiates a new GetAssetConvertHistoryV5Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssetConvertHistoryV5RespWithDefaults() *GetAssetConvertHistoryV5Resp {
	this := GetAssetConvertHistoryV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5Resp) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5Resp) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetAssetConvertHistoryV5Resp) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5Resp) GetData() GetAssetConvertHistoryV5RespData {
	if o == nil || IsNil(o.Data) {
		var ret GetAssetConvertHistoryV5RespData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5Resp) GetDataOk() (*GetAssetConvertHistoryV5RespData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5Resp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetAssetConvertHistoryV5RespData and assigns it to the Data field.
func (o *GetAssetConvertHistoryV5Resp) SetData(v GetAssetConvertHistoryV5RespData) {
	o.Data = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *GetAssetConvertHistoryV5Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetConvertHistoryV5Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *GetAssetConvertHistoryV5Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *GetAssetConvertHistoryV5Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o GetAssetConvertHistoryV5Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssetConvertHistoryV5Resp) ToMap() (map[string]interface{}, error) {
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

type NullableGetAssetConvertHistoryV5Resp struct {
	value *GetAssetConvertHistoryV5Resp
	isSet bool
}

func (v NullableGetAssetConvertHistoryV5Resp) Get() *GetAssetConvertHistoryV5Resp {
	return v.value
}

func (v *NullableGetAssetConvertHistoryV5Resp) Set(val *GetAssetConvertHistoryV5Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetConvertHistoryV5Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetConvertHistoryV5Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetConvertHistoryV5Resp(val *GetAssetConvertHistoryV5Resp) *NullableGetAssetConvertHistoryV5Resp {
	return &NullableGetAssetConvertHistoryV5Resp{value: val, isSet: true}
}

func (v NullableGetAssetConvertHistoryV5Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetConvertHistoryV5Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


