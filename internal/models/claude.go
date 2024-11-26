package models

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/liushuangls/go-anthropic/v2"
)

type ClaudeHandler struct {
	client *anthropic.Client
}

func NewClaudeHandler() *ClaudeHandler {
	apiKey := os.Getenv("CLAUDE_API_KEY")
	client := anthropic.NewClient(apiKey)
	return &ClaudeHandler{
		client: client,
	}
}

func (h *ClaudeHandler) ProcessQuestions(questions []string) ([]string, error) {
	var responses []string

	for _, question := range questions {
		resp, err := h.client.CreateMessages(context.Background(), anthropic.MessagesRequest{
			Model: anthropic.ModelClaude3Haiku20240307,
			MaxTokens: 300,
			System: "You are a friendly AI assistant helping children. Keep answers simple, age-appropriate, and encouraging.",
			Messages: []anthropic.Message{
				anthropic.NewUserTextMessage(question),
			},
		})

		if err != nil {
			var apiErr *anthropic.APIError
			if errors.As(err, &apiErr) {
				return nil, fmt.Errorf("claude API error: type=%s message=%s", apiErr.Type, apiErr.Message)
			}
			return nil, fmt.Errorf("error getting completion: %v", err)
		}

		response := fmt.Sprintf("Kid: %s\nAssistant: %s",
			question,
			resp.Content[0].GetText(),
		)
		responses = append(responses, response)
	}

	return responses, nil
}
