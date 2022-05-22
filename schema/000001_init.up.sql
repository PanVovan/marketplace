CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public.sellers
(
    id uuid NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT seller_pkey PRIMARY KEY (id),
    CONSTRAINT sellers_email_key UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS public.users
(
    id uuid NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS public.baskets
(
    id uuid NOT NULL,
    user_id uuid,
    CONSTRAINT baskets_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.brands
(
    id uuid NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT brands_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.categories
(
    id uuid NOT NULL,
    name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT categories_pkey PRIMARY KEY (id)
);


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
);

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


CREATE TABLE IF NOT EXISTS public.rating
(
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    product_id uuid NOT NULL,
    rate integer NOT NULL DEFAULT 5,
    CONSTRAINT rating_pkey PRIMARY KEY (id),
    CONSTRAINT rating_product_id_fkey FOREIGN KEY (product_id)
        REFERENCES public.products (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT rating_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE SET NULL
);

