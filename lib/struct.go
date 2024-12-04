package libReadWord

type ParaQuiz struct {
	Para string
	Type string
	Quiz []Question
}

type Question struct {
	Subject       string
	Type          string
	Options       []string
	Solution      string
	CorrectOption int
}
