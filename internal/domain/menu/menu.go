package menu

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	id      uuid.UUID
	recipes []uuid.UUID
	date    time.Time
}

func (m *Menu) Id() string {
	return m.id.String()
}

func (m *Menu) Recipes() []uuid.UUID {
	return m.recipes
}

func (m *Menu) Date() time.Time {
	return m.date
}
