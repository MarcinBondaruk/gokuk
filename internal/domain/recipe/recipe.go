package recipe

import (
	"unicode/utf8"

	"github.com/google/uuid"
)

type Recipe struct {
	id          uuid.UUID
	title       string
	description string
	ingredients []Ingredient
	authorId    uuid.UUID
}

func NewRecipe(
	id uuid.UUID,
	title string,
	description string,
	authorId uuid.UUID,
) (*Recipe, error) {
	if utf8.RuneCountInString(title) > 128 {
		return nil, NewInvalidTitleLengthErr()
	}

	if utf8.RuneCountInString(description) > 5000 {
		return nil, NewInvalidDescriptionLengthErr()
	}

	return &Recipe{
		id:          id,
		title:       title,
		description: description,
		authorId:    authorId,
	}, nil
}
