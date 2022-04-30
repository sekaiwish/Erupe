BEGIN;

ALTER TABLE characters
    ADD COLUMN guild_post_checked timestamp without time zone NOT NULL DEFAULT now(),
    ADD COLUMN time_played int NOT NULL DEFAULT 0,
    ADD COLUMN weapon_id int NOT NULL DEFAULT 0;

ALTER TABLE characters
    DROP COLUMN gr_override_mode,
    DROP COLUMN small_gr_level,
    DROP COLUMN gr_override_unk0,
    DROP COLUMN gr_override_unk1;

ALTER TABLE characters
    RENAME COLUMN gr_override_level TO gr;

ALTER TABLE characters
    RENAME COLUMN exp TO hrp;

ALTER TABLE characters
    RENAME COLUMN weapon TO weapon_type;

ALTER TABLE guilds
    ADD COLUMN item_box bytea,
    ADD COLUMN event_rp int NOT NULL DEFAULT 0;

ALTER TABLE guilds
    RENAME COLUMN rp TO rank_rp;

ALTER TABLE guilds
    DROP COLUMN guild_hall;

CREATE TABLE guild_posts
(
    id serial NOT NULL PRIMARY KEY,
    guild_id int NOT NULL,
    author_id int NOT NULL,
    post_type int NOT NULL,
    stamp_id int NOT NULL,
    title text NOT NULL,
    body text NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    liked_by text NOT NULL DEFAULT ''
);

CREATE TABLE guild_alliances
(
    id serial NOT NULL PRIMARY KEY,
    name varchar(24) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    parent_id int NOT NULL,
    sub1_id int,
    sub2_id int
);

END;