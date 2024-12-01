package models

type ModelType int

const (
	OpenAI ModelType = iota // OpenAI models like GPT-3.5, GPT-4
	Claude                  // Anthropic's Claude models
	Llama                   // Meta's Llama models including Llama 3.1 405B
	Qwen                    // Qwen 32B model (preview)
	Nemotron                // Nemotron models
	Gemini                  // Google's Gemini models
	Phi                     // Microsoft's Phi 3.5B model
	Mistral                 // Mistral AI models
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
	case Qwen:
		return "Qwen"
	case Nemotron:
		return "Nemotron"
	case Gemini:
		return "Gemini"
	case Phi:
		return "Phi"
	case Mistral:
		return "Mistral"
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
	case "Qwen":
		return Qwen
	case "Nemotron":
		return Nemotron
	case "Gemini":
		return Gemini
	case "Phi":
		return Phi
	case "Mistral":
		return Mistral
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
	Llama: NewTogetherHandler("meta-llama/Meta-Llama-3.1-70B-Instruct-Turbo"),
	Qwen: NewTogetherHandler("Qwen/QwQ-32B-Preview"),
	Mistral: NewTogetherHandler("mistralai/Mixtral-8x22B-Instruct-v0.1"),
	Nemotron: NewTogetherHandler("nvidia/Llama-3.1-Nemotron-70B-Instruct-HF"),

}
