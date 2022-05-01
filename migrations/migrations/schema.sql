--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4
-- Dumped by pg_dump version 13.4

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
-- Name: products; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA products;


ALTER SCHEMA products OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: orders; Type: TABLE; Schema: products; Owner: postgres
--

CREATE TABLE products.orders (
    id integer NOT NULL,
    widget_id integer,
    transaction_id integer,
    status_id integer,
    quantity integer,
    amount integer,
    created_at date DEFAULT now(),
    updated_at date DEFAULT now()
);


ALTER TABLE products.orders OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: products; Owner: postgres
--

CREATE SEQUENCE products.orders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE products.orders_id_seq OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: products; Owner: postgres
--

ALTER SEQUENCE products.orders_id_seq OWNED BY products.orders.id;


--
-- Name: statuses; Type: TABLE; Schema: products; Owner: postgres
--

CREATE TABLE products.statuses (
    id integer NOT NULL,
    name text,
    created_at date DEFAULT now(),
    updated_at date DEFAULT now()
);


ALTER TABLE products.statuses OWNER TO postgres;

--
-- Name: statuses_id_seq; Type: SEQUENCE; Schema: products; Owner: postgres
--

CREATE SEQUENCE products.statuses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE products.statuses_id_seq OWNER TO postgres;

--
-- Name: statuses_id_seq; Type: SEQUENCE OWNED BY; Schema: products; Owner: postgres
--

ALTER SEQUENCE products.statuses_id_seq OWNED BY products.statuses.id;


--
-- Name: transaction_statuses; Type: TABLE; Schema: products; Owner: postgres
--

CREATE TABLE products.transaction_statuses (
    id integer NOT NULL,
    name text,
    created_at date DEFAULT now(),
    updated_at date DEFAULT now()
);


ALTER TABLE products.transaction_statuses OWNER TO postgres;

--
-- Name: transaction_statuses_id_seq; Type: SEQUENCE; Schema: products; Owner: postgres
--

CREATE SEQUENCE products.transaction_statuses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE products.transaction_statuses_id_seq OWNER TO postgres;

--
-- Name: transaction_statuses_id_seq; Type: SEQUENCE OWNED BY; Schema: products; Owner: postgres
--

ALTER SEQUENCE products.transaction_statuses_id_seq OWNED BY products.transaction_statuses.id;


--
-- Name: transactions; Type: TABLE; Schema: products; Owner: postgres
--

CREATE TABLE products.transactions (
    id integer NOT NULL,
    amount integer,
    currency text,
    last_four text,
    bank_return_code text,
    transaction_status_id integer,
    created_at date DEFAULT now(),
    updated_at date DEFAULT now()
);


ALTER TABLE products.transactions OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: products; Owner: postgres
--

CREATE SEQUENCE products.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE products.transactions_id_seq OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: products; Owner: postgres
--

ALTER SEQUENCE products.transactions_id_seq OWNED BY products.transactions.id;


--
-- Name: users; Type: TABLE; Schema: products; Owner: postgres
--

CREATE TABLE products.users (
    id integer NOT NULL,
    first_name character varying(255),
    last_name character varying(255),
    email text,
    password text,
    created_at date DEFAULT now(),
    updated_at date DEFAULT now()
);


ALTER TABLE products.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: products; Owner: postgres
--

CREATE SEQUENCE products.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE products.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: products; Owner: postgres
--

ALTER SEQUENCE products.users_id_seq OWNED BY products.users.id;


--
-- Name: widgets; Type: TABLE; Schema: products; Owner: postgres
--

CREATE TABLE products.widgets (
    id integer NOT NULL,
    name text DEFAULT ''::text,
    description text DEFAULT ''::text,
    inventory_level integer,
    price integer,
    created_at date DEFAULT now(),
    updated_at date DEFAULT now()
);


ALTER TABLE products.widgets OWNER TO postgres;

--
-- Name: widgets_id_seq; Type: SEQUENCE; Schema: products; Owner: postgres
--

CREATE SEQUENCE products.widgets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE products.widgets_id_seq OWNER TO postgres;

--
-- Name: widgets_id_seq; Type: SEQUENCE OWNED BY; Schema: products; Owner: postgres
--

ALTER SEQUENCE products.widgets_id_seq OWNED BY products.widgets.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: orders id; Type: DEFAULT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.orders ALTER COLUMN id SET DEFAULT nextval('products.orders_id_seq'::regclass);


--
-- Name: statuses id; Type: DEFAULT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.statuses ALTER COLUMN id SET DEFAULT nextval('products.statuses_id_seq'::regclass);


--
-- Name: transaction_statuses id; Type: DEFAULT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.transaction_statuses ALTER COLUMN id SET DEFAULT nextval('products.transaction_statuses_id_seq'::regclass);


--
-- Name: transactions id; Type: DEFAULT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.transactions ALTER COLUMN id SET DEFAULT nextval('products.transactions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.users ALTER COLUMN id SET DEFAULT nextval('products.users_id_seq'::regclass);


--
-- Name: widgets id; Type: DEFAULT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.widgets ALTER COLUMN id SET DEFAULT nextval('products.widgets_id_seq'::regclass);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: statuses statuses_pkey; Type: CONSTRAINT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.statuses
    ADD CONSTRAINT statuses_pkey PRIMARY KEY (id);


--
-- Name: transaction_statuses transaction_statuses_pkey; Type: CONSTRAINT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.transaction_statuses
    ADD CONSTRAINT transaction_statuses_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: widgets widgets_pkey; Type: CONSTRAINT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.widgets
    ADD CONSTRAINT widgets_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: orders products.orders_products.statuses_id_fk; Type: FK CONSTRAINT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.orders
    ADD CONSTRAINT "products.orders_products.statuses_id_fk" FOREIGN KEY (status_id) REFERENCES products.statuses(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: orders products.orders_products.transactions_id_fk; Type: FK CONSTRAINT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.orders
    ADD CONSTRAINT "products.orders_products.transactions_id_fk" FOREIGN KEY (transaction_id) REFERENCES products.transactions(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: orders products.orders_products.widgets_id_fk; Type: FK CONSTRAINT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.orders
    ADD CONSTRAINT "products.orders_products.widgets_id_fk" FOREIGN KEY (widget_id) REFERENCES products.widgets(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: transactions products.transactions_products.transaction_statuses_id_fk; Type: FK CONSTRAINT; Schema: products; Owner: postgres
--

ALTER TABLE ONLY products.transactions
    ADD CONSTRAINT "products.transactions_products.transaction_statuses_id_fk" FOREIGN KEY (transaction_status_id) REFERENCES products.transaction_statuses(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

