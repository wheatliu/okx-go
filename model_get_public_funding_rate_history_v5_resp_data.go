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

// checks if the GetPublicFundingRateHistoryV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPublicFundingRateHistoryV5RespData{}

// GetPublicFundingRateHistoryV5RespData struct for GetPublicFundingRateHistoryV5RespData
type GetPublicFundingRateHistoryV5RespData struct {
	// Formula type  `noRate`: old funding rate formula  `withRate`: new funding rate formula
	FormulaType *string `json:"formulaType,omitempty"`
	// Predicted funding rate
	FundingRate *string `json:"fundingRate,omitempty"`
	// Settlement time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	FundingTime *string `json:"fundingTime,omitempty"`
	// Instrument ID, e.g. `BTC-USD-SWAP`
	InstId *string `json:"instId,omitempty"`
	// Instrument type  `SWAP`
	InstType *string `json:"instType,omitempty"`
	// Funding rate mechanism   `current_period`   `next_period`
	Method *string `json:"method,omitempty"`
	// Actual funding rate
	RealizedRate *string `json:"realizedRate,omitempty"`
}

// NewGetPublicFundingRateHistoryV5RespData instantiates a new GetPublicFundingRateHistoryV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPublicFundingRateHistoryV5RespData() *GetPublicFundingRateHistoryV5RespData {
	this := GetPublicFundingRateHistoryV5RespData{}
	var formulaType string = ""
	this.FormulaType = &formulaType
	var fundingRate string = ""
	this.FundingRate = &fundingRate
	var fundingTime string = ""
	this.FundingTime = &fundingTime
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var method string = ""
	this.Method = &method
	var realizedRate string = ""
	this.RealizedRate = &realizedRate
	return &this
}

// NewGetPublicFundingRateHistoryV5RespDataWithDefaults instantiates a new GetPublicFundingRateHistoryV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPublicFundingRateHistoryV5RespDataWithDefaults() *GetPublicFundingRateHistoryV5RespData {
	this := GetPublicFundingRateHistoryV5RespData{}
	var formulaType string = ""
	this.FormulaType = &formulaType
	var fundingRate string = ""
	this.FundingRate = &fundingRate
	var fundingTime string = ""
	this.FundingTime = &fundingTime
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var method string = ""
	this.Method = &method
	var realizedRate string = ""
	this.RealizedRate = &realizedRate
	return &this
}

// GetFormulaType returns the FormulaType field value if set, zero value otherwise.
func (o *GetPublicFundingRateHistoryV5RespData) GetFormulaType() string {
	if o == nil || IsNil(o.FormulaType) {
		var ret string
		return ret
	}
	return *o.FormulaType
}

// GetFormulaTypeOk returns a tuple with the FormulaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateHistoryV5RespData) GetFormulaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FormulaType) {
		return nil, false
	}
	return o.FormulaType, true
}

// HasFormulaType returns a boolean if a field has been set.
func (o *GetPublicFundingRateHistoryV5RespData) HasFormulaType() bool {
	if o != nil && !IsNil(o.FormulaType) {
		return true
	}

	return false
}

// SetFormulaType gets a reference to the given string and assigns it to the FormulaType field.
func (o *GetPublicFundingRateHistoryV5RespData) SetFormulaType(v string) {
	o.FormulaType = &v
}

// GetFundingRate returns the FundingRate field value if set, zero value otherwise.
func (o *GetPublicFundingRateHistoryV5RespData) GetFundingRate() string {
	if o == nil || IsNil(o.FundingRate) {
		var ret string
		return ret
	}
	return *o.FundingRate
}

// GetFundingRateOk returns a tuple with the FundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateHistoryV5RespData) GetFundingRateOk() (*string, bool) {
	if o == nil || IsNil(o.FundingRate) {
		return nil, false
	}
	return o.FundingRate, true
}

// HasFundingRate returns a boolean if a field has been set.
func (o *GetPublicFundingRateHistoryV5RespData) HasFundingRate() bool {
	if o != nil && !IsNil(o.FundingRate) {
		return true
	}

	return false
}

// SetFundingRate gets a reference to the given string and assigns it to the FundingRate field.
func (o *GetPublicFundingRateHistoryV5RespData) SetFundingRate(v string) {
	o.FundingRate = &v
}

// GetFundingTime returns the FundingTime field value if set, zero value otherwise.
func (o *GetPublicFundingRateHistoryV5RespData) GetFundingTime() string {
	if o == nil || IsNil(o.FundingTime) {
		var ret string
		return ret
	}
	return *o.FundingTime
}

// GetFundingTimeOk returns a tuple with the FundingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateHistoryV5RespData) GetFundingTimeOk() (*string, bool) {
	if o == nil || IsNil(o.FundingTime) {
		return nil, false
	}
	return o.FundingTime, true
}

// HasFundingTime returns a boolean if a field has been set.
func (o *GetPublicFundingRateHistoryV5RespData) HasFundingTime() bool {
	if o != nil && !IsNil(o.FundingTime) {
		return true
	}

	return false
}

// SetFundingTime gets a reference to the given string and assigns it to the FundingTime field.
func (o *GetPublicFundingRateHistoryV5RespData) SetFundingTime(v string) {
	o.FundingTime = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetPublicFundingRateHistoryV5RespData) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateHistoryV5RespData) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetPublicFundingRateHistoryV5RespData) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetPublicFundingRateHistoryV5RespData) SetInstId(v string) {
	o.InstId = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *GetPublicFundingRateHistoryV5RespData) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateHistoryV5RespData) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *GetPublicFundingRateHistoryV5RespData) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *GetPublicFundingRateHistoryV5RespData) SetInstType(v string) {
	o.InstType = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *GetPublicFundingRateHistoryV5RespData) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateHistoryV5RespData) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *GetPublicFundingRateHistoryV5RespData) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *GetPublicFundingRateHistoryV5RespData) SetMethod(v string) {
	o.Method = &v
}

// GetRealizedRate returns the RealizedRate field value if set, zero value otherwise.
func (o *GetPublicFundingRateHistoryV5RespData) GetRealizedRate() string {
	if o == nil || IsNil(o.RealizedRate) {
		var ret string
		return ret
	}
	return *o.RealizedRate
}

// GetRealizedRateOk returns a tuple with the RealizedRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateHistoryV5RespData) GetRealizedRateOk() (*string, bool) {
	if o == nil || IsNil(o.RealizedRate) {
		return nil, false
	}
	return o.RealizedRate, true
}

// HasRealizedRate returns a boolean if a field has been set.
func (o *GetPublicFundingRateHistoryV5RespData) HasRealizedRate() bool {
	if o != nil && !IsNil(o.RealizedRate) {
		return true
	}

	return false
}

// SetRealizedRate gets a reference to the given string and assigns it to the RealizedRate field.
func (o *GetPublicFundingRateHistoryV5RespData) SetRealizedRate(v string) {
	o.RealizedRate = &v
}

func (o GetPublicFundingRateHistoryV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPublicFundingRateHistoryV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FormulaType) {
		toSerialize["formulaType"] = o.FormulaType
	}
	if !IsNil(o.FundingRate) {
		toSerialize["fundingRate"] = o.FundingRate
	}
	if !IsNil(o.FundingTime) {
		toSerialize["fundingTime"] = o.FundingTime
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.RealizedRate) {
		toSerialize["realizedRate"] = o.RealizedRate
	}
	return toSerialize, nil
}

type NullableGetPublicFundingRateHistoryV5RespData struct {
	value *GetPublicFundingRateHistoryV5RespData
	isSet bool
}

func (v NullableGetPublicFundingRateHistoryV5RespData) Get() *GetPublicFundingRateHistoryV5RespData {
	return v.value
}

func (v *NullableGetPublicFundingRateHistoryV5RespData) Set(val *GetPublicFundingRateHistoryV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPublicFundingRateHistoryV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPublicFundingRateHistoryV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPublicFundingRateHistoryV5RespData(val *GetPublicFundingRateHistoryV5RespData) *NullableGetPublicFundingRateHistoryV5RespData {
	return &NullableGetPublicFundingRateHistoryV5RespData{value: val, isSet: true}
}

func (v NullableGetPublicFundingRateHistoryV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPublicFundingRateHistoryV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


