CREATE TABLE IF NOT EXISTS public.product_images
(
    id uuid NOT NULL,
    product_id uuid NOT NULL,
    file character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT product_images_pkey PRIMARY KEY (id),
    CONSTRAINT product_images_product_id_fkey FOREIGN KEY (product_id)
        REFERENCES public.products (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)