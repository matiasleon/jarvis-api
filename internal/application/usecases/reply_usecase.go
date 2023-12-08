package usecases

import "fmt"

type openaiClient interface {
	GetOpenAIResponse(context, inputText string) (string, error)
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

	context := ru.getContextJarvis()

	response, err := ru.openaiClient.GetOpenAIResponse(context, inputMessage)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	return response, nil
}

func (ru *replyUseCase) getContextJarvis() string {
	return "Sos un asistidor financiero. Te llamas Jarvis. Sos un poco gracioso. El usdt sale 1000 pesos argentinos"
}
