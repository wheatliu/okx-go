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

// checks if the GetMarketTradesV5Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarketTradesV5Resp{}

// GetMarketTradesV5Resp struct for GetMarketTradesV5Resp
type GetMarketTradesV5Resp struct {
	Code *string `json:"code,omitempty"`
	Data []GetMarketHistoryTradesV5RespDataInner `json:"data,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewGetMarketTradesV5Resp instantiates a new GetMarketTradesV5Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketTradesV5Resp() *GetMarketTradesV5Resp {
	this := GetMarketTradesV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewGetMarketTradesV5RespWithDefaults instantiates a new GetMarketTradesV5Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketTradesV5RespWithDefaults() *GetMarketTradesV5Resp {
	this := GetMarketTradesV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetMarketTradesV5Resp) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketTradesV5Resp) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetMarketTradesV5Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetMarketTradesV5Resp) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetMarketTradesV5Resp) GetData() []GetMarketHistoryTradesV5RespDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetMarketHistoryTradesV5RespDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketTradesV5Resp) GetDataOk() ([]GetMarketHistoryTradesV5RespDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetMarketTradesV5Resp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetMarketHistoryTradesV5RespDataInner and assigns it to the Data field.
func (o *GetMarketTradesV5Resp) SetData(v []GetMarketHistoryTradesV5RespDataInner) {
	o.Data = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *GetMarketTradesV5Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMarketTradesV5Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *GetMarketTradesV5Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *GetMarketTradesV5Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o GetMarketTradesV5Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketTradesV5Resp) ToMap() (map[string]interface{}, error) {
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

type NullableGetMarketTradesV5Resp struct {
	value *GetMarketTradesV5Resp
	isSet bool
}

func (v NullableGetMarketTradesV5Resp) Get() *GetMarketTradesV5Resp {
	return v.value
}

func (v *NullableGetMarketTradesV5Resp) Set(val *GetMarketTradesV5Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketTradesV5Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketTradesV5Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketTradesV5Resp(val *GetMarketTradesV5Resp) *NullableGetMarketTradesV5Resp {
	return &NullableGetMarketTradesV5Resp{value: val, isSet: true}
}

func (v NullableGetMarketTradesV5Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketTradesV5Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


