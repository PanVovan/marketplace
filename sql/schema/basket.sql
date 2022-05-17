CREATE TABLE IF NOT EXISTS public.baskets
(
    id uuid NOT NULL,
    user_id uuid,
    CONSTRAINT baskets_pkey PRIMARY KEY (id)
);