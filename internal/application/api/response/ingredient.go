package response

type IngredientResponse struct {
	Amount    int    `json:"amount"`
	ProductId string `json:"productId"`
	Unit      string `json:"unit"`
}
