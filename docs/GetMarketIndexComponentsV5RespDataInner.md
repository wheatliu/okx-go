# GetMarketIndexComponentsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]GetMarketIndexComponentsV5RespDataInnerComponentsInner**](GetMarketIndexComponentsV5RespDataInnerComponentsInner.md) | Components | [optional] 
**Index** | Pointer to **string** | Index | [optional] [default to ""]
**Last** | Pointer to **string** | Latest Index Price | [optional] [default to ""]
**Ts** | Pointer to **string** | Data generation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetMarketIndexComponentsV5RespDataInner

`func NewGetMarketIndexComponentsV5RespDataInner() *GetMarketIndexComponentsV5RespDataInner`

NewGetMarketIndexComponentsV5RespDataInner instantiates a new GetMarketIndexComponentsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketIndexComponentsV5RespDataInnerWithDefaults

`func NewGetMarketIndexComponentsV5RespDataInnerWithDefaults() *GetMarketIndexComponentsV5RespDataInner`

NewGetMarketIndexComponentsV5RespDataInnerWithDefaults instantiates a new GetMarketIndexComponentsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *GetMarketIndexComponentsV5RespDataInner) GetComponents() []GetMarketIndexComponentsV5RespDataInnerComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GetMarketIndexComponentsV5RespDataInner) GetComponentsOk() (*[]GetMarketIndexComponentsV5RespDataInnerComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GetMarketIndexComponentsV5RespDataInner) SetComponents(v []GetMarketIndexComponentsV5RespDataInnerComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *GetMarketIndexComponentsV5RespDataInner) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetIndex

`func (o *GetMarketIndexComponentsV5RespDataInner) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GetMarketIndexComponentsV5RespDataInner) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GetMarketIndexComponentsV5RespDataInner) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *GetMarketIndexComponentsV5RespDataInner) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLast

`func (o *GetMarketIndexComponentsV5RespDataInner) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetMarketIndexComponentsV5RespDataInner) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetMarketIndexComponentsV5RespDataInner) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GetMarketIndexComponentsV5RespDataInner) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetTs

`func (o *GetMarketIndexComponentsV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetMarketIndexComponentsV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetMarketIndexComponentsV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetMarketIndexComponentsV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


