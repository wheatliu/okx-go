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

// checks if the CreateAssetWithdrawalV5ReqRcvrInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssetWithdrawalV5ReqRcvrInfo{}

// CreateAssetWithdrawalV5ReqRcvrInfo Recipient information  For the specific entity users to do on-chain withdrawal/lightning withdrawal, this information is required.
type CreateAssetWithdrawalV5ReqRcvrInfo struct {
	// Exchange ID  You can query supported exchanges through the endpoint of   If the exchange is not in the exchange list, fill in '0' in this field.   Apply to walletType = `exchange`
	ExchId *string `json:"exchId,omitempty"`
	// The recipient's country, e.g. `United States`  You must enter an English country name or a two letter country code (ISO 3166-1). Please refer to the `Country Name` and `Country Code` in the country information table below.
	RcvrCountry *string `json:"rcvrCountry,omitempty"`
	// State/Province of the recipient, e.g. `California`
	RcvrCountrySubDivision *string `json:"rcvrCountrySubDivision,omitempty"`
	// Receiver's first name, e.g. `Bruce`
	RcvrFirstName *string `json:"rcvrFirstName,omitempty"`
	// Receiver's last name, e.g. `Wayne`
	RcvrLastName *string `json:"rcvrLastName,omitempty"`
	// Recipient's street address, e.g. `Clementi Avenue 1`
	RcvrStreetName *string `json:"rcvrStreetName,omitempty"`
	// The town/city where the recipient is located, e.g. `San Jose`
	RcvrTownName *string `json:"rcvrTownName,omitempty"`
	// Wallet Type  `exchange`: Withdraw to exchange wallet  `private`: Withdraw to private wallet  For the wallet belongs to business recipient, `rcvrFirstName` may input the company name, `rcvrLastName` may input \"N/A\", location info may input the registered address of the company.
	WalletType *string `json:"walletType,omitempty"`
}

// NewCreateAssetWithdrawalV5ReqRcvrInfo instantiates a new CreateAssetWithdrawalV5ReqRcvrInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssetWithdrawalV5ReqRcvrInfo() *CreateAssetWithdrawalV5ReqRcvrInfo {
	this := CreateAssetWithdrawalV5ReqRcvrInfo{}
	var exchId string = ""
	this.ExchId = &exchId
	var rcvrCountry string = ""
	this.RcvrCountry = &rcvrCountry
	var rcvrCountrySubDivision string = ""
	this.RcvrCountrySubDivision = &rcvrCountrySubDivision
	var rcvrFirstName string = ""
	this.RcvrFirstName = &rcvrFirstName
	var rcvrLastName string = ""
	this.RcvrLastName = &rcvrLastName
	var rcvrStreetName string = ""
	this.RcvrStreetName = &rcvrStreetName
	var rcvrTownName string = ""
	this.RcvrTownName = &rcvrTownName
	var walletType string = ""
	this.WalletType = &walletType
	return &this
}

// NewCreateAssetWithdrawalV5ReqRcvrInfoWithDefaults instantiates a new CreateAssetWithdrawalV5ReqRcvrInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssetWithdrawalV5ReqRcvrInfoWithDefaults() *CreateAssetWithdrawalV5ReqRcvrInfo {
	this := CreateAssetWithdrawalV5ReqRcvrInfo{}
	var exchId string = ""
	this.ExchId = &exchId
	var rcvrCountry string = ""
	this.RcvrCountry = &rcvrCountry
	var rcvrCountrySubDivision string = ""
	this.RcvrCountrySubDivision = &rcvrCountrySubDivision
	var rcvrFirstName string = ""
	this.RcvrFirstName = &rcvrFirstName
	var rcvrLastName string = ""
	this.RcvrLastName = &rcvrLastName
	var rcvrStreetName string = ""
	this.RcvrStreetName = &rcvrStreetName
	var rcvrTownName string = ""
	this.RcvrTownName = &rcvrTownName
	var walletType string = ""
	this.WalletType = &walletType
	return &this
}

// GetExchId returns the ExchId field value if set, zero value otherwise.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetExchId() string {
	if o == nil || IsNil(o.ExchId) {
		var ret string
		return ret
	}
	return *o.ExchId
}

// GetExchIdOk returns a tuple with the ExchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetExchIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExchId) {
		return nil, false
	}
	return o.ExchId, true
}

// HasExchId returns a boolean if a field has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasExchId() bool {
	if o != nil && !IsNil(o.ExchId) {
		return true
	}

	return false
}

// SetExchId gets a reference to the given string and assigns it to the ExchId field.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetExchId(v string) {
	o.ExchId = &v
}

// GetRcvrCountry returns the RcvrCountry field value if set, zero value otherwise.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrCountry() string {
	if o == nil || IsNil(o.RcvrCountry) {
		var ret string
		return ret
	}
	return *o.RcvrCountry
}

// GetRcvrCountryOk returns a tuple with the RcvrCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrCountryOk() (*string, bool) {
	if o == nil || IsNil(o.RcvrCountry) {
		return nil, false
	}
	return o.RcvrCountry, true
}

// HasRcvrCountry returns a boolean if a field has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasRcvrCountry() bool {
	if o != nil && !IsNil(o.RcvrCountry) {
		return true
	}

	return false
}

// SetRcvrCountry gets a reference to the given string and assigns it to the RcvrCountry field.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetRcvrCountry(v string) {
	o.RcvrCountry = &v
}

// GetRcvrCountrySubDivision returns the RcvrCountrySubDivision field value if set, zero value otherwise.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrCountrySubDivision() string {
	if o == nil || IsNil(o.RcvrCountrySubDivision) {
		var ret string
		return ret
	}
	return *o.RcvrCountrySubDivision
}

// GetRcvrCountrySubDivisionOk returns a tuple with the RcvrCountrySubDivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrCountrySubDivisionOk() (*string, bool) {
	if o == nil || IsNil(o.RcvrCountrySubDivision) {
		return nil, false
	}
	return o.RcvrCountrySubDivision, true
}

// HasRcvrCountrySubDivision returns a boolean if a field has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasRcvrCountrySubDivision() bool {
	if o != nil && !IsNil(o.RcvrCountrySubDivision) {
		return true
	}

	return false
}

// SetRcvrCountrySubDivision gets a reference to the given string and assigns it to the RcvrCountrySubDivision field.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetRcvrCountrySubDivision(v string) {
	o.RcvrCountrySubDivision = &v
}

// GetRcvrFirstName returns the RcvrFirstName field value if set, zero value otherwise.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrFirstName() string {
	if o == nil || IsNil(o.RcvrFirstName) {
		var ret string
		return ret
	}
	return *o.RcvrFirstName
}

// GetRcvrFirstNameOk returns a tuple with the RcvrFirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.RcvrFirstName) {
		return nil, false
	}
	return o.RcvrFirstName, true
}

// HasRcvrFirstName returns a boolean if a field has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasRcvrFirstName() bool {
	if o != nil && !IsNil(o.RcvrFirstName) {
		return true
	}

	return false
}

// SetRcvrFirstName gets a reference to the given string and assigns it to the RcvrFirstName field.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetRcvrFirstName(v string) {
	o.RcvrFirstName = &v
}

// GetRcvrLastName returns the RcvrLastName field value if set, zero value otherwise.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrLastName() string {
	if o == nil || IsNil(o.RcvrLastName) {
		var ret string
		return ret
	}
	return *o.RcvrLastName
}

// GetRcvrLastNameOk returns a tuple with the RcvrLastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.RcvrLastName) {
		return nil, false
	}
	return o.RcvrLastName, true
}

// HasRcvrLastName returns a boolean if a field has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasRcvrLastName() bool {
	if o != nil && !IsNil(o.RcvrLastName) {
		return true
	}

	return false
}

// SetRcvrLastName gets a reference to the given string and assigns it to the RcvrLastName field.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetRcvrLastName(v string) {
	o.RcvrLastName = &v
}

// GetRcvrStreetName returns the RcvrStreetName field value if set, zero value otherwise.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrStreetName() string {
	if o == nil || IsNil(o.RcvrStreetName) {
		var ret string
		return ret
	}
	return *o.RcvrStreetName
}

// GetRcvrStreetNameOk returns a tuple with the RcvrStreetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrStreetNameOk() (*string, bool) {
	if o == nil || IsNil(o.RcvrStreetName) {
		return nil, false
	}
	return o.RcvrStreetName, true
}

// HasRcvrStreetName returns a boolean if a field has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasRcvrStreetName() bool {
	if o != nil && !IsNil(o.RcvrStreetName) {
		return true
	}

	return false
}

// SetRcvrStreetName gets a reference to the given string and assigns it to the RcvrStreetName field.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetRcvrStreetName(v string) {
	o.RcvrStreetName = &v
}

// GetRcvrTownName returns the RcvrTownName field value if set, zero value otherwise.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrTownName() string {
	if o == nil || IsNil(o.RcvrTownName) {
		var ret string
		return ret
	}
	return *o.RcvrTownName
}

// GetRcvrTownNameOk returns a tuple with the RcvrTownName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrTownNameOk() (*string, bool) {
	if o == nil || IsNil(o.RcvrTownName) {
		return nil, false
	}
	return o.RcvrTownName, true
}

// HasRcvrTownName returns a boolean if a field has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasRcvrTownName() bool {
	if o != nil && !IsNil(o.RcvrTownName) {
		return true
	}

	return false
}

// SetRcvrTownName gets a reference to the given string and assigns it to the RcvrTownName field.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetRcvrTownName(v string) {
	o.RcvrTownName = &v
}

// GetWalletType returns the WalletType field value if set, zero value otherwise.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetWalletType() string {
	if o == nil || IsNil(o.WalletType) {
		var ret string
		return ret
	}
	return *o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetWalletTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WalletType) {
		return nil, false
	}
	return o.WalletType, true
}

// HasWalletType returns a boolean if a field has been set.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasWalletType() bool {
	if o != nil && !IsNil(o.WalletType) {
		return true
	}

	return false
}

// SetWalletType gets a reference to the given string and assigns it to the WalletType field.
func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetWalletType(v string) {
	o.WalletType = &v
}

func (o CreateAssetWithdrawalV5ReqRcvrInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAssetWithdrawalV5ReqRcvrInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExchId) {
		toSerialize["exchId"] = o.ExchId
	}
	if !IsNil(o.RcvrCountry) {
		toSerialize["rcvrCountry"] = o.RcvrCountry
	}
	if !IsNil(o.RcvrCountrySubDivision) {
		toSerialize["rcvrCountrySubDivision"] = o.RcvrCountrySubDivision
	}
	if !IsNil(o.RcvrFirstName) {
		toSerialize["rcvrFirstName"] = o.RcvrFirstName
	}
	if !IsNil(o.RcvrLastName) {
		toSerialize["rcvrLastName"] = o.RcvrLastName
	}
	if !IsNil(o.RcvrStreetName) {
		toSerialize["rcvrStreetName"] = o.RcvrStreetName
	}
	if !IsNil(o.RcvrTownName) {
		toSerialize["rcvrTownName"] = o.RcvrTownName
	}
	if !IsNil(o.WalletType) {
		toSerialize["walletType"] = o.WalletType
	}
	return toSerialize, nil
}

type NullableCreateAssetWithdrawalV5ReqRcvrInfo struct {
	value *CreateAssetWithdrawalV5ReqRcvrInfo
	isSet bool
}

func (v NullableCreateAssetWithdrawalV5ReqRcvrInfo) Get() *CreateAssetWithdrawalV5ReqRcvrInfo {
	return v.value
}

func (v *NullableCreateAssetWithdrawalV5ReqRcvrInfo) Set(val *CreateAssetWithdrawalV5ReqRcvrInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssetWithdrawalV5ReqRcvrInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssetWithdrawalV5ReqRcvrInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssetWithdrawalV5ReqRcvrInfo(val *CreateAssetWithdrawalV5ReqRcvrInfo) *NullableCreateAssetWithdrawalV5ReqRcvrInfo {
	return &NullableCreateAssetWithdrawalV5ReqRcvrInfo{value: val, isSet: true}
}

func (v NullableCreateAssetWithdrawalV5ReqRcvrInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAssetWithdrawalV5ReqRcvrInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


