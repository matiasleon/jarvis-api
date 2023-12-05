package clients

type openAiResponse struct {
	ID      string    `json:"id"`
	Choices []*Choice `json:"choices"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}
