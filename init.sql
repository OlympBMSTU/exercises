-- index - subject,
-- index subject, level
-- tag index - name, subject name
-- 

create table if not exists subject (
    id serial,
    name varchar(255) 
);

create table if not exists excerciese (
    id serial, 
    author_id integer,
    right_answer varchar(255),
    level integer,
    file_name varchar(255),
    subject varchar(255)
);

create table if not exists tag (
    id serial,
    subject varchar(255),
    name varchar(255),
    CONSTRAINT subjec_tag UNIQUE (subject, name)
);

create table if not exists tag_excerciese(
    id serial,
    excerciese_id integer,
    tag_id integer, 
    CONSTRAINT ex_id_tag_id UNIQUE (excerciese_id, tag_id)
);


-- index lower name
create table if not exists subject(
    id serial,
    name varchar(255) UNIQUE
);


-- lower
-- todo ON CONFLCIT делать и проверка на id tag 
CREATE OR REPLACE FUNCTION add_excerciese(auth_id integer, r_answer varchar(255), lev integer, f_name varchar(255), subj varchar(255), tags varchar(255)[])
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

    INSERT INTO EXCERCIESE(author_id, file_name, right_answer, level, subject) VALUES(auth_id, f_name, r_answer, lev, subj) RETURNING id INTO ex_id;
    FOR i IN 1..array_length(tags, 1) LOOP
        SELECT id from tag where subject = subj and name = tags[i] into t_id;
        if t_id IS null then 
            INSERT INTO TAG(subject, name) VALUES(subj, tags[i]) RETURNING id INTO t_id;
        end if;

        SELECT id FROM tag_excerciese where excerciese_id = ex_id AND tag_id = t_id into tag_ex_id;
        IF tag_ex_id IS NULL THEN 
            INSERT INTO TAG_EXCERCIESE(tag_id, excerciese_id) VALUES(t_id, ex_id);         
        END IF;
    END LOOP;

    RETURN 0;
END;
$$ LANGUAGE plpgsql;

-- SELECT add_excerciese(1, 'dg', 3, 'fdfd', 'mathematic', ARRAY['data', 'nj'])



-- по уму сначала уменьшаем таблицу просредника 
-- select * from (SELECT * FROM Exceriese WHERE subject=$1) ex  join ((SELECT * FROM tag WHERE name=$2) t join tag_excerciese tg on (tg.tag_id = t.id)) tt on (tt.excerciese_id = ex.id)

-- select * from tag t join tag_excerciese tg on (t.id = tg.tag_id) join excerciese_id ex on (ex.id = tg.excerciese_id) 
-- WHERE t.name = 'equations'

        -- raise notice '%', t_id;

        -- ON CONLICT 
        --     DO SELECT id from tag where subject = subj and name = tags[i] into t_id;

        -- INSERT INTO TAG(subject, name) VALUES(subj, tags[i]) RETURNING id INTO t_id


    -- SELECT pg_typeof(tags) INTO data;
    -- raise notice '%', data;
    -- FOREACH tag SLICE 1 in ARRAY tags LOOP
    --     INSERT INTO TAG(subject, name) VALUES(subject, tag) RETURNING id INTO t_id;
    --     INSERT INTO TAG_EXCERCIESE(tag_id, excerciese_id) VALUES(t_id, ex_id);
    -- END LOOP;