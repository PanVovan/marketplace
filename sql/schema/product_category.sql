CREATE TABLE IF NOT EXISTS public.products_categories
(
    products_id uuid NOT NULL,
    categories_id uuid NOT NULL,
    CONSTRAINT products_categories_pkey PRIMARY KEY (products_id, categories_id),
    CONSTRAINT products_categories_categories_id_fkey FOREIGN KEY (categories_id)
        REFERENCES public.categories (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT products_categories_products_id_fkey FOREIGN KEY (products_id)
        REFERENCES public.products (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);