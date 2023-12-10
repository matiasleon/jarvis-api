package usecases

import (
	"fmt"
	"main/internal/application/domain"
)

type cryptoClient interface {
	GetUSDTPrice() (int, error)
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

	context := domain.NewJarvisContextBuilder(domain.InitialJarvisContext)
	contextString := context.Build()

	response, err := ru.openaiClient.GetOpenAIResponse(contextString, inputMessage)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	return response, nil
}
