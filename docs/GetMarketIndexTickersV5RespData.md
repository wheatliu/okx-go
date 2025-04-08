# GetMarketIndexTickersV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**High24h** | Pointer to **string** | Highest price in the past 24 hours | [optional] [default to ""]
**IdxPx** | Pointer to **string** | Latest index price | [optional] [default to ""]
**InstId** | Pointer to **string** | Index | [optional] [default to ""]
**Low24h** | Pointer to **string** | Lowest price in the past 24 hours | [optional] [default to ""]
**Open24h** | Pointer to **string** | Open price in the past 24 hours | [optional] [default to ""]
**SodUtc0** | Pointer to **string** | Open price in the UTC 0 | [optional] [default to ""]
**SodUtc8** | Pointer to **string** | Open price in the UTC 8 | [optional] [default to ""]
**Ts** | Pointer to **string** | Index price update time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetMarketIndexTickersV5RespData

`func NewGetMarketIndexTickersV5RespData() *GetMarketIndexTickersV5RespData`

NewGetMarketIndexTickersV5RespData instantiates a new GetMarketIndexTickersV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketIndexTickersV5RespDataWithDefaults

`func NewGetMarketIndexTickersV5RespDataWithDefaults() *GetMarketIndexTickersV5RespData`

NewGetMarketIndexTickersV5RespDataWithDefaults instantiates a new GetMarketIndexTickersV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHigh24h

`func (o *GetMarketIndexTickersV5RespData) GetHigh24h() string`

GetHigh24h returns the High24h field if non-nil, zero value otherwise.

### GetHigh24hOk

`func (o *GetMarketIndexTickersV5RespData) GetHigh24hOk() (*string, bool)`

GetHigh24hOk returns a tuple with the High24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh24h

`func (o *GetMarketIndexTickersV5RespData) SetHigh24h(v string)`

SetHigh24h sets High24h field to given value.

### HasHigh24h

`func (o *GetMarketIndexTickersV5RespData) HasHigh24h() bool`

HasHigh24h returns a boolean if a field has been set.

### GetIdxPx

`func (o *GetMarketIndexTickersV5RespData) GetIdxPx() string`

GetIdxPx returns the IdxPx field if non-nil, zero value otherwise.

### GetIdxPxOk

`func (o *GetMarketIndexTickersV5RespData) GetIdxPxOk() (*string, bool)`

GetIdxPxOk returns a tuple with the IdxPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdxPx

`func (o *GetMarketIndexTickersV5RespData) SetIdxPx(v string)`

SetIdxPx sets IdxPx field to given value.

### HasIdxPx

`func (o *GetMarketIndexTickersV5RespData) HasIdxPx() bool`

HasIdxPx returns a boolean if a field has been set.

### GetInstId

`func (o *GetMarketIndexTickersV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetMarketIndexTickersV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetMarketIndexTickersV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetMarketIndexTickersV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLow24h

`func (o *GetMarketIndexTickersV5RespData) GetLow24h() string`

GetLow24h returns the Low24h field if non-nil, zero value otherwise.

### GetLow24hOk

`func (o *GetMarketIndexTickersV5RespData) GetLow24hOk() (*string, bool)`

GetLow24hOk returns a tuple with the Low24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow24h

`func (o *GetMarketIndexTickersV5RespData) SetLow24h(v string)`

SetLow24h sets Low24h field to given value.

### HasLow24h

`func (o *GetMarketIndexTickersV5RespData) HasLow24h() bool`

HasLow24h returns a boolean if a field has been set.

### GetOpen24h

`func (o *GetMarketIndexTickersV5RespData) GetOpen24h() string`

GetOpen24h returns the Open24h field if non-nil, zero value otherwise.

### GetOpen24hOk

`func (o *GetMarketIndexTickersV5RespData) GetOpen24hOk() (*string, bool)`

GetOpen24hOk returns a tuple with the Open24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen24h

`func (o *GetMarketIndexTickersV5RespData) SetOpen24h(v string)`

SetOpen24h sets Open24h field to given value.

### HasOpen24h

`func (o *GetMarketIndexTickersV5RespData) HasOpen24h() bool`

HasOpen24h returns a boolean if a field has been set.

### GetSodUtc0

`func (o *GetMarketIndexTickersV5RespData) GetSodUtc0() string`

GetSodUtc0 returns the SodUtc0 field if non-nil, zero value otherwise.

### GetSodUtc0Ok

`func (o *GetMarketIndexTickersV5RespData) GetSodUtc0Ok() (*string, bool)`

GetSodUtc0Ok returns a tuple with the SodUtc0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSodUtc0

`func (o *GetMarketIndexTickersV5RespData) SetSodUtc0(v string)`

SetSodUtc0 sets SodUtc0 field to given value.

### HasSodUtc0

`func (o *GetMarketIndexTickersV5RespData) HasSodUtc0() bool`

HasSodUtc0 returns a boolean if a field has been set.

### GetSodUtc8

`func (o *GetMarketIndexTickersV5RespData) GetSodUtc8() string`

GetSodUtc8 returns the SodUtc8 field if non-nil, zero value otherwise.

### GetSodUtc8Ok

`func (o *GetMarketIndexTickersV5RespData) GetSodUtc8Ok() (*string, bool)`

GetSodUtc8Ok returns a tuple with the SodUtc8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSodUtc8

`func (o *GetMarketIndexTickersV5RespData) SetSodUtc8(v string)`

SetSodUtc8 sets SodUtc8 field to given value.

### HasSodUtc8

`func (o *GetMarketIndexTickersV5RespData) HasSodUtc8() bool`

HasSodUtc8 returns a boolean if a field has been set.

### GetTs

`func (o *GetMarketIndexTickersV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetMarketIndexTickersV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetMarketIndexTickersV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetMarketIndexTickersV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


