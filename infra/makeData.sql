GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO exam_8099;
GRANT USAGE ON SCHEMA public TO exam_8099;

insert into t_question
(
question_id, category_id, question_txt, choice_0, sound, 
usr_id, usr_img, question_type, previous_question, next_question, sequence
)
select question_id, category_id, resource_txt, resource_img, resource_sound, 
3, '/data/usr/20200715/3.png', 2, previous_question, next_question, seq
from c_resource where question_id > 340
order by question_id ASC

insert into m_category_question
(
  question_id, category_id, question_title, in_list
)
select question_id, category_list_id, resource_txt, in_list
from c_resource where question_id > 340
order by question_id ASC

update t_question set choice_1 = c_resource.resource_img
from c_resource
where t_question.question_id = c_resource.question_id -1
and t_question.question_id > 340;

UPDATE
  c_resource AS t1
SET
  seq = t2.seq
FROM (
  SELECT
    question_id,
    ROW_NUMBER() OVER (PARTITION BY category_id ORDER BY question_id ASC) AS seq
  FROM
    c_resource
  WHERE category_id > 25
) AS t2
WHERE 
  t1.question_id = t2.question_id


UPDATE
  c_resource AS t1
SET
  next_question = t2.diff
FROM (
  SELECT
    question_id,
    lead(question_id,1,0) OVER( ORDER BY question_id ) AS diff
  FROM
    c_resource
  WHERE category_id = 16
) AS t2
WHERE 
  t1.question_id = t2.question_id

UPDATE t_question as t1
SET sequence = t2.seq, previous_question = t2.previous_question, next_question = t2.next_question
FROM (
  SELECT
    question_id,seq,previous_question,next_question
  FROM
    c_resource
  WHERE category_id > 8
) AS t2
WHERE t1.question_id = t2.question_id