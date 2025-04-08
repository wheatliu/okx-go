# GetPublicSettlementHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]GetPublicSettlementHistoryV5RespDataInnerDetailsInner**](GetPublicSettlementHistoryV5RespDataInnerDetailsInner.md) | Settlement info | [optional] 
**Ts** | Pointer to **string** | Settlement time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetPublicSettlementHistoryV5RespDataInner

`func NewGetPublicSettlementHistoryV5RespDataInner() *GetPublicSettlementHistoryV5RespDataInner`

NewGetPublicSettlementHistoryV5RespDataInner instantiates a new GetPublicSettlementHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPublicSettlementHistoryV5RespDataInnerWithDefaults

`func NewGetPublicSettlementHistoryV5RespDataInnerWithDefaults() *GetPublicSettlementHistoryV5RespDataInner`

NewGetPublicSettlementHistoryV5RespDataInnerWithDefaults instantiates a new GetPublicSettlementHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *GetPublicSettlementHistoryV5RespDataInner) GetDetails() []GetPublicSettlementHistoryV5RespDataInnerDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetPublicSettlementHistoryV5RespDataInner) GetDetailsOk() (*[]GetPublicSettlementHistoryV5RespDataInnerDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetPublicSettlementHistoryV5RespDataInner) SetDetails(v []GetPublicSettlementHistoryV5RespDataInnerDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetPublicSettlementHistoryV5RespDataInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTs

`func (o *GetPublicSettlementHistoryV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetPublicSettlementHistoryV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetPublicSettlementHistoryV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetPublicSettlementHistoryV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


