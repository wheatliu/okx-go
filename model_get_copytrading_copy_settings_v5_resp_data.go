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

// checks if the GetCopytradingCopySettingsV5RespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCopytradingCopySettingsV5RespData{}

// GetCopytradingCopySettingsV5RespData struct for GetCopytradingCopySettingsV5RespData
type GetCopytradingCopySettingsV5RespData struct {
	// Margin currency
	Ccy *string `json:"ccy,omitempty"`
	// Copy amount in USDT per order.
	CopyAmt *string `json:"copyAmt,omitempty"`
	// Copy contract type setted  `custom`: custom by `instId` which is required；  `copy`: Keep your contracts consistent with this trader by automatically adding or removing contracts when they do
	CopyInstIdType *string `json:"copyInstIdType,omitempty"`
	// Copy margin mode  `cross`: cross  `isolated`: isolated  `copy`: Use the same margin mode as lead trader when opening positions
	CopyMgnMode *string `json:"copyMgnMode,omitempty"`
	// Copy mode  `fixed_amount` `ratio_copy`
	CopyMode *string `json:"copyMode,omitempty"`
	// Copy ratio per order.
	CopyRatio *string `json:"copyRatio,omitempty"`
	// Current copy state   `0`: non-copy, `1`: copy
	CopyState *string `json:"copyState,omitempty"`
	// Maximum total amount in USDT.   The maximum total amount you'll invest at any given time across all orders in this copy trade
	CopyTotalAmt *string `json:"copyTotalAmt,omitempty"`
	// Instrument list. It will return all lead contracts of the lead trader
	InstIds []GetCopytradingCopySettingsV5RespDataInstIdsInner `json:"instIds,omitempty"`
	// Stop loss per order. 0.1 represents 10%
	SlRatio *string `json:"slRatio,omitempty"`
	// Total stop loss in USDT for trader.
	SlTotalAmt *string `json:"slTotalAmt,omitempty"`
	// Action type for open positions  `market_close`: immediately close at market price  `copy_close`：close when trader closes  `manual_close`: close manually
	SubPosCloseType *string `json:"subPosCloseType,omitempty"`
	// Order tag
	Tag *string `json:"tag,omitempty"`
	// Take profit per order. 0.1 represents 10%
	TpRatio *string `json:"tpRatio,omitempty"`
}

// NewGetCopytradingCopySettingsV5RespData instantiates a new GetCopytradingCopySettingsV5RespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCopytradingCopySettingsV5RespData() *GetCopytradingCopySettingsV5RespData {
	this := GetCopytradingCopySettingsV5RespData{}
	var ccy string = ""
	this.Ccy = &ccy
	var copyAmt string = ""
	this.CopyAmt = &copyAmt
	var copyInstIdType string = ""
	this.CopyInstIdType = &copyInstIdType
	var copyMgnMode string = ""
	this.CopyMgnMode = &copyMgnMode
	var copyMode string = ""
	this.CopyMode = &copyMode
	var copyRatio string = ""
	this.CopyRatio = &copyRatio
	var copyState string = ""
	this.CopyState = &copyState
	var copyTotalAmt string = ""
	this.CopyTotalAmt = &copyTotalAmt
	var slRatio string = ""
	this.SlRatio = &slRatio
	var slTotalAmt string = ""
	this.SlTotalAmt = &slTotalAmt
	var subPosCloseType string = ""
	this.SubPosCloseType = &subPosCloseType
	var tag string = ""
	this.Tag = &tag
	var tpRatio string = ""
	this.TpRatio = &tpRatio
	return &this
}

// NewGetCopytradingCopySettingsV5RespDataWithDefaults instantiates a new GetCopytradingCopySettingsV5RespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCopytradingCopySettingsV5RespDataWithDefaults() *GetCopytradingCopySettingsV5RespData {
	this := GetCopytradingCopySettingsV5RespData{}
	var ccy string = ""
	this.Ccy = &ccy
	var copyAmt string = ""
	this.CopyAmt = &copyAmt
	var copyInstIdType string = ""
	this.CopyInstIdType = &copyInstIdType
	var copyMgnMode string = ""
	this.CopyMgnMode = &copyMgnMode
	var copyMode string = ""
	this.CopyMode = &copyMode
	var copyRatio string = ""
	this.CopyRatio = &copyRatio
	var copyState string = ""
	this.CopyState = &copyState
	var copyTotalAmt string = ""
	this.CopyTotalAmt = &copyTotalAmt
	var slRatio string = ""
	this.SlRatio = &slRatio
	var slTotalAmt string = ""
	this.SlTotalAmt = &slTotalAmt
	var subPosCloseType string = ""
	this.SubPosCloseType = &subPosCloseType
	var tag string = ""
	this.Tag = &tag
	var tpRatio string = ""
	this.TpRatio = &tpRatio
	return &this
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetCopytradingCopySettingsV5RespData) SetCcy(v string) {
	o.Ccy = &v
}

// GetCopyAmt returns the CopyAmt field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyAmt() string {
	if o == nil || IsNil(o.CopyAmt) {
		var ret string
		return ret
	}
	return *o.CopyAmt
}

// GetCopyAmtOk returns a tuple with the CopyAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyAmtOk() (*string, bool) {
	if o == nil || IsNil(o.CopyAmt) {
		return nil, false
	}
	return o.CopyAmt, true
}

// HasCopyAmt returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasCopyAmt() bool {
	if o != nil && !IsNil(o.CopyAmt) {
		return true
	}

	return false
}

// SetCopyAmt gets a reference to the given string and assigns it to the CopyAmt field.
func (o *GetCopytradingCopySettingsV5RespData) SetCopyAmt(v string) {
	o.CopyAmt = &v
}

// GetCopyInstIdType returns the CopyInstIdType field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyInstIdType() string {
	if o == nil || IsNil(o.CopyInstIdType) {
		var ret string
		return ret
	}
	return *o.CopyInstIdType
}

// GetCopyInstIdTypeOk returns a tuple with the CopyInstIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyInstIdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CopyInstIdType) {
		return nil, false
	}
	return o.CopyInstIdType, true
}

// HasCopyInstIdType returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasCopyInstIdType() bool {
	if o != nil && !IsNil(o.CopyInstIdType) {
		return true
	}

	return false
}

// SetCopyInstIdType gets a reference to the given string and assigns it to the CopyInstIdType field.
func (o *GetCopytradingCopySettingsV5RespData) SetCopyInstIdType(v string) {
	o.CopyInstIdType = &v
}

// GetCopyMgnMode returns the CopyMgnMode field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyMgnMode() string {
	if o == nil || IsNil(o.CopyMgnMode) {
		var ret string
		return ret
	}
	return *o.CopyMgnMode
}

// GetCopyMgnModeOk returns a tuple with the CopyMgnMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyMgnModeOk() (*string, bool) {
	if o == nil || IsNil(o.CopyMgnMode) {
		return nil, false
	}
	return o.CopyMgnMode, true
}

// HasCopyMgnMode returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasCopyMgnMode() bool {
	if o != nil && !IsNil(o.CopyMgnMode) {
		return true
	}

	return false
}

// SetCopyMgnMode gets a reference to the given string and assigns it to the CopyMgnMode field.
func (o *GetCopytradingCopySettingsV5RespData) SetCopyMgnMode(v string) {
	o.CopyMgnMode = &v
}

// GetCopyMode returns the CopyMode field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyMode() string {
	if o == nil || IsNil(o.CopyMode) {
		var ret string
		return ret
	}
	return *o.CopyMode
}

// GetCopyModeOk returns a tuple with the CopyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyModeOk() (*string, bool) {
	if o == nil || IsNil(o.CopyMode) {
		return nil, false
	}
	return o.CopyMode, true
}

// HasCopyMode returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasCopyMode() bool {
	if o != nil && !IsNil(o.CopyMode) {
		return true
	}

	return false
}

// SetCopyMode gets a reference to the given string and assigns it to the CopyMode field.
func (o *GetCopytradingCopySettingsV5RespData) SetCopyMode(v string) {
	o.CopyMode = &v
}

// GetCopyRatio returns the CopyRatio field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyRatio() string {
	if o == nil || IsNil(o.CopyRatio) {
		var ret string
		return ret
	}
	return *o.CopyRatio
}

// GetCopyRatioOk returns a tuple with the CopyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyRatioOk() (*string, bool) {
	if o == nil || IsNil(o.CopyRatio) {
		return nil, false
	}
	return o.CopyRatio, true
}

// HasCopyRatio returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasCopyRatio() bool {
	if o != nil && !IsNil(o.CopyRatio) {
		return true
	}

	return false
}

// SetCopyRatio gets a reference to the given string and assigns it to the CopyRatio field.
func (o *GetCopytradingCopySettingsV5RespData) SetCopyRatio(v string) {
	o.CopyRatio = &v
}

// GetCopyState returns the CopyState field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyState() string {
	if o == nil || IsNil(o.CopyState) {
		var ret string
		return ret
	}
	return *o.CopyState
}

// GetCopyStateOk returns a tuple with the CopyState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyStateOk() (*string, bool) {
	if o == nil || IsNil(o.CopyState) {
		return nil, false
	}
	return o.CopyState, true
}

// HasCopyState returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasCopyState() bool {
	if o != nil && !IsNil(o.CopyState) {
		return true
	}

	return false
}

// SetCopyState gets a reference to the given string and assigns it to the CopyState field.
func (o *GetCopytradingCopySettingsV5RespData) SetCopyState(v string) {
	o.CopyState = &v
}

// GetCopyTotalAmt returns the CopyTotalAmt field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyTotalAmt() string {
	if o == nil || IsNil(o.CopyTotalAmt) {
		var ret string
		return ret
	}
	return *o.CopyTotalAmt
}

// GetCopyTotalAmtOk returns a tuple with the CopyTotalAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetCopyTotalAmtOk() (*string, bool) {
	if o == nil || IsNil(o.CopyTotalAmt) {
		return nil, false
	}
	return o.CopyTotalAmt, true
}

// HasCopyTotalAmt returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasCopyTotalAmt() bool {
	if o != nil && !IsNil(o.CopyTotalAmt) {
		return true
	}

	return false
}

// SetCopyTotalAmt gets a reference to the given string and assigns it to the CopyTotalAmt field.
func (o *GetCopytradingCopySettingsV5RespData) SetCopyTotalAmt(v string) {
	o.CopyTotalAmt = &v
}

// GetInstIds returns the InstIds field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetInstIds() []GetCopytradingCopySettingsV5RespDataInstIdsInner {
	if o == nil || IsNil(o.InstIds) {
		var ret []GetCopytradingCopySettingsV5RespDataInstIdsInner
		return ret
	}
	return o.InstIds
}

// GetInstIdsOk returns a tuple with the InstIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetInstIdsOk() ([]GetCopytradingCopySettingsV5RespDataInstIdsInner, bool) {
	if o == nil || IsNil(o.InstIds) {
		return nil, false
	}
	return o.InstIds, true
}

// HasInstIds returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasInstIds() bool {
	if o != nil && !IsNil(o.InstIds) {
		return true
	}

	return false
}

// SetInstIds gets a reference to the given []GetCopytradingCopySettingsV5RespDataInstIdsInner and assigns it to the InstIds field.
func (o *GetCopytradingCopySettingsV5RespData) SetInstIds(v []GetCopytradingCopySettingsV5RespDataInstIdsInner) {
	o.InstIds = v
}

// GetSlRatio returns the SlRatio field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetSlRatio() string {
	if o == nil || IsNil(o.SlRatio) {
		var ret string
		return ret
	}
	return *o.SlRatio
}

// GetSlRatioOk returns a tuple with the SlRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetSlRatioOk() (*string, bool) {
	if o == nil || IsNil(o.SlRatio) {
		return nil, false
	}
	return o.SlRatio, true
}

// HasSlRatio returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasSlRatio() bool {
	if o != nil && !IsNil(o.SlRatio) {
		return true
	}

	return false
}

// SetSlRatio gets a reference to the given string and assigns it to the SlRatio field.
func (o *GetCopytradingCopySettingsV5RespData) SetSlRatio(v string) {
	o.SlRatio = &v
}

// GetSlTotalAmt returns the SlTotalAmt field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetSlTotalAmt() string {
	if o == nil || IsNil(o.SlTotalAmt) {
		var ret string
		return ret
	}
	return *o.SlTotalAmt
}

// GetSlTotalAmtOk returns a tuple with the SlTotalAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetSlTotalAmtOk() (*string, bool) {
	if o == nil || IsNil(o.SlTotalAmt) {
		return nil, false
	}
	return o.SlTotalAmt, true
}

// HasSlTotalAmt returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasSlTotalAmt() bool {
	if o != nil && !IsNil(o.SlTotalAmt) {
		return true
	}

	return false
}

// SetSlTotalAmt gets a reference to the given string and assigns it to the SlTotalAmt field.
func (o *GetCopytradingCopySettingsV5RespData) SetSlTotalAmt(v string) {
	o.SlTotalAmt = &v
}

// GetSubPosCloseType returns the SubPosCloseType field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetSubPosCloseType() string {
	if o == nil || IsNil(o.SubPosCloseType) {
		var ret string
		return ret
	}
	return *o.SubPosCloseType
}

// GetSubPosCloseTypeOk returns a tuple with the SubPosCloseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetSubPosCloseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubPosCloseType) {
		return nil, false
	}
	return o.SubPosCloseType, true
}

// HasSubPosCloseType returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasSubPosCloseType() bool {
	if o != nil && !IsNil(o.SubPosCloseType) {
		return true
	}

	return false
}

// SetSubPosCloseType gets a reference to the given string and assigns it to the SubPosCloseType field.
func (o *GetCopytradingCopySettingsV5RespData) SetSubPosCloseType(v string) {
	o.SubPosCloseType = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *GetCopytradingCopySettingsV5RespData) SetTag(v string) {
	o.Tag = &v
}

// GetTpRatio returns the TpRatio field value if set, zero value otherwise.
func (o *GetCopytradingCopySettingsV5RespData) GetTpRatio() string {
	if o == nil || IsNil(o.TpRatio) {
		var ret string
		return ret
	}
	return *o.TpRatio
}

// GetTpRatioOk returns a tuple with the TpRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCopytradingCopySettingsV5RespData) GetTpRatioOk() (*string, bool) {
	if o == nil || IsNil(o.TpRatio) {
		return nil, false
	}
	return o.TpRatio, true
}

// HasTpRatio returns a boolean if a field has been set.
func (o *GetCopytradingCopySettingsV5RespData) HasTpRatio() bool {
	if o != nil && !IsNil(o.TpRatio) {
		return true
	}

	return false
}

// SetTpRatio gets a reference to the given string and assigns it to the TpRatio field.
func (o *GetCopytradingCopySettingsV5RespData) SetTpRatio(v string) {
	o.TpRatio = &v
}

func (o GetCopytradingCopySettingsV5RespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCopytradingCopySettingsV5RespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.CopyAmt) {
		toSerialize["copyAmt"] = o.CopyAmt
	}
	if !IsNil(o.CopyInstIdType) {
		toSerialize["copyInstIdType"] = o.CopyInstIdType
	}
	if !IsNil(o.CopyMgnMode) {
		toSerialize["copyMgnMode"] = o.CopyMgnMode
	}
	if !IsNil(o.CopyMode) {
		toSerialize["copyMode"] = o.CopyMode
	}
	if !IsNil(o.CopyRatio) {
		toSerialize["copyRatio"] = o.CopyRatio
	}
	if !IsNil(o.CopyState) {
		toSerialize["copyState"] = o.CopyState
	}
	if !IsNil(o.CopyTotalAmt) {
		toSerialize["copyTotalAmt"] = o.CopyTotalAmt
	}
	if !IsNil(o.InstIds) {
		toSerialize["instIds"] = o.InstIds
	}
	if !IsNil(o.SlRatio) {
		toSerialize["slRatio"] = o.SlRatio
	}
	if !IsNil(o.SlTotalAmt) {
		toSerialize["slTotalAmt"] = o.SlTotalAmt
	}
	if !IsNil(o.SubPosCloseType) {
		toSerialize["subPosCloseType"] = o.SubPosCloseType
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TpRatio) {
		toSerialize["tpRatio"] = o.TpRatio
	}
	return toSerialize, nil
}

type NullableGetCopytradingCopySettingsV5RespData struct {
	value *GetCopytradingCopySettingsV5RespData
	isSet bool
}

func (v NullableGetCopytradingCopySettingsV5RespData) Get() *GetCopytradingCopySettingsV5RespData {
	return v.value
}

func (v *NullableGetCopytradingCopySettingsV5RespData) Set(val *GetCopytradingCopySettingsV5RespData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCopytradingCopySettingsV5RespData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCopytradingCopySettingsV5RespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCopytradingCopySettingsV5RespData(val *GetCopytradingCopySettingsV5RespData) *NullableGetCopytradingCopySettingsV5RespData {
	return &NullableGetCopytradingCopySettingsV5RespData{value: val, isSet: true}
}

func (v NullableGetCopytradingCopySettingsV5RespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCopytradingCopySettingsV5RespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


