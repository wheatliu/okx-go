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

// checks if the GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData{}

// GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData Risk warning data
type GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData struct {
	// Liquidation instrument ID, e.g. `BTC-USDT`  This field is only valid when there is only one type of collateral and one type of borrowed currency. In other cases, it returns \"\".
	InstId *string `json:"instId,omitempty"`
	// Liquidation price  The unit of the liquidation price is the quote currency of the instrument, e.g. `USDT` in `BTC-USDT`.  This field is only valid when there is only one type of collateral and one type of borrowed currency. In other cases, it returns \"\".
	LiqPx *string `json:"liqPx,omitempty"`
}

// NewGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData instantiates a new GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData() *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData {
	this := GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData{}
	var instId string = ""
	this.InstId = &instId
	var liqPx string = ""
	this.LiqPx = &liqPx
	return &this
}

// NewGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningDataWithDefaults instantiates a new GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningDataWithDefaults() *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData {
	this := GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData{}
	var instId string = ""
	this.InstId = &instId
	var liqPx string = ""
	this.LiqPx = &liqPx
	return &this
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) SetInstId(v string) {
	o.InstId = &v
}

// GetLiqPx returns the LiqPx field value if set, zero value otherwise.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) GetLiqPx() string {
	if o == nil || IsNil(o.LiqPx) {
		var ret string
		return ret
	}
	return *o.LiqPx
}

// GetLiqPxOk returns a tuple with the LiqPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) GetLiqPxOk() (*string, bool) {
	if o == nil || IsNil(o.LiqPx) {
		return nil, false
	}
	return o.LiqPx, true
}

// HasLiqPx returns a boolean if a field has been set.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) HasLiqPx() bool {
	if o != nil && !IsNil(o.LiqPx) {
		return true
	}

	return false
}

// SetLiqPx gets a reference to the given string and assigns it to the LiqPx field.
func (o *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) SetLiqPx(v string) {
	o.LiqPx = &v
}

func (o GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.LiqPx) {
		toSerialize["liqPx"] = o.LiqPx
	}
	return toSerialize, nil
}

type NullableGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData struct {
	value *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData
	isSet bool
}

func (v NullableGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) Get() *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData {
	return v.value
}

func (v *NullableGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) Set(val *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData(val *GetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) *NullableGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData {
	return &NullableGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData{value: val, isSet: true}
}

func (v NullableGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceFlexibleLoanLoanInfoV5RespDataRiskWarningData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


