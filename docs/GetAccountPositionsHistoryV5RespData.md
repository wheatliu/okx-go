# GetAccountPositionsHistoryV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CTime** | Pointer to **string** | Created time of position | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency used for margin | [optional] [default to ""]
**CloseAvgPx** | Pointer to **string** | Average price of closing position | [optional] [default to ""]
**CloseTotalPos** | Pointer to **string** | Position&#39;s cumulative closed volume | [optional] [default to ""]
**Direction** | Pointer to **string** | Direction: &#x60;long&#x60; &#x60;short&#x60;  Only applicable to &#x60;MARGIN/FUTURES/SWAP/OPTION&#x60; | [optional] [default to ""]
**Fee** | Pointer to **string** | Accumulated fee  Negative number represents the user transaction fee charged by the platform.Positive number represents rebate. | [optional] [default to ""]
**FundingFee** | Pointer to **string** | Accumulated funding fee | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage | [optional] [default to ""]
**LiqPenalty** | Pointer to **string** | Accumulated liquidation penalty. It is negative when there is a value. | [optional] [default to ""]
**MgnMode** | Pointer to **string** | Margin mode  &#x60;cross&#x60; &#x60;isolated&#x60; | [optional] [default to ""]
**NonSettleAvgPx** | Pointer to **string** | Non-settlement entry price  The non-settlement entry price only reflects the average price at which the position is opened or increased.  Only applicable to &#x60;cross&#x60; &#x60;FUTURES&#x60; | [optional] [default to ""]
**OpenAvgPx** | Pointer to **string** | Average price of opening position  Under cross-margin mode, the entry price of expiry futures will update at settlement to the last settlement price, and when the position is opened or increased. | [optional] [default to ""]
**OpenMaxPos** | Pointer to **string** | Max quantity of position | [optional] [default to ""]
**Pnl** | Pointer to **string** | Profit and loss | [optional] [default to ""]
**PnlRatio** | Pointer to **string** | Realized P&amp;L ratio | [optional] [default to ""]
**PosId** | Pointer to **string** | Position ID | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position mode side   &#x60;long&#x60;: Hedge mode long   &#x60;short&#x60;: Hedge mode short   &#x60;net&#x60;: Net mode | [optional] [default to ""]
**RealizedPnl** | Pointer to **string** | Realized profit and loss  Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;  &#x60;realizedPnl&#x60;&#x3D;&#x60;pnl&#x60;+&#x60;fee&#x60;+&#x60;fundingFee&#x60;+&#x60;liqPenalty&#x60;+&#x60;settledPnl&#x60; | [optional] [default to ""]
**SettledPnl** | Pointer to **string** | Accumulated settled profit and loss (calculated by settlement price)  Only applicable to &#x60;cross&#x60; &#x60;FUTURES&#x60; | [optional] [default to ""]
**TriggerPx** | Pointer to **string** | trigger mark price. There is value when &#x60;type&#x60; is equal to &#x60;3&#x60;, &#x60;4&#x60; or &#x60;5&#x60;. It is \&quot;\&quot; when &#x60;type&#x60; is equal to &#x60;1&#x60; or &#x60;2&#x60; | [optional] [default to ""]
**Type** | Pointer to **string** | The type of latest close position  &#x60;1&#x60;：Close position partially;&#x60;2&#x60;：Close all;&#x60;3&#x60;：Liquidation;&#x60;4&#x60;：Partial liquidation; &#x60;5&#x60;：ADL;   It is the latest type if there are several types for the same position. | [optional] [default to ""]
**UTime** | Pointer to **string** | Updated time of position | [optional] [default to ""]
**Uly** | Pointer to **string** | Underlying | [optional] [default to ""]

## Methods

### NewGetAccountPositionsHistoryV5RespData

`func NewGetAccountPositionsHistoryV5RespData() *GetAccountPositionsHistoryV5RespData`

NewGetAccountPositionsHistoryV5RespData instantiates a new GetAccountPositionsHistoryV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountPositionsHistoryV5RespDataWithDefaults

`func NewGetAccountPositionsHistoryV5RespDataWithDefaults() *GetAccountPositionsHistoryV5RespData`

NewGetAccountPositionsHistoryV5RespDataWithDefaults instantiates a new GetAccountPositionsHistoryV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCTime

`func (o *GetAccountPositionsHistoryV5RespData) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetAccountPositionsHistoryV5RespData) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetAccountPositionsHistoryV5RespData) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetAccountPositionsHistoryV5RespData) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCcy

`func (o *GetAccountPositionsHistoryV5RespData) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountPositionsHistoryV5RespData) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountPositionsHistoryV5RespData) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountPositionsHistoryV5RespData) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetCloseAvgPx

`func (o *GetAccountPositionsHistoryV5RespData) GetCloseAvgPx() string`

GetCloseAvgPx returns the CloseAvgPx field if non-nil, zero value otherwise.

### GetCloseAvgPxOk

`func (o *GetAccountPositionsHistoryV5RespData) GetCloseAvgPxOk() (*string, bool)`

GetCloseAvgPxOk returns a tuple with the CloseAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseAvgPx

`func (o *GetAccountPositionsHistoryV5RespData) SetCloseAvgPx(v string)`

SetCloseAvgPx sets CloseAvgPx field to given value.

### HasCloseAvgPx

`func (o *GetAccountPositionsHistoryV5RespData) HasCloseAvgPx() bool`

HasCloseAvgPx returns a boolean if a field has been set.

### GetCloseTotalPos

`func (o *GetAccountPositionsHistoryV5RespData) GetCloseTotalPos() string`

GetCloseTotalPos returns the CloseTotalPos field if non-nil, zero value otherwise.

### GetCloseTotalPosOk

`func (o *GetAccountPositionsHistoryV5RespData) GetCloseTotalPosOk() (*string, bool)`

GetCloseTotalPosOk returns a tuple with the CloseTotalPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTotalPos

`func (o *GetAccountPositionsHistoryV5RespData) SetCloseTotalPos(v string)`

SetCloseTotalPos sets CloseTotalPos field to given value.

### HasCloseTotalPos

`func (o *GetAccountPositionsHistoryV5RespData) HasCloseTotalPos() bool`

HasCloseTotalPos returns a boolean if a field has been set.

### GetDirection

`func (o *GetAccountPositionsHistoryV5RespData) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetAccountPositionsHistoryV5RespData) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetAccountPositionsHistoryV5RespData) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetAccountPositionsHistoryV5RespData) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFee

`func (o *GetAccountPositionsHistoryV5RespData) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetAccountPositionsHistoryV5RespData) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetAccountPositionsHistoryV5RespData) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetAccountPositionsHistoryV5RespData) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFundingFee

`func (o *GetAccountPositionsHistoryV5RespData) GetFundingFee() string`

GetFundingFee returns the FundingFee field if non-nil, zero value otherwise.

### GetFundingFeeOk

`func (o *GetAccountPositionsHistoryV5RespData) GetFundingFeeOk() (*string, bool)`

GetFundingFeeOk returns a tuple with the FundingFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingFee

`func (o *GetAccountPositionsHistoryV5RespData) SetFundingFee(v string)`

SetFundingFee sets FundingFee field to given value.

### HasFundingFee

`func (o *GetAccountPositionsHistoryV5RespData) HasFundingFee() bool`

HasFundingFee returns a boolean if a field has been set.

### GetInstId

`func (o *GetAccountPositionsHistoryV5RespData) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAccountPositionsHistoryV5RespData) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAccountPositionsHistoryV5RespData) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAccountPositionsHistoryV5RespData) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetAccountPositionsHistoryV5RespData) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetAccountPositionsHistoryV5RespData) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetAccountPositionsHistoryV5RespData) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetAccountPositionsHistoryV5RespData) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetLever

`func (o *GetAccountPositionsHistoryV5RespData) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetAccountPositionsHistoryV5RespData) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetAccountPositionsHistoryV5RespData) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetAccountPositionsHistoryV5RespData) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetLiqPenalty

`func (o *GetAccountPositionsHistoryV5RespData) GetLiqPenalty() string`

GetLiqPenalty returns the LiqPenalty field if non-nil, zero value otherwise.

### GetLiqPenaltyOk

`func (o *GetAccountPositionsHistoryV5RespData) GetLiqPenaltyOk() (*string, bool)`

GetLiqPenaltyOk returns a tuple with the LiqPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqPenalty

`func (o *GetAccountPositionsHistoryV5RespData) SetLiqPenalty(v string)`

SetLiqPenalty sets LiqPenalty field to given value.

### HasLiqPenalty

`func (o *GetAccountPositionsHistoryV5RespData) HasLiqPenalty() bool`

HasLiqPenalty returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetAccountPositionsHistoryV5RespData) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetAccountPositionsHistoryV5RespData) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetAccountPositionsHistoryV5RespData) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetAccountPositionsHistoryV5RespData) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetNonSettleAvgPx

`func (o *GetAccountPositionsHistoryV5RespData) GetNonSettleAvgPx() string`

GetNonSettleAvgPx returns the NonSettleAvgPx field if non-nil, zero value otherwise.

### GetNonSettleAvgPxOk

`func (o *GetAccountPositionsHistoryV5RespData) GetNonSettleAvgPxOk() (*string, bool)`

GetNonSettleAvgPxOk returns a tuple with the NonSettleAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSettleAvgPx

`func (o *GetAccountPositionsHistoryV5RespData) SetNonSettleAvgPx(v string)`

SetNonSettleAvgPx sets NonSettleAvgPx field to given value.

### HasNonSettleAvgPx

`func (o *GetAccountPositionsHistoryV5RespData) HasNonSettleAvgPx() bool`

HasNonSettleAvgPx returns a boolean if a field has been set.

### GetOpenAvgPx

`func (o *GetAccountPositionsHistoryV5RespData) GetOpenAvgPx() string`

GetOpenAvgPx returns the OpenAvgPx field if non-nil, zero value otherwise.

### GetOpenAvgPxOk

`func (o *GetAccountPositionsHistoryV5RespData) GetOpenAvgPxOk() (*string, bool)`

GetOpenAvgPxOk returns a tuple with the OpenAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAvgPx

`func (o *GetAccountPositionsHistoryV5RespData) SetOpenAvgPx(v string)`

SetOpenAvgPx sets OpenAvgPx field to given value.

### HasOpenAvgPx

`func (o *GetAccountPositionsHistoryV5RespData) HasOpenAvgPx() bool`

HasOpenAvgPx returns a boolean if a field has been set.

### GetOpenMaxPos

`func (o *GetAccountPositionsHistoryV5RespData) GetOpenMaxPos() string`

GetOpenMaxPos returns the OpenMaxPos field if non-nil, zero value otherwise.

### GetOpenMaxPosOk

`func (o *GetAccountPositionsHistoryV5RespData) GetOpenMaxPosOk() (*string, bool)`

GetOpenMaxPosOk returns a tuple with the OpenMaxPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenMaxPos

`func (o *GetAccountPositionsHistoryV5RespData) SetOpenMaxPos(v string)`

SetOpenMaxPos sets OpenMaxPos field to given value.

### HasOpenMaxPos

`func (o *GetAccountPositionsHistoryV5RespData) HasOpenMaxPos() bool`

HasOpenMaxPos returns a boolean if a field has been set.

### GetPnl

`func (o *GetAccountPositionsHistoryV5RespData) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *GetAccountPositionsHistoryV5RespData) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *GetAccountPositionsHistoryV5RespData) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *GetAccountPositionsHistoryV5RespData) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetPnlRatio

`func (o *GetAccountPositionsHistoryV5RespData) GetPnlRatio() string`

GetPnlRatio returns the PnlRatio field if non-nil, zero value otherwise.

### GetPnlRatioOk

`func (o *GetAccountPositionsHistoryV5RespData) GetPnlRatioOk() (*string, bool)`

GetPnlRatioOk returns a tuple with the PnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnlRatio

`func (o *GetAccountPositionsHistoryV5RespData) SetPnlRatio(v string)`

SetPnlRatio sets PnlRatio field to given value.

### HasPnlRatio

`func (o *GetAccountPositionsHistoryV5RespData) HasPnlRatio() bool`

HasPnlRatio returns a boolean if a field has been set.

### GetPosId

`func (o *GetAccountPositionsHistoryV5RespData) GetPosId() string`

GetPosId returns the PosId field if non-nil, zero value otherwise.

### GetPosIdOk

`func (o *GetAccountPositionsHistoryV5RespData) GetPosIdOk() (*string, bool)`

GetPosIdOk returns a tuple with the PosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosId

`func (o *GetAccountPositionsHistoryV5RespData) SetPosId(v string)`

SetPosId sets PosId field to given value.

### HasPosId

`func (o *GetAccountPositionsHistoryV5RespData) HasPosId() bool`

HasPosId returns a boolean if a field has been set.

### GetPosSide

`func (o *GetAccountPositionsHistoryV5RespData) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetAccountPositionsHistoryV5RespData) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetAccountPositionsHistoryV5RespData) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetAccountPositionsHistoryV5RespData) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *GetAccountPositionsHistoryV5RespData) GetRealizedPnl() string`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *GetAccountPositionsHistoryV5RespData) GetRealizedPnlOk() (*string, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *GetAccountPositionsHistoryV5RespData) SetRealizedPnl(v string)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *GetAccountPositionsHistoryV5RespData) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetSettledPnl

`func (o *GetAccountPositionsHistoryV5RespData) GetSettledPnl() string`

GetSettledPnl returns the SettledPnl field if non-nil, zero value otherwise.

### GetSettledPnlOk

`func (o *GetAccountPositionsHistoryV5RespData) GetSettledPnlOk() (*string, bool)`

GetSettledPnlOk returns a tuple with the SettledPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledPnl

`func (o *GetAccountPositionsHistoryV5RespData) SetSettledPnl(v string)`

SetSettledPnl sets SettledPnl field to given value.

### HasSettledPnl

`func (o *GetAccountPositionsHistoryV5RespData) HasSettledPnl() bool`

HasSettledPnl returns a boolean if a field has been set.

### GetTriggerPx

`func (o *GetAccountPositionsHistoryV5RespData) GetTriggerPx() string`

GetTriggerPx returns the TriggerPx field if non-nil, zero value otherwise.

### GetTriggerPxOk

`func (o *GetAccountPositionsHistoryV5RespData) GetTriggerPxOk() (*string, bool)`

GetTriggerPxOk returns a tuple with the TriggerPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPx

`func (o *GetAccountPositionsHistoryV5RespData) SetTriggerPx(v string)`

SetTriggerPx sets TriggerPx field to given value.

### HasTriggerPx

`func (o *GetAccountPositionsHistoryV5RespData) HasTriggerPx() bool`

HasTriggerPx returns a boolean if a field has been set.

### GetType

`func (o *GetAccountPositionsHistoryV5RespData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAccountPositionsHistoryV5RespData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAccountPositionsHistoryV5RespData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAccountPositionsHistoryV5RespData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUTime

`func (o *GetAccountPositionsHistoryV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetAccountPositionsHistoryV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetAccountPositionsHistoryV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetAccountPositionsHistoryV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUly

`func (o *GetAccountPositionsHistoryV5RespData) GetUly() string`

GetUly returns the Uly field if non-nil, zero value otherwise.

### GetUlyOk

`func (o *GetAccountPositionsHistoryV5RespData) GetUlyOk() (*string, bool)`

GetUlyOk returns a tuple with the Uly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUly

`func (o *GetAccountPositionsHistoryV5RespData) SetUly(v string)`

SetUly sets Uly field to given value.

### HasUly

`func (o *GetAccountPositionsHistoryV5RespData) HasUly() bool`

HasUly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


