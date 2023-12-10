package usecases

import (
	"encoding/json"
	"fmt"
	"main/internal/adapters/clients"
	"main/internal/application/domain"
)

type cryptoClient interface {
	GetUSDTMarketsPrice() (*clients.ExchangeResponse, error)
}

type openaiClient interface {
	GetOpenAIResponse(context, inputText string) (string, error)
}

type replyUseCase struct {
	openaiClient openaiClient
	cryptoClient cryptoClient
}

func NewReply(openaiClient openaiClient, cryptoClient cryptoClient) replyUseCase {
	return replyUseCase{
		openaiClient: openaiClient,
		cryptoClient: cryptoClient,
	}
}

func (ru *replyUseCase) Reply(inputMessage string) (string, error) {

	exchangesValues, err := ru.cryptoClient.GetUSDTMarketsPrice()
	if err != nil {
		fmt.Println("Error getting usdt exchanges values")
	}

	// Convert the struct to JSON
	jsonData, err := json.Marshal(exchangesValues)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	context := domain.NewJarvisContextBuilder(domain.InitialJarvisContext).
		Append("Los valores de usdt son " + string(jsonData)).
		Append(exchangesValues.PropertiesGlossary()).
		Build()

	response, err := ru.openaiClient.GetOpenAIResponse(context, inputMessage)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	return response, nil
}
