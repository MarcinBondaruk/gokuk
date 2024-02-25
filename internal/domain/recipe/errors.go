package recipe

import (
	"fmt"
)

type InvalidTitleLength struct {
	Msg  string
	Code string
}

func (e *InvalidTitleLength) Error() string {
	return fmt.Sprintf("recipe error - code: %s, msg: %s", e.Code, e.Msg)
}

func NewInvalidTitleLengthErr() *InvalidTitleLength {
	return &InvalidTitleLength{
		Msg:  "title should have less than 128 runes",
		Code: "2e5766fe-9942-4d7a-a630-123a9caf8e1b",
	}
}

type InvalidDescriptionLength struct {
	Msg  string
	Code string
}

func (e *InvalidDescriptionLength) Error() string {
	return fmt.Sprintf("recipe error - code: %s, msg: %s", e.Code, e.Msg)
}

func NewInvalidDescriptionLengthErr() *InvalidDescriptionLength {
	return &InvalidDescriptionLength{
		Msg:  "description should be shorter than X",
		Code: "47e86261-7fa2-4cc3-b4f8-60038d0f6463",
	}
}
