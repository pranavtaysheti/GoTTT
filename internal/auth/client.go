package auth

import (
	"github.com/google/uuid"
)

type Client interface {
	NewClient() uuid.UUID
	Login() error
}