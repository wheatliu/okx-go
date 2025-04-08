# GetSprdSpreadsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCcy** | Pointer to **string** | Currency instrument is based in. Valid values include BTC, ETH | [optional] [default to ""]
**ExpTime** | Pointer to **string** | Expiry time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Legs** | Pointer to [**[]GetSprdSpreadsV5RespDataLegsInner**](GetSprdSpreadsV5RespDataLegsInner.md) |  | [optional] 
**ListTime** | Pointer to **string** | The timestamp the spread was created. Unix timestamp format in milliseconds, , e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**LotSz** | Pointer to **string** | The minimum order size increment the spread can be traded in the szCcy of the spread. | [optional] [default to ""]
**MinSz** | Pointer to **string** | Minimum order size in the szCcy of the spread. | [optional] [default to ""]
**QuoteCcy** | Pointer to **string** | The currency the spread is priced in, e.g. USDT, USD | [optional] [default to ""]
**SprdId** | Pointer to **string** | spread ID | [optional] [default to ""]
**SprdType** | Pointer to **string** | spread Type. Valid values are &#x60;linear&#x60;, &#x60;inverse&#x60;, &#x60;hybrid&#x60; | [optional] [default to ""]
**State** | Pointer to **string** | Current state of the spread. Valid values include &#x60;live&#x60;, &#x60;expired&#x60;, &#x60;suspend&#x60;. | [optional] [default to ""]
**SzCcy** | Pointer to **string** | The currency the spread order size is submitted to the underlying venue in, e.g. USD, BTC, ETH. | [optional] [default to ""]
**TickSz** | Pointer to **string** | Tick size, e.g. 0.0001 in the quoteCcy of the spread. | [optional] [default to ""]
**UTime** | Pointer to **string** | The timestamp the spread was last updated. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]

## Methods

### NewGetSprdSpreadsV5RespData

`func NewGetSprdSpreadsV5RespData() *GetSprdSpreadsV5RespData`

NewGetSprdSpreadsV5RespData instantiates a new GetSprdSpreadsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSprdSpreadsV5RespDataWithDefaults

`func NewGetSprdSpreadsV5RespDataWithDefaults() *GetSprdSpreadsV5RespData`

NewGetSprdSpreadsV5RespDataWithDefaults instantiates a new GetSprdSpreadsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCcy

`func (o *GetSprdSpreadsV5RespData) GetBaseCcy() string`

GetBaseCcy returns the BaseCcy field if non-nil, zero value otherwise.

### GetBaseCcyOk

`func (o *GetSprdSpreadsV5RespData) GetBaseCcyOk() (*string, bool)`

GetBaseCcyOk returns a tuple with the BaseCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCcy

`func (o *GetSprdSpreadsV5RespData) SetBaseCcy(v string)`

SetBaseCcy sets BaseCcy field to given value.

### HasBaseCcy

`func (o *GetSprdSpreadsV5RespData) HasBaseCcy() bool`

HasBaseCcy returns a boolean if a field has been set.

### GetExpTime

`func (o *GetSprdSpreadsV5RespData) GetExpTime() string`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *GetSprdSpreadsV5RespData) GetExpTimeOk() (*string, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *GetSprdSpreadsV5RespData) SetExpTime(v string)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *GetSprdSpreadsV5RespData) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetLegs

`func (o *GetSprdSpreadsV5RespData) GetLegs() []GetSprdSpreadsV5RespDataLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *GetSprdSpreadsV5RespData) GetLegsOk() (*[]GetSprdSpreadsV5RespDataLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *GetSprdSpreadsV5RespData) SetLegs(v []GetSprdSpreadsV5RespDataLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *GetSprdSpreadsV5RespData) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetListTime

`func (o *GetSprdSpreadsV5RespData) GetListTime() string`

GetListTime returns the ListTime field if non-nil, zero value otherwise.

### GetListTimeOk

`func (o *GetSprdSpreadsV5RespData) GetListTimeOk() (*string, bool)`

GetListTimeOk returns a tuple with the ListTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListTime

`func (o *GetSprdSpreadsV5RespData) SetListTime(v string)`

SetListTime sets ListTime field to given value.

### HasListTime

`func (o *GetSprdSpreadsV5RespData) HasListTime() bool`

HasListTime returns a boolean if a field has been set.

### GetLotSz

`func (o *GetSprdSpreadsV5RespData) GetLotSz() string`

GetLotSz returns the LotSz field if non-nil, zero value otherwise.

### GetLotSzOk

`func (o *GetSprdSpreadsV5RespData) GetLotSzOk() (*string, bool)`

GetLotSzOk returns a tuple with the LotSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotSz

`func (o *GetSprdSpreadsV5RespData) SetLotSz(v string)`

SetLotSz sets LotSz field to given value.

### HasLotSz

`func (o *GetSprdSpreadsV5RespData) HasLotSz() bool`

HasLotSz returns a boolean if a field has been set.

### GetMinSz

`func (o *GetSprdSpreadsV5RespData) GetMinSz() string`

GetMinSz returns the MinSz field if non-nil, zero value otherwise.

### GetMinSzOk

`func (o *GetSprdSpreadsV5RespData) GetMinSzOk() (*string, bool)`

GetMinSzOk returns a tuple with the MinSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSz

`func (o *GetSprdSpreadsV5RespData) SetMinSz(v string)`

SetMinSz sets MinSz field to given value.

### HasMinSz

`func (o *GetSprdSpreadsV5RespData) HasMinSz() bool`

HasMinSz returns a boolean if a field has been set.

### GetQuoteCcy

`func (o *GetSprdSpreadsV5RespData) GetQuoteCcy() string`

GetQuoteCcy returns the QuoteCcy field if non-nil, zero value otherwise.

### GetQuoteCcyOk

`func (o *GetSprdSpreadsV5RespData) GetQuoteCcyOk() (*string, bool)`

GetQuoteCcyOk returns a tuple with the QuoteCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCcy

`func (o *GetSprdSpreadsV5RespData) SetQuoteCcy(v string)`

SetQuoteCcy sets QuoteCcy field to given value.

### HasQuoteCcy

`func (o *GetSprdSpreadsV5RespData) HasQuoteCcy() bool`

HasQuoteCcy returns a boolean if a field has been set.

### GetSprdId

`func (o *GetSprdSpreadsV5RespData) GetSprdId() string`

GetSprdId returns the SprdId field if non-nil, zero value otherwise.

### GetSprdIdOk

`func (o *GetSprdSpreadsV5RespData) GetSprdIdOk() (*string, bool)`

GetSprdIdOk returns a tuple with the SprdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSprdId

`func (o *GetSprdSpreadsV5RespData) SetSprdId(v string)`

SetSprdId sets SprdId field to given value.

### HasSprdId

`func (o *GetSprdSpreadsV5RespData) HasSprdId() bool`

HasSprdId returns a boolean if a field has been set.

### GetSprdType

`func (o *GetSprdSpreadsV5RespData) GetSprdType() string`

GetSprdType returns the SprdType field if non-nil, zero value otherwise.

### GetSprdTypeOk

`func (o *GetSprdSpreadsV5RespData) GetSprdTypeOk() (*string, bool)`

GetSprdTypeOk returns a tuple with the SprdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSprdType

`func (o *GetSprdSpreadsV5RespData) SetSprdType(v string)`

SetSprdType sets SprdType field to given value.

### HasSprdType

`func (o *GetSprdSpreadsV5RespData) HasSprdType() bool`

HasSprdType returns a boolean if a field has been set.

### GetState

`func (o *GetSprdSpreadsV5RespData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetSprdSpreadsV5RespData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetSprdSpreadsV5RespData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetSprdSpreadsV5RespData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSzCcy

`func (o *GetSprdSpreadsV5RespData) GetSzCcy() string`

GetSzCcy returns the SzCcy field if non-nil, zero value otherwise.

### GetSzCcyOk

`func (o *GetSprdSpreadsV5RespData) GetSzCcyOk() (*string, bool)`

GetSzCcyOk returns a tuple with the SzCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSzCcy

`func (o *GetSprdSpreadsV5RespData) SetSzCcy(v string)`

SetSzCcy sets SzCcy field to given value.

### HasSzCcy

`func (o *GetSprdSpreadsV5RespData) HasSzCcy() bool`

HasSzCcy returns a boolean if a field has been set.

### GetTickSz

`func (o *GetSprdSpreadsV5RespData) GetTickSz() string`

GetTickSz returns the TickSz field if non-nil, zero value otherwise.

### GetTickSzOk

`func (o *GetSprdSpreadsV5RespData) GetTickSzOk() (*string, bool)`

GetTickSzOk returns a tuple with the TickSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickSz

`func (o *GetSprdSpreadsV5RespData) SetTickSz(v string)`

SetTickSz sets TickSz field to given value.

### HasTickSz

`func (o *GetSprdSpreadsV5RespData) HasTickSz() bool`

HasTickSz returns a boolean if a field has been set.

### GetUTime

`func (o *GetSprdSpreadsV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetSprdSpreadsV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetSprdSpreadsV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetSprdSpreadsV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


