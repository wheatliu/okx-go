# GetPublicPremiumHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [optional] [default to ""]
**Premium** | Pointer to **string** | Premium between the mid price of perps market and the index price | [optional] [default to ""]
**Ts** | Pointer to **string** | Data generation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicPremiumHistoryV5RespDataInner

`func NewGetPublicPremiumHistoryV5RespDataInner() *GetPublicPremiumHistoryV5RespDataInner`

NewGetPublicPremiumHistoryV5RespDataInner instantiates a new GetPublicPremiumHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicPremiumHistoryV5RespDataInnerWithDefaults

`func NewGetPublicPremiumHistoryV5RespDataInnerWithDefaults() *GetPublicPremiumHistoryV5RespDataInner`

NewGetPublicPremiumHistoryV5RespDataInnerWithDefaults instantiates a new GetPublicPremiumHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstId

`func (o *GetPublicPremiumHistoryV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicPremiumHistoryV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicPremiumHistoryV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicPremiumHistoryV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetPremium

`func (o *GetPublicPremiumHistoryV5RespDataInner) GetPremium() string`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *GetPublicPremiumHistoryV5RespDataInner) GetPremiumOk() (*string, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *GetPublicPremiumHistoryV5RespDataInner) SetPremium(v string)`

SetPremium sets Premium field to given value.

### HasPremium

`func (o *GetPublicPremiumHistoryV5RespDataInner) HasPremium() bool`

HasPremium returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicPremiumHistoryV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicPremiumHistoryV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicPremiumHistoryV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicPremiumHistoryV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


