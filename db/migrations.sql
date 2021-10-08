--
-- PostgreSQL database dump
--

-- Dumped from database version 12.8 (Debian 12.8-1.pgdg110+1)
-- Dumped by pg_dump version 12.8 (Ubuntu 12.8-0ubuntu0.20.04.1)

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

--
-- Data for Name: migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.migrations (id, applied_at) FROM stdin;
20211007194628-make-sellers.sql	2021-10-08 00:38:21.729698+00
20211007211946-make-products.sql	2021-10-08 01:22:36.385334+00
20211007220455-make-variations.sql	2021-10-08 02:36:45.649341+00
\.


--
-- PostgreSQL database dump complete
--

