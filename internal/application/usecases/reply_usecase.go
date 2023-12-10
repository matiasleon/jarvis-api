package usecases

import (
	"fmt"
	"main/internal/application/domain"
	"strconv"
)

type cryptoClient interface {
	GetUSDTPrice() (float64, error)
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

	usdtValue, err := ru.cryptoClient.GetUSDTPrice()

	if err != nil {
		fmt.Println("Error getting usdt price")
	}

	stringValue := strconv.FormatFloat(usdtValue, 'f', -1, 64)
	context := domain.NewJarvisContextBuilder(domain.InitialJarvisContext).Append("usdt vale" + stringValue).Build()

	response, err := ru.openaiClient.GetOpenAIResponse(context, inputMessage)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	return response, nil
}
