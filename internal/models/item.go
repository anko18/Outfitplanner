package models

import (
	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type Item struct {
	ID    uuid.UUID
	Name  string
	Color string
	Print string
	Firm  string
}
