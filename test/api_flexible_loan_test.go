/*
OKX v5 API

Testing FlexibleLoanAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package rest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/openxapi/okx-go/rest"
)

func Test_rest_FlexibleLoanAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FlexibleLoanAPIService CreateFinanceFlexibleLoanAdjustCollateralV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlexibleLoanAPI.CreateFinanceFlexibleLoanAdjustCollateralV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlexibleLoanAPIService CreateFinanceFlexibleLoanMaxLoanV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlexibleLoanAPI.CreateFinanceFlexibleLoanMaxLoanV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlexibleLoanAPIService GetFinanceFlexibleLoanBorrowCurrenciesV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlexibleLoanAPI.GetFinanceFlexibleLoanBorrowCurrenciesV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlexibleLoanAPIService GetFinanceFlexibleLoanCollateralAssetsV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlexibleLoanAPI.GetFinanceFlexibleLoanCollateralAssetsV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlexibleLoanAPIService GetFinanceFlexibleLoanInterestAccruedV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlexibleLoanAPI.GetFinanceFlexibleLoanInterestAccruedV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlexibleLoanAPIService GetFinanceFlexibleLoanLoanHistoryV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlexibleLoanAPI.GetFinanceFlexibleLoanLoanHistoryV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlexibleLoanAPIService GetFinanceFlexibleLoanLoanInfoV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlexibleLoanAPI.GetFinanceFlexibleLoanLoanInfoV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FlexibleLoanAPIService GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FlexibleLoanAPI.GetFinanceFlexibleLoanMaxCollateralRedeemAmountV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
