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

// checks if the CreateAccountPositionBuilderV5RespDataInnerPositionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountPositionBuilderV5RespDataInnerPositionsInner{}

// CreateAccountPositionBuilderV5RespDataInnerPositionsInner struct for CreateAccountPositionBuilderV5RespDataInnerPositionsInner
type CreateAccountPositionBuilderV5RespDataInnerPositionsInner struct {
	// When `instType` is `SPOT`, it represents spot in use.  When `instType` is `SWAP`/`FUTURES`/`OPTION`, it represents position amount.
	Amt *string `json:"amt,omitempty"`
	// Average open price
	AvgPx *string `json:"avgPx,omitempty"`
	// Float P&L
	FloatPnl *string `json:"floatPnl,omitempty"`
	// IMR
	Imr *string `json:"imr,omitempty"`
	// Instrument ID, e.g. `BTC-USDT-SWAP`
	InstId *string `json:"instId,omitempty"`
	// Instrument type  `SPOT`  `SWAP`  `FUTURES`  `OPTION`
	InstType *string `json:"instType,omitempty"`
	// Whether it is a real position  If `instType` is `SWAP`/`FUTURES`/`OPTION`, it is a valid parameter, else it will returns `false`
	IsRealPos *bool `json:"isRealPos,omitempty"`
	// Leverage
	Lever *string `json:"lever,omitempty"`
	// Mark price
	MarkPx *string `json:"markPx,omitempty"`
	// Margin ratio
	MgnRatio *string `json:"mgnRatio,omitempty"`
	// Notional in `USD`
	NotionalUsd *string `json:"notionalUsd,omitempty"`
	// Position side  `long`  `short`  `net`
	PosSide *string `json:"posSide,omitempty"`
}

// NewCreateAccountPositionBuilderV5RespDataInnerPositionsInner instantiates a new CreateAccountPositionBuilderV5RespDataInnerPositionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountPositionBuilderV5RespDataInnerPositionsInner() *CreateAccountPositionBuilderV5RespDataInnerPositionsInner {
	this := CreateAccountPositionBuilderV5RespDataInnerPositionsInner{}
	var amt string = ""
	this.Amt = &amt
	var avgPx string = ""
	this.AvgPx = &avgPx
	var floatPnl string = ""
	this.FloatPnl = &floatPnl
	var imr string = ""
	this.Imr = &imr
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var lever string = ""
	this.Lever = &lever
	var markPx string = ""
	this.MarkPx = &markPx
	var mgnRatio string = ""
	this.MgnRatio = &mgnRatio
	var notionalUsd string = ""
	this.NotionalUsd = &notionalUsd
	var posSide string = ""
	this.PosSide = &posSide
	return &this
}

// NewCreateAccountPositionBuilderV5RespDataInnerPositionsInnerWithDefaults instantiates a new CreateAccountPositionBuilderV5RespDataInnerPositionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountPositionBuilderV5RespDataInnerPositionsInnerWithDefaults() *CreateAccountPositionBuilderV5RespDataInnerPositionsInner {
	this := CreateAccountPositionBuilderV5RespDataInnerPositionsInner{}
	var amt string = ""
	this.Amt = &amt
	var avgPx string = ""
	this.AvgPx = &avgPx
	var floatPnl string = ""
	this.FloatPnl = &floatPnl
	var imr string = ""
	this.Imr = &imr
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var lever string = ""
	this.Lever = &lever
	var markPx string = ""
	this.MarkPx = &markPx
	var mgnRatio string = ""
	this.MgnRatio = &mgnRatio
	var notionalUsd string = ""
	this.NotionalUsd = &notionalUsd
	var posSide string = ""
	this.PosSide = &posSide
	return &this
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetAmt(v string) {
	o.Amt = &v
}

// GetAvgPx returns the AvgPx field value if set, zero value otherwise.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetAvgPx() string {
	if o == nil || IsNil(o.AvgPx) {
		var ret string
		return ret
	}
	return *o.AvgPx
}

// GetAvgPxOk returns a tuple with the AvgPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetAvgPxOk() (*string, bool) {
	if o == nil || IsNil(o.AvgPx) {
		return nil, false
	}
	return o.AvgPx, true
}

// HasAvgPx returns a boolean if a field has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasAvgPx() bool {
	if o != nil && !IsNil(o.AvgPx) {
		return true
	}

	return false
}

// SetAvgPx gets a reference to the given string and assigns it to the AvgPx field.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetAvgPx(v string) {
	o.AvgPx = &v
}

// GetFloatPnl returns the FloatPnl field value if set, zero value otherwise.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetFloatPnl() string {
	if o == nil || IsNil(o.FloatPnl) {
		var ret string
		return ret
	}
	return *o.FloatPnl
}

// GetFloatPnlOk returns a tuple with the FloatPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetFloatPnlOk() (*string, bool) {
	if o == nil || IsNil(o.FloatPnl) {
		return nil, false
	}
	return o.FloatPnl, true
}

// HasFloatPnl returns a boolean if a field has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasFloatPnl() bool {
	if o != nil && !IsNil(o.FloatPnl) {
		return true
	}

	return false
}

// SetFloatPnl gets a reference to the given string and assigns it to the FloatPnl field.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetFloatPnl(v string) {
	o.FloatPnl = &v
}

// GetImr returns the Imr field value if set, zero value otherwise.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetImr() string {
	if o == nil || IsNil(o.Imr) {
		var ret string
		return ret
	}
	return *o.Imr
}

// GetImrOk returns a tuple with the Imr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetImrOk() (*string, bool) {
	if o == nil || IsNil(o.Imr) {
		return nil, false
	}
	return o.Imr, true
}

// HasImr returns a boolean if a field has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasImr() bool {
	if o != nil && !IsNil(o.Imr) {
		return true
	}

	return false
}

// SetImr gets a reference to the given string and assigns it to the Imr field.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetImr(v string) {
	o.Imr = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetInstId(v string) {
	o.InstId = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetInstType(v string) {
	o.InstType = &v
}

// GetIsRealPos returns the IsRealPos field value if set, zero value otherwise.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetIsRealPos() bool {
	if o == nil || IsNil(o.IsRealPos) {
		var ret bool
		return ret
	}
	return *o.IsRealPos
}

// GetIsRealPosOk returns a tuple with the IsRealPos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetIsRealPosOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRealPos) {
		return nil, false
	}
	return o.IsRealPos, true
}

// HasIsRealPos returns a boolean if a field has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasIsRealPos() bool {
	if o != nil && !IsNil(o.IsRealPos) {
		return true
	}

	return false
}

// SetIsRealPos gets a reference to the given bool and assigns it to the IsRealPos field.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetIsRealPos(v bool) {
	o.IsRealPos = &v
}

// GetLever returns the Lever field value if set, zero value otherwise.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetLever() string {
	if o == nil || IsNil(o.Lever) {
		var ret string
		return ret
	}
	return *o.Lever
}

// GetLeverOk returns a tuple with the Lever field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetLeverOk() (*string, bool) {
	if o == nil || IsNil(o.Lever) {
		return nil, false
	}
	return o.Lever, true
}

// HasLever returns a boolean if a field has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasLever() bool {
	if o != nil && !IsNil(o.Lever) {
		return true
	}

	return false
}

// SetLever gets a reference to the given string and assigns it to the Lever field.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetLever(v string) {
	o.Lever = &v
}

// GetMarkPx returns the MarkPx field value if set, zero value otherwise.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetMarkPx() string {
	if o == nil || IsNil(o.MarkPx) {
		var ret string
		return ret
	}
	return *o.MarkPx
}

// GetMarkPxOk returns a tuple with the MarkPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetMarkPxOk() (*string, bool) {
	if o == nil || IsNil(o.MarkPx) {
		return nil, false
	}
	return o.MarkPx, true
}

// HasMarkPx returns a boolean if a field has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasMarkPx() bool {
	if o != nil && !IsNil(o.MarkPx) {
		return true
	}

	return false
}

// SetMarkPx gets a reference to the given string and assigns it to the MarkPx field.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetMarkPx(v string) {
	o.MarkPx = &v
}

// GetMgnRatio returns the MgnRatio field value if set, zero value otherwise.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetMgnRatio() string {
	if o == nil || IsNil(o.MgnRatio) {
		var ret string
		return ret
	}
	return *o.MgnRatio
}

// GetMgnRatioOk returns a tuple with the MgnRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetMgnRatioOk() (*string, bool) {
	if o == nil || IsNil(o.MgnRatio) {
		return nil, false
	}
	return o.MgnRatio, true
}

// HasMgnRatio returns a boolean if a field has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasMgnRatio() bool {
	if o != nil && !IsNil(o.MgnRatio) {
		return true
	}

	return false
}

// SetMgnRatio gets a reference to the given string and assigns it to the MgnRatio field.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetMgnRatio(v string) {
	o.MgnRatio = &v
}

// GetNotionalUsd returns the NotionalUsd field value if set, zero value otherwise.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetNotionalUsd() string {
	if o == nil || IsNil(o.NotionalUsd) {
		var ret string
		return ret
	}
	return *o.NotionalUsd
}

// GetNotionalUsdOk returns a tuple with the NotionalUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetNotionalUsdOk() (*string, bool) {
	if o == nil || IsNil(o.NotionalUsd) {
		return nil, false
	}
	return o.NotionalUsd, true
}

// HasNotionalUsd returns a boolean if a field has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasNotionalUsd() bool {
	if o != nil && !IsNil(o.NotionalUsd) {
		return true
	}

	return false
}

// SetNotionalUsd gets a reference to the given string and assigns it to the NotionalUsd field.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetNotionalUsd(v string) {
	o.NotionalUsd = &v
}

// GetPosSide returns the PosSide field value if set, zero value otherwise.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetPosSide() string {
	if o == nil || IsNil(o.PosSide) {
		var ret string
		return ret
	}
	return *o.PosSide
}

// GetPosSideOk returns a tuple with the PosSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) GetPosSideOk() (*string, bool) {
	if o == nil || IsNil(o.PosSide) {
		return nil, false
	}
	return o.PosSide, true
}

// HasPosSide returns a boolean if a field has been set.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) HasPosSide() bool {
	if o != nil && !IsNil(o.PosSide) {
		return true
	}

	return false
}

// SetPosSide gets a reference to the given string and assigns it to the PosSide field.
func (o *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) SetPosSide(v string) {
	o.PosSide = &v
}

func (o CreateAccountPositionBuilderV5RespDataInnerPositionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountPositionBuilderV5RespDataInnerPositionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.AvgPx) {
		toSerialize["avgPx"] = o.AvgPx
	}
	if !IsNil(o.FloatPnl) {
		toSerialize["floatPnl"] = o.FloatPnl
	}
	if !IsNil(o.Imr) {
		toSerialize["imr"] = o.Imr
	}
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.IsRealPos) {
		toSerialize["isRealPos"] = o.IsRealPos
	}
	if !IsNil(o.Lever) {
		toSerialize["lever"] = o.Lever
	}
	if !IsNil(o.MarkPx) {
		toSerialize["markPx"] = o.MarkPx
	}
	if !IsNil(o.MgnRatio) {
		toSerialize["mgnRatio"] = o.MgnRatio
	}
	if !IsNil(o.NotionalUsd) {
		toSerialize["notionalUsd"] = o.NotionalUsd
	}
	if !IsNil(o.PosSide) {
		toSerialize["posSide"] = o.PosSide
	}
	return toSerialize, nil
}

type NullableCreateAccountPositionBuilderV5RespDataInnerPositionsInner struct {
	value *CreateAccountPositionBuilderV5RespDataInnerPositionsInner
	isSet bool
}

func (v NullableCreateAccountPositionBuilderV5RespDataInnerPositionsInner) Get() *CreateAccountPositionBuilderV5RespDataInnerPositionsInner {
	return v.value
}

func (v *NullableCreateAccountPositionBuilderV5RespDataInnerPositionsInner) Set(val *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountPositionBuilderV5RespDataInnerPositionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountPositionBuilderV5RespDataInnerPositionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountPositionBuilderV5RespDataInnerPositionsInner(val *CreateAccountPositionBuilderV5RespDataInnerPositionsInner) *NullableCreateAccountPositionBuilderV5RespDataInnerPositionsInner {
	return &NullableCreateAccountPositionBuilderV5RespDataInnerPositionsInner{value: val, isSet: true}
}

func (v NullableCreateAccountPositionBuilderV5RespDataInnerPositionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountPositionBuilderV5RespDataInnerPositionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


