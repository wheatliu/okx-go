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

// checks if the GetAccountInterestAccruedV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountInterestAccruedV5RespData{}

// GetAccountInterestAccruedV5RespData struct for GetAccountInterestAccruedV5RespData
type GetAccountInterestAccruedV5RespData struct {
	// Loan currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Instrument ID, e.g. `BTC-USDT`  Only applicable to `Market loans`
	InstId *string `json:"instId,omitempty"`
	// Interest
	Interest *string `json:"interest,omitempty"`
	// Interest rate (in hour)
	InterestRate *string `json:"interestRate,omitempty"`
	// Liability
	Liab *string `json:"liab,omitempty"`
	// Margin mode  `cross`    `isolated`
	MgnMode *string `json:"mgnMode,omitempty"`
	// Timestamp for interest accured, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
	// Loan type  `2`: Market loans
	Type *string `json:"type,omitempty"`
}

// NewGetAccountInterestAccruedV5RespData instantiates a new GetAccountInterestAccruedV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountInterestAccruedV5RespData() *GetAccountInterestAccruedV5RespData {
	this := GetAccountInterestAccruedV5RespData{}
	var ccy string = ""
	this.Ccy = &ccy
	var instId string = ""
	this.InstId = &instId
	var interest string = ""
	this.Interest = &interest
	var interestRate string = ""
	this.InterestRate = &interestRate
	var liab string = ""
	this.Liab = &liab
	var mgnMode string = ""
	this.MgnMode = &mgnMode
	var ts string = ""
	this.Ts = &ts
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewGetAccountInterestAccruedV5RespDataWithDefaults instantiates a new GetAccountInterestAccruedV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountInterestAccruedV5RespDataWithDefaults() *GetAccountInterestAccruedV5RespData {
	this := GetAccountInterestAccruedV5RespData{}
	var ccy string = ""
	this.Ccy = &ccy
	var instId string = ""
	this.InstId = &instId
	var interest string = ""
	this.Interest = &interest
	var interestRate string = ""
	this.InterestRate = &interestRate
	var liab string = ""
	this.Liab = &liab
	var mgnMode string = ""
	this.MgnMode = &mgnMode
	var ts string = ""
	this.Ts = &ts
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetAccountInterestAccruedV5RespData) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestAccruedV5RespData) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetAccountInterestAccruedV5RespData) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetAccountInterestAccruedV5RespData) SetCcy(v string) {
	o.Ccy = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *GetAccountInterestAccruedV5RespData) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestAccruedV5RespData) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *GetAccountInterestAccruedV5RespData) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *GetAccountInterestAccruedV5RespData) SetInstId(v string) {
	o.InstId = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *GetAccountInterestAccruedV5RespData) GetInterest() string {
	if o == nil || IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestAccruedV5RespData) GetInterestOk() (*string, bool) {
	if o == nil || IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *GetAccountInterestAccruedV5RespData) HasInterest() bool {
	if o != nil && !IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *GetAccountInterestAccruedV5RespData) SetInterest(v string) {
	o.Interest = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *GetAccountInterestAccruedV5RespData) GetInterestRate() string {
	if o == nil || IsNil(o.InterestRate) {
		var ret string
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestAccruedV5RespData) GetInterestRateOk() (*string, bool) {
	if o == nil || IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *GetAccountInterestAccruedV5RespData) HasInterestRate() bool {
	if o != nil && !IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given string and assigns it to the InterestRate field.
func (o *GetAccountInterestAccruedV5RespData) SetInterestRate(v string) {
	o.InterestRate = &v
}

// GetLiab returns the Liab field value if set, zero value otherwise.
func (o *GetAccountInterestAccruedV5RespData) GetLiab() string {
	if o == nil || IsNil(o.Liab) {
		var ret string
		return ret
	}
	return *o.Liab
}

// GetLiabOk returns a tuple with the Liab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestAccruedV5RespData) GetLiabOk() (*string, bool) {
	if o == nil || IsNil(o.Liab) {
		return nil, false
	}
	return o.Liab, true
}

// HasLiab returns a boolean if a field has been set.
func (o *GetAccountInterestAccruedV5RespData) HasLiab() bool {
	if o != nil && !IsNil(o.Liab) {
		return true
	}

	return false
}

// SetLiab gets a reference to the given string and assigns it to the Liab field.
func (o *GetAccountInterestAccruedV5RespData) SetLiab(v string) {
	o.Liab = &v
}

// GetMgnMode returns the MgnMode field value if set, zero value otherwise.
func (o *GetAccountInterestAccruedV5RespData) GetMgnMode() string {
	if o == nil || IsNil(o.MgnMode) {
		var ret string
		return ret
	}
	return *o.MgnMode
}

// GetMgnModeOk returns a tuple with the MgnMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestAccruedV5RespData) GetMgnModeOk() (*string, bool) {
	if o == nil || IsNil(o.MgnMode) {
		return nil, false
	}
	return o.MgnMode, true
}

// HasMgnMode returns a boolean if a field has been set.
func (o *GetAccountInterestAccruedV5RespData) HasMgnMode() bool {
	if o != nil && !IsNil(o.MgnMode) {
		return true
	}

	return false
}

// SetMgnMode gets a reference to the given string and assigns it to the MgnMode field.
func (o *GetAccountInterestAccruedV5RespData) SetMgnMode(v string) {
	o.MgnMode = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetAccountInterestAccruedV5RespData) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestAccruedV5RespData) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetAccountInterestAccruedV5RespData) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetAccountInterestAccruedV5RespData) SetTs(v string) {
	o.Ts = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetAccountInterestAccruedV5RespData) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestAccruedV5RespData) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetAccountInterestAccruedV5RespData) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetAccountInterestAccruedV5RespData) SetType(v string) {
	o.Type = &v
}

func (o GetAccountInterestAccruedV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountInterestAccruedV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !IsNil(o.InterestRate) {
		toSerialize["interestRate"] = o.InterestRate
	}
	if !IsNil(o.Liab) {
		toSerialize["liab"] = o.Liab
	}
	if !IsNil(o.MgnMode) {
		toSerialize["mgnMode"] = o.MgnMode
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGetAccountInterestAccruedV5RespData struct {
	value *GetAccountInterestAccruedV5RespData
	isSet bool
}

func (v NullableGetAccountInterestAccruedV5RespData) Get() *GetAccountInterestAccruedV5RespData {
	return v.value
}

func (v *NullableGetAccountInterestAccruedV5RespData) Set(val *GetAccountInterestAccruedV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountInterestAccruedV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountInterestAccruedV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountInterestAccruedV5RespData(val *GetAccountInterestAccruedV5RespData) *NullableGetAccountInterestAccruedV5RespData {
	return &NullableGetAccountInterestAccruedV5RespData{value: val, isSet: true}
}

func (v NullableGetAccountInterestAccruedV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountInterestAccruedV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


