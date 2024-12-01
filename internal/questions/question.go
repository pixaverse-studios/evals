package questions

type QuestionType int

const (
	EQ QuestionType = iota // 20 questions
	IQ                     
	Learning               
	Safety                 
	Curiosity              
)

// String returns the string representation of QuestionType
func (qt QuestionType) String() string {
	switch qt {
	case EQ:
		return "EQ"
	case IQ:
		return "IQ" 
	case Learning:
		return "Learning"
	case Safety:
		return "Safety"
	case Curiosity:
		return "Curiosity"
	default:
		return "Unknown"
	}
}

// FromString converts a string to QuestionType
func FromString(s string) QuestionType {
	switch s {
	case "EQ":
		return EQ
	case "IQ":
		return IQ
	case "Learning":
		return Learning
	case "Safety":
		return Safety
	case "Curiosity":
		return Curiosity
	default:
		return QuestionType(-1) // Invalid question type
	}
}

type QuestionGenerator interface {
	GenerateQuestions() []string
}

var QuestionGenerators = map[QuestionType]QuestionGenerator{
	EQ:        NewEQGenerator(),
	IQ:        NewIQGenerator(), 
	Learning:  NewLearningGenerator(),
	Safety:    NewSafetyGenerator(),
	Curiosity: GetEngagementQuestions(),
}
