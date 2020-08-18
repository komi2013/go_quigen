GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO exam_8099;
GRANT USAGE ON SCHEMA public TO exam_8099;

insert into t_question
(
question_id, category_id, question_txt, choice_0, sound, 
usr_id, usr_img, question_type
)
select question_id, category_id, resource_txt, resource_img, resource_sound, 
3, '/data/usr/20200715/3.png', 2
from c_resource where resource_id > 100
order by question_id
-- question_id is more than 115

update t_question set choice_1 = c_resource.resource_img
from c_resource
where t_question.question_id = c_resource.question_id
and c_resource.resource_id > 100;

