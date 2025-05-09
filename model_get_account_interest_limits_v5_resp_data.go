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

// checks if the GetAccountInterestLimitsV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountInterestLimitsV5RespData{}

// GetAccountInterestLimitsV5RespData struct for GetAccountInterestLimitsV5RespData
type GetAccountInterestLimitsV5RespData struct {
	// Current debt in `USDT`
	Debt *string `json:"debt,omitempty"`
	// Current interest in `USDT`, the unit is `USDT`  Only applicable to `Market loans`
	Interest *string `json:"interest,omitempty"`
	// VIP Loan allocation for the current trading account  1. The unit is percent(%). Range is [0, 100]. Precision is 0.01%  2. If master account did not assign anything, then \"0\"  3. \"\" if shared between master and sub-account
	LoanAlloc *string `json:"loanAlloc,omitempty"`
	// Next deduct time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	NextDiscountTime *string `json:"nextDiscountTime,omitempty"`
	// Next accrual time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	NextInterestTime *string `json:"nextInterestTime,omitempty"`
	// Details for currencies
	Records []GetAccountInterestLimitsV5RespDataRecordsInner `json:"records,omitempty"`
}

// NewGetAccountInterestLimitsV5RespData instantiates a new GetAccountInterestLimitsV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountInterestLimitsV5RespData() *GetAccountInterestLimitsV5RespData {
	this := GetAccountInterestLimitsV5RespData{}
	var debt string = ""
	this.Debt = &debt
	var interest string = ""
	this.Interest = &interest
	var loanAlloc string = ""
	this.LoanAlloc = &loanAlloc
	var nextDiscountTime string = ""
	this.NextDiscountTime = &nextDiscountTime
	var nextInterestTime string = ""
	this.NextInterestTime = &nextInterestTime
	return &this
}

// NewGetAccountInterestLimitsV5RespDataWithDefaults instantiates a new GetAccountInterestLimitsV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountInterestLimitsV5RespDataWithDefaults() *GetAccountInterestLimitsV5RespData {
	this := GetAccountInterestLimitsV5RespData{}
	var debt string = ""
	this.Debt = &debt
	var interest string = ""
	this.Interest = &interest
	var loanAlloc string = ""
	this.LoanAlloc = &loanAlloc
	var nextDiscountTime string = ""
	this.NextDiscountTime = &nextDiscountTime
	var nextInterestTime string = ""
	this.NextInterestTime = &nextInterestTime
	return &this
}

// GetDebt returns the Debt field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespData) GetDebt() string {
	if o == nil || IsNil(o.Debt) {
		var ret string
		return ret
	}
	return *o.Debt
}

// GetDebtOk returns a tuple with the Debt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespData) GetDebtOk() (*string, bool) {
	if o == nil || IsNil(o.Debt) {
		return nil, false
	}
	return o.Debt, true
}

// HasDebt returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespData) HasDebt() bool {
	if o != nil && !IsNil(o.Debt) {
		return true
	}

	return false
}

// SetDebt gets a reference to the given string and assigns it to the Debt field.
func (o *GetAccountInterestLimitsV5RespData) SetDebt(v string) {
	o.Debt = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespData) GetInterest() string {
	if o == nil || IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespData) GetInterestOk() (*string, bool) {
	if o == nil || IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespData) HasInterest() bool {
	if o != nil && !IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *GetAccountInterestLimitsV5RespData) SetInterest(v string) {
	o.Interest = &v
}

// GetLoanAlloc returns the LoanAlloc field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespData) GetLoanAlloc() string {
	if o == nil || IsNil(o.LoanAlloc) {
		var ret string
		return ret
	}
	return *o.LoanAlloc
}

// GetLoanAllocOk returns a tuple with the LoanAlloc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespData) GetLoanAllocOk() (*string, bool) {
	if o == nil || IsNil(o.LoanAlloc) {
		return nil, false
	}
	return o.LoanAlloc, true
}

// HasLoanAlloc returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespData) HasLoanAlloc() bool {
	if o != nil && !IsNil(o.LoanAlloc) {
		return true
	}

	return false
}

// SetLoanAlloc gets a reference to the given string and assigns it to the LoanAlloc field.
func (o *GetAccountInterestLimitsV5RespData) SetLoanAlloc(v string) {
	o.LoanAlloc = &v
}

// GetNextDiscountTime returns the NextDiscountTime field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespData) GetNextDiscountTime() string {
	if o == nil || IsNil(o.NextDiscountTime) {
		var ret string
		return ret
	}
	return *o.NextDiscountTime
}

// GetNextDiscountTimeOk returns a tuple with the NextDiscountTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespData) GetNextDiscountTimeOk() (*string, bool) {
	if o == nil || IsNil(o.NextDiscountTime) {
		return nil, false
	}
	return o.NextDiscountTime, true
}

// HasNextDiscountTime returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespData) HasNextDiscountTime() bool {
	if o != nil && !IsNil(o.NextDiscountTime) {
		return true
	}

	return false
}

// SetNextDiscountTime gets a reference to the given string and assigns it to the NextDiscountTime field.
func (o *GetAccountInterestLimitsV5RespData) SetNextDiscountTime(v string) {
	o.NextDiscountTime = &v
}

// GetNextInterestTime returns the NextInterestTime field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespData) GetNextInterestTime() string {
	if o == nil || IsNil(o.NextInterestTime) {
		var ret string
		return ret
	}
	return *o.NextInterestTime
}

// GetNextInterestTimeOk returns a tuple with the NextInterestTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespData) GetNextInterestTimeOk() (*string, bool) {
	if o == nil || IsNil(o.NextInterestTime) {
		return nil, false
	}
	return o.NextInterestTime, true
}

// HasNextInterestTime returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespData) HasNextInterestTime() bool {
	if o != nil && !IsNil(o.NextInterestTime) {
		return true
	}

	return false
}

// SetNextInterestTime gets a reference to the given string and assigns it to the NextInterestTime field.
func (o *GetAccountInterestLimitsV5RespData) SetNextInterestTime(v string) {
	o.NextInterestTime = &v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespData) GetRecords() []GetAccountInterestLimitsV5RespDataRecordsInner {
	if o == nil || IsNil(o.Records) {
		var ret []GetAccountInterestLimitsV5RespDataRecordsInner
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespData) GetRecordsOk() ([]GetAccountInterestLimitsV5RespDataRecordsInner, bool) {
	if o == nil || IsNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespData) HasRecords() bool {
	if o != nil && !IsNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []GetAccountInterestLimitsV5RespDataRecordsInner and assigns it to the Records field.
func (o *GetAccountInterestLimitsV5RespData) SetRecords(v []GetAccountInterestLimitsV5RespDataRecordsInner) {
	o.Records = v
}

func (o GetAccountInterestLimitsV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountInterestLimitsV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Debt) {
		toSerialize["debt"] = o.Debt
	}
	if !IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !IsNil(o.LoanAlloc) {
		toSerialize["loanAlloc"] = o.LoanAlloc
	}
	if !IsNil(o.NextDiscountTime) {
		toSerialize["nextDiscountTime"] = o.NextDiscountTime
	}
	if !IsNil(o.NextInterestTime) {
		toSerialize["nextInterestTime"] = o.NextInterestTime
	}
	if !IsNil(o.Records) {
		toSerialize["records"] = o.Records
	}
	return toSerialize, nil
}

type NullableGetAccountInterestLimitsV5RespData struct {
	value *GetAccountInterestLimitsV5RespData
	isSet bool
}

func (v NullableGetAccountInterestLimitsV5RespData) Get() *GetAccountInterestLimitsV5RespData {
	return v.value
}

func (v *NullableGetAccountInterestLimitsV5RespData) Set(val *GetAccountInterestLimitsV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountInterestLimitsV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountInterestLimitsV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountInterestLimitsV5RespData(val *GetAccountInterestLimitsV5RespData) *NullableGetAccountInterestLimitsV5RespData {
	return &NullableGetAccountInterestLimitsV5RespData{value: val, isSet: true}
}

func (v NullableGetAccountInterestLimitsV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountInterestLimitsV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


