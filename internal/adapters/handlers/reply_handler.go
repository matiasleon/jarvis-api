package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

type replyUseCase interface {
	Reply(inputMessage string) (response string, err error)
}

type ReplyHandler struct {
	replyUseCase replyUseCase
}

func NewReply(replyUseCase replyUseCase) *ReplyHandler {
	return &ReplyHandler{replyUseCase: replyUseCase}
}

func (ah *ReplyHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	inputText := string(body)

	response, err := ah.replyUseCase.Reply(inputText)
	if err != nil {
		http.Error(w, "Error getting response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(toReplResponse(response))
}
