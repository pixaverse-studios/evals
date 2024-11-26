package eval

import (
	"evals/internal/models"
	"evals/internal/questions"
	"fmt"
	"strings"
)

type Client struct {
	model1 models.ModelType
	model2 models.ModelType
}

func Engine(
	modelOne models.ModelType,
	modelTwo models.ModelType,
) *Client {
	return &Client{
		model1: modelOne,
		model2: modelTwo,
	}
}

func (c *Client) evaluate(qt questions.QuestionType) (score1 string, score2 string, msg string, err error) {
	ques := questions.QuestionGenerators[qt].GenerateQuestions()

	// Get responses from both models
	responses1, err := models.Models[c.model1].ProcessQuestions(ques)
	if err != nil {
		return "", "", "", fmt.Errorf("error processing questions with model 1: %w", err)
	}

	responses2, err := models.Models[c.model2].ProcessQuestions(ques)
	if err != nil {
		return "", "", "", fmt.Errorf("error processing questions with model 2: %w", err)
	}

	ans1 := strings.Join(responses1, "\n")
	ans2 := strings.Join(responses2, "\n")
    
	evaluator := models.NewEvalHandler()

	return evaluator.Benchmark(ans1, ans2)
}

func (c *Client) EvaluateIqBenchmarks() (score1 string, score2 string, msg string, err error) {
	return c.evaluate(questions.IQ)
}

func (c *Client) EvaluateEqBenchmarks() (score1 string, score2 string, msg string, err error) {
	return c.evaluate(questions.EQ)
}

func (c *Client) EvaluateLearningBenchmarks() (score1 string, score2 string, msg string, err error) {
	return c.evaluate(questions.Learning)
}

func (c *Client) EvaluateSafetyBenchmarks() (score1 string, score2 string, msg string, err error) {
	return c.evaluate(questions.Safety)
}