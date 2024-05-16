package models

import (
	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type Outfit struct {
	ID     uuid.UUID
	Items  string
	Season string
	Type   string
}
