package views

type AnswerView struct {
	ID     int      `json:"id"`
	Input  []string `json:"input"`
	Output []string `json:"output"`
}
