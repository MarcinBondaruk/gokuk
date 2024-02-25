package product

import "github.com/google/uuid"

type Product struct {
	id   uuid.UUID
	name string
}

func (p *Product) Id() string {
	return p.id.String()
}

func (p *Product) Name() string {
	return p.name
}
