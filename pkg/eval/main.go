package eval

import (
	"evals/internal/models"
	"evals/internal/questions"
	"fmt"
	"strings"
)

type ChildData struct {
	Name      string
	Age       string
	Interests string
	Goals     string
}

type Client struct {
	model1    models.ModelType
	model2    models.ModelType
	childData ChildData
}

func Engine(
	modelOne models.ModelType,
	modelTwo models.ModelType,
	childData ChildData,
) *Client {
	return &Client{
		model1:    modelOne,
		model2:    modelTwo,
		childData: childData,
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
	evaluator.SetContext(qt, c.childData.Name, c.childData.Age, c.childData.Interests, c.childData.Goals)

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