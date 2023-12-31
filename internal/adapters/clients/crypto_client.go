package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var CryptoMarketEndpoint = "https://criptoya.com/api/usdt/ars/1"

type cryptoClient struct {
	endpoint string
	client   *http.Client
}

func NewCrypto(endpoint string) cryptoClient {
	client := &http.Client{}
	return cryptoClient{endpoint: endpoint, client: client}
}

func (cc *cryptoClient) GetUSDTMarketsPrice() (*ExchangeResponse, error) {

	// do http request
	return cc.getExchangesValuesForUSDT()

}

func (cc *cryptoClient) getExchangesValuesForUSDT() (*ExchangeResponse, error) {
	// do http request
	response, err := cc.client.Get(cc.endpoint)

	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}

	// handling error
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Error: Status code: " + string(rune(response.StatusCode)))
	}

	// handling okey
	var exchangeResponse ExchangeResponse
	err = json.NewDecoder(response.Body).Decode(&exchangeResponse)

	if err != nil {
		return nil, err
	}

	return &exchangeResponse, nil
}
