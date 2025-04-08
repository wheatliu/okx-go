# GetCopytradingPublicStatsV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgSubPosNotional** | Pointer to **string** | Average lead position notional (USDT) | [optional] [default to ""]
**Ccy** | Pointer to **string** | Margin currency | [optional] [default to ""]
**CurCopyTraderPnl** | Pointer to **string** | Current copy trader pnl (USDT) | [optional] [default to ""]
**InvestAmt** | Pointer to **string** | Investment amount (USDT) | [optional] [default to ""]
**LossDays** | Pointer to **string** | Loss days | [optional] [default to ""]
**ProfitDays** | Pointer to **string** | Profit days | [optional] [default to ""]
**WinRatio** | Pointer to **string** | Win ratio | [optional] [default to ""]

## Methods

### NewGetCopytradingPublicStatsV5RespData

`func NewGetCopytradingPublicStatsV5RespData() *GetCopytradingPublicStatsV5RespData`

NewGetCopytradingPublicStatsV5RespData instantiates a new GetCopytradingPublicStatsV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingPublicStatsV5RespDataWithDefaults

`func NewGetCopytradingPublicStatsV5RespDataWithDefaults() *GetCopytradingPublicStatsV5RespData`

NewGetCopytradingPublicStatsV5RespDataWithDefaults instantiates a new GetCopytradingPublicStatsV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgSubPosNotional

`func (o *GetCopytradingPublicStatsV5RespData) GetAvgSubPosNotional() string`

GetAvgSubPosNotional returns the AvgSubPosNotional field if non-nil, zero value otherwise.

### GetAvgSubPosNotionalOk

`func (o *GetCopytradingPublicStatsV5RespData) GetAvgSubPosNotionalOk() (*string, bool)`

GetAvgSubPosNotionalOk returns a tuple with the AvgSubPosNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSubPosNotional

`func (o *GetCopytradingPublicStatsV5RespData) SetAvgSubPosNotional(v string)`

SetAvgSubPosNotional sets AvgSubPosNotional field to given value.

### HasAvgSubPosNotional

`func (o *GetCopytradingPublicStatsV5RespData) HasAvgSubPosNotional() bool`

HasAvgSubPosNotional returns a boolean if a field has been set.

### GetCcy

`func (o *GetCopytradingPublicStatsV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetCopytradingPublicStatsV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetCopytradingPublicStatsV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetCopytradingPublicStatsV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetCurCopyTraderPnl

`func (o *GetCopytradingPublicStatsV5RespData) GetCurCopyTraderPnl() string`

GetCurCopyTraderPnl returns the CurCopyTraderPnl field if non-nil, zero value otherwise.

### GetCurCopyTraderPnlOk

`func (o *GetCopytradingPublicStatsV5RespData) GetCurCopyTraderPnlOk() (*string, bool)`

GetCurCopyTraderPnlOk returns a tuple with the CurCopyTraderPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurCopyTraderPnl

`func (o *GetCopytradingPublicStatsV5RespData) SetCurCopyTraderPnl(v string)`

SetCurCopyTraderPnl sets CurCopyTraderPnl field to given value.

### HasCurCopyTraderPnl

`func (o *GetCopytradingPublicStatsV5RespData) HasCurCopyTraderPnl() bool`

HasCurCopyTraderPnl returns a boolean if a field has been set.

### GetInvestAmt

`func (o *GetCopytradingPublicStatsV5RespData) GetInvestAmt() string`

GetInvestAmt returns the InvestAmt field if non-nil, zero value otherwise.

### GetInvestAmtOk

`func (o *GetCopytradingPublicStatsV5RespData) GetInvestAmtOk() (*string, bool)`

GetInvestAmtOk returns a tuple with the InvestAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestAmt

`func (o *GetCopytradingPublicStatsV5RespData) SetInvestAmt(v string)`

SetInvestAmt sets InvestAmt field to given value.

### HasInvestAmt

`func (o *GetCopytradingPublicStatsV5RespData) HasInvestAmt() bool`

HasInvestAmt returns a boolean if a field has been set.

### GetLossDays

`func (o *GetCopytradingPublicStatsV5RespData) GetLossDays() string`

GetLossDays returns the LossDays field if non-nil, zero value otherwise.

### GetLossDaysOk

`func (o *GetCopytradingPublicStatsV5RespData) GetLossDaysOk() (*string, bool)`

GetLossDaysOk returns a tuple with the LossDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossDays

`func (o *GetCopytradingPublicStatsV5RespData) SetLossDays(v string)`

SetLossDays sets LossDays field to given value.

### HasLossDays

`func (o *GetCopytradingPublicStatsV5RespData) HasLossDays() bool`

HasLossDays returns a boolean if a field has been set.

### GetProfitDays

`func (o *GetCopytradingPublicStatsV5RespData) GetProfitDays() string`

GetProfitDays returns the ProfitDays field if non-nil, zero value otherwise.

### GetProfitDaysOk

`func (o *GetCopytradingPublicStatsV5RespData) GetProfitDaysOk() (*string, bool)`

GetProfitDaysOk returns a tuple with the ProfitDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitDays

`func (o *GetCopytradingPublicStatsV5RespData) SetProfitDays(v string)`

SetProfitDays sets ProfitDays field to given value.

### HasProfitDays

`func (o *GetCopytradingPublicStatsV5RespData) HasProfitDays() bool`

HasProfitDays returns a boolean if a field has been set.

### GetWinRatio

`func (o *GetCopytradingPublicStatsV5RespData) GetWinRatio() string`

GetWinRatio returns the WinRatio field if non-nil, zero value otherwise.

### GetWinRatioOk

`func (o *GetCopytradingPublicStatsV5RespData) GetWinRatioOk() (*string, bool)`

GetWinRatioOk returns a tuple with the WinRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinRatio

`func (o *GetCopytradingPublicStatsV5RespData) SetWinRatio(v string)`

SetWinRatio sets WinRatio field to given value.

### HasWinRatio

`func (o *GetCopytradingPublicStatsV5RespData) HasWinRatio() bool`

HasWinRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


