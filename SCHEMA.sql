BEGIN;

ALTER TABLE public.mail
    ADD COLUMN locked bool NOT NULL DEFAULT false;

CREATE TABLE public.guild_alliances
(
    id serial NOT NULL PRIMARY KEY,
    name varchar(24) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    parent_id int NOT NULL,
    sub1_id int,
    sub2_id int
);

END;