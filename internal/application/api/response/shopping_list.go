package response

type ShoppingListResponse struct {
	Id       string            `json:"id"`
	Products []ProductResponse `json:"products"`
}
