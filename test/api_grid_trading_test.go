/*
OKX v5 API

Testing GridTradingAPIService

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

func Test_rest_GridTradingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GridTradingAPIService CreateTradingBotGridAdjustInvestmentV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.CreateTradingBotGridAdjustInvestmentV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService CreateTradingBotGridAmendOrderAlgoV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.CreateTradingBotGridAmendOrderAlgoV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService CreateTradingBotGridCancelCloseOrderV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.CreateTradingBotGridCancelCloseOrderV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService CreateTradingBotGridClosePositionV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.CreateTradingBotGridClosePositionV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService CreateTradingBotGridComputeMarginBalanceV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.CreateTradingBotGridComputeMarginBalanceV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService CreateTradingBotGridMarginBalanceV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.CreateTradingBotGridMarginBalanceV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService CreateTradingBotGridMinInvestmentV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.CreateTradingBotGridMinInvestmentV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService CreateTradingBotGridOrderAlgoV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.CreateTradingBotGridOrderAlgoV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService CreateTradingBotGridOrderInstantTriggerV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.CreateTradingBotGridOrderInstantTriggerV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService CreateTradingBotGridStopOrderAlgoV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.CreateTradingBotGridStopOrderAlgoV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService CreateTradingBotGridWithdrawIncomeV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.CreateTradingBotGridWithdrawIncomeV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService GetTradingBotGridAiParamV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.GetTradingBotGridAiParamV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService GetTradingBotGridGridQuantityV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.GetTradingBotGridGridQuantityV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService GetTradingBotGridOrdersAlgoDetailsV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.GetTradingBotGridOrdersAlgoDetailsV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService GetTradingBotGridOrdersAlgoHistoryV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.GetTradingBotGridOrdersAlgoHistoryV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService GetTradingBotGridOrdersAlgoPendingV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.GetTradingBotGridOrdersAlgoPendingV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService GetTradingBotGridPositionsV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.GetTradingBotGridPositionsV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService GetTradingBotGridSubOrdersV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.GetTradingBotGridSubOrdersV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GridTradingAPIService GetTradingBotPublicRsiBackTestingV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GridTradingAPI.GetTradingBotPublicRsiBackTestingV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
