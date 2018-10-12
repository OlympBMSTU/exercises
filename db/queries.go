package db

const (
	ADD_SUBJECT  = "INTSERT INTO SUBJECT(name) VALUES($1)"
	GET_SUBJECTS = "SELECT name FROM SUBJECT"

	GET_BY_SUBJECT       = "select * from Excerciese ex WHERE subject = $1 "
	INSERT_EXCERCIESE    = "SELECT add_excerciese($1, $2, $3, $4, $5, $6)"
	GET_EXCERCIESE_BY_ID = "SELECT * FROM EXCERCIESE ex WHERE id = $1"
	insert_excerciese    = "SELECT add_excerciese($1, $2, $3, $4, $5, $6)"

	GET_EXCERCIESE_BY_SUBJECT_AND_TAG = "SELECT ex.* FROM (SELECT * FROM Excerciese WHERE subject=$1) ex join " +
		"((SELECT id as t_id FROM tag WHERE name=$2 AND subject=$3) t join tag_excerciese tg on (tg.tag_id = t.t_id)) tgt on tgt.excerciese_id = ex.id "

	get_excerciese = "SELECT * FROM EXCERCIESE ex WHERE id = $1"

	get_tags_for_excerciese = "SELECT tg.name FROM (SELECT * FROM tag_excerciese WHERE excerciese_id = $1) t join tag tg on (t.tag_id = tg.id)"
	get_by_subject          = "select * from Excerciese ex WHERE subject = '%s' "

	get_by_subject_tag = "SELECT ex.* FROM (SELECT * FROM Excerciese WHERE subject='%s') ex join " +
		"((SELECT id as t_id FROM tag WHERE name='%s' AND subject='%s') t join tag_excerciese tg on (tg.tag_id = t.t_id)) tgt on tgt.excerciese_id = ex.id "

	get_tags = "SELECT name FROM TAG WHERE subject=$1"

	get_tags_by_subject = "SELECT name FROM tag WHERE subject = $1"
)
