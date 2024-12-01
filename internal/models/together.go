package models

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

type TogetherHandler struct {
	client *openai.Client
	model  string
}

func NewTogetherHandler(model string) QuestionProcessor {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Warning: Error loading .env file: %v\n", err)
	}
	apiKey := os.Getenv("TOGETHER_API_KEY")

	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.together.xyz/v1"

	return &TogetherHandler{
		client: openai.NewClientWithConfig(config),
		model:  model,
	}
}

func (h *TogetherHandler) ProcessQuestions(questions []string) ([]string, error) {
	var responses []string

	fmt.Printf("🌟 Starting to process %d questions!\n", len(questions))

	for _, question := range questions {

		resp, err := h.client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: h.model,
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