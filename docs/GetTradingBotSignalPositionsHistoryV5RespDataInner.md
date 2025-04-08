# GetTradingBotSignalPositionsHistoryV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CTime** | Pointer to **string** | Created time of position | [optional] [default to ""]
**CloseAvgPx** | Pointer to **string** | Average price of closing position | [optional] [default to ""]
**Direction** | Pointer to **string** | Direction: &#x60;long&#x60; &#x60;short&#x60; | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage | [optional] [default to ""]
**MgnMode** | Pointer to **string** | Margin mode &#x60;cross&#x60; &#x60;isolated&#x60; | [optional] [default to ""]
**OpenAvgPx** | Pointer to **string** | Average price of opening position | [optional] [default to ""]
**Pnl** | Pointer to **string** | Profit and loss | [optional] [default to ""]
**PnlRatio** | Pointer to **string** | P&amp;L ratio | [optional] [default to ""]
**UTime** | Pointer to **string** | Updated time of position | [optional] [default to ""]
**Uly** | Pointer to **string** | Underlying | [optional] [default to ""]

## Methods

### NewGetTradingBotSignalPositionsHistoryV5RespDataInner

`func NewGetTradingBotSignalPositionsHistoryV5RespDataInner() *GetTradingBotSignalPositionsHistoryV5RespDataInner`

NewGetTradingBotSignalPositionsHistoryV5RespDataInner instantiates a new GetTradingBotSignalPositionsHistoryV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTradingBotSignalPositionsHistoryV5RespDataInnerWithDefaults

`func NewGetTradingBotSignalPositionsHistoryV5RespDataInnerWithDefaults() *GetTradingBotSignalPositionsHistoryV5RespDataInner`

NewGetTradingBotSignalPositionsHistoryV5RespDataInnerWithDefaults instantiates a new GetTradingBotSignalPositionsHistoryV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCTime

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCloseAvgPx

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetCloseAvgPx() string`

GetCloseAvgPx returns the CloseAvgPx field if non-nil, zero value otherwise.

### GetCloseAvgPxOk

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetCloseAvgPxOk() (*string, bool)`

GetCloseAvgPxOk returns a tuple with the CloseAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseAvgPx

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) SetCloseAvgPx(v string)`

SetCloseAvgPx sets CloseAvgPx field to given value.

### HasCloseAvgPx

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) HasCloseAvgPx() bool`

HasCloseAvgPx returns a boolean if a field has been set.

### GetDirection

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetInstId

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetLever

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetOpenAvgPx

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetOpenAvgPx() string`

GetOpenAvgPx returns the OpenAvgPx field if non-nil, zero value otherwise.

### GetOpenAvgPxOk

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetOpenAvgPxOk() (*string, bool)`

GetOpenAvgPxOk returns a tuple with the OpenAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAvgPx

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) SetOpenAvgPx(v string)`

SetOpenAvgPx sets OpenAvgPx field to given value.

### HasOpenAvgPx

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) HasOpenAvgPx() bool`

HasOpenAvgPx returns a boolean if a field has been set.

### GetPnl

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetPnlRatio

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetPnlRatio() string`

GetPnlRatio returns the PnlRatio field if non-nil, zero value otherwise.

### GetPnlRatioOk

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetPnlRatioOk() (*string, bool)`

GetPnlRatioOk returns a tuple with the PnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlRatio

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) SetPnlRatio(v string)`

SetPnlRatio sets PnlRatio field to given value.

### HasPnlRatio

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) HasPnlRatio() bool`

HasPnlRatio returns a boolean if a field has been set.

### GetUTime

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUly

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetUly() string`

GetUly returns the Uly field if non-nil, zero value otherwise.

### GetUlyOk

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) GetUlyOk() (*string, bool)`

GetUlyOk returns a tuple with the Uly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUly

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) SetUly(v string)`

SetUly sets Uly field to given value.

### HasUly

`func (o *GetTradingBotSignalPositionsHistoryV5RespDataInner) HasUly() bool`

HasUly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


