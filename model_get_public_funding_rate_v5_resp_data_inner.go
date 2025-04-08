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

// checks if the GetPublicFundingRateV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPublicFundingRateV5RespDataInner{}

// GetPublicFundingRateV5RespDataInner struct for GetPublicFundingRateV5RespDataInner
type GetPublicFundingRateV5RespDataInner struct {
	// Formula type  `noRate`: old funding rate formula  `withRate`: new funding rate formula
	FormulaType *string `json:"formulaType,omitempty"`
	// Current funding rate
	FundingRate *string `json:"fundingRate,omitempty"`
	// Settlement time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	FundingTime *string `json:"fundingTime,omitempty"`
	// Depth weighted amount (in the unit of quote currency)
	ImpactValue *string `json:"impactValue,omitempty"`
	// Instrument ID, e.g. `BTC-USD-SWAP`
	InstId *string `json:"instId,omitempty"`
	// Instrument type  `SWAP`
	InstType *string `json:"instType,omitempty"`
	// Interest rate
	InterestRate *string `json:"interestRate,omitempty"`
	// The upper limit of the predicted funding rate of the next cycle
	MaxFundingRate *string `json:"maxFundingRate,omitempty"`
	// Funding rate mechanism   `current_period`   `next_period`(no longer supported)
	Method *string `json:"method,omitempty"`
	// The lower limit of the predicted funding rate of the next cycle
	MinFundingRate *string `json:"minFundingRate,omitempty"`
	// Forecasted funding rate for the next period   The nextFundingRate will be \"\" if the method is `current_period`(no longer supported)
	NextFundingRate *string `json:"nextFundingRate,omitempty"`
	// Forecasted funding time for the next period , Unix timestamp format in milliseconds, e.g. `1597026383085`
	NextFundingTime *string `json:"nextFundingTime,omitempty"`
	// Premium between the mid price of perps market and the index price
	Premium *string `json:"premium,omitempty"`
	// If settState = `processing`, it is the funding rate that is being used for current settlement cycle.   If settState = `settled`, it is the funding rate that is being used for previous settlement cycle
	SettFundingRate *string `json:"settFundingRate,omitempty"`
	// Settlement state of funding rate   `processing`   `settled`
	SettState *string `json:"settState,omitempty"`
	// Data return time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewGetPublicFundingRateV5RespDataInner instantiates a new GetPublicFundingRateV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPublicFundingRateV5RespDataInner() *GetPublicFundingRateV5RespDataInner {
	this := GetPublicFundingRateV5RespDataInner{}
	var formulaType string = ""
	this.FormulaType = &formulaType
	var fundingRate string = ""
	this.FundingRate = &fundingRate
	var fundingTime string = ""
	this.FundingTime = &fundingTime
	var impactValue string = ""
	this.ImpactValue = &impactValue
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var interestRate string = ""
	this.InterestRate = &interestRate
	var maxFundingRate string = ""
	this.MaxFundingRate = &maxFundingRate
	var method string = ""
	this.Method = &method
	var minFundingRate string = ""
	this.MinFundingRate = &minFundingRate
	var nextFundingRate string = ""
	this.NextFundingRate = &nextFundingRate
	var nextFundingTime string = ""
	this.NextFundingTime = &nextFundingTime
	var premium string = ""
	this.Premium = &premium
	var settFundingRate string = ""
	this.SettFundingRate = &settFundingRate
	var settState string = ""
	this.SettState = &settState
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetPublicFundingRateV5RespDataInnerWithDefaults instantiates a new GetPublicFundingRateV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPublicFundingRateV5RespDataInnerWithDefaults() *GetPublicFundingRateV5RespDataInner {
	this := GetPublicFundingRateV5RespDataInner{}
	var formulaType string = ""
	this.FormulaType = &formulaType
	var fundingRate string = ""
	this.FundingRate = &fundingRate
	var fundingTime string = ""
	this.FundingTime = &fundingTime
	var impactValue string = ""
	this.ImpactValue = &impactValue
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var interestRate string = ""
	this.InterestRate = &interestRate
	var maxFundingRate string = ""
	this.MaxFundingRate = &maxFundingRate
	var method string = ""
	this.Method = &method
	var minFundingRate string = ""
	this.MinFundingRate = &minFundingRate
	var nextFundingRate string = ""
	this.NextFundingRate = &nextFundingRate
	var nextFundingTime string = ""
	this.NextFundingTime = &nextFundingTime
	var premium string = ""
	this.Premium = &premium
	var settFundingRate string = ""
	this.SettFundingRate = &settFundingRate
	var settState string = ""
	this.SettState = &settState
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetFormulaType returns the FormulaType field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetFormulaType() string {
	if o == nil || IsNil(o.FormulaType) {
		var ret string
		return ret
	}
	return *o.FormulaType
}

// GetFormulaTypeOk returns a tuple with the FormulaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetFormulaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FormulaType) {
		return nil, false
	}
	return o.FormulaType, true
}

// HasFormulaType returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasFormulaType() bool {
	if o != nil && !IsNil(o.FormulaType) {
		return true
	}

	return false
}

// SetFormulaType gets a reference to the given string and assigns it to the FormulaType field.
func (o *GetPublicFundingRateV5RespDataInner) SetFormulaType(v string) {
	o.FormulaType = &v
}

// GetFundingRate returns the FundingRate field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetFundingRate() string {
	if o == nil || IsNil(o.FundingRate) {
		var ret string
		return ret
	}
	return *o.FundingRate
}

// GetFundingRateOk returns a tuple with the FundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetFundingRateOk() (*string, bool) {
	if o == nil || IsNil(o.FundingRate) {
		return nil, false
	}
	return o.FundingRate, true
}

// HasFundingRate returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasFundingRate() bool {
	if o != nil && !IsNil(o.FundingRate) {
		return true
	}

	return false
}

// SetFundingRate gets a reference to the given string and assigns it to the FundingRate field.
func (o *GetPublicFundingRateV5RespDataInner) SetFundingRate(v string) {
	o.FundingRate = &v
}

// GetFundingTime returns the FundingTime field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetFundingTime() string {
	if o == nil || IsNil(o.FundingTime) {
		var ret string
		return ret
	}
	return *o.FundingTime
}

// GetFundingTimeOk returns a tuple with the FundingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetFundingTimeOk() (*string, bool) {
	if o == nil || IsNil(o.FundingTime) {
		return nil, false
	}
	return o.FundingTime, true
}

// HasFundingTime returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasFundingTime() bool {
	if o != nil && !IsNil(o.FundingTime) {
		return true
	}

	return false
}

// SetFundingTime gets a reference to the given string and assigns it to the FundingTime field.
func (o *GetPublicFundingRateV5RespDataInner) SetFundingTime(v string) {
	o.FundingTime = &v
}

// GetImpactValue returns the ImpactValue field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetImpactValue() string {
	if o == nil || IsNil(o.ImpactValue) {
		var ret string
		return ret
	}
	return *o.ImpactValue
}

// GetImpactValueOk returns a tuple with the ImpactValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetImpactValueOk() (*string, bool) {
	if o == nil || IsNil(o.ImpactValue) {
		return nil, false
	}
	return o.ImpactValue, true
}

// HasImpactValue returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasImpactValue() bool {
	if o != nil && !IsNil(o.ImpactValue) {
		return true
	}

	return false
}

// SetImpactValue gets a reference to the given string and assigns it to the ImpactValue field.
func (o *GetPublicFundingRateV5RespDataInner) SetImpactValue(v string) {
	o.ImpactValue = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetPublicFundingRateV5RespDataInner) SetInstId(v string) {
	o.InstId = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *GetPublicFundingRateV5RespDataInner) SetInstType(v string) {
	o.InstType = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetInterestRate() string {
	if o == nil || IsNil(o.InterestRate) {
		var ret string
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetInterestRateOk() (*string, bool) {
	if o == nil || IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasInterestRate() bool {
	if o != nil && !IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given string and assigns it to the InterestRate field.
func (o *GetPublicFundingRateV5RespDataInner) SetInterestRate(v string) {
	o.InterestRate = &v
}

// GetMaxFundingRate returns the MaxFundingRate field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetMaxFundingRate() string {
	if o == nil || IsNil(o.MaxFundingRate) {
		var ret string
		return ret
	}
	return *o.MaxFundingRate
}

// GetMaxFundingRateOk returns a tuple with the MaxFundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetMaxFundingRateOk() (*string, bool) {
	if o == nil || IsNil(o.MaxFundingRate) {
		return nil, false
	}
	return o.MaxFundingRate, true
}

// HasMaxFundingRate returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasMaxFundingRate() bool {
	if o != nil && !IsNil(o.MaxFundingRate) {
		return true
	}

	return false
}

// SetMaxFundingRate gets a reference to the given string and assigns it to the MaxFundingRate field.
func (o *GetPublicFundingRateV5RespDataInner) SetMaxFundingRate(v string) {
	o.MaxFundingRate = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *GetPublicFundingRateV5RespDataInner) SetMethod(v string) {
	o.Method = &v
}

// GetMinFundingRate returns the MinFundingRate field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetMinFundingRate() string {
	if o == nil || IsNil(o.MinFundingRate) {
		var ret string
		return ret
	}
	return *o.MinFundingRate
}

// GetMinFundingRateOk returns a tuple with the MinFundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetMinFundingRateOk() (*string, bool) {
	if o == nil || IsNil(o.MinFundingRate) {
		return nil, false
	}
	return o.MinFundingRate, true
}

// HasMinFundingRate returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasMinFundingRate() bool {
	if o != nil && !IsNil(o.MinFundingRate) {
		return true
	}

	return false
}

// SetMinFundingRate gets a reference to the given string and assigns it to the MinFundingRate field.
func (o *GetPublicFundingRateV5RespDataInner) SetMinFundingRate(v string) {
	o.MinFundingRate = &v
}

// GetNextFundingRate returns the NextFundingRate field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetNextFundingRate() string {
	if o == nil || IsNil(o.NextFundingRate) {
		var ret string
		return ret
	}
	return *o.NextFundingRate
}

// GetNextFundingRateOk returns a tuple with the NextFundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetNextFundingRateOk() (*string, bool) {
	if o == nil || IsNil(o.NextFundingRate) {
		return nil, false
	}
	return o.NextFundingRate, true
}

// HasNextFundingRate returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasNextFundingRate() bool {
	if o != nil && !IsNil(o.NextFundingRate) {
		return true
	}

	return false
}

// SetNextFundingRate gets a reference to the given string and assigns it to the NextFundingRate field.
func (o *GetPublicFundingRateV5RespDataInner) SetNextFundingRate(v string) {
	o.NextFundingRate = &v
}

// GetNextFundingTime returns the NextFundingTime field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetNextFundingTime() string {
	if o == nil || IsNil(o.NextFundingTime) {
		var ret string
		return ret
	}
	return *o.NextFundingTime
}

// GetNextFundingTimeOk returns a tuple with the NextFundingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetNextFundingTimeOk() (*string, bool) {
	if o == nil || IsNil(o.NextFundingTime) {
		return nil, false
	}
	return o.NextFundingTime, true
}

// HasNextFundingTime returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasNextFundingTime() bool {
	if o != nil && !IsNil(o.NextFundingTime) {
		return true
	}

	return false
}

// SetNextFundingTime gets a reference to the given string and assigns it to the NextFundingTime field.
func (o *GetPublicFundingRateV5RespDataInner) SetNextFundingTime(v string) {
	o.NextFundingTime = &v
}

// GetPremium returns the Premium field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetPremium() string {
	if o == nil || IsNil(o.Premium) {
		var ret string
		return ret
	}
	return *o.Premium
}

// GetPremiumOk returns a tuple with the Premium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetPremiumOk() (*string, bool) {
	if o == nil || IsNil(o.Premium) {
		return nil, false
	}
	return o.Premium, true
}

// HasPremium returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasPremium() bool {
	if o != nil && !IsNil(o.Premium) {
		return true
	}

	return false
}

// SetPremium gets a reference to the given string and assigns it to the Premium field.
func (o *GetPublicFundingRateV5RespDataInner) SetPremium(v string) {
	o.Premium = &v
}

// GetSettFundingRate returns the SettFundingRate field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetSettFundingRate() string {
	if o == nil || IsNil(o.SettFundingRate) {
		var ret string
		return ret
	}
	return *o.SettFundingRate
}

// GetSettFundingRateOk returns a tuple with the SettFundingRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetSettFundingRateOk() (*string, bool) {
	if o == nil || IsNil(o.SettFundingRate) {
		return nil, false
	}
	return o.SettFundingRate, true
}

// HasSettFundingRate returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasSettFundingRate() bool {
	if o != nil && !IsNil(o.SettFundingRate) {
		return true
	}

	return false
}

// SetSettFundingRate gets a reference to the given string and assigns it to the SettFundingRate field.
func (o *GetPublicFundingRateV5RespDataInner) SetSettFundingRate(v string) {
	o.SettFundingRate = &v
}

// GetSettState returns the SettState field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetSettState() string {
	if o == nil || IsNil(o.SettState) {
		var ret string
		return ret
	}
	return *o.SettState
}

// GetSettStateOk returns a tuple with the SettState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetSettStateOk() (*string, bool) {
	if o == nil || IsNil(o.SettState) {
		return nil, false
	}
	return o.SettState, true
}

// HasSettState returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasSettState() bool {
	if o != nil && !IsNil(o.SettState) {
		return true
	}

	return false
}

// SetSettState gets a reference to the given string and assigns it to the SettState field.
func (o *GetPublicFundingRateV5RespDataInner) SetSettState(v string) {
	o.SettState = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetPublicFundingRateV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPublicFundingRateV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetPublicFundingRateV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetPublicFundingRateV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetPublicFundingRateV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPublicFundingRateV5RespDataInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ImpactValue) {
		toSerialize["impactValue"] = o.ImpactValue
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.InterestRate) {
		toSerialize["interestRate"] = o.InterestRate
	}
	if !IsNil(o.MaxFundingRate) {
		toSerialize["maxFundingRate"] = o.MaxFundingRate
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.MinFundingRate) {
		toSerialize["minFundingRate"] = o.MinFundingRate
	}
	if !IsNil(o.NextFundingRate) {
		toSerialize["nextFundingRate"] = o.NextFundingRate
	}
	if !IsNil(o.NextFundingTime) {
		toSerialize["nextFundingTime"] = o.NextFundingTime
	}
	if !IsNil(o.Premium) {
		toSerialize["premium"] = o.Premium
	}
	if !IsNil(o.SettFundingRate) {
		toSerialize["settFundingRate"] = o.SettFundingRate
	}
	if !IsNil(o.SettState) {
		toSerialize["settState"] = o.SettState
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetPublicFundingRateV5RespDataInner struct {
	value *GetPublicFundingRateV5RespDataInner
	isSet bool
}

func (v NullableGetPublicFundingRateV5RespDataInner) Get() *GetPublicFundingRateV5RespDataInner {
	return v.value
}

func (v *NullableGetPublicFundingRateV5RespDataInner) Set(val *GetPublicFundingRateV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPublicFundingRateV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPublicFundingRateV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPublicFundingRateV5RespDataInner(val *GetPublicFundingRateV5RespDataInner) *NullableGetPublicFundingRateV5RespDataInner {
	return &NullableGetPublicFundingRateV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetPublicFundingRateV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPublicFundingRateV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


