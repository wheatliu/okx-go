# GetPublicConvertContractCoinV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**Px** | Pointer to **string** | Order price | [optional] [default to ""]
**Sz** | Pointer to **string** | Quantity to buy or sell  It is quantity of contract while converting currency to contract  It is quantity of currency while contract to currency. | [optional] [default to ""]
**Type** | Pointer to **string** | Convert type   &#x60;1&#x60;: Convert currency to contract  &#x60;2&#x60;: Convert contract to currency | [optional] [default to ""]
**Unit** | Pointer to **string** | The unit of currency  &#x60;coin&#x60;  &#x60;usds&#x60;: USDT/USDC | [optional] [default to ""]

## Methods

### NewGetPublicConvertContractCoinV5RespDataInner

`func NewGetPublicConvertContractCoinV5RespDataInner() *GetPublicConvertContractCoinV5RespDataInner`

NewGetPublicConvertContractCoinV5RespDataInner instantiates a new GetPublicConvertContractCoinV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicConvertContractCoinV5RespDataInnerWithDefaults

`func NewGetPublicConvertContractCoinV5RespDataInnerWithDefaults() *GetPublicConvertContractCoinV5RespDataInner`

NewGetPublicConvertContractCoinV5RespDataInnerWithDefaults instantiates a new GetPublicConvertContractCoinV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstId

`func (o *GetPublicConvertContractCoinV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicConvertContractCoinV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicConvertContractCoinV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicConvertContractCoinV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetPx

`func (o *GetPublicConvertContractCoinV5RespDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetPublicConvertContractCoinV5RespDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetPublicConvertContractCoinV5RespDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetPublicConvertContractCoinV5RespDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSz

`func (o *GetPublicConvertContractCoinV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetPublicConvertContractCoinV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetPublicConvertContractCoinV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetPublicConvertContractCoinV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetType

`func (o *GetPublicConvertContractCoinV5RespDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetPublicConvertContractCoinV5RespDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetPublicConvertContractCoinV5RespDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetPublicConvertContractCoinV5RespDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *GetPublicConvertContractCoinV5RespDataInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetPublicConvertContractCoinV5RespDataInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetPublicConvertContractCoinV5RespDataInner) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *GetPublicConvertContractCoinV5RespDataInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


