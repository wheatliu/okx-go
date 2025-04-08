# GetMarketTickersV5RespDataInner

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
**Ts** | Pointer to **string** | Ticker data generation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Vol24h** | Pointer to **string** | 24h trading volume, with a unit of &#x60;contract&#x60;.   If it is a &#x60;derivatives&#x60; contract, the value is the number of contracts.   If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in base currency. | [optional] [default to ""]
**VolCcy24h** | Pointer to **string** | 24h trading volume, with a unit of &#x60;currency&#x60;.   If it is a &#x60;derivatives&#x60; contract, the value is the number of base currency.   If it is &#x60;SPOT&#x60;/&#x60;MARGIN&#x60;, the value is the quantity in quote currency. | [optional] [default to ""]

## Methods

### NewGetMarketTickersV5RespDataInner

`func NewGetMarketTickersV5RespDataInner() *GetMarketTickersV5RespDataInner`

NewGetMarketTickersV5RespDataInner instantiates a new GetMarketTickersV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketTickersV5RespDataInnerWithDefaults

`func NewGetMarketTickersV5RespDataInnerWithDefaults() *GetMarketTickersV5RespDataInner`

NewGetMarketTickersV5RespDataInnerWithDefaults instantiates a new GetMarketTickersV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPx

`func (o *GetMarketTickersV5RespDataInner) GetAskPx() string`

GetAskPx returns the AskPx field if non-nil, zero value otherwise.

### GetAskPxOk

`func (o *GetMarketTickersV5RespDataInner) GetAskPxOk() (*string, bool)`

GetAskPxOk returns a tuple with the AskPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPx

`func (o *GetMarketTickersV5RespDataInner) SetAskPx(v string)`

SetAskPx sets AskPx field to given value.

### HasAskPx

`func (o *GetMarketTickersV5RespDataInner) HasAskPx() bool`

HasAskPx returns a boolean if a field has been set.

### GetAskSz

`func (o *GetMarketTickersV5RespDataInner) GetAskSz() string`

GetAskSz returns the AskSz field if non-nil, zero value otherwise.

### GetAskSzOk

`func (o *GetMarketTickersV5RespDataInner) GetAskSzOk() (*string, bool)`

GetAskSzOk returns a tuple with the AskSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSz

`func (o *GetMarketTickersV5RespDataInner) SetAskSz(v string)`

SetAskSz sets AskSz field to given value.

### HasAskSz

`func (o *GetMarketTickersV5RespDataInner) HasAskSz() bool`

HasAskSz returns a boolean if a field has been set.

### GetBidPx

`func (o *GetMarketTickersV5RespDataInner) GetBidPx() string`

GetBidPx returns the BidPx field if non-nil, zero value otherwise.

### GetBidPxOk

`func (o *GetMarketTickersV5RespDataInner) GetBidPxOk() (*string, bool)`

GetBidPxOk returns a tuple with the BidPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPx

`func (o *GetMarketTickersV5RespDataInner) SetBidPx(v string)`

SetBidPx sets BidPx field to given value.

### HasBidPx

`func (o *GetMarketTickersV5RespDataInner) HasBidPx() bool`

HasBidPx returns a boolean if a field has been set.

### GetBidSz

`func (o *GetMarketTickersV5RespDataInner) GetBidSz() string`

GetBidSz returns the BidSz field if non-nil, zero value otherwise.

### GetBidSzOk

`func (o *GetMarketTickersV5RespDataInner) GetBidSzOk() (*string, bool)`

GetBidSzOk returns a tuple with the BidSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSz

`func (o *GetMarketTickersV5RespDataInner) SetBidSz(v string)`

SetBidSz sets BidSz field to given value.

### HasBidSz

`func (o *GetMarketTickersV5RespDataInner) HasBidSz() bool`

HasBidSz returns a boolean if a field has been set.

### GetHigh24h

`func (o *GetMarketTickersV5RespDataInner) GetHigh24h() string`

GetHigh24h returns the High24h field if non-nil, zero value otherwise.

### GetHigh24hOk

`func (o *GetMarketTickersV5RespDataInner) GetHigh24hOk() (*string, bool)`

GetHigh24hOk returns a tuple with the High24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh24h

`func (o *GetMarketTickersV5RespDataInner) SetHigh24h(v string)`

SetHigh24h sets High24h field to given value.

### HasHigh24h

`func (o *GetMarketTickersV5RespDataInner) HasHigh24h() bool`

HasHigh24h returns a boolean if a field has been set.

### GetInstId

`func (o *GetMarketTickersV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetMarketTickersV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetMarketTickersV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetMarketTickersV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetMarketTickersV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetMarketTickersV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetMarketTickersV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetMarketTickersV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLast

`func (o *GetMarketTickersV5RespDataInner) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetMarketTickersV5RespDataInner) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetMarketTickersV5RespDataInner) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GetMarketTickersV5RespDataInner) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLastSz

`func (o *GetMarketTickersV5RespDataInner) GetLastSz() string`

GetLastSz returns the LastSz field if non-nil, zero value otherwise.

### GetLastSzOk

`func (o *GetMarketTickersV5RespDataInner) GetLastSzOk() (*string, bool)`

GetLastSzOk returns a tuple with the LastSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSz

`func (o *GetMarketTickersV5RespDataInner) SetLastSz(v string)`

SetLastSz sets LastSz field to given value.

### HasLastSz

`func (o *GetMarketTickersV5RespDataInner) HasLastSz() bool`

HasLastSz returns a boolean if a field has been set.

### GetLow24h

`func (o *GetMarketTickersV5RespDataInner) GetLow24h() string`

GetLow24h returns the Low24h field if non-nil, zero value otherwise.

### GetLow24hOk

`func (o *GetMarketTickersV5RespDataInner) GetLow24hOk() (*string, bool)`

GetLow24hOk returns a tuple with the Low24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow24h

`func (o *GetMarketTickersV5RespDataInner) SetLow24h(v string)`

SetLow24h sets Low24h field to given value.

### HasLow24h

`func (o *GetMarketTickersV5RespDataInner) HasLow24h() bool`

HasLow24h returns a boolean if a field has been set.

### GetOpen24h

`func (o *GetMarketTickersV5RespDataInner) GetOpen24h() string`

GetOpen24h returns the Open24h field if non-nil, zero value otherwise.

### GetOpen24hOk

`func (o *GetMarketTickersV5RespDataInner) GetOpen24hOk() (*string, bool)`

GetOpen24hOk returns a tuple with the Open24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen24h

`func (o *GetMarketTickersV5RespDataInner) SetOpen24h(v string)`

SetOpen24h sets Open24h field to given value.

### HasOpen24h

`func (o *GetMarketTickersV5RespDataInner) HasOpen24h() bool`

HasOpen24h returns a boolean if a field has been set.

### GetSodUtc0

`func (o *GetMarketTickersV5RespDataInner) GetSodUtc0() string`

GetSodUtc0 returns the SodUtc0 field if non-nil, zero value otherwise.

### GetSodUtc0Ok

`func (o *GetMarketTickersV5RespDataInner) GetSodUtc0Ok() (*string, bool)`

GetSodUtc0Ok returns a tuple with the SodUtc0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSodUtc0

`func (o *GetMarketTickersV5RespDataInner) SetSodUtc0(v string)`

SetSodUtc0 sets SodUtc0 field to given value.

### HasSodUtc0

`func (o *GetMarketTickersV5RespDataInner) HasSodUtc0() bool`

HasSodUtc0 returns a boolean if a field has been set.

### GetSodUtc8

`func (o *GetMarketTickersV5RespDataInner) GetSodUtc8() string`

GetSodUtc8 returns the SodUtc8 field if non-nil, zero value otherwise.

### GetSodUtc8Ok

`func (o *GetMarketTickersV5RespDataInner) GetSodUtc8Ok() (*string, bool)`

GetSodUtc8Ok returns a tuple with the SodUtc8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSodUtc8

`func (o *GetMarketTickersV5RespDataInner) SetSodUtc8(v string)`

SetSodUtc8 sets SodUtc8 field to given value.

### HasSodUtc8

`func (o *GetMarketTickersV5RespDataInner) HasSodUtc8() bool`

HasSodUtc8 returns a boolean if a field has been set.

### GetTs

`func (o *GetMarketTickersV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetMarketTickersV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetMarketTickersV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetMarketTickersV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetVol24h

`func (o *GetMarketTickersV5RespDataInner) GetVol24h() string`

GetVol24h returns the Vol24h field if non-nil, zero value otherwise.

### GetVol24hOk

`func (o *GetMarketTickersV5RespDataInner) GetVol24hOk() (*string, bool)`

GetVol24hOk returns a tuple with the Vol24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol24h

`func (o *GetMarketTickersV5RespDataInner) SetVol24h(v string)`

SetVol24h sets Vol24h field to given value.

### HasVol24h

`func (o *GetMarketTickersV5RespDataInner) HasVol24h() bool`

HasVol24h returns a boolean if a field has been set.

### GetVolCcy24h

`func (o *GetMarketTickersV5RespDataInner) GetVolCcy24h() string`

GetVolCcy24h returns the VolCcy24h field if non-nil, zero value otherwise.

### GetVolCcy24hOk

`func (o *GetMarketTickersV5RespDataInner) GetVolCcy24hOk() (*string, bool)`

GetVolCcy24hOk returns a tuple with the VolCcy24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolCcy24h

`func (o *GetMarketTickersV5RespDataInner) SetVolCcy24h(v string)`

SetVolCcy24h sets VolCcy24h field to given value.

### HasVolCcy24h

`func (o *GetMarketTickersV5RespDataInner) HasVolCcy24h() bool`

HasVolCcy24h returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


