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

// checks if the GetAccountInterestLimitsV5RespDataInnerRecordsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountInterestLimitsV5RespDataInnerRecordsInner{}

// GetAccountInterestLimitsV5RespDataInnerRecordsInner struct for GetAccountInterestLimitsV5RespDataInnerRecordsInner
type GetAccountInterestLimitsV5RespDataInnerRecordsInner struct {
	// Total remaining quota for master account and sub-accounts
	AllAcctRemainingQuota *string `json:"allAcctRemainingQuota,omitempty"`
	// Available amount for current account (Within the locked quota)  Only applicable to `VIP loans`
	AvailLoan *string `json:"availLoan,omitempty"`
	// Average (hour) interest of already borrowed coin  only applicable to `VIP loans`
	AvgRate *string `json:"avgRate,omitempty"`
	// Loan currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// The remaining quota for the current account.  Only applicable to the case in which the sub-account is assigned the loan allocation
	CurAcctRemainingQuota *string `json:"curAcctRemainingQuota,omitempty"`
	// Interest to be deducted  Only applicable to `Market loans`
	Interest *string `json:"interest,omitempty"`
	// Borrow limit of master account  If loan allocation has been assigned, then it is the borrow limit of the current trading account
	LoanQuota *string `json:"loanQuota,omitempty"`
	// Remaining quota for the platform.  The format like  \"600\" will be returned when it is more than `curAcctRemainingQuota` or `allAcctRemainingQuota`
	PlatRemainingQuota *string `json:"platRemainingQuota,omitempty"`
	// Frozen amount for current account (Within the locked quota)   Only applicable to `VIP loans`
	PosLoan *string `json:"posLoan,omitempty"`
	// Current daily rate
	Rate *string `json:"rate,omitempty"`
	// Available amount across all sub-accounts  If loan allocation has been assigned, then it is the available amount to borrow by the current trading account
	SurplusLmt *string `json:"surplusLmt,omitempty"`
	SurplusLmtDetails *GetAccountInterestLimitsV5RespDataInnerRecordsInnerSurplusLmtDetails `json:"surplusLmtDetails,omitempty"`
	// Borrowed amount across all sub-accounts  If loan allocation has been assigned, then it is the borrowed amount by the current trading account
	UsedLmt *string `json:"usedLmt,omitempty"`
	// Borrowed amount for current account  Only applicable to `VIP loans`
	UsedLoan *string `json:"usedLoan,omitempty"`
}

// NewGetAccountInterestLimitsV5RespDataInnerRecordsInner instantiates a new GetAccountInterestLimitsV5RespDataInnerRecordsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountInterestLimitsV5RespDataInnerRecordsInner() *GetAccountInterestLimitsV5RespDataInnerRecordsInner {
	this := GetAccountInterestLimitsV5RespDataInnerRecordsInner{}
	var allAcctRemainingQuota string = ""
	this.AllAcctRemainingQuota = &allAcctRemainingQuota
	var availLoan string = ""
	this.AvailLoan = &availLoan
	var avgRate string = ""
	this.AvgRate = &avgRate
	var ccy string = ""
	this.Ccy = &ccy
	var curAcctRemainingQuota string = ""
	this.CurAcctRemainingQuota = &curAcctRemainingQuota
	var interest string = ""
	this.Interest = &interest
	var loanQuota string = ""
	this.LoanQuota = &loanQuota
	var platRemainingQuota string = ""
	this.PlatRemainingQuota = &platRemainingQuota
	var posLoan string = ""
	this.PosLoan = &posLoan
	var rate string = ""
	this.Rate = &rate
	var surplusLmt string = ""
	this.SurplusLmt = &surplusLmt
	var usedLmt string = ""
	this.UsedLmt = &usedLmt
	var usedLoan string = ""
	this.UsedLoan = &usedLoan
	return &this
}

// NewGetAccountInterestLimitsV5RespDataInnerRecordsInnerWithDefaults instantiates a new GetAccountInterestLimitsV5RespDataInnerRecordsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountInterestLimitsV5RespDataInnerRecordsInnerWithDefaults() *GetAccountInterestLimitsV5RespDataInnerRecordsInner {
	this := GetAccountInterestLimitsV5RespDataInnerRecordsInner{}
	var allAcctRemainingQuota string = ""
	this.AllAcctRemainingQuota = &allAcctRemainingQuota
	var availLoan string = ""
	this.AvailLoan = &availLoan
	var avgRate string = ""
	this.AvgRate = &avgRate
	var ccy string = ""
	this.Ccy = &ccy
	var curAcctRemainingQuota string = ""
	this.CurAcctRemainingQuota = &curAcctRemainingQuota
	var interest string = ""
	this.Interest = &interest
	var loanQuota string = ""
	this.LoanQuota = &loanQuota
	var platRemainingQuota string = ""
	this.PlatRemainingQuota = &platRemainingQuota
	var posLoan string = ""
	this.PosLoan = &posLoan
	var rate string = ""
	this.Rate = &rate
	var surplusLmt string = ""
	this.SurplusLmt = &surplusLmt
	var usedLmt string = ""
	this.UsedLmt = &usedLmt
	var usedLoan string = ""
	this.UsedLoan = &usedLoan
	return &this
}

// GetAllAcctRemainingQuota returns the AllAcctRemainingQuota field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetAllAcctRemainingQuota() string {
	if o == nil || IsNil(o.AllAcctRemainingQuota) {
		var ret string
		return ret
	}
	return *o.AllAcctRemainingQuota
}

// GetAllAcctRemainingQuotaOk returns a tuple with the AllAcctRemainingQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetAllAcctRemainingQuotaOk() (*string, bool) {
	if o == nil || IsNil(o.AllAcctRemainingQuota) {
		return nil, false
	}
	return o.AllAcctRemainingQuota, true
}

// HasAllAcctRemainingQuota returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasAllAcctRemainingQuota() bool {
	if o != nil && !IsNil(o.AllAcctRemainingQuota) {
		return true
	}

	return false
}

// SetAllAcctRemainingQuota gets a reference to the given string and assigns it to the AllAcctRemainingQuota field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetAllAcctRemainingQuota(v string) {
	o.AllAcctRemainingQuota = &v
}

// GetAvailLoan returns the AvailLoan field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetAvailLoan() string {
	if o == nil || IsNil(o.AvailLoan) {
		var ret string
		return ret
	}
	return *o.AvailLoan
}

// GetAvailLoanOk returns a tuple with the AvailLoan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetAvailLoanOk() (*string, bool) {
	if o == nil || IsNil(o.AvailLoan) {
		return nil, false
	}
	return o.AvailLoan, true
}

// HasAvailLoan returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasAvailLoan() bool {
	if o != nil && !IsNil(o.AvailLoan) {
		return true
	}

	return false
}

// SetAvailLoan gets a reference to the given string and assigns it to the AvailLoan field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetAvailLoan(v string) {
	o.AvailLoan = &v
}

// GetAvgRate returns the AvgRate field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetAvgRate() string {
	if o == nil || IsNil(o.AvgRate) {
		var ret string
		return ret
	}
	return *o.AvgRate
}

// GetAvgRateOk returns a tuple with the AvgRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetAvgRateOk() (*string, bool) {
	if o == nil || IsNil(o.AvgRate) {
		return nil, false
	}
	return o.AvgRate, true
}

// HasAvgRate returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasAvgRate() bool {
	if o != nil && !IsNil(o.AvgRate) {
		return true
	}

	return false
}

// SetAvgRate gets a reference to the given string and assigns it to the AvgRate field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetAvgRate(v string) {
	o.AvgRate = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetCurAcctRemainingQuota returns the CurAcctRemainingQuota field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetCurAcctRemainingQuota() string {
	if o == nil || IsNil(o.CurAcctRemainingQuota) {
		var ret string
		return ret
	}
	return *o.CurAcctRemainingQuota
}

// GetCurAcctRemainingQuotaOk returns a tuple with the CurAcctRemainingQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetCurAcctRemainingQuotaOk() (*string, bool) {
	if o == nil || IsNil(o.CurAcctRemainingQuota) {
		return nil, false
	}
	return o.CurAcctRemainingQuota, true
}

// HasCurAcctRemainingQuota returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasCurAcctRemainingQuota() bool {
	if o != nil && !IsNil(o.CurAcctRemainingQuota) {
		return true
	}

	return false
}

// SetCurAcctRemainingQuota gets a reference to the given string and assigns it to the CurAcctRemainingQuota field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetCurAcctRemainingQuota(v string) {
	o.CurAcctRemainingQuota = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetInterest() string {
	if o == nil || IsNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetInterestOk() (*string, bool) {
	if o == nil || IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasInterest() bool {
	if o != nil && !IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetInterest(v string) {
	o.Interest = &v
}

// GetLoanQuota returns the LoanQuota field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetLoanQuota() string {
	if o == nil || IsNil(o.LoanQuota) {
		var ret string
		return ret
	}
	return *o.LoanQuota
}

// GetLoanQuotaOk returns a tuple with the LoanQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetLoanQuotaOk() (*string, bool) {
	if o == nil || IsNil(o.LoanQuota) {
		return nil, false
	}
	return o.LoanQuota, true
}

// HasLoanQuota returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasLoanQuota() bool {
	if o != nil && !IsNil(o.LoanQuota) {
		return true
	}

	return false
}

// SetLoanQuota gets a reference to the given string and assigns it to the LoanQuota field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetLoanQuota(v string) {
	o.LoanQuota = &v
}

// GetPlatRemainingQuota returns the PlatRemainingQuota field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetPlatRemainingQuota() string {
	if o == nil || IsNil(o.PlatRemainingQuota) {
		var ret string
		return ret
	}
	return *o.PlatRemainingQuota
}

// GetPlatRemainingQuotaOk returns a tuple with the PlatRemainingQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetPlatRemainingQuotaOk() (*string, bool) {
	if o == nil || IsNil(o.PlatRemainingQuota) {
		return nil, false
	}
	return o.PlatRemainingQuota, true
}

// HasPlatRemainingQuota returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasPlatRemainingQuota() bool {
	if o != nil && !IsNil(o.PlatRemainingQuota) {
		return true
	}

	return false
}

// SetPlatRemainingQuota gets a reference to the given string and assigns it to the PlatRemainingQuota field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetPlatRemainingQuota(v string) {
	o.PlatRemainingQuota = &v
}

// GetPosLoan returns the PosLoan field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetPosLoan() string {
	if o == nil || IsNil(o.PosLoan) {
		var ret string
		return ret
	}
	return *o.PosLoan
}

// GetPosLoanOk returns a tuple with the PosLoan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetPosLoanOk() (*string, bool) {
	if o == nil || IsNil(o.PosLoan) {
		return nil, false
	}
	return o.PosLoan, true
}

// HasPosLoan returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasPosLoan() bool {
	if o != nil && !IsNil(o.PosLoan) {
		return true
	}

	return false
}

// SetPosLoan gets a reference to the given string and assigns it to the PosLoan field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetPosLoan(v string) {
	o.PosLoan = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetRate(v string) {
	o.Rate = &v
}

// GetSurplusLmt returns the SurplusLmt field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetSurplusLmt() string {
	if o == nil || IsNil(o.SurplusLmt) {
		var ret string
		return ret
	}
	return *o.SurplusLmt
}

// GetSurplusLmtOk returns a tuple with the SurplusLmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetSurplusLmtOk() (*string, bool) {
	if o == nil || IsNil(o.SurplusLmt) {
		return nil, false
	}
	return o.SurplusLmt, true
}

// HasSurplusLmt returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasSurplusLmt() bool {
	if o != nil && !IsNil(o.SurplusLmt) {
		return true
	}

	return false
}

// SetSurplusLmt gets a reference to the given string and assigns it to the SurplusLmt field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetSurplusLmt(v string) {
	o.SurplusLmt = &v
}

// GetSurplusLmtDetails returns the SurplusLmtDetails field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetSurplusLmtDetails() GetAccountInterestLimitsV5RespDataInnerRecordsInnerSurplusLmtDetails {
	if o == nil || IsNil(o.SurplusLmtDetails) {
		var ret GetAccountInterestLimitsV5RespDataInnerRecordsInnerSurplusLmtDetails
		return ret
	}
	return *o.SurplusLmtDetails
}

// GetSurplusLmtDetailsOk returns a tuple with the SurplusLmtDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetSurplusLmtDetailsOk() (*GetAccountInterestLimitsV5RespDataInnerRecordsInnerSurplusLmtDetails, bool) {
	if o == nil || IsNil(o.SurplusLmtDetails) {
		return nil, false
	}
	return o.SurplusLmtDetails, true
}

// HasSurplusLmtDetails returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasSurplusLmtDetails() bool {
	if o != nil && !IsNil(o.SurplusLmtDetails) {
		return true
	}

	return false
}

// SetSurplusLmtDetails gets a reference to the given GetAccountInterestLimitsV5RespDataInnerRecordsInnerSurplusLmtDetails and assigns it to the SurplusLmtDetails field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetSurplusLmtDetails(v GetAccountInterestLimitsV5RespDataInnerRecordsInnerSurplusLmtDetails) {
	o.SurplusLmtDetails = &v
}

// GetUsedLmt returns the UsedLmt field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetUsedLmt() string {
	if o == nil || IsNil(o.UsedLmt) {
		var ret string
		return ret
	}
	return *o.UsedLmt
}

// GetUsedLmtOk returns a tuple with the UsedLmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetUsedLmtOk() (*string, bool) {
	if o == nil || IsNil(o.UsedLmt) {
		return nil, false
	}
	return o.UsedLmt, true
}

// HasUsedLmt returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasUsedLmt() bool {
	if o != nil && !IsNil(o.UsedLmt) {
		return true
	}

	return false
}

// SetUsedLmt gets a reference to the given string and assigns it to the UsedLmt field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetUsedLmt(v string) {
	o.UsedLmt = &v
}

// GetUsedLoan returns the UsedLoan field value if set, zero value otherwise.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetUsedLoan() string {
	if o == nil || IsNil(o.UsedLoan) {
		var ret string
		return ret
	}
	return *o.UsedLoan
}

// GetUsedLoanOk returns a tuple with the UsedLoan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) GetUsedLoanOk() (*string, bool) {
	if o == nil || IsNil(o.UsedLoan) {
		return nil, false
	}
	return o.UsedLoan, true
}

// HasUsedLoan returns a boolean if a field has been set.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) HasUsedLoan() bool {
	if o != nil && !IsNil(o.UsedLoan) {
		return true
	}

	return false
}

// SetUsedLoan gets a reference to the given string and assigns it to the UsedLoan field.
func (o *GetAccountInterestLimitsV5RespDataInnerRecordsInner) SetUsedLoan(v string) {
	o.UsedLoan = &v
}

func (o GetAccountInterestLimitsV5RespDataInnerRecordsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountInterestLimitsV5RespDataInnerRecordsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllAcctRemainingQuota) {
		toSerialize["allAcctRemainingQuota"] = o.AllAcctRemainingQuota
	}
	if !IsNil(o.AvailLoan) {
		toSerialize["availLoan"] = o.AvailLoan
	}
	if !IsNil(o.AvgRate) {
		toSerialize["avgRate"] = o.AvgRate
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.CurAcctRemainingQuota) {
		toSerialize["curAcctRemainingQuota"] = o.CurAcctRemainingQuota
	}
	if !IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !IsNil(o.LoanQuota) {
		toSerialize["loanQuota"] = o.LoanQuota
	}
	if !IsNil(o.PlatRemainingQuota) {
		toSerialize["platRemainingQuota"] = o.PlatRemainingQuota
	}
	if !IsNil(o.PosLoan) {
		toSerialize["posLoan"] = o.PosLoan
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.SurplusLmt) {
		toSerialize["surplusLmt"] = o.SurplusLmt
	}
	if !IsNil(o.SurplusLmtDetails) {
		toSerialize["surplusLmtDetails"] = o.SurplusLmtDetails
	}
	if !IsNil(o.UsedLmt) {
		toSerialize["usedLmt"] = o.UsedLmt
	}
	if !IsNil(o.UsedLoan) {
		toSerialize["usedLoan"] = o.UsedLoan
	}
	return toSerialize, nil
}

type NullableGetAccountInterestLimitsV5RespDataInnerRecordsInner struct {
	value *GetAccountInterestLimitsV5RespDataInnerRecordsInner
	isSet bool
}

func (v NullableGetAccountInterestLimitsV5RespDataInnerRecordsInner) Get() *GetAccountInterestLimitsV5RespDataInnerRecordsInner {
	return v.value
}

func (v *NullableGetAccountInterestLimitsV5RespDataInnerRecordsInner) Set(val *GetAccountInterestLimitsV5RespDataInnerRecordsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountInterestLimitsV5RespDataInnerRecordsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountInterestLimitsV5RespDataInnerRecordsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountInterestLimitsV5RespDataInnerRecordsInner(val *GetAccountInterestLimitsV5RespDataInnerRecordsInner) *NullableGetAccountInterestLimitsV5RespDataInnerRecordsInner {
	return &NullableGetAccountInterestLimitsV5RespDataInnerRecordsInner{value: val, isSet: true}
}

func (v NullableGetAccountInterestLimitsV5RespDataInnerRecordsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountInterestLimitsV5RespDataInnerRecordsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


