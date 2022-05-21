CREATE TABLE IF NOT EXISTS public.basket_products
(
    id uuid NOT NULL,
    basket_id uuid NOT NULL,
    product_id uuid NOT NULL,
    quantity integer NOT NULL,
    CONSTRAINT basket_products_pkey PRIMARY KEY (id),
    CONSTRAINT basket_products_basket_id_fkey FOREIGN KEY (basket_id)
        REFERENCES public.baskets (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT basket_products_product_id_fkey FOREIGN KEY (product_id)
        REFERENCES public.products (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
);