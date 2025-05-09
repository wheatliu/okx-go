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

// checks if the GetAssetDepositHistoryV5RespDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAssetDepositHistoryV5RespDataInner{}

// GetAssetDepositHistoryV5RespDataInner struct for GetAssetDepositHistoryV5RespDataInner
type GetAssetDepositHistoryV5RespDataInner struct {
	// The actual amount of blockchain confirmed in a single deposit
	ActualDepBlkConfirm *string `json:"actualDepBlkConfirm,omitempty"`
	// Deposit amount
	Amt *string `json:"amt,omitempty"`
	// If `from` is a phone number, this parameter return area code of the phone number
	AreaCodeFrom *string `json:"areaCodeFrom,omitempty"`
	// Currency
	Ccy *string `json:"ccy,omitempty"`
	// Chain name
	Chain *string `json:"chain,omitempty"`
	// Deposit ID
	DepId *string `json:"depId,omitempty"`
	// Deposit account  If the deposit comes from an internal transfer, this field displays the account information of the internal transfer initiator, which can be a mobile phone number, email address, account name, and will return \"\" in other cases
	From *string `json:"from,omitempty"`
	// Internal transfer initiator's withdrawal ID  If the deposit comes from internal transfer, this field displays the withdrawal ID of the internal transfer initiator, and will return \"\" in other cases
	FromWdId *string `json:"fromWdId,omitempty"`
	// Status of deposit  `0`: Waiting for confirmation  `1`: Deposit credited    `2`: Deposit successful   `8`: Pending due to temporary deposit suspension on this crypto currency  `11`: Match the address blacklist  `12`: Account or deposit is frozen  `13`: Sub-account deposit interception  `14`: KYC limit
	State *string `json:"state,omitempty"`
	// Deposit address  If the deposit comes from the on-chain, this field displays the on-chain address, and will return \"\" in other cases
	To *string `json:"to,omitempty"`
	// The timestamp that the deposit record is created, Unix timestamp format in milliseconds, e.g. `1655251200000`
	Ts *string `json:"ts,omitempty"`
	// Hash record of the deposit
	TxId *string `json:"txId,omitempty"`
}

// NewGetAssetDepositHistoryV5RespDataInner instantiates a new GetAssetDepositHistoryV5RespDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssetDepositHistoryV5RespDataInner() *GetAssetDepositHistoryV5RespDataInner {
	this := GetAssetDepositHistoryV5RespDataInner{}
	var actualDepBlkConfirm string = ""
	this.ActualDepBlkConfirm = &actualDepBlkConfirm
	var amt string = ""
	this.Amt = &amt
	var areaCodeFrom string = ""
	this.AreaCodeFrom = &areaCodeFrom
	var ccy string = ""
	this.Ccy = &ccy
	var chain string = ""
	this.Chain = &chain
	var depId string = ""
	this.DepId = &depId
	var from string = ""
	this.From = &from
	var fromWdId string = ""
	this.FromWdId = &fromWdId
	var state string = ""
	this.State = &state
	var to string = ""
	this.To = &to
	var ts string = ""
	this.Ts = &ts
	var txId string = ""
	this.TxId = &txId
	return &this
}

// NewGetAssetDepositHistoryV5RespDataInnerWithDefaults instantiates a new GetAssetDepositHistoryV5RespDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssetDepositHistoryV5RespDataInnerWithDefaults() *GetAssetDepositHistoryV5RespDataInner {
	this := GetAssetDepositHistoryV5RespDataInner{}
	var actualDepBlkConfirm string = ""
	this.ActualDepBlkConfirm = &actualDepBlkConfirm
	var amt string = ""
	this.Amt = &amt
	var areaCodeFrom string = ""
	this.AreaCodeFrom = &areaCodeFrom
	var ccy string = ""
	this.Ccy = &ccy
	var chain string = ""
	this.Chain = &chain
	var depId string = ""
	this.DepId = &depId
	var from string = ""
	this.From = &from
	var fromWdId string = ""
	this.FromWdId = &fromWdId
	var state string = ""
	this.State = &state
	var to string = ""
	this.To = &to
	var ts string = ""
	this.Ts = &ts
	var txId string = ""
	this.TxId = &txId
	return &this
}

// GetActualDepBlkConfirm returns the ActualDepBlkConfirm field value if set, zero value otherwise.
func (o *GetAssetDepositHistoryV5RespDataInner) GetActualDepBlkConfirm() string {
	if o == nil || IsNil(o.ActualDepBlkConfirm) {
		var ret string
		return ret
	}
	return *o.ActualDepBlkConfirm
}

// GetActualDepBlkConfirmOk returns a tuple with the ActualDepBlkConfirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) GetActualDepBlkConfirmOk() (*string, bool) {
	if o == nil || IsNil(o.ActualDepBlkConfirm) {
		return nil, false
	}
	return o.ActualDepBlkConfirm, true
}

// HasActualDepBlkConfirm returns a boolean if a field has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) HasActualDepBlkConfirm() bool {
	if o != nil && !IsNil(o.ActualDepBlkConfirm) {
		return true
	}

	return false
}

// SetActualDepBlkConfirm gets a reference to the given string and assigns it to the ActualDepBlkConfirm field.
func (o *GetAssetDepositHistoryV5RespDataInner) SetActualDepBlkConfirm(v string) {
	o.ActualDepBlkConfirm = &v
}

// GetAmt returns the Amt field value if set, zero value otherwise.
func (o *GetAssetDepositHistoryV5RespDataInner) GetAmt() string {
	if o == nil || IsNil(o.Amt) {
		var ret string
		return ret
	}
	return *o.Amt
}

// GetAmtOk returns a tuple with the Amt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) GetAmtOk() (*string, bool) {
	if o == nil || IsNil(o.Amt) {
		return nil, false
	}
	return o.Amt, true
}

// HasAmt returns a boolean if a field has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) HasAmt() bool {
	if o != nil && !IsNil(o.Amt) {
		return true
	}

	return false
}

// SetAmt gets a reference to the given string and assigns it to the Amt field.
func (o *GetAssetDepositHistoryV5RespDataInner) SetAmt(v string) {
	o.Amt = &v
}

// GetAreaCodeFrom returns the AreaCodeFrom field value if set, zero value otherwise.
func (o *GetAssetDepositHistoryV5RespDataInner) GetAreaCodeFrom() string {
	if o == nil || IsNil(o.AreaCodeFrom) {
		var ret string
		return ret
	}
	return *o.AreaCodeFrom
}

// GetAreaCodeFromOk returns a tuple with the AreaCodeFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) GetAreaCodeFromOk() (*string, bool) {
	if o == nil || IsNil(o.AreaCodeFrom) {
		return nil, false
	}
	return o.AreaCodeFrom, true
}

// HasAreaCodeFrom returns a boolean if a field has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) HasAreaCodeFrom() bool {
	if o != nil && !IsNil(o.AreaCodeFrom) {
		return true
	}

	return false
}

// SetAreaCodeFrom gets a reference to the given string and assigns it to the AreaCodeFrom field.
func (o *GetAssetDepositHistoryV5RespDataInner) SetAreaCodeFrom(v string) {
	o.AreaCodeFrom = &v
}

// GetCcy returns the Ccy field value if set, zero value otherwise.
func (o *GetAssetDepositHistoryV5RespDataInner) GetCcy() string {
	if o == nil || IsNil(o.Ccy) {
		var ret string
		return ret
	}
	return *o.Ccy
}

// GetCcyOk returns a tuple with the Ccy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) GetCcyOk() (*string, bool) {
	if o == nil || IsNil(o.Ccy) {
		return nil, false
	}
	return o.Ccy, true
}

// HasCcy returns a boolean if a field has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) HasCcy() bool {
	if o != nil && !IsNil(o.Ccy) {
		return true
	}

	return false
}

// SetCcy gets a reference to the given string and assigns it to the Ccy field.
func (o *GetAssetDepositHistoryV5RespDataInner) SetCcy(v string) {
	o.Ccy = &v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *GetAssetDepositHistoryV5RespDataInner) GetChain() string {
	if o == nil || IsNil(o.Chain) {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) GetChainOk() (*string, bool) {
	if o == nil || IsNil(o.Chain) {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) HasChain() bool {
	if o != nil && !IsNil(o.Chain) {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *GetAssetDepositHistoryV5RespDataInner) SetChain(v string) {
	o.Chain = &v
}

// GetDepId returns the DepId field value if set, zero value otherwise.
func (o *GetAssetDepositHistoryV5RespDataInner) GetDepId() string {
	if o == nil || IsNil(o.DepId) {
		var ret string
		return ret
	}
	return *o.DepId
}

// GetDepIdOk returns a tuple with the DepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) GetDepIdOk() (*string, bool) {
	if o == nil || IsNil(o.DepId) {
		return nil, false
	}
	return o.DepId, true
}

// HasDepId returns a boolean if a field has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) HasDepId() bool {
	if o != nil && !IsNil(o.DepId) {
		return true
	}

	return false
}

// SetDepId gets a reference to the given string and assigns it to the DepId field.
func (o *GetAssetDepositHistoryV5RespDataInner) SetDepId(v string) {
	o.DepId = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *GetAssetDepositHistoryV5RespDataInner) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *GetAssetDepositHistoryV5RespDataInner) SetFrom(v string) {
	o.From = &v
}

// GetFromWdId returns the FromWdId field value if set, zero value otherwise.
func (o *GetAssetDepositHistoryV5RespDataInner) GetFromWdId() string {
	if o == nil || IsNil(o.FromWdId) {
		var ret string
		return ret
	}
	return *o.FromWdId
}

// GetFromWdIdOk returns a tuple with the FromWdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) GetFromWdIdOk() (*string, bool) {
	if o == nil || IsNil(o.FromWdId) {
		return nil, false
	}
	return o.FromWdId, true
}

// HasFromWdId returns a boolean if a field has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) HasFromWdId() bool {
	if o != nil && !IsNil(o.FromWdId) {
		return true
	}

	return false
}

// SetFromWdId gets a reference to the given string and assigns it to the FromWdId field.
func (o *GetAssetDepositHistoryV5RespDataInner) SetFromWdId(v string) {
	o.FromWdId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetAssetDepositHistoryV5RespDataInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetAssetDepositHistoryV5RespDataInner) SetState(v string) {
	o.State = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *GetAssetDepositHistoryV5RespDataInner) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *GetAssetDepositHistoryV5RespDataInner) SetTo(v string) {
	o.To = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetAssetDepositHistoryV5RespDataInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetAssetDepositHistoryV5RespDataInner) SetTs(v string) {
	o.Ts = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *GetAssetDepositHistoryV5RespDataInner) GetTxId() string {
	if o == nil || IsNil(o.TxId) {
		var ret string
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) GetTxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *GetAssetDepositHistoryV5RespDataInner) HasTxId() bool {
	if o != nil && !IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given string and assigns it to the TxId field.
func (o *GetAssetDepositHistoryV5RespDataInner) SetTxId(v string) {
	o.TxId = &v
}

func (o GetAssetDepositHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssetDepositHistoryV5RespDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActualDepBlkConfirm) {
		toSerialize["actualDepBlkConfirm"] = o.ActualDepBlkConfirm
	}
	if !IsNil(o.Amt) {
		toSerialize["amt"] = o.Amt
	}
	if !IsNil(o.AreaCodeFrom) {
		toSerialize["areaCodeFrom"] = o.AreaCodeFrom
	}
	if !IsNil(o.Ccy) {
		toSerialize["ccy"] = o.Ccy
	}
	if !IsNil(o.Chain) {
		toSerialize["chain"] = o.Chain
	}
	if !IsNil(o.DepId) {
		toSerialize["depId"] = o.DepId
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.FromWdId) {
		toSerialize["fromWdId"] = o.FromWdId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.TxId) {
		toSerialize["txId"] = o.TxId
	}
	return toSerialize, nil
}

type NullableGetAssetDepositHistoryV5RespDataInner struct {
	value *GetAssetDepositHistoryV5RespDataInner
	isSet bool
}

func (v NullableGetAssetDepositHistoryV5RespDataInner) Get() *GetAssetDepositHistoryV5RespDataInner {
	return v.value
}

func (v *NullableGetAssetDepositHistoryV5RespDataInner) Set(val *GetAssetDepositHistoryV5RespDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetDepositHistoryV5RespDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetDepositHistoryV5RespDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetDepositHistoryV5RespDataInner(val *GetAssetDepositHistoryV5RespDataInner) *NullableGetAssetDepositHistoryV5RespDataInner {
	return &NullableGetAssetDepositHistoryV5RespDataInner{value: val, isSet: true}
}

func (v NullableGetAssetDepositHistoryV5RespDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetDepositHistoryV5RespDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


