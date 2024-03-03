package response

type ProductsBulkAddResponse struct {
	ProductsAdded int      `json:"productsAdded"`
	ProductIds    []string `json:"productIds"`
}
