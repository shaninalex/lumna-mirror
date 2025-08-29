--
-- PostgreSQL database dump
--

-- Dumped from database version 15.13 (Debian 15.13-1.pgdg120+1)
-- Dumped by pg_dump version 15.13 (Debian 15.13-1.pgdg120+1)

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
-- Name: keto; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA keto;


ALTER SCHEMA keto OWNER TO postgres;

--
-- Name: kratos; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA kratos;


ALTER SCHEMA kratos OWNER TO postgres;

--
-- Name: test; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA test;


ALTER SCHEMA test OWNER TO postgres;

--
-- Name: btree_gin; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS btree_gin WITH SCHEMA kratos;


--
-- Name: EXTENSION btree_gin; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION btree_gin IS 'support for indexing common datatypes in GIN';


--
-- Name: pg_trgm; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA kratos;


--
-- Name: EXTENSION pg_trgm; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: continuity_containers; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.continuity_containers (
    id uuid NOT NULL,
    identity_id uuid,
    name character varying(255) NOT NULL,
    payload jsonb,
    expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    nid uuid
);


ALTER TABLE kratos.continuity_containers OWNER TO postgres;

--
-- Name: courier_message_dispatches; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.courier_message_dispatches (
    id uuid NOT NULL,
    message_id uuid NOT NULL,
    status character varying(7) NOT NULL,
    error json,
    nid uuid NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE kratos.courier_message_dispatches OWNER TO postgres;

--
-- Name: courier_messages; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.courier_messages (
    id uuid NOT NULL,
    type integer NOT NULL,
    status integer NOT NULL,
    body text NOT NULL,
    subject character varying(255) NOT NULL,
    recipient character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    template_type character varying(255) DEFAULT ''::character varying NOT NULL,
    template_data bytea,
    nid uuid,
    send_count integer DEFAULT 0 NOT NULL,
    channel character varying(32)
);


ALTER TABLE kratos.courier_messages OWNER TO postgres;

--
-- Name: identities; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.identities (
    id uuid NOT NULL,
    schema_id character varying(2048) NOT NULL,
    traits jsonb NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    nid uuid,
    state character varying(255) DEFAULT 'active'::character varying NOT NULL,
    state_changed_at timestamp without time zone,
    metadata_public jsonb,
    metadata_admin jsonb,
    available_aal character varying(4),
    organization_id uuid
);


ALTER TABLE kratos.identities OWNER TO postgres;

--
-- Name: identity_credential_identifiers; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.identity_credential_identifiers (
    id uuid NOT NULL,
    identifier character varying(255) NOT NULL,
    identity_credential_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    nid uuid,
    identity_credential_type_id uuid NOT NULL
);


ALTER TABLE kratos.identity_credential_identifiers OWNER TO postgres;

--
-- Name: identity_credential_types; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.identity_credential_types (
    id uuid NOT NULL,
    name character varying(32) NOT NULL
);


ALTER TABLE kratos.identity_credential_types OWNER TO postgres;

--
-- Name: identity_credentials; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.identity_credentials (
    id uuid NOT NULL,
    config jsonb NOT NULL,
    identity_credential_type_id uuid NOT NULL,
    identity_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    nid uuid,
    version integer DEFAULT 0 NOT NULL
);


ALTER TABLE kratos.identity_credentials OWNER TO postgres;

--
-- Name: identity_login_codes; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.identity_login_codes (
    id uuid NOT NULL,
    code character varying(64) NOT NULL,
    address character varying(255) NOT NULL,
    address_type character(36) NOT NULL,
    used_at timestamp without time zone,
    expires_at timestamp without time zone DEFAULT '2000-01-01 00:00:00'::timestamp without time zone NOT NULL,
    issued_at timestamp without time zone DEFAULT '2000-01-01 00:00:00'::timestamp without time zone NOT NULL,
    selfservice_login_flow_id uuid NOT NULL,
    identity_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    nid uuid NOT NULL
);


ALTER TABLE kratos.identity_login_codes OWNER TO postgres;

--
-- Name: identity_recovery_addresses; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.identity_recovery_addresses (
    id uuid NOT NULL,
    via character varying(16) NOT NULL,
    value character varying(400) NOT NULL,
    identity_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    nid uuid
);


ALTER TABLE kratos.identity_recovery_addresses OWNER TO postgres;

--
-- Name: identity_recovery_codes; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.identity_recovery_codes (
    id uuid NOT NULL,
    code character varying(64) NOT NULL,
    used_at timestamp without time zone,
    identity_recovery_address_id uuid,
    code_type integer NOT NULL,
    expires_at timestamp without time zone DEFAULT '2000-01-01 00:00:00'::timestamp without time zone NOT NULL,
    issued_at timestamp without time zone DEFAULT '2000-01-01 00:00:00'::timestamp without time zone NOT NULL,
    selfservice_recovery_flow_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    nid uuid NOT NULL,
    identity_id uuid NOT NULL
);


ALTER TABLE kratos.identity_recovery_codes OWNER TO postgres;

--
-- Name: identity_recovery_tokens; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.identity_recovery_tokens (
    id uuid NOT NULL,
    token character varying(64) NOT NULL,
    used boolean DEFAULT false NOT NULL,
    used_at timestamp without time zone,
    identity_recovery_address_id uuid,
    selfservice_recovery_flow_id uuid,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    expires_at timestamp without time zone DEFAULT '2000-01-01 00:00:00'::timestamp without time zone NOT NULL,
    issued_at timestamp without time zone DEFAULT '2000-01-01 00:00:00'::timestamp without time zone NOT NULL,
    nid uuid,
    identity_id uuid NOT NULL,
    token_type integer DEFAULT 0 NOT NULL,
    CONSTRAINT identity_recovery_tokens_token_type_ck CHECK (((token_type = 1) OR (token_type = 2)))
);


ALTER TABLE kratos.identity_recovery_tokens OWNER TO postgres;

--
-- Name: identity_registration_codes; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.identity_registration_codes (
    id uuid NOT NULL,
    code character varying(64) NOT NULL,
    address character varying(255) NOT NULL,
    address_type character(36) NOT NULL,
    used_at timestamp without time zone,
    expires_at timestamp without time zone DEFAULT '2000-01-01 00:00:00'::timestamp without time zone NOT NULL,
    issued_at timestamp without time zone DEFAULT '2000-01-01 00:00:00'::timestamp without time zone NOT NULL,
    selfservice_registration_flow_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    nid uuid NOT NULL
);


ALTER TABLE kratos.identity_registration_codes OWNER TO postgres;

--
-- Name: identity_verifiable_addresses; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.identity_verifiable_addresses (
    id uuid NOT NULL,
    status character varying(16) NOT NULL,
    via character varying(16) NOT NULL,
    verified boolean NOT NULL,
    value character varying(400) NOT NULL,
    verified_at timestamp without time zone,
    identity_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    nid uuid
);


ALTER TABLE kratos.identity_verifiable_addresses OWNER TO postgres;

--
-- Name: identity_verification_codes; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.identity_verification_codes (
    id uuid NOT NULL,
    code_hmac character varying(64) NOT NULL,
    used_at timestamp without time zone,
    identity_verifiable_address_id uuid,
    expires_at timestamp without time zone DEFAULT '2000-01-01 00:00:00'::timestamp without time zone NOT NULL,
    issued_at timestamp without time zone DEFAULT '2000-01-01 00:00:00'::timestamp without time zone NOT NULL,
    selfservice_verification_flow_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    nid uuid NOT NULL
);


ALTER TABLE kratos.identity_verification_codes OWNER TO postgres;

--
-- Name: identity_verification_tokens; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.identity_verification_tokens (
    id uuid NOT NULL,
    token character varying(64) NOT NULL,
    used boolean DEFAULT false NOT NULL,
    used_at timestamp without time zone,
    expires_at timestamp without time zone NOT NULL,
    issued_at timestamp without time zone NOT NULL,
    identity_verifiable_address_id uuid NOT NULL,
    selfservice_verification_flow_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    nid uuid
);


ALTER TABLE kratos.identity_verification_tokens OWNER TO postgres;

--
-- Name: networks; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.networks (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE kratos.networks OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.schema_migration (
    version character varying(48) NOT NULL,
    version_self integer DEFAULT 0 NOT NULL
);


ALTER TABLE kratos.schema_migration OWNER TO postgres;

--
-- Name: selfservice_errors; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.selfservice_errors (
    id uuid NOT NULL,
    errors jsonb NOT NULL,
    seen_at timestamp without time zone,
    was_seen boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    csrf_token character varying(255) DEFAULT ''::character varying NOT NULL,
    nid uuid
);


ALTER TABLE kratos.selfservice_errors OWNER TO postgres;

--
-- Name: selfservice_login_flows; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.selfservice_login_flows (
    id uuid NOT NULL,
    request_url text NOT NULL,
    issued_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    active_method character varying(32) NOT NULL,
    csrf_token character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    forced boolean DEFAULT false NOT NULL,
    type character varying(16) DEFAULT 'browser'::character varying NOT NULL,
    ui jsonb,
    nid uuid,
    requested_aal character varying(4) DEFAULT 'aal1'::character varying NOT NULL,
    internal_context jsonb NOT NULL,
    oauth2_login_challenge uuid,
    oauth2_login_challenge_data text,
    state character varying(255),
    submit_count integer DEFAULT 0 NOT NULL,
    organization_id uuid
);


ALTER TABLE kratos.selfservice_login_flows OWNER TO postgres;

--
-- Name: selfservice_recovery_flows; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.selfservice_recovery_flows (
    id uuid NOT NULL,
    request_url text NOT NULL,
    issued_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    active_method character varying(32),
    csrf_token character varying(255) NOT NULL,
    state character varying(32) NOT NULL,
    recovered_identity_id uuid,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    type character varying(16) DEFAULT 'browser'::character varying NOT NULL,
    ui jsonb,
    nid uuid,
    submit_count integer DEFAULT 0 NOT NULL,
    skip_csrf_check boolean DEFAULT false NOT NULL
);


ALTER TABLE kratos.selfservice_recovery_flows OWNER TO postgres;

--
-- Name: selfservice_registration_flows; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.selfservice_registration_flows (
    id uuid NOT NULL,
    request_url text NOT NULL,
    issued_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    active_method character varying(32) NOT NULL,
    csrf_token character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    type character varying(16) DEFAULT 'browser'::character varying NOT NULL,
    ui jsonb,
    nid uuid,
    internal_context jsonb NOT NULL,
    oauth2_login_challenge uuid,
    oauth2_login_challenge_data text,
    state character varying(255),
    submit_count integer DEFAULT 0 NOT NULL,
    organization_id uuid
);


ALTER TABLE kratos.selfservice_registration_flows OWNER TO postgres;

--
-- Name: selfservice_settings_flows; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.selfservice_settings_flows (
    id uuid NOT NULL,
    request_url text NOT NULL,
    issued_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    identity_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    active_method character varying(32),
    state character varying(255) DEFAULT 'show_form'::character varying NOT NULL,
    type character varying(16) DEFAULT 'browser'::character varying NOT NULL,
    ui jsonb,
    nid uuid,
    internal_context jsonb NOT NULL
);


ALTER TABLE kratos.selfservice_settings_flows OWNER TO postgres;

--
-- Name: selfservice_verification_flows; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.selfservice_verification_flows (
    id uuid NOT NULL,
    request_url text NOT NULL,
    issued_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    csrf_token character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    type character varying(16) DEFAULT 'browser'::character varying NOT NULL,
    state character varying(255) DEFAULT 'show_form'::character varying NOT NULL,
    active_method character varying(32),
    ui jsonb,
    nid uuid,
    submit_count integer DEFAULT 0 NOT NULL,
    oauth2_login_challenge text,
    session_id uuid,
    identity_id uuid,
    authentication_methods json
);


ALTER TABLE kratos.selfservice_verification_flows OWNER TO postgres;

--
-- Name: session_devices; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.session_devices (
    id uuid NOT NULL,
    ip_address character varying(50) DEFAULT ''::character varying,
    user_agent character varying(512) DEFAULT ''::character varying,
    location character varying(512) DEFAULT ''::character varying,
    nid uuid NOT NULL,
    session_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE kratos.session_devices OWNER TO postgres;

--
-- Name: session_token_exchanges; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.session_token_exchanges (
    id uuid NOT NULL,
    nid uuid NOT NULL,
    flow_id uuid NOT NULL,
    session_id uuid,
    init_code character varying(64) NOT NULL,
    return_to_code character varying(64) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE kratos.session_token_exchanges OWNER TO postgres;

--
-- Name: sessions; Type: TABLE; Schema: kratos; Owner: postgres
--

CREATE TABLE kratos.sessions (
    id uuid NOT NULL,
    issued_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    authenticated_at timestamp without time zone NOT NULL,
    identity_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    token character varying(39),
    active boolean DEFAULT false,
    nid uuid,
    logout_token character varying(39),
    aal character varying(4) DEFAULT 'aal1'::character varying NOT NULL,
    authentication_methods jsonb NOT NULL
);


ALTER TABLE kratos.sessions OWNER TO postgres;

--
-- Name: epics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.epics (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid,
    project_id uuid,
    title text,
    description text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.epics OWNER TO postgres;

--
-- Name: organizations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.organizations (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id text,
    title text,
    description text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.organizations OWNER TO postgres;

--
-- Name: projects; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.projects (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid,
    organization_id uuid,
    title text,
    project_key text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.projects OWNER TO postgres;

--
-- Name: sprints; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sprints (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid,
    organization_id uuid,
    title text,
    description text,
    start_date timestamp with time zone,
    end_date timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.sprints OWNER TO postgres;

--
-- Name: task_statuses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.task_statuses (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    project_id uuid,
    title text,
    description text,
    complete boolean,
    index bigint,
    config text
);


ALTER TABLE public.task_statuses OWNER TO postgres;

--
-- Name: tasks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tasks (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid,
    epic_id uuid,
    organization_id uuid,
    sprint_id uuid,
    project_id uuid,
    task_status_id uuid,
    assignee text,
    title text,
    completed boolean,
    description text,
    list_index bigint,
    code text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.tasks OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    organization_id uuid,
    settings text,
    code text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: continuity_containers; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.continuity_containers (id, identity_id, name, payload, expires_at, created_at, updated_at, nid) FROM stdin;
\.


--
-- Data for Name: courier_message_dispatches; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.courier_message_dispatches (id, message_id, status, error, nid, created_at, updated_at) FROM stdin;
134aa173-f27d-47ad-a3af-63b03fa5090b	06cd8418-93c3-44c4-98a0-2d9c96e33b82	success	null	ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0	2025-08-24 19:44:29.641745	2025-08-24 19:44:29.641745
\.


--
-- Data for Name: courier_messages; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.courier_messages (id, type, status, body, subject, recipient, created_at, updated_at, template_type, template_data, nid, send_count, channel) FROM stdin;
06cd8418-93c3-44c4-98a0-2d9c96e33b82	1	2	Hi,\n\nplease verify your account by entering the following code:\n\n292569\n\nor clicking the following link:\n\nhttp://localhost:4433/self-service/verification?code=292569&flow=4d6d06b4-65db-4d89-93a8-efed5eab0dc4\n	Please verify your email address	test@test.com	2025-08-24 19:44:28.98192	2025-08-24 19:44:28.98192	verification_code_valid	\\x7b22746f223a227465737440746573742e636f6d222c22766572696669636174696f6e5f75726c223a22687474703a2f2f6c6f63616c686f73743a343433332f73656c662d736572766963652f766572696669636174696f6e3f636f64653d3239323536395c7530303236666c6f773d34643664303662342d363564622d346438392d393361382d656665643565616230646334222c22766572696669636174696f6e5f636f6465223a22323932353639222c226964656e74697479223a7b22637265617465645f6174223a22323032352d30382d32345431363a34343a32382e3934333534375a222c226964223a2235346636636232622d356532302d343932382d616137332d373131613561613866356131222c226d657461646174615f7075626c6963223a6e756c6c2c226f7267616e697a6174696f6e5f6964223a6e756c6c2c227265636f766572795f616464726573736573223a5b7b22637265617465645f6174223a22323032352d30382d32345431363a34343a32382e3935313434375a222c226964223a2230633339363565632d613331352d346435662d623264372d356566623365393963356162222c22757064617465645f6174223a22323032352d30382d32345431363a34343a32382e3935313434375a222c2276616c7565223a227465737440746573742e636f6d222c22766961223a22656d61696c227d5d2c22736368656d615f6964223a2264656661756c74222c22736368656d615f75726c223a22687474703a2f2f6c6f63616c686f73743a343433332f736368656d61732f5a47566d595856736441222c227374617465223a22616374697665222c2273746174655f6368616e6765645f6174223a22323032352d30382d32345431393a34343a32382e3934303232313438372b30333a3030222c22747261697473223a7b22656d61696c223a227465737440746573742e636f6d222c226e616d65223a7b226669727374223a22416c6578222c226c617374223a225368616e696e227d7d2c22757064617465645f6174223a22323032352d30382d32345431363a34343a32382e3934333534375a222c2276657269666961626c655f616464726573736573223a5b7b22637265617465645f6174223a22323032352d30382d32345431363a34343a32382e39343738375a222c226964223a2234376239323036612d626237632d346461302d393834322d323730343162663932336166222c22737461747573223a2270656e64696e67222c22757064617465645f6174223a22323032352d30382d32345431363a34343a32382e39343738375a222c2276616c7565223a227465737440746573742e636f6d222c227665726966696564223a66616c73652c22766961223a22656d61696c227d5d7d2c22726571756573745f75726c223a22687474703a2f2f6c6f63616c686f73743a343433332f73656c662d736572766963652f726567697374726174696f6e2f62726f77736572222c227472616e7369656e745f7061796c6f6164223a7b7d7d	ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0	1	email
\.


--
-- Data for Name: identities; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.identities (id, schema_id, traits, created_at, updated_at, nid, state, state_changed_at, metadata_public, metadata_admin, available_aal, organization_id) FROM stdin;
54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	default	{"name": {"last": "Shanin", "first": "Alex"}, "email": "test@test.com"}	2025-08-24 16:44:28.943547	2025-08-24 16:44:28.943547	ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0	active	2025-08-24 19:44:28.940221	\N	\N	aal1	\N
\.


--
-- Data for Name: identity_credential_identifiers; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.identity_credential_identifiers (id, identifier, identity_credential_id, created_at, updated_at, nid, identity_credential_type_id) FROM stdin;
9130e4fe-a428-4101-b012-054e8cbdcbe6	test@test.com	b9acf651-a9e9-401d-9b84-ee065772c9e9	2025-08-24 16:44:28.959639	2025-08-24 16:44:28.959639	ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0	78c1b41d-8341-4507-aa60-aff1d4369670
\.


--
-- Data for Name: identity_credential_types; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.identity_credential_types (id, name) FROM stdin;
78c1b41d-8341-4507-aa60-aff1d4369670	password
6fa5e2e0-bfce-4631-b62b-cf2b0252b289	oidc
5e29b036-aa47-457f-9fe6-aa8b854a752b	totp
567a0730-7f48-4dd7-a13d-df87a51c245f	lookup_secret
6b213fa0-e6ad-46cb-8878-b088d2ce2e3c	webauthn
14f3b7e2-8725-4068-be39-8a796485fd97	code
8d0ca544-9bf6-45d3-bd75-0bbb3aeba3c7	passkey
\.


--
-- Data for Name: identity_credentials; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.identity_credentials (id, config, identity_credential_type_id, identity_id, created_at, updated_at, nid, version) FROM stdin;
b9acf651-a9e9-401d-9b84-ee065772c9e9	{"hashed_password": "$2a$08$pWbx.AEcw1Ff5OaGi3Hk6ewiHTPNa6XKviqM6E9HCC1FOtayseZTi"}	78c1b41d-8341-4507-aa60-aff1d4369670	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	2025-08-24 16:44:28.956611	2025-08-24 16:44:28.956611	ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0	0
\.


--
-- Data for Name: identity_login_codes; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.identity_login_codes (id, code, address, address_type, used_at, expires_at, issued_at, selfservice_login_flow_id, identity_id, created_at, updated_at, nid) FROM stdin;
\.


--
-- Data for Name: identity_recovery_addresses; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.identity_recovery_addresses (id, via, value, identity_id, created_at, updated_at, nid) FROM stdin;
0c3965ec-a315-4d5f-b2d7-5efb3e99c5ab	email	test@test.com	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	2025-08-24 16:44:28.951447	2025-08-24 16:44:28.951447	ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0
\.


--
-- Data for Name: identity_recovery_codes; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.identity_recovery_codes (id, code, used_at, identity_recovery_address_id, code_type, expires_at, issued_at, selfservice_recovery_flow_id, created_at, updated_at, nid, identity_id) FROM stdin;
\.


--
-- Data for Name: identity_recovery_tokens; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.identity_recovery_tokens (id, token, used, used_at, identity_recovery_address_id, selfservice_recovery_flow_id, created_at, updated_at, expires_at, issued_at, nid, identity_id, token_type) FROM stdin;
\.


--
-- Data for Name: identity_registration_codes; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.identity_registration_codes (id, code, address, address_type, used_at, expires_at, issued_at, selfservice_registration_flow_id, created_at, updated_at, nid) FROM stdin;
\.


--
-- Data for Name: identity_verifiable_addresses; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.identity_verifiable_addresses (id, status, via, verified, value, verified_at, identity_id, created_at, updated_at, nid) FROM stdin;
47b9206a-bb7c-4da0-9842-27041bf923af	completed	email	t	test@test.com	2025-08-24 16:44:39.682165	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	2025-08-24 16:44:28.94787	2025-08-24 16:44:28.94787	ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0
\.


--
-- Data for Name: identity_verification_codes; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.identity_verification_codes (id, code_hmac, used_at, identity_verifiable_address_id, expires_at, issued_at, selfservice_verification_flow_id, created_at, updated_at, nid) FROM stdin;
45494b49-9a6f-40d9-a942-00ade2b9c4d6	bf351207c482822f112f6a208b66c4946d268e6fed4b0e378556f0fce42b1278	2025-08-24 16:44:39.676779	47b9206a-bb7c-4da0-9842-27041bf923af	2025-08-24 17:44:28.977225	2025-08-24 16:44:28.977225	4d6d06b4-65db-4d89-93a8-efed5eab0dc4	2025-08-24 19:44:28.977466	2025-08-24 19:44:28.977466	ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0
\.


--
-- Data for Name: identity_verification_tokens; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.identity_verification_tokens (id, token, used, used_at, expires_at, issued_at, identity_verifiable_address_id, selfservice_verification_flow_id, created_at, updated_at, nid) FROM stdin;
\.


--
-- Data for Name: networks; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.networks (id, created_at, updated_at) FROM stdin;
ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0	2025-08-24 19:44:05.500761	2025-08-24 19:44:05.500761
\.


--
-- Data for Name: schema_migration; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.schema_migration (version, version_self) FROM stdin;
20150100000001000000	0
20191100000001000000	0
20191100000001000001	0
20191100000001000002	0
20191100000001000003	0
20191100000001000004	0
20191100000001000005	0
20191100000002000000	0
20191100000002000001	0
20191100000002000002	0
20191100000002000003	0
20191100000002000004	0
20191100000003000000	0
20191100000004000000	0
20191100000006000000	0
20191100000007000000	0
20191100000008000000	0
20191100000008000001	0
20191100000008000002	0
20191100000008000003	0
20191100000008000004	0
20191100000008000005	0
20191100000010000000	0
20191100000010000001	0
20191100000011000000	0
20191100000012000000	0
20200317160354000000	0
20200317160354000001	0
20200317160354000002	0
20200317160354000003	0
20200317160354000004	0
20200401183443000000	0
20200402142539000000	0
20200402142539000001	0
20200402142539000002	0
20200519101057000000	0
20200519101057000001	0
20200519101057000002	0
20200519101057000003	0
20200519101057000004	0
20200519101057000005	0
20200519101057000006	0
20200519101057000007	0
20200601101000000000	0
20200605111551000000	0
20200605111551000001	0
20200605111551000002	0
20200607165100000000	0
20200607165100000001	0
20200705105359000000	0
20200810141652000000	0
20200810141652000001	0
20200810141652000002	0
20200810141652000003	0
20200810141652000004	0
20200810161022000000	0
20200810161022000001	0
20200810161022000002	0
20200810161022000003	0
20200810161022000004	0
20200810161022000005	0
20200810161022000006	0
20200810161022000007	0
20200810161022000008	0
20200810162450000000	0
20200810162450000001	0
20200810162450000002	0
20200810162450000003	0
20200812124254000000	0
20200812124254000001	0
20200812124254000002	0
20200812124254000003	0
20200812124254000004	0
20200812160551000000	0
20200830121710000000	0
20200830130642000000	0
20200830130642000001	0
20200830130642000002	0
20200830130642000003	0
20200830130642000004	0
20200830130642000005	0
20200830130642000006	0
20200830130642000007	0
20200830130643000000	0
20200830130644000000	0
20200830130644000001	0
20200830130645000000	0
20200830130646000000	0
20200830130646000001	0
20200830130646000002	0
20200830154602000000	0
20200830154602000001	0
20200830154602000002	0
20200830154602000003	0
20200830154602000004	0
20200830172221000000	0
20200830172221000001	0
20200830172221000002	0
20200830172221000003	0
20200831110752000000	0
20200831110752000001	0
20200831110752000002	0
20200831110752000003	0
20200831110752000004	0
20200831110752000005	0
20200831110752000006	0
20200831110752000007	0
20201201161451000000	0
20201201161451000001	0
20210307130558000000	0
20210307130559000000	0
20210307130559000001	0
20210311102338000000	0
20210311102338000001	0
20210311102338000002	0
20210311102338000003	0
20210311102338000004	0
20210311102338000005	0
20210311102338000006	0
20210311102338000007	0
20210311102338000008	0
20210311102338000009	0
20210311102338000010	0
20210311102338000011	0
20210311102338000012	0
20210311102338000013	0
20210311102338000014	0
20210311102338000015	0
20210311102338000016	0
20210311102338000017	0
20210311102338000018	0
20210311102338000019	0
20210311102338000020	0
20210311102338000021	0
20210311102338000022	0
20210311102338000023	0
20210311102338000024	0
20210410175418000000	0
20210410175418000001	0
20210410175418000002	0
20210410175418000003	0
20210410175418000004	0
20210410175418000005	0
20210410175418000006	0
20210410175418000007	0
20210410175418000008	0
20210410175418000009	0
20210410175418000010	0
20210410175418000011	0
20210410175418000012	0
20210410175418000013	0
20210410175418000014	0
20210410175418000015	0
20210410175418000016	0
20210410175418000017	0
20210410175418000018	0
20210410175418000019	0
20210410175418000020	0
20210410175418000021	0
20210410175418000022	0
20210410175418000023	0
20210410175418000024	0
20210410175418000025	0
20210410175418000026	0
20210410175418000027	0
20210410175418000028	0
20210410175418000029	0
20210410175418000030	0
20210410175418000031	0
20210410175418000032	0
20210410175418000033	0
20210410175418000034	0
20210410175418000035	0
20210410175418000036	0
20210410175418000037	0
20210410175418000038	0
20210410175418000039	0
20210410175418000040	0
20210410175418000041	0
20210410175418000042	0
20210410175418000043	0
20210410175418000044	0
20210410175418000045	0
20210410175418000046	0
20210410175418000047	0
20210410175418000048	0
20210410175418000049	0
20210410175418000050	0
20210410175418000051	0
20210410175418000052	0
20210410175418000053	0
20210410175418000054	0
20210410175418000055	0
20210410175418000056	0
20210410175418000057	0
20210410175418000058	0
20210410175418000059	0
20210410175418000060	0
20210410175418000061	0
20210410175418000062	0
20210410175418000063	0
20210410175418000064	0
20210410175418000065	0
20210410175418000066	0
20210410175418000067	0
20210410175418000068	0
20210410175418000069	0
20210410175418000070	0
20210410175418000071	0
20210410175418000072	0
20210410175418000073	0
20210410175418000074	0
20210410175418000075	0
20210410175418000076	0
20210410175418000077	0
20210410175418000078	0
20210410175418000079	0
20210410175418000080	0
20210410175418000081	0
20210410175418000082	0
20210410175418000083	0
20210410175418000084	0
20210410175418000085	0
20210410175418000086	0
20210410175418000087	0
20210410175418000088	0
20210410175418000089	0
20210504121624000000	0
20210504121624000001	0
20210618103120000000	0
20210618103120000001	0
20210618103120000002	0
20210618103120000003	0
20210618103120000004	0
20210805112414000000	0
20210805112414000001	0
20210805112414000002	0
20210805122535000000	0
20210810153530000000	0
20210810153530000001	0
20210810153530000002	0
20210810153530000003	0
20210810153530000004	0
20210813150152000000	0
20210816113956000000	0
20210816142650000000	0
20210816142650000001	0
20210816142650000002	0
20210816142650000003	0
20210816142650000004	0
20210816142650000005	0
20210817181232000000	0
20210817181232000001	0
20210817181232000002	0
20210817181232000003	0
20210817181232000004	0
20210817181232000005	0
20210829131458000000	0
20210913095309000000	0
20210913095309000001	0
20210913095309000002	0
20210913095309000003	0
20210913095309000004	0
20220118104539000000	0
20220118104539000001	0
20220118104539000002	0
20220118104539000003	0
20220301102701000000	0
20220301102701000001	0
20220420102701000000	0
20220512102703000000	0
20220607000001000000	0
20220610155809000000	0
20220802103909000000	0
20220824165300000000	0
20220824165300000001	0
20220824165300000002	0
20220825134336000000	0
20220825134336000001	0
20220901123209000000	0
20220907132836000000	0
20221024182336000000	0
20221205092803000000	0
20221214101328000000	0
20221220124639000000	0
20230104193739000000	0
20230216142104000000	0
20230313141439000000	0
20230313141439000001	0
20230322144139000001	0
20230405000000000001	0
20230614000001000000	0
20230619000000000001	0
20230626000000000001	0
20230703143600000001	0
20230705000000000001	0
20230706000000000001	0
20230707133700000000	0
20230707133700000001	0
20230712173852000000	0
20230811000000000001	0
20230818000000000001	0
20230823000000000001	0
20230907085000000000	0
20230920171028000000	0
20231108111100000000	0
20231130094628000000	0
20240119094628000000	0
20240213095000000000	0
20240214113828000000	0
20240221000000000000	0
20240318143139000000	0
20240325153839000000	0
20240425095000000000	0
20240425095000000001	0
\.


--
-- Data for Name: selfservice_errors; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.selfservice_errors (id, errors, seen_at, was_seen, created_at, updated_at, csrf_token, nid) FROM stdin;
\.


--
-- Data for Name: selfservice_login_flows; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.selfservice_login_flows (id, request_url, issued_at, expires_at, active_method, csrf_token, created_at, updated_at, forced, type, ui, nid, requested_aal, internal_context, oauth2_login_challenge, oauth2_login_challenge_data, state, submit_count, organization_id) FROM stdin;
98496454-fec9-4922-8ab1-64e64c3ec1e4	http://localhost:4433/self-service/login/browser	2025-08-24 16:44:46.170642	2025-08-24 16:54:46.170642	password	iWmxaPxi4GQKetJTDUkdAczB2PLu0WR+GSlxyDOkBsK2bYjmUzOkyp7c2ZbAYDOX/tlC8Px6WXlDeUE1OmfSHg==	2025-08-24 19:44:46.178312	2025-08-24 19:44:46.178312	f	browser	{"nodes": [{"meta": {"label": {"id": 1010002, "text": "Sign in with google", "type": "info", "context": {"provider": "google", "provider_id": "google"}}}, "type": "input", "group": "oidc", "messages": [], "attributes": {"name": "provider", "type": "submit", "value": "google", "disabled": false, "node_type": "input"}}, {"meta": {}, "type": "input", "group": "default", "messages": [], "attributes": {"name": "csrf_token", "type": "hidden", "value": "iWmxaPxi4GQKetJTDUkdAczB2PLu0WR+GSlxyDOkBsK2bYjmUzOkyp7c2ZbAYDOX/tlC8Px6WXlDeUE1OmfSHg==", "disabled": false, "required": true, "node_type": "input"}}, {"meta": {"label": {"id": 1070002, "text": "E-Mail", "type": "info", "context": {"title": "E-Mail"}}}, "type": "input", "group": "default", "messages": [], "attributes": {"name": "identifier", "type": "text", "value": "", "disabled": false, "required": true, "node_type": "input"}}, {"meta": {"label": {"id": 1070001, "text": "Password", "type": "info"}}, "type": "input", "group": "password", "messages": [], "attributes": {"name": "password", "type": "password", "disabled": false, "required": true, "node_type": "input", "autocomplete": "current-password"}}, {"meta": {"label": {"id": 1010022, "text": "Sign in with password", "type": "info"}}, "type": "input", "group": "password", "messages": [], "attributes": {"name": "method", "type": "submit", "value": "password", "disabled": false, "node_type": "input"}}], "action": "http://localhost:4433/self-service/login?flow=98496454-fec9-4922-8ab1-64e64c3ec1e4", "method": "POST"}	ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0	aal1	{}	\N	\N	choose_method	0	\N
\.


--
-- Data for Name: selfservice_recovery_flows; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.selfservice_recovery_flows (id, request_url, issued_at, expires_at, active_method, csrf_token, state, recovered_identity_id, created_at, updated_at, type, ui, nid, submit_count, skip_csrf_check) FROM stdin;
\.


--
-- Data for Name: selfservice_registration_flows; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.selfservice_registration_flows (id, request_url, issued_at, expires_at, active_method, csrf_token, created_at, updated_at, type, ui, nid, internal_context, oauth2_login_challenge, oauth2_login_challenge_data, state, submit_count, organization_id) FROM stdin;
39da5d49-7113-43c9-8458-c3e4c4c316ce	http://localhost:4433/self-service/registration/browser	2025-08-24 16:44:14.536378	2025-08-24 16:54:14.536378		xoLuU19uaYIApCpnxSRrOgG4Eb8XpGwCy1klrjecAeApkFfl/Z6uNCsqu0/0OL9a4FWtJ5BGRvdmLrJAXIzH1g==	2025-08-24 19:44:14.538517	2025-08-24 19:44:14.538517	browser	{"nodes": [{"meta": {"label": {"id": 1040002, "text": "Sign up with google", "type": "info", "context": {"provider": "google", "provider_id": "google"}}}, "type": "input", "group": "oidc", "messages": [], "attributes": {"name": "provider", "type": "submit", "value": "google", "disabled": false, "node_type": "input"}}, {"meta": {}, "type": "input", "group": "default", "messages": [], "attributes": {"name": "csrf_token", "type": "hidden", "value": "xoLuU19uaYIApCpnxSRrOgG4Eb8XpGwCy1klrjecAeApkFfl/Z6uNCsqu0/0OL9a4FWtJ5BGRvdmLrJAXIzH1g==", "disabled": false, "required": true, "node_type": "input"}}, {"meta": {"label": {"id": 1070002, "text": "E-Mail", "type": "info", "context": {"title": "E-Mail"}}}, "type": "input", "group": "password", "messages": [], "attributes": {"name": "traits.email", "type": "email", "disabled": false, "required": true, "node_type": "input", "autocomplete": "email"}}, {"meta": {"label": {"id": 1070001, "text": "Password", "type": "info"}}, "type": "input", "group": "password", "messages": [], "attributes": {"name": "password", "type": "password", "disabled": false, "required": true, "node_type": "input", "autocomplete": "new-password"}}, {"meta": {"label": {"id": 1070002, "text": "First Name", "type": "info", "context": {"title": "First Name"}}}, "type": "input", "group": "password", "messages": [], "attributes": {"name": "traits.name.first", "type": "text", "disabled": false, "required": true, "node_type": "input"}}, {"meta": {"label": {"id": 1070002, "text": "Last Name", "type": "info", "context": {"title": "Last Name"}}}, "type": "input", "group": "password", "messages": [], "attributes": {"name": "traits.name.last", "type": "text", "disabled": false, "required": true, "node_type": "input"}}, {"meta": {"label": {"id": 1040001, "text": "Sign up", "type": "info"}}, "type": "input", "group": "password", "messages": [], "attributes": {"name": "method", "type": "submit", "value": "password", "disabled": false, "node_type": "input"}}], "action": "http://localhost:4433/self-service/registration?flow=39da5d49-7113-43c9-8458-c3e4c4c316ce", "method": "POST"}	ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0	{}	\N	\N	choose_method	0	\N
\.


--
-- Data for Name: selfservice_settings_flows; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.selfservice_settings_flows (id, request_url, issued_at, expires_at, identity_id, created_at, updated_at, active_method, state, type, ui, nid, internal_context) FROM stdin;
\.


--
-- Data for Name: selfservice_verification_flows; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.selfservice_verification_flows (id, request_url, issued_at, expires_at, csrf_token, created_at, updated_at, type, state, active_method, ui, nid, submit_count, oauth2_login_challenge, session_id, identity_id, authentication_methods) FROM stdin;
4d6d06b4-65db-4d89-93a8-efed5eab0dc4	http://localhost:4433/self-service/registration/browser	2025-08-24 16:44:28.973383	2025-08-24 17:44:28.973383	KlDsA5l7cvyZGNy7tyeWkDt06rVeswq5I9D1kdNUMGMVVNWNNio2Ug2+1356DrgGCWxwt0wYN755gMVs2pfkvw==	2025-08-24 19:44:28.973721	2025-08-24 19:44:28.973721	browser	passed_challenge	code	{"nodes": [{"meta": {}, "type": "input", "group": "default", "messages": [], "attributes": {"name": "csrf_token", "type": "hidden", "value": "KlDsA5l7cvyZGNy7tyeWkDt06rVeswq5I9D1kdNUMGMVVNWNNio2Ug2+1356DrgGCWxwt0wYN755gMVs2pfkvw==", "disabled": false, "required": true, "node_type": "input"}}, {"meta": {"label": {"id": 1070009, "text": "Continue", "type": "info"}}, "type": "a", "group": "code", "messages": [], "attributes": {"id": "continue", "href": "http://localhost:4200", "title": {"id": 1070009, "text": "Continue", "type": "info"}, "node_type": "a"}}], "action": "http://localhost:4200", "method": "GET", "messages": [{"id": 1080002, "text": "You successfully verified your email address.", "type": "success"}]}	ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0	1	\N	047967b5-7320-4fd4-bdb4-72e2c10568c7	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	[{"method":"password","aal":"aal1","completed_at":"2025-08-24T16:44:28.966166445Z"}]
\.


--
-- Data for Name: session_devices; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.session_devices (id, ip_address, user_agent, location, nid, session_id, created_at, updated_at) FROM stdin;
02a16111-08be-4322-b229-142dca9f5b9c	[::1]:48390	Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36		ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0	047967b5-7320-4fd4-bdb4-72e2c10568c7	2025-08-24 19:44:28.97027	2025-08-24 19:44:28.97027
d144e99c-6590-448c-8560-a20fa33ba13c	[::1]:48390	Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36		ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0	b80e7fb4-3a46-4400-92b7-5d3974ff1c56	2025-08-24 19:44:53.283864	2025-08-24 19:44:53.283864
\.


--
-- Data for Name: session_token_exchanges; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.session_token_exchanges (id, nid, flow_id, session_id, init_code, return_to_code, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: sessions; Type: TABLE DATA; Schema: kratos; Owner: postgres
--

COPY kratos.sessions (id, issued_at, expires_at, authenticated_at, identity_id, created_at, updated_at, token, active, nid, logout_token, aal, authentication_methods) FROM stdin;
047967b5-7320-4fd4-bdb4-72e2c10568c7	2025-08-24 16:44:28.966167	2025-08-25 16:44:28.966167	2025-08-24 16:44:28.966286	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	2025-08-24 19:44:28.967095	2025-08-24 19:44:28.967095	ory_st_VRb6yqWF8hZ3IP2GI1XdNyCPKtHTbBbN	t	ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0	ory_lo_1CfWDJnn80Bkjlh2BsTObGNzQ0bPZVoX	aal1	[{"aal": "aal1", "method": "password", "completed_at": "2025-08-24T16:44:28.966166445Z"}]
b80e7fb4-3a46-4400-92b7-5d3974ff1c56	2025-08-24 16:44:53.276277	2025-08-25 16:44:53.276277	2025-08-24 16:44:53.276277	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	2025-08-24 19:44:53.281505	2025-08-24 19:44:53.281505	ory_st_6B9Z7SrxS2eOg2l37mmkJADvG3ZfVg2r	t	ab36bc93-bb3a-40a9-8a8c-cce2c447ebf0	ory_lo_QaHgiUlSk28uaP4uQM8JG1PBAUXNfaLW	aal1	[{"aal": "aal1", "method": "password", "completed_at": "2025-08-24T16:44:53.276270399Z"}]
\.


--
-- Data for Name: epics; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.epics (id, user_id, project_id, title, description, created_at, updated_at, deleted_at) FROM stdin;
29e34cea-a50c-419b-a303-baab6254b3ee	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	User Management	Epic covering authentication, user settings, and organization membership	2025-08-24 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	\N
ca71d7bd-3c57-45d0-b5c7-091a2d94d37a	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	Core Project Management	Epic for projects, sprints, and issues	2025-08-24 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	\N
edf3a1af-16c1-4c94-9762-d2fb2ca390e3	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	UI/UX Improvements	Epic for frontend polish and usability	2025-08-24 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	\N
\.


--
-- Data for Name: organizations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.organizations (id, user_id, title, description, created_at, updated_at, deleted_at) FROM stdin;
867d1b19-5332-4a74-88aa-f680a1d9a86f	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	Self Org	Organization for personal projects including Taskiro.	2025-08-24 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	\N
\.


--
-- Data for Name: projects; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.projects (id, user_id, organization_id, title, project_key, created_at, updated_at, deleted_at) FROM stdin;
6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	867d1b19-5332-4a74-88aa-f680a1d9a86f	Taskiro	taskiro	2025-08-24 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	\N
\.


--
-- Data for Name: sprints; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sprints (id, user_id, organization_id, title, description, start_date, end_date, created_at, updated_at, deleted_at) FROM stdin;
8ff8e606-30e9-4c2c-9d4e-984f2ac56c62	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	867d1b19-5332-4a74-88aa-f680a1d9a86f	Sprint 1: Foundation	Set up the base of Taskiro: authentication, projects, issues	2025-08-24 16:45:20.890962+00	2025-09-07 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	\N
b6533f90-bc0e-496a-b34f-324178c61a03	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	867d1b19-5332-4a74-88aa-f680a1d9a86f	Sprint 2: Task Management	Implement issues, epics, and sprint features	2025-09-08 16:45:20.890962+00	2025-09-23 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	\N
\.


--
-- Data for Name: task_statuses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.task_statuses (id, project_id, title, description, complete, index, config) FROM stdin;
f2ce1996-74ee-4cf5-9635-977b62eca7ce	6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	Todo		f	0	
b24af14d-ea9c-48ef-a7ea-09b1e0a6593e	6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	In Progress		f	1	
3a1f3e61-b03c-4a26-9435-f8463b2b2632	6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	Testing		f	2	
393371db-5a39-44f0-957f-69d33e19f78a	6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	Done		t	3	
\.


--
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tasks (id, user_id, epic_id, organization_id, sprint_id, project_id, task_status_id, assignee, title, completed, description, list_index, code, created_at, updated_at, deleted_at) FROM stdin;
b1273597-4a38-4039-9572-955e1b5162c1	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	29e34cea-a50c-419b-a303-baab6254b3ee	867d1b19-5332-4a74-88aa-f680a1d9a86f	8ff8e606-30e9-4c2c-9d4e-984f2ac56c62	6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	f2ce1996-74ee-4cf5-9635-977b62eca7ce	alex	Use only Material SDK	f	Replace default material components with SDK and manual created ui elements	1	task-641697	2025-08-24 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	\N
d227b959-6d86-49aa-a4f0-48fd3766197d	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	29e34cea-a50c-419b-a303-baab6254b3ee	867d1b19-5332-4a74-88aa-f680a1d9a86f	8ff8e606-30e9-4c2c-9d4e-984f2ac56c62	6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	393371db-5a39-44f0-957f-69d33e19f78a	alex	Implement authentication (login/register)	f	Use Ory Kratos for identity management	2	task-487850	2025-08-24 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	\N
c9200d54-d1c5-4409-8eea-9108a978beb3	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	29e34cea-a50c-419b-a303-baab6254b3ee	867d1b19-5332-4a74-88aa-f680a1d9a86f	8ff8e606-30e9-4c2c-9d4e-984f2ac56c62	6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	f2ce1996-74ee-4cf5-9635-977b62eca7ce	alex	Add user profile & settings page	f	Allow users to update their information	3	task-505501	2025-08-24 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	\N
a0aa3f40-e04b-40a2-b889-5e5380bd373c	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	ca71d7bd-3c57-45d0-b5c7-091a2d94d37a	867d1b19-5332-4a74-88aa-f680a1d9a86f	b6533f90-bc0e-496a-b34f-324178c61a03	6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	3a1f3e61-b03c-4a26-9435-f8463b2b2632	alex	Create projects & organizations	f	Implement CRUD for projects/organizations	4	task-440478	2025-08-24 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	\N
6cea1c23-2986-48cf-8574-2ad2704ff452	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	ca71d7bd-3c57-45d0-b5c7-091a2d94d37a	867d1b19-5332-4a74-88aa-f680a1d9a86f	b6533f90-bc0e-496a-b34f-324178c61a03	6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	b24af14d-ea9c-48ef-a7ea-09b1e0a6593e	alex	Add issues & epics	f	Core task tracking functionality	5	task-802210	2025-08-24 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	\N
b84183b3-5669-48da-8f58-0b6cc16e2225	54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	edf3a1af-16c1-4c94-9762-d2fb2ca390e3	867d1b19-5332-4a74-88aa-f680a1d9a86f	b6533f90-bc0e-496a-b34f-324178c61a03	6dfa70dc-2d4a-431f-9df3-f0d9d57a67f2	f2ce1996-74ee-4cf5-9635-977b62eca7ce	alex	Polish dashboard UI	f	Make Taskiro visually appealing with Tailwind & animations	6	task-277977	2025-08-24 16:45:20.890962+00	2025-08-24 16:45:20.890962+00	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, organization_id, settings, code, created_at, updated_at, deleted_at) FROM stdin;
54f6cb2b-5e20-4928-aa73-711a5aa8f5a1	867d1b19-5332-4a74-88aa-f680a1d9a86f		antlernight41401	2025-08-24 16:44:29.074475+00	2025-08-24 16:45:20.895702+00	\N
\.


--
-- Name: continuity_containers continuity_containers_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.continuity_containers
    ADD CONSTRAINT continuity_containers_pkey PRIMARY KEY (id);


--
-- Name: courier_message_dispatches courier_message_dispatches_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.courier_message_dispatches
    ADD CONSTRAINT courier_message_dispatches_pkey PRIMARY KEY (id);


--
-- Name: courier_messages courier_messages_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.courier_messages
    ADD CONSTRAINT courier_messages_pkey PRIMARY KEY (id);


--
-- Name: identities identities_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identities
    ADD CONSTRAINT identities_pkey PRIMARY KEY (id);


--
-- Name: identity_credential_identifiers identity_credential_identifiers_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_credential_identifiers
    ADD CONSTRAINT identity_credential_identifiers_pkey PRIMARY KEY (id);


--
-- Name: identity_credential_types identity_credential_types_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_credential_types
    ADD CONSTRAINT identity_credential_types_pkey PRIMARY KEY (id);


--
-- Name: identity_credentials identity_credentials_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_credentials
    ADD CONSTRAINT identity_credentials_pkey PRIMARY KEY (id);


--
-- Name: identity_login_codes identity_login_codes_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_login_codes
    ADD CONSTRAINT identity_login_codes_pkey PRIMARY KEY (id);


--
-- Name: identity_recovery_addresses identity_recovery_addresses_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_recovery_addresses
    ADD CONSTRAINT identity_recovery_addresses_pkey PRIMARY KEY (id);


--
-- Name: identity_recovery_codes identity_recovery_codes_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_recovery_codes
    ADD CONSTRAINT identity_recovery_codes_pkey PRIMARY KEY (id);


--
-- Name: identity_recovery_tokens identity_recovery_tokens_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_recovery_tokens
    ADD CONSTRAINT identity_recovery_tokens_pkey PRIMARY KEY (id);


--
-- Name: identity_registration_codes identity_registration_codes_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_registration_codes
    ADD CONSTRAINT identity_registration_codes_pkey PRIMARY KEY (id);


--
-- Name: identity_verifiable_addresses identity_verifiable_addresses_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_verifiable_addresses
    ADD CONSTRAINT identity_verifiable_addresses_pkey PRIMARY KEY (id);


--
-- Name: identity_verification_codes identity_verification_codes_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_verification_codes
    ADD CONSTRAINT identity_verification_codes_pkey PRIMARY KEY (id);


--
-- Name: identity_verification_tokens identity_verification_tokens_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_verification_tokens
    ADD CONSTRAINT identity_verification_tokens_pkey PRIMARY KEY (id);


--
-- Name: networks networks_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.networks
    ADD CONSTRAINT networks_pkey PRIMARY KEY (id);


--
-- Name: selfservice_errors selfservice_errors_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_errors
    ADD CONSTRAINT selfservice_errors_pkey PRIMARY KEY (id);


--
-- Name: selfservice_login_flows selfservice_login_requests_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_login_flows
    ADD CONSTRAINT selfservice_login_requests_pkey PRIMARY KEY (id);


--
-- Name: selfservice_settings_flows selfservice_profile_management_requests_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_settings_flows
    ADD CONSTRAINT selfservice_profile_management_requests_pkey PRIMARY KEY (id);


--
-- Name: selfservice_recovery_flows selfservice_recovery_requests_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_recovery_flows
    ADD CONSTRAINT selfservice_recovery_requests_pkey PRIMARY KEY (id);


--
-- Name: selfservice_registration_flows selfservice_registration_requests_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_registration_flows
    ADD CONSTRAINT selfservice_registration_requests_pkey PRIMARY KEY (id);


--
-- Name: selfservice_verification_flows selfservice_verification_requests_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_verification_flows
    ADD CONSTRAINT selfservice_verification_requests_pkey PRIMARY KEY (id);


--
-- Name: session_devices session_devices_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.session_devices
    ADD CONSTRAINT session_devices_pkey PRIMARY KEY (id);


--
-- Name: session_token_exchanges session_token_exchanges_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.session_token_exchanges
    ADD CONSTRAINT session_token_exchanges_pkey PRIMARY KEY (id);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- Name: session_devices unique_session_device; Type: CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.session_devices
    ADD CONSTRAINT unique_session_device UNIQUE (nid, session_id, ip_address, user_agent);


--
-- Name: epics epics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.epics
    ADD CONSTRAINT epics_pkey PRIMARY KEY (id);


--
-- Name: organizations organizations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organizations
    ADD CONSTRAINT organizations_pkey PRIMARY KEY (id);


--
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- Name: sprints sprints_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sprints
    ADD CONSTRAINT sprints_pkey PRIMARY KEY (id);


--
-- Name: task_statuses task_statuses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_statuses
    ADD CONSTRAINT task_statuses_pkey PRIMARY KEY (id);


--
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: continuity_containers_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX continuity_containers_id_nid_idx ON kratos.continuity_containers USING btree (id, nid);


--
-- Name: continuity_containers_identity_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX continuity_containers_identity_id_nid_idx ON kratos.continuity_containers USING btree (identity_id, nid);


--
-- Name: continuity_containers_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX continuity_containers_nid_id_idx ON kratos.continuity_containers USING btree (nid, id);


--
-- Name: courier_message_dispatches_message_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX courier_message_dispatches_message_id_idx ON kratos.courier_message_dispatches USING btree (message_id, created_at DESC);


--
-- Name: courier_message_dispatches_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX courier_message_dispatches_nid_idx ON kratos.courier_message_dispatches USING btree (nid);


--
-- Name: courier_messages_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX courier_messages_id_nid_idx ON kratos.courier_messages USING btree (id, nid);


--
-- Name: courier_messages_nid_created_at_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX courier_messages_nid_created_at_id_idx ON kratos.courier_messages USING btree (nid, created_at DESC);


--
-- Name: courier_messages_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX courier_messages_nid_id_idx ON kratos.courier_messages USING btree (nid, id);


--
-- Name: courier_messages_nid_recipient_created_at_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX courier_messages_nid_recipient_created_at_id_idx ON kratos.courier_messages USING btree (nid, recipient, created_at DESC);


--
-- Name: courier_messages_nid_status_created_at_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX courier_messages_nid_status_created_at_id_idx ON kratos.courier_messages USING btree (nid, status, created_at DESC);


--
-- Name: courier_messages_status_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX courier_messages_status_idx ON kratos.courier_messages USING btree (status);


--
-- Name: identities_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identities_id_nid_idx ON kratos.identities USING btree (id, nid);


--
-- Name: identities_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identities_nid_id_idx ON kratos.identities USING btree (nid, id);


--
-- Name: identity_credential_identifiers_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_credential_identifiers_id_nid_idx ON kratos.identity_credential_identifiers USING btree (id, nid);


--
-- Name: identity_credential_identifiers_identifier_nid_type_uq_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE UNIQUE INDEX identity_credential_identifiers_identifier_nid_type_uq_idx ON kratos.identity_credential_identifiers USING btree (nid, identity_credential_type_id, identifier);


--
-- Name: identity_credential_identifiers_nid_i_ici_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_credential_identifiers_nid_i_ici_idx ON kratos.identity_credential_identifiers USING btree (nid, identifier, identity_credential_id);


--
-- Name: identity_credential_identifiers_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_credential_identifiers_nid_id_idx ON kratos.identity_credential_identifiers USING btree (nid, id);


--
-- Name: identity_credential_identifiers_nid_identity_credential_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_credential_identifiers_nid_identity_credential_id_idx ON kratos.identity_credential_identifiers USING btree (identity_credential_id, nid);


--
-- Name: identity_credential_types_name_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE UNIQUE INDEX identity_credential_types_name_idx ON kratos.identity_credential_types USING btree (name);


--
-- Name: identity_credentials_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_credentials_id_nid_idx ON kratos.identity_credentials USING btree (id, nid);


--
-- Name: identity_credentials_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_credentials_nid_id_idx ON kratos.identity_credentials USING btree (nid, id);


--
-- Name: identity_credentials_nid_identity_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_credentials_nid_identity_id_idx ON kratos.identity_credentials USING btree (identity_id, nid);


--
-- Name: identity_login_codes_flow_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_login_codes_flow_id_idx ON kratos.identity_login_codes USING btree (selfservice_login_flow_id);


--
-- Name: identity_login_codes_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_login_codes_id_nid_idx ON kratos.identity_login_codes USING btree (id, nid);


--
-- Name: identity_login_codes_identity_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_login_codes_identity_id_idx ON kratos.identity_login_codes USING btree (identity_id);


--
-- Name: identity_login_codes_nid_flow_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_login_codes_nid_flow_id_idx ON kratos.identity_login_codes USING btree (nid, selfservice_login_flow_id);


--
-- Name: identity_recovery_addresses_code_uq_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE UNIQUE INDEX identity_recovery_addresses_code_uq_idx ON kratos.identity_recovery_tokens USING btree (token);


--
-- Name: identity_recovery_addresses_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_addresses_id_nid_idx ON kratos.identity_recovery_addresses USING btree (id, nid);


--
-- Name: identity_recovery_addresses_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_addresses_nid_id_idx ON kratos.identity_recovery_addresses USING btree (nid, id);


--
-- Name: identity_recovery_addresses_nid_identity_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_addresses_nid_identity_id_idx ON kratos.identity_recovery_addresses USING btree (identity_id, nid);


--
-- Name: identity_recovery_addresses_status_via_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_addresses_status_via_idx ON kratos.identity_recovery_addresses USING btree (nid, via, value);


--
-- Name: identity_recovery_addresses_status_via_uq_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE UNIQUE INDEX identity_recovery_addresses_status_via_uq_idx ON kratos.identity_recovery_addresses USING btree (nid, via, value);


--
-- Name: identity_recovery_codes_flow_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_codes_flow_id_idx ON kratos.identity_recovery_codes USING btree (selfservice_recovery_flow_id);


--
-- Name: identity_recovery_codes_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_codes_id_nid_idx ON kratos.identity_recovery_codes USING btree (id, nid);


--
-- Name: identity_recovery_codes_identity_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_codes_identity_id_nid_idx ON kratos.identity_recovery_codes USING btree (identity_id, nid);


--
-- Name: identity_recovery_codes_identity_recovery_address_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_codes_identity_recovery_address_id_nid_idx ON kratos.identity_recovery_codes USING btree (identity_recovery_address_id, nid);


--
-- Name: identity_recovery_codes_nid_flow_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_codes_nid_flow_id_idx ON kratos.identity_recovery_codes USING btree (nid, selfservice_recovery_flow_id);


--
-- Name: identity_recovery_tokens_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_tokens_id_nid_idx ON kratos.identity_recovery_tokens USING btree (id, nid);


--
-- Name: identity_recovery_tokens_identity_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_tokens_identity_id_nid_idx ON kratos.identity_recovery_tokens USING btree (identity_id, nid);


--
-- Name: identity_recovery_tokens_identity_recovery_address_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_tokens_identity_recovery_address_id_idx ON kratos.identity_recovery_tokens USING btree (identity_recovery_address_id);


--
-- Name: identity_recovery_tokens_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_tokens_nid_id_idx ON kratos.identity_recovery_tokens USING btree (nid, id);


--
-- Name: identity_recovery_tokens_selfservice_recovery_flow_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_tokens_selfservice_recovery_flow_id_idx ON kratos.identity_recovery_tokens USING btree (selfservice_recovery_flow_id);


--
-- Name: identity_recovery_tokens_token_nid_used_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_recovery_tokens_token_nid_used_idx ON kratos.identity_recovery_tokens USING btree (nid, token, used);


--
-- Name: identity_registration_codes_flow_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_registration_codes_flow_id_idx ON kratos.identity_registration_codes USING btree (selfservice_registration_flow_id);


--
-- Name: identity_registration_codes_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_registration_codes_id_nid_idx ON kratos.identity_registration_codes USING btree (id, nid);


--
-- Name: identity_registration_codes_nid_flow_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_registration_codes_nid_flow_id_idx ON kratos.identity_registration_codes USING btree (nid, selfservice_registration_flow_id);


--
-- Name: identity_verifiable_addresses_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_verifiable_addresses_id_nid_idx ON kratos.identity_verifiable_addresses USING btree (id, nid);


--
-- Name: identity_verifiable_addresses_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_verifiable_addresses_nid_id_idx ON kratos.identity_verifiable_addresses USING btree (nid, id);


--
-- Name: identity_verifiable_addresses_nid_identity_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_verifiable_addresses_nid_identity_id_idx ON kratos.identity_verifiable_addresses USING btree (identity_id, nid);


--
-- Name: identity_verifiable_addresses_status_via_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_verifiable_addresses_status_via_idx ON kratos.identity_verifiable_addresses USING btree (nid, via, value);


--
-- Name: identity_verifiable_addresses_status_via_uq_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE UNIQUE INDEX identity_verifiable_addresses_status_via_uq_idx ON kratos.identity_verifiable_addresses USING btree (nid, via, value);


--
-- Name: identity_verification_codes_flow_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_verification_codes_flow_id_idx ON kratos.identity_verification_codes USING btree (selfservice_verification_flow_id);


--
-- Name: identity_verification_codes_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_verification_codes_id_nid_idx ON kratos.identity_verification_codes USING btree (id, nid);


--
-- Name: identity_verification_codes_nid_flow_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_verification_codes_nid_flow_id_idx ON kratos.identity_verification_codes USING btree (nid, selfservice_verification_flow_id);


--
-- Name: identity_verification_codes_verifiable_address_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_verification_codes_verifiable_address_nid_idx ON kratos.identity_verification_codes USING btree (identity_verifiable_address_id, nid);


--
-- Name: identity_verification_tokens_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_verification_tokens_id_nid_idx ON kratos.identity_verification_tokens USING btree (id, nid);


--
-- Name: identity_verification_tokens_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_verification_tokens_nid_id_idx ON kratos.identity_verification_tokens USING btree (nid, id);


--
-- Name: identity_verification_tokens_token_nid_used_flow_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_verification_tokens_token_nid_used_flow_id_idx ON kratos.identity_verification_tokens USING btree (nid, token, used, selfservice_verification_flow_id);


--
-- Name: identity_verification_tokens_token_uq_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE UNIQUE INDEX identity_verification_tokens_token_uq_idx ON kratos.identity_verification_tokens USING btree (token);


--
-- Name: identity_verification_tokens_verifiable_address_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_verification_tokens_verifiable_address_id_idx ON kratos.identity_verification_tokens USING btree (identity_verifiable_address_id);


--
-- Name: identity_verification_tokens_verification_flow_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX identity_verification_tokens_verification_flow_id_idx ON kratos.identity_verification_tokens USING btree (selfservice_verification_flow_id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON kratos.schema_migration USING btree (version);


--
-- Name: schema_migration_version_self_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX schema_migration_version_self_idx ON kratos.schema_migration USING btree (version_self);


--
-- Name: selfservice_errors_errors_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX selfservice_errors_errors_nid_id_idx ON kratos.selfservice_errors USING btree (nid, id);


--
-- Name: selfservice_login_flows_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX selfservice_login_flows_id_nid_idx ON kratos.selfservice_login_flows USING btree (id, nid);


--
-- Name: selfservice_login_flows_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX selfservice_login_flows_nid_id_idx ON kratos.selfservice_login_flows USING btree (nid, id);


--
-- Name: selfservice_recovery_flows_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX selfservice_recovery_flows_id_nid_idx ON kratos.selfservice_recovery_flows USING btree (id, nid);


--
-- Name: selfservice_recovery_flows_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX selfservice_recovery_flows_nid_id_idx ON kratos.selfservice_recovery_flows USING btree (nid, id);


--
-- Name: selfservice_recovery_flows_recovered_identity_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX selfservice_recovery_flows_recovered_identity_id_nid_idx ON kratos.selfservice_recovery_flows USING btree (recovered_identity_id, nid);


--
-- Name: selfservice_registration_flows_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX selfservice_registration_flows_id_nid_idx ON kratos.selfservice_registration_flows USING btree (id, nid);


--
-- Name: selfservice_registration_flows_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX selfservice_registration_flows_nid_id_idx ON kratos.selfservice_registration_flows USING btree (nid, id);


--
-- Name: selfservice_settings_flows_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX selfservice_settings_flows_id_nid_idx ON kratos.selfservice_settings_flows USING btree (id, nid);


--
-- Name: selfservice_settings_flows_identity_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX selfservice_settings_flows_identity_id_nid_idx ON kratos.selfservice_settings_flows USING btree (identity_id, nid);


--
-- Name: selfservice_settings_flows_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX selfservice_settings_flows_nid_id_idx ON kratos.selfservice_settings_flows USING btree (nid, id);


--
-- Name: selfservice_verification_flows_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX selfservice_verification_flows_id_nid_idx ON kratos.selfservice_verification_flows USING btree (id, nid);


--
-- Name: selfservice_verification_flows_nid_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX selfservice_verification_flows_nid_id_idx ON kratos.selfservice_verification_flows USING btree (nid, id);


--
-- Name: session_devices_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX session_devices_id_nid_idx ON kratos.session_devices USING btree (id, nid);


--
-- Name: session_devices_session_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX session_devices_session_id_nid_idx ON kratos.session_devices USING btree (session_id, nid);


--
-- Name: session_token_exchanges_nid_code_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX session_token_exchanges_nid_code_idx ON kratos.session_token_exchanges USING btree (init_code, nid);


--
-- Name: session_token_exchanges_nid_flow_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX session_token_exchanges_nid_flow_id_idx ON kratos.session_token_exchanges USING btree (flow_id, nid);


--
-- Name: sessions_id_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX sessions_id_nid_idx ON kratos.sessions USING btree (id, nid);


--
-- Name: sessions_identity_id_nid_sorted_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX sessions_identity_id_nid_sorted_idx ON kratos.sessions USING btree (identity_id, nid, authenticated_at DESC);


--
-- Name: sessions_logout_token_uq_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE UNIQUE INDEX sessions_logout_token_uq_idx ON kratos.sessions USING btree (logout_token);


--
-- Name: sessions_nid_created_at_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX sessions_nid_created_at_id_idx ON kratos.sessions USING btree (nid, created_at DESC, id);


--
-- Name: sessions_nid_id_identity_id_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX sessions_nid_id_identity_id_idx ON kratos.sessions USING btree (nid, identity_id, id);


--
-- Name: sessions_token_nid_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE INDEX sessions_token_nid_idx ON kratos.sessions USING btree (nid, token);


--
-- Name: sessions_token_uq_idx; Type: INDEX; Schema: kratos; Owner: postgres
--

CREATE UNIQUE INDEX sessions_token_uq_idx ON kratos.sessions USING btree (token);


--
-- Name: idx_epics_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_epics_deleted_at ON public.epics USING btree (deleted_at);


--
-- Name: idx_organizations_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_organizations_deleted_at ON public.organizations USING btree (deleted_at);


--
-- Name: idx_projects_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_projects_deleted_at ON public.projects USING btree (deleted_at);


--
-- Name: idx_sprints_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sprints_deleted_at ON public.sprints USING btree (deleted_at);


--
-- Name: idx_task_statuses_title; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_task_statuses_title ON public.task_statuses USING btree (title);


--
-- Name: idx_tasks_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tasks_deleted_at ON public.tasks USING btree (deleted_at);


--
-- Name: idx_users_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_users_code ON public.users USING btree (code);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: project_key_uniq; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX project_key_uniq ON public.projects USING btree (organization_id, project_key);


--
-- Name: continuity_containers continuity_containers_identity_id_fkey; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.continuity_containers
    ADD CONSTRAINT continuity_containers_identity_id_fkey FOREIGN KEY (identity_id) REFERENCES kratos.identities(id) ON DELETE CASCADE;


--
-- Name: continuity_containers continuity_containers_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.continuity_containers
    ADD CONSTRAINT continuity_containers_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: courier_message_dispatches courier_message_dispatches_message_id_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.courier_message_dispatches
    ADD CONSTRAINT courier_message_dispatches_message_id_fk FOREIGN KEY (message_id) REFERENCES kratos.courier_messages(id) ON DELETE CASCADE;


--
-- Name: courier_message_dispatches courier_message_dispatches_nid_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.courier_message_dispatches
    ADD CONSTRAINT courier_message_dispatches_nid_fk FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON DELETE CASCADE;


--
-- Name: courier_messages courier_messages_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.courier_messages
    ADD CONSTRAINT courier_messages_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identities identities_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identities
    ADD CONSTRAINT identities_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_credential_identifiers identity_credential_identifiers_identity_credential_id_fkey; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_credential_identifiers
    ADD CONSTRAINT identity_credential_identifiers_identity_credential_id_fkey FOREIGN KEY (identity_credential_id) REFERENCES kratos.identity_credentials(id) ON DELETE CASCADE;


--
-- Name: identity_credential_identifiers identity_credential_identifiers_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_credential_identifiers
    ADD CONSTRAINT identity_credential_identifiers_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_credential_identifiers identity_credential_identifiers_type_id_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_credential_identifiers
    ADD CONSTRAINT identity_credential_identifiers_type_id_fk_idx FOREIGN KEY (identity_credential_type_id) REFERENCES kratos.identity_credential_types(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_credentials identity_credentials_identity_credential_type_id_fkey; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_credentials
    ADD CONSTRAINT identity_credentials_identity_credential_type_id_fkey FOREIGN KEY (identity_credential_type_id) REFERENCES kratos.identity_credential_types(id) ON DELETE CASCADE;


--
-- Name: identity_credentials identity_credentials_identity_id_fkey; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_credentials
    ADD CONSTRAINT identity_credentials_identity_id_fkey FOREIGN KEY (identity_id) REFERENCES kratos.identities(id) ON DELETE CASCADE;


--
-- Name: identity_credentials identity_credentials_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_credentials
    ADD CONSTRAINT identity_credentials_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_login_codes identity_login_codes_identity_id_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_login_codes
    ADD CONSTRAINT identity_login_codes_identity_id_fk FOREIGN KEY (identity_id) REFERENCES kratos.identities(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_login_codes identity_login_codes_networks_id_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_login_codes
    ADD CONSTRAINT identity_login_codes_networks_id_fk FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_login_codes identity_login_codes_selfservice_login_flows_id_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_login_codes
    ADD CONSTRAINT identity_login_codes_selfservice_login_flows_id_fk FOREIGN KEY (selfservice_login_flow_id) REFERENCES kratos.selfservice_login_flows(id) ON DELETE CASCADE;


--
-- Name: identity_recovery_addresses identity_recovery_addresses_identity_id_fkey; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_recovery_addresses
    ADD CONSTRAINT identity_recovery_addresses_identity_id_fkey FOREIGN KEY (identity_id) REFERENCES kratos.identities(id) ON DELETE CASCADE;


--
-- Name: identity_recovery_addresses identity_recovery_addresses_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_recovery_addresses
    ADD CONSTRAINT identity_recovery_addresses_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_recovery_codes identity_recovery_codes_identity_id_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_recovery_codes
    ADD CONSTRAINT identity_recovery_codes_identity_id_fk FOREIGN KEY (identity_id) REFERENCES kratos.identities(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_recovery_codes identity_recovery_codes_identity_recovery_addresses_id_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_recovery_codes
    ADD CONSTRAINT identity_recovery_codes_identity_recovery_addresses_id_fk FOREIGN KEY (identity_recovery_address_id) REFERENCES kratos.identity_recovery_addresses(id) ON DELETE CASCADE;


--
-- Name: identity_recovery_codes identity_recovery_codes_networks_id_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_recovery_codes
    ADD CONSTRAINT identity_recovery_codes_networks_id_fk FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_recovery_codes identity_recovery_codes_selfservice_recovery_flows_id_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_recovery_codes
    ADD CONSTRAINT identity_recovery_codes_selfservice_recovery_flows_id_fk FOREIGN KEY (selfservice_recovery_flow_id) REFERENCES kratos.selfservice_recovery_flows(id) ON DELETE CASCADE;


--
-- Name: identity_recovery_tokens identity_recovery_tokens_identity_id_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_recovery_tokens
    ADD CONSTRAINT identity_recovery_tokens_identity_id_fk_idx FOREIGN KEY (identity_id) REFERENCES kratos.identities(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_recovery_tokens identity_recovery_tokens_identity_recovery_address_id_fkey; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_recovery_tokens
    ADD CONSTRAINT identity_recovery_tokens_identity_recovery_address_id_fkey FOREIGN KEY (identity_recovery_address_id) REFERENCES kratos.identity_recovery_addresses(id) ON DELETE CASCADE;


--
-- Name: identity_recovery_tokens identity_recovery_tokens_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_recovery_tokens
    ADD CONSTRAINT identity_recovery_tokens_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_recovery_tokens identity_recovery_tokens_selfservice_recovery_request_id_fkey; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_recovery_tokens
    ADD CONSTRAINT identity_recovery_tokens_selfservice_recovery_request_id_fkey FOREIGN KEY (selfservice_recovery_flow_id) REFERENCES kratos.selfservice_recovery_flows(id) ON DELETE CASCADE;


--
-- Name: identity_registration_codes identity_registration_codes_networks_id_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_registration_codes
    ADD CONSTRAINT identity_registration_codes_networks_id_fk FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_registration_codes identity_registration_codes_selfservice_registration_flows_id_f; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_registration_codes
    ADD CONSTRAINT identity_registration_codes_selfservice_registration_flows_id_f FOREIGN KEY (selfservice_registration_flow_id) REFERENCES kratos.selfservice_registration_flows(id) ON DELETE CASCADE;


--
-- Name: identity_verifiable_addresses identity_verifiable_addresses_identity_id_fkey; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_verifiable_addresses
    ADD CONSTRAINT identity_verifiable_addresses_identity_id_fkey FOREIGN KEY (identity_id) REFERENCES kratos.identities(id) ON DELETE CASCADE;


--
-- Name: identity_verifiable_addresses identity_verifiable_addresses_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_verifiable_addresses
    ADD CONSTRAINT identity_verifiable_addresses_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_verification_codes identity_verification_codes_identity_verifiable_addresses_id_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_verification_codes
    ADD CONSTRAINT identity_verification_codes_identity_verifiable_addresses_id_fk FOREIGN KEY (identity_verifiable_address_id) REFERENCES kratos.identity_verifiable_addresses(id) ON DELETE CASCADE;


--
-- Name: identity_verification_codes identity_verification_codes_networks_id_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_verification_codes
    ADD CONSTRAINT identity_verification_codes_networks_id_fk FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_verification_codes identity_verification_codes_selfservice_verification_flows_id_f; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_verification_codes
    ADD CONSTRAINT identity_verification_codes_selfservice_verification_flows_id_f FOREIGN KEY (selfservice_verification_flow_id) REFERENCES kratos.selfservice_verification_flows(id) ON DELETE CASCADE;


--
-- Name: identity_verification_tokens identity_verification_tokens_identity_verifiable_address_i_fkey; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_verification_tokens
    ADD CONSTRAINT identity_verification_tokens_identity_verifiable_address_i_fkey FOREIGN KEY (identity_verifiable_address_id) REFERENCES kratos.identity_verifiable_addresses(id) ON DELETE CASCADE;


--
-- Name: identity_verification_tokens identity_verification_tokens_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_verification_tokens
    ADD CONSTRAINT identity_verification_tokens_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: identity_verification_tokens identity_verification_tokens_selfservice_verification_flow_fkey; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.identity_verification_tokens
    ADD CONSTRAINT identity_verification_tokens_selfservice_verification_flow_fkey FOREIGN KEY (selfservice_verification_flow_id) REFERENCES kratos.selfservice_verification_flows(id) ON DELETE CASCADE;


--
-- Name: selfservice_errors selfservice_errors_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_errors
    ADD CONSTRAINT selfservice_errors_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: selfservice_login_flows selfservice_login_flows_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_login_flows
    ADD CONSTRAINT selfservice_login_flows_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: selfservice_settings_flows selfservice_profile_management_requests_identity_id_fkey; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_settings_flows
    ADD CONSTRAINT selfservice_profile_management_requests_identity_id_fkey FOREIGN KEY (identity_id) REFERENCES kratos.identities(id) ON DELETE CASCADE;


--
-- Name: selfservice_recovery_flows selfservice_recovery_flows_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_recovery_flows
    ADD CONSTRAINT selfservice_recovery_flows_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: selfservice_recovery_flows selfservice_recovery_requests_recovered_identity_id_fkey; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_recovery_flows
    ADD CONSTRAINT selfservice_recovery_requests_recovered_identity_id_fkey FOREIGN KEY (recovered_identity_id) REFERENCES kratos.identities(id) ON DELETE CASCADE;


--
-- Name: selfservice_registration_flows selfservice_registration_flows_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_registration_flows
    ADD CONSTRAINT selfservice_registration_flows_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: selfservice_settings_flows selfservice_settings_flows_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_settings_flows
    ADD CONSTRAINT selfservice_settings_flows_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: selfservice_verification_flows selfservice_verification_flows_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.selfservice_verification_flows
    ADD CONSTRAINT selfservice_verification_flows_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: session_devices session_metadata_nid_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.session_devices
    ADD CONSTRAINT session_metadata_nid_fk FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON DELETE CASCADE;


--
-- Name: session_devices session_metadata_sessions_id_fk; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.session_devices
    ADD CONSTRAINT session_metadata_sessions_id_fk FOREIGN KEY (session_id) REFERENCES kratos.sessions(id) ON DELETE CASCADE;


--
-- Name: sessions sessions_identity_id_fkey; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.sessions
    ADD CONSTRAINT sessions_identity_id_fkey FOREIGN KEY (identity_id) REFERENCES kratos.identities(id) ON DELETE CASCADE;


--
-- Name: sessions sessions_nid_fk_idx; Type: FK CONSTRAINT; Schema: kratos; Owner: postgres
--

ALTER TABLE ONLY kratos.sessions
    ADD CONSTRAINT sessions_nid_fk_idx FOREIGN KEY (nid) REFERENCES kratos.networks(id) ON UPDATE RESTRICT ON DELETE CASCADE;


--
-- Name: epics fk_epics_project; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.epics
    ADD CONSTRAINT fk_epics_project FOREIGN KEY (project_id) REFERENCES public.projects(id);


--
-- Name: epics fk_epics_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.epics
    ADD CONSTRAINT fk_epics_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: users fk_organizations_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_organizations_user FOREIGN KEY (organization_id) REFERENCES public.organizations(id);


--
-- Name: users fk_organizations_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_organizations_users FOREIGN KEY (organization_id) REFERENCES public.organizations(id);


--
-- Name: projects fk_projects_organization; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT fk_projects_organization FOREIGN KEY (organization_id) REFERENCES public.organizations(id);


--
-- Name: task_statuses fk_projects_statuses; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_statuses
    ADD CONSTRAINT fk_projects_statuses FOREIGN KEY (project_id) REFERENCES public.projects(id);


--
-- Name: tasks fk_projects_tasks; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT fk_projects_tasks FOREIGN KEY (project_id) REFERENCES public.projects(id);


--
-- Name: projects fk_projects_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT fk_projects_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: sprints fk_sprints_organization; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sprints
    ADD CONSTRAINT fk_sprints_organization FOREIGN KEY (organization_id) REFERENCES public.organizations(id);


--
-- Name: sprints fk_sprints_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sprints
    ADD CONSTRAINT fk_sprints_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: tasks fk_task_statuses_tasks; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT fk_task_statuses_tasks FOREIGN KEY (task_status_id) REFERENCES public.task_statuses(id);


--
-- Name: tasks fk_tasks_epic; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT fk_tasks_epic FOREIGN KEY (epic_id) REFERENCES public.epics(id);


--
-- Name: tasks fk_tasks_organization; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT fk_tasks_organization FOREIGN KEY (organization_id) REFERENCES public.organizations(id);


--
-- Name: tasks fk_tasks_sprint; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT fk_tasks_sprint FOREIGN KEY (sprint_id) REFERENCES public.sprints(id);


--
-- Name: tasks fk_tasks_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT fk_tasks_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

