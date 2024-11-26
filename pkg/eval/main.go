package eval

import (
	"evals/internal/models"
	"evals/internal/questions"
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

func (c*Client) evaluate(qt questions.QuestionType) (err error, score1 string, score2 string, msg string) {
	ques := questions.QuestionGenerators[qt].GenerateQuestions()
    
	evaluationOne := models.QuestionProcessor[c.model1]

}




func (c*Client) EvaluateIqBenchmarks() {

}