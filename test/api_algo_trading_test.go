/*
OKX v5 API

Testing AlgoTradingAPIService

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

func Test_rest_AlgoTradingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AlgoTradingAPIService CreateTradeAmendAlgosV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AlgoTradingAPI.CreateTradeAmendAlgosV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlgoTradingAPIService CreateTradeCancelAdvanceAlgosV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AlgoTradingAPI.CreateTradeCancelAdvanceAlgosV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlgoTradingAPIService CreateTradeCancelAlgosV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AlgoTradingAPI.CreateTradeCancelAlgosV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlgoTradingAPIService CreateTradeOrderAlgoV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AlgoTradingAPI.CreateTradeOrderAlgoV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlgoTradingAPIService GetTradeOrderAlgoV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AlgoTradingAPI.GetTradeOrderAlgoV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlgoTradingAPIService GetTradeOrdersAlgoHistoryV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AlgoTradingAPI.GetTradeOrdersAlgoHistoryV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlgoTradingAPIService GetTradeOrdersAlgoPendingV5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AlgoTradingAPI.GetTradeOrdersAlgoPendingV5(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
