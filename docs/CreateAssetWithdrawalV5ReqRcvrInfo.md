# CreateAssetWithdrawalV5ReqRcvrInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchId** | Pointer to **string** | Exchange ID  You can query supported exchanges through the endpoint of   If the exchange is not in the exchange list, fill in &#39;0&#39; in this field.   Apply to walletType &#x3D; &#x60;exchange&#x60; | [optional] [default to ""]
**RcvrCountry** | Pointer to **string** | The recipient&#39;s country, e.g. &#x60;United States&#x60;  You must enter an English country name or a two letter country code (ISO 3166-1). Please refer to the &#x60;Country Name&#x60; and &#x60;Country Code&#x60; in the country information table below. | [optional] [default to ""]
**RcvrCountrySubDivision** | Pointer to **string** | State/Province of the recipient, e.g. &#x60;California&#x60; | [optional] [default to ""]
**RcvrFirstName** | Pointer to **string** | Receiver&#39;s first name, e.g. &#x60;Bruce&#x60; | [optional] [default to ""]
**RcvrLastName** | Pointer to **string** | Receiver&#39;s last name, e.g. &#x60;Wayne&#x60; | [optional] [default to ""]
**RcvrStreetName** | Pointer to **string** | Recipient&#39;s street address, e.g. &#x60;Clementi Avenue 1&#x60; | [optional] [default to ""]
**RcvrTownName** | Pointer to **string** | The town/city where the recipient is located, e.g. &#x60;San Jose&#x60; | [optional] [default to ""]
**WalletType** | Pointer to **string** | Wallet Type  &#x60;exchange&#x60;: Withdraw to exchange wallet  &#x60;private&#x60;: Withdraw to private wallet  For the wallet belongs to business recipient, &#x60;rcvrFirstName&#x60; may input the company name, &#x60;rcvrLastName&#x60; may input \&quot;N/A\&quot;, location info may input the registered address of the company. | [optional] [default to ""]

## Methods

### NewCreateAssetWithdrawalV5ReqRcvrInfo

`func NewCreateAssetWithdrawalV5ReqRcvrInfo() *CreateAssetWithdrawalV5ReqRcvrInfo`

NewCreateAssetWithdrawalV5ReqRcvrInfo instantiates a new CreateAssetWithdrawalV5ReqRcvrInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssetWithdrawalV5ReqRcvrInfoWithDefaults

`func NewCreateAssetWithdrawalV5ReqRcvrInfoWithDefaults() *CreateAssetWithdrawalV5ReqRcvrInfo`

NewCreateAssetWithdrawalV5ReqRcvrInfoWithDefaults instantiates a new CreateAssetWithdrawalV5ReqRcvrInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchId

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetExchId() string`

GetExchId returns the ExchId field if non-nil, zero value otherwise.

### GetExchIdOk

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetExchIdOk() (*string, bool)`

GetExchIdOk returns a tuple with the ExchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchId

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetExchId(v string)`

SetExchId sets ExchId field to given value.

### HasExchId

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasExchId() bool`

HasExchId returns a boolean if a field has been set.

### GetRcvrCountry

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrCountry() string`

GetRcvrCountry returns the RcvrCountry field if non-nil, zero value otherwise.

### GetRcvrCountryOk

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrCountryOk() (*string, bool)`

GetRcvrCountryOk returns a tuple with the RcvrCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcvrCountry

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetRcvrCountry(v string)`

SetRcvrCountry sets RcvrCountry field to given value.

### HasRcvrCountry

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasRcvrCountry() bool`

HasRcvrCountry returns a boolean if a field has been set.

### GetRcvrCountrySubDivision

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrCountrySubDivision() string`

GetRcvrCountrySubDivision returns the RcvrCountrySubDivision field if non-nil, zero value otherwise.

### GetRcvrCountrySubDivisionOk

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrCountrySubDivisionOk() (*string, bool)`

GetRcvrCountrySubDivisionOk returns a tuple with the RcvrCountrySubDivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcvrCountrySubDivision

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetRcvrCountrySubDivision(v string)`

SetRcvrCountrySubDivision sets RcvrCountrySubDivision field to given value.

### HasRcvrCountrySubDivision

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasRcvrCountrySubDivision() bool`

HasRcvrCountrySubDivision returns a boolean if a field has been set.

### GetRcvrFirstName

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrFirstName() string`

GetRcvrFirstName returns the RcvrFirstName field if non-nil, zero value otherwise.

### GetRcvrFirstNameOk

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrFirstNameOk() (*string, bool)`

GetRcvrFirstNameOk returns a tuple with the RcvrFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcvrFirstName

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetRcvrFirstName(v string)`

SetRcvrFirstName sets RcvrFirstName field to given value.

### HasRcvrFirstName

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasRcvrFirstName() bool`

HasRcvrFirstName returns a boolean if a field has been set.

### GetRcvrLastName

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrLastName() string`

GetRcvrLastName returns the RcvrLastName field if non-nil, zero value otherwise.

### GetRcvrLastNameOk

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrLastNameOk() (*string, bool)`

GetRcvrLastNameOk returns a tuple with the RcvrLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcvrLastName

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetRcvrLastName(v string)`

SetRcvrLastName sets RcvrLastName field to given value.

### HasRcvrLastName

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasRcvrLastName() bool`

HasRcvrLastName returns a boolean if a field has been set.

### GetRcvrStreetName

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrStreetName() string`

GetRcvrStreetName returns the RcvrStreetName field if non-nil, zero value otherwise.

### GetRcvrStreetNameOk

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrStreetNameOk() (*string, bool)`

GetRcvrStreetNameOk returns a tuple with the RcvrStreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcvrStreetName

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetRcvrStreetName(v string)`

SetRcvrStreetName sets RcvrStreetName field to given value.

### HasRcvrStreetName

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasRcvrStreetName() bool`

HasRcvrStreetName returns a boolean if a field has been set.

### GetRcvrTownName

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrTownName() string`

GetRcvrTownName returns the RcvrTownName field if non-nil, zero value otherwise.

### GetRcvrTownNameOk

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetRcvrTownNameOk() (*string, bool)`

GetRcvrTownNameOk returns a tuple with the RcvrTownName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcvrTownName

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetRcvrTownName(v string)`

SetRcvrTownName sets RcvrTownName field to given value.

### HasRcvrTownName

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasRcvrTownName() bool`

HasRcvrTownName returns a boolean if a field has been set.

### GetWalletType

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetWalletType() string`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) GetWalletTypeOk() (*string, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) SetWalletType(v string)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *CreateAssetWithdrawalV5ReqRcvrInfo) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


