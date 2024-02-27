/*
Equinix Fabric API v4

Testing ServiceTokensApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fabricv4

import (
	"context"
	"testing"

	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fabricv4_ServiceTokensApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceTokensApiService CreateServiceToken", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServiceTokensApi.CreateServiceToken(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ServiceTokensApiService CreateServiceTokenAction", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var serviceTokenId string

		resp, httpRes, err := apiClient.ServiceTokensApi.CreateServiceTokenAction(context.Background(), serviceTokenId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ServiceTokensApiService DeleteServiceTokenByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var serviceTokenId string

		httpRes, err := apiClient.ServiceTokensApi.DeleteServiceTokenByUuid(context.Background(), serviceTokenId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ServiceTokensApiService GetServiceTokenByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var serviceTokenId string

		resp, httpRes, err := apiClient.ServiceTokensApi.GetServiceTokenByUuid(context.Background(), serviceTokenId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ServiceTokensApiService GetServiceTokens", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServiceTokensApi.GetServiceTokens(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ServiceTokensApiService SearchServiceTokens", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServiceTokensApi.SearchServiceTokens(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ServiceTokensApiService UpdateServiceTokenByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var serviceTokenId string

		resp, httpRes, err := apiClient.ServiceTokensApi.UpdateServiceTokenByUuid(context.Background(), serviceTokenId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
