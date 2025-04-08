# GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgPx** | Pointer to **string** | Average price of recurring buy, quote currency is &#x60;investmentCcy&#x60; | [optional] [default to ""]
**Ccy** | Pointer to **string** | Recurring buy currency, e.g. &#x60;BTC&#x60; | [optional] [default to ""]
**Profit** | Pointer to **string** | Profit in unit of &#x60;investmentCcy&#x60; | [optional] [default to ""]
**Px** | Pointer to **string** | Current market price, quote currency is &#x60;investmentCcy&#x60; | [optional] [default to ""]
**Ratio** | Pointer to **string** | Proportion of recurring currency assets, e.g. \&quot;0.2\&quot; representing 20% | [optional] [default to ""]
**TotalAmt** | Pointer to **string** | Accumulated quantity in unit of recurring buy currency | [optional] [default to ""]

## Methods

### NewGetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner

`func NewGetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner() *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner`

NewGetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner instantiates a new GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInnerWithDefaults

`func NewGetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInnerWithDefaults() *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner`

NewGetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInnerWithDefaults instantiates a new GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgPx

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetCcy

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetProfit

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) GetProfit() string`

GetProfit returns the Profit field if non-nil, zero value otherwise.

### GetProfitOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) GetProfitOk() (*string, bool)`

GetProfitOk returns a tuple with the Profit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfit

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) SetProfit(v string)`

SetProfit sets Profit field to given value.

### HasProfit

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) HasProfit() bool`

HasProfit returns a boolean if a field has been set.

### GetPx

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetRatio

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) GetRatio() string`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) GetRatioOk() (*string, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) SetRatio(v string)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetTotalAmt

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) GetTotalAmt() string`

GetTotalAmt returns the TotalAmt field if non-nil, zero value otherwise.

### GetTotalAmtOk

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) GetTotalAmtOk() (*string, bool)`

GetTotalAmtOk returns a tuple with the TotalAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmt

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) SetTotalAmt(v string)`

SetTotalAmt sets TotalAmt field to given value.

### HasTotalAmt

`func (o *GetTradingBotRecurringOrdersAlgoDetailsV5RespDataRecurringListInner) HasTotalAmt() bool`

HasTotalAmt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


