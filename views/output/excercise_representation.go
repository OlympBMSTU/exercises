package output

type OutputExercieseView struct {
	Id       int      `json:"id"`
	FilePath string   `json:"file_path"`
	Subject  string   `json:"subject"`
	Tags     []string `json:"tags"`
	Level    int      `json:"level"`
	Author   int      `json:"author"`
}
