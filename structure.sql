--
-- PostgreSQL database dump
--

-- Dumped from database version 10.1
-- Dumped by pg_dump version 10.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: pl_operator; Type: TABLE; Schema: public; Owner: kjell
--

CREATE TABLE pl_operator (
    transaction_id text NOT NULL,
    operator_id integer,
    operator_name text
);


ALTER TABLE pl_operator OWNER TO kjell;

--
-- Name: pl_transaction; Type: TABLE; Schema: public; Owner: kjell
--

CREATE TABLE pl_transaction (
    transaction_id text NOT NULL,
    retail_store_id smallint NOT NULL,
    workstation_id text NOT NULL,
    sequence_number text NOT NULL,
    end_date_time timestamp without time zone,
    operator_id integer NOT NULL,
    currency_code text,
    transaction_type_id text,
    buisness_day_date date NOT NULL,
    control_transaction text
);


ALTER TABLE pl_transaction OWNER TO kjell;

--
-- Name: pl_transaction_operator_id_seq; Type: SEQUENCE; Schema: public; Owner: kjell
--

CREATE SEQUENCE pl_transaction_operator_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE pl_transaction_operator_id_seq OWNER TO kjell;

--
-- Name: pl_transaction_operator_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: kjell
--

ALTER SEQUENCE pl_transaction_operator_id_seq OWNED BY pl_transaction.operator_id;


--
-- Name: pl_transaction operator_id; Type: DEFAULT; Schema: public; Owner: kjell
--

ALTER TABLE ONLY pl_transaction ALTER COLUMN operator_id SET DEFAULT nextval('pl_transaction_operator_id_seq'::regclass);


--
-- Name: pl_operator pl_operator_pkey; Type: CONSTRAINT; Schema: public; Owner: kjell
--

ALTER TABLE ONLY pl_operator
    ADD CONSTRAINT pl_operator_pkey PRIMARY KEY (transaction_id);


--
-- Name: pl_transaction pl_transaction_pkey; Type: CONSTRAINT; Schema: public; Owner: kjell
--

ALTER TABLE ONLY pl_transaction
    ADD CONSTRAINT pl_transaction_pkey PRIMARY KEY (transaction_id);


--
-- Name: pl_operator pl_operator_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: kjell
--

ALTER TABLE ONLY pl_operator
    ADD CONSTRAINT pl_operator_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES pl_transaction(transaction_id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

