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
	Score   float64 `json:"score"`
	Message string  `json:"msg"`
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

func (h *EvalHandler) Evaluate(ans string) (score string, msg string, err error) {
	fmt.Printf("🎯 Starting evaluation of response...\n")

	// Get evaluation prompt based on question type and context
	evalPrompt := prompt.GetPrompt(h.questionType, h.childName, h.childAge, h.interests, h.goals)
    fmt.Print(evalPrompt)
	// Combine system prompt and response into one message
	fullPrompt := fmt.Sprintf("%s\n\nResponse:\n%s\n\nEvaluate this response.", evalPrompt, ans)

	resp, err := h.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.O1Preview20240912,
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
		return "", "", fmt.Errorf("error getting completion: %v", err)
	}
    
	evaluation := resp.Choices[0].Message.Content

	fmt.Print(evaluation)
	
	var evalResp EvalResponse
	if err := json.Unmarshal([]byte(evaluation), &evalResp); err != nil {
		fmt.Printf("⚠️ Warning: Could not parse JSON response, returning raw evaluation\n")
		return "", evaluation, nil // Return full evaluation as message if parsing fails
	}

	fmt.Printf("✨ Evaluation complete! Score: %.1f\n", evalResp.Score)

	return fmt.Sprintf("%.1f", evalResp.Score), evalResp.Message, nil
}