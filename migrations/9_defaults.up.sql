--
-- Name: admins id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.admins ALTER COLUMN id SET DEFAULT nextval('public.admins_id_seq'::regclass);


--
-- Name: item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.item ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);


--
-- Name: order_det id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_det ALTER COLUMN id SET DEFAULT nextval('public.order_det_id_seq'::regclass);


--
-- Name: order_enc id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_enc ALTER COLUMN id SET DEFAULT nextval('public.order_enc_id_seq'::regclass);


--
-- Name: person id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.person ALTER COLUMN id SET DEFAULT nextval('public.person_id_seq'::regclass);


--
-- Name: person_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.person_item ALTER COLUMN id SET DEFAULT nextval('public.person_item_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);

--

