package handlers

import (
	"encoding/json"
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

	var messageInput MessageInput
	err := json.NewDecoder(r.Body).Decode(&messageInput)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	response, err := ah.replyUseCase.Reply(messageInput.Message)
	if err != nil {
		http.Error(w, "Error getting response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(toReplResponse(response))
}
