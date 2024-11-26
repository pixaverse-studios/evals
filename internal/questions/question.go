package questions



type QuestionType int

const (
	EQ QuestionType = iota
	IQ
	Learning
	Safety
)

type QuestionGenerator interface {
	GenerateQuestions() []string
}

var QuestionGenerators = map[QuestionType]QuestionGenerator{
	EQ:       GetEngagementQuestions(),
	IQ:       NewIQGenerator(), 
	Learning: NewLearningGenerator(),
	Safety:   NewSafetyGenerator(),
}



