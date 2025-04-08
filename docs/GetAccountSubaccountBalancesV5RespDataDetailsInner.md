# GetAccountSubaccountBalancesV5RespDataDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccAvgPx** | Pointer to **string** | Spot accumulated cost price. The unit is USD.  | [optional] [default to ""]
**AvailBal** | Pointer to **string** | Available balance of currency | [optional] [default to ""]
**AvailEq** | Pointer to **string** | Available equity of currency  Applicable to &#x60;Spot and futures mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**BorrowFroz** | Pointer to **string** | Potential borrowing IMR of currency in &#x60;USD&#x60;   Applicable to &#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60;. It is \&quot;\&quot; for other margin modes. | [optional] [default to ""]
**CashBal** | Pointer to **string** | Cash balance | [optional] [default to ""]
**Ccy** | Pointer to **string** | Currency | [optional] [default to ""]
**ClSpotInUseAmt** | Pointer to **string** | User-defined spot risk offset amount  Applicable to &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**CrossLiab** | Pointer to **string** | Cross liabilities of currency  Applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**DisEq** | Pointer to **string** | Discount equity of currency in &#x60;USD&#x60;. | [optional] [default to ""]
**Eq** | Pointer to **string** | Equity of currency | [optional] [default to ""]
**EqUsd** | Pointer to **string** | Equity in &#x60;USD&#x60; of currency | [optional] [default to ""]
**FixedBal** | Pointer to **string** | Frozen balance for &#x60;Dip Sniper&#x60; and &#x60;Peak Sniper&#x60; | [optional] [default to ""]
**FrozenBal** | Pointer to **string** | Frozen balance of currency | [optional] [default to ""]
**Imr** | Pointer to **string** | Cross initial margin requirement at the currency level  Applicable to &#x60;Spot and futures mode&#x60; and when there is cross position | [optional] [default to ""]
**Interest** | Pointer to **string** | Accrued interest of currency  It is a positive value, e.g. &#x60;9.01&#x60;  Applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**IsoEq** | Pointer to **string** | Isolated margin equity of currency  Applicable to &#x60;Spot and futures mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**IsoLiab** | Pointer to **string** | Isolated liabilities of currency  Applicable to &#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**IsoUpl** | Pointer to **string** | Isolated unrealized profit and loss of currency  Applicable to &#x60;Spot and futures mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**Liab** | Pointer to **string** | Liabilities of currency  It is a positive value, e.g. &#x60;21625.64&#x60;  Applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**MaxLoan** | Pointer to **string** | Max loan of currency  Applicable to &#x60;cross&#x60; of &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**MaxSpotInUse** | Pointer to **string** | Max possible spot risk offset amount  Applicable to &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**MgnRatio** | Pointer to **string** | Cross margin ratio of currency   The index for measuring the risk of a certain asset in the account.   Applicable to &#x60;Spot and futures mode&#x60; and when there is cross position | [optional] [default to ""]
**Mmr** | Pointer to **string** | Cross maintenance margin requirement at the currency level  Applicable to &#x60;Spot and futures mode&#x60; and when there is cross position | [optional] [default to ""]
**NotionalLever** | Pointer to **string** | Leverage of currency  Applicable to &#x60;Spot and futures mode&#x60; | [optional] [default to ""]
**OpenAvgPx** | Pointer to **string** | Spot average cost price. The unit is USD.  | [optional] [default to ""]
**OrdFrozen** | Pointer to **string** | Margin frozen for open orders   Applicable to &#x60;Spot mode&#x60;/&#x60;Spot and futures mode&#x60;/&#x60;Multi-currency margin&#x60; | [optional] [default to ""]
**RewardBal** | Pointer to **string** | Trial fund balance | [optional] [default to ""]
**SmtSyncEq** | Pointer to **string** | Smart sync equity  The default is \&quot;0\&quot;, only applicable to copy trader. | [optional] [default to ""]
**SpotBal** | Pointer to **string** | Spot balance. The unit is currency, e.g. BTC.  | [optional] [default to ""]
**SpotCopyTradingEq** | Pointer to **string** | Spot smart sync equity.   The default is \&quot;0\&quot;, only applicable to copy trader. | [optional] [default to ""]
**SpotInUseAmt** | Pointer to **string** | Spot in use amount  Applicable to &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**SpotIsoBal** | Pointer to **string** | Spot isolated balance  Applicable to copy trading  Applicable to &#x60;Spot mode&#x60;/&#x60;Spot and futures mode&#x60;. | [optional] [default to ""]
**SpotUpl** | Pointer to **string** | Spot unrealized profit and loss. The unit is USD.  | [optional] [default to ""]
**SpotUplRatio** | Pointer to **string** | Spot unrealized profit and loss ratio.  | [optional] [default to ""]
**StgyEq** | Pointer to **string** | Strategy equity | [optional] [default to ""]
**TotalPnl** | Pointer to **string** | Spot accumulated profit and loss. The unit is USD.  | [optional] [default to ""]
**TotalPnlRatio** | Pointer to **string** | Spot accumulated profit and loss ratio.  | [optional] [default to ""]
**Twap** | Pointer to **string** | Risk indicator of auto liability repayment  Divided into multiple levels from 0 to 5, the larger the number, the more likely the auto repayment will be triggered.   Applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**UTime** | Pointer to **string** | Update time of currency balance information, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Upl** | Pointer to **string** | The sum of the unrealized profit &amp; loss of all margin and derivatives positions of currency.   Applicable to &#x60;Spot and futures mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**UplLiab** | Pointer to **string** | Liabilities due to Unrealized loss of currency  Applicable to &#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountSubaccountBalancesV5RespDataDetailsInner

`func NewGetAccountSubaccountBalancesV5RespDataDetailsInner() *GetAccountSubaccountBalancesV5RespDataDetailsInner`

NewGetAccountSubaccountBalancesV5RespDataDetailsInner instantiates a new GetAccountSubaccountBalancesV5RespDataDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountSubaccountBalancesV5RespDataDetailsInnerWithDefaults

`func NewGetAccountSubaccountBalancesV5RespDataDetailsInnerWithDefaults() *GetAccountSubaccountBalancesV5RespDataDetailsInner`

NewGetAccountSubaccountBalancesV5RespDataDetailsInnerWithDefaults instantiates a new GetAccountSubaccountBalancesV5RespDataDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccAvgPx

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetAccAvgPx() string`

GetAccAvgPx returns the AccAvgPx field if non-nil, zero value otherwise.

### GetAccAvgPxOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetAccAvgPxOk() (*string, bool)`

GetAccAvgPxOk returns a tuple with the AccAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccAvgPx

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetAccAvgPx(v string)`

SetAccAvgPx sets AccAvgPx field to given value.

### HasAccAvgPx

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasAccAvgPx() bool`

HasAccAvgPx returns a boolean if a field has been set.

### GetAvailBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetAvailBal() string`

GetAvailBal returns the AvailBal field if non-nil, zero value otherwise.

### GetAvailBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetAvailBalOk() (*string, bool)`

GetAvailBalOk returns a tuple with the AvailBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetAvailBal(v string)`

SetAvailBal sets AvailBal field to given value.

### HasAvailBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasAvailBal() bool`

HasAvailBal returns a boolean if a field has been set.

### GetAvailEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetAvailEq() string`

GetAvailEq returns the AvailEq field if non-nil, zero value otherwise.

### GetAvailEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetAvailEqOk() (*string, bool)`

GetAvailEqOk returns a tuple with the AvailEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetAvailEq(v string)`

SetAvailEq sets AvailEq field to given value.

### HasAvailEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasAvailEq() bool`

HasAvailEq returns a boolean if a field has been set.

### GetBorrowFroz

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetBorrowFroz() string`

GetBorrowFroz returns the BorrowFroz field if non-nil, zero value otherwise.

### GetBorrowFrozOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetBorrowFrozOk() (*string, bool)`

GetBorrowFrozOk returns a tuple with the BorrowFroz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowFroz

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetBorrowFroz(v string)`

SetBorrowFroz sets BorrowFroz field to given value.

### HasBorrowFroz

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasBorrowFroz() bool`

HasBorrowFroz returns a boolean if a field has been set.

### GetCashBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetCashBal() string`

GetCashBal returns the CashBal field if non-nil, zero value otherwise.

### GetCashBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetCashBalOk() (*string, bool)`

GetCashBalOk returns a tuple with the CashBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetCashBal(v string)`

SetCashBal sets CashBal field to given value.

### HasCashBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasCashBal() bool`

HasCashBal returns a boolean if a field has been set.

### GetCcy

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClSpotInUseAmt

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetClSpotInUseAmt() string`

GetClSpotInUseAmt returns the ClSpotInUseAmt field if non-nil, zero value otherwise.

### GetClSpotInUseAmtOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetClSpotInUseAmtOk() (*string, bool)`

GetClSpotInUseAmtOk returns a tuple with the ClSpotInUseAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClSpotInUseAmt

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetClSpotInUseAmt(v string)`

SetClSpotInUseAmt sets ClSpotInUseAmt field to given value.

### HasClSpotInUseAmt

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasClSpotInUseAmt() bool`

HasClSpotInUseAmt returns a boolean if a field has been set.

### GetCrossLiab

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetCrossLiab() string`

GetCrossLiab returns the CrossLiab field if non-nil, zero value otherwise.

### GetCrossLiabOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetCrossLiabOk() (*string, bool)`

GetCrossLiabOk returns a tuple with the CrossLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossLiab

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetCrossLiab(v string)`

SetCrossLiab sets CrossLiab field to given value.

### HasCrossLiab

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasCrossLiab() bool`

HasCrossLiab returns a boolean if a field has been set.

### GetDisEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetDisEq() string`

GetDisEq returns the DisEq field if non-nil, zero value otherwise.

### GetDisEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetDisEqOk() (*string, bool)`

GetDisEqOk returns a tuple with the DisEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetDisEq(v string)`

SetDisEq sets DisEq field to given value.

### HasDisEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasDisEq() bool`

HasDisEq returns a boolean if a field has been set.

### GetEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetEqUsd

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetEqUsd() string`

GetEqUsd returns the EqUsd field if non-nil, zero value otherwise.

### GetEqUsdOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetEqUsdOk() (*string, bool)`

GetEqUsdOk returns a tuple with the EqUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqUsd

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetEqUsd(v string)`

SetEqUsd sets EqUsd field to given value.

### HasEqUsd

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasEqUsd() bool`

HasEqUsd returns a boolean if a field has been set.

### GetFixedBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetFixedBal() string`

GetFixedBal returns the FixedBal field if non-nil, zero value otherwise.

### GetFixedBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetFixedBalOk() (*string, bool)`

GetFixedBalOk returns a tuple with the FixedBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetFixedBal(v string)`

SetFixedBal sets FixedBal field to given value.

### HasFixedBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasFixedBal() bool`

HasFixedBal returns a boolean if a field has been set.

### GetFrozenBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetFrozenBal() string`

GetFrozenBal returns the FrozenBal field if non-nil, zero value otherwise.

### GetFrozenBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetFrozenBalOk() (*string, bool)`

GetFrozenBalOk returns a tuple with the FrozenBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetFrozenBal(v string)`

SetFrozenBal sets FrozenBal field to given value.

### HasFrozenBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasFrozenBal() bool`

HasFrozenBal returns a boolean if a field has been set.

### GetImr

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetImr() string`

GetImr returns the Imr field if non-nil, zero value otherwise.

### GetImrOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetImrOk() (*string, bool)`

GetImrOk returns a tuple with the Imr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImr

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetImr(v string)`

SetImr sets Imr field to given value.

### HasImr

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasImr() bool`

HasImr returns a boolean if a field has been set.

### GetInterest

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetIsoEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetIsoEq() string`

GetIsoEq returns the IsoEq field if non-nil, zero value otherwise.

### GetIsoEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetIsoEqOk() (*string, bool)`

GetIsoEqOk returns a tuple with the IsoEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetIsoEq(v string)`

SetIsoEq sets IsoEq field to given value.

### HasIsoEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasIsoEq() bool`

HasIsoEq returns a boolean if a field has been set.

### GetIsoLiab

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetIsoLiab() string`

GetIsoLiab returns the IsoLiab field if non-nil, zero value otherwise.

### GetIsoLiabOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetIsoLiabOk() (*string, bool)`

GetIsoLiabOk returns a tuple with the IsoLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoLiab

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetIsoLiab(v string)`

SetIsoLiab sets IsoLiab field to given value.

### HasIsoLiab

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasIsoLiab() bool`

HasIsoLiab returns a boolean if a field has been set.

### GetIsoUpl

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetIsoUpl() string`

GetIsoUpl returns the IsoUpl field if non-nil, zero value otherwise.

### GetIsoUplOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetIsoUplOk() (*string, bool)`

GetIsoUplOk returns a tuple with the IsoUpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoUpl

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetIsoUpl(v string)`

SetIsoUpl sets IsoUpl field to given value.

### HasIsoUpl

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasIsoUpl() bool`

HasIsoUpl returns a boolean if a field has been set.

### GetLiab

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetLiab() string`

GetLiab returns the Liab field if non-nil, zero value otherwise.

### GetLiabOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetLiabOk() (*string, bool)`

GetLiabOk returns a tuple with the Liab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiab

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetLiab(v string)`

SetLiab sets Liab field to given value.

### HasLiab

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasLiab() bool`

HasLiab returns a boolean if a field has been set.

### GetMaxLoan

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetMaxLoan() string`

GetMaxLoan returns the MaxLoan field if non-nil, zero value otherwise.

### GetMaxLoanOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetMaxLoanOk() (*string, bool)`

GetMaxLoanOk returns a tuple with the MaxLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoan

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetMaxLoan(v string)`

SetMaxLoan sets MaxLoan field to given value.

### HasMaxLoan

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasMaxLoan() bool`

HasMaxLoan returns a boolean if a field has been set.

### GetMaxSpotInUse

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetMaxSpotInUse() string`

GetMaxSpotInUse returns the MaxSpotInUse field if non-nil, zero value otherwise.

### GetMaxSpotInUseOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetMaxSpotInUseOk() (*string, bool)`

GetMaxSpotInUseOk returns a tuple with the MaxSpotInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpotInUse

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetMaxSpotInUse(v string)`

SetMaxSpotInUse sets MaxSpotInUse field to given value.

### HasMaxSpotInUse

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasMaxSpotInUse() bool`

HasMaxSpotInUse returns a boolean if a field has been set.

### GetMgnRatio

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetMgnRatio() string`

GetMgnRatio returns the MgnRatio field if non-nil, zero value otherwise.

### GetMgnRatioOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetMgnRatioOk() (*string, bool)`

GetMgnRatioOk returns a tuple with the MgnRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnRatio

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetMgnRatio(v string)`

SetMgnRatio sets MgnRatio field to given value.

### HasMgnRatio

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasMgnRatio() bool`

HasMgnRatio returns a boolean if a field has been set.

### GetMmr

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetMmr() string`

GetMmr returns the Mmr field if non-nil, zero value otherwise.

### GetMmrOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetMmrOk() (*string, bool)`

GetMmrOk returns a tuple with the Mmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmr

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetMmr(v string)`

SetMmr sets Mmr field to given value.

### HasMmr

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasMmr() bool`

HasMmr returns a boolean if a field has been set.

### GetNotionalLever

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetNotionalLever() string`

GetNotionalLever returns the NotionalLever field if non-nil, zero value otherwise.

### GetNotionalLeverOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetNotionalLeverOk() (*string, bool)`

GetNotionalLeverOk returns a tuple with the NotionalLever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalLever

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetNotionalLever(v string)`

SetNotionalLever sets NotionalLever field to given value.

### HasNotionalLever

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasNotionalLever() bool`

HasNotionalLever returns a boolean if a field has been set.

### GetOpenAvgPx

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetOpenAvgPx() string`

GetOpenAvgPx returns the OpenAvgPx field if non-nil, zero value otherwise.

### GetOpenAvgPxOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetOpenAvgPxOk() (*string, bool)`

GetOpenAvgPxOk returns a tuple with the OpenAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAvgPx

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetOpenAvgPx(v string)`

SetOpenAvgPx sets OpenAvgPx field to given value.

### HasOpenAvgPx

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasOpenAvgPx() bool`

HasOpenAvgPx returns a boolean if a field has been set.

### GetOrdFrozen

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetOrdFrozen() string`

GetOrdFrozen returns the OrdFrozen field if non-nil, zero value otherwise.

### GetOrdFrozenOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetOrdFrozenOk() (*string, bool)`

GetOrdFrozenOk returns a tuple with the OrdFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdFrozen

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetOrdFrozen(v string)`

SetOrdFrozen sets OrdFrozen field to given value.

### HasOrdFrozen

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasOrdFrozen() bool`

HasOrdFrozen returns a boolean if a field has been set.

### GetRewardBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetRewardBal() string`

GetRewardBal returns the RewardBal field if non-nil, zero value otherwise.

### GetRewardBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetRewardBalOk() (*string, bool)`

GetRewardBalOk returns a tuple with the RewardBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetRewardBal(v string)`

SetRewardBal sets RewardBal field to given value.

### HasRewardBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasRewardBal() bool`

HasRewardBal returns a boolean if a field has been set.

### GetSmtSyncEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSmtSyncEq() string`

GetSmtSyncEq returns the SmtSyncEq field if non-nil, zero value otherwise.

### GetSmtSyncEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSmtSyncEqOk() (*string, bool)`

GetSmtSyncEqOk returns a tuple with the SmtSyncEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtSyncEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetSmtSyncEq(v string)`

SetSmtSyncEq sets SmtSyncEq field to given value.

### HasSmtSyncEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasSmtSyncEq() bool`

HasSmtSyncEq returns a boolean if a field has been set.

### GetSpotBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSpotBal() string`

GetSpotBal returns the SpotBal field if non-nil, zero value otherwise.

### GetSpotBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSpotBalOk() (*string, bool)`

GetSpotBalOk returns a tuple with the SpotBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetSpotBal(v string)`

SetSpotBal sets SpotBal field to given value.

### HasSpotBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasSpotBal() bool`

HasSpotBal returns a boolean if a field has been set.

### GetSpotCopyTradingEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSpotCopyTradingEq() string`

GetSpotCopyTradingEq returns the SpotCopyTradingEq field if non-nil, zero value otherwise.

### GetSpotCopyTradingEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSpotCopyTradingEqOk() (*string, bool)`

GetSpotCopyTradingEqOk returns a tuple with the SpotCopyTradingEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotCopyTradingEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetSpotCopyTradingEq(v string)`

SetSpotCopyTradingEq sets SpotCopyTradingEq field to given value.

### HasSpotCopyTradingEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasSpotCopyTradingEq() bool`

HasSpotCopyTradingEq returns a boolean if a field has been set.

### GetSpotInUseAmt

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSpotInUseAmt() string`

GetSpotInUseAmt returns the SpotInUseAmt field if non-nil, zero value otherwise.

### GetSpotInUseAmtOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSpotInUseAmtOk() (*string, bool)`

GetSpotInUseAmtOk returns a tuple with the SpotInUseAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInUseAmt

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetSpotInUseAmt(v string)`

SetSpotInUseAmt sets SpotInUseAmt field to given value.

### HasSpotInUseAmt

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasSpotInUseAmt() bool`

HasSpotInUseAmt returns a boolean if a field has been set.

### GetSpotIsoBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSpotIsoBal() string`

GetSpotIsoBal returns the SpotIsoBal field if non-nil, zero value otherwise.

### GetSpotIsoBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSpotIsoBalOk() (*string, bool)`

GetSpotIsoBalOk returns a tuple with the SpotIsoBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotIsoBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetSpotIsoBal(v string)`

SetSpotIsoBal sets SpotIsoBal field to given value.

### HasSpotIsoBal

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasSpotIsoBal() bool`

HasSpotIsoBal returns a boolean if a field has been set.

### GetSpotUpl

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSpotUpl() string`

GetSpotUpl returns the SpotUpl field if non-nil, zero value otherwise.

### GetSpotUplOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSpotUplOk() (*string, bool)`

GetSpotUplOk returns a tuple with the SpotUpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotUpl

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetSpotUpl(v string)`

SetSpotUpl sets SpotUpl field to given value.

### HasSpotUpl

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasSpotUpl() bool`

HasSpotUpl returns a boolean if a field has been set.

### GetSpotUplRatio

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSpotUplRatio() string`

GetSpotUplRatio returns the SpotUplRatio field if non-nil, zero value otherwise.

### GetSpotUplRatioOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetSpotUplRatioOk() (*string, bool)`

GetSpotUplRatioOk returns a tuple with the SpotUplRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotUplRatio

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetSpotUplRatio(v string)`

SetSpotUplRatio sets SpotUplRatio field to given value.

### HasSpotUplRatio

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasSpotUplRatio() bool`

HasSpotUplRatio returns a boolean if a field has been set.

### GetStgyEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetStgyEq() string`

GetStgyEq returns the StgyEq field if non-nil, zero value otherwise.

### GetStgyEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetStgyEqOk() (*string, bool)`

GetStgyEqOk returns a tuple with the StgyEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStgyEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetStgyEq(v string)`

SetStgyEq sets StgyEq field to given value.

### HasStgyEq

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasStgyEq() bool`

HasStgyEq returns a boolean if a field has been set.

### GetTotalPnl

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetTotalPnl() string`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetTotalPnlOk() (*string, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetTotalPnl(v string)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.

### GetTotalPnlRatio

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetTotalPnlRatio() string`

GetTotalPnlRatio returns the TotalPnlRatio field if non-nil, zero value otherwise.

### GetTotalPnlRatioOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetTotalPnlRatioOk() (*string, bool)`

GetTotalPnlRatioOk returns a tuple with the TotalPnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnlRatio

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetTotalPnlRatio(v string)`

SetTotalPnlRatio sets TotalPnlRatio field to given value.

### HasTotalPnlRatio

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasTotalPnlRatio() bool`

HasTotalPnlRatio returns a boolean if a field has been set.

### GetTwap

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetTwap() string`

GetTwap returns the Twap field if non-nil, zero value otherwise.

### GetTwapOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetTwapOk() (*string, bool)`

GetTwapOk returns a tuple with the Twap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwap

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetTwap(v string)`

SetTwap sets Twap field to given value.

### HasTwap

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasTwap() bool`

HasTwap returns a boolean if a field has been set.

### GetUTime

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUpl

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasUpl() bool`

HasUpl returns a boolean if a field has been set.

### GetUplLiab

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetUplLiab() string`

GetUplLiab returns the UplLiab field if non-nil, zero value otherwise.

### GetUplLiabOk

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) GetUplLiabOk() (*string, bool)`

GetUplLiabOk returns a tuple with the UplLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplLiab

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) SetUplLiab(v string)`

SetUplLiab sets UplLiab field to given value.

### HasUplLiab

`func (o *GetAccountSubaccountBalancesV5RespDataDetailsInner) HasUplLiab() bool`

HasUplLiab returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


