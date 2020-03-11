package utils

import (
	"github.com/google/uuid"
)

func TokenGenerator() string {
	uuid := uuid.New()
	return uuid.String()
}
