-- Adminer 4.7.1 PostgreSQL dump

DROP TABLE IF EXISTS "c_resource";
CREATE TABLE "public"."c_resource" (
    "resource_txt" text DEFAULT '' NOT NULL,
    "resource_img" text DEFAULT '' NOT NULL,
    "resource_sound" text DEFAULT '' NOT NULL,
    "category_id" integer DEFAULT '0' NOT NULL,
    "question_id" integer DEFAULT '0' NOT NULL,
    "previous_question" integer DEFAULT '0' NOT NULL,
    "next_question" integer DEFAULT '0' NOT NULL,
    "seq" integer DEFAULT '0' NOT NULL,
    "category_list_id" integer DEFAULT '0' NOT NULL,
    "in_list" smallint DEFAULT '0' NOT NULL,
    CONSTRAINT "c_resource_question_id" PRIMARY KEY ("question_id")
) WITH (oids = false);


DROP TABLE IF EXISTS "h_answer";
CREATE TABLE "public"."h_answer" (
    "question_id" integer DEFAULT '0' NOT NULL,
    "question_txt" text DEFAULT '' NOT NULL,
    "category_id" integer DEFAULT '0' NOT NULL,
    "respondent_id" integer DEFAULT '0' NOT NULL,
    "respondent_img" text DEFAULT '' NOT NULL,
    "correct" smallint DEFAULT '0' NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL
) WITH (oids = false);

CREATE INDEX "h_answer_category_id" ON "public"."h_answer" USING btree ("category_id");

CREATE INDEX "h_answer_question_id" ON "public"."h_answer" USING btree ("question_id");

CREATE INDEX "h_answer_respondent_id" ON "public"."h_answer" USING btree ("respondent_id");


DROP TABLE IF EXISTS "m_category_name";
DROP SEQUENCE IF EXISTS category_id_seq;
CREATE SEQUENCE category_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."m_category_name" (
    "category_id" integer DEFAULT nextval('category_id_seq') NOT NULL,
    "category_name" text DEFAULT '' NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "category_description" text DEFAULT '' NOT NULL,
    CONSTRAINT "m_category_name_category_id" PRIMARY KEY ("category_id")
) WITH (oids = false);


DROP TABLE IF EXISTS "m_category_question";
CREATE TABLE "public"."m_category_question" (
    "question_id" integer DEFAULT '0' NOT NULL,
    "category_id" integer DEFAULT '0' NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "question_title" text DEFAULT '' NOT NULL,
    "in_list" smallint DEFAULT '0' NOT NULL,
    CONSTRAINT "m_category_question_question_id" PRIMARY KEY ("question_id")
) WITH (oids = false);

CREATE INDEX "m_category_question_category_id" ON "public"."m_category_question" USING btree ("category_id");

COMMENT ON COLUMN "public"."m_category_question"."in_list" IS '0=direct show, 1=in list';


DROP TABLE IF EXISTS "m_category_tree";
CREATE TABLE "public"."m_category_tree" (
    "leaf_id" integer DEFAULT '0' NOT NULL,
    "level_1" integer DEFAULT '0' NOT NULL,
    "level_2" integer DEFAULT '0' NOT NULL,
    "level_3" integer DEFAULT '0' NOT NULL,
    "level_4" integer DEFAULT '0' NOT NULL,
    "level_5" integer DEFAULT '0' NOT NULL,
    "level_6" integer DEFAULT '0' NOT NULL,
    "level_7" integer DEFAULT '0' NOT NULL,
    "level_8" integer DEFAULT '0' NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    CONSTRAINT "m_category_tree_leaf_id" PRIMARY KEY ("leaf_id")
) WITH (oids = false);


DROP TABLE IF EXISTS "m_public_message";
DROP SEQUENCE IF EXISTS m_public_message_public_message_id_seq;
CREATE SEQUENCE m_public_message_public_message_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."m_public_message" (
    "public_message_id" integer DEFAULT nextval('m_public_message_public_message_id_seq') NOT NULL,
    "public_content" text DEFAULT '' NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "sender" integer DEFAULT '0' NOT NULL,
    "sender_img" text DEFAULT '' NOT NULL,
    "condition_id" smallint DEFAULT '0' NOT NULL,
    "start_time" timestamp DEFAULT now() NOT NULL,
    "end_time" timestamp DEFAULT now() NOT NULL,
    "img_txt_flg" smallint DEFAULT '0' NOT NULL
) WITH (oids = false);


DROP TABLE IF EXISTS "t_comment";
DROP SEQUENCE IF EXISTS t_comment_comment_id_seq;
CREATE SEQUENCE t_comment_comment_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."t_comment" (
    "comment_id" integer DEFAULT nextval('t_comment_comment_id_seq') NOT NULL,
    "comment_txt" text DEFAULT '' NOT NULL,
    "usr_id" integer DEFAULT '0' NOT NULL,
    "question_id" integer DEFAULT '0' NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "usr_img" text DEFAULT '' NOT NULL,
    CONSTRAINT "t_comment_comment_id" PRIMARY KEY ("comment_id")
) WITH (oids = false);

CREATE INDEX "t_comment_question_id" ON "public"."t_comment" USING btree ("question_id");


DROP TABLE IF EXISTS "t_message";
DROP SEQUENCE IF EXISTS t_message_message_id_seq;
CREATE SEQUENCE t_message_message_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."t_message" (
    "message_id" integer DEFAULT nextval('t_message_message_id_seq') NOT NULL,
    "sender" integer DEFAULT '0' NOT NULL,
    "receiver" integer DEFAULT '0' NOT NULL,
    "message_content" text DEFAULT '' NOT NULL,
    "created_at" timestamp DEFAULT now() NOT NULL,
    "sender_img" text DEFAULT '' NOT NULL,
    "img_txt_flg" smallint DEFAULT '0' NOT NULL,
    CONSTRAINT "t_message_message_id" PRIMARY KEY ("message_id")
) WITH (oids = false);

CREATE INDEX "t_message_sender_receiver" ON "public"."t_message" USING btree ("sender", "receiver");

COMMENT ON COLUMN "public"."t_message"."img_txt_flg" IS '0=txt,1=img';


DROP TABLE IF EXISTS "t_question";
DROP SEQUENCE IF EXISTS t_question_question_id_seq;
CREATE SEQUENCE t_question_question_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 1 CACHE 1;

CREATE TABLE "public"."t_question" (
    "question_id" integer DEFAULT nextval('t_question_question_id_seq') NOT NULL,
    "question_txt" text DEFAULT '' NOT NULL,
    "usr_id" integer DEFAULT '0' NOT NULL,
    "usr_img" text DEFAULT '' NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "choice_0" text DEFAULT '' NOT NULL,
    "choice_1" text DEFAULT '' NOT NULL,
    "choice_2" text DEFAULT '' NOT NULL,
    "choice_3" text DEFAULT '' NOT NULL,
    "reference" text DEFAULT '' NOT NULL,
    "question_type" smallint DEFAULT '0' NOT NULL,
    "category_id" integer DEFAULT '0' NOT NULL,
    "question_img" text DEFAULT '' NOT NULL,
    "previous_question" integer DEFAULT '0' NOT NULL,
    "next_question" integer DEFAULT '0' NOT NULL,
    "sequence" integer DEFAULT '0' NOT NULL,
    "sound" text DEFAULT '' NOT NULL,
    "explanation" text DEFAULT '' NOT NULL,
    CONSTRAINT "t_question_question_id" PRIMARY KEY ("question_id")
) WITH (oids = false);

COMMENT ON COLUMN "public"."t_question"."question_type" IS '0=choice, 1=text, 2=picture';

COMMENT ON COLUMN "public"."t_question"."sequence" IS 'within leaf category';


DROP TABLE IF EXISTS "t_usr";
DROP SEQUENCE IF EXISTS usr_id_seq;
CREATE SEQUENCE usr_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;

CREATE TABLE "public"."t_usr" (
    "usr_id" integer DEFAULT nextval('usr_id_seq') NOT NULL,
    "pv_u_id" text DEFAULT '' NOT NULL,
    "provider" integer DEFAULT '0' NOT NULL,
    "usr_img" text DEFAULT '' NOT NULL,
    "updated_at" timestamp DEFAULT now() NOT NULL,
    "introduce" text DEFAULT '' NOT NULL,
    "push_tokens" json,
    CONSTRAINT "t_usr_pkey" PRIMARY KEY ("usr_id")
) WITH (oids = false);

COMMENT ON COLUMN "public"."t_usr"."provider" IS '1=FB, 2=TW, 3=G+, 4=LN';


-- 2020-08-29 05:03:02.078806+03