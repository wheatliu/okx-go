# GetAssetExchangeListV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchId** | Pointer to **string** | Exchange ID, e.g. &#x60;did:ethr:0xfeb4f99829a9acdf52979abee87e83addf22a7e1&#x60; | [optional] [default to ""]
**ExchName** | Pointer to **string** | Exchange name, e.g. &#x60;1xbet&#x60; | [optional] [default to ""]

## Methods

### NewGetAssetExchangeListV5RespDataInner

`func NewGetAssetExchangeListV5RespDataInner() *GetAssetExchangeListV5RespDataInner`

NewGetAssetExchangeListV5RespDataInner instantiates a new GetAssetExchangeListV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssetExchangeListV5RespDataInnerWithDefaults

`func NewGetAssetExchangeListV5RespDataInnerWithDefaults() *GetAssetExchangeListV5RespDataInner`

NewGetAssetExchangeListV5RespDataInnerWithDefaults instantiates a new GetAssetExchangeListV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchId

`func (o *GetAssetExchangeListV5RespDataInner) GetExchId() string`

GetExchId returns the ExchId field if non-nil, zero value otherwise.

### GetExchIdOk

`func (o *GetAssetExchangeListV5RespDataInner) GetExchIdOk() (*string, bool)`

GetExchIdOk returns a tuple with the ExchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchId

`func (o *GetAssetExchangeListV5RespDataInner) SetExchId(v string)`

SetExchId sets ExchId field to given value.

### HasExchId

`func (o *GetAssetExchangeListV5RespDataInner) HasExchId() bool`

HasExchId returns a boolean if a field has been set.

### GetExchName

`func (o *GetAssetExchangeListV5RespDataInner) GetExchName() string`

GetExchName returns the ExchName field if non-nil, zero value otherwise.

### GetExchNameOk

`func (o *GetAssetExchangeListV5RespDataInner) GetExchNameOk() (*string, bool)`

GetExchNameOk returns a tuple with the ExchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchName

`func (o *GetAssetExchangeListV5RespDataInner) SetExchName(v string)`

SetExchName sets ExchName field to given value.

### HasExchName

`func (o *GetAssetExchangeListV5RespDataInner) HasExchName() bool`

HasExchName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


