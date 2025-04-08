# GetRubikStatOptionOpenInterestVolumeExpiryV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallOI** | Pointer to **string** | Total call open interest (&#x60;coin&#x60; as the unit) | [optional] [default to ""]
**CallVol** | Pointer to **string** | Total call trading volume (&#x60;coin&#x60; as the unit) | [optional] [default to ""]
**ExpTime** | Pointer to **string** | Contract expiry date, the format is &#x60;YYYYMMDD&#x60;, e.g. &#x60;20210623&#x60; | [optional] [default to ""]
**PutOI** | Pointer to **string** | Total put  open interest (&#x60;coin&#x60; as the unit) | [optional] [default to ""]
**PutVol** | Pointer to **string** | Total put trading volume (&#x60;coin&#x60; as the unit) | [optional] [default to ""]
**Ts** | Pointer to **string** | Timestamp | [optional] [default to ""]

## Methods

### NewGetRubikStatOptionOpenInterestVolumeExpiryV5RespData

`func NewGetRubikStatOptionOpenInterestVolumeExpiryV5RespData() *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData`

NewGetRubikStatOptionOpenInterestVolumeExpiryV5RespData instantiates a new GetRubikStatOptionOpenInterestVolumeExpiryV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRubikStatOptionOpenInterestVolumeExpiryV5RespDataWithDefaults

`func NewGetRubikStatOptionOpenInterestVolumeExpiryV5RespDataWithDefaults() *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData`

NewGetRubikStatOptionOpenInterestVolumeExpiryV5RespDataWithDefaults instantiates a new GetRubikStatOptionOpenInterestVolumeExpiryV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallOI

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) GetCallOI() string`

GetCallOI returns the CallOI field if non-nil, zero value otherwise.

### GetCallOIOk

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) GetCallOIOk() (*string, bool)`

GetCallOIOk returns a tuple with the CallOI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallOI

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) SetCallOI(v string)`

SetCallOI sets CallOI field to given value.

### HasCallOI

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) HasCallOI() bool`

HasCallOI returns a boolean if a field has been set.

### GetCallVol

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) GetCallVol() string`

GetCallVol returns the CallVol field if non-nil, zero value otherwise.

### GetCallVolOk

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) GetCallVolOk() (*string, bool)`

GetCallVolOk returns a tuple with the CallVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallVol

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) SetCallVol(v string)`

SetCallVol sets CallVol field to given value.

### HasCallVol

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) HasCallVol() bool`

HasCallVol returns a boolean if a field has been set.

### GetExpTime

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) GetExpTime() string`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) GetExpTimeOk() (*string, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) SetExpTime(v string)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetPutOI

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) GetPutOI() string`

GetPutOI returns the PutOI field if non-nil, zero value otherwise.

### GetPutOIOk

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) GetPutOIOk() (*string, bool)`

GetPutOIOk returns a tuple with the PutOI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutOI

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) SetPutOI(v string)`

SetPutOI sets PutOI field to given value.

### HasPutOI

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) HasPutOI() bool`

HasPutOI returns a boolean if a field has been set.

### GetPutVol

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) GetPutVol() string`

GetPutVol returns the PutVol field if non-nil, zero value otherwise.

### GetPutVolOk

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) GetPutVolOk() (*string, bool)`

GetPutVolOk returns a tuple with the PutVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutVol

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) SetPutVol(v string)`

SetPutVol sets PutVol field to given value.

### HasPutVol

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) HasPutVol() bool`

HasPutVol returns a boolean if a field has been set.

### GetTs

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetRubikStatOptionOpenInterestVolumeExpiryV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


