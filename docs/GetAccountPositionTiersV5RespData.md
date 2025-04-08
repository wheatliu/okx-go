# GetAccountPositionTiersV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstFamily** | Pointer to **string** | Instrument family  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]
**MaxSz** | Pointer to **string** | Max number of positions | [optional] [default to ""]
**PosType** | Pointer to **string** | Limitation of position type, only applicable to cross &#x60;OPTION&#x60; under portfolio margin mode   &#x60;1&#x60;: Contracts of pending orders and open positions for all derivatives instruments. &#x60;2&#x60;: Contracts of pending orders for all derivatives instruments. &#x60;3&#x60;: Pending orders for all derivatives instruments. &#x60;4&#x60;: Contracts of pending orders and open positions for all derivatives instruments on the same side. &#x60;5&#x60;: Pending orders for one derivatives instrument. &#x60;6&#x60;: Contracts of pending orders and open positions for one derivatives instrument. &#x60;7&#x60;: Contracts of one pending order. | [optional] [default to ""]
**Uly** | Pointer to **string** | Underlying  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountPositionTiersV5RespData

`func NewGetAccountPositionTiersV5RespData() *GetAccountPositionTiersV5RespData`

NewGetAccountPositionTiersV5RespData instantiates a new GetAccountPositionTiersV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountPositionTiersV5RespDataWithDefaults

`func NewGetAccountPositionTiersV5RespDataWithDefaults() *GetAccountPositionTiersV5RespData`

NewGetAccountPositionTiersV5RespDataWithDefaults instantiates a new GetAccountPositionTiersV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstFamily

`func (o *GetAccountPositionTiersV5RespData) GetInstFamily() string`

GetInstFamily returns the InstFamily field if non-nil, zero value otherwise.

### GetInstFamilyOk

`func (o *GetAccountPositionTiersV5RespData) GetInstFamilyOk() (*string, bool)`

GetInstFamilyOk returns a tuple with the InstFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstFamily

`func (o *GetAccountPositionTiersV5RespData) SetInstFamily(v string)`

SetInstFamily sets InstFamily field to given value.

### HasInstFamily

`func (o *GetAccountPositionTiersV5RespData) HasInstFamily() bool`

HasInstFamily returns a boolean if a field has been set.

### GetMaxSz

`func (o *GetAccountPositionTiersV5RespData) GetMaxSz() string`

GetMaxSz returns the MaxSz field if non-nil, zero value otherwise.

### GetMaxSzOk

`func (o *GetAccountPositionTiersV5RespData) GetMaxSzOk() (*string, bool)`

GetMaxSzOk returns a tuple with the MaxSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSz

`func (o *GetAccountPositionTiersV5RespData) SetMaxSz(v string)`

SetMaxSz sets MaxSz field to given value.

### HasMaxSz

`func (o *GetAccountPositionTiersV5RespData) HasMaxSz() bool`

HasMaxSz returns a boolean if a field has been set.

### GetPosType

`func (o *GetAccountPositionTiersV5RespData) GetPosType() string`

GetPosType returns the PosType field if non-nil, zero value otherwise.

### GetPosTypeOk

`func (o *GetAccountPositionTiersV5RespData) GetPosTypeOk() (*string, bool)`

GetPosTypeOk returns a tuple with the PosType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosType

`func (o *GetAccountPositionTiersV5RespData) SetPosType(v string)`

SetPosType sets PosType field to given value.

### HasPosType

`func (o *GetAccountPositionTiersV5RespData) HasPosType() bool`

HasPosType returns a boolean if a field has been set.

### GetUly

`func (o *GetAccountPositionTiersV5RespData) GetUly() string`

GetUly returns the Uly field if non-nil, zero value otherwise.

### GetUlyOk

`func (o *GetAccountPositionTiersV5RespData) GetUlyOk() (*string, bool)`

GetUlyOk returns a tuple with the Uly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUly

`func (o *GetAccountPositionTiersV5RespData) SetUly(v string)`

SetUly sets Uly field to given value.

### HasUly

`func (o *GetAccountPositionTiersV5RespData) HasUly() bool`

HasUly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


