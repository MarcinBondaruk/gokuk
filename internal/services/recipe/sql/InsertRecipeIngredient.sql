INSERT INTO  recipe_ingredient(recipe_id, ingredient_id, value, unit)
VALUES (@recipe_id, @ingredient_id, @value, @unit)
RETURNING
    id