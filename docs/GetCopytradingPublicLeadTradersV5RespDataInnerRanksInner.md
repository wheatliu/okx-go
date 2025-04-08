# GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccCopyTraderNum** | Pointer to **string** | Accumulated number of copy traders | [optional] [default to ""]
**Aum** | Pointer to **string** | assets under management | [optional] [default to ""]
**BeginTs** | Pointer to **string** | Begin time of pnl ratio on that day | [optional] [default to ""]
**Ccy** | Pointer to **string** | Margin currency | [optional] [default to ""]
**CopyState** | Pointer to **string** | Current copy state   &#x60;0&#x60;: non-copy, &#x60;1&#x60;: copy | [optional] [default to ""]
**CopyTraderNum** | Pointer to **string** | Current number of copy traders | [optional] [default to ""]
**LeadDays** | Pointer to **string** | Lead days | [optional] [default to ""]
**MaxCopyTraderNum** | Pointer to **string** | Maximum number of copy traders | [optional] [default to ""]
**NickName** | Pointer to **string** | Nick name | [optional] [default to ""]
**Pnl** | Pointer to **string** | Pnl (in USDT) of last 90 days. | [optional] [default to ""]
**PnlRatio** | Pointer to **string** | Pnl ratio on that day | [optional] [default to ""]
**PnlRatios** | Pointer to [**[]GetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerPnlRatiosInner**](GetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerPnlRatiosInner.md) | Pnl ratios | [optional] 
**PortLink** | Pointer to **string** | Portrait link | [optional] [default to ""]
**TraderInsts** | Pointer to **[]string** | Contract list which lead trader is leading | [optional] 
**UniqueCode** | Pointer to **string** | Lead trader unique code | [optional] [default to ""]
**WinRatio** | Pointer to **string** | Win ratio, 0.1 represents 10% | [optional] [default to ""]

## Methods

### NewGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner

`func NewGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner() *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner`

NewGetCopytradingPublicLeadTradersV5RespDataInnerRanksInner instantiates a new GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerWithDefaults

`func NewGetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerWithDefaults() *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner`

NewGetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerWithDefaults instantiates a new GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccCopyTraderNum

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetAccCopyTraderNum() string`

GetAccCopyTraderNum returns the AccCopyTraderNum field if non-nil, zero value otherwise.

### GetAccCopyTraderNumOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetAccCopyTraderNumOk() (*string, bool)`

GetAccCopyTraderNumOk returns a tuple with the AccCopyTraderNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccCopyTraderNum

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetAccCopyTraderNum(v string)`

SetAccCopyTraderNum sets AccCopyTraderNum field to given value.

### HasAccCopyTraderNum

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasAccCopyTraderNum() bool`

HasAccCopyTraderNum returns a boolean if a field has been set.

### GetAum

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetAum() string`

GetAum returns the Aum field if non-nil, zero value otherwise.

### GetAumOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetAumOk() (*string, bool)`

GetAumOk returns a tuple with the Aum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAum

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetAum(v string)`

SetAum sets Aum field to given value.

### HasAum

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasAum() bool`

HasAum returns a boolean if a field has been set.

### GetBeginTs

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetBeginTs() string`

GetBeginTs returns the BeginTs field if non-nil, zero value otherwise.

### GetBeginTsOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetBeginTsOk() (*string, bool)`

GetBeginTsOk returns a tuple with the BeginTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginTs

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetBeginTs(v string)`

SetBeginTs sets BeginTs field to given value.

### HasBeginTs

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasBeginTs() bool`

HasBeginTs returns a boolean if a field has been set.

### GetCcy

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetCopyState

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetCopyState() string`

GetCopyState returns the CopyState field if non-nil, zero value otherwise.

### GetCopyStateOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetCopyStateOk() (*string, bool)`

GetCopyStateOk returns a tuple with the CopyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyState

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetCopyState(v string)`

SetCopyState sets CopyState field to given value.

### HasCopyState

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasCopyState() bool`

HasCopyState returns a boolean if a field has been set.

### GetCopyTraderNum

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetCopyTraderNum() string`

GetCopyTraderNum returns the CopyTraderNum field if non-nil, zero value otherwise.

### GetCopyTraderNumOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetCopyTraderNumOk() (*string, bool)`

GetCopyTraderNumOk returns a tuple with the CopyTraderNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTraderNum

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetCopyTraderNum(v string)`

SetCopyTraderNum sets CopyTraderNum field to given value.

### HasCopyTraderNum

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasCopyTraderNum() bool`

HasCopyTraderNum returns a boolean if a field has been set.

### GetLeadDays

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetLeadDays() string`

GetLeadDays returns the LeadDays field if non-nil, zero value otherwise.

### GetLeadDaysOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetLeadDaysOk() (*string, bool)`

GetLeadDaysOk returns a tuple with the LeadDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadDays

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetLeadDays(v string)`

SetLeadDays sets LeadDays field to given value.

### HasLeadDays

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasLeadDays() bool`

HasLeadDays returns a boolean if a field has been set.

### GetMaxCopyTraderNum

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetMaxCopyTraderNum() string`

GetMaxCopyTraderNum returns the MaxCopyTraderNum field if non-nil, zero value otherwise.

### GetMaxCopyTraderNumOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetMaxCopyTraderNumOk() (*string, bool)`

GetMaxCopyTraderNumOk returns a tuple with the MaxCopyTraderNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCopyTraderNum

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetMaxCopyTraderNum(v string)`

SetMaxCopyTraderNum sets MaxCopyTraderNum field to given value.

### HasMaxCopyTraderNum

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasMaxCopyTraderNum() bool`

HasMaxCopyTraderNum returns a boolean if a field has been set.

### GetNickName

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### GetPnl

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetPnlRatio

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPnlRatio() string`

GetPnlRatio returns the PnlRatio field if non-nil, zero value otherwise.

### GetPnlRatioOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPnlRatioOk() (*string, bool)`

GetPnlRatioOk returns a tuple with the PnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlRatio

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetPnlRatio(v string)`

SetPnlRatio sets PnlRatio field to given value.

### HasPnlRatio

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasPnlRatio() bool`

HasPnlRatio returns a boolean if a field has been set.

### GetPnlRatios

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPnlRatios() []GetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerPnlRatiosInner`

GetPnlRatios returns the PnlRatios field if non-nil, zero value otherwise.

### GetPnlRatiosOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPnlRatiosOk() (*[]GetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerPnlRatiosInner, bool)`

GetPnlRatiosOk returns a tuple with the PnlRatios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlRatios

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetPnlRatios(v []GetCopytradingPublicLeadTradersV5RespDataInnerRanksInnerPnlRatiosInner)`

SetPnlRatios sets PnlRatios field to given value.

### HasPnlRatios

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasPnlRatios() bool`

HasPnlRatios returns a boolean if a field has been set.

### GetPortLink

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPortLink() string`

GetPortLink returns the PortLink field if non-nil, zero value otherwise.

### GetPortLinkOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetPortLinkOk() (*string, bool)`

GetPortLinkOk returns a tuple with the PortLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLink

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetPortLink(v string)`

SetPortLink sets PortLink field to given value.

### HasPortLink

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasPortLink() bool`

HasPortLink returns a boolean if a field has been set.

### GetTraderInsts

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetTraderInsts() []string`

GetTraderInsts returns the TraderInsts field if non-nil, zero value otherwise.

### GetTraderInstsOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetTraderInstsOk() (*[]string, bool)`

GetTraderInstsOk returns a tuple with the TraderInsts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraderInsts

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetTraderInsts(v []string)`

SetTraderInsts sets TraderInsts field to given value.

### HasTraderInsts

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasTraderInsts() bool`

HasTraderInsts returns a boolean if a field has been set.

### GetUniqueCode

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetUniqueCode() string`

GetUniqueCode returns the UniqueCode field if non-nil, zero value otherwise.

### GetUniqueCodeOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetUniqueCodeOk() (*string, bool)`

GetUniqueCodeOk returns a tuple with the UniqueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCode

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetUniqueCode(v string)`

SetUniqueCode sets UniqueCode field to given value.

### HasUniqueCode

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasUniqueCode() bool`

HasUniqueCode returns a boolean if a field has been set.

### GetWinRatio

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetWinRatio() string`

GetWinRatio returns the WinRatio field if non-nil, zero value otherwise.

### GetWinRatioOk

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) GetWinRatioOk() (*string, bool)`

GetWinRatioOk returns a tuple with the WinRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinRatio

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) SetWinRatio(v string)`

SetWinRatio sets WinRatio field to given value.

### HasWinRatio

`func (o *GetCopytradingPublicLeadTradersV5RespDataInnerRanksInner) HasWinRatio() bool`

HasWinRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


