package product

import "github.com/google/uuid"

type Product struct {
	id       uuid.UUID
	name     string
	category string
}

func (p *Product) Id() string {
	return p.id.String()
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Category() string {
	return p.category
}
