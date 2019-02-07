package db

const (
	ADD_SUBJECT         = "INSERT INTO SUBJECT(name) VALUES($1) RETURNING id"
	GET_SUBJECTS        = "SELECT name FROM SUBJECT"
	DELETE_EXERCISE     = "SELECT del_exercise($1)"
	UPDAET_EXERCISE     = "UPDATE exercise SET "
	INSERT_EXERCISE     = "SELECT add_exercise($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	GET_TAGS_BY_SUBJECT = "SELECT name FROM tag WHERE subject = $1"
	GET_COUNT_SUBJECT   = "SELECT COUNT(*) FROM subject where name = $1"
	GET_EXERCISE_BY_ID  = "SELECT * FROM EXERCISE ex WHERE id = $1"

	GET_EXERCISES_BY_SUBJECT = "select * from EXERCISE ex WHERE ex.subject = $1 "

	UPDATE_TAGS_BY_EX = "SELECT update_exercise_tags($1, $2, $3, $4, $5)"

	GET_EXERCISE_BY_SUBJECT_AND_TAG = "SELECT ex.* FROM (SELECT * FROM EXERCISE WHERE subject=$1) ex join " +
		"((SELECT id as t_id FROM tag WHERE name=$2 AND subject=$3) t join tag_EXERCISE tg on (tg.tag_id = t.t_id)) tgt on tgt.EXERCISE_id = ex.id "
)

// INSERT_EXERCISE          = "SELECT add_EXERCISE($1, $2, $3, $4, $5)"
// GET_TAGS_FOR_EXERCISE    = "SELECT tg.name FROM (SELECT * FROM tag_EXERCISE WHERE EXERCISE_id = $1) t join tag tg on (t.tag_id = tg.id)"
