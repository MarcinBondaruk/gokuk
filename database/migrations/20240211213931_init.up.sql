CREATE TABLE users (
    id uuid PRIMARY KEY,
    username VARCHAR(64) NOT NULL,
    passwordHash VARCHAR(64) NOT NULL
);

CREATE TABLE products (
    id uuid PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    type VARCHAR(32) NOT NULL
);

CREATE TABLE recipes (
    id uuid PRIMARY KEY,
    title VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    author_id uuid REFERENCES users(id)
);

CREATE TABLE recipe_ingredient (
    id uuid PRIMARY KEY,
    product_id uuid REFERENCES products(id),
    recipe_id uuid REFERENCES recipes(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    unit VARCHAR(2) NOT NULL
);

CREATE TABLE meal_plan (
    id uuid PRIMARY KEY,
    owner_id uuid REFERENCES users(id),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    CHECK (start_date <= end_date)
);

CREATE TABLE meal_plan_recipe (
    id uuid PRIMARY KEY,
    meal_plan_id uuid REFERENCES meal_plan(id) ON DELETE CASCADE,
    recipe_id uuid REFERENCES recipes(id),
    day DATE NOT NULL
);
