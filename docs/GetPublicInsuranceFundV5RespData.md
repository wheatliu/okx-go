# GetPublicInsuranceFundV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]GetPublicInsuranceFundV5RespDataDetailsInner**](GetPublicInsuranceFundV5RespDataDetailsInner.md) | Insurance fund data | [optional] 
**InstFamily** | Pointer to **string** | Instrument family  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Total** | Pointer to **string** | The total balance of insurance fund, in &#x60;USD&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicInsuranceFundV5RespData

`func NewGetPublicInsuranceFundV5RespData() *GetPublicInsuranceFundV5RespData`

NewGetPublicInsuranceFundV5RespData instantiates a new GetPublicInsuranceFundV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicInsuranceFundV5RespDataWithDefaults

`func NewGetPublicInsuranceFundV5RespDataWithDefaults() *GetPublicInsuranceFundV5RespData`

NewGetPublicInsuranceFundV5RespDataWithDefaults instantiates a new GetPublicInsuranceFundV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *GetPublicInsuranceFundV5RespData) GetDetails() []GetPublicInsuranceFundV5RespDataDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetPublicInsuranceFundV5RespData) GetDetailsOk() (*[]GetPublicInsuranceFundV5RespDataDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetPublicInsuranceFundV5RespData) SetDetails(v []GetPublicInsuranceFundV5RespDataDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetPublicInsuranceFundV5RespData) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetInstFamily

`func (o *GetPublicInsuranceFundV5RespData) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetPublicInsuranceFundV5RespData) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetPublicInsuranceFundV5RespData) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetPublicInsuranceFundV5RespData) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstType

`func (o *GetPublicInsuranceFundV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetPublicInsuranceFundV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetPublicInsuranceFundV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetPublicInsuranceFundV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetTotal

`func (o *GetPublicInsuranceFundV5RespData) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetPublicInsuranceFundV5RespData) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetPublicInsuranceFundV5RespData) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetPublicInsuranceFundV5RespData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


