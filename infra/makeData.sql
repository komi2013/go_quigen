GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO exam_8099;
GRANT USAGE ON SCHEMA public TO exam_8099;

INSERT INTO t_question
(
question_id, category_id, question_txt, choice_0, sound, 
usr_id, usr_img, question_type, previous_question, next_question, sequence
)
select question_id, category_id, resource_txt, resource_img, resource_sound, 
3, '/data/usr/20200715/3.png', 2, previous_question, next_question, seq
from c_resource where question_id > 561
order by question_id ASC

INSERT INTO m_category_question
(
  question_id, category_id, question_title, in_list
)
select question_id, 1, substr(question_txt, 0, 30), 1
from t_question 
where sequence = 1
AND category_id in (2,3,4,5,15,18,19) 
order by question_id ASC


update t_question set choice_1 = c_resource.resource_img
from c_resource
where t_question.question_id = c_resource.question_id -1
and t_question.question_id > 561;

INSERT INTO m_category_tree(leaf_id,level_1,level_2,level_3,level_4)
select leaf_id+25,level_1+25,level_2+25,level_3+25,level_4+25 from m_category_tree 
where leaf_id > 25
order by leaf_id ASC;

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
  WHERE question_id > 561
) AS t2
WHERE 
  t1.question_id = t2.question_id

UPDATE t_question as t1
SET explanation = t2.explanation
FROM (
  SELECT
    explanation
  FROM
    c_resource
  -- WHERE category_id > 8
) AS t2
WHERE t1.question_id = t2.question_id

INSERT INTO t_question
(
question_id, category_id, question_txt, choice_0, choice_1, choice_2, choice_3, 
usr_id, usr_img, question_type, choice_type, explanation
)
SELECT question_id, category_id, resource_txt, choice_0, choice_1, choice_2, choice_3,
3, '/data/usr/20200715/3.png', 3, 0, explanation
FROM c_resource
ORDER BY question_id ASC

update t_question set question_txt = replace(question_txt,'/img/','/'), choice_0 = replace(choice_0,'/img/','/'), choice_1 = replace(choice_1,'/img/','/'),choice_2 = replace(choice_2,'/img/','/'),choice_3 = replace(choice_3,'/img/','/'),explanation = replace(explanation,'/img/','/');
