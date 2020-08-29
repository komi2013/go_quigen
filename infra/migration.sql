INSERT INTO t_question
(
question_id, question_txt, updated_at, usr_id, usr_img
)
select id, 0, txt, open_time, 3, '/data/usr/20200715/3.png'
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
  reference = t2.txt
FROM comment AS t2
WHERE 
  t1.question_id = t2.question_id

INSERT INTO m_category_name 
(
  question_txt, updated_at
)
SELECT txt, '2020-08-29 13:14:00' FROM tag 
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
  next_question = t2.diff
FROM (
  SELECT
    question_id,
    lag(question_id,1,0) OVER(PARTITION BY category_id ORDER BY question_id ) AS diff
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
    ROW_NUMBER() OVER (PARTITION BY category_id ORDER BY question_id ASC) AS seq
  FROM
    t_question
) AS t2
WHERE 
  t1.question_id = t2.question_id


INSERT INTO m_category_question (
  question_id, category_id, updated_at, question_title, in_list)
SELECT question_id, category_id, '2020-08-29 18:24:00', question_txt, 0
FROM t_question
WHERE sequence < 16 AND sequence > 5

INSERT INTO m_category_tree (leaf_id, level_1, updated_at)
SELECT category_id, category_id, '2020-08-29 18:24:00'
FROM m_category_question
