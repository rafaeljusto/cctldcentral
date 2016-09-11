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
    cctld VARCHAR,
    date TIMESTAMP,
    number INTEGER
);

--
-- Name: registered_domains_cctld_date_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY registered_domains
    ADD CONSTRAINT registered_domains_cctld_date_key UNIQUE (cctld, date);

--
-- Name: registered_domains; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE registered_domains TO cctldcentral;