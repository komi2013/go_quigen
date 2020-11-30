GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO exam_8099;
GRANT USAGE ON SCHEMA public TO exam_8099;

INSERT INTO t_question
(
category_id, question_txt, choice_0, usr_id, usr_img
)
select 50, question, choice_0, 3, '/data/usr/20200715/3.png'
from c_word where word_id > 499 AND word_id < 1251
order by word_id ASC

UPDATE
  t_question AS t1
SET
  choice_1 = t2.choice_1, choice_2 = t2.choice_2, choice_3 = t2.choice_3
FROM (
  SELECT
    question_id,
    choice_0,
    lag(choice_0,1,'') OVER(PARTITION BY category_id ORDER BY question_id ) AS choice_1,
    lag(choice_0,2,'') OVER(PARTITION BY category_id ORDER BY question_id ) AS choice_2,
    lead(choice_0,1,'') OVER(PARTITION BY category_id ORDER BY question_id ) AS choice_3
  FROM t_question
  WHERE category_id = 50
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
usr_id, usr_img, choice_type, explanation, sequence, previous_question, next_question
)
SELECT question_id, category_id, resource_txt, choice_0, choice_1, choice_2, choice_3,
3, '/data/usr/20200715/3.png', 0, explanation, seq, previous_question, next_question
FROM c_resource
ORDER BY question_id ASC

update t_question set question_txt = replace(question_txt,'/img/','/'), choice_0 = replace(choice_0,'/img/','/'), choice_1 = replace(choice_1,'/img/','/'),choice_2 = replace(choice_2,'/img/','/'),choice_3 = replace(choice_3,'/img/','/'),explanation = replace(explanation,'/img/','/');

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
AND t2.txt <> ''


INSERT INTO m_category_name 
(
  category_name, updated_at
)
SELECT txt, '2020-10-05 08:08:00' FROM tag 
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
  previous_question = t2.prev, next_question = t2.next, sequence = t2.seq
FROM (
  SELECT
    question_id,
    lag(question_id,1,0) OVER(PARTITION BY category_id ORDER BY question_id ) AS prev,
    lead(question_id,1,0) OVER(PARTITION BY category_id ORDER BY question_id ) AS next,
    ROW_NUMBER() OVER (PARTITION BY category_id ORDER BY question_id ASC) AS seq
  FROM
    t_question
) AS t2
WHERE t1.question_id = t2.question_id
AND category_id in (95,96,97,98,99,100,101)

INSERT INTO m_category_tree (leaf_id, level_1, updated_at)
SELECT category_id, category_id, '2020-10-10 09:30:00'
FROM m_category_name

UPDATE
  m_category_name AS t1
SET
  category_description = t2.description
FROM mt_seo_tag AS t2
WHERE 
  t1.category_name = t2.tag

INSERT INTO m_category_question (
  question_id, category_id, updated_at, question_title, in_list)
SELECT question_id, category_id, '2020-11-26 09:34:00', SUBSTR(question_txt, 0 , 60), 0
FROM t_question
WHERE question_id in (
 ,851	      	
 ,5077	           
 ,684		      
 ,4858	              
 ,2549	              
 ,3717	         
 ,4257	             
 ,2120	            	 
 ,1624	            	 
 ,1172
)

update s_stock as t1 set ratio_0_1 = t3.ratio_0_1, ratio_1_3 = t3.ratio3, ratio_1_9 = t3.ratio9, ratio_1_19 = t3.ratio19
from (
select * from (
SELECT
stock_id,
date,
price,
ratio_0_1,
lead(price,3,price) OVER(PARTITION BY stock_id ORDER BY date DESC ) as last3,
round((lead(price,1,price) OVER(PARTITION BY stock_id ORDER BY date DESC ) / lead(price,3,price) OVER(PARTITION BY stock_id ORDER BY date DESC ) * 100) - 100,2) as ratio3,
lead(price,9,price) OVER(PARTITION BY stock_id ORDER BY date DESC ) as last9,
round((lead(price,1,price) OVER(PARTITION BY stock_id ORDER BY date DESC ) / lead(price,9,price) OVER(PARTITION BY stock_id ORDER BY date DESC ) * 100) - 100,2) as ratio9,
lead(price,19,price) OVER(PARTITION BY stock_id ORDER BY date DESC ) as last19,
round((lead(price,1,price) OVER(PARTITION BY stock_id ORDER BY date DESC ) / lead(price,19,price) OVER(PARTITION BY stock_id ORDER BY date DESC ) * 100) - 100,2) as ratio19
FROM
    s_chart
) as t2
where t2.date = '2020-11-30'
) t3
where t1.stock_id = t3.stock_id



	             
