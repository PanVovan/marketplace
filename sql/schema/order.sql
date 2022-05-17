CREATE TABLE IF NOT EXISTS public.orders
(
    id uuid NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    address character varying(255) COLLATE pg_catalog."default" NOT NULL,
    amount integer NOT NULL,
    status integer NOT NULL DEFAULT 0,
    phone character varying(255) COLLATE pg_catalog."default" NOT NULL,
    comment character varying(255) COLLATE pg_catalog."default",
    user_id uuid,
    CONSTRAINT orders_pkey PRIMARY KEY (id),
    CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);