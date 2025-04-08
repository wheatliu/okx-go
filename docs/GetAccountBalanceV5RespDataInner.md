# GetAccountBalanceV5RespDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjEq** | Pointer to **string** | Adjusted / Effective equity in &#x60;USD&#x60;   The net fiat value of the assets in the account that can provide margins for spot, expiry futures, perpetual futures and options under the cross-margin mode.   In multi-ccy or PM mode, the asset and margin requirement will all be converted to USD value to process the order check or liquidation.   Due to the volatility of each currency market, our platform calculates the actual USD value of each currency based on discount rates to balance market risks.   Applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60; and &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**BorrowFroz** | Pointer to **string** | Potential borrowing IMR of the account in &#x60;USD&#x60;   Only applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60;. It is \&quot;\&quot; for other margin modes. | [optional] [default to ""]
**Details** | Pointer to [**[]GetAccountBalanceV5RespDataInnerDetailsInner**](GetAccountBalanceV5RespDataInnerDetailsInner.md) | Detailed asset information in all currencies | [optional] 
**Imr** | Pointer to **string** | Initial margin requirement in &#x60;USD&#x60;   The sum of initial margins of all open positions and pending orders under cross-margin mode in &#x60;USD&#x60;.   Applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**IsoEq** | Pointer to **string** | Isolated margin equity in &#x60;USD&#x60;  Applicable to &#x60;Spot and futures mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**MgnRatio** | Pointer to **string** | Margin ratio in &#x60;USD&#x60;   Applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**Mmr** | Pointer to **string** | Maintenance margin requirement in &#x60;USD&#x60;   The sum of maintenance margins of all open positions and pending orders under cross-margin mode in &#x60;USD&#x60;.   Applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**NotionalUsd** | Pointer to **string** | Notional value of positions in &#x60;USD&#x60;   Applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**NotionalUsdForBorrow** | Pointer to **string** | Notional value for &#x60;Borrow&#x60; in USD  Applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**NotionalUsdForFutures** | Pointer to **string** | Notional value of positions for &#x60;Expiry Futures&#x60; in USD  Applicable to &#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**NotionalUsdForOption** | Pointer to **string** | Notional value of positions for &#x60;Option&#x60; in USD  Applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**NotionalUsdForSwap** | Pointer to **string** | Notional value of positions for &#x60;Perpetual Futures&#x60; in USD  Applicable to &#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**OrdFroz** | Pointer to **string** | Cross margin frozen for pending orders in &#x60;USD&#x60;   Only applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]
**TotalEq** | Pointer to **string** | The total amount of equity in &#x60;USD&#x60; | [optional] [default to ""]
**UTime** | Pointer to **string** | Update time of account information, millisecond format of Unix timestamp, e.g. &#x60;1597026383085&#x60; | [optional] [default to ""]
**Upl** | Pointer to **string** | Cross-margin info of unrealized profit and loss at the account level in &#x60;USD&#x60;  Applicable to &#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60; | [optional] [default to ""]

## Methods

### NewGetAccountBalanceV5RespDataInner

`func NewGetAccountBalanceV5RespDataInner() *GetAccountBalanceV5RespDataInner`

NewGetAccountBalanceV5RespDataInner instantiates a new GetAccountBalanceV5RespDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountBalanceV5RespDataInnerWithDefaults

`func NewGetAccountBalanceV5RespDataInnerWithDefaults() *GetAccountBalanceV5RespDataInner`

NewGetAccountBalanceV5RespDataInnerWithDefaults instantiates a new GetAccountBalanceV5RespDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjEq

`func (o *GetAccountBalanceV5RespDataInner) GetAdjEq() string`

GetAdjEq returns the AdjEq field if non-nil, zero value otherwise.

### GetAdjEqOk

`func (o *GetAccountBalanceV5RespDataInner) GetAdjEqOk() (*string, bool)`

GetAdjEqOk returns a tuple with the AdjEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjEq

`func (o *GetAccountBalanceV5RespDataInner) SetAdjEq(v string)`

SetAdjEq sets AdjEq field to given value.

### HasAdjEq

`func (o *GetAccountBalanceV5RespDataInner) HasAdjEq() bool`

HasAdjEq returns a boolean if a field has been set.

### GetBorrowFroz

`func (o *GetAccountBalanceV5RespDataInner) GetBorrowFroz() string`

GetBorrowFroz returns the BorrowFroz field if non-nil, zero value otherwise.

### GetBorrowFrozOk

`func (o *GetAccountBalanceV5RespDataInner) GetBorrowFrozOk() (*string, bool)`

GetBorrowFrozOk returns a tuple with the BorrowFroz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowFroz

`func (o *GetAccountBalanceV5RespDataInner) SetBorrowFroz(v string)`

SetBorrowFroz sets BorrowFroz field to given value.

### HasBorrowFroz

`func (o *GetAccountBalanceV5RespDataInner) HasBorrowFroz() bool`

HasBorrowFroz returns a boolean if a field has been set.

### GetDetails

`func (o *GetAccountBalanceV5RespDataInner) GetDetails() []GetAccountBalanceV5RespDataInnerDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetAccountBalanceV5RespDataInner) GetDetailsOk() (*[]GetAccountBalanceV5RespDataInnerDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetAccountBalanceV5RespDataInner) SetDetails(v []GetAccountBalanceV5RespDataInnerDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetAccountBalanceV5RespDataInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetImr

`func (o *GetAccountBalanceV5RespDataInner) GetImr() string`

GetImr returns the Imr field if non-nil, zero value otherwise.

### GetImrOk

`func (o *GetAccountBalanceV5RespDataInner) GetImrOk() (*string, bool)`

GetImrOk returns a tuple with the Imr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImr

`func (o *GetAccountBalanceV5RespDataInner) SetImr(v string)`

SetImr sets Imr field to given value.

### HasImr

`func (o *GetAccountBalanceV5RespDataInner) HasImr() bool`

HasImr returns a boolean if a field has been set.

### GetIsoEq

`func (o *GetAccountBalanceV5RespDataInner) GetIsoEq() string`

GetIsoEq returns the IsoEq field if non-nil, zero value otherwise.

### GetIsoEqOk

`func (o *GetAccountBalanceV5RespDataInner) GetIsoEqOk() (*string, bool)`

GetIsoEqOk returns a tuple with the IsoEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoEq

`func (o *GetAccountBalanceV5RespDataInner) SetIsoEq(v string)`

SetIsoEq sets IsoEq field to given value.

### HasIsoEq

`func (o *GetAccountBalanceV5RespDataInner) HasIsoEq() bool`

HasIsoEq returns a boolean if a field has been set.

### GetMgnRatio

`func (o *GetAccountBalanceV5RespDataInner) GetMgnRatio() string`

GetMgnRatio returns the MgnRatio field if non-nil, zero value otherwise.

### GetMgnRatioOk

`func (o *GetAccountBalanceV5RespDataInner) GetMgnRatioOk() (*string, bool)`

GetMgnRatioOk returns a tuple with the MgnRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnRatio

`func (o *GetAccountBalanceV5RespDataInner) SetMgnRatio(v string)`

SetMgnRatio sets MgnRatio field to given value.

### HasMgnRatio

`func (o *GetAccountBalanceV5RespDataInner) HasMgnRatio() bool`

HasMgnRatio returns a boolean if a field has been set.

### GetMmr

`func (o *GetAccountBalanceV5RespDataInner) GetMmr() string`

GetMmr returns the Mmr field if non-nil, zero value otherwise.

### GetMmrOk

`func (o *GetAccountBalanceV5RespDataInner) GetMmrOk() (*string, bool)`

GetMmrOk returns a tuple with the Mmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmr

`func (o *GetAccountBalanceV5RespDataInner) SetMmr(v string)`

SetMmr sets Mmr field to given value.

### HasMmr

`func (o *GetAccountBalanceV5RespDataInner) HasMmr() bool`

HasMmr returns a boolean if a field has been set.

### GetNotionalUsd

`func (o *GetAccountBalanceV5RespDataInner) GetNotionalUsd() string`

GetNotionalUsd returns the NotionalUsd field if non-nil, zero value otherwise.

### GetNotionalUsdOk

`func (o *GetAccountBalanceV5RespDataInner) GetNotionalUsdOk() (*string, bool)`

GetNotionalUsdOk returns a tuple with the NotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsd

`func (o *GetAccountBalanceV5RespDataInner) SetNotionalUsd(v string)`

SetNotionalUsd sets NotionalUsd field to given value.

### HasNotionalUsd

`func (o *GetAccountBalanceV5RespDataInner) HasNotionalUsd() bool`

HasNotionalUsd returns a boolean if a field has been set.

### GetNotionalUsdForBorrow

`func (o *GetAccountBalanceV5RespDataInner) GetNotionalUsdForBorrow() string`

GetNotionalUsdForBorrow returns the NotionalUsdForBorrow field if non-nil, zero value otherwise.

### GetNotionalUsdForBorrowOk

`func (o *GetAccountBalanceV5RespDataInner) GetNotionalUsdForBorrowOk() (*string, bool)`

GetNotionalUsdForBorrowOk returns a tuple with the NotionalUsdForBorrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsdForBorrow

`func (o *GetAccountBalanceV5RespDataInner) SetNotionalUsdForBorrow(v string)`

SetNotionalUsdForBorrow sets NotionalUsdForBorrow field to given value.

### HasNotionalUsdForBorrow

`func (o *GetAccountBalanceV5RespDataInner) HasNotionalUsdForBorrow() bool`

HasNotionalUsdForBorrow returns a boolean if a field has been set.

### GetNotionalUsdForFutures

`func (o *GetAccountBalanceV5RespDataInner) GetNotionalUsdForFutures() string`

GetNotionalUsdForFutures returns the NotionalUsdForFutures field if non-nil, zero value otherwise.

### GetNotionalUsdForFuturesOk

`func (o *GetAccountBalanceV5RespDataInner) GetNotionalUsdForFuturesOk() (*string, bool)`

GetNotionalUsdForFuturesOk returns a tuple with the NotionalUsdForFutures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsdForFutures

`func (o *GetAccountBalanceV5RespDataInner) SetNotionalUsdForFutures(v string)`

SetNotionalUsdForFutures sets NotionalUsdForFutures field to given value.

### HasNotionalUsdForFutures

`func (o *GetAccountBalanceV5RespDataInner) HasNotionalUsdForFutures() bool`

HasNotionalUsdForFutures returns a boolean if a field has been set.

### GetNotionalUsdForOption

`func (o *GetAccountBalanceV5RespDataInner) GetNotionalUsdForOption() string`

GetNotionalUsdForOption returns the NotionalUsdForOption field if non-nil, zero value otherwise.

### GetNotionalUsdForOptionOk

`func (o *GetAccountBalanceV5RespDataInner) GetNotionalUsdForOptionOk() (*string, bool)`

GetNotionalUsdForOptionOk returns a tuple with the NotionalUsdForOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsdForOption

`func (o *GetAccountBalanceV5RespDataInner) SetNotionalUsdForOption(v string)`

SetNotionalUsdForOption sets NotionalUsdForOption field to given value.

### HasNotionalUsdForOption

`func (o *GetAccountBalanceV5RespDataInner) HasNotionalUsdForOption() bool`

HasNotionalUsdForOption returns a boolean if a field has been set.

### GetNotionalUsdForSwap

`func (o *GetAccountBalanceV5RespDataInner) GetNotionalUsdForSwap() string`

GetNotionalUsdForSwap returns the NotionalUsdForSwap field if non-nil, zero value otherwise.

### GetNotionalUsdForSwapOk

`func (o *GetAccountBalanceV5RespDataInner) GetNotionalUsdForSwapOk() (*string, bool)`

GetNotionalUsdForSwapOk returns a tuple with the NotionalUsdForSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsdForSwap

`func (o *GetAccountBalanceV5RespDataInner) SetNotionalUsdForSwap(v string)`

SetNotionalUsdForSwap sets NotionalUsdForSwap field to given value.

### HasNotionalUsdForSwap

`func (o *GetAccountBalanceV5RespDataInner) HasNotionalUsdForSwap() bool`

HasNotionalUsdForSwap returns a boolean if a field has been set.

### GetOrdFroz

`func (o *GetAccountBalanceV5RespDataInner) GetOrdFroz() string`

GetOrdFroz returns the OrdFroz field if non-nil, zero value otherwise.

### GetOrdFrozOk

`func (o *GetAccountBalanceV5RespDataInner) GetOrdFrozOk() (*string, bool)`

GetOrdFrozOk returns a tuple with the OrdFroz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdFroz

`func (o *GetAccountBalanceV5RespDataInner) SetOrdFroz(v string)`

SetOrdFroz sets OrdFroz field to given value.

### HasOrdFroz

`func (o *GetAccountBalanceV5RespDataInner) HasOrdFroz() bool`

HasOrdFroz returns a boolean if a field has been set.

### GetTotalEq

`func (o *GetAccountBalanceV5RespDataInner) GetTotalEq() string`

GetTotalEq returns the TotalEq field if non-nil, zero value otherwise.

### GetTotalEqOk

`func (o *GetAccountBalanceV5RespDataInner) GetTotalEqOk() (*string, bool)`

GetTotalEqOk returns a tuple with the TotalEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEq

`func (o *GetAccountBalanceV5RespDataInner) SetTotalEq(v string)`

SetTotalEq sets TotalEq field to given value.

### HasTotalEq

`func (o *GetAccountBalanceV5RespDataInner) HasTotalEq() bool`

HasTotalEq returns a boolean if a field has been set.

### GetUTime

`func (o *GetAccountBalanceV5RespDataInner) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetAccountBalanceV5RespDataInner) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetAccountBalanceV5RespDataInner) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetAccountBalanceV5RespDataInner) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUpl

`func (o *GetAccountBalanceV5RespDataInner) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *GetAccountBalanceV5RespDataInner) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *GetAccountBalanceV5RespDataInner) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *GetAccountBalanceV5RespDataInner) HasUpl() bool`

HasUpl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


