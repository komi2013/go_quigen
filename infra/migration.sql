INSERT INTO t_question
(
question_id, question_txt, updated_at, usr_id, usr_img
)
select id, txt, open_time, 3, '/data/usr/20200715/3.png'
from question

UPDATE
  t_question AS t1
SET
  choice_0 = t2.choice_0, choice_1 = t2.choice_1, choice_2 = t2.choice_2, choice_3 = t2.choice_3
FROM choice AS t2
WHERE 
  t1.question_id = t2.question_id

UPDATE
  t_question AS t1
SET
  reference = t2.reference
FROM choice AS t2
WHERE t1.question_id = t2.question_id
AND t2.reference <> ''

UPDATE
  t_question AS t1
SET
  explanation = t2.txt
FROM comment AS t2
WHERE t1.question_id = t2.question_id


INSERT INTO m_category_name 
(
  category_name, updated_at
)
SELECT txt, '2020-09-05 08:08:00' FROM tag 
GROUP BY txt
ORDER BY txt

UPDATE
  tag AS t1
SET
  category_id = t2.category_id
FROM m_category_name AS t2
WHERE 
  t1.txt = t2.category_name


UPDATE
  t_question AS t1
SET
  category_id = t2.category_id
FROM tag AS t2
WHERE 
  t1.question_id = t2.question_id

UPDATE
  t_question AS t1
SET
  previous_question = t2.diff
FROM (
  SELECT
    question_id,
    lag(question_id,1,0) OVER(PARTITION BY category_id ORDER BY updated_at ) AS diff
  FROM
    t_question
) AS t2
WHERE 
  t1.question_id = t2.question_id

UPDATE
  t_question AS t1
SET
  sequence = t2.seq
FROM (
  SELECT
    question_id,
    ROW_NUMBER() OVER (PARTITION BY category_id ORDER BY updated_at ASC) AS seq
  FROM
    t_question
) AS t2
WHERE 
  t1.question_id = t2.question_id

INSERT INTO m_category_tree (leaf_id, level_1, updated_at)
SELECT category_id, category_id, '2020-09-05 09:30:00'
FROM m_category_name

UPDATE
  m_category_name AS t1
SET
  category_description = t2.description
FROM mt_seo_tag AS t2
WHERE 
  t1.category_name = t2.tag

INSERT INTO m_category_question (
  question_id, category_id, updated_at, SUBSTR(question_title, 0 , 60), in_list)
SELECT question_id, category_id, '2020-09-05 09:14:00', question_txt, 1
FROM t_question
WHERE sequence = 1
order by category_id , sequence