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
