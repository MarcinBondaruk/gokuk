package recipe

type Ingredient struct {
	Name     string
	Quantity int
	Unit     string //g kg ml l
}

func NewIngredient(name string, quantity int, unit string) Ingredient {
	return Ingredient{
		Name:     name,
		Quantity: quantity,
		Unit:     unit,
	}
}
