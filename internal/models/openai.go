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

	for v, question := range questions {
		fmt.Printf("📝 Processing question %d of %d: %s\n", v+1, len(questions), question)

		resp, err := h.client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role: openai.ChatMessageRoleSystem,
						Content: "You are a friendly AI assistant helping children. " +
							"Keep answers simple, age-appropriate, and encouraging.",
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
