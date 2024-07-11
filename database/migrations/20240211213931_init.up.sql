CREATE TABLE IF NOT EXISTS user (
    id bigserial NOT NULL PRIMARY KEY,
    username VARCHAR(64) NOT NULL,
    passwordHash VARCHAR(64) NOT NULL,

    created_at timestamp with time zone NOT NULL default current_timestamp
);

CREATE TABLE IF NOT EXISTS ingredient (
    id bigserial NOT NULL PRIMARY KEY,
    name VARCHAR(32) UNIQUE NOT NULL,
    type VARCHAR(32) NOT NULL
);

CREATE TABLE IF NOT EXISTS recipe (
    id bigserial NOT NULL PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    author_id uuid REFERENCES users(id)

    created_at timestamp with time zone NOT NULL default current_timestamp
);

CREATE TABLE IF NOT EXISTS recipe_ingredient (
    id bigserial NOT NULL PRIMARY KEY,
    recipe_id bigint REFERENCES recipes(id) ON DELETE CASCADE,
    ingredient_id bigint REFERENCES ingredient(id),
    value INT NOT NULL,
    unit VARCHAR(3) NOT NULL
);
