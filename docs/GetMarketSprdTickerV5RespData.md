# GetMarketSprdTickerV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskPx** | Pointer to **string** | Best ask price | [optional] [default to ""]
**AskSz** | Pointer to **string** | Best ask size | [optional] [default to ""]
**BidPx** | Pointer to **string** | Best bid price | [optional] [default to ""]
**BidSz** | Pointer to **string** | Best bid size | [optional] [default to ""]
**High24h** | Pointer to **string** | Highest price in the past 24 hours | [optional] [default to ""]
**Last** | Pointer to **string** | Last traded price | [optional] [default to ""]
**LastSz** | Pointer to **string** | Last traded size | [optional] [default to ""]
**Low24h** | Pointer to **string** | Lowest price in the past 24 hours | [optional] [default to ""]
**Open24h** | Pointer to **string** | Open price in the past 24 hours | [optional] [default to ""]
**SprdId** | Pointer to **string** | spread ID | [optional] [default to ""]
**Ts** | Pointer to **string** | Ticker data generation time, Unix timestamp format in milliseconds, e.g. 1597026383085. | [optional] [default to ""]
**Vol24h** | Pointer to **string** | 24h trading volume   The unit is USD for inverse spreads, and the corresponding baseCcy for linear and hybrid spreads. | [optional] [default to ""]

## Methods

### NewGetMarketSprdTickerV5RespData

`func NewGetMarketSprdTickerV5RespData() *GetMarketSprdTickerV5RespData`

NewGetMarketSprdTickerV5RespData instantiates a new GetMarketSprdTickerV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketSprdTickerV5RespDataWithDefaults

`func NewGetMarketSprdTickerV5RespDataWithDefaults() *GetMarketSprdTickerV5RespData`

NewGetMarketSprdTickerV5RespDataWithDefaults instantiates a new GetMarketSprdTickerV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPx

`func (o *GetMarketSprdTickerV5RespData) GetAskPx() string`

GetAskPx returns the AskPx field if non-nil, zero value otherwise.

### GetAskPxOk

`func (o *GetMarketSprdTickerV5RespData) GetAskPxOk() (*string, bool)`

GetAskPxOk returns a tuple with the AskPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPx

`func (o *GetMarketSprdTickerV5RespData) SetAskPx(v string)`

SetAskPx sets AskPx field to given value.

### HasAskPx

`func (o *GetMarketSprdTickerV5RespData) HasAskPx() bool`

HasAskPx returns a boolean if a field has been set.

### GetAskSz

`func (o *GetMarketSprdTickerV5RespData) GetAskSz() string`

GetAskSz returns the AskSz field if non-nil, zero value otherwise.

### GetAskSzOk

`func (o *GetMarketSprdTickerV5RespData) GetAskSzOk() (*string, bool)`

GetAskSzOk returns a tuple with the AskSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSz

`func (o *GetMarketSprdTickerV5RespData) SetAskSz(v string)`

SetAskSz sets AskSz field to given value.

### HasAskSz

`func (o *GetMarketSprdTickerV5RespData) HasAskSz() bool`

HasAskSz returns a boolean if a field has been set.

### GetBidPx

`func (o *GetMarketSprdTickerV5RespData) GetBidPx() string`

GetBidPx returns the BidPx field if non-nil, zero value otherwise.

### GetBidPxOk

`func (o *GetMarketSprdTickerV5RespData) GetBidPxOk() (*string, bool)`

GetBidPxOk returns a tuple with the BidPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPx

`func (o *GetMarketSprdTickerV5RespData) SetBidPx(v string)`

SetBidPx sets BidPx field to given value.

### HasBidPx

`func (o *GetMarketSprdTickerV5RespData) HasBidPx() bool`

HasBidPx returns a boolean if a field has been set.

### GetBidSz

`func (o *GetMarketSprdTickerV5RespData) GetBidSz() string`

GetBidSz returns the BidSz field if non-nil, zero value otherwise.

### GetBidSzOk

`func (o *GetMarketSprdTickerV5RespData) GetBidSzOk() (*string, bool)`

GetBidSzOk returns a tuple with the BidSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSz

`func (o *GetMarketSprdTickerV5RespData) SetBidSz(v string)`

SetBidSz sets BidSz field to given value.

### HasBidSz

`func (o *GetMarketSprdTickerV5RespData) HasBidSz() bool`

HasBidSz returns a boolean if a field has been set.

### GetHigh24h

`func (o *GetMarketSprdTickerV5RespData) GetHigh24h() string`

GetHigh24h returns the High24h field if non-nil, zero value otherwise.

### GetHigh24hOk

`func (o *GetMarketSprdTickerV5RespData) GetHigh24hOk() (*string, bool)`

GetHigh24hOk returns a tuple with the High24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh24h

`func (o *GetMarketSprdTickerV5RespData) SetHigh24h(v string)`

SetHigh24h sets High24h field to given value.

### HasHigh24h

`func (o *GetMarketSprdTickerV5RespData) HasHigh24h() bool`

HasHigh24h returns a boolean if a field has been set.

### GetLast

`func (o *GetMarketSprdTickerV5RespData) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetMarketSprdTickerV5RespData) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetMarketSprdTickerV5RespData) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GetMarketSprdTickerV5RespData) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLastSz

`func (o *GetMarketSprdTickerV5RespData) GetLastSz() string`

GetLastSz returns the LastSz field if non-nil, zero value otherwise.

### GetLastSzOk

`func (o *GetMarketSprdTickerV5RespData) GetLastSzOk() (*string, bool)`

GetLastSzOk returns a tuple with the LastSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSz

`func (o *GetMarketSprdTickerV5RespData) SetLastSz(v string)`

SetLastSz sets LastSz field to given value.

### HasLastSz

`func (o *GetMarketSprdTickerV5RespData) HasLastSz() bool`

HasLastSz returns a boolean if a field has been set.

### GetLow24h

`func (o *GetMarketSprdTickerV5RespData) GetLow24h() string`

GetLow24h returns the Low24h field if non-nil, zero value otherwise.

### GetLow24hOk

`func (o *GetMarketSprdTickerV5RespData) GetLow24hOk() (*string, bool)`

GetLow24hOk returns a tuple with the Low24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow24h

`func (o *GetMarketSprdTickerV5RespData) SetLow24h(v string)`

SetLow24h sets Low24h field to given value.

### HasLow24h

`func (o *GetMarketSprdTickerV5RespData) HasLow24h() bool`

HasLow24h returns a boolean if a field has been set.

### GetOpen24h

`func (o *GetMarketSprdTickerV5RespData) GetOpen24h() string`

GetOpen24h returns the Open24h field if non-nil, zero value otherwise.

### GetOpen24hOk

`func (o *GetMarketSprdTickerV5RespData) GetOpen24hOk() (*string, bool)`

GetOpen24hOk returns a tuple with the Open24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen24h

`func (o *GetMarketSprdTickerV5RespData) SetOpen24h(v string)`

SetOpen24h sets Open24h field to given value.

### HasOpen24h

`func (o *GetMarketSprdTickerV5RespData) HasOpen24h() bool`

HasOpen24h returns a boolean if a field has been set.

### GetSprdId

`func (o *GetMarketSprdTickerV5RespData) GetSprdId() string`

GetSprdId returns the SprdId field if non-nil, zero value otherwise.

### GetSprdIdOk

`func (o *GetMarketSprdTickerV5RespData) GetSprdIdOk() (*string, bool)`

GetSprdIdOk returns a tuple with the SprdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSprdId

`func (o *GetMarketSprdTickerV5RespData) SetSprdId(v string)`

SetSprdId sets SprdId field to given value.

### HasSprdId

`func (o *GetMarketSprdTickerV5RespData) HasSprdId() bool`

HasSprdId returns a boolean if a field has been set.

### GetTs

`func (o *GetMarketSprdTickerV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetMarketSprdTickerV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetMarketSprdTickerV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetMarketSprdTickerV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetVol24h

`func (o *GetMarketSprdTickerV5RespData) GetVol24h() string`

GetVol24h returns the Vol24h field if non-nil, zero value otherwise.

### GetVol24hOk

`func (o *GetMarketSprdTickerV5RespData) GetVol24hOk() (*string, bool)`

GetVol24hOk returns a tuple with the Vol24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol24h

`func (o *GetMarketSprdTickerV5RespData) SetVol24h(v string)`

SetVol24h sets Vol24h field to given value.

### HasVol24h

`func (o *GetMarketSprdTickerV5RespData) HasVol24h() bool`

HasVol24h returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


