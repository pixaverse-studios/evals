package prompt


import "evals/internal/questions"

// PromptMap maps QuestionType to corresponding prompt generation functions
var PromptMap = map[questions.QuestionType]func(string, string, string, string) string{
	questions.EQ:        GetEQPrompt,
	questions.IQ:        GetIQPrompt,
	questions.Learning:  GetLearningPrompt,
	questions.Safety:    GetSafetyPrompt,
	questions.Curiosity: GetCuriosityPrompt,
}

// GetPrompt returns the appropriate prompt based on question type
func GetPrompt(qt questions.QuestionType, childName, childAge, interests, goals string) string {
	if promptFn, ok := PromptMap[qt]; ok {
		return promptFn(childName, childAge, interests, goals)
	}
	return "" // Return empty string for unknown question types
}
