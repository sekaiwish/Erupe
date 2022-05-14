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

CREATE TABLE public.distribution
(
    id serial NOT NULL PRIMARY KEY,
    character_id int,
    type int NOT NULL,
    deadline timestamp without time zone,
    event_name text NOT NULL DEFAULT 'GM Gift!',
    description text NOT NULL DEFAULT '~C05You received a gift!',
    times_acceptable int NOT NULL DEFAULT 1,
    min_hr int NOT NULL DEFAULT 65535,
    max_hr int NOT NULL DEFAULT 65535,
    min_sr int NOT NULL DEFAULT 65535,
    max_sr int NOT NULL DEFAULT 65535,
    min_gr int NOT NULL DEFAULT 65535,
    max_gr int NOT NULL DEFAULT 65535,
    data bytea NOT NULL
);

CREATE TABLE public.distributions_accepted
(
    distribution_id int,
    user_id int
);

END;