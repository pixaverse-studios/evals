package questions

type iqGenerator struct{}

func NewIQGenerator() QuestionGenerator {
	return &iqGenerator{}
}

func (g *iqGenerator) GenerateQuestions() []string {
	return []string{
		"How can I build a paper bridge that can hold my toy car?",
		"Why do birds fly in the sky and not walk everywhere?",
		"What's a fun way to organize my books by color or size?",
		"How does a rainbow appear in the sky after it rains?",
		"What can I make with small crayon pieces that are too short to use?",
		"Why do I sometimes get hiccups when I eat fast?",
		"What's a quick way to count candies in a jar?",
		"How can I build a Lego tower that doesn't fall over easily?",
		"Why does the sea taste salty and not like plain water?",
		"What's a cool trick to remember what I need to bring to school?",
		"How can I tell if an egg is cooked without cracking it?",
		"Why does ice float in a glass of water?",
		"How can I guess the height of a tree by looking at its shadow?",
		"What happens to leaves when they fall off a tree?",
		"How can I pack my school lunchbox so everything fits neatly?",
		"Why do cats swish their tails when they're playing?",
		"Can I make a game using only paper and dice? How?",
		"What's the easiest way to draw a big circle without a compass?",
		"Why does popcorn pop when it's heated?",
		"How do astronauts sleep when they're floating in space?",
		"What do I need to do to grow a little plant in my room?",
		"Why do mirrors show left and right swapped but not upside down?",
		"What's the quickest way to count my toy cars in a pile?",
		"How can I find which way is north if I'm in a park without a map?",
		"Why do birds sing when the sun comes up?",
	}
}
