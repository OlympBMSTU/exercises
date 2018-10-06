package db

const (
	insert_excercise = "INSERT INTO EXCERCIESES(author_id, file_name, right_answer, level) VALUES($1, $2, $3, $4)"
	get_tags         = "SELECT name FROM TAGS WHERE subject=$1"
)
