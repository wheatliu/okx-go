# CreateTradingBotRecurringOrderAlgoV5Req

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoClOrdId** | Pointer to **string** | Client-supplied Algo ID  There will be a value when algo order attaching algoClOrdId is triggered, or it will be \&quot;\&quot;.  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [optional] [default to ""]
**Amt** | **string** | Quantity invested per cycle | [default to ""]
**InvestmentCcy** | **string** | The invested quantity unit, can only be &#x60;USDT&#x60;/&#x60;USDC&#x60; | [default to ""]
**Period** | **string** | Period  &#x60;monthly&#x60;  &#x60;weekly&#x60;  &#x60;daily&#x60;  &#x60;hourly&#x60; | [default to ""]
**RecurringDay** | Pointer to **string** | Recurring buy date  When the period is &#x60;monthly&#x60;, the value range is an integer of [1,28]  When the period is &#x60;weekly&#x60;, the value range is an integer of [1,7]  When the period is &#x60;daily&#x60;/&#x60;hourly&#x60;, the parameter is not required. | [optional] [default to ""]
**RecurringHour** | Pointer to **string** | Recurring buy by hourly  &#x60;1&#x60;/&#x60;4&#x60;/&#x60;8&#x60;/&#x60;12&#x60;, e.g. &#x60;4&#x60; represents \&quot;recurring buy every 4 hour\&quot;  When the period is &#x60;hourly&#x60;, the parameter is required. | [optional] [default to ""]
**RecurringList** | [**[]CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner**](CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner.md) | Recurring buy info | 
**RecurringTime** | **string** | Recurring buy time, the value range is an integer of [0,23]  When the period is &#x60;hourly&#x60;, the parameter is the time of the first investment occurs. | [default to ""]
**StgyName** | **string** | Custom name for trading bot, no more than 40 characters | [default to ""]
**Tag** | Pointer to **string** | Order tag  A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 16 characters. | [optional] [default to ""]
**TdMode** | **string** | Trading mode  Margin mode: &#x60;cross&#x60;  Non-Margin mode: &#x60;cash&#x60; | [default to ""]
**TimeZone** | **string** | UTC time zone, the value range is an integer of [-12,14]  e.g. \&quot;8\&quot; representing UTC+8 (East 8 District), Beijing Time | [default to ""]

## Methods

### NewCreateTradingBotRecurringOrderAlgoV5Req

`func NewCreateTradingBotRecurringOrderAlgoV5Req(amt string, investmentCcy string, period string, recurringList []CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner, recurringTime string, stgyName string, tdMode string, timeZone string, ) *CreateTradingBotRecurringOrderAlgoV5Req`

NewCreateTradingBotRecurringOrderAlgoV5Req instantiates a new CreateTradingBotRecurringOrderAlgoV5Req object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradingBotRecurringOrderAlgoV5ReqWithDefaults

`func NewCreateTradingBotRecurringOrderAlgoV5ReqWithDefaults() *CreateTradingBotRecurringOrderAlgoV5Req`

NewCreateTradingBotRecurringOrderAlgoV5ReqWithDefaults instantiates a new CreateTradingBotRecurringOrderAlgoV5Req object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoClOrdId

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetAlgoClOrdId() string`

GetAlgoClOrdId returns the AlgoClOrdId field if non-nil, zero value otherwise.

### GetAlgoClOrdIdOk

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetAlgoClOrdIdOk() (*string, bool)`

GetAlgoClOrdIdOk returns a tuple with the AlgoClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoClOrdId

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) SetAlgoClOrdId(v string)`

SetAlgoClOrdId sets AlgoClOrdId field to given value.

### HasAlgoClOrdId

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) HasAlgoClOrdId() bool`

HasAlgoClOrdId returns a boolean if a field has been set.

### GetAmt

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetAmt() string`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetAmtOk() (*string, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) SetAmt(v string)`

SetAmt sets Amt field to given value.


### GetInvestmentCcy

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetInvestmentCcy() string`

GetInvestmentCcy returns the InvestmentCcy field if non-nil, zero value otherwise.

### GetInvestmentCcyOk

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetInvestmentCcyOk() (*string, bool)`

GetInvestmentCcyOk returns a tuple with the InvestmentCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentCcy

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) SetInvestmentCcy(v string)`

SetInvestmentCcy sets InvestmentCcy field to given value.


### GetPeriod

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetRecurringDay

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetRecurringDay() string`

GetRecurringDay returns the RecurringDay field if non-nil, zero value otherwise.

### GetRecurringDayOk

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetRecurringDayOk() (*string, bool)`

GetRecurringDayOk returns a tuple with the RecurringDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDay

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) SetRecurringDay(v string)`

SetRecurringDay sets RecurringDay field to given value.

### HasRecurringDay

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) HasRecurringDay() bool`

HasRecurringDay returns a boolean if a field has been set.

### GetRecurringHour

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetRecurringHour() string`

GetRecurringHour returns the RecurringHour field if non-nil, zero value otherwise.

### GetRecurringHourOk

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetRecurringHourOk() (*string, bool)`

GetRecurringHourOk returns a tuple with the RecurringHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringHour

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) SetRecurringHour(v string)`

SetRecurringHour sets RecurringHour field to given value.

### HasRecurringHour

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) HasRecurringHour() bool`

HasRecurringHour returns a boolean if a field has been set.

### GetRecurringList

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetRecurringList() []CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner`

GetRecurringList returns the RecurringList field if non-nil, zero value otherwise.

### GetRecurringListOk

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetRecurringListOk() (*[]CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner, bool)`

GetRecurringListOk returns a tuple with the RecurringList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringList

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) SetRecurringList(v []CreateTradingBotRecurringOrderAlgoV5ReqRecurringListInner)`

SetRecurringList sets RecurringList field to given value.


### GetRecurringTime

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetRecurringTime() string`

GetRecurringTime returns the RecurringTime field if non-nil, zero value otherwise.

### GetRecurringTimeOk

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetRecurringTimeOk() (*string, bool)`

GetRecurringTimeOk returns a tuple with the RecurringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTime

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) SetRecurringTime(v string)`

SetRecurringTime sets RecurringTime field to given value.


### GetStgyName

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetStgyName() string`

GetStgyName returns the StgyName field if non-nil, zero value otherwise.

### GetStgyNameOk

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetStgyNameOk() (*string, bool)`

GetStgyNameOk returns a tuple with the StgyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStgyName

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) SetStgyName(v string)`

SetStgyName sets StgyName field to given value.


### GetTag

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTdMode

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetTdMode() string`

GetTdMode returns the TdMode field if non-nil, zero value otherwise.

### GetTdModeOk

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetTdModeOk() (*string, bool)`

GetTdModeOk returns a tuple with the TdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdMode

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) SetTdMode(v string)`

SetTdMode sets TdMode field to given value.


### GetTimeZone

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CreateTradingBotRecurringOrderAlgoV5Req) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


