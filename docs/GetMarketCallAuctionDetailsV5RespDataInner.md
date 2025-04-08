# GetMarketCallAuctionDetailsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionEndTime** | Pointer to **string** | Call auction end time. Unix timestamp in milliseconds. | [optional] [default to ""]
**EqPx** | Pointer to **string** | Equilibrium price | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**MatchedSz** | Pointer to **string** | Matched size for both buy and sell  The unit is in base currency | [optional] [default to ""]
**State** | Pointer to **string** | Trading state of the symbol  &#x60;call_auction&#x60;  &#x60;continuous_trading&#x60; | [optional] [default to ""]
**Ts** | Pointer to **string** | Data generation time. Unix timestamp in millieseconds. | [optional] [default to ""]
**UnmatchedSz** | Pointer to **string** | Unmatched size | [optional] [default to ""]

## Methods

### NewGetMarketCallAuctionDetailsV5RespDataInner

`func NewGetMarketCallAuctionDetailsV5RespDataInner() *GetMarketCallAuctionDetailsV5RespDataInner`

NewGetMarketCallAuctionDetailsV5RespDataInner instantiates a new GetMarketCallAuctionDetailsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketCallAuctionDetailsV5RespDataInnerWithDefaults

`func NewGetMarketCallAuctionDetailsV5RespDataInnerWithDefaults() *GetMarketCallAuctionDetailsV5RespDataInner`

NewGetMarketCallAuctionDetailsV5RespDataInnerWithDefaults instantiates a new GetMarketCallAuctionDetailsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionEndTime

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetAuctionEndTime() string`

GetAuctionEndTime returns the AuctionEndTime field if non-nil, zero value otherwise.

### GetAuctionEndTimeOk

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetAuctionEndTimeOk() (*string, bool)`

GetAuctionEndTimeOk returns a tuple with the AuctionEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionEndTime

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetAuctionEndTime(v string)`

SetAuctionEndTime sets AuctionEndTime field to given value.

### HasAuctionEndTime

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasAuctionEndTime() bool`

HasAuctionEndTime returns a boolean if a field has been set.

### GetEqPx

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetEqPx() string`

GetEqPx returns the EqPx field if non-nil, zero value otherwise.

### GetEqPxOk

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetEqPxOk() (*string, bool)`

GetEqPxOk returns a tuple with the EqPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqPx

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetEqPx(v string)`

SetEqPx sets EqPx field to given value.

### HasEqPx

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasEqPx() bool`

HasEqPx returns a boolean if a field has been set.

### GetInstId

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetMatchedSz

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetMatchedSz() string`

GetMatchedSz returns the MatchedSz field if non-nil, zero value otherwise.

### GetMatchedSzOk

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetMatchedSzOk() (*string, bool)`

GetMatchedSzOk returns a tuple with the MatchedSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedSz

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetMatchedSz(v string)`

SetMatchedSz sets MatchedSz field to given value.

### HasMatchedSz

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasMatchedSz() bool`

HasMatchedSz returns a boolean if a field has been set.

### GetState

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTs

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetUnmatchedSz

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetUnmatchedSz() string`

GetUnmatchedSz returns the UnmatchedSz field if non-nil, zero value otherwise.

### GetUnmatchedSzOk

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) GetUnmatchedSzOk() (*string, bool)`

GetUnmatchedSzOk returns a tuple with the UnmatchedSz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedSz

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) SetUnmatchedSz(v string)`

SetUnmatchedSz sets UnmatchedSz field to given value.

### HasUnmatchedSz

`func (o *GetMarketCallAuctionDetailsV5RespDataInner) HasUnmatchedSz() bool`

HasUnmatchedSz returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


