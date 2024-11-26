package models

type ModelType int

const (
	OpenAI ModelType = iota // OpenAI models like GPT-3.5, GPT-4
	Claude                  // Anthropic's Claude models
	Llama                   // Meta's Llama models
)

// String returns the string representation of ModelType
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

// FromString converts a string to ModelType
func FromString(s string) ModelType {
	switch s {
	case "OpenAI":
		return OpenAI
	case "Claude":
		return Claude
	case "Llama":
		return Llama
	default:
		return ModelType(-1) // Invalid model type
	}
}

type QuestionProcessor interface {
	ProcessQuestions(questions []string) ([]string, error)
}

var Models = map[ModelType]QuestionProcessor{
	OpenAI: NewOpenAIHandler(),
	Claude: NewClaudeHandler(),
}
