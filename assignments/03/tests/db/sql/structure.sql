CREATE SCHEMA IF NOT EXISTS public;

DROP TABLE IF EXISTS public.products;

CREATE TABLE IF NOT EXISTS public.products (
    id bigserial NOT NULL,
    name varchar(250) NOT NULL,
    price double precision NOT NULL,
    amount numeric NOT NULL
);

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);
