CREATE TABLE IF NOT EXISTS public.products
(
    id uuid NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    price numeric NOT NULL DEFAULT 0,
    rating numeric DEFAULT NULL::numeric,
    brand_id uuid,
    seller_id uuid NOT NULL,
    amount integer NOT NULL DEFAULT 1,
    CONSTRAINT products_pkey PRIMARY KEY (id),
    CONSTRAINT products_brand_id_fkey FOREIGN KEY (brand_id)
        REFERENCES public.brands (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT products_seller_id_fkey FOREIGN KEY (seller_id)
        REFERENCES public.sellers (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);