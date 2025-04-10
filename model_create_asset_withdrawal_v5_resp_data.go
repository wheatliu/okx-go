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

// checks if the CreateAssetWithdrawalV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssetWithdrawalV5RespData{}

// CreateAssetWithdrawalV5RespData struct for CreateAssetWithdrawalV5RespData
type CreateAssetWithdrawalV5RespData struct {
	// Withdrawal amount
	Amt *string `json:"amt,omitempty"`
	// Currency
	Ccy *string `json:"ccy,omitempty"`
	// Chain name, e.g. `USDT-ERC20`, `USDT-TRC20`
	Chain *string `json:"chain,omitempty"`
	// Client-supplied ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.
	ClientId *string `json:"clientId,omitempty"`
	// Withdrawal ID
	WdId *string `json:"wdId,omitempty"`
}

// NewCreateAssetWithdrawalV5RespData instantiates a new CreateAssetWithdrawalV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssetWithdrawalV5RespData() *CreateAssetWithdrawalV5RespData {
	this := CreateAssetWithdrawalV5RespData{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var chain string = ""
	this.Chain = &chain
	var clientId string = ""
	this.ClientId = &clientId
	var wdId string = ""
	this.WdId = &wdId
	return &this
}

// NewCreateAssetWithdrawalV5RespDataWithDefaults instantiates a new CreateAssetWithdrawalV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssetWithdrawalV5RespDataWithDefaults() *CreateAssetWithdrawalV5RespData {
	this := CreateAssetWithdrawalV5RespData{}
	var amt string = ""
	this.Amt = &amt
	var ccy string = ""
	this.Ccy = &ccy
	var chain string = ""
	this.Chain = &chain
	var clientId string = ""
	this.ClientId = &clientId
	var wdId string = ""
	this.WdId = &wdId
	return &this
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *CreateAssetWithdrawalV5RespData) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetWithdrawalV5RespData) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *CreateAssetWithdrawalV5RespData) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *CreateAssetWithdrawalV5RespData) SetAmt(v string) {
	o.Amt = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *CreateAssetWithdrawalV5RespData) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetWithdrawalV5RespData) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *CreateAssetWithdrawalV5RespData) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *CreateAssetWithdrawalV5RespData) SetCcy(v string) {
	o.Ccy = &v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *CreateAssetWithdrawalV5RespData) GetChain() string {
	if o == nil || IsNil(o.Chain) {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetWithdrawalV5RespData) GetChainOk() (*string, bool) {
	if o == nil || IsNil(o.Chain) {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *CreateAssetWithdrawalV5RespData) HasChain() bool {
	if o != nil && !IsNil(o.Chain) {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *CreateAssetWithdrawalV5RespData) SetChain(v string) {
	o.Chain = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreateAssetWithdrawalV5RespData) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetWithdrawalV5RespData) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreateAssetWithdrawalV5RespData) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreateAssetWithdrawalV5RespData) SetClientId(v string) {
	o.ClientId = &v
}

// GetWdId returns the WdId field value if set, zero value otherwise.
func (o *CreateAssetWithdrawalV5RespData) GetWdId() string {
	if o == nil || IsNil(o.WdId) {
		var ret string
		return ret
	}
	return *o.WdId
}

// GetWdIdOk returns a tuple with the WdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetWithdrawalV5RespData) GetWdIdOk() (*string, bool) {
	if o == nil || IsNil(o.WdId) {
		return nil, false
	}
	return o.WdId, true
}

// HasWdId returns a boolean if a field has been set.
func (o *CreateAssetWithdrawalV5RespData) HasWdId() bool {
	if o != nil && !IsNil(o.WdId) {
		return true
	}

	return false
}

// SetWdId gets a reference to the given string and assigns it to the WdId field.
func (o *CreateAssetWithdrawalV5RespData) SetWdId(v string) {
	o.WdId = &v
}

func (o CreateAssetWithdrawalV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAssetWithdrawalV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.Chain) {
		toSerialize["chain"] = o.Chain
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.WdId) {
		toSerialize["wdId"] = o.WdId
	}
	return toSerialize, nil
}

type NullableCreateAssetWithdrawalV5RespData struct {
	value *CreateAssetWithdrawalV5RespData
	isSet bool
}

func (v NullableCreateAssetWithdrawalV5RespData) Get() *CreateAssetWithdrawalV5RespData {
	return v.value
}

func (v *NullableCreateAssetWithdrawalV5RespData) Set(val *CreateAssetWithdrawalV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssetWithdrawalV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssetWithdrawalV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssetWithdrawalV5RespData(val *CreateAssetWithdrawalV5RespData) *NullableCreateAssetWithdrawalV5RespData {
	return &NullableCreateAssetWithdrawalV5RespData{value: val, isSet: true}
}

func (v NullableCreateAssetWithdrawalV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAssetWithdrawalV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


