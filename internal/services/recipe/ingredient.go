package recipe

type Ingredient struct {
	Name  string
	Value int
	Unit  string //g kg ml l
}

func NewIngredient(name string, value int, unit string) Ingredient {
	return Ingredient{
		Name:  name,
		Value: value,
		Unit:  unit,
	}
}
