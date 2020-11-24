package test

import (
	"time"
)

type Contato struct {
	ID         int64     `json:"id"`
	Nome       string    `json:"nome"  validate:"required"`
	Celular    string    `json:"celular" validate:"required"`
	Email      string    `json:"email"`
	DateCreate time.Time `json:"dt_create"`
}
