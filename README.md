# About Gokuk
Domestic meal planner and shopping helper  
API: https://marcinbondaruk.github.io/gokuk/

## Main Functionalities
1. Add Recipe
2. Browse Recipes
3. Create meal plans
4. Create Shopping list based on meal plans or recipes

## Advanced Functionalities
1. Suggest Recipes based on leftovers from previous cooking


## About Project

### Migrations
` go-migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/gokuk?sslmode=disable" up`

* repositories - retrieve and put away domain objects
* query - bypass domain objects
