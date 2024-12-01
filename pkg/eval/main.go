package eval

import (
	"evals/internal/models"
	"evals/internal/questions"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

type ChildData struct {
	Name      string
	Age       string
	Interests string
	Goals     string
}

type Client struct {
	model     models.ModelType
	childData ChildData
}

func Engine(
	model models.ModelType,
	childData ChildData,
) *Client {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Warning: Error loading .env file: %v\n", err)
	}
	return &Client{
		model:     model,
		childData: childData,
	}
}

func (c *Client) evaluate(qt questions.QuestionType) (score string, msg string, err error) {
	ques := questions.QuestionGenerators[qt].GenerateQuestions()

	responses, err := models.Models[c.model].ProcessQuestions(ques)
	if err != nil {
		return "", "", fmt.Errorf("error processing questions: %w", err)
	}

	ans := strings.Join(responses, "\n")
    
	evaluator := models.NewEvalHandler()
	evaluator.SetContext(qt, c.childData.Name, c.childData.Age, c.childData.Interests, c.childData.Goals)

	return evaluator.Evaluate(ans)
}

func (c *Client) EvaluateIqBenchmarks() (score string, msg string, err error) {
	return c.evaluate(questions.IQ)
}

func (c *Client) EvaluateEqBenchmarks() (score string, msg string, err error) {
	return c.evaluate(questions.EQ)
}

func (c *Client) EvaluateLearningBenchmarks() (score string, msg string, err error) {
	return c.evaluate(questions.Learning)
}

func (c *Client) EvaluateSafetyBenchmarks() (score string, msg string, err error) {
	return c.evaluate(questions.Safety)
}

func (c *Client) EvaluateCuriosityBenchmarks() (score string, msg string, err error) {
	return c.evaluate(questions.Curiosity)
}

