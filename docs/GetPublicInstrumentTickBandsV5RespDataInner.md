# GetPublicInstrumentTickBandsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstFamily** | Pointer to **string** | Instrument family | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**TickBand** | Pointer to [**[]GetPublicInstrumentTickBandsV5RespDataInnerTickBandInner**](GetPublicInstrumentTickBandsV5RespDataInnerTickBandInner.md) | Tick size band | [optional] 

## Methods

### NewGetPublicInstrumentTickBandsV5RespDataInner

`func NewGetPublicInstrumentTickBandsV5RespDataInner() *GetPublicInstrumentTickBandsV5RespDataInner`

NewGetPublicInstrumentTickBandsV5RespDataInner instantiates a new GetPublicInstrumentTickBandsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicInstrumentTickBandsV5RespDataInnerWithDefaults

`func NewGetPublicInstrumentTickBandsV5RespDataInnerWithDefaults() *GetPublicInstrumentTickBandsV5RespDataInner`

NewGetPublicInstrumentTickBandsV5RespDataInnerWithDefaults instantiates a new GetPublicInstrumentTickBandsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstFamily

`func (o *GetPublicInstrumentTickBandsV5RespDataInner) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetPublicInstrumentTickBandsV5RespDataInner) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetPublicInstrumentTickBandsV5RespDataInner) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetPublicInstrumentTickBandsV5RespDataInner) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetInstType

`func (o *GetPublicInstrumentTickBandsV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetPublicInstrumentTickBandsV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetPublicInstrumentTickBandsV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetPublicInstrumentTickBandsV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetTickBand

`func (o *GetPublicInstrumentTickBandsV5RespDataInner) GetTickBand() []GetPublicInstrumentTickBandsV5RespDataInnerTickBandInner`

GetTickBand returns the TickBand field if non-nil, zero value otherwise.

### GetTickBandOk

`func (o *GetPublicInstrumentTickBandsV5RespDataInner) GetTickBandOk() (*[]GetPublicInstrumentTickBandsV5RespDataInnerTickBandInner, bool)`

GetTickBandOk returns a tuple with the TickBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickBand

`func (o *GetPublicInstrumentTickBandsV5RespDataInner) SetTickBand(v []GetPublicInstrumentTickBandsV5RespDataInnerTickBandInner)`

SetTickBand sets TickBand field to given value.

### HasTickBand

`func (o *GetPublicInstrumentTickBandsV5RespDataInner) HasTickBand() bool`

HasTickBand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


