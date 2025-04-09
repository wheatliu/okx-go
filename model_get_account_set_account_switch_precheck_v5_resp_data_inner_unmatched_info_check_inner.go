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

// checks if the GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner{}

// GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner struct for GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner
type GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner struct {
	// Unmatched position list (posId)   Applicable when type is related to positions, return [] for other scenarios
	PosList []string `json:"posList,omitempty"`
	// Total assets   Only applicable when type is `asset_validation`, return \"\" for other scenarios
	TotalAsset *string `json:"totalAsset,omitempty"`
	// Unmatched information type   `asset_validation`: asset validation   `pending_orders`: order book pending orders   `pending_algos`: pending algo orders and trading bots, such as iceberg, recurring buy and twap   `isolated_margin`: isolated margin (quick margin and manual transfers)   `isolated_contract`: isolated contract (manual transfers)   `contract_long_short`: contract positions in hedge mode   `cross_margin`: cross margin positions   `cross_option_buyer`: cross options buyer   `isolated_option`: isolated options (only applicable to spot mode)   `growth_fund`: positions with trial funds   `all_positions`: all positions   `spot_lead_copy_only_simple_single`: copy trader and customize lead trader can only use spot mode or spot and futures mode   `stop_spot_custom`: spot customize copy trading   `stop_futures_custom`: contract customize copy trading   `lead_portfolio`: lead trader can not switch to portfolio margin mode   `futures_smart_sync`: you can not switch to spot mode when having smart contract sync   `vip_fixed_loan`: vip loan   `repay_borrowings`: borrowings   `compliance_restriction`: due to compliance restrictions, margin trading services are unavailable   `compliance_kyc2`: Due to compliance restrictions, margin trading services are unavailable. If you are not a resident of this region, please complete kyc2 identity verification.
	Type *string `json:"type,omitempty"`
}

// NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner instantiates a new GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner() *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner {
	this := GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner{}
	var totalAsset string = ""
	this.TotalAsset = &totalAsset
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInnerWithDefaults instantiates a new GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInnerWithDefaults() *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner {
	this := GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner{}
	var totalAsset string = ""
	this.TotalAsset = &totalAsset
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetPosList returns the PosList field value if set, zero value otherwise.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) GetPosList() []string {
	if o == nil || IsNil(o.PosList) {
		var ret []string
		return ret
	}
	return o.PosList
}

// GetPosListOk returns a tuple with the PosList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) GetPosListOk() ([]string, bool) {
	if o == nil || IsNil(o.PosList) {
		return nil, false
	}
	return o.PosList, true
}

// HasPosList returns a boolean if a field has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) HasPosList() bool {
	if o != nil && !IsNil(o.PosList) {
		return true
	}

	return false
}

// SetPosList gets a reference to the given []string and assigns it to the PosList field.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) SetPosList(v []string) {
	o.PosList = v
}

// GetTotalAsset returns the TotalAsset field value if set, zero value otherwise.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) GetTotalAsset() string {
	if o == nil || IsNil(o.TotalAsset) {
		var ret string
		return ret
	}
	return *o.TotalAsset
}

// GetTotalAssetOk returns a tuple with the TotalAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) GetTotalAssetOk() (*string, bool) {
	if o == nil || IsNil(o.TotalAsset) {
		return nil, false
	}
	return o.TotalAsset, true
}

// HasTotalAsset returns a boolean if a field has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) HasTotalAsset() bool {
	if o != nil && !IsNil(o.TotalAsset) {
		return true
	}

	return false
}

// SetTotalAsset gets a reference to the given string and assigns it to the TotalAsset field.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) SetTotalAsset(v string) {
	o.TotalAsset = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) SetType(v string) {
	o.Type = &v
}

func (o GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PosList) {
		toSerialize["posList"] = o.PosList
	}
	if !IsNil(o.TotalAsset) {
		toSerialize["totalAsset"] = o.TotalAsset
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner struct {
	value *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner
	isSet bool
}

func (v NullableGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) Get() *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner {
	return v.value
}

func (v *NullableGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) Set(val *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner(val *GetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) *NullableGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner {
	return &NullableGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner{value: val, isSet: true}
}

func (v NullableGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountSetAccountSwitchPrecheckV5RespDataInnerUnmatchedInfoCheckInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


