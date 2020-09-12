--
-- PostgreSQL database dump
--

-- Dumped from database version 11.9
-- Dumped by pg_dump version 11.9

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: c_resource; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.c_resource (
    resource_txt text DEFAULT ''::text NOT NULL,
    resource_img text DEFAULT ''::text NOT NULL,
    resource_sound text DEFAULT ''::text NOT NULL,
    category_id integer DEFAULT 0 NOT NULL,
    question_id integer DEFAULT 0 NOT NULL,
    previous_question integer DEFAULT 0 NOT NULL,
    next_question integer DEFAULT 0 NOT NULL,
    seq integer DEFAULT 0 NOT NULL,
    category_list_id integer DEFAULT 0 NOT NULL,
    in_list smallint DEFAULT '0'::smallint NOT NULL
);


ALTER TABLE public.c_resource OWNER TO postgres;

--
-- Name: m_category_tree; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_category_tree (
    leaf_id integer DEFAULT 0 NOT NULL,
    level_1 integer DEFAULT 0 NOT NULL,
    level_2 integer DEFAULT 0 NOT NULL,
    level_3 integer DEFAULT 0 NOT NULL,
    level_4 integer DEFAULT 0 NOT NULL,
    level_5 integer DEFAULT 0 NOT NULL,
    level_6 integer DEFAULT 0 NOT NULL,
    level_7 integer DEFAULT 0 NOT NULL,
    level_8 integer DEFAULT 0 NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.m_category_tree OWNER TO postgres;

--
-- Name: category_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.category_id_seq OWNER TO postgres;

--
-- Name: category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.category_id_seq OWNED BY public.m_category_tree.leaf_id;


--
-- Name: h_answer; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.h_answer (
    question_id integer DEFAULT 0 NOT NULL,
    question_txt text DEFAULT ''::text NOT NULL,
    category_id integer DEFAULT 0 NOT NULL,
    respondent_id integer DEFAULT 0 NOT NULL,
    respondent_img text DEFAULT ''::text NOT NULL,
    correct smallint DEFAULT '0'::smallint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.h_answer OWNER TO postgres;

--
-- Name: m_category_name; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_category_name (
    category_id integer DEFAULT nextval('public.category_id_seq'::regclass) NOT NULL,
    category_name text DEFAULT ''::text NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    category_description text DEFAULT ''::text NOT NULL
);


ALTER TABLE public.m_category_name OWNER TO postgres;

--
-- Name: m_category_question; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_category_question (
    question_id integer DEFAULT 0 NOT NULL,
    category_id integer DEFAULT 0 NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    question_title text DEFAULT ''::text NOT NULL,
    in_list smallint DEFAULT '0'::smallint NOT NULL
);


ALTER TABLE public.m_category_question OWNER TO postgres;

--
-- Name: COLUMN m_category_question.in_list; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.m_category_question.in_list IS '0=direct show, 1=in list';


--
-- Name: m_public_message; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.m_public_message (
    public_message_id integer NOT NULL,
    public_content text DEFAULT ''::text NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    sender integer DEFAULT 0 NOT NULL,
    sender_img text DEFAULT ''::text NOT NULL,
    condition_id smallint DEFAULT '0'::smallint NOT NULL,
    start_time timestamp without time zone DEFAULT now() NOT NULL,
    end_time timestamp without time zone DEFAULT now() NOT NULL,
    img_txt_flg smallint DEFAULT '0'::smallint NOT NULL
);


ALTER TABLE public.m_public_message OWNER TO postgres;

--
-- Name: m_public_message_public_message_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.m_public_message_public_message_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.m_public_message_public_message_id_seq OWNER TO postgres;

--
-- Name: m_public_message_public_message_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.m_public_message_public_message_id_seq OWNED BY public.m_public_message.public_message_id;


--
-- Name: nologin_usr_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.nologin_usr_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
    CYCLE;


ALTER TABLE public.nologin_usr_id_seq OWNER TO postgres;

--
-- Name: t_comment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_comment (
    comment_id integer NOT NULL,
    comment_txt text DEFAULT ''::text NOT NULL,
    usr_id integer DEFAULT 0 NOT NULL,
    question_id integer DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    usr_img text DEFAULT ''::text NOT NULL
);


ALTER TABLE public.t_comment OWNER TO postgres;

--
-- Name: t_comment_comment_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_comment_comment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_comment_comment_id_seq OWNER TO postgres;

--
-- Name: t_comment_comment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_comment_comment_id_seq OWNED BY public.t_comment.comment_id;


--
-- Name: t_message; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_message (
    message_id integer NOT NULL,
    sender integer DEFAULT 0 NOT NULL,
    receiver integer DEFAULT 0 NOT NULL,
    message_content text DEFAULT ''::text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    sender_img text DEFAULT ''::text NOT NULL,
    img_txt_flg smallint DEFAULT '0'::smallint NOT NULL
);


ALTER TABLE public.t_message OWNER TO postgres;

--
-- Name: COLUMN t_message.img_txt_flg; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.t_message.img_txt_flg IS '0=txt,1=img';


--
-- Name: t_message_message_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_message_message_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_message_message_id_seq OWNER TO postgres;

--
-- Name: t_message_message_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_message_message_id_seq OWNED BY public.t_message.message_id;


--
-- Name: t_question; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_question (
    question_id integer NOT NULL,
    question_txt text DEFAULT ''::text NOT NULL,
    usr_id integer DEFAULT 0 NOT NULL,
    usr_img text DEFAULT ''::text NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    choice_0 text DEFAULT ''::text NOT NULL,
    choice_1 text DEFAULT ''::text NOT NULL,
    choice_2 text DEFAULT ''::text NOT NULL,
    choice_3 text DEFAULT ''::text NOT NULL,
    reference text DEFAULT ''::text NOT NULL,
    question_type smallint DEFAULT '0'::smallint NOT NULL,
    category_id integer DEFAULT 0 NOT NULL,
    question_img text DEFAULT ''::text NOT NULL,
    previous_question integer DEFAULT 0 NOT NULL,
    next_question integer DEFAULT 0 NOT NULL,
    sequence integer DEFAULT 0 NOT NULL,
    sound text DEFAULT ''::text NOT NULL,
    explanation text DEFAULT ''::text NOT NULL
);


ALTER TABLE public.t_question OWNER TO postgres;

--
-- Name: COLUMN t_question.question_type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.t_question.question_type IS '0=choice, 1=text, 2=picture';


--
-- Name: COLUMN t_question.sequence; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.t_question.sequence IS 'within leaf category';


--
-- Name: t_question_question_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.t_question_question_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_question_question_id_seq OWNER TO postgres;

--
-- Name: t_question_question_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.t_question_question_id_seq OWNED BY public.t_question.question_id;


--
-- Name: usr_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.usr_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.usr_id_seq OWNER TO postgres;

--
-- Name: t_usr; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_usr (
    usr_id integer DEFAULT nextval('public.usr_id_seq'::regclass) NOT NULL,
    pv_u_id text DEFAULT ''::text NOT NULL,
    provider integer DEFAULT 0 NOT NULL,
    usr_img text DEFAULT ''::text NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    introduce text DEFAULT ''::text NOT NULL,
    push_tokens json
);


ALTER TABLE public.t_usr OWNER TO postgres;

--
-- Name: COLUMN t_usr.provider; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.t_usr.provider IS '1=FB, 2=TW, 3=G+, 4=LN';


--
-- Name: m_public_message public_message_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_public_message ALTER COLUMN public_message_id SET DEFAULT nextval('public.m_public_message_public_message_id_seq'::regclass);


--
-- Name: t_comment comment_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_comment ALTER COLUMN comment_id SET DEFAULT nextval('public.t_comment_comment_id_seq'::regclass);


--
-- Name: t_message message_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_message ALTER COLUMN message_id SET DEFAULT nextval('public.t_message_message_id_seq'::regclass);


--
-- Name: t_question question_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_question ALTER COLUMN question_id SET DEFAULT nextval('public.t_question_question_id_seq'::regclass);


--
-- Name: c_resource c_resource_question_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.c_resource
    ADD CONSTRAINT c_resource_question_id PRIMARY KEY (question_id);


--
-- Name: m_category_name m_category_name_category_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_category_name
    ADD CONSTRAINT m_category_name_category_id PRIMARY KEY (category_id);


--
-- Name: m_category_question m_category_question_question_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_category_question
    ADD CONSTRAINT m_category_question_question_id PRIMARY KEY (question_id);


--
-- Name: m_category_tree m_category_tree_leaf_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_category_tree
    ADD CONSTRAINT m_category_tree_leaf_id PRIMARY KEY (leaf_id);


--
-- Name: m_public_message m_public_message_public_message_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.m_public_message
    ADD CONSTRAINT m_public_message_public_message_id PRIMARY KEY (public_message_id);


--
-- Name: t_comment t_comment_comment_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_comment
    ADD CONSTRAINT t_comment_comment_id PRIMARY KEY (comment_id);


--
-- Name: t_message t_message_message_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_message
    ADD CONSTRAINT t_message_message_id PRIMARY KEY (message_id);


--
-- Name: t_question t_question_question_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_question
    ADD CONSTRAINT t_question_question_id PRIMARY KEY (question_id);


--
-- Name: t_usr t_usr_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_usr
    ADD CONSTRAINT t_usr_pkey PRIMARY KEY (usr_id);


--
-- Name: h_answer_category_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX h_answer_category_id ON public.h_answer USING btree (category_id);


--
-- Name: h_answer_question_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX h_answer_question_id ON public.h_answer USING btree (question_id);


--
-- Name: h_answer_respondent_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX h_answer_respondent_id ON public.h_answer USING btree (respondent_id);


--
-- Name: m_category_question_category_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX m_category_question_category_id ON public.m_category_question USING btree (category_id);


--
-- Name: t_comment_question_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX t_comment_question_id ON public.t_comment USING btree (question_id);


--
-- Name: t_message_sender_receiver; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX t_message_sender_receiver ON public.t_message USING btree (sender, receiver);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

GRANT USAGE ON SCHEMA public TO exam_8099;


--
-- Name: TABLE c_resource; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.c_resource TO exam_8099;


--
-- Name: TABLE m_category_tree; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.m_category_tree TO exam_8099;


--
-- Name: SEQUENCE category_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.category_id_seq TO exam_8099;


--
-- Name: TABLE h_answer; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.h_answer TO exam_8099;


--
-- Name: TABLE m_category_name; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.m_category_name TO exam_8099;


--
-- Name: TABLE m_category_question; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.m_category_question TO exam_8099;


--
-- Name: TABLE m_public_message; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.m_public_message TO exam_8099;


--
-- Name: SEQUENCE m_public_message_public_message_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.m_public_message_public_message_id_seq TO exam_8099;


--
-- Name: SEQUENCE nologin_usr_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.nologin_usr_id_seq TO exam_8099;


--
-- Name: TABLE t_comment; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.t_comment TO exam_8099;


--
-- Name: SEQUENCE t_comment_comment_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.t_comment_comment_id_seq TO exam_8099;


--
-- Name: TABLE t_message; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.t_message TO exam_8099;


--
-- Name: SEQUENCE t_message_message_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.t_message_message_id_seq TO exam_8099;


--
-- Name: TABLE t_question; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.t_question TO exam_8099;


--
-- Name: SEQUENCE t_question_question_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.t_question_question_id_seq TO exam_8099;


--
-- Name: SEQUENCE usr_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.usr_id_seq TO exam_8099;


--
-- Name: TABLE t_usr; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.t_usr TO exam_8099;


--
-- PostgreSQL database dump complete
--

