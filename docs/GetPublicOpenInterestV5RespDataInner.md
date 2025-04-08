# GetPublicOpenInterestV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Oi** | Pointer to **string** | Open interest in number of contracts | [optional] [default to ""]
**OiCcy** | Pointer to **string** | Open interest in number of coin | [optional] [default to ""]
**OiUsd** | Pointer to **string** | Open interest in number of USD | [optional] [default to ""]
**Ts** | Pointer to **string** | Data return time,  Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicOpenInterestV5RespDataInner

`func NewGetPublicOpenInterestV5RespDataInner() *GetPublicOpenInterestV5RespDataInner`

NewGetPublicOpenInterestV5RespDataInner instantiates a new GetPublicOpenInterestV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicOpenInterestV5RespDataInnerWithDefaults

`func NewGetPublicOpenInterestV5RespDataInnerWithDefaults() *GetPublicOpenInterestV5RespDataInner`

NewGetPublicOpenInterestV5RespDataInnerWithDefaults instantiates a new GetPublicOpenInterestV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstId

`func (o *GetPublicOpenInterestV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicOpenInterestV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicOpenInterestV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicOpenInterestV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetPublicOpenInterestV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetPublicOpenInterestV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetPublicOpenInterestV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetPublicOpenInterestV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetOi

`func (o *GetPublicOpenInterestV5RespDataInner) GetOi() string`

GetOi returns the Oi field if non-nil, zero value otherwise.

### GetOiOk

`func (o *GetPublicOpenInterestV5RespDataInner) GetOiOk() (*string, bool)`

GetOiOk returns a tuple with the Oi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOi

`func (o *GetPublicOpenInterestV5RespDataInner) SetOi(v string)`

SetOi sets Oi field to given value.

### HasOi

`func (o *GetPublicOpenInterestV5RespDataInner) HasOi() bool`

HasOi returns a boolean if a field has been set.

### GetOiCcy

`func (o *GetPublicOpenInterestV5RespDataInner) GetOiCcy() string`

GetOiCcy returns the OiCcy field if non-nil, zero value otherwise.

### GetOiCcyOk

`func (o *GetPublicOpenInterestV5RespDataInner) GetOiCcyOk() (*string, bool)`

GetOiCcyOk returns a tuple with the OiCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOiCcy

`func (o *GetPublicOpenInterestV5RespDataInner) SetOiCcy(v string)`

SetOiCcy sets OiCcy field to given value.

### HasOiCcy

`func (o *GetPublicOpenInterestV5RespDataInner) HasOiCcy() bool`

HasOiCcy returns a boolean if a field has been set.

### GetOiUsd

`func (o *GetPublicOpenInterestV5RespDataInner) GetOiUsd() string`

GetOiUsd returns the OiUsd field if non-nil, zero value otherwise.

### GetOiUsdOk

`func (o *GetPublicOpenInterestV5RespDataInner) GetOiUsdOk() (*string, bool)`

GetOiUsdOk returns a tuple with the OiUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOiUsd

`func (o *GetPublicOpenInterestV5RespDataInner) SetOiUsd(v string)`

SetOiUsd sets OiUsd field to given value.

### HasOiUsd

`func (o *GetPublicOpenInterestV5RespDataInner) HasOiUsd() bool`

HasOiUsd returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicOpenInterestV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicOpenInterestV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicOpenInterestV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicOpenInterestV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


