package questions


type IqEval struct {
	Question string
	Options  []string
	Answer   int
}

func GetIqEvalQuestions() []IqEval {
	return []IqEval{
		{
			Question: "Which number should come next in the pattern? 2, 4, 8, 16, ...",
			Options:  []string{"24", "32", "28", "20"},
			Answer:   1, // 32
		},
		{
			Question: "Complete the analogy: Book is to Reading as Fork is to:",
			Options:  []string{"Drawing", "Writing", "Eating", "Walking"},
			Answer:   2, // Eating
		},
		{
			Question: "Which figure completes the pattern? [Visual Pattern Question]",
			Options:  []string{"Square", "Circle", "Triangle", "Rectangle"},
			Answer:   0, // Square
		},
		{
			Question: "If all Zips are Zaps, and some Zaps are Zops, then:",
			Options:  []string{
				"All Zips are definitely Zops",
				"Some Zips might be Zops",
				"No Zips are Zops",
				"All Zops are Zips",
			},
			Answer: 1, // Some Zips might be Zops
		},
		{
			Question: "What number is missing? 7, 12, 19, 28, ?",
			Options:  []string{"35", "39", "42", "45"},
			Answer:   1, // 39
		},
	}
}
