package http_handler

import (
	"encoding/json"
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/services/recipe"
)

type RecipeHandler struct {
	recipeService recipe.RecipeService
}

func NewRecipeHandler(recipeService recipe.RecipeService) RecipeHandler {
	return RecipeHandler{
		recipeService: recipeService,
	}
}

func (rh *RecipeHandler) GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := rh.recipeService.GetRecipes()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(recipes)
}
