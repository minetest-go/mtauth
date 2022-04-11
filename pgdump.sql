--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1 (Debian 14.1-1.pgdg110+1)
-- Dumped by pg_dump version 14.1 (Debian 14.1-1.pgdg110+1)

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

SET default_table_access_method = heap;

--
-- Name: auth; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.auth (
    id integer NOT NULL,
    name text,
    password text,
    last_login integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.auth OWNER TO postgres;

--
-- Name: auth_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.auth_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_id_seq OWNER TO postgres;

--
-- Name: auth_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.auth_id_seq OWNED BY public.auth.id;


--
-- Name: user_privileges; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_privileges (
    id integer NOT NULL,
    privilege text NOT NULL
);


ALTER TABLE public.user_privileges OWNER TO postgres;

--
-- Name: auth id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth ALTER COLUMN id SET DEFAULT nextval('public.auth_id_seq'::regclass);


--
-- Data for Name: auth; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.auth (id, name, password, last_login) FROM stdin;
1	test	#1#gZ/lKmPydwJpCw7ItG51vw#KMGXLc/ioCCCaoCG/4OMdxGQRBkdJbItGcbifljwodZ5N0+0cuuKEPW+i0p8kuW56c6jCyh992gKC509peBBeqeSnw7wVsXDOsOzt/cG+vUj5zUpX1vkWoGM65HN6fi1CHtXYYNdQps6wTBSyF6klaSEB1gqzZDntIFbezvNPO2h0w2in7zt3PrWsxL9LdRAuh8LAIPXbFgrct93aYnPlKU6T+ObdXH6qoTRu+Sc4lUj2rAY4ZpGwhMc1GKNtrMzh5C6YUDMsR8Hr9VJmP8llKrLXvYzwjr0FTSxSTuhjGf7Ns7KFSq2KdR/0VZ7jTnL87ntbsy9Pj8qlFVMkd/jVQ	1649667042
\.


--
-- Data for Name: user_privileges; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_privileges (id, privilege) FROM stdin;
1	interact
1	shout
\.


--
-- Name: auth_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.auth_id_seq', 1, true);


--
-- Name: auth auth_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth
    ADD CONSTRAINT auth_name_key UNIQUE (name);


--
-- Name: auth auth_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth
    ADD CONSTRAINT auth_pkey PRIMARY KEY (id);


--
-- Name: user_privileges user_privileges_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_privileges
    ADD CONSTRAINT user_privileges_pkey PRIMARY KEY (id, privilege);


--
-- Name: user_privileges fk_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_privileges
    ADD CONSTRAINT fk_id FOREIGN KEY (id) REFERENCES public.auth(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

