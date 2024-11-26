package models


type QuestionProcessor interface {
	ProcessQuestions(questions []string) ([]string, error)
}

type ModelType int

func (m ModelType) String() string {
	switch m {
	case OpenAI:
		return "OpenAI"
	case Claude:
		return "Claude" 
	case Llama:
		return "Llama"
	default:
		return "Unknown"
	}
}


const (
	OpenAI ModelType = iota
	Claude
	Llama
)

var Models = map[ModelType]QuestionProcessor{
	OpenAI: NewOpenAIHandler(),
	Claude: NewClaudeHandler(),
}
