# GetPublicPriceLimitV5RespDataInner

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

### NewGetPublicPriceLimitV5RespDataInner

`func NewGetPublicPriceLimitV5RespDataInner() *GetPublicPriceLimitV5RespDataInner`

NewGetPublicPriceLimitV5RespDataInner instantiates a new GetPublicPriceLimitV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicPriceLimitV5RespDataInnerWithDefaults

`func NewGetPublicPriceLimitV5RespDataInnerWithDefaults() *GetPublicPriceLimitV5RespDataInner`

NewGetPublicPriceLimitV5RespDataInnerWithDefaults instantiates a new GetPublicPriceLimitV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyLmt

`func (o *GetPublicPriceLimitV5RespDataInner) GetBuyLmt() string`

GetBuyLmt returns the BuyLmt field if non-nil, zero value otherwise.

### GetBuyLmtOk

`func (o *GetPublicPriceLimitV5RespDataInner) GetBuyLmtOk() (*string, bool)`

GetBuyLmtOk returns a tuple with the BuyLmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyLmt

`func (o *GetPublicPriceLimitV5RespDataInner) SetBuyLmt(v string)`

SetBuyLmt sets BuyLmt field to given value.

### HasBuyLmt

`func (o *GetPublicPriceLimitV5RespDataInner) HasBuyLmt() bool`

HasBuyLmt returns a boolean if a field has been set.

### GetEnabled

`func (o *GetPublicPriceLimitV5RespDataInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetPublicPriceLimitV5RespDataInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetPublicPriceLimitV5RespDataInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetPublicPriceLimitV5RespDataInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInstId

`func (o *GetPublicPriceLimitV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicPriceLimitV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicPriceLimitV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicPriceLimitV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetPublicPriceLimitV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetPublicPriceLimitV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetPublicPriceLimitV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetPublicPriceLimitV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetSellLmt

`func (o *GetPublicPriceLimitV5RespDataInner) GetSellLmt() string`

GetSellLmt returns the SellLmt field if non-nil, zero value otherwise.

### GetSellLmtOk

`func (o *GetPublicPriceLimitV5RespDataInner) GetSellLmtOk() (*string, bool)`

GetSellLmtOk returns a tuple with the SellLmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellLmt

`func (o *GetPublicPriceLimitV5RespDataInner) SetSellLmt(v string)`

SetSellLmt sets SellLmt field to given value.

### HasSellLmt

`func (o *GetPublicPriceLimitV5RespDataInner) HasSellLmt() bool`

HasSellLmt returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicPriceLimitV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicPriceLimitV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicPriceLimitV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicPriceLimitV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


