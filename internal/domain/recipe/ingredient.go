package recipe

type Ingredient struct {
	name     string
	quantity int
	unit     string //g kg ml l
}

func NewIngredient(name string, quantity int, unit string) *Ingredient {
	return &Ingredient{
		name:     name,
		quantity: quantity,
		unit:     unit,
	}
}
