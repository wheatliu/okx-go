/*
Okx Rest API

OpenAPI specification for Okx cryptocurrency exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateTradingBotGridOrderAlgoV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTradingBotGridOrderAlgoV5Req{}

// CreateTradingBotGridOrderAlgoV5Req struct for CreateTradingBotGridOrderAlgoV5Req
type CreateTradingBotGridOrderAlgoV5Req struct {
	// Client-supplied Algo ID  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters.
	AlgoClOrdId *string `json:"algoClOrdId,omitempty"`
	// Algo order type  `grid`: Spot grid  `contract_grid`: Contract grid
	AlgoOrdType string `json:"algoOrdType"`
	// Grid quantity
	GridNum string `json:"gridNum"`
	// Instrument ID, e.g. `BTC-USDT-SWAP`
	InstId string `json:"instId"`
	// Upper price of price range
	MaxPx string `json:"maxPx"`
	// Lower price of price range
	MinPx string `json:"minPx"`
	// Profit sharing ratio, it only supports these values  `0`,`0.1`,`0.2`,`0.3`   0.1 represents 10%
	ProfitSharingRatio *string `json:"profitSharingRatio,omitempty"`
	// Grid type  `1`: Arithmetic, `2`: Geometric  Default is Arithmetic
	RunType *string `json:"runType,omitempty"`
	// SL tigger price  Applicable to `Spot grid`/`Contract grid`
	SlTriggerPx *string `json:"slTriggerPx,omitempty"`
	// Order tag
	Tag *string `json:"tag,omitempty"`
	// TP tigger price  Applicable to `Spot grid`/`Contract grid`
	TpTriggerPx *string `json:"tpTriggerPx,omitempty"`
	// Trigger Parameters  Applicable to `Spot grid`/`Contract grid`
	TriggerParams []CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner `json:"triggerParams,omitempty"`
}

type _CreateTradingBotGridOrderAlgoV5Req CreateTradingBotGridOrderAlgoV5Req

// NewCreateTradingBotGridOrderAlgoV5Req instantiates a new CreateTradingBotGridOrderAlgoV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradingBotGridOrderAlgoV5Req(algoOrdType string, gridNum string, instId string, maxPx string, minPx string) *CreateTradingBotGridOrderAlgoV5Req {
	this := CreateTradingBotGridOrderAlgoV5Req{}
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	this.AlgoOrdType = algoOrdType
	this.GridNum = gridNum
	this.InstId = instId
	this.MaxPx = maxPx
	this.MinPx = minPx
	var profitSharingRatio string = ""
	this.ProfitSharingRatio = &profitSharingRatio
	var runType string = ""
	this.RunType = &runType
	var slTriggerPx string = ""
	this.SlTriggerPx = &slTriggerPx
	var tag string = ""
	this.Tag = &tag
	var tpTriggerPx string = ""
	this.TpTriggerPx = &tpTriggerPx
	return &this
}

// NewCreateTradingBotGridOrderAlgoV5ReqWithDefaults instantiates a new CreateTradingBotGridOrderAlgoV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradingBotGridOrderAlgoV5ReqWithDefaults() *CreateTradingBotGridOrderAlgoV5Req {
	this := CreateTradingBotGridOrderAlgoV5Req{}
	var algoClOrdId string = ""
	this.AlgoClOrdId = &algoClOrdId
	var algoOrdType string = ""
	this.AlgoOrdType = algoOrdType
	var gridNum string = ""
	this.GridNum = gridNum
	var instId string = ""
	this.InstId = instId
	var maxPx string = ""
	this.MaxPx = maxPx
	var minPx string = ""
	this.MinPx = minPx
	var profitSharingRatio string = ""
	this.ProfitSharingRatio = &profitSharingRatio
	var runType string = ""
	this.RunType = &runType
	var slTriggerPx string = ""
	this.SlTriggerPx = &slTriggerPx
	var tag string = ""
	this.Tag = &tag
	var tpTriggerPx string = ""
	this.TpTriggerPx = &tpTriggerPx
	return &this
}

// GetAlgoClOrdId returns the AlgoClOrdId field value if set, zero value otherwise.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetAlgoClOrdId() string {
	if o == nil || IsNil(o.AlgoClOrdId) {
		var ret string
		return ret
	}
	return *o.AlgoClOrdId
}

// GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetAlgoClOrdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlgoClOrdId) {
		return nil, false
	}
	return o.AlgoClOrdId, true
}

// HasAlgoClOrdId returns a boolean if a field has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) HasAlgoClOrdId() bool {
	if o != nil && !IsNil(o.AlgoClOrdId) {
		return true
	}

	return false
}

// SetAlgoClOrdId gets a reference to the given string and assigns it to the AlgoClOrdId field.
func (o *CreateTradingBotGridOrderAlgoV5Req) SetAlgoClOrdId(v string) {
	o.AlgoClOrdId = &v
}

// GetAlgoOrdType returns the AlgoOrdType field value
func (o *CreateTradingBotGridOrderAlgoV5Req) GetAlgoOrdType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlgoOrdType
}

// GetAlgoOrdTypeOk returns a tuple with the AlgoOrdType field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetAlgoOrdTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlgoOrdType, true
}

// SetAlgoOrdType sets field value
func (o *CreateTradingBotGridOrderAlgoV5Req) SetAlgoOrdType(v string) {
	o.AlgoOrdType = v
}

// GetGridNum returns the GridNum field value
func (o *CreateTradingBotGridOrderAlgoV5Req) GetGridNum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GridNum
}

// GetGridNumOk returns a tuple with the GridNum field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetGridNumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GridNum, true
}

// SetGridNum sets field value
func (o *CreateTradingBotGridOrderAlgoV5Req) SetGridNum(v string) {
	o.GridNum = v
}

// GetInstId returns the InstId field value
func (o *CreateTradingBotGridOrderAlgoV5Req) GetInstId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetInstIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstId, true
}

// SetInstId sets field value
func (o *CreateTradingBotGridOrderAlgoV5Req) SetInstId(v string) {
	o.InstId = v
}

// GetMaxPx returns the MaxPx field value
func (o *CreateTradingBotGridOrderAlgoV5Req) GetMaxPx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxPx
}

// GetMaxPxOk returns a tuple with the MaxPx field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetMaxPxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxPx, true
}

// SetMaxPx sets field value
func (o *CreateTradingBotGridOrderAlgoV5Req) SetMaxPx(v string) {
	o.MaxPx = v
}

// GetMinPx returns the MinPx field value
func (o *CreateTradingBotGridOrderAlgoV5Req) GetMinPx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinPx
}

// GetMinPxOk returns a tuple with the MinPx field value
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetMinPxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinPx, true
}

// SetMinPx sets field value
func (o *CreateTradingBotGridOrderAlgoV5Req) SetMinPx(v string) {
	o.MinPx = v
}

// GetProfitSharingRatio returns the ProfitSharingRatio field value if set, zero value otherwise.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetProfitSharingRatio() string {
	if o == nil || IsNil(o.ProfitSharingRatio) {
		var ret string
		return ret
	}
	return *o.ProfitSharingRatio
}

// GetProfitSharingRatioOk returns a tuple with the ProfitSharingRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetProfitSharingRatioOk() (*string, bool) {
	if o == nil || IsNil(o.ProfitSharingRatio) {
		return nil, false
	}
	return o.ProfitSharingRatio, true
}

// HasProfitSharingRatio returns a boolean if a field has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) HasProfitSharingRatio() bool {
	if o != nil && !IsNil(o.ProfitSharingRatio) {
		return true
	}

	return false
}

// SetProfitSharingRatio gets a reference to the given string and assigns it to the ProfitSharingRatio field.
func (o *CreateTradingBotGridOrderAlgoV5Req) SetProfitSharingRatio(v string) {
	o.ProfitSharingRatio = &v
}

// GetRunType returns the RunType field value if set, zero value otherwise.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetRunType() string {
	if o == nil || IsNil(o.RunType) {
		var ret string
		return ret
	}
	return *o.RunType
}

// GetRunTypeOk returns a tuple with the RunType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetRunTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RunType) {
		return nil, false
	}
	return o.RunType, true
}

// HasRunType returns a boolean if a field has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) HasRunType() bool {
	if o != nil && !IsNil(o.RunType) {
		return true
	}

	return false
}

// SetRunType gets a reference to the given string and assigns it to the RunType field.
func (o *CreateTradingBotGridOrderAlgoV5Req) SetRunType(v string) {
	o.RunType = &v
}

// GetSlTriggerPx returns the SlTriggerPx field value if set, zero value otherwise.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetSlTriggerPx() string {
	if o == nil || IsNil(o.SlTriggerPx) {
		var ret string
		return ret
	}
	return *o.SlTriggerPx
}

// GetSlTriggerPxOk returns a tuple with the SlTriggerPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetSlTriggerPxOk() (*string, bool) {
	if o == nil || IsNil(o.SlTriggerPx) {
		return nil, false
	}
	return o.SlTriggerPx, true
}

// HasSlTriggerPx returns a boolean if a field has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) HasSlTriggerPx() bool {
	if o != nil && !IsNil(o.SlTriggerPx) {
		return true
	}

	return false
}

// SetSlTriggerPx gets a reference to the given string and assigns it to the SlTriggerPx field.
func (o *CreateTradingBotGridOrderAlgoV5Req) SetSlTriggerPx(v string) {
	o.SlTriggerPx = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *CreateTradingBotGridOrderAlgoV5Req) SetTag(v string) {
	o.Tag = &v
}

// GetTpTriggerPx returns the TpTriggerPx field value if set, zero value otherwise.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetTpTriggerPx() string {
	if o == nil || IsNil(o.TpTriggerPx) {
		var ret string
		return ret
	}
	return *o.TpTriggerPx
}

// GetTpTriggerPxOk returns a tuple with the TpTriggerPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetTpTriggerPxOk() (*string, bool) {
	if o == nil || IsNil(o.TpTriggerPx) {
		return nil, false
	}
	return o.TpTriggerPx, true
}

// HasTpTriggerPx returns a boolean if a field has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) HasTpTriggerPx() bool {
	if o != nil && !IsNil(o.TpTriggerPx) {
		return true
	}

	return false
}

// SetTpTriggerPx gets a reference to the given string and assigns it to the TpTriggerPx field.
func (o *CreateTradingBotGridOrderAlgoV5Req) SetTpTriggerPx(v string) {
	o.TpTriggerPx = &v
}

// GetTriggerParams returns the TriggerParams field value if set, zero value otherwise.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetTriggerParams() []CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner {
	if o == nil || IsNil(o.TriggerParams) {
		var ret []CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner
		return ret
	}
	return o.TriggerParams
}

// GetTriggerParamsOk returns a tuple with the TriggerParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) GetTriggerParamsOk() ([]CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner, bool) {
	if o == nil || IsNil(o.TriggerParams) {
		return nil, false
	}
	return o.TriggerParams, true
}

// HasTriggerParams returns a boolean if a field has been set.
func (o *CreateTradingBotGridOrderAlgoV5Req) HasTriggerParams() bool {
	if o != nil && !IsNil(o.TriggerParams) {
		return true
	}

	return false
}

// SetTriggerParams gets a reference to the given []CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner and assigns it to the TriggerParams field.
func (o *CreateTradingBotGridOrderAlgoV5Req) SetTriggerParams(v []CreateTradingBotGridOrderAlgoV5ReqTriggerParamsInner) {
	o.TriggerParams = v
}

func (o CreateTradingBotGridOrderAlgoV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTradingBotGridOrderAlgoV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlgoClOrdId) {
		toSerialize["algoClOrdId"] = o.AlgoClOrdId
	}
	toSerialize["algoOrdType"] = o.AlgoOrdType
	toSerialize["gridNum"] = o.GridNum
	toSerialize["instId"] = o.InstId
	toSerialize["maxPx"] = o.MaxPx
	toSerialize["minPx"] = o.MinPx
	if !IsNil(o.ProfitSharingRatio) {
		toSerialize["profitSharingRatio"] = o.ProfitSharingRatio
	}
	if !IsNil(o.RunType) {
		toSerialize["runType"] = o.RunType
	}
	if !IsNil(o.SlTriggerPx) {
		toSerialize["slTriggerPx"] = o.SlTriggerPx
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TpTriggerPx) {
		toSerialize["tpTriggerPx"] = o.TpTriggerPx
	}
	if !IsNil(o.TriggerParams) {
		toSerialize["triggerParams"] = o.TriggerParams
	}
	return toSerialize, nil
}

func (o *CreateTradingBotGridOrderAlgoV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"algoOrdType",
		"gridNum",
		"instId",
		"maxPx",
		"minPx",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateTradingBotGridOrderAlgoV5Req := _CreateTradingBotGridOrderAlgoV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTradingBotGridOrderAlgoV5Req)

	if err != nil {
		return err
	}

	*o = CreateTradingBotGridOrderAlgoV5Req(varCreateTradingBotGridOrderAlgoV5Req)

	return err
}

type NullableCreateTradingBotGridOrderAlgoV5Req struct {
	value *CreateTradingBotGridOrderAlgoV5Req
	isSet bool
}

func (v NullableCreateTradingBotGridOrderAlgoV5Req) Get() *CreateTradingBotGridOrderAlgoV5Req {
	return v.value
}

func (v *NullableCreateTradingBotGridOrderAlgoV5Req) Set(val *CreateTradingBotGridOrderAlgoV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradingBotGridOrderAlgoV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradingBotGridOrderAlgoV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradingBotGridOrderAlgoV5Req(val *CreateTradingBotGridOrderAlgoV5Req) *NullableCreateTradingBotGridOrderAlgoV5Req {
	return &NullableCreateTradingBotGridOrderAlgoV5Req{value: val, isSet: true}
}

func (v NullableCreateTradingBotGridOrderAlgoV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradingBotGridOrderAlgoV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


