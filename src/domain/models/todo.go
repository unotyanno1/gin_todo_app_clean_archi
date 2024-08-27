package models

import (
	"errors"

	"gorm.io/gorm"

)

type Todo struct {
	*gorm.Model
	Content string
}

func (t *Todo) Validate() error {
	if t.Content == "" {
		return errors.New("content empty")
	}
	return nil
}