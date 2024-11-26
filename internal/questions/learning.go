package questions

type learningGenerator struct{}

func NewLearningGenerator() QuestionGenerator {
	return &learningGenerator{}
}

func (g *learningGenerator) GenerateQuestions() []string {
	questions := []string{
		"How can I learn to play football? Does the ball look like a planet?",
		"I want to draw a car. How do I make it with shapes like circles and squares?",
		"How can I ride my bike faster? Is wind pushing against me like when I run?",
		"I want to count my toys faster. How can I group them to count quickly?",
		"How can I grow a plant at home? Does it need water like I need food?",
		"Why do frogs jump so high? Can I learn how to jump like them?",
		"How can I throw a ball better? Does gravity make it fall?",
		"Why do some boats float and others sink? Can I try making a paper boat?",
		"How do I bake cookies? Is mixing the dough like a science experiment?",
		"How can I learn to sing like a bird? Does sound travel like my voice?",
		"How can I build a tall tower of blocks without it falling?",
		"Why do stars twinkle? Can I make a star craft to learn more?",
		"How can I run faster? Do my legs work like springs?",
		"How can I make my writing neater? Do I use shapes when I write?",
		"How can I tie my shoes by myself? Is it like weaving a thread?",
		"Why do rainbows have so many colors? Can I make a rainbow at home?",
		"How can I practice adding big numbers? Can I count candies to help?",
		"How can I learn to read a clock? Does it have something to do with counting by fives?",
		"How can I learn the names of the planets? Is there a fun song to help me remember?",
		"Why does a paper plane fly? Does how I fold it make it go farther?",
		"How can I learn to whistle? Is it about blowing air just right?",
		"Why does my shadow change size? Can I make funny shapes with it?",
		"How can I stack cups super fast? Is it like solving a puzzle?",
		"Why does the moon look different each night? Can I draw what I see?",
		"How can I learn to skip a rock on water? Is it about how the rock spins?",
	}
	return questions
}
