--
-- Name: products stock_check; Type: CHECK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE public.products
    DROP CONSTRAINT stock_check;


--
-- Name: order_enc admin_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_enc
    DROP CONSTRAINT admin_fk;


--
-- Name: person_item admin_person_item_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.person_item
    DROP CONSTRAINT admin_person_item_fk;


--
-- Name: person_item item_person_item_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.person_item
    DROP CONSTRAINT item_person_item_fk;


--
-- Name: order_det order_enc_order_det_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_det
    DROP CONSTRAINT order_enc_order_det_fk;


--
-- Name: order_enc person_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_enc
    DROP CONSTRAINT person_fk;


--
-- Name: person_item person_person_item_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.person_item
    DROP CONSTRAINT person_person_item_fk;


--
-- Name: order_det product_order_det_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_det
    DROP CONSTRAINT product_order_det_fk;


