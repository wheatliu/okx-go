# GetPublicEstimatedSettlementInfoV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstSettlePx** | Pointer to **string** | Estimated settlement price | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;XRP-USDT-250307&#x60; | [optional] [default to ""]
**NextSettleTime** | Pointer to **string** | Next settlement time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Ts** | Pointer to **string** | Data return time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicEstimatedSettlementInfoV5RespDataInner

`func NewGetPublicEstimatedSettlementInfoV5RespDataInner() *GetPublicEstimatedSettlementInfoV5RespDataInner`

NewGetPublicEstimatedSettlementInfoV5RespDataInner instantiates a new GetPublicEstimatedSettlementInfoV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicEstimatedSettlementInfoV5RespDataInnerWithDefaults

`func NewGetPublicEstimatedSettlementInfoV5RespDataInnerWithDefaults() *GetPublicEstimatedSettlementInfoV5RespDataInner`

NewGetPublicEstimatedSettlementInfoV5RespDataInnerWithDefaults instantiates a new GetPublicEstimatedSettlementInfoV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstSettlePx

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) GetEstSettlePx() string`

GetEstSettlePx returns the EstSettlePx field if non-nil, zero value otherwise.

### GetEstSettlePxOk

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) GetEstSettlePxOk() (*string, bool)`

GetEstSettlePxOk returns a tuple with the EstSettlePx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstSettlePx

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) SetEstSettlePx(v string)`

SetEstSettlePx sets EstSettlePx field to given value.

### HasEstSettlePx

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) HasEstSettlePx() bool`

HasEstSettlePx returns a boolean if a field has been set.

### GetInstId

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetNextSettleTime

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) GetNextSettleTime() string`

GetNextSettleTime returns the NextSettleTime field if non-nil, zero value otherwise.

### GetNextSettleTimeOk

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) GetNextSettleTimeOk() (*string, bool)`

GetNextSettleTimeOk returns a tuple with the NextSettleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSettleTime

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) SetNextSettleTime(v string)`

SetNextSettleTime sets NextSettleTime field to given value.

### HasNextSettleTime

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) HasNextSettleTime() bool`

HasNextSettleTime returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicEstimatedSettlementInfoV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


