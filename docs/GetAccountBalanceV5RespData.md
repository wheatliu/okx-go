# GetAccountBalanceV5RespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjEq** | Pointer to **string** | Adjusted / Effective equity in &#x60;USD&#x60;   The net fiat value of the assets in the account that can provide margins for spot, expiry futures, perpetual futures and options under the cross-margin mode.   In multi-ccy or PM mode, the asset and margin requirement will all be converted to USD value to process the order check or liquidation.   Due to the volatility of each currency market, our platform calculates the actual USD value of each currency based on discount rates to balance market risks.   Applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60; and &#x60;Portfolio margin&#x60; | [optional] [default to ""]
**BorrowFroz** | Pointer to **string** | Potential borrowing IMR of the account in &#x60;USD&#x60;   Only applicable to &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60;. It is \&quot;\&quot; for other margin modes. | [optional] [default to ""]
**Details** | Pointer to [**[]GetAccountBalanceV5RespDataDetailsInner**](GetAccountBalanceV5RespDataDetailsInner.md) | Detailed asset information in all currencies | [optional] 
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

### NewGetAccountBalanceV5RespData

`func NewGetAccountBalanceV5RespData() *GetAccountBalanceV5RespData`

NewGetAccountBalanceV5RespData instantiates a new GetAccountBalanceV5RespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountBalanceV5RespDataWithDefaults

`func NewGetAccountBalanceV5RespDataWithDefaults() *GetAccountBalanceV5RespData`

NewGetAccountBalanceV5RespDataWithDefaults instantiates a new GetAccountBalanceV5RespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjEq

`func (o *GetAccountBalanceV5RespData) GetAdjEq() string`

GetAdjEq returns the AdjEq field if non-nil, zero value otherwise.

### GetAdjEqOk

`func (o *GetAccountBalanceV5RespData) GetAdjEqOk() (*string, bool)`

GetAdjEqOk returns a tuple with the AdjEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjEq

`func (o *GetAccountBalanceV5RespData) SetAdjEq(v string)`

SetAdjEq sets AdjEq field to given value.

### HasAdjEq

`func (o *GetAccountBalanceV5RespData) HasAdjEq() bool`

HasAdjEq returns a boolean if a field has been set.

### GetBorrowFroz

`func (o *GetAccountBalanceV5RespData) GetBorrowFroz() string`

GetBorrowFroz returns the BorrowFroz field if non-nil, zero value otherwise.

### GetBorrowFrozOk

`func (o *GetAccountBalanceV5RespData) GetBorrowFrozOk() (*string, bool)`

GetBorrowFrozOk returns a tuple with the BorrowFroz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowFroz

`func (o *GetAccountBalanceV5RespData) SetBorrowFroz(v string)`

SetBorrowFroz sets BorrowFroz field to given value.

### HasBorrowFroz

`func (o *GetAccountBalanceV5RespData) HasBorrowFroz() bool`

HasBorrowFroz returns a boolean if a field has been set.

### GetDetails

`func (o *GetAccountBalanceV5RespData) GetDetails() []GetAccountBalanceV5RespDataDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetAccountBalanceV5RespData) GetDetailsOk() (*[]GetAccountBalanceV5RespDataDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetAccountBalanceV5RespData) SetDetails(v []GetAccountBalanceV5RespDataDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetAccountBalanceV5RespData) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetImr

`func (o *GetAccountBalanceV5RespData) GetImr() string`

GetImr returns the Imr field if non-nil, zero value otherwise.

### GetImrOk

`func (o *GetAccountBalanceV5RespData) GetImrOk() (*string, bool)`

GetImrOk returns a tuple with the Imr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImr

`func (o *GetAccountBalanceV5RespData) SetImr(v string)`

SetImr sets Imr field to given value.

### HasImr

`func (o *GetAccountBalanceV5RespData) HasImr() bool`

HasImr returns a boolean if a field has been set.

### GetIsoEq

`func (o *GetAccountBalanceV5RespData) GetIsoEq() string`

GetIsoEq returns the IsoEq field if non-nil, zero value otherwise.

### GetIsoEqOk

`func (o *GetAccountBalanceV5RespData) GetIsoEqOk() (*string, bool)`

GetIsoEqOk returns a tuple with the IsoEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoEq

`func (o *GetAccountBalanceV5RespData) SetIsoEq(v string)`

SetIsoEq sets IsoEq field to given value.

### HasIsoEq

`func (o *GetAccountBalanceV5RespData) HasIsoEq() bool`

HasIsoEq returns a boolean if a field has been set.

### GetMgnRatio

`func (o *GetAccountBalanceV5RespData) GetMgnRatio() string`

GetMgnRatio returns the MgnRatio field if non-nil, zero value otherwise.

### GetMgnRatioOk

`func (o *GetAccountBalanceV5RespData) GetMgnRatioOk() (*string, bool)`

GetMgnRatioOk returns a tuple with the MgnRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgnRatio

`func (o *GetAccountBalanceV5RespData) SetMgnRatio(v string)`

SetMgnRatio sets MgnRatio field to given value.

### HasMgnRatio

`func (o *GetAccountBalanceV5RespData) HasMgnRatio() bool`

HasMgnRatio returns a boolean if a field has been set.

### GetMmr

`func (o *GetAccountBalanceV5RespData) GetMmr() string`

GetMmr returns the Mmr field if non-nil, zero value otherwise.

### GetMmrOk

`func (o *GetAccountBalanceV5RespData) GetMmrOk() (*string, bool)`

GetMmrOk returns a tuple with the Mmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmr

`func (o *GetAccountBalanceV5RespData) SetMmr(v string)`

SetMmr sets Mmr field to given value.

### HasMmr

`func (o *GetAccountBalanceV5RespData) HasMmr() bool`

HasMmr returns a boolean if a field has been set.

### GetNotionalUsd

`func (o *GetAccountBalanceV5RespData) GetNotionalUsd() string`

GetNotionalUsd returns the NotionalUsd field if non-nil, zero value otherwise.

### GetNotionalUsdOk

`func (o *GetAccountBalanceV5RespData) GetNotionalUsdOk() (*string, bool)`

GetNotionalUsdOk returns a tuple with the NotionalUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsd

`func (o *GetAccountBalanceV5RespData) SetNotionalUsd(v string)`

SetNotionalUsd sets NotionalUsd field to given value.

### HasNotionalUsd

`func (o *GetAccountBalanceV5RespData) HasNotionalUsd() bool`

HasNotionalUsd returns a boolean if a field has been set.

### GetNotionalUsdForBorrow

`func (o *GetAccountBalanceV5RespData) GetNotionalUsdForBorrow() string`

GetNotionalUsdForBorrow returns the NotionalUsdForBorrow field if non-nil, zero value otherwise.

### GetNotionalUsdForBorrowOk

`func (o *GetAccountBalanceV5RespData) GetNotionalUsdForBorrowOk() (*string, bool)`

GetNotionalUsdForBorrowOk returns a tuple with the NotionalUsdForBorrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsdForBorrow

`func (o *GetAccountBalanceV5RespData) SetNotionalUsdForBorrow(v string)`

SetNotionalUsdForBorrow sets NotionalUsdForBorrow field to given value.

### HasNotionalUsdForBorrow

`func (o *GetAccountBalanceV5RespData) HasNotionalUsdForBorrow() bool`

HasNotionalUsdForBorrow returns a boolean if a field has been set.

### GetNotionalUsdForFutures

`func (o *GetAccountBalanceV5RespData) GetNotionalUsdForFutures() string`

GetNotionalUsdForFutures returns the NotionalUsdForFutures field if non-nil, zero value otherwise.

### GetNotionalUsdForFuturesOk

`func (o *GetAccountBalanceV5RespData) GetNotionalUsdForFuturesOk() (*string, bool)`

GetNotionalUsdForFuturesOk returns a tuple with the NotionalUsdForFutures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsdForFutures

`func (o *GetAccountBalanceV5RespData) SetNotionalUsdForFutures(v string)`

SetNotionalUsdForFutures sets NotionalUsdForFutures field to given value.

### HasNotionalUsdForFutures

`func (o *GetAccountBalanceV5RespData) HasNotionalUsdForFutures() bool`

HasNotionalUsdForFutures returns a boolean if a field has been set.

### GetNotionalUsdForOption

`func (o *GetAccountBalanceV5RespData) GetNotionalUsdForOption() string`

GetNotionalUsdForOption returns the NotionalUsdForOption field if non-nil, zero value otherwise.

### GetNotionalUsdForOptionOk

`func (o *GetAccountBalanceV5RespData) GetNotionalUsdForOptionOk() (*string, bool)`

GetNotionalUsdForOptionOk returns a tuple with the NotionalUsdForOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsdForOption

`func (o *GetAccountBalanceV5RespData) SetNotionalUsdForOption(v string)`

SetNotionalUsdForOption sets NotionalUsdForOption field to given value.

### HasNotionalUsdForOption

`func (o *GetAccountBalanceV5RespData) HasNotionalUsdForOption() bool`

HasNotionalUsdForOption returns a boolean if a field has been set.

### GetNotionalUsdForSwap

`func (o *GetAccountBalanceV5RespData) GetNotionalUsdForSwap() string`

GetNotionalUsdForSwap returns the NotionalUsdForSwap field if non-nil, zero value otherwise.

### GetNotionalUsdForSwapOk

`func (o *GetAccountBalanceV5RespData) GetNotionalUsdForSwapOk() (*string, bool)`

GetNotionalUsdForSwapOk returns a tuple with the NotionalUsdForSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotionalUsdForSwap

`func (o *GetAccountBalanceV5RespData) SetNotionalUsdForSwap(v string)`

SetNotionalUsdForSwap sets NotionalUsdForSwap field to given value.

### HasNotionalUsdForSwap

`func (o *GetAccountBalanceV5RespData) HasNotionalUsdForSwap() bool`

HasNotionalUsdForSwap returns a boolean if a field has been set.

### GetOrdFroz

`func (o *GetAccountBalanceV5RespData) GetOrdFroz() string`

GetOrdFroz returns the OrdFroz field if non-nil, zero value otherwise.

### GetOrdFrozOk

`func (o *GetAccountBalanceV5RespData) GetOrdFrozOk() (*string, bool)`

GetOrdFrozOk returns a tuple with the OrdFroz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdFroz

`func (o *GetAccountBalanceV5RespData) SetOrdFroz(v string)`

SetOrdFroz sets OrdFroz field to given value.

### HasOrdFroz

`func (o *GetAccountBalanceV5RespData) HasOrdFroz() bool`

HasOrdFroz returns a boolean if a field has been set.

### GetTotalEq

`func (o *GetAccountBalanceV5RespData) GetTotalEq() string`

GetTotalEq returns the TotalEq field if non-nil, zero value otherwise.

### GetTotalEqOk

`func (o *GetAccountBalanceV5RespData) GetTotalEqOk() (*string, bool)`

GetTotalEqOk returns a tuple with the TotalEq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEq

`func (o *GetAccountBalanceV5RespData) SetTotalEq(v string)`

SetTotalEq sets TotalEq field to given value.

### HasTotalEq

`func (o *GetAccountBalanceV5RespData) HasTotalEq() bool`

HasTotalEq returns a boolean if a field has been set.

### GetUTime

`func (o *GetAccountBalanceV5RespData) GetUTime() string`

GetUTime returns the UTime field if non-nil, zero value otherwise.

### GetUTimeOk

`func (o *GetAccountBalanceV5RespData) GetUTimeOk() (*string, bool)`

GetUTimeOk returns a tuple with the UTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUTime

`func (o *GetAccountBalanceV5RespData) SetUTime(v string)`

SetUTime sets UTime field to given value.

### HasUTime

`func (o *GetAccountBalanceV5RespData) HasUTime() bool`

HasUTime returns a boolean if a field has been set.

### GetUpl

`func (o *GetAccountBalanceV5RespData) GetUpl() string`

GetUpl returns the Upl field if non-nil, zero value otherwise.

### GetUplOk

`func (o *GetAccountBalanceV5RespData) GetUplOk() (*string, bool)`

GetUplOk returns a tuple with the Upl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpl

`func (o *GetAccountBalanceV5RespData) SetUpl(v string)`

SetUpl sets Upl field to given value.

### HasUpl

`func (o *GetAccountBalanceV5RespData) HasUpl() bool`

HasUpl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


