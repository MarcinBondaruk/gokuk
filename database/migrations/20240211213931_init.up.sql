CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    username VARCHAR(64) NOT NULL,
    passwordHash VARCHAR(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    id uuid PRIMARY KEY,
    name VARCHAR(32) UNIQUE NOT NULL,
    type VARCHAR(32) NOT NULL
);

CREATE TABLE IF NOT EXISTS recipes (
    id uuid PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    author_id uuid REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS recipe_ingredient (
    name VARCHAR(128) NOT NULL UNIQUE,
    recipe_id uuid REFERENCES recipes(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    unit VARCHAR(3) NOT NULL
);
