# GetAccountBillsArchiveV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bal** | Pointer to **string** | Balance at the account level | [optional] [default to ""]
**BalChg** | Pointer to **string** | Change in balance amount at the account level | [optional] [default to ""]
**BillId** | Pointer to **string** | Bill ID | [optional] [default to ""]
**Ccy** | Pointer to **string** | Account balance currency | [optional] [default to ""]
**ClOrdId** | Pointer to **string** | Client Order ID as assigned by the client   A combination of case-sensitive alphanumerics, all numbers, or all letters of up to 32 characters. | [optional] [default to ""]
**ExecType** | Pointer to **string** | Liquidity taker or maker  &#x60;T&#x60;: taker &#x60;M&#x60;: maker | [optional] [default to ""]
**Fee** | Pointer to **string** | Fee  Negative number represents the user transaction fee charged by the platform.   Positive number represents rebate.   | [optional] [default to ""]
**FillFwdPx** | Pointer to **string** | Forward price when filled   Only applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**FillIdxPx** | Pointer to **string** | Index price at the moment of trade execution   For cross currency spot pairs, it returns baseCcy-USDT index price. For example, for LTC-ETH, this field returns the index price of LTC-USDT. | [optional] [default to ""]
**FillMarkPx** | Pointer to **string** | Mark price when filled   Applicable to FUTURES/SWAP/OPTIONS, return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**FillMarkVol** | Pointer to **string** | Mark volatility when filled   Only applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**FillPxUsd** | Pointer to **string** | Options price when filled, in the unit of USD   Only applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**FillPxVol** | Pointer to **string** | Implied volatility when filled   Only applicable to options; return \&quot;\&quot; for other instrument types | [optional] [default to ""]
**FillTime** | Pointer to **string** | Last filled time | [optional] [default to ""]
**From** | Pointer to **string** | The remitting account  &#x60;6&#x60;: Funding account  &#x60;18&#x60;: Trading account  Only applicable to &#x60;transfer&#x60;. When bill type is not &#x60;transfer&#x60;, the field returns \&quot;\&quot;. | [optional] [default to ""]
**InstId** | Pointer to **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [optional] [default to ""]
**InstType** | Pointer to **string** | Instrument type | [optional] [default to ""]
**Interest** | Pointer to **string** | Interest | [optional] [default to ""]
**MgnMode** | Pointer to **string** | Margin mode  &#x60;isolated&#x60; &#x60;cross&#x60; &#x60;cash&#x60;  When bills are not generated by trading, the field returns \&quot;\&quot; | [optional] [default to ""]
**Notes** | Pointer to **string** | Notes | [optional] [default to ""]
**OrdId** | Pointer to **string** | Order ID  Return order ID when the type is &#x60;2&#x60;/&#x60;5&#x60;/&#x60;9&#x60;  Return \&quot;\&quot; when there is no order. | [optional] [default to ""]
**Pnl** | Pointer to **string** | Profit and loss | [optional] [default to ""]
**PosBal** | Pointer to **string** | Balance at the position level | [optional] [default to ""]
**PosBalChg** | Pointer to **string** | Change in balance amount at the position level | [optional] [default to ""]
**Px** | Pointer to **string** | Price which related to subType  &#x60;1&#x60;: Buy &#x60;2&#x60;: Sell &#x60;3&#x60;: Open long &#x60;4&#x60;: Open short &#x60;5&#x60;: Close long &#x60;6&#x60;: Close short &#x60;204&#x60;: block trade buy &#x60;205&#x60;: block trade sell &#x60;206&#x60;: block trade open long &#x60;207&#x60;: block trade open short  &#x60;208&#x60;: block trade close long &#x60;209&#x60;: block trade close short &#x60;114&#x60;: Forced repayment buy &#x60;115&#x60;: Forced repayment sell  &#x60;100&#x60;: Partial liquidation close long &#x60;101&#x60;: Partial liquidation close short &#x60;102&#x60;: Partial liquidation buy &#x60;103&#x60;: Partial liquidation sell &#x60;104&#x60;: Liquidation long &#x60;105&#x60;: Liquidation short &#x60;106&#x60;: Liquidation buy &#x60;107&#x60;: Liquidation sell &#x60;16&#x60;: Repay forcibly &#x60;17&#x60;: Repay interest by borrowing forcibly &#x60;110&#x60;: Liquidation transfer in &#x60;111&#x60;: Liquidation transfer out  &#x60;112&#x60;: Delivery long &#x60;113&#x60;: Delivery short  &#x60;170&#x60;: Exercised &#x60;171&#x60;: Counterparty exercised &#x60;172&#x60;: Expired OTM  &#x60;173&#x60;: Funding fee expense &#x60;174&#x60;: Funding fee income | [optional] [default to ""]
**SubType** | Pointer to **string** | Bill subtype | [optional] [default to ""]
**Sz** | Pointer to **string** | Quantity  For &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;, it is fill quantity or position quantity, the unit is contract. The value is always positive.  For other scenarios. the unit is account balance currency(&#x60;ccy&#x60;). | [optional] [default to ""]
**Tag** | Pointer to **string** | Order tag | [optional] [default to ""]
**To** | Pointer to **string** | The beneficiary account  &#x60;6&#x60;: Funding account  &#x60;18&#x60;: Trading account  Only applicable to &#x60;transfer&#x60;. When bill type is not &#x60;transfer&#x60;, the field returns \&quot;\&quot;. | [optional] [default to ""]
**TradeId** | Pointer to **string** | Last traded ID | [optional] [default to ""]
**Ts** | Pointer to **string** | The time when the balance complete update, Unix timestamp format in milliseconds, e.g.&#x60;1597026383085&#x60; | [optional] [default to ""]
**Type** | Pointer to **string** | Bill type | [optional] [default to ""]

## Methods

### NewGetAccountBillsArchiveV5RespDataInner

`func NewGetAccountBillsArchiveV5RespDataInner() *GetAccountBillsArchiveV5RespDataInner`

NewGetAccountBillsArchiveV5RespDataInner instantiates a new GetAccountBillsArchiveV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountBillsArchiveV5RespDataInnerWithDefaults

`func NewGetAccountBillsArchiveV5RespDataInnerWithDefaults() *GetAccountBillsArchiveV5RespDataInner`

NewGetAccountBillsArchiveV5RespDataInnerWithDefaults instantiates a new GetAccountBillsArchiveV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBal

`func (o *GetAccountBillsArchiveV5RespDataInner) GetBal() string`

GetBal returns the Bal field if non-nil, zero value otherwise.

### GetBalOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetBalOk() (*string, bool)`

GetBalOk returns a tuple with the Bal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBal

`func (o *GetAccountBillsArchiveV5RespDataInner) SetBal(v string)`

SetBal sets Bal field to given value.

### HasBal

`func (o *GetAccountBillsArchiveV5RespDataInner) HasBal() bool`

HasBal returns a boolean if a field has been set.

### GetBalChg

`func (o *GetAccountBillsArchiveV5RespDataInner) GetBalChg() string`

GetBalChg returns the BalChg field if non-nil, zero value otherwise.

### GetBalChgOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetBalChgOk() (*string, bool)`

GetBalChgOk returns a tuple with the BalChg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalChg

`func (o *GetAccountBillsArchiveV5RespDataInner) SetBalChg(v string)`

SetBalChg sets BalChg field to given value.

### HasBalChg

`func (o *GetAccountBillsArchiveV5RespDataInner) HasBalChg() bool`

HasBalChg returns a boolean if a field has been set.

### GetBillId

`func (o *GetAccountBillsArchiveV5RespDataInner) GetBillId() string`

GetBillId returns the BillId field if non-nil, zero value otherwise.

### GetBillIdOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetBillIdOk() (*string, bool)`

GetBillIdOk returns a tuple with the BillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillId

`func (o *GetAccountBillsArchiveV5RespDataInner) SetBillId(v string)`

SetBillId sets BillId field to given value.

### HasBillId

`func (o *GetAccountBillsArchiveV5RespDataInner) HasBillId() bool`

HasBillId returns a boolean if a field has been set.

### GetCcy

`func (o *GetAccountBillsArchiveV5RespDataInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountBillsArchiveV5RespDataInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountBillsArchiveV5RespDataInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClOrdId

`func (o *GetAccountBillsArchiveV5RespDataInner) GetClOrdId() string`

GetClOrdId returns the ClOrdId field if non-nil, zero value otherwise.

### GetClOrdIdOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetClOrdIdOk() (*string, bool)`

GetClOrdIdOk returns a tuple with the ClOrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClOrdId

`func (o *GetAccountBillsArchiveV5RespDataInner) SetClOrdId(v string)`

SetClOrdId sets ClOrdId field to given value.

### HasClOrdId

`func (o *GetAccountBillsArchiveV5RespDataInner) HasClOrdId() bool`

HasClOrdId returns a boolean if a field has been set.

### GetExecType

`func (o *GetAccountBillsArchiveV5RespDataInner) GetExecType() string`

GetExecType returns the ExecType field if non-nil, zero value otherwise.

### GetExecTypeOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetExecTypeOk() (*string, bool)`

GetExecTypeOk returns a tuple with the ExecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecType

`func (o *GetAccountBillsArchiveV5RespDataInner) SetExecType(v string)`

SetExecType sets ExecType field to given value.

### HasExecType

`func (o *GetAccountBillsArchiveV5RespDataInner) HasExecType() bool`

HasExecType returns a boolean if a field has been set.

### GetFee

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetAccountBillsArchiveV5RespDataInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetAccountBillsArchiveV5RespDataInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFillFwdPx

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillFwdPx() string`

GetFillFwdPx returns the FillFwdPx field if non-nil, zero value otherwise.

### GetFillFwdPxOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillFwdPxOk() (*string, bool)`

GetFillFwdPxOk returns a tuple with the FillFwdPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillFwdPx

`func (o *GetAccountBillsArchiveV5RespDataInner) SetFillFwdPx(v string)`

SetFillFwdPx sets FillFwdPx field to given value.

### HasFillFwdPx

`func (o *GetAccountBillsArchiveV5RespDataInner) HasFillFwdPx() bool`

HasFillFwdPx returns a boolean if a field has been set.

### GetFillIdxPx

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillIdxPx() string`

GetFillIdxPx returns the FillIdxPx field if non-nil, zero value otherwise.

### GetFillIdxPxOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillIdxPxOk() (*string, bool)`

GetFillIdxPxOk returns a tuple with the FillIdxPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillIdxPx

`func (o *GetAccountBillsArchiveV5RespDataInner) SetFillIdxPx(v string)`

SetFillIdxPx sets FillIdxPx field to given value.

### HasFillIdxPx

`func (o *GetAccountBillsArchiveV5RespDataInner) HasFillIdxPx() bool`

HasFillIdxPx returns a boolean if a field has been set.

### GetFillMarkPx

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillMarkPx() string`

GetFillMarkPx returns the FillMarkPx field if non-nil, zero value otherwise.

### GetFillMarkPxOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillMarkPxOk() (*string, bool)`

GetFillMarkPxOk returns a tuple with the FillMarkPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillMarkPx

`func (o *GetAccountBillsArchiveV5RespDataInner) SetFillMarkPx(v string)`

SetFillMarkPx sets FillMarkPx field to given value.

### HasFillMarkPx

`func (o *GetAccountBillsArchiveV5RespDataInner) HasFillMarkPx() bool`

HasFillMarkPx returns a boolean if a field has been set.

### GetFillMarkVol

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillMarkVol() string`

GetFillMarkVol returns the FillMarkVol field if non-nil, zero value otherwise.

### GetFillMarkVolOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillMarkVolOk() (*string, bool)`

GetFillMarkVolOk returns a tuple with the FillMarkVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillMarkVol

`func (o *GetAccountBillsArchiveV5RespDataInner) SetFillMarkVol(v string)`

SetFillMarkVol sets FillMarkVol field to given value.

### HasFillMarkVol

`func (o *GetAccountBillsArchiveV5RespDataInner) HasFillMarkVol() bool`

HasFillMarkVol returns a boolean if a field has been set.

### GetFillPxUsd

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillPxUsd() string`

GetFillPxUsd returns the FillPxUsd field if non-nil, zero value otherwise.

### GetFillPxUsdOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillPxUsdOk() (*string, bool)`

GetFillPxUsdOk returns a tuple with the FillPxUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPxUsd

`func (o *GetAccountBillsArchiveV5RespDataInner) SetFillPxUsd(v string)`

SetFillPxUsd sets FillPxUsd field to given value.

### HasFillPxUsd

`func (o *GetAccountBillsArchiveV5RespDataInner) HasFillPxUsd() bool`

HasFillPxUsd returns a boolean if a field has been set.

### GetFillPxVol

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillPxVol() string`

GetFillPxVol returns the FillPxVol field if non-nil, zero value otherwise.

### GetFillPxVolOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillPxVolOk() (*string, bool)`

GetFillPxVolOk returns a tuple with the FillPxVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPxVol

`func (o *GetAccountBillsArchiveV5RespDataInner) SetFillPxVol(v string)`

SetFillPxVol sets FillPxVol field to given value.

### HasFillPxVol

`func (o *GetAccountBillsArchiveV5RespDataInner) HasFillPxVol() bool`

HasFillPxVol returns a boolean if a field has been set.

### GetFillTime

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillTime() string`

GetFillTime returns the FillTime field if non-nil, zero value otherwise.

### GetFillTimeOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFillTimeOk() (*string, bool)`

GetFillTimeOk returns a tuple with the FillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillTime

`func (o *GetAccountBillsArchiveV5RespDataInner) SetFillTime(v string)`

SetFillTime sets FillTime field to given value.

### HasFillTime

`func (o *GetAccountBillsArchiveV5RespDataInner) HasFillTime() bool`

HasFillTime returns a boolean if a field has been set.

### GetFrom

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetAccountBillsArchiveV5RespDataInner) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetAccountBillsArchiveV5RespDataInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetInstId

`func (o *GetAccountBillsArchiveV5RespDataInner) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *GetAccountBillsArchiveV5RespDataInner) SetInstId(v string)`

SetInstId sets InstId field to given value.

### HasInstId

`func (o *GetAccountBillsArchiveV5RespDataInner) HasInstId() bool`

HasInstId returns a boolean if a field has been set.

### GetInstType

`func (o *GetAccountBillsArchiveV5RespDataInner) GetInstType() string`

GetInstType returns the InstType field if non-nil, zero value otherwise.

### GetInstTypeOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetInstTypeOk() (*string, bool)`

GetInstTypeOk returns a tuple with the InstType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstType

`func (o *GetAccountBillsArchiveV5RespDataInner) SetInstType(v string)`

SetInstType sets InstType field to given value.

### HasInstType

`func (o *GetAccountBillsArchiveV5RespDataInner) HasInstType() bool`

HasInstType returns a boolean if a field has been set.

### GetInterest

`func (o *GetAccountBillsArchiveV5RespDataInner) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *GetAccountBillsArchiveV5RespDataInner) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *GetAccountBillsArchiveV5RespDataInner) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetMgnMode

`func (o *GetAccountBillsArchiveV5RespDataInner) GetMgnMode() string`

GetMgnMode returns the MgnMode field if non-nil, zero value otherwise.

### GetMgnModeOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetMgnModeOk() (*string, bool)`

GetMgnModeOk returns a tuple with the MgnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnMode

`func (o *GetAccountBillsArchiveV5RespDataInner) SetMgnMode(v string)`

SetMgnMode sets MgnMode field to given value.

### HasMgnMode

`func (o *GetAccountBillsArchiveV5RespDataInner) HasMgnMode() bool`

HasMgnMode returns a boolean if a field has been set.

### GetNotes

`func (o *GetAccountBillsArchiveV5RespDataInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetAccountBillsArchiveV5RespDataInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetAccountBillsArchiveV5RespDataInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOrdId

`func (o *GetAccountBillsArchiveV5RespDataInner) GetOrdId() string`

GetOrdId returns the OrdId field if non-nil, zero value otherwise.

### GetOrdIdOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetOrdIdOk() (*string, bool)`

GetOrdIdOk returns a tuple with the OrdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdId

`func (o *GetAccountBillsArchiveV5RespDataInner) SetOrdId(v string)`

SetOrdId sets OrdId field to given value.

### HasOrdId

`func (o *GetAccountBillsArchiveV5RespDataInner) HasOrdId() bool`

HasOrdId returns a boolean if a field has been set.

### GetPnl

`func (o *GetAccountBillsArchiveV5RespDataInner) GetPnl() string`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetPnlOk() (*string, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *GetAccountBillsArchiveV5RespDataInner) SetPnl(v string)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *GetAccountBillsArchiveV5RespDataInner) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### GetPosBal

`func (o *GetAccountBillsArchiveV5RespDataInner) GetPosBal() string`

GetPosBal returns the PosBal field if non-nil, zero value otherwise.

### GetPosBalOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetPosBalOk() (*string, bool)`

GetPosBalOk returns a tuple with the PosBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosBal

`func (o *GetAccountBillsArchiveV5RespDataInner) SetPosBal(v string)`

SetPosBal sets PosBal field to given value.

### HasPosBal

`func (o *GetAccountBillsArchiveV5RespDataInner) HasPosBal() bool`

HasPosBal returns a boolean if a field has been set.

### GetPosBalChg

`func (o *GetAccountBillsArchiveV5RespDataInner) GetPosBalChg() string`

GetPosBalChg returns the PosBalChg field if non-nil, zero value otherwise.

### GetPosBalChgOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetPosBalChgOk() (*string, bool)`

GetPosBalChgOk returns a tuple with the PosBalChg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosBalChg

`func (o *GetAccountBillsArchiveV5RespDataInner) SetPosBalChg(v string)`

SetPosBalChg sets PosBalChg field to given value.

### HasPosBalChg

`func (o *GetAccountBillsArchiveV5RespDataInner) HasPosBalChg() bool`

HasPosBalChg returns a boolean if a field has been set.

### GetPx

`func (o *GetAccountBillsArchiveV5RespDataInner) GetPx() string`

GetPx returns the Px field if non-nil, zero value otherwise.

### GetPxOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetPxOk() (*string, bool)`

GetPxOk returns a tuple with the Px field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPx

`func (o *GetAccountBillsArchiveV5RespDataInner) SetPx(v string)`

SetPx sets Px field to given value.

### HasPx

`func (o *GetAccountBillsArchiveV5RespDataInner) HasPx() bool`

HasPx returns a boolean if a field has been set.

### GetSubType

`func (o *GetAccountBillsArchiveV5RespDataInner) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetAccountBillsArchiveV5RespDataInner) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *GetAccountBillsArchiveV5RespDataInner) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetSz

`func (o *GetAccountBillsArchiveV5RespDataInner) GetSz() string`

GetSz returns the Sz field if non-nil, zero value otherwise.

### GetSzOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetSzOk() (*string, bool)`

GetSzOk returns a tuple with the Sz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSz

`func (o *GetAccountBillsArchiveV5RespDataInner) SetSz(v string)`

SetSz sets Sz field to given value.

### HasSz

`func (o *GetAccountBillsArchiveV5RespDataInner) HasSz() bool`

HasSz returns a boolean if a field has been set.

### GetTag

`func (o *GetAccountBillsArchiveV5RespDataInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetAccountBillsArchiveV5RespDataInner) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetAccountBillsArchiveV5RespDataInner) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTo

`func (o *GetAccountBillsArchiveV5RespDataInner) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetAccountBillsArchiveV5RespDataInner) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GetAccountBillsArchiveV5RespDataInner) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTradeId

`func (o *GetAccountBillsArchiveV5RespDataInner) GetTradeId() string`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetTradeIdOk() (*string, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *GetAccountBillsArchiveV5RespDataInner) SetTradeId(v string)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *GetAccountBillsArchiveV5RespDataInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetTs

`func (o *GetAccountBillsArchiveV5RespDataInner) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetAccountBillsArchiveV5RespDataInner) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetAccountBillsArchiveV5RespDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetType

`func (o *GetAccountBillsArchiveV5RespDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAccountBillsArchiveV5RespDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAccountBillsArchiveV5RespDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAccountBillsArchiveV5RespDataInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


