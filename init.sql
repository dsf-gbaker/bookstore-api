DROP TABLE IF EXISTS authors CASCADE;
CREATE TABLE public.authors
(
    id SERIAL,
    firstname character varying(255) COLLATE pg_catalog."default" NOT NULL,
    lastname character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT derpy_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.authors
    OWNER to postgres;

DROP TABLE IF EXISTS books;
CREATE TABLE public.books
(
    id SERIAL,
    title character varying(255) COLLATE pg_catalog."default" NOT NULL,
    author_id integer,
    CONSTRAINT "fkAuthor" FOREIGN KEY (author_id)
        REFERENCES public.authors (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.books
    OWNER to postgres;


INSERT INTO authors (firstname,lastname) VALUES ('George','Washington');
INSERT INTO authors (firstname,lastname) VALUES ('Penelope','Penultimate');
INSERT INTO authors (firstname,lastname) VALUES ('Magnus','Carlson');
INSERT INTO books (title, author_id) VALUES ('My Journey With Jonny Appleseed',1);
INSERT INTO books (title, author_id) VALUES ('The Secret to Super Powers',2);
INSERT INTO books (title, author_id) VALUES ('My ''PEN'' Ultimate Book on Screenwriting',2);
INSERT INTO books (title, author_id) VALUES ('Power Chess: A Guide to Rage Quiting',3);
