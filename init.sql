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
    is_broken boolean default false,
    class integer,
    position INTEGER,
    mark   INTEGER,
    type_olymp INTEGER,
    answer jsonb,
    created timestamp default NOW()
);

-- [{id:1, input: "dfd", output: "fd"}, {id: 2, input: "ds", output: "df"}]

-- DELETE FROM TAGS WHERE id int (SELECT tag_id FROM tax_excercise te JOIN (SELECT id FROM tag WHERE tag_nmae = '' AND subject = '') WHERE C 

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
CREATE OR REPLACE FUNCTION add_exercise(auth_id integer, lev integer, f_name varchar(255),
     subj varchar(255), tags varchar(255)[], cls integer, pos integer, mrk integer, typ_ol integer,
    answ jsonb)
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

    INSERT INTO EXERCISE(author_id, file_name, level, subject, tags, class, position, mark, type_olymp, answer) VALUES(auth_id, f_name, lev, subj, tags, cls, mrk, pos, typ_ol, answ) RETURNING id INTO ex_id;
    FOR i IN 1..array_length(tags, 1) LOOP
        SELECT id from tag where subject = subj and name = tags[i] into t_id;
        if t_id IS null then 
            INSERT INTO TAG(subject, name) VALUES(subj, tags[i]) RETURNING id INTO t_id;
        end if;

        -- whats haooens here if tag is duplicated ???
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


-- CREATE OR REPLACE PROCEDURE fix_tags() 
-- DECLARE subj varchar(255);
-- DECLARE to_add, to_remove varchar(255)[];
-- BEGIN 
--     if new_subj is null then
--         select array(select unnest(old) except select unnest(new)) into to_remove;
--         select array(select unnest(new) except select unnest(old])) into to_add;
--         subj = ex.subj;
--     else 
--         to_remove = old; -- with subj
--         select array(select unnest(new) except select unnest(old])) into to_add;
--         subj = new_subj;
--     end if;



-- if tags exist empty

-- CREATE OR REPLACE FUNCTION update_exercise_tag(ex_id integer, tags_input varchar(255)[])
-- RETURNS INTEGER AS $$a
-- DECLARE size_existing_tags_arr integer;
-- DECLARE size_input_tags_arr integer;
-- DECLARE count_tag integer; 
-- DECLARE tag_exist boolean;
-- DECLARE tags_arr varchar(255)[];
-- DECLARE moved_tags varchar(255)[];
-- DECLARE new_tags varchar(255)[];
-- BEGIN 
--     -- TODO for for search that not exist 
--     -- new tags - is new arr todo, create arr thhat contains tags that differs from 
--     -- input then and that new from input 
--     -- then delete differs, new add to tag
--     SELECT tags FROM exercise WHERE id = ex_id INTO tags_arr;

--     FOR i in 1..array_length(tags_arr, 1) LOOP
--         tag_exist = FALSE;
--         FOR j in 1..array_length(tags_input, 1) LOOP 
--             IF tags_arr[i] = tags_input[j] THEN
--                 tag_exist = TRUE;
--                 SELECT array_remove(new_tags, tags_arr[i]);
--                 EXIT;
--             END IF;
--         END LOOP;
--         -- RAISE NOTICE '%' new_tags;
--         IF TAG_EXIST = FALSE THEN 
--             SELECT array_append(moved_tags, tags_arr[i]);
--             -- SELECT COUNT (*) FROM tag_exercise WHERE exercise_id = ex_id into count_tag;
--             -- IF count_tag == 0 THEN
--         END IF;
--     END LOOP;

--     RAISE NOTICE '%', moved_tags;
--     RAISE NOTICE '%', new_tags;

--     FOR 
--     -- FOR i in 1..array_length(moved_tags, 1) LOOP
--     --     SELECT id FROM tags where subject = subj and name = moved_tags[i] into tag_id;

--     --     DELETE FROM 
--     --     IF (SELECT COUNT(*) FROM tag_exercise where exercise_id == ex_id) = O THEN 
--     --         DELETE FROM tags 
--     -- END 
--     -- SE
--     RETURN 0;
-- END;
-- $$LANGUAGE plpgsql;

-- TRIGGER FOR UPDATE 
-- delete old tags - also delete counter
-- delete tags with 0
-- insert new tags
-- also if new entity exist new subject 

CREATE OR REPLACE FUNCTION update_exercise_tags(ex_id integer, old_subj varchar(255), new_subj varchar(255), tags_to_add varchar(255)[], tags_to_remove varchar(255)[])
RETURNS INTEGER AS $$
DECLARE subj varchar(255);
DECLARE tg_id integer;
DECLARE tag_ex_id integer;
BEGIN 
    -- SELECT subject from exercise WHERE id = ex_id INTO subj;
    if new_subj IS NULL OR new_subj = '' THEN 
        new_subj = old_subj;
    END IF;
    -- RAISE NOTICE '% %', old_subj, new_subj;
    -- RAISE NOTICE '%', subj;
    -- RAISE NOTICE '% %', tags_to_add, tags_to_remove;
    IF array_length(tags_to_remove, 1) IS NOT NULL AND array_length(tags_to_remove, 1) > 0 THEN
        FOR i in 1..array_length(tags_to_remove, 1) LOOP
            -- RAISE NOTICE '% %', i, tags_to_remove[i];

            SELECT id FROM tag WHERE subject = old_subj AND name = tags_to_remove[i] INTO tg_id;
            -- RAISE NOTICE '%', tg_id;
            DELETE FROM tag_exercise WHERE tag_id = tg_id and exercise_id = ex_id;
            IF (SELECT COUNT(*) FROM tag_exercise WHERE tag_id = tg_id) = 0 THEN
                DELETE FROM tag WHERE id = tg_id;
            END IF;
        END LOOP;
    END IF;

    IF array_length(tags_to_add, 1) IS NOT NULL AND array_length(tags_to_add, 1) > 0 THEN
        FOR i in 1..array_length(tags_to_add, 1) LOOP 
            SELECT id from tag where subject = new_subj and name = tags_to_add[i] into tg_id;
            if tg_id IS null then 
                INSERT INTO TAG(subject, name) VALUES(new_subj, tags_to_add[i]) RETURNING id INTO tg_id;
            end if;

            --- or check
            
            SELECT id FROM tag_exercise where exercise_id = ex_id AND tag_id = tg_id into tag_ex_id;
            
            -- we've checked this at back, but new tag subject
            IF tag_ex_id IS NULL THEN 
                INSERT INTO TAG_EXERCISE(tag_id, exercise_id) VALUES(tg_id, ex_id);         
            END IF;
        END LOOP;
    END IF;

    RETURN 0;
END;
$$LANGUAGE plpgsql;


-- insert into subject(name) values('mathematic'), ('physics');

-- del fr tag_ex where ex_id = and id in (select id from  not in join)) 

-- 1 5 2 32 3 
-- 2 3 23

-- 1 5 32 - to delete 
-- 23 to add 