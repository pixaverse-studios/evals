package models

import (
	"context"
	"fmt"
	"os"
	"github.com/sashabaranov/go-openai"
)

type OpenAIHandler struct {
	client *openai.Client
}

func NewOpenAIHandler() *OpenAIHandler {
	apiKey := os.Getenv("OPENAI_API_KEY")
	return &OpenAIHandler{
		client: openai.NewClient(apiKey),
	}
}

func (h *OpenAIHandler) ProcessQuestions(questions []string) ([]string, error) {
	var responses []string

	for _, question := range questions {
		resp, err := h.client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role: openai.ChatMessageRoleSystem,
						Content: "You are a friendly AI assistant helping children. " +
							"Keep answers simple, age-appropriate, and encouraging. " ,
					},
					{
						Role:    openai.ChatMessageRoleUser,
						Content: question,
					},
				},
			},
		)

		if err != nil {
			return nil, fmt.Errorf("error getting completion: %v", err)
		}

		response := fmt.Sprintf("Kid: %s\nAssistant: %s", 
			question,
			resp.Choices[0].Message.Content,
		)
		responses = append(responses, response)
	}

	return responses, nil
}
