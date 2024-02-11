CREATE TABLE products (
    id uuid PRIMARY KEY,
    name VARCHAR(32)
);

CREATE TABLE recipes (
    id uuid PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE recipe_ingredient (
    product_id REFERENCES products(id),
    recipe_id REFERENCES recipes(id) ON DELETE CASCADE,
    amout INT,
    unit VARCHAR(2)
);
