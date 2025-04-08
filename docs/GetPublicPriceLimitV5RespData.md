# GetPublicPriceLimitV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyLmt** | Pointer to **string** | Highest buy limit   Return \&quot;\&quot; when enabled is false | [optional] [default to ""]
**Enabled** | Pointer to **bool** | Whether price limit is effective   &#x60;true&#x60;: the price limit is effective   &#x60;false&#x60;: the price limit is not effective | [optional] 
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**SellLmt** | Pointer to **string** | Lowest sell limit   Return \&quot;\&quot; when enabled is false | [optional] [default to ""]
**Ts** | Pointer to **string** | Data return time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicPriceLimitV5RespData

`func NewGetPublicPriceLimitV5RespData() *GetPublicPriceLimitV5RespData`

NewGetPublicPriceLimitV5RespData instantiates a new GetPublicPriceLimitV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicPriceLimitV5RespDataWithDefaults

`func NewGetPublicPriceLimitV5RespDataWithDefaults() *GetPublicPriceLimitV5RespData`

NewGetPublicPriceLimitV5RespDataWithDefaults instantiates a new GetPublicPriceLimitV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyLmt

`func (o *GetPublicPriceLimitV5RespData) GetBuyLmt() string`

GetBuyLmt returns the BuyLmt field if non-nil, zero value otherwise.

### GetBuyLmtOk

`func (o *GetPublicPriceLimitV5RespData) GetBuyLmtOk() (*string, bool)`

GetBuyLmtOk returns a tuple with the BuyLmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyLmt

`func (o *GetPublicPriceLimitV5RespData) SetBuyLmt(v string)`

SetBuyLmt sets BuyLmt field to given value.

### HasBuyLmt

`func (o *GetPublicPriceLimitV5RespData) HasBuyLmt() bool`

HasBuyLmt returns a boolean if a field has been set.

### GetEnabled

`func (o *GetPublicPriceLimitV5RespData) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetPublicPriceLimitV5RespData) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetPublicPriceLimitV5RespData) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetPublicPriceLimitV5RespData) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInstId

`func (o *GetPublicPriceLimitV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicPriceLimitV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicPriceLimitV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicPriceLimitV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetPublicPriceLimitV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetPublicPriceLimitV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetPublicPriceLimitV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetPublicPriceLimitV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetSellLmt

`func (o *GetPublicPriceLimitV5RespData) GetSellLmt() string`

GetSellLmt returns the SellLmt field if non-nil, zero value otherwise.

### GetSellLmtOk

`func (o *GetPublicPriceLimitV5RespData) GetSellLmtOk() (*string, bool)`

GetSellLmtOk returns a tuple with the SellLmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellLmt

`func (o *GetPublicPriceLimitV5RespData) SetSellLmt(v string)`

SetSellLmt sets SellLmt field to given value.

### HasSellLmt

`func (o *GetPublicPriceLimitV5RespData) HasSellLmt() bool`

HasSellLmt returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicPriceLimitV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicPriceLimitV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicPriceLimitV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicPriceLimitV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


