/*
OKX v5 API

Testing EthStakingAPIService

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

func Test_rest_EthStakingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EthStakingAPIService CreateFinanceStakingDefiEthPurchaseV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EthStakingAPI.CreateFinanceStakingDefiEthPurchaseV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EthStakingAPIService CreateFinanceStakingDefiEthRedeemV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EthStakingAPI.CreateFinanceStakingDefiEthRedeemV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EthStakingAPIService GetFinanceStakingDefiEthApyHistoryV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EthStakingAPI.GetFinanceStakingDefiEthApyHistoryV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EthStakingAPIService GetFinanceStakingDefiEthBalanceV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EthStakingAPI.GetFinanceStakingDefiEthBalanceV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EthStakingAPIService GetFinanceStakingDefiEthProductInfoV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EthStakingAPI.GetFinanceStakingDefiEthProductInfoV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EthStakingAPIService GetFinanceStakingDefiEthPurchaseRedeemHistoryV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EthStakingAPI.GetFinanceStakingDefiEthPurchaseRedeemHistoryV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
