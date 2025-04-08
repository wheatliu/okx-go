# GetCopytradingPublicLeadTradersV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataVer** | Pointer to **string** | Data version | [optional] [default to ""]
**Ranks** | Pointer to [**[]GetCopytradingPublicLeadTradersV5RespDataRanksInner**](GetCopytradingPublicLeadTradersV5RespDataRanksInner.md) | The rank information of lead traders | [optional] 
**TotalPage** | Pointer to **string** | Total number of pages | [optional] [default to ""]

## Methods

### NewGetCopytradingPublicLeadTradersV5RespData

`func NewGetCopytradingPublicLeadTradersV5RespData() *GetCopytradingPublicLeadTradersV5RespData`

NewGetCopytradingPublicLeadTradersV5RespData instantiates a new GetCopytradingPublicLeadTradersV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingPublicLeadTradersV5RespDataWithDefaults

`func NewGetCopytradingPublicLeadTradersV5RespDataWithDefaults() *GetCopytradingPublicLeadTradersV5RespData`

NewGetCopytradingPublicLeadTradersV5RespDataWithDefaults instantiates a new GetCopytradingPublicLeadTradersV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataVer

`func (o *GetCopytradingPublicLeadTradersV5RespData) GetDataVer() string`

GetDataVer returns the DataVer field if non-nil, zero value otherwise.

### GetDataVerOk

`func (o *GetCopytradingPublicLeadTradersV5RespData) GetDataVerOk() (*string, bool)`

GetDataVerOk returns a tuple with the DataVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVer

`func (o *GetCopytradingPublicLeadTradersV5RespData) SetDataVer(v string)`

SetDataVer sets DataVer field to given value.

### HasDataVer

`func (o *GetCopytradingPublicLeadTradersV5RespData) HasDataVer() bool`

HasDataVer returns a boolean if a field has been set.

### GetRanks

`func (o *GetCopytradingPublicLeadTradersV5RespData) GetRanks() []GetCopytradingPublicLeadTradersV5RespDataRanksInner`

GetRanks returns the Ranks field if non-nil, zero value otherwise.

### GetRanksOk

`func (o *GetCopytradingPublicLeadTradersV5RespData) GetRanksOk() (*[]GetCopytradingPublicLeadTradersV5RespDataRanksInner, bool)`

GetRanksOk returns a tuple with the Ranks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanks

`func (o *GetCopytradingPublicLeadTradersV5RespData) SetRanks(v []GetCopytradingPublicLeadTradersV5RespDataRanksInner)`

SetRanks sets Ranks field to given value.

### HasRanks

`func (o *GetCopytradingPublicLeadTradersV5RespData) HasRanks() bool`

HasRanks returns a boolean if a field has been set.

### GetTotalPage

`func (o *GetCopytradingPublicLeadTradersV5RespData) GetTotalPage() string`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *GetCopytradingPublicLeadTradersV5RespData) GetTotalPageOk() (*string, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *GetCopytradingPublicLeadTradersV5RespData) SetTotalPage(v string)`

SetTotalPage sets TotalPage field to given value.

### HasTotalPage

`func (o *GetCopytradingPublicLeadTradersV5RespData) HasTotalPage() bool`

HasTotalPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


