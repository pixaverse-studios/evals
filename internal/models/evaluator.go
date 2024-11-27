package models

import (
	"context"
	"fmt"
	"os"
	"evals/internal/questions"
	"evals/internal/prompts"
	"github.com/sashabaranov/go-openai"
)

type EvalHandler struct {
	client *openai.Client
	questionType questions.QuestionType
	childName string
	childAge string 
	interests string
	goals string
}

func NewEvalHandler() *EvalHandler {
	apiKey := os.Getenv("OPENAI_API_KEY")
	return &EvalHandler{
		client: openai.NewClient(apiKey),
	}
}

func (h *EvalHandler) SetContext(qt questions.QuestionType, name, age, interests, goals string) {
	h.questionType = qt
	h.childName = name
	h.childAge = age
	h.interests = interests 
	h.goals = goals
}

func (h *EvalHandler) Benchmark(ans1 string, ans2 string) (score1 string, score2 string, msg string, err error) {
	// Get evaluation prompt based on question type and context
	evalPrompt := prompt.GetPrompt(h.questionType, h.childName, h.childAge, h.interests, h.goals)

	resp, err := h.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.O1Preview,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: evalPrompt,
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
