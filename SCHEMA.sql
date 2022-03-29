BEGIN;

ALTER TABLE guilds
    ADD COLUMN item_box bytea;

CREATE TABLE guild_posts
(
    id serial NOT NULL PRIMARY KEY,
    guild_id int NOT NULL,
    author_id int NOT NULL,
    post_type int NOT NULL,
    stamp_id int NOT NULL,
    post bytea NOT NULL,
    likes int NOT NULL DEFAULT 0,
    created_at int NOT NULL DEFAULT CAST(EXTRACT(epoch FROM now()) AS int),
    liked_by text
);

END;