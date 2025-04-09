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

// checks if the GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner{}

// GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner struct for GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner
type GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner struct {
	// Begin copying time. Unix timestamp format in milliseconds, e.g.1597026383085
	BeginCopyTime *string `json:"beginCopyTime,omitempty"`
	// Nick name
	NickName *string `json:"nickName,omitempty"`
	// Copy trading profit and loss
	Pnl *string `json:"pnl,omitempty"`
	// Copy trader portrait link
	PortLink *string `json:"portLink,omitempty"`
}

// NewGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner instantiates a new GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner() *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner {
	this := GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner{}
	var beginCopyTime string = ""
	this.BeginCopyTime = &beginCopyTime
	var nickName string = ""
	this.NickName = &nickName
	var pnl string = ""
	this.Pnl = &pnl
	var portLink string = ""
	this.PortLink = &portLink
	return &this
}

// NewGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInnerWithDefaults instantiates a new GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInnerWithDefaults() *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner {
	this := GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner{}
	var beginCopyTime string = ""
	this.BeginCopyTime = &beginCopyTime
	var nickName string = ""
	this.NickName = &nickName
	var pnl string = ""
	this.Pnl = &pnl
	var portLink string = ""
	this.PortLink = &portLink
	return &this
}

// GetBeginCopyTime returns the BeginCopyTime field value if set, zero value otherwise.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) GetBeginCopyTime() string {
	if o == nil || IsNil(o.BeginCopyTime) {
		var ret string
		return ret
	}
	return *o.BeginCopyTime
}

// GetBeginCopyTimeOk returns a tuple with the BeginCopyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) GetBeginCopyTimeOk() (*string, bool) {
	if o == nil || IsNil(o.BeginCopyTime) {
		return nil, false
	}
	return o.BeginCopyTime, true
}

// HasBeginCopyTime returns a boolean if a field has been set.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) HasBeginCopyTime() bool {
	if o != nil && !IsNil(o.BeginCopyTime) {
		return true
	}

	return false
}

// SetBeginCopyTime gets a reference to the given string and assigns it to the BeginCopyTime field.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) SetBeginCopyTime(v string) {
	o.BeginCopyTime = &v
}

// GetNickName returns the NickName field value if set, zero value otherwise.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) GetNickName() string {
	if o == nil || IsNil(o.NickName) {
		var ret string
		return ret
	}
	return *o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) GetNickNameOk() (*string, bool) {
	if o == nil || IsNil(o.NickName) {
		return nil, false
	}
	return o.NickName, true
}

// HasNickName returns a boolean if a field has been set.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) HasNickName() bool {
	if o != nil && !IsNil(o.NickName) {
		return true
	}

	return false
}

// SetNickName gets a reference to the given string and assigns it to the NickName field.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) SetNickName(v string) {
	o.NickName = &v
}

// GetPnl returns the Pnl field value if set, zero value otherwise.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) GetPnl() string {
	if o == nil || IsNil(o.Pnl) {
		var ret string
		return ret
	}
	return *o.Pnl
}

// GetPnlOk returns a tuple with the Pnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) GetPnlOk() (*string, bool) {
	if o == nil || IsNil(o.Pnl) {
		return nil, false
	}
	return o.Pnl, true
}

// HasPnl returns a boolean if a field has been set.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) HasPnl() bool {
	if o != nil && !IsNil(o.Pnl) {
		return true
	}

	return false
}

// SetPnl gets a reference to the given string and assigns it to the Pnl field.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) SetPnl(v string) {
	o.Pnl = &v
}

// GetPortLink returns the PortLink field value if set, zero value otherwise.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) GetPortLink() string {
	if o == nil || IsNil(o.PortLink) {
		var ret string
		return ret
	}
	return *o.PortLink
}

// GetPortLinkOk returns a tuple with the PortLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) GetPortLinkOk() (*string, bool) {
	if o == nil || IsNil(o.PortLink) {
		return nil, false
	}
	return o.PortLink, true
}

// HasPortLink returns a boolean if a field has been set.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) HasPortLink() bool {
	if o != nil && !IsNil(o.PortLink) {
		return true
	}

	return false
}

// SetPortLink gets a reference to the given string and assigns it to the PortLink field.
func (o *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) SetPortLink(v string) {
	o.PortLink = &v
}

func (o GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BeginCopyTime) {
		toSerialize["beginCopyTime"] = o.BeginCopyTime
	}
	if !IsNil(o.NickName) {
		toSerialize["nickName"] = o.NickName
	}
	if !IsNil(o.Pnl) {
		toSerialize["pnl"] = o.Pnl
	}
	if !IsNil(o.PortLink) {
		toSerialize["portLink"] = o.PortLink
	}
	return toSerialize, nil
}

type NullableGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner struct {
	value *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner
	isSet bool
}

func (v NullableGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) Get() *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner {
	return v.value
}

func (v *NullableGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) Set(val *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner(val *GetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) *NullableGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner {
	return &NullableGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner{value: val, isSet: true}
}

func (v NullableGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCopytradingPublicCopyTradersV5RespDataInnerCopyTradersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


