CREATE TABLE IF NOT EXISTS public.order_items
(
    id uuid NOT NULL,
    order_id uuid NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    price numeric NOT NULL,
    quantity integer NOT NULL,
    CONSTRAINT order_items_pkey PRIMARY KEY (id),
    CONSTRAINT order_items_order_id_fkey FOREIGN KEY (order_id)
        REFERENCES public.orders (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);