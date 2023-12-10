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

	// clients
	openAiClient := clients.NewOpenAI(clients.OpenAiUrl)
	cryptoClient := clients.NewCrypto(clients.CryptoMarketEndpoint)

	replyUseCase := usecases.NewReply(&openAiClient, &cryptoClient)
	replyHandler := handlers.NewReply(&replyUseCase)

	http.HandleFunc("/reply", replyHandler.Handle)
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode("pong")
	})

	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(port, nil)
}
