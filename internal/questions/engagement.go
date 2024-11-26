package questions

type engagementGenerator struct{}

func GetEngagementQuestions() QuestionGenerator {
	return &engagementGenerator{}
}

func (g *engagementGenerator) GenerateQuestions() []string {
	questions := []string{
		"Why is the sky blue, but sunsets are orange? Does the sun change color?",
		"Let's pretend we're space explorers! What do we need to pack for our trip?",
		"What would happen if it rained candy instead of water? Would plants grow candy?",
		"Why do we have seasons? Does the Earth wear different clothes like we do?",
		"If I dug a hole all the way through the Earth, where would I come out?",
		"Why do stars twinkle but the moon doesn't? Are stars really winking at us?",
		"If animals could talk, what do you think your pet would say to you?",
		"Why does my shadow follow me around? Can I trick it into going somewhere else?",
		"If I jumped on the moon, would I fly like a superhero? Why is that?",
		"Why do birds fly south in the winter? Do they have a map in their heads?",
		"What would happen if the sun didn't rise one day? Would it always be night?",
		"Why do we yawn? Does it help us do something special?",
		"If I could shrink down to the size of an ant, what would the world look like?",
		"How do fish breathe underwater? Do they have secret lungs?",
		"If I plant a candy, will a candy tree grow? Why or why not?",
		"Why does ice float in water? Isn't it heavier than water?",
		"Why do balloons float up, but we can't? Is it because they're magic?",
		"If I could go back in time, what would dinosaurs look like? Could I see one?",
		"Why does lightning come before thunder? Are they racing each other?",
		"What would happen if I lived underwater? Could I grow fins like a fish?",
		"Why do magnets stick to some things but not others? Is it a superpower?",
		"If I could build a robot, what should I teach it to do for me?",
		"Why do leaves change color in autumn? Do trees have feelings about it?",
		"What's inside a bubble? Why does it pop when I touch it?",
		"If the Earth spun faster, would our days be shorter? What would happen to us?",
	}
	return questions
}
