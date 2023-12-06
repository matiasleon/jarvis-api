package clients

const UserRole = "user"

type OpenAiRequest struct {
	Model     string     `json:"model"`
	Messages  []*Message `json:"messages"`
	MaxTokens int        `json:"max_tokens"`
}

func createOpenAiRequest(inputText string) OpenAiRequest {

	messages := []*Message{
		createUserMessage(inputText),
	}

	return OpenAiRequest{
		Model:     OpenAIModel,
		Messages:  messages,
		MaxTokens: 500,
	}
}

func createUserMessage(inputText string) *Message {
	return &Message{
		Role:    UserRole,
		Content: inputText,
	}
}
