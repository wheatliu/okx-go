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

// checks if the GetCopytradingProfitSharingDetailsV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCopytradingProfitSharingDetailsV5RespDataInner{}

// GetCopytradingProfitSharingDetailsV5RespDataInner struct for GetCopytradingProfitSharingDetailsV5RespDataInner
type GetCopytradingProfitSharingDetailsV5RespDataInner struct {
	// The currency of profit sharing.
	Ccy *string `json:"ccy,omitempty"`
	// Instrument type
	InstType *string `json:"instType,omitempty"`
	// Nickname of copy trader.
	NickName *string `json:"nickName,omitempty"`
	// Portrait link
	PortLink *string `json:"portLink,omitempty"`
	// Profit sharing amount. It would be 0 if there is no any profit sharing.
	ProfitSharingAmt *string `json:"profitSharingAmt,omitempty"`
	// Profit sharing ID.
	ProfitSharingId *string `json:"profitSharingId,omitempty"`
	// Profit sharing time.
	Ts *string `json:"ts,omitempty"`
}

// NewGetCopytradingProfitSharingDetailsV5RespDataInner instantiates a new GetCopytradingProfitSharingDetailsV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCopytradingProfitSharingDetailsV5RespDataInner() *GetCopytradingProfitSharingDetailsV5RespDataInner {
	this := GetCopytradingProfitSharingDetailsV5RespDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var instType string = ""
	this.InstType = &instType
	var nickName string = ""
	this.NickName = &nickName
	var portLink string = ""
	this.PortLink = &portLink
	var profitSharingAmt string = ""
	this.ProfitSharingAmt = &profitSharingAmt
	var profitSharingId string = ""
	this.ProfitSharingId = &profitSharingId
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetCopytradingProfitSharingDetailsV5RespDataInnerWithDefaults instantiates a new GetCopytradingProfitSharingDetailsV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCopytradingProfitSharingDetailsV5RespDataInnerWithDefaults() *GetCopytradingProfitSharingDetailsV5RespDataInner {
	this := GetCopytradingProfitSharingDetailsV5RespDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var instType string = ""
	this.InstType = &instType
	var nickName string = ""
	this.NickName = &nickName
	var portLink string = ""
	this.PortLink = &portLink
	var profitSharingAmt string = ""
	this.ProfitSharingAmt = &profitSharingAmt
	var profitSharingId string = ""
	this.ProfitSharingId = &profitSharingId
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) SetInstType(v string) {
	o.InstType = &v
}

// GetNickName returns the NickName field value if set, zero value otherwise.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetNickName() string {
	if o == nil || IsNil(o.NickName) {
		var ret string
		return ret
	}
	return *o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetNickNameOk() (*string, bool) {
	if o == nil || IsNil(o.NickName) {
		return nil, false
	}
	return o.NickName, true
}

// HasNickName returns a boolean if a field has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) HasNickName() bool {
	if o != nil && !IsNil(o.NickName) {
		return true
	}

	return false
}

// SetNickName gets a reference to the given string and assigns it to the NickName field.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) SetNickName(v string) {
	o.NickName = &v
}

// GetPortLink returns the PortLink field value if set, zero value otherwise.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetPortLink() string {
	if o == nil || IsNil(o.PortLink) {
		var ret string
		return ret
	}
	return *o.PortLink
}

// GetPortLinkOk returns a tuple with the PortLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetPortLinkOk() (*string, bool) {
	if o == nil || IsNil(o.PortLink) {
		return nil, false
	}
	return o.PortLink, true
}

// HasPortLink returns a boolean if a field has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) HasPortLink() bool {
	if o != nil && !IsNil(o.PortLink) {
		return true
	}

	return false
}

// SetPortLink gets a reference to the given string and assigns it to the PortLink field.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) SetPortLink(v string) {
	o.PortLink = &v
}

// GetProfitSharingAmt returns the ProfitSharingAmt field value if set, zero value otherwise.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetProfitSharingAmt() string {
	if o == nil || IsNil(o.ProfitSharingAmt) {
		var ret string
		return ret
	}
	return *o.ProfitSharingAmt
}

// GetProfitSharingAmtOk returns a tuple with the ProfitSharingAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetProfitSharingAmtOk() (*string, bool) {
	if o == nil || IsNil(o.ProfitSharingAmt) {
		return nil, false
	}
	return o.ProfitSharingAmt, true
}

// HasProfitSharingAmt returns a boolean if a field has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) HasProfitSharingAmt() bool {
	if o != nil && !IsNil(o.ProfitSharingAmt) {
		return true
	}

	return false
}

// SetProfitSharingAmt gets a reference to the given string and assigns it to the ProfitSharingAmt field.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) SetProfitSharingAmt(v string) {
	o.ProfitSharingAmt = &v
}

// GetProfitSharingId returns the ProfitSharingId field value if set, zero value otherwise.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetProfitSharingId() string {
	if o == nil || IsNil(o.ProfitSharingId) {
		var ret string
		return ret
	}
	return *o.ProfitSharingId
}

// GetProfitSharingIdOk returns a tuple with the ProfitSharingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetProfitSharingIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfitSharingId) {
		return nil, false
	}
	return o.ProfitSharingId, true
}

// HasProfitSharingId returns a boolean if a field has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) HasProfitSharingId() bool {
	if o != nil && !IsNil(o.ProfitSharingId) {
		return true
	}

	return false
}

// SetProfitSharingId gets a reference to the given string and assigns it to the ProfitSharingId field.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) SetProfitSharingId(v string) {
	o.ProfitSharingId = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetCopytradingProfitSharingDetailsV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetCopytradingProfitSharingDetailsV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCopytradingProfitSharingDetailsV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.NickName) {
		toSerialize["nickName"] = o.NickName
	}
	if !IsNil(o.PortLink) {
		toSerialize["portLink"] = o.PortLink
	}
	if !IsNil(o.ProfitSharingAmt) {
		toSerialize["profitSharingAmt"] = o.ProfitSharingAmt
	}
	if !IsNil(o.ProfitSharingId) {
		toSerialize["profitSharingId"] = o.ProfitSharingId
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetCopytradingProfitSharingDetailsV5RespDataInner struct {
	value *GetCopytradingProfitSharingDetailsV5RespDataInner
	isSet bool
}

func (v NullableGetCopytradingProfitSharingDetailsV5RespDataInner) Get() *GetCopytradingProfitSharingDetailsV5RespDataInner {
	return v.value
}

func (v *NullableGetCopytradingProfitSharingDetailsV5RespDataInner) Set(val *GetCopytradingProfitSharingDetailsV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCopytradingProfitSharingDetailsV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCopytradingProfitSharingDetailsV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCopytradingProfitSharingDetailsV5RespDataInner(val *GetCopytradingProfitSharingDetailsV5RespDataInner) *NullableGetCopytradingProfitSharingDetailsV5RespDataInner {
	return &NullableGetCopytradingProfitSharingDetailsV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetCopytradingProfitSharingDetailsV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCopytradingProfitSharingDetailsV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


