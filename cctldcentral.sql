--
-- Roles
--

CREATE ROLE cctldcentral;
ALTER ROLE cctldcentral WITH NOSUPERUSER INHERIT NOCREATEROLE NOCREATEDB LOGIN NOREPLICATION NOBYPASSRLS PASSWORD 'md578e0745d7f14ffd47c0a6bff808da2a4';

--
-- Database creation
--

CREATE DATABASE cctldcentral WITH TEMPLATE = template0 OWNER = postgres;
GRANT CONNECT ON DATABASE cctldcentral TO postgres;
GRANT CONNECT ON DATABASE cctldcentral TO cctldcentral;

\connect cctldcentral

--
-- Name: registered_domains; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE registered_domains (
    date TIMESTAMP,
    number INTEGER,
    cctld VARCHAR
);

--
-- Name: registered_domains_date_cctld_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY registered_domains
    ADD CONSTRAINT registered_domains_date_cctld_key UNIQUE (date, cctld);

--
-- Name: registered_domains; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE registered_domains TO cctldcentral;