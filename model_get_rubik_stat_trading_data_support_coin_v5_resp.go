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

// checks if the GetRubikStatTradingDataSupportCoinV5Resp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRubikStatTradingDataSupportCoinV5Resp{}

// GetRubikStatTradingDataSupportCoinV5Resp struct for GetRubikStatTradingDataSupportCoinV5Resp
type GetRubikStatTradingDataSupportCoinV5Resp struct {
	Code *string `json:"code,omitempty"`
	Data []GetRubikStatTradingDataSupportCoinV5RespDataInner `json:"data,omitempty"`
	Msg *string `json:"msg,omitempty"`
}

// NewGetRubikStatTradingDataSupportCoinV5Resp instantiates a new GetRubikStatTradingDataSupportCoinV5Resp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRubikStatTradingDataSupportCoinV5Resp() *GetRubikStatTradingDataSupportCoinV5Resp {
	this := GetRubikStatTradingDataSupportCoinV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// NewGetRubikStatTradingDataSupportCoinV5RespWithDefaults instantiates a new GetRubikStatTradingDataSupportCoinV5Resp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRubikStatTradingDataSupportCoinV5RespWithDefaults() *GetRubikStatTradingDataSupportCoinV5Resp {
	this := GetRubikStatTradingDataSupportCoinV5Resp{}
	var code string = ""
	this.Code = &code
	var msg string = ""
	this.Msg = &msg
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetRubikStatTradingDataSupportCoinV5Resp) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatTradingDataSupportCoinV5Resp) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetRubikStatTradingDataSupportCoinV5Resp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetRubikStatTradingDataSupportCoinV5Resp) SetCode(v string) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetRubikStatTradingDataSupportCoinV5Resp) GetData() []GetRubikStatTradingDataSupportCoinV5RespDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetRubikStatTradingDataSupportCoinV5RespDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatTradingDataSupportCoinV5Resp) GetDataOk() ([]GetRubikStatTradingDataSupportCoinV5RespDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetRubikStatTradingDataSupportCoinV5Resp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetRubikStatTradingDataSupportCoinV5RespDataInner and assigns it to the Data field.
func (o *GetRubikStatTradingDataSupportCoinV5Resp) SetData(v []GetRubikStatTradingDataSupportCoinV5RespDataInner) {
	o.Data = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *GetRubikStatTradingDataSupportCoinV5Resp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRubikStatTradingDataSupportCoinV5Resp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *GetRubikStatTradingDataSupportCoinV5Resp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *GetRubikStatTradingDataSupportCoinV5Resp) SetMsg(v string) {
	o.Msg = &v
}

func (o GetRubikStatTradingDataSupportCoinV5Resp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRubikStatTradingDataSupportCoinV5Resp) ToMap() (map[string]interface{}, error) {
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

type NullableGetRubikStatTradingDataSupportCoinV5Resp struct {
	value *GetRubikStatTradingDataSupportCoinV5Resp
	isSet bool
}

func (v NullableGetRubikStatTradingDataSupportCoinV5Resp) Get() *GetRubikStatTradingDataSupportCoinV5Resp {
	return v.value
}

func (v *NullableGetRubikStatTradingDataSupportCoinV5Resp) Set(val *GetRubikStatTradingDataSupportCoinV5Resp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRubikStatTradingDataSupportCoinV5Resp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRubikStatTradingDataSupportCoinV5Resp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRubikStatTradingDataSupportCoinV5Resp(val *GetRubikStatTradingDataSupportCoinV5Resp) *NullableGetRubikStatTradingDataSupportCoinV5Resp {
	return &NullableGetRubikStatTradingDataSupportCoinV5Resp{value: val, isSet: true}
}

func (v NullableGetRubikStatTradingDataSupportCoinV5Resp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRubikStatTradingDataSupportCoinV5Resp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


