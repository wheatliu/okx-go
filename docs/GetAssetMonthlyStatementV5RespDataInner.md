# GetAssetMonthlyStatementV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileHref** | Pointer to **string** | Download file link | [optional] [default to ""]
**State** | Pointer to **string** | Download link status   \&quot;finished\&quot; \&quot;ongoing\&quot; | [optional] [default to ""]
**Ts** | Pointer to **int32** | Download link generation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] 

## Methods

### NewGetAssetMonthlyStatementV5RespDataInner

`func NewGetAssetMonthlyStatementV5RespDataInner() *GetAssetMonthlyStatementV5RespDataInner`

NewGetAssetMonthlyStatementV5RespDataInner instantiates a new GetAssetMonthlyStatementV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetMonthlyStatementV5RespDataInnerWithDefaults

`func NewGetAssetMonthlyStatementV5RespDataInnerWithDefaults() *GetAssetMonthlyStatementV5RespDataInner`

NewGetAssetMonthlyStatementV5RespDataInnerWithDefaults instantiates a new GetAssetMonthlyStatementV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileHref

`func (o *GetAssetMonthlyStatementV5RespDataInner) GetFileHref() string`

GetFileHref returns the FileHref field if non-nil, zero value otherwise.

### GetFileHrefOk

`func (o *GetAssetMonthlyStatementV5RespDataInner) GetFileHrefOk() (*string, bool)`

GetFileHrefOk returns a tuple with the FileHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHref

`func (o *GetAssetMonthlyStatementV5RespDataInner) SetFileHref(v string)`

SetFileHref sets FileHref field to given value.

### HasFileHref

`func (o *GetAssetMonthlyStatementV5RespDataInner) HasFileHref() bool`

HasFileHref returns a boolean if a field has been set.

### GetState

`func (o *GetAssetMonthlyStatementV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetAssetMonthlyStatementV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetAssetMonthlyStatementV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetAssetMonthlyStatementV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTs

`func (o *GetAssetMonthlyStatementV5RespDataInner) GetTs() int32`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAssetMonthlyStatementV5RespDataInner) GetTsOk() (*int32, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAssetMonthlyStatementV5RespDataInner) SetTs(v int32)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAssetMonthlyStatementV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


