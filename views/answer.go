package views

type AnswerView struct {
	ID     int      `json:"id"`
	Input  []string `json:"input"`
	Output []string `json:"output"`
	Mark   int      `json:"mark"`
}

// func (view *AnswerView) UnmarshalJSON(data []byte) error {
// 	ans := AnswerView{}
// 	err := json.Unmarshal(data, ans)
// 	if err != nil {
// 		return err
// 	}

// 	view.ID = ans.ID
// 	view.Input = ans.Input
// 	view.Output = ans.Output
// 	return nil
// }
