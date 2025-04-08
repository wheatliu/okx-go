# GetMarketPlatform24VolumeV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **string** | Data return time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**VolCny** | Pointer to **string** | 24-hour total trading volume from the order book trading in \&quot;CNY\&quot; | [optional] [default to ""]
**VolUsd** | Pointer to **string** | 24-hour total trading volume from the order book trading in \&quot;USD\&quot; | [optional] [default to ""]

## Methods

### NewGetMarketPlatform24VolumeV5RespDataInner

`func NewGetMarketPlatform24VolumeV5RespDataInner() *GetMarketPlatform24VolumeV5RespDataInner`

NewGetMarketPlatform24VolumeV5RespDataInner instantiates a new GetMarketPlatform24VolumeV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketPlatform24VolumeV5RespDataInnerWithDefaults

`func NewGetMarketPlatform24VolumeV5RespDataInnerWithDefaults() *GetMarketPlatform24VolumeV5RespDataInner`

NewGetMarketPlatform24VolumeV5RespDataInnerWithDefaults instantiates a new GetMarketPlatform24VolumeV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *GetMarketPlatform24VolumeV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetMarketPlatform24VolumeV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetMarketPlatform24VolumeV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetMarketPlatform24VolumeV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetVolCny

`func (o *GetMarketPlatform24VolumeV5RespDataInner) GetVolCny() string`

GetVolCny returns the VolCny field if non-nil, zero value otherwise.

### GetVolCnyOk

`func (o *GetMarketPlatform24VolumeV5RespDataInner) GetVolCnyOk() (*string, bool)`

GetVolCnyOk returns a tuple with the VolCny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolCny

`func (o *GetMarketPlatform24VolumeV5RespDataInner) SetVolCny(v string)`

SetVolCny sets VolCny field to given value.

### HasVolCny

`func (o *GetMarketPlatform24VolumeV5RespDataInner) HasVolCny() bool`

HasVolCny returns a boolean if a field has been set.

### GetVolUsd

`func (o *GetMarketPlatform24VolumeV5RespDataInner) GetVolUsd() string`

GetVolUsd returns the VolUsd field if non-nil, zero value otherwise.

### GetVolUsdOk

`func (o *GetMarketPlatform24VolumeV5RespDataInner) GetVolUsdOk() (*string, bool)`

GetVolUsdOk returns a tuple with the VolUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolUsd

`func (o *GetMarketPlatform24VolumeV5RespDataInner) SetVolUsd(v string)`

SetVolUsd sets VolUsd field to given value.

### HasVolUsd

`func (o *GetMarketPlatform24VolumeV5RespDataInner) HasVolUsd() bool`

HasVolUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


