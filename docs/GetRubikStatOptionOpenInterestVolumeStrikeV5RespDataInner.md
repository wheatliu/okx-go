# GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallOI** | Pointer to **string** | Total call open interest (&#x60;coin&#x60; as the unit) | [optional] [default to ""]
**CallVol** | Pointer to **string** | Total call trading volume (&#x60;coin&#x60; as the unit) | [optional] [default to ""]
**PutOI** | Pointer to **string** | Total put open interest (&#x60;coin&#x60; as the unit) | [optional] [default to ""]
**PutVol** | Pointer to **string** | Total put trading volume (&#x60;coin&#x60; as the unit) | [optional] [default to ""]
**Strike** | Pointer to **string** | Strike price | [optional] [default to ""]
**Ts** | Pointer to **string** | Timestamp | [optional] [default to ""]

## Methods

### NewGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner

`func NewGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner() *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner`

NewGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner instantiates a new GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInnerWithDefaults

`func NewGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInnerWithDefaults() *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner`

NewGetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInnerWithDefaults instantiates a new GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallOI

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetCallOI() string`

GetCallOI returns the CallOI field if non-nil, zero value otherwise.

### GetCallOIOk

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetCallOIOk() (*string, bool)`

GetCallOIOk returns a tuple with the CallOI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallOI

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) SetCallOI(v string)`

SetCallOI sets CallOI field to given value.

### HasCallOI

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) HasCallOI() bool`

HasCallOI returns a boolean if a field has been set.

### GetCallVol

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetCallVol() string`

GetCallVol returns the CallVol field if non-nil, zero value otherwise.

### GetCallVolOk

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetCallVolOk() (*string, bool)`

GetCallVolOk returns a tuple with the CallVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallVol

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) SetCallVol(v string)`

SetCallVol sets CallVol field to given value.

### HasCallVol

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) HasCallVol() bool`

HasCallVol returns a boolean if a field has been set.

### GetPutOI

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetPutOI() string`

GetPutOI returns the PutOI field if non-nil, zero value otherwise.

### GetPutOIOk

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetPutOIOk() (*string, bool)`

GetPutOIOk returns a tuple with the PutOI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutOI

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) SetPutOI(v string)`

SetPutOI sets PutOI field to given value.

### HasPutOI

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) HasPutOI() bool`

HasPutOI returns a boolean if a field has been set.

### GetPutVol

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetPutVol() string`

GetPutVol returns the PutVol field if non-nil, zero value otherwise.

### GetPutVolOk

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetPutVolOk() (*string, bool)`

GetPutVolOk returns a tuple with the PutVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutVol

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) SetPutVol(v string)`

SetPutVol sets PutVol field to given value.

### HasPutVol

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) HasPutVol() bool`

HasPutVol returns a boolean if a field has been set.

### GetStrike

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetStrike() string`

GetStrike returns the Strike field if non-nil, zero value otherwise.

### GetStrikeOk

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetStrikeOk() (*string, bool)`

GetStrikeOk returns a tuple with the Strike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrike

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) SetStrike(v string)`

SetStrike sets Strike field to given value.

### HasStrike

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) HasStrike() bool`

HasStrike returns a boolean if a field has been set.

### GetTs

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetRubikStatOptionOpenInterestVolumeStrikeV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


