# GetPublicEstimatedPriceV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USD-200214&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [optional] [default to ""]
**SettlePx** | Pointer to **string** | Estimated delivery/exercise price | [optional] [default to ""]
**Ts** | Pointer to **string** | Data return time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicEstimatedPriceV5RespDataInner

`func NewGetPublicEstimatedPriceV5RespDataInner() *GetPublicEstimatedPriceV5RespDataInner`

NewGetPublicEstimatedPriceV5RespDataInner instantiates a new GetPublicEstimatedPriceV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicEstimatedPriceV5RespDataInnerWithDefaults

`func NewGetPublicEstimatedPriceV5RespDataInnerWithDefaults() *GetPublicEstimatedPriceV5RespDataInner`

NewGetPublicEstimatedPriceV5RespDataInnerWithDefaults instantiates a new GetPublicEstimatedPriceV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstId

`func (o *GetPublicEstimatedPriceV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicEstimatedPriceV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicEstimatedPriceV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicEstimatedPriceV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetPublicEstimatedPriceV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetPublicEstimatedPriceV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetPublicEstimatedPriceV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetPublicEstimatedPriceV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetSettlePx

`func (o *GetPublicEstimatedPriceV5RespDataInner) GetSettlePx() string`

GetSettlePx returns the SettlePx field if non-nil, zero value otherwise.

### GetSettlePxOk

`func (o *GetPublicEstimatedPriceV5RespDataInner) GetSettlePxOk() (*string, bool)`

GetSettlePxOk returns a tuple with the SettlePx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlePx

`func (o *GetPublicEstimatedPriceV5RespDataInner) SetSettlePx(v string)`

SetSettlePx sets SettlePx field to given value.

### HasSettlePx

`func (o *GetPublicEstimatedPriceV5RespDataInner) HasSettlePx() bool`

HasSettlePx returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicEstimatedPriceV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicEstimatedPriceV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicEstimatedPriceV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicEstimatedPriceV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


