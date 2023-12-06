package clients

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var OpenAiUrl = "https://api.openai.com/v1/chat/completions"
var OpenAIModel = "gpt-3.5-turbo"

type openAIClient struct {
	endpoint string
	client   *http.Client
}

func NewOpenAI(endpoint string) openAIClient {
	client := &http.Client{}
	return openAIClient{endpoint: endpoint, client: client}
}

func (oc *openAIClient) GetOpenAIResponse(inputText string) (string, error) {
	// create request body
	requestBody := createOpenAiRequest(inputText)
	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return "", err
	}

	// create http request
	request, err := http.NewRequest("POST", oc.endpoint, bytes.NewReader(requestJson))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	// Set the Content-Type header
	request.Header.Set("Content-Type", "application/json")

	// Set additional headers as needed
	apiKey := os.Getenv("OPEN_AI_KEY")
	request.Header.Set("Authorization", "Bearer "+apiKey)

	//do api call
	response, err := oc.client.Do(request)

	if err != nil {
		fmt.Println("Error making request:", err)
		return "", err
	}

	responseString, err := oc.handleResponse(*response)

	if err != nil {
		fmt.Println("Error handling response:", err)
		return "", err
	}

	defer response.Body.Close()

	return responseString, nil
}

func (oc *openAIClient) handleResponse(response http.Response) (string, error) {

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error calling openai")

		jsonData, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		fmt.Println("Status code: " + response.Status + "desc: " + string(jsonData))

		return "", errors.New("error calling OPEN AI")
	}

	return oc.handleOkeyResponse(response)
}

func (oc *openAIClient) handleOkeyResponse(response http.Response) (string, error) {
	var openAiResponse openAiResponse
	err := json.NewDecoder(response.Body).Decode(&openAiResponse)

	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return "", nil
	}

	if len(openAiResponse.Choices) == 0 {
		return "empty", nil
	}

	return openAiResponse.Choices[0].Message.Content, nil
}
