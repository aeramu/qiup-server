package service

import (
	"github.com/google/uuid"
)

//GenerateUUID create uuid
func GenerateUUID() string {
	return uuid.New().String()
}
