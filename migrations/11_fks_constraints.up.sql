--
-- Name: products stock_check; Type: CHECK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE public.products
    ADD CONSTRAINT stock_check CHECK (((stock)::numeric >= (0)::numeric)) NOT VALID;


--
-- Name: order_enc admin_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_enc
    ADD CONSTRAINT admin_fk FOREIGN KEY (created_by) REFERENCES public.admins(id);


--
-- Name: person_item admin_person_item_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.person_item
    ADD CONSTRAINT admin_person_item_fk FOREIGN KEY (created_by) REFERENCES public.admins(id);


--
-- Name: person_item item_person_item_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.person_item
    ADD CONSTRAINT item_person_item_fk FOREIGN KEY (item_id) REFERENCES public.item(id);


--
-- Name: order_det order_enc_order_det_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_det
    ADD CONSTRAINT order_enc_order_det_fk FOREIGN KEY (oreder_enc_id) REFERENCES public.order_enc(id);


--
-- Name: order_enc person_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_enc
    ADD CONSTRAINT person_fk FOREIGN KEY (person_id) REFERENCES public.person(id);


--
-- Name: person_item person_person_item_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.person_item
    ADD CONSTRAINT person_person_item_fk FOREIGN KEY (person_id) REFERENCES public.person(id);


--
-- Name: order_det product_order_det_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_det
    ADD CONSTRAINT product_order_det_fk FOREIGN KEY (product_id) REFERENCES public.products(id);


