# GetAccountBillsHistoryArchiveV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileHref** | Pointer to **string** | Download file link.   The expiration of every link is 5 and a half hours. If you already apply the files for the same quarter, then it don’t need to apply again within 30 days. | [optional] [default to ""]
**State** | Pointer to **string** | Download link status   \&quot;finished\&quot; \&quot;ongoing\&quot; \&quot;failed\&quot;: Failed, please apply again | [optional] [default to ""]
**Ts** | Pointer to **string** | The first request time when the server receives. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountBillsHistoryArchiveV5RespDataInner

`func NewGetAccountBillsHistoryArchiveV5RespDataInner() *GetAccountBillsHistoryArchiveV5RespDataInner`

NewGetAccountBillsHistoryArchiveV5RespDataInner instantiates a new GetAccountBillsHistoryArchiveV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountBillsHistoryArchiveV5RespDataInnerWithDefaults

`func NewGetAccountBillsHistoryArchiveV5RespDataInnerWithDefaults() *GetAccountBillsHistoryArchiveV5RespDataInner`

NewGetAccountBillsHistoryArchiveV5RespDataInnerWithDefaults instantiates a new GetAccountBillsHistoryArchiveV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileHref

`func (o *GetAccountBillsHistoryArchiveV5RespDataInner) GetFileHref() string`

GetFileHref returns the FileHref field if non-nil, zero value otherwise.

### GetFileHrefOk

`func (o *GetAccountBillsHistoryArchiveV5RespDataInner) GetFileHrefOk() (*string, bool)`

GetFileHrefOk returns a tuple with the FileHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHref

`func (o *GetAccountBillsHistoryArchiveV5RespDataInner) SetFileHref(v string)`

SetFileHref sets FileHref field to given value.

### HasFileHref

`func (o *GetAccountBillsHistoryArchiveV5RespDataInner) HasFileHref() bool`

HasFileHref returns a boolean if a field has been set.

### GetState

`func (o *GetAccountBillsHistoryArchiveV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetAccountBillsHistoryArchiveV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetAccountBillsHistoryArchiveV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetAccountBillsHistoryArchiveV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTs

`func (o *GetAccountBillsHistoryArchiveV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAccountBillsHistoryArchiveV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAccountBillsHistoryArchiveV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAccountBillsHistoryArchiveV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


