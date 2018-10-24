--
-- Name: person; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.person (
    id integer NOT NULL,
    first_name character varying(255),
    last_name character varying(255),
    birthdate date,
    passport_number character varying(255),
    country_origin character varying(255),
    country_birth character varying(255),
    language character varying(255),
    gender character varying(10),
    transafer text,
    mastercouncil character varying(255),
    image text,
    status integer,
    section character varying(10),
    "position" character varying(100),
    notes text,
    updated_by integer
);


--
-- Name: person_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.person_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: person_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.person_id_seq OWNED BY public.person.id;


