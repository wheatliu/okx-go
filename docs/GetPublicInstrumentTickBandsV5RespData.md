# GetPublicInstrumentTickBandsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstFamily** | Pointer to **string** | Instrument family | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**TickBand** | Pointer to [**[]GetPublicInstrumentTickBandsV5RespDataTickBandInner**](GetPublicInstrumentTickBandsV5RespDataTickBandInner.md) | Tick size band | [optional] 

## Methods

### NewGetPublicInstrumentTickBandsV5RespData

`func NewGetPublicInstrumentTickBandsV5RespData() *GetPublicInstrumentTickBandsV5RespData`

NewGetPublicInstrumentTickBandsV5RespData instantiates a new GetPublicInstrumentTickBandsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicInstrumentTickBandsV5RespDataWithDefaults

`func NewGetPublicInstrumentTickBandsV5RespDataWithDefaults() *GetPublicInstrumentTickBandsV5RespData`

NewGetPublicInstrumentTickBandsV5RespDataWithDefaults instantiates a new GetPublicInstrumentTickBandsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstFamily

`func (o *GetPublicInstrumentTickBandsV5RespData) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetPublicInstrumentTickBandsV5RespData) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetPublicInstrumentTickBandsV5RespData) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetPublicInstrumentTickBandsV5RespData) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstType

`func (o *GetPublicInstrumentTickBandsV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetPublicInstrumentTickBandsV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetPublicInstrumentTickBandsV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetPublicInstrumentTickBandsV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetTickBand

`func (o *GetPublicInstrumentTickBandsV5RespData) GetTickBand() []GetPublicInstrumentTickBandsV5RespDataTickBandInner`

GetTickBand returns the TickBand field if non-nil, zero value otherwise.

### GetTickBandOk

`func (o *GetPublicInstrumentTickBandsV5RespData) GetTickBandOk() (*[]GetPublicInstrumentTickBandsV5RespDataTickBandInner, bool)`

GetTickBandOk returns a tuple with the TickBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickBand

`func (o *GetPublicInstrumentTickBandsV5RespData) SetTickBand(v []GetPublicInstrumentTickBandsV5RespDataTickBandInner)`

SetTickBand sets TickBand field to given value.

### HasTickBand

`func (o *GetPublicInstrumentTickBandsV5RespData) HasTickBand() bool`

HasTickBand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


