package handlers

type ReplyResponse struct {
	Response string `json:"response"`
}

func toReplResponse(inputText string) ReplyResponse {
	return ReplyResponse{Response: inputText}
}
