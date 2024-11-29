package models

import (
	"context"
	"encoding/json"
	"evals/internal/prompts"
	"evals/internal/questions"
	"fmt"
	"os"

	"github.com/joho/godotenv"
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

type EvalResponse struct {
	ScoreModel1 float64 `json:"score_model1"`
	ScoreModel2 float64 `json:"score_model2"`
	Message     string  `json:"msg"`
}

func NewEvalHandler() *EvalHandler {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Warning: Error loading .env file: %v\n", err)
	}
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
	fmt.Printf("🎯 Starting evaluation of responses...\n")

	// Get evaluation prompt based on question type and context
	evalPrompt := prompt.GetPrompt(h.questionType, h.childName, h.childAge, h.interests, h.goals)

	// Combine system prompt and responses into one message
	fullPrompt := fmt.Sprintf("%s\n\nResponse 1:\n%s\n\nResponse 2:\n%s\n\nCompare and rate these responses.", evalPrompt, ans1, ans2)

	resp, err := h.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.O1Preview,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fullPrompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("❌ Error during evaluation: %v\n", err)
		return "", "", "", fmt.Errorf("error getting completion: %v", err)
	}

	evaluation := resp.Choices[0].Message.Content
	
	// Parse JSON response
	var evalResp EvalResponse
	if err := json.Unmarshal([]byte(evaluation), &evalResp); err != nil {
		fmt.Printf("⚠️ Warning: Could not parse JSON response, returning raw evaluation\n")
		return "", "", evaluation, nil // Return full evaluation as message if parsing fails
	}

	fmt.Printf("✨ Evaluation complete! Scores - Model 1: %.1f, Model 2: %.1f\n", evalResp.ScoreModel1, evalResp.ScoreModel2)

	return fmt.Sprintf("%.1f", evalResp.ScoreModel1), fmt.Sprintf("%.1f", evalResp.ScoreModel2), evalResp.Message, nil
}