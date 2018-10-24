--
-- Name: order_det; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_det (
    id integer NOT NULL,
    oreder_enc_id integer NOT NULL,
    product_id integer NOT NULL,
    " quantity" integer NOT NULL,
    price double precision NOT NULL,
    amount double precision NOT NULL
);


--
-- Name: order_det_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_det_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
--
-- Name: order_det_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_det_id_seq OWNED BY public.order_det.id;


