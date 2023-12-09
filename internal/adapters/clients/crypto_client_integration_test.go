package clients

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetExchangesValuesForUSDTShouldReturnFloatWhenApiCallIsSuccessful(t *testing.T) {

	// init
	cryptoClient := NewCrypto(CryptoMarketEndpoint)

	// execute
	result, err := cryptoClient.getExchangesValuesForUSDT()

	// assert
	assert.NotNil(t, result)
	assert.NoError(t, err)
}
