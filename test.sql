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
-- Name: poslog; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog (
    filename text NOT NULL,
    retail_store_id integer NOT NULL,
    buisness_day_date date NOT NULL,
    transaction_count integer NOT NULL,
    id bigint NOT NULL
);


ALTER TABLE poslog OWNER TO poslog;

--
-- Name: poslog_id_seq; Type: SEQUENCE; Schema: public; Owner: poslog
--

CREATE SEQUENCE poslog_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE poslog_id_seq OWNER TO poslog;

--
-- Name: poslog_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: poslog
--

ALTER SEQUENCE poslog_id_seq OWNED BY poslog.id;


--
-- Name: poslog_merchandise_hierarchy; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog_merchandise_hierarchy (
    transaction_id text NOT NULL,
    sequence_number integer NOT NULL,
    level text NOT NULL,
    level_value integer NOT NULL,
    column5 text,
    column6 text,
    column7 text,
    column8 text,
    column9 text,
    column10 text
);


ALTER TABLE poslog_merchandise_hierarchy OWNER TO poslog;

--
-- Name: poslog_pos_identity; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog_pos_identity (
    transaction_id text NOT NULL,
    sequence_number integer NOT NULL,
    pos_id_type text NOT NULL,
    pos_item_id text,
    qualifier text
);


ALTER TABLE poslog_pos_identity OWNER TO poslog;

--
-- Name: poslog_transaction; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog_transaction (
    transaction_id text NOT NULL,
    retail_store_id smallint NOT NULL,
    workstation_id smallint NOT NULL,
    sequence_number smallint NOT NULL,
    end_date_time timestamp with time zone NOT NULL,
    currency_code text,
    buisness_day_date date NOT NULL
);


ALTER TABLE poslog_transaction OWNER TO poslog;

--
-- Name: poslog_transaction_counts; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog_transaction_counts (
    transaction_id text NOT NULL,
    loyalty_reward_count integer,
    sale_count integer,
    tax_count integer,
    tender_count integer
);


ALTER TABLE poslog_transaction_counts OWNER TO poslog;

--
-- Name: poslog_transaction_operatorid; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog_transaction_operatorid (
    transaction_id text NOT NULL,
    operator_id integer,
    operator_name text
);


ALTER TABLE poslog_transaction_operatorid OWNER TO poslog;

--
-- Name: poslog_transaction_retailtransaction; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog_transaction_retailtransaction (
    transaction_id text NOT NULL,
    version text,
    receipt_date_time timestamp with time zone NOT NULL,
    transaction_count integer
);


ALTER TABLE poslog_transaction_retailtransaction OWNER TO poslog;

--
-- Name: poslog_transaction_retailtransaction_lineitem; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog_transaction_retailtransaction_lineitem (
    transaction_id text NOT NULL,
    entry_mehtod text NOT NULL,
    void_flag boolean,
    sequence_number integer NOT NULL
);


ALTER TABLE poslog_transaction_retailtransaction_lineitem OWNER TO poslog;

--
-- Name: poslog_transaction_retailtransaction_lineitem_loyaltyreward; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog_transaction_retailtransaction_lineitem_loyaltyreward (
    transaction_id text NOT NULL,
    sequence_number integer NOT NULL,
    promotion_id integer,
    event_id integer,
    reason_code text
);


ALTER TABLE poslog_transaction_retailtransaction_lineitem_loyaltyreward OWNER TO poslog;

--
-- Name: poslog_transaction_retailtransaction_lineitem_sale; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog_transaction_retailtransaction_lineitem_sale (
    transaction_id text NOT NULL,
    sequence_number integer NOT NULL,
    item_type text,
    description text,
    discount_amount money,
    extended_amount money,
    extended_discount_amount money,
    item_id text NOT NULL,
    quantity numeric,
    regular_sale_unit_price money
);


ALTER TABLE poslog_transaction_retailtransaction_lineitem_sale OWNER TO poslog;

--
-- Name: poslog_transaction_retailtransaction_lineitem_tax; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog_transaction_retailtransaction_lineitem_tax (
    transaction_id text NOT NULL,
    sequence_number integer NOT NULL,
    amount money,
    percent numeric,
    reason text,
    taxable_amount money
);


ALTER TABLE poslog_transaction_retailtransaction_lineitem_tax OWNER TO poslog;

--
-- Name: poslog_transaction_retailtransaction_lineitem_tender; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog_transaction_retailtransaction_lineitem_tender (
    sequence_number integer NOT NULL,
    tender_type text NOT NULL,
    tender_code text NOT NULL,
    "amount " money,
    "authorization" text,
    cashback money,
    tender_id text,
    transaction_id text NOT NULL
);


ALTER TABLE poslog_transaction_retailtransaction_lineitem_tender OWNER TO poslog;

--
-- Name: poslog_transaction_retailtransaction_total; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog_transaction_retailtransaction_total (
    transaction_id text,
    total_type text,
    text money
);


ALTER TABLE poslog_transaction_retailtransaction_total OWNER TO poslog;

--
-- Name: poslog_transaction_retailtransaction_transactionlink; Type: TABLE; Schema: public; Owner: poslog
--

CREATE TABLE poslog_transaction_retailtransaction_transactionlink (
    transaction_id text NOT NULL,
    entry_method text NOT NULL,
    reason_code text NOT NULL,
    buisness_day_date date,
    retail_store_id integer,
    sequence_number integer NOT NULL,
    workstation_id integer
);


ALTER TABLE poslog_transaction_retailtransaction_transactionlink OWNER TO poslog;

--
-- Name: poslog id; Type: DEFAULT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog ALTER COLUMN id SET DEFAULT nextval('poslog_id_seq'::regclass);


--
-- Name: poslog_merchandise_hierarchy poslog_merchandise_hierarchy_pkey; Type: CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_merchandise_hierarchy
    ADD CONSTRAINT poslog_merchandise_hierarchy_pkey PRIMARY KEY (transaction_id, sequence_number);


--
-- Name: poslog poslog_pkey; Type: CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog
    ADD CONSTRAINT poslog_pkey PRIMARY KEY (id);


--
-- Name: poslog_pos_identity poslog_pos_identity_pkey; Type: CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_pos_identity
    ADD CONSTRAINT poslog_pos_identity_pkey PRIMARY KEY (transaction_id, sequence_number);


--
-- Name: poslog_transaction_counts poslog_transaction_counts_pkey; Type: CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_counts
    ADD CONSTRAINT poslog_transaction_counts_pkey PRIMARY KEY (transaction_id);


--
-- Name: poslog_transaction_operatorid poslog_transaction_operator_id_pkey; Type: CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_operatorid
    ADD CONSTRAINT poslog_transaction_operator_id_pkey PRIMARY KEY (transaction_id);


--
-- Name: poslog_transaction poslog_transaction_pkey; Type: CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction
    ADD CONSTRAINT poslog_transaction_pkey PRIMARY KEY (transaction_id);


--
-- Name: poslog_transaction_retailtransaction_lineitem_loyaltyreward poslog_transaction_retailtransaction_lineitem_loyaltyrewar_pkey; Type: CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction_lineitem_loyaltyreward
    ADD CONSTRAINT poslog_transaction_retailtransaction_lineitem_loyaltyrewar_pkey PRIMARY KEY (transaction_id, sequence_number);


--
-- Name: poslog_transaction_retailtransaction_lineitem poslog_transaction_retailtransaction_lineitem_pkey; Type: CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction_lineitem
    ADD CONSTRAINT poslog_transaction_retailtransaction_lineitem_pkey PRIMARY KEY (transaction_id, sequence_number);


--
-- Name: poslog_transaction_retailtransaction_lineitem_sale poslog_transaction_retailtransaction_lineitem_sale_pkey; Type: CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction_lineitem_sale
    ADD CONSTRAINT poslog_transaction_retailtransaction_lineitem_sale_pkey PRIMARY KEY (transaction_id, sequence_number);


--
-- Name: poslog_transaction_retailtransaction_lineitem_tax poslog_transaction_retailtransaction_lineitem_tax_pkey; Type: CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction_lineitem_tax
    ADD CONSTRAINT poslog_transaction_retailtransaction_lineitem_tax_pkey PRIMARY KEY (transaction_id, sequence_number);


--
-- Name: poslog_transaction_retailtransaction_lineitem_tender poslog_transaction_retailtransaction_lineitem_tender_pkey; Type: CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction_lineitem_tender
    ADD CONSTRAINT poslog_transaction_retailtransaction_lineitem_tender_pkey PRIMARY KEY (transaction_id, sequence_number);


--
-- Name: poslog_transaction_retailtransaction poslog_transaction_retailtransaction_pkey; Type: CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction
    ADD CONSTRAINT poslog_transaction_retailtransaction_pkey PRIMARY KEY (transaction_id);


--
-- Name: poslog_transaction_retailtransaction_transactionlink poslog_transaction_retailtransaction_transactionlink_pkey; Type: CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction_transactionlink
    ADD CONSTRAINT poslog_transaction_retailtransaction_transactionlink_pkey PRIMARY KEY (transaction_id);


--
-- Name: poslog_filename_retail_store_id_buisness_day_date_transaction_c; Type: INDEX; Schema: public; Owner: poslog
--

CREATE UNIQUE INDEX poslog_filename_retail_store_id_buisness_day_date_transaction_c ON poslog USING btree (filename, retail_store_id, buisness_day_date, transaction_count);


--
-- Name: poslog_transaction_buisness_day_date_idx; Type: INDEX; Schema: public; Owner: poslog
--

CREATE INDEX poslog_transaction_buisness_day_date_idx ON poslog_transaction USING btree (buisness_day_date);


--
-- Name: poslog_transaction_retail_store_id_buisness_day_date_idx; Type: INDEX; Schema: public; Owner: poslog
--

CREATE INDEX poslog_transaction_retail_store_id_buisness_day_date_idx ON poslog_transaction USING btree (retail_store_id, buisness_day_date);


--
-- Name: poslog_transaction_retailtransaction_lineitem_tender_tender_cod; Type: INDEX; Schema: public; Owner: poslog
--

CREATE INDEX poslog_transaction_retailtransaction_lineitem_tender_tender_cod ON poslog_transaction_retailtransaction_lineitem_tender USING btree (tender_code);


--
-- Name: poslog_transaction_retailtransaction_lineitem_tender_tender_typ; Type: INDEX; Schema: public; Owner: poslog
--

CREATE INDEX poslog_transaction_retailtransaction_lineitem_tender_tender_typ ON poslog_transaction_retailtransaction_lineitem_tender USING btree (tender_type);


--
-- Name: poslog_transaction_retailtransaction_lineitem_transaction_id_se; Type: INDEX; Schema: public; Owner: poslog
--

CREATE UNIQUE INDEX poslog_transaction_retailtransaction_lineitem_transaction_id_se ON poslog_transaction_retailtransaction_lineitem USING btree (transaction_id, sequence_number);


--
-- Name: poslog_merchandise_hierarchy poslog_merchandise_hierarchy_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_merchandise_hierarchy
    ADD CONSTRAINT poslog_merchandise_hierarchy_transaction_id_fkey FOREIGN KEY (transaction_id, sequence_number) REFERENCES poslog_transaction_retailtransaction_lineitem(transaction_id, sequence_number) ON DELETE CASCADE;


--
-- Name: poslog_pos_identity poslog_pos_identity_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_pos_identity
    ADD CONSTRAINT poslog_pos_identity_transaction_id_fkey FOREIGN KEY (transaction_id, sequence_number) REFERENCES poslog_transaction_retailtransaction_lineitem(transaction_id, sequence_number) ON DELETE CASCADE;


--
-- Name: poslog_transaction_counts poslog_transaction_counts_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_counts
    ADD CONSTRAINT poslog_transaction_counts_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES poslog_transaction(transaction_id) ON DELETE CASCADE;


--
-- Name: poslog_transaction_operatorid poslog_transaction_operator_id_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_operatorid
    ADD CONSTRAINT poslog_transaction_operator_id_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES poslog_transaction(transaction_id) ON DELETE CASCADE;


--
-- Name: poslog_transaction_retailtransaction_lineitem_tax poslog_transaction_retailtransaction_linei_transaction_id_fkey1; Type: FK CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction_lineitem_tax
    ADD CONSTRAINT poslog_transaction_retailtransaction_linei_transaction_id_fkey1 FOREIGN KEY (transaction_id, sequence_number) REFERENCES poslog_transaction_retailtransaction_lineitem(transaction_id, sequence_number) ON UPDATE CASCADE;


--
-- Name: poslog_transaction_retailtransaction_lineitem_loyaltyreward poslog_transaction_retailtransaction_linei_transaction_id_fkey2; Type: FK CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction_lineitem_loyaltyreward
    ADD CONSTRAINT poslog_transaction_retailtransaction_linei_transaction_id_fkey2 FOREIGN KEY (transaction_id, sequence_number) REFERENCES poslog_transaction_retailtransaction_lineitem(transaction_id, sequence_number) ON DELETE CASCADE;


--
-- Name: poslog_transaction_retailtransaction_lineitem_tender poslog_transaction_retailtransaction_linei_transaction_id_fkey3; Type: FK CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction_lineitem_tender
    ADD CONSTRAINT poslog_transaction_retailtransaction_linei_transaction_id_fkey3 FOREIGN KEY (transaction_id, sequence_number) REFERENCES poslog_transaction_retailtransaction_lineitem(transaction_id, sequence_number) ON DELETE CASCADE;


--
-- Name: poslog_transaction_retailtransaction_lineitem_sale poslog_transaction_retailtransaction_linei_transaction_id_fkey4; Type: FK CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction_lineitem_sale
    ADD CONSTRAINT poslog_transaction_retailtransaction_linei_transaction_id_fkey4 FOREIGN KEY (transaction_id, sequence_number) REFERENCES poslog_transaction_retailtransaction_lineitem(transaction_id, sequence_number) ON DELETE CASCADE;


--
-- Name: poslog_merchandise_hierarchy poslog_transaction_retailtransaction_linei_transaction_id_fkey5; Type: FK CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_merchandise_hierarchy
    ADD CONSTRAINT poslog_transaction_retailtransaction_linei_transaction_id_fkey5 FOREIGN KEY (transaction_id, sequence_number) REFERENCES poslog_transaction_retailtransaction_lineitem_sale(transaction_id, sequence_number) ON DELETE CASCADE;


--
-- Name: poslog_transaction_retailtransaction_lineitem poslog_transaction_retailtransaction_lineit_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction_lineitem
    ADD CONSTRAINT poslog_transaction_retailtransaction_lineit_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES poslog_transaction_retailtransaction(transaction_id) ON DELETE CASCADE;


--
-- Name: poslog_transaction_retailtransaction_total poslog_transaction_retailtransaction_total_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction_total
    ADD CONSTRAINT poslog_transaction_retailtransaction_total_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES poslog_transaction_retailtransaction(transaction_id) ON DELETE CASCADE;


--
-- Name: poslog_transaction_retailtransaction_transactionlink poslog_transaction_retailtransaction_transa_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction_transactionlink
    ADD CONSTRAINT poslog_transaction_retailtransaction_transa_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES poslog_transaction_retailtransaction(transaction_id) ON DELETE CASCADE;


--
-- Name: poslog_transaction_retailtransaction poslog_transaction_retailtransaction_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: poslog
--

ALTER TABLE ONLY poslog_transaction_retailtransaction
    ADD CONSTRAINT poslog_transaction_retailtransaction_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES poslog_transaction(transaction_id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

