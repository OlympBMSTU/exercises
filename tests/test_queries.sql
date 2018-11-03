select * from (SELECT * FROM Exceriese WHERE subject='physics') ex
  join ((SELECT * FROM tag WHERE name='array' AND subject='physics') t join tag_excerciese tg on (tg.tag_id = t.id)) tt on (tt.excerciese_id = ex.id) ORDER BY ex.level 


 SELECT * FROM (SELECT * FROM tag WHERE name='array' AND subject='physics') t join tag_excerciese tg on (tg.tag_id = t.id)) tt on (tt.excerciese_id = ex.id) ORDER BY ex.level 


  SELECT * FROM (SELECT * FROM tag WHERE name='array' AND subject='physics') foo


   SELECT * FROM ((SELECT * FROM Excerciese WHERE subject='physics') ex join 
    ((SELECT * FROM tag WHERE name='array' AND subject='physics') t join tag_excerciese tg on (tg.tag_id = t.id)) tgt on tgt.excerciese_id = ex.id) f



SELECT ex.* FROM (SELECT * FROM Excerciese WHERE subject='physics') ex join 
( (SELECT id as t_id FROM tag WHERE name='array' AND subject='physics') t join tag_excerciese tg on (tg.tag_id = t.t_id)) tgt on tgt.excerciese_id = ex.id