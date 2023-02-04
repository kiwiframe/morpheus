package domain

import (
	"time"

	"github.com/google/uuid"
)

type Verification struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	ExpiresAt time.Time
}

type NewVerificationParams struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	ExpiresAt time.Time
}

func NewVerification(params *NewVerificationParams) (*Verification, error) {
	v := &Verification{
		ID:        params.ID,
		UserID:    params.UserID,
		Token:     params.Token,
		ExpiresAt: params.ExpiresAt,
	}
	return v, nil
}
