package models


type QuestionProcessor interface {
	ProcessQuestions(questions []string) ([]string, error)
}

type ModelType int

const (
	OpenAI ModelType = iota
	Claude
	Llama
)

var Models = map[ModelType]QuestionProcessor{
	OpenAI: NewOpenAIHandler(),
	Claude: NewClaudeHandler(),
}
