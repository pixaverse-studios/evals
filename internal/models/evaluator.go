package models

import (
	"context"
	"fmt"
	"os"
	"github.com/sashabaranov/go-openai"
)

type EvalHandler struct {
	client *openai.Client
}

func NewEvalHandler() *OpenAIHandler {
	apiKey := os.Getenv("OPENAI_API_KEY")
	return &OpenAIHandler{
		client: openai.NewClient(apiKey),
	}
}
func (h *OpenAIHandler) Benchmark(ans1 string, ans2 string) (score1 string, score2 string, msg string, err error) {
	resp, err := h.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "You are an AI evaluator. Compare two AI responses and rate them on a scale of 1-10. " +
						"Provide scores and brief explanation focusing on clarity, accuracy and helpfulness.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Response 1:\n%s\n\nResponse 2:\n%s\n\nCompare and rate these responses.", ans1, ans2),
				},
			},
		},
	)

	if err != nil {
		return "", "", "", fmt.Errorf("error getting completion: %v", err)
	}

	evaluation := resp.Choices[0].Message.Content
	
	// Extract scores and message from evaluation
	var s1, s2 float64
	_, err = fmt.Sscanf(evaluation, "Response 1: %f\nResponse 2: %f", &s1, &s2)
	if err != nil {
		return "", "", evaluation, nil // Return full evaluation as message if parsing fails
	}

	return fmt.Sprintf("%.1f", s1), fmt.Sprintf("%.1f", s2), evaluation, nil
}
