# GetPublicSettlementHistoryV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]GetPublicSettlementHistoryV5RespDataDetailsInner**](GetPublicSettlementHistoryV5RespDataDetailsInner.md) | Settlement info | [optional] 
**Ts** | Pointer to **string** | Settlement time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicSettlementHistoryV5RespData

`func NewGetPublicSettlementHistoryV5RespData() *GetPublicSettlementHistoryV5RespData`

NewGetPublicSettlementHistoryV5RespData instantiates a new GetPublicSettlementHistoryV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicSettlementHistoryV5RespDataWithDefaults

`func NewGetPublicSettlementHistoryV5RespDataWithDefaults() *GetPublicSettlementHistoryV5RespData`

NewGetPublicSettlementHistoryV5RespDataWithDefaults instantiates a new GetPublicSettlementHistoryV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *GetPublicSettlementHistoryV5RespData) GetDetails() []GetPublicSettlementHistoryV5RespDataDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetPublicSettlementHistoryV5RespData) GetDetailsOk() (*[]GetPublicSettlementHistoryV5RespDataDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetPublicSettlementHistoryV5RespData) SetDetails(v []GetPublicSettlementHistoryV5RespDataDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetPublicSettlementHistoryV5RespData) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicSettlementHistoryV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicSettlementHistoryV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicSettlementHistoryV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicSettlementHistoryV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


