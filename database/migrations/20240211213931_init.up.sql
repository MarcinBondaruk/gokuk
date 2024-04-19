CREATE TABLE users (
    id uuid PRIMARY KEY,
    username VARCHAR(64) NOT NULL,
    passwordHash VARCHAR(64) NOT NULL
);

CREATE TABLE products (
    id uuid PRIMARY KEY,
    name VARCHAR(32) UNIQUE NOT NULL,
    type VARCHAR(32) NOT NULL
);

CREATE TABLE recipes (
    id uuid PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    author_id uuid REFERENCES users(id)
);

CREATE TABLE recipe_ingredient (
    name VARCHAR(128) NOT NULL UNIQUE,
    recipe_id uuid REFERENCES recipes(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    unit VARCHAR(3) NOT NULL
);
