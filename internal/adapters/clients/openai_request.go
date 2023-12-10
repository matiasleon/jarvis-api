package clients

const UserRole = "user"
const System = "system"

type OpenAiRequest struct {
	Model     string     `json:"model"`
	Messages  []*Message `json:"messages"`
	MaxTokens int        `json:"max_tokens"`
}

func createOpenAiRequest(context, inputText string) OpenAiRequest {

	messages := []*Message{
		createSystemMessage(context),
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

func createSystemMessage(context string) *Message {
	return &Message{
		Role:    System,
		Content: context,
	}
}
