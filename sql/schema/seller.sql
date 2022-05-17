CREATE TABLE IF NOT EXISTS public.sellers
(
    id uuid NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT seller_pkey PRIMARY KEY (id),
    CONSTRAINT sellers_email_key UNIQUE (email)
);