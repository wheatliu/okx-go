# GetPublicInsuranceFundV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]GetPublicInsuranceFundV5RespDataInnerDetailsInner**](GetPublicInsuranceFundV5RespDataInnerDetailsInner.md) | Insurance fund data | [optional] 
**InstFamily** | Pointer to **string** | Instrument family  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Total** | Pointer to **string** | The total balance of insurance fund, in &#x60;USD&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicInsuranceFundV5RespDataInner

`func NewGetPublicInsuranceFundV5RespDataInner() *GetPublicInsuranceFundV5RespDataInner`

NewGetPublicInsuranceFundV5RespDataInner instantiates a new GetPublicInsuranceFundV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicInsuranceFundV5RespDataInnerWithDefaults

`func NewGetPublicInsuranceFundV5RespDataInnerWithDefaults() *GetPublicInsuranceFundV5RespDataInner`

NewGetPublicInsuranceFundV5RespDataInnerWithDefaults instantiates a new GetPublicInsuranceFundV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *GetPublicInsuranceFundV5RespDataInner) GetDetails() []GetPublicInsuranceFundV5RespDataInnerDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetPublicInsuranceFundV5RespDataInner) GetDetailsOk() (*[]GetPublicInsuranceFundV5RespDataInnerDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetPublicInsuranceFundV5RespDataInner) SetDetails(v []GetPublicInsuranceFundV5RespDataInnerDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetPublicInsuranceFundV5RespDataInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetInstFamily

`func (o *GetPublicInsuranceFundV5RespDataInner) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetPublicInsuranceFundV5RespDataInner) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetPublicInsuranceFundV5RespDataInner) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetPublicInsuranceFundV5RespDataInner) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstType

`func (o *GetPublicInsuranceFundV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetPublicInsuranceFundV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetPublicInsuranceFundV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetPublicInsuranceFundV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetTotal

`func (o *GetPublicInsuranceFundV5RespDataInner) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetPublicInsuranceFundV5RespDataInner) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetPublicInsuranceFundV5RespDataInner) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetPublicInsuranceFundV5RespDataInner) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


