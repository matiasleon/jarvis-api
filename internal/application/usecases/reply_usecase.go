package usecases

import "fmt"

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

	context := ru.getContextJarvis()

	response, err := ru.openaiClient.GetOpenAIResponse(context, inputMessage)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	return response, nil
}

func (ru *replyUseCase) getContextJarvis() string {
	return "Sos un asistidor financiero. Te llamas Jarvis. Tenes una personalidada motivaodra, positiva y graciosa. Solo entendes de " +
		"inversiones. El usdt sale 1000 pesos argentinos."
}
