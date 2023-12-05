package main

import (
	"encoding/json"
	"fmt"
	"main/internal/adapters/clients"
	"main/internal/adapters/handlers"
	"main/internal/application/usecases"
	"net/http"
)

func main() {

	openAiClient := clients.NewOpenAI(clients.OpenAiUrl)
	replyUseCase := usecases.NewReply(&openAiClient)
	replyHandler := handlers.NewReply(&replyUseCase)

	http.HandleFunc("/reply", replyHandler.Handle)
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode("pong")
	})

	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(port, nil)
}
