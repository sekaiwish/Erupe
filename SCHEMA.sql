BEGIN;

ALTER TABLE characters
    ADD COLUMN guild_post_checked int NOT NULL DEFAULT CAST(EXTRACT(epoch FROM now()) AS int);

ALTER TABLE guilds
    ADD COLUMN item_box bytea;

ALTER TABLE guilds
    DROP COLUMN guild_hall;

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

/*
(SELECT id FROM guild_alliances ga WHERE
	 	ga.parent_id = g.id OR
	 	ga.sub1_id = g.id OR
	 	ga.sub2_id = g.id
	) AS alliance_id,
*/