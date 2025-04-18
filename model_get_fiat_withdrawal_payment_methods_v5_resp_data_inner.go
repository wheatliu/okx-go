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

// checks if the GetFiatWithdrawalPaymentMethodsV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFiatWithdrawalPaymentMethodsV5RespDataInner{}

// GetFiatWithdrawalPaymentMethodsV5RespDataInner struct for GetFiatWithdrawalPaymentMethodsV5RespDataInner
type GetFiatWithdrawalPaymentMethodsV5RespDataInner struct {
	// An array containing information about payment accounts associated with the currency and method.
	Accounts []GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner `json:"accounts,omitempty"`
	// Fiat currency
	Ccy *string `json:"ccy,omitempty"`
	// The fee rate for each deposit, expressed as a percentage    e.g. `0.02` represents 2 percent fee for each transaction.
	FeeRate *string `json:"feeRate,omitempty"`
	Limits *GetFiatDepositPaymentMethodsV5RespDataInnerLimits `json:"limits,omitempty"`
	// The minimum fee for each deposit
	MinFee *string `json:"minFee,omitempty"`
	// The payment method associated with the currency  `TR_BANKS`  `PIX`  `SEPA`
	PaymentMethod *string `json:"paymentMethod,omitempty"`
}

// NewGetFiatWithdrawalPaymentMethodsV5RespDataInner instantiates a new GetFiatWithdrawalPaymentMethodsV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFiatWithdrawalPaymentMethodsV5RespDataInner() *GetFiatWithdrawalPaymentMethodsV5RespDataInner {
	this := GetFiatWithdrawalPaymentMethodsV5RespDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var feeRate string = ""
	this.FeeRate = &feeRate
	var minFee string = ""
	this.MinFee = &minFee
	var paymentMethod string = ""
	this.PaymentMethod = &paymentMethod
	return &this
}

// NewGetFiatWithdrawalPaymentMethodsV5RespDataInnerWithDefaults instantiates a new GetFiatWithdrawalPaymentMethodsV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFiatWithdrawalPaymentMethodsV5RespDataInnerWithDefaults() *GetFiatWithdrawalPaymentMethodsV5RespDataInner {
	this := GetFiatWithdrawalPaymentMethodsV5RespDataInner{}
	var ccy string = ""
	this.Ccy = &ccy
	var feeRate string = ""
	this.FeeRate = &feeRate
	var minFee string = ""
	this.MinFee = &minFee
	var paymentMethod string = ""
	this.PaymentMethod = &paymentMethod
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) GetAccounts() []GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner {
	if o == nil || IsNil(o.Accounts) {
		var ret []GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) GetAccountsOk() ([]GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner and assigns it to the Accounts field.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) SetAccounts(v []GetFiatDepositPaymentMethodsV5RespDataInnerAccountsInner) {
	o.Accounts = v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetFeeRate returns the FeeRate field value if set, zero value otherwise.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) GetFeeRate() string {
	if o == nil || IsNil(o.FeeRate) {
		var ret string
		return ret
	}
	return *o.FeeRate
}

// GetFeeRateOk returns a tuple with the FeeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) GetFeeRateOk() (*string, bool) {
	if o == nil || IsNil(o.FeeRate) {
		return nil, false
	}
	return o.FeeRate, true
}

// HasFeeRate returns a boolean if a field has been set.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) HasFeeRate() bool {
	if o != nil && !IsNil(o.FeeRate) {
		return true
	}

	return false
}

// SetFeeRate gets a reference to the given string and assigns it to the FeeRate field.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) SetFeeRate(v string) {
	o.FeeRate = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) GetLimits() GetFiatDepositPaymentMethodsV5RespDataInnerLimits {
	if o == nil || IsNil(o.Limits) {
		var ret GetFiatDepositPaymentMethodsV5RespDataInnerLimits
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) GetLimitsOk() (*GetFiatDepositPaymentMethodsV5RespDataInnerLimits, bool) {
	if o == nil || IsNil(o.Limits) {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) HasLimits() bool {
	if o != nil && !IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given GetFiatDepositPaymentMethodsV5RespDataInnerLimits and assigns it to the Limits field.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) SetLimits(v GetFiatDepositPaymentMethodsV5RespDataInnerLimits) {
	o.Limits = &v
}

// GetMinFee returns the MinFee field value if set, zero value otherwise.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) GetMinFee() string {
	if o == nil || IsNil(o.MinFee) {
		var ret string
		return ret
	}
	return *o.MinFee
}

// GetMinFeeOk returns a tuple with the MinFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) GetMinFeeOk() (*string, bool) {
	if o == nil || IsNil(o.MinFee) {
		return nil, false
	}
	return o.MinFee, true
}

// HasMinFee returns a boolean if a field has been set.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) HasMinFee() bool {
	if o != nil && !IsNil(o.MinFee) {
		return true
	}

	return false
}

// SetMinFee gets a reference to the given string and assigns it to the MinFee field.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) SetMinFee(v string) {
	o.MinFee = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *GetFiatWithdrawalPaymentMethodsV5RespDataInner) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

func (o GetFiatWithdrawalPaymentMethodsV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFiatWithdrawalPaymentMethodsV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.FeeRate) {
		toSerialize["feeRate"] = o.FeeRate
	}
	if !IsNil(o.Limits) {
		toSerialize["limits"] = o.Limits
	}
	if !IsNil(o.MinFee) {
		toSerialize["minFee"] = o.MinFee
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	return toSerialize, nil
}

type NullableGetFiatWithdrawalPaymentMethodsV5RespDataInner struct {
	value *GetFiatWithdrawalPaymentMethodsV5RespDataInner
	isSet bool
}

func (v NullableGetFiatWithdrawalPaymentMethodsV5RespDataInner) Get() *GetFiatWithdrawalPaymentMethodsV5RespDataInner {
	return v.value
}

func (v *NullableGetFiatWithdrawalPaymentMethodsV5RespDataInner) Set(val *GetFiatWithdrawalPaymentMethodsV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFiatWithdrawalPaymentMethodsV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFiatWithdrawalPaymentMethodsV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFiatWithdrawalPaymentMethodsV5RespDataInner(val *GetFiatWithdrawalPaymentMethodsV5RespDataInner) *NullableGetFiatWithdrawalPaymentMethodsV5RespDataInner {
	return &NullableGetFiatWithdrawalPaymentMethodsV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetFiatWithdrawalPaymentMethodsV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFiatWithdrawalPaymentMethodsV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


