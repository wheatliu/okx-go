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

// checks if the GetFinanceStakingDefiOrdersActiveV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFinanceStakingDefiOrdersActiveV5RespData{}

// GetFinanceStakingDefiOrdersActiveV5RespData struct for GetFinanceStakingDefiOrdersActiveV5RespData
type GetFinanceStakingDefiOrdersActiveV5RespData struct {
	// Estimated APY  If the estimated APY is 7% , this field is 0.07  Retain to 4 decimal places (truncated)
	Apy *string `json:"apy,omitempty"`
	// Deadline for cancellation of redemption application
	CancelRedemptionDeadline *string `json:"cancelRedemptionDeadline,omitempty"`
	// Currency, e.g. `BTC`
	Ccy *string `json:"ccy,omitempty"`
	// Earning data
	EarningData []GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner `json:"earningData,omitempty"`
	// Estimated redemption settlement time
	EstSettlementTime *string `json:"estSettlementTime,omitempty"`
	// Fast redemption data
	FastRedemptionData []GetFinanceStakingDefiOrdersActiveV5RespDataFastRedemptionDataInner `json:"fastRedemptionData,omitempty"`
	// Investment data
	InvestData []GetFinanceStakingDefiOrdersActiveV5RespDataInvestDataInner `json:"investData,omitempty"`
	// Order ID
	OrdId *string `json:"ordId,omitempty"`
	// Product ID
	ProductId *string `json:"productId,omitempty"`
	// Protocol
	Protocol *string `json:"protocol,omitempty"`
	// Protocol type  `defi`: on-chain earn
	ProtocolType *string `json:"protocolType,omitempty"`
	// Order purchased time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	PurchasedTime *string `json:"purchasedTime,omitempty"`
	// Order state  8: Pending   13: Cancelling   9: Onchain   1: Earning   2: Redeeming
	State *string `json:"state,omitempty"`
	// Order tag
	Tag *string `json:"tag,omitempty"`
	// Protocol term  It will return the days of fixed term and will return `0` for flexible product
	Term *string `json:"term,omitempty"`
}

// NewGetFinanceStakingDefiOrdersActiveV5RespData instantiates a new GetFinanceStakingDefiOrdersActiveV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFinanceStakingDefiOrdersActiveV5RespData() *GetFinanceStakingDefiOrdersActiveV5RespData {
	this := GetFinanceStakingDefiOrdersActiveV5RespData{}
	var apy string = ""
	this.Apy = &apy
	var cancelRedemptionDeadline string = ""
	this.CancelRedemptionDeadline = &cancelRedemptionDeadline
	var ccy string = ""
	this.Ccy = &ccy
	var estSettlementTime string = ""
	this.EstSettlementTime = &estSettlementTime
	var ordId string = ""
	this.OrdId = &ordId
	var productId string = ""
	this.ProductId = &productId
	var protocol string = ""
	this.Protocol = &protocol
	var protocolType string = ""
	this.ProtocolType = &protocolType
	var purchasedTime string = ""
	this.PurchasedTime = &purchasedTime
	var state string = ""
	this.State = &state
	var tag string = ""
	this.Tag = &tag
	var term string = ""
	this.Term = &term
	return &this
}

// NewGetFinanceStakingDefiOrdersActiveV5RespDataWithDefaults instantiates a new GetFinanceStakingDefiOrdersActiveV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFinanceStakingDefiOrdersActiveV5RespDataWithDefaults() *GetFinanceStakingDefiOrdersActiveV5RespData {
	this := GetFinanceStakingDefiOrdersActiveV5RespData{}
	var apy string = ""
	this.Apy = &apy
	var cancelRedemptionDeadline string = ""
	this.CancelRedemptionDeadline = &cancelRedemptionDeadline
	var ccy string = ""
	this.Ccy = &ccy
	var estSettlementTime string = ""
	this.EstSettlementTime = &estSettlementTime
	var ordId string = ""
	this.OrdId = &ordId
	var productId string = ""
	this.ProductId = &productId
	var protocol string = ""
	this.Protocol = &protocol
	var protocolType string = ""
	this.ProtocolType = &protocolType
	var purchasedTime string = ""
	this.PurchasedTime = &purchasedTime
	var state string = ""
	this.State = &state
	var tag string = ""
	this.Tag = &tag
	var term string = ""
	this.Term = &term
	return &this
}

// GetApy returns the Apy field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetApy() string {
	if o == nil || IsNil(o.Apy) {
		var ret string
		return ret
	}
	return *o.Apy
}

// GetApyOk returns a tuple with the Apy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetApyOk() (*string, bool) {
	if o == nil || IsNil(o.Apy) {
		return nil, false
	}
	return o.Apy, true
}

// HasApy returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasApy() bool {
	if o != nil && !IsNil(o.Apy) {
		return true
	}

	return false
}

// SetApy gets a reference to the given string and assigns it to the Apy field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetApy(v string) {
	o.Apy = &v
}

// GetCancelRedemptionDeadline returns the CancelRedemptionDeadline field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetCancelRedemptionDeadline() string {
	if o == nil || IsNil(o.CancelRedemptionDeadline) {
		var ret string
		return ret
	}
	return *o.CancelRedemptionDeadline
}

// GetCancelRedemptionDeadlineOk returns a tuple with the CancelRedemptionDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetCancelRedemptionDeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.CancelRedemptionDeadline) {
		return nil, false
	}
	return o.CancelRedemptionDeadline, true
}

// HasCancelRedemptionDeadline returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasCancelRedemptionDeadline() bool {
	if o != nil && !IsNil(o.CancelRedemptionDeadline) {
		return true
	}

	return false
}

// SetCancelRedemptionDeadline gets a reference to the given string and assigns it to the CancelRedemptionDeadline field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetCancelRedemptionDeadline(v string) {
	o.CancelRedemptionDeadline = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetCcy(v string) {
	o.Ccy = &v
}

// GetEarningData returns the EarningData field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetEarningData() []GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner {
	if o == nil || IsNil(o.EarningData) {
		var ret []GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner
		return ret
	}
	return o.EarningData
}

// GetEarningDataOk returns a tuple with the EarningData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetEarningDataOk() ([]GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner, bool) {
	if o == nil || IsNil(o.EarningData) {
		return nil, false
	}
	return o.EarningData, true
}

// HasEarningData returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasEarningData() bool {
	if o != nil && !IsNil(o.EarningData) {
		return true
	}

	return false
}

// SetEarningData gets a reference to the given []GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner and assigns it to the EarningData field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetEarningData(v []GetFinanceStakingDefiOrdersActiveV5RespDataEarningDataInner) {
	o.EarningData = v
}

// GetEstSettlementTime returns the EstSettlementTime field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetEstSettlementTime() string {
	if o == nil || IsNil(o.EstSettlementTime) {
		var ret string
		return ret
	}
	return *o.EstSettlementTime
}

// GetEstSettlementTimeOk returns a tuple with the EstSettlementTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetEstSettlementTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EstSettlementTime) {
		return nil, false
	}
	return o.EstSettlementTime, true
}

// HasEstSettlementTime returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasEstSettlementTime() bool {
	if o != nil && !IsNil(o.EstSettlementTime) {
		return true
	}

	return false
}

// SetEstSettlementTime gets a reference to the given string and assigns it to the EstSettlementTime field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetEstSettlementTime(v string) {
	o.EstSettlementTime = &v
}

// GetFastRedemptionData returns the FastRedemptionData field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetFastRedemptionData() []GetFinanceStakingDefiOrdersActiveV5RespDataFastRedemptionDataInner {
	if o == nil || IsNil(o.FastRedemptionData) {
		var ret []GetFinanceStakingDefiOrdersActiveV5RespDataFastRedemptionDataInner
		return ret
	}
	return o.FastRedemptionData
}

// GetFastRedemptionDataOk returns a tuple with the FastRedemptionData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetFastRedemptionDataOk() ([]GetFinanceStakingDefiOrdersActiveV5RespDataFastRedemptionDataInner, bool) {
	if o == nil || IsNil(o.FastRedemptionData) {
		return nil, false
	}
	return o.FastRedemptionData, true
}

// HasFastRedemptionData returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasFastRedemptionData() bool {
	if o != nil && !IsNil(o.FastRedemptionData) {
		return true
	}

	return false
}

// SetFastRedemptionData gets a reference to the given []GetFinanceStakingDefiOrdersActiveV5RespDataFastRedemptionDataInner and assigns it to the FastRedemptionData field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetFastRedemptionData(v []GetFinanceStakingDefiOrdersActiveV5RespDataFastRedemptionDataInner) {
	o.FastRedemptionData = v
}

// GetInvestData returns the InvestData field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetInvestData() []GetFinanceStakingDefiOrdersActiveV5RespDataInvestDataInner {
	if o == nil || IsNil(o.InvestData) {
		var ret []GetFinanceStakingDefiOrdersActiveV5RespDataInvestDataInner
		return ret
	}
	return o.InvestData
}

// GetInvestDataOk returns a tuple with the InvestData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetInvestDataOk() ([]GetFinanceStakingDefiOrdersActiveV5RespDataInvestDataInner, bool) {
	if o == nil || IsNil(o.InvestData) {
		return nil, false
	}
	return o.InvestData, true
}

// HasInvestData returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasInvestData() bool {
	if o != nil && !IsNil(o.InvestData) {
		return true
	}

	return false
}

// SetInvestData gets a reference to the given []GetFinanceStakingDefiOrdersActiveV5RespDataInvestDataInner and assigns it to the InvestData field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetInvestData(v []GetFinanceStakingDefiOrdersActiveV5RespDataInvestDataInner) {
	o.InvestData = v
}

// GetOrdId returns the OrdId field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetOrdId() string {
	if o == nil || IsNil(o.OrdId) {
		var ret string
		return ret
	}
	return *o.OrdId
}

// GetOrdIdOk returns a tuple with the OrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrdId) {
		return nil, false
	}
	return o.OrdId, true
}

// HasOrdId returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasOrdId() bool {
	if o != nil && !IsNil(o.OrdId) {
		return true
	}

	return false
}

// SetOrdId gets a reference to the given string and assigns it to the OrdId field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetOrdId(v string) {
	o.OrdId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetProductId(v string) {
	o.ProductId = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetProtocol(v string) {
	o.Protocol = &v
}

// GetProtocolType returns the ProtocolType field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetProtocolType() string {
	if o == nil || IsNil(o.ProtocolType) {
		var ret string
		return ret
	}
	return *o.ProtocolType
}

// GetProtocolTypeOk returns a tuple with the ProtocolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetProtocolTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolType) {
		return nil, false
	}
	return o.ProtocolType, true
}

// HasProtocolType returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasProtocolType() bool {
	if o != nil && !IsNil(o.ProtocolType) {
		return true
	}

	return false
}

// SetProtocolType gets a reference to the given string and assigns it to the ProtocolType field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetProtocolType(v string) {
	o.ProtocolType = &v
}

// GetPurchasedTime returns the PurchasedTime field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetPurchasedTime() string {
	if o == nil || IsNil(o.PurchasedTime) {
		var ret string
		return ret
	}
	return *o.PurchasedTime
}

// GetPurchasedTimeOk returns a tuple with the PurchasedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetPurchasedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PurchasedTime) {
		return nil, false
	}
	return o.PurchasedTime, true
}

// HasPurchasedTime returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasPurchasedTime() bool {
	if o != nil && !IsNil(o.PurchasedTime) {
		return true
	}

	return false
}

// SetPurchasedTime gets a reference to the given string and assigns it to the PurchasedTime field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetPurchasedTime(v string) {
	o.PurchasedTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetState(v string) {
	o.State = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetTag(v string) {
	o.Tag = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetTerm() string {
	if o == nil || IsNil(o.Term) {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) GetTermOk() (*string, bool) {
	if o == nil || IsNil(o.Term) {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) HasTerm() bool {
	if o != nil && !IsNil(o.Term) {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *GetFinanceStakingDefiOrdersActiveV5RespData) SetTerm(v string) {
	o.Term = &v
}

func (o GetFinanceStakingDefiOrdersActiveV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFinanceStakingDefiOrdersActiveV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Apy) {
		toSerialize["apy"] = o.Apy
	}
	if !IsNil(o.CancelRedemptionDeadline) {
		toSerialize["cancelRedemptionDeadline"] = o.CancelRedemptionDeadline
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.EarningData) {
		toSerialize["earningData"] = o.EarningData
	}
	if !IsNil(o.EstSettlementTime) {
		toSerialize["estSettlementTime"] = o.EstSettlementTime
	}
	if !IsNil(o.FastRedemptionData) {
		toSerialize["fastRedemptionData"] = o.FastRedemptionData
	}
	if !IsNil(o.InvestData) {
		toSerialize["investData"] = o.InvestData
	}
	if !IsNil(o.OrdId) {
		toSerialize["ordId"] = o.OrdId
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
	if !IsNil(o.PurchasedTime) {
		toSerialize["purchasedTime"] = o.PurchasedTime
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Term) {
		toSerialize["term"] = o.Term
	}
	return toSerialize, nil
}

type NullableGetFinanceStakingDefiOrdersActiveV5RespData struct {
	value *GetFinanceStakingDefiOrdersActiveV5RespData
	isSet bool
}

func (v NullableGetFinanceStakingDefiOrdersActiveV5RespData) Get() *GetFinanceStakingDefiOrdersActiveV5RespData {
	return v.value
}

func (v *NullableGetFinanceStakingDefiOrdersActiveV5RespData) Set(val *GetFinanceStakingDefiOrdersActiveV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFinanceStakingDefiOrdersActiveV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFinanceStakingDefiOrdersActiveV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFinanceStakingDefiOrdersActiveV5RespData(val *GetFinanceStakingDefiOrdersActiveV5RespData) *NullableGetFinanceStakingDefiOrdersActiveV5RespData {
	return &NullableGetFinanceStakingDefiOrdersActiveV5RespData{value: val, isSet: true}
}

func (v NullableGetFinanceStakingDefiOrdersActiveV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFinanceStakingDefiOrdersActiveV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


