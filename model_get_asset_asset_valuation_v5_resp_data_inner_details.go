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

// checks if the GetAssetAssetValuationV5RespDataInnerDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAssetAssetValuationV5RespDataInnerDetails{}

// GetAssetAssetValuationV5RespDataInnerDetails Asset valuation details for each account
type GetAssetAssetValuationV5RespDataInnerDetails struct {
	// [Deprecated] Classic account
	// Deprecated
	Classic *string `json:"classic,omitempty"`
	// Earn account
	Earn *string `json:"earn,omitempty"`
	// Funding account
	Funding *string `json:"funding,omitempty"`
	// Trading account
	Trading *string `json:"trading,omitempty"`
}

// NewGetAssetAssetValuationV5RespDataInnerDetails instantiates a new GetAssetAssetValuationV5RespDataInnerDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssetAssetValuationV5RespDataInnerDetails() *GetAssetAssetValuationV5RespDataInnerDetails {
	this := GetAssetAssetValuationV5RespDataInnerDetails{}
	var classic string = ""
	this.Classic = &classic
	var earn string = ""
	this.Earn = &earn
	var funding string = ""
	this.Funding = &funding
	var trading string = ""
	this.Trading = &trading
	return &this
}

// NewGetAssetAssetValuationV5RespDataInnerDetailsWithDefaults instantiates a new GetAssetAssetValuationV5RespDataInnerDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssetAssetValuationV5RespDataInnerDetailsWithDefaults() *GetAssetAssetValuationV5RespDataInnerDetails {
	this := GetAssetAssetValuationV5RespDataInnerDetails{}
	var classic string = ""
	this.Classic = &classic
	var earn string = ""
	this.Earn = &earn
	var funding string = ""
	this.Funding = &funding
	var trading string = ""
	this.Trading = &trading
	return &this
}

// GetClassic returns the Classic field value if set, zero value otherwise.
// Deprecated
func (o *GetAssetAssetValuationV5RespDataInnerDetails) GetClassic() string {
	if o == nil || IsNil(o.Classic) {
		var ret string
		return ret
	}
	return *o.Classic
}

// GetClassicOk returns a tuple with the Classic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GetAssetAssetValuationV5RespDataInnerDetails) GetClassicOk() (*string, bool) {
	if o == nil || IsNil(o.Classic) {
		return nil, false
	}
	return o.Classic, true
}

// HasClassic returns a boolean if a field has been set.
func (o *GetAssetAssetValuationV5RespDataInnerDetails) HasClassic() bool {
	if o != nil && !IsNil(o.Classic) {
		return true
	}

	return false
}

// SetClassic gets a reference to the given string and assigns it to the Classic field.
// Deprecated
func (o *GetAssetAssetValuationV5RespDataInnerDetails) SetClassic(v string) {
	o.Classic = &v
}

// GetEarn returns the Earn field value if set, zero value otherwise.
func (o *GetAssetAssetValuationV5RespDataInnerDetails) GetEarn() string {
	if o == nil || IsNil(o.Earn) {
		var ret string
		return ret
	}
	return *o.Earn
}

// GetEarnOk returns a tuple with the Earn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetAssetValuationV5RespDataInnerDetails) GetEarnOk() (*string, bool) {
	if o == nil || IsNil(o.Earn) {
		return nil, false
	}
	return o.Earn, true
}

// HasEarn returns a boolean if a field has been set.
func (o *GetAssetAssetValuationV5RespDataInnerDetails) HasEarn() bool {
	if o != nil && !IsNil(o.Earn) {
		return true
	}

	return false
}

// SetEarn gets a reference to the given string and assigns it to the Earn field.
func (o *GetAssetAssetValuationV5RespDataInnerDetails) SetEarn(v string) {
	o.Earn = &v
}

// GetFunding returns the Funding field value if set, zero value otherwise.
func (o *GetAssetAssetValuationV5RespDataInnerDetails) GetFunding() string {
	if o == nil || IsNil(o.Funding) {
		var ret string
		return ret
	}
	return *o.Funding
}

// GetFundingOk returns a tuple with the Funding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetAssetValuationV5RespDataInnerDetails) GetFundingOk() (*string, bool) {
	if o == nil || IsNil(o.Funding) {
		return nil, false
	}
	return o.Funding, true
}

// HasFunding returns a boolean if a field has been set.
func (o *GetAssetAssetValuationV5RespDataInnerDetails) HasFunding() bool {
	if o != nil && !IsNil(o.Funding) {
		return true
	}

	return false
}

// SetFunding gets a reference to the given string and assigns it to the Funding field.
func (o *GetAssetAssetValuationV5RespDataInnerDetails) SetFunding(v string) {
	o.Funding = &v
}

// GetTrading returns the Trading field value if set, zero value otherwise.
func (o *GetAssetAssetValuationV5RespDataInnerDetails) GetTrading() string {
	if o == nil || IsNil(o.Trading) {
		var ret string
		return ret
	}
	return *o.Trading
}

// GetTradingOk returns a tuple with the Trading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetAssetValuationV5RespDataInnerDetails) GetTradingOk() (*string, bool) {
	if o == nil || IsNil(o.Trading) {
		return nil, false
	}
	return o.Trading, true
}

// HasTrading returns a boolean if a field has been set.
func (o *GetAssetAssetValuationV5RespDataInnerDetails) HasTrading() bool {
	if o != nil && !IsNil(o.Trading) {
		return true
	}

	return false
}

// SetTrading gets a reference to the given string and assigns it to the Trading field.
func (o *GetAssetAssetValuationV5RespDataInnerDetails) SetTrading(v string) {
	o.Trading = &v
}

func (o GetAssetAssetValuationV5RespDataInnerDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssetAssetValuationV5RespDataInnerDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Classic) {
		toSerialize["classic"] = o.Classic
	}
	if !IsNil(o.Earn) {
		toSerialize["earn"] = o.Earn
	}
	if !IsNil(o.Funding) {
		toSerialize["funding"] = o.Funding
	}
	if !IsNil(o.Trading) {
		toSerialize["trading"] = o.Trading
	}
	return toSerialize, nil
}

type NullableGetAssetAssetValuationV5RespDataInnerDetails struct {
	value *GetAssetAssetValuationV5RespDataInnerDetails
	isSet bool
}

func (v NullableGetAssetAssetValuationV5RespDataInnerDetails) Get() *GetAssetAssetValuationV5RespDataInnerDetails {
	return v.value
}

func (v *NullableGetAssetAssetValuationV5RespDataInnerDetails) Set(val *GetAssetAssetValuationV5RespDataInnerDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetAssetValuationV5RespDataInnerDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetAssetValuationV5RespDataInnerDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetAssetValuationV5RespDataInnerDetails(val *GetAssetAssetValuationV5RespDataInnerDetails) *NullableGetAssetAssetValuationV5RespDataInnerDetails {
	return &NullableGetAssetAssetValuationV5RespDataInnerDetails{value: val, isSet: true}
}

func (v NullableGetAssetAssetValuationV5RespDataInnerDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetAssetValuationV5RespDataInnerDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


