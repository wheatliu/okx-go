# GetMarketOptionInstrumentFamilyTradesV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptType** | Pointer to **string** | Option type, C: Call P: Put | [optional] [default to ""]
**TradeInfo** | Pointer to [**[]GetMarketOptionInstrumentFamilyTradesV5RespDataInnerTradeInfoInner**](GetMarketOptionInstrumentFamilyTradesV5RespDataInnerTradeInfoInner.md) | The list trade data | [optional] 
**Vol24h** | Pointer to **string** | 24h trading volume, with a unit of contract. | [optional] [default to ""]

## Methods

### NewGetMarketOptionInstrumentFamilyTradesV5RespDataInner

`func NewGetMarketOptionInstrumentFamilyTradesV5RespDataInner() *GetMarketOptionInstrumentFamilyTradesV5RespDataInner`

NewGetMarketOptionInstrumentFamilyTradesV5RespDataInner instantiates a new GetMarketOptionInstrumentFamilyTradesV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketOptionInstrumentFamilyTradesV5RespDataInnerWithDefaults

`func NewGetMarketOptionInstrumentFamilyTradesV5RespDataInnerWithDefaults() *GetMarketOptionInstrumentFamilyTradesV5RespDataInner`

NewGetMarketOptionInstrumentFamilyTradesV5RespDataInnerWithDefaults instantiates a new GetMarketOptionInstrumentFamilyTradesV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptType

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataInner) GetOptType() string`

GetOptType returns the OptType field if non-nil, zero value otherwise.

### GetOptTypeOk

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataInner) GetOptTypeOk() (*string, bool)`

GetOptTypeOk returns a tuple with the OptType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptType

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataInner) SetOptType(v string)`

SetOptType sets OptType field to given value.

### HasOptType

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataInner) HasOptType() bool`

HasOptType returns a boolean if a field has been set.

### GetTradeInfo

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataInner) GetTradeInfo() []GetMarketOptionInstrumentFamilyTradesV5RespDataInnerTradeInfoInner`

GetTradeInfo returns the TradeInfo field if non-nil, zero value otherwise.

### GetTradeInfoOk

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataInner) GetTradeInfoOk() (*[]GetMarketOptionInstrumentFamilyTradesV5RespDataInnerTradeInfoInner, bool)`

GetTradeInfoOk returns a tuple with the TradeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeInfo

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataInner) SetTradeInfo(v []GetMarketOptionInstrumentFamilyTradesV5RespDataInnerTradeInfoInner)`

SetTradeInfo sets TradeInfo field to given value.

### HasTradeInfo

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataInner) HasTradeInfo() bool`

HasTradeInfo returns a boolean if a field has been set.

### GetVol24h

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataInner) GetVol24h() string`

GetVol24h returns the Vol24h field if non-nil, zero value otherwise.

### GetVol24hOk

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataInner) GetVol24hOk() (*string, bool)`

GetVol24hOk returns a tuple with the Vol24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol24h

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataInner) SetVol24h(v string)`

SetVol24h sets Vol24h field to given value.

### HasVol24h

`func (o *GetMarketOptionInstrumentFamilyTradesV5RespDataInner) HasVol24h() bool`

HasVol24h returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


