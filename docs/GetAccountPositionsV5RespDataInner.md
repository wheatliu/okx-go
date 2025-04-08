# GetAccountPositionsV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adl** | Pointer to **string** | Auto-deleveraging (ADL) indicator  Divided into 5 levels, from 1 to 5, the smaller the number, the weaker the adl intensity. | [optional] [default to ""]
**AvailPos** | Pointer to **string** | Position that can be closed   Only applicable to &#x60;MARGIN&#x60; and &#x60;OPTION&#x60;.  For &#x60;MARGIN&#x60; position, the rest of sz will be &#x60;SPOT&#x60; trading after the liability is repaid while closing the position. Please get the available reduce-only amount from \&quot;Get maximum available tradable amount\&quot; if you want to reduce the amount of &#x60;SPOT&#x60; trading as much as possible. | [optional] [default to ""]
**AvgPx** | Pointer to **string** | Average open price  Under cross-margin mode, the entry price of expiry futures will update at settlement to the last settlement price, and when the position is opened or increased. | [optional] [default to ""]
**BaseBal** | Pointer to **string** | Base currency balance, only applicable to &#x60;MARGIN&#x60;（Quick Margin Mode）(Deprecated) | [optional] [default to ""]
**BaseBorrowed** | Pointer to **string** | Base currency amount already borrowed, only applicable to MARGIN(Quick Margin Mode）(Deprecated) | [optional] [default to ""]
**BaseInterest** | Pointer to **string** | Base Interest, undeducted interest that has been incurred, only applicable to MARGIN(Quick Margin Mode）(Deprecated) | [optional] [default to ""]
**BePx** | Pointer to **string** | Breakeven price | [optional] [default to ""]
**BizRefId** | Pointer to **string** | External business id, e.g. experience coupon id | [optional] [default to ""]
**BizRefType** | Pointer to **string** | External business type | [optional] [default to ""]
**CTime** | Pointer to **string** | Creation time, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency used for margin | [optional] [default to ""]
**ClSpotInUseAmt** | Pointer to **string** | User-defined spot risk offset amount  Applicable to &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**CloseOrderAlgo** | Pointer to [**[]GetAccountPositionsV5RespDataInnerCloseOrderAlgoInner**](GetAccountPositionsV5RespDataInnerCloseOrderAlgoInner.md) | Close position algo orders attached to the position. This array will have values only after you request \&quot;Place algo order\&quot; with &#x60;closeFraction&#x60;&#x3D;1. | [optional] 
**DeltaBS** | Pointer to **string** | delta: Black-Scholes Greeks in dollars, only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**DeltaPA** | Pointer to **string** | delta: Greeks in coins, only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**Fee** | Pointer to **string** | Accumulated fee  Negative number represents the user transaction fee charged by the platform.Positive number represents rebate. | [optional] [default to ""]
**FundingFee** | Pointer to **string** | Accumulated funding fee | [optional] [default to ""]
**GammaBS** | Pointer to **string** | gamma: Black-Scholes Greeks in dollars, only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**GammaPA** | Pointer to **string** | gamma: Greeks in coins, only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**IdxPx** | Pointer to **string** | Latest underlying index price | [optional] [default to ""]
**Imr** | Pointer to **string** | Initial margin requirement, only applicable to &#x60;cross&#x60;. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Interest** | Pointer to **string** | Interest. Undeducted interest that has been incurred. | [optional] [default to ""]
**Last** | Pointer to **string** | Latest traded price | [optional] [default to ""]
**Lever** | Pointer to **string** | Leverage   Not applicable to &#x60;OPTION&#x60; and positions of cross margin mode under &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**Liab** | Pointer to **string** | Liabilities, only applicable to &#x60;MARGIN&#x60;. | [optional] [default to ""]
**LiabCcy** | Pointer to **string** | Liabilities currency, only applicable to &#x60;MARGIN&#x60;. | [optional] [default to ""]
**LiqPenalty** | Pointer to **string** | Accumulated liquidation penalty. It is negative when there is a value. | [optional] [default to ""]
**LiqPx** | Pointer to **string** | Estimated liquidation price    Not applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**Margin** | Pointer to **string** | Margin, can be added or reduced. Only applicable to &#x60;isolated&#x60;. | [optional] [default to ""]
**MarkPx** | Pointer to **string** | Latest Mark price | [optional] [default to ""]
**MaxSpotInUseAmt** | Pointer to **string** | Max possible spot risk offset amount  Applicable to &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**MgnMode** | Pointer to **string** | Margin mode  &#x60;cross&#x60;    &#x60;isolated&#x60; | [optional] [default to ""]
**MgnRatio** | Pointer to **string** | Margin ratio | [optional] [default to ""]
**Mmr** | Pointer to **string** | Maintenance margin requirement | [optional] [default to ""]
**NonSettleAvgPx** | Pointer to **string** | Non-settlement entry price  The non-settlement entry price only reflects the average price at which the position is opened or increased.  Applicable to &#x60;cross&#x60; &#x60;FUTURES&#x60; positions. | [optional] [default to ""]
**NotionalUsd** | Pointer to **string** | Notional value of positions in &#x60;USD&#x60; | [optional] [default to ""]
**OptVal** | Pointer to **string** | Option Value, only applicable to &#x60;OPTION&#x60;. | [optional] [default to ""]
**PendingCloseOrdLiabVal** | Pointer to **string** | The amount of close orders of isolated margin liability. | [optional] [default to ""]
**Pnl** | Pointer to **string** | Accumulated pnl of closing order(s) | [optional] [default to ""]
**Pos** | Pointer to **string** | Quantity of positions. In the isolated margin mode, when doing manual transfers, a position with pos of &#x60;0&#x60; will be generated after the deposit is transferred | [optional] [default to ""]
**PosCcy** | Pointer to **string** | Position currency, only applicable to &#x60;MARGIN&#x60; positions. | [optional] [default to ""]
**PosId** | Pointer to **string** | Position ID | [optional] [default to ""]
**PosSide** | Pointer to **string** | Position side  &#x60;long&#x60;, &#x60;pos&#x60; is positive    &#x60;short&#x60;, &#x60;pos&#x60; is positive    &#x60;net&#x60; (&#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;: positive &#x60;pos&#x60; means long position and negative &#x60;pos&#x60; means short position. For &#x60;MARGIN&#x60;, &#x60;pos&#x60; is always positive, &#x60;posCcy&#x60; being base currency means long position, &#x60;posCcy&#x60; being quote currency means short position.) | [optional] [default to ""]
**QuoteBal** | Pointer to **string** | Quote currency balance, only applicable to &#x60;MARGIN&#x60;（Quick Margin Mode）(Deprecated) | [optional] [default to ""]
**QuoteBorrowed** | Pointer to **string** | Quote currency amount already borrowed, only applicable to MARGIN(Quick Margin Mode）(Deprecated) | [optional] [default to ""]
**QuoteInterest** | Pointer to **string** | Quote Interest, undeducted interest that has been incurred, only applicable to MARGIN(Quick Margin Mode）(Deprecated) | [optional] [default to ""]
**RealizedPnl** | Pointer to **string** | Realized profit and loss  Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;  &#x60;realizedPnl&#x60;&#x3D;&#x60;pnl&#x60;+&#x60;fee&#x60;+&#x60;fundingFee&#x60;+&#x60;liqPenalty&#x60;+&#x60;settledPnl&#x60; | [optional] [default to ""]
**SettledPnl** | Pointer to **string** | Accumulated settled profit and loss (calculated by settlement price)  Only applicable to &#x60;cross&#x60; &#x60;FUTURES&#x60; | [optional] [default to ""]
**SpotInUseAmt** | Pointer to **string** | Spot in use amount  Applicable to &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**SpotInUseCcy** | Pointer to **string** | Spot in use unit, e.g. &#x60;BTC&#x60;  Applicable to &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**ThetaBS** | Pointer to **string** | theta：Black-Scholes Greeks in dollars, only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**ThetaPA** | Pointer to **string** | theta：Greeks in coins, only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last trade ID | [optional] [default to ""]
**UTime** | Pointer to **string** | Latest time position was adjusted, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Upl** | Pointer to **string** | Unrealized profit and loss calculated by mark price. | [optional] [default to ""]
**UplLastPx** | Pointer to **string** | Unrealized profit and loss calculated by last price. Main usage is showing, actual value is upl. | [optional] [default to ""]
**UplRatio** | Pointer to **string** | Unrealized profit and loss ratio calculated by mark price. | [optional] [default to ""]
**UplRatioLastPx** | Pointer to **string** | Unrealized profit and loss ratio calculated by last price. | [optional] [default to ""]
**UsdPx** | Pointer to **string** | Latest USD price of the &#x60;ccy&#x60; on the market, only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**VegaBS** | Pointer to **string** | vega：Black-Scholes Greeks in dollars, only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]
**VegaPA** | Pointer to **string** | vega：Greeks in coins, only applicable to &#x60;OPTION&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountPositionsV5RespDataInner

`func NewGetAccountPositionsV5RespDataInner() *GetAccountPositionsV5RespDataInner`

NewGetAccountPositionsV5RespDataInner instantiates a new GetAccountPositionsV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountPositionsV5RespDataInnerWithDefaults

`func NewGetAccountPositionsV5RespDataInnerWithDefaults() *GetAccountPositionsV5RespDataInner`

NewGetAccountPositionsV5RespDataInnerWithDefaults instantiates a new GetAccountPositionsV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdl

`func (o *GetAccountPositionsV5RespDataInner) GetAdl() string`

GetAdl returns the Adl field if non-nil, zero value otherwise.

### GetAdlOk

`func (o *GetAccountPositionsV5RespDataInner) GetAdlOk() (*string, bool)`

GetAdlOk returns a tuple with the Adl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdl

`func (o *GetAccountPositionsV5RespDataInner) SetAdl(v string)`

SetAdl sets Adl field to given value.

### HasAdl

`func (o *GetAccountPositionsV5RespDataInner) HasAdl() bool`

HasAdl returns a boolean if a field has been set.

### GetAvailPos

`func (o *GetAccountPositionsV5RespDataInner) GetAvailPos() string`

GetAvailPos returns the AvailPos field if non-nil, zero value otherwise.

### GetAvailPosOk

`func (o *GetAccountPositionsV5RespDataInner) GetAvailPosOk() (*string, bool)`

GetAvailPosOk returns a tuple with the AvailPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailPos

`func (o *GetAccountPositionsV5RespDataInner) SetAvailPos(v string)`

SetAvailPos sets AvailPos field to given value.

### HasAvailPos

`func (o *GetAccountPositionsV5RespDataInner) HasAvailPos() bool`

HasAvailPos returns a boolean if a field has been set.

### GetAvgPx

`func (o *GetAccountPositionsV5RespDataInner) GetAvgPx() string`

GetAvgPx returns the AvgPx field if non-nil, zero value otherwise.

### GetAvgPxOk

`func (o *GetAccountPositionsV5RespDataInner) GetAvgPxOk() (*string, bool)`

GetAvgPxOk returns a tuple with the AvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPx

`func (o *GetAccountPositionsV5RespDataInner) SetAvgPx(v string)`

SetAvgPx sets AvgPx field to given value.

### HasAvgPx

`func (o *GetAccountPositionsV5RespDataInner) HasAvgPx() bool`

HasAvgPx returns a boolean if a field has been set.

### GetBaseBal

`func (o *GetAccountPositionsV5RespDataInner) GetBaseBal() string`

GetBaseBal returns the BaseBal field if non-nil, zero value otherwise.

### GetBaseBalOk

`func (o *GetAccountPositionsV5RespDataInner) GetBaseBalOk() (*string, bool)`

GetBaseBalOk returns a tuple with the BaseBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBal

`func (o *GetAccountPositionsV5RespDataInner) SetBaseBal(v string)`

SetBaseBal sets BaseBal field to given value.

### HasBaseBal

`func (o *GetAccountPositionsV5RespDataInner) HasBaseBal() bool`

HasBaseBal returns a boolean if a field has been set.

### GetBaseBorrowed

`func (o *GetAccountPositionsV5RespDataInner) GetBaseBorrowed() string`

GetBaseBorrowed returns the BaseBorrowed field if non-nil, zero value otherwise.

### GetBaseBorrowedOk

`func (o *GetAccountPositionsV5RespDataInner) GetBaseBorrowedOk() (*string, bool)`

GetBaseBorrowedOk returns a tuple with the BaseBorrowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBorrowed

`func (o *GetAccountPositionsV5RespDataInner) SetBaseBorrowed(v string)`

SetBaseBorrowed sets BaseBorrowed field to given value.

### HasBaseBorrowed

`func (o *GetAccountPositionsV5RespDataInner) HasBaseBorrowed() bool`

HasBaseBorrowed returns a boolean if a field has been set.

### GetBaseInterest

`func (o *GetAccountPositionsV5RespDataInner) GetBaseInterest() string`

GetBaseInterest returns the BaseInterest field if non-nil, zero value otherwise.

### GetBaseInterestOk

`func (o *GetAccountPositionsV5RespDataInner) GetBaseInterestOk() (*string, bool)`

GetBaseInterestOk returns a tuple with the BaseInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseInterest

`func (o *GetAccountPositionsV5RespDataInner) SetBaseInterest(v string)`

SetBaseInterest sets BaseInterest field to given value.

### HasBaseInterest

`func (o *GetAccountPositionsV5RespDataInner) HasBaseInterest() bool`

HasBaseInterest returns a boolean if a field has been set.

### GetBePx

`func (o *GetAccountPositionsV5RespDataInner) GetBePx() string`

GetBePx returns the BePx field if non-nil, zero value otherwise.

### GetBePxOk

`func (o *GetAccountPositionsV5RespDataInner) GetBePxOk() (*string, bool)`

GetBePxOk returns a tuple with the BePx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBePx

`func (o *GetAccountPositionsV5RespDataInner) SetBePx(v string)`

SetBePx sets BePx field to given value.

### HasBePx

`func (o *GetAccountPositionsV5RespDataInner) HasBePx() bool`

HasBePx returns a boolean if a field has been set.

### GetBizRefId

`func (o *GetAccountPositionsV5RespDataInner) GetBizRefId() string`

GetBizRefId returns the BizRefId field if non-nil, zero value otherwise.

### GetBizRefIdOk

`func (o *GetAccountPositionsV5RespDataInner) GetBizRefIdOk() (*string, bool)`

GetBizRefIdOk returns a tuple with the BizRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizRefId

`func (o *GetAccountPositionsV5RespDataInner) SetBizRefId(v string)`

SetBizRefId sets BizRefId field to given value.

### HasBizRefId

`func (o *GetAccountPositionsV5RespDataInner) HasBizRefId() bool`

HasBizRefId returns a boolean if a field has been set.

### GetBizRefType

`func (o *GetAccountPositionsV5RespDataInner) GetBizRefType() string`

GetBizRefType returns the BizRefType field if non-nil, zero value otherwise.

### GetBizRefTypeOk

`func (o *GetAccountPositionsV5RespDataInner) GetBizRefTypeOk() (*string, bool)`

GetBizRefTypeOk returns a tuple with the BizRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizRefType

`func (o *GetAccountPositionsV5RespDataInner) SetBizRefType(v string)`

SetBizRefType sets BizRefType field to given value.

### HasBizRefType

`func (o *GetAccountPositionsV5RespDataInner) HasBizRefType() bool`

HasBizRefType returns a boolean if a field has been set.

### GetCTime

`func (o *GetAccountPositionsV5RespDataInner) GetCTime() string`

GetCTime returns the CTime field if non-nil, zero value otherwise.

### GetCTimeOk

`func (o *GetAccountPositionsV5RespDataInner) GetCTimeOk() (*string, bool)`

GetCTimeOk returns a tuple with the CTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTime

`func (o *GetAccountPositionsV5RespDataInner) SetCTime(v string)`

SetCTime sets CTime field to given value.

### HasCTime

`func (o *GetAccountPositionsV5RespDataInner) HasCTime() bool`

HasCTime returns a boolean if a field has been set.

### GetCcy

`func (o *GetAccountPositionsV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountPositionsV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountPositionsV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountPositionsV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClSpotInUseAmt

`func (o *GetAccountPositionsV5RespDataInner) GetClSpotInUseAmt() string`

GetClSpotInUseAmt returns the ClSpotInUseAmt field if non-nil, zero value otherwise.

### GetClSpotInUseAmtOk

`func (o *GetAccountPositionsV5RespDataInner) GetClSpotInUseAmtOk() (*string, bool)`

GetClSpotInUseAmtOk returns a tuple with the ClSpotInUseAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClSpotInUseAmt

`func (o *GetAccountPositionsV5RespDataInner) SetClSpotInUseAmt(v string)`

SetClSpotInUseAmt sets ClSpotInUseAmt field to given value.

### HasClSpotInUseAmt

`func (o *GetAccountPositionsV5RespDataInner) HasClSpotInUseAmt() bool`

HasClSpotInUseAmt returns a boolean if a field has been set.

### GetCloseOrderAlgo

`func (o *GetAccountPositionsV5RespDataInner) GetCloseOrderAlgo() []GetAccountPositionsV5RespDataInnerCloseOrderAlgoInner`

GetCloseOrderAlgo returns the CloseOrderAlgo field if non-nil, zero value otherwise.

### GetCloseOrderAlgoOk

`func (o *GetAccountPositionsV5RespDataInner) GetCloseOrderAlgoOk() (*[]GetAccountPositionsV5RespDataInnerCloseOrderAlgoInner, bool)`

GetCloseOrderAlgoOk returns a tuple with the CloseOrderAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseOrderAlgo

`func (o *GetAccountPositionsV5RespDataInner) SetCloseOrderAlgo(v []GetAccountPositionsV5RespDataInnerCloseOrderAlgoInner)`

SetCloseOrderAlgo sets CloseOrderAlgo field to given value.

### HasCloseOrderAlgo

`func (o *GetAccountPositionsV5RespDataInner) HasCloseOrderAlgo() bool`

HasCloseOrderAlgo returns a boolean if a field has been set.

### GetDeltaBS

`func (o *GetAccountPositionsV5RespDataInner) GetDeltaBS() string`

GetDeltaBS returns the DeltaBS field if non-nil, zero value otherwise.

### GetDeltaBSOk

`func (o *GetAccountPositionsV5RespDataInner) GetDeltaBSOk() (*string, bool)`

GetDeltaBSOk returns a tuple with the DeltaBS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaBS

`func (o *GetAccountPositionsV5RespDataInner) SetDeltaBS(v string)`

SetDeltaBS sets DeltaBS field to given value.

### HasDeltaBS

`func (o *GetAccountPositionsV5RespDataInner) HasDeltaBS() bool`

HasDeltaBS returns a boolean if a field has been set.

### GetDeltaPA

`func (o *GetAccountPositionsV5RespDataInner) GetDeltaPA() string`

GetDeltaPA returns the DeltaPA field if non-nil, zero value otherwise.

### GetDeltaPAOk

`func (o *GetAccountPositionsV5RespDataInner) GetDeltaPAOk() (*string, bool)`

GetDeltaPAOk returns a tuple with the DeltaPA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaPA

`func (o *GetAccountPositionsV5RespDataInner) SetDeltaPA(v string)`

SetDeltaPA sets DeltaPA field to given value.

### HasDeltaPA

`func (o *GetAccountPositionsV5RespDataInner) HasDeltaPA() bool`

HasDeltaPA returns a boolean if a field has been set.

### GetFee

`func (o *GetAccountPositionsV5RespDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetAccountPositionsV5RespDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetAccountPositionsV5RespDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetAccountPositionsV5RespDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFundingFee

`func (o *GetAccountPositionsV5RespDataInner) GetFundingFee() string`

GetFundingFee returns the FundingFee field if non-nil, zero value otherwise.

### GetFundingFeeOk

`func (o *GetAccountPositionsV5RespDataInner) GetFundingFeeOk() (*string, bool)`

GetFundingFeeOk returns a tuple with the FundingFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingFee

`func (o *GetAccountPositionsV5RespDataInner) SetFundingFee(v string)`

SetFundingFee sets FundingFee field to given value.

### HasFundingFee

`func (o *GetAccountPositionsV5RespDataInner) HasFundingFee() bool`

HasFundingFee returns a boolean if a field has been set.

### GetGammaBS

`func (o *GetAccountPositionsV5RespDataInner) GetGammaBS() string`

GetGammaBS returns the GammaBS field if non-nil, zero value otherwise.

### GetGammaBSOk

`func (o *GetAccountPositionsV5RespDataInner) GetGammaBSOk() (*string, bool)`

GetGammaBSOk returns a tuple with the GammaBS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGammaBS

`func (o *GetAccountPositionsV5RespDataInner) SetGammaBS(v string)`

SetGammaBS sets GammaBS field to given value.

### HasGammaBS

`func (o *GetAccountPositionsV5RespDataInner) HasGammaBS() bool`

HasGammaBS returns a boolean if a field has been set.

### GetGammaPA

`func (o *GetAccountPositionsV5RespDataInner) GetGammaPA() string`

GetGammaPA returns the GammaPA field if non-nil, zero value otherwise.

### GetGammaPAOk

`func (o *GetAccountPositionsV5RespDataInner) GetGammaPAOk() (*string, bool)`

GetGammaPAOk returns a tuple with the GammaPA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGammaPA

`func (o *GetAccountPositionsV5RespDataInner) SetGammaPA(v string)`

SetGammaPA sets GammaPA field to given value.

### HasGammaPA

`func (o *GetAccountPositionsV5RespDataInner) HasGammaPA() bool`

HasGammaPA returns a boolean if a field has been set.

### GetIdxPx

`func (o *GetAccountPositionsV5RespDataInner) GetIdxPx() string`

GetIdxPx returns the IdxPx field if non-nil, zero value otherwise.

### GetIdxPxOk

`func (o *GetAccountPositionsV5RespDataInner) GetIdxPxOk() (*string, bool)`

GetIdxPxOk returns a tuple with the IdxPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdxPx

`func (o *GetAccountPositionsV5RespDataInner) SetIdxPx(v string)`

SetIdxPx sets IdxPx field to given value.

### HasIdxPx

`func (o *GetAccountPositionsV5RespDataInner) HasIdxPx() bool`

HasIdxPx returns a boolean if a field has been set.

### GetImr

`func (o *GetAccountPositionsV5RespDataInner) GetImr() string`

GetImr returns the Imr field if non-nil, zero value otherwise.

### GetImrOk

`func (o *GetAccountPositionsV5RespDataInner) GetImrOk() (*string, bool)`

GetImrOk returns a tuple with the Imr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImr

`func (o *GetAccountPositionsV5RespDataInner) SetImr(v string)`

SetImr sets Imr field to given value.

### HasImr

`func (o *GetAccountPositionsV5RespDataInner) HasImr() bool`

HasImr returns a boolean if a field has been set.

### GetInstId

`func (o *GetAccountPositionsV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAccountPositionsV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAccountPositionsV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAccountPositionsV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetAccountPositionsV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetAccountPositionsV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetAccountPositionsV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetAccountPositionsV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetInterest

`func (o *GetAccountPositionsV5RespDataInner) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *GetAccountPositionsV5RespDataInner) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *GetAccountPositionsV5RespDataInner) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *GetAccountPositionsV5RespDataInner) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetLast

`func (o *GetAccountPositionsV5RespDataInner) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetAccountPositionsV5RespDataInner) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetAccountPositionsV5RespDataInner) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *GetAccountPositionsV5RespDataInner) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLever

`func (o *GetAccountPositionsV5RespDataInner) GetLever() string`

GetLever returns the Lever field if non-nil, zero value otherwise.

### GetLeverOk

`func (o *GetAccountPositionsV5RespDataInner) GetLeverOk() (*string, bool)`

GetLeverOk returns a tuple with the Lever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLever

`func (o *GetAccountPositionsV5RespDataInner) SetLever(v string)`

SetLever sets Lever field to given value.

### HasLever

`func (o *GetAccountPositionsV5RespDataInner) HasLever() bool`

HasLever returns a boolean if a field has been set.

### GetLiab

`func (o *GetAccountPositionsV5RespDataInner) GetLiab() string`

GetLiab returns the Liab field if non-nil, zero value otherwise.

### GetLiabOk

`func (o *GetAccountPositionsV5RespDataInner) GetLiabOk() (*string, bool)`

GetLiabOk returns a tuple with the Liab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiab

`func (o *GetAccountPositionsV5RespDataInner) SetLiab(v string)`

SetLiab sets Liab field to given value.

### HasLiab

`func (o *GetAccountPositionsV5RespDataInner) HasLiab() bool`

HasLiab returns a boolean if a field has been set.

### GetLiabCcy

`func (o *GetAccountPositionsV5RespDataInner) GetLiabCcy() string`

GetLiabCcy returns the LiabCcy field if non-nil, zero value otherwise.

### GetLiabCcyOk

`func (o *GetAccountPositionsV5RespDataInner) GetLiabCcyOk() (*string, bool)`

GetLiabCcyOk returns a tuple with the LiabCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabCcy

`func (o *GetAccountPositionsV5RespDataInner) SetLiabCcy(v string)`

SetLiabCcy sets LiabCcy field to given value.

### HasLiabCcy

`func (o *GetAccountPositionsV5RespDataInner) HasLiabCcy() bool`

HasLiabCcy returns a boolean if a field has been set.

### GetLiqPenalty

`func (o *GetAccountPositionsV5RespDataInner) GetLiqPenalty() string`

GetLiqPenalty returns the LiqPenalty field if non-nil, zero value otherwise.

### GetLiqPenaltyOk

`func (o *GetAccountPositionsV5RespDataInner) GetLiqPenaltyOk() (*string, bool)`

GetLiqPenaltyOk returns a tuple with the LiqPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqPenalty

`func (o *GetAccountPositionsV5RespDataInner) SetLiqPenalty(v string)`

SetLiqPenalty sets LiqPenalty field to given value.

### HasLiqPenalty

`func (o *GetAccountPositionsV5RespDataInner) HasLiqPenalty() bool`

HasLiqPenalty returns a boolean if a field has been set.

### GetLiqPx

`func (o *GetAccountPositionsV5RespDataInner) GetLiqPx() string`

GetLiqPx returns the LiqPx field if non-nil, zero value otherwise.

### GetLiqPxOk

`func (o *GetAccountPositionsV5RespDataInner) GetLiqPxOk() (*string, bool)`

GetLiqPxOk returns a tuple with the LiqPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiqPx

`func (o *GetAccountPositionsV5RespDataInner) SetLiqPx(v string)`

SetLiqPx sets LiqPx field to given value.

### HasLiqPx

`func (o *GetAccountPositionsV5RespDataInner) HasLiqPx() bool`

HasLiqPx returns a boolean if a field has been set.

### GetMargin

`func (o *GetAccountPositionsV5RespDataInner) GetMargin() string`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *GetAccountPositionsV5RespDataInner) GetMarginOk() (*string, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *GetAccountPositionsV5RespDataInner) SetMargin(v string)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *GetAccountPositionsV5RespDataInner) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetMarkPx

`func (o *GetAccountPositionsV5RespDataInner) GetMarkPx() string`

GetMarkPx returns the MarkPx field if non-nil, zero value otherwise.

### GetMarkPxOk

`func (o *GetAccountPositionsV5RespDataInner) GetMarkPxOk() (*string, bool)`

GetMarkPxOk returns a tuple with the MarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPx

`func (o *GetAccountPositionsV5RespDataInner) SetMarkPx(v string)`

SetMarkPx sets MarkPx field to given value.

### HasMarkPx

`func (o *GetAccountPositionsV5RespDataInner) HasMarkPx() bool`

HasMarkPx returns a boolean if a field has been set.

### GetMaxSpotInUseAmt

`func (o *GetAccountPositionsV5RespDataInner) GetMaxSpotInUseAmt() string`

GetMaxSpotInUseAmt returns the MaxSpotInUseAmt field if non-nil, zero value otherwise.

### GetMaxSpotInUseAmtOk

`func (o *GetAccountPositionsV5RespDataInner) GetMaxSpotInUseAmtOk() (*string, bool)`

GetMaxSpotInUseAmtOk returns a tuple with the MaxSpotInUseAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpotInUseAmt

`func (o *GetAccountPositionsV5RespDataInner) SetMaxSpotInUseAmt(v string)`

SetMaxSpotInUseAmt sets MaxSpotInUseAmt field to given value.

### HasMaxSpotInUseAmt

`func (o *GetAccountPositionsV5RespDataInner) HasMaxSpotInUseAmt() bool`

HasMaxSpotInUseAmt returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetAccountPositionsV5RespDataInner) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetAccountPositionsV5RespDataInner) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetAccountPositionsV5RespDataInner) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetAccountPositionsV5RespDataInner) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetMgnRatio

`func (o *GetAccountPositionsV5RespDataInner) GetMgnRatio() string`

GetMgnRatio returns the MgnRatio field if non-nil, zero value otherwise.

### GetMgnRatioOk

`func (o *GetAccountPositionsV5RespDataInner) GetMgnRatioOk() (*string, bool)`

GetMgnRatioOk returns a tuple with the MgnRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnRatio

`func (o *GetAccountPositionsV5RespDataInner) SetMgnRatio(v string)`

SetMgnRatio sets MgnRatio field to given value.

### HasMgnRatio

`func (o *GetAccountPositionsV5RespDataInner) HasMgnRatio() bool`

HasMgnRatio returns a boolean if a field has been set.

### GetMmr

`func (o *GetAccountPositionsV5RespDataInner) GetMmr() string`

GetMmr returns the Mmr field if non-nil, zero value otherwise.

### GetMmrOk

`func (o *GetAccountPositionsV5RespDataInner) GetMmrOk() (*string, bool)`

GetMmrOk returns a tuple with the Mmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmr

`func (o *GetAccountPositionsV5RespDataInner) SetMmr(v string)`

SetMmr sets Mmr field to given value.

### HasMmr

`func (o *GetAccountPositionsV5RespDataInner) HasMmr() bool`

HasMmr returns a boolean if a field has been set.

### GetNonSettleAvgPx

`func (o *GetAccountPositionsV5RespDataInner) GetNonSettleAvgPx() string`

GetNonSettleAvgPx returns the NonSettleAvgPx field if non-nil, zero value otherwise.

### GetNonSettleAvgPxOk

`func (o *GetAccountPositionsV5RespDataInner) GetNonSettleAvgPxOk() (*string, bool)`

GetNonSettleAvgPxOk returns a tuple with the NonSettleAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonSettleAvgPx

`func (o *GetAccountPositionsV5RespDataInner) SetNonSettleAvgPx(v string)`

SetNonSettleAvgPx sets NonSettleAvgPx field to given value.

### HasNonSettleAvgPx

`func (o *GetAccountPositionsV5RespDataInner) HasNonSettleAvgPx() bool`

HasNonSettleAvgPx returns a boolean if a field has been set.

### GetNotionalUsd

`func (o *GetAccountPositionsV5RespDataInner) GetNotionalUsd() string`

GetNotionalUsd returns the NotionalUsd field if non-nil, zero value otherwise.

### GetNotionalUsdOk

`func (o *GetAccountPositionsV5RespDataInner) GetNotionalUsdOk() (*string, bool)`

GetNotionalUsdOk returns a tuple with the NotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsd

`func (o *GetAccountPositionsV5RespDataInner) SetNotionalUsd(v string)`

SetNotionalUsd sets NotionalUsd field to given value.

### HasNotionalUsd

`func (o *GetAccountPositionsV5RespDataInner) HasNotionalUsd() bool`

HasNotionalUsd returns a boolean if a field has been set.

### GetOptVal

`func (o *GetAccountPositionsV5RespDataInner) GetOptVal() string`

GetOptVal returns the OptVal field if non-nil, zero value otherwise.

### GetOptValOk

`func (o *GetAccountPositionsV5RespDataInner) GetOptValOk() (*string, bool)`

GetOptValOk returns a tuple with the OptVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptVal

`func (o *GetAccountPositionsV5RespDataInner) SetOptVal(v string)`

SetOptVal sets OptVal field to given value.

### HasOptVal

`func (o *GetAccountPositionsV5RespDataInner) HasOptVal() bool`

HasOptVal returns a boolean if a field has been set.

### GetPendingCloseOrdLiabVal

`func (o *GetAccountPositionsV5RespDataInner) GetPendingCloseOrdLiabVal() string`

GetPendingCloseOrdLiabVal returns the PendingCloseOrdLiabVal field if non-nil, zero value otherwise.

### GetPendingCloseOrdLiabValOk

`func (o *GetAccountPositionsV5RespDataInner) GetPendingCloseOrdLiabValOk() (*string, bool)`

GetPendingCloseOrdLiabValOk returns a tuple with the PendingCloseOrdLiabVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCloseOrdLiabVal

`func (o *GetAccountPositionsV5RespDataInner) SetPendingCloseOrdLiabVal(v string)`

SetPendingCloseOrdLiabVal sets PendingCloseOrdLiabVal field to given value.

### HasPendingCloseOrdLiabVal

`func (o *GetAccountPositionsV5RespDataInner) HasPendingCloseOrdLiabVal() bool`

HasPendingCloseOrdLiabVal returns a boolean if a field has been set.

### GetPnl

`func (o *GetAccountPositionsV5RespDataInner) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *GetAccountPositionsV5RespDataInner) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *GetAccountPositionsV5RespDataInner) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *GetAccountPositionsV5RespDataInner) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetPos

`func (o *GetAccountPositionsV5RespDataInner) GetPos() string`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *GetAccountPositionsV5RespDataInner) GetPosOk() (*string, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *GetAccountPositionsV5RespDataInner) SetPos(v string)`

SetPos sets Pos field to given value.

### HasPos

`func (o *GetAccountPositionsV5RespDataInner) HasPos() bool`

HasPos returns a boolean if a field has been set.

### GetPosCcy

`func (o *GetAccountPositionsV5RespDataInner) GetPosCcy() string`

GetPosCcy returns the PosCcy field if non-nil, zero value otherwise.

### GetPosCcyOk

`func (o *GetAccountPositionsV5RespDataInner) GetPosCcyOk() (*string, bool)`

GetPosCcyOk returns a tuple with the PosCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosCcy

`func (o *GetAccountPositionsV5RespDataInner) SetPosCcy(v string)`

SetPosCcy sets PosCcy field to given value.

### HasPosCcy

`func (o *GetAccountPositionsV5RespDataInner) HasPosCcy() bool`

HasPosCcy returns a boolean if a field has been set.

### GetPosId

`func (o *GetAccountPositionsV5RespDataInner) GetPosId() string`

GetPosId returns the PosId field if non-nil, zero value otherwise.

### GetPosIdOk

`func (o *GetAccountPositionsV5RespDataInner) GetPosIdOk() (*string, bool)`

GetPosIdOk returns a tuple with the PosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosId

`func (o *GetAccountPositionsV5RespDataInner) SetPosId(v string)`

SetPosId sets PosId field to given value.

### HasPosId

`func (o *GetAccountPositionsV5RespDataInner) HasPosId() bool`

HasPosId returns a boolean if a field has been set.

### GetPosSide

`func (o *GetAccountPositionsV5RespDataInner) GetPosSide() string`

GetPosSide returns the PosSide field if non-nil, zero value otherwise.

### GetPosSideOk

`func (o *GetAccountPositionsV5RespDataInner) GetPosSideOk() (*string, bool)`

GetPosSideOk returns a tuple with the PosSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosSide

`func (o *GetAccountPositionsV5RespDataInner) SetPosSide(v string)`

SetPosSide sets PosSide field to given value.

### HasPosSide

`func (o *GetAccountPositionsV5RespDataInner) HasPosSide() bool`

HasPosSide returns a boolean if a field has been set.

### GetQuoteBal

`func (o *GetAccountPositionsV5RespDataInner) GetQuoteBal() string`

GetQuoteBal returns the QuoteBal field if non-nil, zero value otherwise.

### GetQuoteBalOk

`func (o *GetAccountPositionsV5RespDataInner) GetQuoteBalOk() (*string, bool)`

GetQuoteBalOk returns a tuple with the QuoteBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteBal

`func (o *GetAccountPositionsV5RespDataInner) SetQuoteBal(v string)`

SetQuoteBal sets QuoteBal field to given value.

### HasQuoteBal

`func (o *GetAccountPositionsV5RespDataInner) HasQuoteBal() bool`

HasQuoteBal returns a boolean if a field has been set.

### GetQuoteBorrowed

`func (o *GetAccountPositionsV5RespDataInner) GetQuoteBorrowed() string`

GetQuoteBorrowed returns the QuoteBorrowed field if non-nil, zero value otherwise.

### GetQuoteBorrowedOk

`func (o *GetAccountPositionsV5RespDataInner) GetQuoteBorrowedOk() (*string, bool)`

GetQuoteBorrowedOk returns a tuple with the QuoteBorrowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteBorrowed

`func (o *GetAccountPositionsV5RespDataInner) SetQuoteBorrowed(v string)`

SetQuoteBorrowed sets QuoteBorrowed field to given value.

### HasQuoteBorrowed

`func (o *GetAccountPositionsV5RespDataInner) HasQuoteBorrowed() bool`

HasQuoteBorrowed returns a boolean if a field has been set.

### GetQuoteInterest

`func (o *GetAccountPositionsV5RespDataInner) GetQuoteInterest() string`

GetQuoteInterest returns the QuoteInterest field if non-nil, zero value otherwise.

### GetQuoteInterestOk

`func (o *GetAccountPositionsV5RespDataInner) GetQuoteInterestOk() (*string, bool)`

GetQuoteInterestOk returns a tuple with the QuoteInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteInterest

`func (o *GetAccountPositionsV5RespDataInner) SetQuoteInterest(v string)`

SetQuoteInterest sets QuoteInterest field to given value.

### HasQuoteInterest

`func (o *GetAccountPositionsV5RespDataInner) HasQuoteInterest() bool`

HasQuoteInterest returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *GetAccountPositionsV5RespDataInner) GetRealizedPnl() string`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *GetAccountPositionsV5RespDataInner) GetRealizedPnlOk() (*string, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *GetAccountPositionsV5RespDataInner) SetRealizedPnl(v string)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *GetAccountPositionsV5RespDataInner) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetSettledPnl

`func (o *GetAccountPositionsV5RespDataInner) GetSettledPnl() string`

GetSettledPnl returns the SettledPnl field if non-nil, zero value otherwise.

### GetSettledPnlOk

`func (o *GetAccountPositionsV5RespDataInner) GetSettledPnlOk() (*string, bool)`

GetSettledPnlOk returns a tuple with the SettledPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledPnl

`func (o *GetAccountPositionsV5RespDataInner) SetSettledPnl(v string)`

SetSettledPnl sets SettledPnl field to given value.

### HasSettledPnl

`func (o *GetAccountPositionsV5RespDataInner) HasSettledPnl() bool`

HasSettledPnl returns a boolean if a field has been set.

### GetSpotInUseAmt

`func (o *GetAccountPositionsV5RespDataInner) GetSpotInUseAmt() string`

GetSpotInUseAmt returns the SpotInUseAmt field if non-nil, zero value otherwise.

### GetSpotInUseAmtOk

`func (o *GetAccountPositionsV5RespDataInner) GetSpotInUseAmtOk() (*string, bool)`

GetSpotInUseAmtOk returns a tuple with the SpotInUseAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInUseAmt

`func (o *GetAccountPositionsV5RespDataInner) SetSpotInUseAmt(v string)`

SetSpotInUseAmt sets SpotInUseAmt field to given value.

### HasSpotInUseAmt

`func (o *GetAccountPositionsV5RespDataInner) HasSpotInUseAmt() bool`

HasSpotInUseAmt returns a boolean if a field has been set.

### GetSpotInUseCcy

`func (o *GetAccountPositionsV5RespDataInner) GetSpotInUseCcy() string`

GetSpotInUseCcy returns the SpotInUseCcy field if non-nil, zero value otherwise.

### GetSpotInUseCcyOk

`func (o *GetAccountPositionsV5RespDataInner) GetSpotInUseCcyOk() (*string, bool)`

GetSpotInUseCcyOk returns a tuple with the SpotInUseCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInUseCcy

`func (o *GetAccountPositionsV5RespDataInner) SetSpotInUseCcy(v string)`

SetSpotInUseCcy sets SpotInUseCcy field to given value.

### HasSpotInUseCcy

`func (o *GetAccountPositionsV5RespDataInner) HasSpotInUseCcy() bool`

HasSpotInUseCcy returns a boolean if a field has been set.

### GetThetaBS

`func (o *GetAccountPositionsV5RespDataInner) GetThetaBS() string`

GetThetaBS returns the ThetaBS field if non-nil, zero value otherwise.

### GetThetaBSOk

`func (o *GetAccountPositionsV5RespDataInner) GetThetaBSOk() (*string, bool)`

GetThetaBSOk returns a tuple with the ThetaBS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThetaBS

`func (o *GetAccountPositionsV5RespDataInner) SetThetaBS(v string)`

SetThetaBS sets ThetaBS field to given value.

### HasThetaBS

`func (o *GetAccountPositionsV5RespDataInner) HasThetaBS() bool`

HasThetaBS returns a boolean if a field has been set.

### GetThetaPA

`func (o *GetAccountPositionsV5RespDataInner) GetThetaPA() string`

GetThetaPA returns the ThetaPA field if non-nil, zero value otherwise.

### GetThetaPAOk

`func (o *GetAccountPositionsV5RespDataInner) GetThetaPAOk() (*string, bool)`

GetThetaPAOk returns a tuple with the ThetaPA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThetaPA

`func (o *GetAccountPositionsV5RespDataInner) SetThetaPA(v string)`

SetThetaPA sets ThetaPA field to given value.

### HasThetaPA

`func (o *GetAccountPositionsV5RespDataInner) HasThetaPA() bool`

HasThetaPA returns a boolean if a field has been set.

### GetTradeId

`func (o *GetAccountPositionsV5RespDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetAccountPositionsV5RespDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetAccountPositionsV5RespDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetAccountPositionsV5RespDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetUTime

`func (o *GetAccountPositionsV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetAccountPositionsV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetAccountPositionsV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetAccountPositionsV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUpl

`func (o *GetAccountPositionsV5RespDataInner) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *GetAccountPositionsV5RespDataInner) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *GetAccountPositionsV5RespDataInner) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *GetAccountPositionsV5RespDataInner) HasUpl() bool`

HasUpl returns a boolean if a field has been set.

### GetUplLastPx

`func (o *GetAccountPositionsV5RespDataInner) GetUplLastPx() string`

GetUplLastPx returns the UplLastPx field if non-nil, zero value otherwise.

### GetUplLastPxOk

`func (o *GetAccountPositionsV5RespDataInner) GetUplLastPxOk() (*string, bool)`

GetUplLastPxOk returns a tuple with the UplLastPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplLastPx

`func (o *GetAccountPositionsV5RespDataInner) SetUplLastPx(v string)`

SetUplLastPx sets UplLastPx field to given value.

### HasUplLastPx

`func (o *GetAccountPositionsV5RespDataInner) HasUplLastPx() bool`

HasUplLastPx returns a boolean if a field has been set.

### GetUplRatio

`func (o *GetAccountPositionsV5RespDataInner) GetUplRatio() string`

GetUplRatio returns the UplRatio field if non-nil, zero value otherwise.

### GetUplRatioOk

`func (o *GetAccountPositionsV5RespDataInner) GetUplRatioOk() (*string, bool)`

GetUplRatioOk returns a tuple with the UplRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplRatio

`func (o *GetAccountPositionsV5RespDataInner) SetUplRatio(v string)`

SetUplRatio sets UplRatio field to given value.

### HasUplRatio

`func (o *GetAccountPositionsV5RespDataInner) HasUplRatio() bool`

HasUplRatio returns a boolean if a field has been set.

### GetUplRatioLastPx

`func (o *GetAccountPositionsV5RespDataInner) GetUplRatioLastPx() string`

GetUplRatioLastPx returns the UplRatioLastPx field if non-nil, zero value otherwise.

### GetUplRatioLastPxOk

`func (o *GetAccountPositionsV5RespDataInner) GetUplRatioLastPxOk() (*string, bool)`

GetUplRatioLastPxOk returns a tuple with the UplRatioLastPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplRatioLastPx

`func (o *GetAccountPositionsV5RespDataInner) SetUplRatioLastPx(v string)`

SetUplRatioLastPx sets UplRatioLastPx field to given value.

### HasUplRatioLastPx

`func (o *GetAccountPositionsV5RespDataInner) HasUplRatioLastPx() bool`

HasUplRatioLastPx returns a boolean if a field has been set.

### GetUsdPx

`func (o *GetAccountPositionsV5RespDataInner) GetUsdPx() string`

GetUsdPx returns the UsdPx field if non-nil, zero value otherwise.

### GetUsdPxOk

`func (o *GetAccountPositionsV5RespDataInner) GetUsdPxOk() (*string, bool)`

GetUsdPxOk returns a tuple with the UsdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsdPx

`func (o *GetAccountPositionsV5RespDataInner) SetUsdPx(v string)`

SetUsdPx sets UsdPx field to given value.

### HasUsdPx

`func (o *GetAccountPositionsV5RespDataInner) HasUsdPx() bool`

HasUsdPx returns a boolean if a field has been set.

### GetVegaBS

`func (o *GetAccountPositionsV5RespDataInner) GetVegaBS() string`

GetVegaBS returns the VegaBS field if non-nil, zero value otherwise.

### GetVegaBSOk

`func (o *GetAccountPositionsV5RespDataInner) GetVegaBSOk() (*string, bool)`

GetVegaBSOk returns a tuple with the VegaBS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVegaBS

`func (o *GetAccountPositionsV5RespDataInner) SetVegaBS(v string)`

SetVegaBS sets VegaBS field to given value.

### HasVegaBS

`func (o *GetAccountPositionsV5RespDataInner) HasVegaBS() bool`

HasVegaBS returns a boolean if a field has been set.

### GetVegaPA

`func (o *GetAccountPositionsV5RespDataInner) GetVegaPA() string`

GetVegaPA returns the VegaPA field if non-nil, zero value otherwise.

### GetVegaPAOk

`func (o *GetAccountPositionsV5RespDataInner) GetVegaPAOk() (*string, bool)`

GetVegaPAOk returns a tuple with the VegaPA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVegaPA

`func (o *GetAccountPositionsV5RespDataInner) SetVegaPA(v string)`

SetVegaPA sets VegaPA field to given value.

### HasVegaPA

`func (o *GetAccountPositionsV5RespDataInner) HasVegaPA() bool`

HasVegaPA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


