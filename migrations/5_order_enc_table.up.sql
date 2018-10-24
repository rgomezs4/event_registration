--
-- Name: order_enc; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_enc (
    id integer NOT NULL,
    person_id integer NOT NULL,
    " payment_method" integer NOT NULL,
    total double precision NOT NULL,
    comment text,
    btc_address text,
    created_by integer NOT NULL
);


--
-- Name: order_enc_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_enc_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

--
-- Name: order_enc_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_enc_id_seq OWNED BY public.order_enc.id;


