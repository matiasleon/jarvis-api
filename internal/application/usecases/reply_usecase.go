package usecases

import "fmt"

type openaiClient interface {
	GetOpenAIResponse(inputText string) (string, error)
}

type replyUseCase struct {
	openaiClient openaiClient
}

func NewReply(openaiClient openaiClient) replyUseCase {
	return replyUseCase{openaiClient: openaiClient}
}

func (ru *replyUseCase) Reply(inputMessage string) (string, error) {

	// anteriormente podemos obtener la "conciencia" de jarvis
	// osea el contexto para open ai
	// obtener tambien todos el historico de mensajes anteriores
	response, err := ru.openaiClient.GetOpenAIResponse(inputMessage)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	return response, nil
}
