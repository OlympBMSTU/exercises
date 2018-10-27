package db

const (
	ADD_SUBJECT  = "INSERT INTO SUBJECT(name) VALUES($1)"
	GET_SUBJECTS = "SELECT name FROM SUBJECT"

	GET_EXERCISES_BY_SUBJECT = "select * from Excerciese ex WHERE subject = $1 "
	INSERT_EXERCISE          = "SELECT add_excerciese($1, $2, $3, $4, $5, $6)"
	GET_EXERCISE_BY_ID       = "SELECT * FROM EXCERCIESE ex WHERE id = $1"
	GET_TAGS_BY_SUBJECT      = "SELECT name FROM tag WHERE subject = $1"
	GET_TAGS_FOR_EXERCISE    = "SELECT tg.name FROM (SELECT * FROM tag_excerciese WHERE excerciese_id = $1) t join tag tg on (t.tag_id = tg.id)"

	GET_EXERCISE_BY_SUBJECT_AND_TAG = "SELECT ex.* FROM (SELECT * FROM Excerciese WHERE subject=$1) ex join " +
		"((SELECT id as t_id FROM tag WHERE name=$2 AND subject=$3) t join tag_excerciese tg on (tg.tag_id = t.t_id)) tgt on tgt.excerciese_id = ex.id "
)
