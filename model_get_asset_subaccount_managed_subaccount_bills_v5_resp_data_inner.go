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

// checks if the GetAssetSubaccountManagedSubaccountBillsV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAssetSubaccountManagedSubaccountBillsV5RespDataInner{}

// GetAssetSubaccountManagedSubaccountBillsV5RespDataInner struct for GetAssetSubaccountManagedSubaccountBillsV5RespDataInner
type GetAssetSubaccountManagedSubaccountBillsV5RespDataInner struct {
	// Transfer amount
	Amt *string `json:"amt,omitempty"`
	// Bill ID
	BillId *string `json:"billId,omitempty"`
	// Transfer currency
	Ccy *string `json:"ccy,omitempty"`
	// Sub-account name
	SubAcct *string `json:"subAcct,omitempty"`
	// Sub-account UID
	SubUid *string `json:"subUid,omitempty"`
	// Bill ID creation time, Unix timestamp in millisecond format, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
	// Bill type
	Type *string `json:"type,omitempty"`
}

// NewGetAssetSubaccountManagedSubaccountBillsV5RespDataInner instantiates a new GetAssetSubaccountManagedSubaccountBillsV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssetSubaccountManagedSubaccountBillsV5RespDataInner() *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner {
	this := GetAssetSubaccountManagedSubaccountBillsV5RespDataInner{}
	var amt string = ""
	this.Amt = &amt
	var billId string = ""
	this.BillId = &billId
	var ccy string = ""
	this.Ccy = &ccy
	var subAcct string = ""
	this.SubAcct = &subAcct
	var subUid string = ""
	this.SubUid = &subUid
	var ts string = ""
	this.Ts = &ts
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewGetAssetSubaccountManagedSubaccountBillsV5RespDataInnerWithDefaults instantiates a new GetAssetSubaccountManagedSubaccountBillsV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssetSubaccountManagedSubaccountBillsV5RespDataInnerWithDefaults() *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner {
	this := GetAssetSubaccountManagedSubaccountBillsV5RespDataInner{}
	var amt string = ""
	this.Amt = &amt
	var billId string = ""
	this.BillId = &billId
	var ccy string = ""
	this.Ccy = &ccy
	var subAcct string = ""
	this.SubAcct = &subAcct
	var subUid string = ""
	this.SubUid = &subUid
	var ts string = ""
	this.Ts = &ts
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) SetAmt(v string) {
	o.Amt = &v
}

// GetBillId returns the BillId field value if set, zero value otherwise.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetBillId() string {
	if o == nil || IsNil(o.BillId) {
		var ret string
		return ret
	}
	return *o.BillId
}

// GetBillIdOk returns a tuple with the BillId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetBillIdOk() (*string, bool) {
	if o == nil || IsNil(o.BillId) {
		return nil, false
	}
	return o.BillId, true
}

// HasBillId returns a boolean if a field has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) HasBillId() bool {
	if o != nil && !IsNil(o.BillId) {
		return true
	}

	return false
}

// SetBillId gets a reference to the given string and assigns it to the BillId field.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) SetBillId(v string) {
	o.BillId = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetSubAcct returns the SubAcct field value if set, zero value otherwise.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetSubAcct() string {
	if o == nil || IsNil(o.SubAcct) {
		var ret string
		return ret
	}
	return *o.SubAcct
}

// GetSubAcctOk returns a tuple with the SubAcct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetSubAcctOk() (*string, bool) {
	if o == nil || IsNil(o.SubAcct) {
		return nil, false
	}
	return o.SubAcct, true
}

// HasSubAcct returns a boolean if a field has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) HasSubAcct() bool {
	if o != nil && !IsNil(o.SubAcct) {
		return true
	}

	return false
}

// SetSubAcct gets a reference to the given string and assigns it to the SubAcct field.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) SetSubAcct(v string) {
	o.SubAcct = &v
}

// GetSubUid returns the SubUid field value if set, zero value otherwise.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetSubUid() string {
	if o == nil || IsNil(o.SubUid) {
		var ret string
		return ret
	}
	return *o.SubUid
}

// GetSubUidOk returns a tuple with the SubUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetSubUidOk() (*string, bool) {
	if o == nil || IsNil(o.SubUid) {
		return nil, false
	}
	return o.SubUid, true
}

// HasSubUid returns a boolean if a field has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) HasSubUid() bool {
	if o != nil && !IsNil(o.SubUid) {
		return true
	}

	return false
}

// SetSubUid gets a reference to the given string and assigns it to the SubUid field.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) SetSubUid(v string) {
	o.SubUid = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) SetType(v string) {
	o.Type = &v
}

func (o GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.BillId) {
		toSerialize["billId"] = o.BillId
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.SubAcct) {
		toSerialize["subAcct"] = o.SubAcct
	}
	if !IsNil(o.SubUid) {
		toSerialize["subUid"] = o.SubUid
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGetAssetSubaccountManagedSubaccountBillsV5RespDataInner struct {
	value *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner
	isSet bool
}

func (v NullableGetAssetSubaccountManagedSubaccountBillsV5RespDataInner) Get() *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner {
	return v.value
}

func (v *NullableGetAssetSubaccountManagedSubaccountBillsV5RespDataInner) Set(val *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetSubaccountManagedSubaccountBillsV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetSubaccountManagedSubaccountBillsV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetSubaccountManagedSubaccountBillsV5RespDataInner(val *GetAssetSubaccountManagedSubaccountBillsV5RespDataInner) *NullableGetAssetSubaccountManagedSubaccountBillsV5RespDataInner {
	return &NullableGetAssetSubaccountManagedSubaccountBillsV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetAssetSubaccountManagedSubaccountBillsV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetSubaccountManagedSubaccountBillsV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


