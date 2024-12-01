package models

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/liushuangls/go-anthropic/v2"
)

type ClaudeHandler struct {
	client *anthropic.Client
}

// NewClaudeHandler creates a new Claude handler that implements QuestionProcessor
func NewClaudeHandler() QuestionProcessor {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Warning: Error loading .env file: %v\n", err)
	}
	apiKey := os.Getenv("CLAUDE_API_KEY")
	client := anthropic.NewClient(apiKey)
	return &ClaudeHandler{
		client: client,
	}
}

func (h *ClaudeHandler) ProcessQuestions(questions []string) ([]string, error) {
	var responses []string

	fmt.Printf("🌟 Starting to process %d questions!\n", len(questions))

	for v, question := range questions {
		fmt.Printf("📝 Processing question %d of %d: %s\n", v+1, len(questions), question)
		resp, err := h.client.CreateMessages(context.Background(), anthropic.MessagesRequest{
			Model: anthropic.ModelClaude3Dot5Sonnet20241022,//claude-3-5-sonnet-20241022
			MaxTokens: 300,
			System: Prompt(),
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
		fmt.Printf("🤖 %s\n", response)
		responses = append(responses, response)
	}

	fmt.Printf("✨ Successfully answered all %d questions!\n", len(questions))



	return responses, nil
}
