# GetAccountSubaccountBalancesV5RespDataInnerDetailsInner

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

### NewGetAccountSubaccountBalancesV5RespDataInnerDetailsInner

`func NewGetAccountSubaccountBalancesV5RespDataInnerDetailsInner() *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner`

NewGetAccountSubaccountBalancesV5RespDataInnerDetailsInner instantiates a new GetAccountSubaccountBalancesV5RespDataInnerDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountSubaccountBalancesV5RespDataInnerDetailsInnerWithDefaults

`func NewGetAccountSubaccountBalancesV5RespDataInnerDetailsInnerWithDefaults() *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner`

NewGetAccountSubaccountBalancesV5RespDataInnerDetailsInnerWithDefaults instantiates a new GetAccountSubaccountBalancesV5RespDataInnerDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccAvgPx

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetAccAvgPx() string`

GetAccAvgPx returns the AccAvgPx field if non-nil, zero value otherwise.

### GetAccAvgPxOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetAccAvgPxOk() (*string, bool)`

GetAccAvgPxOk returns a tuple with the AccAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccAvgPx

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetAccAvgPx(v string)`

SetAccAvgPx sets AccAvgPx field to given value.

### HasAccAvgPx

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasAccAvgPx() bool`

HasAccAvgPx returns a boolean if a field has been set.

### GetAvailBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetAvailBal() string`

GetAvailBal returns the AvailBal field if non-nil, zero value otherwise.

### GetAvailBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetAvailBalOk() (*string, bool)`

GetAvailBalOk returns a tuple with the AvailBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetAvailBal(v string)`

SetAvailBal sets AvailBal field to given value.

### HasAvailBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasAvailBal() bool`

HasAvailBal returns a boolean if a field has been set.

### GetAvailEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetAvailEq() string`

GetAvailEq returns the AvailEq field if non-nil, zero value otherwise.

### GetAvailEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetAvailEqOk() (*string, bool)`

GetAvailEqOk returns a tuple with the AvailEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetAvailEq(v string)`

SetAvailEq sets AvailEq field to given value.

### HasAvailEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasAvailEq() bool`

HasAvailEq returns a boolean if a field has been set.

### GetBorrowFroz

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetBorrowFroz() string`

GetBorrowFroz returns the BorrowFroz field if non-nil, zero value otherwise.

### GetBorrowFrozOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetBorrowFrozOk() (*string, bool)`

GetBorrowFrozOk returns a tuple with the BorrowFroz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowFroz

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetBorrowFroz(v string)`

SetBorrowFroz sets BorrowFroz field to given value.

### HasBorrowFroz

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasBorrowFroz() bool`

HasBorrowFroz returns a boolean if a field has been set.

### GetCashBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetCashBal() string`

GetCashBal returns the CashBal field if non-nil, zero value otherwise.

### GetCashBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetCashBalOk() (*string, bool)`

GetCashBalOk returns a tuple with the CashBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetCashBal(v string)`

SetCashBal sets CashBal field to given value.

### HasCashBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasCashBal() bool`

HasCashBal returns a boolean if a field has been set.

### GetCcy

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClSpotInUseAmt

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetClSpotInUseAmt() string`

GetClSpotInUseAmt returns the ClSpotInUseAmt field if non-nil, zero value otherwise.

### GetClSpotInUseAmtOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetClSpotInUseAmtOk() (*string, bool)`

GetClSpotInUseAmtOk returns a tuple with the ClSpotInUseAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClSpotInUseAmt

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetClSpotInUseAmt(v string)`

SetClSpotInUseAmt sets ClSpotInUseAmt field to given value.

### HasClSpotInUseAmt

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasClSpotInUseAmt() bool`

HasClSpotInUseAmt returns a boolean if a field has been set.

### GetCrossLiab

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetCrossLiab() string`

GetCrossLiab returns the CrossLiab field if non-nil, zero value otherwise.

### GetCrossLiabOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetCrossLiabOk() (*string, bool)`

GetCrossLiabOk returns a tuple with the CrossLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossLiab

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetCrossLiab(v string)`

SetCrossLiab sets CrossLiab field to given value.

### HasCrossLiab

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasCrossLiab() bool`

HasCrossLiab returns a boolean if a field has been set.

### GetDisEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetDisEq() string`

GetDisEq returns the DisEq field if non-nil, zero value otherwise.

### GetDisEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetDisEqOk() (*string, bool)`

GetDisEqOk returns a tuple with the DisEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetDisEq(v string)`

SetDisEq sets DisEq field to given value.

### HasDisEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasDisEq() bool`

HasDisEq returns a boolean if a field has been set.

### GetEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetEqUsd

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetEqUsd() string`

GetEqUsd returns the EqUsd field if non-nil, zero value otherwise.

### GetEqUsdOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetEqUsdOk() (*string, bool)`

GetEqUsdOk returns a tuple with the EqUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqUsd

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetEqUsd(v string)`

SetEqUsd sets EqUsd field to given value.

### HasEqUsd

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasEqUsd() bool`

HasEqUsd returns a boolean if a field has been set.

### GetFixedBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetFixedBal() string`

GetFixedBal returns the FixedBal field if non-nil, zero value otherwise.

### GetFixedBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetFixedBalOk() (*string, bool)`

GetFixedBalOk returns a tuple with the FixedBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetFixedBal(v string)`

SetFixedBal sets FixedBal field to given value.

### HasFixedBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasFixedBal() bool`

HasFixedBal returns a boolean if a field has been set.

### GetFrozenBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetFrozenBal() string`

GetFrozenBal returns the FrozenBal field if non-nil, zero value otherwise.

### GetFrozenBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetFrozenBalOk() (*string, bool)`

GetFrozenBalOk returns a tuple with the FrozenBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetFrozenBal(v string)`

SetFrozenBal sets FrozenBal field to given value.

### HasFrozenBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasFrozenBal() bool`

HasFrozenBal returns a boolean if a field has been set.

### GetImr

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetImr() string`

GetImr returns the Imr field if non-nil, zero value otherwise.

### GetImrOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetImrOk() (*string, bool)`

GetImrOk returns a tuple with the Imr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImr

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetImr(v string)`

SetImr sets Imr field to given value.

### HasImr

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasImr() bool`

HasImr returns a boolean if a field has been set.

### GetInterest

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetIsoEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetIsoEq() string`

GetIsoEq returns the IsoEq field if non-nil, zero value otherwise.

### GetIsoEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetIsoEqOk() (*string, bool)`

GetIsoEqOk returns a tuple with the IsoEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetIsoEq(v string)`

SetIsoEq sets IsoEq field to given value.

### HasIsoEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasIsoEq() bool`

HasIsoEq returns a boolean if a field has been set.

### GetIsoLiab

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetIsoLiab() string`

GetIsoLiab returns the IsoLiab field if non-nil, zero value otherwise.

### GetIsoLiabOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetIsoLiabOk() (*string, bool)`

GetIsoLiabOk returns a tuple with the IsoLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoLiab

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetIsoLiab(v string)`

SetIsoLiab sets IsoLiab field to given value.

### HasIsoLiab

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasIsoLiab() bool`

HasIsoLiab returns a boolean if a field has been set.

### GetIsoUpl

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetIsoUpl() string`

GetIsoUpl returns the IsoUpl field if non-nil, zero value otherwise.

### GetIsoUplOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetIsoUplOk() (*string, bool)`

GetIsoUplOk returns a tuple with the IsoUpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoUpl

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetIsoUpl(v string)`

SetIsoUpl sets IsoUpl field to given value.

### HasIsoUpl

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasIsoUpl() bool`

HasIsoUpl returns a boolean if a field has been set.

### GetLiab

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetLiab() string`

GetLiab returns the Liab field if non-nil, zero value otherwise.

### GetLiabOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetLiabOk() (*string, bool)`

GetLiabOk returns a tuple with the Liab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiab

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetLiab(v string)`

SetLiab sets Liab field to given value.

### HasLiab

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasLiab() bool`

HasLiab returns a boolean if a field has been set.

### GetMaxLoan

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetMaxLoan() string`

GetMaxLoan returns the MaxLoan field if non-nil, zero value otherwise.

### GetMaxLoanOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetMaxLoanOk() (*string, bool)`

GetMaxLoanOk returns a tuple with the MaxLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoan

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetMaxLoan(v string)`

SetMaxLoan sets MaxLoan field to given value.

### HasMaxLoan

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasMaxLoan() bool`

HasMaxLoan returns a boolean if a field has been set.

### GetMaxSpotInUse

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetMaxSpotInUse() string`

GetMaxSpotInUse returns the MaxSpotInUse field if non-nil, zero value otherwise.

### GetMaxSpotInUseOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetMaxSpotInUseOk() (*string, bool)`

GetMaxSpotInUseOk returns a tuple with the MaxSpotInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpotInUse

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetMaxSpotInUse(v string)`

SetMaxSpotInUse sets MaxSpotInUse field to given value.

### HasMaxSpotInUse

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasMaxSpotInUse() bool`

HasMaxSpotInUse returns a boolean if a field has been set.

### GetMgnRatio

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetMgnRatio() string`

GetMgnRatio returns the MgnRatio field if non-nil, zero value otherwise.

### GetMgnRatioOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetMgnRatioOk() (*string, bool)`

GetMgnRatioOk returns a tuple with the MgnRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnRatio

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetMgnRatio(v string)`

SetMgnRatio sets MgnRatio field to given value.

### HasMgnRatio

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasMgnRatio() bool`

HasMgnRatio returns a boolean if a field has been set.

### GetMmr

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetMmr() string`

GetMmr returns the Mmr field if non-nil, zero value otherwise.

### GetMmrOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetMmrOk() (*string, bool)`

GetMmrOk returns a tuple with the Mmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmr

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetMmr(v string)`

SetMmr sets Mmr field to given value.

### HasMmr

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasMmr() bool`

HasMmr returns a boolean if a field has been set.

### GetNotionalLever

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetNotionalLever() string`

GetNotionalLever returns the NotionalLever field if non-nil, zero value otherwise.

### GetNotionalLeverOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetNotionalLeverOk() (*string, bool)`

GetNotionalLeverOk returns a tuple with the NotionalLever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalLever

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetNotionalLever(v string)`

SetNotionalLever sets NotionalLever field to given value.

### HasNotionalLever

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasNotionalLever() bool`

HasNotionalLever returns a boolean if a field has been set.

### GetOpenAvgPx

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetOpenAvgPx() string`

GetOpenAvgPx returns the OpenAvgPx field if non-nil, zero value otherwise.

### GetOpenAvgPxOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetOpenAvgPxOk() (*string, bool)`

GetOpenAvgPxOk returns a tuple with the OpenAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAvgPx

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetOpenAvgPx(v string)`

SetOpenAvgPx sets OpenAvgPx field to given value.

### HasOpenAvgPx

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasOpenAvgPx() bool`

HasOpenAvgPx returns a boolean if a field has been set.

### GetOrdFrozen

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetOrdFrozen() string`

GetOrdFrozen returns the OrdFrozen field if non-nil, zero value otherwise.

### GetOrdFrozenOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetOrdFrozenOk() (*string, bool)`

GetOrdFrozenOk returns a tuple with the OrdFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdFrozen

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetOrdFrozen(v string)`

SetOrdFrozen sets OrdFrozen field to given value.

### HasOrdFrozen

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasOrdFrozen() bool`

HasOrdFrozen returns a boolean if a field has been set.

### GetRewardBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetRewardBal() string`

GetRewardBal returns the RewardBal field if non-nil, zero value otherwise.

### GetRewardBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetRewardBalOk() (*string, bool)`

GetRewardBalOk returns a tuple with the RewardBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetRewardBal(v string)`

SetRewardBal sets RewardBal field to given value.

### HasRewardBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasRewardBal() bool`

HasRewardBal returns a boolean if a field has been set.

### GetSmtSyncEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSmtSyncEq() string`

GetSmtSyncEq returns the SmtSyncEq field if non-nil, zero value otherwise.

### GetSmtSyncEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSmtSyncEqOk() (*string, bool)`

GetSmtSyncEqOk returns a tuple with the SmtSyncEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtSyncEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetSmtSyncEq(v string)`

SetSmtSyncEq sets SmtSyncEq field to given value.

### HasSmtSyncEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasSmtSyncEq() bool`

HasSmtSyncEq returns a boolean if a field has been set.

### GetSpotBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSpotBal() string`

GetSpotBal returns the SpotBal field if non-nil, zero value otherwise.

### GetSpotBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSpotBalOk() (*string, bool)`

GetSpotBalOk returns a tuple with the SpotBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetSpotBal(v string)`

SetSpotBal sets SpotBal field to given value.

### HasSpotBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasSpotBal() bool`

HasSpotBal returns a boolean if a field has been set.

### GetSpotCopyTradingEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSpotCopyTradingEq() string`

GetSpotCopyTradingEq returns the SpotCopyTradingEq field if non-nil, zero value otherwise.

### GetSpotCopyTradingEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSpotCopyTradingEqOk() (*string, bool)`

GetSpotCopyTradingEqOk returns a tuple with the SpotCopyTradingEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotCopyTradingEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetSpotCopyTradingEq(v string)`

SetSpotCopyTradingEq sets SpotCopyTradingEq field to given value.

### HasSpotCopyTradingEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasSpotCopyTradingEq() bool`

HasSpotCopyTradingEq returns a boolean if a field has been set.

### GetSpotInUseAmt

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSpotInUseAmt() string`

GetSpotInUseAmt returns the SpotInUseAmt field if non-nil, zero value otherwise.

### GetSpotInUseAmtOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSpotInUseAmtOk() (*string, bool)`

GetSpotInUseAmtOk returns a tuple with the SpotInUseAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInUseAmt

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetSpotInUseAmt(v string)`

SetSpotInUseAmt sets SpotInUseAmt field to given value.

### HasSpotInUseAmt

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasSpotInUseAmt() bool`

HasSpotInUseAmt returns a boolean if a field has been set.

### GetSpotIsoBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSpotIsoBal() string`

GetSpotIsoBal returns the SpotIsoBal field if non-nil, zero value otherwise.

### GetSpotIsoBalOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSpotIsoBalOk() (*string, bool)`

GetSpotIsoBalOk returns a tuple with the SpotIsoBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotIsoBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetSpotIsoBal(v string)`

SetSpotIsoBal sets SpotIsoBal field to given value.

### HasSpotIsoBal

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasSpotIsoBal() bool`

HasSpotIsoBal returns a boolean if a field has been set.

### GetSpotUpl

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSpotUpl() string`

GetSpotUpl returns the SpotUpl field if non-nil, zero value otherwise.

### GetSpotUplOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSpotUplOk() (*string, bool)`

GetSpotUplOk returns a tuple with the SpotUpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotUpl

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetSpotUpl(v string)`

SetSpotUpl sets SpotUpl field to given value.

### HasSpotUpl

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasSpotUpl() bool`

HasSpotUpl returns a boolean if a field has been set.

### GetSpotUplRatio

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSpotUplRatio() string`

GetSpotUplRatio returns the SpotUplRatio field if non-nil, zero value otherwise.

### GetSpotUplRatioOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetSpotUplRatioOk() (*string, bool)`

GetSpotUplRatioOk returns a tuple with the SpotUplRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotUplRatio

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetSpotUplRatio(v string)`

SetSpotUplRatio sets SpotUplRatio field to given value.

### HasSpotUplRatio

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasSpotUplRatio() bool`

HasSpotUplRatio returns a boolean if a field has been set.

### GetStgyEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetStgyEq() string`

GetStgyEq returns the StgyEq field if non-nil, zero value otherwise.

### GetStgyEqOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetStgyEqOk() (*string, bool)`

GetStgyEqOk returns a tuple with the StgyEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStgyEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetStgyEq(v string)`

SetStgyEq sets StgyEq field to given value.

### HasStgyEq

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasStgyEq() bool`

HasStgyEq returns a boolean if a field has been set.

### GetTotalPnl

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetTotalPnl() string`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetTotalPnlOk() (*string, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetTotalPnl(v string)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.

### GetTotalPnlRatio

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetTotalPnlRatio() string`

GetTotalPnlRatio returns the TotalPnlRatio field if non-nil, zero value otherwise.

### GetTotalPnlRatioOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetTotalPnlRatioOk() (*string, bool)`

GetTotalPnlRatioOk returns a tuple with the TotalPnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnlRatio

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetTotalPnlRatio(v string)`

SetTotalPnlRatio sets TotalPnlRatio field to given value.

### HasTotalPnlRatio

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasTotalPnlRatio() bool`

HasTotalPnlRatio returns a boolean if a field has been set.

### GetTwap

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetTwap() string`

GetTwap returns the Twap field if non-nil, zero value otherwise.

### GetTwapOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetTwapOk() (*string, bool)`

GetTwapOk returns a tuple with the Twap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwap

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetTwap(v string)`

SetTwap sets Twap field to given value.

### HasTwap

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasTwap() bool`

HasTwap returns a boolean if a field has been set.

### GetUTime

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUpl

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasUpl() bool`

HasUpl returns a boolean if a field has been set.

### GetUplLiab

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetUplLiab() string`

GetUplLiab returns the UplLiab field if non-nil, zero value otherwise.

### GetUplLiabOk

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) GetUplLiabOk() (*string, bool)`

GetUplLiabOk returns a tuple with the UplLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplLiab

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) SetUplLiab(v string)`

SetUplLiab sets UplLiab field to given value.

### HasUplLiab

`func (o *GetAccountSubaccountBalancesV5RespDataInnerDetailsInner) HasUplLiab() bool`

HasUplLiab returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


