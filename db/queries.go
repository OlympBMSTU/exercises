package db

const (
	insert_excerciese = "SELECT add_excerciese($1, $2, $3, $4, $5, $6)"

	get_excerciese = "SELECT * FROM EXCERCIESE WHERE id = $1"

	// limit offset order by
	get_excercieses = "select * from (SELECT * FROM Exceriese WHERE subject=$1) ex  join " +
		"((SELECT * FROM tag WHERE name=$2) t join tag_excerciese tg on (tg.tag_id = t.id)) tt on (tt.excerciese_id = ex.id)"

	get_tags = "SELECT name FROM TAG WHERE subject=$1"

	get_subjects = "SELECT subject FROM SUBJECTS"
)

// insert_excercise = "INSERT INTO EXCERCIESE(author_id, file_name, right_answer, level) VALUES($1, $2, $3, $4) RETURNING id"

// subquery for excerciese on tags -> 1
// get_excercise = "SELECT * FROM EXCERCIESES ex JOIN TAG_EXCERCIESE tg ON ex.id = tg.excerciese_id JOIN TAG t ON (t.id = tg.excerciese_id) WHERE id = $1"
