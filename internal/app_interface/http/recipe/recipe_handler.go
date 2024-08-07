package recipe

import (
	"encoding/json"
	"net/http"

	"github.com/MarcinBondaruk/gokuk/internal/services/recipe"
)

type RecipeHandler struct {
	recipeService *recipe.RecipeService
}

func NewRecipeHandler(recipeService *recipe.RecipeService) *RecipeHandler {
	return &RecipeHandler{
		recipeService: recipeService,
	}
}

func (rh *RecipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var request CreateRecipeRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = request.Validate(); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = rh.recipeService.CreateRecipe(r.Context(), request.ToCommand())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Recipe created"))
}

func (rh *RecipeHandler) GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := rh.recipeService.GetRecipes(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipes)
}
