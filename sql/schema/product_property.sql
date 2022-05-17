CREATE TABLE IF NOT EXISTS public.product_properties
(
    id uuid NOT NULL,
    product_id uuid NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    value character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT product_properties_pkey PRIMARY KEY (id),
    CONSTRAINT product_properties_product_id_fkey FOREIGN KEY (product_id)
        REFERENCES public.products (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);