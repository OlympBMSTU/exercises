package db

const (
	ADD_SUBJECT              = "INSERT INTO SUBJECT(name) VALUES($1) RETURNING id"
	GET_SUBJECTS             = "SELECT name FROM SUBJECT"
	DELETE_EXERCISE          = "SELECT del_exercise($1)"
	UPDAET_EXERCISE          = "UPDATE exercise SET "
	GET_EXERCISES_BY_SUBJECT = "select * from EXERCISE ex WHERE ex.subject = $1 "
	INSERT_EXERCISE          = "SELECT add_EXERCISE($1, $2, $3, $4, $5)"
	GET_EXERCISE_BY_ID       = "SELECT * FROM EXERCISE ex WHERE id = $1"
	GET_TAGS_BY_SUBJECT      = "SELECT name FROM tag WHERE subject = $1"
	GET_TAGS_FOR_EXERCISE    = "SELECT tg.name FROM (SELECT * FROM tag_EXERCISE WHERE EXERCISE_id = $1) t join tag tg on (t.tag_id = tg.id)"
	GET_COUNT_SUBJECT        = "SELECT COUNT(*) FROM subject where name = $1"
	UPDATE_TAGS_BY_EX        = "SELECT update_exercise_tags($1, $2, $3, $4, $5)"

	GET_EXERCISE_BY_SUBJECT_AND_TAG = "SELECT ex.* FROM (SELECT * FROM EXERCISE WHERE subject=$1) ex join " +
		"((SELECT id as t_id FROM tag WHERE name=$2 AND subject=$3) t join tag_EXERCISE tg on (tg.tag_id = t.t_id)) tgt on tgt.EXERCISE_id = ex.id "

	GET_SEROND_ROUND_EX = ""

	GET_LIST_SECOND_ROUND_EX = ""

	ADD_SECOND_ROUND_EX = "SELECT add_exercise($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
)

// id serial,
// author_id integer,
// level integer,
// file_name varchar(255),
// subject varchar(255),
// tags varchar(255)[],
// is_broken boolean default false,
// class integer,
// position INTEGER,
// mark   INTEGER,
// type_olymp INTEGER,
// answer jsonb

// firstly sp - is hard way, because we have to check it and other;
// also go slic - ref, so check it
// also test when is empty now its return all to delete
// todo fix procedure and parser
// after that check all update
// is borken request
//SELECT * FROM (select unnest(tags), 'mathematic' subj  from exercise where id =1) tag_names join tag on (name=unnest, subject=subj);
// array of two may fall on cycle
// SELECT * FROM (select unnest(tags), 'mathematic' subj  from exercise where id =1) tag_names join tag on (name=unnest, subject=subj);

// SELECT * FROM (select unnest(tags), 'mathematic' subj  from exercise where id =1) tn join tag on (tn.unnest=tag.name)

// DELETE FROM tag_exercise where tag_id in (SELECT id FROM (select unnest(tags), 'methematic' subj  from exercise where id =1) tn join tag on (tn.unnest=tag.name and tn.subj=tag.subject)) and ex_id=1
// DELETE FROM tag WHERE id in (SELECT t.id FROM tag t LEFT JOIN tag_exercise tg ON (t.id = tg.tag_id AND tg.exercise_id=1)  WHERE tag_id IS NULL);

// DELETE FROM tags WHERE id in (SELECT t.id FROM tag t LEFT JOIN tag_exercise tg ON (t.id = tg.tag_id AND tg.exercise_id=1)  WHERE tag_id IS NULL)   ex_id = '' join SELECT id from tag where name='' and subj='')

// INSERT TAGS
