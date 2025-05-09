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

// checks if the GetFinanceStakingDefiOffersV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceStakingDefiOffersV5RespDataInner{}

// GetFinanceStakingDefiOffersV5RespDataInner struct for GetFinanceStakingDefiOffersV5RespDataInner
type GetFinanceStakingDefiOffersV5RespDataInner struct {
	// Estimated annualization  If the annualization is 7% , this field is 0.07
	Apy *string `json:"apy,omitempty"`
	// Currency type, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Whether the protocol supports early redemption
	EarlyRedeem *bool `json:"earlyRedeem,omitempty"`
	// Earning data
	EarningData []GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner `json:"earningData,omitempty"`
	// Fast redemption daily limit  If fast redemption is not supported, it will return ''.
	FastRedemptionDailyLimit *string `json:"fastRedemptionDailyLimit,omitempty"`
	// Current target currency information available for investment
	InvestData []GetFinanceStakingDefiOffersV5RespDataInnerInvestDataInner `json:"investData,omitempty"`
	// Product ID
	ProductId *string `json:"productId,omitempty"`
	// Protocol
	Protocol *string `json:"protocol,omitempty"`
	// Protocol type  `defi`: on-chain earn
	ProtocolType *string `json:"protocolType,omitempty"`
	// Redemption Period, format in [min time,max time]  `H`: Hour, `D`: Day  e.g. [\"1H\",\"24H\"] represents redemption period is between 1 Hour and 24 Hours.  [\"14D\",\"14D\"] represents redemption period is 14 days.
	RedeemPeriod []string `json:"redeemPeriod,omitempty"`
	// Product state  `purchasable`: Purchasable  `sold_out`: Sold out  `Stop`: Suspension of subscription
	State *string `json:"state,omitempty"`
	// Protocol term  It will return the days of fixed term and will return `0` for flexible product
	Term *string `json:"term,omitempty"`
}

// NewGetFinanceStakingDefiOffersV5RespDataInner instantiates a new GetFinanceStakingDefiOffersV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceStakingDefiOffersV5RespDataInner() *GetFinanceStakingDefiOffersV5RespDataInner {
	this := GetFinanceStakingDefiOffersV5RespDataInner{}
	var apy string = ""
	this.Apy = &apy
	var ccy string = ""
	this.Ccy = &ccy
	var fastRedemptionDailyLimit string = ""
	this.FastRedemptionDailyLimit = &fastRedemptionDailyLimit
	var productId string = ""
	this.ProductId = &productId
	var protocol string = ""
	this.Protocol = &protocol
	var protocolType string = ""
	this.ProtocolType = &protocolType
	var state string = ""
	this.State = &state
	var term string = ""
	this.Term = &term
	return &this
}

// NewGetFinanceStakingDefiOffersV5RespDataInnerWithDefaults instantiates a new GetFinanceStakingDefiOffersV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceStakingDefiOffersV5RespDataInnerWithDefaults() *GetFinanceStakingDefiOffersV5RespDataInner {
	this := GetFinanceStakingDefiOffersV5RespDataInner{}
	var apy string = ""
	this.Apy = &apy
	var ccy string = ""
	this.Ccy = &ccy
	var fastRedemptionDailyLimit string = ""
	this.FastRedemptionDailyLimit = &fastRedemptionDailyLimit
	var productId string = ""
	this.ProductId = &productId
	var protocol string = ""
	this.Protocol = &protocol
	var protocolType string = ""
	this.ProtocolType = &protocolType
	var state string = ""
	this.State = &state
	var term string = ""
	this.Term = &term
	return &this
}

// GetApy returns the Apy field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetApy() string {
	if o == nil || IsNil(o.Apy) {
		var ret string
		return ret
	}
	return *o.Apy
}

// GetApyOk returns a tuple with the Apy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetApyOk() (*string, bool) {
	if o == nil || IsNil(o.Apy) {
		return nil, false
	}
	return o.Apy, true
}

// HasApy returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasApy() bool {
	if o != nil && !IsNil(o.Apy) {
		return true
	}

	return false
}

// SetApy gets a reference to the given string and assigns it to the Apy field.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetApy(v string) {
	o.Apy = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetEarlyRedeem returns the EarlyRedeem field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetEarlyRedeem() bool {
	if o == nil || IsNil(o.EarlyRedeem) {
		var ret bool
		return ret
	}
	return *o.EarlyRedeem
}

// GetEarlyRedeemOk returns a tuple with the EarlyRedeem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetEarlyRedeemOk() (*bool, bool) {
	if o == nil || IsNil(o.EarlyRedeem) {
		return nil, false
	}
	return o.EarlyRedeem, true
}

// HasEarlyRedeem returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasEarlyRedeem() bool {
	if o != nil && !IsNil(o.EarlyRedeem) {
		return true
	}

	return false
}

// SetEarlyRedeem gets a reference to the given bool and assigns it to the EarlyRedeem field.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetEarlyRedeem(v bool) {
	o.EarlyRedeem = &v
}

// GetEarningData returns the EarningData field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetEarningData() []GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner {
	if o == nil || IsNil(o.EarningData) {
		var ret []GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner
		return ret
	}
	return o.EarningData
}

// GetEarningDataOk returns a tuple with the EarningData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetEarningDataOk() ([]GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner, bool) {
	if o == nil || IsNil(o.EarningData) {
		return nil, false
	}
	return o.EarningData, true
}

// HasEarningData returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasEarningData() bool {
	if o != nil && !IsNil(o.EarningData) {
		return true
	}

	return false
}

// SetEarningData gets a reference to the given []GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner and assigns it to the EarningData field.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetEarningData(v []GetFinanceStakingDefiOffersV5RespDataInnerEarningDataInner) {
	o.EarningData = v
}

// GetFastRedemptionDailyLimit returns the FastRedemptionDailyLimit field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetFastRedemptionDailyLimit() string {
	if o == nil || IsNil(o.FastRedemptionDailyLimit) {
		var ret string
		return ret
	}
	return *o.FastRedemptionDailyLimit
}

// GetFastRedemptionDailyLimitOk returns a tuple with the FastRedemptionDailyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetFastRedemptionDailyLimitOk() (*string, bool) {
	if o == nil || IsNil(o.FastRedemptionDailyLimit) {
		return nil, false
	}
	return o.FastRedemptionDailyLimit, true
}

// HasFastRedemptionDailyLimit returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasFastRedemptionDailyLimit() bool {
	if o != nil && !IsNil(o.FastRedemptionDailyLimit) {
		return true
	}

	return false
}

// SetFastRedemptionDailyLimit gets a reference to the given string and assigns it to the FastRedemptionDailyLimit field.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetFastRedemptionDailyLimit(v string) {
	o.FastRedemptionDailyLimit = &v
}

// GetInvestData returns the InvestData field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetInvestData() []GetFinanceStakingDefiOffersV5RespDataInnerInvestDataInner {
	if o == nil || IsNil(o.InvestData) {
		var ret []GetFinanceStakingDefiOffersV5RespDataInnerInvestDataInner
		return ret
	}
	return o.InvestData
}

// GetInvestDataOk returns a tuple with the InvestData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetInvestDataOk() ([]GetFinanceStakingDefiOffersV5RespDataInnerInvestDataInner, bool) {
	if o == nil || IsNil(o.InvestData) {
		return nil, false
	}
	return o.InvestData, true
}

// HasInvestData returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasInvestData() bool {
	if o != nil && !IsNil(o.InvestData) {
		return true
	}

	return false
}

// SetInvestData gets a reference to the given []GetFinanceStakingDefiOffersV5RespDataInnerInvestDataInner and assigns it to the InvestData field.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetInvestData(v []GetFinanceStakingDefiOffersV5RespDataInnerInvestDataInner) {
	o.InvestData = v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetProductId(v string) {
	o.ProductId = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetProtocolType returns the ProtocolType field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetProtocolType() string {
	if o == nil || IsNil(o.ProtocolType) {
		var ret string
		return ret
	}
	return *o.ProtocolType
}

// GetProtocolTypeOk returns a tuple with the ProtocolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetProtocolTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolType) {
		return nil, false
	}
	return o.ProtocolType, true
}

// HasProtocolType returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasProtocolType() bool {
	if o != nil && !IsNil(o.ProtocolType) {
		return true
	}

	return false
}

// SetProtocolType gets a reference to the given string and assigns it to the ProtocolType field.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetProtocolType(v string) {
	o.ProtocolType = &v
}

// GetRedeemPeriod returns the RedeemPeriod field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetRedeemPeriod() []string {
	if o == nil || IsNil(o.RedeemPeriod) {
		var ret []string
		return ret
	}
	return o.RedeemPeriod
}

// GetRedeemPeriodOk returns a tuple with the RedeemPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetRedeemPeriodOk() ([]string, bool) {
	if o == nil || IsNil(o.RedeemPeriod) {
		return nil, false
	}
	return o.RedeemPeriod, true
}

// HasRedeemPeriod returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasRedeemPeriod() bool {
	if o != nil && !IsNil(o.RedeemPeriod) {
		return true
	}

	return false
}

// SetRedeemPeriod gets a reference to the given []string and assigns it to the RedeemPeriod field.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetRedeemPeriod(v []string) {
	o.RedeemPeriod = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetState(v string) {
	o.State = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetTerm() string {
	if o == nil || IsNil(o.Term) {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) GetTermOk() (*string, bool) {
	if o == nil || IsNil(o.Term) {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) HasTerm() bool {
	if o != nil && !IsNil(o.Term) {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *GetFinanceStakingDefiOffersV5RespDataInner) SetTerm(v string) {
	o.Term = &v
}

func (o GetFinanceStakingDefiOffersV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceStakingDefiOffersV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apy) {
		toSerialize["apy"] = o.Apy
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.EarlyRedeem) {
		toSerialize["earlyRedeem"] = o.EarlyRedeem
	}
	if !IsNil(o.EarningData) {
		toSerialize["earningData"] = o.EarningData
	}
	if !IsNil(o.FastRedemptionDailyLimit) {
		toSerialize["fastRedemptionDailyLimit"] = o.FastRedemptionDailyLimit
	}
	if !IsNil(o.InvestData) {
		toSerialize["investData"] = o.InvestData
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.ProtocolType) {
		toSerialize["protocolType"] = o.ProtocolType
	}
	if !IsNil(o.RedeemPeriod) {
		toSerialize["redeemPeriod"] = o.RedeemPeriod
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Term) {
		toSerialize["term"] = o.Term
	}
	return toSerialize, nil
}

type NullableGetFinanceStakingDefiOffersV5RespDataInner struct {
	value *GetFinanceStakingDefiOffersV5RespDataInner
	isSet bool
}

func (v NullableGetFinanceStakingDefiOffersV5RespDataInner) Get() *GetFinanceStakingDefiOffersV5RespDataInner {
	return v.value
}

func (v *NullableGetFinanceStakingDefiOffersV5RespDataInner) Set(val *GetFinanceStakingDefiOffersV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceStakingDefiOffersV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceStakingDefiOffersV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceStakingDefiOffersV5RespDataInner(val *GetFinanceStakingDefiOffersV5RespDataInner) *NullableGetFinanceStakingDefiOffersV5RespDataInner {
	return &NullableGetFinanceStakingDefiOffersV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetFinanceStakingDefiOffersV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceStakingDefiOffersV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


