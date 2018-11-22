-- index - subject,
-- index subject, level
-- tag index - name, subject name
-- 

create table if not exists exercise (
    id serial, 
    author_id integer,
    level integer,
    file_name varchar(255),
    subject varchar(255),
    tags varchar(255)[],
    is_wrong boolean default false
);

create table if not exists tag (
    id serial,
    subject varchar(255),
    name varchar(255),
    CONSTRAINT subjec_tag UNIQUE (subject, name)
);

create table if not exists tag_exercise(
    id serial,
    exercise_id integer,
    tag_id integer, 
    CONSTRAINT ex_id_tag_id UNIQUE (exercise_id, tag_id)
);


-- index lower name
create table if not exists subject(
    id serial,
    name varchar(255) UNIQUE
);


-- lower
-- todo ON CONFLCIT делать и проверка на id tag 
CREATE OR REPLACE FUNCTION add_exercise(auth_id integer, lev integer, f_name varchar(255), subj varchar(255), tags varchar(255)[])
RETURNS integer AS $$
DECLARE ex_id integer;
DECLARE t_id INTEGER;
DECLARE tag text;
DECLARE subj_id INTEGER;
DECLARE tag_ex_id INTEGER;
DECLARE data text;
BEGIN
    -- check that this subject exists
    SELECT id FROM SUBJECT WHERE lower(name)=lower(subj) INTO subj_id;
    IF subj_id IS NULL THEN
        RETURN -1;
    END IF;

    INSERT INTO EXERCISE(author_id, file_name, level, subject, tags) VALUES(auth_id, f_name, lev, subj, tags) RETURNING id INTO ex_id;
    FOR i IN 1..array_length(tags, 1) LOOP
        SELECT id from tag where subject = subj and name = tags[i] into t_id;
        if t_id IS null then 
            INSERT INTO TAG(subject, name) VALUES(subj, tags[i]) RETURNING id INTO t_id;
        end if;

        SELECT id FROM tag_exercise where exercise_id = ex_id AND tag_id = t_id into tag_ex_id;
        IF tag_ex_id IS NULL THEN 
            INSERT INTO TAG_EXERCISE(tag_id, exercise_id) VALUES(t_id, ex_id);         
        END IF;
    END LOOP;

    RETURN ex_id;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION del_exercise(ex_id integer) 
RETURNS INTEGER AS $$
DECLARE subj varchar(255);
DECLARE tags_arr varchar(255)[];
DECLARE count_ex integer;
DECLARE tg_id integer;
DECLARE count_deleted integer := 0;
BEGIN
    SELECT tags, subject FROM exercise WHERE id = ex_id into tags_arr, subj;
    DELETE FROM tag_exercise WHERE exercise_id = ex_id;
    FOR i IN 1..array_length(tags_arr, 1) LOOP 
        -- RAISE NOTICE '%', i;
        SELECT id FROM tag WHERE name = tags_arr[i] INTO tg_id;
        SELECT COUNT(*) FROM tag_exercise WHERE tag_id = tg_id into count_ex;
        IF count_ex = 0 THEN
            DELETE FROM tag WHERE id = tg_id;
            count_deleted := count_deleted + 1;
        END IF;
    END LOOP;
    DELETE FROM exercise WHERE id = ex_id;
    RETURN count_deleted;
END;
$$ LANGUAGE plpgsql;


-- if tags exist empty

CREATE OR REPLACE FUNCTION update_exercise_tag(ex_id integer, tags_input varchar(255)[])
RETURNS INTEGER AS $$a
DECLARE size_existing_tags_arr integer;
DECLARE size_input_tags_arr integer;
DECLARE count_tag integer; 
DECLARE tag_exist boolean;
DECLARE tags_arr varchar(255)[];
DECLARE moved_tags varchar(255)[];
DECLARE new_tags varchar(255)[];
BEGIN 
    -- TODO for for search that not exist 
    -- new tags - is new arr todo, create arr thhat contains tags that differs from 
    -- input then and that new from input 
    -- then delete differs, new add to tag
    SELECT tags FROM exercise WHERE id = ex_id INTO tags_arr;

    FOR i in 1..array_length(tags_arr, 1) LOOP
        tag_exist = FALSE;
        FOR j in 1..array_length(tags_input, 1) LOOP 
            IF tags_arr[i] = tags_input[j] THEN
                tag_exist = TRUE;
                SELECT array_remove(new_tags, tags_arr[i]);
                EXIT;
            END IF;
        END LOOP;
        -- RAISE NOTICE '%' new_tags;
        IF TAG_EXIST = FALSE THEN 
            SELECT array_append(moved_tags, tags_arr[i]);
            -- SELECT COUNT (*) FROM tag_exercise WHERE exercise_id = ex_id into count_tag;
            -- IF count_tag == 0 THEN
        END IF;
    END LOOP;

    RAISE NOTICE '%', moved_tags;
    RAISE NOTICE '%', new_tags;

    FOR 
    -- FOR i in 1..array_length(moved_tags, 1) LOOP
    --     SELECT id FROM tags where subject = subj and name = moved_tags[i] into tag_id;

    --     DELETE FROM 
    --     IF (SELECT COUNT(*) FROM tag_exercise where exercise_id == ex_id) = O THEN 
    --         DELETE FROM tags 
    -- END 
    -- SE
    RETURN 0;
END;
$$LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_exercise_tags(ex_id integer, tags_to_add varchar(255)[], tags_to_remove varchar(255)[])
RETURNS INTEGER AS $$

BEGIN 

    RETURN 0;
END;


1 5 2 32 3 
2 3 23

1 5 32 - to delete 
23 to add 