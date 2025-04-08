# GetMarketTickerV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskPx** | Pointer to **string** | Best ask price | [optional] [default to ""]
**AskSz** | Pointer to **string** | Best ask size | [optional] [default to ""]
**BidPx** | Pointer to **string** | Best bid price | [optional] [default to ""]
**BidSz** | Pointer to **string** | Best bid size | [optional] [default to ""]
**High24h** | Pointer to **string** | Highest price in the past 24 hours | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Last** | Pointer to **string** | Last traded price | [optional] [default to ""]
**LastSz** | Pointer to **string** | Last traded size. 0 represents there is no trading volume | [optional] [default to ""]
**Low24h** | Pointer to **string** | Lowest price in the past 24 hours | [optional] [default to ""]
**Open24h** | Pointer to **string** | Open price in the past 24 hours | [optional] [default to ""]
**SodUtc0** | Pointer to **string** | Open price in the UTC 0 | [optional] [default to ""]
**SodUtc8** | Pointer to **string** | Open price in the UTC 8 | [optional] [default to ""]
**Ts** | Pointer to **string** | Ticker data generation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60;. | [optional] [default to ""]
**Vol24h** | Pointer to **string** | 24h trading volume, with a unit of &#x60;contract&#x60;.   If it is a &#x60;derivatives&#x60; contract, the value is the number of contracts.   If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in base currency. | [optional] [default to ""]
**VolCcy24h** | Pointer to **string** | 24h trading volume, with a unit of &#x60;currency&#x60;.   If it is a &#x60;derivatives&#x60; contract, the value is the number of base currency.   If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in quote currency. | [optional] [default to ""]

## Methods

### NewGetMarketTickerV5RespData

`func NewGetMarketTickerV5RespData() *GetMarketTickerV5RespData`

NewGetMarketTickerV5RespData instantiates a new GetMarketTickerV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketTickerV5RespDataWithDefaults

`func NewGetMarketTickerV5RespDataWithDefaults() *GetMarketTickerV5RespData`

NewGetMarketTickerV5RespDataWithDefaults instantiates a new GetMarketTickerV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPx

`func (o *GetMarketTickerV5RespData) GetAskPx() string`

GetAskPx returns the AskPx field if non-nil, zero value otherwise.

### GetAskPxOk

`func (o *GetMarketTickerV5RespData) GetAskPxOk() (*string, bool)`

GetAskPxOk returns a tuple with the AskPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPx

`func (o *GetMarketTickerV5RespData) SetAskPx(v string)`

SetAskPx sets AskPx field to given value.

### HasAskPx

`func (o *GetMarketTickerV5RespData) HasAskPx() bool`

HasAskPx returns a boolean if a field has been set.

### GetAskSz

`func (o *GetMarketTickerV5RespData) GetAskSz() string`

GetAskSz returns the AskSz field if non-nil, zero value otherwise.

### GetAskSzOk

`func (o *GetMarketTickerV5RespData) GetAskSzOk() (*string, bool)`

GetAskSzOk returns a tuple with the AskSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSz

`func (o *GetMarketTickerV5RespData) SetAskSz(v string)`

SetAskSz sets AskSz field to given value.

### HasAskSz

`func (o *GetMarketTickerV5RespData) HasAskSz() bool`

HasAskSz returns a boolean if a field has been set.

### GetBidPx

`func (o *GetMarketTickerV5RespData) GetBidPx() string`

GetBidPx returns the BidPx field if non-nil, zero value otherwise.

### GetBidPxOk

`func (o *GetMarketTickerV5RespData) GetBidPxOk() (*string, bool)`

GetBidPxOk returns a tuple with the BidPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPx

`func (o *GetMarketTickerV5RespData) SetBidPx(v string)`

SetBidPx sets BidPx field to given value.

### HasBidPx

`func (o *GetMarketTickerV5RespData) HasBidPx() bool`

HasBidPx returns a boolean if a field has been set.

### GetBidSz

`func (o *GetMarketTickerV5RespData) GetBidSz() string`

GetBidSz returns the BidSz field if non-nil, zero value otherwise.

### GetBidSzOk

`func (o *GetMarketTickerV5RespData) GetBidSzOk() (*string, bool)`

GetBidSzOk returns a tuple with the BidSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSz

`func (o *GetMarketTickerV5RespData) SetBidSz(v string)`

SetBidSz sets BidSz field to given value.

### HasBidSz

`func (o *GetMarketTickerV5RespData) HasBidSz() bool`

HasBidSz returns a boolean if a field has been set.

### GetHigh24h

`func (o *GetMarketTickerV5RespData) GetHigh24h() string`

GetHigh24h returns the High24h field if non-nil, zero value otherwise.

### GetHigh24hOk

`func (o *GetMarketTickerV5RespData) GetHigh24hOk() (*string, bool)`

GetHigh24hOk returns a tuple with the High24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh24h

`func (o *GetMarketTickerV5RespData) SetHigh24h(v string)`

SetHigh24h sets High24h field to given value.

### HasHigh24h

`func (o *GetMarketTickerV5RespData) HasHigh24h() bool`

HasHigh24h returns a boolean if a field has been set.

### GetInstId

`func (o *GetMarketTickerV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetMarketTickerV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetMarketTickerV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetMarketTickerV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetMarketTickerV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetMarketTickerV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetMarketTickerV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetMarketTickerV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLast

`func (o *GetMarketTickerV5RespData) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetMarketTickerV5RespData) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetMarketTickerV5RespData) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GetMarketTickerV5RespData) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLastSz

`func (o *GetMarketTickerV5RespData) GetLastSz() string`

GetLastSz returns the LastSz field if non-nil, zero value otherwise.

### GetLastSzOk

`func (o *GetMarketTickerV5RespData) GetLastSzOk() (*string, bool)`

GetLastSzOk returns a tuple with the LastSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSz

`func (o *GetMarketTickerV5RespData) SetLastSz(v string)`

SetLastSz sets LastSz field to given value.

### HasLastSz

`func (o *GetMarketTickerV5RespData) HasLastSz() bool`

HasLastSz returns a boolean if a field has been set.

### GetLow24h

`func (o *GetMarketTickerV5RespData) GetLow24h() string`

GetLow24h returns the Low24h field if non-nil, zero value otherwise.

### GetLow24hOk

`func (o *GetMarketTickerV5RespData) GetLow24hOk() (*string, bool)`

GetLow24hOk returns a tuple with the Low24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow24h

`func (o *GetMarketTickerV5RespData) SetLow24h(v string)`

SetLow24h sets Low24h field to given value.

### HasLow24h

`func (o *GetMarketTickerV5RespData) HasLow24h() bool`

HasLow24h returns a boolean if a field has been set.

### GetOpen24h

`func (o *GetMarketTickerV5RespData) GetOpen24h() string`

GetOpen24h returns the Open24h field if non-nil, zero value otherwise.

### GetOpen24hOk

`func (o *GetMarketTickerV5RespData) GetOpen24hOk() (*string, bool)`

GetOpen24hOk returns a tuple with the Open24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen24h

`func (o *GetMarketTickerV5RespData) SetOpen24h(v string)`

SetOpen24h sets Open24h field to given value.

### HasOpen24h

`func (o *GetMarketTickerV5RespData) HasOpen24h() bool`

HasOpen24h returns a boolean if a field has been set.

### GetSodUtc0

`func (o *GetMarketTickerV5RespData) GetSodUtc0() string`

GetSodUtc0 returns the SodUtc0 field if non-nil, zero value otherwise.

### GetSodUtc0Ok

`func (o *GetMarketTickerV5RespData) GetSodUtc0Ok() (*string, bool)`

GetSodUtc0Ok returns a tuple with the SodUtc0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSodUtc0

`func (o *GetMarketTickerV5RespData) SetSodUtc0(v string)`

SetSodUtc0 sets SodUtc0 field to given value.

### HasSodUtc0

`func (o *GetMarketTickerV5RespData) HasSodUtc0() bool`

HasSodUtc0 returns a boolean if a field has been set.

### GetSodUtc8

`func (o *GetMarketTickerV5RespData) GetSodUtc8() string`

GetSodUtc8 returns the SodUtc8 field if non-nil, zero value otherwise.

### GetSodUtc8Ok

`func (o *GetMarketTickerV5RespData) GetSodUtc8Ok() (*string, bool)`

GetSodUtc8Ok returns a tuple with the SodUtc8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSodUtc8

`func (o *GetMarketTickerV5RespData) SetSodUtc8(v string)`

SetSodUtc8 sets SodUtc8 field to given value.

### HasSodUtc8

`func (o *GetMarketTickerV5RespData) HasSodUtc8() bool`

HasSodUtc8 returns a boolean if a field has been set.

### GetTs

`func (o *GetMarketTickerV5RespData) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetMarketTickerV5RespData) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetMarketTickerV5RespData) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetMarketTickerV5RespData) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetVol24h

`func (o *GetMarketTickerV5RespData) GetVol24h() string`

GetVol24h returns the Vol24h field if non-nil, zero value otherwise.

### GetVol24hOk

`func (o *GetMarketTickerV5RespData) GetVol24hOk() (*string, bool)`

GetVol24hOk returns a tuple with the Vol24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol24h

`func (o *GetMarketTickerV5RespData) SetVol24h(v string)`

SetVol24h sets Vol24h field to given value.

### HasVol24h

`func (o *GetMarketTickerV5RespData) HasVol24h() bool`

HasVol24h returns a boolean if a field has been set.

### GetVolCcy24h

`func (o *GetMarketTickerV5RespData) GetVolCcy24h() string`

GetVolCcy24h returns the VolCcy24h field if non-nil, zero value otherwise.

### GetVolCcy24hOk

`func (o *GetMarketTickerV5RespData) GetVolCcy24hOk() (*string, bool)`

GetVolCcy24hOk returns a tuple with the VolCcy24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolCcy24h

`func (o *GetMarketTickerV5RespData) SetVolCcy24h(v string)`

SetVolCcy24h sets VolCcy24h field to given value.

### HasVolCcy24h

`func (o *GetMarketTickerV5RespData) HasVolCcy24h() bool`

HasVolCcy24h returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


