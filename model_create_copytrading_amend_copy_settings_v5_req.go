/*
OKX v5 API

OpenAPI specification for Okx exchange - Rest API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateCopytradingAmendCopySettingsV5Req type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCopytradingAmendCopySettingsV5Req{}

// CreateCopytradingAmendCopySettingsV5Req struct for CreateCopytradingAmendCopySettingsV5Req
type CreateCopytradingAmendCopySettingsV5Req struct {
	// Copy amount per order in USDT
	CopyAmt *string `json:"copyAmt,omitempty"`
	// Copy contract type setted  `custom`: custom by `instId` which is required；  `copy`: Keep your contracts consistent with this trader by automatically adding or removing contracts when they do
	CopyInstIdType string `json:"copyInstIdType"`
	// Copy margin mode  `cross`: cross  `isolated`: isolated  `copy`: Use the same margin mode as lead trader when opening positions
	CopyMgnMode string `json:"copyMgnMode"`
	// Copy mode  `fixed_amount`: set the same fixed amount for each order, and `copyAmt` is required；  `ratio_copy`: set amount as a multiple of the lead trader’s order value, and `copyRatio` is required   The default is `fixed_amount`
	CopyMode *string `json:"copyMode,omitempty"`
	// Copy ratio per order.
	CopyRatio *string `json:"copyRatio,omitempty"`
	// Maximum total amount in USDT.   The maximum total amount you'll invest at any given time across all orders in this copy trade  You won’t copy new orders if you exceed this amount
	CopyTotalAmt string `json:"copyTotalAmt"`
	// Instrument ID.   If there are multiple instruments, separate them with commas.
	InstId *string `json:"instId,omitempty"`
	// Instrument type  `SWAP`
	InstType *string `json:"instType,omitempty"`
	// Stop loss per order. 0.1 represents 10%
	SlRatio *string `json:"slRatio,omitempty"`
	// Total stop loss in USDT for trader.  If your net loss (total profit - total loss) reaches this amount, you'll stop copying this trader
	SlTotalAmt *string `json:"slTotalAmt,omitempty"`
	// Action type for open positions  `market_close`: immediately close at market price  `copy_close`：close when trader closes  `manual_close`: close manually  The default is `copy_close`
	SubPosCloseType string `json:"subPosCloseType"`
	// Order tag  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters.
	Tag *string `json:"tag,omitempty"`
	// Take profit per order. 0.1 represents 10%
	TpRatio *string `json:"tpRatio,omitempty"`
	// Lead trader unique code  A combination of case-sensitive alphanumerics, all numbers and the length is 16 characters, e.g. 213E8C92DC61EFAC
	UniqueCode string `json:"uniqueCode"`
}

type _CreateCopytradingAmendCopySettingsV5Req CreateCopytradingAmendCopySettingsV5Req

// NewCreateCopytradingAmendCopySettingsV5Req instantiates a new CreateCopytradingAmendCopySettingsV5Req object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCopytradingAmendCopySettingsV5Req(copyInstIdType string, copyMgnMode string, copyTotalAmt string, subPosCloseType string, uniqueCode string) *CreateCopytradingAmendCopySettingsV5Req {
	this := CreateCopytradingAmendCopySettingsV5Req{}
	var copyAmt string = ""
	this.CopyAmt = &copyAmt
	this.CopyInstIdType = copyInstIdType
	this.CopyMgnMode = copyMgnMode
	var copyMode string = ""
	this.CopyMode = &copyMode
	var copyRatio string = ""
	this.CopyRatio = &copyRatio
	this.CopyTotalAmt = copyTotalAmt
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var slRatio string = ""
	this.SlRatio = &slRatio
	var slTotalAmt string = ""
	this.SlTotalAmt = &slTotalAmt
	this.SubPosCloseType = subPosCloseType
	var tag string = ""
	this.Tag = &tag
	var tpRatio string = ""
	this.TpRatio = &tpRatio
	this.UniqueCode = uniqueCode
	return &this
}

// NewCreateCopytradingAmendCopySettingsV5ReqWithDefaults instantiates a new CreateCopytradingAmendCopySettingsV5Req object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCopytradingAmendCopySettingsV5ReqWithDefaults() *CreateCopytradingAmendCopySettingsV5Req {
	this := CreateCopytradingAmendCopySettingsV5Req{}
	var copyAmt string = ""
	this.CopyAmt = &copyAmt
	var copyInstIdType string = ""
	this.CopyInstIdType = copyInstIdType
	var copyMgnMode string = ""
	this.CopyMgnMode = copyMgnMode
	var copyMode string = ""
	this.CopyMode = &copyMode
	var copyRatio string = ""
	this.CopyRatio = &copyRatio
	var copyTotalAmt string = ""
	this.CopyTotalAmt = copyTotalAmt
	var instId string = ""
	this.InstId = &instId
	var instType string = ""
	this.InstType = &instType
	var slRatio string = ""
	this.SlRatio = &slRatio
	var slTotalAmt string = ""
	this.SlTotalAmt = &slTotalAmt
	var subPosCloseType string = ""
	this.SubPosCloseType = subPosCloseType
	var tag string = ""
	this.Tag = &tag
	var tpRatio string = ""
	this.TpRatio = &tpRatio
	var uniqueCode string = ""
	this.UniqueCode = uniqueCode
	return &this
}

// GetCopyAmt returns the CopyAmt field value if set, zero value otherwise.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetCopyAmt() string {
	if o == nil || IsNil(o.CopyAmt) {
		var ret string
		return ret
	}
	return *o.CopyAmt
}

// GetCopyAmtOk returns a tuple with the CopyAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetCopyAmtOk() (*string, bool) {
	if o == nil || IsNil(o.CopyAmt) {
		return nil, false
	}
	return o.CopyAmt, true
}

// HasCopyAmt returns a boolean if a field has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) HasCopyAmt() bool {
	if o != nil && !IsNil(o.CopyAmt) {
		return true
	}

	return false
}

// SetCopyAmt gets a reference to the given string and assigns it to the CopyAmt field.
func (o *CreateCopytradingAmendCopySettingsV5Req) SetCopyAmt(v string) {
	o.CopyAmt = &v
}

// GetCopyInstIdType returns the CopyInstIdType field value
func (o *CreateCopytradingAmendCopySettingsV5Req) GetCopyInstIdType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CopyInstIdType
}

// GetCopyInstIdTypeOk returns a tuple with the CopyInstIdType field value
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetCopyInstIdTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyInstIdType, true
}

// SetCopyInstIdType sets field value
func (o *CreateCopytradingAmendCopySettingsV5Req) SetCopyInstIdType(v string) {
	o.CopyInstIdType = v
}

// GetCopyMgnMode returns the CopyMgnMode field value
func (o *CreateCopytradingAmendCopySettingsV5Req) GetCopyMgnMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CopyMgnMode
}

// GetCopyMgnModeOk returns a tuple with the CopyMgnMode field value
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetCopyMgnModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyMgnMode, true
}

// SetCopyMgnMode sets field value
func (o *CreateCopytradingAmendCopySettingsV5Req) SetCopyMgnMode(v string) {
	o.CopyMgnMode = v
}

// GetCopyMode returns the CopyMode field value if set, zero value otherwise.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetCopyMode() string {
	if o == nil || IsNil(o.CopyMode) {
		var ret string
		return ret
	}
	return *o.CopyMode
}

// GetCopyModeOk returns a tuple with the CopyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetCopyModeOk() (*string, bool) {
	if o == nil || IsNil(o.CopyMode) {
		return nil, false
	}
	return o.CopyMode, true
}

// HasCopyMode returns a boolean if a field has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) HasCopyMode() bool {
	if o != nil && !IsNil(o.CopyMode) {
		return true
	}

	return false
}

// SetCopyMode gets a reference to the given string and assigns it to the CopyMode field.
func (o *CreateCopytradingAmendCopySettingsV5Req) SetCopyMode(v string) {
	o.CopyMode = &v
}

// GetCopyRatio returns the CopyRatio field value if set, zero value otherwise.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetCopyRatio() string {
	if o == nil || IsNil(o.CopyRatio) {
		var ret string
		return ret
	}
	return *o.CopyRatio
}

// GetCopyRatioOk returns a tuple with the CopyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetCopyRatioOk() (*string, bool) {
	if o == nil || IsNil(o.CopyRatio) {
		return nil, false
	}
	return o.CopyRatio, true
}

// HasCopyRatio returns a boolean if a field has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) HasCopyRatio() bool {
	if o != nil && !IsNil(o.CopyRatio) {
		return true
	}

	return false
}

// SetCopyRatio gets a reference to the given string and assigns it to the CopyRatio field.
func (o *CreateCopytradingAmendCopySettingsV5Req) SetCopyRatio(v string) {
	o.CopyRatio = &v
}

// GetCopyTotalAmt returns the CopyTotalAmt field value
func (o *CreateCopytradingAmendCopySettingsV5Req) GetCopyTotalAmt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CopyTotalAmt
}

// GetCopyTotalAmtOk returns a tuple with the CopyTotalAmt field value
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetCopyTotalAmtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyTotalAmt, true
}

// SetCopyTotalAmt sets field value
func (o *CreateCopytradingAmendCopySettingsV5Req) SetCopyTotalAmt(v string) {
	o.CopyTotalAmt = v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetInstId() string {
	if o == nil || IsNil(o.InstId) {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetInstIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstId) {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) HasInstId() bool {
	if o != nil && !IsNil(o.InstId) {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *CreateCopytradingAmendCopySettingsV5Req) SetInstId(v string) {
	o.InstId = &v
}

// GetInstType returns the InstType field value if set, zero value otherwise.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetInstType() string {
	if o == nil || IsNil(o.InstType) {
		var ret string
		return ret
	}
	return *o.InstType
}

// GetInstTypeOk returns a tuple with the InstType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetInstTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstType) {
		return nil, false
	}
	return o.InstType, true
}

// HasInstType returns a boolean if a field has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) HasInstType() bool {
	if o != nil && !IsNil(o.InstType) {
		return true
	}

	return false
}

// SetInstType gets a reference to the given string and assigns it to the InstType field.
func (o *CreateCopytradingAmendCopySettingsV5Req) SetInstType(v string) {
	o.InstType = &v
}

// GetSlRatio returns the SlRatio field value if set, zero value otherwise.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetSlRatio() string {
	if o == nil || IsNil(o.SlRatio) {
		var ret string
		return ret
	}
	return *o.SlRatio
}

// GetSlRatioOk returns a tuple with the SlRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetSlRatioOk() (*string, bool) {
	if o == nil || IsNil(o.SlRatio) {
		return nil, false
	}
	return o.SlRatio, true
}

// HasSlRatio returns a boolean if a field has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) HasSlRatio() bool {
	if o != nil && !IsNil(o.SlRatio) {
		return true
	}

	return false
}

// SetSlRatio gets a reference to the given string and assigns it to the SlRatio field.
func (o *CreateCopytradingAmendCopySettingsV5Req) SetSlRatio(v string) {
	o.SlRatio = &v
}

// GetSlTotalAmt returns the SlTotalAmt field value if set, zero value otherwise.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetSlTotalAmt() string {
	if o == nil || IsNil(o.SlTotalAmt) {
		var ret string
		return ret
	}
	return *o.SlTotalAmt
}

// GetSlTotalAmtOk returns a tuple with the SlTotalAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetSlTotalAmtOk() (*string, bool) {
	if o == nil || IsNil(o.SlTotalAmt) {
		return nil, false
	}
	return o.SlTotalAmt, true
}

// HasSlTotalAmt returns a boolean if a field has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) HasSlTotalAmt() bool {
	if o != nil && !IsNil(o.SlTotalAmt) {
		return true
	}

	return false
}

// SetSlTotalAmt gets a reference to the given string and assigns it to the SlTotalAmt field.
func (o *CreateCopytradingAmendCopySettingsV5Req) SetSlTotalAmt(v string) {
	o.SlTotalAmt = &v
}

// GetSubPosCloseType returns the SubPosCloseType field value
func (o *CreateCopytradingAmendCopySettingsV5Req) GetSubPosCloseType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubPosCloseType
}

// GetSubPosCloseTypeOk returns a tuple with the SubPosCloseType field value
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetSubPosCloseTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubPosCloseType, true
}

// SetSubPosCloseType sets field value
func (o *CreateCopytradingAmendCopySettingsV5Req) SetSubPosCloseType(v string) {
	o.SubPosCloseType = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *CreateCopytradingAmendCopySettingsV5Req) SetTag(v string) {
	o.Tag = &v
}

// GetTpRatio returns the TpRatio field value if set, zero value otherwise.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetTpRatio() string {
	if o == nil || IsNil(o.TpRatio) {
		var ret string
		return ret
	}
	return *o.TpRatio
}

// GetTpRatioOk returns a tuple with the TpRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetTpRatioOk() (*string, bool) {
	if o == nil || IsNil(o.TpRatio) {
		return nil, false
	}
	return o.TpRatio, true
}

// HasTpRatio returns a boolean if a field has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) HasTpRatio() bool {
	if o != nil && !IsNil(o.TpRatio) {
		return true
	}

	return false
}

// SetTpRatio gets a reference to the given string and assigns it to the TpRatio field.
func (o *CreateCopytradingAmendCopySettingsV5Req) SetTpRatio(v string) {
	o.TpRatio = &v
}

// GetUniqueCode returns the UniqueCode field value
func (o *CreateCopytradingAmendCopySettingsV5Req) GetUniqueCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniqueCode
}

// GetUniqueCodeOk returns a tuple with the UniqueCode field value
// and a boolean to check if the value has been set.
func (o *CreateCopytradingAmendCopySettingsV5Req) GetUniqueCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueCode, true
}

// SetUniqueCode sets field value
func (o *CreateCopytradingAmendCopySettingsV5Req) SetUniqueCode(v string) {
	o.UniqueCode = v
}

func (o CreateCopytradingAmendCopySettingsV5Req) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCopytradingAmendCopySettingsV5Req) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CopyAmt) {
		toSerialize["copyAmt"] = o.CopyAmt
	}
	toSerialize["copyInstIdType"] = o.CopyInstIdType
	toSerialize["copyMgnMode"] = o.CopyMgnMode
	if !IsNil(o.CopyMode) {
		toSerialize["copyMode"] = o.CopyMode
	}
	if !IsNil(o.CopyRatio) {
		toSerialize["copyRatio"] = o.CopyRatio
	}
	toSerialize["copyTotalAmt"] = o.CopyTotalAmt
	if !IsNil(o.InstId) {
		toSerialize["instId"] = o.InstId
	}
	if !IsNil(o.InstType) {
		toSerialize["instType"] = o.InstType
	}
	if !IsNil(o.SlRatio) {
		toSerialize["slRatio"] = o.SlRatio
	}
	if !IsNil(o.SlTotalAmt) {
		toSerialize["slTotalAmt"] = o.SlTotalAmt
	}
	toSerialize["subPosCloseType"] = o.SubPosCloseType
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TpRatio) {
		toSerialize["tpRatio"] = o.TpRatio
	}
	toSerialize["uniqueCode"] = o.UniqueCode
	return toSerialize, nil
}

func (o *CreateCopytradingAmendCopySettingsV5Req) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"copyInstIdType",
		"copyMgnMode",
		"copyTotalAmt",
		"subPosCloseType",
		"uniqueCode",
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

	varCreateCopytradingAmendCopySettingsV5Req := _CreateCopytradingAmendCopySettingsV5Req{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCopytradingAmendCopySettingsV5Req)

	if err != nil {
		return err
	}

	*o = CreateCopytradingAmendCopySettingsV5Req(varCreateCopytradingAmendCopySettingsV5Req)

	return err
}

type NullableCreateCopytradingAmendCopySettingsV5Req struct {
	value *CreateCopytradingAmendCopySettingsV5Req
	isSet bool
}

func (v NullableCreateCopytradingAmendCopySettingsV5Req) Get() *CreateCopytradingAmendCopySettingsV5Req {
	return v.value
}

func (v *NullableCreateCopytradingAmendCopySettingsV5Req) Set(val *CreateCopytradingAmendCopySettingsV5Req) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCopytradingAmendCopySettingsV5Req) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCopytradingAmendCopySettingsV5Req) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCopytradingAmendCopySettingsV5Req(val *CreateCopytradingAmendCopySettingsV5Req) *NullableCreateCopytradingAmendCopySettingsV5Req {
	return &NullableCreateCopytradingAmendCopySettingsV5Req{value: val, isSet: true}
}

func (v NullableCreateCopytradingAmendCopySettingsV5Req) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCopytradingAmendCopySettingsV5Req) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


