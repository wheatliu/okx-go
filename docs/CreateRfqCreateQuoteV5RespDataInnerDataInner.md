# CreateRfqCreateQuoteV5RespDataInnerDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CTime** | Pointer to **string** | The timestamp the Quote was created, Unix timestamp format in milliseconds. | [optional] [default to ""]
**Ccy** | Pointer to **string** | Margin currency.   Only applicable to &#x60;cross&#x60; &#x60;MARGIN&#x60; orders in &#x60;Spot and futures mode&#x60;. The parameter will be ignored in other scenarios. | [optional] [default to ""]
**ClQuoteId** | Pointer to **string** | Client-supplied Quote ID.   This attribute is treated as client sensitive information. It will not be exposed to the Taker, only return empty string. | [optional] [default to ""]
**ClRfqId** | Pointer to **string** | Client-supplied RFQ ID.   This attribute is treated as client sensitive information. It will not be exposed to the Maker, only return empty string. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [optional] [default to ""]
**Legs** | Pointer to [**[]CreateRfqCreateQuoteV5RespDataInnerDataInnerLegsInner**](CreateRfqCreateQuoteV5RespDataInnerDataInnerLegsInner.md) | The legs of the Quote. | [optional] 
**PosSide** | Pointer to **string** | Position side.   The default is &#x60;net&#x60; in the net mode. If not specified, return \&quot;\&quot;, which is equivalent to net.   It can only be &#x60;long&#x60; or &#x60;short&#x60; in the long/short mode. If not specified, return \&quot;\&quot;, which corresponds to the direction that opens new positions for the trade (buy &#x3D;&gt; long, sell &#x3D;&gt; short).   Only applicable to FUTURES/SWAP. | [optional] [default to ""]
**Px** | Pointer to **string** | The price of the leg. | [optional] [default to ""]
**QuoteId** | Pointer to **string** | Quote ID. | [optional] [default to ""]
**QuoteSide** | Pointer to **string** | The trading direction of the Quote.   Its value can be &#x60;buy&#x60; or &#x60;sell&#x60;. For example, if quoteSide is &#x60;buy&#x60;, all the legs are executed in their leg sides; otherwise, all the legs are executed in the opposite of their leg sides. | [optional] [default to ""]
**Reason** | Pointer to **string** | Reasons of state. Valid values can be &#x60;mmp_canceled&#x60;. | [optional] [default to ""]
**RfqId** | Pointer to **string** | RFQ ID | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of the leg. Valid values can be buy or sell. | [optional] [default to ""]
**State** | Pointer to **string** | The status of the quote. Valid values can be &#x60;active&#x60; &#x60;canceled&#x60; &#x60;pending_fill&#x60; &#x60;filled&#x60; &#x60;expired&#x60; or &#x60;failed&#x60;. | [optional] [default to ""]
**Sz** | Pointer to **string** | Size of the leg in contracts or spot. | [optional] [default to ""]
**Tag** | Pointer to **string** | Quote tag.   The block trade associated with the Quote will have the same tag. | [optional] [default to ""]
**TdMode** | Pointer to **string** | Trade mode   Margin mode: &#x60;cross&#x60; &#x60;isolated&#x60;   Non-Margin mode: &#x60;cash&#x60;.   If not provided, tdMode will inherit default values set by the system shown below:   Spot and futures mode &amp; SPOT: &#x60;cash&#x60;   Buy options in Spot and futures mode and Multi-currency Margin: &#x60;isolated&#x60;   Other cases: &#x60;cross&#x60; | [optional] [default to ""]
**TgtCcy** | Pointer to **string** | Defines the unit of the “sz” attribute.   Only applicable to instType &#x3D; SPOT.   The valid enumerations are &#x60;base_ccy&#x60; and &#x60;quote_ccy&#x60;. When not specified this is equal to &#x60;base_ccy&#x60; by default. | [optional] [default to ""]
**TraderCode** | Pointer to **string** | A unique identifier of maker. | [optional] [default to ""]
**UTime** | Pointer to **string** | The timestamp the Quote was last updated, Unix timestamp format in milliseconds. | [optional] [default to ""]
**ValidUntil** | Pointer to **string** | The timestamp the Quote expires. Unix timestamp format in milliseconds. | [optional] [default to ""]

## Methods

### NewCreateRfqCreateQuoteV5RespDataInnerDataInner

`func NewCreateRfqCreateQuoteV5RespDataInnerDataInner() *CreateRfqCreateQuoteV5RespDataInnerDataInner`

NewCreateRfqCreateQuoteV5RespDataInnerDataInner instantiates a new CreateRfqCreateQuoteV5RespDataInnerDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRfqCreateQuoteV5RespDataInnerDataInnerWithDefaults

`func NewCreateRfqCreateQuoteV5RespDataInnerDataInnerWithDefaults() *CreateRfqCreateQuoteV5RespDataInnerDataInner`

NewCreateRfqCreateQuoteV5RespDataInnerDataInnerWithDefaults instantiates a new CreateRfqCreateQuoteV5RespDataInnerDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCTime

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCcy

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClQuoteId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetClQuoteId() string`

GetClQuoteId returns the ClQuoteId field if non-nil, zero value otherwise.

### GetClQuoteIdOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetClQuoteIdOk() (*string, bool)`

GetClQuoteIdOk returns a tuple with the ClQuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClQuoteId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetClQuoteId(v string)`

SetClQuoteId sets ClQuoteId field to given value.

### HasClQuoteId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasClQuoteId() bool`

HasClQuoteId returns a boolean if a field has been set.

### GetClRfqId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetClRfqId() string`

GetClRfqId returns the ClRfqId field if non-nil, zero value otherwise.

### GetClRfqIdOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetClRfqIdOk() (*string, bool)`

GetClRfqIdOk returns a tuple with the ClRfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClRfqId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetClRfqId(v string)`

SetClRfqId sets ClRfqId field to given value.

### HasClRfqId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasClRfqId() bool`

HasClRfqId returns a boolean if a field has been set.

### GetInstId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLegs

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetLegs() []CreateRfqCreateQuoteV5RespDataInnerDataInnerLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetLegsOk() (*[]CreateRfqCreateQuoteV5RespDataInnerDataInnerLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetLegs(v []CreateRfqCreateQuoteV5RespDataInnerDataInnerLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetPosSide

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetPx

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetQuoteId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetQuoteSide

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetQuoteSide() string`

GetQuoteSide returns the QuoteSide field if non-nil, zero value otherwise.

### GetQuoteSideOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetQuoteSideOk() (*string, bool)`

GetQuoteSideOk returns a tuple with the QuoteSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteSide

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetQuoteSide(v string)`

SetQuoteSide sets QuoteSide field to given value.

### HasQuoteSide

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasQuoteSide() bool`

HasQuoteSide returns a boolean if a field has been set.

### GetReason

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRfqId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetRfqId() string`

GetRfqId returns the RfqId field if non-nil, zero value otherwise.

### GetRfqIdOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetRfqIdOk() (*string, bool)`

GetRfqIdOk returns a tuple with the RfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetRfqId(v string)`

SetRfqId sets RfqId field to given value.

### HasRfqId

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasRfqId() bool`

HasRfqId returns a boolean if a field has been set.

### GetSide

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetState

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSz

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTdMode

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.

### HasTdMode

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasTdMode() bool`

HasTdMode returns a boolean if a field has been set.

### GetTgtCcy

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetTgtCcy() string`

GetTgtCcy returns the TgtCcy field if non-nil, zero value otherwise.

### GetTgtCcyOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetTgtCcyOk() (*string, bool)`

GetTgtCcyOk returns a tuple with the TgtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCcy

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetTgtCcy(v string)`

SetTgtCcy sets TgtCcy field to given value.

### HasTgtCcy

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasTgtCcy() bool`

HasTgtCcy returns a boolean if a field has been set.

### GetTraderCode

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetTraderCode() string`

GetTraderCode returns the TraderCode field if non-nil, zero value otherwise.

### GetTraderCodeOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetTraderCodeOk() (*string, bool)`

GetTraderCodeOk returns a tuple with the TraderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraderCode

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetTraderCode(v string)`

SetTraderCode sets TraderCode field to given value.

### HasTraderCode

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasTraderCode() bool`

HasTraderCode returns a boolean if a field has been set.

### GetUTime

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetValidUntil

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetValidUntil() string`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) GetValidUntilOk() (*string, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) SetValidUntil(v string)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *CreateRfqCreateQuoteV5RespDataInnerDataInner) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


