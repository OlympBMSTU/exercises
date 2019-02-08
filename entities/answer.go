package entities

type Answer struct {
	ID     int      `json:"id"`
	Input  []string `json:"input"`
	Output []string `json:"output"`
	Mark   int      `json:"mark"`
}
