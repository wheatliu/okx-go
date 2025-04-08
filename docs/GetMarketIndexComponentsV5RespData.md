# GetMarketIndexComponentsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]GetMarketIndexComponentsV5RespDataComponentsInner**](GetMarketIndexComponentsV5RespDataComponentsInner.md) | Components | [optional] 
**Index** | Pointer to **string** | Index | [optional] [default to ""]
**Last** | Pointer to **string** | Latest Index Price | [optional] [default to ""]
**Ts** | Pointer to **string** | Data generation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetMarketIndexComponentsV5RespData

`func NewGetMarketIndexComponentsV5RespData() *GetMarketIndexComponentsV5RespData`

NewGetMarketIndexComponentsV5RespData instantiates a new GetMarketIndexComponentsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketIndexComponentsV5RespDataWithDefaults

`func NewGetMarketIndexComponentsV5RespDataWithDefaults() *GetMarketIndexComponentsV5RespData`

NewGetMarketIndexComponentsV5RespDataWithDefaults instantiates a new GetMarketIndexComponentsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *GetMarketIndexComponentsV5RespData) GetComponents() []GetMarketIndexComponentsV5RespDataComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GetMarketIndexComponentsV5RespData) GetComponentsOk() (*[]GetMarketIndexComponentsV5RespDataComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GetMarketIndexComponentsV5RespData) SetComponents(v []GetMarketIndexComponentsV5RespDataComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *GetMarketIndexComponentsV5RespData) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetIndex

`func (o *GetMarketIndexComponentsV5RespData) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *GetMarketIndexComponentsV5RespData) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *GetMarketIndexComponentsV5RespData) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *GetMarketIndexComponentsV5RespData) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLast

`func (o *GetMarketIndexComponentsV5RespData) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetMarketIndexComponentsV5RespData) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetMarketIndexComponentsV5RespData) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GetMarketIndexComponentsV5RespData) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetTs

`func (o *GetMarketIndexComponentsV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetMarketIndexComponentsV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetMarketIndexComponentsV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetMarketIndexComponentsV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


