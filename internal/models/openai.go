package models

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

type OpenAIHandler struct {
	client *openai.Client
}

func NewOpenAIHandler() QuestionProcessor {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Warning: Error loading .env file: %v\n", err)
	}
	apiKey := os.Getenv("OPENAI_API_KEY")

	return &OpenAIHandler{
		client: openai.NewClient(apiKey),
	}
}

func (h *OpenAIHandler) ProcessQuestions(questions []string) ([]string, error) {
	var responses []string

	fmt.Printf("🌟 Starting to process %d questions!\n", len(questions))

	for _, question := range questions {

		resp, err := h.client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT4o20240513,
				MaxTokens: 300,
				Messages: []openai.ChatCompletionMessage{
					{
						Role: openai.ChatMessageRoleSystem,
						Content: Prompt(),
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

	fmt.Printf("✨ Successfully answered all %d questions!\n", len(questions))

	return responses, nil
}
