# GetAccountBalanceV5RespDataInnerDetailsInner

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
**CollateralEnabled** | Pointer to **bool** | &#x60;true&#x60;: Collateral enabled  &#x60;false&#x60;: Collateral disabled  Applicable to &#x60;Multi-currency margin&#x60;   | [optional] 
**CrossLiab** | Pointer to **string** | Cross liabilities of currency  Applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**DisEq** | Pointer to **string** | Discount equity of currency in &#x60;USD&#x60;.  Applicable to &#x60;Spot mode&#x60;(enabled spot borrow)/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
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
**OrdFrozen** | Pointer to **string** | Margin frozen for open orders  Applicable to &#x60;Spot mode&#x60;/&#x60;Spot and futures mode&#x60;/&#x60;Multi-currency margin&#x60; | [optional] [default to ""]
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

### NewGetAccountBalanceV5RespDataInnerDetailsInner

`func NewGetAccountBalanceV5RespDataInnerDetailsInner() *GetAccountBalanceV5RespDataInnerDetailsInner`

NewGetAccountBalanceV5RespDataInnerDetailsInner instantiates a new GetAccountBalanceV5RespDataInnerDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountBalanceV5RespDataInnerDetailsInnerWithDefaults

`func NewGetAccountBalanceV5RespDataInnerDetailsInnerWithDefaults() *GetAccountBalanceV5RespDataInnerDetailsInner`

NewGetAccountBalanceV5RespDataInnerDetailsInnerWithDefaults instantiates a new GetAccountBalanceV5RespDataInnerDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccAvgPx

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetAccAvgPx() string`

GetAccAvgPx returns the AccAvgPx field if non-nil, zero value otherwise.

### GetAccAvgPxOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetAccAvgPxOk() (*string, bool)`

GetAccAvgPxOk returns a tuple with the AccAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccAvgPx

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetAccAvgPx(v string)`

SetAccAvgPx sets AccAvgPx field to given value.

### HasAccAvgPx

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasAccAvgPx() bool`

HasAccAvgPx returns a boolean if a field has been set.

### GetAvailBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetAvailBal() string`

GetAvailBal returns the AvailBal field if non-nil, zero value otherwise.

### GetAvailBalOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetAvailBalOk() (*string, bool)`

GetAvailBalOk returns a tuple with the AvailBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetAvailBal(v string)`

SetAvailBal sets AvailBal field to given value.

### HasAvailBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasAvailBal() bool`

HasAvailBal returns a boolean if a field has been set.

### GetAvailEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetAvailEq() string`

GetAvailEq returns the AvailEq field if non-nil, zero value otherwise.

### GetAvailEqOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetAvailEqOk() (*string, bool)`

GetAvailEqOk returns a tuple with the AvailEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetAvailEq(v string)`

SetAvailEq sets AvailEq field to given value.

### HasAvailEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasAvailEq() bool`

HasAvailEq returns a boolean if a field has been set.

### GetBorrowFroz

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetBorrowFroz() string`

GetBorrowFroz returns the BorrowFroz field if non-nil, zero value otherwise.

### GetBorrowFrozOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetBorrowFrozOk() (*string, bool)`

GetBorrowFrozOk returns a tuple with the BorrowFroz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowFroz

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetBorrowFroz(v string)`

SetBorrowFroz sets BorrowFroz field to given value.

### HasBorrowFroz

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasBorrowFroz() bool`

HasBorrowFroz returns a boolean if a field has been set.

### GetCashBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetCashBal() string`

GetCashBal returns the CashBal field if non-nil, zero value otherwise.

### GetCashBalOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetCashBalOk() (*string, bool)`

GetCashBalOk returns a tuple with the CashBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetCashBal(v string)`

SetCashBal sets CashBal field to given value.

### HasCashBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasCashBal() bool`

HasCashBal returns a boolean if a field has been set.

### GetCcy

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetCcy() string`

GetCcy returns the Ccy field if non-nil, zero value otherwise.

### GetCcyOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetCcyOk() (*string, bool)`

GetCcyOk returns a tuple with the Ccy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcy

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetCcy(v string)`

SetCcy sets Ccy field to given value.

### HasCcy

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasCcy() bool`

HasCcy returns a boolean if a field has been set.

### GetClSpotInUseAmt

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetClSpotInUseAmt() string`

GetClSpotInUseAmt returns the ClSpotInUseAmt field if non-nil, zero value otherwise.

### GetClSpotInUseAmtOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetClSpotInUseAmtOk() (*string, bool)`

GetClSpotInUseAmtOk returns a tuple with the ClSpotInUseAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClSpotInUseAmt

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetClSpotInUseAmt(v string)`

SetClSpotInUseAmt sets ClSpotInUseAmt field to given value.

### HasClSpotInUseAmt

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasClSpotInUseAmt() bool`

HasClSpotInUseAmt returns a boolean if a field has been set.

### GetCollateralEnabled

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetCollateralEnabled() bool`

GetCollateralEnabled returns the CollateralEnabled field if non-nil, zero value otherwise.

### GetCollateralEnabledOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetCollateralEnabledOk() (*bool, bool)`

GetCollateralEnabledOk returns a tuple with the CollateralEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralEnabled

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetCollateralEnabled(v bool)`

SetCollateralEnabled sets CollateralEnabled field to given value.

### HasCollateralEnabled

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasCollateralEnabled() bool`

HasCollateralEnabled returns a boolean if a field has been set.

### GetCrossLiab

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetCrossLiab() string`

GetCrossLiab returns the CrossLiab field if non-nil, zero value otherwise.

### GetCrossLiabOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetCrossLiabOk() (*string, bool)`

GetCrossLiabOk returns a tuple with the CrossLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossLiab

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetCrossLiab(v string)`

SetCrossLiab sets CrossLiab field to given value.

### HasCrossLiab

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasCrossLiab() bool`

HasCrossLiab returns a boolean if a field has been set.

### GetDisEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetDisEq() string`

GetDisEq returns the DisEq field if non-nil, zero value otherwise.

### GetDisEqOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetDisEqOk() (*string, bool)`

GetDisEqOk returns a tuple with the DisEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetDisEq(v string)`

SetDisEq sets DisEq field to given value.

### HasDisEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasDisEq() bool`

HasDisEq returns a boolean if a field has been set.

### GetEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetEq() string`

GetEq returns the Eq field if non-nil, zero value otherwise.

### GetEqOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetEqOk() (*string, bool)`

GetEqOk returns a tuple with the Eq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetEq(v string)`

SetEq sets Eq field to given value.

### HasEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasEq() bool`

HasEq returns a boolean if a field has been set.

### GetEqUsd

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetEqUsd() string`

GetEqUsd returns the EqUsd field if non-nil, zero value otherwise.

### GetEqUsdOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetEqUsdOk() (*string, bool)`

GetEqUsdOk returns a tuple with the EqUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqUsd

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetEqUsd(v string)`

SetEqUsd sets EqUsd field to given value.

### HasEqUsd

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasEqUsd() bool`

HasEqUsd returns a boolean if a field has been set.

### GetFixedBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetFixedBal() string`

GetFixedBal returns the FixedBal field if non-nil, zero value otherwise.

### GetFixedBalOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetFixedBalOk() (*string, bool)`

GetFixedBalOk returns a tuple with the FixedBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetFixedBal(v string)`

SetFixedBal sets FixedBal field to given value.

### HasFixedBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasFixedBal() bool`

HasFixedBal returns a boolean if a field has been set.

### GetFrozenBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetFrozenBal() string`

GetFrozenBal returns the FrozenBal field if non-nil, zero value otherwise.

### GetFrozenBalOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetFrozenBalOk() (*string, bool)`

GetFrozenBalOk returns a tuple with the FrozenBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetFrozenBal(v string)`

SetFrozenBal sets FrozenBal field to given value.

### HasFrozenBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasFrozenBal() bool`

HasFrozenBal returns a boolean if a field has been set.

### GetImr

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetImr() string`

GetImr returns the Imr field if non-nil, zero value otherwise.

### GetImrOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetImrOk() (*string, bool)`

GetImrOk returns a tuple with the Imr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImr

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetImr(v string)`

SetImr sets Imr field to given value.

### HasImr

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasImr() bool`

HasImr returns a boolean if a field has been set.

### GetInterest

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetIsoEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetIsoEq() string`

GetIsoEq returns the IsoEq field if non-nil, zero value otherwise.

### GetIsoEqOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetIsoEqOk() (*string, bool)`

GetIsoEqOk returns a tuple with the IsoEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetIsoEq(v string)`

SetIsoEq sets IsoEq field to given value.

### HasIsoEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasIsoEq() bool`

HasIsoEq returns a boolean if a field has been set.

### GetIsoLiab

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetIsoLiab() string`

GetIsoLiab returns the IsoLiab field if non-nil, zero value otherwise.

### GetIsoLiabOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetIsoLiabOk() (*string, bool)`

GetIsoLiabOk returns a tuple with the IsoLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoLiab

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetIsoLiab(v string)`

SetIsoLiab sets IsoLiab field to given value.

### HasIsoLiab

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasIsoLiab() bool`

HasIsoLiab returns a boolean if a field has been set.

### GetIsoUpl

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetIsoUpl() string`

GetIsoUpl returns the IsoUpl field if non-nil, zero value otherwise.

### GetIsoUplOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetIsoUplOk() (*string, bool)`

GetIsoUplOk returns a tuple with the IsoUpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoUpl

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetIsoUpl(v string)`

SetIsoUpl sets IsoUpl field to given value.

### HasIsoUpl

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasIsoUpl() bool`

HasIsoUpl returns a boolean if a field has been set.

### GetLiab

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetLiab() string`

GetLiab returns the Liab field if non-nil, zero value otherwise.

### GetLiabOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetLiabOk() (*string, bool)`

GetLiabOk returns a tuple with the Liab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiab

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetLiab(v string)`

SetLiab sets Liab field to given value.

### HasLiab

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasLiab() bool`

HasLiab returns a boolean if a field has been set.

### GetMaxLoan

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetMaxLoan() string`

GetMaxLoan returns the MaxLoan field if non-nil, zero value otherwise.

### GetMaxLoanOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetMaxLoanOk() (*string, bool)`

GetMaxLoanOk returns a tuple with the MaxLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoan

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetMaxLoan(v string)`

SetMaxLoan sets MaxLoan field to given value.

### HasMaxLoan

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasMaxLoan() bool`

HasMaxLoan returns a boolean if a field has been set.

### GetMaxSpotInUse

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetMaxSpotInUse() string`

GetMaxSpotInUse returns the MaxSpotInUse field if non-nil, zero value otherwise.

### GetMaxSpotInUseOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetMaxSpotInUseOk() (*string, bool)`

GetMaxSpotInUseOk returns a tuple with the MaxSpotInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpotInUse

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetMaxSpotInUse(v string)`

SetMaxSpotInUse sets MaxSpotInUse field to given value.

### HasMaxSpotInUse

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasMaxSpotInUse() bool`

HasMaxSpotInUse returns a boolean if a field has been set.

### GetMgnRatio

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetMgnRatio() string`

GetMgnRatio returns the MgnRatio field if non-nil, zero value otherwise.

### GetMgnRatioOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetMgnRatioOk() (*string, bool)`

GetMgnRatioOk returns a tuple with the MgnRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnRatio

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetMgnRatio(v string)`

SetMgnRatio sets MgnRatio field to given value.

### HasMgnRatio

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasMgnRatio() bool`

HasMgnRatio returns a boolean if a field has been set.

### GetMmr

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetMmr() string`

GetMmr returns the Mmr field if non-nil, zero value otherwise.

### GetMmrOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetMmrOk() (*string, bool)`

GetMmrOk returns a tuple with the Mmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmr

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetMmr(v string)`

SetMmr sets Mmr field to given value.

### HasMmr

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasMmr() bool`

HasMmr returns a boolean if a field has been set.

### GetNotionalLever

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetNotionalLever() string`

GetNotionalLever returns the NotionalLever field if non-nil, zero value otherwise.

### GetNotionalLeverOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetNotionalLeverOk() (*string, bool)`

GetNotionalLeverOk returns a tuple with the NotionalLever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalLever

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetNotionalLever(v string)`

SetNotionalLever sets NotionalLever field to given value.

### HasNotionalLever

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasNotionalLever() bool`

HasNotionalLever returns a boolean if a field has been set.

### GetOpenAvgPx

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetOpenAvgPx() string`

GetOpenAvgPx returns the OpenAvgPx field if non-nil, zero value otherwise.

### GetOpenAvgPxOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetOpenAvgPxOk() (*string, bool)`

GetOpenAvgPxOk returns a tuple with the OpenAvgPx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAvgPx

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetOpenAvgPx(v string)`

SetOpenAvgPx sets OpenAvgPx field to given value.

### HasOpenAvgPx

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasOpenAvgPx() bool`

HasOpenAvgPx returns a boolean if a field has been set.

### GetOrdFrozen

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetOrdFrozen() string`

GetOrdFrozen returns the OrdFrozen field if non-nil, zero value otherwise.

### GetOrdFrozenOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetOrdFrozenOk() (*string, bool)`

GetOrdFrozenOk returns a tuple with the OrdFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdFrozen

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetOrdFrozen(v string)`

SetOrdFrozen sets OrdFrozen field to given value.

### HasOrdFrozen

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasOrdFrozen() bool`

HasOrdFrozen returns a boolean if a field has been set.

### GetRewardBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetRewardBal() string`

GetRewardBal returns the RewardBal field if non-nil, zero value otherwise.

### GetRewardBalOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetRewardBalOk() (*string, bool)`

GetRewardBalOk returns a tuple with the RewardBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetRewardBal(v string)`

SetRewardBal sets RewardBal field to given value.

### HasRewardBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasRewardBal() bool`

HasRewardBal returns a boolean if a field has been set.

### GetSmtSyncEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSmtSyncEq() string`

GetSmtSyncEq returns the SmtSyncEq field if non-nil, zero value otherwise.

### GetSmtSyncEqOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSmtSyncEqOk() (*string, bool)`

GetSmtSyncEqOk returns a tuple with the SmtSyncEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtSyncEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetSmtSyncEq(v string)`

SetSmtSyncEq sets SmtSyncEq field to given value.

### HasSmtSyncEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasSmtSyncEq() bool`

HasSmtSyncEq returns a boolean if a field has been set.

### GetSpotBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSpotBal() string`

GetSpotBal returns the SpotBal field if non-nil, zero value otherwise.

### GetSpotBalOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSpotBalOk() (*string, bool)`

GetSpotBalOk returns a tuple with the SpotBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetSpotBal(v string)`

SetSpotBal sets SpotBal field to given value.

### HasSpotBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasSpotBal() bool`

HasSpotBal returns a boolean if a field has been set.

### GetSpotCopyTradingEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSpotCopyTradingEq() string`

GetSpotCopyTradingEq returns the SpotCopyTradingEq field if non-nil, zero value otherwise.

### GetSpotCopyTradingEqOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSpotCopyTradingEqOk() (*string, bool)`

GetSpotCopyTradingEqOk returns a tuple with the SpotCopyTradingEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotCopyTradingEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetSpotCopyTradingEq(v string)`

SetSpotCopyTradingEq sets SpotCopyTradingEq field to given value.

### HasSpotCopyTradingEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasSpotCopyTradingEq() bool`

HasSpotCopyTradingEq returns a boolean if a field has been set.

### GetSpotInUseAmt

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSpotInUseAmt() string`

GetSpotInUseAmt returns the SpotInUseAmt field if non-nil, zero value otherwise.

### GetSpotInUseAmtOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSpotInUseAmtOk() (*string, bool)`

GetSpotInUseAmtOk returns a tuple with the SpotInUseAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInUseAmt

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetSpotInUseAmt(v string)`

SetSpotInUseAmt sets SpotInUseAmt field to given value.

### HasSpotInUseAmt

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasSpotInUseAmt() bool`

HasSpotInUseAmt returns a boolean if a field has been set.

### GetSpotIsoBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSpotIsoBal() string`

GetSpotIsoBal returns the SpotIsoBal field if non-nil, zero value otherwise.

### GetSpotIsoBalOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSpotIsoBalOk() (*string, bool)`

GetSpotIsoBalOk returns a tuple with the SpotIsoBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotIsoBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetSpotIsoBal(v string)`

SetSpotIsoBal sets SpotIsoBal field to given value.

### HasSpotIsoBal

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasSpotIsoBal() bool`

HasSpotIsoBal returns a boolean if a field has been set.

### GetSpotUpl

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSpotUpl() string`

GetSpotUpl returns the SpotUpl field if non-nil, zero value otherwise.

### GetSpotUplOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSpotUplOk() (*string, bool)`

GetSpotUplOk returns a tuple with the SpotUpl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotUpl

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetSpotUpl(v string)`

SetSpotUpl sets SpotUpl field to given value.

### HasSpotUpl

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasSpotUpl() bool`

HasSpotUpl returns a boolean if a field has been set.

### GetSpotUplRatio

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSpotUplRatio() string`

GetSpotUplRatio returns the SpotUplRatio field if non-nil, zero value otherwise.

### GetSpotUplRatioOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetSpotUplRatioOk() (*string, bool)`

GetSpotUplRatioOk returns a tuple with the SpotUplRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotUplRatio

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetSpotUplRatio(v string)`

SetSpotUplRatio sets SpotUplRatio field to given value.

### HasSpotUplRatio

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasSpotUplRatio() bool`

HasSpotUplRatio returns a boolean if a field has been set.

### GetStgyEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetStgyEq() string`

GetStgyEq returns the StgyEq field if non-nil, zero value otherwise.

### GetStgyEqOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetStgyEqOk() (*string, bool)`

GetStgyEqOk returns a tuple with the StgyEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStgyEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetStgyEq(v string)`

SetStgyEq sets StgyEq field to given value.

### HasStgyEq

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasStgyEq() bool`

HasStgyEq returns a boolean if a field has been set.

### GetTotalPnl

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetTotalPnl() string`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetTotalPnlOk() (*string, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetTotalPnl(v string)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.

### GetTotalPnlRatio

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetTotalPnlRatio() string`

GetTotalPnlRatio returns the TotalPnlRatio field if non-nil, zero value otherwise.

### GetTotalPnlRatioOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetTotalPnlRatioOk() (*string, bool)`

GetTotalPnlRatioOk returns a tuple with the TotalPnlRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnlRatio

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetTotalPnlRatio(v string)`

SetTotalPnlRatio sets TotalPnlRatio field to given value.

### HasTotalPnlRatio

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasTotalPnlRatio() bool`

HasTotalPnlRatio returns a boolean if a field has been set.

### GetTwap

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetTwap() string`

GetTwap returns the Twap field if non-nil, zero value otherwise.

### GetTwapOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetTwapOk() (*string, bool)`

GetTwapOk returns a tuple with the Twap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwap

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetTwap(v string)`

SetTwap sets Twap field to given value.

### HasTwap

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasTwap() bool`

HasTwap returns a boolean if a field has been set.

### GetUTime

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUpl

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasUpl() bool`

HasUpl returns a boolean if a field has been set.

### GetUplLiab

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetUplLiab() string`

GetUplLiab returns the UplLiab field if non-nil, zero value otherwise.

### GetUplLiabOk

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) GetUplLiabOk() (*string, bool)`

GetUplLiabOk returns a tuple with the UplLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplLiab

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) SetUplLiab(v string)`

SetUplLiab sets UplLiab field to given value.

### HasUplLiab

`func (o *GetAccountBalanceV5RespDataInnerDetailsInner) HasUplLiab() bool`

HasUplLiab returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


