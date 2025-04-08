# GetAssetAssetValuationV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**GetAssetAssetValuationV5RespDataInnerDetails**](GetAssetAssetValuationV5RespDataInnerDetails.md) |  | [optional] 
**TotalBal** | Pointer to **string** | Valuation of total account assets | [optional] [default to ""]
**Ts** | Pointer to **string** | Unix timestamp format in milliseconds, e.g.&#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetAssetAssetValuationV5RespDataInner

`func NewGetAssetAssetValuationV5RespDataInner() *GetAssetAssetValuationV5RespDataInner`

NewGetAssetAssetValuationV5RespDataInner instantiates a new GetAssetAssetValuationV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetAssetValuationV5RespDataInnerWithDefaults

`func NewGetAssetAssetValuationV5RespDataInnerWithDefaults() *GetAssetAssetValuationV5RespDataInner`

NewGetAssetAssetValuationV5RespDataInnerWithDefaults instantiates a new GetAssetAssetValuationV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *GetAssetAssetValuationV5RespDataInner) GetDetails() GetAssetAssetValuationV5RespDataInnerDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetAssetAssetValuationV5RespDataInner) GetDetailsOk() (*GetAssetAssetValuationV5RespDataInnerDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetAssetAssetValuationV5RespDataInner) SetDetails(v GetAssetAssetValuationV5RespDataInnerDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetAssetAssetValuationV5RespDataInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTotalBal

`func (o *GetAssetAssetValuationV5RespDataInner) GetTotalBal() string`

GetTotalBal returns the TotalBal field if non-nil, zero value otherwise.

### GetTotalBalOk

`func (o *GetAssetAssetValuationV5RespDataInner) GetTotalBalOk() (*string, bool)`

GetTotalBalOk returns a tuple with the TotalBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBal

`func (o *GetAssetAssetValuationV5RespDataInner) SetTotalBal(v string)`

SetTotalBal sets TotalBal field to given value.

### HasTotalBal

`func (o *GetAssetAssetValuationV5RespDataInner) HasTotalBal() bool`

HasTotalBal returns a boolean if a field has been set.

### GetTs

`func (o *GetAssetAssetValuationV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAssetAssetValuationV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAssetAssetValuationV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAssetAssetValuationV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


