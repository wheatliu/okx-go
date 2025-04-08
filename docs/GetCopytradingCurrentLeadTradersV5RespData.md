# GetCopytradingCurrentLeadTradersV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeginCopyTime** | Pointer to **string** | Begin copying time. Unix timestamp format in milliseconds, e.g.1597026383085 | [optional] [default to ""]
**Ccy** | Pointer to **string** | margin currency | [optional] [default to ""]
**CopyTotalAmt** | Pointer to **string** | Copy total amount | [optional] [default to ""]
**CopyTotalPnl** | Pointer to **string** | Copy total pnl | [optional] [default to ""]
**LeadMode** | Pointer to **string** | Lead mode &#x60;public&#x60; &#x60;private&#x60; | [optional] [default to ""]
**Margin** | Pointer to **string** | Margin for copy trading | [optional] [default to ""]
**NickName** | Pointer to **string** | Nick name | [optional] [default to ""]
**PortLink** | Pointer to **string** | Portrait link | [optional] [default to ""]
**ProfitSharingRatio** | Pointer to **string** | Profit sharing ratio. 0.1 represents 10% | [optional] [default to ""]
**TodayPnl** | Pointer to **string** | Today pnl | [optional] [default to ""]
**UniqueCode** | Pointer to **string** | Lead trader unique code | [optional] [default to ""]
**Upl** | Pointer to **string** | Unrealized profit &amp; loss | [optional] [default to ""]

## Methods

### NewGetCopytradingCurrentLeadTradersV5RespData

`func NewGetCopytradingCurrentLeadTradersV5RespData() *GetCopytradingCurrentLeadTradersV5RespData`

NewGetCopytradingCurrentLeadTradersV5RespData instantiates a new GetCopytradingCurrentLeadTradersV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCopytradingCurrentLeadTradersV5RespDataWithDefaults

`func NewGetCopytradingCurrentLeadTradersV5RespDataWithDefaults() *GetCopytradingCurrentLeadTradersV5RespData`

NewGetCopytradingCurrentLeadTradersV5RespDataWithDefaults instantiates a new GetCopytradingCurrentLeadTradersV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeginCopyTime

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetBeginCopyTime() string`

GetBeginCopyTime returns the BeginCopyTime field if non-nil, zero value otherwise.

### GetBeginCopyTimeOk

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetBeginCopyTimeOk() (*string, bool)`

GetBeginCopyTimeOk returns a tuple with the BeginCopyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeginCopyTime

`func (o *GetCopytradingCurrentLeadTradersV5RespData) SetBeginCopyTime(v string)`

SetBeginCopyTime sets BeginCopyTime field to given value.

### HasBeginCopyTime

`func (o *GetCopytradingCurrentLeadTradersV5RespData) HasBeginCopyTime() bool`

HasBeginCopyTime returns a boolean if a field has been set.

### GetCcy

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetCopytradingCurrentLeadTradersV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetCopytradingCurrentLeadTradersV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetCopyTotalAmt

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetCopyTotalAmt() string`

GetCopyTotalAmt returns the CopyTotalAmt field if non-nil, zero value otherwise.

### GetCopyTotalAmtOk

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetCopyTotalAmtOk() (*string, bool)`

GetCopyTotalAmtOk returns a tuple with the CopyTotalAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTotalAmt

`func (o *GetCopytradingCurrentLeadTradersV5RespData) SetCopyTotalAmt(v string)`

SetCopyTotalAmt sets CopyTotalAmt field to given value.

### HasCopyTotalAmt

`func (o *GetCopytradingCurrentLeadTradersV5RespData) HasCopyTotalAmt() bool`

HasCopyTotalAmt returns a boolean if a field has been set.

### GetCopyTotalPnl

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetCopyTotalPnl() string`

GetCopyTotalPnl returns the CopyTotalPnl field if non-nil, zero value otherwise.

### GetCopyTotalPnlOk

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetCopyTotalPnlOk() (*string, bool)`

GetCopyTotalPnlOk returns a tuple with the CopyTotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTotalPnl

`func (o *GetCopytradingCurrentLeadTradersV5RespData) SetCopyTotalPnl(v string)`

SetCopyTotalPnl sets CopyTotalPnl field to given value.

### HasCopyTotalPnl

`func (o *GetCopytradingCurrentLeadTradersV5RespData) HasCopyTotalPnl() bool`

HasCopyTotalPnl returns a boolean if a field has been set.

### GetLeadMode

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetLeadMode() string`

GetLeadMode returns the LeadMode field if non-nil, zero value otherwise.

### GetLeadModeOk

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetLeadModeOk() (*string, bool)`

GetLeadModeOk returns a tuple with the LeadMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadMode

`func (o *GetCopytradingCurrentLeadTradersV5RespData) SetLeadMode(v string)`

SetLeadMode sets LeadMode field to given value.

### HasLeadMode

`func (o *GetCopytradingCurrentLeadTradersV5RespData) HasLeadMode() bool`

HasLeadMode returns a boolean if a field has been set.

### GetMargin

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetMargin() string`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetMarginOk() (*string, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *GetCopytradingCurrentLeadTradersV5RespData) SetMargin(v string)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *GetCopytradingCurrentLeadTradersV5RespData) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetNickName

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *GetCopytradingCurrentLeadTradersV5RespData) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *GetCopytradingCurrentLeadTradersV5RespData) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### GetPortLink

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetPortLink() string`

GetPortLink returns the PortLink field if non-nil, zero value otherwise.

### GetPortLinkOk

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetPortLinkOk() (*string, bool)`

GetPortLinkOk returns a tuple with the PortLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLink

`func (o *GetCopytradingCurrentLeadTradersV5RespData) SetPortLink(v string)`

SetPortLink sets PortLink field to given value.

### HasPortLink

`func (o *GetCopytradingCurrentLeadTradersV5RespData) HasPortLink() bool`

HasPortLink returns a boolean if a field has been set.

### GetProfitSharingRatio

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetProfitSharingRatio() string`

GetProfitSharingRatio returns the ProfitSharingRatio field if non-nil, zero value otherwise.

### GetProfitSharingRatioOk

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetProfitSharingRatioOk() (*string, bool)`

GetProfitSharingRatioOk returns a tuple with the ProfitSharingRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitSharingRatio

`func (o *GetCopytradingCurrentLeadTradersV5RespData) SetProfitSharingRatio(v string)`

SetProfitSharingRatio sets ProfitSharingRatio field to given value.

### HasProfitSharingRatio

`func (o *GetCopytradingCurrentLeadTradersV5RespData) HasProfitSharingRatio() bool`

HasProfitSharingRatio returns a boolean if a field has been set.

### GetTodayPnl

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetTodayPnl() string`

GetTodayPnl returns the TodayPnl field if non-nil, zero value otherwise.

### GetTodayPnlOk

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetTodayPnlOk() (*string, bool)`

GetTodayPnlOk returns a tuple with the TodayPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodayPnl

`func (o *GetCopytradingCurrentLeadTradersV5RespData) SetTodayPnl(v string)`

SetTodayPnl sets TodayPnl field to given value.

### HasTodayPnl

`func (o *GetCopytradingCurrentLeadTradersV5RespData) HasTodayPnl() bool`

HasTodayPnl returns a boolean if a field has been set.

### GetUniqueCode

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetUniqueCode() string`

GetUniqueCode returns the UniqueCode field if non-nil, zero value otherwise.

### GetUniqueCodeOk

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetUniqueCodeOk() (*string, bool)`

GetUniqueCodeOk returns a tuple with the UniqueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCode

`func (o *GetCopytradingCurrentLeadTradersV5RespData) SetUniqueCode(v string)`

SetUniqueCode sets UniqueCode field to given value.

### HasUniqueCode

`func (o *GetCopytradingCurrentLeadTradersV5RespData) HasUniqueCode() bool`

HasUniqueCode returns a boolean if a field has been set.

### GetUpl

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *GetCopytradingCurrentLeadTradersV5RespData) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *GetCopytradingCurrentLeadTradersV5RespData) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *GetCopytradingCurrentLeadTradersV5RespData) HasUpl() bool`

HasUpl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


