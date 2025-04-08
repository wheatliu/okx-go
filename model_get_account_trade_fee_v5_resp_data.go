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

// checks if the GetAccountTradeFeeV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountTradeFeeV5RespData{}

// GetAccountTradeFeeV5RespData struct for GetAccountTradeFeeV5RespData
type GetAccountTradeFeeV5RespData struct {
	// Currency category. Note: this parameter is already deprecated
	// Deprecated
	Category *string `json:"category,omitempty"`
	// Delivery fee rate
	Delivery *string `json:"delivery,omitempty"`
	// Fee rate for exercising the option
	Exercise *string `json:"exercise,omitempty"`
	// Details of fiat fee rate
	Fiat []GetAccountTradeFeeV5RespDataFiatInner `json:"fiat,omitempty"`
	// Instrument type
	InstType *string `json:"instType,omitempty"`
	// Fee rate Level
	Level *string `json:"level,omitempty"`
	// For `SPOT`/`MARGIN`, it is maker fee rate of the USDT trading pairs.   For `FUTURES`/`SWAP`/`OPTION`, it is the fee rate of crypto-margined contracts
	Maker *string `json:"maker,omitempty"`
	// Maker fee rate of USDT-margined contracts, only applicable to `FUTURES`/`SWAP`
	MakerU *string `json:"makerU,omitempty"`
	// For `SPOT`/`MARGIN`, it is maker fee rate of the USDⓈ&Crypto trading pairs.  For `FUTURES`/`SWAP`, it is the fee rate of USDC-margined contracts
	MakerUSDC *string `json:"makerUSDC,omitempty"`
	// Trading rule types   `normal`: normal trading   `pre_market`: pre-market trading
	RuleType *string `json:"ruleType,omitempty"`
	// For `SPOT`/`MARGIN`, it is taker fee rate of the USDT trading pairs.   For `FUTURES`/`SWAP`/`OPTION`, it is the fee rate of crypto-margined contracts
	Taker *string `json:"taker,omitempty"`
	// Taker fee rate of USDT-margined contracts, only applicable to `FUTURES`/`SWAP`
	TakerU *string `json:"takerU,omitempty"`
	// For `SPOT`/`MARGIN`, it is taker fee rate of the USDⓈ&Crypto trading pairs.  For `FUTURES`/`SWAP`, it is the fee rate of USDC-margined contracts
	TakerUSDC *string `json:"takerUSDC,omitempty"`
	// Data return time, Unix timestamp format in milliseconds, e.g. `1597026383085`
	Ts *string `json:"ts,omitempty"`
}

// NewGetAccountTradeFeeV5RespData instantiates a new GetAccountTradeFeeV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountTradeFeeV5RespData() *GetAccountTradeFeeV5RespData {
	this := GetAccountTradeFeeV5RespData{}
	var category string = ""
	this.Category = &category
	var delivery string = ""
	this.Delivery = &delivery
	var exercise string = ""
	this.Exercise = &exercise
	var instType string = ""
	this.InstType = &instType
	var level string = ""
	this.Level = &level
	var maker string = ""
	this.Maker = &maker
	var makerU string = ""
	this.MakerU = &makerU
	var makerUSDC string = ""
	this.MakerUSDC = &makerUSDC
	var ruleType string = ""
	this.RuleType = &ruleType
	var taker string = ""
	this.Taker = &taker
	var takerU string = ""
	this.TakerU = &takerU
	var takerUSDC string = ""
	this.TakerUSDC = &takerUSDC
	var ts string = ""
	this.Ts = &ts
	return &this
}

// NewGetAccountTradeFeeV5RespDataWithDefaults instantiates a new GetAccountTradeFeeV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountTradeFeeV5RespDataWithDefaults() *GetAccountTradeFeeV5RespData {
	this := GetAccountTradeFeeV5RespData{}
	var category string = ""
	this.Category = &category
	var delivery string = ""
	this.Delivery = &delivery
	var exercise string = ""
	this.Exercise = &exercise
	var instType string = ""
	this.InstType = &instType
	var level string = ""
	this.Level = &level
	var maker string = ""
	this.Maker = &maker
	var makerU string = ""
	this.MakerU = &makerU
	var makerUSDC string = ""
	this.MakerUSDC = &makerUSDC
	var ruleType string = ""
	this.RuleType = &ruleType
	var taker string = ""
	this.Taker = &taker
	var takerU string = ""
	this.TakerU = &takerU
	var takerUSDC string = ""
	this.TakerUSDC = &takerUSDC
	var ts string = ""
	this.Ts = &ts
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
// Deprecated
func (o *GetAccountTradeFeeV5RespData) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GetAccountTradeFeeV5RespData) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
// Deprecated
func (o *GetAccountTradeFeeV5RespData) SetCategory(v string) {
	o.Category = &v
}

// GetDelivery returns the Delivery field value if set, zero value otherwise.
func (o *GetAccountTradeFeeV5RespData) GetDelivery() string {
	if o == nil || IsNil(o.Delivery) {
		var ret string
		return ret
	}
	return *o.Delivery
}

// GetDeliveryOk returns a tuple with the Delivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountTradeFeeV5RespData) GetDeliveryOk() (*string, bool) {
	if o == nil || IsNil(o.Delivery) {
		return nil, false
	}
	return o.Delivery, true
}

// HasDelivery returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasDelivery() bool {
	if o != nil && !IsNil(o.Delivery) {
		return true
	}

	return false
}

// SetDelivery gets a reference to the given string and assigns it to the Delivery field.
func (o *GetAccountTradeFeeV5RespData) SetDelivery(v string) {
	o.Delivery = &v
}

// GetExercise returns the Exercise field value if set, zero value otherwise.
func (o *GetAccountTradeFeeV5RespData) GetExercise() string {
	if o == nil || IsNil(o.Exercise) {
		var ret string
		return ret
	}
	return *o.Exercise
}

// GetExerciseOk returns a tuple with the Exercise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountTradeFeeV5RespData) GetExerciseOk() (*string, bool) {
	if o == nil || IsNil(o.Exercise) {
		return nil, false
	}
	return o.Exercise, true
}

// HasExercise returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasExercise() bool {
	if o != nil && !IsNil(o.Exercise) {
		return true
	}

	return false
}

// SetExercise gets a reference to the given string and assigns it to the Exercise field.
func (o *GetAccountTradeFeeV5RespData) SetExercise(v string) {
	o.Exercise = &v
}

// GetFiat returns the Fiat field value if set, zero value otherwise.
func (o *GetAccountTradeFeeV5RespData) GetFiat() []GetAccountTradeFeeV5RespDataFiatInner {
	if o == nil || IsNil(o.Fiat) {
		var ret []GetAccountTradeFeeV5RespDataFiatInner
		return ret
	}
	return o.Fiat
}

// GetFiatOk returns a tuple with the Fiat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountTradeFeeV5RespData) GetFiatOk() ([]GetAccountTradeFeeV5RespDataFiatInner, bool) {
	if o == nil || IsNil(o.Fiat) {
		return nil, false
	}
	return o.Fiat, true
}

// HasFiat returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasFiat() bool {
	if o != nil && !IsNil(o.Fiat) {
		return true
	}

	return false
}

// SetFiat gets a reference to the given []GetAccountTradeFeeV5RespDataFiatInner and assigns it to the Fiat field.
func (o *GetAccountTradeFeeV5RespData) SetFiat(v []GetAccountTradeFeeV5RespDataFiatInner) {
	o.Fiat = v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *GetAccountTradeFeeV5RespData) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountTradeFeeV5RespData) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *GetAccountTradeFeeV5RespData) SetInstType(v string) {
	o.InstType = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *GetAccountTradeFeeV5RespData) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountTradeFeeV5RespData) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *GetAccountTradeFeeV5RespData) SetLevel(v string) {
	o.Level = &v
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *GetAccountTradeFeeV5RespData) GetMaker() string {
	if o == nil || IsNil(o.Maker) {
		var ret string
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountTradeFeeV5RespData) GetMakerOk() (*string, bool) {
	if o == nil || IsNil(o.Maker) {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasMaker() bool {
	if o != nil && !IsNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given string and assigns it to the Maker field.
func (o *GetAccountTradeFeeV5RespData) SetMaker(v string) {
	o.Maker = &v
}

// GetMakerU returns the MakerU field value if set, zero value otherwise.
func (o *GetAccountTradeFeeV5RespData) GetMakerU() string {
	if o == nil || IsNil(o.MakerU) {
		var ret string
		return ret
	}
	return *o.MakerU
}

// GetMakerUOk returns a tuple with the MakerU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountTradeFeeV5RespData) GetMakerUOk() (*string, bool) {
	if o == nil || IsNil(o.MakerU) {
		return nil, false
	}
	return o.MakerU, true
}

// HasMakerU returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasMakerU() bool {
	if o != nil && !IsNil(o.MakerU) {
		return true
	}

	return false
}

// SetMakerU gets a reference to the given string and assigns it to the MakerU field.
func (o *GetAccountTradeFeeV5RespData) SetMakerU(v string) {
	o.MakerU = &v
}

// GetMakerUSDC returns the MakerUSDC field value if set, zero value otherwise.
func (o *GetAccountTradeFeeV5RespData) GetMakerUSDC() string {
	if o == nil || IsNil(o.MakerUSDC) {
		var ret string
		return ret
	}
	return *o.MakerUSDC
}

// GetMakerUSDCOk returns a tuple with the MakerUSDC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountTradeFeeV5RespData) GetMakerUSDCOk() (*string, bool) {
	if o == nil || IsNil(o.MakerUSDC) {
		return nil, false
	}
	return o.MakerUSDC, true
}

// HasMakerUSDC returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasMakerUSDC() bool {
	if o != nil && !IsNil(o.MakerUSDC) {
		return true
	}

	return false
}

// SetMakerUSDC gets a reference to the given string and assigns it to the MakerUSDC field.
func (o *GetAccountTradeFeeV5RespData) SetMakerUSDC(v string) {
	o.MakerUSDC = &v
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *GetAccountTradeFeeV5RespData) GetRuleType() string {
	if o == nil || IsNil(o.RuleType) {
		var ret string
		return ret
	}
	return *o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountTradeFeeV5RespData) GetRuleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RuleType) {
		return nil, false
	}
	return o.RuleType, true
}

// HasRuleType returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasRuleType() bool {
	if o != nil && !IsNil(o.RuleType) {
		return true
	}

	return false
}

// SetRuleType gets a reference to the given string and assigns it to the RuleType field.
func (o *GetAccountTradeFeeV5RespData) SetRuleType(v string) {
	o.RuleType = &v
}

// GetTaker returns the Taker field value if set, zero value otherwise.
func (o *GetAccountTradeFeeV5RespData) GetTaker() string {
	if o == nil || IsNil(o.Taker) {
		var ret string
		return ret
	}
	return *o.Taker
}

// GetTakerOk returns a tuple with the Taker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountTradeFeeV5RespData) GetTakerOk() (*string, bool) {
	if o == nil || IsNil(o.Taker) {
		return nil, false
	}
	return o.Taker, true
}

// HasTaker returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasTaker() bool {
	if o != nil && !IsNil(o.Taker) {
		return true
	}

	return false
}

// SetTaker gets a reference to the given string and assigns it to the Taker field.
func (o *GetAccountTradeFeeV5RespData) SetTaker(v string) {
	o.Taker = &v
}

// GetTakerU returns the TakerU field value if set, zero value otherwise.
func (o *GetAccountTradeFeeV5RespData) GetTakerU() string {
	if o == nil || IsNil(o.TakerU) {
		var ret string
		return ret
	}
	return *o.TakerU
}

// GetTakerUOk returns a tuple with the TakerU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountTradeFeeV5RespData) GetTakerUOk() (*string, bool) {
	if o == nil || IsNil(o.TakerU) {
		return nil, false
	}
	return o.TakerU, true
}

// HasTakerU returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasTakerU() bool {
	if o != nil && !IsNil(o.TakerU) {
		return true
	}

	return false
}

// SetTakerU gets a reference to the given string and assigns it to the TakerU field.
func (o *GetAccountTradeFeeV5RespData) SetTakerU(v string) {
	o.TakerU = &v
}

// GetTakerUSDC returns the TakerUSDC field value if set, zero value otherwise.
func (o *GetAccountTradeFeeV5RespData) GetTakerUSDC() string {
	if o == nil || IsNil(o.TakerUSDC) {
		var ret string
		return ret
	}
	return *o.TakerUSDC
}

// GetTakerUSDCOk returns a tuple with the TakerUSDC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountTradeFeeV5RespData) GetTakerUSDCOk() (*string, bool) {
	if o == nil || IsNil(o.TakerUSDC) {
		return nil, false
	}
	return o.TakerUSDC, true
}

// HasTakerUSDC returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasTakerUSDC() bool {
	if o != nil && !IsNil(o.TakerUSDC) {
		return true
	}

	return false
}

// SetTakerUSDC gets a reference to the given string and assigns it to the TakerUSDC field.
func (o *GetAccountTradeFeeV5RespData) SetTakerUSDC(v string) {
	o.TakerUSDC = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetAccountTradeFeeV5RespData) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountTradeFeeV5RespData) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetAccountTradeFeeV5RespData) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetAccountTradeFeeV5RespData) SetTs(v string) {
	o.Ts = &v
}

func (o GetAccountTradeFeeV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountTradeFeeV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Delivery) {
		toSerialize["delivery"] = o.Delivery
	}
	if !IsNil(o.Exercise) {
		toSerialize["exercise"] = o.Exercise
	}
	if !IsNil(o.Fiat) {
		toSerialize["fiat"] = o.Fiat
	}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Maker) {
		toSerialize["maker"] = o.Maker
	}
	if !IsNil(o.MakerU) {
		toSerialize["makerU"] = o.MakerU
	}
	if !IsNil(o.MakerUSDC) {
		toSerialize["makerUSDC"] = o.MakerUSDC
	}
	if !IsNil(o.RuleType) {
		toSerialize["ruleType"] = o.RuleType
	}
	if !IsNil(o.Taker) {
		toSerialize["taker"] = o.Taker
	}
	if !IsNil(o.TakerU) {
		toSerialize["takerU"] = o.TakerU
	}
	if !IsNil(o.TakerUSDC) {
		toSerialize["takerUSDC"] = o.TakerUSDC
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetAccountTradeFeeV5RespData struct {
	value *GetAccountTradeFeeV5RespData
	isSet bool
}

func (v NullableGetAccountTradeFeeV5RespData) Get() *GetAccountTradeFeeV5RespData {
	return v.value
}

func (v *NullableGetAccountTradeFeeV5RespData) Set(val *GetAccountTradeFeeV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountTradeFeeV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountTradeFeeV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountTradeFeeV5RespData(val *GetAccountTradeFeeV5RespData) *NullableGetAccountTradeFeeV5RespData {
	return &NullableGetAccountTradeFeeV5RespData{value: val, isSet: true}
}

func (v NullableGetAccountTradeFeeV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountTradeFeeV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


