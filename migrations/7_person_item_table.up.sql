--
-- Name: person_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.person_item (
    id integer NOT NULL,
    person_id integer NOT NULL,
    item_id integer NOT NULL,
    created_by integer NOT NULL
);


--
-- Name: person_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.person_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: person_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.person_item_id_seq OWNED BY public.person_item.id;


