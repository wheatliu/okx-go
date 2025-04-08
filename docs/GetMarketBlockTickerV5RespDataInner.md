# GetMarketBlockTickerV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Ts** | Pointer to **string** | Block ticker data generation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Vol24h** | Pointer to **string** | 24h trading volume, with a unit of &#x60;contract&#x60;.   If it is a &#x60;derivatives&#x60; contract, the value is the number of contracts.   If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in base currency. | [optional] [default to ""]
**VolCcy24h** | Pointer to **string** | 24h trading volume, with a unit of &#x60;currency&#x60;.   If it is a &#x60;derivatives&#x60; contract, the value is the number of base currency.   If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in quote currency. | [optional] [default to ""]

## Methods

### NewGetMarketBlockTickerV5RespDataInner

`func NewGetMarketBlockTickerV5RespDataInner() *GetMarketBlockTickerV5RespDataInner`

NewGetMarketBlockTickerV5RespDataInner instantiates a new GetMarketBlockTickerV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketBlockTickerV5RespDataInnerWithDefaults

`func NewGetMarketBlockTickerV5RespDataInnerWithDefaults() *GetMarketBlockTickerV5RespDataInner`

NewGetMarketBlockTickerV5RespDataInnerWithDefaults instantiates a new GetMarketBlockTickerV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstId

`func (o *GetMarketBlockTickerV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetMarketBlockTickerV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetMarketBlockTickerV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetMarketBlockTickerV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetMarketBlockTickerV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetMarketBlockTickerV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetMarketBlockTickerV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetMarketBlockTickerV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetTs

`func (o *GetMarketBlockTickerV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetMarketBlockTickerV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetMarketBlockTickerV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetMarketBlockTickerV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetVol24h

`func (o *GetMarketBlockTickerV5RespDataInner) GetVol24h() string`

GetVol24h returns the Vol24h field if non-nil, zero value otherwise.

### GetVol24hOk

`func (o *GetMarketBlockTickerV5RespDataInner) GetVol24hOk() (*string, bool)`

GetVol24hOk returns a tuple with the Vol24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol24h

`func (o *GetMarketBlockTickerV5RespDataInner) SetVol24h(v string)`

SetVol24h sets Vol24h field to given value.

### HasVol24h

`func (o *GetMarketBlockTickerV5RespDataInner) HasVol24h() bool`

HasVol24h returns a boolean if a field has been set.

### GetVolCcy24h

`func (o *GetMarketBlockTickerV5RespDataInner) GetVolCcy24h() string`

GetVolCcy24h returns the VolCcy24h field if non-nil, zero value otherwise.

### GetVolCcy24hOk

`func (o *GetMarketBlockTickerV5RespDataInner) GetVolCcy24hOk() (*string, bool)`

GetVolCcy24hOk returns a tuple with the VolCcy24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolCcy24h

`func (o *GetMarketBlockTickerV5RespDataInner) SetVolCcy24h(v string)`

SetVolCcy24h sets VolCcy24h field to given value.

### HasVolCcy24h

`func (o *GetMarketBlockTickerV5RespDataInner) HasVolCcy24h() bool`

HasVolCcy24h returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


