package handlers

import (
	"github.com/google/uuid"
)


func getUUID(id string) uuid.UUID {
	return uuid.MustParse(id)
}
