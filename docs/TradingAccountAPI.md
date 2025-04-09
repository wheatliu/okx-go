# \TradingAccountAPI

All URIs are relative to *https://www.okx.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountAccountLevelSwitchPresetV5**](TradingAccountAPI.md#CreateAccountAccountLevelSwitchPresetV5) | **Post** /api/v5/account/account-level-switch-preset | Pre-set the required information for account mode switching. When switching from &#x60;Portfolio margin mode&#x60; back to &#x60;Spot and futures mode&#x60; / &#x60;Multi-currency margin mode&#x60;, and if there are existing cross-margin contract positions, it is mandatory to pre-set leverage.  If the user does not follow the required settings, they will receive an error message during the pre-check or when setting the account mode.  
[**CreateAccountActivateOptionV5**](TradingAccountAPI.md#CreateAccountActivateOptionV5) | **Post** /api/v5/account/activate-option | 
[**CreateAccountBillsHistoryArchiveV5**](TradingAccountAPI.md#CreateAccountBillsHistoryArchiveV5) | **Post** /api/v5/account/bills-history-archive | Apply for bill data since 1 February, 2021 except for the current quarter.  
[**CreateAccountMmpConfigV5**](TradingAccountAPI.md#CreateAccountMmpConfigV5) | **Post** /api/v5/account/mmp-config | This endpoint is used to set MMP configure    Only applicable to Option in Portfolio Margin mode, and MMP privilege is required.  
[**CreateAccountMmpResetV5**](TradingAccountAPI.md#CreateAccountMmpResetV5) | **Post** /api/v5/account/mmp-reset | You can unfreeze by this endpoint once MMP is triggered.    Only applicable to Option in Portfolio Margin mode, and MMP privilege is required.  
[**CreateAccountPositionBuilderV5**](TradingAccountAPI.md#CreateAccountPositionBuilderV5) | **Post** /api/v5/account/position-builder | Calculates portfolio margin information for virtual position/assets or current position of the user.   You can add up to 200 virtual positions and 200 virtual assets in one request.  
[**CreateAccountPositionMarginBalanceV5**](TradingAccountAPI.md#CreateAccountPositionMarginBalanceV5) | **Post** /api/v5/account/position/margin-balance | Increase or decrease the margin of the isolated position. Margin reduction may result in the change of the actual leverage.  
[**CreateAccountQuickMarginBorrowRepayV5**](TradingAccountAPI.md#CreateAccountQuickMarginBorrowRepayV5) | **Post** /api/v5/account/quick-margin-borrow-repay | Please note that this endpoint will be deprecated soon.  
[**CreateAccountSetAccountLevelV5**](TradingAccountAPI.md#CreateAccountSetAccountLevelV5) | **Post** /api/v5/account/set-account-level | You need to set on the Web/App for the first set of every account mode. If users plan to switch account modes while holding positions, they should first call the preset endpoint to conduct necessary settings, then call the precheck endpoint to get unmatched information, margin check, and other related information, and finally call the account mode switch endpoint to switch account modes.  
[**CreateAccountSetAutoLoanV5**](TradingAccountAPI.md#CreateAccountSetAutoLoanV5) | **Post** /api/v5/account/set-auto-loan | Only applicable to &#x60;Multi-currency margin&#x60; and &#x60;Portfolio margin&#x60;  
[**CreateAccountSetAutoRepayV5**](TradingAccountAPI.md#CreateAccountSetAutoRepayV5) | **Post** /api/v5/account/set-auto-repay | Only applicable to &#x60;Spot mode&#x60; (enabled borrowing)  
[**CreateAccountSetCollateralAssetsV5**](TradingAccountAPI.md#CreateAccountSetCollateralAssetsV5) | **Post** /api/v5/account/set-collateral-assets | 
[**CreateAccountSetGreeksV5**](TradingAccountAPI.md#CreateAccountSetGreeksV5) | **Post** /api/v5/account/set-greeks | Set the display type of Greeks.  
[**CreateAccountSetIsolatedModeV5**](TradingAccountAPI.md#CreateAccountSetIsolatedModeV5) | **Post** /api/v5/account/set-isolated-mode | You can set the currency margin and futures/perpetual Isolated margin trading mode  
[**CreateAccountSetLeverageV5**](TradingAccountAPI.md#CreateAccountSetLeverageV5) | **Post** /api/v5/account/set-leverage |   There are 10 different scenarios for leverage setting:      1. Set leverage for &#x60;MARGIN&#x60; instruments under &#x60;isolated-margin&#x60; trade mode at pairs level.    2. Set leverage for &#x60;MARGIN&#x60; instruments under &#x60;cross-margin&#x60; trade mode and Spot mode (enabled borrow) at currency level.    3. Set leverage for &#x60;MARGIN&#x60; instruments under &#x60;cross-margin&#x60; trade mode and Spot and futures mode account mode at pairs level.    4. Set leverage for &#x60;MARGIN&#x60; instruments under &#x60;cross-margin&#x60; trade mode and Multi-currency margin at currency level.    5. Set leverage for &#x60;MARGIN&#x60; instruments under &#x60;cross-margin&#x60; trade mode and Portfolio margin at currency level.    6. Set leverage for &#x60;FUTURES&#x60; instruments under &#x60;cross-margin&#x60; trade mode at underlying level.    7. Set leverage for &#x60;FUTURES&#x60; instruments under &#x60;isolated-margin&#x60; trade mode and buy/sell position mode at contract level.    8. Set leverage for &#x60;FUTURES&#x60; instruments under &#x60;isolated-margin&#x60; trade mode and long/short position mode at contract and position side level.    9. Set leverage for &#x60;SWAP&#x60; instruments under &#x60;cross-margin&#x60; trade at contract level.    10. Set leverage for &#x60;SWAP&#x60; instruments under &#x60;isolated-margin&#x60; trade mode and buy/sell position mode at contract level.    11. Set leverage for &#x60;SWAP&#x60; instruments under &#x60;isolated-margin&#x60; trade mode and long/short position mode at contract and position side level.       Note that the request parameter &#x60;posSide&#x60; is only required when margin mode is isolated in long/short position mode for FUTURES/SWAP instruments (see scenario 8 and 11 above).    Please refer to the request examples on the right for each case.     
[**CreateAccountSetPositionModeV5**](TradingAccountAPI.md#CreateAccountSetPositionModeV5) | **Post** /api/v5/account/set-position-mode | Spot and futures mode and Multi-currency mode: &#x60;FUTURES&#x60; and &#x60;SWAP&#x60; support both &#x60;long/short&#x60; mode and &#x60;net&#x60; mode. In &#x60;net&#x60; mode, users can only have positions in one direction; In &#x60;long/short&#x60; mode, users can hold positions in long and short directions.   Portfolio margin mode: &#x60;FUTURES&#x60; and &#x60;SWAP&#x60; only support &#x60;net&#x60; mode  
[**CreateAccountSetRiskOffsetAmtV5**](TradingAccountAPI.md#CreateAccountSetRiskOffsetAmtV5) | **Post** /api/v5/account/set-riskOffset-amt | Set risk offset amount. This does not represent the actual spot risk offset amount. Only applicable to Portfolio Margin Mode.  
[**CreateAccountSpotManualBorrowRepayV5**](TradingAccountAPI.md#CreateAccountSpotManualBorrowRepayV5) | **Post** /api/v5/account/spot-manual-borrow-repay | Only applicable to &#x60;Spot mode&#x60; (enabled borrowing)  
[**GetAccountAccountPositionRiskV5**](TradingAccountAPI.md#GetAccountAccountPositionRiskV5) | **Get** /api/v5/account/account-position-risk | Get account and position risk  
[**GetAccountAdjustLeverageInfoV5**](TradingAccountAPI.md#GetAccountAdjustLeverageInfoV5) | **Get** /api/v5/account/adjust-leverage-info | 
[**GetAccountBalanceV5**](TradingAccountAPI.md#GetAccountBalanceV5) | **Get** /api/v5/account/balance | Retrieve a list of assets (with non-zero balance), remaining balance, and available amount in the trading account.  
[**GetAccountBillsArchiveV5**](TradingAccountAPI.md#GetAccountBillsArchiveV5) | **Get** /api/v5/account/bills-archive | Retrieve the account’s bills. The bill refers to all transaction records that result in changing the balance of an account. Pagination is supported, and the response is sorted with most recent first. This endpoint can retrieve data from the last 1 year since July 1, 2024.  
[**GetAccountBillsHistoryArchiveV5**](TradingAccountAPI.md#GetAccountBillsHistoryArchiveV5) | **Get** /api/v5/account/bills-history-archive | Apply for bill data since 1 February, 2021 except for the current quarter.  
[**GetAccountBillsV5**](TradingAccountAPI.md#GetAccountBillsV5) | **Get** /api/v5/account/bills | Retrieve the bills of the account. The bill refers to all transaction records that result in changing the balance of an account. Pagination is supported, and the response is sorted with the most recent first. This endpoint can retrieve data from the last 7 days.  
[**GetAccountCollateralAssetsV5**](TradingAccountAPI.md#GetAccountCollateralAssetsV5) | **Get** /api/v5/account/collateral-assets | 
[**GetAccountConfigV5**](TradingAccountAPI.md#GetAccountConfigV5) | **Get** /api/v5/account/config | Retrieve current account configuration.  
[**GetAccountGreeksV5**](TradingAccountAPI.md#GetAccountGreeksV5) | **Get** /api/v5/account/greeks | Retrieve a greeks list of all assets in the account.  
[**GetAccountInstrumentsV5**](TradingAccountAPI.md#GetAccountInstrumentsV5) | **Get** /api/v5/account/instruments | Retrieve available instruments info of current account.  
[**GetAccountInterestAccruedV5**](TradingAccountAPI.md#GetAccountInterestAccruedV5) | **Get** /api/v5/account/interest-accrued | Get interest accrued data. Only data within the last one year can be obtained.  
[**GetAccountInterestLimitsV5**](TradingAccountAPI.md#GetAccountInterestLimitsV5) | **Get** /api/v5/account/interest-limits | 
[**GetAccountInterestRateV5**](TradingAccountAPI.md#GetAccountInterestRateV5) | **Get** /api/v5/account/interest-rate | Get the user&#39;s current leveraged currency borrowing market interest rate  
[**GetAccountLeverageInfoV5**](TradingAccountAPI.md#GetAccountLeverageInfoV5) | **Get** /api/v5/account/leverage-info | 
[**GetAccountMaxAvailSizeV5**](TradingAccountAPI.md#GetAccountMaxAvailSizeV5) | **Get** /api/v5/account/max-avail-size | Available balance for isolated margin positions and SPOT, available equity for cross margin positions.  
[**GetAccountMaxLoanV5**](TradingAccountAPI.md#GetAccountMaxLoanV5) | **Get** /api/v5/account/max-loan | 
[**GetAccountMaxSizeV5**](TradingAccountAPI.md#GetAccountMaxSizeV5) | **Get** /api/v5/account/max-size | The maximum quantity to buy or sell. It corresponds to the \&quot;sz\&quot; from placement.  
[**GetAccountMaxWithdrawalV5**](TradingAccountAPI.md#GetAccountMaxWithdrawalV5) | **Get** /api/v5/account/max-withdrawal | Retrieve the maximum transferable amount from trading account to funding account. If no currency is specified, the transferable amount of all owned currencies will be returned.  
[**GetAccountMmpConfigV5**](TradingAccountAPI.md#GetAccountMmpConfigV5) | **Get** /api/v5/account/mmp-config | This endpoint is used to get MMP configure information    Only applicable to Option in Portfolio Margin mode, and MMP privilege is required.  
[**GetAccountPositionTiersV5**](TradingAccountAPI.md#GetAccountPositionTiersV5) | **Get** /api/v5/account/position-tiers | Retrieve cross position limitation of SWAP/FUTURES/OPTION under Portfolio margin mode.  
[**GetAccountPositionsHistoryV5**](TradingAccountAPI.md#GetAccountPositionsHistoryV5) | **Get** /api/v5/account/positions-history | Retrieve the updated position data for the last 3 months. Return in reverse chronological order using utime. Getting positions history is supported under Portfolio margin mode since .  
[**GetAccountPositionsV5**](TradingAccountAPI.md#GetAccountPositionsV5) | **Get** /api/v5/account/positions | Retrieve information on your positions. When the account is in &#x60;net&#x60; mode, &#x60;net&#x60; positions will be displayed, and when the account is in &#x60;long/short&#x60; mode, &#x60;long&#x60; or &#x60;short&#x60; positions will be displayed. Return in reverse chronological order using ctime.  
[**GetAccountQuickMarginBorrowRepayHistoryV5**](TradingAccountAPI.md#GetAccountQuickMarginBorrowRepayHistoryV5) | **Get** /api/v5/account/quick-margin-borrow-repay-history | Get record in the past 3 months.  
[**GetAccountRiskStateV5**](TradingAccountAPI.md#GetAccountRiskStateV5) | **Get** /api/v5/account/risk-state | Only applicable to Portfolio margin account  
[**GetAccountSetAccountSwitchPrecheckV5**](TradingAccountAPI.md#GetAccountSetAccountSwitchPrecheckV5) | **Get** /api/v5/account/set-account-switch-precheck | Retrieve precheck information for account mode switching.  
[**GetAccountSpotBorrowRepayHistoryV5**](TradingAccountAPI.md#GetAccountSpotBorrowRepayHistoryV5) | **Get** /api/v5/account/spot-borrow-repay-history | Retrieve the borrow/repay history under &#x60;Spot mode&#x60;  
[**GetAccountTradeFeeV5**](TradingAccountAPI.md#GetAccountTradeFeeV5) | **Get** /api/v5/account/trade-fee | 



## CreateAccountAccountLevelSwitchPresetV5

> CreateAccountAccountLevelSwitchPresetV5Resp CreateAccountAccountLevelSwitchPresetV5(ctx).CreateAccountAccountLevelSwitchPresetV5Req(createAccountAccountLevelSwitchPresetV5Req).Execute()

Pre-set the required information for account mode switching. When switching from `Portfolio margin mode` back to `Spot and futures mode` / `Multi-currency margin mode`, and if there are existing cross-margin contract positions, it is mandatory to pre-set leverage.  If the user does not follow the required settings, they will receive an error message during the pre-check or when setting the account mode.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountAccountLevelSwitchPresetV5Req := *openapiclient.NewCreateAccountAccountLevelSwitchPresetV5Req("AcctLv_example") // CreateAccountAccountLevelSwitchPresetV5Req | The request body for CreateAccountAccountLevelSwitchPresetV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountAccountLevelSwitchPresetV5(context.Background()).CreateAccountAccountLevelSwitchPresetV5Req(createAccountAccountLevelSwitchPresetV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountAccountLevelSwitchPresetV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountAccountLevelSwitchPresetV5`: CreateAccountAccountLevelSwitchPresetV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountAccountLevelSwitchPresetV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountAccountLevelSwitchPresetV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountAccountLevelSwitchPresetV5Req** | [**CreateAccountAccountLevelSwitchPresetV5Req**](CreateAccountAccountLevelSwitchPresetV5Req.md) | The request body for CreateAccountAccountLevelSwitchPresetV5 | 

### Return type

[**CreateAccountAccountLevelSwitchPresetV5Resp**](CreateAccountAccountLevelSwitchPresetV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountActivateOptionV5

> CreateAccountActivateOptionV5Resp CreateAccountActivateOptionV5(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountActivateOptionV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountActivateOptionV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountActivateOptionV5`: CreateAccountActivateOptionV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountActivateOptionV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountActivateOptionV5Request struct via the builder pattern


### Return type

[**CreateAccountActivateOptionV5Resp**](CreateAccountActivateOptionV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountBillsHistoryArchiveV5

> CreateAccountBillsHistoryArchiveV5Resp CreateAccountBillsHistoryArchiveV5(ctx).CreateAccountBillsHistoryArchiveV5Req(createAccountBillsHistoryArchiveV5Req).Execute()

Apply for bill data since 1 February, 2021 except for the current quarter.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountBillsHistoryArchiveV5Req := *openapiclient.NewCreateAccountBillsHistoryArchiveV5Req("Quarter_example", "Year_example") // CreateAccountBillsHistoryArchiveV5Req | The request body for CreateAccountBillsHistoryArchiveV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountBillsHistoryArchiveV5(context.Background()).CreateAccountBillsHistoryArchiveV5Req(createAccountBillsHistoryArchiveV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountBillsHistoryArchiveV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountBillsHistoryArchiveV5`: CreateAccountBillsHistoryArchiveV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountBillsHistoryArchiveV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountBillsHistoryArchiveV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountBillsHistoryArchiveV5Req** | [**CreateAccountBillsHistoryArchiveV5Req**](CreateAccountBillsHistoryArchiveV5Req.md) | The request body for CreateAccountBillsHistoryArchiveV5 | 

### Return type

[**CreateAccountBillsHistoryArchiveV5Resp**](CreateAccountBillsHistoryArchiveV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountMmpConfigV5

> CreateAccountMmpConfigV5Resp CreateAccountMmpConfigV5(ctx).CreateAccountMmpConfigV5Req(createAccountMmpConfigV5Req).Execute()

This endpoint is used to set MMP configure    Only applicable to Option in Portfolio Margin mode, and MMP privilege is required.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountMmpConfigV5Req := *openapiclient.NewCreateAccountMmpConfigV5Req("FrozenInterval_example", "InstFamily_example", "QtyLimit_example", "TimeInterval_example") // CreateAccountMmpConfigV5Req | The request body for CreateAccountMmpConfigV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountMmpConfigV5(context.Background()).CreateAccountMmpConfigV5Req(createAccountMmpConfigV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountMmpConfigV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountMmpConfigV5`: CreateAccountMmpConfigV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountMmpConfigV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountMmpConfigV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountMmpConfigV5Req** | [**CreateAccountMmpConfigV5Req**](CreateAccountMmpConfigV5Req.md) | The request body for CreateAccountMmpConfigV5 | 

### Return type

[**CreateAccountMmpConfigV5Resp**](CreateAccountMmpConfigV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountMmpResetV5

> CreateAccountMmpResetV5Resp CreateAccountMmpResetV5(ctx).CreateAccountMmpResetV5Req(createAccountMmpResetV5Req).Execute()

You can unfreeze by this endpoint once MMP is triggered.    Only applicable to Option in Portfolio Margin mode, and MMP privilege is required.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountMmpResetV5Req := *openapiclient.NewCreateAccountMmpResetV5Req("InstFamily_example") // CreateAccountMmpResetV5Req | The request body for CreateAccountMmpResetV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountMmpResetV5(context.Background()).CreateAccountMmpResetV5Req(createAccountMmpResetV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountMmpResetV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountMmpResetV5`: CreateAccountMmpResetV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountMmpResetV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountMmpResetV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountMmpResetV5Req** | [**CreateAccountMmpResetV5Req**](CreateAccountMmpResetV5Req.md) | The request body for CreateAccountMmpResetV5 | 

### Return type

[**CreateAccountMmpResetV5Resp**](CreateAccountMmpResetV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountPositionBuilderV5

> CreateAccountPositionBuilderV5Resp CreateAccountPositionBuilderV5(ctx).CreateAccountPositionBuilderV5Req(createAccountPositionBuilderV5Req).Execute()

Calculates portfolio margin information for virtual position/assets or current position of the user.   You can add up to 200 virtual positions and 200 virtual assets in one request.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountPositionBuilderV5Req := *openapiclient.NewCreateAccountPositionBuilderV5Req() // CreateAccountPositionBuilderV5Req | The request body for CreateAccountPositionBuilderV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountPositionBuilderV5(context.Background()).CreateAccountPositionBuilderV5Req(createAccountPositionBuilderV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountPositionBuilderV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountPositionBuilderV5`: CreateAccountPositionBuilderV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountPositionBuilderV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountPositionBuilderV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountPositionBuilderV5Req** | [**CreateAccountPositionBuilderV5Req**](CreateAccountPositionBuilderV5Req.md) | The request body for CreateAccountPositionBuilderV5 | 

### Return type

[**CreateAccountPositionBuilderV5Resp**](CreateAccountPositionBuilderV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountPositionMarginBalanceV5

> CreateAccountPositionMarginBalanceV5Resp CreateAccountPositionMarginBalanceV5(ctx).CreateAccountPositionMarginBalanceV5Req(createAccountPositionMarginBalanceV5Req).Execute()

Increase or decrease the margin of the isolated position. Margin reduction may result in the change of the actual leverage.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountPositionMarginBalanceV5Req := *openapiclient.NewCreateAccountPositionMarginBalanceV5Req("Amt_example", "InstId_example", "PosSide_example", "Type_example") // CreateAccountPositionMarginBalanceV5Req | The request body for CreateAccountPositionMarginBalanceV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountPositionMarginBalanceV5(context.Background()).CreateAccountPositionMarginBalanceV5Req(createAccountPositionMarginBalanceV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountPositionMarginBalanceV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountPositionMarginBalanceV5`: CreateAccountPositionMarginBalanceV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountPositionMarginBalanceV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountPositionMarginBalanceV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountPositionMarginBalanceV5Req** | [**CreateAccountPositionMarginBalanceV5Req**](CreateAccountPositionMarginBalanceV5Req.md) | The request body for CreateAccountPositionMarginBalanceV5 | 

### Return type

[**CreateAccountPositionMarginBalanceV5Resp**](CreateAccountPositionMarginBalanceV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountQuickMarginBorrowRepayV5

> CreateAccountQuickMarginBorrowRepayV5Resp CreateAccountQuickMarginBorrowRepayV5(ctx).CreateAccountQuickMarginBorrowRepayV5Req(createAccountQuickMarginBorrowRepayV5Req).Execute()

Please note that this endpoint will be deprecated soon.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountQuickMarginBorrowRepayV5Req := *openapiclient.NewCreateAccountQuickMarginBorrowRepayV5Req("Amt_example", "Ccy_example", "InstId_example", "Side_example") // CreateAccountQuickMarginBorrowRepayV5Req | The request body for CreateAccountQuickMarginBorrowRepayV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountQuickMarginBorrowRepayV5(context.Background()).CreateAccountQuickMarginBorrowRepayV5Req(createAccountQuickMarginBorrowRepayV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountQuickMarginBorrowRepayV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountQuickMarginBorrowRepayV5`: CreateAccountQuickMarginBorrowRepayV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountQuickMarginBorrowRepayV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountQuickMarginBorrowRepayV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountQuickMarginBorrowRepayV5Req** | [**CreateAccountQuickMarginBorrowRepayV5Req**](CreateAccountQuickMarginBorrowRepayV5Req.md) | The request body for CreateAccountQuickMarginBorrowRepayV5 | 

### Return type

[**CreateAccountQuickMarginBorrowRepayV5Resp**](CreateAccountQuickMarginBorrowRepayV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountSetAccountLevelV5

> CreateAccountSetAccountLevelV5Resp CreateAccountSetAccountLevelV5(ctx).CreateAccountSetAccountLevelV5Req(createAccountSetAccountLevelV5Req).Execute()

You need to set on the Web/App for the first set of every account mode. If users plan to switch account modes while holding positions, they should first call the preset endpoint to conduct necessary settings, then call the precheck endpoint to get unmatched information, margin check, and other related information, and finally call the account mode switch endpoint to switch account modes.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountSetAccountLevelV5Req := *openapiclient.NewCreateAccountSetAccountLevelV5Req("AcctLv_example") // CreateAccountSetAccountLevelV5Req | The request body for CreateAccountSetAccountLevelV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountSetAccountLevelV5(context.Background()).CreateAccountSetAccountLevelV5Req(createAccountSetAccountLevelV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountSetAccountLevelV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountSetAccountLevelV5`: CreateAccountSetAccountLevelV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountSetAccountLevelV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountSetAccountLevelV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountSetAccountLevelV5Req** | [**CreateAccountSetAccountLevelV5Req**](CreateAccountSetAccountLevelV5Req.md) | The request body for CreateAccountSetAccountLevelV5 | 

### Return type

[**CreateAccountSetAccountLevelV5Resp**](CreateAccountSetAccountLevelV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountSetAutoLoanV5

> CreateAccountSetAutoLoanV5Resp CreateAccountSetAutoLoanV5(ctx).CreateAccountSetAutoLoanV5Req(createAccountSetAutoLoanV5Req).Execute()

Only applicable to `Multi-currency margin` and `Portfolio margin`  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountSetAutoLoanV5Req := *openapiclient.NewCreateAccountSetAutoLoanV5Req() // CreateAccountSetAutoLoanV5Req | The request body for CreateAccountSetAutoLoanV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountSetAutoLoanV5(context.Background()).CreateAccountSetAutoLoanV5Req(createAccountSetAutoLoanV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountSetAutoLoanV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountSetAutoLoanV5`: CreateAccountSetAutoLoanV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountSetAutoLoanV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountSetAutoLoanV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountSetAutoLoanV5Req** | [**CreateAccountSetAutoLoanV5Req**](CreateAccountSetAutoLoanV5Req.md) | The request body for CreateAccountSetAutoLoanV5 | 

### Return type

[**CreateAccountSetAutoLoanV5Resp**](CreateAccountSetAutoLoanV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountSetAutoRepayV5

> CreateAccountSetAutoRepayV5Resp CreateAccountSetAutoRepayV5(ctx).CreateAccountSetAutoRepayV5Req(createAccountSetAutoRepayV5Req).Execute()

Only applicable to `Spot mode` (enabled borrowing)  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountSetAutoRepayV5Req := *openapiclient.NewCreateAccountSetAutoRepayV5Req(false) // CreateAccountSetAutoRepayV5Req | The request body for CreateAccountSetAutoRepayV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountSetAutoRepayV5(context.Background()).CreateAccountSetAutoRepayV5Req(createAccountSetAutoRepayV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountSetAutoRepayV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountSetAutoRepayV5`: CreateAccountSetAutoRepayV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountSetAutoRepayV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountSetAutoRepayV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountSetAutoRepayV5Req** | [**CreateAccountSetAutoRepayV5Req**](CreateAccountSetAutoRepayV5Req.md) | The request body for CreateAccountSetAutoRepayV5 | 

### Return type

[**CreateAccountSetAutoRepayV5Resp**](CreateAccountSetAutoRepayV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountSetCollateralAssetsV5

> CreateAccountSetCollateralAssetsV5Resp CreateAccountSetCollateralAssetsV5(ctx).CreateAccountSetCollateralAssetsV5Req(createAccountSetCollateralAssetsV5Req).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountSetCollateralAssetsV5Req := *openapiclient.NewCreateAccountSetCollateralAssetsV5Req() // CreateAccountSetCollateralAssetsV5Req | The request body for CreateAccountSetCollateralAssetsV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountSetCollateralAssetsV5(context.Background()).CreateAccountSetCollateralAssetsV5Req(createAccountSetCollateralAssetsV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountSetCollateralAssetsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountSetCollateralAssetsV5`: CreateAccountSetCollateralAssetsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountSetCollateralAssetsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountSetCollateralAssetsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountSetCollateralAssetsV5Req** | [**CreateAccountSetCollateralAssetsV5Req**](CreateAccountSetCollateralAssetsV5Req.md) | The request body for CreateAccountSetCollateralAssetsV5 | 

### Return type

[**CreateAccountSetCollateralAssetsV5Resp**](CreateAccountSetCollateralAssetsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountSetGreeksV5

> CreateAccountSetGreeksV5Resp CreateAccountSetGreeksV5(ctx).CreateAccountSetGreeksV5Req(createAccountSetGreeksV5Req).Execute()

Set the display type of Greeks.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountSetGreeksV5Req := *openapiclient.NewCreateAccountSetGreeksV5Req("GreeksType_example") // CreateAccountSetGreeksV5Req | The request body for CreateAccountSetGreeksV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountSetGreeksV5(context.Background()).CreateAccountSetGreeksV5Req(createAccountSetGreeksV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountSetGreeksV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountSetGreeksV5`: CreateAccountSetGreeksV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountSetGreeksV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountSetGreeksV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountSetGreeksV5Req** | [**CreateAccountSetGreeksV5Req**](CreateAccountSetGreeksV5Req.md) | The request body for CreateAccountSetGreeksV5 | 

### Return type

[**CreateAccountSetGreeksV5Resp**](CreateAccountSetGreeksV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountSetIsolatedModeV5

> CreateAccountSetIsolatedModeV5Resp CreateAccountSetIsolatedModeV5(ctx).CreateAccountSetIsolatedModeV5Req(createAccountSetIsolatedModeV5Req).Execute()

You can set the currency margin and futures/perpetual Isolated margin trading mode  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountSetIsolatedModeV5Req := *openapiclient.NewCreateAccountSetIsolatedModeV5Req("IsoMode_example", "Type_example") // CreateAccountSetIsolatedModeV5Req | The request body for CreateAccountSetIsolatedModeV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountSetIsolatedModeV5(context.Background()).CreateAccountSetIsolatedModeV5Req(createAccountSetIsolatedModeV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountSetIsolatedModeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountSetIsolatedModeV5`: CreateAccountSetIsolatedModeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountSetIsolatedModeV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountSetIsolatedModeV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountSetIsolatedModeV5Req** | [**CreateAccountSetIsolatedModeV5Req**](CreateAccountSetIsolatedModeV5Req.md) | The request body for CreateAccountSetIsolatedModeV5 | 

### Return type

[**CreateAccountSetIsolatedModeV5Resp**](CreateAccountSetIsolatedModeV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountSetLeverageV5

> CreateAccountSetLeverageV5Resp CreateAccountSetLeverageV5(ctx).CreateAccountSetLeverageV5Req(createAccountSetLeverageV5Req).Execute()

  There are 10 different scenarios for leverage setting:      1. Set leverage for `MARGIN` instruments under `isolated-margin` trade mode at pairs level.    2. Set leverage for `MARGIN` instruments under `cross-margin` trade mode and Spot mode (enabled borrow) at currency level.    3. Set leverage for `MARGIN` instruments under `cross-margin` trade mode and Spot and futures mode account mode at pairs level.    4. Set leverage for `MARGIN` instruments under `cross-margin` trade mode and Multi-currency margin at currency level.    5. Set leverage for `MARGIN` instruments under `cross-margin` trade mode and Portfolio margin at currency level.    6. Set leverage for `FUTURES` instruments under `cross-margin` trade mode at underlying level.    7. Set leverage for `FUTURES` instruments under `isolated-margin` trade mode and buy/sell position mode at contract level.    8. Set leverage for `FUTURES` instruments under `isolated-margin` trade mode and long/short position mode at contract and position side level.    9. Set leverage for `SWAP` instruments under `cross-margin` trade at contract level.    10. Set leverage for `SWAP` instruments under `isolated-margin` trade mode and buy/sell position mode at contract level.    11. Set leverage for `SWAP` instruments under `isolated-margin` trade mode and long/short position mode at contract and position side level.       Note that the request parameter `posSide` is only required when margin mode is isolated in long/short position mode for FUTURES/SWAP instruments (see scenario 8 and 11 above).    Please refer to the request examples on the right for each case.     



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountSetLeverageV5Req := *openapiclient.NewCreateAccountSetLeverageV5Req("Lever_example", "MgnMode_example") // CreateAccountSetLeverageV5Req | The request body for CreateAccountSetLeverageV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountSetLeverageV5(context.Background()).CreateAccountSetLeverageV5Req(createAccountSetLeverageV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountSetLeverageV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountSetLeverageV5`: CreateAccountSetLeverageV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountSetLeverageV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountSetLeverageV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountSetLeverageV5Req** | [**CreateAccountSetLeverageV5Req**](CreateAccountSetLeverageV5Req.md) | The request body for CreateAccountSetLeverageV5 | 

### Return type

[**CreateAccountSetLeverageV5Resp**](CreateAccountSetLeverageV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountSetPositionModeV5

> CreateAccountSetPositionModeV5Resp CreateAccountSetPositionModeV5(ctx).CreateAccountSetPositionModeV5Req(createAccountSetPositionModeV5Req).Execute()

Spot and futures mode and Multi-currency mode: `FUTURES` and `SWAP` support both `long/short` mode and `net` mode. In `net` mode, users can only have positions in one direction; In `long/short` mode, users can hold positions in long and short directions.   Portfolio margin mode: `FUTURES` and `SWAP` only support `net` mode  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountSetPositionModeV5Req := *openapiclient.NewCreateAccountSetPositionModeV5Req("PosMode_example") // CreateAccountSetPositionModeV5Req | The request body for CreateAccountSetPositionModeV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountSetPositionModeV5(context.Background()).CreateAccountSetPositionModeV5Req(createAccountSetPositionModeV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountSetPositionModeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountSetPositionModeV5`: CreateAccountSetPositionModeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountSetPositionModeV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountSetPositionModeV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountSetPositionModeV5Req** | [**CreateAccountSetPositionModeV5Req**](CreateAccountSetPositionModeV5Req.md) | The request body for CreateAccountSetPositionModeV5 | 

### Return type

[**CreateAccountSetPositionModeV5Resp**](CreateAccountSetPositionModeV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountSetRiskOffsetAmtV5

> CreateAccountSetRiskOffsetAmtV5Resp CreateAccountSetRiskOffsetAmtV5(ctx).CreateAccountSetRiskOffsetAmtV5Req(createAccountSetRiskOffsetAmtV5Req).Execute()

Set risk offset amount. This does not represent the actual spot risk offset amount. Only applicable to Portfolio Margin Mode.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountSetRiskOffsetAmtV5Req := *openapiclient.NewCreateAccountSetRiskOffsetAmtV5Req("Ccy_example", "ClSpotInUseAmt_example") // CreateAccountSetRiskOffsetAmtV5Req | The request body for CreateAccountSetRiskOffsetAmtV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountSetRiskOffsetAmtV5(context.Background()).CreateAccountSetRiskOffsetAmtV5Req(createAccountSetRiskOffsetAmtV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountSetRiskOffsetAmtV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountSetRiskOffsetAmtV5`: CreateAccountSetRiskOffsetAmtV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountSetRiskOffsetAmtV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountSetRiskOffsetAmtV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountSetRiskOffsetAmtV5Req** | [**CreateAccountSetRiskOffsetAmtV5Req**](CreateAccountSetRiskOffsetAmtV5Req.md) | The request body for CreateAccountSetRiskOffsetAmtV5 | 

### Return type

[**CreateAccountSetRiskOffsetAmtV5Resp**](CreateAccountSetRiskOffsetAmtV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountSpotManualBorrowRepayV5

> CreateAccountSpotManualBorrowRepayV5Resp CreateAccountSpotManualBorrowRepayV5(ctx).CreateAccountSpotManualBorrowRepayV5Req(createAccountSpotManualBorrowRepayV5Req).Execute()

Only applicable to `Spot mode` (enabled borrowing)  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	createAccountSpotManualBorrowRepayV5Req := *openapiclient.NewCreateAccountSpotManualBorrowRepayV5Req("Amt_example", "Ccy_example", "Side_example") // CreateAccountSpotManualBorrowRepayV5Req | The request body for CreateAccountSpotManualBorrowRepayV5

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.CreateAccountSpotManualBorrowRepayV5(context.Background()).CreateAccountSpotManualBorrowRepayV5Req(createAccountSpotManualBorrowRepayV5Req).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.CreateAccountSpotManualBorrowRepayV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountSpotManualBorrowRepayV5`: CreateAccountSpotManualBorrowRepayV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.CreateAccountSpotManualBorrowRepayV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountSpotManualBorrowRepayV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountSpotManualBorrowRepayV5Req** | [**CreateAccountSpotManualBorrowRepayV5Req**](CreateAccountSpotManualBorrowRepayV5Req.md) | The request body for CreateAccountSpotManualBorrowRepayV5 | 

### Return type

[**CreateAccountSpotManualBorrowRepayV5Resp**](CreateAccountSpotManualBorrowRepayV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAccountPositionRiskV5

> GetAccountAccountPositionRiskV5Resp GetAccountAccountPositionRiskV5(ctx).InstType(instType).Execute()

Get account and position risk  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	instType := "instType_example" // string | Instrument type  `MARGIN`  `SWAP`  `FUTURES`  `OPTION` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountAccountPositionRiskV5(context.Background()).InstType(instType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountAccountPositionRiskV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAccountPositionRiskV5`: GetAccountAccountPositionRiskV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountAccountPositionRiskV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAccountPositionRiskV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]

### Return type

[**GetAccountAccountPositionRiskV5Resp**](GetAccountAccountPositionRiskV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAdjustLeverageInfoV5

> GetAccountAdjustLeverageInfoV5Resp GetAccountAdjustLeverageInfoV5(ctx).InstType(instType).MgnMode(mgnMode).Lever(lever).InstId(instId).Ccy(ccy).PosSide(posSide).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	instType := "instType_example" // string | Instrument type  `MARGIN`  `SWAP`  `FUTURES` (default to "")
	mgnMode := "mgnMode_example" // string | Margin mode  `isolated`  `cross` (default to "")
	lever := "lever_example" // string | Leverage (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. BTC-USDT  It is required for these scenarioes: `SWAP` and `FUTURES`, Margin isolation, Margin cross in `Spot and futures mode`. (optional) (default to "")
	ccy := "ccy_example" // string | Currency used for margin, e.g. BTC  It is required for isolated margin and cross margin in `Spot and futures mode`, `Multi-currency margin` and `Portfolio margin` (optional) (default to "")
	posSide := "posSide_example" // string | posSide  `net`: The default value  `long`  `short` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountAdjustLeverageInfoV5(context.Background()).InstType(instType).MgnMode(mgnMode).Lever(lever).InstId(instId).Ccy(ccy).PosSide(posSide).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountAdjustLeverageInfoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAdjustLeverageInfoV5`: GetAccountAdjustLeverageInfoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountAdjustLeverageInfoV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAdjustLeverageInfoV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60; | [default to &quot;&quot;]
 **mgnMode** | **string** | Margin mode  &#x60;isolated&#x60;  &#x60;cross&#x60; | [default to &quot;&quot;]
 **lever** | **string** | Leverage | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. BTC-USDT  It is required for these scenarioes: &#x60;SWAP&#x60; and &#x60;FUTURES&#x60;, Margin isolation, Margin cross in &#x60;Spot and futures mode&#x60;. | [default to &quot;&quot;]
 **ccy** | **string** | Currency used for margin, e.g. BTC  It is required for isolated margin and cross margin in &#x60;Spot and futures mode&#x60;, &#x60;Multi-currency margin&#x60; and &#x60;Portfolio margin&#x60; | [default to &quot;&quot;]
 **posSide** | **string** | posSide  &#x60;net&#x60;: The default value  &#x60;long&#x60;  &#x60;short&#x60; | [default to &quot;&quot;]

### Return type

[**GetAccountAdjustLeverageInfoV5Resp**](GetAccountAdjustLeverageInfoV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountBalanceV5

> GetAccountBalanceV5Resp GetAccountBalanceV5(ctx).Ccy(ccy).Execute()

Retrieve a list of assets (with non-zero balance), remaining balance, and available amount in the trading account.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	ccy := "ccy_example" // string | Single currency or multiple currencies (no more than 20) separated with comma, e.g. `BTC` or `BTC,ETH`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountBalanceV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountBalanceV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountBalanceV5`: GetAccountBalanceV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountBalanceV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountBalanceV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Single currency or multiple currencies (no more than 20) separated with comma, e.g. &#x60;BTC&#x60; or &#x60;BTC,ETH&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAccountBalanceV5Resp**](GetAccountBalanceV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountBillsArchiveV5

> GetAccountBillsArchiveV5Resp GetAccountBillsArchiveV5(ctx).InstType(instType).InstId(instId).Ccy(ccy).MgnMode(mgnMode).CtType(ctType).Type_(type_).SubType(subType).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()

Retrieve the account’s bills. The bill refers to all transaction records that result in changing the balance of an account. Pagination is supported, and the response is sorted with most recent first. This endpoint can retrieve data from the last 1 year since July 1, 2024.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	instType := "instType_example" // string | Instrument type  `SPOT`  `MARGIN`  `SWAP`  `FUTURES`  `OPTION` (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT` (optional) (default to "")
	ccy := "ccy_example" // string | Bill currency (optional) (default to "")
	mgnMode := "mgnMode_example" // string | Margin mode  `isolated`  `cross` (optional) (default to "")
	ctType := "ctType_example" // string | Contract type  `linear`  `inverse`  Only applicable to `FUTURES`/`SWAP` (optional) (default to "")
	type_ := "type__example" // string | Bill type  `1`: Transfer  `2`: Trade  `3`: Delivery  `4`: Forced repayment  `5`: Liquidation  `6`: Margin transfer  `7`: Interest deduction  `8`: Funding fee  `9`: ADL  `10`: Clawback  `11`: System token conversion  `12`: Strategy transfer  `13`: DDH  `14`: Block trade  `15`: Quick Margin  `16`: Borrowing  `22`: Repay  `24`: Spread trading  `26`: Structured products  `27`: Convert  `28`: Easy convert  `29`: One-click repay  `30`: Simple trade  `33`: Loans  `34`: Settlement  `250`: Copy trader profit sharing expenses  `251`: Copy trader profit sharing refund (optional) (default to "")
	subType := "subType_example" // string | Bill subtype  `1`: Buy  `2`: Sell  `3`: Open long  `4`: Open short  `5`: Close long  `6`: Close short  `9`: Interest deduction for Market loans  `11`: Transfer in  `12`: Transfer out  `14`: Interest deduction for VIP loans  `160`: Manual margin increase  `161`: Manual margin decrease  `162`: Auto margin increase  `114`: Forced repayment buy  `115`: Forced repayment sell  `118`: System token conversion transfer in  `119`: System token conversion transfer out  `100`: Partial liquidation close long  `101`: Partial liquidation close short  `102`: Partial liquidation buy  `103`: Partial liquidation sell  `104`: Liquidation long  `105`: Liquidation short  `106`: Liquidation buy  `107`: Liquidation sell  `108`: Clawback  `110`: Liquidation transfer in  `111`: Liquidation transfer out  `125`: ADL close long  `126`: ADL close short  `127`: ADL buy  `128`: ADL sell  `131`: ddh buy  `132`: ddh sell  `170`: Exercised(ITM buy side)  `171`: Counterparty exercised(ITM sell side)  `172`: Expired(Non-ITM buy and sell side)  `112`: Delivery long  `113`: Delivery short  `117`: Delivery/Exercise clawback  `173`: Funding fee expense  `174`: Funding fee income  `200`:System transfer in  `201`: Manually transfer in  `202`: System transfer out  `203`: Manually transfer out  `204`: block trade buy  `205`: block trade sell  `206`: block trade open long  `207`: block trade open short  `208`: block trade close long  `209`: block trade close short  `210`: Manual Borrowing of quick margin  `211`: Manual Repayment of quick margin  `212`: Auto borrow of quick margin  `213`: Auto repay of quick margin  `220`: Transfer in when using USDT to buy OPTION  `221`: Transfer out when using USDT to buy OPTION  `16`: Repay forcibly  `17`: Repay interest by borrowing forcibly  `224`: Repayment transfer in  `225`: Repayment transfer out  `236`: Easy convert in  `237`: Easy convert out  `250`: Profit sharing expenses  `251`: Profit sharing refund  `280`: SPOT profit sharing expenses  `281`: SPOT profit sharing refund  `270`: Spread trading buy  `271`: Spread trading sell  `272`: Spread trading open long  `273`: Spread trading open short  `274`: Spread trading close long  `275`: Spread trading close short  `280`: SPOT profit sharing expenses  `281`: SPOT profit sharing refund   `284`: Copy trade automatic transfer in  `285`: Copy trade manual transfer in  `286`: Copy trade automatic transfer out  `287`: Copy trade manual transfer out  `290`: Crypto dust auto-transfer out  `293`: Fixed loan interest deduction  `294`: Fixed loan interest refund  `295`: Fixed loan overdue penalty  `296`: From structured order placements  `297`: To structured order placements  `298`: From structured settlements  `299`: To structured settlements  `306`: Manual borrow  `307`: Auto borrow  `308`: Manual repay  `309`: Auto repay  `312`: Auto offset  `318`: Convert in  `319`: Convert out  `320`: Simple buy  `321`: Simple sell  `332`: Margin transfer in isolated margin position   `333`: Margin transfer out isolated margin position  `334`: Margin loss when closing isolated margin position  `348`: [Credit line] Forced repayment  `350`: [Credit line] Forced repayment refund  `352`: [Credit line] Forced repayment penalty fee deduction  `353`: [Credit line] Forced repayment penalty fee (pending deduction)  `356`: [Credit line] Auto conversion (pending deduction)  `357`: [Credit line] Auto Conversion Transfer to Funding  `355`: Settlement PnL (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested bill ID. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested bill ID. (optional) (default to "")
	begin := "begin_example" // string | Filter with a begin timestamp  `ts`. Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	end := "end_example" // string | Filter with an end timestamp  `ts`. Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountBillsArchiveV5(context.Background()).InstType(instType).InstId(instId).Ccy(ccy).MgnMode(mgnMode).CtType(ctType).Type_(type_).SubType(subType).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountBillsArchiveV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountBillsArchiveV5`: GetAccountBillsArchiveV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountBillsArchiveV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountBillsArchiveV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **ccy** | **string** | Bill currency | [default to &quot;&quot;]
 **mgnMode** | **string** | Margin mode  &#x60;isolated&#x60;  &#x60;cross&#x60; | [default to &quot;&quot;]
 **ctType** | **string** | Contract type  &#x60;linear&#x60;  &#x60;inverse&#x60;  Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60; | [default to &quot;&quot;]
 **type_** | **string** | Bill type  &#x60;1&#x60;: Transfer  &#x60;2&#x60;: Trade  &#x60;3&#x60;: Delivery  &#x60;4&#x60;: Forced repayment  &#x60;5&#x60;: Liquidation  &#x60;6&#x60;: Margin transfer  &#x60;7&#x60;: Interest deduction  &#x60;8&#x60;: Funding fee  &#x60;9&#x60;: ADL  &#x60;10&#x60;: Clawback  &#x60;11&#x60;: System token conversion  &#x60;12&#x60;: Strategy transfer  &#x60;13&#x60;: DDH  &#x60;14&#x60;: Block trade  &#x60;15&#x60;: Quick Margin  &#x60;16&#x60;: Borrowing  &#x60;22&#x60;: Repay  &#x60;24&#x60;: Spread trading  &#x60;26&#x60;: Structured products  &#x60;27&#x60;: Convert  &#x60;28&#x60;: Easy convert  &#x60;29&#x60;: One-click repay  &#x60;30&#x60;: Simple trade  &#x60;33&#x60;: Loans  &#x60;34&#x60;: Settlement  &#x60;250&#x60;: Copy trader profit sharing expenses  &#x60;251&#x60;: Copy trader profit sharing refund | [default to &quot;&quot;]
 **subType** | **string** | Bill subtype  &#x60;1&#x60;: Buy  &#x60;2&#x60;: Sell  &#x60;3&#x60;: Open long  &#x60;4&#x60;: Open short  &#x60;5&#x60;: Close long  &#x60;6&#x60;: Close short  &#x60;9&#x60;: Interest deduction for Market loans  &#x60;11&#x60;: Transfer in  &#x60;12&#x60;: Transfer out  &#x60;14&#x60;: Interest deduction for VIP loans  &#x60;160&#x60;: Manual margin increase  &#x60;161&#x60;: Manual margin decrease  &#x60;162&#x60;: Auto margin increase  &#x60;114&#x60;: Forced repayment buy  &#x60;115&#x60;: Forced repayment sell  &#x60;118&#x60;: System token conversion transfer in  &#x60;119&#x60;: System token conversion transfer out  &#x60;100&#x60;: Partial liquidation close long  &#x60;101&#x60;: Partial liquidation close short  &#x60;102&#x60;: Partial liquidation buy  &#x60;103&#x60;: Partial liquidation sell  &#x60;104&#x60;: Liquidation long  &#x60;105&#x60;: Liquidation short  &#x60;106&#x60;: Liquidation buy  &#x60;107&#x60;: Liquidation sell  &#x60;108&#x60;: Clawback  &#x60;110&#x60;: Liquidation transfer in  &#x60;111&#x60;: Liquidation transfer out  &#x60;125&#x60;: ADL close long  &#x60;126&#x60;: ADL close short  &#x60;127&#x60;: ADL buy  &#x60;128&#x60;: ADL sell  &#x60;131&#x60;: ddh buy  &#x60;132&#x60;: ddh sell  &#x60;170&#x60;: Exercised(ITM buy side)  &#x60;171&#x60;: Counterparty exercised(ITM sell side)  &#x60;172&#x60;: Expired(Non-ITM buy and sell side)  &#x60;112&#x60;: Delivery long  &#x60;113&#x60;: Delivery short  &#x60;117&#x60;: Delivery/Exercise clawback  &#x60;173&#x60;: Funding fee expense  &#x60;174&#x60;: Funding fee income  &#x60;200&#x60;:System transfer in  &#x60;201&#x60;: Manually transfer in  &#x60;202&#x60;: System transfer out  &#x60;203&#x60;: Manually transfer out  &#x60;204&#x60;: block trade buy  &#x60;205&#x60;: block trade sell  &#x60;206&#x60;: block trade open long  &#x60;207&#x60;: block trade open short  &#x60;208&#x60;: block trade close long  &#x60;209&#x60;: block trade close short  &#x60;210&#x60;: Manual Borrowing of quick margin  &#x60;211&#x60;: Manual Repayment of quick margin  &#x60;212&#x60;: Auto borrow of quick margin  &#x60;213&#x60;: Auto repay of quick margin  &#x60;220&#x60;: Transfer in when using USDT to buy OPTION  &#x60;221&#x60;: Transfer out when using USDT to buy OPTION  &#x60;16&#x60;: Repay forcibly  &#x60;17&#x60;: Repay interest by borrowing forcibly  &#x60;224&#x60;: Repayment transfer in  &#x60;225&#x60;: Repayment transfer out  &#x60;236&#x60;: Easy convert in  &#x60;237&#x60;: Easy convert out  &#x60;250&#x60;: Profit sharing expenses  &#x60;251&#x60;: Profit sharing refund  &#x60;280&#x60;: SPOT profit sharing expenses  &#x60;281&#x60;: SPOT profit sharing refund  &#x60;270&#x60;: Spread trading buy  &#x60;271&#x60;: Spread trading sell  &#x60;272&#x60;: Spread trading open long  &#x60;273&#x60;: Spread trading open short  &#x60;274&#x60;: Spread trading close long  &#x60;275&#x60;: Spread trading close short  &#x60;280&#x60;: SPOT profit sharing expenses  &#x60;281&#x60;: SPOT profit sharing refund   &#x60;284&#x60;: Copy trade automatic transfer in  &#x60;285&#x60;: Copy trade manual transfer in  &#x60;286&#x60;: Copy trade automatic transfer out  &#x60;287&#x60;: Copy trade manual transfer out  &#x60;290&#x60;: Crypto dust auto-transfer out  &#x60;293&#x60;: Fixed loan interest deduction  &#x60;294&#x60;: Fixed loan interest refund  &#x60;295&#x60;: Fixed loan overdue penalty  &#x60;296&#x60;: From structured order placements  &#x60;297&#x60;: To structured order placements  &#x60;298&#x60;: From structured settlements  &#x60;299&#x60;: To structured settlements  &#x60;306&#x60;: Manual borrow  &#x60;307&#x60;: Auto borrow  &#x60;308&#x60;: Manual repay  &#x60;309&#x60;: Auto repay  &#x60;312&#x60;: Auto offset  &#x60;318&#x60;: Convert in  &#x60;319&#x60;: Convert out  &#x60;320&#x60;: Simple buy  &#x60;321&#x60;: Simple sell  &#x60;332&#x60;: Margin transfer in isolated margin position   &#x60;333&#x60;: Margin transfer out isolated margin position  &#x60;334&#x60;: Margin loss when closing isolated margin position  &#x60;348&#x60;: [Credit line] Forced repayment  &#x60;350&#x60;: [Credit line] Forced repayment refund  &#x60;352&#x60;: [Credit line] Forced repayment penalty fee deduction  &#x60;353&#x60;: [Credit line] Forced repayment penalty fee (pending deduction)  &#x60;356&#x60;: [Credit line] Auto conversion (pending deduction)  &#x60;357&#x60;: [Credit line] Auto Conversion Transfer to Funding  &#x60;355&#x60;: Settlement PnL | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested bill ID. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested bill ID. | [default to &quot;&quot;]
 **begin** | **string** | Filter with a begin timestamp  &#x60;ts&#x60;. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **end** | **string** | Filter with an end timestamp  &#x60;ts&#x60;. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAccountBillsArchiveV5Resp**](GetAccountBillsArchiveV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountBillsHistoryArchiveV5

> GetAccountBillsHistoryArchiveV5Resp GetAccountBillsHistoryArchiveV5(ctx).Year(year).Quarter(quarter).Execute()

Apply for bill data since 1 February, 2021 except for the current quarter.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	year := "year_example" // string | 4 digits year (default to "")
	quarter := "quarter_example" // string | Quarter, valid value is `Q1`, `Q2`, `Q3`, `Q4` (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountBillsHistoryArchiveV5(context.Background()).Year(year).Quarter(quarter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountBillsHistoryArchiveV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountBillsHistoryArchiveV5`: GetAccountBillsHistoryArchiveV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountBillsHistoryArchiveV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountBillsHistoryArchiveV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **string** | 4 digits year | [default to &quot;&quot;]
 **quarter** | **string** | Quarter, valid value is &#x60;Q1&#x60;, &#x60;Q2&#x60;, &#x60;Q3&#x60;, &#x60;Q4&#x60; | [default to &quot;&quot;]

### Return type

[**GetAccountBillsHistoryArchiveV5Resp**](GetAccountBillsHistoryArchiveV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountBillsV5

> GetAccountBillsV5Resp GetAccountBillsV5(ctx).InstType(instType).InstId(instId).Ccy(ccy).MgnMode(mgnMode).CtType(ctType).Type_(type_).SubType(subType).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()

Retrieve the bills of the account. The bill refers to all transaction records that result in changing the balance of an account. Pagination is supported, and the response is sorted with the most recent first. This endpoint can retrieve data from the last 7 days.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	instType := "instType_example" // string | Instrument type  `SPOT`  `MARGIN`  `SWAP`  `FUTURES`  `OPTION` (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT` (optional) (default to "")
	ccy := "ccy_example" // string | Bill currency (optional) (default to "")
	mgnMode := "mgnMode_example" // string | Margin mode  `isolated`  `cross` (optional) (default to "")
	ctType := "ctType_example" // string | Contract type  `linear`  `inverse`  Only applicable to `FUTURES`/`SWAP` (optional) (default to "")
	type_ := "type__example" // string | Bill type  `1`: Transfer  `2`: Trade  `3`: Delivery  `4`: Forced repayment  `5`: Liquidation  `6`: Margin transfer  `7`: Interest deduction  `8`: Funding fee  `9`: ADL  `10`: Clawback  `11`: System token conversion  `12`: Strategy transfer  `13`: DDH  `14`: Block trade  `15`: Quick Margin  `16`: Borrowing  `22`: Repay  `24`: Spread trading  `26`: Structured products  `27`: Convert  `28`: Easy convert  `29`: One-click repay  `30`: Simple trade  `33`: Loans  `34`: Settlement  `250`: Copy trader profit sharing expenses  `251`: Copy trader profit sharing refund (optional) (default to "")
	subType := "subType_example" // string | Bill subtype  `1`: Buy  `2`: Sell  `3`: Open long  `4`: Open short  `5`: Close long  `6`: Close short  `9`: Interest deduction for Market loans  `11`: Transfer in  `12`: Transfer out  `14`: Interest deduction for VIP loans  `160`: Manual margin increase  `161`: Manual margin decrease  `162`: Auto margin increase  `114`: Forced repayment buy  `115`: Forced repayment sell  `118`: System token conversion transfer in  `119`: System token conversion transfer out  `100`: Partial liquidation close long  `101`: Partial liquidation close short  `102`: Partial liquidation buy  `103`: Partial liquidation sell  `104`: Liquidation long  `105`: Liquidation short  `106`: Liquidation buy  `107`: Liquidation sell  `108`: Clawback  `110`: Liquidation transfer in  `111`: Liquidation transfer out  `125`: ADL close long  `126`: ADL close short  `127`: ADL buy  `128`: ADL sell  `131`: ddh buy  `132`: ddh sell  `170`: Exercised(ITM buy side)  `171`: Counterparty exercised(ITM sell side)  `172`: Expired(Non-ITM buy and sell side)  `112`: Delivery long  `113`: Delivery short  `117`: Delivery/Exercise clawback  `173`: Funding fee expense  `174`: Funding fee income  `200`:System transfer in  `201`: Manually transfer in  `202`: System transfer out  `203`: Manually transfer out  `204`: block trade buy  `205`: block trade sell  `206`: block trade open long  `207`: block trade open short  `208`: block trade close long  `209`: block trade close short  `210`: Manual Borrowing of quick margin  `211`: Manual Repayment of quick margin  `212`: Auto borrow of quick margin  `213`: Auto repay of quick margin  `220`: Transfer in when using USDT to buy OPTION  `221`: Transfer out when using USDT to buy OPTION  `16`: Repay forcibly  `17`: Repay interest by borrowing forcibly  `224`: Repayment transfer in  `225`: Repayment transfer out  `236`: Easy convert in  `237`: Easy convert out  `250`: Profit sharing expenses  `251`: Profit sharing refund  `280`: SPOT profit sharing expenses  `281`: SPOT profit sharing refund  `270`: Spread trading buy  `271`: Spread trading sell  `272`: Spread trading open long  `273`: Spread trading open short  `274`: Spread trading close long  `275`: Spread trading close short  `280`: SPOT profit sharing expenses  `281`: SPOT profit sharing refund   `284`: Copy trade automatic transfer in  `285`: Copy trade manual transfer in  `286`: Copy trade automatic transfer out  `287`: Copy trade manual transfer out  `290`: Crypto dust auto-transfer out  `293`: Fixed loan interest deduction  `294`: Fixed loan interest refund  `295`: Fixed loan overdue penalty  `296`: From structured order placements  `297`: To structured order placements  `298`: From structured settlements  `299`: To structured settlements  `306`: Manual borrow  `307`: Auto borrow  `308`: Manual repay  `309`: Auto repay  `312`: Auto offset  `318`: Convert in  `319`: Convert out  `320`: Simple buy  `321`: Simple sell  `332`: Margin transfer in isolated margin position   `333`: Margin transfer out isolated margin position  `334`: Margin loss when closing isolated margin position  `348`: [Credit line] Forced repayment  `350`: [Credit line] Forced repayment refund  `352`: [Credit line] Forced repayment penalty fee deduction  `353`: [Credit line] Forced repayment penalty fee (pending deduction)  `356`: [Credit line] Auto conversion (pending deduction)  `357`: [Credit line] Auto Conversion Transfer to Funding  `355`: Settlement PnL (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested bill ID. (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested bill ID. (optional) (default to "")
	begin := "begin_example" // string | Filter with a begin timestamp `ts`. Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	end := "end_example" // string | Filter with an end timestamp `ts`. Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountBillsV5(context.Background()).InstType(instType).InstId(instId).Ccy(ccy).MgnMode(mgnMode).CtType(ctType).Type_(type_).SubType(subType).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountBillsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountBillsV5`: GetAccountBillsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountBillsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountBillsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60; | [default to &quot;&quot;]
 **ccy** | **string** | Bill currency | [default to &quot;&quot;]
 **mgnMode** | **string** | Margin mode  &#x60;isolated&#x60;  &#x60;cross&#x60; | [default to &quot;&quot;]
 **ctType** | **string** | Contract type  &#x60;linear&#x60;  &#x60;inverse&#x60;  Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60; | [default to &quot;&quot;]
 **type_** | **string** | Bill type  &#x60;1&#x60;: Transfer  &#x60;2&#x60;: Trade  &#x60;3&#x60;: Delivery  &#x60;4&#x60;: Forced repayment  &#x60;5&#x60;: Liquidation  &#x60;6&#x60;: Margin transfer  &#x60;7&#x60;: Interest deduction  &#x60;8&#x60;: Funding fee  &#x60;9&#x60;: ADL  &#x60;10&#x60;: Clawback  &#x60;11&#x60;: System token conversion  &#x60;12&#x60;: Strategy transfer  &#x60;13&#x60;: DDH  &#x60;14&#x60;: Block trade  &#x60;15&#x60;: Quick Margin  &#x60;16&#x60;: Borrowing  &#x60;22&#x60;: Repay  &#x60;24&#x60;: Spread trading  &#x60;26&#x60;: Structured products  &#x60;27&#x60;: Convert  &#x60;28&#x60;: Easy convert  &#x60;29&#x60;: One-click repay  &#x60;30&#x60;: Simple trade  &#x60;33&#x60;: Loans  &#x60;34&#x60;: Settlement  &#x60;250&#x60;: Copy trader profit sharing expenses  &#x60;251&#x60;: Copy trader profit sharing refund | [default to &quot;&quot;]
 **subType** | **string** | Bill subtype  &#x60;1&#x60;: Buy  &#x60;2&#x60;: Sell  &#x60;3&#x60;: Open long  &#x60;4&#x60;: Open short  &#x60;5&#x60;: Close long  &#x60;6&#x60;: Close short  &#x60;9&#x60;: Interest deduction for Market loans  &#x60;11&#x60;: Transfer in  &#x60;12&#x60;: Transfer out  &#x60;14&#x60;: Interest deduction for VIP loans  &#x60;160&#x60;: Manual margin increase  &#x60;161&#x60;: Manual margin decrease  &#x60;162&#x60;: Auto margin increase  &#x60;114&#x60;: Forced repayment buy  &#x60;115&#x60;: Forced repayment sell  &#x60;118&#x60;: System token conversion transfer in  &#x60;119&#x60;: System token conversion transfer out  &#x60;100&#x60;: Partial liquidation close long  &#x60;101&#x60;: Partial liquidation close short  &#x60;102&#x60;: Partial liquidation buy  &#x60;103&#x60;: Partial liquidation sell  &#x60;104&#x60;: Liquidation long  &#x60;105&#x60;: Liquidation short  &#x60;106&#x60;: Liquidation buy  &#x60;107&#x60;: Liquidation sell  &#x60;108&#x60;: Clawback  &#x60;110&#x60;: Liquidation transfer in  &#x60;111&#x60;: Liquidation transfer out  &#x60;125&#x60;: ADL close long  &#x60;126&#x60;: ADL close short  &#x60;127&#x60;: ADL buy  &#x60;128&#x60;: ADL sell  &#x60;131&#x60;: ddh buy  &#x60;132&#x60;: ddh sell  &#x60;170&#x60;: Exercised(ITM buy side)  &#x60;171&#x60;: Counterparty exercised(ITM sell side)  &#x60;172&#x60;: Expired(Non-ITM buy and sell side)  &#x60;112&#x60;: Delivery long  &#x60;113&#x60;: Delivery short  &#x60;117&#x60;: Delivery/Exercise clawback  &#x60;173&#x60;: Funding fee expense  &#x60;174&#x60;: Funding fee income  &#x60;200&#x60;:System transfer in  &#x60;201&#x60;: Manually transfer in  &#x60;202&#x60;: System transfer out  &#x60;203&#x60;: Manually transfer out  &#x60;204&#x60;: block trade buy  &#x60;205&#x60;: block trade sell  &#x60;206&#x60;: block trade open long  &#x60;207&#x60;: block trade open short  &#x60;208&#x60;: block trade close long  &#x60;209&#x60;: block trade close short  &#x60;210&#x60;: Manual Borrowing of quick margin  &#x60;211&#x60;: Manual Repayment of quick margin  &#x60;212&#x60;: Auto borrow of quick margin  &#x60;213&#x60;: Auto repay of quick margin  &#x60;220&#x60;: Transfer in when using USDT to buy OPTION  &#x60;221&#x60;: Transfer out when using USDT to buy OPTION  &#x60;16&#x60;: Repay forcibly  &#x60;17&#x60;: Repay interest by borrowing forcibly  &#x60;224&#x60;: Repayment transfer in  &#x60;225&#x60;: Repayment transfer out  &#x60;236&#x60;: Easy convert in  &#x60;237&#x60;: Easy convert out  &#x60;250&#x60;: Profit sharing expenses  &#x60;251&#x60;: Profit sharing refund  &#x60;280&#x60;: SPOT profit sharing expenses  &#x60;281&#x60;: SPOT profit sharing refund  &#x60;270&#x60;: Spread trading buy  &#x60;271&#x60;: Spread trading sell  &#x60;272&#x60;: Spread trading open long  &#x60;273&#x60;: Spread trading open short  &#x60;274&#x60;: Spread trading close long  &#x60;275&#x60;: Spread trading close short  &#x60;280&#x60;: SPOT profit sharing expenses  &#x60;281&#x60;: SPOT profit sharing refund   &#x60;284&#x60;: Copy trade automatic transfer in  &#x60;285&#x60;: Copy trade manual transfer in  &#x60;286&#x60;: Copy trade automatic transfer out  &#x60;287&#x60;: Copy trade manual transfer out  &#x60;290&#x60;: Crypto dust auto-transfer out  &#x60;293&#x60;: Fixed loan interest deduction  &#x60;294&#x60;: Fixed loan interest refund  &#x60;295&#x60;: Fixed loan overdue penalty  &#x60;296&#x60;: From structured order placements  &#x60;297&#x60;: To structured order placements  &#x60;298&#x60;: From structured settlements  &#x60;299&#x60;: To structured settlements  &#x60;306&#x60;: Manual borrow  &#x60;307&#x60;: Auto borrow  &#x60;308&#x60;: Manual repay  &#x60;309&#x60;: Auto repay  &#x60;312&#x60;: Auto offset  &#x60;318&#x60;: Convert in  &#x60;319&#x60;: Convert out  &#x60;320&#x60;: Simple buy  &#x60;321&#x60;: Simple sell  &#x60;332&#x60;: Margin transfer in isolated margin position   &#x60;333&#x60;: Margin transfer out isolated margin position  &#x60;334&#x60;: Margin loss when closing isolated margin position  &#x60;348&#x60;: [Credit line] Forced repayment  &#x60;350&#x60;: [Credit line] Forced repayment refund  &#x60;352&#x60;: [Credit line] Forced repayment penalty fee deduction  &#x60;353&#x60;: [Credit line] Forced repayment penalty fee (pending deduction)  &#x60;356&#x60;: [Credit line] Auto conversion (pending deduction)  &#x60;357&#x60;: [Credit line] Auto Conversion Transfer to Funding  &#x60;355&#x60;: Settlement PnL | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested bill ID. | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested bill ID. | [default to &quot;&quot;]
 **begin** | **string** | Filter with a begin timestamp &#x60;ts&#x60;. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **end** | **string** | Filter with an end timestamp &#x60;ts&#x60;. Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAccountBillsV5Resp**](GetAccountBillsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountCollateralAssetsV5

> GetAccountCollateralAssetsV5Resp GetAccountCollateralAssetsV5(ctx).Ccy(ccy).CollateralEnabled(collateralEnabled).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	ccy := "ccy_example" // string | Single currency or multiple currencies (no more than 20) separated with comma, e.g. \"BTC\" or \"BTC,ETH\". (optional) (default to "")
	collateralEnabled := true // bool | Whether or not to be a collateral asset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountCollateralAssetsV5(context.Background()).Ccy(ccy).CollateralEnabled(collateralEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountCollateralAssetsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountCollateralAssetsV5`: GetAccountCollateralAssetsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountCollateralAssetsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountCollateralAssetsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Single currency or multiple currencies (no more than 20) separated with comma, e.g. \&quot;BTC\&quot; or \&quot;BTC,ETH\&quot;. | [default to &quot;&quot;]
 **collateralEnabled** | **bool** | Whether or not to be a collateral asset | 

### Return type

[**GetAccountCollateralAssetsV5Resp**](GetAccountCollateralAssetsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountConfigV5

> GetAccountConfigV5Resp GetAccountConfigV5(ctx).Execute()

Retrieve current account configuration.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountConfigV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountConfigV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountConfigV5`: GetAccountConfigV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountConfigV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountConfigV5Request struct via the builder pattern


### Return type

[**GetAccountConfigV5Resp**](GetAccountConfigV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountGreeksV5

> GetAccountGreeksV5Resp GetAccountGreeksV5(ctx).Ccy(ccy).Execute()

Retrieve a greeks list of all assets in the account.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	ccy := "ccy_example" // string | Single currency, e.g. `BTC`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountGreeksV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountGreeksV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountGreeksV5`: GetAccountGreeksV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountGreeksV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountGreeksV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Single currency, e.g. &#x60;BTC&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAccountGreeksV5Resp**](GetAccountGreeksV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInstrumentsV5

> GetAccountInstrumentsV5Resp GetAccountInstrumentsV5(ctx).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).Execute()

Retrieve available instruments info of current account.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	instType := "instType_example" // string | Instrument type  `SPOT`: Spot  `MARGIN`: Margin  `SWAP`: Perpetual Futures  `FUTURES`: Expiry Futures  `OPTION`: Option (default to "")
	uly := "uly_example" // string | Underlying   Only applicable to `FUTURES`/`SWAP`/`OPTION`.If instType is `OPTION`, either `uly` or `instFamily` is required. (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family  Only applicable to `FUTURES`/`SWAP`/`OPTION`. If instType is `OPTION`, either `uly` or `instFamily` is required. (optional) (default to "")
	instId := "instId_example" // string | Instrument ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountInstrumentsV5(context.Background()).InstType(instType).Uly(uly).InstFamily(instFamily).InstId(instId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountInstrumentsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInstrumentsV5`: GetAccountInstrumentsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountInstrumentsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInstrumentsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;: Spot  &#x60;MARGIN&#x60;: Margin  &#x60;SWAP&#x60;: Perpetual Futures  &#x60;FUTURES&#x60;: Expiry Futures  &#x60;OPTION&#x60;: Option | [default to &quot;&quot;]
 **uly** | **string** | Underlying   Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;.If instType is &#x60;OPTION&#x60;, either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family  Only applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60;. If instType is &#x60;OPTION&#x60;, either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID | [default to &quot;&quot;]

### Return type

[**GetAccountInstrumentsV5Resp**](GetAccountInstrumentsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInterestAccruedV5

> GetAccountInterestAccruedV5Resp GetAccountInterestAccruedV5(ctx).Type_(type_).Ccy(ccy).InstId(instId).MgnMode(mgnMode).After(after).Before(before).Limit(limit).Execute()

Get interest accrued data. Only data within the last one year can be obtained.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	type_ := "type__example" // string | Loan type  `2`: Market loans  Default is `2` (optional) (default to "")
	ccy := "ccy_example" // string | Loan currency, e.g. `BTC`  Only applicable to `Market loans`  Only applicable to`MARGIN` (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT`  Only applicable to `Market loans` (optional) (default to "")
	mgnMode := "mgnMode_example" // string | Margin mode  `cross`    `isolated`  Only applicable to `Market loans` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested timestamp, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`. The default is `100`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountInterestAccruedV5(context.Background()).Type_(type_).Ccy(ccy).InstId(instId).MgnMode(mgnMode).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountInterestAccruedV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInterestAccruedV5`: GetAccountInterestAccruedV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountInterestAccruedV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInterestAccruedV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Loan type  &#x60;2&#x60;: Market loans  Default is &#x60;2&#x60; | [default to &quot;&quot;]
 **ccy** | **string** | Loan currency, e.g. &#x60;BTC&#x60;  Only applicable to &#x60;Market loans&#x60;  Only applicable to&#x60;MARGIN&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60;  Only applicable to &#x60;Market loans&#x60; | [default to &quot;&quot;]
 **mgnMode** | **string** | Margin mode  &#x60;cross&#x60;    &#x60;isolated&#x60;  Only applicable to &#x60;Market loans&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested timestamp, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;. The default is &#x60;100&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAccountInterestAccruedV5Resp**](GetAccountInterestAccruedV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInterestLimitsV5

> GetAccountInterestLimitsV5Resp GetAccountInterestLimitsV5(ctx).Type_(type_).Ccy(ccy).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	type_ := "type__example" // string | Loan type  `2`: Market loans  Default is `2` (optional) (default to "")
	ccy := "ccy_example" // string | Loan currency, e.g. `BTC` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountInterestLimitsV5(context.Background()).Type_(type_).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountInterestLimitsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInterestLimitsV5`: GetAccountInterestLimitsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountInterestLimitsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInterestLimitsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Loan type  &#x60;2&#x60;: Market loans  Default is &#x60;2&#x60; | [default to &quot;&quot;]
 **ccy** | **string** | Loan currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]

### Return type

[**GetAccountInterestLimitsV5Resp**](GetAccountInterestLimitsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInterestRateV5

> GetAccountInterestRateV5Resp GetAccountInterestRateV5(ctx).Ccy(ccy).Execute()

Get the user's current leveraged currency borrowing market interest rate  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	ccy := "ccy_example" // string | Currency, e.g. `BTC` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountInterestRateV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountInterestRateV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInterestRateV5`: GetAccountInterestRateV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountInterestRateV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInterestRateV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]

### Return type

[**GetAccountInterestRateV5Resp**](GetAccountInterestRateV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountLeverageInfoV5

> GetAccountLeverageInfoV5Resp GetAccountLeverageInfoV5(ctx).MgnMode(mgnMode).InstId(instId).Ccy(ccy).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	mgnMode := "mgnMode_example" // string | Margin mode  `cross` `isolated` (default to "")
	instId := "instId_example" // string | Instrument ID  Single instrument ID or multiple instrument IDs (no more than 20) separated with comma (optional) (default to "")
	ccy := "ccy_example" // string | Currency，used for getting leverage of currency level.  Applicable to `cross` `MARGIN` of `Spot mode`/`Multi-currency margin`/`Portfolio margin`.  Supported single currency or multiple currencies (no more than 20) separated with comma. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountLeverageInfoV5(context.Background()).MgnMode(mgnMode).InstId(instId).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountLeverageInfoV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountLeverageInfoV5`: GetAccountLeverageInfoV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountLeverageInfoV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountLeverageInfoV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mgnMode** | **string** | Margin mode  &#x60;cross&#x60; &#x60;isolated&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID  Single instrument ID or multiple instrument IDs (no more than 20) separated with comma | [default to &quot;&quot;]
 **ccy** | **string** | Currency，used for getting leverage of currency level.  Applicable to &#x60;cross&#x60; &#x60;MARGIN&#x60; of &#x60;Spot mode&#x60;/&#x60;Multi-currency margin&#x60;/&#x60;Portfolio margin&#x60;.  Supported single currency or multiple currencies (no more than 20) separated with comma. | [default to &quot;&quot;]

### Return type

[**GetAccountLeverageInfoV5Resp**](GetAccountLeverageInfoV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountMaxAvailSizeV5

> GetAccountMaxAvailSizeV5Resp GetAccountMaxAvailSizeV5(ctx).InstId(instId).TdMode(tdMode).Ccy(ccy).ReduceOnly(reduceOnly).Px(px).Execute()

Available balance for isolated margin positions and SPOT, available equity for cross margin positions.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	instId := "instId_example" // string | Single instrument or multiple instruments (no more than 5) separated with comma, e.g. `BTC-USDT,ETH-USDT` (default to "")
	tdMode := "tdMode_example" // string | Trade mode  `cross`  `isolated`  `cash`  `spot_isolated` (default to "")
	ccy := "ccy_example" // string | Currency used for margin  Applicable to `isolated` `MARGIN` and `cross` `MARGIN` in `Spot and futures mode`. (optional) (default to "")
	reduceOnly := true // bool | Whether to reduce position only   Only applicable to `MARGIN` (optional)
	px := "px_example" // string | The price of closing position.   Only applicable to reduceOnly `MARGIN`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountMaxAvailSizeV5(context.Background()).InstId(instId).TdMode(tdMode).Ccy(ccy).ReduceOnly(reduceOnly).Px(px).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountMaxAvailSizeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountMaxAvailSizeV5`: GetAccountMaxAvailSizeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountMaxAvailSizeV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountMaxAvailSizeV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Single instrument or multiple instruments (no more than 5) separated with comma, e.g. &#x60;BTC-USDT,ETH-USDT&#x60; | [default to &quot;&quot;]
 **tdMode** | **string** | Trade mode  &#x60;cross&#x60;  &#x60;isolated&#x60;  &#x60;cash&#x60;  &#x60;spot_isolated&#x60; | [default to &quot;&quot;]
 **ccy** | **string** | Currency used for margin  Applicable to &#x60;isolated&#x60; &#x60;MARGIN&#x60; and &#x60;cross&#x60; &#x60;MARGIN&#x60; in &#x60;Spot and futures mode&#x60;. | [default to &quot;&quot;]
 **reduceOnly** | **bool** | Whether to reduce position only   Only applicable to &#x60;MARGIN&#x60; | 
 **px** | **string** | The price of closing position.   Only applicable to reduceOnly &#x60;MARGIN&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAccountMaxAvailSizeV5Resp**](GetAccountMaxAvailSizeV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountMaxLoanV5

> GetAccountMaxLoanV5Resp GetAccountMaxLoanV5(ctx).MgnMode(mgnMode).InstId(instId).Ccy(ccy).MgnCcy(mgnCcy).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	mgnMode := "mgnMode_example" // string | Margin mode  `isolated` `cross` (default to "")
	instId := "instId_example" // string | Single instrument or multiple instruments (no more than 5) separated with comma, e.g. `BTC-USDT,ETH-USDT` (optional) (default to "")
	ccy := "ccy_example" // string | Currency  Applicable to get Max loan of manual borrow for the currency in `Spot mode` (enabled borrowing) (optional) (default to "")
	mgnCcy := "mgnCcy_example" // string | Margin currency  Applicable to `isolated` `MARGIN` and `cross` `MARGIN` in `Spot and futures mode`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountMaxLoanV5(context.Background()).MgnMode(mgnMode).InstId(instId).Ccy(ccy).MgnCcy(mgnCcy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountMaxLoanV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountMaxLoanV5`: GetAccountMaxLoanV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountMaxLoanV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountMaxLoanV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mgnMode** | **string** | Margin mode  &#x60;isolated&#x60; &#x60;cross&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Single instrument or multiple instruments (no more than 5) separated with comma, e.g. &#x60;BTC-USDT,ETH-USDT&#x60; | [default to &quot;&quot;]
 **ccy** | **string** | Currency  Applicable to get Max loan of manual borrow for the currency in &#x60;Spot mode&#x60; (enabled borrowing) | [default to &quot;&quot;]
 **mgnCcy** | **string** | Margin currency  Applicable to &#x60;isolated&#x60; &#x60;MARGIN&#x60; and &#x60;cross&#x60; &#x60;MARGIN&#x60; in &#x60;Spot and futures mode&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAccountMaxLoanV5Resp**](GetAccountMaxLoanV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountMaxSizeV5

> GetAccountMaxSizeV5Resp GetAccountMaxSizeV5(ctx).InstId(instId).TdMode(tdMode).Ccy(ccy).Px(px).Leverage(leverage).Execute()

The maximum quantity to buy or sell. It corresponds to the \"sz\" from placement.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	instId := "instId_example" // string | Single instrument or multiple instruments (no more than 5) in the smae instrument type separated with comma, e.g. `BTC-USDT,ETH-USDT` (default to "")
	tdMode := "tdMode_example" // string | Trade mode  `cross`  `isolated`  `cash`  `spot_isolated` (default to "")
	ccy := "ccy_example" // string | Currency used for margin   Applicable to `isolated` `MARGIN` and `cross` `MARGIN` orders in `Spot and futures mode`. (optional) (default to "")
	px := "px_example" // string | Price  When the price is not specified, it will be calculated according to the current limit price for `FUTURES` and `SWAP`, the last traded price for other instrument types.  The parameter will be ignored when multiple instruments are specified. (optional) (default to "")
	leverage := "leverage_example" // string | Leverage for instrument  The default is current leverage  Only applicable to `MARGIN/FUTURES/SWAP` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountMaxSizeV5(context.Background()).InstId(instId).TdMode(tdMode).Ccy(ccy).Px(px).Leverage(leverage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountMaxSizeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountMaxSizeV5`: GetAccountMaxSizeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountMaxSizeV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountMaxSizeV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Single instrument or multiple instruments (no more than 5) in the smae instrument type separated with comma, e.g. &#x60;BTC-USDT,ETH-USDT&#x60; | [default to &quot;&quot;]
 **tdMode** | **string** | Trade mode  &#x60;cross&#x60;  &#x60;isolated&#x60;  &#x60;cash&#x60;  &#x60;spot_isolated&#x60; | [default to &quot;&quot;]
 **ccy** | **string** | Currency used for margin   Applicable to &#x60;isolated&#x60; &#x60;MARGIN&#x60; and &#x60;cross&#x60; &#x60;MARGIN&#x60; orders in &#x60;Spot and futures mode&#x60;. | [default to &quot;&quot;]
 **px** | **string** | Price  When the price is not specified, it will be calculated according to the current limit price for &#x60;FUTURES&#x60; and &#x60;SWAP&#x60;, the last traded price for other instrument types.  The parameter will be ignored when multiple instruments are specified. | [default to &quot;&quot;]
 **leverage** | **string** | Leverage for instrument  The default is current leverage  Only applicable to &#x60;MARGIN/FUTURES/SWAP&#x60; | [default to &quot;&quot;]

### Return type

[**GetAccountMaxSizeV5Resp**](GetAccountMaxSizeV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountMaxWithdrawalV5

> GetAccountMaxWithdrawalV5Resp GetAccountMaxWithdrawalV5(ctx).Ccy(ccy).Execute()

Retrieve the maximum transferable amount from trading account to funding account. If no currency is specified, the transferable amount of all owned currencies will be returned.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	ccy := "ccy_example" // string | Single currency or multiple currencies (no more than 20) separated with comma, e.g. `BTC` or `BTC,ETH`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountMaxWithdrawalV5(context.Background()).Ccy(ccy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountMaxWithdrawalV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountMaxWithdrawalV5`: GetAccountMaxWithdrawalV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountMaxWithdrawalV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountMaxWithdrawalV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Single currency or multiple currencies (no more than 20) separated with comma, e.g. &#x60;BTC&#x60; or &#x60;BTC,ETH&#x60;. | [default to &quot;&quot;]

### Return type

[**GetAccountMaxWithdrawalV5Resp**](GetAccountMaxWithdrawalV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountMmpConfigV5

> GetAccountMmpConfigV5Resp GetAccountMmpConfigV5(ctx).InstFamily(instFamily).Execute()

This endpoint is used to get MMP configure information    Only applicable to Option in Portfolio Margin mode, and MMP privilege is required.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	instFamily := "instFamily_example" // string | Instrument Family (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountMmpConfigV5(context.Background()).InstFamily(instFamily).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountMmpConfigV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountMmpConfigV5`: GetAccountMmpConfigV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountMmpConfigV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountMmpConfigV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instFamily** | **string** | Instrument Family | [default to &quot;&quot;]

### Return type

[**GetAccountMmpConfigV5Resp**](GetAccountMmpConfigV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountPositionTiersV5

> GetAccountPositionTiersV5Resp GetAccountPositionTiersV5(ctx).InstType(instType).Uly(uly).InstFamily(instFamily).Execute()

Retrieve cross position limitation of SWAP/FUTURES/OPTION under Portfolio margin mode.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	instType := "instType_example" // string | Instrument type  `SWAP`  `FUTURES`  `OPTION` (default to "")
	uly := "uly_example" // string | Single underlying or multiple underlyings (no more than 3) separated with comma.  Either `uly` or `instFamily` is required. If both are passed, `instFamily` will be used. (optional) (default to "")
	instFamily := "instFamily_example" // string | Single instrument family or instrument families (no more than 5) separated with comma.  Either `uly` or `instFamily` is required. If both are passed, `instFamily` will be used. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountPositionTiersV5(context.Background()).InstType(instType).Uly(uly).InstFamily(instFamily).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountPositionTiersV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountPositionTiersV5`: GetAccountPositionTiersV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountPositionTiersV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountPositionTiersV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Single underlying or multiple underlyings (no more than 3) separated with comma.  Either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. If both are passed, &#x60;instFamily&#x60; will be used. | [default to &quot;&quot;]
 **instFamily** | **string** | Single instrument family or instrument families (no more than 5) separated with comma.  Either &#x60;uly&#x60; or &#x60;instFamily&#x60; is required. If both are passed, &#x60;instFamily&#x60; will be used. | [default to &quot;&quot;]

### Return type

[**GetAccountPositionTiersV5Resp**](GetAccountPositionTiersV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountPositionsHistoryV5

> GetAccountPositionsHistoryV5Resp GetAccountPositionsHistoryV5(ctx).InstType(instType).InstId(instId).MgnMode(mgnMode).Type_(type_).PosId(posId).After(after).Before(before).Limit(limit).Execute()

Retrieve the updated position data for the last 3 months. Return in reverse chronological order using utime. Getting positions history is supported under Portfolio margin mode since .  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	instType := "instType_example" // string | Instrument type  `MARGIN`  `SWAP`  `FUTURES`  `OPTION` (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USD-SWAP` (optional) (default to "")
	mgnMode := "mgnMode_example" // string | Margin mode  `cross` `isolated` (optional) (default to "")
	type_ := "type__example" // string | The type of latest close position  `1`: Close position partially;`2`：Close all;`3`：Liquidation;`4`：Partial liquidation; `5`：ADL;   It is the latest type if there are several types for the same position. (optional) (default to "")
	posId := "posId_example" // string | Position ID. There is attribute expiration. The posId will be expired if it is more than 30 days after the last full close position, then position will use new posId. (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `uTime`, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `uTime`, Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. All records that have the same `uTime` will be returned at the current request (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountPositionsHistoryV5(context.Background()).InstType(instType).InstId(instId).MgnMode(mgnMode).Type_(type_).PosId(posId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountPositionsHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountPositionsHistoryV5`: GetAccountPositionsHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountPositionsHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountPositionsHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USD-SWAP&#x60; | [default to &quot;&quot;]
 **mgnMode** | **string** | Margin mode  &#x60;cross&#x60; &#x60;isolated&#x60; | [default to &quot;&quot;]
 **type_** | **string** | The type of latest close position  &#x60;1&#x60;: Close position partially;&#x60;2&#x60;：Close all;&#x60;3&#x60;：Liquidation;&#x60;4&#x60;：Partial liquidation; &#x60;5&#x60;：ADL;   It is the latest type if there are several types for the same position. | [default to &quot;&quot;]
 **posId** | **string** | Position ID. There is attribute expiration. The posId will be expired if it is more than 30 days after the last full close position, then position will use new posId. | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;uTime&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;uTime&#x60;, Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. All records that have the same &#x60;uTime&#x60; will be returned at the current request | [default to &quot;&quot;]

### Return type

[**GetAccountPositionsHistoryV5Resp**](GetAccountPositionsHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountPositionsV5

> GetAccountPositionsV5Resp GetAccountPositionsV5(ctx).InstType(instType).InstId(instId).PosId(posId).Execute()

Retrieve information on your positions. When the account is in `net` mode, `net` positions will be displayed, and when the account is in `long/short` mode, `long` or `short` positions will be displayed. Return in reverse chronological order using ctime.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	instType := "instType_example" // string | Instrument type  `MARGIN`  `SWAP`  `FUTURES`  `OPTION`  `instId` will be checked against `instType` when both parameters are passed. (optional) (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT-SWAP`. Single instrument ID or multiple instrument IDs (no more than 10) separated with comma (optional) (default to "")
	posId := "posId_example" // string | Single position ID or multiple position IDs (no more than 20) separated with comma.   There is attribute expiration, the posId and position information will be cleared if it is more than 30 days after the last full close position. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountPositionsV5(context.Background()).InstType(instType).InstId(instId).PosId(posId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountPositionsV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountPositionsV5`: GetAccountPositionsV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountPositionsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountPositionsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60;  &#x60;instId&#x60; will be checked against &#x60;instType&#x60; when both parameters are passed. | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT-SWAP&#x60;. Single instrument ID or multiple instrument IDs (no more than 10) separated with comma | [default to &quot;&quot;]
 **posId** | **string** | Single position ID or multiple position IDs (no more than 20) separated with comma.   There is attribute expiration, the posId and position information will be cleared if it is more than 30 days after the last full close position. | [default to &quot;&quot;]

### Return type

[**GetAccountPositionsV5Resp**](GetAccountPositionsV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountQuickMarginBorrowRepayHistoryV5

> GetAccountQuickMarginBorrowRepayHistoryV5Resp GetAccountQuickMarginBorrowRepayHistoryV5(ctx).InstId(instId).Ccy(ccy).Side(side).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()

Get record in the past 3 months.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	instId := "instId_example" // string | Instrument ID, e.g. BTC-USDT (optional) (default to "")
	ccy := "ccy_example" // string | Loan currency, e.g. `BTC` (optional) (default to "")
	side := "side_example" // string | `borrow`  `repay` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `refId` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `refId` (optional) (default to "")
	begin := "begin_example" // string | Filter with a begin timestamp. Unix timestamp format in milliseconds, e.g. 1597026383085 (optional) (default to "")
	end := "end_example" // string | Filter with an end timestamp. Unix timestamp format in milliseconds, e.g. 1597026383085 (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is `100`; The default is `100` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountQuickMarginBorrowRepayHistoryV5(context.Background()).InstId(instId).Ccy(ccy).Side(side).After(after).Before(before).Begin(begin).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountQuickMarginBorrowRepayHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountQuickMarginBorrowRepayHistoryV5`: GetAccountQuickMarginBorrowRepayHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountQuickMarginBorrowRepayHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountQuickMarginBorrowRepayHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instId** | **string** | Instrument ID, e.g. BTC-USDT | [default to &quot;&quot;]
 **ccy** | **string** | Loan currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]
 **side** | **string** | &#x60;borrow&#x60;  &#x60;repay&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;refId&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;refId&#x60; | [default to &quot;&quot;]
 **begin** | **string** | Filter with a begin timestamp. Unix timestamp format in milliseconds, e.g. 1597026383085 | [default to &quot;&quot;]
 **end** | **string** | Filter with an end timestamp. Unix timestamp format in milliseconds, e.g. 1597026383085 | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is &#x60;100&#x60;; The default is &#x60;100&#x60; | [default to &quot;&quot;]

### Return type

[**GetAccountQuickMarginBorrowRepayHistoryV5Resp**](GetAccountQuickMarginBorrowRepayHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountRiskStateV5

> GetAccountRiskStateV5Resp GetAccountRiskStateV5(ctx).Execute()

Only applicable to Portfolio margin account  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountRiskStateV5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountRiskStateV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountRiskStateV5`: GetAccountRiskStateV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountRiskStateV5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRiskStateV5Request struct via the builder pattern


### Return type

[**GetAccountRiskStateV5Resp**](GetAccountRiskStateV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountSetAccountSwitchPrecheckV5

> GetAccountSetAccountSwitchPrecheckV5Resp GetAccountSetAccountSwitchPrecheckV5(ctx).AcctLv(acctLv).Execute()

Retrieve precheck information for account mode switching.  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	acctLv := "acctLv_example" // string | Account mode   `1`: Spot mode   `2`: Spot and futures mode   `3`: Multi-currency margin code   `4`: Portfolio margin mode (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountSetAccountSwitchPrecheckV5(context.Background()).AcctLv(acctLv).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountSetAccountSwitchPrecheckV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountSetAccountSwitchPrecheckV5`: GetAccountSetAccountSwitchPrecheckV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountSetAccountSwitchPrecheckV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSetAccountSwitchPrecheckV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acctLv** | **string** | Account mode   &#x60;1&#x60;: Spot mode   &#x60;2&#x60;: Spot and futures mode   &#x60;3&#x60;: Multi-currency margin code   &#x60;4&#x60;: Portfolio margin mode | [default to &quot;&quot;]

### Return type

[**GetAccountSetAccountSwitchPrecheckV5Resp**](GetAccountSetAccountSwitchPrecheckV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountSpotBorrowRepayHistoryV5

> GetAccountSpotBorrowRepayHistoryV5Resp GetAccountSpotBorrowRepayHistoryV5(ctx).Ccy(ccy).Type_(type_).After(after).Before(before).Limit(limit).Execute()

Retrieve the borrow/repay history under `Spot mode`  



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	ccy := "ccy_example" // string | Currency, e.g. `BTC` (optional) (default to "")
	type_ := "type__example" // string | Event type  `auto_borrow`  `auto_repay`  `manual_borrow`  `manual_repay` (optional) (default to "")
	after := "after_example" // string | Pagination of data to return records earlier than the requested `ts` (included), Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	before := "before_example" // string | Pagination of data to return records newer than the requested `ts`(included), Unix timestamp format in milliseconds, e.g. `1597026383085` (optional) (default to "")
	limit := "limit_example" // string | Number of results per request. The maximum is 100. The default is 100. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountSpotBorrowRepayHistoryV5(context.Background()).Ccy(ccy).Type_(type_).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountSpotBorrowRepayHistoryV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountSpotBorrowRepayHistoryV5`: GetAccountSpotBorrowRepayHistoryV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountSpotBorrowRepayHistoryV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSpotBorrowRepayHistoryV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ccy** | **string** | Currency, e.g. &#x60;BTC&#x60; | [default to &quot;&quot;]
 **type_** | **string** | Event type  &#x60;auto_borrow&#x60;  &#x60;auto_repay&#x60;  &#x60;manual_borrow&#x60;  &#x60;manual_repay&#x60; | [default to &quot;&quot;]
 **after** | **string** | Pagination of data to return records earlier than the requested &#x60;ts&#x60; (included), Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **before** | **string** | Pagination of data to return records newer than the requested &#x60;ts&#x60;(included), Unix timestamp format in milliseconds, e.g. &#x60;1597026383085&#x60; | [default to &quot;&quot;]
 **limit** | **string** | Number of results per request. The maximum is 100. The default is 100. | [default to &quot;&quot;]

### Return type

[**GetAccountSpotBorrowRepayHistoryV5Resp**](GetAccountSpotBorrowRepayHistoryV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTradeFeeV5

> GetAccountTradeFeeV5Resp GetAccountTradeFeeV5(ctx).InstType(instType).RuleType(ruleType).InstId(instId).Uly(uly).InstFamily(instFamily).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func main() {
	instType := "instType_example" // string | Instrument type  `SPOT`  `MARGIN`  `SWAP`  `FUTURES`  `OPTION` (default to "")
	ruleType := "ruleType_example" // string | Trading rule types   `normal`: normal trading   `pre_market`: pre-market trading   ruleType can not be passed through together with instId/instFamily/uly (default to "")
	instId := "instId_example" // string | Instrument ID, e.g. `BTC-USDT`  Applicable to `SPOT`/`MARGIN` (optional) (default to "")
	uly := "uly_example" // string | Underlying, e.g. `BTC-USD`  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")
	instFamily := "instFamily_example" // string | Instrument family, e.g. `BTC-USD`  Applicable to `FUTURES`/`SWAP`/`OPTION` (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TradingAccountAPI.GetAccountTradeFeeV5(context.Background()).InstType(instType).RuleType(ruleType).InstId(instId).Uly(uly).InstFamily(instFamily).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradingAccountAPI.GetAccountTradeFeeV5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountTradeFeeV5`: GetAccountTradeFeeV5Resp
	fmt.Fprintf(os.Stdout, "Response from `TradingAccountAPI.GetAccountTradeFeeV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTradeFeeV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instType** | **string** | Instrument type  &#x60;SPOT&#x60;  &#x60;MARGIN&#x60;  &#x60;SWAP&#x60;  &#x60;FUTURES&#x60;  &#x60;OPTION&#x60; | [default to &quot;&quot;]
 **ruleType** | **string** | Trading rule types   &#x60;normal&#x60;: normal trading   &#x60;pre_market&#x60;: pre-market trading   ruleType can not be passed through together with instId/instFamily/uly | [default to &quot;&quot;]
 **instId** | **string** | Instrument ID, e.g. &#x60;BTC-USDT&#x60;  Applicable to &#x60;SPOT&#x60;/&#x60;MARGIN&#x60; | [default to &quot;&quot;]
 **uly** | **string** | Underlying, e.g. &#x60;BTC-USD&#x60;  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]
 **instFamily** | **string** | Instrument family, e.g. &#x60;BTC-USD&#x60;  Applicable to &#x60;FUTURES&#x60;/&#x60;SWAP&#x60;/&#x60;OPTION&#x60; | [default to &quot;&quot;]

### Return type

[**GetAccountTradeFeeV5Resp**](GetAccountTradeFeeV5Resp.md)

### Authorization

[APIKey](../README.md#APIKey), [Passphrase](../README.md#Passphrase)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

