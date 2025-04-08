# GetRfqRfqsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPartialExecution** | Pointer to **bool** | Whether the RFQ can be partially filled provided that the shape of legs stays the same.    Valid value is &#x60;true&#x60; or &#x60;false&#x60;. &#x60;false&#x60; by default. | [optional] 
**CTime** | Pointer to **string** | The timestamp the RFQ was created. Unix timestamp format in milliseconds. | [optional] [default to ""]
**Ccy** | Pointer to **string** | Margin currency.   Only applicable to &#x60;cross&#x60; &#x60;MARGIN&#x60; orders in &#x60;Spot and futures mode&#x60;. The parameter will be ignored in other scenarios. | [optional] [default to ""]
**ClRfqId** | Pointer to **string** | Client-supplied RFQ ID.   This attribute is treated as client sensitive information. It will not be exposed to the Maker, only return empty string. | [optional] [default to ""]
**Counterparties** | Pointer to **[]string** | The list of counterparties traderCode the RFQ was broadcasted to. | [optional] 
**FlowType** | Pointer to **string** | Identify the type of the RFQ.   Only applicable to Makers, return \&quot;\&quot; for Takers | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. BTC-USDT-SWAP | [optional] [default to ""]
**Legs** | Pointer to [**[]GetRfqRfqsV5RespDataInnerLegsInner**](GetRfqRfqsV5RespDataInnerLegsInner.md) | Legs of RFQ | [optional] 
**PosSide** | Pointer to **string** | Position side.   The default is &#x60;net&#x60; in the net mode. If not specified, return \&quot;\&quot;, which is equivalent to net.   It can only be &#x60;long&#x60; or &#x60;short&#x60; in the long/short mode. If not specified, return \&quot;\&quot;, which corresponds to the direction that opens new positions for the trade (buy &#x3D;&gt; long, sell &#x3D;&gt; short).   Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;. | [optional] [default to ""]
**RfqId** | Pointer to **string** | RFQ ID. | [optional] [default to ""]
**Side** | Pointer to **string** | The direction of the leg. Valid values can be buy or sell. | [optional] [default to ""]
**State** | Pointer to **string** | The status of the RFQ.   Valid values can be &#x60;active&#x60; &#x60;canceled&#x60; &#x60;pending_fill&#x60; &#x60;filled&#x60; &#x60;expired&#x60; &#x60;failed&#x60; &#x60;traded_away&#x60;.   &#x60;traded_away&#x60; only applies to Maker | [optional] [default to ""]
**Sz** | Pointer to **string** | Size of the leg in contracts or spot. | [optional] [default to ""]
**Tag** | Pointer to **string** | RFQ tag.   The block trade associated with the RFQ will have the same tag. | [optional] [default to ""]
**TdMode** | Pointer to **string** | Trade mode   Margin mode: &#x60;cross&#x60; &#x60;isolated&#x60;   Non-Margin mode: &#x60;cash&#x60;.   If not provided, tdMode will inherit default values set by the system shown below:   Spot and futures mode &amp; SPOT: &#x60;cash&#x60;   Buy options in Spot and futures mode and Multi-currency Margin: &#x60;isolated&#x60;   Other cases: &#x60;cross&#x60; | [optional] [default to ""]
**TgtCcy** | Pointer to **string** | Defines the unit of the “sz” attribute.   Only applicable to instType &#x3D; SPOT.   The valid enumerations are base_ccy and quote_ccy. When not specified this is equal to base_ccy by default. | [optional] [default to ""]
**TraderCode** | Pointer to **string** | A unique identifier of taker. Empty if the anonymous parameter of the RFQ is set to be &#x60;true&#x60;. | [optional] [default to ""]
**UTime** | Pointer to **string** | The timestamp the RFQ was last updated. Unix timestamp format in milliseconds. | [optional] [default to ""]
**ValidUntil** | Pointer to **string** | The timestamp the RFQ expires. Unix timestamp format in milliseconds. | [optional] [default to ""]

## Methods

### NewGetRfqRfqsV5RespDataInner

`func NewGetRfqRfqsV5RespDataInner() *GetRfqRfqsV5RespDataInner`

NewGetRfqRfqsV5RespDataInner instantiates a new GetRfqRfqsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRfqRfqsV5RespDataInnerWithDefaults

`func NewGetRfqRfqsV5RespDataInnerWithDefaults() *GetRfqRfqsV5RespDataInner`

NewGetRfqRfqsV5RespDataInnerWithDefaults instantiates a new GetRfqRfqsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPartialExecution

`func (o *GetRfqRfqsV5RespDataInner) GetAllowPartialExecution() bool`

GetAllowPartialExecution returns the AllowPartialExecution field if non-nil, zero value otherwise.

### GetAllowPartialExecutionOk

`func (o *GetRfqRfqsV5RespDataInner) GetAllowPartialExecutionOk() (*bool, bool)`

GetAllowPartialExecutionOk returns a tuple with the AllowPartialExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPartialExecution

`func (o *GetRfqRfqsV5RespDataInner) SetAllowPartialExecution(v bool)`

SetAllowPartialExecution sets AllowPartialExecution field to given value.

### HasAllowPartialExecution

`func (o *GetRfqRfqsV5RespDataInner) HasAllowPartialExecution() bool`

HasAllowPartialExecution returns a boolean if a field has been set.

### GetCTime

`func (o *GetRfqRfqsV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetRfqRfqsV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetRfqRfqsV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetRfqRfqsV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCcy

`func (o *GetRfqRfqsV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetRfqRfqsV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetRfqRfqsV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetRfqRfqsV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClRfqId

`func (o *GetRfqRfqsV5RespDataInner) GetClRfqId() string`

GetClRfqId returns the ClRfqId field if non-nil, zero value otherwise.

### GetClRfqIdOk

`func (o *GetRfqRfqsV5RespDataInner) GetClRfqIdOk() (*string, bool)`

GetClRfqIdOk returns a tuple with the ClRfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClRfqId

`func (o *GetRfqRfqsV5RespDataInner) SetClRfqId(v string)`

SetClRfqId sets ClRfqId field to given value.

### HasClRfqId

`func (o *GetRfqRfqsV5RespDataInner) HasClRfqId() bool`

HasClRfqId returns a boolean if a field has been set.

### GetCounterparties

`func (o *GetRfqRfqsV5RespDataInner) GetCounterparties() []string`

GetCounterparties returns the Counterparties field if non-nil, zero value otherwise.

### GetCounterpartiesOk

`func (o *GetRfqRfqsV5RespDataInner) GetCounterpartiesOk() (*[]string, bool)`

GetCounterpartiesOk returns a tuple with the Counterparties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterparties

`func (o *GetRfqRfqsV5RespDataInner) SetCounterparties(v []string)`

SetCounterparties sets Counterparties field to given value.

### HasCounterparties

`func (o *GetRfqRfqsV5RespDataInner) HasCounterparties() bool`

HasCounterparties returns a boolean if a field has been set.

### GetFlowType

`func (o *GetRfqRfqsV5RespDataInner) GetFlowType() string`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *GetRfqRfqsV5RespDataInner) GetFlowTypeOk() (*string, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *GetRfqRfqsV5RespDataInner) SetFlowType(v string)`

SetFlowType sets FlowType field to given value.

### HasFlowType

`func (o *GetRfqRfqsV5RespDataInner) HasFlowType() bool`

HasFlowType returns a boolean if a field has been set.

### GetInstId

`func (o *GetRfqRfqsV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetRfqRfqsV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetRfqRfqsV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetRfqRfqsV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLegs

`func (o *GetRfqRfqsV5RespDataInner) GetLegs() []GetRfqRfqsV5RespDataInnerLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *GetRfqRfqsV5RespDataInner) GetLegsOk() (*[]GetRfqRfqsV5RespDataInnerLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *GetRfqRfqsV5RespDataInner) SetLegs(v []GetRfqRfqsV5RespDataInnerLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *GetRfqRfqsV5RespDataInner) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetPosSide

`func (o *GetRfqRfqsV5RespDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetRfqRfqsV5RespDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetRfqRfqsV5RespDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetRfqRfqsV5RespDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetRfqId

`func (o *GetRfqRfqsV5RespDataInner) GetRfqId() string`

GetRfqId returns the RfqId field if non-nil, zero value otherwise.

### GetRfqIdOk

`func (o *GetRfqRfqsV5RespDataInner) GetRfqIdOk() (*string, bool)`

GetRfqIdOk returns a tuple with the RfqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfqId

`func (o *GetRfqRfqsV5RespDataInner) SetRfqId(v string)`

SetRfqId sets RfqId field to given value.

### HasRfqId

`func (o *GetRfqRfqsV5RespDataInner) HasRfqId() bool`

HasRfqId returns a boolean if a field has been set.

### GetSide

`func (o *GetRfqRfqsV5RespDataInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetRfqRfqsV5RespDataInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetRfqRfqsV5RespDataInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetRfqRfqsV5RespDataInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetState

`func (o *GetRfqRfqsV5RespDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetRfqRfqsV5RespDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetRfqRfqsV5RespDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetRfqRfqsV5RespDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSz

`func (o *GetRfqRfqsV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetRfqRfqsV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetRfqRfqsV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetRfqRfqsV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *GetRfqRfqsV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetRfqRfqsV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetRfqRfqsV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetRfqRfqsV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTdMode

`func (o *GetRfqRfqsV5RespDataInner) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *GetRfqRfqsV5RespDataInner) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *GetRfqRfqsV5RespDataInner) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.

### HasTdMode

`func (o *GetRfqRfqsV5RespDataInner) HasTdMode() bool`

HasTdMode returns a boolean if a field has been set.

### GetTgtCcy

`func (o *GetRfqRfqsV5RespDataInner) GetTgtCcy() string`

GetTgtCcy returns the TgtCcy field if non-nil, zero value otherwise.

### GetTgtCcyOk

`func (o *GetRfqRfqsV5RespDataInner) GetTgtCcyOk() (*string, bool)`

GetTgtCcyOk returns a tuple with the TgtCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtCcy

`func (o *GetRfqRfqsV5RespDataInner) SetTgtCcy(v string)`

SetTgtCcy sets TgtCcy field to given value.

### HasTgtCcy

`func (o *GetRfqRfqsV5RespDataInner) HasTgtCcy() bool`

HasTgtCcy returns a boolean if a field has been set.

### GetTraderCode

`func (o *GetRfqRfqsV5RespDataInner) GetTraderCode() string`

GetTraderCode returns the TraderCode field if non-nil, zero value otherwise.

### GetTraderCodeOk

`func (o *GetRfqRfqsV5RespDataInner) GetTraderCodeOk() (*string, bool)`

GetTraderCodeOk returns a tuple with the TraderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraderCode

`func (o *GetRfqRfqsV5RespDataInner) SetTraderCode(v string)`

SetTraderCode sets TraderCode field to given value.

### HasTraderCode

`func (o *GetRfqRfqsV5RespDataInner) HasTraderCode() bool`

HasTraderCode returns a boolean if a field has been set.

### GetUTime

`func (o *GetRfqRfqsV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetRfqRfqsV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetRfqRfqsV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetRfqRfqsV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetValidUntil

`func (o *GetRfqRfqsV5RespDataInner) GetValidUntil() string`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *GetRfqRfqsV5RespDataInner) GetValidUntilOk() (*string, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *GetRfqRfqsV5RespDataInner) SetValidUntil(v string)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *GetRfqRfqsV5RespDataInner) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


