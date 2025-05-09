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

// checks if the GetUsersSubaccountListV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUsersSubaccountListV5RespData{}

// GetUsersSubaccountListV5RespData struct for GetUsersSubaccountListV5RespData
type GetUsersSubaccountListV5RespData struct {
	// Whether the sub-account has the right to transfer out.   `true`: can transfer out   `false`: cannot transfer out
	CanTransOut *bool `json:"canTransOut,omitempty"`
	// Sub-account status  `true`: Normal  `false`: Frozen (global)
	Enable *bool `json:"enable,omitempty"`
	// The first level sub-account.   For subAcctLv: 1, firstLvSubAcct is equal to subAcct  For subAcctLv: 2, subAcct belongs to firstLvSubAcct.
	FirstLvSubAcct *string `json:"firstLvSubAcct,omitempty"`
	// Frozen functions  `trading`  `convert`  `transfer`  `withdrawal`  `deposit`  `flexible_loan`
	FrozenFunc []string `json:"frozenFunc,omitempty"`
	// If the sub-account switches on the Google Authenticator for login authentication.   `true`: On  `false`: Off
	GAuth *bool `json:"gAuth,omitempty"`
	// Whether it is dma broker sub-account.   `true`: Dma broker sub-account   `false`: It is not dma broker sub-account.
	IfDma *bool `json:"ifDma,omitempty"`
	// Sub-account note
	Label *string `json:"label,omitempty"`
	// Mobile number that linked with the sub-account.
	Mobile *string `json:"mobile,omitempty"`
	// Sub-account name
	SubAcct *string `json:"subAcct,omitempty"`
	// Sub-account level   `1`: First level sub-account  `2`: Second level sub-account.
	SubAcctLv *string `json:"subAcctLv,omitempty"`
	// Sub-account creation time, Unix timestamp in millisecond format. e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
	// Sub-account type   `1`: Standard sub-account   `2`: Managed trading sub-account   `5`: Custody trading sub-account - Copper  `9`: Managed trading sub-account - Copper   `12`: Custody trading sub-account - Komainu
	Type *string `json:"type,omitempty"`
	// Sub-account uid
	Uid *string `json:"uid,omitempty"`
}

// NewGetUsersSubaccountListV5RespData instantiates a new GetUsersSubaccountListV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUsersSubaccountListV5RespData() *GetUsersSubaccountListV5RespData {
	this := GetUsersSubaccountListV5RespData{}
	var firstLvSubAcct string = ""
	this.FirstLvSubAcct = &firstLvSubAcct
	var label string = ""
	this.Label = &label
	var mobile string = ""
	this.Mobile = &mobile
	var subAcct string = ""
	this.SubAcct = &subAcct
	var subAcctLv string = ""
	this.SubAcctLv = &subAcctLv
	var ts string = ""
	this.Ts = &ts
	var type_ string = ""
	this.Type = &type_
	var uid string = ""
	this.Uid = &uid
	return &this
}

// NewGetUsersSubaccountListV5RespDataWithDefaults instantiates a new GetUsersSubaccountListV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUsersSubaccountListV5RespDataWithDefaults() *GetUsersSubaccountListV5RespData {
	this := GetUsersSubaccountListV5RespData{}
	var firstLvSubAcct string = ""
	this.FirstLvSubAcct = &firstLvSubAcct
	var label string = ""
	this.Label = &label
	var mobile string = ""
	this.Mobile = &mobile
	var subAcct string = ""
	this.SubAcct = &subAcct
	var subAcctLv string = ""
	this.SubAcctLv = &subAcctLv
	var ts string = ""
	this.Ts = &ts
	var type_ string = ""
	this.Type = &type_
	var uid string = ""
	this.Uid = &uid
	return &this
}

// GetCanTransOut returns the CanTransOut field value if set, zero value otherwise.
func (o *GetUsersSubaccountListV5RespData) GetCanTransOut() bool {
	if o == nil || IsNil(o.CanTransOut) {
		var ret bool
		return ret
	}
	return *o.CanTransOut
}

// GetCanTransOutOk returns a tuple with the CanTransOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersSubaccountListV5RespData) GetCanTransOutOk() (*bool, bool) {
	if o == nil || IsNil(o.CanTransOut) {
		return nil, false
	}
	return o.CanTransOut, true
}

// HasCanTransOut returns a boolean if a field has been set.
func (o *GetUsersSubaccountListV5RespData) HasCanTransOut() bool {
	if o != nil && !IsNil(o.CanTransOut) {
		return true
	}

	return false
}

// SetCanTransOut gets a reference to the given bool and assigns it to the CanTransOut field.
func (o *GetUsersSubaccountListV5RespData) SetCanTransOut(v bool) {
	o.CanTransOut = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *GetUsersSubaccountListV5RespData) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersSubaccountListV5RespData) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *GetUsersSubaccountListV5RespData) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *GetUsersSubaccountListV5RespData) SetEnable(v bool) {
	o.Enable = &v
}

// GetFirstLvSubAcct returns the FirstLvSubAcct field value if set, zero value otherwise.
func (o *GetUsersSubaccountListV5RespData) GetFirstLvSubAcct() string {
	if o == nil || IsNil(o.FirstLvSubAcct) {
		var ret string
		return ret
	}
	return *o.FirstLvSubAcct
}

// GetFirstLvSubAcctOk returns a tuple with the FirstLvSubAcct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersSubaccountListV5RespData) GetFirstLvSubAcctOk() (*string, bool) {
	if o == nil || IsNil(o.FirstLvSubAcct) {
		return nil, false
	}
	return o.FirstLvSubAcct, true
}

// HasFirstLvSubAcct returns a boolean if a field has been set.
func (o *GetUsersSubaccountListV5RespData) HasFirstLvSubAcct() bool {
	if o != nil && !IsNil(o.FirstLvSubAcct) {
		return true
	}

	return false
}

// SetFirstLvSubAcct gets a reference to the given string and assigns it to the FirstLvSubAcct field.
func (o *GetUsersSubaccountListV5RespData) SetFirstLvSubAcct(v string) {
	o.FirstLvSubAcct = &v
}

// GetFrozenFunc returns the FrozenFunc field value if set, zero value otherwise.
func (o *GetUsersSubaccountListV5RespData) GetFrozenFunc() []string {
	if o == nil || IsNil(o.FrozenFunc) {
		var ret []string
		return ret
	}
	return o.FrozenFunc
}

// GetFrozenFuncOk returns a tuple with the FrozenFunc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersSubaccountListV5RespData) GetFrozenFuncOk() ([]string, bool) {
	if o == nil || IsNil(o.FrozenFunc) {
		return nil, false
	}
	return o.FrozenFunc, true
}

// HasFrozenFunc returns a boolean if a field has been set.
func (o *GetUsersSubaccountListV5RespData) HasFrozenFunc() bool {
	if o != nil && !IsNil(o.FrozenFunc) {
		return true
	}

	return false
}

// SetFrozenFunc gets a reference to the given []string and assigns it to the FrozenFunc field.
func (o *GetUsersSubaccountListV5RespData) SetFrozenFunc(v []string) {
	o.FrozenFunc = v
}

// GetGAuth returns the GAuth field value if set, zero value otherwise.
func (o *GetUsersSubaccountListV5RespData) GetGAuth() bool {
	if o == nil || IsNil(o.GAuth) {
		var ret bool
		return ret
	}
	return *o.GAuth
}

// GetGAuthOk returns a tuple with the GAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersSubaccountListV5RespData) GetGAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.GAuth) {
		return nil, false
	}
	return o.GAuth, true
}

// HasGAuth returns a boolean if a field has been set.
func (o *GetUsersSubaccountListV5RespData) HasGAuth() bool {
	if o != nil && !IsNil(o.GAuth) {
		return true
	}

	return false
}

// SetGAuth gets a reference to the given bool and assigns it to the GAuth field.
func (o *GetUsersSubaccountListV5RespData) SetGAuth(v bool) {
	o.GAuth = &v
}

// GetIfDma returns the IfDma field value if set, zero value otherwise.
func (o *GetUsersSubaccountListV5RespData) GetIfDma() bool {
	if o == nil || IsNil(o.IfDma) {
		var ret bool
		return ret
	}
	return *o.IfDma
}

// GetIfDmaOk returns a tuple with the IfDma field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersSubaccountListV5RespData) GetIfDmaOk() (*bool, bool) {
	if o == nil || IsNil(o.IfDma) {
		return nil, false
	}
	return o.IfDma, true
}

// HasIfDma returns a boolean if a field has been set.
func (o *GetUsersSubaccountListV5RespData) HasIfDma() bool {
	if o != nil && !IsNil(o.IfDma) {
		return true
	}

	return false
}

// SetIfDma gets a reference to the given bool and assigns it to the IfDma field.
func (o *GetUsersSubaccountListV5RespData) SetIfDma(v bool) {
	o.IfDma = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetUsersSubaccountListV5RespData) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersSubaccountListV5RespData) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetUsersSubaccountListV5RespData) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetUsersSubaccountListV5RespData) SetLabel(v string) {
	o.Label = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *GetUsersSubaccountListV5RespData) GetMobile() string {
	if o == nil || IsNil(o.Mobile) {
		var ret string
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersSubaccountListV5RespData) GetMobileOk() (*string, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *GetUsersSubaccountListV5RespData) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given string and assigns it to the Mobile field.
func (o *GetUsersSubaccountListV5RespData) SetMobile(v string) {
	o.Mobile = &v
}

// GetSubAcct returns the SubAcct field value if set, zero value otherwise.
func (o *GetUsersSubaccountListV5RespData) GetSubAcct() string {
	if o == nil || IsNil(o.SubAcct) {
		var ret string
		return ret
	}
	return *o.SubAcct
}

// GetSubAcctOk returns a tuple with the SubAcct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersSubaccountListV5RespData) GetSubAcctOk() (*string, bool) {
	if o == nil || IsNil(o.SubAcct) {
		return nil, false
	}
	return o.SubAcct, true
}

// HasSubAcct returns a boolean if a field has been set.
func (o *GetUsersSubaccountListV5RespData) HasSubAcct() bool {
	if o != nil && !IsNil(o.SubAcct) {
		return true
	}

	return false
}

// SetSubAcct gets a reference to the given string and assigns it to the SubAcct field.
func (o *GetUsersSubaccountListV5RespData) SetSubAcct(v string) {
	o.SubAcct = &v
}

// GetSubAcctLv returns the SubAcctLv field value if set, zero value otherwise.
func (o *GetUsersSubaccountListV5RespData) GetSubAcctLv() string {
	if o == nil || IsNil(o.SubAcctLv) {
		var ret string
		return ret
	}
	return *o.SubAcctLv
}

// GetSubAcctLvOk returns a tuple with the SubAcctLv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersSubaccountListV5RespData) GetSubAcctLvOk() (*string, bool) {
	if o == nil || IsNil(o.SubAcctLv) {
		return nil, false
	}
	return o.SubAcctLv, true
}

// HasSubAcctLv returns a boolean if a field has been set.
func (o *GetUsersSubaccountListV5RespData) HasSubAcctLv() bool {
	if o != nil && !IsNil(o.SubAcctLv) {
		return true
	}

	return false
}

// SetSubAcctLv gets a reference to the given string and assigns it to the SubAcctLv field.
func (o *GetUsersSubaccountListV5RespData) SetSubAcctLv(v string) {
	o.SubAcctLv = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetUsersSubaccountListV5RespData) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersSubaccountListV5RespData) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetUsersSubaccountListV5RespData) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetUsersSubaccountListV5RespData) SetTs(v string) {
	o.Ts = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetUsersSubaccountListV5RespData) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersSubaccountListV5RespData) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetUsersSubaccountListV5RespData) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetUsersSubaccountListV5RespData) SetType(v string) {
	o.Type = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *GetUsersSubaccountListV5RespData) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsersSubaccountListV5RespData) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *GetUsersSubaccountListV5RespData) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *GetUsersSubaccountListV5RespData) SetUid(v string) {
	o.Uid = &v
}

func (o GetUsersSubaccountListV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUsersSubaccountListV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanTransOut) {
		toSerialize["canTransOut"] = o.CanTransOut
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.FirstLvSubAcct) {
		toSerialize["firstLvSubAcct"] = o.FirstLvSubAcct
	}
	if !IsNil(o.FrozenFunc) {
		toSerialize["frozenFunc"] = o.FrozenFunc
	}
	if !IsNil(o.GAuth) {
		toSerialize["gAuth"] = o.GAuth
	}
	if !IsNil(o.IfDma) {
		toSerialize["ifDma"] = o.IfDma
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !IsNil(o.SubAcct) {
		toSerialize["subAcct"] = o.SubAcct
	}
	if !IsNil(o.SubAcctLv) {
		toSerialize["subAcctLv"] = o.SubAcctLv
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableGetUsersSubaccountListV5RespData struct {
	value *GetUsersSubaccountListV5RespData
	isSet bool
}

func (v NullableGetUsersSubaccountListV5RespData) Get() *GetUsersSubaccountListV5RespData {
	return v.value
}

func (v *NullableGetUsersSubaccountListV5RespData) Set(val *GetUsersSubaccountListV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsersSubaccountListV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsersSubaccountListV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsersSubaccountListV5RespData(val *GetUsersSubaccountListV5RespData) *NullableGetUsersSubaccountListV5RespData {
	return &NullableGetUsersSubaccountListV5RespData{value: val, isSet: true}
}

func (v NullableGetUsersSubaccountListV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsersSubaccountListV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


