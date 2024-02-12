CREATE TABLE users (
    id uuid PRIMARY KEY,
    username VARCHAR(64),
    password VARCHAHR(64)
);

CREATE TABLE products (
    id uuid PRIMARY KEY,
    name VARCHAR(32) NOT NULL
);

CREATE TABLE recipes (
    id uuid PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    author_id REFERENCES users(id)
);

CREATE TABLE recipe_ingredient (
    id uuid PRIMARY KEY,
    product_id REFERENCES products(id),
    recipe_id REFERENCES recipes(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    unit VARCHAR(2) NOT NULL
);

CREATE TABLE meal_plan (
    id uuid PRIMARY KEY,
    owner_id REFERENCES users(id)
);

CREATE TABLE meal_plan_recipe (
    id uuid PRIMARY KEY,
    meal_plan_id REFERENCES meal_plan(id) ON DELETE CASCADE,
    recipe_id REFERENCES recipes(id)
);
