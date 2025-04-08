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

// checks if the GetAssetDepositAddressV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAssetDepositAddressV5RespDataInner{}

// GetAssetDepositAddressV5RespDataInner struct for GetAssetDepositAddressV5RespDataInner
type GetAssetDepositAddressV5RespDataInner struct {
	// Deposit address
	Addr *string `json:"addr,omitempty"`
	// Deposit address attachment (This will not be returned if the currency does not require this)  e.g. `TONCOIN` attached tag name is `comment`, the return will be `{'comment':'123456'}`
	AddrEx map[string]interface{} `json:"addrEx,omitempty"`
	// Currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Chain name, e.g. `USDT-ERC20`, `USDT-TRC20`
	Chain *string `json:"chain,omitempty"`
	// Last 6 digits of contract address
	CtAddr *string `json:"ctAddr,omitempty"`
	// Deposit memo (This will not be returned if the currency does not require a memo for deposit)
	Memo *string `json:"memo,omitempty"`
	// Deposit payment ID (This will not be returned if the currency does not require a payment_id for deposit)
	PmtId *string `json:"pmtId,omitempty"`
	// Return `true` if the current deposit address is selected by the website page
	Selected *bool `json:"selected,omitempty"`
	// Deposit tag (This will not be returned if the currency does not require a tag for deposit)
	Tag *string `json:"tag,omitempty"`
	// The beneficiary account  `6`: Funding account `18`: Trading account  The users under some entity (e.g. Brazil) only support deposit to trading account.
	To *string `json:"to,omitempty"`
	// Verified name (for recipient)
	VerifiedName *string `json:"verifiedName,omitempty"`
}

// NewGetAssetDepositAddressV5RespDataInner instantiates a new GetAssetDepositAddressV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssetDepositAddressV5RespDataInner() *GetAssetDepositAddressV5RespDataInner {
	this := GetAssetDepositAddressV5RespDataInner{}
	var addr string = ""
	this.Addr = &addr
	var ccy string = ""
	this.Ccy = &ccy
	var chain string = ""
	this.Chain = &chain
	var ctAddr string = ""
	this.CtAddr = &ctAddr
	var memo string = ""
	this.Memo = &memo
	var pmtId string = ""
	this.PmtId = &pmtId
	var tag string = ""
	this.Tag = &tag
	var to string = ""
	this.To = &to
	var verifiedName string = ""
	this.VerifiedName = &verifiedName
	return &this
}

// NewGetAssetDepositAddressV5RespDataInnerWithDefaults instantiates a new GetAssetDepositAddressV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssetDepositAddressV5RespDataInnerWithDefaults() *GetAssetDepositAddressV5RespDataInner {
	this := GetAssetDepositAddressV5RespDataInner{}
	var addr string = ""
	this.Addr = &addr
	var ccy string = ""
	this.Ccy = &ccy
	var chain string = ""
	this.Chain = &chain
	var ctAddr string = ""
	this.CtAddr = &ctAddr
	var memo string = ""
	this.Memo = &memo
	var pmtId string = ""
	this.PmtId = &pmtId
	var tag string = ""
	this.Tag = &tag
	var to string = ""
	this.To = &to
	var verifiedName string = ""
	this.VerifiedName = &verifiedName
	return &this
}

// GetAddr returns the Addr field value if set, zero value otherwise.
func (o *GetAssetDepositAddressV5RespDataInner) GetAddr() string {
	if o == nil || IsNil(o.Addr) {
		var ret string
		return ret
	}
	return *o.Addr
}

// GetAddrOk returns a tuple with the Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositAddressV5RespDataInner) GetAddrOk() (*string, bool) {
	if o == nil || IsNil(o.Addr) {
		return nil, false
	}
	return o.Addr, true
}

// HasAddr returns a boolean if a field has been set.
func (o *GetAssetDepositAddressV5RespDataInner) HasAddr() bool {
	if o != nil && !IsNil(o.Addr) {
		return true
	}

	return false
}

// SetAddr gets a reference to the given string and assigns it to the Addr field.
func (o *GetAssetDepositAddressV5RespDataInner) SetAddr(v string) {
	o.Addr = &v
}

// GetAddrEx returns the AddrEx field value if set, zero value otherwise.
func (o *GetAssetDepositAddressV5RespDataInner) GetAddrEx() map[string]interface{} {
	if o == nil || IsNil(o.AddrEx) {
		var ret map[string]interface{}
		return ret
	}
	return o.AddrEx
}

// GetAddrExOk returns a tuple with the AddrEx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositAddressV5RespDataInner) GetAddrExOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AddrEx) {
		return map[string]interface{}{}, false
	}
	return o.AddrEx, true
}

// HasAddrEx returns a boolean if a field has been set.
func (o *GetAssetDepositAddressV5RespDataInner) HasAddrEx() bool {
	if o != nil && !IsNil(o.AddrEx) {
		return true
	}

	return false
}

// SetAddrEx gets a reference to the given map[string]interface{} and assigns it to the AddrEx field.
func (o *GetAssetDepositAddressV5RespDataInner) SetAddrEx(v map[string]interface{}) {
	o.AddrEx = v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetAssetDepositAddressV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositAddressV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetAssetDepositAddressV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetAssetDepositAddressV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *GetAssetDepositAddressV5RespDataInner) GetChain() string {
	if o == nil || IsNil(o.Chain) {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositAddressV5RespDataInner) GetChainOk() (*string, bool) {
	if o == nil || IsNil(o.Chain) {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *GetAssetDepositAddressV5RespDataInner) HasChain() bool {
	if o != nil && !IsNil(o.Chain) {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *GetAssetDepositAddressV5RespDataInner) SetChain(v string) {
	o.Chain = &v
}

// GetCtAddr returns the CtAddr field value if set, zero value otherwise.
func (o *GetAssetDepositAddressV5RespDataInner) GetCtAddr() string {
	if o == nil || IsNil(o.CtAddr) {
		var ret string
		return ret
	}
	return *o.CtAddr
}

// GetCtAddrOk returns a tuple with the CtAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositAddressV5RespDataInner) GetCtAddrOk() (*string, bool) {
	if o == nil || IsNil(o.CtAddr) {
		return nil, false
	}
	return o.CtAddr, true
}

// HasCtAddr returns a boolean if a field has been set.
func (o *GetAssetDepositAddressV5RespDataInner) HasCtAddr() bool {
	if o != nil && !IsNil(o.CtAddr) {
		return true
	}

	return false
}

// SetCtAddr gets a reference to the given string and assigns it to the CtAddr field.
func (o *GetAssetDepositAddressV5RespDataInner) SetCtAddr(v string) {
	o.CtAddr = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *GetAssetDepositAddressV5RespDataInner) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositAddressV5RespDataInner) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *GetAssetDepositAddressV5RespDataInner) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *GetAssetDepositAddressV5RespDataInner) SetMemo(v string) {
	o.Memo = &v
}

// GetPmtId returns the PmtId field value if set, zero value otherwise.
func (o *GetAssetDepositAddressV5RespDataInner) GetPmtId() string {
	if o == nil || IsNil(o.PmtId) {
		var ret string
		return ret
	}
	return *o.PmtId
}

// GetPmtIdOk returns a tuple with the PmtId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositAddressV5RespDataInner) GetPmtIdOk() (*string, bool) {
	if o == nil || IsNil(o.PmtId) {
		return nil, false
	}
	return o.PmtId, true
}

// HasPmtId returns a boolean if a field has been set.
func (o *GetAssetDepositAddressV5RespDataInner) HasPmtId() bool {
	if o != nil && !IsNil(o.PmtId) {
		return true
	}

	return false
}

// SetPmtId gets a reference to the given string and assigns it to the PmtId field.
func (o *GetAssetDepositAddressV5RespDataInner) SetPmtId(v string) {
	o.PmtId = &v
}

// GetSelected returns the Selected field value if set, zero value otherwise.
func (o *GetAssetDepositAddressV5RespDataInner) GetSelected() bool {
	if o == nil || IsNil(o.Selected) {
		var ret bool
		return ret
	}
	return *o.Selected
}

// GetSelectedOk returns a tuple with the Selected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositAddressV5RespDataInner) GetSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Selected) {
		return nil, false
	}
	return o.Selected, true
}

// HasSelected returns a boolean if a field has been set.
func (o *GetAssetDepositAddressV5RespDataInner) HasSelected() bool {
	if o != nil && !IsNil(o.Selected) {
		return true
	}

	return false
}

// SetSelected gets a reference to the given bool and assigns it to the Selected field.
func (o *GetAssetDepositAddressV5RespDataInner) SetSelected(v bool) {
	o.Selected = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetAssetDepositAddressV5RespDataInner) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositAddressV5RespDataInner) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetAssetDepositAddressV5RespDataInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetAssetDepositAddressV5RespDataInner) SetTag(v string) {
	o.Tag = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *GetAssetDepositAddressV5RespDataInner) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositAddressV5RespDataInner) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *GetAssetDepositAddressV5RespDataInner) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *GetAssetDepositAddressV5RespDataInner) SetTo(v string) {
	o.To = &v
}

// GetVerifiedName returns the VerifiedName field value if set, zero value otherwise.
func (o *GetAssetDepositAddressV5RespDataInner) GetVerifiedName() string {
	if o == nil || IsNil(o.VerifiedName) {
		var ret string
		return ret
	}
	return *o.VerifiedName
}

// GetVerifiedNameOk returns a tuple with the VerifiedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositAddressV5RespDataInner) GetVerifiedNameOk() (*string, bool) {
	if o == nil || IsNil(o.VerifiedName) {
		return nil, false
	}
	return o.VerifiedName, true
}

// HasVerifiedName returns a boolean if a field has been set.
func (o *GetAssetDepositAddressV5RespDataInner) HasVerifiedName() bool {
	if o != nil && !IsNil(o.VerifiedName) {
		return true
	}

	return false
}

// SetVerifiedName gets a reference to the given string and assigns it to the VerifiedName field.
func (o *GetAssetDepositAddressV5RespDataInner) SetVerifiedName(v string) {
	o.VerifiedName = &v
}

func (o GetAssetDepositAddressV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssetDepositAddressV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addr) {
		toSerialize["addr"] = o.Addr
	}
	if !IsNil(o.AddrEx) {
		toSerialize["addrEx"] = o.AddrEx
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.Chain) {
		toSerialize["chain"] = o.Chain
	}
	if !IsNil(o.CtAddr) {
		toSerialize["ctAddr"] = o.CtAddr
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.PmtId) {
		toSerialize["pmtId"] = o.PmtId
	}
	if !IsNil(o.Selected) {
		toSerialize["selected"] = o.Selected
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.VerifiedName) {
		toSerialize["verifiedName"] = o.VerifiedName
	}
	return toSerialize, nil
}

type NullableGetAssetDepositAddressV5RespDataInner struct {
	value *GetAssetDepositAddressV5RespDataInner
	isSet bool
}

func (v NullableGetAssetDepositAddressV5RespDataInner) Get() *GetAssetDepositAddressV5RespDataInner {
	return v.value
}

func (v *NullableGetAssetDepositAddressV5RespDataInner) Set(val *GetAssetDepositAddressV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetDepositAddressV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetDepositAddressV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetDepositAddressV5RespDataInner(val *GetAssetDepositAddressV5RespDataInner) *NullableGetAssetDepositAddressV5RespDataInner {
	return &NullableGetAssetDepositAddressV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetAssetDepositAddressV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetDepositAddressV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


