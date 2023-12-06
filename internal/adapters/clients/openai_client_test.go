package clients

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenGetOpenAIResponseInvokeCorrectlyMustReturnSimpleString(t *testing.T) {

	// init
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with a predefined JSON response
		fmt.Fprintln(w, fakeResponse())
	}))
	defer server.Close()

	openaiClient := NewOpenAI(server.URL)

	// execute
	result, err := openaiClient.GetOpenAIResponse("test input")

	// assert
	assert.EqualValues(t, "Hello", result)
	assert.NoError(t, err)
}

func TestWhenGetOpenAIResponseInvokeBadRequestMustReturnError(t *testing.T) {

	// init
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "some bad request desc", http.StatusBadRequest)
		return

	}))
	defer server.Close()

	openaiClient := NewOpenAI(server.URL)

	// execute
	result, err := openaiClient.GetOpenAIResponse("test input")

	// assert
	assert.EqualValues(t, "", result)
	assert.Error(t, err)
}

func fakeResponse() string {
	return `{
		"id": "chatcmpl-123",
		"object": "chat.completion",
		"created": 1677652288,
		"model": "gpt-3.5-turbo-0613",
		"system_fingerprint": "fp_44709d6fcb",
		"choices": [{
			"index": 0,
			"message": {
				"role": "assistant",
				"content": "Hello"
			},
			"finish_reason": "stop"
		}],
		"usage": {
			"prompt_tokens": 9,
			"completion_tokens": 12,
			"total_tokens": 21
		}
	}`
}
