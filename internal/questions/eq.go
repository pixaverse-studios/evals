package questions

type eqGenerator struct{}

func NewEQGenerator() QuestionGenerator {
	return &eqGenerator{}
}

func (g *eqGenerator) GenerateQuestions() []string {
	questions := []string{
		// 20 questions
		"I saw my mom smiling when dad told her, 'I love you.' Why does that make her happy?",
		"My best friend didn't sit with me at lunch today. Does that mean they're upset with me?",
		"I saw a kid crying because they dropped their ice cream. How can I help them feel better?",
		"My brother gets mad when I use his toys without asking. Why does that make him angry?",
		"When I shared my favorite snack, my friend said, 'Thank you!' Why did that make me feel good?",
		"My teacher said sorry to me when she forgot my name. Why is saying sorry important?",
		"I saw a kid at school sitting alone during recess. Do you think they might feel lonely?",
		"A kid in my class always makes fun of me. Why are they being mean to me?",
		"My grandma hugs me really tight whenever I visit her. Why does hugging show love?",
		"I was scared to go on stage, but my dad said, 'You're brave!' What does being brave mean?",
		"I saw two people getting married in a park. Why do people get married?",
		"My little sister hides under her blanket when it's dark. Why is she afraid at night?",
		"We had a big party for my birthday last week. Why do birthdays make people so happy?",
		"My friend looked sad when I didn't invite them to play. How can I make it up to them?",
		"I feel butterflies in my stomach before a big test. What should I do to feel better?",
		"I gave my teacher a drawing, and she said, 'This is beautiful!' Why does saying nice things matter?",
		"I saw a kid getting teased on the playground. What should I do to help them?",
		"My friend told me a big secret. Why do people share their secrets with someone they trust?",
		"When I broke my friend's toy, they didn't talk to me. How can I make things right?",
		"I saw a bird with a broken wing, and it made me feel sad. Why is it important to care for animals?",
	}
	return questions
}
